# PeripheralsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acceleration** | Pointer to **[]float32** |  | [optional] 
**Gyroscope** | Pointer to **[]float32** |  | [optional] 
**Magnetic** | Pointer to **[]float32** |  | [optional] 
**Orientation** | Pointer to **[]float32** |  | [optional] 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**Proximity** | Pointer to **NullableFloat32** |  | [optional] 
**Light** | Pointer to **NullableFloat32** |  | [optional] 
**Pressure** | Pointer to **NullableFloat32** |  | [optional] 
**Humidity** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewPeripheralsData

`func NewPeripheralsData() *PeripheralsData`

NewPeripheralsData instantiates a new PeripheralsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeripheralsDataWithDefaults

`func NewPeripheralsDataWithDefaults() *PeripheralsData`

NewPeripheralsDataWithDefaults instantiates a new PeripheralsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceleration

`func (o *PeripheralsData) GetAcceleration() []float32`

GetAcceleration returns the Acceleration field if non-nil, zero value otherwise.

### GetAccelerationOk

`func (o *PeripheralsData) GetAccelerationOk() (*[]float32, bool)`

GetAccelerationOk returns a tuple with the Acceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleration

`func (o *PeripheralsData) SetAcceleration(v []float32)`

SetAcceleration sets Acceleration field to given value.

### HasAcceleration

`func (o *PeripheralsData) HasAcceleration() bool`

HasAcceleration returns a boolean if a field has been set.

### SetAccelerationNil

`func (o *PeripheralsData) SetAccelerationNil(b bool)`

 SetAccelerationNil sets the value for Acceleration to be an explicit nil

### UnsetAcceleration
`func (o *PeripheralsData) UnsetAcceleration()`

UnsetAcceleration ensures that no value is present for Acceleration, not even an explicit nil
### GetGyroscope

`func (o *PeripheralsData) GetGyroscope() []float32`

GetGyroscope returns the Gyroscope field if non-nil, zero value otherwise.

### GetGyroscopeOk

`func (o *PeripheralsData) GetGyroscopeOk() (*[]float32, bool)`

GetGyroscopeOk returns a tuple with the Gyroscope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGyroscope

`func (o *PeripheralsData) SetGyroscope(v []float32)`

SetGyroscope sets Gyroscope field to given value.

### HasGyroscope

`func (o *PeripheralsData) HasGyroscope() bool`

HasGyroscope returns a boolean if a field has been set.

### SetGyroscopeNil

`func (o *PeripheralsData) SetGyroscopeNil(b bool)`

 SetGyroscopeNil sets the value for Gyroscope to be an explicit nil

### UnsetGyroscope
`func (o *PeripheralsData) UnsetGyroscope()`

UnsetGyroscope ensures that no value is present for Gyroscope, not even an explicit nil
### GetMagnetic

`func (o *PeripheralsData) GetMagnetic() []float32`

GetMagnetic returns the Magnetic field if non-nil, zero value otherwise.

### GetMagneticOk

`func (o *PeripheralsData) GetMagneticOk() (*[]float32, bool)`

GetMagneticOk returns a tuple with the Magnetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetic

`func (o *PeripheralsData) SetMagnetic(v []float32)`

SetMagnetic sets Magnetic field to given value.

### HasMagnetic

`func (o *PeripheralsData) HasMagnetic() bool`

HasMagnetic returns a boolean if a field has been set.

### SetMagneticNil

`func (o *PeripheralsData) SetMagneticNil(b bool)`

 SetMagneticNil sets the value for Magnetic to be an explicit nil

### UnsetMagnetic
`func (o *PeripheralsData) UnsetMagnetic()`

UnsetMagnetic ensures that no value is present for Magnetic, not even an explicit nil
### GetOrientation

`func (o *PeripheralsData) GetOrientation() []float32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *PeripheralsData) GetOrientationOk() (*[]float32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *PeripheralsData) SetOrientation(v []float32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *PeripheralsData) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### SetOrientationNil

`func (o *PeripheralsData) SetOrientationNil(b bool)`

 SetOrientationNil sets the value for Orientation to be an explicit nil

### UnsetOrientation
`func (o *PeripheralsData) UnsetOrientation()`

UnsetOrientation ensures that no value is present for Orientation, not even an explicit nil
### GetTemperature

`func (o *PeripheralsData) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *PeripheralsData) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *PeripheralsData) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *PeripheralsData) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *PeripheralsData) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *PeripheralsData) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetProximity

`func (o *PeripheralsData) GetProximity() float32`

GetProximity returns the Proximity field if non-nil, zero value otherwise.

### GetProximityOk

`func (o *PeripheralsData) GetProximityOk() (*float32, bool)`

GetProximityOk returns a tuple with the Proximity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximity

`func (o *PeripheralsData) SetProximity(v float32)`

SetProximity sets Proximity field to given value.

### HasProximity

`func (o *PeripheralsData) HasProximity() bool`

HasProximity returns a boolean if a field has been set.

### SetProximityNil

`func (o *PeripheralsData) SetProximityNil(b bool)`

 SetProximityNil sets the value for Proximity to be an explicit nil

### UnsetProximity
`func (o *PeripheralsData) UnsetProximity()`

UnsetProximity ensures that no value is present for Proximity, not even an explicit nil
### GetLight

`func (o *PeripheralsData) GetLight() float32`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *PeripheralsData) GetLightOk() (*float32, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *PeripheralsData) SetLight(v float32)`

SetLight sets Light field to given value.

### HasLight

`func (o *PeripheralsData) HasLight() bool`

HasLight returns a boolean if a field has been set.

### SetLightNil

`func (o *PeripheralsData) SetLightNil(b bool)`

 SetLightNil sets the value for Light to be an explicit nil

### UnsetLight
`func (o *PeripheralsData) UnsetLight()`

UnsetLight ensures that no value is present for Light, not even an explicit nil
### GetPressure

`func (o *PeripheralsData) GetPressure() float32`

GetPressure returns the Pressure field if non-nil, zero value otherwise.

### GetPressureOk

`func (o *PeripheralsData) GetPressureOk() (*float32, bool)`

GetPressureOk returns a tuple with the Pressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressure

`func (o *PeripheralsData) SetPressure(v float32)`

SetPressure sets Pressure field to given value.

### HasPressure

`func (o *PeripheralsData) HasPressure() bool`

HasPressure returns a boolean if a field has been set.

### SetPressureNil

`func (o *PeripheralsData) SetPressureNil(b bool)`

 SetPressureNil sets the value for Pressure to be an explicit nil

### UnsetPressure
`func (o *PeripheralsData) UnsetPressure()`

UnsetPressure ensures that no value is present for Pressure, not even an explicit nil
### GetHumidity

`func (o *PeripheralsData) GetHumidity() float32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *PeripheralsData) GetHumidityOk() (*float32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *PeripheralsData) SetHumidity(v float32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *PeripheralsData) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### SetHumidityNil

`func (o *PeripheralsData) SetHumidityNil(b bool)`

 SetHumidityNil sets the value for Humidity to be an explicit nil

### UnsetHumidity
`func (o *PeripheralsData) UnsetHumidity()`

UnsetHumidity ensures that no value is present for Humidity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


