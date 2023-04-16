# InstanceCreateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | the name of the device | [optional] 
**Key** | Pointer to **NullableString** | Key used to encrypt the Instance | [optional] 
**Flavor** | **string** | the flavor id | 
**Project** | **string** | project UUID | 
**Os** | **string** | OS Version | 
**Osbuild** | Pointer to **NullableString** | OS Build | [optional] 
**Patches** | Pointer to **[]string** | list of patches to apply | [optional] 
**Fwpackage** | Pointer to **NullableString** | URL or image id | [optional] 
**OrigFwPackageUrl** | Pointer to **NullableString** | URL that firmware package used to create this instance is available at | [optional] 
**Encrypt** | Pointer to **NullableBool** |  | [optional] 
**OverrideWifiMAC** | Pointer to **NullableString** |  | [optional] 
**Volume** | Pointer to [**VolumeOptions**](VolumeOptions.md) |  | [optional] 
**Snapshot** | Pointer to **NullableString** | Snapshot ID for this instance to be cloned from if defined | [optional] 
**BootOptions** | Pointer to [**InstanceBootOptions**](InstanceBootOptions.md) |  | [optional] 
**Device** | Pointer to [**Model**](Model.md) |  | [optional] 

## Methods

### NewInstanceCreateOptions

`func NewInstanceCreateOptions(flavor string, project string, os string, ) *InstanceCreateOptions`

NewInstanceCreateOptions instantiates a new InstanceCreateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateOptionsWithDefaults

`func NewInstanceCreateOptionsWithDefaults() *InstanceCreateOptions`

NewInstanceCreateOptionsWithDefaults instantiates a new InstanceCreateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceCreateOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceCreateOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceCreateOptions) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceCreateOptions) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *InstanceCreateOptions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InstanceCreateOptions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InstanceCreateOptions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InstanceCreateOptions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *InstanceCreateOptions) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *InstanceCreateOptions) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetFlavor

`func (o *InstanceCreateOptions) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *InstanceCreateOptions) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *InstanceCreateOptions) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetProject

`func (o *InstanceCreateOptions) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InstanceCreateOptions) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InstanceCreateOptions) SetProject(v string)`

SetProject sets Project field to given value.


### GetOs

`func (o *InstanceCreateOptions) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstanceCreateOptions) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstanceCreateOptions) SetOs(v string)`

SetOs sets Os field to given value.


### GetOsbuild

`func (o *InstanceCreateOptions) GetOsbuild() string`

GetOsbuild returns the Osbuild field if non-nil, zero value otherwise.

### GetOsbuildOk

`func (o *InstanceCreateOptions) GetOsbuildOk() (*string, bool)`

GetOsbuildOk returns a tuple with the Osbuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsbuild

`func (o *InstanceCreateOptions) SetOsbuild(v string)`

SetOsbuild sets Osbuild field to given value.

### HasOsbuild

`func (o *InstanceCreateOptions) HasOsbuild() bool`

HasOsbuild returns a boolean if a field has been set.

### SetOsbuildNil

`func (o *InstanceCreateOptions) SetOsbuildNil(b bool)`

 SetOsbuildNil sets the value for Osbuild to be an explicit nil

### UnsetOsbuild
`func (o *InstanceCreateOptions) UnsetOsbuild()`

UnsetOsbuild ensures that no value is present for Osbuild, not even an explicit nil
### GetPatches

`func (o *InstanceCreateOptions) GetPatches() []string`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *InstanceCreateOptions) GetPatchesOk() (*[]string, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *InstanceCreateOptions) SetPatches(v []string)`

SetPatches sets Patches field to given value.

### HasPatches

`func (o *InstanceCreateOptions) HasPatches() bool`

HasPatches returns a boolean if a field has been set.

### SetPatchesNil

`func (o *InstanceCreateOptions) SetPatchesNil(b bool)`

 SetPatchesNil sets the value for Patches to be an explicit nil

### UnsetPatches
`func (o *InstanceCreateOptions) UnsetPatches()`

UnsetPatches ensures that no value is present for Patches, not even an explicit nil
### GetFwpackage

`func (o *InstanceCreateOptions) GetFwpackage() string`

GetFwpackage returns the Fwpackage field if non-nil, zero value otherwise.

### GetFwpackageOk

`func (o *InstanceCreateOptions) GetFwpackageOk() (*string, bool)`

GetFwpackageOk returns a tuple with the Fwpackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwpackage

`func (o *InstanceCreateOptions) SetFwpackage(v string)`

SetFwpackage sets Fwpackage field to given value.

### HasFwpackage

`func (o *InstanceCreateOptions) HasFwpackage() bool`

HasFwpackage returns a boolean if a field has been set.

### SetFwpackageNil

`func (o *InstanceCreateOptions) SetFwpackageNil(b bool)`

 SetFwpackageNil sets the value for Fwpackage to be an explicit nil

### UnsetFwpackage
`func (o *InstanceCreateOptions) UnsetFwpackage()`

UnsetFwpackage ensures that no value is present for Fwpackage, not even an explicit nil
### GetOrigFwPackageUrl

`func (o *InstanceCreateOptions) GetOrigFwPackageUrl() string`

GetOrigFwPackageUrl returns the OrigFwPackageUrl field if non-nil, zero value otherwise.

### GetOrigFwPackageUrlOk

`func (o *InstanceCreateOptions) GetOrigFwPackageUrlOk() (*string, bool)`

GetOrigFwPackageUrlOk returns a tuple with the OrigFwPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigFwPackageUrl

`func (o *InstanceCreateOptions) SetOrigFwPackageUrl(v string)`

SetOrigFwPackageUrl sets OrigFwPackageUrl field to given value.

### HasOrigFwPackageUrl

`func (o *InstanceCreateOptions) HasOrigFwPackageUrl() bool`

HasOrigFwPackageUrl returns a boolean if a field has been set.

### SetOrigFwPackageUrlNil

`func (o *InstanceCreateOptions) SetOrigFwPackageUrlNil(b bool)`

 SetOrigFwPackageUrlNil sets the value for OrigFwPackageUrl to be an explicit nil

### UnsetOrigFwPackageUrl
`func (o *InstanceCreateOptions) UnsetOrigFwPackageUrl()`

UnsetOrigFwPackageUrl ensures that no value is present for OrigFwPackageUrl, not even an explicit nil
### GetEncrypt

`func (o *InstanceCreateOptions) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *InstanceCreateOptions) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *InstanceCreateOptions) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *InstanceCreateOptions) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### SetEncryptNil

`func (o *InstanceCreateOptions) SetEncryptNil(b bool)`

 SetEncryptNil sets the value for Encrypt to be an explicit nil

### UnsetEncrypt
`func (o *InstanceCreateOptions) UnsetEncrypt()`

UnsetEncrypt ensures that no value is present for Encrypt, not even an explicit nil
### GetOverrideWifiMAC

`func (o *InstanceCreateOptions) GetOverrideWifiMAC() string`

GetOverrideWifiMAC returns the OverrideWifiMAC field if non-nil, zero value otherwise.

### GetOverrideWifiMACOk

`func (o *InstanceCreateOptions) GetOverrideWifiMACOk() (*string, bool)`

GetOverrideWifiMACOk returns a tuple with the OverrideWifiMAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideWifiMAC

`func (o *InstanceCreateOptions) SetOverrideWifiMAC(v string)`

SetOverrideWifiMAC sets OverrideWifiMAC field to given value.

### HasOverrideWifiMAC

`func (o *InstanceCreateOptions) HasOverrideWifiMAC() bool`

HasOverrideWifiMAC returns a boolean if a field has been set.

### SetOverrideWifiMACNil

`func (o *InstanceCreateOptions) SetOverrideWifiMACNil(b bool)`

 SetOverrideWifiMACNil sets the value for OverrideWifiMAC to be an explicit nil

### UnsetOverrideWifiMAC
`func (o *InstanceCreateOptions) UnsetOverrideWifiMAC()`

UnsetOverrideWifiMAC ensures that no value is present for OverrideWifiMAC, not even an explicit nil
### GetVolume

`func (o *InstanceCreateOptions) GetVolume() VolumeOptions`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *InstanceCreateOptions) GetVolumeOk() (*VolumeOptions, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *InstanceCreateOptions) SetVolume(v VolumeOptions)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *InstanceCreateOptions) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetSnapshot

`func (o *InstanceCreateOptions) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *InstanceCreateOptions) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *InstanceCreateOptions) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *InstanceCreateOptions) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### SetSnapshotNil

`func (o *InstanceCreateOptions) SetSnapshotNil(b bool)`

 SetSnapshotNil sets the value for Snapshot to be an explicit nil

### UnsetSnapshot
`func (o *InstanceCreateOptions) UnsetSnapshot()`

UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
### GetBootOptions

`func (o *InstanceCreateOptions) GetBootOptions() InstanceBootOptions`

GetBootOptions returns the BootOptions field if non-nil, zero value otherwise.

### GetBootOptionsOk

`func (o *InstanceCreateOptions) GetBootOptionsOk() (*InstanceBootOptions, bool)`

GetBootOptionsOk returns a tuple with the BootOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptions

`func (o *InstanceCreateOptions) SetBootOptions(v InstanceBootOptions)`

SetBootOptions sets BootOptions field to given value.

### HasBootOptions

`func (o *InstanceCreateOptions) HasBootOptions() bool`

HasBootOptions returns a boolean if a field has been set.

### GetDevice

`func (o *InstanceCreateOptions) GetDevice() Model`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InstanceCreateOptions) GetDeviceOk() (*Model, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InstanceCreateOptions) SetDevice(v Model)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InstanceCreateOptions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


