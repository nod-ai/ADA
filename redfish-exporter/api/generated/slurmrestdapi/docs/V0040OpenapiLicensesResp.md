# V0040OpenapiLicensesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licenses** | [**[]V0040License**](V0040License.md) |  | 
**LastUpdate** | [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiLicensesResp

`func NewV0040OpenapiLicensesResp(licenses []V0040License, lastUpdate V0040Uint64NoVal, ) *V0040OpenapiLicensesResp`

NewV0040OpenapiLicensesResp instantiates a new V0040OpenapiLicensesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiLicensesRespWithDefaults

`func NewV0040OpenapiLicensesRespWithDefaults() *V0040OpenapiLicensesResp`

NewV0040OpenapiLicensesRespWithDefaults instantiates a new V0040OpenapiLicensesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenses

`func (o *V0040OpenapiLicensesResp) GetLicenses() []V0040License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0040OpenapiLicensesResp) GetLicensesOk() (*[]V0040License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0040OpenapiLicensesResp) SetLicenses(v []V0040License)`

SetLicenses sets Licenses field to given value.


### GetLastUpdate

`func (o *V0040OpenapiLicensesResp) GetLastUpdate() V0040Uint64NoVal`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0040OpenapiLicensesResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0040OpenapiLicensesResp) SetLastUpdate(v V0040Uint64NoVal)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0040OpenapiLicensesResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiLicensesResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiLicensesResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiLicensesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiLicensesResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiLicensesResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiLicensesResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiLicensesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiLicensesResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiLicensesResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiLicensesResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiLicensesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


