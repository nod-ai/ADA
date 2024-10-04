/**
 * Copyright 2024 Advanced Micro Devices, Inc.  All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
**/

package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// Define a struct that matches the JSON structure
type Payload struct {
	OdataType string  `json:"@odata.type"`
	Name      string  `json:"Name"`
	Id        string  `json:"Id"`
	Events    []Event `json:"Events"`
	Context   string  `json:"Context"`
}

type Event struct {
	EventType         string            `json:"EventType"`
	EventId           string            `json:"EventId"`
	EventTimestamp    string            `json:"EventTimestamp"`
	Severity          string            `json:"Severity"`
	Message           string            `json:"Message"`
	MessageId         string            `json:"MessageId"`
	MessageArgs       []string          `json:"MessageArgs"`
	OriginOfCondition OriginOfCondition `json:"OriginOfCondition"`
}

type OriginOfCondition struct {
	OdataId string `json:"@odata.id"`
}

type Server struct {
	listenIP     string
	listenPort   string
	listener     net.Listener
	shutdownChan chan struct{}
}

func NewServer(listenIP string, listenPort string) *Server {
	return &Server{
		listenIP:     listenIP,
		listenPort:   listenPort,
		shutdownChan: make(chan struct{}),
	}
}

func (s *Server) Start(AppConfig Config) error {
	var err error
	var listener net.Listener
	if AppConfig.SystemInformation.UseSSL {
		listener, err = s.startTLS(AppConfig.CertificateDetails.CertFile, AppConfig.CertificateDetails.KeyFile)
	} else {
		listener, err = s.startWithoutTLS()
	}

	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
		return err
	}

	s.listener = listener // Store the listener in the Server struct

	go s.acceptLoop(AppConfig)

	log.Printf("Listening on %s:%s via %s",
		s.listenIP,
		s.listenPort,
		func() string {
			if AppConfig.SystemInformation.UseSSL {
				return "HTTPS"
			}
			return "HTTP"
		}())

	<-s.shutdownChan
	log.Println("Shutting down listener...")
	listener.Close()
	return nil

}

func (s *Server) startTLS(certFile, keyFile string) (net.Listener, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Println("Failed to load certificates")
		return nil, err
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	return tls.Listen("tcp", fmt.Sprintf("%s:%s", s.listenIP, s.listenPort), config)
}

func (s *Server) startWithoutTLS() (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf("%s:%s", s.listenIP, s.listenPort))
}

func (s *Server) acceptLoop(AppConfig Config) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// Check if the error is due to the listener being closed
			if errors.Is(err, net.ErrClosed) {
				log.Println("Listener closed, stopping accept loop")
				break
			}
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		log.Printf("Connection from %s connected", conn.RemoteAddr())

		go s.handleConnection(AppConfig, conn)
	}

}

func (s *Server) handleConnection(AppConfig Config, conn net.Conn) {
	defer conn.Close()

	eventCount := &AppConfig.eventCount
	dataBuffer := &AppConfig.dataBuffer

	reader := bufio.NewReader(conn)

	for {
		// Read the HTTP request from the connection
		req, err := http.ReadRequest(reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %v", err)
				sendErrorResponse(conn, nil)
			}
			break
		}

		err = processRequest(AppConfig, conn, req, eventCount, dataBuffer)
		if err != nil {
			log.Printf("Error processing request: %v", err)
			sendErrorResponse(conn, req)
			break
		}
	}
}

func processRequest(AppConfig Config, conn net.Conn, req *http.Request, eventCount *int, dataBuffer *[]byte) error {
	// Extract method, headers, and payload
	method := req.Method
	headers := req.Header

	// Read the payload
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("error reading payload: %w", err)
	}
	req.Body.Close()

	// Unmarshal the JSON payload into the struct
	var p Payload
	err = json.Unmarshal(payload, &p)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	// Log the extracted information
	var eventType string
	log.Printf("Method: %s", method)
	log.Printf("Headers: %v", headers)
	for _, event := range p.Events {
		eventType = event.EventType
		eventId := event.EventId
		severity := event.Severity
		message := event.Message
		messageId := event.MessageId
		messageArgs := event.MessageArgs
		originOfCondition := event.OriginOfCondition.OdataId

		log.Printf("Event Type: %s", eventType)
		log.Printf("Event ID: %s", eventId)
		log.Printf("Severity: %s", severity)
		log.Printf("Message: %s", message)
		log.Printf("Message ID: %s", messageId)
		log.Printf("Message Args: %v", messageArgs)
		log.Printf("Origin Of Condition: %s", originOfCondition)
		for _, triggerEvent := range AppConfig.TriggerEvents {
			if eventType == triggerEvent.EventType && eventId == triggerEvent.EventId && severity == triggerEvent.Severity {
				log.Printf("Matched Trigger Event: %s with action %s", triggerEvent.EventId, triggerEvent.Action)
				// TODO: Add the SLURM integration here
				// Sending event belings to redfish_utils. Each server may have different slurm node associated, and redfish_servers has the info/map.
				break
			}
		}
	}

	// Append data to dataBuffer and increment eventCount
	*dataBuffer = append(*dataBuffer, payload...)
	*eventCount++

	// Extract the IP address
	ip, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		ip = conn.RemoteAddr().String() // Fallback to full address if splitting fails
	}

	// Update metrics using variables from metrics.go
	timestamp := float64(time.Now().Unix())
	eventCountMetric.WithLabelValues(ip, eventType).Inc()
	eventProcessingTimeMetric.WithLabelValues(ip, eventType).Set(timestamp)

	// Send a 200 OK response
	response := &http.Response{
		Status:        "200 OK",
		StatusCode:    http.StatusOK,
		Proto:         req.Proto,
		ProtoMajor:    req.ProtoMajor,
		ProtoMinor:    req.ProtoMinor,
		Header:        make(http.Header),
		Body:          io.NopCloser(bytes.NewBufferString("OK")),
		ContentLength: int64(len("OK")),
	}
	response.Header.Set("Content-Type", "text/plain")
	err = response.Write(conn)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
	return nil
}

func sendErrorResponse(conn net.Conn, req *http.Request) {
	response := &http.Response{
		Status:        "500 Internal Server Error",
		StatusCode:    http.StatusInternalServerError,
		Proto:         "HTTP/1.1",
		Header:        make(http.Header),
		Body:          io.NopCloser(bytes.NewBufferString("Internal Server Error")),
		ContentLength: int64(len("Internal Server Error")),
	}
	response.Header.Set("Content-Type", "text/plain")
	if req != nil {
		response.Proto = req.Proto
		response.ProtoMajor = req.ProtoMajor
		response.ProtoMinor = req.ProtoMinor
	}
	err := response.Write(conn)
	if err != nil {
		log.Printf("Error writing error response: %v", err)
	}
}
