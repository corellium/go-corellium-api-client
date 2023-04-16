# ProxyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePort** | Pointer to **NullableFloat32** | The device port to use for proxying. | [optional] 
**FirstAvailable** | Pointer to **NullableBool** | If &#x60;true&#x60;, the first available port will be used if &#x60;devicePort&#x60; is not available. | [optional] 
**Expose** | Pointer to **NullableBool** | If &#x60;true&#x60;, the proxy will be exposed to the external interface. | [optional] 
**RouterPort** | Pointer to **NullableFloat32** | The router port to use for proxying. | [optional] 

## Methods

### NewProxyConfig

`func NewProxyConfig() *ProxyConfig`

NewProxyConfig instantiates a new ProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyConfigWithDefaults

`func NewProxyConfigWithDefaults() *ProxyConfig`

NewProxyConfigWithDefaults instantiates a new ProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePort

`func (o *ProxyConfig) GetDevicePort() float32`

GetDevicePort returns the DevicePort field if non-nil, zero value otherwise.

### GetDevicePortOk

`func (o *ProxyConfig) GetDevicePortOk() (*float32, bool)`

GetDevicePortOk returns a tuple with the DevicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePort

`func (o *ProxyConfig) SetDevicePort(v float32)`

SetDevicePort sets DevicePort field to given value.

### HasDevicePort

`func (o *ProxyConfig) HasDevicePort() bool`

HasDevicePort returns a boolean if a field has been set.

### SetDevicePortNil

`func (o *ProxyConfig) SetDevicePortNil(b bool)`

 SetDevicePortNil sets the value for DevicePort to be an explicit nil

### UnsetDevicePort
`func (o *ProxyConfig) UnsetDevicePort()`

UnsetDevicePort ensures that no value is present for DevicePort, not even an explicit nil
### GetFirstAvailable

`func (o *ProxyConfig) GetFirstAvailable() bool`

GetFirstAvailable returns the FirstAvailable field if non-nil, zero value otherwise.

### GetFirstAvailableOk

`func (o *ProxyConfig) GetFirstAvailableOk() (*bool, bool)`

GetFirstAvailableOk returns a tuple with the FirstAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailable

`func (o *ProxyConfig) SetFirstAvailable(v bool)`

SetFirstAvailable sets FirstAvailable field to given value.

### HasFirstAvailable

`func (o *ProxyConfig) HasFirstAvailable() bool`

HasFirstAvailable returns a boolean if a field has been set.

### SetFirstAvailableNil

`func (o *ProxyConfig) SetFirstAvailableNil(b bool)`

 SetFirstAvailableNil sets the value for FirstAvailable to be an explicit nil

### UnsetFirstAvailable
`func (o *ProxyConfig) UnsetFirstAvailable()`

UnsetFirstAvailable ensures that no value is present for FirstAvailable, not even an explicit nil
### GetExpose

`func (o *ProxyConfig) GetExpose() bool`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *ProxyConfig) GetExposeOk() (*bool, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *ProxyConfig) SetExpose(v bool)`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *ProxyConfig) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### SetExposeNil

`func (o *ProxyConfig) SetExposeNil(b bool)`

 SetExposeNil sets the value for Expose to be an explicit nil

### UnsetExpose
`func (o *ProxyConfig) UnsetExpose()`

UnsetExpose ensures that no value is present for Expose, not even an explicit nil
### GetRouterPort

`func (o *ProxyConfig) GetRouterPort() float32`

GetRouterPort returns the RouterPort field if non-nil, zero value otherwise.

### GetRouterPortOk

`func (o *ProxyConfig) GetRouterPortOk() (*float32, bool)`

GetRouterPortOk returns a tuple with the RouterPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPort

`func (o *ProxyConfig) SetRouterPort(v float32)`

SetRouterPort sets RouterPort field to given value.

### HasRouterPort

`func (o *ProxyConfig) HasRouterPort() bool`

HasRouterPort returns a boolean if a field has been set.

### SetRouterPortNil

`func (o *ProxyConfig) SetRouterPortNil(b bool)`

 SetRouterPortNil sets the value for RouterPort to be an explicit nil

### UnsetRouterPort
`func (o *ProxyConfig) UnsetRouterPort()`

UnsetRouterPort ensures that no value is present for RouterPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


