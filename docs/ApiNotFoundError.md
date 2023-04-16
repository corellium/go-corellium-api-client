# ApiNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error text | 
**ErrorID** | **string** | Error ID | 
**Name** | Pointer to **NullableString** | Name of the object requested | [optional] 
**Params** | Pointer to **map[string]interface{}** | params provided by user | [optional] 

## Methods

### NewApiNotFoundError

`func NewApiNotFoundError(error_ string, errorID string, ) *ApiNotFoundError`

NewApiNotFoundError instantiates a new ApiNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNotFoundErrorWithDefaults

`func NewApiNotFoundErrorWithDefaults() *ApiNotFoundError`

NewApiNotFoundErrorWithDefaults instantiates a new ApiNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiNotFoundError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiNotFoundError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiNotFoundError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorID

`func (o *ApiNotFoundError) GetErrorID() string`

GetErrorID returns the ErrorID field if non-nil, zero value otherwise.

### GetErrorIDOk

`func (o *ApiNotFoundError) GetErrorIDOk() (*string, bool)`

GetErrorIDOk returns a tuple with the ErrorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorID

`func (o *ApiNotFoundError) SetErrorID(v string)`

SetErrorID sets ErrorID field to given value.


### GetName

`func (o *ApiNotFoundError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiNotFoundError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiNotFoundError) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiNotFoundError) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiNotFoundError) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiNotFoundError) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParams

`func (o *ApiNotFoundError) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ApiNotFoundError) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ApiNotFoundError) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ApiNotFoundError) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *ApiNotFoundError) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *ApiNotFoundError) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


