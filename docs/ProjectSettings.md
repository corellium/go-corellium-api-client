# ProjectSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternetAccess** | Pointer to **NullableBool** |  | [optional] 
**Dhcp** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProjectSettings

`func NewProjectSettings() *ProjectSettings`

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

### HasInternetAccess

`func (o *ProjectSettings) HasInternetAccess() bool`

HasInternetAccess returns a boolean if a field has been set.

### SetInternetAccessNil

`func (o *ProjectSettings) SetInternetAccessNil(b bool)`

 SetInternetAccessNil sets the value for InternetAccess to be an explicit nil

### UnsetInternetAccess
`func (o *ProjectSettings) UnsetInternetAccess()`

UnsetInternetAccess ensures that no value is present for InternetAccess, not even an explicit nil
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

### HasDhcp

`func (o *ProjectSettings) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### SetDhcpNil

`func (o *ProjectSettings) SetDhcpNil(b bool)`

 SetDhcpNil sets the value for Dhcp to be an explicit nil

### UnsetDhcp
`func (o *ProjectSettings) UnsetDhcp()`

UnsetDhcp ensures that no value is present for Dhcp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


