# InstanceBootOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootArgs** | Pointer to **NullableString** |  | [optional] 
**RestoreBootArgs** | Pointer to **NullableString** |  | [optional] 
**Udid** | Pointer to **NullableString** | Boot udid | [optional] 
**Ecid** | Pointer to **NullableString** | Assigned ECID | [optional] 
**RandomSeed** | Pointer to **NullableString** | Random seed to provide to boot if any | [optional] 
**Pac** | Pointer to **NullableBool** | Enable PAC | [optional] 
**Aprr** | Pointer to **NullableBool** | Enable APRR | [optional] 
**AdditionalTags** | Pointer to [**[]InstanceBootOptionsAdditionalTag**](InstanceBootOptionsAdditionalTag.md) |  | [optional] 

## Methods

### NewInstanceBootOptions

`func NewInstanceBootOptions() *InstanceBootOptions`

NewInstanceBootOptions instantiates a new InstanceBootOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceBootOptionsWithDefaults

`func NewInstanceBootOptionsWithDefaults() *InstanceBootOptions`

NewInstanceBootOptionsWithDefaults instantiates a new InstanceBootOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootArgs

`func (o *InstanceBootOptions) GetBootArgs() string`

GetBootArgs returns the BootArgs field if non-nil, zero value otherwise.

### GetBootArgsOk

`func (o *InstanceBootOptions) GetBootArgsOk() (*string, bool)`

GetBootArgsOk returns a tuple with the BootArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootArgs

`func (o *InstanceBootOptions) SetBootArgs(v string)`

SetBootArgs sets BootArgs field to given value.

### HasBootArgs

`func (o *InstanceBootOptions) HasBootArgs() bool`

HasBootArgs returns a boolean if a field has been set.

### SetBootArgsNil

`func (o *InstanceBootOptions) SetBootArgsNil(b bool)`

 SetBootArgsNil sets the value for BootArgs to be an explicit nil

### UnsetBootArgs
`func (o *InstanceBootOptions) UnsetBootArgs()`

UnsetBootArgs ensures that no value is present for BootArgs, not even an explicit nil
### GetRestoreBootArgs

`func (o *InstanceBootOptions) GetRestoreBootArgs() string`

GetRestoreBootArgs returns the RestoreBootArgs field if non-nil, zero value otherwise.

### GetRestoreBootArgsOk

`func (o *InstanceBootOptions) GetRestoreBootArgsOk() (*string, bool)`

GetRestoreBootArgsOk returns a tuple with the RestoreBootArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreBootArgs

`func (o *InstanceBootOptions) SetRestoreBootArgs(v string)`

SetRestoreBootArgs sets RestoreBootArgs field to given value.

### HasRestoreBootArgs

`func (o *InstanceBootOptions) HasRestoreBootArgs() bool`

HasRestoreBootArgs returns a boolean if a field has been set.

### SetRestoreBootArgsNil

`func (o *InstanceBootOptions) SetRestoreBootArgsNil(b bool)`

 SetRestoreBootArgsNil sets the value for RestoreBootArgs to be an explicit nil

### UnsetRestoreBootArgs
`func (o *InstanceBootOptions) UnsetRestoreBootArgs()`

UnsetRestoreBootArgs ensures that no value is present for RestoreBootArgs, not even an explicit nil
### GetUdid

`func (o *InstanceBootOptions) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *InstanceBootOptions) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *InstanceBootOptions) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *InstanceBootOptions) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### SetUdidNil

`func (o *InstanceBootOptions) SetUdidNil(b bool)`

 SetUdidNil sets the value for Udid to be an explicit nil

### UnsetUdid
`func (o *InstanceBootOptions) UnsetUdid()`

UnsetUdid ensures that no value is present for Udid, not even an explicit nil
### GetEcid

`func (o *InstanceBootOptions) GetEcid() string`

GetEcid returns the Ecid field if non-nil, zero value otherwise.

### GetEcidOk

`func (o *InstanceBootOptions) GetEcidOk() (*string, bool)`

GetEcidOk returns a tuple with the Ecid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcid

`func (o *InstanceBootOptions) SetEcid(v string)`

SetEcid sets Ecid field to given value.

### HasEcid

`func (o *InstanceBootOptions) HasEcid() bool`

HasEcid returns a boolean if a field has been set.

### SetEcidNil

`func (o *InstanceBootOptions) SetEcidNil(b bool)`

 SetEcidNil sets the value for Ecid to be an explicit nil

### UnsetEcid
`func (o *InstanceBootOptions) UnsetEcid()`

UnsetEcid ensures that no value is present for Ecid, not even an explicit nil
### GetRandomSeed

`func (o *InstanceBootOptions) GetRandomSeed() string`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *InstanceBootOptions) GetRandomSeedOk() (*string, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *InstanceBootOptions) SetRandomSeed(v string)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *InstanceBootOptions) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### SetRandomSeedNil

`func (o *InstanceBootOptions) SetRandomSeedNil(b bool)`

 SetRandomSeedNil sets the value for RandomSeed to be an explicit nil

### UnsetRandomSeed
`func (o *InstanceBootOptions) UnsetRandomSeed()`

UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
### GetPac

`func (o *InstanceBootOptions) GetPac() bool`

GetPac returns the Pac field if non-nil, zero value otherwise.

### GetPacOk

`func (o *InstanceBootOptions) GetPacOk() (*bool, bool)`

GetPacOk returns a tuple with the Pac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPac

`func (o *InstanceBootOptions) SetPac(v bool)`

SetPac sets Pac field to given value.

### HasPac

`func (o *InstanceBootOptions) HasPac() bool`

HasPac returns a boolean if a field has been set.

### SetPacNil

`func (o *InstanceBootOptions) SetPacNil(b bool)`

 SetPacNil sets the value for Pac to be an explicit nil

### UnsetPac
`func (o *InstanceBootOptions) UnsetPac()`

UnsetPac ensures that no value is present for Pac, not even an explicit nil
### GetAprr

`func (o *InstanceBootOptions) GetAprr() bool`

GetAprr returns the Aprr field if non-nil, zero value otherwise.

### GetAprrOk

`func (o *InstanceBootOptions) GetAprrOk() (*bool, bool)`

GetAprrOk returns a tuple with the Aprr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprr

`func (o *InstanceBootOptions) SetAprr(v bool)`

SetAprr sets Aprr field to given value.

### HasAprr

`func (o *InstanceBootOptions) HasAprr() bool`

HasAprr returns a boolean if a field has been set.

### SetAprrNil

`func (o *InstanceBootOptions) SetAprrNil(b bool)`

 SetAprrNil sets the value for Aprr to be an explicit nil

### UnsetAprr
`func (o *InstanceBootOptions) UnsetAprr()`

UnsetAprr ensures that no value is present for Aprr, not even an explicit nil
### GetAdditionalTags

`func (o *InstanceBootOptions) GetAdditionalTags() []InstanceBootOptionsAdditionalTag`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *InstanceBootOptions) GetAdditionalTagsOk() (*[]InstanceBootOptionsAdditionalTag, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *InstanceBootOptions) SetAdditionalTags(v []InstanceBootOptionsAdditionalTag)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *InstanceBootOptions) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### SetAdditionalTagsNil

`func (o *InstanceBootOptions) SetAdditionalTagsNil(b bool)`

 SetAdditionalTagsNil sets the value for AdditionalTags to be an explicit nil

### UnsetAdditionalTags
`func (o *InstanceBootOptions) UnsetAdditionalTags()`

UnsetAdditionalTags ensures that no value is present for AdditionalTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


