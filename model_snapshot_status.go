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

// checks if the SnapshotStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotStatus{}

// SnapshotStatus 
type SnapshotStatus struct {
	// 
	Task NullableString `json:"task,omitempty"`
	// 
	Created NullableBool `json:"created,omitempty"`
}

// NewSnapshotStatus instantiates a new SnapshotStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotStatus() *SnapshotStatus {
	this := SnapshotStatus{}
	return &this
}

// NewSnapshotStatusWithDefaults instantiates a new SnapshotStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotStatusWithDefaults() *SnapshotStatus {
	this := SnapshotStatus{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotStatus) GetTask() string {
	if o == nil || IsNil(o.Task.Get()) {
		var ret string
		return ret
	}
	return *o.Task.Get()
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotStatus) GetTaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Task.Get(), o.Task.IsSet()
}

// HasTask returns a boolean if a field has been set.
func (o *SnapshotStatus) HasTask() bool {
	if o != nil && o.Task.IsSet() {
		return true
	}

	return false
}

// SetTask gets a reference to the given NullableString and assigns it to the Task field.
func (o *SnapshotStatus) SetTask(v string) {
	o.Task.Set(&v)
}
// SetTaskNil sets the value for Task to be an explicit nil
func (o *SnapshotStatus) SetTaskNil() {
	o.Task.Set(nil)
}

// UnsetTask ensures that no value is present for Task, not even an explicit nil
func (o *SnapshotStatus) UnsetTask() {
	o.Task.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotStatus) GetCreated() bool {
	if o == nil || IsNil(o.Created.Get()) {
		var ret bool
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotStatus) GetCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *SnapshotStatus) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableBool and assigns it to the Created field.
func (o *SnapshotStatus) SetCreated(v bool) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *SnapshotStatus) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *SnapshotStatus) UnsetCreated() {
	o.Created.Unset()
}

func (o SnapshotStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Task.IsSet() {
		toSerialize["task"] = o.Task.Get()
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	return toSerialize, nil
}

type NullableSnapshotStatus struct {
	value *SnapshotStatus
	isSet bool
}

func (v NullableSnapshotStatus) Get() *SnapshotStatus {
	return v.value
}

func (v *NullableSnapshotStatus) Set(val *SnapshotStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotStatus(val *SnapshotStatus) *NullableSnapshotStatus {
	return &NullableSnapshotStatus{value: val, isSet: true}
}

func (v NullableSnapshotStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


