# V0040ExtSensorsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumedEnergy** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Temperature** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**EnergyUpdateTime** | Pointer to **int64** |  | [optional] 
**CurrentWatts** | Pointer to **int32** |  | [optional] 

## Methods

### NewV0040ExtSensorsData

`func NewV0040ExtSensorsData() *V0040ExtSensorsData`

NewV0040ExtSensorsData instantiates a new V0040ExtSensorsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040ExtSensorsDataWithDefaults

`func NewV0040ExtSensorsDataWithDefaults() *V0040ExtSensorsData`

NewV0040ExtSensorsDataWithDefaults instantiates a new V0040ExtSensorsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumedEnergy

`func (o *V0040ExtSensorsData) GetConsumedEnergy() V0040Uint64NoVal`

GetConsumedEnergy returns the ConsumedEnergy field if non-nil, zero value otherwise.

### GetConsumedEnergyOk

`func (o *V0040ExtSensorsData) GetConsumedEnergyOk() (*V0040Uint64NoVal, bool)`

GetConsumedEnergyOk returns a tuple with the ConsumedEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedEnergy

`func (o *V0040ExtSensorsData) SetConsumedEnergy(v V0040Uint64NoVal)`

SetConsumedEnergy sets ConsumedEnergy field to given value.

### HasConsumedEnergy

`func (o *V0040ExtSensorsData) HasConsumedEnergy() bool`

HasConsumedEnergy returns a boolean if a field has been set.

### GetTemperature

`func (o *V0040ExtSensorsData) GetTemperature() V0040Uint32NoVal`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *V0040ExtSensorsData) GetTemperatureOk() (*V0040Uint32NoVal, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *V0040ExtSensorsData) SetTemperature(v V0040Uint32NoVal)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *V0040ExtSensorsData) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetEnergyUpdateTime

`func (o *V0040ExtSensorsData) GetEnergyUpdateTime() int64`

GetEnergyUpdateTime returns the EnergyUpdateTime field if non-nil, zero value otherwise.

### GetEnergyUpdateTimeOk

`func (o *V0040ExtSensorsData) GetEnergyUpdateTimeOk() (*int64, bool)`

GetEnergyUpdateTimeOk returns a tuple with the EnergyUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyUpdateTime

`func (o *V0040ExtSensorsData) SetEnergyUpdateTime(v int64)`

SetEnergyUpdateTime sets EnergyUpdateTime field to given value.

### HasEnergyUpdateTime

`func (o *V0040ExtSensorsData) HasEnergyUpdateTime() bool`

HasEnergyUpdateTime returns a boolean if a field has been set.

### GetCurrentWatts

`func (o *V0040ExtSensorsData) GetCurrentWatts() int32`

GetCurrentWatts returns the CurrentWatts field if non-nil, zero value otherwise.

### GetCurrentWattsOk

`func (o *V0040ExtSensorsData) GetCurrentWattsOk() (*int32, bool)`

GetCurrentWattsOk returns a tuple with the CurrentWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWatts

`func (o *V0040ExtSensorsData) SetCurrentWatts(v int32)`

SetCurrentWatts sets CurrentWatts field to given value.

### HasCurrentWatts

`func (o *V0040ExtSensorsData) HasCurrentWatts() bool`

HasCurrentWatts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


