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

// checks if the AgentProfilesReturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentProfilesReturn{}

// AgentProfilesReturn 
type AgentProfilesReturn struct {
	// 
	Profiles []string `json:"profiles"`
}

// NewAgentProfilesReturn instantiates a new AgentProfilesReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentProfilesReturn(profiles []string) *AgentProfilesReturn {
	this := AgentProfilesReturn{}
	this.Profiles = profiles
	return &this
}

// NewAgentProfilesReturnWithDefaults instantiates a new AgentProfilesReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentProfilesReturnWithDefaults() *AgentProfilesReturn {
	this := AgentProfilesReturn{}
	return &this
}

// GetProfiles returns the Profiles field value
func (o *AgentProfilesReturn) GetProfiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value
// and a boolean to check if the value has been set.
func (o *AgentProfilesReturn) GetProfilesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profiles, true
}

// SetProfiles sets field value
func (o *AgentProfilesReturn) SetProfiles(v []string) {
	o.Profiles = v
}

func (o AgentProfilesReturn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentProfilesReturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profiles"] = o.Profiles
	return toSerialize, nil
}

type NullableAgentProfilesReturn struct {
	value *AgentProfilesReturn
	isSet bool
}

func (v NullableAgentProfilesReturn) Get() *AgentProfilesReturn {
	return v.value
}

func (v *NullableAgentProfilesReturn) Set(val *AgentProfilesReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentProfilesReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentProfilesReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentProfilesReturn(val *AgentProfilesReturn) *NullableAgentProfilesReturn {
	return &NullableAgentProfilesReturn{value: val, isSet: true}
}

func (v NullableAgentProfilesReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentProfilesReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


