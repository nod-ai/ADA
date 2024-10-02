# V0040OpenapiPartitionResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partitions** | [**[]V0040PartitionInfo**](V0040PartitionInfo.md) |  | 
**LastUpdate** | [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiPartitionResp

`func NewV0040OpenapiPartitionResp(partitions []V0040PartitionInfo, lastUpdate V0040Uint64NoVal, ) *V0040OpenapiPartitionResp`

NewV0040OpenapiPartitionResp instantiates a new V0040OpenapiPartitionResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiPartitionRespWithDefaults

`func NewV0040OpenapiPartitionRespWithDefaults() *V0040OpenapiPartitionResp`

NewV0040OpenapiPartitionRespWithDefaults instantiates a new V0040OpenapiPartitionResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitions

`func (o *V0040OpenapiPartitionResp) GetPartitions() []V0040PartitionInfo`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0040OpenapiPartitionResp) GetPartitionsOk() (*[]V0040PartitionInfo, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0040OpenapiPartitionResp) SetPartitions(v []V0040PartitionInfo)`

SetPartitions sets Partitions field to given value.


### GetLastUpdate

`func (o *V0040OpenapiPartitionResp) GetLastUpdate() V0040Uint64NoVal`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0040OpenapiPartitionResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0040OpenapiPartitionResp) SetLastUpdate(v V0040Uint64NoVal)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0040OpenapiPartitionResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiPartitionResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiPartitionResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiPartitionResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiPartitionResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiPartitionResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiPartitionResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiPartitionResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiPartitionResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiPartitionResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiPartitionResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiPartitionResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


