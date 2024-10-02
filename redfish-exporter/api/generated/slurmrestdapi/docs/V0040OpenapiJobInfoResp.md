# V0040OpenapiJobInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]V0040JobInfo**](V0040JobInfo.md) |  | 
**LastBackfill** | [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | 
**LastUpdate** | [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiJobInfoResp

`func NewV0040OpenapiJobInfoResp(jobs []V0040JobInfo, lastBackfill V0040Uint64NoVal, lastUpdate V0040Uint64NoVal, ) *V0040OpenapiJobInfoResp`

NewV0040OpenapiJobInfoResp instantiates a new V0040OpenapiJobInfoResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiJobInfoRespWithDefaults

`func NewV0040OpenapiJobInfoRespWithDefaults() *V0040OpenapiJobInfoResp`

NewV0040OpenapiJobInfoRespWithDefaults instantiates a new V0040OpenapiJobInfoResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *V0040OpenapiJobInfoResp) GetJobs() []V0040JobInfo`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0040OpenapiJobInfoResp) GetJobsOk() (*[]V0040JobInfo, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0040OpenapiJobInfoResp) SetJobs(v []V0040JobInfo)`

SetJobs sets Jobs field to given value.


### GetLastBackfill

`func (o *V0040OpenapiJobInfoResp) GetLastBackfill() V0040Uint64NoVal`

GetLastBackfill returns the LastBackfill field if non-nil, zero value otherwise.

### GetLastBackfillOk

`func (o *V0040OpenapiJobInfoResp) GetLastBackfillOk() (*V0040Uint64NoVal, bool)`

GetLastBackfillOk returns a tuple with the LastBackfill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackfill

`func (o *V0040OpenapiJobInfoResp) SetLastBackfill(v V0040Uint64NoVal)`

SetLastBackfill sets LastBackfill field to given value.


### GetLastUpdate

`func (o *V0040OpenapiJobInfoResp) GetLastUpdate() V0040Uint64NoVal`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0040OpenapiJobInfoResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0040OpenapiJobInfoResp) SetLastUpdate(v V0040Uint64NoVal)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0040OpenapiJobInfoResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiJobInfoResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiJobInfoResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiJobInfoResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiJobInfoResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiJobInfoResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiJobInfoResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiJobInfoResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiJobInfoResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiJobInfoResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiJobInfoResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiJobInfoResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


