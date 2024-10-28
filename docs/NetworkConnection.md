# NetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | UUIDv4 | 
**Name** | **string** | User specified name for this network connection. e.g. My Network Connection | 
**Config** | Pointer to **map[string]interface{}** | An object containing network connection configuration data. Will vary based on network provider type. | [optional] 
**Provider** | Pointer to **NullableString** | One of the registered [network provider types](#get-/v1/network/providers) | [optional] 
**CreatedAt** | **string** | Created Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | 
**UpdatedAt** | **string** | Updated Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | 
**CreatedBy** | **string** | UUIDv4 of the user who created this record. | 
**UpdatedBy** | **string** | UUIDv4 of the user who last updated this record. | 

## Methods

### NewNetworkConnection

`func NewNetworkConnection(identifier string, name string, createdAt string, updatedAt string, createdBy string, updatedBy string, ) *NetworkConnection`

NewNetworkConnection instantiates a new NetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConnectionWithDefaults

`func NewNetworkConnectionWithDefaults() *NetworkConnection`

NewNetworkConnectionWithDefaults instantiates a new NetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *NetworkConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NetworkConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NetworkConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *NetworkConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkConnection) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *NetworkConnection) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkConnection) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkConnection) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkConnection) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *NetworkConnection) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *NetworkConnection) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetProvider

`func (o *NetworkConnection) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NetworkConnection) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NetworkConnection) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NetworkConnection) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *NetworkConnection) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *NetworkConnection) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetCreatedAt

`func (o *NetworkConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedBy

`func (o *NetworkConnection) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NetworkConnection) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NetworkConnection) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *NetworkConnection) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *NetworkConnection) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *NetworkConnection) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


