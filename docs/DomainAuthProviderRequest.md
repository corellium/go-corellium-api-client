# DomainAuthProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | **string** | Provider Type | 
**Enabled** | **bool** | Enabled/Disabled | 
**Label** | Pointer to **NullableString** | Login Button Text | [optional] 
**Config** | Pointer to [**OpenIDConfig**](OpenIDConfig.md) |  | [optional] 

## Methods

### NewDomainAuthProviderRequest

`func NewDomainAuthProviderRequest(providerType string, enabled bool, ) *DomainAuthProviderRequest`

NewDomainAuthProviderRequest instantiates a new DomainAuthProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAuthProviderRequestWithDefaults

`func NewDomainAuthProviderRequestWithDefaults() *DomainAuthProviderRequest`

NewDomainAuthProviderRequestWithDefaults instantiates a new DomainAuthProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *DomainAuthProviderRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DomainAuthProviderRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DomainAuthProviderRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetEnabled

`func (o *DomainAuthProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DomainAuthProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DomainAuthProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLabel

`func (o *DomainAuthProviderRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DomainAuthProviderRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DomainAuthProviderRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DomainAuthProviderRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *DomainAuthProviderRequest) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *DomainAuthProviderRequest) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetConfig

`func (o *DomainAuthProviderRequest) GetConfig() OpenIDConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DomainAuthProviderRequest) GetConfigOk() (*OpenIDConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DomainAuthProviderRequest) SetConfig(v OpenIDConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DomainAuthProviderRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


