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

func startListener(AppConfig Config, shutdownChan <-chan struct{}) {
	var listener net.Listener
	var err error
	//TODO test with certificates
	if AppConfig.SystemInformation.UseSSL {
		cert, err := tls.LoadX509KeyPair(AppConfig.CertificateDetails.CertFile, AppConfig.CertificateDetails.KeyFile)
		if err != nil {
			log.Fatalf("Failed to load certificates: %v", err)
		}
		config := &tls.Config{Certificates: []tls.Certificate{cert}}
		listener, err = tls.Listen("tcp", fmt.Sprintf("%s:%d", AppConfig.SystemInformation.ListenerIP, AppConfig.SystemInformation.ListenerPort), config)
	} else {
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", AppConfig.SystemInformation.ListenerIP, AppConfig.SystemInformation.ListenerPort))
	}

	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
	}

	log.Printf("Listening on %s:%d via %s", AppConfig.SystemInformation.ListenerIP, AppConfig.SystemInformation.ListenerPort, func() string {
		if AppConfig.SystemInformation.UseSSL {
			return "HTTPS"
		}
		return "HTTP"
	}())

	go func() {
		<-shutdownChan
		log.Println("Shutting down listener...")
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-shutdownChan:
				return
			default:
				log.Printf("Failed to accept connection: %v", err)
				continue
			}
		}
		go processData(AppConfig, conn)
	}
}

func processData(AppConfig Config, conn net.Conn) {
	var connStreamOut net.Conn

	if AppConfig.SystemInformation.UseSSL {
		connStreamOut = tls.Server(conn, AppConfig.context)
	} else {
		connStreamOut = conn
	}

	// Ensure global variables are used
	globalEventCount := &AppConfig.eventCount
	globalDataBuffer := &AppConfig.dataBuffer

	// Example of processing data (this part should be expanded based on actual requirements)
	handleConnection(AppConfig, connStreamOut, globalEventCount, globalDataBuffer)
}

func handleConnection(AppConfig Config, conn net.Conn, eventCount *int, dataBuffer *[]byte) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	remote, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		log.Printf("Error splitting host and port: %v", err)
	}
	log.Printf("Socket from " + remote + " connected")
	reader := bufio.NewReader(conn)
	for {
		// Read the HTTP request from the connection
		req, err := http.ReadRequest(reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %v", err)
				sendErrorResponse(conn, req)
			}
			break
		}

		// Extract method, headers, and payload
		method := req.Method
		headers := req.Header
		payload := make([]byte, req.ContentLength)
		_, err = io.ReadFull(req.Body, payload)
		if err != nil {
			log.Printf("Error reading payload: %v", err)
			sendErrorResponse(conn, req)
			break
		}
		req.Body.Close()

		// Log the extracted information
		var eventType string

		// Unmarshal the JSON payload into the struct
		var p Payload
		err = json.Unmarshal(payload, &p)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %v", err)
			sendErrorResponse(conn, req)
			return
		}
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
					log.Printf("Matched Trigger Event: %+v", triggerEvent)
					//TODO add the slurm integration here
					break
				}
			}
		}

		// Get the current Unix timestamp
		timestamp := float64(time.Now().Unix())
		// Append data to dataBuffer and increment eventCount
		*dataBuffer = append(*dataBuffer, payload...)
		*eventCount++
		eventCountMetric.WithLabelValues(remote, eventType).Inc()
		eventProcessingTimeMetric.WithLabelValues(remote, eventType).Set(timestamp)
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
		break
	}
}

type StreamOutput struct {
	Method  string
	Headers map[string]string
	Payload string
}

func sendErrorResponse(conn net.Conn, req *http.Request) {
	response := &http.Response{
		Status:        "500 Internal Server Error",
		StatusCode:    http.StatusInternalServerError,
		Proto:         req.Proto,
		ProtoMajor:    req.ProtoMajor,
		ProtoMinor:    req.ProtoMinor,
		Header:        make(http.Header),
		Body:          io.NopCloser(bytes.NewBufferString("Internal Server Error")),
		ContentLength: int64(len("Internal Server Error")),
	}
	response.Header.Set("Content-Type", "text/plain")
	err := response.Write(conn)
	if err != nil {
		log.Printf("Error writing error response: %v", err)
	}
}
