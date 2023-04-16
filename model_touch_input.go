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

// checks if the TouchInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TouchInput{}

// TouchInput 
type TouchInput struct {
	// Finger Positions
	Position map[string]interface{} `json:"position,omitempty"`
}

// NewTouchInput instantiates a new TouchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTouchInput() *TouchInput {
	this := TouchInput{}
	return &this
}

// NewTouchInputWithDefaults instantiates a new TouchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTouchInputWithDefaults() *TouchInput {
	this := TouchInput{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *TouchInput) GetPosition() map[string]interface{} {
	if o == nil || IsNil(o.Position) {
		var ret map[string]interface{}
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TouchInput) GetPositionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Position) {
		return map[string]interface{}{}, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TouchInput) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given map[string]interface{} and assigns it to the Position field.
func (o *TouchInput) SetPosition(v map[string]interface{}) {
	o.Position = v
}

func (o TouchInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TouchInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableTouchInput struct {
	value *TouchInput
	isSet bool
}

func (v NullableTouchInput) Get() *TouchInput {
	return v.value
}

func (v *NullableTouchInput) Set(val *TouchInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTouchInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTouchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTouchInput(val *TouchInput) *NullableTouchInput {
	return &NullableTouchInput{value: val, isSet: true}
}

func (v NullableTouchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTouchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


