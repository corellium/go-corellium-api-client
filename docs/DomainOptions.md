# DomainOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpRequired** | Pointer to **NullableBool** | if true, totp is required | [optional] 
**TrialExtension** | Pointer to [**TrialExtension**](TrialExtension.md) |  | [optional] 

## Methods

### NewDomainOptions

`func NewDomainOptions() *DomainOptions`

NewDomainOptions instantiates a new DomainOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainOptionsWithDefaults

`func NewDomainOptionsWithDefaults() *DomainOptions`

NewDomainOptionsWithDefaults instantiates a new DomainOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpRequired

`func (o *DomainOptions) GetTotpRequired() bool`

GetTotpRequired returns the TotpRequired field if non-nil, zero value otherwise.

### GetTotpRequiredOk

`func (o *DomainOptions) GetTotpRequiredOk() (*bool, bool)`

GetTotpRequiredOk returns a tuple with the TotpRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpRequired

`func (o *DomainOptions) SetTotpRequired(v bool)`

SetTotpRequired sets TotpRequired field to given value.

### HasTotpRequired

`func (o *DomainOptions) HasTotpRequired() bool`

HasTotpRequired returns a boolean if a field has been set.

### SetTotpRequiredNil

`func (o *DomainOptions) SetTotpRequiredNil(b bool)`

 SetTotpRequiredNil sets the value for TotpRequired to be an explicit nil

### UnsetTotpRequired
`func (o *DomainOptions) UnsetTotpRequired()`

UnsetTotpRequired ensures that no value is present for TotpRequired, not even an explicit nil
### GetTrialExtension

`func (o *DomainOptions) GetTrialExtension() TrialExtension`

GetTrialExtension returns the TrialExtension field if non-nil, zero value otherwise.

### GetTrialExtensionOk

`func (o *DomainOptions) GetTrialExtensionOk() (*TrialExtension, bool)`

GetTrialExtensionOk returns a tuple with the TrialExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialExtension

`func (o *DomainOptions) SetTrialExtension(v TrialExtension)`

SetTrialExtension sets TrialExtension field to given value.

### HasTrialExtension

`func (o *DomainOptions) HasTrialExtension() bool`

HasTrialExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


