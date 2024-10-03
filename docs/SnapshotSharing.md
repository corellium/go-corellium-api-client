# SnapshotSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharingType** | **string** |  | 
**Password** | Pointer to **NullableString** | Password with using passwordPublicLink | [optional] 
**SharedBy** | Pointer to [**SnapshotOwner**](SnapshotOwner.md) |  | [optional] 
**Members** | Pointer to [**[]SnapshotMember**](SnapshotMember.md) | The members who have access to the snapshot | [optional] 

## Methods

### NewSnapshotSharing

`func NewSnapshotSharing(sharingType string, ) *SnapshotSharing`

NewSnapshotSharing instantiates a new SnapshotSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSharingWithDefaults

`func NewSnapshotSharingWithDefaults() *SnapshotSharing`

NewSnapshotSharingWithDefaults instantiates a new SnapshotSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharingType

`func (o *SnapshotSharing) GetSharingType() string`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *SnapshotSharing) GetSharingTypeOk() (*string, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingType

`func (o *SnapshotSharing) SetSharingType(v string)`

SetSharingType sets SharingType field to given value.


### GetPassword

`func (o *SnapshotSharing) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SnapshotSharing) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SnapshotSharing) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SnapshotSharing) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SnapshotSharing) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SnapshotSharing) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSharedBy

`func (o *SnapshotSharing) GetSharedBy() SnapshotOwner`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *SnapshotSharing) GetSharedByOk() (*SnapshotOwner, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *SnapshotSharing) SetSharedBy(v SnapshotOwner)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *SnapshotSharing) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetMembers

`func (o *SnapshotSharing) GetMembers() []SnapshotMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SnapshotSharing) GetMembersOk() (*[]SnapshotMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SnapshotSharing) SetMembers(v []SnapshotMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *SnapshotSharing) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *SnapshotSharing) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *SnapshotSharing) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


