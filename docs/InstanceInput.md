# InstanceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **map[string]interface{}** | Finger Positions | 
**Buttons** | Pointer to [**[]TouchInputButtonsInner**](TouchInputButtonsInner.md) | Buttons to be held when this position is reached. Either a Button enum object or an ASCII character. No change if not defined. | [optional] 
**Normalized** | Pointer to **NullableBool** | true if you want to use normalized x,y coordinates from 0-10000 (eg 5000,5000 is always center of screen) | [optional] 
**Wait** | Pointer to **NullableFloat32** | Duration to wait before this input is executed.  Instant if not defined. | [optional] 
**Start** | **map[string]interface{}** | Finger Positions | 
**End** | **map[string]interface{}** | Finger Positions | 
**BezierPoints** | Pointer to **[]map[string]interface{}** | array of per-finger intermediate touch positions, up to 10 depending on model.  Straight line if not defined. | [optional] 
**StartButtons** | [**[]TouchInputButtonsInner**](TouchInputButtonsInner.md) | Buttons to be held during this curve. | 
**EndButtons** | Pointer to [**[]TouchInputButtonsInner**](TouchInputButtonsInner.md) | Buttons to be left held after this curve.  Same as startButtons if not defined. | [optional] 
**Duration** | **float32** | Duration to execute this curve over. | 
**Required** | **string** | text to type | 
**KeyDuration** | Pointer to **NullableFloat32** | How long to take to type each key.  150ms if not defined. | [optional] 

## Methods

### NewInstanceInput

`func NewInstanceInput(position map[string]interface{}, start map[string]interface{}, end map[string]interface{}, startButtons []TouchInputButtonsInner, duration float32, required string, ) *InstanceInput`

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


### GetButtons

`func (o *InstanceInput) GetButtons() []TouchInputButtonsInner`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *InstanceInput) GetButtonsOk() (*[]TouchInputButtonsInner, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *InstanceInput) SetButtons(v []TouchInputButtonsInner)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *InstanceInput) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### SetButtonsNil

`func (o *InstanceInput) SetButtonsNil(b bool)`

 SetButtonsNil sets the value for Buttons to be an explicit nil

### UnsetButtons
`func (o *InstanceInput) UnsetButtons()`

UnsetButtons ensures that no value is present for Buttons, not even an explicit nil
### GetNormalized

`func (o *InstanceInput) GetNormalized() bool`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *InstanceInput) GetNormalizedOk() (*bool, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *InstanceInput) SetNormalized(v bool)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *InstanceInput) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### SetNormalizedNil

`func (o *InstanceInput) SetNormalizedNil(b bool)`

 SetNormalizedNil sets the value for Normalized to be an explicit nil

### UnsetNormalized
`func (o *InstanceInput) UnsetNormalized()`

UnsetNormalized ensures that no value is present for Normalized, not even an explicit nil
### GetWait

`func (o *InstanceInput) GetWait() float32`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *InstanceInput) GetWaitOk() (*float32, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *InstanceInput) SetWait(v float32)`

SetWait sets Wait field to given value.

### HasWait

`func (o *InstanceInput) HasWait() bool`

HasWait returns a boolean if a field has been set.

### SetWaitNil

`func (o *InstanceInput) SetWaitNil(b bool)`

 SetWaitNil sets the value for Wait to be an explicit nil

### UnsetWait
`func (o *InstanceInput) UnsetWait()`

UnsetWait ensures that no value is present for Wait, not even an explicit nil
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


### GetBezierPoints

`func (o *InstanceInput) GetBezierPoints() []map[string]interface{}`

GetBezierPoints returns the BezierPoints field if non-nil, zero value otherwise.

### GetBezierPointsOk

`func (o *InstanceInput) GetBezierPointsOk() (*[]map[string]interface{}, bool)`

GetBezierPointsOk returns a tuple with the BezierPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBezierPoints

`func (o *InstanceInput) SetBezierPoints(v []map[string]interface{})`

SetBezierPoints sets BezierPoints field to given value.

### HasBezierPoints

`func (o *InstanceInput) HasBezierPoints() bool`

HasBezierPoints returns a boolean if a field has been set.

### SetBezierPointsNil

`func (o *InstanceInput) SetBezierPointsNil(b bool)`

 SetBezierPointsNil sets the value for BezierPoints to be an explicit nil

### UnsetBezierPoints
`func (o *InstanceInput) UnsetBezierPoints()`

UnsetBezierPoints ensures that no value is present for BezierPoints, not even an explicit nil
### GetStartButtons

`func (o *InstanceInput) GetStartButtons() []TouchInputButtonsInner`

GetStartButtons returns the StartButtons field if non-nil, zero value otherwise.

### GetStartButtonsOk

`func (o *InstanceInput) GetStartButtonsOk() (*[]TouchInputButtonsInner, bool)`

GetStartButtonsOk returns a tuple with the StartButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartButtons

`func (o *InstanceInput) SetStartButtons(v []TouchInputButtonsInner)`

SetStartButtons sets StartButtons field to given value.


### GetEndButtons

`func (o *InstanceInput) GetEndButtons() []TouchInputButtonsInner`

GetEndButtons returns the EndButtons field if non-nil, zero value otherwise.

### GetEndButtonsOk

`func (o *InstanceInput) GetEndButtonsOk() (*[]TouchInputButtonsInner, bool)`

GetEndButtonsOk returns a tuple with the EndButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndButtons

`func (o *InstanceInput) SetEndButtons(v []TouchInputButtonsInner)`

SetEndButtons sets EndButtons field to given value.

### HasEndButtons

`func (o *InstanceInput) HasEndButtons() bool`

HasEndButtons returns a boolean if a field has been set.

### SetEndButtonsNil

`func (o *InstanceInput) SetEndButtonsNil(b bool)`

 SetEndButtonsNil sets the value for EndButtons to be an explicit nil

### UnsetEndButtons
`func (o *InstanceInput) UnsetEndButtons()`

UnsetEndButtons ensures that no value is present for EndButtons, not even an explicit nil
### GetDuration

`func (o *InstanceInput) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InstanceInput) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InstanceInput) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetRequired

`func (o *InstanceInput) GetRequired() string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *InstanceInput) GetRequiredOk() (*string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *InstanceInput) SetRequired(v string)`

SetRequired sets Required field to given value.


### GetKeyDuration

`func (o *InstanceInput) GetKeyDuration() float32`

GetKeyDuration returns the KeyDuration field if non-nil, zero value otherwise.

### GetKeyDurationOk

`func (o *InstanceInput) GetKeyDurationOk() (*float32, bool)`

GetKeyDurationOk returns a tuple with the KeyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDuration

`func (o *InstanceInput) SetKeyDuration(v float32)`

SetKeyDuration sets KeyDuration field to given value.

### HasKeyDuration

`func (o *InstanceInput) HasKeyDuration() bool`

HasKeyDuration returns a boolean if a field has been set.

### SetKeyDurationNil

`func (o *InstanceInput) SetKeyDurationNil(b bool)`

 SetKeyDurationNil sets the value for KeyDuration to be an explicit nil

### UnsetKeyDuration
`func (o *InstanceInput) UnsetKeyDuration()`

UnsetKeyDuration ensures that no value is present for KeyDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


