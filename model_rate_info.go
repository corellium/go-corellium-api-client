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

// checks if the RateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateInfo{}

// RateInfo 
type RateInfo struct {
	// The amount per second, in microcents (USD), that this instance charges to be running.
	OnRateMicrocents NullableInt32 `json:"onRateMicrocents,omitempty"`
	// The amount per second, in microcents (USD), that this instance charges to be stored.
	OffRateMicrocents NullableInt32 `json:"offRateMicrocents,omitempty"`
}

// NewRateInfo instantiates a new RateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateInfo() *RateInfo {
	this := RateInfo{}
	return &this
}

// NewRateInfoWithDefaults instantiates a new RateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateInfoWithDefaults() *RateInfo {
	this := RateInfo{}
	return &this
}

// GetOnRateMicrocents returns the OnRateMicrocents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RateInfo) GetOnRateMicrocents() int32 {
	if o == nil || IsNil(o.OnRateMicrocents.Get()) {
		var ret int32
		return ret
	}
	return *o.OnRateMicrocents.Get()
}

// GetOnRateMicrocentsOk returns a tuple with the OnRateMicrocents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RateInfo) GetOnRateMicrocentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnRateMicrocents.Get(), o.OnRateMicrocents.IsSet()
}

// HasOnRateMicrocents returns a boolean if a field has been set.
func (o *RateInfo) HasOnRateMicrocents() bool {
	if o != nil && o.OnRateMicrocents.IsSet() {
		return true
	}

	return false
}

// SetOnRateMicrocents gets a reference to the given NullableInt32 and assigns it to the OnRateMicrocents field.
func (o *RateInfo) SetOnRateMicrocents(v int32) {
	o.OnRateMicrocents.Set(&v)
}
// SetOnRateMicrocentsNil sets the value for OnRateMicrocents to be an explicit nil
func (o *RateInfo) SetOnRateMicrocentsNil() {
	o.OnRateMicrocents.Set(nil)
}

// UnsetOnRateMicrocents ensures that no value is present for OnRateMicrocents, not even an explicit nil
func (o *RateInfo) UnsetOnRateMicrocents() {
	o.OnRateMicrocents.Unset()
}

// GetOffRateMicrocents returns the OffRateMicrocents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RateInfo) GetOffRateMicrocents() int32 {
	if o == nil || IsNil(o.OffRateMicrocents.Get()) {
		var ret int32
		return ret
	}
	return *o.OffRateMicrocents.Get()
}

// GetOffRateMicrocentsOk returns a tuple with the OffRateMicrocents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RateInfo) GetOffRateMicrocentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OffRateMicrocents.Get(), o.OffRateMicrocents.IsSet()
}

// HasOffRateMicrocents returns a boolean if a field has been set.
func (o *RateInfo) HasOffRateMicrocents() bool {
	if o != nil && o.OffRateMicrocents.IsSet() {
		return true
	}

	return false
}

// SetOffRateMicrocents gets a reference to the given NullableInt32 and assigns it to the OffRateMicrocents field.
func (o *RateInfo) SetOffRateMicrocents(v int32) {
	o.OffRateMicrocents.Set(&v)
}
// SetOffRateMicrocentsNil sets the value for OffRateMicrocents to be an explicit nil
func (o *RateInfo) SetOffRateMicrocentsNil() {
	o.OffRateMicrocents.Set(nil)
}

// UnsetOffRateMicrocents ensures that no value is present for OffRateMicrocents, not even an explicit nil
func (o *RateInfo) UnsetOffRateMicrocents() {
	o.OffRateMicrocents.Unset()
}

func (o RateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OnRateMicrocents.IsSet() {
		toSerialize["onRateMicrocents"] = o.OnRateMicrocents.Get()
	}
	if o.OffRateMicrocents.IsSet() {
		toSerialize["offRateMicrocents"] = o.OffRateMicrocents.Get()
	}
	return toSerialize, nil
}

type NullableRateInfo struct {
	value *RateInfo
	isSet bool
}

func (v NullableRateInfo) Get() *RateInfo {
	return v.value
}

func (v *NullableRateInfo) Set(val *RateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateInfo(val *RateInfo) *NullableRateInfo {
	return &NullableRateInfo{value: val, isSet: true}
}

func (v NullableRateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


