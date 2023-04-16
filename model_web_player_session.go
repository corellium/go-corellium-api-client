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

// checks if the WebPlayerSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebPlayerSession{}

// WebPlayerSession 
type WebPlayerSession struct {
	// New Session Identifier
	Identifier string `json:"identifier"`
	// Session Token
	Token string `json:"token"`
	// Expiration in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z
	Expiration float32 `json:"expiration"`
}

// NewWebPlayerSession instantiates a new WebPlayerSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebPlayerSession(identifier string, token string, expiration float32) *WebPlayerSession {
	this := WebPlayerSession{}
	this.Identifier = identifier
	this.Token = token
	this.Expiration = expiration
	return &this
}

// NewWebPlayerSessionWithDefaults instantiates a new WebPlayerSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebPlayerSessionWithDefaults() *WebPlayerSession {
	this := WebPlayerSession{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *WebPlayerSession) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *WebPlayerSession) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *WebPlayerSession) SetIdentifier(v string) {
	o.Identifier = v
}

// GetToken returns the Token field value
func (o *WebPlayerSession) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *WebPlayerSession) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *WebPlayerSession) SetToken(v string) {
	o.Token = v
}

// GetExpiration returns the Expiration field value
func (o *WebPlayerSession) GetExpiration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *WebPlayerSession) GetExpirationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *WebPlayerSession) SetExpiration(v float32) {
	o.Expiration = v
}

func (o WebPlayerSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebPlayerSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	toSerialize["token"] = o.Token
	toSerialize["expiration"] = o.Expiration
	return toSerialize, nil
}

type NullableWebPlayerSession struct {
	value *WebPlayerSession
	isSet bool
}

func (v NullableWebPlayerSession) Get() *WebPlayerSession {
	return v.value
}

func (v *NullableWebPlayerSession) Set(val *WebPlayerSession) {
	v.value = val
	v.isSet = true
}

func (v NullableWebPlayerSession) IsSet() bool {
	return v.isSet
}

func (v *NullableWebPlayerSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebPlayerSession(val *WebPlayerSession) *NullableWebPlayerSession {
	return &NullableWebPlayerSession{value: val, isSet: true}
}

func (v NullableWebPlayerSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebPlayerSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


