/*
Corellium API

REST API to manage your virtual devices.

API version: 6.1.0-20784
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"time"
)

// checks if the Firmware type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Firmware{}

// Firmware 
type Firmware struct {
	// 
	Version NullableString `json:"version,omitempty"`
	// 
	Buildid NullableString `json:"buildid,omitempty"`
	// Android only flavor
	AndroidFlavor NullableString `json:"AndroidFlavor,omitempty"`
	// Android only API version
	APIVersion NullableString `json:"APIVersion,omitempty"`
	// 
	Sha256sum NullableString `json:"sha256sum,omitempty"`
	// 
	Sha1sum NullableString `json:"sha1sum,omitempty"`
	// 
	Md5sum NullableString `json:"md5sum,omitempty"`
	// 
	Size NullableInt32 `json:"size,omitempty"`
	// 
	UniqueId NullableString `json:"uniqueId,omitempty"`
	// Firmware metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Release Date
	Releasedate NullableTime `json:"releasedate,omitempty"`
	// Date uploaded
	Uploaddate NullableTime `json:"uploaddate,omitempty"`
	// URL firmware is available at
	Url NullableString `json:"url,omitempty"`
	// URL firmware is available at from vendor
	OrigUrl NullableString `json:"orig_url,omitempty"`
	// 
	Filename NullableString `json:"filename,omitempty"`
}

// NewFirmware instantiates a new Firmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmware() *Firmware {
	this := Firmware{}
	return &this
}

// NewFirmwareWithDefaults instantiates a new Firmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareWithDefaults() *Firmware {
	this := Firmware{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *Firmware) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *Firmware) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *Firmware) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *Firmware) UnsetVersion() {
	o.Version.Unset()
}

// GetBuildid returns the Buildid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetBuildid() string {
	if o == nil || IsNil(o.Buildid.Get()) {
		var ret string
		return ret
	}
	return *o.Buildid.Get()
}

// GetBuildidOk returns a tuple with the Buildid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetBuildidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buildid.Get(), o.Buildid.IsSet()
}

// HasBuildid returns a boolean if a field has been set.
func (o *Firmware) HasBuildid() bool {
	if o != nil && o.Buildid.IsSet() {
		return true
	}

	return false
}

// SetBuildid gets a reference to the given NullableString and assigns it to the Buildid field.
func (o *Firmware) SetBuildid(v string) {
	o.Buildid.Set(&v)
}
// SetBuildidNil sets the value for Buildid to be an explicit nil
func (o *Firmware) SetBuildidNil() {
	o.Buildid.Set(nil)
}

// UnsetBuildid ensures that no value is present for Buildid, not even an explicit nil
func (o *Firmware) UnsetBuildid() {
	o.Buildid.Unset()
}

// GetAndroidFlavor returns the AndroidFlavor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetAndroidFlavor() string {
	if o == nil || IsNil(o.AndroidFlavor.Get()) {
		var ret string
		return ret
	}
	return *o.AndroidFlavor.Get()
}

// GetAndroidFlavorOk returns a tuple with the AndroidFlavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetAndroidFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AndroidFlavor.Get(), o.AndroidFlavor.IsSet()
}

// HasAndroidFlavor returns a boolean if a field has been set.
func (o *Firmware) HasAndroidFlavor() bool {
	if o != nil && o.AndroidFlavor.IsSet() {
		return true
	}

	return false
}

// SetAndroidFlavor gets a reference to the given NullableString and assigns it to the AndroidFlavor field.
func (o *Firmware) SetAndroidFlavor(v string) {
	o.AndroidFlavor.Set(&v)
}
// SetAndroidFlavorNil sets the value for AndroidFlavor to be an explicit nil
func (o *Firmware) SetAndroidFlavorNil() {
	o.AndroidFlavor.Set(nil)
}

// UnsetAndroidFlavor ensures that no value is present for AndroidFlavor, not even an explicit nil
func (o *Firmware) UnsetAndroidFlavor() {
	o.AndroidFlavor.Unset()
}

// GetAPIVersion returns the APIVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetAPIVersion() string {
	if o == nil || IsNil(o.APIVersion.Get()) {
		var ret string
		return ret
	}
	return *o.APIVersion.Get()
}

// GetAPIVersionOk returns a tuple with the APIVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetAPIVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.APIVersion.Get(), o.APIVersion.IsSet()
}

// HasAPIVersion returns a boolean if a field has been set.
func (o *Firmware) HasAPIVersion() bool {
	if o != nil && o.APIVersion.IsSet() {
		return true
	}

	return false
}

// SetAPIVersion gets a reference to the given NullableString and assigns it to the APIVersion field.
func (o *Firmware) SetAPIVersion(v string) {
	o.APIVersion.Set(&v)
}
// SetAPIVersionNil sets the value for APIVersion to be an explicit nil
func (o *Firmware) SetAPIVersionNil() {
	o.APIVersion.Set(nil)
}

// UnsetAPIVersion ensures that no value is present for APIVersion, not even an explicit nil
func (o *Firmware) UnsetAPIVersion() {
	o.APIVersion.Unset()
}

// GetSha256sum returns the Sha256sum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetSha256sum() string {
	if o == nil || IsNil(o.Sha256sum.Get()) {
		var ret string
		return ret
	}
	return *o.Sha256sum.Get()
}

// GetSha256sumOk returns a tuple with the Sha256sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetSha256sumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha256sum.Get(), o.Sha256sum.IsSet()
}

// HasSha256sum returns a boolean if a field has been set.
func (o *Firmware) HasSha256sum() bool {
	if o != nil && o.Sha256sum.IsSet() {
		return true
	}

	return false
}

// SetSha256sum gets a reference to the given NullableString and assigns it to the Sha256sum field.
func (o *Firmware) SetSha256sum(v string) {
	o.Sha256sum.Set(&v)
}
// SetSha256sumNil sets the value for Sha256sum to be an explicit nil
func (o *Firmware) SetSha256sumNil() {
	o.Sha256sum.Set(nil)
}

// UnsetSha256sum ensures that no value is present for Sha256sum, not even an explicit nil
func (o *Firmware) UnsetSha256sum() {
	o.Sha256sum.Unset()
}

// GetSha1sum returns the Sha1sum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetSha1sum() string {
	if o == nil || IsNil(o.Sha1sum.Get()) {
		var ret string
		return ret
	}
	return *o.Sha1sum.Get()
}

// GetSha1sumOk returns a tuple with the Sha1sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetSha1sumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha1sum.Get(), o.Sha1sum.IsSet()
}

// HasSha1sum returns a boolean if a field has been set.
func (o *Firmware) HasSha1sum() bool {
	if o != nil && o.Sha1sum.IsSet() {
		return true
	}

	return false
}

// SetSha1sum gets a reference to the given NullableString and assigns it to the Sha1sum field.
func (o *Firmware) SetSha1sum(v string) {
	o.Sha1sum.Set(&v)
}
// SetSha1sumNil sets the value for Sha1sum to be an explicit nil
func (o *Firmware) SetSha1sumNil() {
	o.Sha1sum.Set(nil)
}

// UnsetSha1sum ensures that no value is present for Sha1sum, not even an explicit nil
func (o *Firmware) UnsetSha1sum() {
	o.Sha1sum.Unset()
}

// GetMd5sum returns the Md5sum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetMd5sum() string {
	if o == nil || IsNil(o.Md5sum.Get()) {
		var ret string
		return ret
	}
	return *o.Md5sum.Get()
}

// GetMd5sumOk returns a tuple with the Md5sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetMd5sumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Md5sum.Get(), o.Md5sum.IsSet()
}

// HasMd5sum returns a boolean if a field has been set.
func (o *Firmware) HasMd5sum() bool {
	if o != nil && o.Md5sum.IsSet() {
		return true
	}

	return false
}

// SetMd5sum gets a reference to the given NullableString and assigns it to the Md5sum field.
func (o *Firmware) SetMd5sum(v string) {
	o.Md5sum.Set(&v)
}
// SetMd5sumNil sets the value for Md5sum to be an explicit nil
func (o *Firmware) SetMd5sumNil() {
	o.Md5sum.Set(nil)
}

// UnsetMd5sum ensures that no value is present for Md5sum, not even an explicit nil
func (o *Firmware) UnsetMd5sum() {
	o.Md5sum.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *Firmware) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *Firmware) SetSize(v int32) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *Firmware) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *Firmware) UnsetSize() {
	o.Size.Unset()
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId.Get()) {
		var ret string
		return ret
	}
	return *o.UniqueId.Get()
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniqueId.Get(), o.UniqueId.IsSet()
}

// HasUniqueId returns a boolean if a field has been set.
func (o *Firmware) HasUniqueId() bool {
	if o != nil && o.UniqueId.IsSet() {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given NullableString and assigns it to the UniqueId field.
func (o *Firmware) SetUniqueId(v string) {
	o.UniqueId.Set(&v)
}
// SetUniqueIdNil sets the value for UniqueId to be an explicit nil
func (o *Firmware) SetUniqueIdNil() {
	o.UniqueId.Set(nil)
}

// UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
func (o *Firmware) UnsetUniqueId() {
	o.UniqueId.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Firmware) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Firmware) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetReleasedate returns the Releasedate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetReleasedate() time.Time {
	if o == nil || IsNil(o.Releasedate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Releasedate.Get()
}

// GetReleasedateOk returns a tuple with the Releasedate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetReleasedateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Releasedate.Get(), o.Releasedate.IsSet()
}

// HasReleasedate returns a boolean if a field has been set.
func (o *Firmware) HasReleasedate() bool {
	if o != nil && o.Releasedate.IsSet() {
		return true
	}

	return false
}

// SetReleasedate gets a reference to the given NullableTime and assigns it to the Releasedate field.
func (o *Firmware) SetReleasedate(v time.Time) {
	o.Releasedate.Set(&v)
}
// SetReleasedateNil sets the value for Releasedate to be an explicit nil
func (o *Firmware) SetReleasedateNil() {
	o.Releasedate.Set(nil)
}

// UnsetReleasedate ensures that no value is present for Releasedate, not even an explicit nil
func (o *Firmware) UnsetReleasedate() {
	o.Releasedate.Unset()
}

// GetUploaddate returns the Uploaddate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetUploaddate() time.Time {
	if o == nil || IsNil(o.Uploaddate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Uploaddate.Get()
}

// GetUploaddateOk returns a tuple with the Uploaddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetUploaddateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uploaddate.Get(), o.Uploaddate.IsSet()
}

// HasUploaddate returns a boolean if a field has been set.
func (o *Firmware) HasUploaddate() bool {
	if o != nil && o.Uploaddate.IsSet() {
		return true
	}

	return false
}

// SetUploaddate gets a reference to the given NullableTime and assigns it to the Uploaddate field.
func (o *Firmware) SetUploaddate(v time.Time) {
	o.Uploaddate.Set(&v)
}
// SetUploaddateNil sets the value for Uploaddate to be an explicit nil
func (o *Firmware) SetUploaddateNil() {
	o.Uploaddate.Set(nil)
}

// UnsetUploaddate ensures that no value is present for Uploaddate, not even an explicit nil
func (o *Firmware) UnsetUploaddate() {
	o.Uploaddate.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Firmware) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Firmware) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Firmware) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Firmware) UnsetUrl() {
	o.Url.Unset()
}

// GetOrigUrl returns the OrigUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetOrigUrl() string {
	if o == nil || IsNil(o.OrigUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OrigUrl.Get()
}

// GetOrigUrlOk returns a tuple with the OrigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetOrigUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrigUrl.Get(), o.OrigUrl.IsSet()
}

// HasOrigUrl returns a boolean if a field has been set.
func (o *Firmware) HasOrigUrl() bool {
	if o != nil && o.OrigUrl.IsSet() {
		return true
	}

	return false
}

// SetOrigUrl gets a reference to the given NullableString and assigns it to the OrigUrl field.
func (o *Firmware) SetOrigUrl(v string) {
	o.OrigUrl.Set(&v)
}
// SetOrigUrlNil sets the value for OrigUrl to be an explicit nil
func (o *Firmware) SetOrigUrlNil() {
	o.OrigUrl.Set(nil)
}

// UnsetOrigUrl ensures that no value is present for OrigUrl, not even an explicit nil
func (o *Firmware) UnsetOrigUrl() {
	o.OrigUrl.Unset()
}

// GetFilename returns the Filename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Firmware) GetFilename() string {
	if o == nil || IsNil(o.Filename.Get()) {
		var ret string
		return ret
	}
	return *o.Filename.Get()
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Firmware) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filename.Get(), o.Filename.IsSet()
}

// HasFilename returns a boolean if a field has been set.
func (o *Firmware) HasFilename() bool {
	if o != nil && o.Filename.IsSet() {
		return true
	}

	return false
}

// SetFilename gets a reference to the given NullableString and assigns it to the Filename field.
func (o *Firmware) SetFilename(v string) {
	o.Filename.Set(&v)
}
// SetFilenameNil sets the value for Filename to be an explicit nil
func (o *Firmware) SetFilenameNil() {
	o.Filename.Set(nil)
}

// UnsetFilename ensures that no value is present for Filename, not even an explicit nil
func (o *Firmware) UnsetFilename() {
	o.Filename.Unset()
}

func (o Firmware) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Firmware) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Buildid.IsSet() {
		toSerialize["buildid"] = o.Buildid.Get()
	}
	if o.AndroidFlavor.IsSet() {
		toSerialize["AndroidFlavor"] = o.AndroidFlavor.Get()
	}
	if o.APIVersion.IsSet() {
		toSerialize["APIVersion"] = o.APIVersion.Get()
	}
	if o.Sha256sum.IsSet() {
		toSerialize["sha256sum"] = o.Sha256sum.Get()
	}
	if o.Sha1sum.IsSet() {
		toSerialize["sha1sum"] = o.Sha1sum.Get()
	}
	if o.Md5sum.IsSet() {
		toSerialize["md5sum"] = o.Md5sum.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.UniqueId.IsSet() {
		toSerialize["uniqueId"] = o.UniqueId.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Releasedate.IsSet() {
		toSerialize["releasedate"] = o.Releasedate.Get()
	}
	if o.Uploaddate.IsSet() {
		toSerialize["uploaddate"] = o.Uploaddate.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.OrigUrl.IsSet() {
		toSerialize["orig_url"] = o.OrigUrl.Get()
	}
	if o.Filename.IsSet() {
		toSerialize["filename"] = o.Filename.Get()
	}
	return toSerialize, nil
}

type NullableFirmware struct {
	value *Firmware
	isSet bool
}

func (v NullableFirmware) Get() *Firmware {
	return v.value
}

func (v *NullableFirmware) Set(val *Firmware) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmware) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmware(val *Firmware) *NullableFirmware {
	return &NullableFirmware{value: val, isSet: true}
}

func (v NullableFirmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


