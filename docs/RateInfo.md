# RateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnRateMicrocents** | Pointer to **NullableInt32** | The amount per second, in microcents (USD), that this instance charges to be running. | [optional] 
**OffRateMicrocents** | Pointer to **NullableInt32** | The amount per second, in microcents (USD), that this instance charges to be stored. | [optional] 

## Methods

### NewRateInfo

`func NewRateInfo() *RateInfo`

NewRateInfo instantiates a new RateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateInfoWithDefaults

`func NewRateInfoWithDefaults() *RateInfo`

NewRateInfoWithDefaults instantiates a new RateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnRateMicrocents

`func (o *RateInfo) GetOnRateMicrocents() int32`

GetOnRateMicrocents returns the OnRateMicrocents field if non-nil, zero value otherwise.

### GetOnRateMicrocentsOk

`func (o *RateInfo) GetOnRateMicrocentsOk() (*int32, bool)`

GetOnRateMicrocentsOk returns a tuple with the OnRateMicrocents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRateMicrocents

`func (o *RateInfo) SetOnRateMicrocents(v int32)`

SetOnRateMicrocents sets OnRateMicrocents field to given value.

### HasOnRateMicrocents

`func (o *RateInfo) HasOnRateMicrocents() bool`

HasOnRateMicrocents returns a boolean if a field has been set.

### SetOnRateMicrocentsNil

`func (o *RateInfo) SetOnRateMicrocentsNil(b bool)`

 SetOnRateMicrocentsNil sets the value for OnRateMicrocents to be an explicit nil

### UnsetOnRateMicrocents
`func (o *RateInfo) UnsetOnRateMicrocents()`

UnsetOnRateMicrocents ensures that no value is present for OnRateMicrocents, not even an explicit nil
### GetOffRateMicrocents

`func (o *RateInfo) GetOffRateMicrocents() int32`

GetOffRateMicrocents returns the OffRateMicrocents field if non-nil, zero value otherwise.

### GetOffRateMicrocentsOk

`func (o *RateInfo) GetOffRateMicrocentsOk() (*int32, bool)`

GetOffRateMicrocentsOk returns a tuple with the OffRateMicrocents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffRateMicrocents

`func (o *RateInfo) SetOffRateMicrocents(v int32)`

SetOffRateMicrocents sets OffRateMicrocents field to given value.

### HasOffRateMicrocents

`func (o *RateInfo) HasOffRateMicrocents() bool`

HasOffRateMicrocents returns a boolean if a field has been set.

### SetOffRateMicrocentsNil

`func (o *RateInfo) SetOffRateMicrocentsNil(b bool)`

 SetOffRateMicrocentsNil sets the value for OffRateMicrocents to be an explicit nil

### UnsetOffRateMicrocents
`func (o *RateInfo) UnsetOffRateMicrocents()`

UnsetOffRateMicrocents ensures that no value is present for OffRateMicrocents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


