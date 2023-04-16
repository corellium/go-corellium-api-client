# ApiConflictError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error text | 
**ErrorID** | **string** | Error ID | 
**Object** | Pointer to **map[string]interface{}** | Object that is conflicted with | [optional] 

## Methods

### NewApiConflictError

`func NewApiConflictError(error_ string, errorID string, ) *ApiConflictError`

NewApiConflictError instantiates a new ApiConflictError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiConflictErrorWithDefaults

`func NewApiConflictErrorWithDefaults() *ApiConflictError`

NewApiConflictErrorWithDefaults instantiates a new ApiConflictError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiConflictError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiConflictError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiConflictError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorID

`func (o *ApiConflictError) GetErrorID() string`

GetErrorID returns the ErrorID field if non-nil, zero value otherwise.

### GetErrorIDOk

`func (o *ApiConflictError) GetErrorIDOk() (*string, bool)`

GetErrorIDOk returns a tuple with the ErrorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorID

`func (o *ApiConflictError) SetErrorID(v string)`

SetErrorID sets ErrorID field to given value.


### GetObject

`func (o *ApiConflictError) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ApiConflictError) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ApiConflictError) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *ApiConflictError) HasObject() bool`

HasObject returns a boolean if a field has been set.

### SetObjectNil

`func (o *ApiConflictError) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *ApiConflictError) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


