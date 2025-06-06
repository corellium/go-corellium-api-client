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

// checks if the KernelThread type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KernelThread{}

// KernelThread 
type KernelThread struct {
	// Kernel Thread ID
	KernelId NullableString `json:"kernelId,omitempty"`
	// Task ID
	Tid NullableFloat32 `json:"tid,omitempty"`
	// Array of stack addresses
	Stack []string `json:"stack,omitempty"`
}

// NewKernelThread instantiates a new KernelThread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKernelThread() *KernelThread {
	this := KernelThread{}
	return &this
}

// NewKernelThreadWithDefaults instantiates a new KernelThread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKernelThreadWithDefaults() *KernelThread {
	this := KernelThread{}
	return &this
}

// GetKernelId returns the KernelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KernelThread) GetKernelId() string {
	if o == nil || IsNil(o.KernelId.Get()) {
		var ret string
		return ret
	}
	return *o.KernelId.Get()
}

// GetKernelIdOk returns a tuple with the KernelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KernelThread) GetKernelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KernelId.Get(), o.KernelId.IsSet()
}

// HasKernelId returns a boolean if a field has been set.
func (o *KernelThread) HasKernelId() bool {
	if o != nil && o.KernelId.IsSet() {
		return true
	}

	return false
}

// SetKernelId gets a reference to the given NullableString and assigns it to the KernelId field.
func (o *KernelThread) SetKernelId(v string) {
	o.KernelId.Set(&v)
}
// SetKernelIdNil sets the value for KernelId to be an explicit nil
func (o *KernelThread) SetKernelIdNil() {
	o.KernelId.Set(nil)
}

// UnsetKernelId ensures that no value is present for KernelId, not even an explicit nil
func (o *KernelThread) UnsetKernelId() {
	o.KernelId.Unset()
}

// GetTid returns the Tid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KernelThread) GetTid() float32 {
	if o == nil || IsNil(o.Tid.Get()) {
		var ret float32
		return ret
	}
	return *o.Tid.Get()
}

// GetTidOk returns a tuple with the Tid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KernelThread) GetTidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tid.Get(), o.Tid.IsSet()
}

// HasTid returns a boolean if a field has been set.
func (o *KernelThread) HasTid() bool {
	if o != nil && o.Tid.IsSet() {
		return true
	}

	return false
}

// SetTid gets a reference to the given NullableFloat32 and assigns it to the Tid field.
func (o *KernelThread) SetTid(v float32) {
	o.Tid.Set(&v)
}
// SetTidNil sets the value for Tid to be an explicit nil
func (o *KernelThread) SetTidNil() {
	o.Tid.Set(nil)
}

// UnsetTid ensures that no value is present for Tid, not even an explicit nil
func (o *KernelThread) UnsetTid() {
	o.Tid.Unset()
}

// GetStack returns the Stack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KernelThread) GetStack() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KernelThread) GetStackOk() ([]string, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *KernelThread) HasStack() bool {
	if o != nil && IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given []string and assigns it to the Stack field.
func (o *KernelThread) SetStack(v []string) {
	o.Stack = v
}

func (o KernelThread) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KernelThread) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.KernelId.IsSet() {
		toSerialize["kernelId"] = o.KernelId.Get()
	}
	if o.Tid.IsSet() {
		toSerialize["tid"] = o.Tid.Get()
	}
	if o.Stack != nil {
		toSerialize["stack"] = o.Stack
	}
	return toSerialize, nil
}

type NullableKernelThread struct {
	value *KernelThread
	isSet bool
}

func (v NullableKernelThread) Get() *KernelThread {
	return v.value
}

func (v *NullableKernelThread) Set(val *KernelThread) {
	v.value = val
	v.isSet = true
}

func (v NullableKernelThread) IsSet() bool {
	return v.isSet
}

func (v *NullableKernelThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKernelThread(val *KernelThread) *NullableKernelThread {
	return &NullableKernelThread{value: val, isSet: true}
}

func (v NullableKernelThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKernelThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


