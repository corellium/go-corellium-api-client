# ProjectNetworkState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | State of the project network connection. One of: disconnected, disconnecting, connecting, connected, error | 
**Details** | Pointer to **map[string]interface{}** | Additional parameters for the network state. | [optional] 

## Methods

### NewProjectNetworkState

`func NewProjectNetworkState(state string, ) *ProjectNetworkState`

NewProjectNetworkState instantiates a new ProjectNetworkState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectNetworkStateWithDefaults

`func NewProjectNetworkStateWithDefaults() *ProjectNetworkState`

NewProjectNetworkStateWithDefaults instantiates a new ProjectNetworkState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ProjectNetworkState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectNetworkState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectNetworkState) SetState(v string)`

SetState sets State field to given value.


### GetDetails

`func (o *ProjectNetworkState) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProjectNetworkState) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProjectNetworkState) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProjectNetworkState) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ProjectNetworkState) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ProjectNetworkState) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


