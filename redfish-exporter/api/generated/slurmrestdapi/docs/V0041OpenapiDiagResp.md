# V0041OpenapiDiagResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | [**V0041StatsMsg**](V0041StatsMsg.md) |  | 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiDiagResp

`func NewV0041OpenapiDiagResp(statistics V0041StatsMsg, ) *V0041OpenapiDiagResp`

NewV0041OpenapiDiagResp instantiates a new V0041OpenapiDiagResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiDiagRespWithDefaults

`func NewV0041OpenapiDiagRespWithDefaults() *V0041OpenapiDiagResp`

NewV0041OpenapiDiagRespWithDefaults instantiates a new V0041OpenapiDiagResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *V0041OpenapiDiagResp) GetStatistics() V0041StatsMsg`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0041OpenapiDiagResp) GetStatisticsOk() (*V0041StatsMsg, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0041OpenapiDiagResp) SetStatistics(v V0041StatsMsg)`

SetStatistics sets Statistics field to given value.


### GetMeta

`func (o *V0041OpenapiDiagResp) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiDiagResp) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiDiagResp) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiDiagResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiDiagResp) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiDiagResp) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiDiagResp) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiDiagResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiDiagResp) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiDiagResp) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiDiagResp) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiDiagResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


