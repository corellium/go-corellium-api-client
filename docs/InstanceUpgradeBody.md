# InstanceUpgradeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | **string** | iOS version | 
**Osbuild** | Pointer to **NullableString** | (optional) iOS build ID | [optional] 

## Methods

### NewInstanceUpgradeBody

`func NewInstanceUpgradeBody(os string, ) *InstanceUpgradeBody`

NewInstanceUpgradeBody instantiates a new InstanceUpgradeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUpgradeBodyWithDefaults

`func NewInstanceUpgradeBodyWithDefaults() *InstanceUpgradeBody`

NewInstanceUpgradeBodyWithDefaults instantiates a new InstanceUpgradeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *InstanceUpgradeBody) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstanceUpgradeBody) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstanceUpgradeBody) SetOs(v string)`

SetOs sets Os field to given value.


### GetOsbuild

`func (o *InstanceUpgradeBody) GetOsbuild() string`

GetOsbuild returns the Osbuild field if non-nil, zero value otherwise.

### GetOsbuildOk

`func (o *InstanceUpgradeBody) GetOsbuildOk() (*string, bool)`

GetOsbuildOk returns a tuple with the Osbuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsbuild

`func (o *InstanceUpgradeBody) SetOsbuild(v string)`

SetOsbuild sets Osbuild field to given value.

### HasOsbuild

`func (o *InstanceUpgradeBody) HasOsbuild() bool`

HasOsbuild returns a boolean if a field has been set.

### SetOsbuildNil

`func (o *InstanceUpgradeBody) SetOsbuildNil(b bool)`

 SetOsbuildNil sets the value for Osbuild to be an explicit nil

### UnsetOsbuild
`func (o *InstanceUpgradeBody) UnsetOsbuild()`

UnsetOsbuild ensures that no value is present for Osbuild, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


