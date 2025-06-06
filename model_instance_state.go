/*
Corellium API

REST API to manage your virtual devices.

API version: 7.3.0-27831
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"fmt"
)

// InstanceState Current Instance State
type InstanceState string

// List of InstanceState
const (
	ON InstanceState = "on"
	OFF InstanceState = "off"
	BOOTING InstanceState = "booting"
	DELETING InstanceState = "deleting"
	CREATING InstanceState = "creating"
	RESTORING InstanceState = "restoring"
	PAUSED InstanceState = "paused"
	REBOOTING InstanceState = "rebooting"
	ERROR InstanceState = "error"
)

// All allowed values of InstanceState enum
var AllowedInstanceStateEnumValues = []InstanceState{
	"on",
	"off",
	"booting",
	"deleting",
	"creating",
	"restoring",
	"paused",
	"rebooting",
	"error",
}

func (v *InstanceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceState(value)
	for _, existing := range AllowedInstanceStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceState", value)
}

// NewInstanceStateFromValue returns a pointer to a valid InstanceState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceStateFromValue(v string) (*InstanceState, error) {
	ev := InstanceState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceState: valid values are %v", v, AllowedInstanceStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceState) IsValid() bool {
	for _, existing := range AllowedInstanceStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceState value
func (v InstanceState) Ptr() *InstanceState {
	return &v
}

type NullableInstanceState struct {
	value *InstanceState
	isSet bool
}

func (v NullableInstanceState) Get() *InstanceState {
	return v.value
}

func (v *NullableInstanceState) Set(val *InstanceState) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceState) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceState(val *InstanceState) *NullableInstanceState {
	return &NullableInstanceState{value: val, isSet: true}
}

func (v NullableInstanceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

