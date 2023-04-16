# InstanceStopOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Soft** | Pointer to **NullableBool** | Request VM OS power down | [optional] 

## Methods

### NewInstanceStopOptions

`func NewInstanceStopOptions() *InstanceStopOptions`

NewInstanceStopOptions instantiates a new InstanceStopOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceStopOptionsWithDefaults

`func NewInstanceStopOptionsWithDefaults() *InstanceStopOptions`

NewInstanceStopOptionsWithDefaults instantiates a new InstanceStopOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoft

`func (o *InstanceStopOptions) GetSoft() bool`

GetSoft returns the Soft field if non-nil, zero value otherwise.

### GetSoftOk

`func (o *InstanceStopOptions) GetSoftOk() (*bool, bool)`

GetSoftOk returns a tuple with the Soft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoft

`func (o *InstanceStopOptions) SetSoft(v bool)`

SetSoft sets Soft field to given value.

### HasSoft

`func (o *InstanceStopOptions) HasSoft() bool`

HasSoft returns a boolean if a field has been set.

### SetSoftNil

`func (o *InstanceStopOptions) SetSoftNil(b bool)`

 SetSoftNil sets the value for Soft to be an explicit nil

### UnsetSoft
`func (o *InstanceStopOptions) UnsetSoft()`

UnsetSoft ensures that no value is present for Soft, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


