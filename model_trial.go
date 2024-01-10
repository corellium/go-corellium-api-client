/*
Corellium API

REST API to manage your virtual devices.

API version: 5.7.1-19698
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the Trial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Trial{}

// Trial 
type Trial struct {
	// Maintenance message
	DefaultDuration NullableFloat32 `json:"defaultDuration,omitempty"`
	// Maintenance header
	DefaultHours NullableFloat32 `json:"defaultHours,omitempty"`
	// Maintenance header
	DefaultDevices NullableFloat32 `json:"defaultDevices,omitempty"`
	// Maintenance header
	DefaultCores NullableFloat32 `json:"defaultCores,omitempty"`
}

// NewTrial instantiates a new Trial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrial() *Trial {
	this := Trial{}
	return &this
}

// NewTrialWithDefaults instantiates a new Trial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialWithDefaults() *Trial {
	this := Trial{}
	return &this
}

// GetDefaultDuration returns the DefaultDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trial) GetDefaultDuration() float32 {
	if o == nil || IsNil(o.DefaultDuration.Get()) {
		var ret float32
		return ret
	}
	return *o.DefaultDuration.Get()
}

// GetDefaultDurationOk returns a tuple with the DefaultDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trial) GetDefaultDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDuration.Get(), o.DefaultDuration.IsSet()
}

// HasDefaultDuration returns a boolean if a field has been set.
func (o *Trial) HasDefaultDuration() bool {
	if o != nil && o.DefaultDuration.IsSet() {
		return true
	}

	return false
}

// SetDefaultDuration gets a reference to the given NullableFloat32 and assigns it to the DefaultDuration field.
func (o *Trial) SetDefaultDuration(v float32) {
	o.DefaultDuration.Set(&v)
}
// SetDefaultDurationNil sets the value for DefaultDuration to be an explicit nil
func (o *Trial) SetDefaultDurationNil() {
	o.DefaultDuration.Set(nil)
}

// UnsetDefaultDuration ensures that no value is present for DefaultDuration, not even an explicit nil
func (o *Trial) UnsetDefaultDuration() {
	o.DefaultDuration.Unset()
}

// GetDefaultHours returns the DefaultHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trial) GetDefaultHours() float32 {
	if o == nil || IsNil(o.DefaultHours.Get()) {
		var ret float32
		return ret
	}
	return *o.DefaultHours.Get()
}

// GetDefaultHoursOk returns a tuple with the DefaultHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trial) GetDefaultHoursOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultHours.Get(), o.DefaultHours.IsSet()
}

// HasDefaultHours returns a boolean if a field has been set.
func (o *Trial) HasDefaultHours() bool {
	if o != nil && o.DefaultHours.IsSet() {
		return true
	}

	return false
}

// SetDefaultHours gets a reference to the given NullableFloat32 and assigns it to the DefaultHours field.
func (o *Trial) SetDefaultHours(v float32) {
	o.DefaultHours.Set(&v)
}
// SetDefaultHoursNil sets the value for DefaultHours to be an explicit nil
func (o *Trial) SetDefaultHoursNil() {
	o.DefaultHours.Set(nil)
}

// UnsetDefaultHours ensures that no value is present for DefaultHours, not even an explicit nil
func (o *Trial) UnsetDefaultHours() {
	o.DefaultHours.Unset()
}

// GetDefaultDevices returns the DefaultDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trial) GetDefaultDevices() float32 {
	if o == nil || IsNil(o.DefaultDevices.Get()) {
		var ret float32
		return ret
	}
	return *o.DefaultDevices.Get()
}

// GetDefaultDevicesOk returns a tuple with the DefaultDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trial) GetDefaultDevicesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDevices.Get(), o.DefaultDevices.IsSet()
}

// HasDefaultDevices returns a boolean if a field has been set.
func (o *Trial) HasDefaultDevices() bool {
	if o != nil && o.DefaultDevices.IsSet() {
		return true
	}

	return false
}

// SetDefaultDevices gets a reference to the given NullableFloat32 and assigns it to the DefaultDevices field.
func (o *Trial) SetDefaultDevices(v float32) {
	o.DefaultDevices.Set(&v)
}
// SetDefaultDevicesNil sets the value for DefaultDevices to be an explicit nil
func (o *Trial) SetDefaultDevicesNil() {
	o.DefaultDevices.Set(nil)
}

// UnsetDefaultDevices ensures that no value is present for DefaultDevices, not even an explicit nil
func (o *Trial) UnsetDefaultDevices() {
	o.DefaultDevices.Unset()
}

// GetDefaultCores returns the DefaultCores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trial) GetDefaultCores() float32 {
	if o == nil || IsNil(o.DefaultCores.Get()) {
		var ret float32
		return ret
	}
	return *o.DefaultCores.Get()
}

// GetDefaultCoresOk returns a tuple with the DefaultCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trial) GetDefaultCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultCores.Get(), o.DefaultCores.IsSet()
}

// HasDefaultCores returns a boolean if a field has been set.
func (o *Trial) HasDefaultCores() bool {
	if o != nil && o.DefaultCores.IsSet() {
		return true
	}

	return false
}

// SetDefaultCores gets a reference to the given NullableFloat32 and assigns it to the DefaultCores field.
func (o *Trial) SetDefaultCores(v float32) {
	o.DefaultCores.Set(&v)
}
// SetDefaultCoresNil sets the value for DefaultCores to be an explicit nil
func (o *Trial) SetDefaultCoresNil() {
	o.DefaultCores.Set(nil)
}

// UnsetDefaultCores ensures that no value is present for DefaultCores, not even an explicit nil
func (o *Trial) UnsetDefaultCores() {
	o.DefaultCores.Unset()
}

func (o Trial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultDuration.IsSet() {
		toSerialize["defaultDuration"] = o.DefaultDuration.Get()
	}
	if o.DefaultHours.IsSet() {
		toSerialize["defaultHours"] = o.DefaultHours.Get()
	}
	if o.DefaultDevices.IsSet() {
		toSerialize["defaultDevices"] = o.DefaultDevices.Get()
	}
	if o.DefaultCores.IsSet() {
		toSerialize["defaultCores"] = o.DefaultCores.Get()
	}
	return toSerialize, nil
}

type NullableTrial struct {
	value *Trial
	isSet bool
}

func (v NullableTrial) Get() *Trial {
	return v.value
}

func (v *NullableTrial) Set(val *Trial) {
	v.value = val
	v.isSet = true
}

func (v NullableTrial) IsSet() bool {
	return v.isSet
}

func (v *NullableTrial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrial(val *Trial) *NullableTrial {
	return &NullableTrial{value: val, isSet: true}
}

func (v NullableTrial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


