# CouponOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Amount** | **float32** | Amount | 
**Used** | **bool** | Is Used. If true, this coupon has been used and cannot be used again | 

## Methods

### NewCouponOptions

`func NewCouponOptions(type_ string, amount float32, used bool, ) *CouponOptions`

NewCouponOptions instantiates a new CouponOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponOptionsWithDefaults

`func NewCouponOptionsWithDefaults() *CouponOptions`

NewCouponOptionsWithDefaults instantiates a new CouponOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponOptions) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *CouponOptions) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CouponOptions) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CouponOptions) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetUsed

`func (o *CouponOptions) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *CouponOptions) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *CouponOptions) SetUsed(v bool)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


