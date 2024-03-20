# SnapshotUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user. | 
**InviteSentAt** | **float32** | The date when the invite was sent. | 
**Timestamp** | **float32** | The timestamp. | 
**UserId** | **string** | The user ID. | 
**Label** | **string** | The label. | 

## Methods

### NewSnapshotUser

`func NewSnapshotUser(email string, inviteSentAt float32, timestamp float32, userId string, label string, ) *SnapshotUser`

NewSnapshotUser instantiates a new SnapshotUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotUserWithDefaults

`func NewSnapshotUserWithDefaults() *SnapshotUser`

NewSnapshotUserWithDefaults instantiates a new SnapshotUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SnapshotUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SnapshotUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SnapshotUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetInviteSentAt

`func (o *SnapshotUser) GetInviteSentAt() float32`

GetInviteSentAt returns the InviteSentAt field if non-nil, zero value otherwise.

### GetInviteSentAtOk

`func (o *SnapshotUser) GetInviteSentAtOk() (*float32, bool)`

GetInviteSentAtOk returns a tuple with the InviteSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteSentAt

`func (o *SnapshotUser) SetInviteSentAt(v float32)`

SetInviteSentAt sets InviteSentAt field to given value.


### GetTimestamp

`func (o *SnapshotUser) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SnapshotUser) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SnapshotUser) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetUserId

`func (o *SnapshotUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SnapshotUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SnapshotUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetLabel

`func (o *SnapshotUser) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SnapshotUser) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SnapshotUser) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


