# \ImagesApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateImage**](ImagesApi.md#V1CreateImage) | **Post** /v1/images | Create a new Image
[**V1DeleteImage**](ImagesApi.md#V1DeleteImage) | **Delete** /v2/images/{imageId} | Delete Image
[**V1GetImage**](ImagesApi.md#V1GetImage) | **Get** /v1/images/{imageId} | Get Image Metadata
[**V1GetImages**](ImagesApi.md#V1GetImages) | **Get** /v1/images | Get all Images Metadata
[**V1UploadImageData**](ImagesApi.md#V1UploadImageData) | **Post** /v1/images/{imageId} | Upload Image Data



## V1CreateImage

> Image V1CreateImage(ctx).Type_(type_).Encoding(encoding).Encapsulated(encapsulated).Name(name).Project(project).Instance(instance).File(file).Execute()

Create a new Image

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
    type_ := "type__example" // string | Image type
    encoding := "encoding_example" // string | How the file is stored
    encapsulated := true // bool | set to false if the uploaded file is not encapsulated inside an outer zipfile (optional)
    name := "name_example" // string | Image name (optional)
    project := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID (optional)
    instance := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Instance ID (optional)
    file := os.NewFile(1234, "some_file") // *os.File | Optionally the actual file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.V1CreateImage(context.Background()).Type_(type_).Encoding(encoding).Encapsulated(encapsulated).Name(name).Project(project).Instance(instance).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.V1CreateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.V1CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Image type | 
 **encoding** | **string** | How the file is stored | 
 **encapsulated** | **bool** | set to false if the uploaded file is not encapsulated inside an outer zipfile | 
 **name** | **string** | Image name | 
 **project** | **string** | Project ID | 
 **instance** | **string** | Instance ID | 
 **file** | ***os.File** | Optionally the actual file | 

### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteImage

> V1DeleteImage(ctx, imageId).Execute()

Delete Image

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
    imageId := "imageId_example" // string | Image ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesApi.V1DeleteImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.V1DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteImageRequest struct via the builder pattern


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


## V1GetImage

> Image V1GetImage(ctx, imageId).Execute()

Get Image Metadata

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
    imageId := "imageId_example" // string | Image ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.V1GetImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.V1GetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.V1GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetImages

> []Image V1GetImages(ctx).Project(project).Execute()

Get all Images Metadata

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
    project := "project_example" // string | Optionally filter by project - uuid (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.V1GetImages(context.Background()).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.V1GetImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetImages`: []Image
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.V1GetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Optionally filter by project - uuid | 

### Return type

[**[]Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UploadImageData

> Image V1UploadImageData(ctx, imageId).Body(body).Execute()

Upload Image Data



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
    imageId := "imageId_example" // string | Image ID - uuid
    body := "body_example" // string | Uploaded Image

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.V1UploadImageData(context.Background(), imageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.V1UploadImageData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UploadImageData`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.V1UploadImageData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UploadImageDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Uploaded Image | 

### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: binary
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

