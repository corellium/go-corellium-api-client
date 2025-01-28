# Assessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ReportId** | Pointer to **string** |  | [optional] 
**Bundle** | Pointer to **map[string]interface{}** |  | [optional] 
**Results** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**WordlistId** | Pointer to **string** |  | [optional] 
**Wordlist** | Pointer to **map[string]interface{}** |  | [optional] 
**InstanceId** | **string** |  | 
**BundleId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewAssessment

`func NewAssessment(instanceId string, bundleId string, status string, ) *Assessment`

NewAssessment instantiates a new Assessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentWithDefaults

`func NewAssessmentWithDefaults() *Assessment`

NewAssessmentWithDefaults instantiates a new Assessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Assessment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Assessment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Assessment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Assessment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Assessment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Assessment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Assessment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Assessment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Assessment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Assessment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Assessment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Assessment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetReportId

`func (o *Assessment) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *Assessment) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *Assessment) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *Assessment) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetBundle

`func (o *Assessment) GetBundle() map[string]interface{}`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *Assessment) GetBundleOk() (*map[string]interface{}, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *Assessment) SetBundle(v map[string]interface{})`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *Assessment) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetResults

`func (o *Assessment) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Assessment) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Assessment) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *Assessment) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetMetadata

`func (o *Assessment) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Assessment) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Assessment) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Assessment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWordlistId

`func (o *Assessment) GetWordlistId() string`

GetWordlistId returns the WordlistId field if non-nil, zero value otherwise.

### GetWordlistIdOk

`func (o *Assessment) GetWordlistIdOk() (*string, bool)`

GetWordlistIdOk returns a tuple with the WordlistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordlistId

`func (o *Assessment) SetWordlistId(v string)`

SetWordlistId sets WordlistId field to given value.

### HasWordlistId

`func (o *Assessment) HasWordlistId() bool`

HasWordlistId returns a boolean if a field has been set.

### GetWordlist

`func (o *Assessment) GetWordlist() map[string]interface{}`

GetWordlist returns the Wordlist field if non-nil, zero value otherwise.

### GetWordlistOk

`func (o *Assessment) GetWordlistOk() (*map[string]interface{}, bool)`

GetWordlistOk returns a tuple with the Wordlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordlist

`func (o *Assessment) SetWordlist(v map[string]interface{})`

SetWordlist sets Wordlist field to given value.

### HasWordlist

`func (o *Assessment) HasWordlist() bool`

HasWordlist returns a boolean if a field has been set.

### GetInstanceId

`func (o *Assessment) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Assessment) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Assessment) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetBundleId

`func (o *Assessment) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *Assessment) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *Assessment) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetStatus

`func (o *Assessment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Assessment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Assessment) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


