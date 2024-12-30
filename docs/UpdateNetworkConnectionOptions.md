# UpdateNetworkConnectionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User specified name for this network connection. e.g. My Network Connection | 
**Config** | Pointer to **map[string]interface{}** | An object containing network connection configuration data. Will vary based on network provider type. See examples. | [optional] 

## Methods

### NewUpdateNetworkConnectionOptions

`func NewUpdateNetworkConnectionOptions(name string, ) *UpdateNetworkConnectionOptions`

NewUpdateNetworkConnectionOptions instantiates a new UpdateNetworkConnectionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkConnectionOptionsWithDefaults

`func NewUpdateNetworkConnectionOptionsWithDefaults() *UpdateNetworkConnectionOptions`

NewUpdateNetworkConnectionOptionsWithDefaults instantiates a new UpdateNetworkConnectionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkConnectionOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkConnectionOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkConnectionOptions) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *UpdateNetworkConnectionOptions) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNetworkConnectionOptions) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNetworkConnectionOptions) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNetworkConnectionOptions) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *UpdateNetworkConnectionOptions) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *UpdateNetworkConnectionOptions) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


