# \HyperTraceApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1BtracePreauthorize**](HyperTraceApi.md#V1BtracePreauthorize) | **Get** /v1/instances/{instanceId}/btrace-authorize | Pre-authorize an btrace download
[**V1ClearHyperTrace**](HyperTraceApi.md#V1ClearHyperTrace) | **Delete** /v1/instances/{instanceId}/btrace | Clear HyperTrace logs
[**V1Kcrange**](HyperTraceApi.md#V1Kcrange) | **Get** /v1/instances/{instanceId}/btrace-kcrange | Get Kernel extension ranges
[**V1StartHyperTrace**](HyperTraceApi.md#V1StartHyperTrace) | **Post** /v1/instances/{instanceId}/btrace/enable | Start HyperTrace on an instance
[**V1StopHyperTrace**](HyperTraceApi.md#V1StopHyperTrace) | **Post** /v1/instances/{instanceId}/btrace/disable | Stop HyperTrace on an instance.



## V1BtracePreauthorize

> map[string]interface{} V1BtracePreauthorize(ctx, instanceId).Execute()

Pre-authorize an btrace download

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/aimoda/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperTraceApi.V1BtracePreauthorize(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperTraceApi.V1BtracePreauthorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BtracePreauthorize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HyperTraceApi.V1BtracePreauthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BtracePreauthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ClearHyperTrace

> V1ClearHyperTrace(ctx, instanceId).Execute()

Clear HyperTrace logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/aimoda/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HyperTraceApi.V1ClearHyperTrace(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperTraceApi.V1ClearHyperTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClearHyperTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Kcrange

> []Kcrange V1Kcrange(ctx, instanceId).Execute()

Get Kernel extension ranges

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/aimoda/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperTraceApi.V1Kcrange(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperTraceApi.V1Kcrange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1Kcrange`: []Kcrange
    fmt.Fprintf(os.Stdout, "Response from `HyperTraceApi.V1Kcrange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KcrangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Kcrange**](Kcrange.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StartHyperTrace

> V1StartHyperTrace(ctx, instanceId).BtraceEnableOptions(btraceEnableOptions).Execute()

Start HyperTrace on an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/aimoda/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    btraceEnableOptions := *openapiclient.NewBtraceEnableOptions() // BtraceEnableOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HyperTraceApi.V1StartHyperTrace(context.Background(), instanceId).BtraceEnableOptions(btraceEnableOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperTraceApi.V1StartHyperTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StartHyperTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **btraceEnableOptions** | [**BtraceEnableOptions**](BtraceEnableOptions.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StopHyperTrace

> V1StopHyperTrace(ctx, instanceId).Execute()

Stop HyperTrace on an instance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/aimoda/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HyperTraceApi.V1StopHyperTrace(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperTraceApi.V1StopHyperTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StopHyperTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

