/*
Corellium API

REST API to manage your virtual devices.

API version: 5.5.0-18750
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"fmt"
)

// Button Button definition
type Button string

// List of Button
const (
	FINGER Button = "finger"
	HOME_BUTTON Button = "homeButton"
	HOLD_BUTTON Button = "holdButton"
	VOLUME_UP Button = "volumeUp"
	VOLUME_DOWN Button = "volumeDown"
	RINGER_SWITCH Button = "ringerSwitch"
	BACK_BUTTON Button = "backButton"
	OVERVIEW_BUTTON Button = "overviewButton"
)

// All allowed values of Button enum
var AllowedButtonEnumValues = []Button{
	"finger",
	"homeButton",
	"holdButton",
	"volumeUp",
	"volumeDown",
	"ringerSwitch",
	"backButton",
	"overviewButton",
}

func (v *Button) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Button(value)
	for _, existing := range AllowedButtonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Button", value)
}

// NewButtonFromValue returns a pointer to a valid Button
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewButtonFromValue(v string) (*Button, error) {
	ev := Button(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Button: valid values are %v", v, AllowedButtonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Button) IsValid() bool {
	for _, existing := range AllowedButtonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Button value
func (v Button) Ptr() *Button {
	return &v
}

type NullableButton struct {
	value *Button
	isSet bool
}

func (v NullableButton) Get() *Button {
	return v.value
}

func (v *NullableButton) Set(val *Button) {
	v.value = val
	v.isSet = true
}

func (v NullableButton) IsSet() bool {
	return v.isSet
}

func (v *NullableButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButton(val *Button) *NullableButton {
	return &NullableButton{value: val, isSet: true}
}

func (v NullableButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

