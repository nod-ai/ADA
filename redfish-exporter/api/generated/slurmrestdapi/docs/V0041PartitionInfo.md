# V0041PartitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**V0041PartitionInfoNodes**](V0041PartitionInfoNodes.md) |  | [optional] 
**Accounts** | Pointer to [**V0041PartitionInfoAccounts**](V0041PartitionInfoAccounts.md) |  | [optional] 
**Groups** | Pointer to [**V0041PartitionInfoGroups**](V0041PartitionInfoGroups.md) |  | [optional] 
**Qos** | Pointer to [**V0041PartitionInfoQos**](V0041PartitionInfoQos.md) |  | [optional] 
**Alternate** | Pointer to **string** |  | [optional] 
**Tres** | Pointer to [**V0041PartitionInfoTres**](V0041PartitionInfoTres.md) |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**SelectType** | Pointer to **[]string** | Scheduler consumable resource selection types | [optional] 
**Cpus** | Pointer to [**V0041PartitionInfoCpus**](V0041PartitionInfoCpus.md) |  | [optional] 
**Defaults** | Pointer to [**V0041PartitionInfoDefaults**](V0041PartitionInfoDefaults.md) |  | [optional] 
**GraceTime** | Pointer to **int32** |  | [optional] 
**Maximums** | Pointer to [**V0041PartitionInfoMaximums**](V0041PartitionInfoMaximums.md) |  | [optional] 
**Minimums** | Pointer to [**V0041PartitionInfoMinimums**](V0041PartitionInfoMinimums.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeSets** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**V0041PartitionInfoPriority**](V0041PartitionInfoPriority.md) |  | [optional] 
**Timeouts** | Pointer to [**V0041PartitionInfoTimeouts**](V0041PartitionInfoTimeouts.md) |  | [optional] 
**Partition** | Pointer to [**V0041PartitionInfoPartition**](V0041PartitionInfoPartition.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 

## Methods

### NewV0041PartitionInfo

`func NewV0041PartitionInfo() *V0041PartitionInfo`

NewV0041PartitionInfo instantiates a new V0041PartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041PartitionInfoWithDefaults

`func NewV0041PartitionInfoWithDefaults() *V0041PartitionInfo`

NewV0041PartitionInfoWithDefaults instantiates a new V0041PartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0041PartitionInfo) GetNodes() V0041PartitionInfoNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041PartitionInfo) GetNodesOk() (*V0041PartitionInfoNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041PartitionInfo) SetNodes(v V0041PartitionInfoNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041PartitionInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAccounts

`func (o *V0041PartitionInfo) GetAccounts() V0041PartitionInfoAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041PartitionInfo) GetAccountsOk() (*V0041PartitionInfoAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041PartitionInfo) SetAccounts(v V0041PartitionInfoAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0041PartitionInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetGroups

`func (o *V0041PartitionInfo) GetGroups() V0041PartitionInfoGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0041PartitionInfo) GetGroupsOk() (*V0041PartitionInfoGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0041PartitionInfo) SetGroups(v V0041PartitionInfoGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0041PartitionInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetQos

`func (o *V0041PartitionInfo) GetQos() V0041PartitionInfoQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041PartitionInfo) GetQosOk() (*V0041PartitionInfoQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041PartitionInfo) SetQos(v V0041PartitionInfoQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041PartitionInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetAlternate

`func (o *V0041PartitionInfo) GetAlternate() string`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *V0041PartitionInfo) GetAlternateOk() (*string, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *V0041PartitionInfo) SetAlternate(v string)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *V0041PartitionInfo) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetTres

`func (o *V0041PartitionInfo) GetTres() V0041PartitionInfoTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041PartitionInfo) GetTresOk() (*V0041PartitionInfoTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041PartitionInfo) SetTres(v V0041PartitionInfoTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041PartitionInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetCluster

`func (o *V0041PartitionInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041PartitionInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041PartitionInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041PartitionInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSelectType

`func (o *V0041PartitionInfo) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0041PartitionInfo) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0041PartitionInfo) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0041PartitionInfo) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetCpus

`func (o *V0041PartitionInfo) GetCpus() V0041PartitionInfoCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041PartitionInfo) GetCpusOk() (*V0041PartitionInfoCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041PartitionInfo) SetCpus(v V0041PartitionInfoCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0041PartitionInfo) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDefaults

`func (o *V0041PartitionInfo) GetDefaults() V0041PartitionInfoDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *V0041PartitionInfo) GetDefaultsOk() (*V0041PartitionInfoDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *V0041PartitionInfo) SetDefaults(v V0041PartitionInfoDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *V0041PartitionInfo) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetGraceTime

`func (o *V0041PartitionInfo) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0041PartitionInfo) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0041PartitionInfo) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0041PartitionInfo) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMaximums

`func (o *V0041PartitionInfo) GetMaximums() V0041PartitionInfoMaximums`

GetMaximums returns the Maximums field if non-nil, zero value otherwise.

### GetMaximumsOk

`func (o *V0041PartitionInfo) GetMaximumsOk() (*V0041PartitionInfoMaximums, bool)`

GetMaximumsOk returns a tuple with the Maximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximums

`func (o *V0041PartitionInfo) SetMaximums(v V0041PartitionInfoMaximums)`

SetMaximums sets Maximums field to given value.

### HasMaximums

`func (o *V0041PartitionInfo) HasMaximums() bool`

HasMaximums returns a boolean if a field has been set.

### GetMinimums

`func (o *V0041PartitionInfo) GetMinimums() V0041PartitionInfoMinimums`

GetMinimums returns the Minimums field if non-nil, zero value otherwise.

### GetMinimumsOk

`func (o *V0041PartitionInfo) GetMinimumsOk() (*V0041PartitionInfoMinimums, bool)`

GetMinimumsOk returns a tuple with the Minimums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimums

`func (o *V0041PartitionInfo) SetMinimums(v V0041PartitionInfoMinimums)`

SetMinimums sets Minimums field to given value.

### HasMinimums

`func (o *V0041PartitionInfo) HasMinimums() bool`

HasMinimums returns a boolean if a field has been set.

### GetName

`func (o *V0041PartitionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041PartitionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041PartitionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041PartitionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeSets

`func (o *V0041PartitionInfo) GetNodeSets() string`

GetNodeSets returns the NodeSets field if non-nil, zero value otherwise.

### GetNodeSetsOk

`func (o *V0041PartitionInfo) GetNodeSetsOk() (*string, bool)`

GetNodeSetsOk returns a tuple with the NodeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSets

`func (o *V0041PartitionInfo) SetNodeSets(v string)`

SetNodeSets sets NodeSets field to given value.

### HasNodeSets

`func (o *V0041PartitionInfo) HasNodeSets() bool`

HasNodeSets returns a boolean if a field has been set.

### GetPriority

`func (o *V0041PartitionInfo) GetPriority() V0041PartitionInfoPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041PartitionInfo) GetPriorityOk() (*V0041PartitionInfoPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041PartitionInfo) SetPriority(v V0041PartitionInfoPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041PartitionInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeouts

`func (o *V0041PartitionInfo) GetTimeouts() V0041PartitionInfoTimeouts`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *V0041PartitionInfo) GetTimeoutsOk() (*V0041PartitionInfoTimeouts, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *V0041PartitionInfo) SetTimeouts(v V0041PartitionInfoTimeouts)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *V0041PartitionInfo) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetPartition

`func (o *V0041PartitionInfo) GetPartition() V0041PartitionInfoPartition`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041PartitionInfo) GetPartitionOk() (*V0041PartitionInfoPartition, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041PartitionInfo) SetPartition(v V0041PartitionInfoPartition)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041PartitionInfo) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0041PartitionInfo) GetSuspendTime() V0041Uint32NoValStruct`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0041PartitionInfo) GetSuspendTimeOk() (*V0041Uint32NoValStruct, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0041PartitionInfo) SetSuspendTime(v V0041Uint32NoValStruct)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0041PartitionInfo) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


