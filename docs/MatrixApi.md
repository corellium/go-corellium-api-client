# \MatrixApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssessment**](MatrixApi.md#CreateAssessment) | **Post** /v1/services/matrix/{instanceId}/assessments | Create assessment
[**DeleteAssessment**](MatrixApi.md#DeleteAssessment) | **Delete** /v1/services/matrix/{instanceId}/assessments/{assessmentId} | Delete assessment
[**DownloadReport**](MatrixApi.md#DownloadReport) | **Get** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/download | Download report
[**GetAssessmentById**](MatrixApi.md#GetAssessmentById) | **Get** /v1/services/matrix/{instanceId}/assessments/{assessmentId} | Get assessment by ID
[**GetAssessmentsByInstanceId**](MatrixApi.md#GetAssessmentsByInstanceId) | **Get** /v1/services/matrix/{instanceId}/instances/{instanceId}/assessments | Get assessments by instanceId
[**RunTests**](MatrixApi.md#RunTests) | **Post** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/test | Update assessment state and execute MATRIX tests
[**StartMonitoring**](MatrixApi.md#StartMonitoring) | **Post** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/start | Update assessment state and begin device monitoring
[**StopMonitoring**](MatrixApi.md#StopMonitoring) | **Post** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/stop | Update assessment state and stop device monitoring



## CreateAssessment

> AssessmentIdStatus CreateAssessment(ctx, instanceId).CreateAssessmentDto(createAssessmentDto).Execute()

Create assessment

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    createAssessmentDto := *openapiclient.NewCreateAssessmentDto("InstanceId_example", "com.android.egg") // CreateAssessmentDto | Create a new assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixApi.CreateAssessment(context.Background(), instanceId).CreateAssessmentDto(createAssessmentDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.CreateAssessment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssessment`: AssessmentIdStatus
    fmt.Fprintf(os.Stdout, "Response from `MatrixApi.CreateAssessment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssessmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAssessmentDto** | [**CreateAssessmentDto**](CreateAssessmentDto.md) | Create a new assessment | 

### Return type

[**AssessmentIdStatus**](AssessmentIdStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssessment

> DeleteAssessment(ctx, instanceId, assessmentId).Execute()

Delete assessment

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MatrixApi.DeleteAssessment(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.DeleteAssessment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssessmentRequest struct via the builder pattern


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


## DownloadReport

> string DownloadReport(ctx, instanceId, assessmentId).Format(format).Execute()

Download report

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment
    format := "format_example" // string | Assessment report download format (default to "html")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixApi.DownloadReport(context.Background(), instanceId, assessmentId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.DownloadReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadReport`: string
    fmt.Fprintf(os.Stdout, "Response from `MatrixApi.DownloadReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Assessment report download format | [default to &quot;html&quot;]

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssessmentById

> Assessment GetAssessmentById(ctx, instanceId, assessmentId).Execute()

Get assessment by ID

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixApi.GetAssessmentById(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.GetAssessmentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssessmentById`: Assessment
    fmt.Fprintf(os.Stdout, "Response from `MatrixApi.GetAssessmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssessmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Assessment**](Assessment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssessmentsByInstanceId

> []Assessment GetAssessmentsByInstanceId(ctx, instanceId).Execute()

Get assessments by instanceId

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixApi.GetAssessmentsByInstanceId(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.GetAssessmentsByInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssessmentsByInstanceId`: []Assessment
    fmt.Fprintf(os.Stdout, "Response from `MatrixApi.GetAssessmentsByInstanceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssessmentsByInstanceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Assessment**](Assessment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunTests

> AssessmentIdStatus RunTests(ctx, instanceId, assessmentId).TestAssessmentDto(testAssessmentDto).Execute()

Update assessment state and execute MATRIX tests

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment
    testAssessmentDto := *openapiclient.NewTestAssessmentDto() // TestAssessmentDto | Execute MATRIX tests (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixApi.RunTests(context.Background(), instanceId, assessmentId).TestAssessmentDto(testAssessmentDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.RunTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunTests`: AssessmentIdStatus
    fmt.Fprintf(os.Stdout, "Response from `MatrixApi.RunTests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testAssessmentDto** | [**TestAssessmentDto**](TestAssessmentDto.md) | Execute MATRIX tests | 

### Return type

[**AssessmentIdStatus**](AssessmentIdStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMonitoring

> StartMonitoring(ctx, instanceId, assessmentId).Execute()

Update assessment state and begin device monitoring

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MatrixApi.StartMonitoring(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.StartMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartMonitoringRequest struct via the builder pattern


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


## StopMonitoring

> StopMonitoring(ctx, instanceId, assessmentId).Execute()

Update assessment state and stop device monitoring

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
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MatrixApi.StopMonitoring(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixApi.StopMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopMonitoringRequest struct via the builder pattern


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

