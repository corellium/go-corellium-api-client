# AgentAppsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]AgentApp**](AgentApp.md) |  | [optional] 
**Frontmost** | Pointer to **NullableString** | bundleID of frontmost app or empty string if none are frontmost | [optional] 

## Methods

### NewAgentAppsList

`func NewAgentAppsList() *AgentAppsList`

NewAgentAppsList instantiates a new AgentAppsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAppsListWithDefaults

`func NewAgentAppsListWithDefaults() *AgentAppsList`

NewAgentAppsListWithDefaults instantiates a new AgentAppsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *AgentAppsList) GetApps() []AgentApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AgentAppsList) GetAppsOk() (*[]AgentApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AgentAppsList) SetApps(v []AgentApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AgentAppsList) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *AgentAppsList) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *AgentAppsList) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetFrontmost

`func (o *AgentAppsList) GetFrontmost() string`

GetFrontmost returns the Frontmost field if non-nil, zero value otherwise.

### GetFrontmostOk

`func (o *AgentAppsList) GetFrontmostOk() (*string, bool)`

GetFrontmostOk returns a tuple with the Frontmost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontmost

`func (o *AgentAppsList) SetFrontmost(v string)`

SetFrontmost sets Frontmost field to given value.

### HasFrontmost

`func (o *AgentAppsList) HasFrontmost() bool`

HasFrontmost returns a boolean if a field has been set.

### SetFrontmostNil

`func (o *AgentAppsList) SetFrontmostNil(b bool)`

 SetFrontmostNil sets the value for Frontmost to be an explicit nil

### UnsetFrontmost
`func (o *AgentAppsList) UnsetFrontmost()`

UnsetFrontmost ensures that no value is present for Frontmost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


