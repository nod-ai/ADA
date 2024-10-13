package slurm

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/nod-ai/ADA/redfish-exporter/api/generated/slurmrestdapi"
)

var (
	defaultSlurmUsername   = "root"
	slurmRestClientTimeout = time.Minute * 5
	maxRetries             = 5
	baseDelay              = 1 * time.Second
)

// Retryable is a function type that represents any API call function you we want to retry
type Retryable func() (interface{}, *http.Response, error)

// CallWithRetry is a generic retry function that takes a function, retry count, and delay
func CallWithRetry(call Retryable, maxRetries int, baseDelay time.Duration) (interface{}, *http.Response, error) {
	var resp *http.Response
	var err error
	var res interface{}

	for i := 0; i < maxRetries; i++ {
		res, resp, err = call()
		if err == nil && resp.StatusCode == http.StatusOK {
			return res, resp, nil // success
		}

		// exponential backoff delay
		delay := time.Duration(math.Pow(2, float64(i))) * baseDelay
		log.Printf("attempt %d: Error: %v, retrying in %v\n", i+1, err, delay)
		time.Sleep(delay)
	}

	return nil, nil, fmt.Errorf("after %d retries, error: %w", maxRetries, err)
}

type SlurmServerConfig struct {
	URL         string
	Username    string
	BearerToken string
}

type Client struct {
	apiClient *slurmrestdapi.APIClient // slurm URL to client mapping
}

var apiCl *Client // singleton client

func NewClient(slurmControlNode, slurmUser, slurmToken string) (*Client, error) {
	slConfig := &SlurmServerConfig{
		URL:         slurmControlNode,
		Username:    defaultSlurmUsername,
		BearerToken: slurmToken,
	}
	if slurmUser != "" {
		slConfig.Username = slurmUser
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
	apiCall := func() (interface{}, *http.Response, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		jreq := c.apiClient.SlurmAPI.SlurmV0040PostNode(ctx, nodeName)
		req := slurmrestdapi.V0040UpdateNodeMsg{State: []string{"resume"}}
		jreq = jreq.V0040UpdateNodeMsg(req)
		res, resp, err := c.apiClient.SlurmAPI.SlurmV0040PostNodeExecute(jreq)
		cancel()
		if err != nil {
			return res, resp, err
		} else if resp.StatusCode != 200 {
			return res, resp, fmt.Errorf("invalid status code: %v", resp.StatusCode)
		}
		return res, resp, nil
	}

	_, resp, err := CallWithRetry(apiCall, maxRetries, baseDelay)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) DrainNode(nodeName string) error {
	apiCall := func() (interface{}, *http.Response, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		jreq := c.apiClient.SlurmAPI.SlurmV0040PostNode(ctx, nodeName)
		req := slurmrestdapi.V0040UpdateNodeMsg{State: []string{"drain"}}
		jreq = jreq.V0040UpdateNodeMsg(req)
		res, resp, err := c.apiClient.SlurmAPI.SlurmV0040PostNodeExecute(jreq)
		cancel()
		if err != nil {
			return res, resp, err
		} else if resp.StatusCode != 200 {
			return res, resp, fmt.Errorf("invalid status code: %v", resp.StatusCode)
		}
		return res, resp, nil
	}

	_, resp, err := CallWithRetry(apiCall, maxRetries, baseDelay)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) GetNodes() ([]string, error) {
	var nodes []string
	apiCall := func() (interface{}, *http.Response, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		jreq := c.apiClient.SlurmAPI.SlurmV0039GetNodes(ctx)
		res, resp, err := c.apiClient.SlurmAPI.SlurmV0039GetNodesExecute(jreq)
		cancel()
		if err != nil {
			return res, resp, err
		} else if resp.StatusCode != 200 {
			return res, resp, fmt.Errorf("invalid status code: %v", resp.StatusCode)
		}
		return res, resp, nil
	}

	res, resp, err := CallWithRetry(apiCall, maxRetries, baseDelay)
	if err != nil {
		return nodes, err
	}
	defer resp.Body.Close()

	log.Printf("[slurm] get nodes: %+v\n", nodes)
	temp := res.(*slurmrestdapi.V0039NodesResponse)
	for _, node := range temp.GetNodes() {
		nodes = append(nodes, *node.Name)
	}
	return nodes, nil
}

func (c *Client) getConnectionStatus() error {
	apiCall := func() (interface{}, *http.Response, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		jreq := c.apiClient.SlurmAPI.SlurmV0039Ping(ctx)
		res, resp, err := c.apiClient.SlurmAPI.SlurmV0039PingExecute(jreq)
		cancel()
		if err != nil {
			return res, resp, err
		} else if resp.StatusCode != 200 {
			return res, resp, fmt.Errorf("invalid status code: %v", resp.StatusCode)
		}
		return res, resp, nil
	}

	_, resp, err := CallWithRetry(apiCall, maxRetries, baseDelay)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

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
