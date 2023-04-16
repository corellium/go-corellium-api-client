# WebPlayerCreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Project Identifier | 
**InstanceId** | **string** | VM Instance Identifier | 
**ExpiresIn** | **float32** | Number of seconds token remains valid | 
**Features** | [**Features**](Features.md) |  | 
**ClientId** | Pointer to **NullableString** | Optional string value supplied by client for tracking purposes | [optional] 

## Methods

### NewWebPlayerCreateSessionRequest

`func NewWebPlayerCreateSessionRequest(projectId string, instanceId string, expiresIn float32, features Features, ) *WebPlayerCreateSessionRequest`

NewWebPlayerCreateSessionRequest instantiates a new WebPlayerCreateSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPlayerCreateSessionRequestWithDefaults

`func NewWebPlayerCreateSessionRequestWithDefaults() *WebPlayerCreateSessionRequest`

NewWebPlayerCreateSessionRequestWithDefaults instantiates a new WebPlayerCreateSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *WebPlayerCreateSessionRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebPlayerCreateSessionRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebPlayerCreateSessionRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetInstanceId

`func (o *WebPlayerCreateSessionRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *WebPlayerCreateSessionRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *WebPlayerCreateSessionRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetExpiresIn

`func (o *WebPlayerCreateSessionRequest) GetExpiresIn() float32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *WebPlayerCreateSessionRequest) GetExpiresInOk() (*float32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *WebPlayerCreateSessionRequest) SetExpiresIn(v float32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetFeatures

`func (o *WebPlayerCreateSessionRequest) GetFeatures() Features`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *WebPlayerCreateSessionRequest) GetFeaturesOk() (*Features, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *WebPlayerCreateSessionRequest) SetFeatures(v Features)`

SetFeatures sets Features field to given value.


### GetClientId

`func (o *WebPlayerCreateSessionRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WebPlayerCreateSessionRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WebPlayerCreateSessionRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *WebPlayerCreateSessionRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *WebPlayerCreateSessionRequest) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *WebPlayerCreateSessionRequest) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


