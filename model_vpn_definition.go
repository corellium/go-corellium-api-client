/*
Corellium API

REST API to manage your virtual devices.

API version: 4.3.1-16664
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the VpnDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpnDefinition{}

// VpnDefinition 
type VpnDefinition struct {
	// 
	Proxy []map[string]interface{} `json:"proxy,omitempty"`
	// 
	Listeners []map[string]interface{} `json:"listeners,omitempty"`
}

// NewVpnDefinition instantiates a new VpnDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpnDefinition() *VpnDefinition {
	this := VpnDefinition{}
	return &this
}

// NewVpnDefinitionWithDefaults instantiates a new VpnDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpnDefinitionWithDefaults() *VpnDefinition {
	this := VpnDefinition{}
	return &this
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VpnDefinition) GetProxy() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnDefinition) GetProxyOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *VpnDefinition) HasProxy() bool {
	if o != nil && IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given []map[string]interface{} and assigns it to the Proxy field.
func (o *VpnDefinition) SetProxy(v []map[string]interface{}) {
	o.Proxy = v
}

// GetListeners returns the Listeners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VpnDefinition) GetListeners() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Listeners
}

// GetListenersOk returns a tuple with the Listeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnDefinition) GetListenersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Listeners) {
		return nil, false
	}
	return o.Listeners, true
}

// HasListeners returns a boolean if a field has been set.
func (o *VpnDefinition) HasListeners() bool {
	if o != nil && IsNil(o.Listeners) {
		return true
	}

	return false
}

// SetListeners gets a reference to the given []map[string]interface{} and assigns it to the Listeners field.
func (o *VpnDefinition) SetListeners(v []map[string]interface{}) {
	o.Listeners = v
}

func (o VpnDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpnDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	if o.Listeners != nil {
		toSerialize["listeners"] = o.Listeners
	}
	return toSerialize, nil
}

type NullableVpnDefinition struct {
	value *VpnDefinition
	isSet bool
}

func (v NullableVpnDefinition) Get() *VpnDefinition {
	return v.value
}

func (v *NullableVpnDefinition) Set(val *VpnDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnDefinition(val *VpnDefinition) *NullableVpnDefinition {
	return &NullableVpnDefinition{value: val, isSet: true}
}

func (v NullableVpnDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


