# V0041PartitionInfoMaximums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpusPerNode** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**CpusPerSocket** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**MemoryPerCpu** | Pointer to **int64** |  | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**Nodes** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Shares** | Pointer to **int32** |  | [optional] 
**Oversubscribe** | Pointer to [**V0041PartitionInfoMaximumsOversubscribe**](V0041PartitionInfoMaximumsOversubscribe.md) |  | [optional] 
**Time** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**OverTimeLimit** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 

## Methods

### NewV0041PartitionInfoMaximums

`func NewV0041PartitionInfoMaximums() *V0041PartitionInfoMaximums`

NewV0041PartitionInfoMaximums instantiates a new V0041PartitionInfoMaximums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041PartitionInfoMaximumsWithDefaults

`func NewV0041PartitionInfoMaximumsWithDefaults() *V0041PartitionInfoMaximums`

NewV0041PartitionInfoMaximumsWithDefaults instantiates a new V0041PartitionInfoMaximums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpusPerNode

`func (o *V0041PartitionInfoMaximums) GetCpusPerNode() V0041Uint32NoValStruct`

GetCpusPerNode returns the CpusPerNode field if non-nil, zero value otherwise.

### GetCpusPerNodeOk

`func (o *V0041PartitionInfoMaximums) GetCpusPerNodeOk() (*V0041Uint32NoValStruct, bool)`

GetCpusPerNodeOk returns a tuple with the CpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerNode

`func (o *V0041PartitionInfoMaximums) SetCpusPerNode(v V0041Uint32NoValStruct)`

SetCpusPerNode sets CpusPerNode field to given value.

### HasCpusPerNode

`func (o *V0041PartitionInfoMaximums) HasCpusPerNode() bool`

HasCpusPerNode returns a boolean if a field has been set.

### GetCpusPerSocket

`func (o *V0041PartitionInfoMaximums) GetCpusPerSocket() V0041Uint32NoValStruct`

GetCpusPerSocket returns the CpusPerSocket field if non-nil, zero value otherwise.

### GetCpusPerSocketOk

`func (o *V0041PartitionInfoMaximums) GetCpusPerSocketOk() (*V0041Uint32NoValStruct, bool)`

GetCpusPerSocketOk returns a tuple with the CpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerSocket

`func (o *V0041PartitionInfoMaximums) SetCpusPerSocket(v V0041Uint32NoValStruct)`

SetCpusPerSocket sets CpusPerSocket field to given value.

### HasCpusPerSocket

`func (o *V0041PartitionInfoMaximums) HasCpusPerSocket() bool`

HasCpusPerSocket returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0041PartitionInfoMaximums) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0041PartitionInfoMaximums) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0041PartitionInfoMaximums) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0041PartitionInfoMaximums) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0041PartitionInfoMaximums) GetPartitionMemoryPerCpu() V0041Uint64NoValStruct`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0041PartitionInfoMaximums) GetPartitionMemoryPerCpuOk() (*V0041Uint64NoValStruct, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0041PartitionInfoMaximums) SetPartitionMemoryPerCpu(v V0041Uint64NoValStruct)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0041PartitionInfoMaximums) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0041PartitionInfoMaximums) GetPartitionMemoryPerNode() V0041Uint64NoValStruct`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0041PartitionInfoMaximums) GetPartitionMemoryPerNodeOk() (*V0041Uint64NoValStruct, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0041PartitionInfoMaximums) SetPartitionMemoryPerNode(v V0041Uint64NoValStruct)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0041PartitionInfoMaximums) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetNodes

`func (o *V0041PartitionInfoMaximums) GetNodes() V0041Uint32NoValStruct`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041PartitionInfoMaximums) GetNodesOk() (*V0041Uint32NoValStruct, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041PartitionInfoMaximums) SetNodes(v V0041Uint32NoValStruct)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041PartitionInfoMaximums) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetShares

`func (o *V0041PartitionInfoMaximums) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0041PartitionInfoMaximums) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0041PartitionInfoMaximums) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0041PartitionInfoMaximums) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0041PartitionInfoMaximums) GetOversubscribe() V0041PartitionInfoMaximumsOversubscribe`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0041PartitionInfoMaximums) GetOversubscribeOk() (*V0041PartitionInfoMaximumsOversubscribe, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0041PartitionInfoMaximums) SetOversubscribe(v V0041PartitionInfoMaximumsOversubscribe)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0041PartitionInfoMaximums) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetTime

`func (o *V0041PartitionInfoMaximums) GetTime() V0041Uint32NoValStruct`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0041PartitionInfoMaximums) GetTimeOk() (*V0041Uint32NoValStruct, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0041PartitionInfoMaximums) SetTime(v V0041Uint32NoValStruct)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0041PartitionInfoMaximums) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0041PartitionInfoMaximums) GetOverTimeLimit() V0041Uint16NoValStruct`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0041PartitionInfoMaximums) GetOverTimeLimitOk() (*V0041Uint16NoValStruct, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0041PartitionInfoMaximums) SetOverTimeLimit(v V0041Uint16NoValStruct)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0041PartitionInfoMaximums) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


