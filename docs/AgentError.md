# AgentError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The error encountered by the agent | 
**ErrorID** | **string** | AgentError | 
**OriginalError** | Pointer to **map[string]interface{}** | The full error encountered by the agent | [optional] 

## Methods

### NewAgentError

`func NewAgentError(error_ string, errorID string, ) *AgentError`

NewAgentError instantiates a new AgentError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentErrorWithDefaults

`func NewAgentErrorWithDefaults() *AgentError`

NewAgentErrorWithDefaults instantiates a new AgentError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AgentError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorID

`func (o *AgentError) GetErrorID() string`

GetErrorID returns the ErrorID field if non-nil, zero value otherwise.

### GetErrorIDOk

`func (o *AgentError) GetErrorIDOk() (*string, bool)`

GetErrorIDOk returns a tuple with the ErrorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorID

`func (o *AgentError) SetErrorID(v string)`

SetErrorID sets ErrorID field to given value.


### GetOriginalError

`func (o *AgentError) GetOriginalError() map[string]interface{}`

GetOriginalError returns the OriginalError field if non-nil, zero value otherwise.

### GetOriginalErrorOk

`func (o *AgentError) GetOriginalErrorOk() (*map[string]interface{}, bool)`

GetOriginalErrorOk returns a tuple with the OriginalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalError

`func (o *AgentError) SetOriginalError(v map[string]interface{})`

SetOriginalError sets OriginalError field to given value.

### HasOriginalError

`func (o *AgentError) HasOriginalError() bool`

HasOriginalError returns a boolean if a field has been set.

### SetOriginalErrorNil

`func (o *AgentError) SetOriginalErrorNil(b bool)`

 SetOriginalErrorNil sets the value for OriginalError to be an explicit nil

### UnsetOriginalError
`func (o *AgentError) UnsetOriginalError()`

UnsetOriginalError ensures that no value is present for OriginalError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


