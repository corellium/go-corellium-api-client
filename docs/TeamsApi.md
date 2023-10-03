# \TeamsApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AddUserToTeam**](TeamsApi.md#V1AddUserToTeam) | **Put** /v1/teams/{teamId}/users/{userId} | Add user to team
[**V1RemoveUserFromTeam**](TeamsApi.md#V1RemoveUserFromTeam) | **Delete** /v1/teams/{teamId}/users/{userId} | Remove user from team
[**V1TeamChange**](TeamsApi.md#V1TeamChange) | **Patch** /v1/teams/{teamId} | Update team
[**V1TeamCreate**](TeamsApi.md#V1TeamCreate) | **Post** /v1/teams | Create team
[**V1TeamDelete**](TeamsApi.md#V1TeamDelete) | **Delete** /v1/teams/{teamId} | Delete team
[**V1Teams**](TeamsApi.md#V1Teams) | **Get** /v1/teams | Get teams



## V1AddUserToTeam

> V1AddUserToTeam(ctx, teamId, userId).Execute()

Add user to team



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
    teamId := "teamId_example" // string | Team ID - uuid
    userId := "userId_example" // string | User ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TeamsApi.V1AddUserToTeam(context.Background(), teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.V1AddUserToTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 
**userId** | **string** | User ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AddUserToTeamRequest struct via the builder pattern


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


## V1RemoveUserFromTeam

> V1RemoveUserFromTeam(ctx, teamId, userId).Execute()

Remove user from team



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
    teamId := "teamId_example" // string | Team ID - uuid
    userId := "userId_example" // string | User ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TeamsApi.V1RemoveUserFromTeam(context.Background(), teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.V1RemoveUserFromTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 
**userId** | **string** | User ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RemoveUserFromTeamRequest struct via the builder pattern


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


## V1TeamChange

> V1TeamChange(ctx, teamId).CreateTeam(createTeam).Execute()

Update team



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
    teamId := "teamId_example" // string | Team ID - uuid
    createTeam := *openapiclient.NewCreateTeam("Name_example") // CreateTeam | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TeamsApi.V1TeamChange(context.Background(), teamId).CreateTeam(createTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.V1TeamChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTeam** | [**CreateTeam**](CreateTeam.md) |  | 

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


## V1TeamCreate

> TeamCreate V1TeamCreate(ctx).CreateTeam(createTeam).Execute()

Create team



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
    createTeam := *openapiclient.NewCreateTeam("Name_example") // CreateTeam | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.V1TeamCreate(context.Background()).CreateTeam(createTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.V1TeamCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TeamCreate`: TeamCreate
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.V1TeamCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeam** | [**CreateTeam**](CreateTeam.md) |  | 

### Return type

[**TeamCreate**](TeamCreate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TeamDelete

> V1TeamDelete(ctx, teamId).Execute()

Delete team



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
    teamId := "teamId_example" // string | Team ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TeamsApi.V1TeamDelete(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.V1TeamDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamDeleteRequest struct via the builder pattern


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


## V1Teams

> []Team V1Teams(ctx).Execute()

Get teams



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
    resp, r, err := apiClient.TeamsApi.V1Teams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.V1Teams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1Teams`: []Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.V1Teams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamsRequest struct via the builder pattern


### Return type

[**[]Team**](Team.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

