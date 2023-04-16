# TrialExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **NullableFloat32** | how many additional days? | [optional] 
**Reason** | Pointer to **NullableString** | why does the user want more time? | [optional] 
**Denied** | Pointer to **NullableBool** |  | [optional] 
**DeniedReason** | Pointer to **NullableString** | explanation for why the request was denied | [optional] 

## Methods

### NewTrialExtension

`func NewTrialExtension() *TrialExtension`

NewTrialExtension instantiates a new TrialExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialExtensionWithDefaults

`func NewTrialExtensionWithDefaults() *TrialExtension`

NewTrialExtensionWithDefaults instantiates a new TrialExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *TrialExtension) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TrialExtension) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TrialExtension) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TrialExtension) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TrialExtension) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TrialExtension) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetReason

`func (o *TrialExtension) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TrialExtension) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TrialExtension) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TrialExtension) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *TrialExtension) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *TrialExtension) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetDenied

`func (o *TrialExtension) GetDenied() bool`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *TrialExtension) GetDeniedOk() (*bool, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *TrialExtension) SetDenied(v bool)`

SetDenied sets Denied field to given value.

### HasDenied

`func (o *TrialExtension) HasDenied() bool`

HasDenied returns a boolean if a field has been set.

### SetDeniedNil

`func (o *TrialExtension) SetDeniedNil(b bool)`

 SetDeniedNil sets the value for Denied to be an explicit nil

### UnsetDenied
`func (o *TrialExtension) UnsetDenied()`

UnsetDenied ensures that no value is present for Denied, not even an explicit nil
### GetDeniedReason

`func (o *TrialExtension) GetDeniedReason() string`

GetDeniedReason returns the DeniedReason field if non-nil, zero value otherwise.

### GetDeniedReasonOk

`func (o *TrialExtension) GetDeniedReasonOk() (*string, bool)`

GetDeniedReasonOk returns a tuple with the DeniedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedReason

`func (o *TrialExtension) SetDeniedReason(v string)`

SetDeniedReason sets DeniedReason field to given value.

### HasDeniedReason

`func (o *TrialExtension) HasDeniedReason() bool`

HasDeniedReason returns a boolean if a field has been set.

### SetDeniedReasonNil

`func (o *TrialExtension) SetDeniedReasonNil(b bool)`

 SetDeniedReasonNil sets the value for DeniedReason to be an explicit nil

### UnsetDeniedReason
`func (o *TrialExtension) UnsetDeniedReason()`

UnsetDeniedReason ensures that no value is present for DeniedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


