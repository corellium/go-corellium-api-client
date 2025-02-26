# UserSnapshots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedWithUser** | [**SharedSnapshotsInfo**](SharedSnapshotsInfo.md) |  | 
**SharedByUser** | [**SharedSnapshotsInfo**](SharedSnapshotsInfo.md) |  | 

## Methods

### NewUserSnapshots

`func NewUserSnapshots(sharedWithUser SharedSnapshotsInfo, sharedByUser SharedSnapshotsInfo, ) *UserSnapshots`

NewUserSnapshots instantiates a new UserSnapshots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSnapshotsWithDefaults

`func NewUserSnapshotsWithDefaults() *UserSnapshots`

NewUserSnapshotsWithDefaults instantiates a new UserSnapshots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedWithUser

`func (o *UserSnapshots) GetSharedWithUser() SharedSnapshotsInfo`

GetSharedWithUser returns the SharedWithUser field if non-nil, zero value otherwise.

### GetSharedWithUserOk

`func (o *UserSnapshots) GetSharedWithUserOk() (*SharedSnapshotsInfo, bool)`

GetSharedWithUserOk returns a tuple with the SharedWithUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithUser

`func (o *UserSnapshots) SetSharedWithUser(v SharedSnapshotsInfo)`

SetSharedWithUser sets SharedWithUser field to given value.


### GetSharedByUser

`func (o *UserSnapshots) GetSharedByUser() SharedSnapshotsInfo`

GetSharedByUser returns the SharedByUser field if non-nil, zero value otherwise.

### GetSharedByUserOk

`func (o *UserSnapshots) GetSharedByUserOk() (*SharedSnapshotsInfo, bool)`

GetSharedByUserOk returns a tuple with the SharedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedByUser

`func (o *UserSnapshots) SetSharedByUser(v SharedSnapshotsInfo)`

SetSharedByUser sets SharedByUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


