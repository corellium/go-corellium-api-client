# ActivityExportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **NullableString** | Event to filter on | [optional] 
**OrderBy** | Pointer to **NullableString** | Column to order by | [optional] 
**Instance** | Pointer to **NullableString** | Instance ID | [optional] 
**Project** | Pointer to **NullableString** | Project ID | [optional] 
**Actor** | Pointer to **NullableString** | User ID | [optional] 
**Search** | Pointer to **NullableString** | Arbitrary search string against activity metadata | [optional] 
**To** | Pointer to **NullableString** | Include activities up to given ISO 8601 Date | [optional] 
**From** | Pointer to **NullableString** | Include activities from given ISO 8601 Date | [optional] 

## Methods

### NewActivityExportDto

`func NewActivityExportDto() *ActivityExportDto`

NewActivityExportDto instantiates a new ActivityExportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityExportDtoWithDefaults

`func NewActivityExportDtoWithDefaults() *ActivityExportDto`

NewActivityExportDtoWithDefaults instantiates a new ActivityExportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *ActivityExportDto) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ActivityExportDto) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ActivityExportDto) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ActivityExportDto) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ActivityExportDto) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ActivityExportDto) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetOrderBy

`func (o *ActivityExportDto) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ActivityExportDto) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ActivityExportDto) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ActivityExportDto) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *ActivityExportDto) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *ActivityExportDto) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetInstance

`func (o *ActivityExportDto) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ActivityExportDto) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ActivityExportDto) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ActivityExportDto) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *ActivityExportDto) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *ActivityExportDto) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetProject

`func (o *ActivityExportDto) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ActivityExportDto) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ActivityExportDto) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ActivityExportDto) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *ActivityExportDto) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *ActivityExportDto) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetActor

`func (o *ActivityExportDto) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ActivityExportDto) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ActivityExportDto) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *ActivityExportDto) HasActor() bool`

HasActor returns a boolean if a field has been set.

### SetActorNil

`func (o *ActivityExportDto) SetActorNil(b bool)`

 SetActorNil sets the value for Actor to be an explicit nil

### UnsetActor
`func (o *ActivityExportDto) UnsetActor()`

UnsetActor ensures that no value is present for Actor, not even an explicit nil
### GetSearch

`func (o *ActivityExportDto) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ActivityExportDto) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ActivityExportDto) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ActivityExportDto) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ActivityExportDto) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ActivityExportDto) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetTo

`func (o *ActivityExportDto) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ActivityExportDto) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ActivityExportDto) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ActivityExportDto) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ActivityExportDto) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ActivityExportDto) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetFrom

`func (o *ActivityExportDto) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ActivityExportDto) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ActivityExportDto) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ActivityExportDto) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ActivityExportDto) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ActivityExportDto) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


