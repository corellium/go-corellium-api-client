# InstanceServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpn** | Pointer to [**VpnDefinition**](VpnDefinition.md) |  | [optional] 

## Methods

### NewInstanceServices

`func NewInstanceServices() *InstanceServices`

NewInstanceServices instantiates a new InstanceServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceServicesWithDefaults

`func NewInstanceServicesWithDefaults() *InstanceServices`

NewInstanceServicesWithDefaults instantiates a new InstanceServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpn

`func (o *InstanceServices) GetVpn() VpnDefinition`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InstanceServices) GetVpnOk() (*VpnDefinition, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InstanceServices) SetVpn(v VpnDefinition)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InstanceServices) HasVpn() bool`

HasVpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


