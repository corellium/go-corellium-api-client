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

// checks if the DomainOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainOptions{}

// DomainOptions 
type DomainOptions struct {
	// if true, totp is required
	TotpRequired NullableBool `json:"totpRequired,omitempty"`
	TrialExtension *TrialExtension `json:"trialExtension,omitempty"`
	SnapshotPermissions *SnapshotPermissions `json:"snapshotPermissions,omitempty"`
}

// NewDomainOptions instantiates a new DomainOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainOptions() *DomainOptions {
	this := DomainOptions{}
	return &this
}

// NewDomainOptionsWithDefaults instantiates a new DomainOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainOptionsWithDefaults() *DomainOptions {
	this := DomainOptions{}
	return &this
}

// GetTotpRequired returns the TotpRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainOptions) GetTotpRequired() bool {
	if o == nil || IsNil(o.TotpRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.TotpRequired.Get()
}

// GetTotpRequiredOk returns a tuple with the TotpRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainOptions) GetTotpRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotpRequired.Get(), o.TotpRequired.IsSet()
}

// HasTotpRequired returns a boolean if a field has been set.
func (o *DomainOptions) HasTotpRequired() bool {
	if o != nil && o.TotpRequired.IsSet() {
		return true
	}

	return false
}

// SetTotpRequired gets a reference to the given NullableBool and assigns it to the TotpRequired field.
func (o *DomainOptions) SetTotpRequired(v bool) {
	o.TotpRequired.Set(&v)
}
// SetTotpRequiredNil sets the value for TotpRequired to be an explicit nil
func (o *DomainOptions) SetTotpRequiredNil() {
	o.TotpRequired.Set(nil)
}

// UnsetTotpRequired ensures that no value is present for TotpRequired, not even an explicit nil
func (o *DomainOptions) UnsetTotpRequired() {
	o.TotpRequired.Unset()
}

// GetTrialExtension returns the TrialExtension field value if set, zero value otherwise.
func (o *DomainOptions) GetTrialExtension() TrialExtension {
	if o == nil || IsNil(o.TrialExtension) {
		var ret TrialExtension
		return ret
	}
	return *o.TrialExtension
}

// GetTrialExtensionOk returns a tuple with the TrialExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainOptions) GetTrialExtensionOk() (*TrialExtension, bool) {
	if o == nil || IsNil(o.TrialExtension) {
		return nil, false
	}
	return o.TrialExtension, true
}

// HasTrialExtension returns a boolean if a field has been set.
func (o *DomainOptions) HasTrialExtension() bool {
	if o != nil && !IsNil(o.TrialExtension) {
		return true
	}

	return false
}

// SetTrialExtension gets a reference to the given TrialExtension and assigns it to the TrialExtension field.
func (o *DomainOptions) SetTrialExtension(v TrialExtension) {
	o.TrialExtension = &v
}

// GetSnapshotPermissions returns the SnapshotPermissions field value if set, zero value otherwise.
func (o *DomainOptions) GetSnapshotPermissions() SnapshotPermissions {
	if o == nil || IsNil(o.SnapshotPermissions) {
		var ret SnapshotPermissions
		return ret
	}
	return *o.SnapshotPermissions
}

// GetSnapshotPermissionsOk returns a tuple with the SnapshotPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainOptions) GetSnapshotPermissionsOk() (*SnapshotPermissions, bool) {
	if o == nil || IsNil(o.SnapshotPermissions) {
		return nil, false
	}
	return o.SnapshotPermissions, true
}

// HasSnapshotPermissions returns a boolean if a field has been set.
func (o *DomainOptions) HasSnapshotPermissions() bool {
	if o != nil && !IsNil(o.SnapshotPermissions) {
		return true
	}

	return false
}

// SetSnapshotPermissions gets a reference to the given SnapshotPermissions and assigns it to the SnapshotPermissions field.
func (o *DomainOptions) SetSnapshotPermissions(v SnapshotPermissions) {
	o.SnapshotPermissions = &v
}

func (o DomainOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TotpRequired.IsSet() {
		toSerialize["totpRequired"] = o.TotpRequired.Get()
	}
	if !IsNil(o.TrialExtension) {
		toSerialize["trialExtension"] = o.TrialExtension
	}
	if !IsNil(o.SnapshotPermissions) {
		toSerialize["snapshotPermissions"] = o.SnapshotPermissions
	}
	return toSerialize, nil
}

type NullableDomainOptions struct {
	value *DomainOptions
	isSet bool
}

func (v NullableDomainOptions) Get() *DomainOptions {
	return v.value
}

func (v *NullableDomainOptions) Set(val *DomainOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainOptions(val *DomainOptions) *NullableDomainOptions {
	return &NullableDomainOptions{value: val, isSet: true}
}

func (v NullableDomainOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


