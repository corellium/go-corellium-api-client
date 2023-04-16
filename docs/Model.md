# Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Flavor** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Model** | **string** |  | 
**BoardConfig** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Cpid** | Pointer to **NullableFloat32** |  | [optional] 
**Bdid** | Pointer to **NullableFloat32** |  | [optional] 
**Peripherals** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewModel

`func NewModel(type_ string, name string, flavor string, model string, ) *Model`

NewModel instantiates a new Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelWithDefaults

`func NewModelWithDefaults() *Model`

NewModelWithDefaults instantiates a new Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Model) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Model) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Model) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Model) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model) SetName(v string)`

SetName sets Name field to given value.


### GetFlavor

`func (o *Model) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Model) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Model) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetDescription

`func (o *Model) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Model) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Model) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Model) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Model) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Model) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetModel

`func (o *Model) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Model) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Model) SetModel(v string)`

SetModel sets Model field to given value.


### GetBoardConfig

`func (o *Model) GetBoardConfig() string`

GetBoardConfig returns the BoardConfig field if non-nil, zero value otherwise.

### GetBoardConfigOk

`func (o *Model) GetBoardConfigOk() (*string, bool)`

GetBoardConfigOk returns a tuple with the BoardConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardConfig

`func (o *Model) SetBoardConfig(v string)`

SetBoardConfig sets BoardConfig field to given value.

### HasBoardConfig

`func (o *Model) HasBoardConfig() bool`

HasBoardConfig returns a boolean if a field has been set.

### SetBoardConfigNil

`func (o *Model) SetBoardConfigNil(b bool)`

 SetBoardConfigNil sets the value for BoardConfig to be an explicit nil

### UnsetBoardConfig
`func (o *Model) UnsetBoardConfig()`

UnsetBoardConfig ensures that no value is present for BoardConfig, not even an explicit nil
### GetPlatform

`func (o *Model) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Model) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Model) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Model) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Model) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Model) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetCpid

`func (o *Model) GetCpid() float32`

GetCpid returns the Cpid field if non-nil, zero value otherwise.

### GetCpidOk

`func (o *Model) GetCpidOk() (*float32, bool)`

GetCpidOk returns a tuple with the Cpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpid

`func (o *Model) SetCpid(v float32)`

SetCpid sets Cpid field to given value.

### HasCpid

`func (o *Model) HasCpid() bool`

HasCpid returns a boolean if a field has been set.

### SetCpidNil

`func (o *Model) SetCpidNil(b bool)`

 SetCpidNil sets the value for Cpid to be an explicit nil

### UnsetCpid
`func (o *Model) UnsetCpid()`

UnsetCpid ensures that no value is present for Cpid, not even an explicit nil
### GetBdid

`func (o *Model) GetBdid() float32`

GetBdid returns the Bdid field if non-nil, zero value otherwise.

### GetBdidOk

`func (o *Model) GetBdidOk() (*float32, bool)`

GetBdidOk returns a tuple with the Bdid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdid

`func (o *Model) SetBdid(v float32)`

SetBdid sets Bdid field to given value.

### HasBdid

`func (o *Model) HasBdid() bool`

HasBdid returns a boolean if a field has been set.

### SetBdidNil

`func (o *Model) SetBdidNil(b bool)`

 SetBdidNil sets the value for Bdid to be an explicit nil

### UnsetBdid
`func (o *Model) UnsetBdid()`

UnsetBdid ensures that no value is present for Bdid, not even an explicit nil
### GetPeripherals

`func (o *Model) GetPeripherals() bool`

GetPeripherals returns the Peripherals field if non-nil, zero value otherwise.

### GetPeripheralsOk

`func (o *Model) GetPeripheralsOk() (*bool, bool)`

GetPeripheralsOk returns a tuple with the Peripherals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeripherals

`func (o *Model) SetPeripherals(v bool)`

SetPeripherals sets Peripherals field to given value.

### HasPeripherals

`func (o *Model) HasPeripherals() bool`

HasPeripherals returns a boolean if a field has been set.

### SetPeripheralsNil

`func (o *Model) SetPeripheralsNil(b bool)`

 SetPeripheralsNil sets the value for Peripherals to be an explicit nil

### UnsetPeripherals
`func (o *Model) UnsetPeripherals()`

UnsetPeripherals ensures that no value is present for Peripherals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


