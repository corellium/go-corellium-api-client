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

// checks if the Model type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model{}

// Model 
type Model struct {
	// 
	Type string `json:"type"`
	// 
	Name string `json:"name"`
	// 
	Flavor string `json:"flavor"`
	// 
	Description NullableString `json:"description,omitempty"`
	// 
	Model string `json:"model"`
	// 
	BoardConfig NullableString `json:"boardConfig,omitempty"`
	// 
	Platform NullableString `json:"platform,omitempty"`
	// 
	Cpid NullableFloat32 `json:"cpid,omitempty"`
	// 
	Bdid NullableFloat32 `json:"bdid,omitempty"`
	// 
	Peripherals NullableBool `json:"peripherals,omitempty"`
}

// NewModel instantiates a new Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel(type_ string, name string, flavor string, model string) *Model {
	this := Model{}
	this.Type = type_
	this.Name = name
	this.Flavor = flavor
	this.Model = model
	return &this
}

// NewModelWithDefaults instantiates a new Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWithDefaults() *Model {
	this := Model{}
	return &this
}

// GetType returns the Type field value
func (o *Model) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Model) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Model) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *Model) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Model) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Model) SetName(v string) {
	o.Name = v
}

// GetFlavor returns the Flavor field value
func (o *Model) GetFlavor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value
// and a boolean to check if the value has been set.
func (o *Model) GetFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flavor, true
}

// SetFlavor sets field value
func (o *Model) SetFlavor(v string) {
	o.Flavor = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Model) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Model) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Model) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Model) UnsetDescription() {
	o.Description.Unset()
}

// GetModel returns the Model field value
func (o *Model) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *Model) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *Model) SetModel(v string) {
	o.Model = v
}

// GetBoardConfig returns the BoardConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model) GetBoardConfig() string {
	if o == nil || IsNil(o.BoardConfig.Get()) {
		var ret string
		return ret
	}
	return *o.BoardConfig.Get()
}

// GetBoardConfigOk returns a tuple with the BoardConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetBoardConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardConfig.Get(), o.BoardConfig.IsSet()
}

// HasBoardConfig returns a boolean if a field has been set.
func (o *Model) HasBoardConfig() bool {
	if o != nil && o.BoardConfig.IsSet() {
		return true
	}

	return false
}

// SetBoardConfig gets a reference to the given NullableString and assigns it to the BoardConfig field.
func (o *Model) SetBoardConfig(v string) {
	o.BoardConfig.Set(&v)
}
// SetBoardConfigNil sets the value for BoardConfig to be an explicit nil
func (o *Model) SetBoardConfigNil() {
	o.BoardConfig.Set(nil)
}

// UnsetBoardConfig ensures that no value is present for BoardConfig, not even an explicit nil
func (o *Model) UnsetBoardConfig() {
	o.BoardConfig.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *Model) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *Model) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *Model) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *Model) UnsetPlatform() {
	o.Platform.Unset()
}

// GetCpid returns the Cpid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model) GetCpid() float32 {
	if o == nil || IsNil(o.Cpid.Get()) {
		var ret float32
		return ret
	}
	return *o.Cpid.Get()
}

// GetCpidOk returns a tuple with the Cpid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetCpidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpid.Get(), o.Cpid.IsSet()
}

// HasCpid returns a boolean if a field has been set.
func (o *Model) HasCpid() bool {
	if o != nil && o.Cpid.IsSet() {
		return true
	}

	return false
}

// SetCpid gets a reference to the given NullableFloat32 and assigns it to the Cpid field.
func (o *Model) SetCpid(v float32) {
	o.Cpid.Set(&v)
}
// SetCpidNil sets the value for Cpid to be an explicit nil
func (o *Model) SetCpidNil() {
	o.Cpid.Set(nil)
}

// UnsetCpid ensures that no value is present for Cpid, not even an explicit nil
func (o *Model) UnsetCpid() {
	o.Cpid.Unset()
}

// GetBdid returns the Bdid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model) GetBdid() float32 {
	if o == nil || IsNil(o.Bdid.Get()) {
		var ret float32
		return ret
	}
	return *o.Bdid.Get()
}

// GetBdidOk returns a tuple with the Bdid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetBdidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bdid.Get(), o.Bdid.IsSet()
}

// HasBdid returns a boolean if a field has been set.
func (o *Model) HasBdid() bool {
	if o != nil && o.Bdid.IsSet() {
		return true
	}

	return false
}

// SetBdid gets a reference to the given NullableFloat32 and assigns it to the Bdid field.
func (o *Model) SetBdid(v float32) {
	o.Bdid.Set(&v)
}
// SetBdidNil sets the value for Bdid to be an explicit nil
func (o *Model) SetBdidNil() {
	o.Bdid.Set(nil)
}

// UnsetBdid ensures that no value is present for Bdid, not even an explicit nil
func (o *Model) UnsetBdid() {
	o.Bdid.Unset()
}

// GetPeripherals returns the Peripherals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model) GetPeripherals() bool {
	if o == nil || IsNil(o.Peripherals.Get()) {
		var ret bool
		return ret
	}
	return *o.Peripherals.Get()
}

// GetPeripheralsOk returns a tuple with the Peripherals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetPeripheralsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peripherals.Get(), o.Peripherals.IsSet()
}

// HasPeripherals returns a boolean if a field has been set.
func (o *Model) HasPeripherals() bool {
	if o != nil && o.Peripherals.IsSet() {
		return true
	}

	return false
}

// SetPeripherals gets a reference to the given NullableBool and assigns it to the Peripherals field.
func (o *Model) SetPeripherals(v bool) {
	o.Peripherals.Set(&v)
}
// SetPeripheralsNil sets the value for Peripherals to be an explicit nil
func (o *Model) SetPeripheralsNil() {
	o.Peripherals.Set(nil)
}

// UnsetPeripherals ensures that no value is present for Peripherals, not even an explicit nil
func (o *Model) UnsetPeripherals() {
	o.Peripherals.Unset()
}

func (o Model) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["flavor"] = o.Flavor
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["model"] = o.Model
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
	if o.Peripherals.IsSet() {
		toSerialize["peripherals"] = o.Peripherals.Get()
	}
	return toSerialize, nil
}

type NullableModel struct {
	value *Model
	isSet bool
}

func (v NullableModel) Get() *Model {
	return v.value
}

func (v *NullableModel) Set(val *Model) {
	v.value = val
	v.isSet = true
}

func (v NullableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel(val *Model) *NullableModel {
	return &NullableModel{value: val, isSet: true}
}

func (v NullableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


