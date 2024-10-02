# V0040PartitionInfoMaximums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpusPerNode** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**CpusPerSocket** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**MemoryPerCpu** | Pointer to **int64** |  | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Nodes** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Shares** | Pointer to **int32** |  | [optional] 
**Oversubscribe** | Pointer to [**V0041PartitionInfoMaximumsOversubscribe**](V0041PartitionInfoMaximumsOversubscribe.md) |  | [optional] 
**Time** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**OverTimeLimit** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 

## Methods

### NewV0040PartitionInfoMaximums

`func NewV0040PartitionInfoMaximums() *V0040PartitionInfoMaximums`

NewV0040PartitionInfoMaximums instantiates a new V0040PartitionInfoMaximums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040PartitionInfoMaximumsWithDefaults

`func NewV0040PartitionInfoMaximumsWithDefaults() *V0040PartitionInfoMaximums`

NewV0040PartitionInfoMaximumsWithDefaults instantiates a new V0040PartitionInfoMaximums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpusPerNode

`func (o *V0040PartitionInfoMaximums) GetCpusPerNode() V0040Uint32NoVal`

GetCpusPerNode returns the CpusPerNode field if non-nil, zero value otherwise.

### GetCpusPerNodeOk

`func (o *V0040PartitionInfoMaximums) GetCpusPerNodeOk() (*V0040Uint32NoVal, bool)`

GetCpusPerNodeOk returns a tuple with the CpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerNode

`func (o *V0040PartitionInfoMaximums) SetCpusPerNode(v V0040Uint32NoVal)`

SetCpusPerNode sets CpusPerNode field to given value.

### HasCpusPerNode

`func (o *V0040PartitionInfoMaximums) HasCpusPerNode() bool`

HasCpusPerNode returns a boolean if a field has been set.

### GetCpusPerSocket

`func (o *V0040PartitionInfoMaximums) GetCpusPerSocket() V0040Uint32NoVal`

GetCpusPerSocket returns the CpusPerSocket field if non-nil, zero value otherwise.

### GetCpusPerSocketOk

`func (o *V0040PartitionInfoMaximums) GetCpusPerSocketOk() (*V0040Uint32NoVal, bool)`

GetCpusPerSocketOk returns a tuple with the CpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerSocket

`func (o *V0040PartitionInfoMaximums) SetCpusPerSocket(v V0040Uint32NoVal)`

SetCpusPerSocket sets CpusPerSocket field to given value.

### HasCpusPerSocket

`func (o *V0040PartitionInfoMaximums) HasCpusPerSocket() bool`

HasCpusPerSocket returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0040PartitionInfoMaximums) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0040PartitionInfoMaximums) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0040PartitionInfoMaximums) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0040PartitionInfoMaximums) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerCpu() V0040Uint64NoVal`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerCpuOk() (*V0040Uint64NoVal, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0040PartitionInfoMaximums) SetPartitionMemoryPerCpu(v V0040Uint64NoVal)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0040PartitionInfoMaximums) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerNode() V0040Uint64NoVal`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0040PartitionInfoMaximums) GetPartitionMemoryPerNodeOk() (*V0040Uint64NoVal, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0040PartitionInfoMaximums) SetPartitionMemoryPerNode(v V0040Uint64NoVal)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0040PartitionInfoMaximums) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetNodes

`func (o *V0040PartitionInfoMaximums) GetNodes() V0040Uint32NoVal`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0040PartitionInfoMaximums) GetNodesOk() (*V0040Uint32NoVal, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0040PartitionInfoMaximums) SetNodes(v V0040Uint32NoVal)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0040PartitionInfoMaximums) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetShares

`func (o *V0040PartitionInfoMaximums) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0040PartitionInfoMaximums) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0040PartitionInfoMaximums) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0040PartitionInfoMaximums) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0040PartitionInfoMaximums) GetOversubscribe() V0041PartitionInfoMaximumsOversubscribe`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0040PartitionInfoMaximums) GetOversubscribeOk() (*V0041PartitionInfoMaximumsOversubscribe, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0040PartitionInfoMaximums) SetOversubscribe(v V0041PartitionInfoMaximumsOversubscribe)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0040PartitionInfoMaximums) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetTime

`func (o *V0040PartitionInfoMaximums) GetTime() V0040Uint32NoVal`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0040PartitionInfoMaximums) GetTimeOk() (*V0040Uint32NoVal, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0040PartitionInfoMaximums) SetTime(v V0040Uint32NoVal)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0040PartitionInfoMaximums) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0040PartitionInfoMaximums) GetOverTimeLimit() V0040Uint16NoVal`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0040PartitionInfoMaximums) GetOverTimeLimitOk() (*V0040Uint16NoVal, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0040PartitionInfoMaximums) SetOverTimeLimit(v V0040Uint16NoVal)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0040PartitionInfoMaximums) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


