# MeteredSubscriptionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurstRate** | Pointer to **NullableFloat32** | Unit price | [optional] 
**AllocatedHours** | Pointer to **NullableFloat32** | Allocated hours | [optional] 
**PlanBilledAmount** | Pointer to **NullableFloat32** | Plan cost | [optional] 
**PlanBilledUnits** | Pointer to **NullableFloat32** | Units included in plan | [optional] 
**BurstBilledAmount** | Pointer to **NullableFloat32** | Amount billed | [optional] 
**BurstBilledUnits** | Pointer to **NullableFloat32** | Units billed | [optional] 
**BurstOutstandingAmount** | Pointer to **NullableFloat32** | Outstanding amount | [optional] 
**BurstOutstandingUnits** | Pointer to **NullableFloat32** | Outstanding units | [optional] 
**TotalUsageAmount** | Pointer to **NullableFloat32** | Total cost in cents for current period | [optional] 
**TotalUsedUnits** | Pointer to **NullableFloat32** | Total used units for current period | [optional] 

## Methods

### NewMeteredSubscriptionUsage

`func NewMeteredSubscriptionUsage() *MeteredSubscriptionUsage`

NewMeteredSubscriptionUsage instantiates a new MeteredSubscriptionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteredSubscriptionUsageWithDefaults

`func NewMeteredSubscriptionUsageWithDefaults() *MeteredSubscriptionUsage`

NewMeteredSubscriptionUsageWithDefaults instantiates a new MeteredSubscriptionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurstRate

`func (o *MeteredSubscriptionUsage) GetBurstRate() float32`

GetBurstRate returns the BurstRate field if non-nil, zero value otherwise.

### GetBurstRateOk

`func (o *MeteredSubscriptionUsage) GetBurstRateOk() (*float32, bool)`

GetBurstRateOk returns a tuple with the BurstRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstRate

`func (o *MeteredSubscriptionUsage) SetBurstRate(v float32)`

SetBurstRate sets BurstRate field to given value.

### HasBurstRate

`func (o *MeteredSubscriptionUsage) HasBurstRate() bool`

HasBurstRate returns a boolean if a field has been set.

### SetBurstRateNil

`func (o *MeteredSubscriptionUsage) SetBurstRateNil(b bool)`

 SetBurstRateNil sets the value for BurstRate to be an explicit nil

### UnsetBurstRate
`func (o *MeteredSubscriptionUsage) UnsetBurstRate()`

UnsetBurstRate ensures that no value is present for BurstRate, not even an explicit nil
### GetAllocatedHours

`func (o *MeteredSubscriptionUsage) GetAllocatedHours() float32`

GetAllocatedHours returns the AllocatedHours field if non-nil, zero value otherwise.

### GetAllocatedHoursOk

`func (o *MeteredSubscriptionUsage) GetAllocatedHoursOk() (*float32, bool)`

GetAllocatedHoursOk returns a tuple with the AllocatedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedHours

`func (o *MeteredSubscriptionUsage) SetAllocatedHours(v float32)`

SetAllocatedHours sets AllocatedHours field to given value.

### HasAllocatedHours

`func (o *MeteredSubscriptionUsage) HasAllocatedHours() bool`

HasAllocatedHours returns a boolean if a field has been set.

### SetAllocatedHoursNil

`func (o *MeteredSubscriptionUsage) SetAllocatedHoursNil(b bool)`

 SetAllocatedHoursNil sets the value for AllocatedHours to be an explicit nil

### UnsetAllocatedHours
`func (o *MeteredSubscriptionUsage) UnsetAllocatedHours()`

UnsetAllocatedHours ensures that no value is present for AllocatedHours, not even an explicit nil
### GetPlanBilledAmount

`func (o *MeteredSubscriptionUsage) GetPlanBilledAmount() float32`

GetPlanBilledAmount returns the PlanBilledAmount field if non-nil, zero value otherwise.

### GetPlanBilledAmountOk

`func (o *MeteredSubscriptionUsage) GetPlanBilledAmountOk() (*float32, bool)`

GetPlanBilledAmountOk returns a tuple with the PlanBilledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBilledAmount

`func (o *MeteredSubscriptionUsage) SetPlanBilledAmount(v float32)`

SetPlanBilledAmount sets PlanBilledAmount field to given value.

### HasPlanBilledAmount

`func (o *MeteredSubscriptionUsage) HasPlanBilledAmount() bool`

HasPlanBilledAmount returns a boolean if a field has been set.

### SetPlanBilledAmountNil

`func (o *MeteredSubscriptionUsage) SetPlanBilledAmountNil(b bool)`

 SetPlanBilledAmountNil sets the value for PlanBilledAmount to be an explicit nil

### UnsetPlanBilledAmount
`func (o *MeteredSubscriptionUsage) UnsetPlanBilledAmount()`

UnsetPlanBilledAmount ensures that no value is present for PlanBilledAmount, not even an explicit nil
### GetPlanBilledUnits

`func (o *MeteredSubscriptionUsage) GetPlanBilledUnits() float32`

GetPlanBilledUnits returns the PlanBilledUnits field if non-nil, zero value otherwise.

### GetPlanBilledUnitsOk

`func (o *MeteredSubscriptionUsage) GetPlanBilledUnitsOk() (*float32, bool)`

GetPlanBilledUnitsOk returns a tuple with the PlanBilledUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBilledUnits

`func (o *MeteredSubscriptionUsage) SetPlanBilledUnits(v float32)`

SetPlanBilledUnits sets PlanBilledUnits field to given value.

### HasPlanBilledUnits

`func (o *MeteredSubscriptionUsage) HasPlanBilledUnits() bool`

HasPlanBilledUnits returns a boolean if a field has been set.

### SetPlanBilledUnitsNil

`func (o *MeteredSubscriptionUsage) SetPlanBilledUnitsNil(b bool)`

 SetPlanBilledUnitsNil sets the value for PlanBilledUnits to be an explicit nil

### UnsetPlanBilledUnits
`func (o *MeteredSubscriptionUsage) UnsetPlanBilledUnits()`

UnsetPlanBilledUnits ensures that no value is present for PlanBilledUnits, not even an explicit nil
### GetBurstBilledAmount

`func (o *MeteredSubscriptionUsage) GetBurstBilledAmount() float32`

GetBurstBilledAmount returns the BurstBilledAmount field if non-nil, zero value otherwise.

### GetBurstBilledAmountOk

`func (o *MeteredSubscriptionUsage) GetBurstBilledAmountOk() (*float32, bool)`

GetBurstBilledAmountOk returns a tuple with the BurstBilledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBilledAmount

`func (o *MeteredSubscriptionUsage) SetBurstBilledAmount(v float32)`

SetBurstBilledAmount sets BurstBilledAmount field to given value.

### HasBurstBilledAmount

`func (o *MeteredSubscriptionUsage) HasBurstBilledAmount() bool`

HasBurstBilledAmount returns a boolean if a field has been set.

### SetBurstBilledAmountNil

`func (o *MeteredSubscriptionUsage) SetBurstBilledAmountNil(b bool)`

 SetBurstBilledAmountNil sets the value for BurstBilledAmount to be an explicit nil

### UnsetBurstBilledAmount
`func (o *MeteredSubscriptionUsage) UnsetBurstBilledAmount()`

UnsetBurstBilledAmount ensures that no value is present for BurstBilledAmount, not even an explicit nil
### GetBurstBilledUnits

`func (o *MeteredSubscriptionUsage) GetBurstBilledUnits() float32`

GetBurstBilledUnits returns the BurstBilledUnits field if non-nil, zero value otherwise.

### GetBurstBilledUnitsOk

`func (o *MeteredSubscriptionUsage) GetBurstBilledUnitsOk() (*float32, bool)`

GetBurstBilledUnitsOk returns a tuple with the BurstBilledUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBilledUnits

`func (o *MeteredSubscriptionUsage) SetBurstBilledUnits(v float32)`

SetBurstBilledUnits sets BurstBilledUnits field to given value.

### HasBurstBilledUnits

`func (o *MeteredSubscriptionUsage) HasBurstBilledUnits() bool`

HasBurstBilledUnits returns a boolean if a field has been set.

### SetBurstBilledUnitsNil

`func (o *MeteredSubscriptionUsage) SetBurstBilledUnitsNil(b bool)`

 SetBurstBilledUnitsNil sets the value for BurstBilledUnits to be an explicit nil

### UnsetBurstBilledUnits
`func (o *MeteredSubscriptionUsage) UnsetBurstBilledUnits()`

UnsetBurstBilledUnits ensures that no value is present for BurstBilledUnits, not even an explicit nil
### GetBurstOutstandingAmount

`func (o *MeteredSubscriptionUsage) GetBurstOutstandingAmount() float32`

GetBurstOutstandingAmount returns the BurstOutstandingAmount field if non-nil, zero value otherwise.

### GetBurstOutstandingAmountOk

`func (o *MeteredSubscriptionUsage) GetBurstOutstandingAmountOk() (*float32, bool)`

GetBurstOutstandingAmountOk returns a tuple with the BurstOutstandingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstOutstandingAmount

`func (o *MeteredSubscriptionUsage) SetBurstOutstandingAmount(v float32)`

SetBurstOutstandingAmount sets BurstOutstandingAmount field to given value.

### HasBurstOutstandingAmount

`func (o *MeteredSubscriptionUsage) HasBurstOutstandingAmount() bool`

HasBurstOutstandingAmount returns a boolean if a field has been set.

### SetBurstOutstandingAmountNil

`func (o *MeteredSubscriptionUsage) SetBurstOutstandingAmountNil(b bool)`

 SetBurstOutstandingAmountNil sets the value for BurstOutstandingAmount to be an explicit nil

### UnsetBurstOutstandingAmount
`func (o *MeteredSubscriptionUsage) UnsetBurstOutstandingAmount()`

UnsetBurstOutstandingAmount ensures that no value is present for BurstOutstandingAmount, not even an explicit nil
### GetBurstOutstandingUnits

`func (o *MeteredSubscriptionUsage) GetBurstOutstandingUnits() float32`

GetBurstOutstandingUnits returns the BurstOutstandingUnits field if non-nil, zero value otherwise.

### GetBurstOutstandingUnitsOk

`func (o *MeteredSubscriptionUsage) GetBurstOutstandingUnitsOk() (*float32, bool)`

GetBurstOutstandingUnitsOk returns a tuple with the BurstOutstandingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstOutstandingUnits

`func (o *MeteredSubscriptionUsage) SetBurstOutstandingUnits(v float32)`

SetBurstOutstandingUnits sets BurstOutstandingUnits field to given value.

### HasBurstOutstandingUnits

`func (o *MeteredSubscriptionUsage) HasBurstOutstandingUnits() bool`

HasBurstOutstandingUnits returns a boolean if a field has been set.

### SetBurstOutstandingUnitsNil

`func (o *MeteredSubscriptionUsage) SetBurstOutstandingUnitsNil(b bool)`

 SetBurstOutstandingUnitsNil sets the value for BurstOutstandingUnits to be an explicit nil

### UnsetBurstOutstandingUnits
`func (o *MeteredSubscriptionUsage) UnsetBurstOutstandingUnits()`

UnsetBurstOutstandingUnits ensures that no value is present for BurstOutstandingUnits, not even an explicit nil
### GetTotalUsageAmount

`func (o *MeteredSubscriptionUsage) GetTotalUsageAmount() float32`

GetTotalUsageAmount returns the TotalUsageAmount field if non-nil, zero value otherwise.

### GetTotalUsageAmountOk

`func (o *MeteredSubscriptionUsage) GetTotalUsageAmountOk() (*float32, bool)`

GetTotalUsageAmountOk returns a tuple with the TotalUsageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageAmount

`func (o *MeteredSubscriptionUsage) SetTotalUsageAmount(v float32)`

SetTotalUsageAmount sets TotalUsageAmount field to given value.

### HasTotalUsageAmount

`func (o *MeteredSubscriptionUsage) HasTotalUsageAmount() bool`

HasTotalUsageAmount returns a boolean if a field has been set.

### SetTotalUsageAmountNil

`func (o *MeteredSubscriptionUsage) SetTotalUsageAmountNil(b bool)`

 SetTotalUsageAmountNil sets the value for TotalUsageAmount to be an explicit nil

### UnsetTotalUsageAmount
`func (o *MeteredSubscriptionUsage) UnsetTotalUsageAmount()`

UnsetTotalUsageAmount ensures that no value is present for TotalUsageAmount, not even an explicit nil
### GetTotalUsedUnits

`func (o *MeteredSubscriptionUsage) GetTotalUsedUnits() float32`

GetTotalUsedUnits returns the TotalUsedUnits field if non-nil, zero value otherwise.

### GetTotalUsedUnitsOk

`func (o *MeteredSubscriptionUsage) GetTotalUsedUnitsOk() (*float32, bool)`

GetTotalUsedUnitsOk returns a tuple with the TotalUsedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsedUnits

`func (o *MeteredSubscriptionUsage) SetTotalUsedUnits(v float32)`

SetTotalUsedUnits sets TotalUsedUnits field to given value.

### HasTotalUsedUnits

`func (o *MeteredSubscriptionUsage) HasTotalUsedUnits() bool`

HasTotalUsedUnits returns a boolean if a field has been set.

### SetTotalUsedUnitsNil

`func (o *MeteredSubscriptionUsage) SetTotalUsedUnitsNil(b bool)`

 SetTotalUsedUnitsNil sets the value for TotalUsedUnits to be an explicit nil

### UnsetTotalUsedUnits
`func (o *MeteredSubscriptionUsage) UnsetTotalUsedUnits()`

UnsetTotalUsedUnits ensures that no value is present for TotalUsedUnits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


