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
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Redfish Event Listener/Exporter")

	// Setup configuration
	AppConfig := setupConfig()

	// Log the initialized config
	log.Printf("Initialized Config: %+v", AppConfig)

	// Channel to signal shutdown
	shutdownChan := make(chan struct{})

	// Subscribe the listener to the event stream for all servers
	subscriptionMap, err := CreateEventSubscriptionsForAllServers(AppConfig)
	if err != nil {
		// TODO detect if error is due to already being subscribed, and delete/re-add subscription.
		log.Fatalf("Failed to create subscriptions: %v", err)
	}

	// Start the listener
	go startListener(AppConfig, shutdownChan)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting metrics server on :2112")
		if err := http.ListenAndServe(":2112", nil); err != nil && err != http.ErrServerClosed {
			log.Printf("Metrics server error: %v", err)
		}
	}()

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal
	<-sigChan
	log.Println("Received shutdown signal. Closing listener...")

	// Signal shutdown to all components
	close(shutdownChan)

	// Give some time for the listener to close gracefully
	time.Sleep(time.Second)

	// Unsubscribe the listener from all servers
	log.Println("Unsubscribing from servers...")
	DeleteAllSubscriptions(AppConfig, subscriptionMap)

	// Perform any additional shutdown steps here
	log.Println("Shutdown complete")
}
