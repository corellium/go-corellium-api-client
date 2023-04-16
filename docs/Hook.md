# Hook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableString** | Identifier | [optional] 
**Label** | Pointer to **NullableString** | Label | [optional] 
**Address** | Pointer to **NullableString** | Address | [optional] 
**Patch** | Pointer to **NullableString** | Patch | [optional] 
**PatchType** | Pointer to **NullableString** | Patch Type | [optional] 
**Enabled** | Pointer to **NullableBool** | If true, instances can call required hooks | [optional] 
**CreatedAt** | Pointer to **NullableString** | Created Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | [optional] 
**UpdatedAt** | Pointer to **NullableString** | Updated Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | [optional] 
**InstanceId** | Pointer to **NullableString** | Instance Id | [optional] 

## Methods

### NewHook

`func NewHook() *Hook`

NewHook instantiates a new Hook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookWithDefaults

`func NewHookWithDefaults() *Hook`

NewHookWithDefaults instantiates a new Hook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *Hook) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Hook) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Hook) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Hook) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *Hook) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Hook) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetLabel

`func (o *Hook) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Hook) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Hook) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Hook) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Hook) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Hook) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetAddress

`func (o *Hook) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Hook) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Hook) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Hook) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Hook) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Hook) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPatch

`func (o *Hook) GetPatch() string`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *Hook) GetPatchOk() (*string, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *Hook) SetPatch(v string)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *Hook) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### SetPatchNil

`func (o *Hook) SetPatchNil(b bool)`

 SetPatchNil sets the value for Patch to be an explicit nil

### UnsetPatch
`func (o *Hook) UnsetPatch()`

UnsetPatch ensures that no value is present for Patch, not even an explicit nil
### GetPatchType

`func (o *Hook) GetPatchType() string`

GetPatchType returns the PatchType field if non-nil, zero value otherwise.

### GetPatchTypeOk

`func (o *Hook) GetPatchTypeOk() (*string, bool)`

GetPatchTypeOk returns a tuple with the PatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchType

`func (o *Hook) SetPatchType(v string)`

SetPatchType sets PatchType field to given value.

### HasPatchType

`func (o *Hook) HasPatchType() bool`

HasPatchType returns a boolean if a field has been set.

### SetPatchTypeNil

`func (o *Hook) SetPatchTypeNil(b bool)`

 SetPatchTypeNil sets the value for PatchType to be an explicit nil

### UnsetPatchType
`func (o *Hook) UnsetPatchType()`

UnsetPatchType ensures that no value is present for PatchType, not even an explicit nil
### GetEnabled

`func (o *Hook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Hook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Hook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Hook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *Hook) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *Hook) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetCreatedAt

`func (o *Hook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Hook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Hook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Hook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Hook) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Hook) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Hook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Hook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Hook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Hook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Hook) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Hook) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetInstanceId

`func (o *Hook) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Hook) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Hook) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *Hook) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *Hook) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *Hook) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


