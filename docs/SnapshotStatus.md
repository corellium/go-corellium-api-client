# SnapshotStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to **NullableString** |  | [optional] 
**Created** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSnapshotStatus

`func NewSnapshotStatus() *SnapshotStatus`

NewSnapshotStatus instantiates a new SnapshotStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotStatusWithDefaults

`func NewSnapshotStatusWithDefaults() *SnapshotStatus`

NewSnapshotStatusWithDefaults instantiates a new SnapshotStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *SnapshotStatus) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *SnapshotStatus) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *SnapshotStatus) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *SnapshotStatus) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *SnapshotStatus) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *SnapshotStatus) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetCreated

`func (o *SnapshotStatus) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SnapshotStatus) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SnapshotStatus) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SnapshotStatus) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SnapshotStatus) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SnapshotStatus) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


