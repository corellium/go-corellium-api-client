# TrialRequestMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Company** | Pointer to **NullableString** | provided company name | [optional] 
**Malware** | Pointer to **NullableBool** | option from sign up form, interesting using for malware | [optional] 
**Research** | Pointer to **NullableBool** | option from sign up form, interesting using for research | [optional] 
**Pentest** | Pointer to **NullableBool** | option from sign up form, interesting using for pentesting | [optional] 
**Other** | Pointer to **NullableString** | option from sign up form, interesting using for other area added here | [optional] 

## Methods

### NewTrialRequestMetadata

`func NewTrialRequestMetadata() *TrialRequestMetadata`

NewTrialRequestMetadata instantiates a new TrialRequestMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialRequestMetadataWithDefaults

`func NewTrialRequestMetadataWithDefaults() *TrialRequestMetadata`

NewTrialRequestMetadataWithDefaults instantiates a new TrialRequestMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrialRequestMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrialRequestMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrialRequestMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrialRequestMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TrialRequestMetadata) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TrialRequestMetadata) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *TrialRequestMetadata) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TrialRequestMetadata) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TrialRequestMetadata) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TrialRequestMetadata) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *TrialRequestMetadata) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *TrialRequestMetadata) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetMalware

`func (o *TrialRequestMetadata) GetMalware() bool`

GetMalware returns the Malware field if non-nil, zero value otherwise.

### GetMalwareOk

`func (o *TrialRequestMetadata) GetMalwareOk() (*bool, bool)`

GetMalwareOk returns a tuple with the Malware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalware

`func (o *TrialRequestMetadata) SetMalware(v bool)`

SetMalware sets Malware field to given value.

### HasMalware

`func (o *TrialRequestMetadata) HasMalware() bool`

HasMalware returns a boolean if a field has been set.

### SetMalwareNil

`func (o *TrialRequestMetadata) SetMalwareNil(b bool)`

 SetMalwareNil sets the value for Malware to be an explicit nil

### UnsetMalware
`func (o *TrialRequestMetadata) UnsetMalware()`

UnsetMalware ensures that no value is present for Malware, not even an explicit nil
### GetResearch

`func (o *TrialRequestMetadata) GetResearch() bool`

GetResearch returns the Research field if non-nil, zero value otherwise.

### GetResearchOk

`func (o *TrialRequestMetadata) GetResearchOk() (*bool, bool)`

GetResearchOk returns a tuple with the Research field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearch

`func (o *TrialRequestMetadata) SetResearch(v bool)`

SetResearch sets Research field to given value.

### HasResearch

`func (o *TrialRequestMetadata) HasResearch() bool`

HasResearch returns a boolean if a field has been set.

### SetResearchNil

`func (o *TrialRequestMetadata) SetResearchNil(b bool)`

 SetResearchNil sets the value for Research to be an explicit nil

### UnsetResearch
`func (o *TrialRequestMetadata) UnsetResearch()`

UnsetResearch ensures that no value is present for Research, not even an explicit nil
### GetPentest

`func (o *TrialRequestMetadata) GetPentest() bool`

GetPentest returns the Pentest field if non-nil, zero value otherwise.

### GetPentestOk

`func (o *TrialRequestMetadata) GetPentestOk() (*bool, bool)`

GetPentestOk returns a tuple with the Pentest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPentest

`func (o *TrialRequestMetadata) SetPentest(v bool)`

SetPentest sets Pentest field to given value.

### HasPentest

`func (o *TrialRequestMetadata) HasPentest() bool`

HasPentest returns a boolean if a field has been set.

### SetPentestNil

`func (o *TrialRequestMetadata) SetPentestNil(b bool)`

 SetPentestNil sets the value for Pentest to be an explicit nil

### UnsetPentest
`func (o *TrialRequestMetadata) UnsetPentest()`

UnsetPentest ensures that no value is present for Pentest, not even an explicit nil
### GetOther

`func (o *TrialRequestMetadata) GetOther() string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *TrialRequestMetadata) GetOtherOk() (*string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *TrialRequestMetadata) SetOther(v string)`

SetOther sets Other field to given value.

### HasOther

`func (o *TrialRequestMetadata) HasOther() bool`

HasOther returns a boolean if a field has been set.

### SetOtherNil

`func (o *TrialRequestMetadata) SetOtherNil(b bool)`

 SetOtherNil sets the value for Other to be an explicit nil

### UnsetOther
`func (o *TrialRequestMetadata) UnsetOther()`

UnsetOther ensures that no value is present for Other, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


