# \NetConnectApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateNetworkConnection**](NetConnectApi.md#V1CreateNetworkConnection) | **Post** /v1/network/connections | Create a new Network Connection
[**V1DeleteNetworkConnection**](NetConnectApi.md#V1DeleteNetworkConnection) | **Delete** /v1/network/connections/{id} | Delete an existing Network Connection
[**V1ListNetworkConnections**](NetConnectApi.md#V1ListNetworkConnections) | **Get** /v1/network/connections | List available network connections
[**V1ListNetworkProviders**](NetConnectApi.md#V1ListNetworkProviders) | **Get** /v1/network/providers | List available network providers
[**V1PartialUpdateNetworkConnection**](NetConnectApi.md#V1PartialUpdateNetworkConnection) | **Patch** /v1/network/connections/{id} | Update Network Connection (partial)
[**V1UpdateNetworkConnection**](NetConnectApi.md#V1UpdateNetworkConnection) | **Put** /v1/network/connections/{id} | Update Network Connection



## V1CreateNetworkConnection

> V1CreateNetworkConnection(ctx).CreateNetworkConnectionOptions(createNetworkConnectionOptions).Execute()

Create a new Network Connection



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
    createNetworkConnectionOptions := *openapiclient.NewCreateNetworkConnectionOptions("Identifier_example", "Name_example", "Provider_example") // CreateNetworkConnectionOptions | Network Connection Options

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetConnectApi.V1CreateNetworkConnection(context.Background()).CreateNetworkConnectionOptions(createNetworkConnectionOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetConnectApi.V1CreateNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkConnectionOptions** | [**CreateNetworkConnectionOptions**](CreateNetworkConnectionOptions.md) | Network Connection Options | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteNetworkConnection

> V1DeleteNetworkConnection(ctx, id).Execute()

Delete an existing Network Connection



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
    id := "id_example" // string | Network Connection Identifier - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetConnectApi.V1DeleteNetworkConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetConnectApi.V1DeleteNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network Connection Identifier - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteNetworkConnectionRequest struct via the builder pattern


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


## V1ListNetworkConnections

> NetworkConnectionOffsetPaginationResult V1ListNetworkConnections(ctx).Execute()

List available network connections



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetConnectApi.V1ListNetworkConnections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetConnectApi.V1ListNetworkConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ListNetworkConnections`: NetworkConnectionOffsetPaginationResult
    fmt.Fprintf(os.Stdout, "Response from `NetConnectApi.V1ListNetworkConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListNetworkConnectionsRequest struct via the builder pattern


### Return type

[**NetworkConnectionOffsetPaginationResult**](NetworkConnectionOffsetPaginationResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListNetworkProviders

> NetworkConnectionProviderOffsetPaginationResult V1ListNetworkProviders(ctx).Execute()

List available network providers



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetConnectApi.V1ListNetworkProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetConnectApi.V1ListNetworkProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ListNetworkProviders`: NetworkConnectionProviderOffsetPaginationResult
    fmt.Fprintf(os.Stdout, "Response from `NetConnectApi.V1ListNetworkProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListNetworkProvidersRequest struct via the builder pattern


### Return type

[**NetworkConnectionProviderOffsetPaginationResult**](NetworkConnectionProviderOffsetPaginationResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PartialUpdateNetworkConnection

> V1PartialUpdateNetworkConnection(ctx, id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Execute()

Update Network Connection (partial)



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
    id := "id_example" // string | Network Connection Identifier - uuid
    updateNetworkConnectionOptions := *openapiclient.NewUpdateNetworkConnectionOptions("Name_example") // UpdateNetworkConnectionOptions | Network Connection Options

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetConnectApi.V1PartialUpdateNetworkConnection(context.Background(), id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetConnectApi.V1PartialUpdateNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network Connection Identifier - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PartialUpdateNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkConnectionOptions** | [**UpdateNetworkConnectionOptions**](UpdateNetworkConnectionOptions.md) | Network Connection Options | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateNetworkConnection

> V1UpdateNetworkConnection(ctx, id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Execute()

Update Network Connection



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
    id := "id_example" // string | Network Connection Identifier - uuid
    updateNetworkConnectionOptions := *openapiclient.NewUpdateNetworkConnectionOptions("Name_example") // UpdateNetworkConnectionOptions | Network Connection Options

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetConnectApi.V1UpdateNetworkConnection(context.Background(), id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetConnectApi.V1UpdateNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network Connection Identifier - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkConnectionOptions** | [**UpdateNetworkConnectionOptions**](UpdateNetworkConnectionOptions.md) | Network Connection Options | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

