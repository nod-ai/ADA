# V0039PartitionInfoMaximums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpusPerNode** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**CpusPerSocket** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**MemoryPerCpu** | Pointer to **int64** |  | [optional] 
**Nodes** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**Shares** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**OverTimeLimit** | Pointer to [**V0039Uint16NoVal**](V0039Uint16NoVal.md) |  | [optional] 

## Methods

### NewV0039PartitionInfoMaximums

`func NewV0039PartitionInfoMaximums() *V0039PartitionInfoMaximums`

NewV0039PartitionInfoMaximums instantiates a new V0039PartitionInfoMaximums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039PartitionInfoMaximumsWithDefaults

`func NewV0039PartitionInfoMaximumsWithDefaults() *V0039PartitionInfoMaximums`

NewV0039PartitionInfoMaximumsWithDefaults instantiates a new V0039PartitionInfoMaximums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpusPerNode

`func (o *V0039PartitionInfoMaximums) GetCpusPerNode() V0039Uint32NoVal`

GetCpusPerNode returns the CpusPerNode field if non-nil, zero value otherwise.

### GetCpusPerNodeOk

`func (o *V0039PartitionInfoMaximums) GetCpusPerNodeOk() (*V0039Uint32NoVal, bool)`

GetCpusPerNodeOk returns a tuple with the CpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerNode

`func (o *V0039PartitionInfoMaximums) SetCpusPerNode(v V0039Uint32NoVal)`

SetCpusPerNode sets CpusPerNode field to given value.

### HasCpusPerNode

`func (o *V0039PartitionInfoMaximums) HasCpusPerNode() bool`

HasCpusPerNode returns a boolean if a field has been set.

### GetCpusPerSocket

`func (o *V0039PartitionInfoMaximums) GetCpusPerSocket() V0039Uint32NoVal`

GetCpusPerSocket returns the CpusPerSocket field if non-nil, zero value otherwise.

### GetCpusPerSocketOk

`func (o *V0039PartitionInfoMaximums) GetCpusPerSocketOk() (*V0039Uint32NoVal, bool)`

GetCpusPerSocketOk returns a tuple with the CpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerSocket

`func (o *V0039PartitionInfoMaximums) SetCpusPerSocket(v V0039Uint32NoVal)`

SetCpusPerSocket sets CpusPerSocket field to given value.

### HasCpusPerSocket

`func (o *V0039PartitionInfoMaximums) HasCpusPerSocket() bool`

HasCpusPerSocket returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0039PartitionInfoMaximums) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0039PartitionInfoMaximums) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0039PartitionInfoMaximums) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0039PartitionInfoMaximums) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetNodes

`func (o *V0039PartitionInfoMaximums) GetNodes() V0039Uint32NoVal`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0039PartitionInfoMaximums) GetNodesOk() (*V0039Uint32NoVal, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0039PartitionInfoMaximums) SetNodes(v V0039Uint32NoVal)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0039PartitionInfoMaximums) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetShares

`func (o *V0039PartitionInfoMaximums) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0039PartitionInfoMaximums) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0039PartitionInfoMaximums) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0039PartitionInfoMaximums) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTime

`func (o *V0039PartitionInfoMaximums) GetTime() V0039Uint32NoVal`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0039PartitionInfoMaximums) GetTimeOk() (*V0039Uint32NoVal, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0039PartitionInfoMaximums) SetTime(v V0039Uint32NoVal)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0039PartitionInfoMaximums) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0039PartitionInfoMaximums) GetOverTimeLimit() V0039Uint16NoVal`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0039PartitionInfoMaximums) GetOverTimeLimitOk() (*V0039Uint16NoVal, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0039PartitionInfoMaximums) SetOverTimeLimit(v V0039Uint16NoVal)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0039PartitionInfoMaximums) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


