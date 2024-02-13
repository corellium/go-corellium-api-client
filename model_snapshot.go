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

// checks if the Snapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Snapshot{}

// Snapshot 
type Snapshot struct {
	// Snapshot ID
	Id NullableString `json:"id,omitempty"`
	// Snapshot name
	Name NullableString `json:"name,omitempty"`
	// Instance that this snapshot is of
	Instance NullableString `json:"instance,omitempty"`
	Status *SnapshotStatus `json:"status,omitempty"`
	// UNIX Date that the snapshot was created
	Date NullableFloat32 `json:"date,omitempty"`
	// 
	Fresh NullableBool `json:"fresh,omitempty"`
	// Live snapshot (included state and memory)
	Live NullableBool `json:"live,omitempty"`
	// 
	Local NullableBool `json:"local,omitempty"`
}

// NewSnapshot instantiates a new Snapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshot() *Snapshot {
	this := Snapshot{}
	return &this
}

// NewSnapshotWithDefaults instantiates a new Snapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotWithDefaults() *Snapshot {
	this := Snapshot{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Snapshot) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Snapshot) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Snapshot) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Snapshot) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Snapshot) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Snapshot) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Snapshot) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Snapshot) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Snapshot) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Snapshot) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Snapshot) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Snapshot) UnsetName() {
	o.Name.Unset()
}

// GetInstance returns the Instance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Snapshot) GetInstance() string {
	if o == nil || IsNil(o.Instance.Get()) {
		var ret string
		return ret
	}
	return *o.Instance.Get()
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Snapshot) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instance.Get(), o.Instance.IsSet()
}

// HasInstance returns a boolean if a field has been set.
func (o *Snapshot) HasInstance() bool {
	if o != nil && o.Instance.IsSet() {
		return true
	}

	return false
}

// SetInstance gets a reference to the given NullableString and assigns it to the Instance field.
func (o *Snapshot) SetInstance(v string) {
	o.Instance.Set(&v)
}
// SetInstanceNil sets the value for Instance to be an explicit nil
func (o *Snapshot) SetInstanceNil() {
	o.Instance.Set(nil)
}

// UnsetInstance ensures that no value is present for Instance, not even an explicit nil
func (o *Snapshot) UnsetInstance() {
	o.Instance.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Snapshot) GetStatus() SnapshotStatus {
	if o == nil || IsNil(o.Status) {
		var ret SnapshotStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetStatusOk() (*SnapshotStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Snapshot) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SnapshotStatus and assigns it to the Status field.
func (o *Snapshot) SetStatus(v SnapshotStatus) {
	o.Status = &v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Snapshot) GetDate() float32 {
	if o == nil || IsNil(o.Date.Get()) {
		var ret float32
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Snapshot) GetDateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *Snapshot) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableFloat32 and assigns it to the Date field.
func (o *Snapshot) SetDate(v float32) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *Snapshot) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *Snapshot) UnsetDate() {
	o.Date.Unset()
}

// GetFresh returns the Fresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Snapshot) GetFresh() bool {
	if o == nil || IsNil(o.Fresh.Get()) {
		var ret bool
		return ret
	}
	return *o.Fresh.Get()
}

// GetFreshOk returns a tuple with the Fresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Snapshot) GetFreshOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fresh.Get(), o.Fresh.IsSet()
}

// HasFresh returns a boolean if a field has been set.
func (o *Snapshot) HasFresh() bool {
	if o != nil && o.Fresh.IsSet() {
		return true
	}

	return false
}

// SetFresh gets a reference to the given NullableBool and assigns it to the Fresh field.
func (o *Snapshot) SetFresh(v bool) {
	o.Fresh.Set(&v)
}
// SetFreshNil sets the value for Fresh to be an explicit nil
func (o *Snapshot) SetFreshNil() {
	o.Fresh.Set(nil)
}

// UnsetFresh ensures that no value is present for Fresh, not even an explicit nil
func (o *Snapshot) UnsetFresh() {
	o.Fresh.Unset()
}

// GetLive returns the Live field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Snapshot) GetLive() bool {
	if o == nil || IsNil(o.Live.Get()) {
		var ret bool
		return ret
	}
	return *o.Live.Get()
}

// GetLiveOk returns a tuple with the Live field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Snapshot) GetLiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Live.Get(), o.Live.IsSet()
}

// HasLive returns a boolean if a field has been set.
func (o *Snapshot) HasLive() bool {
	if o != nil && o.Live.IsSet() {
		return true
	}

	return false
}

// SetLive gets a reference to the given NullableBool and assigns it to the Live field.
func (o *Snapshot) SetLive(v bool) {
	o.Live.Set(&v)
}
// SetLiveNil sets the value for Live to be an explicit nil
func (o *Snapshot) SetLiveNil() {
	o.Live.Set(nil)
}

// UnsetLive ensures that no value is present for Live, not even an explicit nil
func (o *Snapshot) UnsetLive() {
	o.Live.Unset()
}

// GetLocal returns the Local field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Snapshot) GetLocal() bool {
	if o == nil || IsNil(o.Local.Get()) {
		var ret bool
		return ret
	}
	return *o.Local.Get()
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Snapshot) GetLocalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Local.Get(), o.Local.IsSet()
}

// HasLocal returns a boolean if a field has been set.
func (o *Snapshot) HasLocal() bool {
	if o != nil && o.Local.IsSet() {
		return true
	}

	return false
}

// SetLocal gets a reference to the given NullableBool and assigns it to the Local field.
func (o *Snapshot) SetLocal(v bool) {
	o.Local.Set(&v)
}
// SetLocalNil sets the value for Local to be an explicit nil
func (o *Snapshot) SetLocalNil() {
	o.Local.Set(nil)
}

// UnsetLocal ensures that no value is present for Local, not even an explicit nil
func (o *Snapshot) UnsetLocal() {
	o.Local.Unset()
}

func (o Snapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Instance.IsSet() {
		toSerialize["instance"] = o.Instance.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Fresh.IsSet() {
		toSerialize["fresh"] = o.Fresh.Get()
	}
	if o.Live.IsSet() {
		toSerialize["live"] = o.Live.Get()
	}
	if o.Local.IsSet() {
		toSerialize["local"] = o.Local.Get()
	}
	return toSerialize, nil
}

type NullableSnapshot struct {
	value *Snapshot
	isSet bool
}

func (v NullableSnapshot) Get() *Snapshot {
	return v.value
}

func (v *NullableSnapshot) Set(val *Snapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshot(val *Snapshot) *NullableSnapshot {
	return &NullableSnapshot{value: val, isSet: true}
}

func (v NullableSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


