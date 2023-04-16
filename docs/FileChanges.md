# FileChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **NullableString** | Optional New path | [optional] 
**Mode** | Pointer to **NullableFloat32** | Optional New Mode | [optional] 
**Uid** | Pointer to **NullableFloat32** | Optional New Owner UID | [optional] 
**Gid** | Pointer to **NullableFloat32** | Optional New Group GID | [optional] 

## Methods

### NewFileChanges

`func NewFileChanges() *FileChanges`

NewFileChanges instantiates a new FileChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileChangesWithDefaults

`func NewFileChangesWithDefaults() *FileChanges`

NewFileChangesWithDefaults instantiates a new FileChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FileChanges) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileChanges) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileChanges) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileChanges) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *FileChanges) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *FileChanges) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetMode

`func (o *FileChanges) GetMode() float32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FileChanges) GetModeOk() (*float32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FileChanges) SetMode(v float32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FileChanges) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *FileChanges) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *FileChanges) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetUid

`func (o *FileChanges) GetUid() float32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FileChanges) GetUidOk() (*float32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FileChanges) SetUid(v float32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *FileChanges) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *FileChanges) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *FileChanges) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetGid

`func (o *FileChanges) GetGid() float32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *FileChanges) GetGidOk() (*float32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *FileChanges) SetGid(v float32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *FileChanges) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *FileChanges) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *FileChanges) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


