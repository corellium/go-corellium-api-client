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

// checks if the ProjectUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectUsage{}

// ProjectUsage 
type ProjectUsage struct {
	// 
	Cores NullableFloat32 `json:"cores,omitempty"`
	// 
	Instances NullableFloat32 `json:"instances,omitempty"`
	// 
	Ram NullableFloat32 `json:"ram,omitempty"`
	// 
	Gpu NullableFloat32 `json:"gpu,omitempty"`
}

// NewProjectUsage instantiates a new ProjectUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUsage() *ProjectUsage {
	this := ProjectUsage{}
	return &this
}

// NewProjectUsageWithDefaults instantiates a new ProjectUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUsageWithDefaults() *ProjectUsage {
	this := ProjectUsage{}
	return &this
}

// GetCores returns the Cores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectUsage) GetCores() float32 {
	if o == nil || IsNil(o.Cores.Get()) {
		var ret float32
		return ret
	}
	return *o.Cores.Get()
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectUsage) GetCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cores.Get(), o.Cores.IsSet()
}

// HasCores returns a boolean if a field has been set.
func (o *ProjectUsage) HasCores() bool {
	if o != nil && o.Cores.IsSet() {
		return true
	}

	return false
}

// SetCores gets a reference to the given NullableFloat32 and assigns it to the Cores field.
func (o *ProjectUsage) SetCores(v float32) {
	o.Cores.Set(&v)
}
// SetCoresNil sets the value for Cores to be an explicit nil
func (o *ProjectUsage) SetCoresNil() {
	o.Cores.Set(nil)
}

// UnsetCores ensures that no value is present for Cores, not even an explicit nil
func (o *ProjectUsage) UnsetCores() {
	o.Cores.Unset()
}

// GetInstances returns the Instances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectUsage) GetInstances() float32 {
	if o == nil || IsNil(o.Instances.Get()) {
		var ret float32
		return ret
	}
	return *o.Instances.Get()
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectUsage) GetInstancesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances.Get(), o.Instances.IsSet()
}

// HasInstances returns a boolean if a field has been set.
func (o *ProjectUsage) HasInstances() bool {
	if o != nil && o.Instances.IsSet() {
		return true
	}

	return false
}

// SetInstances gets a reference to the given NullableFloat32 and assigns it to the Instances field.
func (o *ProjectUsage) SetInstances(v float32) {
	o.Instances.Set(&v)
}
// SetInstancesNil sets the value for Instances to be an explicit nil
func (o *ProjectUsage) SetInstancesNil() {
	o.Instances.Set(nil)
}

// UnsetInstances ensures that no value is present for Instances, not even an explicit nil
func (o *ProjectUsage) UnsetInstances() {
	o.Instances.Unset()
}

// GetRam returns the Ram field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectUsage) GetRam() float32 {
	if o == nil || IsNil(o.Ram.Get()) {
		var ret float32
		return ret
	}
	return *o.Ram.Get()
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectUsage) GetRamOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ram.Get(), o.Ram.IsSet()
}

// HasRam returns a boolean if a field has been set.
func (o *ProjectUsage) HasRam() bool {
	if o != nil && o.Ram.IsSet() {
		return true
	}

	return false
}

// SetRam gets a reference to the given NullableFloat32 and assigns it to the Ram field.
func (o *ProjectUsage) SetRam(v float32) {
	o.Ram.Set(&v)
}
// SetRamNil sets the value for Ram to be an explicit nil
func (o *ProjectUsage) SetRamNil() {
	o.Ram.Set(nil)
}

// UnsetRam ensures that no value is present for Ram, not even an explicit nil
func (o *ProjectUsage) UnsetRam() {
	o.Ram.Unset()
}

// GetGpu returns the Gpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectUsage) GetGpu() float32 {
	if o == nil || IsNil(o.Gpu.Get()) {
		var ret float32
		return ret
	}
	return *o.Gpu.Get()
}

// GetGpuOk returns a tuple with the Gpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectUsage) GetGpuOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gpu.Get(), o.Gpu.IsSet()
}

// HasGpu returns a boolean if a field has been set.
func (o *ProjectUsage) HasGpu() bool {
	if o != nil && o.Gpu.IsSet() {
		return true
	}

	return false
}

// SetGpu gets a reference to the given NullableFloat32 and assigns it to the Gpu field.
func (o *ProjectUsage) SetGpu(v float32) {
	o.Gpu.Set(&v)
}
// SetGpuNil sets the value for Gpu to be an explicit nil
func (o *ProjectUsage) SetGpuNil() {
	o.Gpu.Set(nil)
}

// UnsetGpu ensures that no value is present for Gpu, not even an explicit nil
func (o *ProjectUsage) UnsetGpu() {
	o.Gpu.Unset()
}

func (o ProjectUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cores.IsSet() {
		toSerialize["cores"] = o.Cores.Get()
	}
	if o.Instances.IsSet() {
		toSerialize["instances"] = o.Instances.Get()
	}
	if o.Ram.IsSet() {
		toSerialize["ram"] = o.Ram.Get()
	}
	if o.Gpu.IsSet() {
		toSerialize["gpu"] = o.Gpu.Get()
	}
	return toSerialize, nil
}

type NullableProjectUsage struct {
	value *ProjectUsage
	isSet bool
}

func (v NullableProjectUsage) Get() *ProjectUsage {
	return v.value
}

func (v *NullableProjectUsage) Set(val *ProjectUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUsage(val *ProjectUsage) *NullableProjectUsage {
	return &NullableProjectUsage{value: val, isSet: true}
}

func (v NullableProjectUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


