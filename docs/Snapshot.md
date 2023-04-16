# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Snapshot ID | [optional] 
**Name** | Pointer to **NullableString** | Snapshot name | [optional] 
**Instance** | Pointer to **NullableString** | Instance that this snapshot is of | [optional] 
**Status** | Pointer to [**SnapshotStatus**](SnapshotStatus.md) |  | [optional] 
**Date** | Pointer to **NullableFloat32** | UNIX Date that the snapshot was created | [optional] 
**Fresh** | Pointer to **NullableBool** |  | [optional] 
**Live** | Pointer to **NullableBool** | Live snapshot (included state and memory) | [optional] 
**Local** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Snapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Snapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Snapshot) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Snapshot) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Snapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Snapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Snapshot) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Snapshot) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInstance

`func (o *Snapshot) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Snapshot) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Snapshot) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *Snapshot) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *Snapshot) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *Snapshot) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetStatus

`func (o *Snapshot) GetStatus() SnapshotStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Snapshot) GetStatusOk() (*SnapshotStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Snapshot) SetStatus(v SnapshotStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Snapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDate

`func (o *Snapshot) GetDate() float32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Snapshot) GetDateOk() (*float32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Snapshot) SetDate(v float32)`

SetDate sets Date field to given value.

### HasDate

`func (o *Snapshot) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *Snapshot) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *Snapshot) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetFresh

`func (o *Snapshot) GetFresh() bool`

GetFresh returns the Fresh field if non-nil, zero value otherwise.

### GetFreshOk

`func (o *Snapshot) GetFreshOk() (*bool, bool)`

GetFreshOk returns a tuple with the Fresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFresh

`func (o *Snapshot) SetFresh(v bool)`

SetFresh sets Fresh field to given value.

### HasFresh

`func (o *Snapshot) HasFresh() bool`

HasFresh returns a boolean if a field has been set.

### SetFreshNil

`func (o *Snapshot) SetFreshNil(b bool)`

 SetFreshNil sets the value for Fresh to be an explicit nil

### UnsetFresh
`func (o *Snapshot) UnsetFresh()`

UnsetFresh ensures that no value is present for Fresh, not even an explicit nil
### GetLive

`func (o *Snapshot) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *Snapshot) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *Snapshot) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *Snapshot) HasLive() bool`

HasLive returns a boolean if a field has been set.

### SetLiveNil

`func (o *Snapshot) SetLiveNil(b bool)`

 SetLiveNil sets the value for Live to be an explicit nil

### UnsetLive
`func (o *Snapshot) UnsetLive()`

UnsetLive ensures that no value is present for Live, not even an explicit nil
### GetLocal

`func (o *Snapshot) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *Snapshot) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *Snapshot) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *Snapshot) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### SetLocalNil

`func (o *Snapshot) SetLocalNil(b bool)`

 SetLocalNil sets the value for Local to be an explicit nil

### UnsetLocal
`func (o *Snapshot) UnsetLocal()`

UnsetLocal ensures that no value is present for Local, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


