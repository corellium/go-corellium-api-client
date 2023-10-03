# \ProjectsApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AddProjectKey**](ProjectsApi.md#V1AddProjectKey) | **Post** /v1/projects/{projectId}/keys | Add Project Authorized Key
[**V1CreateProject**](ProjectsApi.md#V1CreateProject) | **Post** /v1/projects | Create a Project
[**V1DeleteProject**](ProjectsApi.md#V1DeleteProject) | **Delete** /v1/projects/{projectId} | Delete a Project
[**V1GetProject**](ProjectsApi.md#V1GetProject) | **Get** /v1/projects/{projectId} | Get a Project
[**V1GetProjectInstances**](ProjectsApi.md#V1GetProjectInstances) | **Get** /v1/projects/{projectId}/instances | Get Instances in Project
[**V1GetProjectKeys**](ProjectsApi.md#V1GetProjectKeys) | **Get** /v1/projects/{projectId}/keys | Get Project Authorized Keys
[**V1GetProjectVpnConfig**](ProjectsApi.md#V1GetProjectVpnConfig) | **Get** /v1/projects/{projectId}/vpnconfig/{format} | Get Project VPN Configuration
[**V1GetProjects**](ProjectsApi.md#V1GetProjects) | **Get** /v1/projects | Get Projects
[**V1RemoveProjectKey**](ProjectsApi.md#V1RemoveProjectKey) | **Delete** /v1/projects/{projectId}/keys/{keyId} | Delete Project Authorized Key
[**V1UpdateProject**](ProjectsApi.md#V1UpdateProject) | **Patch** /v1/projects/{projectId} | Update a Project
[**V1UpdateProjectSettings**](ProjectsApi.md#V1UpdateProjectSettings) | **Patch** /v1/projects/{projectId}/settings | Change Project Settings



## V1AddProjectKey

> ProjectKey V1AddProjectKey(ctx, projectId).ProjectKey(projectKey).Execute()

Add Project Authorized Key

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
    projectId := "projectId_example" // string | Project ID - uuid
    projectKey := *openapiclient.NewProjectKey("Kind_example", "Key_example") // ProjectKey | Key to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1AddProjectKey(context.Background(), projectId).ProjectKey(projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1AddProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AddProjectKey`: ProjectKey
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1AddProjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AddProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectKey** | [**ProjectKey**](ProjectKey.md) | Key to add | 

### Return type

[**ProjectKey**](ProjectKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateProject

> Project V1CreateProject(ctx).Project(project).Execute()

Create a Project

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
    project := *openapiclient.NewProject("Id_example") // Project | Project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1CreateProject(context.Background()).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) | Project | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteProject

> V1DeleteProject(ctx, projectId).Execute()

Delete a Project

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
    projectId := "projectId_example" // string | Project ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.V1DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteProjectRequest struct via the builder pattern


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


## V1GetProject

> Project V1GetProject(ctx, projectId).Execute()

Get a Project

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
    projectId := "projectId_example" // string | Project ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectInstances

> []Instance V1GetProjectInstances(ctx, projectId).Name(name).ReturnAttr(returnAttr).Execute()

Get Instances in Project

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
    projectId := "projectId_example" // string | Project ID - uuid
    name := "name_example" // string | Filter by project name (optional)
    returnAttr := []string{"Inner_example"} // []string | Attributes to include in instance return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1GetProjectInstances(context.Background(), projectId).Name(name).ReturnAttr(returnAttr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1GetProjectInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectInstances`: []Instance
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1GetProjectInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by project name | 
 **returnAttr** | **[]string** | Attributes to include in instance return | 

### Return type

[**[]Instance**](Instance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectKeys

> []ProjectKey V1GetProjectKeys(ctx, projectId).Execute()

Get Project Authorized Keys

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
    projectId := "projectId_example" // string | Project ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1GetProjectKeys(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1GetProjectKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectKeys`: []ProjectKey
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1GetProjectKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectKey**](ProjectKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectVpnConfig

> string V1GetProjectVpnConfig(ctx, projectId, format).Execute()

Get Project VPN Configuration

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
    projectId := "projectId_example" // string | Project ID - uuid
    format := "format_example" // string | VPN Config format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1GetProjectVpnConfig(context.Background(), projectId, format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1GetProjectVpnConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectVpnConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1GetProjectVpnConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**format** | **string** | VPN Config format | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectVpnConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-openvpn-profile, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjects

> []Project V1GetProjects(ctx).Name(name).IdsOnly(idsOnly).Execute()

Get Projects

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
    name := "name_example" // string | Filter by project name (optional)
    idsOnly := true // bool | Only include id of project in results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1GetProjects(context.Background()).Name(name).IdsOnly(idsOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjects`: []Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by project name | 
 **idsOnly** | **bool** | Only include id of project in results | 

### Return type

[**[]Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RemoveProjectKey

> V1RemoveProjectKey(ctx, projectId, keyId).Execute()

Delete Project Authorized Key

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
    projectId := "projectId_example" // string | Project ID - uuid
    keyId := "keyId_example" // string | Key ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.V1RemoveProjectKey(context.Background(), projectId, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1RemoveProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**keyId** | **string** | Key ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RemoveProjectKeyRequest struct via the builder pattern


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


## V1UpdateProject

> Project V1UpdateProject(ctx, projectId).Project(project).Execute()

Update a Project

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
    projectId := "projectId_example" // string | Project ID - uuid
    project := *openapiclient.NewProject("Id_example") // Project | Updated Project Settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.V1UpdateProject(context.Background(), projectId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.V1UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) | Updated Project Settings | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateProjectSettings

> V1UpdateProjectSettings(ctx, projectId).ProjectSettings(projectSettings).Execute()

Change Project Settings

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
    projectId := "projectId_example" // string | Project ID - uuid
    projectSettings := *openapiclient.NewProjectSettings() // ProjectSettings | New settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.V1UpdateProjectSettings(context.Background(), projectId).ProjectSettings(projectSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.V1UpdateProjectSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateProjectSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectSettings** | [**ProjectSettings**](ProjectSettings.md) | New settings | 

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

