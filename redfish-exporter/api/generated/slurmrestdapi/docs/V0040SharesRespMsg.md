# V0040SharesRespMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]V0040AssocSharesObjWrap**](V0040AssocSharesObjWrap.md) |  | [optional] 
**TotalShares** | Pointer to **int64** | Total number of shares | [optional] 

## Methods

### NewV0040SharesRespMsg

`func NewV0040SharesRespMsg() *V0040SharesRespMsg`

NewV0040SharesRespMsg instantiates a new V0040SharesRespMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040SharesRespMsgWithDefaults

`func NewV0040SharesRespMsgWithDefaults() *V0040SharesRespMsg`

NewV0040SharesRespMsgWithDefaults instantiates a new V0040SharesRespMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *V0040SharesRespMsg) GetShares() []V0040AssocSharesObjWrap`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0040SharesRespMsg) GetSharesOk() (*[]V0040AssocSharesObjWrap, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0040SharesRespMsg) SetShares(v []V0040AssocSharesObjWrap)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0040SharesRespMsg) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTotalShares

`func (o *V0040SharesRespMsg) GetTotalShares() int64`

GetTotalShares returns the TotalShares field if non-nil, zero value otherwise.

### GetTotalSharesOk

`func (o *V0040SharesRespMsg) GetTotalSharesOk() (*int64, bool)`

GetTotalSharesOk returns a tuple with the TotalShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShares

`func (o *V0040SharesRespMsg) SetTotalShares(v int64)`

SetTotalShares sets TotalShares field to given value.

### HasTotalShares

`func (o *V0040SharesRespMsg) HasTotalShares() bool`

HasTotalShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


