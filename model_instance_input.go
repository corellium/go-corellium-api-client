/*
Corellium API

REST API to manage your virtual devices.

API version: 6.0.0-20323
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"fmt"
)

// InstanceInput - 
type InstanceInput struct {
	TextInput *TextInput
	TouchCurveInput *TouchCurveInput
	TouchInput *TouchInput
}

// TextInputAsInstanceInput is a convenience function that returns TextInput wrapped in InstanceInput
func TextInputAsInstanceInput(v *TextInput) InstanceInput {
	return InstanceInput{
		TextInput: v,
	}
}

// TouchCurveInputAsInstanceInput is a convenience function that returns TouchCurveInput wrapped in InstanceInput
func TouchCurveInputAsInstanceInput(v *TouchCurveInput) InstanceInput {
	return InstanceInput{
		TouchCurveInput: v,
	}
}

// TouchInputAsInstanceInput is a convenience function that returns TouchInput wrapped in InstanceInput
func TouchInputAsInstanceInput(v *TouchInput) InstanceInput {
	return InstanceInput{
		TouchInput: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InstanceInput) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TextInput
	err = newStrictDecoder(data).Decode(&dst.TextInput)
	if err == nil {
		jsonTextInput, _ := json.Marshal(dst.TextInput)
		if string(jsonTextInput) == "{}" { // empty struct
			dst.TextInput = nil
		} else {
			match++
		}
	} else {
		dst.TextInput = nil
	}

	// try to unmarshal data into TouchCurveInput
	err = newStrictDecoder(data).Decode(&dst.TouchCurveInput)
	if err == nil {
		jsonTouchCurveInput, _ := json.Marshal(dst.TouchCurveInput)
		if string(jsonTouchCurveInput) == "{}" { // empty struct
			dst.TouchCurveInput = nil
		} else {
			match++
		}
	} else {
		dst.TouchCurveInput = nil
	}

	// try to unmarshal data into TouchInput
	err = newStrictDecoder(data).Decode(&dst.TouchInput)
	if err == nil {
		jsonTouchInput, _ := json.Marshal(dst.TouchInput)
		if string(jsonTouchInput) == "{}" { // empty struct
			dst.TouchInput = nil
		} else {
			match++
		}
	} else {
		dst.TouchInput = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TextInput = nil
		dst.TouchCurveInput = nil
		dst.TouchInput = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InstanceInput)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InstanceInput)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InstanceInput) MarshalJSON() ([]byte, error) {
	if src.TextInput != nil {
		return json.Marshal(&src.TextInput)
	}

	if src.TouchCurveInput != nil {
		return json.Marshal(&src.TouchCurveInput)
	}

	if src.TouchInput != nil {
		return json.Marshal(&src.TouchInput)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InstanceInput) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TextInput != nil {
		return obj.TextInput
	}

	if obj.TouchCurveInput != nil {
		return obj.TouchCurveInput
	}

	if obj.TouchInput != nil {
		return obj.TouchInput
	}

	// all schemas are nil
	return nil
}

type NullableInstanceInput struct {
	value *InstanceInput
	isSet bool
}

func (v NullableInstanceInput) Get() *InstanceInput {
	return v.value
}

func (v *NullableInstanceInput) Set(val *InstanceInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceInput(val *InstanceInput) *NullableInstanceInput {
	return &NullableInstanceInput{value: val, isSet: true}
}

func (v NullableInstanceInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


