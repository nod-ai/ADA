# \SlurmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmV0039CancelJob**](SlurmAPI.md#SlurmV0039CancelJob) | **Delete** /slurm/v0.0.39/job/{job_id} | cancel or signal job
[**SlurmV0039DeleteNode**](SlurmAPI.md#SlurmV0039DeleteNode) | **Delete** /slurm/v0.0.39/node/{node_name} | delete node
[**SlurmV0039Diag**](SlurmAPI.md#SlurmV0039Diag) | **Get** /slurm/v0.0.39/diag | get diagnostics
[**SlurmV0039GetJob**](SlurmAPI.md#SlurmV0039GetJob) | **Get** /slurm/v0.0.39/job/{job_id} | get job info
[**SlurmV0039GetJobs**](SlurmAPI.md#SlurmV0039GetJobs) | **Get** /slurm/v0.0.39/jobs | get list of jobs
[**SlurmV0039GetNode**](SlurmAPI.md#SlurmV0039GetNode) | **Get** /slurm/v0.0.39/node/{node_name} | get node info
[**SlurmV0039GetNodes**](SlurmAPI.md#SlurmV0039GetNodes) | **Get** /slurm/v0.0.39/nodes | get all node info
[**SlurmV0039GetPartition**](SlurmAPI.md#SlurmV0039GetPartition) | **Get** /slurm/v0.0.39/partition/{partition_name} | get partition info
[**SlurmV0039GetPartitions**](SlurmAPI.md#SlurmV0039GetPartitions) | **Get** /slurm/v0.0.39/partitions | get all partition info
[**SlurmV0039GetReservation**](SlurmAPI.md#SlurmV0039GetReservation) | **Get** /slurm/v0.0.39/reservation/{reservation_name} | get reservation info
[**SlurmV0039GetReservations**](SlurmAPI.md#SlurmV0039GetReservations) | **Get** /slurm/v0.0.39/reservations | get all reservation info
[**SlurmV0039Ping**](SlurmAPI.md#SlurmV0039Ping) | **Get** /slurm/v0.0.39/ping | ping test
[**SlurmV0039SlurmctldGetLicenses**](SlurmAPI.md#SlurmV0039SlurmctldGetLicenses) | **Get** /slurm/v0.0.39/licenses | get all Slurm tracked license info
[**SlurmV0039SubmitJob**](SlurmAPI.md#SlurmV0039SubmitJob) | **Post** /slurm/v0.0.39/job/submit | submit new job
[**SlurmV0039UpdateJob**](SlurmAPI.md#SlurmV0039UpdateJob) | **Post** /slurm/v0.0.39/job/{job_id} | update job
[**SlurmV0039UpdateNode**](SlurmAPI.md#SlurmV0039UpdateNode) | **Post** /slurm/v0.0.39/node/{node_name} | update node properties
[**SlurmV0040DeleteJob**](SlurmAPI.md#SlurmV0040DeleteJob) | **Delete** /slurm/v0.0.40/job/{job_id} | cancel or signal job
[**SlurmV0040DeleteNode**](SlurmAPI.md#SlurmV0040DeleteNode) | **Delete** /slurm/v0.0.40/node/{node_name} | delete node
[**SlurmV0040GetDiag**](SlurmAPI.md#SlurmV0040GetDiag) | **Get** /slurm/v0.0.40/diag | get diagnostics
[**SlurmV0040GetJob**](SlurmAPI.md#SlurmV0040GetJob) | **Get** /slurm/v0.0.40/job/{job_id} | get job info
[**SlurmV0040GetJobs**](SlurmAPI.md#SlurmV0040GetJobs) | **Get** /slurm/v0.0.40/jobs | get list of jobs
[**SlurmV0040GetLicenses**](SlurmAPI.md#SlurmV0040GetLicenses) | **Get** /slurm/v0.0.40/licenses | get all Slurm tracked license info
[**SlurmV0040GetNode**](SlurmAPI.md#SlurmV0040GetNode) | **Get** /slurm/v0.0.40/node/{node_name} | get node info
[**SlurmV0040GetNodes**](SlurmAPI.md#SlurmV0040GetNodes) | **Get** /slurm/v0.0.40/nodes | get all node info
[**SlurmV0040GetPartition**](SlurmAPI.md#SlurmV0040GetPartition) | **Get** /slurm/v0.0.40/partition/{partition_name} | get partition info
[**SlurmV0040GetPartitions**](SlurmAPI.md#SlurmV0040GetPartitions) | **Get** /slurm/v0.0.40/partitions | get all partition info
[**SlurmV0040GetPing**](SlurmAPI.md#SlurmV0040GetPing) | **Get** /slurm/v0.0.40/ping | ping test
[**SlurmV0040GetReconfigure**](SlurmAPI.md#SlurmV0040GetReconfigure) | **Get** /slurm/v0.0.40/reconfigure | request slurmctld reconfigure
[**SlurmV0040GetReservation**](SlurmAPI.md#SlurmV0040GetReservation) | **Get** /slurm/v0.0.40/reservation/{reservation_name} | get reservation info
[**SlurmV0040GetReservations**](SlurmAPI.md#SlurmV0040GetReservations) | **Get** /slurm/v0.0.40/reservations | get all reservation info
[**SlurmV0040GetShares**](SlurmAPI.md#SlurmV0040GetShares) | **Get** /slurm/v0.0.40/shares | get fairshare info
[**SlurmV0040PostJob**](SlurmAPI.md#SlurmV0040PostJob) | **Post** /slurm/v0.0.40/job/{job_id} | update job
[**SlurmV0040PostJobSubmit**](SlurmAPI.md#SlurmV0040PostJobSubmit) | **Post** /slurm/v0.0.40/job/submit | submit new job
[**SlurmV0040PostNode**](SlurmAPI.md#SlurmV0040PostNode) | **Post** /slurm/v0.0.40/node/{node_name} | update node properties
[**SlurmV0041DeleteJob**](SlurmAPI.md#SlurmV0041DeleteJob) | **Delete** /slurm/v0.0.41/job/{job_id} | cancel or signal job
[**SlurmV0041DeleteNode**](SlurmAPI.md#SlurmV0041DeleteNode) | **Delete** /slurm/v0.0.41/node/{node_name} | delete node
[**SlurmV0041GetDiag**](SlurmAPI.md#SlurmV0041GetDiag) | **Get** /slurm/v0.0.41/diag | get diagnostics
[**SlurmV0041GetJob**](SlurmAPI.md#SlurmV0041GetJob) | **Get** /slurm/v0.0.41/job/{job_id} | get job info
[**SlurmV0041GetJobs**](SlurmAPI.md#SlurmV0041GetJobs) | **Get** /slurm/v0.0.41/jobs | get list of jobs
[**SlurmV0041GetLicenses**](SlurmAPI.md#SlurmV0041GetLicenses) | **Get** /slurm/v0.0.41/licenses | get all Slurm tracked license info
[**SlurmV0041GetNode**](SlurmAPI.md#SlurmV0041GetNode) | **Get** /slurm/v0.0.41/node/{node_name} | get node info
[**SlurmV0041GetNodes**](SlurmAPI.md#SlurmV0041GetNodes) | **Get** /slurm/v0.0.41/nodes | get all node info
[**SlurmV0041GetPartition**](SlurmAPI.md#SlurmV0041GetPartition) | **Get** /slurm/v0.0.41/partition/{partition_name} | get partition info
[**SlurmV0041GetPartitions**](SlurmAPI.md#SlurmV0041GetPartitions) | **Get** /slurm/v0.0.41/partitions | get all partition info
[**SlurmV0041GetPing**](SlurmAPI.md#SlurmV0041GetPing) | **Get** /slurm/v0.0.41/ping | ping test
[**SlurmV0041GetReconfigure**](SlurmAPI.md#SlurmV0041GetReconfigure) | **Get** /slurm/v0.0.41/reconfigure | request slurmctld reconfigure
[**SlurmV0041GetReservation**](SlurmAPI.md#SlurmV0041GetReservation) | **Get** /slurm/v0.0.41/reservation/{reservation_name} | get reservation info
[**SlurmV0041GetReservations**](SlurmAPI.md#SlurmV0041GetReservations) | **Get** /slurm/v0.0.41/reservations | get all reservation info
[**SlurmV0041GetShares**](SlurmAPI.md#SlurmV0041GetShares) | **Get** /slurm/v0.0.41/shares | get fairshare info
[**SlurmV0041PostJob**](SlurmAPI.md#SlurmV0041PostJob) | **Post** /slurm/v0.0.41/job/{job_id} | update job
[**SlurmV0041PostJobSubmit**](SlurmAPI.md#SlurmV0041PostJobSubmit) | **Post** /slurm/v0.0.41/job/submit | submit new job
[**SlurmV0041PostNode**](SlurmAPI.md#SlurmV0041PostNode) | **Post** /slurm/v0.0.41/node/{node_name} | update node properties



## SlurmV0039CancelJob

> Status SlurmV0039CancelJob(ctx, jobId).Signal(signal).Execute()

cancel or signal job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Slurm Job ID
	signal := "signal_example" // string | signal to send to job (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039CancelJob(context.Background(), jobId).Signal(signal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039CancelJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039CancelJob`: Status
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039CancelJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039CancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | signal to send to job | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039DeleteNode

> Status SlurmV0039DeleteNode(ctx, nodeName).Execute()

delete node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Slurm Node Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039DeleteNode`: Status
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039Diag

> V0039Diag SlurmV0039Diag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039Diag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039Diag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039Diag`: V0039Diag
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039Diag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039DiagRequest struct via the builder pattern


### Return type

[**V0039Diag**](V0039Diag.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetJob

> V0039JobsResponse SlurmV0039GetJob(ctx, jobId).Execute()

get job info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Slurm JobID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetJob`: V0039JobsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0039JobsResponse**](V0039JobsResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetJobs

> V0039JobsResponse SlurmV0039GetJobs(ctx).UpdateTime(updateTime).Execute()

get list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetJobs(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetJobs`: V0039JobsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039JobsResponse**](V0039JobsResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetNode

> V0039NodesResponse SlurmV0039GetNode(ctx, nodeName).Execute()

get node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Slurm Node Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetNode`: V0039NodesResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0039NodesResponse**](V0039NodesResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetNodes

> V0039NodesResponse SlurmV0039GetNodes(ctx).UpdateTime(updateTime).Execute()

get all node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetNodes(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetNodes`: V0039NodesResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039NodesResponse**](V0039NodesResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetPartition

> V0039PartitionsResponse SlurmV0039GetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

get partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partitionName := "partitionName_example" // string | Slurm Partition Name
	updateTime := int64(789) // int64 | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetPartition`: V0039PartitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0039PartitionsResponse**](V0039PartitionsResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetPartitions

> V0039PartitionsResponse SlurmV0039GetPartitions(ctx).UpdateTime(updateTime).Execute()

get all partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetPartitions(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetPartitions`: V0039PartitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039PartitionsResponse**](V0039PartitionsResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetReservation

> V0039ReservationsResponse SlurmV0039GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	reservationName := "reservationName_example" // string | Slurm Reservation Name
	updateTime := int64(789) // int64 | Filter if no reservation (not limited to reservation in URL) changed since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetReservation`: V0039ReservationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0039ReservationsResponse**](V0039ReservationsResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039GetReservations

> V0039ReservationsResponse SlurmV0039GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039GetReservations`: V0039ReservationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0039ReservationsResponse**](V0039ReservationsResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039Ping

> V0039Pings SlurmV0039Ping(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039Ping`: V0039Pings
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039PingRequest struct via the builder pattern


### Return type

[**V0039Pings**](V0039Pings.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039SlurmctldGetLicenses

> V0039LicensesInfo SlurmV0039SlurmctldGetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039SlurmctldGetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039SlurmctldGetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039SlurmctldGetLicenses`: V0039LicensesInfo
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039SlurmctldGetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039SlurmctldGetLicensesRequest struct via the builder pattern


### Return type

[**V0039LicensesInfo**](V0039LicensesInfo.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039SubmitJob

> V0039JobSubmissionResponse SlurmV0039SubmitJob(ctx).V0039JobSubmission(v0039JobSubmission).Execute()

submit new job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0039JobSubmission := *openapiclient.NewV0039JobSubmission() // V0039JobSubmission | submit new job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039SubmitJob(context.Background()).V0039JobSubmission(v0039JobSubmission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039SubmitJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039SubmitJob`: V0039JobSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039SubmitJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039SubmitJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0039JobSubmission** | [**V0039JobSubmission**](V0039JobSubmission.md) | submit new job | 

### Return type

[**V0039JobSubmissionResponse**](V0039JobSubmissionResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039UpdateJob

> V0039JobUpdateResponse SlurmV0039UpdateJob(ctx, jobId).V0039JobDescMsg(v0039JobDescMsg).Execute()

update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Slurm Job ID
	v0039JobDescMsg := *openapiclient.NewV0039JobDescMsg([]string{"Environment_example"}) // V0039JobDescMsg | update job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039UpdateJob(context.Background(), jobId).V0039JobDescMsg(v0039JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039UpdateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039UpdateJob`: V0039JobUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039UpdateJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039UpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0039JobDescMsg** | [**V0039JobDescMsg**](V0039JobDescMsg.md) | update job | 

### Return type

[**V0039JobUpdateResponse**](V0039JobUpdateResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0039UpdateNode

> Status SlurmV0039UpdateNode(ctx, nodeName).V0039UpdateNodeMsg(v0039UpdateNodeMsg).Execute()

update node properties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Slurm Node Name
	v0039UpdateNodeMsg := *openapiclient.NewV0039UpdateNodeMsg() // V0039UpdateNodeMsg | update node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0039UpdateNode(context.Background(), nodeName).V0039UpdateNodeMsg(v0039UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0039UpdateNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0039UpdateNode`: Status
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0039UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0039UpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0039UpdateNodeMsg** | [**V0039UpdateNodeMsg**](V0039UpdateNodeMsg.md) | update node | 

### Return type

[**Status**](Status.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040DeleteJob

> V0040OpenapiResp SlurmV0040DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

cancel or signal job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | JobId
	signal := "signal_example" // string | Signal to send to Job (optional)
	flags := "flags_example" // string | Signalling flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040DeleteJob`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | JobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040DeleteNode

> V0040OpenapiResp SlurmV0040DeleteNode(ctx, nodeName).Execute()

delete node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040DeleteNode`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetDiag

> V0040OpenapiDiagResp SlurmV0040GetDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetDiag`: V0040OpenapiDiagResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetDiagRequest struct via the builder pattern


### Return type

[**V0040OpenapiDiagResp**](V0040OpenapiDiagResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetJob

> V0040OpenapiJobInfoResp SlurmV0040GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

get job info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | JobId
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetJob`: V0040OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | JobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiJobInfoResp**](V0040OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetJobs

> V0040OpenapiJobInfoResp SlurmV0040GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetJobs`: V0040OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiJobInfoResp**](V0040OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetLicenses

> V0040OpenapiLicensesResp SlurmV0040GetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetLicenses`: V0040OpenapiLicensesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetLicensesRequest struct via the builder pattern


### Return type

[**V0040OpenapiLicensesResp**](V0040OpenapiLicensesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetNode

> V0040OpenapiNodesResp SlurmV0040GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

get node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetNode`: V0040OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiNodesResp**](V0040OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetNodes

> V0040OpenapiNodesResp SlurmV0040GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get all node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetNodes`: V0040OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0040OpenapiNodesResp**](V0040OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetPartition

> V0040OpenapiPartitionResp SlurmV0040GetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

get partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partitionName := "partitionName_example" // string | Slurm Partition Name
	updateTime := int64(789) // int64 | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetPartition`: V0040OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0040OpenapiPartitionResp**](V0040OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetPartitions

> V0040OpenapiPartitionResp SlurmV0040GetPartitions(ctx).UpdateTime(updateTime).Execute()

get all partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetPartitions(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetPartitions`: V0040OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0040OpenapiPartitionResp**](V0040OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetPing

> V0040OpenapiPingArrayResp SlurmV0040GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetPing`: V0040OpenapiPingArrayResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetPingRequest struct via the builder pattern


### Return type

[**V0040OpenapiPingArrayResp**](V0040OpenapiPingArrayResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetReconfigure

> V0040OpenapiResp SlurmV0040GetReconfigure(ctx).Execute()

request slurmctld reconfigure

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetReconfigure`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetReconfigureRequest struct via the builder pattern


### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetReservation

> V0040OpenapiReservationResp SlurmV0040GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	reservationName := "reservationName_example" // string | Slurm Reservation Name
	updateTime := int64(789) // int64 | Filter if no reservation (not limited to reservation in URL) changed since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetReservation`: V0040OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0040OpenapiReservationResp**](V0040OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetReservations

> V0040OpenapiReservationResp SlurmV0040GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetReservations`: V0040OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0040OpenapiReservationResp**](V0040OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040GetShares

> V0040OpenapiSharesResp SlurmV0040GetShares(ctx).Accounts(accounts).Users(users).Execute()

get fairshare info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accounts := "accounts_example" // string | Accounts to query (optional)
	users := "users_example" // string | Users to query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040GetShares`: V0040OpenapiSharesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**V0040OpenapiSharesResp**](V0040OpenapiSharesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040PostJob

> V0040OpenapiJobPostResponse SlurmV0040PostJob(ctx, jobId).V0040JobDescMsg(v0040JobDescMsg).Execute()

update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | JobId
	v0040JobDescMsg := *openapiclient.NewV0040JobDescMsg() // V0040JobDescMsg | update job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040PostJob(context.Background(), jobId).V0040JobDescMsg(v0040JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040PostJob`: V0040OpenapiJobPostResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | JobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0040JobDescMsg** | [**V0040JobDescMsg**](V0040JobDescMsg.md) | update job | 

### Return type

[**V0040OpenapiJobPostResponse**](V0040OpenapiJobPostResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040PostJobSubmit

> V0040OpenapiJobSubmitResponse SlurmV0040PostJobSubmit(ctx).V0040JobSubmitReq(v0040JobSubmitReq).Execute()

submit new job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040JobSubmitReq := *openapiclient.NewV0040JobSubmitReq() // V0040JobSubmitReq | submit new job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040PostJobSubmit(context.Background()).V0040JobSubmitReq(v0040JobSubmitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040PostJobSubmit`: V0040OpenapiJobSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040JobSubmitReq** | [**V0040JobSubmitReq**](V0040JobSubmitReq.md) | submit new job | 

### Return type

[**V0040OpenapiJobSubmitResponse**](V0040OpenapiJobSubmitResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0040PostNode

> V0040OpenapiResp SlurmV0040PostNode(ctx, nodeName).V0040UpdateNodeMsg(v0040UpdateNodeMsg).Execute()

update node properties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	v0040UpdateNodeMsg := *openapiclient.NewV0040UpdateNodeMsg() // V0040UpdateNodeMsg | update node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0040PostNode(context.Background(), nodeName).V0040UpdateNodeMsg(v0040UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0040PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0040PostNode`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0040PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0040PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0040UpdateNodeMsg** | [**V0040UpdateNodeMsg**](V0040UpdateNodeMsg.md) | update node | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteJob

> V0041OpenapiResp SlurmV0041DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

cancel or signal job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | JobId
	signal := "signal_example" // string | Signal to send to Job (optional)
	flags := "flags_example" // string | Signalling flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteJob`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | JobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteNode

> V0041OpenapiResp SlurmV0041DeleteNode(ctx, nodeName).Execute()

delete node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteNode`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetDiag

> V0041OpenapiDiagResp SlurmV0041GetDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetDiag`: V0041OpenapiDiagResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetDiagRequest struct via the builder pattern


### Return type

[**V0041OpenapiDiagResp**](V0041OpenapiDiagResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJob

> V0041OpenapiJobInfoResp SlurmV0041GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

get job info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | JobId
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJob`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | JobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJobs

> V0041OpenapiJobInfoResp SlurmV0041GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJobs`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetLicenses

> V0041OpenapiLicensesResp SlurmV0041GetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetLicenses`: V0041OpenapiLicensesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetLicensesRequest struct via the builder pattern


### Return type

[**V0041OpenapiLicensesResp**](V0041OpenapiLicensesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetNode

> V0041OpenapiNodesResp SlurmV0041GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

get node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetNode`: V0041OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiNodesResp**](V0041OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetNodes

> V0041OpenapiNodesResp SlurmV0041GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get all node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetNodes`: V0041OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiNodesResp**](V0041OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPartition

> V0041OpenapiPartitionResp SlurmV0041GetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

get partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partitionName := "partitionName_example" // string | Slurm Partition Name
	updateTime := int64(789) // int64 | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPartition`: V0041OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0041OpenapiPartitionResp**](V0041OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPartitions

> V0041OpenapiPartitionResp SlurmV0041GetPartitions(ctx).UpdateTime(updateTime).Execute()

get all partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPartitions(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPartitions`: V0041OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0041OpenapiPartitionResp**](V0041OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPing

> V0041OpenapiPingArrayResp SlurmV0041GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPing`: V0041OpenapiPingArrayResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPingRequest struct via the builder pattern


### Return type

[**V0041OpenapiPingArrayResp**](V0041OpenapiPingArrayResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReconfigure

> V0041OpenapiResp SlurmV0041GetReconfigure(ctx).Execute()

request slurmctld reconfigure

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReconfigure`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReconfigureRequest struct via the builder pattern


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReservation

> V0041OpenapiReservationResp SlurmV0041GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	reservationName := "reservationName_example" // string | Slurm Reservation Name
	updateTime := int64(789) // int64 | Filter if no reservation (not limited to reservation in URL) changed since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReservation`: V0041OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0041OpenapiReservationResp**](V0041OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReservations

> V0041OpenapiReservationResp SlurmV0041GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReservations`: V0041OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0041OpenapiReservationResp**](V0041OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetShares

> V0041OpenapiSharesResp SlurmV0041GetShares(ctx).Accounts(accounts).Users(users).Execute()

get fairshare info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accounts := "accounts_example" // string | Accounts to query (optional)
	users := "users_example" // string | Users to query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetShares`: V0041OpenapiSharesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**V0041OpenapiSharesResp**](V0041OpenapiSharesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJob

> V0041OpenapiJobPostResponse SlurmV0041PostJob(ctx, jobId).V0041JobDescMsg(v0041JobDescMsg).Execute()

update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | JobId
	v0041JobDescMsg := *openapiclient.NewV0041JobDescMsg() // V0041JobDescMsg | update job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJob(context.Background(), jobId).V0041JobDescMsg(v0041JobDescMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJob`: V0041OpenapiJobPostResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | JobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0041JobDescMsg** | [**V0041JobDescMsg**](V0041JobDescMsg.md) | update job | 

### Return type

[**V0041OpenapiJobPostResponse**](V0041OpenapiJobPostResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJobSubmit

> V0041OpenapiJobSubmitResponse SlurmV0041PostJobSubmit(ctx).V0041JobSubmitReq(v0041JobSubmitReq).Execute()

submit new job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041JobSubmitReq := *openapiclient.NewV0041JobSubmitReq() // V0041JobSubmitReq | submit new job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJobSubmit(context.Background()).V0041JobSubmitReq(v0041JobSubmitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJobSubmit`: V0041OpenapiJobSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041JobSubmitReq** | [**V0041JobSubmitReq**](V0041JobSubmitReq.md) | submit new job | 

### Return type

[**V0041OpenapiJobSubmitResponse**](V0041OpenapiJobSubmitResponse.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostNode

> V0041OpenapiResp SlurmV0041PostNode(ctx, nodeName).V0041UpdateNodeMsg(v0041UpdateNodeMsg).Execute()

update node properties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodeName := "nodeName_example" // string | Node name
	v0041UpdateNodeMsg := *openapiclient.NewV0041UpdateNodeMsg() // V0041UpdateNodeMsg | update node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostNode(context.Background(), nodeName).V0041UpdateNodeMsg(v0041UpdateNodeMsg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostNode`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0041UpdateNodeMsg** | [**V0041UpdateNodeMsg**](V0041UpdateNodeMsg.md) | update node | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

