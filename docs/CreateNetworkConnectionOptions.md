# CreateNetworkConnectionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | UUIDv4 | 
**Name** | **string** | User specified name for this network connection. e.g. My Network Connection | 
**Provider** | **string** | One of the registered [network provider types](#get-/v1/network/providers) | 
**Config** | Pointer to **map[string]interface{}** | An object containing network connection configuration data. Will vary based on network provider type. See examples. | [optional] 

## Methods

### NewCreateNetworkConnectionOptions

`func NewCreateNetworkConnectionOptions(identifier string, name string, provider string, ) *CreateNetworkConnectionOptions`

NewCreateNetworkConnectionOptions instantiates a new CreateNetworkConnectionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkConnectionOptionsWithDefaults

`func NewCreateNetworkConnectionOptionsWithDefaults() *CreateNetworkConnectionOptions`

NewCreateNetworkConnectionOptionsWithDefaults instantiates a new CreateNetworkConnectionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CreateNetworkConnectionOptions) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateNetworkConnectionOptions) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateNetworkConnectionOptions) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *CreateNetworkConnectionOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkConnectionOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkConnectionOptions) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *CreateNetworkConnectionOptions) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateNetworkConnectionOptions) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateNetworkConnectionOptions) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetConfig

`func (o *CreateNetworkConnectionOptions) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworkConnectionOptions) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworkConnectionOptions) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNetworkConnectionOptions) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *CreateNetworkConnectionOptions) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *CreateNetworkConnectionOptions) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


