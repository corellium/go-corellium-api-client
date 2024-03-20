# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Snapshot ID | 
**Name** | **string** | Snapshot name | 
**Instance** | **string** | Instance that this snapshot is of | 
**Status** | [**SnapshotStatus**](SnapshotStatus.md) |  | 
**Date** | **float32** | UNIX Date that the snapshot was created | 
**Fresh** | **bool** |  | 
**Live** | **bool** | Live snapshot (included state and memory) | 
**Local** | **bool** |  | 
**Sharing** | Pointer to [**SnapshotSharing**](SnapshotSharing.md) |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot(id string, name string, instance string, status SnapshotStatus, date float32, fresh bool, live bool, local bool, ) *Snapshot`

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


### GetSharing

`func (o *Snapshot) GetSharing() SnapshotSharing`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *Snapshot) GetSharingOk() (*SnapshotSharing, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *Snapshot) SetSharing(v SnapshotSharing)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *Snapshot) HasSharing() bool`

HasSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


