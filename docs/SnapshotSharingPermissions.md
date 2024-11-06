# SnapshotSharingPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudEnabled** | Pointer to **NullableBool** | Indicates if snapshot permissions are enabled by the cloud admin | [optional] 
**DomainEnabled** | Pointer to **NullableBool** | Indicates if snapshot permissions are enabled by the domain admin | [optional] 
**PublicLink** | Pointer to **NullableBool** | Indicates if public link access is enabled | [optional] 
**DomainRestrictedLink** | Pointer to **NullableBool** | Indicates if domain-restricted link access is enabled | [optional] 
**PasswordPublicLink** | Pointer to **NullableBool** | Indicates if password-protected public link access is enabled | [optional] 
**EmailInvite** | Pointer to **NullableBool** | Indicates if email invite access is enabled | [optional] 

## Methods

### NewSnapshotSharingPermissions

`func NewSnapshotSharingPermissions() *SnapshotSharingPermissions`

NewSnapshotSharingPermissions instantiates a new SnapshotSharingPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSharingPermissionsWithDefaults

`func NewSnapshotSharingPermissionsWithDefaults() *SnapshotSharingPermissions`

NewSnapshotSharingPermissionsWithDefaults instantiates a new SnapshotSharingPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEnabled

`func (o *SnapshotSharingPermissions) GetCloudEnabled() bool`

GetCloudEnabled returns the CloudEnabled field if non-nil, zero value otherwise.

### GetCloudEnabledOk

`func (o *SnapshotSharingPermissions) GetCloudEnabledOk() (*bool, bool)`

GetCloudEnabledOk returns a tuple with the CloudEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEnabled

`func (o *SnapshotSharingPermissions) SetCloudEnabled(v bool)`

SetCloudEnabled sets CloudEnabled field to given value.

### HasCloudEnabled

`func (o *SnapshotSharingPermissions) HasCloudEnabled() bool`

HasCloudEnabled returns a boolean if a field has been set.

### SetCloudEnabledNil

`func (o *SnapshotSharingPermissions) SetCloudEnabledNil(b bool)`

 SetCloudEnabledNil sets the value for CloudEnabled to be an explicit nil

### UnsetCloudEnabled
`func (o *SnapshotSharingPermissions) UnsetCloudEnabled()`

UnsetCloudEnabled ensures that no value is present for CloudEnabled, not even an explicit nil
### GetDomainEnabled

`func (o *SnapshotSharingPermissions) GetDomainEnabled() bool`

GetDomainEnabled returns the DomainEnabled field if non-nil, zero value otherwise.

### GetDomainEnabledOk

`func (o *SnapshotSharingPermissions) GetDomainEnabledOk() (*bool, bool)`

GetDomainEnabledOk returns a tuple with the DomainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEnabled

`func (o *SnapshotSharingPermissions) SetDomainEnabled(v bool)`

SetDomainEnabled sets DomainEnabled field to given value.

### HasDomainEnabled

`func (o *SnapshotSharingPermissions) HasDomainEnabled() bool`

HasDomainEnabled returns a boolean if a field has been set.

### SetDomainEnabledNil

`func (o *SnapshotSharingPermissions) SetDomainEnabledNil(b bool)`

 SetDomainEnabledNil sets the value for DomainEnabled to be an explicit nil

### UnsetDomainEnabled
`func (o *SnapshotSharingPermissions) UnsetDomainEnabled()`

UnsetDomainEnabled ensures that no value is present for DomainEnabled, not even an explicit nil
### GetPublicLink

`func (o *SnapshotSharingPermissions) GetPublicLink() bool`

GetPublicLink returns the PublicLink field if non-nil, zero value otherwise.

### GetPublicLinkOk

`func (o *SnapshotSharingPermissions) GetPublicLinkOk() (*bool, bool)`

GetPublicLinkOk returns a tuple with the PublicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLink

`func (o *SnapshotSharingPermissions) SetPublicLink(v bool)`

SetPublicLink sets PublicLink field to given value.

### HasPublicLink

`func (o *SnapshotSharingPermissions) HasPublicLink() bool`

HasPublicLink returns a boolean if a field has been set.

### SetPublicLinkNil

`func (o *SnapshotSharingPermissions) SetPublicLinkNil(b bool)`

 SetPublicLinkNil sets the value for PublicLink to be an explicit nil

### UnsetPublicLink
`func (o *SnapshotSharingPermissions) UnsetPublicLink()`

UnsetPublicLink ensures that no value is present for PublicLink, not even an explicit nil
### GetDomainRestrictedLink

`func (o *SnapshotSharingPermissions) GetDomainRestrictedLink() bool`

GetDomainRestrictedLink returns the DomainRestrictedLink field if non-nil, zero value otherwise.

### GetDomainRestrictedLinkOk

`func (o *SnapshotSharingPermissions) GetDomainRestrictedLinkOk() (*bool, bool)`

GetDomainRestrictedLinkOk returns a tuple with the DomainRestrictedLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRestrictedLink

`func (o *SnapshotSharingPermissions) SetDomainRestrictedLink(v bool)`

SetDomainRestrictedLink sets DomainRestrictedLink field to given value.

### HasDomainRestrictedLink

`func (o *SnapshotSharingPermissions) HasDomainRestrictedLink() bool`

HasDomainRestrictedLink returns a boolean if a field has been set.

### SetDomainRestrictedLinkNil

`func (o *SnapshotSharingPermissions) SetDomainRestrictedLinkNil(b bool)`

 SetDomainRestrictedLinkNil sets the value for DomainRestrictedLink to be an explicit nil

### UnsetDomainRestrictedLink
`func (o *SnapshotSharingPermissions) UnsetDomainRestrictedLink()`

UnsetDomainRestrictedLink ensures that no value is present for DomainRestrictedLink, not even an explicit nil
### GetPasswordPublicLink

`func (o *SnapshotSharingPermissions) GetPasswordPublicLink() bool`

GetPasswordPublicLink returns the PasswordPublicLink field if non-nil, zero value otherwise.

### GetPasswordPublicLinkOk

`func (o *SnapshotSharingPermissions) GetPasswordPublicLinkOk() (*bool, bool)`

GetPasswordPublicLinkOk returns a tuple with the PasswordPublicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPublicLink

`func (o *SnapshotSharingPermissions) SetPasswordPublicLink(v bool)`

SetPasswordPublicLink sets PasswordPublicLink field to given value.

### HasPasswordPublicLink

`func (o *SnapshotSharingPermissions) HasPasswordPublicLink() bool`

HasPasswordPublicLink returns a boolean if a field has been set.

### SetPasswordPublicLinkNil

`func (o *SnapshotSharingPermissions) SetPasswordPublicLinkNil(b bool)`

 SetPasswordPublicLinkNil sets the value for PasswordPublicLink to be an explicit nil

### UnsetPasswordPublicLink
`func (o *SnapshotSharingPermissions) UnsetPasswordPublicLink()`

UnsetPasswordPublicLink ensures that no value is present for PasswordPublicLink, not even an explicit nil
### GetEmailInvite

`func (o *SnapshotSharingPermissions) GetEmailInvite() bool`

GetEmailInvite returns the EmailInvite field if non-nil, zero value otherwise.

### GetEmailInviteOk

`func (o *SnapshotSharingPermissions) GetEmailInviteOk() (*bool, bool)`

GetEmailInviteOk returns a tuple with the EmailInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailInvite

`func (o *SnapshotSharingPermissions) SetEmailInvite(v bool)`

SetEmailInvite sets EmailInvite field to given value.

### HasEmailInvite

`func (o *SnapshotSharingPermissions) HasEmailInvite() bool`

HasEmailInvite returns a boolean if a field has been set.

### SetEmailInviteNil

`func (o *SnapshotSharingPermissions) SetEmailInviteNil(b bool)`

 SetEmailInviteNil sets the value for EmailInvite to be an explicit nil

### UnsetEmailInvite
`func (o *SnapshotSharingPermissions) UnsetEmailInvite()`

UnsetEmailInvite ensures that no value is present for EmailInvite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


