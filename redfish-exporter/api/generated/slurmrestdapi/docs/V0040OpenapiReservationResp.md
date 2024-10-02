# V0040OpenapiReservationResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservations** | [**[]V0040ReservationInfo**](V0040ReservationInfo.md) |  | 
**LastUpdate** | [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiReservationResp

`func NewV0040OpenapiReservationResp(reservations []V0040ReservationInfo, lastUpdate V0040Uint64NoVal, ) *V0040OpenapiReservationResp`

NewV0040OpenapiReservationResp instantiates a new V0040OpenapiReservationResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiReservationRespWithDefaults

`func NewV0040OpenapiReservationRespWithDefaults() *V0040OpenapiReservationResp`

NewV0040OpenapiReservationRespWithDefaults instantiates a new V0040OpenapiReservationResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservations

`func (o *V0040OpenapiReservationResp) GetReservations() []V0040ReservationInfo`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *V0040OpenapiReservationResp) GetReservationsOk() (*[]V0040ReservationInfo, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *V0040OpenapiReservationResp) SetReservations(v []V0040ReservationInfo)`

SetReservations sets Reservations field to given value.


### GetLastUpdate

`func (o *V0040OpenapiReservationResp) GetLastUpdate() V0040Uint64NoVal`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0040OpenapiReservationResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0040OpenapiReservationResp) SetLastUpdate(v V0040Uint64NoVal)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0040OpenapiReservationResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiReservationResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiReservationResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiReservationResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiReservationResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiReservationResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiReservationResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiReservationResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiReservationResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiReservationResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiReservationResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiReservationResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


