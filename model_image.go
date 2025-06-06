/*
Corellium API

REST API to manage your virtual devices.

API version: 7.3.0-27831
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"time"
)

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image 
type Image struct {
	// 
	Status string `json:"status"`
	// Image ID
	Id NullableString `json:"id,omitempty"`
	// 
	Name NullableString `json:"name,omitempty"`
	// Type of image
	Type NullableString `json:"type,omitempty"`
	// 
	Filename NullableString `json:"filename,omitempty"`
	// 
	Uniqueid NullableString `json:"uniqueid,omitempty"`
	// 
	Size NullableFloat32 `json:"size,omitempty"`
	// project ID
	Project NullableString `json:"project,omitempty"`
	// When Image was created
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// When Image was last updated
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
}

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(status string) *Image {
	this := Image{}
	this.Status = status
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetStatus returns the Status field value
func (o *Image) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Image) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Image) SetStatus(v string) {
	o.Status = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Image) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Image) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Image) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Image) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Image) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Image) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Image) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Image) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Image) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Image) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Image) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Image) UnsetType() {
	o.Type.Unset()
}

// GetFilename returns the Filename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetFilename() string {
	if o == nil || IsNil(o.Filename.Get()) {
		var ret string
		return ret
	}
	return *o.Filename.Get()
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filename.Get(), o.Filename.IsSet()
}

// HasFilename returns a boolean if a field has been set.
func (o *Image) HasFilename() bool {
	if o != nil && o.Filename.IsSet() {
		return true
	}

	return false
}

// SetFilename gets a reference to the given NullableString and assigns it to the Filename field.
func (o *Image) SetFilename(v string) {
	o.Filename.Set(&v)
}
// SetFilenameNil sets the value for Filename to be an explicit nil
func (o *Image) SetFilenameNil() {
	o.Filename.Set(nil)
}

// UnsetFilename ensures that no value is present for Filename, not even an explicit nil
func (o *Image) UnsetFilename() {
	o.Filename.Unset()
}

// GetUniqueid returns the Uniqueid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetUniqueid() string {
	if o == nil || IsNil(o.Uniqueid.Get()) {
		var ret string
		return ret
	}
	return *o.Uniqueid.Get()
}

// GetUniqueidOk returns a tuple with the Uniqueid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetUniqueidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uniqueid.Get(), o.Uniqueid.IsSet()
}

// HasUniqueid returns a boolean if a field has been set.
func (o *Image) HasUniqueid() bool {
	if o != nil && o.Uniqueid.IsSet() {
		return true
	}

	return false
}

// SetUniqueid gets a reference to the given NullableString and assigns it to the Uniqueid field.
func (o *Image) SetUniqueid(v string) {
	o.Uniqueid.Set(&v)
}
// SetUniqueidNil sets the value for Uniqueid to be an explicit nil
func (o *Image) SetUniqueidNil() {
	o.Uniqueid.Set(nil)
}

// UnsetUniqueid ensures that no value is present for Uniqueid, not even an explicit nil
func (o *Image) UnsetUniqueid() {
	o.Uniqueid.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetSize() float32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret float32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *Image) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableFloat32 and assigns it to the Size field.
func (o *Image) SetSize(v float32) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *Image) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *Image) UnsetSize() {
	o.Size.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *Image) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *Image) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *Image) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *Image) UnsetProject() {
	o.Project.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Image) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *Image) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Image) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Image) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Image) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *Image) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Image) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Image) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Filename.IsSet() {
		toSerialize["filename"] = o.Filename.Get()
	}
	if o.Uniqueid.IsSet() {
		toSerialize["uniqueid"] = o.Uniqueid.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


