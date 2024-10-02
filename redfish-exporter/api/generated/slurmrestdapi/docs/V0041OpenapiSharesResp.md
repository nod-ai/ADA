# V0041OpenapiSharesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | [**V0041SharesRespMsg**](V0041SharesRespMsg.md) |  | 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiSharesResp

`func NewV0041OpenapiSharesResp(shares V0041SharesRespMsg, ) *V0041OpenapiSharesResp`

NewV0041OpenapiSharesResp instantiates a new V0041OpenapiSharesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSharesRespWithDefaults

`func NewV0041OpenapiSharesRespWithDefaults() *V0041OpenapiSharesResp`

NewV0041OpenapiSharesRespWithDefaults instantiates a new V0041OpenapiSharesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *V0041OpenapiSharesResp) GetShares() V0041SharesRespMsg`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0041OpenapiSharesResp) GetSharesOk() (*V0041SharesRespMsg, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0041OpenapiSharesResp) SetShares(v V0041SharesRespMsg)`

SetShares sets Shares field to given value.


### GetMeta

`func (o *V0041OpenapiSharesResp) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiSharesResp) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiSharesResp) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiSharesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiSharesResp) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiSharesResp) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiSharesResp) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiSharesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiSharesResp) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiSharesResp) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiSharesResp) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiSharesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


