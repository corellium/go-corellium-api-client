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
	"strings"
)


// SnapshotsApiService SnapshotsApi service
type SnapshotsApiService service

type SnapshotsApiV1CreateSnapshotRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	instanceId string
	snapshotCreationOptions *SnapshotCreationOptions
}

// 
func (r SnapshotsApiV1CreateSnapshotRequest) SnapshotCreationOptions(snapshotCreationOptions SnapshotCreationOptions) SnapshotsApiV1CreateSnapshotRequest {
	r.snapshotCreationOptions = &snapshotCreationOptions
	return r
}

func (r SnapshotsApiV1CreateSnapshotRequest) Execute() (*Snapshot, *http.Response, error) {
	return r.ApiService.V1CreateSnapshotExecute(r)
}

/*
V1CreateSnapshot Create Instance Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId Instance ID - uuid
 @return SnapshotsApiV1CreateSnapshotRequest
*/
func (a *SnapshotsApiService) V1CreateSnapshot(ctx context.Context, instanceId string) SnapshotsApiV1CreateSnapshotRequest {
	return SnapshotsApiV1CreateSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return Snapshot
func (a *SnapshotsApiService) V1CreateSnapshotExecute(r SnapshotsApiV1CreateSnapshotRequest) (*Snapshot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1CreateSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotCreationOptions == nil {
		return localVarReturnValue, nil, reportError("snapshotCreationOptions is required and must be specified")
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
	localVarPostBody = r.snapshotCreationOptions
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SnapshotsApiV1DeleteInstanceSnapshotRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	instanceId string
	snapshotId string
}

func (r SnapshotsApiV1DeleteInstanceSnapshotRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1DeleteInstanceSnapshotExecute(r)
}

/*
V1DeleteInstanceSnapshot Delete a Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId Instance ID - uuid
 @param snapshotId Snapshot ID - uuid
 @return SnapshotsApiV1DeleteInstanceSnapshotRequest
*/
func (a *SnapshotsApiService) V1DeleteInstanceSnapshot(ctx context.Context, instanceId string, snapshotId string) SnapshotsApiV1DeleteInstanceSnapshotRequest {
	return SnapshotsApiV1DeleteInstanceSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
func (a *SnapshotsApiService) V1DeleteInstanceSnapshotExecute(r SnapshotsApiV1DeleteInstanceSnapshotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1DeleteInstanceSnapshot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SnapshotsApiV1DeleteSnapshotRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	snapshotId string
}

func (r SnapshotsApiV1DeleteSnapshotRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1DeleteSnapshotExecute(r)
}

/*
V1DeleteSnapshot Delete a Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param snapshotId Snapshot ID - uuid
 @return SnapshotsApiV1DeleteSnapshotRequest
*/
func (a *SnapshotsApiService) V1DeleteSnapshot(ctx context.Context, snapshotId string) SnapshotsApiV1DeleteSnapshotRequest {
	return SnapshotsApiV1DeleteSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
func (a *SnapshotsApiService) V1DeleteSnapshotExecute(r SnapshotsApiV1DeleteSnapshotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1DeleteSnapshot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SnapshotsApiV1GetInstanceSnapshotRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	instanceId string
	snapshotId string
}

func (r SnapshotsApiV1GetInstanceSnapshotRequest) Execute() (*Snapshot, *http.Response, error) {
	return r.ApiService.V1GetInstanceSnapshotExecute(r)
}

/*
V1GetInstanceSnapshot Get Instance Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId Instance ID - uuid
 @param snapshotId Snapshot ID - uuid
 @return SnapshotsApiV1GetInstanceSnapshotRequest
*/
func (a *SnapshotsApiService) V1GetInstanceSnapshot(ctx context.Context, instanceId string, snapshotId string) SnapshotsApiV1GetInstanceSnapshotRequest {
	return SnapshotsApiV1GetInstanceSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
//  @return Snapshot
func (a *SnapshotsApiService) V1GetInstanceSnapshotExecute(r SnapshotsApiV1GetInstanceSnapshotRequest) (*Snapshot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1GetInstanceSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SnapshotsApiV1GetInstanceSnapshotsRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	instanceId string
}

func (r SnapshotsApiV1GetInstanceSnapshotsRequest) Execute() ([]Snapshot, *http.Response, error) {
	return r.ApiService.V1GetInstanceSnapshotsExecute(r)
}

/*
V1GetInstanceSnapshots Get Instance Snapshots

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId Instance ID - uuid
 @return SnapshotsApiV1GetInstanceSnapshotsRequest
*/
func (a *SnapshotsApiService) V1GetInstanceSnapshots(ctx context.Context, instanceId string) SnapshotsApiV1GetInstanceSnapshotsRequest {
	return SnapshotsApiV1GetInstanceSnapshotsRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return []Snapshot
func (a *SnapshotsApiService) V1GetInstanceSnapshotsExecute(r SnapshotsApiV1GetInstanceSnapshotsRequest) ([]Snapshot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1GetInstanceSnapshots")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SnapshotsApiV1GetSnapshotRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	snapshotId string
}

func (r SnapshotsApiV1GetSnapshotRequest) Execute() (*Snapshot, *http.Response, error) {
	return r.ApiService.V1GetSnapshotExecute(r)
}

/*
V1GetSnapshot Get Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param snapshotId Snapshot ID - uuid
 @return SnapshotsApiV1GetSnapshotRequest
*/
func (a *SnapshotsApiService) V1GetSnapshot(ctx context.Context, snapshotId string) SnapshotsApiV1GetSnapshotRequest {
	return SnapshotsApiV1GetSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
//  @return Snapshot
func (a *SnapshotsApiService) V1GetSnapshotExecute(r SnapshotsApiV1GetSnapshotRequest) (*Snapshot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1GetSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SnapshotsApiV1RenameInstanceSnapshotRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	instanceId string
	snapshotId string
	snapshotCreationOptions *SnapshotCreationOptions
}

// 
func (r SnapshotsApiV1RenameInstanceSnapshotRequest) SnapshotCreationOptions(snapshotCreationOptions SnapshotCreationOptions) SnapshotsApiV1RenameInstanceSnapshotRequest {
	r.snapshotCreationOptions = &snapshotCreationOptions
	return r
}

func (r SnapshotsApiV1RenameInstanceSnapshotRequest) Execute() (*Snapshot, *http.Response, error) {
	return r.ApiService.V1RenameInstanceSnapshotExecute(r)
}

/*
V1RenameInstanceSnapshot Rename a Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId Instance ID - uuid
 @param snapshotId Snapshot ID - uuid
 @return SnapshotsApiV1RenameInstanceSnapshotRequest
*/
func (a *SnapshotsApiService) V1RenameInstanceSnapshot(ctx context.Context, instanceId string, snapshotId string) SnapshotsApiV1RenameInstanceSnapshotRequest {
	return SnapshotsApiV1RenameInstanceSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
//  @return Snapshot
func (a *SnapshotsApiService) V1RenameInstanceSnapshotExecute(r SnapshotsApiV1RenameInstanceSnapshotRequest) (*Snapshot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1RenameInstanceSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotCreationOptions == nil {
		return localVarReturnValue, nil, reportError("snapshotCreationOptions is required and must be specified")
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
	localVarPostBody = r.snapshotCreationOptions
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type SnapshotsApiV1RestoreInstanceSnapshotRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	instanceId string
	snapshotId string
}

func (r SnapshotsApiV1RestoreInstanceSnapshotRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1RestoreInstanceSnapshotExecute(r)
}

/*
V1RestoreInstanceSnapshot Restore a Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId Instance ID - uuid
 @param snapshotId Snapshot ID - uuid
 @return SnapshotsApiV1RestoreInstanceSnapshotRequest
*/
func (a *SnapshotsApiService) V1RestoreInstanceSnapshot(ctx context.Context, instanceId string, snapshotId string) SnapshotsApiV1RestoreInstanceSnapshotRequest {
	return SnapshotsApiV1RestoreInstanceSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
func (a *SnapshotsApiService) V1RestoreInstanceSnapshotExecute(r SnapshotsApiV1RestoreInstanceSnapshotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1RestoreInstanceSnapshot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{instanceId}/snapshots/{snapshotId}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SnapshotsApiV1SnapshotRenameRequest struct {
	ctx context.Context
	ApiService *SnapshotsApiService
	snapshotId string
	snapshotCreationOptions *SnapshotCreationOptions
}

// 
func (r SnapshotsApiV1SnapshotRenameRequest) SnapshotCreationOptions(snapshotCreationOptions SnapshotCreationOptions) SnapshotsApiV1SnapshotRenameRequest {
	r.snapshotCreationOptions = &snapshotCreationOptions
	return r
}

func (r SnapshotsApiV1SnapshotRenameRequest) Execute() (*Snapshot, *http.Response, error) {
	return r.ApiService.V1SnapshotRenameExecute(r)
}

/*
V1SnapshotRename Rename a Snapshot

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param snapshotId Snapshot ID - uuid
 @return SnapshotsApiV1SnapshotRenameRequest
*/
func (a *SnapshotsApiService) V1SnapshotRename(ctx context.Context, snapshotId string) SnapshotsApiV1SnapshotRenameRequest {
	return SnapshotsApiV1SnapshotRenameRequest{
		ApiService: a,
		ctx: ctx,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
//  @return Snapshot
func (a *SnapshotsApiService) V1SnapshotRenameExecute(r SnapshotsApiV1SnapshotRenameRequest) (*Snapshot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.V1SnapshotRename")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotCreationOptions == nil {
		return localVarReturnValue, nil, reportError("snapshotCreationOptions is required and must be specified")
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
	localVarPostBody = r.snapshotCreationOptions
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UserError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
