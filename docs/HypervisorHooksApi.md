# \HypervisorHooksApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClearHyperTraceHooks**](HypervisorHooksApi.md#V1ClearHyperTraceHooks) | **Post** /v1/instances/{instanceId}/hooks/clear | Clear Hooks on an instance
[**V1CreateHook**](HypervisorHooksApi.md#V1CreateHook) | **Post** /v1/instances/{instanceId}/hooks | Create hypervisor hook for Instance
[**V1DeleteHook**](HypervisorHooksApi.md#V1DeleteHook) | **Delete** /v1/hooks/{hookId} | Delete an existing hypervisor hook
[**V1ExecuteHyperTraceHooks**](HypervisorHooksApi.md#V1ExecuteHyperTraceHooks) | **Post** /v1/instances/{instanceId}/hooks/execute | Execute Hooks on an instance
[**V1GetHookById**](HypervisorHooksApi.md#V1GetHookById) | **Get** /v1/hooks/{hookId} | Get hypervisor hook by id
[**V1GetHooks**](HypervisorHooksApi.md#V1GetHooks) | **Get** /v1/instances/{instanceId}/hooks | Get all hypervisor hooks for Instance
[**V1UpdateHook**](HypervisorHooksApi.md#V1UpdateHook) | **Put** /v1/hooks/{hookId} | Update an existing hypervisor hook



## V1ClearHyperTraceHooks

> V1ClearHyperTraceHooks(ctx, instanceId).Execute()

Clear Hooks on an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorHooksApi.V1ClearHyperTraceHooks(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorHooksApi.V1ClearHyperTraceHooks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1ClearHyperTraceHooksRequest struct via the builder pattern


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


## V1CreateHook

> Hook V1CreateHook(ctx, instanceId).V1CreateHookParameters(v1CreateHookParameters).Execute()

Create hypervisor hook for Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    v1CreateHookParameters := *openapiclient.NewV1CreateHookParameters("Label_example", "Address_example", "Patch_example", "PatchType_example") // V1CreateHookParameters | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorHooksApi.V1CreateHook(context.Background(), instanceId).V1CreateHookParameters(v1CreateHookParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorHooksApi.V1CreateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `HypervisorHooksApi.V1CreateHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateHookParameters** | [**V1CreateHookParameters**](V1CreateHookParameters.md) | application/json | 

### Return type

[**Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteHook

> V1DeleteHook(ctx, hookId).Execute()

Delete an existing hypervisor hook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    hookId := "hookId_example" // string | Hook ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorHooksApi.V1DeleteHook(context.Background(), hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorHooksApi.V1DeleteHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Hook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteHookRequest struct via the builder pattern


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


## V1ExecuteHyperTraceHooks

> V1ExecuteHyperTraceHooks(ctx, instanceId).Execute()

Execute Hooks on an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorHooksApi.V1ExecuteHyperTraceHooks(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorHooksApi.V1ExecuteHyperTraceHooks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1ExecuteHyperTraceHooksRequest struct via the builder pattern


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


## V1GetHookById

> Hook V1GetHookById(ctx, hookId).Execute()

Get hypervisor hook by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    hookId := "hookId_example" // string | Hook Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorHooksApi.V1GetHookById(context.Background(), hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorHooksApi.V1GetHookById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetHookById`: Hook
    fmt.Fprintf(os.Stdout, "Response from `HypervisorHooksApi.V1GetHookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Hook Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetHookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetHooks

> []Hook V1GetHooks(ctx, instanceId).Limit(limit).Offset(offset).Sort(sort).Execute()

Get all hypervisor hooks for Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    limit := float32(8.14) // float32 | limit for pagination results, defaults to 20 (optional)
    offset := float32(8.14) // float32 | offset for pagination results, defaults to 0 (optional)
    sort := "sort_example" // string | sort ASC or DESC, defaults to DESC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorHooksApi.V1GetHooks(context.Background(), instanceId).Limit(limit).Offset(offset).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorHooksApi.V1GetHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetHooks`: []Hook
    fmt.Fprintf(os.Stdout, "Response from `HypervisorHooksApi.V1GetHooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** | limit for pagination results, defaults to 20 | 
 **offset** | **float32** | offset for pagination results, defaults to 0 | 
 **sort** | **string** | sort ASC or DESC, defaults to DESC | 

### Return type

[**[]Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateHook

> Hook V1UpdateHook(ctx, hookId).V1CreateHookParameters(v1CreateHookParameters).Execute()

Update an existing hypervisor hook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    hookId := "hookId_example" // string | Hook ID
    v1CreateHookParameters := *openapiclient.NewV1CreateHookParameters("Label_example", "Address_example", "Patch_example", "PatchType_example") // V1CreateHookParameters | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorHooksApi.V1UpdateHook(context.Background(), hookId).V1CreateHookParameters(v1CreateHookParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorHooksApi.V1UpdateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `HypervisorHooksApi.V1UpdateHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Hook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateHookParameters** | [**V1CreateHookParameters**](V1CreateHookParameters.md) | application/json | 

### Return type

[**Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

