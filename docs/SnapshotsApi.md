# \SnapshotsApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateSnapshot**](SnapshotsApi.md#V1CreateSnapshot) | **Post** /v1/instances/{instanceId}/snapshots | Create Instance Snapshot
[**V1DeleteInstanceSnapshot**](SnapshotsApi.md#V1DeleteInstanceSnapshot) | **Delete** /v1/instances/{instanceId}/snapshots/{snapshotId} | Delete a Snapshot
[**V1DeleteSnapshot**](SnapshotsApi.md#V1DeleteSnapshot) | **Delete** /v1/snapshots/{snapshotId} | Delete a Snapshot
[**V1DeleteSnapshotPermissions**](SnapshotsApi.md#V1DeleteSnapshotPermissions) | **Delete** /v1/snapshots/{snapshotId}/permissions | Remove a user from the list of users who have access to the snapshot
[**V1GetInstanceSnapshot**](SnapshotsApi.md#V1GetInstanceSnapshot) | **Get** /v1/instances/{instanceId}/snapshots/{snapshotId} | Get Instance Snapshot
[**V1GetInstanceSnapshots**](SnapshotsApi.md#V1GetInstanceSnapshots) | **Get** /v1/instances/{instanceId}/snapshots | Get Instance Snapshots
[**V1GetSharedSnapshots**](SnapshotsApi.md#V1GetSharedSnapshots) | **Get** /v1/snapshots | Fetch snapshots shared with and shared by the requesting user
[**V1GetSnapshot**](SnapshotsApi.md#V1GetSnapshot) | **Get** /v1/snapshots/{snapshotId} | Get Snapshot
[**V1RenameInstanceSnapshot**](SnapshotsApi.md#V1RenameInstanceSnapshot) | **Patch** /v1/instances/{instanceId}/snapshots/{snapshotId} | Rename a Snapshot
[**V1RestoreInstanceSnapshot**](SnapshotsApi.md#V1RestoreInstanceSnapshot) | **Post** /v1/instances/{instanceId}/snapshots/{snapshotId}/restore | Restore a Snapshot
[**V1SetSnapshotPermissions**](SnapshotsApi.md#V1SetSnapshotPermissions) | **Post** /v1/snapshots/{snapshotId}/permissions | Add a user to the list of users who have access to the snapshot
[**V1SnapshotRename**](SnapshotsApi.md#V1SnapshotRename) | **Patch** /v1/snapshots/{snapshotId} | Rename a Snapshot



## V1CreateSnapshot

> Snapshot V1CreateSnapshot(ctx, instanceId).SnapshotCreationOptions(snapshotCreationOptions).Execute()

Create Instance Snapshot

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
    snapshotCreationOptions := *openapiclient.NewSnapshotCreationOptions("Name_example") // SnapshotCreationOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.V1CreateSnapshot(context.Background(), instanceId).SnapshotCreationOptions(snapshotCreationOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1CreateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotCreationOptions** | [**SnapshotCreationOptions**](SnapshotCreationOptions.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteInstanceSnapshot

> V1DeleteInstanceSnapshot(ctx, instanceId, snapshotId).Execute()

Delete a Snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SnapshotsApi.V1DeleteInstanceSnapshot(context.Background(), instanceId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1DeleteInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteInstanceSnapshotRequest struct via the builder pattern


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


## V1DeleteSnapshot

> V1DeleteSnapshot(ctx, snapshotId).Execute()

Delete a Snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SnapshotsApi.V1DeleteSnapshot(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1DeleteSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteSnapshotRequest struct via the builder pattern


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


## V1DeleteSnapshotPermissions

> Snapshot V1DeleteSnapshotPermissions(ctx, snapshotId).DeleteSnapshotPermissionsRequestPayload(deleteSnapshotPermissionsRequestPayload).Execute()

Remove a user from the list of users who have access to the snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    deleteSnapshotPermissionsRequestPayload := *openapiclient.NewDeleteSnapshotPermissionsRequestPayload() // DeleteSnapshotPermissionsRequestPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.V1DeleteSnapshotPermissions(context.Background(), snapshotId).DeleteSnapshotPermissionsRequestPayload(deleteSnapshotPermissionsRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1DeleteSnapshotPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DeleteSnapshotPermissions`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1DeleteSnapshotPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteSnapshotPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteSnapshotPermissionsRequestPayload** | [**DeleteSnapshotPermissionsRequestPayload**](DeleteSnapshotPermissionsRequestPayload.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceSnapshot

> Snapshot V1GetInstanceSnapshot(ctx, instanceId, snapshotId).Execute()

Get Instance Snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.V1GetInstanceSnapshot(context.Background(), instanceId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1GetInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1GetInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceSnapshots

> []Snapshot V1GetInstanceSnapshots(ctx, instanceId).Execute()

Get Instance Snapshots

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
    resp, r, err := apiClient.SnapshotsApi.V1GetInstanceSnapshots(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1GetInstanceSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceSnapshots`: []Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1GetInstanceSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetSharedSnapshots

> Snapshot V1GetSharedSnapshots(ctx).Execute()

Fetch snapshots shared with and shared by the requesting user

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
    resp, r, err := apiClient.SnapshotsApi.V1GetSharedSnapshots(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1GetSharedSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetSharedSnapshots`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1GetSharedSnapshots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetSharedSnapshotsRequest struct via the builder pattern


### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetSnapshot

> Snapshot V1GetSnapshot(ctx, snapshotId).Execute()

Get Snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.V1GetSnapshot(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1GetSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RenameInstanceSnapshot

> Snapshot V1RenameInstanceSnapshot(ctx, instanceId, snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()

Rename a Snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    snapshotCreationOptions := *openapiclient.NewSnapshotCreationOptions("Name_example") // SnapshotCreationOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.V1RenameInstanceSnapshot(context.Background(), instanceId, snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1RenameInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1RenameInstanceSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1RenameInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RenameInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snapshotCreationOptions** | [**SnapshotCreationOptions**](SnapshotCreationOptions.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RestoreInstanceSnapshot

> V1RestoreInstanceSnapshot(ctx, instanceId, snapshotId).Execute()

Restore a Snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SnapshotsApi.V1RestoreInstanceSnapshot(context.Background(), instanceId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1RestoreInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RestoreInstanceSnapshotRequest struct via the builder pattern


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


## V1SetSnapshotPermissions

> Snapshot V1SetSnapshotPermissions(ctx, snapshotId).PostSnapshotPermissionsRequestPayload(postSnapshotPermissionsRequestPayload).Execute()

Add a user to the list of users who have access to the snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    postSnapshotPermissionsRequestPayload := *openapiclient.NewPostSnapshotPermissionsRequestPayload("Email_example") // PostSnapshotPermissionsRequestPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.V1SetSnapshotPermissions(context.Background(), snapshotId).PostSnapshotPermissionsRequestPayload(postSnapshotPermissionsRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1SetSnapshotPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SetSnapshotPermissions`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1SetSnapshotPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SetSnapshotPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postSnapshotPermissionsRequestPayload** | [**PostSnapshotPermissionsRequestPayload**](PostSnapshotPermissionsRequestPayload.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SnapshotRename

> Snapshot V1SnapshotRename(ctx, snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()

Rename a Snapshot

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
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    snapshotCreationOptions := *openapiclient.NewSnapshotCreationOptions("Name_example") // SnapshotCreationOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.V1SnapshotRename(context.Background(), snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.V1SnapshotRename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SnapshotRename`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.V1SnapshotRename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SnapshotRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotCreationOptions** | [**SnapshotCreationOptions**](SnapshotCreationOptions.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

