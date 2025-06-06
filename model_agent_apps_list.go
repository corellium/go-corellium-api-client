/*
Corellium API

REST API to manage your virtual devices.

API version: 7.3.0-27831
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the AgentAppsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentAppsList{}

// AgentAppsList 
type AgentAppsList struct {
	// 
	Apps []AgentApp `json:"apps,omitempty"`
	// bundleID of frontmost app or empty string if none are frontmost
	Frontmost NullableString `json:"frontmost,omitempty"`
}

// NewAgentAppsList instantiates a new AgentAppsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentAppsList() *AgentAppsList {
	this := AgentAppsList{}
	return &this
}

// NewAgentAppsListWithDefaults instantiates a new AgentAppsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentAppsListWithDefaults() *AgentAppsList {
	this := AgentAppsList{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentAppsList) GetApps() []AgentApp {
	if o == nil {
		var ret []AgentApp
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentAppsList) GetAppsOk() ([]AgentApp, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *AgentAppsList) HasApps() bool {
	if o != nil && IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []AgentApp and assigns it to the Apps field.
func (o *AgentAppsList) SetApps(v []AgentApp) {
	o.Apps = v
}

// GetFrontmost returns the Frontmost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentAppsList) GetFrontmost() string {
	if o == nil || IsNil(o.Frontmost.Get()) {
		var ret string
		return ret
	}
	return *o.Frontmost.Get()
}

// GetFrontmostOk returns a tuple with the Frontmost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentAppsList) GetFrontmostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frontmost.Get(), o.Frontmost.IsSet()
}

// HasFrontmost returns a boolean if a field has been set.
func (o *AgentAppsList) HasFrontmost() bool {
	if o != nil && o.Frontmost.IsSet() {
		return true
	}

	return false
}

// SetFrontmost gets a reference to the given NullableString and assigns it to the Frontmost field.
func (o *AgentAppsList) SetFrontmost(v string) {
	o.Frontmost.Set(&v)
}
// SetFrontmostNil sets the value for Frontmost to be an explicit nil
func (o *AgentAppsList) SetFrontmostNil() {
	o.Frontmost.Set(nil)
}

// UnsetFrontmost ensures that no value is present for Frontmost, not even an explicit nil
func (o *AgentAppsList) UnsetFrontmost() {
	o.Frontmost.Unset()
}

func (o AgentAppsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentAppsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.Frontmost.IsSet() {
		toSerialize["frontmost"] = o.Frontmost.Get()
	}
	return toSerialize, nil
}

type NullableAgentAppsList struct {
	value *AgentAppsList
	isSet bool
}

func (v NullableAgentAppsList) Get() *AgentAppsList {
	return v.value
}

func (v *NullableAgentAppsList) Set(val *AgentAppsList) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentAppsList) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentAppsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentAppsList(val *AgentAppsList) *NullableAgentAppsList {
	return &NullableAgentAppsList{value: val, isSet: true}
}

func (v NullableAgentAppsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentAppsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


