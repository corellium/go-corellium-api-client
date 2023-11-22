/*
Corellium API

REST API to manage your virtual devices.

API version: 5.6.0-19122
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the PasswordResetBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordResetBody{}

// PasswordResetBody 
type PasswordResetBody struct {
	// Password reset token
	Token string `json:"token"`
	// Password reset totpToken
	TotpToken string `json:"totpToken"`
	// new password
	NewPassword string `json:"new-password"`
}

// NewPasswordResetBody instantiates a new PasswordResetBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordResetBody(token string, totpToken string, newPassword string) *PasswordResetBody {
	this := PasswordResetBody{}
	this.Token = token
	this.TotpToken = totpToken
	this.NewPassword = newPassword
	return &this
}

// NewPasswordResetBodyWithDefaults instantiates a new PasswordResetBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordResetBodyWithDefaults() *PasswordResetBody {
	this := PasswordResetBody{}
	return &this
}

// GetToken returns the Token field value
func (o *PasswordResetBody) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PasswordResetBody) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PasswordResetBody) SetToken(v string) {
	o.Token = v
}

// GetTotpToken returns the TotpToken field value
func (o *PasswordResetBody) GetTotpToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotpToken
}

// GetTotpTokenOk returns a tuple with the TotpToken field value
// and a boolean to check if the value has been set.
func (o *PasswordResetBody) GetTotpTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotpToken, true
}

// SetTotpToken sets field value
func (o *PasswordResetBody) SetTotpToken(v string) {
	o.TotpToken = v
}

// GetNewPassword returns the NewPassword field value
func (o *PasswordResetBody) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *PasswordResetBody) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *PasswordResetBody) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o PasswordResetBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordResetBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["totpToken"] = o.TotpToken
	toSerialize["new-password"] = o.NewPassword
	return toSerialize, nil
}

type NullablePasswordResetBody struct {
	value *PasswordResetBody
	isSet bool
}

func (v NullablePasswordResetBody) Get() *PasswordResetBody {
	return v.value
}

func (v *NullablePasswordResetBody) Set(val *PasswordResetBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordResetBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordResetBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordResetBody(val *PasswordResetBody) *NullablePasswordResetBody {
	return &NullablePasswordResetBody{value: val, isSet: true}
}

func (v NullablePasswordResetBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordResetBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


