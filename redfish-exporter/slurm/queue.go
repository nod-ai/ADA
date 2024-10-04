package slurm

import (
	"context"
	"log"
	"strings"
)

const (
	Drain = "DrainNode"
)

type eventsActionReq struct {
	action        string
	slurmNodeName string
}

type SlurmQueue struct {
	ctx   context.Context
	queue chan *eventsActionReq
}

func InitSlurmQueue(ctx context.Context) *SlurmQueue {
	return &SlurmQueue{ctx: ctx, queue: make(chan *eventsActionReq)}
}

func (q *SlurmQueue) Add(action, slurmNodeName string) {
	q.queue <- &eventsActionReq{action: action, slurmNodeName: slurmNodeName}
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
			q.performEventAction(actionReq)
		}
	}
}

func (q *SlurmQueue) performEventAction(req *eventsActionReq) {
	if len(strings.TrimSpace(req.slurmNodeName)) == 0 {
		return
	}

	slurmClient := GetClient()
	if slurmClient == nil {
		return
	}

	if req.action == Drain {
		err := slurmClient.DrainNode(req.slurmNodeName)
		if err != nil {
			log.Printf("Error draining node: %v", err)
		}
	}
}
