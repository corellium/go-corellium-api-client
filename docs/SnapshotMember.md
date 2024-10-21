# SnapshotMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the member | 
**InviteSentAt** | **float32** | The date when the invite was sent | 
**SharedOn** | **float32** | UNIX Date of when the snapshot was first shared with member | 
**UserId** | Pointer to **NullableString** | The member&#39;s user ID | [optional] 
**Label** | Pointer to **NullableString** | The user&#39;s label or name | [optional] 

## Methods

### NewSnapshotMember

`func NewSnapshotMember(email string, inviteSentAt float32, sharedOn float32, ) *SnapshotMember`

NewSnapshotMember instantiates a new SnapshotMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotMemberWithDefaults

`func NewSnapshotMemberWithDefaults() *SnapshotMember`

NewSnapshotMemberWithDefaults instantiates a new SnapshotMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SnapshotMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SnapshotMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SnapshotMember) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetInviteSentAt

`func (o *SnapshotMember) GetInviteSentAt() float32`

GetInviteSentAt returns the InviteSentAt field if non-nil, zero value otherwise.

### GetInviteSentAtOk

`func (o *SnapshotMember) GetInviteSentAtOk() (*float32, bool)`

GetInviteSentAtOk returns a tuple with the InviteSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteSentAt

`func (o *SnapshotMember) SetInviteSentAt(v float32)`

SetInviteSentAt sets InviteSentAt field to given value.


### GetSharedOn

`func (o *SnapshotMember) GetSharedOn() float32`

GetSharedOn returns the SharedOn field if non-nil, zero value otherwise.

### GetSharedOnOk

`func (o *SnapshotMember) GetSharedOnOk() (*float32, bool)`

GetSharedOnOk returns a tuple with the SharedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOn

`func (o *SnapshotMember) SetSharedOn(v float32)`

SetSharedOn sets SharedOn field to given value.


### GetUserId

`func (o *SnapshotMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SnapshotMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SnapshotMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SnapshotMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SnapshotMember) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SnapshotMember) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetLabel

`func (o *SnapshotMember) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SnapshotMember) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SnapshotMember) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SnapshotMember) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *SnapshotMember) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *SnapshotMember) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


