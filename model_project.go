/*
Corellium API

REST API to manage your virtual devices.

API version: 4.3.1-16664
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project 
type Project struct {
	// Project Identifier
	Id string `json:"id"`
	// Project Name
	Name NullableString `json:"name,omitempty"`
	Settings *ProjectSettings `json:"settings,omitempty"`
	Quotas *ProjectQuota `json:"quotas,omitempty"`
	QuotasUsed *ProjectUsage `json:"quotasUsed,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(id string) *Project {
	this := Project{}
	this.Id = id
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetId returns the Id field value
func (o *Project) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Project) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Project) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Project) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Project) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Project) UnsetName() {
	o.Name.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Project) GetSettings() ProjectSettings {
	if o == nil || IsNil(o.Settings) {
		var ret ProjectSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSettingsOk() (*ProjectSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Project) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ProjectSettings and assigns it to the Settings field.
func (o *Project) SetSettings(v ProjectSettings) {
	o.Settings = &v
}

// GetQuotas returns the Quotas field value if set, zero value otherwise.
func (o *Project) GetQuotas() ProjectQuota {
	if o == nil || IsNil(o.Quotas) {
		var ret ProjectQuota
		return ret
	}
	return *o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetQuotasOk() (*ProjectQuota, bool) {
	if o == nil || IsNil(o.Quotas) {
		return nil, false
	}
	return o.Quotas, true
}

// HasQuotas returns a boolean if a field has been set.
func (o *Project) HasQuotas() bool {
	if o != nil && !IsNil(o.Quotas) {
		return true
	}

	return false
}

// SetQuotas gets a reference to the given ProjectQuota and assigns it to the Quotas field.
func (o *Project) SetQuotas(v ProjectQuota) {
	o.Quotas = &v
}

// GetQuotasUsed returns the QuotasUsed field value if set, zero value otherwise.
func (o *Project) GetQuotasUsed() ProjectUsage {
	if o == nil || IsNil(o.QuotasUsed) {
		var ret ProjectUsage
		return ret
	}
	return *o.QuotasUsed
}

// GetQuotasUsedOk returns a tuple with the QuotasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetQuotasUsedOk() (*ProjectUsage, bool) {
	if o == nil || IsNil(o.QuotasUsed) {
		return nil, false
	}
	return o.QuotasUsed, true
}

// HasQuotasUsed returns a boolean if a field has been set.
func (o *Project) HasQuotasUsed() bool {
	if o != nil && !IsNil(o.QuotasUsed) {
		return true
	}

	return false
}

// SetQuotasUsed gets a reference to the given ProjectUsage and assigns it to the QuotasUsed field.
func (o *Project) SetQuotasUsed(v ProjectUsage) {
	o.QuotasUsed = &v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Quotas) {
		toSerialize["quotas"] = o.Quotas
	}
	if !IsNil(o.QuotasUsed) {
		toSerialize["quotasUsed"] = o.QuotasUsed
	}
	return toSerialize, nil
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


