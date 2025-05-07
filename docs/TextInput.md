# TextInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | **string** | text to type | 
**KeyDuration** | Pointer to **NullableFloat32** | How long to take to type each key.  150ms if not defined. | [optional] 

## Methods

### NewTextInput

`func NewTextInput(required string, ) *TextInput`

NewTextInput instantiates a new TextInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextInputWithDefaults

`func NewTextInputWithDefaults() *TextInput`

NewTextInputWithDefaults instantiates a new TextInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *TextInput) GetRequired() string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *TextInput) GetRequiredOk() (*string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *TextInput) SetRequired(v string)`

SetRequired sets Required field to given value.


### GetKeyDuration

`func (o *TextInput) GetKeyDuration() float32`

GetKeyDuration returns the KeyDuration field if non-nil, zero value otherwise.

### GetKeyDurationOk

`func (o *TextInput) GetKeyDurationOk() (*float32, bool)`

GetKeyDurationOk returns a tuple with the KeyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDuration

`func (o *TextInput) SetKeyDuration(v float32)`

SetKeyDuration sets KeyDuration field to given value.

### HasKeyDuration

`func (o *TextInput) HasKeyDuration() bool`

HasKeyDuration returns a boolean if a field has been set.

### SetKeyDurationNil

`func (o *TextInput) SetKeyDurationNil(b bool)`

 SetKeyDurationNil sets the value for KeyDuration to be an explicit nil

### UnsetKeyDuration
`func (o *TextInput) UnsetKeyDuration()`

UnsetKeyDuration ensures that no value is present for KeyDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


