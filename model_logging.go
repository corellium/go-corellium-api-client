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

// checks if the Logging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Logging{}

// Logging 
type Logging struct {
	// Denotes whether it's in development
	Development NullableBool `json:"development,omitempty"`
	// Denotes whether to throw warnings
	ThrowWarnings NullableBool `json:"throwWarnings,omitempty"`
}

// NewLogging instantiates a new Logging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogging() *Logging {
	this := Logging{}
	return &this
}

// NewLoggingWithDefaults instantiates a new Logging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingWithDefaults() *Logging {
	this := Logging{}
	return &this
}

// GetDevelopment returns the Development field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Logging) GetDevelopment() bool {
	if o == nil || IsNil(o.Development.Get()) {
		var ret bool
		return ret
	}
	return *o.Development.Get()
}

// GetDevelopmentOk returns a tuple with the Development field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logging) GetDevelopmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Development.Get(), o.Development.IsSet()
}

// HasDevelopment returns a boolean if a field has been set.
func (o *Logging) HasDevelopment() bool {
	if o != nil && o.Development.IsSet() {
		return true
	}

	return false
}

// SetDevelopment gets a reference to the given NullableBool and assigns it to the Development field.
func (o *Logging) SetDevelopment(v bool) {
	o.Development.Set(&v)
}
// SetDevelopmentNil sets the value for Development to be an explicit nil
func (o *Logging) SetDevelopmentNil() {
	o.Development.Set(nil)
}

// UnsetDevelopment ensures that no value is present for Development, not even an explicit nil
func (o *Logging) UnsetDevelopment() {
	o.Development.Unset()
}

// GetThrowWarnings returns the ThrowWarnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Logging) GetThrowWarnings() bool {
	if o == nil || IsNil(o.ThrowWarnings.Get()) {
		var ret bool
		return ret
	}
	return *o.ThrowWarnings.Get()
}

// GetThrowWarningsOk returns a tuple with the ThrowWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logging) GetThrowWarningsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThrowWarnings.Get(), o.ThrowWarnings.IsSet()
}

// HasThrowWarnings returns a boolean if a field has been set.
func (o *Logging) HasThrowWarnings() bool {
	if o != nil && o.ThrowWarnings.IsSet() {
		return true
	}

	return false
}

// SetThrowWarnings gets a reference to the given NullableBool and assigns it to the ThrowWarnings field.
func (o *Logging) SetThrowWarnings(v bool) {
	o.ThrowWarnings.Set(&v)
}
// SetThrowWarningsNil sets the value for ThrowWarnings to be an explicit nil
func (o *Logging) SetThrowWarningsNil() {
	o.ThrowWarnings.Set(nil)
}

// UnsetThrowWarnings ensures that no value is present for ThrowWarnings, not even an explicit nil
func (o *Logging) UnsetThrowWarnings() {
	o.ThrowWarnings.Unset()
}

func (o Logging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Logging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Development.IsSet() {
		toSerialize["development"] = o.Development.Get()
	}
	if o.ThrowWarnings.IsSet() {
		toSerialize["throwWarnings"] = o.ThrowWarnings.Get()
	}
	return toSerialize, nil
}

type NullableLogging struct {
	value *Logging
	isSet bool
}

func (v NullableLogging) Get() *Logging {
	return v.value
}

func (v *NullableLogging) Set(val *Logging) {
	v.value = val
	v.isSet = true
}

func (v NullableLogging) IsSet() bool {
	return v.isSet
}

func (v *NullableLogging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogging(val *Logging) *NullableLogging {
	return &NullableLogging{value: val, isSet: true}
}

func (v NullableLogging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


