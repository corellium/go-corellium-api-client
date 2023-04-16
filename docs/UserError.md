# UserError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error text | 
**ErrorID** | **string** | Error ID | 
**Field** | Pointer to **NullableString** | Field associated with error | [optional] 

## Methods

### NewUserError

`func NewUserError(error_ string, errorID string, ) *UserError`

NewUserError instantiates a new UserError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserErrorWithDefaults

`func NewUserErrorWithDefaults() *UserError`

NewUserErrorWithDefaults instantiates a new UserError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UserError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorID

`func (o *UserError) GetErrorID() string`

GetErrorID returns the ErrorID field if non-nil, zero value otherwise.

### GetErrorIDOk

`func (o *UserError) GetErrorIDOk() (*string, bool)`

GetErrorIDOk returns a tuple with the ErrorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorID

`func (o *UserError) SetErrorID(v string)`

SetErrorID sets ErrorID field to given value.


### GetField

`func (o *UserError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *UserError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *UserError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *UserError) HasField() bool`

HasField returns a boolean if a field has been set.

### SetFieldNil

`func (o *UserError) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *UserError) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


