# SubscriberInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponOptions** | Pointer to [**CouponOptions**](CouponOptions.md) |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**Trial** | Pointer to [**Trial**](Trial.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Name | [optional] 
**Email** | Pointer to **NullableString** | Email | [optional] 
**CouponCode** | **string** | Coupon code | 
**Domain** | Pointer to **NullableString** | Domain | [optional] 
**Aggregate** | **string** | Aggregate | 
**UseBy** | Pointer to **NullableString** | Used By | [optional] 
**Active** | **bool** | Is Active flag | 
**Reusable** | **bool** | Is reusable flag | 
**CreatedAt** | **string** | Created Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | 
**UpdatedAt** | **string** | Updated Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z | 

## Methods

### NewSubscriberInvite

`func NewSubscriberInvite(couponCode string, aggregate string, active bool, reusable bool, createdAt string, updatedAt string, ) *SubscriberInvite`

NewSubscriberInvite instantiates a new SubscriberInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberInviteWithDefaults

`func NewSubscriberInviteWithDefaults() *SubscriberInvite`

NewSubscriberInviteWithDefaults instantiates a new SubscriberInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponOptions

`func (o *SubscriberInvite) GetCouponOptions() CouponOptions`

GetCouponOptions returns the CouponOptions field if non-nil, zero value otherwise.

### GetCouponOptionsOk

`func (o *SubscriberInvite) GetCouponOptionsOk() (*CouponOptions, bool)`

GetCouponOptionsOk returns a tuple with the CouponOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponOptions

`func (o *SubscriberInvite) SetCouponOptions(v CouponOptions)`

SetCouponOptions sets CouponOptions field to given value.

### HasCouponOptions

`func (o *SubscriberInvite) HasCouponOptions() bool`

HasCouponOptions returns a boolean if a field has been set.

### GetPlan

`func (o *SubscriberInvite) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriberInvite) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriberInvite) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriberInvite) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetTrial

`func (o *SubscriberInvite) GetTrial() Trial`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *SubscriberInvite) GetTrialOk() (*Trial, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *SubscriberInvite) SetTrial(v Trial)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *SubscriberInvite) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetName

`func (o *SubscriberInvite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriberInvite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriberInvite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriberInvite) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SubscriberInvite) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SubscriberInvite) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *SubscriberInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubscriberInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubscriberInvite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubscriberInvite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SubscriberInvite) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SubscriberInvite) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCouponCode

`func (o *SubscriberInvite) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *SubscriberInvite) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *SubscriberInvite) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.


### GetDomain

`func (o *SubscriberInvite) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SubscriberInvite) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SubscriberInvite) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SubscriberInvite) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SubscriberInvite) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SubscriberInvite) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetAggregate

`func (o *SubscriberInvite) GetAggregate() string`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *SubscriberInvite) GetAggregateOk() (*string, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *SubscriberInvite) SetAggregate(v string)`

SetAggregate sets Aggregate field to given value.


### GetUseBy

`func (o *SubscriberInvite) GetUseBy() string`

GetUseBy returns the UseBy field if non-nil, zero value otherwise.

### GetUseByOk

`func (o *SubscriberInvite) GetUseByOk() (*string, bool)`

GetUseByOk returns a tuple with the UseBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBy

`func (o *SubscriberInvite) SetUseBy(v string)`

SetUseBy sets UseBy field to given value.

### HasUseBy

`func (o *SubscriberInvite) HasUseBy() bool`

HasUseBy returns a boolean if a field has been set.

### SetUseByNil

`func (o *SubscriberInvite) SetUseByNil(b bool)`

 SetUseByNil sets the value for UseBy to be an explicit nil

### UnsetUseBy
`func (o *SubscriberInvite) UnsetUseBy()`

UnsetUseBy ensures that no value is present for UseBy, not even an explicit nil
### GetActive

`func (o *SubscriberInvite) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubscriberInvite) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubscriberInvite) SetActive(v bool)`

SetActive sets Active field to given value.


### GetReusable

`func (o *SubscriberInvite) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *SubscriberInvite) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *SubscriberInvite) SetReusable(v bool)`

SetReusable sets Reusable field to given value.


### GetCreatedAt

`func (o *SubscriberInvite) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriberInvite) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriberInvite) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriberInvite) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriberInvite) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriberInvite) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


