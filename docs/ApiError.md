# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error text | 
**ErrorID** | **string** | Error ID | 
**OriginalError** | Pointer to **NullableString** | Upstream error description | [optional] 

## Methods

### NewApiError

`func NewApiError(error_ string, errorID string, ) *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorID

`func (o *ApiError) GetErrorID() string`

GetErrorID returns the ErrorID field if non-nil, zero value otherwise.

### GetErrorIDOk

`func (o *ApiError) GetErrorIDOk() (*string, bool)`

GetErrorIDOk returns a tuple with the ErrorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorID

`func (o *ApiError) SetErrorID(v string)`

SetErrorID sets ErrorID field to given value.


### GetOriginalError

`func (o *ApiError) GetOriginalError() string`

GetOriginalError returns the OriginalError field if non-nil, zero value otherwise.

### GetOriginalErrorOk

`func (o *ApiError) GetOriginalErrorOk() (*string, bool)`

GetOriginalErrorOk returns a tuple with the OriginalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalError

`func (o *ApiError) SetOriginalError(v string)`

SetOriginalError sets OriginalError field to given value.

### HasOriginalError

`func (o *ApiError) HasOriginalError() bool`

HasOriginalError returns a boolean if a field has been set.

### SetOriginalErrorNil

`func (o *ApiError) SetOriginalErrorNil(b bool)`

 SetOriginalErrorNil sets the value for OriginalError to be an explicit nil

### UnsetOriginalError
`func (o *ApiError) UnsetOriginalError()`

UnsetOriginalError ensures that no value is present for OriginalError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


