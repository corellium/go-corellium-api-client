# AgentAppStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | Pointer to **NullableBool** |  | [optional] 
**BundleID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgentAppStatus

`func NewAgentAppStatus() *AgentAppStatus`

NewAgentAppStatus instantiates a new AgentAppStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAppStatusWithDefaults

`func NewAgentAppStatusWithDefaults() *AgentAppStatus`

NewAgentAppStatusWithDefaults instantiates a new AgentAppStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *AgentAppStatus) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *AgentAppStatus) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *AgentAppStatus) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *AgentAppStatus) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### SetRunningNil

`func (o *AgentAppStatus) SetRunningNil(b bool)`

 SetRunningNil sets the value for Running to be an explicit nil

### UnsetRunning
`func (o *AgentAppStatus) UnsetRunning()`

UnsetRunning ensures that no value is present for Running, not even an explicit nil
### GetBundleID

`func (o *AgentAppStatus) GetBundleID() string`

GetBundleID returns the BundleID field if non-nil, zero value otherwise.

### GetBundleIDOk

`func (o *AgentAppStatus) GetBundleIDOk() (*string, bool)`

GetBundleIDOk returns a tuple with the BundleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleID

`func (o *AgentAppStatus) SetBundleID(v string)`

SetBundleID sets BundleID field to given value.

### HasBundleID

`func (o *AgentAppStatus) HasBundleID() bool`

HasBundleID returns a boolean if a field has been set.

### SetBundleIDNil

`func (o *AgentAppStatus) SetBundleIDNil(b bool)`

 SetBundleIDNil sets the value for BundleID to be an explicit nil

### UnsetBundleID
`func (o *AgentAppStatus) UnsetBundleID()`

UnsetBundleID ensures that no value is present for BundleID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


