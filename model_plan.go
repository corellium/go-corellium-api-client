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

// checks if the Plan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plan{}

// Plan 
type Plan struct {
	// Plan ID
	PlanId NullableString `json:"planId,omitempty"`
	// Plan Name
	Name NullableString `json:"name,omitempty"`
}

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan() *Plan {
	this := Plan{}
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	return &this
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plan) GetPlanId() string {
	if o == nil || IsNil(o.PlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plan) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *Plan) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *Plan) SetPlanId(v string) {
	o.PlanId.Set(&v)
}
// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *Plan) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *Plan) UnsetPlanId() {
	o.PlanId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plan) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Plan) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Plan) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Plan) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Plan) UnsetName() {
	o.Name.Unset()
}

func (o Plan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Plan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PlanId.IsSet() {
		toSerialize["planId"] = o.PlanId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


