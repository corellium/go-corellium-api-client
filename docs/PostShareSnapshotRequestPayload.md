# PostShareSnapshotRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharingType** | **string** |  | 
**Password** | Pointer to **NullableString** | Password for passwordPublicLink. | [optional] 

## Methods

### NewPostShareSnapshotRequestPayload

`func NewPostShareSnapshotRequestPayload(sharingType string, ) *PostShareSnapshotRequestPayload`

NewPostShareSnapshotRequestPayload instantiates a new PostShareSnapshotRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostShareSnapshotRequestPayloadWithDefaults

`func NewPostShareSnapshotRequestPayloadWithDefaults() *PostShareSnapshotRequestPayload`

NewPostShareSnapshotRequestPayloadWithDefaults instantiates a new PostShareSnapshotRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharingType

`func (o *PostShareSnapshotRequestPayload) GetSharingType() string`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *PostShareSnapshotRequestPayload) GetSharingTypeOk() (*string, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingType

`func (o *PostShareSnapshotRequestPayload) SetSharingType(v string)`

SetSharingType sets SharingType field to given value.


### GetPassword

`func (o *PostShareSnapshotRequestPayload) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostShareSnapshotRequestPayload) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostShareSnapshotRequestPayload) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostShareSnapshotRequestPayload) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PostShareSnapshotRequestPayload) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PostShareSnapshotRequestPayload) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


