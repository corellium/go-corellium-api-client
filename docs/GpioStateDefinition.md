# GpioStateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitCount** | **float32** | count of bits in members of this bank | 
**Banks** | [**[][]Bit**]([]Bit.md) | bits for each bank of this name as an array of arrays | 

## Methods

### NewGpioStateDefinition

`func NewGpioStateDefinition(bitCount float32, banks [][]Bit, ) *GpioStateDefinition`

NewGpioStateDefinition instantiates a new GpioStateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpioStateDefinitionWithDefaults

`func NewGpioStateDefinitionWithDefaults() *GpioStateDefinition`

NewGpioStateDefinitionWithDefaults instantiates a new GpioStateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitCount

`func (o *GpioStateDefinition) GetBitCount() float32`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *GpioStateDefinition) GetBitCountOk() (*float32, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *GpioStateDefinition) SetBitCount(v float32)`

SetBitCount sets BitCount field to given value.


### GetBanks

`func (o *GpioStateDefinition) GetBanks() [][]Bit`

GetBanks returns the Banks field if non-nil, zero value otherwise.

### GetBanksOk

`func (o *GpioStateDefinition) GetBanksOk() (*[][]Bit, bool)`

GetBanksOk returns a tuple with the Banks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanks

`func (o *GpioStateDefinition) SetBanks(v [][]Bit)`

SetBanks sets Banks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


