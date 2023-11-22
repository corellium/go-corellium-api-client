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

// checks if the AgentInstallBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentInstallBody{}

// AgentInstallBody 
type AgentInstallBody struct {
	// path to app to install
	Path NullableString `json:"path,omitempty"`
}

// NewAgentInstallBody instantiates a new AgentInstallBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentInstallBody() *AgentInstallBody {
	this := AgentInstallBody{}
	return &this
}

// NewAgentInstallBodyWithDefaults instantiates a new AgentInstallBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentInstallBodyWithDefaults() *AgentInstallBody {
	this := AgentInstallBody{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInstallBody) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInstallBody) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *AgentInstallBody) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *AgentInstallBody) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *AgentInstallBody) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *AgentInstallBody) UnsetPath() {
	o.Path.Unset()
}

func (o AgentInstallBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentInstallBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	return toSerialize, nil
}

type NullableAgentInstallBody struct {
	value *AgentInstallBody
	isSet bool
}

func (v NullableAgentInstallBody) Get() *AgentInstallBody {
	return v.value
}

func (v *NullableAgentInstallBody) Set(val *AgentInstallBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentInstallBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentInstallBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentInstallBody(val *AgentInstallBody) *NullableAgentInstallBody {
	return &NullableAgentInstallBody{value: val, isSet: true}
}

func (v NullableAgentInstallBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentInstallBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


