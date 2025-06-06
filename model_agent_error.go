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

// checks if the AgentError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentError{}

// AgentError 
type AgentError struct {
	// The error encountered by the agent
	Error string `json:"error"`
	// AgentError
	ErrorID string `json:"errorID"`
	// The full error encountered by the agent
	OriginalError map[string]interface{} `json:"originalError,omitempty"`
}

// NewAgentError instantiates a new AgentError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentError(error_ string, errorID string) *AgentError {
	this := AgentError{}
	this.Error = error_
	this.ErrorID = errorID
	return &this
}

// NewAgentErrorWithDefaults instantiates a new AgentError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentErrorWithDefaults() *AgentError {
	this := AgentError{}
	return &this
}

// GetError returns the Error field value
func (o *AgentError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AgentError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AgentError) SetError(v string) {
	o.Error = v
}

// GetErrorID returns the ErrorID field value
func (o *AgentError) GetErrorID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorID
}

// GetErrorIDOk returns a tuple with the ErrorID field value
// and a boolean to check if the value has been set.
func (o *AgentError) GetErrorIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorID, true
}

// SetErrorID sets field value
func (o *AgentError) SetErrorID(v string) {
	o.ErrorID = v
}

// GetOriginalError returns the OriginalError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentError) GetOriginalError() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.OriginalError
}

// GetOriginalErrorOk returns a tuple with the OriginalError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentError) GetOriginalErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OriginalError) {
		return map[string]interface{}{}, false
	}
	return o.OriginalError, true
}

// HasOriginalError returns a boolean if a field has been set.
func (o *AgentError) HasOriginalError() bool {
	if o != nil && IsNil(o.OriginalError) {
		return true
	}

	return false
}

// SetOriginalError gets a reference to the given map[string]interface{} and assigns it to the OriginalError field.
func (o *AgentError) SetOriginalError(v map[string]interface{}) {
	o.OriginalError = v
}

func (o AgentError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["errorID"] = o.ErrorID
	if o.OriginalError != nil {
		toSerialize["originalError"] = o.OriginalError
	}
	return toSerialize, nil
}

type NullableAgentError struct {
	value *AgentError
	isSet bool
}

func (v NullableAgentError) Get() *AgentError {
	return v.value
}

func (v *NullableAgentError) Set(val *AgentError) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentError) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentError(val *AgentError) *NullableAgentError {
	return &NullableAgentError{value: val, isSet: true}
}

func (v NullableAgentError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


