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
	"fmt"
	"log"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
)

// Create a new connection to a redfish server
func newRedfishClient(server Server) (*gofish.APIClient, error) {
	clientConfig := gofish.ClientConfig{
		Endpoint: server.IP,
		Username: server.Username,
		Password: server.Password,
		Insecure: true, // TODO Set Based on login type
	}

	client, err := gofish.Connect(clientConfig)
	if err != nil {
		log.Printf("Error connecting to server %s: %v", server.IP, err)
		return nil, err
	}

	log.Printf("Successfully connected to server %s", server.IP)
	return client, nil
}

// Create subscriptions for all servers and return their URIs
// Rollback if any subscription attempt fails
func CreateEventSubscriptionsForAllServers(AppConfig Config) (map[string]string, error) {
	subscriptionMap := make(map[string]string)

	// TODO Add concurrency
	for _, server := range AppConfig.ServerInformation.Servers {
		// Establish a connection to the server
		c, err := newRedfishClient(server)
		if err != nil {
			log.Printf("Failed to connect to server %s: %v", server.IP, err)
			// Rollback subscriptions in case of failure
			DeleteAllSubscriptions(AppConfig, subscriptionMap)
			return nil, fmt.Errorf("subscription failed on server %s: %v, cleaning up previous subscriptions", server.IP, err)
		}

		defer c.Logout()

		eventService, err := c.Service.EventService()
		if err != nil {
			log.Printf("Failed to get event service for server %s: %v", server.IP, err)
			DeleteAllSubscriptions(AppConfig, subscriptionMap)
			return nil, fmt.Errorf("subscription failed on server %s, cleaning up previous subscriptions", server.IP)
		}

		var subscriptionURI string
		// Decide which subscription to create based on the Redfish version
		if isV1_5() {
			subscriptionURI, err = createV1_5Subscription(eventService, AppConfig)
		} else {
			subscriptionURI, err = createLegacySubscription(eventService, AppConfig)
		}

		if err != nil {
			log.Printf("Failed to create event subscription on server %s: %v", server.IP, err)
			DeleteAllSubscriptions(AppConfig, subscriptionMap)
			return nil, fmt.Errorf("subscription failed on server %s, cleaning up previous subscriptions", server.IP)
		}

		log.Printf("Successfully created subscription on server %s: %s", server.IP, subscriptionURI)
		subscriptionMap[server.IP] = subscriptionURI
	}

	return subscriptionMap, nil

}

func isV1_5() bool {
	// TODO Logic to determine if Redfish server is <v1.5 or higher
	// We assume false until we get version info on the servers.
	return false
}

// Create V1.5 subscription
func createV1_5Subscription(eventService *redfish.EventService, AppConfig Config) (string, error) {
	subscriptionURI, err := eventService.CreateEventSubscriptionInstance(
		AppConfig.SubscriptionPayload.Destination,
		AppConfig.SubscriptionPayload.RegistryPrefixes,
		AppConfig.SubscriptionPayload.ResourceTypes,
		AppConfig.SubscriptionPayload.HTTPHeaders,
		AppConfig.SubscriptionPayload.Protocol,
		AppConfig.SubscriptionPayload.Context,
		AppConfig.SubscriptionPayload.DeliveryRetryPolicy,
		AppConfig.SubscriptionPayload.Oem,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create v1.5 subscription: %w", err)
	}

	return subscriptionURI, nil
}

// Create legacy subscription
func createLegacySubscription(eventService *redfish.EventService, AppConfig Config) (string, error) {
	subscriptionURI, err := eventService.CreateEventSubscription(
		AppConfig.SubscriptionPayload.Destination,
		AppConfig.SubscriptionPayload.EventTypes,
		AppConfig.SubscriptionPayload.HTTPHeaders,
		AppConfig.SubscriptionPayload.Protocol,
		AppConfig.SubscriptionPayload.Context,
		AppConfig.SubscriptionPayload.Oem,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create legacy subscription: %w", err)
	}

	return subscriptionURI, nil
}

// Delete all event subscriptions stored in the map
func DeleteAllSubscriptions(AppConfig Config, subscriptionMap map[string]string) {
	for serverIP, subscriptionURI := range subscriptionMap {
		server := getServerCredentials(AppConfig, serverIP)
		c, err := newRedfishClient(server)
		if err != nil {
			log.Printf("Failed to connect to server %s for cleanup: %v", serverIP, err)
			continue
		}

		defer c.Logout()

		eventService, err := c.Service.EventService()
		if err != nil {
			log.Printf("Failed to get event service for cleanup on server %s: %v", serverIP, err)
			continue
		}

		err = eventService.DeleteEventSubscription(subscriptionURI)
		if err != nil {
			log.Printf("Failed to delete event subscription on server %s: %v", serverIP, err)
		} else {
			log.Printf("Successfully deleted event subscription from server %s", serverIP)
		}
	}
}

// Retrieve the server's credentials from the config based on IP
func getServerCredentials(AppConfig Config, serverIP string) Server {
	for _, server := range AppConfig.ServerInformation.Servers {
		if server.IP == serverIP {
			return server
		}
	}
	return Server{}
}
