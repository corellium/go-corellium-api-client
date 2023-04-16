# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Project Identifier | 
**Name** | Pointer to **NullableString** | Project Name | [optional] 
**Settings** | Pointer to [**ProjectSettings**](ProjectSettings.md) |  | [optional] 
**Quotas** | Pointer to [**ProjectQuota**](ProjectQuota.md) |  | [optional] 
**QuotasUsed** | Pointer to [**ProjectUsage**](ProjectUsage.md) |  | [optional] 

## Methods

### NewProject

`func NewProject(id string, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Project) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Project) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSettings

`func (o *Project) GetSettings() ProjectSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Project) GetSettingsOk() (*ProjectSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Project) SetSettings(v ProjectSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Project) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetQuotas

`func (o *Project) GetQuotas() ProjectQuota`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *Project) GetQuotasOk() (*ProjectQuota, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *Project) SetQuotas(v ProjectQuota)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *Project) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetQuotasUsed

`func (o *Project) GetQuotasUsed() ProjectUsage`

GetQuotasUsed returns the QuotasUsed field if non-nil, zero value otherwise.

### GetQuotasUsedOk

`func (o *Project) GetQuotasUsedOk() (*ProjectUsage, bool)`

GetQuotasUsedOk returns a tuple with the QuotasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotasUsed

`func (o *Project) SetQuotasUsed(v ProjectUsage)`

SetQuotasUsed sets QuotasUsed field to given value.

### HasQuotasUsed

`func (o *Project) HasQuotasUsed() bool`

HasQuotasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


