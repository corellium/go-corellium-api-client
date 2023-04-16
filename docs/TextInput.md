# TextInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **NullableString** | text to type | [optional] 

## Methods

### NewTextInput

`func NewTextInput() *TextInput`

NewTextInput instantiates a new TextInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextInputWithDefaults

`func NewTextInputWithDefaults() *TextInput`

NewTextInputWithDefaults instantiates a new TextInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *TextInput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextInput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextInput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TextInput) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *TextInput) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *TextInput) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


