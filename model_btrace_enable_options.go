/*
Corellium API

REST API to manage your virtual devices.

API version: 5.5.0-18750
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the BtraceEnableOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BtraceEnableOptions{}

// BtraceEnableOptions 
type BtraceEnableOptions struct {
	// 
	Ranges [][]string `json:"ranges,omitempty"`
}

// NewBtraceEnableOptions instantiates a new BtraceEnableOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtraceEnableOptions() *BtraceEnableOptions {
	this := BtraceEnableOptions{}
	return &this
}

// NewBtraceEnableOptionsWithDefaults instantiates a new BtraceEnableOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtraceEnableOptionsWithDefaults() *BtraceEnableOptions {
	this := BtraceEnableOptions{}
	return &this
}

// GetRanges returns the Ranges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BtraceEnableOptions) GetRanges() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BtraceEnableOptions) GetRangesOk() ([][]string, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *BtraceEnableOptions) HasRanges() bool {
	if o != nil && IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given [][]string and assigns it to the Ranges field.
func (o *BtraceEnableOptions) SetRanges(v [][]string) {
	o.Ranges = v
}

func (o BtraceEnableOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BtraceEnableOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	return toSerialize, nil
}

type NullableBtraceEnableOptions struct {
	value *BtraceEnableOptions
	isSet bool
}

func (v NullableBtraceEnableOptions) Get() *BtraceEnableOptions {
	return v.value
}

func (v *NullableBtraceEnableOptions) Set(val *BtraceEnableOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBtraceEnableOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBtraceEnableOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtraceEnableOptions(val *BtraceEnableOptions) *NullableBtraceEnableOptions {
	return &NullableBtraceEnableOptions{value: val, isSet: true}
}

func (v NullableBtraceEnableOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtraceEnableOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


