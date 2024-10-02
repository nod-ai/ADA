# V0041OpenapiLicensesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licenses** | [**[]V0041LicensesInner**](V0041LicensesInner.md) |  | 
**LastUpdate** | [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiLicensesResp

`func NewV0041OpenapiLicensesResp(licenses []V0041LicensesInner, lastUpdate V0041Uint64NoValStruct, ) *V0041OpenapiLicensesResp`

NewV0041OpenapiLicensesResp instantiates a new V0041OpenapiLicensesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiLicensesRespWithDefaults

`func NewV0041OpenapiLicensesRespWithDefaults() *V0041OpenapiLicensesResp`

NewV0041OpenapiLicensesRespWithDefaults instantiates a new V0041OpenapiLicensesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenses

`func (o *V0041OpenapiLicensesResp) GetLicenses() []V0041LicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0041OpenapiLicensesResp) GetLicensesOk() (*[]V0041LicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0041OpenapiLicensesResp) SetLicenses(v []V0041LicensesInner)`

SetLicenses sets Licenses field to given value.


### GetLastUpdate

`func (o *V0041OpenapiLicensesResp) GetLastUpdate() V0041Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0041OpenapiLicensesResp) GetLastUpdateOk() (*V0041Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0041OpenapiLicensesResp) SetLastUpdate(v V0041Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0041OpenapiLicensesResp) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiLicensesResp) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiLicensesResp) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiLicensesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiLicensesResp) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiLicensesResp) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiLicensesResp) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiLicensesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiLicensesResp) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiLicensesResp) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiLicensesResp) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiLicensesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


