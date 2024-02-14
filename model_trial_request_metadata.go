/*
Corellium API

REST API to manage your virtual devices.

API version: 6.0.0-20323
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the TrialRequestMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrialRequestMetadata{}

// TrialRequestMetadata 
type TrialRequestMetadata struct {
	// 
	Name NullableString `json:"name,omitempty"`
	// provided company name
	Company NullableString `json:"company,omitempty"`
	// option from sign up form, interesting using for malware
	Malware NullableBool `json:"malware,omitempty"`
	// option from sign up form, interesting using for research
	Research NullableBool `json:"research,omitempty"`
	// option from sign up form, interesting using for pentesting
	Pentest NullableBool `json:"pentest,omitempty"`
	// option from sign up form, interesting using for other area added here
	Other NullableString `json:"other,omitempty"`
}

// NewTrialRequestMetadata instantiates a new TrialRequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialRequestMetadata() *TrialRequestMetadata {
	this := TrialRequestMetadata{}
	return &this
}

// NewTrialRequestMetadataWithDefaults instantiates a new TrialRequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialRequestMetadataWithDefaults() *TrialRequestMetadata {
	this := TrialRequestMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestMetadata) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TrialRequestMetadata) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TrialRequestMetadata) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TrialRequestMetadata) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TrialRequestMetadata) UnsetName() {
	o.Name.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestMetadata) GetCompany() string {
	if o == nil || IsNil(o.Company.Get()) {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestMetadata) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *TrialRequestMetadata) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *TrialRequestMetadata) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *TrialRequestMetadata) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *TrialRequestMetadata) UnsetCompany() {
	o.Company.Unset()
}

// GetMalware returns the Malware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestMetadata) GetMalware() bool {
	if o == nil || IsNil(o.Malware.Get()) {
		var ret bool
		return ret
	}
	return *o.Malware.Get()
}

// GetMalwareOk returns a tuple with the Malware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestMetadata) GetMalwareOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Malware.Get(), o.Malware.IsSet()
}

// HasMalware returns a boolean if a field has been set.
func (o *TrialRequestMetadata) HasMalware() bool {
	if o != nil && o.Malware.IsSet() {
		return true
	}

	return false
}

// SetMalware gets a reference to the given NullableBool and assigns it to the Malware field.
func (o *TrialRequestMetadata) SetMalware(v bool) {
	o.Malware.Set(&v)
}
// SetMalwareNil sets the value for Malware to be an explicit nil
func (o *TrialRequestMetadata) SetMalwareNil() {
	o.Malware.Set(nil)
}

// UnsetMalware ensures that no value is present for Malware, not even an explicit nil
func (o *TrialRequestMetadata) UnsetMalware() {
	o.Malware.Unset()
}

// GetResearch returns the Research field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestMetadata) GetResearch() bool {
	if o == nil || IsNil(o.Research.Get()) {
		var ret bool
		return ret
	}
	return *o.Research.Get()
}

// GetResearchOk returns a tuple with the Research field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestMetadata) GetResearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Research.Get(), o.Research.IsSet()
}

// HasResearch returns a boolean if a field has been set.
func (o *TrialRequestMetadata) HasResearch() bool {
	if o != nil && o.Research.IsSet() {
		return true
	}

	return false
}

// SetResearch gets a reference to the given NullableBool and assigns it to the Research field.
func (o *TrialRequestMetadata) SetResearch(v bool) {
	o.Research.Set(&v)
}
// SetResearchNil sets the value for Research to be an explicit nil
func (o *TrialRequestMetadata) SetResearchNil() {
	o.Research.Set(nil)
}

// UnsetResearch ensures that no value is present for Research, not even an explicit nil
func (o *TrialRequestMetadata) UnsetResearch() {
	o.Research.Unset()
}

// GetPentest returns the Pentest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestMetadata) GetPentest() bool {
	if o == nil || IsNil(o.Pentest.Get()) {
		var ret bool
		return ret
	}
	return *o.Pentest.Get()
}

// GetPentestOk returns a tuple with the Pentest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestMetadata) GetPentestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pentest.Get(), o.Pentest.IsSet()
}

// HasPentest returns a boolean if a field has been set.
func (o *TrialRequestMetadata) HasPentest() bool {
	if o != nil && o.Pentest.IsSet() {
		return true
	}

	return false
}

// SetPentest gets a reference to the given NullableBool and assigns it to the Pentest field.
func (o *TrialRequestMetadata) SetPentest(v bool) {
	o.Pentest.Set(&v)
}
// SetPentestNil sets the value for Pentest to be an explicit nil
func (o *TrialRequestMetadata) SetPentestNil() {
	o.Pentest.Set(nil)
}

// UnsetPentest ensures that no value is present for Pentest, not even an explicit nil
func (o *TrialRequestMetadata) UnsetPentest() {
	o.Pentest.Unset()
}

// GetOther returns the Other field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestMetadata) GetOther() string {
	if o == nil || IsNil(o.Other.Get()) {
		var ret string
		return ret
	}
	return *o.Other.Get()
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestMetadata) GetOtherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Other.Get(), o.Other.IsSet()
}

// HasOther returns a boolean if a field has been set.
func (o *TrialRequestMetadata) HasOther() bool {
	if o != nil && o.Other.IsSet() {
		return true
	}

	return false
}

// SetOther gets a reference to the given NullableString and assigns it to the Other field.
func (o *TrialRequestMetadata) SetOther(v string) {
	o.Other.Set(&v)
}
// SetOtherNil sets the value for Other to be an explicit nil
func (o *TrialRequestMetadata) SetOtherNil() {
	o.Other.Set(nil)
}

// UnsetOther ensures that no value is present for Other, not even an explicit nil
func (o *TrialRequestMetadata) UnsetOther() {
	o.Other.Unset()
}

func (o TrialRequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrialRequestMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.Malware.IsSet() {
		toSerialize["malware"] = o.Malware.Get()
	}
	if o.Research.IsSet() {
		toSerialize["research"] = o.Research.Get()
	}
	if o.Pentest.IsSet() {
		toSerialize["pentest"] = o.Pentest.Get()
	}
	if o.Other.IsSet() {
		toSerialize["other"] = o.Other.Get()
	}
	return toSerialize, nil
}

type NullableTrialRequestMetadata struct {
	value *TrialRequestMetadata
	isSet bool
}

func (v NullableTrialRequestMetadata) Get() *TrialRequestMetadata {
	return v.value
}

func (v *NullableTrialRequestMetadata) Set(val *TrialRequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialRequestMetadata(val *TrialRequestMetadata) *NullableTrialRequestMetadata {
	return &NullableTrialRequestMetadata{value: val, isSet: true}
}

func (v NullableTrialRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


