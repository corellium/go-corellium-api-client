# \DomainAuthProviderApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateDomainAuthProvider**](DomainAuthProviderApi.md#V1CreateDomainAuthProvider) | **Post** /v1/domain/{domainId}/auth | Create a new auth provider for a domain
[**V1DeleteDomainAuthProvider**](DomainAuthProviderApi.md#V1DeleteDomainAuthProvider) | **Delete** /v1/domain/{domainId}/auth/{providerId} | Delete an auth provider from a domain
[**V1GetDomainAuthProviders**](DomainAuthProviderApi.md#V1GetDomainAuthProviders) | **Get** /v1/domain/{domainId}/auth | Return all configured auth providers for a domain (including globally configured providers)
[**V1UpdateDomainAuthProvider**](DomainAuthProviderApi.md#V1UpdateDomainAuthProvider) | **Put** /v1/domain/{domainId}/auth/{providerId} | Update an auth provider for a domain



## V1CreateDomainAuthProvider

> DomainAuthProviderResponse V1CreateDomainAuthProvider(ctx, domainId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()

Create a new auth provider for a domain

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
    domainId := "domainId_example" // string | Domain ID - uuid
    domainAuthProviderRequest := *openapiclient.NewDomainAuthProviderRequest("ProviderType_example", false) // DomainAuthProviderRequest | Request Data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAuthProviderApi.V1CreateDomainAuthProvider(context.Background(), domainId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAuthProviderApi.V1CreateDomainAuthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateDomainAuthProvider`: DomainAuthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAuthProviderApi.V1CreateDomainAuthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateDomainAuthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainAuthProviderRequest** | [**DomainAuthProviderRequest**](DomainAuthProviderRequest.md) | Request Data | 

### Return type

[**DomainAuthProviderResponse**](DomainAuthProviderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteDomainAuthProvider

> map[string]interface{} V1DeleteDomainAuthProvider(ctx, domainId, providerId).Execute()

Delete an auth provider from a domain

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
    domainId := "domainId_example" // string | Domain ID - uuid
    providerId := "providerId_example" // string | Provider ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAuthProviderApi.V1DeleteDomainAuthProvider(context.Background(), domainId, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAuthProviderApi.V1DeleteDomainAuthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DeleteDomainAuthProvider`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DomainAuthProviderApi.V1DeleteDomainAuthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 
**providerId** | **string** | Provider ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteDomainAuthProviderRequest struct via the builder pattern


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


## V1GetDomainAuthProviders

> []DomainAuthProviderResponse V1GetDomainAuthProviders(ctx, domainId).Execute()

Return all configured auth providers for a domain (including globally configured providers)

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
    domainId := "domainId_example" // string | Domain ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAuthProviderApi.V1GetDomainAuthProviders(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAuthProviderApi.V1GetDomainAuthProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetDomainAuthProviders`: []DomainAuthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAuthProviderApi.V1GetDomainAuthProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetDomainAuthProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainAuthProviderResponse**](DomainAuthProviderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateDomainAuthProvider

> DomainAuthProviderResponse V1UpdateDomainAuthProvider(ctx, domainId, providerId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()

Update an auth provider for a domain

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
    domainId := "domainId_example" // string | Domain ID - uuid
    providerId := "providerId_example" // string | Provider ID - uuid
    domainAuthProviderRequest := *openapiclient.NewDomainAuthProviderRequest("ProviderType_example", false) // DomainAuthProviderRequest | Request Data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAuthProviderApi.V1UpdateDomainAuthProvider(context.Background(), domainId, providerId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAuthProviderApi.V1UpdateDomainAuthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateDomainAuthProvider`: DomainAuthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAuthProviderApi.V1UpdateDomainAuthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 
**providerId** | **string** | Provider ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateDomainAuthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domainAuthProviderRequest** | [**DomainAuthProviderRequest**](DomainAuthProviderRequest.md) | Request Data | 

### Return type

[**DomainAuthProviderResponse**](DomainAuthProviderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

