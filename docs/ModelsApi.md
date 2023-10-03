# \ModelsApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetModelSoftware**](ModelsApi.md#V1GetModelSoftware) | **Get** /v1/models/{model}/software | Get Software for Model
[**V1GetModels**](ModelsApi.md#V1GetModels) | **Get** /v1/models | Get Supported Models



## V1GetModelSoftware

> []Firmware V1GetModelSoftware(ctx, model).Execute()

Get Software for Model

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
    model := "model_example" // string | Model to list available software for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.V1GetModelSoftware(context.Background(), model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.V1GetModelSoftware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetModelSoftware`: []Firmware
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.V1GetModelSoftware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model to list available software for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetModelSoftwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Firmware**](Firmware.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetModels

> []Model V1GetModels(ctx).Execute()

Get Supported Models

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
    resp, r, err := apiClient.ModelsApi.V1GetModels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.V1GetModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetModels`: []Model
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.V1GetModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetModelsRequest struct via the builder pattern


### Return type

[**[]Model**](Model.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

