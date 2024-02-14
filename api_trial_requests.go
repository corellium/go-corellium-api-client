/*
Corellium API

REST API to manage your virtual devices.

API version: 6.0.0-20323
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// TrialRequestsApiService TrialRequestsApi service
type TrialRequestsApiService service

type TrialRequestsApiV1CreateSubscriberInviteRequest struct {
	ctx context.Context
	ApiService *TrialRequestsApiService
	subscriberInvite *SubscriberInvite
}

// Payload of Subscriber Invite
func (r TrialRequestsApiV1CreateSubscriberInviteRequest) SubscriberInvite(subscriberInvite SubscriberInvite) TrialRequestsApiV1CreateSubscriberInviteRequest {
	r.subscriberInvite = &subscriberInvite
	return r
}

func (r TrialRequestsApiV1CreateSubscriberInviteRequest) Execute() (*SubscriberInvite, *http.Response, error) {
	return r.ApiService.V1CreateSubscriberInviteExecute(r)
}

/*
V1CreateSubscriberInvite Create Subscriber Invite

Create Subscriber Invite

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TrialRequestsApiV1CreateSubscriberInviteRequest
*/
func (a *TrialRequestsApiService) V1CreateSubscriberInvite(ctx context.Context) TrialRequestsApiV1CreateSubscriberInviteRequest {
	return TrialRequestsApiV1CreateSubscriberInviteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscriberInvite
func (a *TrialRequestsApiService) V1CreateSubscriberInviteExecute(r TrialRequestsApiV1CreateSubscriberInviteRequest) (*SubscriberInvite, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriberInvite
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrialRequestsApiService.V1CreateSubscriberInvite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/billing/invites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriberInvite == nil {
		return localVarReturnValue, nil, reportError("subscriberInvite is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.subscriberInvite
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
