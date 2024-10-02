# V0039Pings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**Pings** | Pointer to [**[]V0039ControllerPing**](V0039ControllerPing.md) |  | [optional] 

## Methods

### NewV0039Pings

`func NewV0039Pings() *V0039Pings`

NewV0039Pings instantiates a new V0039Pings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039PingsWithDefaults

`func NewV0039PingsWithDefaults() *V0039Pings`

NewV0039PingsWithDefaults instantiates a new V0039Pings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039Pings) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039Pings) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039Pings) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039Pings) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039Pings) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039Pings) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039Pings) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039Pings) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039Pings) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039Pings) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039Pings) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039Pings) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetPings

`func (o *V0039Pings) GetPings() []V0039ControllerPing`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0039Pings) GetPingsOk() (*[]V0039ControllerPing, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0039Pings) SetPings(v []V0039ControllerPing)`

SetPings sets Pings field to given value.

### HasPings

`func (o *V0039Pings) HasPings() bool`

HasPings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


