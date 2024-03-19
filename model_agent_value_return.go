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

// checks if the AgentValueReturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentValueReturn{}

// AgentValueReturn 
type AgentValueReturn struct {
	// The requested value
	Value NullableString `json:"value,omitempty"`
}

// NewAgentValueReturn instantiates a new AgentValueReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentValueReturn() *AgentValueReturn {
	this := AgentValueReturn{}
	return &this
}

// NewAgentValueReturnWithDefaults instantiates a new AgentValueReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentValueReturnWithDefaults() *AgentValueReturn {
	this := AgentValueReturn{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentValueReturn) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentValueReturn) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *AgentValueReturn) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *AgentValueReturn) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *AgentValueReturn) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *AgentValueReturn) UnsetValue() {
	o.Value.Unset()
}

func (o AgentValueReturn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentValueReturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableAgentValueReturn struct {
	value *AgentValueReturn
	isSet bool
}

func (v NullableAgentValueReturn) Get() *AgentValueReturn {
	return v.value
}

func (v *NullableAgentValueReturn) Set(val *AgentValueReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentValueReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentValueReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentValueReturn(val *AgentValueReturn) *NullableAgentValueReturn {
	return &NullableAgentValueReturn{value: val, isSet: true}
}

func (v NullableAgentValueReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentValueReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


