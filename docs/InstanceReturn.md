# InstanceReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Instance ID | 
**State** | [**InstanceState**](InstanceState.md) |  | 

## Methods

### NewInstanceReturn

`func NewInstanceReturn(id string, state InstanceState, ) *InstanceReturn`

NewInstanceReturn instantiates a new InstanceReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceReturnWithDefaults

`func NewInstanceReturnWithDefaults() *InstanceReturn`

NewInstanceReturnWithDefaults instantiates a new InstanceReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceReturn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceReturn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceReturn) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *InstanceReturn) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceReturn) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceReturn) SetState(v InstanceState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


