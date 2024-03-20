/*
Corellium API

REST API to manage your virtual devices.

API version: 6.1.0-20784
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the InstanceBootOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceBootOptions{}

// InstanceBootOptions 
type InstanceBootOptions struct {
	// 
	BootArgs NullableString `json:"bootArgs,omitempty"`
	// 
	RestoreBootArgs NullableString `json:"restoreBootArgs,omitempty"`
	// Boot udid
	Udid NullableString `json:"udid,omitempty"`
	// Assigned ECID
	Ecid NullableString `json:"ecid,omitempty"`
	// Random seed to provide to boot if any
	RandomSeed NullableString `json:"randomSeed,omitempty"`
	// Enable PAC
	Pac NullableBool `json:"pac,omitempty"`
	// Enable APRR
	Aprr NullableBool `json:"aprr,omitempty"`
	// 
	AdditionalTags []InstanceBootOptionsAdditionalTag `json:"additionalTags,omitempty"`
}

// NewInstanceBootOptions instantiates a new InstanceBootOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceBootOptions() *InstanceBootOptions {
	this := InstanceBootOptions{}
	return &this
}

// NewInstanceBootOptionsWithDefaults instantiates a new InstanceBootOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceBootOptionsWithDefaults() *InstanceBootOptions {
	this := InstanceBootOptions{}
	return &this
}

// GetBootArgs returns the BootArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetBootArgs() string {
	if o == nil || IsNil(o.BootArgs.Get()) {
		var ret string
		return ret
	}
	return *o.BootArgs.Get()
}

// GetBootArgsOk returns a tuple with the BootArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetBootArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BootArgs.Get(), o.BootArgs.IsSet()
}

// HasBootArgs returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasBootArgs() bool {
	if o != nil && o.BootArgs.IsSet() {
		return true
	}

	return false
}

// SetBootArgs gets a reference to the given NullableString and assigns it to the BootArgs field.
func (o *InstanceBootOptions) SetBootArgs(v string) {
	o.BootArgs.Set(&v)
}
// SetBootArgsNil sets the value for BootArgs to be an explicit nil
func (o *InstanceBootOptions) SetBootArgsNil() {
	o.BootArgs.Set(nil)
}

// UnsetBootArgs ensures that no value is present for BootArgs, not even an explicit nil
func (o *InstanceBootOptions) UnsetBootArgs() {
	o.BootArgs.Unset()
}

// GetRestoreBootArgs returns the RestoreBootArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetRestoreBootArgs() string {
	if o == nil || IsNil(o.RestoreBootArgs.Get()) {
		var ret string
		return ret
	}
	return *o.RestoreBootArgs.Get()
}

// GetRestoreBootArgsOk returns a tuple with the RestoreBootArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetRestoreBootArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestoreBootArgs.Get(), o.RestoreBootArgs.IsSet()
}

// HasRestoreBootArgs returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasRestoreBootArgs() bool {
	if o != nil && o.RestoreBootArgs.IsSet() {
		return true
	}

	return false
}

// SetRestoreBootArgs gets a reference to the given NullableString and assigns it to the RestoreBootArgs field.
func (o *InstanceBootOptions) SetRestoreBootArgs(v string) {
	o.RestoreBootArgs.Set(&v)
}
// SetRestoreBootArgsNil sets the value for RestoreBootArgs to be an explicit nil
func (o *InstanceBootOptions) SetRestoreBootArgsNil() {
	o.RestoreBootArgs.Set(nil)
}

// UnsetRestoreBootArgs ensures that no value is present for RestoreBootArgs, not even an explicit nil
func (o *InstanceBootOptions) UnsetRestoreBootArgs() {
	o.RestoreBootArgs.Unset()
}

// GetUdid returns the Udid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetUdid() string {
	if o == nil || IsNil(o.Udid.Get()) {
		var ret string
		return ret
	}
	return *o.Udid.Get()
}

// GetUdidOk returns a tuple with the Udid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetUdidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Udid.Get(), o.Udid.IsSet()
}

// HasUdid returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasUdid() bool {
	if o != nil && o.Udid.IsSet() {
		return true
	}

	return false
}

// SetUdid gets a reference to the given NullableString and assigns it to the Udid field.
func (o *InstanceBootOptions) SetUdid(v string) {
	o.Udid.Set(&v)
}
// SetUdidNil sets the value for Udid to be an explicit nil
func (o *InstanceBootOptions) SetUdidNil() {
	o.Udid.Set(nil)
}

// UnsetUdid ensures that no value is present for Udid, not even an explicit nil
func (o *InstanceBootOptions) UnsetUdid() {
	o.Udid.Unset()
}

// GetEcid returns the Ecid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetEcid() string {
	if o == nil || IsNil(o.Ecid.Get()) {
		var ret string
		return ret
	}
	return *o.Ecid.Get()
}

// GetEcidOk returns a tuple with the Ecid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetEcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ecid.Get(), o.Ecid.IsSet()
}

// HasEcid returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasEcid() bool {
	if o != nil && o.Ecid.IsSet() {
		return true
	}

	return false
}

// SetEcid gets a reference to the given NullableString and assigns it to the Ecid field.
func (o *InstanceBootOptions) SetEcid(v string) {
	o.Ecid.Set(&v)
}
// SetEcidNil sets the value for Ecid to be an explicit nil
func (o *InstanceBootOptions) SetEcidNil() {
	o.Ecid.Set(nil)
}

// UnsetEcid ensures that no value is present for Ecid, not even an explicit nil
func (o *InstanceBootOptions) UnsetEcid() {
	o.Ecid.Unset()
}

// GetRandomSeed returns the RandomSeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetRandomSeed() string {
	if o == nil || IsNil(o.RandomSeed.Get()) {
		var ret string
		return ret
	}
	return *o.RandomSeed.Get()
}

// GetRandomSeedOk returns a tuple with the RandomSeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetRandomSeedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RandomSeed.Get(), o.RandomSeed.IsSet()
}

// HasRandomSeed returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasRandomSeed() bool {
	if o != nil && o.RandomSeed.IsSet() {
		return true
	}

	return false
}

// SetRandomSeed gets a reference to the given NullableString and assigns it to the RandomSeed field.
func (o *InstanceBootOptions) SetRandomSeed(v string) {
	o.RandomSeed.Set(&v)
}
// SetRandomSeedNil sets the value for RandomSeed to be an explicit nil
func (o *InstanceBootOptions) SetRandomSeedNil() {
	o.RandomSeed.Set(nil)
}

// UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
func (o *InstanceBootOptions) UnsetRandomSeed() {
	o.RandomSeed.Unset()
}

// GetPac returns the Pac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetPac() bool {
	if o == nil || IsNil(o.Pac.Get()) {
		var ret bool
		return ret
	}
	return *o.Pac.Get()
}

// GetPacOk returns a tuple with the Pac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetPacOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pac.Get(), o.Pac.IsSet()
}

// HasPac returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasPac() bool {
	if o != nil && o.Pac.IsSet() {
		return true
	}

	return false
}

// SetPac gets a reference to the given NullableBool and assigns it to the Pac field.
func (o *InstanceBootOptions) SetPac(v bool) {
	o.Pac.Set(&v)
}
// SetPacNil sets the value for Pac to be an explicit nil
func (o *InstanceBootOptions) SetPacNil() {
	o.Pac.Set(nil)
}

// UnsetPac ensures that no value is present for Pac, not even an explicit nil
func (o *InstanceBootOptions) UnsetPac() {
	o.Pac.Unset()
}

// GetAprr returns the Aprr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetAprr() bool {
	if o == nil || IsNil(o.Aprr.Get()) {
		var ret bool
		return ret
	}
	return *o.Aprr.Get()
}

// GetAprrOk returns a tuple with the Aprr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetAprrOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aprr.Get(), o.Aprr.IsSet()
}

// HasAprr returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasAprr() bool {
	if o != nil && o.Aprr.IsSet() {
		return true
	}

	return false
}

// SetAprr gets a reference to the given NullableBool and assigns it to the Aprr field.
func (o *InstanceBootOptions) SetAprr(v bool) {
	o.Aprr.Set(&v)
}
// SetAprrNil sets the value for Aprr to be an explicit nil
func (o *InstanceBootOptions) SetAprrNil() {
	o.Aprr.Set(nil)
}

// UnsetAprr ensures that no value is present for Aprr, not even an explicit nil
func (o *InstanceBootOptions) UnsetAprr() {
	o.Aprr.Unset()
}

// GetAdditionalTags returns the AdditionalTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceBootOptions) GetAdditionalTags() []InstanceBootOptionsAdditionalTag {
	if o == nil {
		var ret []InstanceBootOptionsAdditionalTag
		return ret
	}
	return o.AdditionalTags
}

// GetAdditionalTagsOk returns a tuple with the AdditionalTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceBootOptions) GetAdditionalTagsOk() ([]InstanceBootOptionsAdditionalTag, bool) {
	if o == nil || IsNil(o.AdditionalTags) {
		return nil, false
	}
	return o.AdditionalTags, true
}

// HasAdditionalTags returns a boolean if a field has been set.
func (o *InstanceBootOptions) HasAdditionalTags() bool {
	if o != nil && IsNil(o.AdditionalTags) {
		return true
	}

	return false
}

// SetAdditionalTags gets a reference to the given []InstanceBootOptionsAdditionalTag and assigns it to the AdditionalTags field.
func (o *InstanceBootOptions) SetAdditionalTags(v []InstanceBootOptionsAdditionalTag) {
	o.AdditionalTags = v
}

func (o InstanceBootOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceBootOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BootArgs.IsSet() {
		toSerialize["bootArgs"] = o.BootArgs.Get()
	}
	if o.RestoreBootArgs.IsSet() {
		toSerialize["restoreBootArgs"] = o.RestoreBootArgs.Get()
	}
	if o.Udid.IsSet() {
		toSerialize["udid"] = o.Udid.Get()
	}
	if o.Ecid.IsSet() {
		toSerialize["ecid"] = o.Ecid.Get()
	}
	if o.RandomSeed.IsSet() {
		toSerialize["randomSeed"] = o.RandomSeed.Get()
	}
	if o.Pac.IsSet() {
		toSerialize["pac"] = o.Pac.Get()
	}
	if o.Aprr.IsSet() {
		toSerialize["aprr"] = o.Aprr.Get()
	}
	if o.AdditionalTags != nil {
		toSerialize["additionalTags"] = o.AdditionalTags
	}
	return toSerialize, nil
}

type NullableInstanceBootOptions struct {
	value *InstanceBootOptions
	isSet bool
}

func (v NullableInstanceBootOptions) Get() *InstanceBootOptions {
	return v.value
}

func (v *NullableInstanceBootOptions) Set(val *InstanceBootOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceBootOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceBootOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceBootOptions(val *InstanceBootOptions) *NullableInstanceBootOptions {
	return &NullableInstanceBootOptions{value: val, isSet: true}
}

func (v NullableInstanceBootOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceBootOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


