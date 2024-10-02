# V0041OpenapiReservationResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservations** | [**[]V0041ReservationInfoMsgInner**](V0041ReservationInfoMsgInner.md) |  | 
**LastUpdate** | [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiReservationResp

`func NewV0041OpenapiReservationResp(reservations []V0041ReservationInfoMsgInner, lastUpdate V0041Uint64NoValStruct, ) *V0041OpenapiReservationResp`

NewV0041OpenapiReservationResp instantiates a new V0041OpenapiReservationResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiReservationRespWithDefaults

`func NewV0041OpenapiReservationRespWithDefaults() *V0041OpenapiReservationResp`

NewV0041OpenapiReservationRespWithDefaults instantiates a new V0041OpenapiReservationResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservations

`func (o *V0041OpenapiReservationResp) GetReservations() []V0041ReservationInfoMsgInner`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *V0041OpenapiReservationResp) GetReservationsOk() (*[]V0041ReservationInfoMsgInner, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *V0041OpenapiReservationResp) SetReservations(v []V0041ReservationInfoMsgInner)`

SetReservations sets Reservations field to given value.


### GetLastUpdate

`func (o *V0041OpenapiReservationResp) GetLastUpdate() V0041Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0041OpenapiReservationResp) GetLastUpdateOk() (*V0041Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0041OpenapiReservationResp) SetLastUpdate(v V0041Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0041OpenapiReservationResp) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiReservationResp) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiReservationResp) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiReservationResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiReservationResp) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiReservationResp) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiReservationResp) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiReservationResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiReservationResp) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiReservationResp) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiReservationResp) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiReservationResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


