# VolumeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocate** | Pointer to **NullableFloat32** | GB | [optional] 
**Partitions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ComputeNode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVolumeOptions

`func NewVolumeOptions() *VolumeOptions`

NewVolumeOptions instantiates a new VolumeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeOptionsWithDefaults

`func NewVolumeOptionsWithDefaults() *VolumeOptions`

NewVolumeOptionsWithDefaults instantiates a new VolumeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocate

`func (o *VolumeOptions) GetAllocate() float32`

GetAllocate returns the Allocate field if non-nil, zero value otherwise.

### GetAllocateOk

`func (o *VolumeOptions) GetAllocateOk() (*float32, bool)`

GetAllocateOk returns a tuple with the Allocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocate

`func (o *VolumeOptions) SetAllocate(v float32)`

SetAllocate sets Allocate field to given value.

### HasAllocate

`func (o *VolumeOptions) HasAllocate() bool`

HasAllocate returns a boolean if a field has been set.

### SetAllocateNil

`func (o *VolumeOptions) SetAllocateNil(b bool)`

 SetAllocateNil sets the value for Allocate to be an explicit nil

### UnsetAllocate
`func (o *VolumeOptions) UnsetAllocate()`

UnsetAllocate ensures that no value is present for Allocate, not even an explicit nil
### GetPartitions

`func (o *VolumeOptions) GetPartitions() []map[string]interface{}`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *VolumeOptions) GetPartitionsOk() (*[]map[string]interface{}, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *VolumeOptions) SetPartitions(v []map[string]interface{})`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *VolumeOptions) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### SetPartitionsNil

`func (o *VolumeOptions) SetPartitionsNil(b bool)`

 SetPartitionsNil sets the value for Partitions to be an explicit nil

### UnsetPartitions
`func (o *VolumeOptions) UnsetPartitions()`

UnsetPartitions ensures that no value is present for Partitions, not even an explicit nil
### GetComputeNode

`func (o *VolumeOptions) GetComputeNode() string`

GetComputeNode returns the ComputeNode field if non-nil, zero value otherwise.

### GetComputeNodeOk

`func (o *VolumeOptions) GetComputeNodeOk() (*string, bool)`

GetComputeNodeOk returns a tuple with the ComputeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNode

`func (o *VolumeOptions) SetComputeNode(v string)`

SetComputeNode sets ComputeNode field to given value.

### HasComputeNode

`func (o *VolumeOptions) HasComputeNode() bool`

HasComputeNode returns a boolean if a field has been set.

### SetComputeNodeNil

`func (o *VolumeOptions) SetComputeNodeNil(b bool)`

 SetComputeNodeNil sets the value for ComputeNode to be an explicit nil

### UnsetComputeNode
`func (o *VolumeOptions) UnsetComputeNode()`

UnsetComputeNode ensures that no value is present for ComputeNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


