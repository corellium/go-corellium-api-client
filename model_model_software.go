/*
Corellium API

REST API to manage your virtual devices.

API version: 7.3.0-27831
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the ModelSoftware type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSoftware{}

// ModelSoftware 
type ModelSoftware struct {
	// 
	Name NullableString `json:"name,omitempty"`
	// 
	BoardConfig NullableString `json:"boardConfig,omitempty"`
	// 
	Platform NullableString `json:"platform,omitempty"`
	// 
	Cpid NullableFloat32 `json:"cpid,omitempty"`
	// 
	Bdid NullableFloat32 `json:"bdid,omitempty"`
	// 
	Firmwares []Firmware `json:"firmwares,omitempty"`
}

// NewModelSoftware instantiates a new ModelSoftware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSoftware() *ModelSoftware {
	this := ModelSoftware{}
	return &this
}

// NewModelSoftwareWithDefaults instantiates a new ModelSoftware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSoftwareWithDefaults() *ModelSoftware {
	this := ModelSoftware{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSoftware) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSoftware) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelSoftware) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelSoftware) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelSoftware) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelSoftware) UnsetName() {
	o.Name.Unset()
}

// GetBoardConfig returns the BoardConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSoftware) GetBoardConfig() string {
	if o == nil || IsNil(o.BoardConfig.Get()) {
		var ret string
		return ret
	}
	return *o.BoardConfig.Get()
}

// GetBoardConfigOk returns a tuple with the BoardConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSoftware) GetBoardConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardConfig.Get(), o.BoardConfig.IsSet()
}

// HasBoardConfig returns a boolean if a field has been set.
func (o *ModelSoftware) HasBoardConfig() bool {
	if o != nil && o.BoardConfig.IsSet() {
		return true
	}

	return false
}

// SetBoardConfig gets a reference to the given NullableString and assigns it to the BoardConfig field.
func (o *ModelSoftware) SetBoardConfig(v string) {
	o.BoardConfig.Set(&v)
}
// SetBoardConfigNil sets the value for BoardConfig to be an explicit nil
func (o *ModelSoftware) SetBoardConfigNil() {
	o.BoardConfig.Set(nil)
}

// UnsetBoardConfig ensures that no value is present for BoardConfig, not even an explicit nil
func (o *ModelSoftware) UnsetBoardConfig() {
	o.BoardConfig.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSoftware) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSoftware) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *ModelSoftware) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *ModelSoftware) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *ModelSoftware) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *ModelSoftware) UnsetPlatform() {
	o.Platform.Unset()
}

// GetCpid returns the Cpid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSoftware) GetCpid() float32 {
	if o == nil || IsNil(o.Cpid.Get()) {
		var ret float32
		return ret
	}
	return *o.Cpid.Get()
}

// GetCpidOk returns a tuple with the Cpid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSoftware) GetCpidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpid.Get(), o.Cpid.IsSet()
}

// HasCpid returns a boolean if a field has been set.
func (o *ModelSoftware) HasCpid() bool {
	if o != nil && o.Cpid.IsSet() {
		return true
	}

	return false
}

// SetCpid gets a reference to the given NullableFloat32 and assigns it to the Cpid field.
func (o *ModelSoftware) SetCpid(v float32) {
	o.Cpid.Set(&v)
}
// SetCpidNil sets the value for Cpid to be an explicit nil
func (o *ModelSoftware) SetCpidNil() {
	o.Cpid.Set(nil)
}

// UnsetCpid ensures that no value is present for Cpid, not even an explicit nil
func (o *ModelSoftware) UnsetCpid() {
	o.Cpid.Unset()
}

// GetBdid returns the Bdid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSoftware) GetBdid() float32 {
	if o == nil || IsNil(o.Bdid.Get()) {
		var ret float32
		return ret
	}
	return *o.Bdid.Get()
}

// GetBdidOk returns a tuple with the Bdid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSoftware) GetBdidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bdid.Get(), o.Bdid.IsSet()
}

// HasBdid returns a boolean if a field has been set.
func (o *ModelSoftware) HasBdid() bool {
	if o != nil && o.Bdid.IsSet() {
		return true
	}

	return false
}

// SetBdid gets a reference to the given NullableFloat32 and assigns it to the Bdid field.
func (o *ModelSoftware) SetBdid(v float32) {
	o.Bdid.Set(&v)
}
// SetBdidNil sets the value for Bdid to be an explicit nil
func (o *ModelSoftware) SetBdidNil() {
	o.Bdid.Set(nil)
}

// UnsetBdid ensures that no value is present for Bdid, not even an explicit nil
func (o *ModelSoftware) UnsetBdid() {
	o.Bdid.Unset()
}

// GetFirmwares returns the Firmwares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSoftware) GetFirmwares() []Firmware {
	if o == nil {
		var ret []Firmware
		return ret
	}
	return o.Firmwares
}

// GetFirmwaresOk returns a tuple with the Firmwares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSoftware) GetFirmwaresOk() ([]Firmware, bool) {
	if o == nil || IsNil(o.Firmwares) {
		return nil, false
	}
	return o.Firmwares, true
}

// HasFirmwares returns a boolean if a field has been set.
func (o *ModelSoftware) HasFirmwares() bool {
	if o != nil && IsNil(o.Firmwares) {
		return true
	}

	return false
}

// SetFirmwares gets a reference to the given []Firmware and assigns it to the Firmwares field.
func (o *ModelSoftware) SetFirmwares(v []Firmware) {
	o.Firmwares = v
}

func (o ModelSoftware) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSoftware) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.BoardConfig.IsSet() {
		toSerialize["boardConfig"] = o.BoardConfig.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Cpid.IsSet() {
		toSerialize["cpid"] = o.Cpid.Get()
	}
	if o.Bdid.IsSet() {
		toSerialize["bdid"] = o.Bdid.Get()
	}
	if o.Firmwares != nil {
		toSerialize["firmwares"] = o.Firmwares
	}
	return toSerialize, nil
}

type NullableModelSoftware struct {
	value *ModelSoftware
	isSet bool
}

func (v NullableModelSoftware) Get() *ModelSoftware {
	return v.value
}

func (v *NullableModelSoftware) Set(val *ModelSoftware) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSoftware) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSoftware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSoftware(val *ModelSoftware) *NullableModelSoftware {
	return &NullableModelSoftware{value: val, isSet: true}
}

func (v NullableModelSoftware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSoftware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


