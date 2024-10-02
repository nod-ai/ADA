# V0039ReservationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**Reservations** | Pointer to [**[]V0039ReservationInfo**](V0039ReservationInfo.md) |  | [optional] 

## Methods

### NewV0039ReservationsResponse

`func NewV0039ReservationsResponse() *V0039ReservationsResponse`

NewV0039ReservationsResponse instantiates a new V0039ReservationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039ReservationsResponseWithDefaults

`func NewV0039ReservationsResponseWithDefaults() *V0039ReservationsResponse`

NewV0039ReservationsResponseWithDefaults instantiates a new V0039ReservationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039ReservationsResponse) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039ReservationsResponse) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039ReservationsResponse) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039ReservationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039ReservationsResponse) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039ReservationsResponse) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039ReservationsResponse) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039ReservationsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039ReservationsResponse) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039ReservationsResponse) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039ReservationsResponse) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039ReservationsResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetReservations

`func (o *V0039ReservationsResponse) GetReservations() []V0039ReservationInfo`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *V0039ReservationsResponse) GetReservationsOk() (*[]V0039ReservationInfo, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *V0039ReservationsResponse) SetReservations(v []V0039ReservationInfo)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *V0039ReservationsResponse) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


