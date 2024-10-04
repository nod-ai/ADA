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
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/nod-ai/ADA/redfish-exporter/slurm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		enableSlurm = flag.Bool("enable-slurm", false, "Enable slurm")
	)
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Redfish Event Listener/Exporter")

	// Setup configuration
	AppConfig := setupConfig()

	// Log the initialized config
	log.Printf("Initialized Config: %+v", AppConfig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var slurmQueue *slurm.SlurmQueue
	if *enableSlurm {
		if len(strings.TrimSpace(AppConfig.SlurmToken)) == 0 {
			log.Fatalf("Provide slurm token to enable slurm")
		}
		if len(strings.TrimSpace(AppConfig.SlurmControlNode)) == 0 {
			log.Fatalf("Provide slurm control node IP:Port to enable slurm")
		}
		_, err := slurm.NewClient(AppConfig.SlurmControlNode, AppConfig.SlurmToken)
		if err != nil {
			log.Fatalf("failed to create slurm client, err: %+v", err)
		}

		slurmQueue = slurm.InitSlurmQueue(ctx)
		go slurmQueue.ProcessEventActionQueue()
	}

	// Subscribe the listener to the event stream for all servers
	subscriptionMap, err := CreateSubscriptionsForAllServers(AppConfig.RedfishServers, AppConfig.SubscriptionPayload)
	if err != nil {
		// TODO detect if error is due to already being subscribed, and delete/re-add subscription.
		log.Fatalf("Failed to create subscriptions: %v", err)
	}

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the listener
	listener := NewServer(AppConfig.SystemInformation.ListenerIP, AppConfig.SystemInformation.ListenerPort, slurmQueue)
	go func() {
		if err := listener.Start(AppConfig); err != nil {
			log.Printf("Server error: %v", err)
			close(sigChan)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("Starting metrics server on :%d", AppConfig.SystemInformation.MetricsPort)
		portStr := strconv.Itoa(AppConfig.SystemInformation.MetricsPort)
		if err := http.ListenAndServe(":"+portStr, nil); err != nil && err != http.ErrServerClosed {
			log.Printf("Metrics server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-sigChan
	log.Println("Received shutdown signal. Closing listener...")

	// Signal shutdown to all components
	close(listener.shutdownChan)

	// Give some time for the listener to close gracefully
	time.Sleep(time.Second)

	// Unsubscribe the listener from all servers
	log.Println("Unsubscribing from servers...")
	DeleteSubscriptionsFromAllServers(AppConfig.RedfishServers, subscriptionMap)

	cancel()

	time.Sleep(time.Second)

	// Perform any additional shutdown steps here
	log.Println("Shutdown complete")
}
