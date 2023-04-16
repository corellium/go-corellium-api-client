# VpnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proxy** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Listeners** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnDefinition

`func NewVpnDefinition() *VpnDefinition`

NewVpnDefinition instantiates a new VpnDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnDefinitionWithDefaults

`func NewVpnDefinitionWithDefaults() *VpnDefinition`

NewVpnDefinitionWithDefaults instantiates a new VpnDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxy

`func (o *VpnDefinition) GetProxy() []map[string]interface{}`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *VpnDefinition) GetProxyOk() (*[]map[string]interface{}, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *VpnDefinition) SetProxy(v []map[string]interface{})`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *VpnDefinition) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *VpnDefinition) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *VpnDefinition) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetListeners

`func (o *VpnDefinition) GetListeners() []map[string]interface{}`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *VpnDefinition) GetListenersOk() (*[]map[string]interface{}, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *VpnDefinition) SetListeners(v []map[string]interface{})`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *VpnDefinition) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *VpnDefinition) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *VpnDefinition) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


