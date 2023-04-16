# InstanceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **map[string]interface{}** | Finger Positions | [optional] 
**Start** | Pointer to **map[string]interface{}** | Finger Positions | [optional] 
**End** | Pointer to **map[string]interface{}** | Finger Positions | [optional] 
**Text** | Pointer to **NullableString** | text to type | [optional] 

## Methods

### NewInstanceInput

`func NewInstanceInput() *InstanceInput`

NewInstanceInput instantiates a new InstanceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInputWithDefaults

`func NewInstanceInputWithDefaults() *InstanceInput`

NewInstanceInputWithDefaults instantiates a new InstanceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *InstanceInput) GetPosition() map[string]interface{}`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *InstanceInput) GetPositionOk() (*map[string]interface{}, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *InstanceInput) SetPosition(v map[string]interface{})`

SetPosition sets Position field to given value.

### HasPosition

`func (o *InstanceInput) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStart

`func (o *InstanceInput) GetStart() map[string]interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *InstanceInput) GetStartOk() (*map[string]interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *InstanceInput) SetStart(v map[string]interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *InstanceInput) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *InstanceInput) GetEnd() map[string]interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *InstanceInput) GetEndOk() (*map[string]interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *InstanceInput) SetEnd(v map[string]interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *InstanceInput) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetText

`func (o *InstanceInput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InstanceInput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InstanceInput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *InstanceInput) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *InstanceInput) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *InstanceInput) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


