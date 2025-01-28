# ActivityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ActivityEntry**](ActivityEntry.md) | List of activity entries | [optional] 
**Total** | Pointer to **NullableFloat32** | Total number of entries | [optional] 

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

### GetItems

`func (o *ActivityList) GetItems() []ActivityEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ActivityList) GetItemsOk() (*[]ActivityEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ActivityList) SetItems(v []ActivityEntry)`

SetItems sets Items field to given value.

### HasItems

`func (o *ActivityList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *ActivityList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *ActivityList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotal

`func (o *ActivityList) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ActivityList) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ActivityList) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ActivityList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ActivityList) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ActivityList) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


