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

// checks if the SnapshotInvitationTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotInvitationTypes{}

// SnapshotInvitationTypes 
type SnapshotInvitationTypes struct {
	// Indicates if public link access is enabled
	PublicLink NullableBool `json:"publicLink,omitempty"`
	// Indicates if domain-restricted link access is enabled
	DomainRestrictedLink NullableBool `json:"domainRestrictedLink,omitempty"`
	// Indicates if password-protected public link access is enabled
	PasswordPublicLink NullableBool `json:"passwordPublicLink,omitempty"`
	// Indicates if email invite access is enabled
	EmailInvite NullableBool `json:"emailInvite,omitempty"`
}

// NewSnapshotInvitationTypes instantiates a new SnapshotInvitationTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotInvitationTypes() *SnapshotInvitationTypes {
	this := SnapshotInvitationTypes{}
	return &this
}

// NewSnapshotInvitationTypesWithDefaults instantiates a new SnapshotInvitationTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotInvitationTypesWithDefaults() *SnapshotInvitationTypes {
	this := SnapshotInvitationTypes{}
	return &this
}

// GetPublicLink returns the PublicLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInvitationTypes) GetPublicLink() bool {
	if o == nil || IsNil(o.PublicLink.Get()) {
		var ret bool
		return ret
	}
	return *o.PublicLink.Get()
}

// GetPublicLinkOk returns a tuple with the PublicLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInvitationTypes) GetPublicLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicLink.Get(), o.PublicLink.IsSet()
}

// HasPublicLink returns a boolean if a field has been set.
func (o *SnapshotInvitationTypes) HasPublicLink() bool {
	if o != nil && o.PublicLink.IsSet() {
		return true
	}

	return false
}

// SetPublicLink gets a reference to the given NullableBool and assigns it to the PublicLink field.
func (o *SnapshotInvitationTypes) SetPublicLink(v bool) {
	o.PublicLink.Set(&v)
}
// SetPublicLinkNil sets the value for PublicLink to be an explicit nil
func (o *SnapshotInvitationTypes) SetPublicLinkNil() {
	o.PublicLink.Set(nil)
}

// UnsetPublicLink ensures that no value is present for PublicLink, not even an explicit nil
func (o *SnapshotInvitationTypes) UnsetPublicLink() {
	o.PublicLink.Unset()
}

// GetDomainRestrictedLink returns the DomainRestrictedLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInvitationTypes) GetDomainRestrictedLink() bool {
	if o == nil || IsNil(o.DomainRestrictedLink.Get()) {
		var ret bool
		return ret
	}
	return *o.DomainRestrictedLink.Get()
}

// GetDomainRestrictedLinkOk returns a tuple with the DomainRestrictedLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInvitationTypes) GetDomainRestrictedLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainRestrictedLink.Get(), o.DomainRestrictedLink.IsSet()
}

// HasDomainRestrictedLink returns a boolean if a field has been set.
func (o *SnapshotInvitationTypes) HasDomainRestrictedLink() bool {
	if o != nil && o.DomainRestrictedLink.IsSet() {
		return true
	}

	return false
}

// SetDomainRestrictedLink gets a reference to the given NullableBool and assigns it to the DomainRestrictedLink field.
func (o *SnapshotInvitationTypes) SetDomainRestrictedLink(v bool) {
	o.DomainRestrictedLink.Set(&v)
}
// SetDomainRestrictedLinkNil sets the value for DomainRestrictedLink to be an explicit nil
func (o *SnapshotInvitationTypes) SetDomainRestrictedLinkNil() {
	o.DomainRestrictedLink.Set(nil)
}

// UnsetDomainRestrictedLink ensures that no value is present for DomainRestrictedLink, not even an explicit nil
func (o *SnapshotInvitationTypes) UnsetDomainRestrictedLink() {
	o.DomainRestrictedLink.Unset()
}

// GetPasswordPublicLink returns the PasswordPublicLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInvitationTypes) GetPasswordPublicLink() bool {
	if o == nil || IsNil(o.PasswordPublicLink.Get()) {
		var ret bool
		return ret
	}
	return *o.PasswordPublicLink.Get()
}

// GetPasswordPublicLinkOk returns a tuple with the PasswordPublicLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInvitationTypes) GetPasswordPublicLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPublicLink.Get(), o.PasswordPublicLink.IsSet()
}

// HasPasswordPublicLink returns a boolean if a field has been set.
func (o *SnapshotInvitationTypes) HasPasswordPublicLink() bool {
	if o != nil && o.PasswordPublicLink.IsSet() {
		return true
	}

	return false
}

// SetPasswordPublicLink gets a reference to the given NullableBool and assigns it to the PasswordPublicLink field.
func (o *SnapshotInvitationTypes) SetPasswordPublicLink(v bool) {
	o.PasswordPublicLink.Set(&v)
}
// SetPasswordPublicLinkNil sets the value for PasswordPublicLink to be an explicit nil
func (o *SnapshotInvitationTypes) SetPasswordPublicLinkNil() {
	o.PasswordPublicLink.Set(nil)
}

// UnsetPasswordPublicLink ensures that no value is present for PasswordPublicLink, not even an explicit nil
func (o *SnapshotInvitationTypes) UnsetPasswordPublicLink() {
	o.PasswordPublicLink.Unset()
}

// GetEmailInvite returns the EmailInvite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInvitationTypes) GetEmailInvite() bool {
	if o == nil || IsNil(o.EmailInvite.Get()) {
		var ret bool
		return ret
	}
	return *o.EmailInvite.Get()
}

// GetEmailInviteOk returns a tuple with the EmailInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInvitationTypes) GetEmailInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailInvite.Get(), o.EmailInvite.IsSet()
}

// HasEmailInvite returns a boolean if a field has been set.
func (o *SnapshotInvitationTypes) HasEmailInvite() bool {
	if o != nil && o.EmailInvite.IsSet() {
		return true
	}

	return false
}

// SetEmailInvite gets a reference to the given NullableBool and assigns it to the EmailInvite field.
func (o *SnapshotInvitationTypes) SetEmailInvite(v bool) {
	o.EmailInvite.Set(&v)
}
// SetEmailInviteNil sets the value for EmailInvite to be an explicit nil
func (o *SnapshotInvitationTypes) SetEmailInviteNil() {
	o.EmailInvite.Set(nil)
}

// UnsetEmailInvite ensures that no value is present for EmailInvite, not even an explicit nil
func (o *SnapshotInvitationTypes) UnsetEmailInvite() {
	o.EmailInvite.Unset()
}

func (o SnapshotInvitationTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotInvitationTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicLink.IsSet() {
		toSerialize["publicLink"] = o.PublicLink.Get()
	}
	if o.DomainRestrictedLink.IsSet() {
		toSerialize["domainRestrictedLink"] = o.DomainRestrictedLink.Get()
	}
	if o.PasswordPublicLink.IsSet() {
		toSerialize["passwordPublicLink"] = o.PasswordPublicLink.Get()
	}
	if o.EmailInvite.IsSet() {
		toSerialize["emailInvite"] = o.EmailInvite.Get()
	}
	return toSerialize, nil
}

type NullableSnapshotInvitationTypes struct {
	value *SnapshotInvitationTypes
	isSet bool
}

func (v NullableSnapshotInvitationTypes) Get() *SnapshotInvitationTypes {
	return v.value
}

func (v *NullableSnapshotInvitationTypes) Set(val *SnapshotInvitationTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotInvitationTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotInvitationTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotInvitationTypes(val *SnapshotInvitationTypes) *NullableSnapshotInvitationTypes {
	return &NullableSnapshotInvitationTypes{value: val, isSet: true}
}

func (v NullableSnapshotInvitationTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotInvitationTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


