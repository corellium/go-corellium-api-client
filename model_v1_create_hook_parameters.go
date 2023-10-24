/*
Corellium API

REST API to manage your virtual devices.

API version: 5.5.0-18750
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the V1CreateHookParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CreateHookParameters{}

// V1CreateHookParameters 
type V1CreateHookParameters struct {
	// Label
	Label string `json:"label"`
	// Address
	Address string `json:"address"`
	// Patch
	Patch string `json:"patch"`
	// Patch Type
	PatchType string `json:"patchType"`
}

// NewV1CreateHookParameters instantiates a new V1CreateHookParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateHookParameters(label string, address string, patch string, patchType string) *V1CreateHookParameters {
	this := V1CreateHookParameters{}
	this.Label = label
	this.Address = address
	this.Patch = patch
	this.PatchType = patchType
	return &this
}

// NewV1CreateHookParametersWithDefaults instantiates a new V1CreateHookParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateHookParametersWithDefaults() *V1CreateHookParameters {
	this := V1CreateHookParameters{}
	return &this
}

// GetLabel returns the Label field value
func (o *V1CreateHookParameters) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *V1CreateHookParameters) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *V1CreateHookParameters) SetLabel(v string) {
	o.Label = v
}

// GetAddress returns the Address field value
func (o *V1CreateHookParameters) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *V1CreateHookParameters) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *V1CreateHookParameters) SetAddress(v string) {
	o.Address = v
}

// GetPatch returns the Patch field value
func (o *V1CreateHookParameters) GetPatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Patch
}

// GetPatchOk returns a tuple with the Patch field value
// and a boolean to check if the value has been set.
func (o *V1CreateHookParameters) GetPatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patch, true
}

// SetPatch sets field value
func (o *V1CreateHookParameters) SetPatch(v string) {
	o.Patch = v
}

// GetPatchType returns the PatchType field value
func (o *V1CreateHookParameters) GetPatchType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatchType
}

// GetPatchTypeOk returns a tuple with the PatchType field value
// and a boolean to check if the value has been set.
func (o *V1CreateHookParameters) GetPatchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatchType, true
}

// SetPatchType sets field value
func (o *V1CreateHookParameters) SetPatchType(v string) {
	o.PatchType = v
}

func (o V1CreateHookParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CreateHookParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["address"] = o.Address
	toSerialize["patch"] = o.Patch
	toSerialize["patchType"] = o.PatchType
	return toSerialize, nil
}

type NullableV1CreateHookParameters struct {
	value *V1CreateHookParameters
	isSet bool
}

func (v NullableV1CreateHookParameters) Get() *V1CreateHookParameters {
	return v.value
}

func (v *NullableV1CreateHookParameters) Set(val *V1CreateHookParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateHookParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateHookParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateHookParameters(val *V1CreateHookParameters) *NullableV1CreateHookParameters {
	return &NullableV1CreateHookParameters{value: val, isSet: true}
}

func (v NullableV1CreateHookParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateHookParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


