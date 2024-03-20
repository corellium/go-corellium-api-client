/*
Corellium API

REST API to manage your virtual devices.

API version: 6.1.0-20784
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the PostSnapshotPermissionsRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSnapshotPermissionsRequestPayload{}

// PostSnapshotPermissionsRequestPayload 
type PostSnapshotPermissionsRequestPayload struct {
	// The email of the user.
	Email string `json:"email"`
	// The user ID.
	UserId NullableString `json:"userId,omitempty"`
}

// NewPostSnapshotPermissionsRequestPayload instantiates a new PostSnapshotPermissionsRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSnapshotPermissionsRequestPayload(email string) *PostSnapshotPermissionsRequestPayload {
	this := PostSnapshotPermissionsRequestPayload{}
	this.Email = email
	return &this
}

// NewPostSnapshotPermissionsRequestPayloadWithDefaults instantiates a new PostSnapshotPermissionsRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSnapshotPermissionsRequestPayloadWithDefaults() *PostSnapshotPermissionsRequestPayload {
	this := PostSnapshotPermissionsRequestPayload{}
	return &this
}

// GetEmail returns the Email field value
func (o *PostSnapshotPermissionsRequestPayload) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PostSnapshotPermissionsRequestPayload) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PostSnapshotPermissionsRequestPayload) SetEmail(v string) {
	o.Email = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSnapshotPermissionsRequestPayload) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSnapshotPermissionsRequestPayload) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *PostSnapshotPermissionsRequestPayload) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *PostSnapshotPermissionsRequestPayload) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *PostSnapshotPermissionsRequestPayload) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *PostSnapshotPermissionsRequestPayload) UnsetUserId() {
	o.UserId.Unset()
}

func (o PostSnapshotPermissionsRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSnapshotPermissionsRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	return toSerialize, nil
}

type NullablePostSnapshotPermissionsRequestPayload struct {
	value *PostSnapshotPermissionsRequestPayload
	isSet bool
}

func (v NullablePostSnapshotPermissionsRequestPayload) Get() *PostSnapshotPermissionsRequestPayload {
	return v.value
}

func (v *NullablePostSnapshotPermissionsRequestPayload) Set(val *PostSnapshotPermissionsRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSnapshotPermissionsRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSnapshotPermissionsRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSnapshotPermissionsRequestPayload(val *PostSnapshotPermissionsRequestPayload) *NullablePostSnapshotPermissionsRequestPayload {
	return &NullablePostSnapshotPermissionsRequestPayload{value: val, isSet: true}
}

func (v NullablePostSnapshotPermissionsRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSnapshotPermissionsRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


