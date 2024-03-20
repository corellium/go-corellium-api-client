# SnapshotSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the snapshot. | 
**Users** | [**[]SnapshotUser**](SnapshotUser.md) | The users who have access to the snapshot. | 

## Methods

### NewSnapshotSharing

`func NewSnapshotSharing(url string, users []SnapshotUser, ) *SnapshotSharing`

NewSnapshotSharing instantiates a new SnapshotSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSharingWithDefaults

`func NewSnapshotSharingWithDefaults() *SnapshotSharing`

NewSnapshotSharingWithDefaults instantiates a new SnapshotSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SnapshotSharing) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SnapshotSharing) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SnapshotSharing) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsers

`func (o *SnapshotSharing) GetUsers() []SnapshotUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SnapshotSharing) GetUsersOk() (*[]SnapshotUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SnapshotSharing) SetUsers(v []SnapshotUser)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


