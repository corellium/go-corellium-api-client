# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error text | 
**ErrorID** | **string** | Error ID | 
**Field** | Pointer to **NullableString** | Field associated with error | [optional] 

## Methods

### NewValidationError

`func NewValidationError(error_ string, errorID string, ) *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ValidationError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidationError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidationError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorID

`func (o *ValidationError) GetErrorID() string`

GetErrorID returns the ErrorID field if non-nil, zero value otherwise.

### GetErrorIDOk

`func (o *ValidationError) GetErrorIDOk() (*string, bool)`

GetErrorIDOk returns a tuple with the ErrorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorID

`func (o *ValidationError) SetErrorID(v string)`

SetErrorID sets ErrorID field to given value.


### GetField

`func (o *ValidationError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValidationError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValidationError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ValidationError) HasField() bool`

HasField returns a boolean if a field has been set.

### SetFieldNil

`func (o *ValidationError) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *ValidationError) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


