# DomainAuthProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Provider ID | 
**Id** | Pointer to **NullableString** | Provider ID for backward compatibility with frontend | [optional] 
**ProviderType** | **string** | Provider Type | 
**Provider** | Pointer to **NullableString** | Provider Type for backward compatibility with frontend | [optional] 
**Label** | **string** | Login Button Text | 
**Name** | Pointer to **NullableString** | Title Text for the Frontend&#39;s Provider Configuration Form | [optional] 
**AuthorizationUrl** | Pointer to **NullableString** | Coordinator endpoint for Authorizing with this provider | [optional] 
**Default** | **bool** | True: Configured in coordinator.json, False: Custom Domain Provider | 
**Enabled** | **bool** | Enabled/Disabled | 
**Config** | Pointer to [**OpenIDConfig**](OpenIDConfig.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** | Optional User ID of creator | [optional] 
**CreatedAt** | **string** | Created Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | 
**UpdatedAt** | **string** | Updated Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | 

## Methods

### NewDomainAuthProviderResponse

`func NewDomainAuthProviderResponse(identifier string, providerType string, label string, default_ bool, enabled bool, createdAt string, updatedAt string, ) *DomainAuthProviderResponse`

NewDomainAuthProviderResponse instantiates a new DomainAuthProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAuthProviderResponseWithDefaults

`func NewDomainAuthProviderResponseWithDefaults() *DomainAuthProviderResponse`

NewDomainAuthProviderResponseWithDefaults instantiates a new DomainAuthProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *DomainAuthProviderResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DomainAuthProviderResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DomainAuthProviderResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetId

`func (o *DomainAuthProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainAuthProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainAuthProviderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainAuthProviderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DomainAuthProviderResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DomainAuthProviderResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProviderType

`func (o *DomainAuthProviderResponse) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DomainAuthProviderResponse) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DomainAuthProviderResponse) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProvider

`func (o *DomainAuthProviderResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DomainAuthProviderResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DomainAuthProviderResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DomainAuthProviderResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *DomainAuthProviderResponse) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *DomainAuthProviderResponse) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetLabel

`func (o *DomainAuthProviderResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DomainAuthProviderResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DomainAuthProviderResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *DomainAuthProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainAuthProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainAuthProviderResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainAuthProviderResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DomainAuthProviderResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DomainAuthProviderResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAuthorizationUrl

`func (o *DomainAuthProviderResponse) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *DomainAuthProviderResponse) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *DomainAuthProviderResponse) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *DomainAuthProviderResponse) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.

### SetAuthorizationUrlNil

`func (o *DomainAuthProviderResponse) SetAuthorizationUrlNil(b bool)`

 SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil

### UnsetAuthorizationUrl
`func (o *DomainAuthProviderResponse) UnsetAuthorizationUrl()`

UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
### GetDefault

`func (o *DomainAuthProviderResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DomainAuthProviderResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DomainAuthProviderResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetEnabled

`func (o *DomainAuthProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DomainAuthProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DomainAuthProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetConfig

`func (o *DomainAuthProviderResponse) GetConfig() OpenIDConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DomainAuthProviderResponse) GetConfigOk() (*OpenIDConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DomainAuthProviderResponse) SetConfig(v OpenIDConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DomainAuthProviderResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DomainAuthProviderResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DomainAuthProviderResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DomainAuthProviderResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DomainAuthProviderResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *DomainAuthProviderResponse) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *DomainAuthProviderResponse) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *DomainAuthProviderResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainAuthProviderResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainAuthProviderResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DomainAuthProviderResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DomainAuthProviderResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DomainAuthProviderResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


