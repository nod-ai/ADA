# V0039Diag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**Statistics** | Pointer to [**V0039StatsMsg**](V0039StatsMsg.md) |  | [optional] 

## Methods

### NewV0039Diag

`func NewV0039Diag() *V0039Diag`

NewV0039Diag instantiates a new V0039Diag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039DiagWithDefaults

`func NewV0039DiagWithDefaults() *V0039Diag`

NewV0039DiagWithDefaults instantiates a new V0039Diag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039Diag) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039Diag) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039Diag) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039Diag) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039Diag) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039Diag) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039Diag) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039Diag) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039Diag) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039Diag) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039Diag) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039Diag) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetStatistics

`func (o *V0039Diag) GetStatistics() V0039StatsMsg`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0039Diag) GetStatisticsOk() (*V0039StatsMsg, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0039Diag) SetStatistics(v V0039StatsMsg)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0039Diag) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


