# SnapshotPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Indicates if snapshot permissions are enabled | [optional] 
**InvitationTypes** | Pointer to [**SnapshotInvitationTypes**](SnapshotInvitationTypes.md) |  | [optional] 

## Methods

### NewSnapshotPermissions

`func NewSnapshotPermissions() *SnapshotPermissions`

NewSnapshotPermissions instantiates a new SnapshotPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotPermissionsWithDefaults

`func NewSnapshotPermissionsWithDefaults() *SnapshotPermissions`

NewSnapshotPermissionsWithDefaults instantiates a new SnapshotPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SnapshotPermissions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnapshotPermissions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnapshotPermissions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnapshotPermissions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *SnapshotPermissions) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *SnapshotPermissions) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInvitationTypes

`func (o *SnapshotPermissions) GetInvitationTypes() SnapshotInvitationTypes`

GetInvitationTypes returns the InvitationTypes field if non-nil, zero value otherwise.

### GetInvitationTypesOk

`func (o *SnapshotPermissions) GetInvitationTypesOk() (*SnapshotInvitationTypes, bool)`

GetInvitationTypesOk returns a tuple with the InvitationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationTypes

`func (o *SnapshotPermissions) SetInvitationTypes(v SnapshotInvitationTypes)`

SetInvitationTypes sets InvitationTypes field to given value.

### HasInvitationTypes

`func (o *SnapshotPermissions) HasInvitationTypes() bool`

HasInvitationTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


