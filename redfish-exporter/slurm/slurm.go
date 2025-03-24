package slurm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os/exec"
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
		jreq := c.apiClient.SlurmAPI.SlurmV0039UpdateNode(ctx, nodeName)
		req := slurmrestdapi.V0039UpdateNodeMsg{State: []string{"resume"}}
		jreq = jreq.V0039UpdateNodeMsg(req)
		res, resp, err := c.apiClient.SlurmAPI.SlurmV0039UpdateNodeExecute(jreq)
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

func (c *Client) DrainNodeWithAPI(nodeName, reason, excludeStr, scontrolPath string) error {
	apiCall := func() (interface{}, *http.Response, error) {
		uid := "0"
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		//slurm_v0039_update_node
		jreq := c.apiClient.SlurmAPI.SlurmV0039UpdateNode(ctx, nodeName)
		req := slurmrestdapi.V0039UpdateNodeMsg{State: []string{"drain"}, Reason: &reason, ReasonUid: &uid}

		jreq = jreq.V0039UpdateNodeMsg(req)
		res, resp, err := c.apiClient.SlurmAPI.SlurmV0039UpdateNodeExecute(jreq)
		cancel()
		if err != nil {
			return res, resp, err
		} else if resp.StatusCode != 200 {
			return res, resp, fmt.Errorf("invalid status code: %v", resp.StatusCode)
		}
		return res, resp, nil
	}

	curReason, err := c.GetNodeReasonWithAPI(nodeName)
	if err != nil {
		return err
	}
	log.Printf("node: %v, Reason: %v", nodeName, curReason)
	if strings.Contains(curReason, excludeStr) {
		return fmt.Errorf("%s: not draining node: %s | current reason: %s", ExlcudeReasonSet, nodeName, curReason)
	}
	_, resp, err := CallWithRetry(apiCall, maxRetries, baseDelay)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func DrainNodeWithScontrol(nodeName, reason, excludeStr, scontrolPath string) error {

	if excludeStr != "" {
		curReason, err := GetNodeReasonWithScontrol(nodeName, scontrolPath)
		if err != nil {
			log.Printf("GetNodeReasonWithScontrol returned err: %v\n", err)
			return err
		}

		if curReason != "" {
			re := regexp.MustCompile(excludeStr)
			match := re.FindAllString(curReason, -1)

			if len(match) != 0 {
				log.Printf("exlcudStr: %v, curReason: %v", excludeStr, curReason)
				log.Printf("match: %v | len: %v", match, len(match))
				return fmt.Errorf("%s: not draining node: %s | current reason: %s", ExlcudeReasonSet, nodeName, curReason)
			}
		}
	}
	cmd := fmt.Sprintf("%s update NodeName=%s State=DRAIN Reason=\"%s\"", scontrolPath, nodeName, reason)
	res := LocalCommandOutput(cmd)
	log.Printf("Drain node result: %s", res)
	return nil
}

func (c *Client) GetNodesWithAPI() ([]string, error) {
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

func (c *Client) GetNodeReasonWithAPI(nodeName string) (string, error) {
	var reason string
	apiCall := func() (interface{}, *http.Response, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		jreq := c.apiClient.SlurmAPI.SlurmV0039GetNode(ctx, nodeName)
		res, resp, err := c.apiClient.SlurmAPI.SlurmV0039GetNodeExecute(jreq)
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
		return reason, err
	}
	defer resp.Body.Close()

	temp := res.(*slurmrestdapi.V0039NodesResponse)
	nodes := temp.GetNodes()
	if len(nodes) != 1 {
		return reason, fmt.Errorf("GetNodeReason failed")
	}

	reason = *nodes[0].Reason
	log.Printf("[slurm] get node reasons(%s): %+v\n", nodeName, reason)
	return reason, nil
}

func GetNodeReasonWithScontrol(nodeName, scontrolPath string) (string, error) {
	type scontrolShowNode struct {
		Meta struct {
			Plugins struct {
				DataParser        string `json:"data_parser"`
				AccountingStorage string `json:"accounting_storage"`
			} `json:"plugins"`
			Command []string `json:"command"`
			Slurm   struct {
				Version struct {
					Major int `json:"major"`
					Micro int `json:"micro"`
					Minor int `json:"minor"`
				} `json:"version"`
				Release string `json:"release"`
			} `json:"Slurm"`
		} `json:"meta"`
		Nodes []struct {
			Architecture              string `json:"architecture"`
			BurstbufferNetworkAddress string `json:"burstbuffer_network_address"`
			Boards                    int    `json:"boards"`
			BootTime                  int    `json:"boot_time"`
			ClusterName               string `json:"cluster_name"`
			Cores                     int    `json:"cores"`
			SpecializedCores          int    `json:"specialized_cores"`
			CPUBinding                int    `json:"cpu_binding"`
			CPULoad                   struct {
				Set      bool `json:"set"`
				Infinite bool `json:"infinite"`
				Number   int  `json:"number"`
			} `json:"cpu_load"`
			FreeMem struct {
				Set      bool `json:"set"`
				Infinite bool `json:"infinite"`
				Number   int  `json:"number"`
			} `json:"free_mem"`
			CPUs            int    `json:"cpus"`
			EffectiveCPUs   int    `json:"effective_cpus"`
			SpecializedCPUs string `json:"specialized_cpus"`
			Energy          struct {
				AverageWatts       int `json:"average_watts"`
				BaseConsumedEnergy int `json:"base_consumed_energy"`
				ConsumedEnergy     int `json:"consumed_energy"`
				CurrentWatts       struct {
					Set      bool `json:"set"`
					Infinite bool `json:"infinite"`
					Number   int  `json:"number"`
				} `json:"current_watts"`
				PreviousConsumedEnergy int `json:"previous_consumed_energy"`
				LastCollected          int `json:"last_collected"`
			} `json:"energy"`
			ExternalSensors struct {
				ConsumedEnergy struct {
					Set      bool `json:"set"`
					Infinite bool `json:"infinite"`
					Number   int  `json:"number"`
				} `json:"consumed_energy"`
				Temperature struct {
					Set      bool `json:"set"`
					Infinite bool `json:"infinite"`
					Number   int  `json:"number"`
				} `json:"temperature"`
				EnergyUpdateTime int `json:"energy_update_time"`
				CurrentWatts     int `json:"current_watts"`
			} `json:"external_sensors"`
			Extra string `json:"extra"`
			Power struct {
				MaximumWatts struct {
					Set      bool `json:"set"`
					Infinite bool `json:"infinite"`
					Number   int  `json:"number"`
				} `json:"maximum_watts"`
				CurrentWatts    int `json:"current_watts"`
				TotalEnergy     int `json:"total_energy"`
				NewMaximumWatts int `json:"new_maximum_watts"`
				PeakWatts       int `json:"peak_watts"`
				LowestWatts     int `json:"lowest_watts"`
				NewJobTime      int `json:"new_job_time"`
				State           int `json:"state"`
				TimeStartDay    int `json:"time_start_day"`
			} `json:"power"`
			Features             []interface{} `json:"features"`
			ActiveFeatures       []interface{} `json:"active_features"`
			Gres                 string        `json:"gres"`
			GresDrained          string        `json:"gres_drained"`
			GresUsed             string        `json:"gres_used"`
			LastBusy             int           `json:"last_busy"`
			McsLabel             string        `json:"mcs_label"`
			SpecializedMemory    int           `json:"specialized_memory"`
			Name                 string        `json:"name"`
			NextStateAfterReboot []string      `json:"next_state_after_reboot"`
			Address              string        `json:"address"`
			Hostname             string        `json:"hostname"`
			State                []string      `json:"state"`
			OperatingSystem      string        `json:"operating_system"`
			Owner                string        `json:"owner"`
			Partitions           []string      `json:"partitions"`
			Port                 int           `json:"port"`
			RealMemory           int           `json:"real_memory"`
			Comment              string        `json:"comment"`
			Reason               string        `json:"reason"`
			ReasonChangedAt      int           `json:"reason_changed_at"`
			ReasonSetByUser      string        `json:"reason_set_by_user"`
			ResumeAfter          struct {
				Set      bool `json:"set"`
				Infinite bool `json:"infinite"`
				Number   int  `json:"number"`
			} `json:"resume_after"`
			Reservation     string  `json:"reservation"`
			AllocMemory     int     `json:"alloc_memory"`
			AllocCPUs       int     `json:"alloc_cpus"`
			AllocIdleCPUs   int     `json:"alloc_idle_cpus"`
			TresUsed        string  `json:"tres_used"`
			TresWeighted    float64 `json:"tres_weighted"`
			SlurmdStartTime int     `json:"slurmd_start_time"`
			Sockets         int     `json:"sockets"`
			Threads         int     `json:"threads"`
			TemporaryDisk   int     `json:"temporary_disk"`
			Weight          int     `json:"weight"`
			Tres            string  `json:"tres"`
			Version         string  `json:"version"`
		} `json:"nodes"`
		Warnings []interface{} `json:"warnings"`
		Errors   []interface{} `json:"errors"`
	}

	cmd := fmt.Sprintf("%s show node %s --json", scontrolPath, nodeName)
	ret := LocalCommandOutput(cmd)

	if ret == "" {
		return "", fmt.Errorf("failed to get current node reason")
	}

	res := scontrolShowNode{}
	err := json.Unmarshal([]byte(ret), &res)
	if err != nil {
		return "", err
	}
	if len(res.Nodes) != 1 {
		return "", fmt.Errorf("show node failed for %s", nodeName)
	}
	log.Printf("get node reasons(%s): %+v\n", nodeName, res.Nodes[0].Reason)
	return res.Nodes[0].Reason, nil
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

// LocalCommandOutput runs a command on a node and returns output in string format
func LocalCommandOutput(command string) string {
	log.Printf("Running cmd: %s\n", command)
	out, _ := exec.Command("bash", "-c", command).CombinedOutput()
	return strings.TrimSpace(string(out))
}
