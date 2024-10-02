# V0041OpenapiPartitionResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partitions** | [**[]V0041PartitionInfo**](V0041PartitionInfo.md) |  | 
**LastUpdate** | [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiPartitionResp

`func NewV0041OpenapiPartitionResp(partitions []V0041PartitionInfo, lastUpdate V0041Uint64NoValStruct, ) *V0041OpenapiPartitionResp`

NewV0041OpenapiPartitionResp instantiates a new V0041OpenapiPartitionResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiPartitionRespWithDefaults

`func NewV0041OpenapiPartitionRespWithDefaults() *V0041OpenapiPartitionResp`

NewV0041OpenapiPartitionRespWithDefaults instantiates a new V0041OpenapiPartitionResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitions

`func (o *V0041OpenapiPartitionResp) GetPartitions() []V0041PartitionInfo`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0041OpenapiPartitionResp) GetPartitionsOk() (*[]V0041PartitionInfo, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0041OpenapiPartitionResp) SetPartitions(v []V0041PartitionInfo)`

SetPartitions sets Partitions field to given value.


### GetLastUpdate

`func (o *V0041OpenapiPartitionResp) GetLastUpdate() V0041Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0041OpenapiPartitionResp) GetLastUpdateOk() (*V0041Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0041OpenapiPartitionResp) SetLastUpdate(v V0041Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0041OpenapiPartitionResp) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiPartitionResp) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiPartitionResp) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiPartitionResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiPartitionResp) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiPartitionResp) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiPartitionResp) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiPartitionResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiPartitionResp) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiPartitionResp) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiPartitionResp) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiPartitionResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


