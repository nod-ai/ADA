# V0040OpenapiDiagResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | [**V0040StatsMsg**](V0040StatsMsg.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiDiagResp

`func NewV0040OpenapiDiagResp(statistics V0040StatsMsg, ) *V0040OpenapiDiagResp`

NewV0040OpenapiDiagResp instantiates a new V0040OpenapiDiagResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiDiagRespWithDefaults

`func NewV0040OpenapiDiagRespWithDefaults() *V0040OpenapiDiagResp`

NewV0040OpenapiDiagRespWithDefaults instantiates a new V0040OpenapiDiagResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *V0040OpenapiDiagResp) GetStatistics() V0040StatsMsg`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0040OpenapiDiagResp) GetStatisticsOk() (*V0040StatsMsg, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0040OpenapiDiagResp) SetStatistics(v V0040StatsMsg)`

SetStatistics sets Statistics field to given value.


### GetMeta

`func (o *V0040OpenapiDiagResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiDiagResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiDiagResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiDiagResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiDiagResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiDiagResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiDiagResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiDiagResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiDiagResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiDiagResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiDiagResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiDiagResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


