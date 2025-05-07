# ActivityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Activity**](Activity.md) | List of activities | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewActivityList

`func NewActivityList() *ActivityList`

NewActivityList instantiates a new ActivityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityListWithDefaults

`func NewActivityListWithDefaults() *ActivityList`

NewActivityListWithDefaults instantiates a new ActivityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ActivityList) GetData() []Activity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActivityList) GetDataOk() (*[]Activity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActivityList) SetData(v []Activity)`

SetData sets Data field to given value.

### HasData

`func (o *ActivityList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ActivityList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ActivityList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPagination

`func (o *ActivityList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ActivityList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ActivityList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ActivityList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


