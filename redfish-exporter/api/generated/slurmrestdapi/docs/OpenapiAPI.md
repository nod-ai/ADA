# \OpenapiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiGet**](OpenapiAPI.md#OpenapiGet) | **Get** /openapi | Retrieve OpenAPI Specification
[**OpenapiJsonGet**](OpenapiAPI.md#OpenapiJsonGet) | **Get** /openapi.json | Retrieve OpenAPI Specification
[**OpenapiV3Get**](OpenapiAPI.md#OpenapiV3Get) | **Get** /openapi/v3 | Retrieve OpenAPI Specification
[**OpenapiYamlGet**](OpenapiAPI.md#OpenapiYamlGet) | **Get** /openapi.yaml | Retrieve OpenAPI Specification



## OpenapiGet

> OpenapiGet(ctx).Execute()

Retrieve OpenAPI Specification

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
	r, err := apiClient.OpenapiAPI.OpenapiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenapiAPI.OpenapiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenapiJsonGet

> OpenapiJsonGet(ctx).Execute()

Retrieve OpenAPI Specification

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
	r, err := apiClient.OpenapiAPI.OpenapiJsonGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenapiAPI.OpenapiJsonGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiJsonGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenapiV3Get

> OpenapiV3Get(ctx).Execute()

Retrieve OpenAPI Specification

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
	r, err := apiClient.OpenapiAPI.OpenapiV3Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenapiAPI.OpenapiV3Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiV3GetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenapiYamlGet

> OpenapiYamlGet(ctx).Execute()

Retrieve OpenAPI Specification

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
	r, err := apiClient.OpenapiAPI.OpenapiYamlGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenapiAPI.OpenapiYamlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiYamlGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

