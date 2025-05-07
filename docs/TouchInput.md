# TouchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **map[string]interface{}** | Finger Positions | 
**Buttons** | Pointer to [**[]TouchInputButtonsInner**](TouchInputButtonsInner.md) | Buttons to be held when this position is reached. Either a Button enum object or an ASCII character. No change if not defined. | [optional] 
**Normalized** | Pointer to **NullableBool** | true if you want to use normalized x,y coordinates from 0-10000 (eg 5000,5000 is always center of screen) | [optional] 
**Wait** | Pointer to **NullableFloat32** | Duration to wait before this input is executed.  Instant if not defined. | [optional] 

## Methods

### NewTouchInput

`func NewTouchInput(position map[string]interface{}, ) *TouchInput`

NewTouchInput instantiates a new TouchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTouchInputWithDefaults

`func NewTouchInputWithDefaults() *TouchInput`

NewTouchInputWithDefaults instantiates a new TouchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *TouchInput) GetPosition() map[string]interface{}`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TouchInput) GetPositionOk() (*map[string]interface{}, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TouchInput) SetPosition(v map[string]interface{})`

SetPosition sets Position field to given value.


### GetButtons

`func (o *TouchInput) GetButtons() []TouchInputButtonsInner`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *TouchInput) GetButtonsOk() (*[]TouchInputButtonsInner, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *TouchInput) SetButtons(v []TouchInputButtonsInner)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *TouchInput) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### SetButtonsNil

`func (o *TouchInput) SetButtonsNil(b bool)`

 SetButtonsNil sets the value for Buttons to be an explicit nil

### UnsetButtons
`func (o *TouchInput) UnsetButtons()`

UnsetButtons ensures that no value is present for Buttons, not even an explicit nil
### GetNormalized

`func (o *TouchInput) GetNormalized() bool`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *TouchInput) GetNormalizedOk() (*bool, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *TouchInput) SetNormalized(v bool)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *TouchInput) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### SetNormalizedNil

`func (o *TouchInput) SetNormalizedNil(b bool)`

 SetNormalizedNil sets the value for Normalized to be an explicit nil

### UnsetNormalized
`func (o *TouchInput) UnsetNormalized()`

UnsetNormalized ensures that no value is present for Normalized, not even an explicit nil
### GetWait

`func (o *TouchInput) GetWait() float32`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *TouchInput) GetWaitOk() (*float32, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *TouchInput) SetWait(v float32)`

SetWait sets Wait field to given value.

### HasWait

`func (o *TouchInput) HasWait() bool`

HasWait returns a boolean if a field has been set.

### SetWaitNil

`func (o *TouchInput) SetWaitNil(b bool)`

 SetWaitNil sets the value for Wait to be an explicit nil

### UnsetWait
`func (o *TouchInput) UnsetWait()`

UnsetWait ensures that no value is present for Wait, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


