# ConfigResponseMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Maintenance message | [optional] 
**Header** | Pointer to **NullableString** | Maintenance header | [optional] 

## Methods

### NewConfigResponseMaintenance

`func NewConfigResponseMaintenance() *ConfigResponseMaintenance`

NewConfigResponseMaintenance instantiates a new ConfigResponseMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigResponseMaintenanceWithDefaults

`func NewConfigResponseMaintenanceWithDefaults() *ConfigResponseMaintenance`

NewConfigResponseMaintenanceWithDefaults instantiates a new ConfigResponseMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConfigResponseMaintenance) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigResponseMaintenance) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigResponseMaintenance) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigResponseMaintenance) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ConfigResponseMaintenance) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ConfigResponseMaintenance) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetHeader

`func (o *ConfigResponseMaintenance) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ConfigResponseMaintenance) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ConfigResponseMaintenance) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ConfigResponseMaintenance) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *ConfigResponseMaintenance) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *ConfigResponseMaintenance) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


