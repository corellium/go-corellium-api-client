# TrialRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Email** | Pointer to **NullableString** | The user&#39;s email address. | [optional] 
**Name** | Pointer to **NullableString** | The user&#39;s full name. | [optional] 
**Metadata** | Pointer to [**TrialRequestMetadata**](TrialRequestMetadata.md) |  | [optional] 
**Enterprise** | Pointer to **NullableBool** | If true, create an enterprise domain. | [optional] 
**Token** | Pointer to **NullableString** | Stripe payment token. | [optional] 

## Methods

### NewTrialRequestOptions

`func NewTrialRequestOptions() *TrialRequestOptions`

NewTrialRequestOptions instantiates a new TrialRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialRequestOptionsWithDefaults

`func NewTrialRequestOptionsWithDefaults() *TrialRequestOptions`

NewTrialRequestOptionsWithDefaults instantiates a new TrialRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TrialRequestOptions) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TrialRequestOptions) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TrialRequestOptions) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TrialRequestOptions) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *TrialRequestOptions) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrialRequestOptions) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrialRequestOptions) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TrialRequestOptions) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *TrialRequestOptions) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *TrialRequestOptions) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *TrialRequestOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrialRequestOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrialRequestOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrialRequestOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TrialRequestOptions) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TrialRequestOptions) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMetadata

`func (o *TrialRequestOptions) GetMetadata() TrialRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TrialRequestOptions) GetMetadataOk() (*TrialRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TrialRequestOptions) SetMetadata(v TrialRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TrialRequestOptions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEnterprise

`func (o *TrialRequestOptions) GetEnterprise() bool`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *TrialRequestOptions) GetEnterpriseOk() (*bool, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *TrialRequestOptions) SetEnterprise(v bool)`

SetEnterprise sets Enterprise field to given value.

### HasEnterprise

`func (o *TrialRequestOptions) HasEnterprise() bool`

HasEnterprise returns a boolean if a field has been set.

### SetEnterpriseNil

`func (o *TrialRequestOptions) SetEnterpriseNil(b bool)`

 SetEnterpriseNil sets the value for Enterprise to be an explicit nil

### UnsetEnterprise
`func (o *TrialRequestOptions) UnsetEnterprise()`

UnsetEnterprise ensures that no value is present for Enterprise, not even an explicit nil
### GetToken

`func (o *TrialRequestOptions) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TrialRequestOptions) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TrialRequestOptions) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TrialRequestOptions) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *TrialRequestOptions) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *TrialRequestOptions) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


