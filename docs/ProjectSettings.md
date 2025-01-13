# ProjectSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternetAccess** | **bool** |  | 
**Connection** | Pointer to **NullableString** | UUIDv4 network connection identifier or null for no vpn connection | [optional] 
**Dhcp** | **bool** |  | 

## Methods

### NewProjectSettings

`func NewProjectSettings(internetAccess bool, dhcp bool, ) *ProjectSettings`

NewProjectSettings instantiates a new ProjectSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSettingsWithDefaults

`func NewProjectSettingsWithDefaults() *ProjectSettings`

NewProjectSettingsWithDefaults instantiates a new ProjectSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternetAccess

`func (o *ProjectSettings) GetInternetAccess() bool`

GetInternetAccess returns the InternetAccess field if non-nil, zero value otherwise.

### GetInternetAccessOk

`func (o *ProjectSettings) GetInternetAccessOk() (*bool, bool)`

GetInternetAccessOk returns a tuple with the InternetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAccess

`func (o *ProjectSettings) SetInternetAccess(v bool)`

SetInternetAccess sets InternetAccess field to given value.


### GetConnection

`func (o *ProjectSettings) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ProjectSettings) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ProjectSettings) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ProjectSettings) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### SetConnectionNil

`func (o *ProjectSettings) SetConnectionNil(b bool)`

 SetConnectionNil sets the value for Connection to be an explicit nil

### UnsetConnection
`func (o *ProjectSettings) UnsetConnection()`

UnsetConnection ensures that no value is present for Connection, not even an explicit nil
### GetDhcp

`func (o *ProjectSettings) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *ProjectSettings) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *ProjectSettings) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


