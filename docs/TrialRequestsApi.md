# \TrialRequestsApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateSubscriberInvite**](TrialRequestsApi.md#V1CreateSubscriberInvite) | **Post** /v1/billing/invites | Create Subscriber Invite



## V1CreateSubscriberInvite

> SubscriberInvite V1CreateSubscriberInvite(ctx).SubscriberInvite(subscriberInvite).Execute()

Create Subscriber Invite



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
    subscriberInvite := *openapiclient.NewSubscriberInvite("CouponCode_example", "Aggregate_example", false, false, "CreatedAt_example", "UpdatedAt_example") // SubscriberInvite | Payload of Subscriber Invite

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrialRequestsApi.V1CreateSubscriberInvite(context.Background()).SubscriberInvite(subscriberInvite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialRequestsApi.V1CreateSubscriberInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateSubscriberInvite`: SubscriberInvite
    fmt.Fprintf(os.Stdout, "Response from `TrialRequestsApi.V1CreateSubscriberInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateSubscriberInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberInvite** | [**SubscriberInvite**](SubscriberInvite.md) | Payload of Subscriber Invite | 

### Return type

[**SubscriberInvite**](SubscriberInvite.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

