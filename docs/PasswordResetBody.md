# PasswordResetBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | Password reset token | 
**TotpToken** | **string** | Password reset totpToken | 
**NewPassword** | **string** | new password | 

## Methods

### NewPasswordResetBody

`func NewPasswordResetBody(token string, totpToken string, newPassword string, ) *PasswordResetBody`

NewPasswordResetBody instantiates a new PasswordResetBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordResetBodyWithDefaults

`func NewPasswordResetBodyWithDefaults() *PasswordResetBody`

NewPasswordResetBodyWithDefaults instantiates a new PasswordResetBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PasswordResetBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PasswordResetBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PasswordResetBody) SetToken(v string)`

SetToken sets Token field to given value.


### GetTotpToken

`func (o *PasswordResetBody) GetTotpToken() string`

GetTotpToken returns the TotpToken field if non-nil, zero value otherwise.

### GetTotpTokenOk

`func (o *PasswordResetBody) GetTotpTokenOk() (*string, bool)`

GetTotpTokenOk returns a tuple with the TotpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpToken

`func (o *PasswordResetBody) SetTotpToken(v string)`

SetTotpToken sets TotpToken field to given value.


### GetNewPassword

`func (o *PasswordResetBody) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PasswordResetBody) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PasswordResetBody) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


