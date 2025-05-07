# AgentApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Running** | Pointer to **NullableBool** |  | [optional] 
**DiskUsage** | Pointer to **NullableFloat32** |  | [optional] 
**Date** | Pointer to **NullableFloat32** |  | [optional] 
**ApplicationType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**BundleID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgentApp

`func NewAgentApp() *AgentApp`

NewAgentApp instantiates a new AgentApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAppWithDefaults

`func NewAgentAppWithDefaults() *AgentApp`

NewAgentAppWithDefaults instantiates a new AgentApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *AgentApp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentApp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentApp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AgentApp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AgentApp) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AgentApp) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetRunning

`func (o *AgentApp) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *AgentApp) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *AgentApp) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *AgentApp) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### SetRunningNil

`func (o *AgentApp) SetRunningNil(b bool)`

 SetRunningNil sets the value for Running to be an explicit nil

### UnsetRunning
`func (o *AgentApp) UnsetRunning()`

UnsetRunning ensures that no value is present for Running, not even an explicit nil
### GetDiskUsage

`func (o *AgentApp) GetDiskUsage() float32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *AgentApp) GetDiskUsageOk() (*float32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *AgentApp) SetDiskUsage(v float32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *AgentApp) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### SetDiskUsageNil

`func (o *AgentApp) SetDiskUsageNil(b bool)`

 SetDiskUsageNil sets the value for DiskUsage to be an explicit nil

### UnsetDiskUsage
`func (o *AgentApp) UnsetDiskUsage()`

UnsetDiskUsage ensures that no value is present for DiskUsage, not even an explicit nil
### GetDate

`func (o *AgentApp) GetDate() float32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AgentApp) GetDateOk() (*float32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AgentApp) SetDate(v float32)`

SetDate sets Date field to given value.

### HasDate

`func (o *AgentApp) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *AgentApp) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *AgentApp) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetApplicationType

`func (o *AgentApp) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *AgentApp) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *AgentApp) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *AgentApp) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *AgentApp) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *AgentApp) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetName

`func (o *AgentApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentApp) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgentApp) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgentApp) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBundleID

`func (o *AgentApp) GetBundleID() string`

GetBundleID returns the BundleID field if non-nil, zero value otherwise.

### GetBundleIDOk

`func (o *AgentApp) GetBundleIDOk() (*string, bool)`

GetBundleIDOk returns a tuple with the BundleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleID

`func (o *AgentApp) SetBundleID(v string)`

SetBundleID sets BundleID field to given value.

### HasBundleID

`func (o *AgentApp) HasBundleID() bool`

HasBundleID returns a boolean if a field has been set.

### SetBundleIDNil

`func (o *AgentApp) SetBundleIDNil(b bool)`

 SetBundleIDNil sets the value for BundleID to be an explicit nil

### UnsetBundleID
`func (o *AgentApp) UnsetBundleID()`

UnsetBundleID ensures that no value is present for BundleID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


