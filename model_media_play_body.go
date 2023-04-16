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

// checks if the MediaPlayBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaPlayBody{}

// MediaPlayBody 
type MediaPlayBody struct {
	// Image ID
	ImageId NullableString `json:"imageId,omitempty"`
	// URL to a media file
	Url NullableString `json:"url,omitempty"`
}

// NewMediaPlayBody instantiates a new MediaPlayBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaPlayBody() *MediaPlayBody {
	this := MediaPlayBody{}
	return &this
}

// NewMediaPlayBodyWithDefaults instantiates a new MediaPlayBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaPlayBodyWithDefaults() *MediaPlayBody {
	this := MediaPlayBody{}
	return &this
}

// GetImageId returns the ImageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaPlayBody) GetImageId() string {
	if o == nil || IsNil(o.ImageId.Get()) {
		var ret string
		return ret
	}
	return *o.ImageId.Get()
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaPlayBody) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageId.Get(), o.ImageId.IsSet()
}

// HasImageId returns a boolean if a field has been set.
func (o *MediaPlayBody) HasImageId() bool {
	if o != nil && o.ImageId.IsSet() {
		return true
	}

	return false
}

// SetImageId gets a reference to the given NullableString and assigns it to the ImageId field.
func (o *MediaPlayBody) SetImageId(v string) {
	o.ImageId.Set(&v)
}
// SetImageIdNil sets the value for ImageId to be an explicit nil
func (o *MediaPlayBody) SetImageIdNil() {
	o.ImageId.Set(nil)
}

// UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
func (o *MediaPlayBody) UnsetImageId() {
	o.ImageId.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaPlayBody) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaPlayBody) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MediaPlayBody) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MediaPlayBody) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *MediaPlayBody) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MediaPlayBody) UnsetUrl() {
	o.Url.Unset()
}

func (o MediaPlayBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaPlayBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageId.IsSet() {
		toSerialize["imageId"] = o.ImageId.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableMediaPlayBody struct {
	value *MediaPlayBody
	isSet bool
}

func (v NullableMediaPlayBody) Get() *MediaPlayBody {
	return v.value
}

func (v *NullableMediaPlayBody) Set(val *MediaPlayBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaPlayBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaPlayBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaPlayBody(val *MediaPlayBody) *NullableMediaPlayBody {
	return &NullableMediaPlayBody{value: val, isSet: true}
}

func (v NullableMediaPlayBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaPlayBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


