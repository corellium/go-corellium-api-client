# NetworkStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Network Provider | 
**Name** | **string** | User defined name | 
**State** | **string** | Current state of the network connection | 
**Params** | **map[string]interface{}** | Additional state information | 

## Methods

### NewNetworkStatusResponse

`func NewNetworkStatusResponse(type_ string, name string, state string, params map[string]interface{}, ) *NetworkStatusResponse`

NewNetworkStatusResponse instantiates a new NetworkStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkStatusResponseWithDefaults

`func NewNetworkStatusResponseWithDefaults() *NetworkStatusResponse`

NewNetworkStatusResponseWithDefaults instantiates a new NetworkStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkStatusResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkStatusResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkStatusResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NetworkStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkStatusResponse) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *NetworkStatusResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetworkStatusResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetworkStatusResponse) SetState(v string)`

SetState sets State field to given value.


### GetParams

`func (o *NetworkStatusResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NetworkStatusResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NetworkStatusResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


