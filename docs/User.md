# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID | 
**Label** | **string** | User Label | 
**Name** | **string** | User Name | 
**Email** | **string** | User Email | 
**Administrator** | Pointer to **NullableBool** | the flag that specifies whether user is Administrator or not | [optional] 
**CanEditUserAttributes** | Pointer to **NullableBool** | Flag to determine if user attributes are editable. | [optional] 

## Methods

### NewUser

`func NewUser(id string, label string, name string, email string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *User) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *User) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *User) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAdministrator

`func (o *User) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *User) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *User) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *User) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### SetAdministratorNil

`func (o *User) SetAdministratorNil(b bool)`

 SetAdministratorNil sets the value for Administrator to be an explicit nil

### UnsetAdministrator
`func (o *User) UnsetAdministrator()`

UnsetAdministrator ensures that no value is present for Administrator, not even an explicit nil
### GetCanEditUserAttributes

`func (o *User) GetCanEditUserAttributes() bool`

GetCanEditUserAttributes returns the CanEditUserAttributes field if non-nil, zero value otherwise.

### GetCanEditUserAttributesOk

`func (o *User) GetCanEditUserAttributesOk() (*bool, bool)`

GetCanEditUserAttributesOk returns a tuple with the CanEditUserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditUserAttributes

`func (o *User) SetCanEditUserAttributes(v bool)`

SetCanEditUserAttributes sets CanEditUserAttributes field to given value.

### HasCanEditUserAttributes

`func (o *User) HasCanEditUserAttributes() bool`

HasCanEditUserAttributes returns a boolean if a field has been set.

### SetCanEditUserAttributesNil

`func (o *User) SetCanEditUserAttributesNil(b bool)`

 SetCanEditUserAttributesNil sets the value for CanEditUserAttributes to be an explicit nil

### UnsetCanEditUserAttributes
`func (o *User) UnsetCanEditUserAttributes()`

UnsetCanEditUserAttributes ensures that no value is present for CanEditUserAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


