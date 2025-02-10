# ActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | **string** | Resource filter | 
**Limit** | **float32** | Number of entries to return | 
**OrderBy** | **string** | Sorting order (&#39;-timestamp&#39; or &#39;timestamp&#39;) | 
**Page** | Pointer to **NullableFloat32** | Paginated results. Must be a positive integer. If not provided, defaults to the first page. | [optional] 
**Instance** | Pointer to **NullableString** | Instance identifier used to filter instance resources | [optional] 

## Methods

### NewActivityRequest

`func NewActivityRequest(resource string, limit float32, orderBy string, ) *ActivityRequest`

NewActivityRequest instantiates a new ActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityRequestWithDefaults

`func NewActivityRequestWithDefaults() *ActivityRequest`

NewActivityRequestWithDefaults instantiates a new ActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *ActivityRequest) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ActivityRequest) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ActivityRequest) SetResource(v string)`

SetResource sets Resource field to given value.


### GetLimit

`func (o *ActivityRequest) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ActivityRequest) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ActivityRequest) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetOrderBy

`func (o *ActivityRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ActivityRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ActivityRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.


### GetPage

`func (o *ActivityRequest) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ActivityRequest) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ActivityRequest) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ActivityRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *ActivityRequest) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *ActivityRequest) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil
### GetInstance

`func (o *ActivityRequest) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ActivityRequest) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ActivityRequest) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ActivityRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *ActivityRequest) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *ActivityRequest) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


