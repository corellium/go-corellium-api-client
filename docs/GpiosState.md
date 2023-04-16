# GpiosState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Led** | Pointer to [**GpioStateDefinition**](GpioStateDefinition.md) |  | [optional] 
**Button** | Pointer to [**GpioStateDefinition**](GpioStateDefinition.md) |  | [optional] 
**Switch** | Pointer to [**GpioStateDefinition**](GpioStateDefinition.md) |  | [optional] 

## Methods

### NewGpiosState

`func NewGpiosState() *GpiosState`

NewGpiosState instantiates a new GpiosState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpiosStateWithDefaults

`func NewGpiosStateWithDefaults() *GpiosState`

NewGpiosStateWithDefaults instantiates a new GpiosState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLed

`func (o *GpiosState) GetLed() GpioStateDefinition`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *GpiosState) GetLedOk() (*GpioStateDefinition, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *GpiosState) SetLed(v GpioStateDefinition)`

SetLed sets Led field to given value.

### HasLed

`func (o *GpiosState) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetButton

`func (o *GpiosState) GetButton() GpioStateDefinition`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *GpiosState) GetButtonOk() (*GpioStateDefinition, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *GpiosState) SetButton(v GpioStateDefinition)`

SetButton sets Button field to given value.

### HasButton

`func (o *GpiosState) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetSwitch

`func (o *GpiosState) GetSwitch() GpioStateDefinition`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *GpiosState) GetSwitchOk() (*GpioStateDefinition, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *GpiosState) SetSwitch(v GpioStateDefinition)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *GpiosState) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


