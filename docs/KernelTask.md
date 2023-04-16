# KernelTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KernelId** | Pointer to **NullableString** | Kernel Task ID | [optional] 
**Name** | Pointer to **NullableString** | Thread name | [optional] 
**Pid** | Pointer to **NullableInt32** | Process ID of task | [optional] 
**Threads** | Pointer to [**[]KernelThread**](KernelThread.md) |  | [optional] 

## Methods

### NewKernelTask

`func NewKernelTask() *KernelTask`

NewKernelTask instantiates a new KernelTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKernelTaskWithDefaults

`func NewKernelTaskWithDefaults() *KernelTask`

NewKernelTaskWithDefaults instantiates a new KernelTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKernelId

`func (o *KernelTask) GetKernelId() string`

GetKernelId returns the KernelId field if non-nil, zero value otherwise.

### GetKernelIdOk

`func (o *KernelTask) GetKernelIdOk() (*string, bool)`

GetKernelIdOk returns a tuple with the KernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelId

`func (o *KernelTask) SetKernelId(v string)`

SetKernelId sets KernelId field to given value.

### HasKernelId

`func (o *KernelTask) HasKernelId() bool`

HasKernelId returns a boolean if a field has been set.

### SetKernelIdNil

`func (o *KernelTask) SetKernelIdNil(b bool)`

 SetKernelIdNil sets the value for KernelId to be an explicit nil

### UnsetKernelId
`func (o *KernelTask) UnsetKernelId()`

UnsetKernelId ensures that no value is present for KernelId, not even an explicit nil
### GetName

`func (o *KernelTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KernelTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KernelTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KernelTask) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KernelTask) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KernelTask) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPid

`func (o *KernelTask) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *KernelTask) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *KernelTask) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *KernelTask) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *KernelTask) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *KernelTask) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetThreads

`func (o *KernelTask) GetThreads() []KernelThread`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *KernelTask) GetThreadsOk() (*[]KernelThread, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *KernelTask) SetThreads(v []KernelThread)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *KernelTask) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### SetThreadsNil

`func (o *KernelTask) SetThreadsNil(b bool)`

 SetThreadsNil sets the value for Threads to be an explicit nil

### UnsetThreads
`func (o *KernelTask) UnsetThreads()`

UnsetThreads ensures that no value is present for Threads, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


