# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Status) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Status) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Status) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Status) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Status) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Status) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Status) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Status) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *Status) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Status) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Status) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Status) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


