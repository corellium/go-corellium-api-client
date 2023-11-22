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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User /_**
type User struct {
	// User ID
	Id string `json:"id"`
	// User Label
	Label string `json:"label"`
	// User Name
	Name string `json:"name"`
	// User Email
	Email string `json:"email"`
	// the flag that specifies whether user is Administrator or not
	Administrator NullableBool `json:"administrator,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(id string, label string, name string, email string) *User {
	this := User{}
	this.Id = id
	this.Label = label
	this.Name = name
	this.Email = email
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *User) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *User) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *User) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *User) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *User) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetAdministrator() bool {
	if o == nil || IsNil(o.Administrator.Get()) {
		var ret bool
		return ret
	}
	return *o.Administrator.Get()
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetAdministratorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Administrator.Get(), o.Administrator.IsSet()
}

// HasAdministrator returns a boolean if a field has been set.
func (o *User) HasAdministrator() bool {
	if o != nil && o.Administrator.IsSet() {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given NullableBool and assigns it to the Administrator field.
func (o *User) SetAdministrator(v bool) {
	o.Administrator.Set(&v)
}
// SetAdministratorNil sets the value for Administrator to be an explicit nil
func (o *User) SetAdministratorNil() {
	o.Administrator.Set(nil)
}

// UnsetAdministrator ensures that no value is present for Administrator, not even an explicit nil
func (o *User) UnsetAdministrator() {
	o.Administrator.Unset()
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	if o.Administrator.IsSet() {
		toSerialize["administrator"] = o.Administrator.Get()
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


