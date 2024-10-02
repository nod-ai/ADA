package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var eventCountMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "RedFishEvents_recieved",
		Help: "Total number of events processed",
	},
	[]string{"SourceIP", "EventType"}, // Define the labels you want to use
)

var eventProcessingTimeMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "RedFishEvents_processing_time",
		Help: "Time taken to process events",
	},
	[]string{"SourceIP", "EventType"}, // Define the labels you want to use
)

func init() {
	// Register the counter with Prometheus's default registry
	prometheus.MustRegister(eventCountMetric)
	// Register the gauge with Prometheus's default registry
	prometheus.MustRegister(eventProcessingTimeMetric)
}
