# V0040AssocSharesObjWrap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | assocation id | [optional] 
**Cluster** | Pointer to **string** | cluster name | [optional] 
**Name** | Pointer to **string** | share name | [optional] 
**Parent** | Pointer to **string** | parent name | [optional] 
**Partition** | Pointer to **string** | partition name | [optional] 
**SharesNormalized** | Pointer to [**V0040Float64NoVal**](V0040Float64NoVal.md) |  | [optional] 
**Shares** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Tres** | Pointer to [**V0040AssocSharesObjWrapTres**](V0040AssocSharesObjWrapTres.md) |  | [optional] 
**EffectiveUsage** | Pointer to **float64** | effective, normalized usage | [optional] 
**UsageNormalized** | Pointer to [**V0040Float64NoVal**](V0040Float64NoVal.md) |  | [optional] 
**Usage** | Pointer to **int64** | measure of tresbillableunits usage | [optional] 
**Fairshare** | Pointer to [**V0041AssocSharesObjListInnerFairshare**](V0041AssocSharesObjListInnerFairshare.md) |  | [optional] 
**Type** | Pointer to **[]string** | user or account association | [optional] 

## Methods

### NewV0040AssocSharesObjWrap

`func NewV0040AssocSharesObjWrap() *V0040AssocSharesObjWrap`

NewV0040AssocSharesObjWrap instantiates a new V0040AssocSharesObjWrap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AssocSharesObjWrapWithDefaults

`func NewV0040AssocSharesObjWrapWithDefaults() *V0040AssocSharesObjWrap`

NewV0040AssocSharesObjWrapWithDefaults instantiates a new V0040AssocSharesObjWrap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0040AssocSharesObjWrap) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0040AssocSharesObjWrap) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0040AssocSharesObjWrap) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0040AssocSharesObjWrap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCluster

`func (o *V0040AssocSharesObjWrap) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0040AssocSharesObjWrap) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0040AssocSharesObjWrap) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0040AssocSharesObjWrap) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetName

`func (o *V0040AssocSharesObjWrap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040AssocSharesObjWrap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040AssocSharesObjWrap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040AssocSharesObjWrap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *V0040AssocSharesObjWrap) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0040AssocSharesObjWrap) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0040AssocSharesObjWrap) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0040AssocSharesObjWrap) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPartition

`func (o *V0040AssocSharesObjWrap) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0040AssocSharesObjWrap) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0040AssocSharesObjWrap) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0040AssocSharesObjWrap) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSharesNormalized

`func (o *V0040AssocSharesObjWrap) GetSharesNormalized() V0040Float64NoVal`

GetSharesNormalized returns the SharesNormalized field if non-nil, zero value otherwise.

### GetSharesNormalizedOk

`func (o *V0040AssocSharesObjWrap) GetSharesNormalizedOk() (*V0040Float64NoVal, bool)`

GetSharesNormalizedOk returns a tuple with the SharesNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesNormalized

`func (o *V0040AssocSharesObjWrap) SetSharesNormalized(v V0040Float64NoVal)`

SetSharesNormalized sets SharesNormalized field to given value.

### HasSharesNormalized

`func (o *V0040AssocSharesObjWrap) HasSharesNormalized() bool`

HasSharesNormalized returns a boolean if a field has been set.

### GetShares

`func (o *V0040AssocSharesObjWrap) GetShares() V0040Uint32NoVal`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0040AssocSharesObjWrap) GetSharesOk() (*V0040Uint32NoVal, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0040AssocSharesObjWrap) SetShares(v V0040Uint32NoVal)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0040AssocSharesObjWrap) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTres

`func (o *V0040AssocSharesObjWrap) GetTres() V0040AssocSharesObjWrapTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0040AssocSharesObjWrap) GetTresOk() (*V0040AssocSharesObjWrapTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0040AssocSharesObjWrap) SetTres(v V0040AssocSharesObjWrapTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0040AssocSharesObjWrap) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetEffectiveUsage

`func (o *V0040AssocSharesObjWrap) GetEffectiveUsage() float64`

GetEffectiveUsage returns the EffectiveUsage field if non-nil, zero value otherwise.

### GetEffectiveUsageOk

`func (o *V0040AssocSharesObjWrap) GetEffectiveUsageOk() (*float64, bool)`

GetEffectiveUsageOk returns a tuple with the EffectiveUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveUsage

`func (o *V0040AssocSharesObjWrap) SetEffectiveUsage(v float64)`

SetEffectiveUsage sets EffectiveUsage field to given value.

### HasEffectiveUsage

`func (o *V0040AssocSharesObjWrap) HasEffectiveUsage() bool`

HasEffectiveUsage returns a boolean if a field has been set.

### GetUsageNormalized

`func (o *V0040AssocSharesObjWrap) GetUsageNormalized() V0040Float64NoVal`

GetUsageNormalized returns the UsageNormalized field if non-nil, zero value otherwise.

### GetUsageNormalizedOk

`func (o *V0040AssocSharesObjWrap) GetUsageNormalizedOk() (*V0040Float64NoVal, bool)`

GetUsageNormalizedOk returns a tuple with the UsageNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageNormalized

`func (o *V0040AssocSharesObjWrap) SetUsageNormalized(v V0040Float64NoVal)`

SetUsageNormalized sets UsageNormalized field to given value.

### HasUsageNormalized

`func (o *V0040AssocSharesObjWrap) HasUsageNormalized() bool`

HasUsageNormalized returns a boolean if a field has been set.

### GetUsage

`func (o *V0040AssocSharesObjWrap) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *V0040AssocSharesObjWrap) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *V0040AssocSharesObjWrap) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *V0040AssocSharesObjWrap) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetFairshare

`func (o *V0040AssocSharesObjWrap) GetFairshare() V0041AssocSharesObjListInnerFairshare`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0040AssocSharesObjWrap) GetFairshareOk() (*V0041AssocSharesObjListInnerFairshare, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0040AssocSharesObjWrap) SetFairshare(v V0041AssocSharesObjListInnerFairshare)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0040AssocSharesObjWrap) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.

### GetType

`func (o *V0040AssocSharesObjWrap) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0040AssocSharesObjWrap) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0040AssocSharesObjWrap) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *V0040AssocSharesObjWrap) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


