# SharedSnapshotsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UUID of the snapshot | 
**Name** | **string** | Snapshot name | 
**Model** | **string** | Device model | 
**SharedOn** | **float32** | UNIX Date of when the snapshot was first shared with member | 
**SharedWithMember** | Pointer to **NullableString** | The member who the snapshot was shared with. Only present in sharedWithUser | [optional] 
**SharedBy** | Pointer to [**SnapshotOwner**](SnapshotOwner.md) |  | [optional] 

## Methods

### NewSharedSnapshotsInfo

`func NewSharedSnapshotsInfo(id string, name string, model string, sharedOn float32, ) *SharedSnapshotsInfo`

NewSharedSnapshotsInfo instantiates a new SharedSnapshotsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedSnapshotsInfoWithDefaults

`func NewSharedSnapshotsInfoWithDefaults() *SharedSnapshotsInfo`

NewSharedSnapshotsInfoWithDefaults instantiates a new SharedSnapshotsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedSnapshotsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedSnapshotsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedSnapshotsInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SharedSnapshotsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedSnapshotsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedSnapshotsInfo) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *SharedSnapshotsInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SharedSnapshotsInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SharedSnapshotsInfo) SetModel(v string)`

SetModel sets Model field to given value.


### GetSharedOn

`func (o *SharedSnapshotsInfo) GetSharedOn() float32`

GetSharedOn returns the SharedOn field if non-nil, zero value otherwise.

### GetSharedOnOk

`func (o *SharedSnapshotsInfo) GetSharedOnOk() (*float32, bool)`

GetSharedOnOk returns a tuple with the SharedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOn

`func (o *SharedSnapshotsInfo) SetSharedOn(v float32)`

SetSharedOn sets SharedOn field to given value.


### GetSharedWithMember

`func (o *SharedSnapshotsInfo) GetSharedWithMember() string`

GetSharedWithMember returns the SharedWithMember field if non-nil, zero value otherwise.

### GetSharedWithMemberOk

`func (o *SharedSnapshotsInfo) GetSharedWithMemberOk() (*string, bool)`

GetSharedWithMemberOk returns a tuple with the SharedWithMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithMember

`func (o *SharedSnapshotsInfo) SetSharedWithMember(v string)`

SetSharedWithMember sets SharedWithMember field to given value.

### HasSharedWithMember

`func (o *SharedSnapshotsInfo) HasSharedWithMember() bool`

HasSharedWithMember returns a boolean if a field has been set.

### SetSharedWithMemberNil

`func (o *SharedSnapshotsInfo) SetSharedWithMemberNil(b bool)`

 SetSharedWithMemberNil sets the value for SharedWithMember to be an explicit nil

### UnsetSharedWithMember
`func (o *SharedSnapshotsInfo) UnsetSharedWithMember()`

UnsetSharedWithMember ensures that no value is present for SharedWithMember, not even an explicit nil
### GetSharedBy

`func (o *SharedSnapshotsInfo) GetSharedBy() SnapshotOwner`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *SharedSnapshotsInfo) GetSharedByOk() (*SnapshotOwner, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *SharedSnapshotsInfo) SetSharedBy(v SnapshotOwner)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *SharedSnapshotsInfo) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


