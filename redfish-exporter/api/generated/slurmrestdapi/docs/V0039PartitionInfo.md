# V0039PartitionInfo

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
**Cpus** | Pointer to [**V0041PartitionInfoCpus**](V0041PartitionInfoCpus.md) |  | [optional] 
**Defaults** | Pointer to [**V0039PartitionInfoDefaults**](V0039PartitionInfoDefaults.md) |  | [optional] 
**GraceTime** | Pointer to **int32** |  | [optional] 
**Maximums** | Pointer to [**V0039PartitionInfoMaximums**](V0039PartitionInfoMaximums.md) |  | [optional] 
**Minimums** | Pointer to [**V0041PartitionInfoMinimums**](V0041PartitionInfoMinimums.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeSets** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**V0041PartitionInfoPriority**](V0041PartitionInfoPriority.md) |  | [optional] 
**Timeouts** | Pointer to [**V0039PartitionInfoTimeouts**](V0039PartitionInfoTimeouts.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 

## Methods

### NewV0039PartitionInfo

`func NewV0039PartitionInfo() *V0039PartitionInfo`

NewV0039PartitionInfo instantiates a new V0039PartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039PartitionInfoWithDefaults

`func NewV0039PartitionInfoWithDefaults() *V0039PartitionInfo`

NewV0039PartitionInfoWithDefaults instantiates a new V0039PartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0039PartitionInfo) GetNodes() V0041PartitionInfoNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0039PartitionInfo) GetNodesOk() (*V0041PartitionInfoNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0039PartitionInfo) SetNodes(v V0041PartitionInfoNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0039PartitionInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAccounts

`func (o *V0039PartitionInfo) GetAccounts() V0041PartitionInfoAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0039PartitionInfo) GetAccountsOk() (*V0041PartitionInfoAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0039PartitionInfo) SetAccounts(v V0041PartitionInfoAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0039PartitionInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetGroups

`func (o *V0039PartitionInfo) GetGroups() V0041PartitionInfoGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0039PartitionInfo) GetGroupsOk() (*V0041PartitionInfoGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0039PartitionInfo) SetGroups(v V0041PartitionInfoGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0039PartitionInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetQos

`func (o *V0039PartitionInfo) GetQos() V0041PartitionInfoQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0039PartitionInfo) GetQosOk() (*V0041PartitionInfoQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0039PartitionInfo) SetQos(v V0041PartitionInfoQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0039PartitionInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetAlternate

`func (o *V0039PartitionInfo) GetAlternate() string`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *V0039PartitionInfo) GetAlternateOk() (*string, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *V0039PartitionInfo) SetAlternate(v string)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *V0039PartitionInfo) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetTres

`func (o *V0039PartitionInfo) GetTres() V0041PartitionInfoTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0039PartitionInfo) GetTresOk() (*V0041PartitionInfoTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0039PartitionInfo) SetTres(v V0041PartitionInfoTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0039PartitionInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetCluster

`func (o *V0039PartitionInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0039PartitionInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0039PartitionInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0039PartitionInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCpus

`func (o *V0039PartitionInfo) GetCpus() V0041PartitionInfoCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0039PartitionInfo) GetCpusOk() (*V0041PartitionInfoCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0039PartitionInfo) SetCpus(v V0041PartitionInfoCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0039PartitionInfo) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDefaults

`func (o *V0039PartitionInfo) GetDefaults() V0039PartitionInfoDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *V0039PartitionInfo) GetDefaultsOk() (*V0039PartitionInfoDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *V0039PartitionInfo) SetDefaults(v V0039PartitionInfoDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *V0039PartitionInfo) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetGraceTime

`func (o *V0039PartitionInfo) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0039PartitionInfo) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0039PartitionInfo) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0039PartitionInfo) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMaximums

`func (o *V0039PartitionInfo) GetMaximums() V0039PartitionInfoMaximums`

GetMaximums returns the Maximums field if non-nil, zero value otherwise.

### GetMaximumsOk

`func (o *V0039PartitionInfo) GetMaximumsOk() (*V0039PartitionInfoMaximums, bool)`

GetMaximumsOk returns a tuple with the Maximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximums

`func (o *V0039PartitionInfo) SetMaximums(v V0039PartitionInfoMaximums)`

SetMaximums sets Maximums field to given value.

### HasMaximums

`func (o *V0039PartitionInfo) HasMaximums() bool`

HasMaximums returns a boolean if a field has been set.

### GetMinimums

`func (o *V0039PartitionInfo) GetMinimums() V0041PartitionInfoMinimums`

GetMinimums returns the Minimums field if non-nil, zero value otherwise.

### GetMinimumsOk

`func (o *V0039PartitionInfo) GetMinimumsOk() (*V0041PartitionInfoMinimums, bool)`

GetMinimumsOk returns a tuple with the Minimums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimums

`func (o *V0039PartitionInfo) SetMinimums(v V0041PartitionInfoMinimums)`

SetMinimums sets Minimums field to given value.

### HasMinimums

`func (o *V0039PartitionInfo) HasMinimums() bool`

HasMinimums returns a boolean if a field has been set.

### GetName

`func (o *V0039PartitionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0039PartitionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0039PartitionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0039PartitionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeSets

`func (o *V0039PartitionInfo) GetNodeSets() string`

GetNodeSets returns the NodeSets field if non-nil, zero value otherwise.

### GetNodeSetsOk

`func (o *V0039PartitionInfo) GetNodeSetsOk() (*string, bool)`

GetNodeSetsOk returns a tuple with the NodeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSets

`func (o *V0039PartitionInfo) SetNodeSets(v string)`

SetNodeSets sets NodeSets field to given value.

### HasNodeSets

`func (o *V0039PartitionInfo) HasNodeSets() bool`

HasNodeSets returns a boolean if a field has been set.

### GetPriority

`func (o *V0039PartitionInfo) GetPriority() V0041PartitionInfoPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0039PartitionInfo) GetPriorityOk() (*V0041PartitionInfoPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0039PartitionInfo) SetPriority(v V0041PartitionInfoPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0039PartitionInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeouts

`func (o *V0039PartitionInfo) GetTimeouts() V0039PartitionInfoTimeouts`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *V0039PartitionInfo) GetTimeoutsOk() (*V0039PartitionInfoTimeouts, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *V0039PartitionInfo) SetTimeouts(v V0039PartitionInfoTimeouts)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *V0039PartitionInfo) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0039PartitionInfo) GetSuspendTime() V0039Uint32NoVal`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0039PartitionInfo) GetSuspendTimeOk() (*V0039Uint32NoVal, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0039PartitionInfo) SetSuspendTime(v V0039Uint32NoVal)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0039PartitionInfo) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


