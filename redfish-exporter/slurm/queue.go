package slurm

import (
	"context"
	"log"
	"strings"

	"github.com/nod-ai/ADA/redfish-exporter/metrics"
)

const (
	Drain = "DrainNode"
)

type eventsActionReq struct {
	redfishServerIP string
	slurmNodeName   string
	severity        string
	action          string
}

type SlurmQueue struct {
	ctx   context.Context
	queue chan *eventsActionReq
}

func InitSlurmQueue(ctx context.Context) *SlurmQueue {
	return &SlurmQueue{ctx: ctx, queue: make(chan *eventsActionReq)}
}

func (q *SlurmQueue) Add(redfishServerIP, slurmNodeName, severity, action string) {
	q.queue <- &eventsActionReq{
		redfishServerIP: redfishServerIP,
		slurmNodeName:   slurmNodeName,
		severity:        severity,
		action:          action,
	}
}

func (q *SlurmQueue) ProcessEventActionQueue() {
	log.Println("Starting slurm queue")
	for {
		select {
		case <-q.ctx.Done():
			log.Printf("Context done, stopping slurm queue processing")
			close(q.queue)
			return
		case actionReq := <-q.queue:
			log.Printf("Processing events action req from slurm queue: %v", actionReq)
			if err := q.performEventAction(actionReq); err != nil {
				metrics.SlurmAPIFailureMetric.WithLabelValues(
					actionReq.redfishServerIP,
					actionReq.slurmNodeName,
					actionReq.severity,
					actionReq.action).Inc()
				return
			}
			metrics.SlurmAPISuccessMetric.
				WithLabelValues(
					actionReq.redfishServerIP,
					actionReq.slurmNodeName,
					actionReq.severity,
					actionReq.action).Inc()
		}
	}
}

func (q *SlurmQueue) performEventAction(req *eventsActionReq) error {
	if len(strings.TrimSpace(req.slurmNodeName)) == 0 {
		return nil
	}

	slurmClient := GetClient()
	if slurmClient == nil {
		return nil
	}

	if req.action == Drain {
		err := slurmClient.DrainNode(req.slurmNodeName)
		if err != nil {
			log.Printf("Error draining node: %v", err)
			return err
		}
	}

	log.Printf("Performed action: %v on slurm node: %v successfully", req.action, req.slurmNodeName)
	return nil
}
