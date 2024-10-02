package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/stmcginnis/gofish/redfish"
)

type Config struct {
	Information struct {
		Updated     string
		Description string
	}
	SystemInformation struct {
		ListenerIP   string
		ListenerPort int
		UseSSL       bool
	}
	CertificateDetails struct {
		CertFile string
		KeyFile  string
	}
	SubscriptionPayload SubscriptionPayload
	ServerInformation   struct {
		Servers []Server
	}
	TriggerEvents []TriggerEvent
	context       *tls.Config
	eventCount    int
	dataBuffer    []byte
}

type SubscriptionPayload struct {
	Destination         string                           `json:"Destination,omitempty"`
	EventTypes          []redfish.EventType              `json:"EventTypes,omitempty"`
	RegistryPrefixes    []string                         `json:"RegistryPrefixes,omitempty"`
	ResourceTypes       []string                         `json:"ResourceTypes,omitempty"`
	DeliveryRetryPolicy redfish.DeliveryRetryPolicy      `json:"DeliveryRetryPolicy,omitempty"`
	HTTPHeaders         map[string]string                `json:"HttpHeaders,omitempty"`
	Oem                 interface{}                      `json:"Oem,omitempty"`
	Protocol            redfish.EventDestinationProtocol `json:"Protocol,omitempty"`
	Context             string                           `json:"Context,omitempty"`
}

type Server struct {
	IP        string `json:"ip"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LoginType string `json:"loginType"`
}

type TriggerEvent struct {
	EventId   string `json:"EventId"`
	EventType string `json:"EventType"`
	Severity  string `json:"Severity"`
}

const (
	DefaultListenerPort = "8080"
	DefaultUseSSL       = "false"
)

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
	listenerPortStr := os.Getenv("LISTENER_PORT")
	if listenerPortStr == "" {
		listenerPortStr = DefaultListenerPort
	}
	listenerPort, err := strconv.Atoi(listenerPortStr)
	if err != nil {
		log.Fatalf("Failed to parse LISTENER_PORT: %v", err)
	}
	AppConfig.SystemInformation.ListenerPort = listenerPort

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
	// Read and parse the SERVERS environment variable
	serversJSON := os.Getenv("SERVERS")
	if serversJSON == "" {
		log.Println("SERVERS environment variable is not set or is empty")
		return AppConfig
	}
	if err := json.Unmarshal([]byte(serversJSON), &AppConfig.ServerInformation.Servers); err != nil {
		log.Fatalf("Failed to parse SERVERS: %v", err)
	}

	return AppConfig
}
