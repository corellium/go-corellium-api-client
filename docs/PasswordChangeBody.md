# PasswordChangeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** | userId | 
**OldPassword** | **string** | old password | 
**NewPassword** | **string** | new password | 

## Methods

### NewPasswordChangeBody

`func NewPasswordChangeBody(user string, oldPassword string, newPassword string, ) *PasswordChangeBody`

NewPasswordChangeBody instantiates a new PasswordChangeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordChangeBodyWithDefaults

`func NewPasswordChangeBodyWithDefaults() *PasswordChangeBody`

NewPasswordChangeBodyWithDefaults instantiates a new PasswordChangeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PasswordChangeBody) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PasswordChangeBody) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PasswordChangeBody) SetUser(v string)`

SetUser sets User field to given value.


### GetOldPassword

`func (o *PasswordChangeBody) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *PasswordChangeBody) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *PasswordChangeBody) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.


### GetNewPassword

`func (o *PasswordChangeBody) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PasswordChangeBody) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PasswordChangeBody) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


