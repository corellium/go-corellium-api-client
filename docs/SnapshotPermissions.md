# SnapshotPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Indicates if snapshot permissions are enabled | [optional] 
**PublicLink** | Pointer to **NullableBool** | Indicates if public link access is enabled | [optional] 
**DomainRestrictedLink** | Pointer to **NullableBool** | Indicates if domain-restricted link access is enabled | [optional] 
**PasswordPublicLink** | Pointer to **NullableBool** | Indicates if password-protected public link access is enabled | [optional] 
**EmailInvite** | Pointer to **NullableBool** | Indicates if email invite access is enabled | [optional] 

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
### GetPublicLink

`func (o *SnapshotPermissions) GetPublicLink() bool`

GetPublicLink returns the PublicLink field if non-nil, zero value otherwise.

### GetPublicLinkOk

`func (o *SnapshotPermissions) GetPublicLinkOk() (*bool, bool)`

GetPublicLinkOk returns a tuple with the PublicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLink

`func (o *SnapshotPermissions) SetPublicLink(v bool)`

SetPublicLink sets PublicLink field to given value.

### HasPublicLink

`func (o *SnapshotPermissions) HasPublicLink() bool`

HasPublicLink returns a boolean if a field has been set.

### SetPublicLinkNil

`func (o *SnapshotPermissions) SetPublicLinkNil(b bool)`

 SetPublicLinkNil sets the value for PublicLink to be an explicit nil

### UnsetPublicLink
`func (o *SnapshotPermissions) UnsetPublicLink()`

UnsetPublicLink ensures that no value is present for PublicLink, not even an explicit nil
### GetDomainRestrictedLink

`func (o *SnapshotPermissions) GetDomainRestrictedLink() bool`

GetDomainRestrictedLink returns the DomainRestrictedLink field if non-nil, zero value otherwise.

### GetDomainRestrictedLinkOk

`func (o *SnapshotPermissions) GetDomainRestrictedLinkOk() (*bool, bool)`

GetDomainRestrictedLinkOk returns a tuple with the DomainRestrictedLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRestrictedLink

`func (o *SnapshotPermissions) SetDomainRestrictedLink(v bool)`

SetDomainRestrictedLink sets DomainRestrictedLink field to given value.

### HasDomainRestrictedLink

`func (o *SnapshotPermissions) HasDomainRestrictedLink() bool`

HasDomainRestrictedLink returns a boolean if a field has been set.

### SetDomainRestrictedLinkNil

`func (o *SnapshotPermissions) SetDomainRestrictedLinkNil(b bool)`

 SetDomainRestrictedLinkNil sets the value for DomainRestrictedLink to be an explicit nil

### UnsetDomainRestrictedLink
`func (o *SnapshotPermissions) UnsetDomainRestrictedLink()`

UnsetDomainRestrictedLink ensures that no value is present for DomainRestrictedLink, not even an explicit nil
### GetPasswordPublicLink

`func (o *SnapshotPermissions) GetPasswordPublicLink() bool`

GetPasswordPublicLink returns the PasswordPublicLink field if non-nil, zero value otherwise.

### GetPasswordPublicLinkOk

`func (o *SnapshotPermissions) GetPasswordPublicLinkOk() (*bool, bool)`

GetPasswordPublicLinkOk returns a tuple with the PasswordPublicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPublicLink

`func (o *SnapshotPermissions) SetPasswordPublicLink(v bool)`

SetPasswordPublicLink sets PasswordPublicLink field to given value.

### HasPasswordPublicLink

`func (o *SnapshotPermissions) HasPasswordPublicLink() bool`

HasPasswordPublicLink returns a boolean if a field has been set.

### SetPasswordPublicLinkNil

`func (o *SnapshotPermissions) SetPasswordPublicLinkNil(b bool)`

 SetPasswordPublicLinkNil sets the value for PasswordPublicLink to be an explicit nil

### UnsetPasswordPublicLink
`func (o *SnapshotPermissions) UnsetPasswordPublicLink()`

UnsetPasswordPublicLink ensures that no value is present for PasswordPublicLink, not even an explicit nil
### GetEmailInvite

`func (o *SnapshotPermissions) GetEmailInvite() bool`

GetEmailInvite returns the EmailInvite field if non-nil, zero value otherwise.

### GetEmailInviteOk

`func (o *SnapshotPermissions) GetEmailInviteOk() (*bool, bool)`

GetEmailInviteOk returns a tuple with the EmailInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailInvite

`func (o *SnapshotPermissions) SetEmailInvite(v bool)`

SetEmailInvite sets EmailInvite field to given value.

### HasEmailInvite

`func (o *SnapshotPermissions) HasEmailInvite() bool`

HasEmailInvite returns a boolean if a field has been set.

### SetEmailInviteNil

`func (o *SnapshotPermissions) SetEmailInviteNil(b bool)`

 SetEmailInviteNil sets the value for EmailInvite to be an explicit nil

### UnsetEmailInvite
`func (o *SnapshotPermissions) UnsetEmailInvite()`

UnsetEmailInvite ensures that no value is present for EmailInvite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


