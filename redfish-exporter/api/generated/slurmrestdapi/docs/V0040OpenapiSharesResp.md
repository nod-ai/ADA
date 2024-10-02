# V0040OpenapiSharesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | [**V0040SharesRespMsg**](V0040SharesRespMsg.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiSharesResp

`func NewV0040OpenapiSharesResp(shares V0040SharesRespMsg, ) *V0040OpenapiSharesResp`

NewV0040OpenapiSharesResp instantiates a new V0040OpenapiSharesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiSharesRespWithDefaults

`func NewV0040OpenapiSharesRespWithDefaults() *V0040OpenapiSharesResp`

NewV0040OpenapiSharesRespWithDefaults instantiates a new V0040OpenapiSharesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *V0040OpenapiSharesResp) GetShares() V0040SharesRespMsg`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0040OpenapiSharesResp) GetSharesOk() (*V0040SharesRespMsg, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0040OpenapiSharesResp) SetShares(v V0040SharesRespMsg)`

SetShares sets Shares field to given value.


### GetMeta

`func (o *V0040OpenapiSharesResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiSharesResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiSharesResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiSharesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiSharesResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiSharesResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiSharesResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiSharesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiSharesResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiSharesResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiSharesResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiSharesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


