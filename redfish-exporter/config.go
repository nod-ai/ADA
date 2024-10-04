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
	"crypto/tls"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	DefaultListenerPort = "8080"
	DefaultMetricsPort  = "2112"
	DefaultUseSSL       = "false"
)

type Config struct {
	Information struct {
		Updated     string
		Description string
	}
	SystemInformation struct {
		ListenerIP   string
		ListenerPort string
		UseSSL       bool
		MetricsPort  int
	}
	CertificateDetails struct {
		CertFile string
		KeyFile  string
	}
	SlurmToken          string
	SlurmControlNode    string
	SubscriptionPayload SubscriptionPayload
	RedfishServers      []RedfishServer
	TriggerEvents       []TriggerEvent
	context             *tls.Config
	eventCount          int
	dataBuffer          []byte
}

type TriggerEvent struct {
	EventId string `json:"EventId"`
	Action  string `json:"Action"`
}

func setupConfig() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file. Assuming environment variables are set.")
	}

	// Initialize Config object
	var AppConfig Config

	// Read environment variables
	AppConfig.Information.Updated = os.Getenv("UPDATED")
	AppConfig.Information.Description = os.Getenv("DESCRIPTION")
	AppConfig.SystemInformation.ListenerIP = os.Getenv("LISTENER_IP")

	// Read and parse LISTENER_PORT with a default value
	listenerPort := os.Getenv("LISTENER_PORT")
	if listenerPort == "" {
		listenerPort = DefaultListenerPort
	}
	AppConfig.SystemInformation.ListenerPort = listenerPort

	// Metrics Port Configuration
	metricsPortStr := os.Getenv("METRICS_PORT")
	if metricsPortStr == "" {
		metricsPortStr = DefaultMetricsPort
	}
	metricsPort, err := strconv.Atoi(metricsPortStr)
	if err != nil {
		log.Fatalf("Failed to parse METRICS_PORT: %v", err)
	}
	AppConfig.SystemInformation.MetricsPort = metricsPort
	// Read and parse USE_SSL with a default value
	useSSLStr := os.Getenv("USE_SSL")
	if useSSLStr == "" {
		useSSLStr = DefaultUseSSL
	}
	useSSL, err := strconv.ParseBool(useSSLStr)
	if err != nil {
		log.Fatalf("Failed to parse USE_SSL: %v", err)
	}
	AppConfig.SystemInformation.UseSSL = useSSL

	AppConfig.CertificateDetails.CertFile = os.Getenv("CERTFILE")
	AppConfig.CertificateDetails.KeyFile = os.Getenv("KEYFILE")

	AppConfig.SlurmToken = os.Getenv("SLURM_TOKEN")
	AppConfig.SlurmControlNode = os.Getenv("SLURM_CONTROL_NODE")

	subscriptionPayloadJSON := os.Getenv("SUBSCRIPTION_PAYLOAD")
	if err := json.Unmarshal([]byte(subscriptionPayloadJSON), &AppConfig.SubscriptionPayload); err != nil {
		log.Fatalf("Failed to parse SUBSCRIPTION_PAYLOAD: %v", err)
	}

	triggerEventsJSON := os.Getenv("TRIGGER_EVENTS")
	if triggerEventsJSON != "" {
		err = json.Unmarshal([]byte(triggerEventsJSON), &AppConfig.TriggerEvents)
		if err != nil {
			log.Fatalf("Failed to unmarshal TRIGGER_EVENTS: %v", err)
		}
	}
	// Read and parse the REDFISH_SERVERS environment variable
	redfishServersJSON := os.Getenv("REDFISH_SERVERS")
	if redfishServersJSON == "" {
		log.Println("REDFISH_SERVERS environment variable is not set or is empty")
		return AppConfig
	}
	if err := json.Unmarshal([]byte(redfishServersJSON), &AppConfig.RedfishServers); err != nil {
		log.Fatalf("Failed to parse REDFISH_SERVERS: %v", err)
	}

	return AppConfig
}
