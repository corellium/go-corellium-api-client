# ActivityEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Unique identifier for the activity entry | [optional] 
**Resource** | Pointer to **NullableString** | Resource associated with the activity | [optional] 
**Timestamp** | Pointer to **NullableString** | Timestamp of the activity | [optional] 

## Methods

### NewActivityEntry

`func NewActivityEntry() *ActivityEntry`

NewActivityEntry instantiates a new ActivityEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityEntryWithDefaults

`func NewActivityEntryWithDefaults() *ActivityEntry`

NewActivityEntryWithDefaults instantiates a new ActivityEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ActivityEntry) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActivityEntry) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResource

`func (o *ActivityEntry) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ActivityEntry) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ActivityEntry) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ActivityEntry) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ActivityEntry) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ActivityEntry) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetTimestamp

`func (o *ActivityEntry) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityEntry) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityEntry) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ActivityEntry) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ActivityEntry) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ActivityEntry) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


