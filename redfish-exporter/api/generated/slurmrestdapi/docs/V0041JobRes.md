# V0041JobRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectType** | **[]string** | Scheduling consumption resource selection type | 
**Nodes** | Pointer to [**V0041JobResNodes**](V0041JobResNodes.md) |  | [optional] 
**Cpus** | **int32** | Number of processors in the allocation | 
**ThreadsPerCore** | [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | 

## Methods

### NewV0041JobRes

`func NewV0041JobRes(selectType []string, cpus int32, threadsPerCore V0041Uint16NoValStruct, ) *V0041JobRes`

NewV0041JobRes instantiates a new V0041JobRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041JobResWithDefaults

`func NewV0041JobResWithDefaults() *V0041JobRes`

NewV0041JobResWithDefaults instantiates a new V0041JobRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectType

`func (o *V0041JobRes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0041JobRes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0041JobRes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.


### GetNodes

`func (o *V0041JobRes) GetNodes() V0041JobResNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041JobRes) GetNodesOk() (*V0041JobResNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041JobRes) SetNodes(v V0041JobResNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041JobRes) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetCpus

`func (o *V0041JobRes) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041JobRes) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041JobRes) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetThreadsPerCore

`func (o *V0041JobRes) GetThreadsPerCore() V0041Uint16NoValStruct`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0041JobRes) GetThreadsPerCoreOk() (*V0041Uint16NoValStruct, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0041JobRes) SetThreadsPerCore(v V0041Uint16NoValStruct)`

SetThreadsPerCore sets ThreadsPerCore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


