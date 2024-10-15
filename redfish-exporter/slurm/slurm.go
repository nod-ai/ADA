package slurm

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"regexp"
	"strings"
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
	helper    map[string]reflect.Method
}

var singletonAPICl *Client // singleton client

func NewClient(slurmControlNode, slurmUser, slurmToken string) (*Client, error) {
	c := &Client{}
	slConfig := &SlurmServerConfig{
		URL:         slurmControlNode,
		Username:    defaultSlurmUsername,
		BearerToken: slurmToken,
	}
	if slurmUser != "" {
		slConfig.Username = slurmUser
	}
	cl := createRestClient(slConfig)

	// populate the methods required for ping and node update operations
	t := reflect.TypeOf(cl.SlurmAPI)
	postNodeRe := regexp.MustCompile(fmt.Sprintf(`%s$`, methodPostNode))
	pingRe := regexp.MustCompile(fmt.Sprintf(`%s$`, methodPing))
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if postNodeRe.MatchString(method.Name) {
			postNodeExecuteMethod, found := t.MethodByName(withExecuteSuffix(method.Name))
			if !found {
				return nil, fmt.Errorf("could not find PostNodeExecute method from Slurm REST APIs")
			}

			if _, found := c.helper[methodPostNode]; !found {
				c.helper[methodPostNode] = method
				c.helper[withExecuteSuffix(methodPostNode)] = postNodeExecuteMethod
			}
		} else if pingRe.MatchString(method.Name) {
			pingExecuteMethod, found := t.MethodByName(withExecuteSuffix(method.Name))
			if !found {
				return nil, fmt.Errorf("could not find PingExecute method from Slurm REST APIs")
			}

			if _, found := c.helper[methodPing]; !found {
				c.helper[methodPing] = method
				c.helper[withExecuteSuffix(methodPing)] = pingExecuteMethod
			}
		}
	}

	c.apiClient = cl

	log.Printf("[slurm] created slurm client for node: %v\n", slurmControlNode)
	err := c.getConnectionStatus()
	if err != nil {
		log.Printf("[slurm] error in getting the connection status of the slurm node: %v, err: %+v\n", slurmControlNode, err)
	}

	singletonAPICl = c
	return c, err
}

func GetClient() *Client {
	return singletonAPICl
}

func (c *Client) ResumeNode(nodeName string) error {
	return c.updateNodeState(nodeName, "resume")
}

func (c *Client) DrainNode(nodeName string) error {
	return c.updateNodeState(nodeName, "drain")
}

func (c *Client) getConnectionStatus() error {
	apiCall := func() (interface{}, *http.Response, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		// Step 1: Call the Ping method using reflection
		pingVals := c.helper["Ping"].Func.Call([]reflect.Value{
			reflect.ValueOf(c.apiClient.SlurmAPI),
			reflect.ValueOf(ctx),
		})

		// Check if the call produced results
		if len(pingVals) == 0 {
			return nil, nil, fmt.Errorf("Ping call returned no values")
		}

		// Step 2: Execute the Ping method with the request
		pingResp := c.helper["PingExecute"].Func.Call([]reflect.Value{
			reflect.ValueOf(c.apiClient.SlurmAPI),
			pingVals[0],
		})

		// Extract and return the response and error
		resp, _ := pingResp[1].Interface().(*http.Response)
		err, _ := pingResp[2].Interface().(error)
		return pingResp[0].Interface(), resp, err
	}

	_, resp, err := CallWithRetry(apiCall, maxRetries, baseDelay)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Printf("[slurm] ping success: %v\n", resp.StatusCode)
	return nil
}

func (c *Client) updateNodeState(nodeName, state string) error {
	apiCall := func() (interface{}, *http.Response, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		// Step 1: Call the PostNode method using reflection
		postNodeVals := c.helper["PostNode"].Func.Call([]reflect.Value{
			reflect.ValueOf(c.apiClient.SlurmAPI),
			reflect.ValueOf(ctx),
			reflect.ValueOf(nodeName),
		})

		// Check if the call produced results
		if len(postNodeVals) == 0 {
			return nil, nil, fmt.Errorf("PostNode call returned no values")
		}

		// Step 2: Find and call the UpdateNodeMsg method on the request object
		newInstance := reflect.New(postNodeVals[0].Type()).Elem()
		instanceType := newInstance.Type()

		for i := 0; i < instanceType.NumMethod(); i++ {
			method := instanceType.Method(i)
			if strings.Contains(method.Name, "UpdateNodeMsg") {
				// Create a new UpdateNodeMsg request
				updateNodeMsgReq := createUpdateNodeMsgRequest(method.Type)
				if updateNodeMsgReq.IsValid() {
					updateNodeMsgReq.FieldByName("State").Set(reflect.ValueOf([]string{state}))
				}

				// Step 3: Call UpdateNodeMsg with the request
				updatedNodeVals := method.Func.Call([]reflect.Value{postNodeVals[0], updateNodeMsgReq})
				if len(updatedNodeVals) == 0 {
					return nil, nil, fmt.Errorf("UpdateNodeMsg call returned no values")
				}

				// Step 4: Execute the PostNode method with the updated request
				postNodeResp := c.helper["PostNodeExecute"].Func.Call([]reflect.Value{
					reflect.ValueOf(c.apiClient.SlurmAPI),
					updatedNodeVals[0],
				})

				if len(postNodeResp) < 3 {
					return nil, nil, fmt.Errorf("PostNodeExecute call returned insufficient values")
				}

				// Extract and return the response and error
				resp, _ := postNodeResp[1].Interface().(*http.Response)
				err, _ := postNodeResp[2].Interface().(error)
				return postNodeResp[0].Interface(), resp, err
			}
		}

		return nil, nil, fmt.Errorf("no suitable UpdateNodeMsg method found")
	}

	// Retry the API call
	_, resp, err := CallWithRetry(apiCall, maxRetries, baseDelay)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Helper function to create an UpdateNodeMsg request using reflection
func createUpdateNodeMsgRequest(methodType reflect.Type) reflect.Value {
	for j := 1; j < methodType.NumIn(); j++ { // Start from 1 to skip the receiver
		paramType := methodType.In(j)
		if strings.Contains(paramType.Name(), "UpdateNodeMsg") {
			return reflect.New(paramType).Elem()
		}
	}
	return reflect.Value{}
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
