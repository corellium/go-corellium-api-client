# NetworkConnectionOffsetPaginationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | Total number of items. | 
**Count** | **float32** | The number of items returned in this result. | 
**Limit** | **float32** | The maximum number of items that could be returned for this result. | 
**Offset** | **float32** | That starting item 0-indexed. | 
**Results** | [**[]NetworkConnection**](NetworkConnection.md) | Array of network connections associated with this domain. | 

## Methods

### NewNetworkConnectionOffsetPaginationResult

`func NewNetworkConnectionOffsetPaginationResult(total float32, count float32, limit float32, offset float32, results []NetworkConnection, ) *NetworkConnectionOffsetPaginationResult`

NewNetworkConnectionOffsetPaginationResult instantiates a new NetworkConnectionOffsetPaginationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConnectionOffsetPaginationResultWithDefaults

`func NewNetworkConnectionOffsetPaginationResultWithDefaults() *NetworkConnectionOffsetPaginationResult`

NewNetworkConnectionOffsetPaginationResultWithDefaults instantiates a new NetworkConnectionOffsetPaginationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *NetworkConnectionOffsetPaginationResult) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NetworkConnectionOffsetPaginationResult) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NetworkConnectionOffsetPaginationResult) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetCount

`func (o *NetworkConnectionOffsetPaginationResult) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NetworkConnectionOffsetPaginationResult) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NetworkConnectionOffsetPaginationResult) SetCount(v float32)`

SetCount sets Count field to given value.


### GetLimit

`func (o *NetworkConnectionOffsetPaginationResult) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NetworkConnectionOffsetPaginationResult) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NetworkConnectionOffsetPaginationResult) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *NetworkConnectionOffsetPaginationResult) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NetworkConnectionOffsetPaginationResult) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NetworkConnectionOffsetPaginationResult) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetResults

`func (o *NetworkConnectionOffsetPaginationResult) GetResults() []NetworkConnection`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NetworkConnectionOffsetPaginationResult) GetResultsOk() (*[]NetworkConnection, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NetworkConnectionOffsetPaginationResult) SetResults(v []NetworkConnection)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


