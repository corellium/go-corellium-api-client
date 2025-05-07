# TouchCurveInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **map[string]interface{}** | Finger Positions | 
**End** | **map[string]interface{}** | Finger Positions | 
**BezierPoints** | Pointer to **[]map[string]interface{}** | array of per-finger intermediate touch positions, up to 10 depending on model.  Straight line if not defined. | [optional] 
**StartButtons** | [**[]TouchInputButtonsInner**](TouchInputButtonsInner.md) | Buttons to be held during this curve. | 
**EndButtons** | Pointer to [**[]TouchInputButtonsInner**](TouchInputButtonsInner.md) | Buttons to be left held after this curve.  Same as startButtons if not defined. | [optional] 
**Normalized** | Pointer to **NullableBool** | true if you want to use normalized x,y coordinates from 0-10000 (eg 5000,5000 is always center of screen) | [optional] 
**Wait** | Pointer to **NullableFloat32** | Duration to wait before this input is executed.  Instant if not defined. | [optional] 
**Duration** | **float32** | Duration to execute this curve over. | 

## Methods

### NewTouchCurveInput

`func NewTouchCurveInput(start map[string]interface{}, end map[string]interface{}, startButtons []TouchInputButtonsInner, duration float32, ) *TouchCurveInput`

NewTouchCurveInput instantiates a new TouchCurveInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTouchCurveInputWithDefaults

`func NewTouchCurveInputWithDefaults() *TouchCurveInput`

NewTouchCurveInputWithDefaults instantiates a new TouchCurveInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TouchCurveInput) GetStart() map[string]interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TouchCurveInput) GetStartOk() (*map[string]interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TouchCurveInput) SetStart(v map[string]interface{})`

SetStart sets Start field to given value.


### GetEnd

`func (o *TouchCurveInput) GetEnd() map[string]interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TouchCurveInput) GetEndOk() (*map[string]interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TouchCurveInput) SetEnd(v map[string]interface{})`

SetEnd sets End field to given value.


### GetBezierPoints

`func (o *TouchCurveInput) GetBezierPoints() []map[string]interface{}`

GetBezierPoints returns the BezierPoints field if non-nil, zero value otherwise.

### GetBezierPointsOk

`func (o *TouchCurveInput) GetBezierPointsOk() (*[]map[string]interface{}, bool)`

GetBezierPointsOk returns a tuple with the BezierPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBezierPoints

`func (o *TouchCurveInput) SetBezierPoints(v []map[string]interface{})`

SetBezierPoints sets BezierPoints field to given value.

### HasBezierPoints

`func (o *TouchCurveInput) HasBezierPoints() bool`

HasBezierPoints returns a boolean if a field has been set.

### SetBezierPointsNil

`func (o *TouchCurveInput) SetBezierPointsNil(b bool)`

 SetBezierPointsNil sets the value for BezierPoints to be an explicit nil

### UnsetBezierPoints
`func (o *TouchCurveInput) UnsetBezierPoints()`

UnsetBezierPoints ensures that no value is present for BezierPoints, not even an explicit nil
### GetStartButtons

`func (o *TouchCurveInput) GetStartButtons() []TouchInputButtonsInner`

GetStartButtons returns the StartButtons field if non-nil, zero value otherwise.

### GetStartButtonsOk

`func (o *TouchCurveInput) GetStartButtonsOk() (*[]TouchInputButtonsInner, bool)`

GetStartButtonsOk returns a tuple with the StartButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartButtons

`func (o *TouchCurveInput) SetStartButtons(v []TouchInputButtonsInner)`

SetStartButtons sets StartButtons field to given value.


### GetEndButtons

`func (o *TouchCurveInput) GetEndButtons() []TouchInputButtonsInner`

GetEndButtons returns the EndButtons field if non-nil, zero value otherwise.

### GetEndButtonsOk

`func (o *TouchCurveInput) GetEndButtonsOk() (*[]TouchInputButtonsInner, bool)`

GetEndButtonsOk returns a tuple with the EndButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndButtons

`func (o *TouchCurveInput) SetEndButtons(v []TouchInputButtonsInner)`

SetEndButtons sets EndButtons field to given value.

### HasEndButtons

`func (o *TouchCurveInput) HasEndButtons() bool`

HasEndButtons returns a boolean if a field has been set.

### SetEndButtonsNil

`func (o *TouchCurveInput) SetEndButtonsNil(b bool)`

 SetEndButtonsNil sets the value for EndButtons to be an explicit nil

### UnsetEndButtons
`func (o *TouchCurveInput) UnsetEndButtons()`

UnsetEndButtons ensures that no value is present for EndButtons, not even an explicit nil
### GetNormalized

`func (o *TouchCurveInput) GetNormalized() bool`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *TouchCurveInput) GetNormalizedOk() (*bool, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *TouchCurveInput) SetNormalized(v bool)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *TouchCurveInput) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### SetNormalizedNil

`func (o *TouchCurveInput) SetNormalizedNil(b bool)`

 SetNormalizedNil sets the value for Normalized to be an explicit nil

### UnsetNormalized
`func (o *TouchCurveInput) UnsetNormalized()`

UnsetNormalized ensures that no value is present for Normalized, not even an explicit nil
### GetWait

`func (o *TouchCurveInput) GetWait() float32`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *TouchCurveInput) GetWaitOk() (*float32, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *TouchCurveInput) SetWait(v float32)`

SetWait sets Wait field to given value.

### HasWait

`func (o *TouchCurveInput) HasWait() bool`

HasWait returns a boolean if a field has been set.

### SetWaitNil

`func (o *TouchCurveInput) SetWaitNil(b bool)`

 SetWaitNil sets the value for Wait to be an explicit nil

### UnsetWait
`func (o *TouchCurveInput) UnsetWait()`

UnsetWait ensures that no value is present for Wait, not even an explicit nil
### GetDuration

`func (o *TouchCurveInput) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TouchCurveInput) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TouchCurveInput) SetDuration(v float32)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


