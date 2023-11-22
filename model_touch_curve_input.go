/*
Corellium API

REST API to manage your virtual devices.

API version: 5.6.0-19122
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the TouchCurveInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TouchCurveInput{}

// TouchCurveInput 
type TouchCurveInput struct {
	// Finger Positions
	Start map[string]interface{} `json:"start,omitempty"`
	// Finger Positions
	End map[string]interface{} `json:"end,omitempty"`
}

// NewTouchCurveInput instantiates a new TouchCurveInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTouchCurveInput() *TouchCurveInput {
	this := TouchCurveInput{}
	return &this
}

// NewTouchCurveInputWithDefaults instantiates a new TouchCurveInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTouchCurveInputWithDefaults() *TouchCurveInput {
	this := TouchCurveInput{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *TouchCurveInput) GetStart() map[string]interface{} {
	if o == nil || IsNil(o.Start) {
		var ret map[string]interface{}
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TouchCurveInput) GetStartOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Start) {
		return map[string]interface{}{}, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *TouchCurveInput) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given map[string]interface{} and assigns it to the Start field.
func (o *TouchCurveInput) SetStart(v map[string]interface{}) {
	o.Start = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *TouchCurveInput) GetEnd() map[string]interface{} {
	if o == nil || IsNil(o.End) {
		var ret map[string]interface{}
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TouchCurveInput) GetEndOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.End) {
		return map[string]interface{}{}, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *TouchCurveInput) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given map[string]interface{} and assigns it to the End field.
func (o *TouchCurveInput) SetEnd(v map[string]interface{}) {
	o.End = v
}

func (o TouchCurveInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TouchCurveInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableTouchCurveInput struct {
	value *TouchCurveInput
	isSet bool
}

func (v NullableTouchCurveInput) Get() *TouchCurveInput {
	return v.value
}

func (v *NullableTouchCurveInput) Set(val *TouchCurveInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTouchCurveInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTouchCurveInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTouchCurveInput(val *TouchCurveInput) *NullableTouchCurveInput {
	return &NullableTouchCurveInput{value: val, isSet: true}
}

func (v NullableTouchCurveInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTouchCurveInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


