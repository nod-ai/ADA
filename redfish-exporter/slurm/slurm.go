package slurm

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nod-ai/ADA/redfish-exporter/api/generated/slurmrestdapi"
)

const (
	slurmRestClientTimeout = time.Minute * 5
	defaultScript          = "#!/bin/bash\n"
)

type SlurmServerConfig struct {
	URL         string
	Username    string
	BearerToken string
}

type Client struct {
	apiClient *slurmrestdapi.APIClient // slurm URL to client mapping
}

var apiCl *Client // singleton client

func NewClient(slurmControlNode, slurmToken string) (*Client, error) {
	slConfig := &SlurmServerConfig{
		URL:         slurmControlNode,
		Username:    "root",
		BearerToken: slurmToken,
	}
	cl := createRestClient(slConfig)
	c := &Client{apiClient: cl}

	log.Printf("[slurm] created slurm client for node: %v\n", slurmControlNode)
	err := c.getConnectionStatus()
	if err != nil {
		log.Printf("[slurm] error in getting the connection status of the slurm node: %v, err: %+v\n", slurmControlNode, err)
	}

	apiCl = c
	return c, err
}

func GetClient() *Client {
	return apiCl
}

func (c *Client) ResumeNode(nodeName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	jreq := c.apiClient.SlurmAPI.SlurmV0040PostNode(ctx, nodeName)
	req := slurmrestdapi.V0040UpdateNodeMsg{State: []string{"resume"}}
	jreq = jreq.V0040UpdateNodeMsg(req)
	_, resp, err := c.apiClient.SlurmAPI.SlurmV0040PostNodeExecute(jreq)
	cancel()
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code: %v", resp.StatusCode)
	}

	return nil
}

func (c *Client) DrainNode(nodeName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	jreq := c.apiClient.SlurmAPI.SlurmV0040PostNode(ctx, nodeName)
	req := slurmrestdapi.V0040UpdateNodeMsg{State: []string{"drain"}}
	jreq = jreq.V0040UpdateNodeMsg(req)
	_, resp, err := c.apiClient.SlurmAPI.SlurmV0040PostNodeExecute(jreq)
	cancel()
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code: %v", resp.StatusCode)
	}

	return nil
}

func GetNodes(client *slurmrestdapi.APIClient) ([]string, error) {
	res := []string{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	jreq := client.SlurmAPI.SlurmV0039GetNodes(ctx)
	nodes, resp, err := client.SlurmAPI.SlurmV0039GetNodesExecute(jreq)
	cancel()
	if err != nil {
		return res, err
	} else if resp.StatusCode != 200 {
		return res, fmt.Errorf("invalid status code: %v", resp.StatusCode)
	}

	log.Printf("[slurm] get nodes: %+v\n", nodes)
	for _, node := range nodes.GetNodes() {
		res = append(res, *node.Name)
	}
	return res, nil
}

func (c *Client) getConnectionStatus() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	jreq := c.apiClient.SlurmAPI.SlurmV0039Ping(ctx)
	_, resp, err := c.apiClient.SlurmAPI.SlurmV0039PingExecute(jreq)
	cancel()
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code: %v", resp.StatusCode)
	}

	log.Printf("[slurm] ping success: %v\n", resp.StatusCode)
	return nil
}

func createRestClient(c *SlurmServerConfig) *slurmrestdapi.APIClient {
	cfg := slurmrestdapi.NewConfiguration()
	cfg.HTTPClient = &http.Client{Timeout: slurmRestClientTimeout}
	cfg.Scheme = "http"
	cfg.Host = c.URL

	cfg.AddDefaultHeader("X-SLURM-USER-NAME", c.Username)
	cfg.AddDefaultHeader("X-SLURM-USER-TOKEN", c.BearerToken)

	client := slurmrestdapi.NewAPIClient(cfg)
	return client
}
