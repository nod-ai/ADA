# V0040PowerMgmtData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumWatts** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**CurrentWatts** | Pointer to **int32** |  | [optional] 
**TotalEnergy** | Pointer to **int64** |  | [optional] 
**NewMaximumWatts** | Pointer to **int32** |  | [optional] 
**PeakWatts** | Pointer to **int32** |  | [optional] 
**LowestWatts** | Pointer to **int32** |  | [optional] 
**NewJobTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**TimeStartDay** | Pointer to **int64** |  | [optional] 

## Methods

### NewV0040PowerMgmtData

`func NewV0040PowerMgmtData() *V0040PowerMgmtData`

NewV0040PowerMgmtData instantiates a new V0040PowerMgmtData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040PowerMgmtDataWithDefaults

`func NewV0040PowerMgmtDataWithDefaults() *V0040PowerMgmtData`

NewV0040PowerMgmtDataWithDefaults instantiates a new V0040PowerMgmtData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumWatts

`func (o *V0040PowerMgmtData) GetMaximumWatts() V0040Uint32NoVal`

GetMaximumWatts returns the MaximumWatts field if non-nil, zero value otherwise.

### GetMaximumWattsOk

`func (o *V0040PowerMgmtData) GetMaximumWattsOk() (*V0040Uint32NoVal, bool)`

GetMaximumWattsOk returns a tuple with the MaximumWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumWatts

`func (o *V0040PowerMgmtData) SetMaximumWatts(v V0040Uint32NoVal)`

SetMaximumWatts sets MaximumWatts field to given value.

### HasMaximumWatts

`func (o *V0040PowerMgmtData) HasMaximumWatts() bool`

HasMaximumWatts returns a boolean if a field has been set.

### GetCurrentWatts

`func (o *V0040PowerMgmtData) GetCurrentWatts() int32`

GetCurrentWatts returns the CurrentWatts field if non-nil, zero value otherwise.

### GetCurrentWattsOk

`func (o *V0040PowerMgmtData) GetCurrentWattsOk() (*int32, bool)`

GetCurrentWattsOk returns a tuple with the CurrentWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWatts

`func (o *V0040PowerMgmtData) SetCurrentWatts(v int32)`

SetCurrentWatts sets CurrentWatts field to given value.

### HasCurrentWatts

`func (o *V0040PowerMgmtData) HasCurrentWatts() bool`

HasCurrentWatts returns a boolean if a field has been set.

### GetTotalEnergy

`func (o *V0040PowerMgmtData) GetTotalEnergy() int64`

GetTotalEnergy returns the TotalEnergy field if non-nil, zero value otherwise.

### GetTotalEnergyOk

`func (o *V0040PowerMgmtData) GetTotalEnergyOk() (*int64, bool)`

GetTotalEnergyOk returns a tuple with the TotalEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnergy

`func (o *V0040PowerMgmtData) SetTotalEnergy(v int64)`

SetTotalEnergy sets TotalEnergy field to given value.

### HasTotalEnergy

`func (o *V0040PowerMgmtData) HasTotalEnergy() bool`

HasTotalEnergy returns a boolean if a field has been set.

### GetNewMaximumWatts

`func (o *V0040PowerMgmtData) GetNewMaximumWatts() int32`

GetNewMaximumWatts returns the NewMaximumWatts field if non-nil, zero value otherwise.

### GetNewMaximumWattsOk

`func (o *V0040PowerMgmtData) GetNewMaximumWattsOk() (*int32, bool)`

GetNewMaximumWattsOk returns a tuple with the NewMaximumWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMaximumWatts

`func (o *V0040PowerMgmtData) SetNewMaximumWatts(v int32)`

SetNewMaximumWatts sets NewMaximumWatts field to given value.

### HasNewMaximumWatts

`func (o *V0040PowerMgmtData) HasNewMaximumWatts() bool`

HasNewMaximumWatts returns a boolean if a field has been set.

### GetPeakWatts

`func (o *V0040PowerMgmtData) GetPeakWatts() int32`

GetPeakWatts returns the PeakWatts field if non-nil, zero value otherwise.

### GetPeakWattsOk

`func (o *V0040PowerMgmtData) GetPeakWattsOk() (*int32, bool)`

GetPeakWattsOk returns a tuple with the PeakWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakWatts

`func (o *V0040PowerMgmtData) SetPeakWatts(v int32)`

SetPeakWatts sets PeakWatts field to given value.

### HasPeakWatts

`func (o *V0040PowerMgmtData) HasPeakWatts() bool`

HasPeakWatts returns a boolean if a field has been set.

### GetLowestWatts

`func (o *V0040PowerMgmtData) GetLowestWatts() int32`

GetLowestWatts returns the LowestWatts field if non-nil, zero value otherwise.

### GetLowestWattsOk

`func (o *V0040PowerMgmtData) GetLowestWattsOk() (*int32, bool)`

GetLowestWattsOk returns a tuple with the LowestWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestWatts

`func (o *V0040PowerMgmtData) SetLowestWatts(v int32)`

SetLowestWatts sets LowestWatts field to given value.

### HasLowestWatts

`func (o *V0040PowerMgmtData) HasLowestWatts() bool`

HasLowestWatts returns a boolean if a field has been set.

### GetNewJobTime

`func (o *V0040PowerMgmtData) GetNewJobTime() V0040Uint64NoVal`

GetNewJobTime returns the NewJobTime field if non-nil, zero value otherwise.

### GetNewJobTimeOk

`func (o *V0040PowerMgmtData) GetNewJobTimeOk() (*V0040Uint64NoVal, bool)`

GetNewJobTimeOk returns a tuple with the NewJobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewJobTime

`func (o *V0040PowerMgmtData) SetNewJobTime(v V0040Uint64NoVal)`

SetNewJobTime sets NewJobTime field to given value.

### HasNewJobTime

`func (o *V0040PowerMgmtData) HasNewJobTime() bool`

HasNewJobTime returns a boolean if a field has been set.

### GetState

`func (o *V0040PowerMgmtData) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0040PowerMgmtData) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0040PowerMgmtData) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *V0040PowerMgmtData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeStartDay

`func (o *V0040PowerMgmtData) GetTimeStartDay() int64`

GetTimeStartDay returns the TimeStartDay field if non-nil, zero value otherwise.

### GetTimeStartDayOk

`func (o *V0040PowerMgmtData) GetTimeStartDayOk() (*int64, bool)`

GetTimeStartDayOk returns a tuple with the TimeStartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStartDay

`func (o *V0040PowerMgmtData) SetTimeStartDay(v int64)`

SetTimeStartDay sets TimeStartDay field to given value.

### HasTimeStartDay

`func (o *V0040PowerMgmtData) HasTimeStartDay() bool`

HasTimeStartDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


