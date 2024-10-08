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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var EventCountMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "RedFishEvents_recieved",
		Help: "Total number of events processed",
	},
	[]string{"SourceIP", "EventType"}, // Define the labels you want to use
)

var EventProcessingTimeMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "RedFishEvents_processing_time",
		Help: "Time taken to process events",
	},
	[]string{"SourceIP", "EventType"}, // Define the labels you want to use
)

var SlurmAPIFailureMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "SlurmAPI_failure",
		Help: "Total number of Slurm API calls that failed",
	},
	[]string{"SourceIP", "SlurmNodeName", "EventSeverity", "EventAction"}, // Define the labels you want to use
)

var SlurmAPISuccessMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "SlurmAPI_success",
		Help: "Total number of Slurm API calls that succeeded",
	},
	[]string{"SourceIP", "SlurmNodeName", "EventSeverity", "EventAction"}, // Define the labels you want to use
)

func init() {
	// Register the counter with Prometheus's default registry
	prometheus.MustRegister(EventCountMetric)
	// Register the gauge with Prometheus's default registry
	prometheus.MustRegister(EventProcessingTimeMetric)
	prometheus.MustRegister(SlurmAPIFailureMetric)
	prometheus.MustRegister(SlurmAPISuccessMetric)
}
