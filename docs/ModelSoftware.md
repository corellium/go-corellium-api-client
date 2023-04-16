# ModelSoftware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**BoardConfig** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Cpid** | Pointer to **NullableFloat32** |  | [optional] 
**Bdid** | Pointer to **NullableFloat32** |  | [optional] 
**Firmwares** | Pointer to [**[]Firmware**](Firmware.md) |  | [optional] 

## Methods

### NewModelSoftware

`func NewModelSoftware() *ModelSoftware`

NewModelSoftware instantiates a new ModelSoftware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSoftwareWithDefaults

`func NewModelSoftwareWithDefaults() *ModelSoftware`

NewModelSoftwareWithDefaults instantiates a new ModelSoftware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelSoftware) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSoftware) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSoftware) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSoftware) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelSoftware) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelSoftware) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBoardConfig

`func (o *ModelSoftware) GetBoardConfig() string`

GetBoardConfig returns the BoardConfig field if non-nil, zero value otherwise.

### GetBoardConfigOk

`func (o *ModelSoftware) GetBoardConfigOk() (*string, bool)`

GetBoardConfigOk returns a tuple with the BoardConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardConfig

`func (o *ModelSoftware) SetBoardConfig(v string)`

SetBoardConfig sets BoardConfig field to given value.

### HasBoardConfig

`func (o *ModelSoftware) HasBoardConfig() bool`

HasBoardConfig returns a boolean if a field has been set.

### SetBoardConfigNil

`func (o *ModelSoftware) SetBoardConfigNil(b bool)`

 SetBoardConfigNil sets the value for BoardConfig to be an explicit nil

### UnsetBoardConfig
`func (o *ModelSoftware) UnsetBoardConfig()`

UnsetBoardConfig ensures that no value is present for BoardConfig, not even an explicit nil
### GetPlatform

`func (o *ModelSoftware) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ModelSoftware) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ModelSoftware) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ModelSoftware) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ModelSoftware) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ModelSoftware) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetCpid

`func (o *ModelSoftware) GetCpid() float32`

GetCpid returns the Cpid field if non-nil, zero value otherwise.

### GetCpidOk

`func (o *ModelSoftware) GetCpidOk() (*float32, bool)`

GetCpidOk returns a tuple with the Cpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpid

`func (o *ModelSoftware) SetCpid(v float32)`

SetCpid sets Cpid field to given value.

### HasCpid

`func (o *ModelSoftware) HasCpid() bool`

HasCpid returns a boolean if a field has been set.

### SetCpidNil

`func (o *ModelSoftware) SetCpidNil(b bool)`

 SetCpidNil sets the value for Cpid to be an explicit nil

### UnsetCpid
`func (o *ModelSoftware) UnsetCpid()`

UnsetCpid ensures that no value is present for Cpid, not even an explicit nil
### GetBdid

`func (o *ModelSoftware) GetBdid() float32`

GetBdid returns the Bdid field if non-nil, zero value otherwise.

### GetBdidOk

`func (o *ModelSoftware) GetBdidOk() (*float32, bool)`

GetBdidOk returns a tuple with the Bdid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdid

`func (o *ModelSoftware) SetBdid(v float32)`

SetBdid sets Bdid field to given value.

### HasBdid

`func (o *ModelSoftware) HasBdid() bool`

HasBdid returns a boolean if a field has been set.

### SetBdidNil

`func (o *ModelSoftware) SetBdidNil(b bool)`

 SetBdidNil sets the value for Bdid to be an explicit nil

### UnsetBdid
`func (o *ModelSoftware) UnsetBdid()`

UnsetBdid ensures that no value is present for Bdid, not even an explicit nil
### GetFirmwares

`func (o *ModelSoftware) GetFirmwares() []Firmware`

GetFirmwares returns the Firmwares field if non-nil, zero value otherwise.

### GetFirmwaresOk

`func (o *ModelSoftware) GetFirmwaresOk() (*[]Firmware, bool)`

GetFirmwaresOk returns a tuple with the Firmwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwares

`func (o *ModelSoftware) SetFirmwares(v []Firmware)`

SetFirmwares sets Firmwares field to given value.

### HasFirmwares

`func (o *ModelSoftware) HasFirmwares() bool`

HasFirmwares returns a boolean if a field has been set.

### SetFirmwaresNil

`func (o *ModelSoftware) SetFirmwaresNil(b bool)`

 SetFirmwaresNil sets the value for Firmwares to be an explicit nil

### UnsetFirmwares
`func (o *ModelSoftware) UnsetFirmwares()`

UnsetFirmwares ensures that no value is present for Firmwares, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


