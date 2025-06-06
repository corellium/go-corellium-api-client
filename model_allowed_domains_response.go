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

// checks if the AllowedDomainsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowedDomainsResponse{}

// AllowedDomainsResponse 
type AllowedDomainsResponse struct {
	// List of allowed domains
	AllowedDomains []string `json:"allowedDomains,omitempty"`
}

// NewAllowedDomainsResponse instantiates a new AllowedDomainsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedDomainsResponse() *AllowedDomainsResponse {
	this := AllowedDomainsResponse{}
	return &this
}

// NewAllowedDomainsResponseWithDefaults instantiates a new AllowedDomainsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedDomainsResponseWithDefaults() *AllowedDomainsResponse {
	this := AllowedDomainsResponse{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllowedDomainsResponse) GetAllowedDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllowedDomainsResponse) GetAllowedDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedDomains) {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *AllowedDomainsResponse) HasAllowedDomains() bool {
	if o != nil && IsNil(o.AllowedDomains) {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *AllowedDomainsResponse) SetAllowedDomains(v []string) {
	o.AllowedDomains = v
}

func (o AllowedDomainsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowedDomainsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedDomains != nil {
		toSerialize["allowedDomains"] = o.AllowedDomains
	}
	return toSerialize, nil
}

type NullableAllowedDomainsResponse struct {
	value *AllowedDomainsResponse
	isSet bool
}

func (v NullableAllowedDomainsResponse) Get() *AllowedDomainsResponse {
	return v.value
}

func (v *NullableAllowedDomainsResponse) Set(val *AllowedDomainsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedDomainsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedDomainsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedDomainsResponse(val *AllowedDomainsResponse) *NullableAllowedDomainsResponse {
	return &NullableAllowedDomainsResponse{value: val, isSet: true}
}

func (v NullableAllowedDomainsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedDomainsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


