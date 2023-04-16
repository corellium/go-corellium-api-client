# PatchInstanceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | the name of the device | [optional] 
**State** | Pointer to **NullableString** | the desired state of the device | [optional] 
**BootOptions** | Pointer to [**InstanceBootOptions**](InstanceBootOptions.md) |  | [optional] 
**Proxy** | Pointer to [**[]ProxyConfig**](ProxyConfig.md) |  | [optional] 

## Methods

### NewPatchInstanceOptions

`func NewPatchInstanceOptions() *PatchInstanceOptions`

NewPatchInstanceOptions instantiates a new PatchInstanceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchInstanceOptionsWithDefaults

`func NewPatchInstanceOptionsWithDefaults() *PatchInstanceOptions`

NewPatchInstanceOptionsWithDefaults instantiates a new PatchInstanceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchInstanceOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchInstanceOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchInstanceOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchInstanceOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchInstanceOptions) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchInstanceOptions) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *PatchInstanceOptions) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchInstanceOptions) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchInstanceOptions) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatchInstanceOptions) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *PatchInstanceOptions) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *PatchInstanceOptions) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetBootOptions

`func (o *PatchInstanceOptions) GetBootOptions() InstanceBootOptions`

GetBootOptions returns the BootOptions field if non-nil, zero value otherwise.

### GetBootOptionsOk

`func (o *PatchInstanceOptions) GetBootOptionsOk() (*InstanceBootOptions, bool)`

GetBootOptionsOk returns a tuple with the BootOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptions

`func (o *PatchInstanceOptions) SetBootOptions(v InstanceBootOptions)`

SetBootOptions sets BootOptions field to given value.

### HasBootOptions

`func (o *PatchInstanceOptions) HasBootOptions() bool`

HasBootOptions returns a boolean if a field has been set.

### GetProxy

`func (o *PatchInstanceOptions) GetProxy() []ProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PatchInstanceOptions) GetProxyOk() (*[]ProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PatchInstanceOptions) SetProxy(v []ProxyConfig)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PatchInstanceOptions) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PatchInstanceOptions) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PatchInstanceOptions) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


