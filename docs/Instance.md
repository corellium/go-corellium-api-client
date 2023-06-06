# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Instance Identifier | [optional] 
**Name** | Pointer to **NullableString** | Instance Name | [optional] 
**Key** | Pointer to **NullableString** | Key used to encrypt the Instance | [optional] 
**Flavor** | Pointer to **NullableString** | The type of virtual machine this is | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** | The projectId of the project this instance belongs to | [optional] 
**State** | Pointer to [**InstanceState**](InstanceState.md) |  | [optional] 
**StateChanged** | Pointer to **NullableTime** | Time the state of the instance last changed | [optional] 
**StartedAt** | Pointer to **NullableString** | Time the instance was started | [optional] 
**UserTask** | Pointer to **NullableString** | Currently executing User Task | [optional] 
**TaskState** | Pointer to **NullableString** | Current task state if any | [optional] 
**Error** | Pointer to **NullableString** | Current error state if any | [optional] 
**BootOptions** | Pointer to [**InstanceBootOptions**](InstanceBootOptions.md) |  | [optional] 
**ServiceIp** | Pointer to **NullableString** | Services IP Address | [optional] 
**WifiIp** | Pointer to **NullableString** | LAN IP Address | [optional] 
**SecondaryIp** | Pointer to **NullableString** | Secondary Inteface LAN IP Address (if supported) | [optional] 
**Services** | Pointer to [**InstanceServices**](InstanceServices.md) |  | [optional] 
**Panicked** | Pointer to **NullableBool** |  | [optional] 
**Created** | Pointer to **NullableTime** | Time instance was created | [optional] 
**Model** | Pointer to **NullableString** | Model of virtual machine device | [optional] 
**Fwpackage** | Pointer to **NullableString** | URL that package used to create this instance is available at | [optional] 
**Os** | Pointer to **NullableString** |  | [optional] 
**Agent** | Pointer to [**NullableInstanceAgentState**](InstanceAgentState.md) |  | [optional] 
**Netmon** | Pointer to [**InstanceNetmonState**](InstanceNetmonState.md) |  | [optional] 
**Netdump** | Pointer to [**InstanceNetdumpState**](InstanceNetdumpState.md) |  | [optional] 
**ExposePort** | Pointer to **NullableString** |  | [optional] 
**Fault** | Pointer to **NullableBool** |  | [optional] 
**Patches** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to [**CreatedBy**](CreatedBy.md) |  | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Instance) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Instance) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Instance) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Instance) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Instance) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *Instance) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Instance) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Instance) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Instance) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *Instance) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *Instance) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetFlavor

`func (o *Instance) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Instance) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Instance) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *Instance) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *Instance) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *Instance) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetType

`func (o *Instance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Instance) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Instance) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Instance) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProject

`func (o *Instance) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Instance) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Instance) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Instance) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *Instance) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *Instance) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetState

`func (o *Instance) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *Instance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChanged

`func (o *Instance) GetStateChanged() time.Time`

GetStateChanged returns the StateChanged field if non-nil, zero value otherwise.

### GetStateChangedOk

`func (o *Instance) GetStateChangedOk() (*time.Time, bool)`

GetStateChangedOk returns a tuple with the StateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChanged

`func (o *Instance) SetStateChanged(v time.Time)`

SetStateChanged sets StateChanged field to given value.

### HasStateChanged

`func (o *Instance) HasStateChanged() bool`

HasStateChanged returns a boolean if a field has been set.

### SetStateChangedNil

`func (o *Instance) SetStateChangedNil(b bool)`

 SetStateChangedNil sets the value for StateChanged to be an explicit nil

### UnsetStateChanged
`func (o *Instance) UnsetStateChanged()`

UnsetStateChanged ensures that no value is present for StateChanged, not even an explicit nil
### GetStartedAt

`func (o *Instance) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Instance) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Instance) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Instance) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Instance) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Instance) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetUserTask

`func (o *Instance) GetUserTask() string`

GetUserTask returns the UserTask field if non-nil, zero value otherwise.

### GetUserTaskOk

`func (o *Instance) GetUserTaskOk() (*string, bool)`

GetUserTaskOk returns a tuple with the UserTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTask

`func (o *Instance) SetUserTask(v string)`

SetUserTask sets UserTask field to given value.

### HasUserTask

`func (o *Instance) HasUserTask() bool`

HasUserTask returns a boolean if a field has been set.

### SetUserTaskNil

`func (o *Instance) SetUserTaskNil(b bool)`

 SetUserTaskNil sets the value for UserTask to be an explicit nil

### UnsetUserTask
`func (o *Instance) UnsetUserTask()`

UnsetUserTask ensures that no value is present for UserTask, not even an explicit nil
### GetTaskState

`func (o *Instance) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *Instance) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *Instance) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *Instance) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *Instance) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *Instance) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetError

`func (o *Instance) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Instance) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Instance) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Instance) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Instance) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Instance) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetBootOptions

`func (o *Instance) GetBootOptions() InstanceBootOptions`

GetBootOptions returns the BootOptions field if non-nil, zero value otherwise.

### GetBootOptionsOk

`func (o *Instance) GetBootOptionsOk() (*InstanceBootOptions, bool)`

GetBootOptionsOk returns a tuple with the BootOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptions

`func (o *Instance) SetBootOptions(v InstanceBootOptions)`

SetBootOptions sets BootOptions field to given value.

### HasBootOptions

`func (o *Instance) HasBootOptions() bool`

HasBootOptions returns a boolean if a field has been set.

### GetServiceIp

`func (o *Instance) GetServiceIp() string`

GetServiceIp returns the ServiceIp field if non-nil, zero value otherwise.

### GetServiceIpOk

`func (o *Instance) GetServiceIpOk() (*string, bool)`

GetServiceIpOk returns a tuple with the ServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIp

`func (o *Instance) SetServiceIp(v string)`

SetServiceIp sets ServiceIp field to given value.

### HasServiceIp

`func (o *Instance) HasServiceIp() bool`

HasServiceIp returns a boolean if a field has been set.

### SetServiceIpNil

`func (o *Instance) SetServiceIpNil(b bool)`

 SetServiceIpNil sets the value for ServiceIp to be an explicit nil

### UnsetServiceIp
`func (o *Instance) UnsetServiceIp()`

UnsetServiceIp ensures that no value is present for ServiceIp, not even an explicit nil
### GetWifiIp

`func (o *Instance) GetWifiIp() string`

GetWifiIp returns the WifiIp field if non-nil, zero value otherwise.

### GetWifiIpOk

`func (o *Instance) GetWifiIpOk() (*string, bool)`

GetWifiIpOk returns a tuple with the WifiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiIp

`func (o *Instance) SetWifiIp(v string)`

SetWifiIp sets WifiIp field to given value.

### HasWifiIp

`func (o *Instance) HasWifiIp() bool`

HasWifiIp returns a boolean if a field has been set.

### SetWifiIpNil

`func (o *Instance) SetWifiIpNil(b bool)`

 SetWifiIpNil sets the value for WifiIp to be an explicit nil

### UnsetWifiIp
`func (o *Instance) UnsetWifiIp()`

UnsetWifiIp ensures that no value is present for WifiIp, not even an explicit nil
### GetSecondaryIp

`func (o *Instance) GetSecondaryIp() string`

GetSecondaryIp returns the SecondaryIp field if non-nil, zero value otherwise.

### GetSecondaryIpOk

`func (o *Instance) GetSecondaryIpOk() (*string, bool)`

GetSecondaryIpOk returns a tuple with the SecondaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIp

`func (o *Instance) SetSecondaryIp(v string)`

SetSecondaryIp sets SecondaryIp field to given value.

### HasSecondaryIp

`func (o *Instance) HasSecondaryIp() bool`

HasSecondaryIp returns a boolean if a field has been set.

### SetSecondaryIpNil

`func (o *Instance) SetSecondaryIpNil(b bool)`

 SetSecondaryIpNil sets the value for SecondaryIp to be an explicit nil

### UnsetSecondaryIp
`func (o *Instance) UnsetSecondaryIp()`

UnsetSecondaryIp ensures that no value is present for SecondaryIp, not even an explicit nil
### GetServices

`func (o *Instance) GetServices() InstanceServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Instance) GetServicesOk() (*InstanceServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Instance) SetServices(v InstanceServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *Instance) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetPanicked

`func (o *Instance) GetPanicked() bool`

GetPanicked returns the Panicked field if non-nil, zero value otherwise.

### GetPanickedOk

`func (o *Instance) GetPanickedOk() (*bool, bool)`

GetPanickedOk returns a tuple with the Panicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanicked

`func (o *Instance) SetPanicked(v bool)`

SetPanicked sets Panicked field to given value.

### HasPanicked

`func (o *Instance) HasPanicked() bool`

HasPanicked returns a boolean if a field has been set.

### SetPanickedNil

`func (o *Instance) SetPanickedNil(b bool)`

 SetPanickedNil sets the value for Panicked to be an explicit nil

### UnsetPanicked
`func (o *Instance) UnsetPanicked()`

UnsetPanicked ensures that no value is present for Panicked, not even an explicit nil
### GetCreated

`func (o *Instance) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Instance) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Instance) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Instance) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Instance) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Instance) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModel

`func (o *Instance) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Instance) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Instance) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Instance) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Instance) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Instance) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetFwpackage

`func (o *Instance) GetFwpackage() string`

GetFwpackage returns the Fwpackage field if non-nil, zero value otherwise.

### GetFwpackageOk

`func (o *Instance) GetFwpackageOk() (*string, bool)`

GetFwpackageOk returns a tuple with the Fwpackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwpackage

`func (o *Instance) SetFwpackage(v string)`

SetFwpackage sets Fwpackage field to given value.

### HasFwpackage

`func (o *Instance) HasFwpackage() bool`

HasFwpackage returns a boolean if a field has been set.

### SetFwpackageNil

`func (o *Instance) SetFwpackageNil(b bool)`

 SetFwpackageNil sets the value for Fwpackage to be an explicit nil

### UnsetFwpackage
`func (o *Instance) UnsetFwpackage()`

UnsetFwpackage ensures that no value is present for Fwpackage, not even an explicit nil
### GetOs

`func (o *Instance) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Instance) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Instance) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Instance) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *Instance) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *Instance) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetAgent

`func (o *Instance) GetAgent() InstanceAgentState`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *Instance) GetAgentOk() (*InstanceAgentState, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *Instance) SetAgent(v InstanceAgentState)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *Instance) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### SetAgentNil

`func (o *Instance) SetAgentNil(b bool)`

 SetAgentNil sets the value for Agent to be an explicit nil

### UnsetAgent
`func (o *Instance) UnsetAgent()`

UnsetAgent ensures that no value is present for Agent, not even an explicit nil
### GetNetmon

`func (o *Instance) GetNetmon() InstanceNetmonState`

GetNetmon returns the Netmon field if non-nil, zero value otherwise.

### GetNetmonOk

`func (o *Instance) GetNetmonOk() (*InstanceNetmonState, bool)`

GetNetmonOk returns a tuple with the Netmon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmon

`func (o *Instance) SetNetmon(v InstanceNetmonState)`

SetNetmon sets Netmon field to given value.

### HasNetmon

`func (o *Instance) HasNetmon() bool`

HasNetmon returns a boolean if a field has been set.

### GetNetdump

`func (o *Instance) GetNetdump() InstanceNetdumpState`

GetNetdump returns the Netdump field if non-nil, zero value otherwise.

### GetNetdumpOk

`func (o *Instance) GetNetdumpOk() (*InstanceNetdumpState, bool)`

GetNetdumpOk returns a tuple with the Netdump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetdump

`func (o *Instance) SetNetdump(v InstanceNetdumpState)`

SetNetdump sets Netdump field to given value.

### HasNetdump

`func (o *Instance) HasNetdump() bool`

HasNetdump returns a boolean if a field has been set.

### GetExposePort

`func (o *Instance) GetExposePort() string`

GetExposePort returns the ExposePort field if non-nil, zero value otherwise.

### GetExposePortOk

`func (o *Instance) GetExposePortOk() (*string, bool)`

GetExposePortOk returns a tuple with the ExposePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposePort

`func (o *Instance) SetExposePort(v string)`

SetExposePort sets ExposePort field to given value.

### HasExposePort

`func (o *Instance) HasExposePort() bool`

HasExposePort returns a boolean if a field has been set.

### SetExposePortNil

`func (o *Instance) SetExposePortNil(b bool)`

 SetExposePortNil sets the value for ExposePort to be an explicit nil

### UnsetExposePort
`func (o *Instance) UnsetExposePort()`

UnsetExposePort ensures that no value is present for ExposePort, not even an explicit nil
### GetFault

`func (o *Instance) GetFault() bool`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *Instance) GetFaultOk() (*bool, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *Instance) SetFault(v bool)`

SetFault sets Fault field to given value.

### HasFault

`func (o *Instance) HasFault() bool`

HasFault returns a boolean if a field has been set.

### SetFaultNil

`func (o *Instance) SetFaultNil(b bool)`

 SetFaultNil sets the value for Fault to be an explicit nil

### UnsetFault
`func (o *Instance) UnsetFault()`

UnsetFault ensures that no value is present for Fault, not even an explicit nil
### GetPatches

`func (o *Instance) GetPatches() []string`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *Instance) GetPatchesOk() (*[]string, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *Instance) SetPatches(v []string)`

SetPatches sets Patches field to given value.

### HasPatches

`func (o *Instance) HasPatches() bool`

HasPatches returns a boolean if a field has been set.

### SetPatchesNil

`func (o *Instance) SetPatchesNil(b bool)`

 SetPatchesNil sets the value for Patches to be an explicit nil

### UnsetPatches
`func (o *Instance) UnsetPatches()`

UnsetPatches ensures that no value is present for Patches, not even an explicit nil
### GetCreatedBy

`func (o *Instance) GetCreatedBy() CreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Instance) GetCreatedByOk() (*CreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Instance) SetCreatedBy(v CreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Instance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


