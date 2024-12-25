# \DomainApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CheckSubdomainExistence**](DomainApi.md#V1CheckSubdomainExistence) | **Post** /v1/domain/check | Check the existence of a subdomain



## V1CheckSubdomainExistence

> CheckSubdomainResponse V1CheckSubdomainExistence(ctx).V1CheckSubdomainExistenceParameters(v1CheckSubdomainExistenceParameters).Execute()

Check the existence of a subdomain

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
    v1CheckSubdomainExistenceParameters := *openapiclient.NewV1CheckSubdomainExistenceParameters("Domain_example") // V1CheckSubdomainExistenceParameters | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainApi.V1CheckSubdomainExistence(context.Background()).V1CheckSubdomainExistenceParameters(v1CheckSubdomainExistenceParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.V1CheckSubdomainExistence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CheckSubdomainExistence`: CheckSubdomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.V1CheckSubdomainExistence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CheckSubdomainExistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CheckSubdomainExistenceParameters** | [**V1CheckSubdomainExistenceParameters**](V1CheckSubdomainExistenceParameters.md) | application/json | 

### Return type

[**CheckSubdomainResponse**](CheckSubdomainResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

