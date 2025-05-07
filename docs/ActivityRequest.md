# ActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **NullableString** | Resource filter | [optional] 
**Limit** | Pointer to **NullableFloat32** | Number of entries to return, defaults to 10 | [optional] 
**OrderBy** | Pointer to **NullableString** | Sorting order (&#39;-createdAt&#39; or &#39;createdAt&#39;) | [optional] 
**Page** | Pointer to **NullableFloat32** | Paginated results. Must be a positive integer. If not provided, defaults to the first page. | [optional] 
**Instance** | Pointer to **NullableString** | Instance identifier used to filter instance resources | [optional] 
**Project** | Pointer to **NullableString** | Instance identifier used to filter instance resources | [optional] 
**Actor** | Pointer to **NullableString** | Actor identifier used to filter actor resources | [optional] 
**Search** | Pointer to **NullableString** | Last filter applied and is a fuzzy match on results | [optional] 
**To** | Pointer to **NullableString** | Date to filter to, keyed off of createdAt | [optional] 
**From** | Pointer to **NullableString** | Date to filter from, keyed off of createdAt | [optional] 

## Methods

### NewActivityRequest

`func NewActivityRequest() *ActivityRequest`

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

### HasResource

`func (o *ActivityRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ActivityRequest) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ActivityRequest) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
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

### HasLimit

`func (o *ActivityRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ActivityRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ActivityRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
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

### HasOrderBy

`func (o *ActivityRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *ActivityRequest) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *ActivityRequest) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
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
### GetProject

`func (o *ActivityRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ActivityRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ActivityRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ActivityRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *ActivityRequest) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *ActivityRequest) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetActor

`func (o *ActivityRequest) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ActivityRequest) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ActivityRequest) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *ActivityRequest) HasActor() bool`

HasActor returns a boolean if a field has been set.

### SetActorNil

`func (o *ActivityRequest) SetActorNil(b bool)`

 SetActorNil sets the value for Actor to be an explicit nil

### UnsetActor
`func (o *ActivityRequest) UnsetActor()`

UnsetActor ensures that no value is present for Actor, not even an explicit nil
### GetSearch

`func (o *ActivityRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ActivityRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ActivityRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ActivityRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ActivityRequest) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ActivityRequest) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetTo

`func (o *ActivityRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ActivityRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ActivityRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ActivityRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ActivityRequest) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ActivityRequest) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetFrom

`func (o *ActivityRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ActivityRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ActivityRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ActivityRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ActivityRequest) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ActivityRequest) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


