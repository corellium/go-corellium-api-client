/*
Corellium API

REST API to manage your virtual devices.

API version: 5.6.0-19122
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"fmt"
)

// ConfigResponseMaintenance - Denotes whether the server is undergoing maintenance
type ConfigResponseMaintenance struct {
	Maintenance *Maintenance
	Bool *bool
}

// MaintenanceAsConfigResponseMaintenance is a convenience function that returns Maintenance wrapped in ConfigResponseMaintenance
func MaintenanceAsConfigResponseMaintenance(v *Maintenance) ConfigResponseMaintenance {
	return ConfigResponseMaintenance{
		Maintenance: v,
	}
}

// boolAsConfigResponseMaintenance is a convenience function that returns bool wrapped in ConfigResponseMaintenance
func BoolAsConfigResponseMaintenance(v *bool) ConfigResponseMaintenance {
	return ConfigResponseMaintenance{
		Bool: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConfigResponseMaintenance) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into Maintenance
	err = newStrictDecoder(data).Decode(&dst.Maintenance)
	if err == nil {
		jsonMaintenance, _ := json.Marshal(dst.Maintenance)
		if string(jsonMaintenance) == "{}" { // empty struct
			dst.Maintenance = nil
		} else {
			match++
		}
	} else {
		dst.Maintenance = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Maintenance = nil
		dst.Bool = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ConfigResponseMaintenance)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ConfigResponseMaintenance)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConfigResponseMaintenance) MarshalJSON() ([]byte, error) {
	if src.Maintenance != nil {
		return json.Marshal(&src.Maintenance)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConfigResponseMaintenance) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Maintenance != nil {
		return obj.Maintenance
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

type NullableConfigResponseMaintenance struct {
	value *ConfigResponseMaintenance
	isSet bool
}

func (v NullableConfigResponseMaintenance) Get() *ConfigResponseMaintenance {
	return v.value
}

func (v *NullableConfigResponseMaintenance) Set(val *ConfigResponseMaintenance) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigResponseMaintenance) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigResponseMaintenance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigResponseMaintenance(val *ConfigResponseMaintenance) *NullableConfigResponseMaintenance {
	return &NullableConfigResponseMaintenance{value: val, isSet: true}
}

func (v NullableConfigResponseMaintenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigResponseMaintenance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


