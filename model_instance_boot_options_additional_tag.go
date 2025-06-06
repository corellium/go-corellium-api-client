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

// InstanceBootOptionsAdditionalTag ### Instance Boot Option * kalloc: Enable kalloc/kfree trace access via GDB (Enterprise only) * gpu: Enable cloud GPU acceleration (Extra costs incurred, cloud only) * no-keyboard: Enable keyboard passthrough from web interface * nodevmode: Disable developer mode on iOS16 and greater * sep-cons-ext: Patch SEPOS to print debug messages to console * iboot-jailbreak: Patch iBoot to disable signature checks * llb-jailbreak: Patch LLB to disable signature checks * rom-jailbreak: Patch BootROM to disable signature checks
type InstanceBootOptionsAdditionalTag string

// List of InstanceBootOptionsAdditionalTag
const (
	KALLOC InstanceBootOptionsAdditionalTag = "kalloc"
	GPU InstanceBootOptionsAdditionalTag = "gpu"
	NO_KEYBOARD InstanceBootOptionsAdditionalTag = "no-keyboard"
	NODEVMODE InstanceBootOptionsAdditionalTag = "nodevmode"
	SEP_CONS_EXT InstanceBootOptionsAdditionalTag = "sep-cons-ext"
	IBOOT_JAILBREAK InstanceBootOptionsAdditionalTag = "iboot-jailbreak"
	LLB_JAILBREAK InstanceBootOptionsAdditionalTag = "llb-jailbreak"
	ROM_JAILBREAK InstanceBootOptionsAdditionalTag = "rom-jailbreak"
)

// All allowed values of InstanceBootOptionsAdditionalTag enum
var AllowedInstanceBootOptionsAdditionalTagEnumValues = []InstanceBootOptionsAdditionalTag{
	"kalloc",
	"gpu",
	"no-keyboard",
	"nodevmode",
	"sep-cons-ext",
	"iboot-jailbreak",
	"llb-jailbreak",
	"rom-jailbreak",
}

func (v *InstanceBootOptionsAdditionalTag) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceBootOptionsAdditionalTag(value)
	for _, existing := range AllowedInstanceBootOptionsAdditionalTagEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceBootOptionsAdditionalTag", value)
}

// NewInstanceBootOptionsAdditionalTagFromValue returns a pointer to a valid InstanceBootOptionsAdditionalTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceBootOptionsAdditionalTagFromValue(v string) (*InstanceBootOptionsAdditionalTag, error) {
	ev := InstanceBootOptionsAdditionalTag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceBootOptionsAdditionalTag: valid values are %v", v, AllowedInstanceBootOptionsAdditionalTagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceBootOptionsAdditionalTag) IsValid() bool {
	for _, existing := range AllowedInstanceBootOptionsAdditionalTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceBootOptionsAdditionalTag value
func (v InstanceBootOptionsAdditionalTag) Ptr() *InstanceBootOptionsAdditionalTag {
	return &v
}

type NullableInstanceBootOptionsAdditionalTag struct {
	value *InstanceBootOptionsAdditionalTag
	isSet bool
}

func (v NullableInstanceBootOptionsAdditionalTag) Get() *InstanceBootOptionsAdditionalTag {
	return v.value
}

func (v *NullableInstanceBootOptionsAdditionalTag) Set(val *InstanceBootOptionsAdditionalTag) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceBootOptionsAdditionalTag) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceBootOptionsAdditionalTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceBootOptionsAdditionalTag(val *InstanceBootOptionsAdditionalTag) *NullableInstanceBootOptionsAdditionalTag {
	return &NullableInstanceBootOptionsAdditionalTag{value: val, isSet: true}
}

func (v NullableInstanceBootOptionsAdditionalTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceBootOptionsAdditionalTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

