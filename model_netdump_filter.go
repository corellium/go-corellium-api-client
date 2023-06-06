/*
Corellium API

REST API to manage your virtual devices.

API version: 5.0.0-17089
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the NetdumpFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetdumpFilter{}

// NetdumpFilter 
type NetdumpFilter struct {
	// 
	PortRanges []string `json:"portRanges,omitempty"`
	// 
	SrcPorts []string `json:"srcPorts,omitempty"`
	// 
	DstPorts []string `json:"dstPorts,omitempty"`
	// 
	Ports []string `json:"ports,omitempty"`
	// 
	Protocols []string `json:"protocols,omitempty"`
	// 
	Processes []string `json:"processes,omitempty"`
}

// NewNetdumpFilter instantiates a new NetdumpFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetdumpFilter() *NetdumpFilter {
	this := NetdumpFilter{}
	return &this
}

// NewNetdumpFilterWithDefaults instantiates a new NetdumpFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetdumpFilterWithDefaults() *NetdumpFilter {
	this := NetdumpFilter{}
	return &this
}

// GetPortRanges returns the PortRanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetdumpFilter) GetPortRanges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PortRanges
}

// GetPortRangesOk returns a tuple with the PortRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetdumpFilter) GetPortRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.PortRanges) {
		return nil, false
	}
	return o.PortRanges, true
}

// HasPortRanges returns a boolean if a field has been set.
func (o *NetdumpFilter) HasPortRanges() bool {
	if o != nil && IsNil(o.PortRanges) {
		return true
	}

	return false
}

// SetPortRanges gets a reference to the given []string and assigns it to the PortRanges field.
func (o *NetdumpFilter) SetPortRanges(v []string) {
	o.PortRanges = v
}

// GetSrcPorts returns the SrcPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetdumpFilter) GetSrcPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SrcPorts
}

// GetSrcPortsOk returns a tuple with the SrcPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetdumpFilter) GetSrcPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.SrcPorts) {
		return nil, false
	}
	return o.SrcPorts, true
}

// HasSrcPorts returns a boolean if a field has been set.
func (o *NetdumpFilter) HasSrcPorts() bool {
	if o != nil && IsNil(o.SrcPorts) {
		return true
	}

	return false
}

// SetSrcPorts gets a reference to the given []string and assigns it to the SrcPorts field.
func (o *NetdumpFilter) SetSrcPorts(v []string) {
	o.SrcPorts = v
}

// GetDstPorts returns the DstPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetdumpFilter) GetDstPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DstPorts
}

// GetDstPortsOk returns a tuple with the DstPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetdumpFilter) GetDstPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.DstPorts) {
		return nil, false
	}
	return o.DstPorts, true
}

// HasDstPorts returns a boolean if a field has been set.
func (o *NetdumpFilter) HasDstPorts() bool {
	if o != nil && IsNil(o.DstPorts) {
		return true
	}

	return false
}

// SetDstPorts gets a reference to the given []string and assigns it to the DstPorts field.
func (o *NetdumpFilter) SetDstPorts(v []string) {
	o.DstPorts = v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetdumpFilter) GetPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetdumpFilter) GetPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *NetdumpFilter) HasPorts() bool {
	if o != nil && IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *NetdumpFilter) SetPorts(v []string) {
	o.Ports = v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetdumpFilter) GetProtocols() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetdumpFilter) GetProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.Protocols) {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *NetdumpFilter) HasProtocols() bool {
	if o != nil && IsNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *NetdumpFilter) SetProtocols(v []string) {
	o.Protocols = v
}

// GetProcesses returns the Processes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetdumpFilter) GetProcesses() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetdumpFilter) GetProcessesOk() ([]string, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// HasProcesses returns a boolean if a field has been set.
func (o *NetdumpFilter) HasProcesses() bool {
	if o != nil && IsNil(o.Processes) {
		return true
	}

	return false
}

// SetProcesses gets a reference to the given []string and assigns it to the Processes field.
func (o *NetdumpFilter) SetProcesses(v []string) {
	o.Processes = v
}

func (o NetdumpFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetdumpFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PortRanges != nil {
		toSerialize["portRanges"] = o.PortRanges
	}
	if o.SrcPorts != nil {
		toSerialize["srcPorts"] = o.SrcPorts
	}
	if o.DstPorts != nil {
		toSerialize["dstPorts"] = o.DstPorts
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.Protocols != nil {
		toSerialize["protocols"] = o.Protocols
	}
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	return toSerialize, nil
}

type NullableNetdumpFilter struct {
	value *NetdumpFilter
	isSet bool
}

func (v NullableNetdumpFilter) Get() *NetdumpFilter {
	return v.value
}

func (v *NullableNetdumpFilter) Set(val *NetdumpFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableNetdumpFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableNetdumpFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetdumpFilter(val *NetdumpFilter) *NullableNetdumpFilter {
	return &NullableNetdumpFilter{value: val, isSet: true}
}

func (v NullableNetdumpFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetdumpFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


