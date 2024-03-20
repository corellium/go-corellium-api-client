# PostSnapshotPermissionsRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user. | 
**UserId** | Pointer to **NullableString** | The user ID. | [optional] 

## Methods

### NewPostSnapshotPermissionsRequestPayload

`func NewPostSnapshotPermissionsRequestPayload(email string, ) *PostSnapshotPermissionsRequestPayload`

NewPostSnapshotPermissionsRequestPayload instantiates a new PostSnapshotPermissionsRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSnapshotPermissionsRequestPayloadWithDefaults

`func NewPostSnapshotPermissionsRequestPayloadWithDefaults() *PostSnapshotPermissionsRequestPayload`

NewPostSnapshotPermissionsRequestPayloadWithDefaults instantiates a new PostSnapshotPermissionsRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostSnapshotPermissionsRequestPayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostSnapshotPermissionsRequestPayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostSnapshotPermissionsRequestPayload) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUserId

`func (o *PostSnapshotPermissionsRequestPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostSnapshotPermissionsRequestPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostSnapshotPermissionsRequestPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PostSnapshotPermissionsRequestPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *PostSnapshotPermissionsRequestPayload) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *PostSnapshotPermissionsRequestPayload) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


