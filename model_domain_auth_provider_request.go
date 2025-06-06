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

// checks if the DomainAuthProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainAuthProviderRequest{}

// DomainAuthProviderRequest 
type DomainAuthProviderRequest struct {
	// Provider Type
	ProviderType string `json:"providerType"`
	// Enabled/Disabled
	Enabled bool `json:"enabled"`
	// Login Button Text
	Label NullableString `json:"label,omitempty"`
	Config *OpenIDConfig `json:"config,omitempty"`
}

// NewDomainAuthProviderRequest instantiates a new DomainAuthProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainAuthProviderRequest(providerType string, enabled bool) *DomainAuthProviderRequest {
	this := DomainAuthProviderRequest{}
	this.ProviderType = providerType
	this.Enabled = enabled
	return &this
}

// NewDomainAuthProviderRequestWithDefaults instantiates a new DomainAuthProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainAuthProviderRequestWithDefaults() *DomainAuthProviderRequest {
	this := DomainAuthProviderRequest{}
	return &this
}

// GetProviderType returns the ProviderType field value
func (o *DomainAuthProviderRequest) GetProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderRequest) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *DomainAuthProviderRequest) SetProviderType(v string) {
	o.ProviderType = v
}

// GetEnabled returns the Enabled field value
func (o *DomainAuthProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DomainAuthProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainAuthProviderRequest) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainAuthProviderRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *DomainAuthProviderRequest) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *DomainAuthProviderRequest) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *DomainAuthProviderRequest) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *DomainAuthProviderRequest) UnsetLabel() {
	o.Label.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *DomainAuthProviderRequest) GetConfig() OpenIDConfig {
	if o == nil || IsNil(o.Config) {
		var ret OpenIDConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderRequest) GetConfigOk() (*OpenIDConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *DomainAuthProviderRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OpenIDConfig and assigns it to the Config field.
func (o *DomainAuthProviderRequest) SetConfig(v OpenIDConfig) {
	o.Config = &v
}

func (o DomainAuthProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainAuthProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerType"] = o.ProviderType
	toSerialize["enabled"] = o.Enabled
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableDomainAuthProviderRequest struct {
	value *DomainAuthProviderRequest
	isSet bool
}

func (v NullableDomainAuthProviderRequest) Get() *DomainAuthProviderRequest {
	return v.value
}

func (v *NullableDomainAuthProviderRequest) Set(val *DomainAuthProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainAuthProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainAuthProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainAuthProviderRequest(val *DomainAuthProviderRequest) *NullableDomainAuthProviderRequest {
	return &NullableDomainAuthProviderRequest{value: val, isSet: true}
}

func (v NullableDomainAuthProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainAuthProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


