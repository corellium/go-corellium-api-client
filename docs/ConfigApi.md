# \ConfigApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetConfig**](ConfigApi.md#V1GetConfig) | **Get** /v1/config | Get all configs



## V1GetConfig

> ConfigResponse V1GetConfig(ctx).Execute()

Get all configs

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.V1GetConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.V1GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetConfig`: ConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.V1GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetConfigRequest struct via the builder pattern


### Return type

[**ConfigResponse**](ConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

