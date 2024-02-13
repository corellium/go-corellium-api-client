/*
Corellium API

REST API to manage your virtual devices.

API version: 6.0.0-20323
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the TeamCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamCreate{}

// TeamCreate 
type TeamCreate struct {
	// ID of newly created team
	Id NullableString `json:"id,omitempty"`
}

// NewTeamCreate instantiates a new TeamCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamCreate() *TeamCreate {
	this := TeamCreate{}
	return &this
}

// NewTeamCreateWithDefaults instantiates a new TeamCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamCreateWithDefaults() *TeamCreate {
	this := TeamCreate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamCreate) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamCreate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TeamCreate) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TeamCreate) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TeamCreate) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TeamCreate) UnsetId() {
	o.Id.Unset()
}

func (o TeamCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return toSerialize, nil
}

type NullableTeamCreate struct {
	value *TeamCreate
	isSet bool
}

func (v NullableTeamCreate) Get() *TeamCreate {
	return v.value
}

func (v *NullableTeamCreate) Set(val *TeamCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamCreate(val *TeamCreate) *NullableTeamCreate {
	return &NullableTeamCreate{value: val, isSet: true}
}

func (v NullableTeamCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


