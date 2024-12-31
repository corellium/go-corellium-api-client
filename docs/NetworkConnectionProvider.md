# NetworkConnectionProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human friendly name for Network Connection Provider. | 
**Type** | **string** | Internal identifier for Network Connection Provider. | 

## Methods

### NewNetworkConnectionProvider

`func NewNetworkConnectionProvider(name string, type_ string, ) *NetworkConnectionProvider`

NewNetworkConnectionProvider instantiates a new NetworkConnectionProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConnectionProviderWithDefaults

`func NewNetworkConnectionProviderWithDefaults() *NetworkConnectionProvider`

NewNetworkConnectionProviderWithDefaults instantiates a new NetworkConnectionProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkConnectionProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkConnectionProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkConnectionProvider) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NetworkConnectionProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkConnectionProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkConnectionProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


