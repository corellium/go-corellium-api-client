# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Id** | Pointer to **NullableString** | Image ID | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** | Type of image | [optional] 
**Filename** | Pointer to **NullableString** |  | [optional] 
**Uniqueid** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableFloat32** |  | [optional] 
**Project** | Pointer to **NullableString** | project ID | [optional] 
**CreatedAt** | Pointer to **NullableTime** | When Image was created | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | When Image was last updated | [optional] 

## Methods

### NewImage

`func NewImage(status string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Image) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Image) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Image) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Image) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Image) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Image) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Image) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Image) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Image) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Image) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Image) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Image) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Image) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Image) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Image) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFilename

`func (o *Image) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Image) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Image) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Image) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *Image) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *Image) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetUniqueid

`func (o *Image) GetUniqueid() string`

GetUniqueid returns the Uniqueid field if non-nil, zero value otherwise.

### GetUniqueidOk

`func (o *Image) GetUniqueidOk() (*string, bool)`

GetUniqueidOk returns a tuple with the Uniqueid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueid

`func (o *Image) SetUniqueid(v string)`

SetUniqueid sets Uniqueid field to given value.

### HasUniqueid

`func (o *Image) HasUniqueid() bool`

HasUniqueid returns a boolean if a field has been set.

### SetUniqueidNil

`func (o *Image) SetUniqueidNil(b bool)`

 SetUniqueidNil sets the value for Uniqueid to be an explicit nil

### UnsetUniqueid
`func (o *Image) UnsetUniqueid()`

UnsetUniqueid ensures that no value is present for Uniqueid, not even an explicit nil
### GetSize

`func (o *Image) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Image) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Image) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Image) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Image) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Image) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetProject

`func (o *Image) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Image) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Image) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Image) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *Image) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *Image) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetCreatedAt

`func (o *Image) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Image) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Image) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Image) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Image) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Image) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Image) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Image) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Image) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Image) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Image) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Image) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


