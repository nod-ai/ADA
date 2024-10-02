# V0041SharesRespMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]V0041AssocSharesObjListInner**](V0041AssocSharesObjListInner.md) |  | [optional] 
**TotalShares** | Pointer to **int64** | Total number of shares | [optional] 

## Methods

### NewV0041SharesRespMsg

`func NewV0041SharesRespMsg() *V0041SharesRespMsg`

NewV0041SharesRespMsg instantiates a new V0041SharesRespMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041SharesRespMsgWithDefaults

`func NewV0041SharesRespMsgWithDefaults() *V0041SharesRespMsg`

NewV0041SharesRespMsgWithDefaults instantiates a new V0041SharesRespMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *V0041SharesRespMsg) GetShares() []V0041AssocSharesObjListInner`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0041SharesRespMsg) GetSharesOk() (*[]V0041AssocSharesObjListInner, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0041SharesRespMsg) SetShares(v []V0041AssocSharesObjListInner)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0041SharesRespMsg) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTotalShares

`func (o *V0041SharesRespMsg) GetTotalShares() int64`

GetTotalShares returns the TotalShares field if non-nil, zero value otherwise.

### GetTotalSharesOk

`func (o *V0041SharesRespMsg) GetTotalSharesOk() (*int64, bool)`

GetTotalSharesOk returns a tuple with the TotalShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShares

`func (o *V0041SharesRespMsg) SetTotalShares(v int64)`

SetTotalShares sets TotalShares field to given value.

### HasTotalShares

`func (o *V0041SharesRespMsg) HasTotalShares() bool`

HasTotalShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


