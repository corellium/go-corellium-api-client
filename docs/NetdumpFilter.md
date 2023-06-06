# NetdumpFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortRanges** | Pointer to **[]string** |  | [optional] 
**SrcPorts** | Pointer to **[]string** |  | [optional] 
**DstPorts** | Pointer to **[]string** |  | [optional] 
**Ports** | Pointer to **[]string** |  | [optional] 
**Protocols** | Pointer to **[]string** |  | [optional] 
**Processes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetdumpFilter

`func NewNetdumpFilter() *NetdumpFilter`

NewNetdumpFilter instantiates a new NetdumpFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetdumpFilterWithDefaults

`func NewNetdumpFilterWithDefaults() *NetdumpFilter`

NewNetdumpFilterWithDefaults instantiates a new NetdumpFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortRanges

`func (o *NetdumpFilter) GetPortRanges() []string`

GetPortRanges returns the PortRanges field if non-nil, zero value otherwise.

### GetPortRangesOk

`func (o *NetdumpFilter) GetPortRangesOk() (*[]string, bool)`

GetPortRangesOk returns a tuple with the PortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRanges

`func (o *NetdumpFilter) SetPortRanges(v []string)`

SetPortRanges sets PortRanges field to given value.

### HasPortRanges

`func (o *NetdumpFilter) HasPortRanges() bool`

HasPortRanges returns a boolean if a field has been set.

### SetPortRangesNil

`func (o *NetdumpFilter) SetPortRangesNil(b bool)`

 SetPortRangesNil sets the value for PortRanges to be an explicit nil

### UnsetPortRanges
`func (o *NetdumpFilter) UnsetPortRanges()`

UnsetPortRanges ensures that no value is present for PortRanges, not even an explicit nil
### GetSrcPorts

`func (o *NetdumpFilter) GetSrcPorts() []string`

GetSrcPorts returns the SrcPorts field if non-nil, zero value otherwise.

### GetSrcPortsOk

`func (o *NetdumpFilter) GetSrcPortsOk() (*[]string, bool)`

GetSrcPortsOk returns a tuple with the SrcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPorts

`func (o *NetdumpFilter) SetSrcPorts(v []string)`

SetSrcPorts sets SrcPorts field to given value.

### HasSrcPorts

`func (o *NetdumpFilter) HasSrcPorts() bool`

HasSrcPorts returns a boolean if a field has been set.

### SetSrcPortsNil

`func (o *NetdumpFilter) SetSrcPortsNil(b bool)`

 SetSrcPortsNil sets the value for SrcPorts to be an explicit nil

### UnsetSrcPorts
`func (o *NetdumpFilter) UnsetSrcPorts()`

UnsetSrcPorts ensures that no value is present for SrcPorts, not even an explicit nil
### GetDstPorts

`func (o *NetdumpFilter) GetDstPorts() []string`

GetDstPorts returns the DstPorts field if non-nil, zero value otherwise.

### GetDstPortsOk

`func (o *NetdumpFilter) GetDstPortsOk() (*[]string, bool)`

GetDstPortsOk returns a tuple with the DstPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPorts

`func (o *NetdumpFilter) SetDstPorts(v []string)`

SetDstPorts sets DstPorts field to given value.

### HasDstPorts

`func (o *NetdumpFilter) HasDstPorts() bool`

HasDstPorts returns a boolean if a field has been set.

### SetDstPortsNil

`func (o *NetdumpFilter) SetDstPortsNil(b bool)`

 SetDstPortsNil sets the value for DstPorts to be an explicit nil

### UnsetDstPorts
`func (o *NetdumpFilter) UnsetDstPorts()`

UnsetDstPorts ensures that no value is present for DstPorts, not even an explicit nil
### GetPorts

`func (o *NetdumpFilter) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NetdumpFilter) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NetdumpFilter) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *NetdumpFilter) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *NetdumpFilter) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *NetdumpFilter) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetProtocols

`func (o *NetdumpFilter) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *NetdumpFilter) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *NetdumpFilter) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *NetdumpFilter) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### SetProtocolsNil

`func (o *NetdumpFilter) SetProtocolsNil(b bool)`

 SetProtocolsNil sets the value for Protocols to be an explicit nil

### UnsetProtocols
`func (o *NetdumpFilter) UnsetProtocols()`

UnsetProtocols ensures that no value is present for Protocols, not even an explicit nil
### GetProcesses

`func (o *NetdumpFilter) GetProcesses() []string`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *NetdumpFilter) GetProcessesOk() (*[]string, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *NetdumpFilter) SetProcesses(v []string)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *NetdumpFilter) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### SetProcessesNil

`func (o *NetdumpFilter) SetProcessesNil(b bool)`

 SetProcessesNil sets the value for Processes to be an explicit nil

### UnsetProcesses
`func (o *NetdumpFilter) UnsetProcesses()`

UnsetProcesses ensures that no value is present for Processes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


