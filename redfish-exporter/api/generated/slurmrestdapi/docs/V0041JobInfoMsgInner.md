# V0041JobInfoMsgInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AccrueTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**AdminComment** | Pointer to **string** |  | [optional] 
**AllocatingNode** | Pointer to **string** |  | [optional] 
**ArrayJobId** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**ArrayTaskId** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**ArrayMaxTasks** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**ArrayTaskString** | Pointer to **string** |  | [optional] 
**AssociationId** | Pointer to **int32** |  | [optional] 
**BatchFeatures** | Pointer to **string** |  | [optional] 
**BatchFlag** | Pointer to **bool** |  | [optional] 
**BatchHost** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**BurstBuffer** | Pointer to **string** |  | [optional] 
**BurstBufferState** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ClusterFeatures** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 
**Contiguous** | Pointer to **bool** |  | [optional] 
**CoreSpec** | Pointer to **int32** |  | [optional] 
**ThreadSpec** | Pointer to **int32** |  | [optional] 
**CoresPerSocket** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**BillableTres** | Pointer to [**V0041Float64NoValStruct**](V0041Float64NoValStruct.md) |  | [optional] 
**CpusPerTask** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**CpuFrequencyMinimum** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**CpuFrequencyMaximum** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**CpuFrequencyGovernor** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**CpusPerTres** | Pointer to **string** |  | [optional] 
**Cron** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**DelayBoot** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Dependency** | Pointer to **string** |  | [optional] 
**DerivedExitCode** | Pointer to [**V0041ProcessExitCodeVerbose**](V0041ProcessExitCodeVerbose.md) |  | [optional] 
**EligibleTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**EndTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**ExcludedNodes** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to [**V0041ProcessExitCodeVerbose**](V0041ProcessExitCodeVerbose.md) |  | [optional] 
**Extra** | Pointer to **string** |  | [optional] 
**FailedNode** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**FederationOrigin** | Pointer to **string** |  | [optional] 
**FederationSiblingsActive** | Pointer to **string** |  | [optional] 
**FederationSiblingsViable** | Pointer to **string** |  | [optional] 
**GresDetail** | Pointer to **[]string** |  | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**HetJobId** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**HetJobIdSet** | Pointer to **string** |  | [optional] 
**HetJobOffset** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**JobId** | Pointer to **int32** |  | [optional] 
**JobResources** | Pointer to [**V0041JobRes**](V0041JobRes.md) |  | [optional] 
**JobSizeStr** | Pointer to **[]string** |  | [optional] 
**JobState** | Pointer to **[]string** |  | [optional] 
**LastSchedEvaluation** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**Licenses** | Pointer to **string** |  | [optional] 
**MailType** | Pointer to **[]string** |  | [optional] 
**MailUser** | Pointer to **string** |  | [optional] 
**MaxCpus** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**MaxNodes** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**McsLabel** | Pointer to **string** |  | [optional] 
**MemoryPerTres** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**Nice** | Pointer to **int32** |  | [optional] 
**TasksPerCore** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**TasksPerTres** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**TasksPerNode** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**TasksPerSocket** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**TasksPerBoard** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**Cpus** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**NodeCount** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Tasks** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Prefer** | Pointer to **string** |  | [optional] 
**MemoryPerCpu** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**MinimumCpusPerNode** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**MinimumTmpDiskPerNode** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Power** | Pointer to [**V0041JobInfoMsgInnerPower**](V0041JobInfoMsgInnerPower.md) |  | [optional] 
**PreemptTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**PreemptableTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**PreSusTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**Priority** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Profile** | Pointer to **[]string** |  | [optional] 
**Qos** | Pointer to **string** |  | [optional] 
**Reboot** | Pointer to **bool** |  | [optional] 
**RequiredNodes** | Pointer to **string** |  | [optional] 
**MinimumSwitches** | Pointer to **int32** |  | [optional] 
**Requeue** | Pointer to **bool** |  | [optional] 
**ResizeTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**RestartCnt** | Pointer to **int32** |  | [optional] 
**ResvName** | Pointer to **string** |  | [optional] 
**ScheduledNodes** | Pointer to **string** |  | [optional] 
**SelinuxContext** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **[]string** |  | [optional] 
**Exclusive** | Pointer to **[]string** |  | [optional] 
**Oversubscribe** | Pointer to **bool** |  | [optional] 
**ShowFlags** | Pointer to **[]string** |  | [optional] 
**SocketsPerBoard** | Pointer to **int32** |  | [optional] 
**SocketsPerNode** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**StartTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**StateDescription** | Pointer to **string** |  | [optional] 
**StateReason** | Pointer to **string** |  | [optional] 
**StandardError** | Pointer to **string** |  | [optional] 
**StandardInput** | Pointer to **string** |  | [optional] 
**StandardOutput** | Pointer to **string** |  | [optional] 
**SubmitTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**SystemComment** | Pointer to **string** |  | [optional] 
**TimeLimit** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**TimeMinimum** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**ThreadsPerCore** | Pointer to [**V0041Uint16NoValStruct**](V0041Uint16NoValStruct.md) |  | [optional] 
**TresBind** | Pointer to **string** |  | [optional] 
**TresFreq** | Pointer to **string** |  | [optional] 
**TresPerJob** | Pointer to **string** |  | [optional] 
**TresPerNode** | Pointer to **string** |  | [optional] 
**TresPerSocket** | Pointer to **string** |  | [optional] 
**TresPerTask** | Pointer to **string** |  | [optional] 
**TresReqStr** | Pointer to **string** |  | [optional] 
**TresAllocStr** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**MaximumSwitchWaitTime** | Pointer to **int32** |  | [optional] 
**Wckey** | Pointer to **string** |  | [optional] 
**CurrentWorkingDirectory** | Pointer to **string** |  | [optional] 

## Methods

### NewV0041JobInfoMsgInner

`func NewV0041JobInfoMsgInner() *V0041JobInfoMsgInner`

NewV0041JobInfoMsgInner instantiates a new V0041JobInfoMsgInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041JobInfoMsgInnerWithDefaults

`func NewV0041JobInfoMsgInnerWithDefaults() *V0041JobInfoMsgInner`

NewV0041JobInfoMsgInnerWithDefaults instantiates a new V0041JobInfoMsgInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0041JobInfoMsgInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0041JobInfoMsgInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0041JobInfoMsgInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0041JobInfoMsgInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccrueTime

`func (o *V0041JobInfoMsgInner) GetAccrueTime() V0041Uint64NoValStruct`

GetAccrueTime returns the AccrueTime field if non-nil, zero value otherwise.

### GetAccrueTimeOk

`func (o *V0041JobInfoMsgInner) GetAccrueTimeOk() (*V0041Uint64NoValStruct, bool)`

GetAccrueTimeOk returns a tuple with the AccrueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrueTime

`func (o *V0041JobInfoMsgInner) SetAccrueTime(v V0041Uint64NoValStruct)`

SetAccrueTime sets AccrueTime field to given value.

### HasAccrueTime

`func (o *V0041JobInfoMsgInner) HasAccrueTime() bool`

HasAccrueTime returns a boolean if a field has been set.

### GetAdminComment

`func (o *V0041JobInfoMsgInner) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *V0041JobInfoMsgInner) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *V0041JobInfoMsgInner) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *V0041JobInfoMsgInner) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAllocatingNode

`func (o *V0041JobInfoMsgInner) GetAllocatingNode() string`

GetAllocatingNode returns the AllocatingNode field if non-nil, zero value otherwise.

### GetAllocatingNodeOk

`func (o *V0041JobInfoMsgInner) GetAllocatingNodeOk() (*string, bool)`

GetAllocatingNodeOk returns a tuple with the AllocatingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatingNode

`func (o *V0041JobInfoMsgInner) SetAllocatingNode(v string)`

SetAllocatingNode sets AllocatingNode field to given value.

### HasAllocatingNode

`func (o *V0041JobInfoMsgInner) HasAllocatingNode() bool`

HasAllocatingNode returns a boolean if a field has been set.

### GetArrayJobId

`func (o *V0041JobInfoMsgInner) GetArrayJobId() V0041Uint32NoValStruct`

GetArrayJobId returns the ArrayJobId field if non-nil, zero value otherwise.

### GetArrayJobIdOk

`func (o *V0041JobInfoMsgInner) GetArrayJobIdOk() (*V0041Uint32NoValStruct, bool)`

GetArrayJobIdOk returns a tuple with the ArrayJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayJobId

`func (o *V0041JobInfoMsgInner) SetArrayJobId(v V0041Uint32NoValStruct)`

SetArrayJobId sets ArrayJobId field to given value.

### HasArrayJobId

`func (o *V0041JobInfoMsgInner) HasArrayJobId() bool`

HasArrayJobId returns a boolean if a field has been set.

### GetArrayTaskId

`func (o *V0041JobInfoMsgInner) GetArrayTaskId() V0041Uint32NoValStruct`

GetArrayTaskId returns the ArrayTaskId field if non-nil, zero value otherwise.

### GetArrayTaskIdOk

`func (o *V0041JobInfoMsgInner) GetArrayTaskIdOk() (*V0041Uint32NoValStruct, bool)`

GetArrayTaskIdOk returns a tuple with the ArrayTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskId

`func (o *V0041JobInfoMsgInner) SetArrayTaskId(v V0041Uint32NoValStruct)`

SetArrayTaskId sets ArrayTaskId field to given value.

### HasArrayTaskId

`func (o *V0041JobInfoMsgInner) HasArrayTaskId() bool`

HasArrayTaskId returns a boolean if a field has been set.

### GetArrayMaxTasks

`func (o *V0041JobInfoMsgInner) GetArrayMaxTasks() V0041Uint32NoValStruct`

GetArrayMaxTasks returns the ArrayMaxTasks field if non-nil, zero value otherwise.

### GetArrayMaxTasksOk

`func (o *V0041JobInfoMsgInner) GetArrayMaxTasksOk() (*V0041Uint32NoValStruct, bool)`

GetArrayMaxTasksOk returns a tuple with the ArrayMaxTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayMaxTasks

`func (o *V0041JobInfoMsgInner) SetArrayMaxTasks(v V0041Uint32NoValStruct)`

SetArrayMaxTasks sets ArrayMaxTasks field to given value.

### HasArrayMaxTasks

`func (o *V0041JobInfoMsgInner) HasArrayMaxTasks() bool`

HasArrayMaxTasks returns a boolean if a field has been set.

### GetArrayTaskString

`func (o *V0041JobInfoMsgInner) GetArrayTaskString() string`

GetArrayTaskString returns the ArrayTaskString field if non-nil, zero value otherwise.

### GetArrayTaskStringOk

`func (o *V0041JobInfoMsgInner) GetArrayTaskStringOk() (*string, bool)`

GetArrayTaskStringOk returns a tuple with the ArrayTaskString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskString

`func (o *V0041JobInfoMsgInner) SetArrayTaskString(v string)`

SetArrayTaskString sets ArrayTaskString field to given value.

### HasArrayTaskString

`func (o *V0041JobInfoMsgInner) HasArrayTaskString() bool`

HasArrayTaskString returns a boolean if a field has been set.

### GetAssociationId

`func (o *V0041JobInfoMsgInner) GetAssociationId() int32`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *V0041JobInfoMsgInner) GetAssociationIdOk() (*int32, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *V0041JobInfoMsgInner) SetAssociationId(v int32)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *V0041JobInfoMsgInner) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0041JobInfoMsgInner) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0041JobInfoMsgInner) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0041JobInfoMsgInner) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0041JobInfoMsgInner) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBatchFlag

`func (o *V0041JobInfoMsgInner) GetBatchFlag() bool`

GetBatchFlag returns the BatchFlag field if non-nil, zero value otherwise.

### GetBatchFlagOk

`func (o *V0041JobInfoMsgInner) GetBatchFlagOk() (*bool, bool)`

GetBatchFlagOk returns a tuple with the BatchFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFlag

`func (o *V0041JobInfoMsgInner) SetBatchFlag(v bool)`

SetBatchFlag sets BatchFlag field to given value.

### HasBatchFlag

`func (o *V0041JobInfoMsgInner) HasBatchFlag() bool`

HasBatchFlag returns a boolean if a field has been set.

### GetBatchHost

`func (o *V0041JobInfoMsgInner) GetBatchHost() string`

GetBatchHost returns the BatchHost field if non-nil, zero value otherwise.

### GetBatchHostOk

`func (o *V0041JobInfoMsgInner) GetBatchHostOk() (*string, bool)`

GetBatchHostOk returns a tuple with the BatchHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchHost

`func (o *V0041JobInfoMsgInner) SetBatchHost(v string)`

SetBatchHost sets BatchHost field to given value.

### HasBatchHost

`func (o *V0041JobInfoMsgInner) HasBatchHost() bool`

HasBatchHost returns a boolean if a field has been set.

### GetFlags

`func (o *V0041JobInfoMsgInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041JobInfoMsgInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041JobInfoMsgInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041JobInfoMsgInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0041JobInfoMsgInner) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0041JobInfoMsgInner) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0041JobInfoMsgInner) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0041JobInfoMsgInner) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetBurstBufferState

`func (o *V0041JobInfoMsgInner) GetBurstBufferState() string`

GetBurstBufferState returns the BurstBufferState field if non-nil, zero value otherwise.

### GetBurstBufferStateOk

`func (o *V0041JobInfoMsgInner) GetBurstBufferStateOk() (*string, bool)`

GetBurstBufferStateOk returns a tuple with the BurstBufferState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBufferState

`func (o *V0041JobInfoMsgInner) SetBurstBufferState(v string)`

SetBurstBufferState sets BurstBufferState field to given value.

### HasBurstBufferState

`func (o *V0041JobInfoMsgInner) HasBurstBufferState() bool`

HasBurstBufferState returns a boolean if a field has been set.

### GetCluster

`func (o *V0041JobInfoMsgInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041JobInfoMsgInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041JobInfoMsgInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041JobInfoMsgInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterFeatures

`func (o *V0041JobInfoMsgInner) GetClusterFeatures() string`

GetClusterFeatures returns the ClusterFeatures field if non-nil, zero value otherwise.

### GetClusterFeaturesOk

`func (o *V0041JobInfoMsgInner) GetClusterFeaturesOk() (*string, bool)`

GetClusterFeaturesOk returns a tuple with the ClusterFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFeatures

`func (o *V0041JobInfoMsgInner) SetClusterFeatures(v string)`

SetClusterFeatures sets ClusterFeatures field to given value.

### HasClusterFeatures

`func (o *V0041JobInfoMsgInner) HasClusterFeatures() bool`

HasClusterFeatures returns a boolean if a field has been set.

### GetCommand

`func (o *V0041JobInfoMsgInner) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0041JobInfoMsgInner) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0041JobInfoMsgInner) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0041JobInfoMsgInner) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetComment

`func (o *V0041JobInfoMsgInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0041JobInfoMsgInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0041JobInfoMsgInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0041JobInfoMsgInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContainer

`func (o *V0041JobInfoMsgInner) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0041JobInfoMsgInner) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0041JobInfoMsgInner) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0041JobInfoMsgInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *V0041JobInfoMsgInner) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *V0041JobInfoMsgInner) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *V0041JobInfoMsgInner) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *V0041JobInfoMsgInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetContiguous

`func (o *V0041JobInfoMsgInner) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *V0041JobInfoMsgInner) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *V0041JobInfoMsgInner) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *V0041JobInfoMsgInner) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetCoreSpec

`func (o *V0041JobInfoMsgInner) GetCoreSpec() int32`

GetCoreSpec returns the CoreSpec field if non-nil, zero value otherwise.

### GetCoreSpecOk

`func (o *V0041JobInfoMsgInner) GetCoreSpecOk() (*int32, bool)`

GetCoreSpecOk returns a tuple with the CoreSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpec

`func (o *V0041JobInfoMsgInner) SetCoreSpec(v int32)`

SetCoreSpec sets CoreSpec field to given value.

### HasCoreSpec

`func (o *V0041JobInfoMsgInner) HasCoreSpec() bool`

HasCoreSpec returns a boolean if a field has been set.

### GetThreadSpec

`func (o *V0041JobInfoMsgInner) GetThreadSpec() int32`

GetThreadSpec returns the ThreadSpec field if non-nil, zero value otherwise.

### GetThreadSpecOk

`func (o *V0041JobInfoMsgInner) GetThreadSpecOk() (*int32, bool)`

GetThreadSpecOk returns a tuple with the ThreadSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpec

`func (o *V0041JobInfoMsgInner) SetThreadSpec(v int32)`

SetThreadSpec sets ThreadSpec field to given value.

### HasThreadSpec

`func (o *V0041JobInfoMsgInner) HasThreadSpec() bool`

HasThreadSpec returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *V0041JobInfoMsgInner) GetCoresPerSocket() V0041Uint16NoValStruct`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *V0041JobInfoMsgInner) GetCoresPerSocketOk() (*V0041Uint16NoValStruct, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *V0041JobInfoMsgInner) SetCoresPerSocket(v V0041Uint16NoValStruct)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *V0041JobInfoMsgInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetBillableTres

`func (o *V0041JobInfoMsgInner) GetBillableTres() V0041Float64NoValStruct`

GetBillableTres returns the BillableTres field if non-nil, zero value otherwise.

### GetBillableTresOk

`func (o *V0041JobInfoMsgInner) GetBillableTresOk() (*V0041Float64NoValStruct, bool)`

GetBillableTresOk returns a tuple with the BillableTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTres

`func (o *V0041JobInfoMsgInner) SetBillableTres(v V0041Float64NoValStruct)`

SetBillableTres sets BillableTres field to given value.

### HasBillableTres

`func (o *V0041JobInfoMsgInner) HasBillableTres() bool`

HasBillableTres returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0041JobInfoMsgInner) GetCpusPerTask() V0041Uint16NoValStruct`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0041JobInfoMsgInner) GetCpusPerTaskOk() (*V0041Uint16NoValStruct, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0041JobInfoMsgInner) SetCpusPerTask(v V0041Uint16NoValStruct)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0041JobInfoMsgInner) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetCpuFrequencyMinimum

`func (o *V0041JobInfoMsgInner) GetCpuFrequencyMinimum() V0041Uint32NoValStruct`

GetCpuFrequencyMinimum returns the CpuFrequencyMinimum field if non-nil, zero value otherwise.

### GetCpuFrequencyMinimumOk

`func (o *V0041JobInfoMsgInner) GetCpuFrequencyMinimumOk() (*V0041Uint32NoValStruct, bool)`

GetCpuFrequencyMinimumOk returns a tuple with the CpuFrequencyMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMinimum

`func (o *V0041JobInfoMsgInner) SetCpuFrequencyMinimum(v V0041Uint32NoValStruct)`

SetCpuFrequencyMinimum sets CpuFrequencyMinimum field to given value.

### HasCpuFrequencyMinimum

`func (o *V0041JobInfoMsgInner) HasCpuFrequencyMinimum() bool`

HasCpuFrequencyMinimum returns a boolean if a field has been set.

### GetCpuFrequencyMaximum

`func (o *V0041JobInfoMsgInner) GetCpuFrequencyMaximum() V0041Uint32NoValStruct`

GetCpuFrequencyMaximum returns the CpuFrequencyMaximum field if non-nil, zero value otherwise.

### GetCpuFrequencyMaximumOk

`func (o *V0041JobInfoMsgInner) GetCpuFrequencyMaximumOk() (*V0041Uint32NoValStruct, bool)`

GetCpuFrequencyMaximumOk returns a tuple with the CpuFrequencyMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMaximum

`func (o *V0041JobInfoMsgInner) SetCpuFrequencyMaximum(v V0041Uint32NoValStruct)`

SetCpuFrequencyMaximum sets CpuFrequencyMaximum field to given value.

### HasCpuFrequencyMaximum

`func (o *V0041JobInfoMsgInner) HasCpuFrequencyMaximum() bool`

HasCpuFrequencyMaximum returns a boolean if a field has been set.

### GetCpuFrequencyGovernor

`func (o *V0041JobInfoMsgInner) GetCpuFrequencyGovernor() V0041Uint32NoValStruct`

GetCpuFrequencyGovernor returns the CpuFrequencyGovernor field if non-nil, zero value otherwise.

### GetCpuFrequencyGovernorOk

`func (o *V0041JobInfoMsgInner) GetCpuFrequencyGovernorOk() (*V0041Uint32NoValStruct, bool)`

GetCpuFrequencyGovernorOk returns a tuple with the CpuFrequencyGovernor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyGovernor

`func (o *V0041JobInfoMsgInner) SetCpuFrequencyGovernor(v V0041Uint32NoValStruct)`

SetCpuFrequencyGovernor sets CpuFrequencyGovernor field to given value.

### HasCpuFrequencyGovernor

`func (o *V0041JobInfoMsgInner) HasCpuFrequencyGovernor() bool`

HasCpuFrequencyGovernor returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *V0041JobInfoMsgInner) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *V0041JobInfoMsgInner) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *V0041JobInfoMsgInner) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *V0041JobInfoMsgInner) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetCron

`func (o *V0041JobInfoMsgInner) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *V0041JobInfoMsgInner) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *V0041JobInfoMsgInner) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *V0041JobInfoMsgInner) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDeadline

`func (o *V0041JobInfoMsgInner) GetDeadline() V0041Uint64NoValStruct`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0041JobInfoMsgInner) GetDeadlineOk() (*V0041Uint64NoValStruct, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0041JobInfoMsgInner) SetDeadline(v V0041Uint64NoValStruct)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0041JobInfoMsgInner) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0041JobInfoMsgInner) GetDelayBoot() V0041Uint32NoValStruct`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0041JobInfoMsgInner) GetDelayBootOk() (*V0041Uint32NoValStruct, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0041JobInfoMsgInner) SetDelayBoot(v V0041Uint32NoValStruct)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0041JobInfoMsgInner) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0041JobInfoMsgInner) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0041JobInfoMsgInner) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0041JobInfoMsgInner) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0041JobInfoMsgInner) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *V0041JobInfoMsgInner) GetDerivedExitCode() V0041ProcessExitCodeVerbose`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *V0041JobInfoMsgInner) GetDerivedExitCodeOk() (*V0041ProcessExitCodeVerbose, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *V0041JobInfoMsgInner) SetDerivedExitCode(v V0041ProcessExitCodeVerbose)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *V0041JobInfoMsgInner) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetEligibleTime

`func (o *V0041JobInfoMsgInner) GetEligibleTime() V0041Uint64NoValStruct`

GetEligibleTime returns the EligibleTime field if non-nil, zero value otherwise.

### GetEligibleTimeOk

`func (o *V0041JobInfoMsgInner) GetEligibleTimeOk() (*V0041Uint64NoValStruct, bool)`

GetEligibleTimeOk returns a tuple with the EligibleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleTime

`func (o *V0041JobInfoMsgInner) SetEligibleTime(v V0041Uint64NoValStruct)`

SetEligibleTime sets EligibleTime field to given value.

### HasEligibleTime

`func (o *V0041JobInfoMsgInner) HasEligibleTime() bool`

HasEligibleTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V0041JobInfoMsgInner) GetEndTime() V0041Uint64NoValStruct`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0041JobInfoMsgInner) GetEndTimeOk() (*V0041Uint64NoValStruct, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0041JobInfoMsgInner) SetEndTime(v V0041Uint64NoValStruct)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0041JobInfoMsgInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *V0041JobInfoMsgInner) GetExcludedNodes() string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *V0041JobInfoMsgInner) GetExcludedNodesOk() (*string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *V0041JobInfoMsgInner) SetExcludedNodes(v string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *V0041JobInfoMsgInner) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExitCode

`func (o *V0041JobInfoMsgInner) GetExitCode() V0041ProcessExitCodeVerbose`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0041JobInfoMsgInner) GetExitCodeOk() (*V0041ProcessExitCodeVerbose, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0041JobInfoMsgInner) SetExitCode(v V0041ProcessExitCodeVerbose)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0041JobInfoMsgInner) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExtra

`func (o *V0041JobInfoMsgInner) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0041JobInfoMsgInner) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0041JobInfoMsgInner) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0041JobInfoMsgInner) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFailedNode

`func (o *V0041JobInfoMsgInner) GetFailedNode() string`

GetFailedNode returns the FailedNode field if non-nil, zero value otherwise.

### GetFailedNodeOk

`func (o *V0041JobInfoMsgInner) GetFailedNodeOk() (*string, bool)`

GetFailedNodeOk returns a tuple with the FailedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNode

`func (o *V0041JobInfoMsgInner) SetFailedNode(v string)`

SetFailedNode sets FailedNode field to given value.

### HasFailedNode

`func (o *V0041JobInfoMsgInner) HasFailedNode() bool`

HasFailedNode returns a boolean if a field has been set.

### GetFeatures

`func (o *V0041JobInfoMsgInner) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0041JobInfoMsgInner) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0041JobInfoMsgInner) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0041JobInfoMsgInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFederationOrigin

`func (o *V0041JobInfoMsgInner) GetFederationOrigin() string`

GetFederationOrigin returns the FederationOrigin field if non-nil, zero value otherwise.

### GetFederationOriginOk

`func (o *V0041JobInfoMsgInner) GetFederationOriginOk() (*string, bool)`

GetFederationOriginOk returns a tuple with the FederationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationOrigin

`func (o *V0041JobInfoMsgInner) SetFederationOrigin(v string)`

SetFederationOrigin sets FederationOrigin field to given value.

### HasFederationOrigin

`func (o *V0041JobInfoMsgInner) HasFederationOrigin() bool`

HasFederationOrigin returns a boolean if a field has been set.

### GetFederationSiblingsActive

`func (o *V0041JobInfoMsgInner) GetFederationSiblingsActive() string`

GetFederationSiblingsActive returns the FederationSiblingsActive field if non-nil, zero value otherwise.

### GetFederationSiblingsActiveOk

`func (o *V0041JobInfoMsgInner) GetFederationSiblingsActiveOk() (*string, bool)`

GetFederationSiblingsActiveOk returns a tuple with the FederationSiblingsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsActive

`func (o *V0041JobInfoMsgInner) SetFederationSiblingsActive(v string)`

SetFederationSiblingsActive sets FederationSiblingsActive field to given value.

### HasFederationSiblingsActive

`func (o *V0041JobInfoMsgInner) HasFederationSiblingsActive() bool`

HasFederationSiblingsActive returns a boolean if a field has been set.

### GetFederationSiblingsViable

`func (o *V0041JobInfoMsgInner) GetFederationSiblingsViable() string`

GetFederationSiblingsViable returns the FederationSiblingsViable field if non-nil, zero value otherwise.

### GetFederationSiblingsViableOk

`func (o *V0041JobInfoMsgInner) GetFederationSiblingsViableOk() (*string, bool)`

GetFederationSiblingsViableOk returns a tuple with the FederationSiblingsViable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsViable

`func (o *V0041JobInfoMsgInner) SetFederationSiblingsViable(v string)`

SetFederationSiblingsViable sets FederationSiblingsViable field to given value.

### HasFederationSiblingsViable

`func (o *V0041JobInfoMsgInner) HasFederationSiblingsViable() bool`

HasFederationSiblingsViable returns a boolean if a field has been set.

### GetGresDetail

`func (o *V0041JobInfoMsgInner) GetGresDetail() []string`

GetGresDetail returns the GresDetail field if non-nil, zero value otherwise.

### GetGresDetailOk

`func (o *V0041JobInfoMsgInner) GetGresDetailOk() (*[]string, bool)`

GetGresDetailOk returns a tuple with the GresDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDetail

`func (o *V0041JobInfoMsgInner) SetGresDetail(v []string)`

SetGresDetail sets GresDetail field to given value.

### HasGresDetail

`func (o *V0041JobInfoMsgInner) HasGresDetail() bool`

HasGresDetail returns a boolean if a field has been set.

### GetGroupId

`func (o *V0041JobInfoMsgInner) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V0041JobInfoMsgInner) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V0041JobInfoMsgInner) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V0041JobInfoMsgInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *V0041JobInfoMsgInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *V0041JobInfoMsgInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *V0041JobInfoMsgInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *V0041JobInfoMsgInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetHetJobId

`func (o *V0041JobInfoMsgInner) GetHetJobId() V0041Uint32NoValStruct`

GetHetJobId returns the HetJobId field if non-nil, zero value otherwise.

### GetHetJobIdOk

`func (o *V0041JobInfoMsgInner) GetHetJobIdOk() (*V0041Uint32NoValStruct, bool)`

GetHetJobIdOk returns a tuple with the HetJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobId

`func (o *V0041JobInfoMsgInner) SetHetJobId(v V0041Uint32NoValStruct)`

SetHetJobId sets HetJobId field to given value.

### HasHetJobId

`func (o *V0041JobInfoMsgInner) HasHetJobId() bool`

HasHetJobId returns a boolean if a field has been set.

### GetHetJobIdSet

`func (o *V0041JobInfoMsgInner) GetHetJobIdSet() string`

GetHetJobIdSet returns the HetJobIdSet field if non-nil, zero value otherwise.

### GetHetJobIdSetOk

`func (o *V0041JobInfoMsgInner) GetHetJobIdSetOk() (*string, bool)`

GetHetJobIdSetOk returns a tuple with the HetJobIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobIdSet

`func (o *V0041JobInfoMsgInner) SetHetJobIdSet(v string)`

SetHetJobIdSet sets HetJobIdSet field to given value.

### HasHetJobIdSet

`func (o *V0041JobInfoMsgInner) HasHetJobIdSet() bool`

HasHetJobIdSet returns a boolean if a field has been set.

### GetHetJobOffset

`func (o *V0041JobInfoMsgInner) GetHetJobOffset() V0041Uint32NoValStruct`

GetHetJobOffset returns the HetJobOffset field if non-nil, zero value otherwise.

### GetHetJobOffsetOk

`func (o *V0041JobInfoMsgInner) GetHetJobOffsetOk() (*V0041Uint32NoValStruct, bool)`

GetHetJobOffsetOk returns a tuple with the HetJobOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobOffset

`func (o *V0041JobInfoMsgInner) SetHetJobOffset(v V0041Uint32NoValStruct)`

SetHetJobOffset sets HetJobOffset field to given value.

### HasHetJobOffset

`func (o *V0041JobInfoMsgInner) HasHetJobOffset() bool`

HasHetJobOffset returns a boolean if a field has been set.

### GetJobId

`func (o *V0041JobInfoMsgInner) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0041JobInfoMsgInner) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0041JobInfoMsgInner) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0041JobInfoMsgInner) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobResources

`func (o *V0041JobInfoMsgInner) GetJobResources() V0041JobRes`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *V0041JobInfoMsgInner) GetJobResourcesOk() (*V0041JobRes, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *V0041JobInfoMsgInner) SetJobResources(v V0041JobRes)`

SetJobResources sets JobResources field to given value.

### HasJobResources

`func (o *V0041JobInfoMsgInner) HasJobResources() bool`

HasJobResources returns a boolean if a field has been set.

### GetJobSizeStr

`func (o *V0041JobInfoMsgInner) GetJobSizeStr() []string`

GetJobSizeStr returns the JobSizeStr field if non-nil, zero value otherwise.

### GetJobSizeStrOk

`func (o *V0041JobInfoMsgInner) GetJobSizeStrOk() (*[]string, bool)`

GetJobSizeStrOk returns a tuple with the JobSizeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSizeStr

`func (o *V0041JobInfoMsgInner) SetJobSizeStr(v []string)`

SetJobSizeStr sets JobSizeStr field to given value.

### HasJobSizeStr

`func (o *V0041JobInfoMsgInner) HasJobSizeStr() bool`

HasJobSizeStr returns a boolean if a field has been set.

### GetJobState

`func (o *V0041JobInfoMsgInner) GetJobState() []string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *V0041JobInfoMsgInner) GetJobStateOk() (*[]string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *V0041JobInfoMsgInner) SetJobState(v []string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *V0041JobInfoMsgInner) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetLastSchedEvaluation

`func (o *V0041JobInfoMsgInner) GetLastSchedEvaluation() V0041Uint64NoValStruct`

GetLastSchedEvaluation returns the LastSchedEvaluation field if non-nil, zero value otherwise.

### GetLastSchedEvaluationOk

`func (o *V0041JobInfoMsgInner) GetLastSchedEvaluationOk() (*V0041Uint64NoValStruct, bool)`

GetLastSchedEvaluationOk returns a tuple with the LastSchedEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedEvaluation

`func (o *V0041JobInfoMsgInner) SetLastSchedEvaluation(v V0041Uint64NoValStruct)`

SetLastSchedEvaluation sets LastSchedEvaluation field to given value.

### HasLastSchedEvaluation

`func (o *V0041JobInfoMsgInner) HasLastSchedEvaluation() bool`

HasLastSchedEvaluation returns a boolean if a field has been set.

### GetLicenses

`func (o *V0041JobInfoMsgInner) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0041JobInfoMsgInner) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0041JobInfoMsgInner) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0041JobInfoMsgInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMailType

`func (o *V0041JobInfoMsgInner) GetMailType() []string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *V0041JobInfoMsgInner) GetMailTypeOk() (*[]string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *V0041JobInfoMsgInner) SetMailType(v []string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *V0041JobInfoMsgInner) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *V0041JobInfoMsgInner) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *V0041JobInfoMsgInner) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *V0041JobInfoMsgInner) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *V0041JobInfoMsgInner) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMaxCpus

`func (o *V0041JobInfoMsgInner) GetMaxCpus() V0041Uint32NoValStruct`

GetMaxCpus returns the MaxCpus field if non-nil, zero value otherwise.

### GetMaxCpusOk

`func (o *V0041JobInfoMsgInner) GetMaxCpusOk() (*V0041Uint32NoValStruct, bool)`

GetMaxCpusOk returns a tuple with the MaxCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpus

`func (o *V0041JobInfoMsgInner) SetMaxCpus(v V0041Uint32NoValStruct)`

SetMaxCpus sets MaxCpus field to given value.

### HasMaxCpus

`func (o *V0041JobInfoMsgInner) HasMaxCpus() bool`

HasMaxCpus returns a boolean if a field has been set.

### GetMaxNodes

`func (o *V0041JobInfoMsgInner) GetMaxNodes() V0041Uint32NoValStruct`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *V0041JobInfoMsgInner) GetMaxNodesOk() (*V0041Uint32NoValStruct, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *V0041JobInfoMsgInner) SetMaxNodes(v V0041Uint32NoValStruct)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *V0041JobInfoMsgInner) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0041JobInfoMsgInner) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0041JobInfoMsgInner) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0041JobInfoMsgInner) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0041JobInfoMsgInner) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *V0041JobInfoMsgInner) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *V0041JobInfoMsgInner) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *V0041JobInfoMsgInner) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *V0041JobInfoMsgInner) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *V0041JobInfoMsgInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041JobInfoMsgInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041JobInfoMsgInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041JobInfoMsgInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *V0041JobInfoMsgInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *V0041JobInfoMsgInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *V0041JobInfoMsgInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *V0041JobInfoMsgInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodes

`func (o *V0041JobInfoMsgInner) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041JobInfoMsgInner) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041JobInfoMsgInner) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041JobInfoMsgInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNice

`func (o *V0041JobInfoMsgInner) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0041JobInfoMsgInner) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0041JobInfoMsgInner) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0041JobInfoMsgInner) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0041JobInfoMsgInner) GetTasksPerCore() V0041Uint16NoValStruct`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0041JobInfoMsgInner) GetTasksPerCoreOk() (*V0041Uint16NoValStruct, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0041JobInfoMsgInner) SetTasksPerCore(v V0041Uint16NoValStruct)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0041JobInfoMsgInner) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerTres

`func (o *V0041JobInfoMsgInner) GetTasksPerTres() V0041Uint16NoValStruct`

GetTasksPerTres returns the TasksPerTres field if non-nil, zero value otherwise.

### GetTasksPerTresOk

`func (o *V0041JobInfoMsgInner) GetTasksPerTresOk() (*V0041Uint16NoValStruct, bool)`

GetTasksPerTresOk returns a tuple with the TasksPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerTres

`func (o *V0041JobInfoMsgInner) SetTasksPerTres(v V0041Uint16NoValStruct)`

SetTasksPerTres sets TasksPerTres field to given value.

### HasTasksPerTres

`func (o *V0041JobInfoMsgInner) HasTasksPerTres() bool`

HasTasksPerTres returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *V0041JobInfoMsgInner) GetTasksPerNode() V0041Uint16NoValStruct`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *V0041JobInfoMsgInner) GetTasksPerNodeOk() (*V0041Uint16NoValStruct, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *V0041JobInfoMsgInner) SetTasksPerNode(v V0041Uint16NoValStruct)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *V0041JobInfoMsgInner) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0041JobInfoMsgInner) GetTasksPerSocket() V0041Uint16NoValStruct`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0041JobInfoMsgInner) GetTasksPerSocketOk() (*V0041Uint16NoValStruct, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0041JobInfoMsgInner) SetTasksPerSocket(v V0041Uint16NoValStruct)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0041JobInfoMsgInner) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *V0041JobInfoMsgInner) GetTasksPerBoard() V0041Uint16NoValStruct`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *V0041JobInfoMsgInner) GetTasksPerBoardOk() (*V0041Uint16NoValStruct, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *V0041JobInfoMsgInner) SetTasksPerBoard(v V0041Uint16NoValStruct)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *V0041JobInfoMsgInner) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetCpus

`func (o *V0041JobInfoMsgInner) GetCpus() V0041Uint32NoValStruct`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041JobInfoMsgInner) GetCpusOk() (*V0041Uint32NoValStruct, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041JobInfoMsgInner) SetCpus(v V0041Uint32NoValStruct)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0041JobInfoMsgInner) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0041JobInfoMsgInner) GetNodeCount() V0041Uint32NoValStruct`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0041JobInfoMsgInner) GetNodeCountOk() (*V0041Uint32NoValStruct, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0041JobInfoMsgInner) SetNodeCount(v V0041Uint32NoValStruct)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0041JobInfoMsgInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTasks

`func (o *V0041JobInfoMsgInner) GetTasks() V0041Uint32NoValStruct`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0041JobInfoMsgInner) GetTasksOk() (*V0041Uint32NoValStruct, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0041JobInfoMsgInner) SetTasks(v V0041Uint32NoValStruct)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0041JobInfoMsgInner) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPartition

`func (o *V0041JobInfoMsgInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041JobInfoMsgInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041JobInfoMsgInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041JobInfoMsgInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPrefer

`func (o *V0041JobInfoMsgInner) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *V0041JobInfoMsgInner) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *V0041JobInfoMsgInner) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *V0041JobInfoMsgInner) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0041JobInfoMsgInner) GetMemoryPerCpu() V0041Uint64NoValStruct`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0041JobInfoMsgInner) GetMemoryPerCpuOk() (*V0041Uint64NoValStruct, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0041JobInfoMsgInner) SetMemoryPerCpu(v V0041Uint64NoValStruct)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0041JobInfoMsgInner) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0041JobInfoMsgInner) GetMemoryPerNode() V0041Uint64NoValStruct`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0041JobInfoMsgInner) GetMemoryPerNodeOk() (*V0041Uint64NoValStruct, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0041JobInfoMsgInner) SetMemoryPerNode(v V0041Uint64NoValStruct)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0041JobInfoMsgInner) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0041JobInfoMsgInner) GetMinimumCpusPerNode() V0041Uint16NoValStruct`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0041JobInfoMsgInner) GetMinimumCpusPerNodeOk() (*V0041Uint16NoValStruct, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0041JobInfoMsgInner) SetMinimumCpusPerNode(v V0041Uint16NoValStruct)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0041JobInfoMsgInner) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMinimumTmpDiskPerNode

`func (o *V0041JobInfoMsgInner) GetMinimumTmpDiskPerNode() V0041Uint32NoValStruct`

GetMinimumTmpDiskPerNode returns the MinimumTmpDiskPerNode field if non-nil, zero value otherwise.

### GetMinimumTmpDiskPerNodeOk

`func (o *V0041JobInfoMsgInner) GetMinimumTmpDiskPerNodeOk() (*V0041Uint32NoValStruct, bool)`

GetMinimumTmpDiskPerNodeOk returns a tuple with the MinimumTmpDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTmpDiskPerNode

`func (o *V0041JobInfoMsgInner) SetMinimumTmpDiskPerNode(v V0041Uint32NoValStruct)`

SetMinimumTmpDiskPerNode sets MinimumTmpDiskPerNode field to given value.

### HasMinimumTmpDiskPerNode

`func (o *V0041JobInfoMsgInner) HasMinimumTmpDiskPerNode() bool`

HasMinimumTmpDiskPerNode returns a boolean if a field has been set.

### GetPower

`func (o *V0041JobInfoMsgInner) GetPower() V0041JobInfoMsgInnerPower`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *V0041JobInfoMsgInner) GetPowerOk() (*V0041JobInfoMsgInnerPower, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *V0041JobInfoMsgInner) SetPower(v V0041JobInfoMsgInnerPower)`

SetPower sets Power field to given value.

### HasPower

`func (o *V0041JobInfoMsgInner) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPreemptTime

`func (o *V0041JobInfoMsgInner) GetPreemptTime() V0041Uint64NoValStruct`

GetPreemptTime returns the PreemptTime field if non-nil, zero value otherwise.

### GetPreemptTimeOk

`func (o *V0041JobInfoMsgInner) GetPreemptTimeOk() (*V0041Uint64NoValStruct, bool)`

GetPreemptTimeOk returns a tuple with the PreemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptTime

`func (o *V0041JobInfoMsgInner) SetPreemptTime(v V0041Uint64NoValStruct)`

SetPreemptTime sets PreemptTime field to given value.

### HasPreemptTime

`func (o *V0041JobInfoMsgInner) HasPreemptTime() bool`

HasPreemptTime returns a boolean if a field has been set.

### GetPreemptableTime

`func (o *V0041JobInfoMsgInner) GetPreemptableTime() V0041Uint64NoValStruct`

GetPreemptableTime returns the PreemptableTime field if non-nil, zero value otherwise.

### GetPreemptableTimeOk

`func (o *V0041JobInfoMsgInner) GetPreemptableTimeOk() (*V0041Uint64NoValStruct, bool)`

GetPreemptableTimeOk returns a tuple with the PreemptableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptableTime

`func (o *V0041JobInfoMsgInner) SetPreemptableTime(v V0041Uint64NoValStruct)`

SetPreemptableTime sets PreemptableTime field to given value.

### HasPreemptableTime

`func (o *V0041JobInfoMsgInner) HasPreemptableTime() bool`

HasPreemptableTime returns a boolean if a field has been set.

### GetPreSusTime

`func (o *V0041JobInfoMsgInner) GetPreSusTime() V0041Uint64NoValStruct`

GetPreSusTime returns the PreSusTime field if non-nil, zero value otherwise.

### GetPreSusTimeOk

`func (o *V0041JobInfoMsgInner) GetPreSusTimeOk() (*V0041Uint64NoValStruct, bool)`

GetPreSusTimeOk returns a tuple with the PreSusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSusTime

`func (o *V0041JobInfoMsgInner) SetPreSusTime(v V0041Uint64NoValStruct)`

SetPreSusTime sets PreSusTime field to given value.

### HasPreSusTime

`func (o *V0041JobInfoMsgInner) HasPreSusTime() bool`

HasPreSusTime returns a boolean if a field has been set.

### GetPriority

`func (o *V0041JobInfoMsgInner) GetPriority() V0041Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041JobInfoMsgInner) GetPriorityOk() (*V0041Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041JobInfoMsgInner) SetPriority(v V0041Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041JobInfoMsgInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *V0041JobInfoMsgInner) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V0041JobInfoMsgInner) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V0041JobInfoMsgInner) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V0041JobInfoMsgInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *V0041JobInfoMsgInner) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041JobInfoMsgInner) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041JobInfoMsgInner) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041JobInfoMsgInner) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *V0041JobInfoMsgInner) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *V0041JobInfoMsgInner) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *V0041JobInfoMsgInner) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *V0041JobInfoMsgInner) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *V0041JobInfoMsgInner) GetRequiredNodes() string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *V0041JobInfoMsgInner) GetRequiredNodesOk() (*string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *V0041JobInfoMsgInner) SetRequiredNodes(v string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *V0041JobInfoMsgInner) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetMinimumSwitches

`func (o *V0041JobInfoMsgInner) GetMinimumSwitches() int32`

GetMinimumSwitches returns the MinimumSwitches field if non-nil, zero value otherwise.

### GetMinimumSwitchesOk

`func (o *V0041JobInfoMsgInner) GetMinimumSwitchesOk() (*int32, bool)`

GetMinimumSwitchesOk returns a tuple with the MinimumSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSwitches

`func (o *V0041JobInfoMsgInner) SetMinimumSwitches(v int32)`

SetMinimumSwitches sets MinimumSwitches field to given value.

### HasMinimumSwitches

`func (o *V0041JobInfoMsgInner) HasMinimumSwitches() bool`

HasMinimumSwitches returns a boolean if a field has been set.

### GetRequeue

`func (o *V0041JobInfoMsgInner) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0041JobInfoMsgInner) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0041JobInfoMsgInner) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0041JobInfoMsgInner) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetResizeTime

`func (o *V0041JobInfoMsgInner) GetResizeTime() V0041Uint64NoValStruct`

GetResizeTime returns the ResizeTime field if non-nil, zero value otherwise.

### GetResizeTimeOk

`func (o *V0041JobInfoMsgInner) GetResizeTimeOk() (*V0041Uint64NoValStruct, bool)`

GetResizeTimeOk returns a tuple with the ResizeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeTime

`func (o *V0041JobInfoMsgInner) SetResizeTime(v V0041Uint64NoValStruct)`

SetResizeTime sets ResizeTime field to given value.

### HasResizeTime

`func (o *V0041JobInfoMsgInner) HasResizeTime() bool`

HasResizeTime returns a boolean if a field has been set.

### GetRestartCnt

`func (o *V0041JobInfoMsgInner) GetRestartCnt() int32`

GetRestartCnt returns the RestartCnt field if non-nil, zero value otherwise.

### GetRestartCntOk

`func (o *V0041JobInfoMsgInner) GetRestartCntOk() (*int32, bool)`

GetRestartCntOk returns a tuple with the RestartCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCnt

`func (o *V0041JobInfoMsgInner) SetRestartCnt(v int32)`

SetRestartCnt sets RestartCnt field to given value.

### HasRestartCnt

`func (o *V0041JobInfoMsgInner) HasRestartCnt() bool`

HasRestartCnt returns a boolean if a field has been set.

### GetResvName

`func (o *V0041JobInfoMsgInner) GetResvName() string`

GetResvName returns the ResvName field if non-nil, zero value otherwise.

### GetResvNameOk

`func (o *V0041JobInfoMsgInner) GetResvNameOk() (*string, bool)`

GetResvNameOk returns a tuple with the ResvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResvName

`func (o *V0041JobInfoMsgInner) SetResvName(v string)`

SetResvName sets ResvName field to given value.

### HasResvName

`func (o *V0041JobInfoMsgInner) HasResvName() bool`

HasResvName returns a boolean if a field has been set.

### GetScheduledNodes

`func (o *V0041JobInfoMsgInner) GetScheduledNodes() string`

GetScheduledNodes returns the ScheduledNodes field if non-nil, zero value otherwise.

### GetScheduledNodesOk

`func (o *V0041JobInfoMsgInner) GetScheduledNodesOk() (*string, bool)`

GetScheduledNodesOk returns a tuple with the ScheduledNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledNodes

`func (o *V0041JobInfoMsgInner) SetScheduledNodes(v string)`

SetScheduledNodes sets ScheduledNodes field to given value.

### HasScheduledNodes

`func (o *V0041JobInfoMsgInner) HasScheduledNodes() bool`

HasScheduledNodes returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *V0041JobInfoMsgInner) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *V0041JobInfoMsgInner) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *V0041JobInfoMsgInner) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *V0041JobInfoMsgInner) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetShared

`func (o *V0041JobInfoMsgInner) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *V0041JobInfoMsgInner) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *V0041JobInfoMsgInner) SetShared(v []string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *V0041JobInfoMsgInner) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetExclusive

`func (o *V0041JobInfoMsgInner) GetExclusive() []string`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *V0041JobInfoMsgInner) GetExclusiveOk() (*[]string, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *V0041JobInfoMsgInner) SetExclusive(v []string)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *V0041JobInfoMsgInner) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0041JobInfoMsgInner) GetOversubscribe() bool`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0041JobInfoMsgInner) GetOversubscribeOk() (*bool, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0041JobInfoMsgInner) SetOversubscribe(v bool)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0041JobInfoMsgInner) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetShowFlags

`func (o *V0041JobInfoMsgInner) GetShowFlags() []string`

GetShowFlags returns the ShowFlags field if non-nil, zero value otherwise.

### GetShowFlagsOk

`func (o *V0041JobInfoMsgInner) GetShowFlagsOk() (*[]string, bool)`

GetShowFlagsOk returns a tuple with the ShowFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFlags

`func (o *V0041JobInfoMsgInner) SetShowFlags(v []string)`

SetShowFlags sets ShowFlags field to given value.

### HasShowFlags

`func (o *V0041JobInfoMsgInner) HasShowFlags() bool`

HasShowFlags returns a boolean if a field has been set.

### GetSocketsPerBoard

`func (o *V0041JobInfoMsgInner) GetSocketsPerBoard() int32`

GetSocketsPerBoard returns the SocketsPerBoard field if non-nil, zero value otherwise.

### GetSocketsPerBoardOk

`func (o *V0041JobInfoMsgInner) GetSocketsPerBoardOk() (*int32, bool)`

GetSocketsPerBoardOk returns a tuple with the SocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerBoard

`func (o *V0041JobInfoMsgInner) SetSocketsPerBoard(v int32)`

SetSocketsPerBoard sets SocketsPerBoard field to given value.

### HasSocketsPerBoard

`func (o *V0041JobInfoMsgInner) HasSocketsPerBoard() bool`

HasSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0041JobInfoMsgInner) GetSocketsPerNode() V0041Uint16NoValStruct`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0041JobInfoMsgInner) GetSocketsPerNodeOk() (*V0041Uint16NoValStruct, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0041JobInfoMsgInner) SetSocketsPerNode(v V0041Uint16NoValStruct)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0041JobInfoMsgInner) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetStartTime

`func (o *V0041JobInfoMsgInner) GetStartTime() V0041Uint64NoValStruct`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0041JobInfoMsgInner) GetStartTimeOk() (*V0041Uint64NoValStruct, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0041JobInfoMsgInner) SetStartTime(v V0041Uint64NoValStruct)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0041JobInfoMsgInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStateDescription

`func (o *V0041JobInfoMsgInner) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *V0041JobInfoMsgInner) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *V0041JobInfoMsgInner) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *V0041JobInfoMsgInner) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetStateReason

`func (o *V0041JobInfoMsgInner) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *V0041JobInfoMsgInner) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *V0041JobInfoMsgInner) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *V0041JobInfoMsgInner) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStandardError

`func (o *V0041JobInfoMsgInner) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0041JobInfoMsgInner) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0041JobInfoMsgInner) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0041JobInfoMsgInner) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0041JobInfoMsgInner) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0041JobInfoMsgInner) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0041JobInfoMsgInner) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0041JobInfoMsgInner) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0041JobInfoMsgInner) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0041JobInfoMsgInner) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0041JobInfoMsgInner) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0041JobInfoMsgInner) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetSubmitTime

`func (o *V0041JobInfoMsgInner) GetSubmitTime() V0041Uint64NoValStruct`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *V0041JobInfoMsgInner) GetSubmitTimeOk() (*V0041Uint64NoValStruct, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *V0041JobInfoMsgInner) SetSubmitTime(v V0041Uint64NoValStruct)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *V0041JobInfoMsgInner) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0041JobInfoMsgInner) GetSuspendTime() V0041Uint64NoValStruct`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0041JobInfoMsgInner) GetSuspendTimeOk() (*V0041Uint64NoValStruct, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0041JobInfoMsgInner) SetSuspendTime(v V0041Uint64NoValStruct)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0041JobInfoMsgInner) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.

### GetSystemComment

`func (o *V0041JobInfoMsgInner) GetSystemComment() string`

GetSystemComment returns the SystemComment field if non-nil, zero value otherwise.

### GetSystemCommentOk

`func (o *V0041JobInfoMsgInner) GetSystemCommentOk() (*string, bool)`

GetSystemCommentOk returns a tuple with the SystemComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemComment

`func (o *V0041JobInfoMsgInner) SetSystemComment(v string)`

SetSystemComment sets SystemComment field to given value.

### HasSystemComment

`func (o *V0041JobInfoMsgInner) HasSystemComment() bool`

HasSystemComment returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0041JobInfoMsgInner) GetTimeLimit() V0041Uint32NoValStruct`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0041JobInfoMsgInner) GetTimeLimitOk() (*V0041Uint32NoValStruct, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0041JobInfoMsgInner) SetTimeLimit(v V0041Uint32NoValStruct)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0041JobInfoMsgInner) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0041JobInfoMsgInner) GetTimeMinimum() V0041Uint32NoValStruct`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0041JobInfoMsgInner) GetTimeMinimumOk() (*V0041Uint32NoValStruct, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0041JobInfoMsgInner) SetTimeMinimum(v V0041Uint32NoValStruct)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0041JobInfoMsgInner) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0041JobInfoMsgInner) GetThreadsPerCore() V0041Uint16NoValStruct`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0041JobInfoMsgInner) GetThreadsPerCoreOk() (*V0041Uint16NoValStruct, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0041JobInfoMsgInner) SetThreadsPerCore(v V0041Uint16NoValStruct)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0041JobInfoMsgInner) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTresBind

`func (o *V0041JobInfoMsgInner) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *V0041JobInfoMsgInner) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *V0041JobInfoMsgInner) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *V0041JobInfoMsgInner) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *V0041JobInfoMsgInner) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *V0041JobInfoMsgInner) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *V0041JobInfoMsgInner) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *V0041JobInfoMsgInner) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *V0041JobInfoMsgInner) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *V0041JobInfoMsgInner) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *V0041JobInfoMsgInner) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *V0041JobInfoMsgInner) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *V0041JobInfoMsgInner) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *V0041JobInfoMsgInner) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *V0041JobInfoMsgInner) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *V0041JobInfoMsgInner) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *V0041JobInfoMsgInner) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *V0041JobInfoMsgInner) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *V0041JobInfoMsgInner) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *V0041JobInfoMsgInner) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *V0041JobInfoMsgInner) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *V0041JobInfoMsgInner) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *V0041JobInfoMsgInner) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *V0041JobInfoMsgInner) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetTresReqStr

`func (o *V0041JobInfoMsgInner) GetTresReqStr() string`

GetTresReqStr returns the TresReqStr field if non-nil, zero value otherwise.

### GetTresReqStrOk

`func (o *V0041JobInfoMsgInner) GetTresReqStrOk() (*string, bool)`

GetTresReqStrOk returns a tuple with the TresReqStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresReqStr

`func (o *V0041JobInfoMsgInner) SetTresReqStr(v string)`

SetTresReqStr sets TresReqStr field to given value.

### HasTresReqStr

`func (o *V0041JobInfoMsgInner) HasTresReqStr() bool`

HasTresReqStr returns a boolean if a field has been set.

### GetTresAllocStr

`func (o *V0041JobInfoMsgInner) GetTresAllocStr() string`

GetTresAllocStr returns the TresAllocStr field if non-nil, zero value otherwise.

### GetTresAllocStrOk

`func (o *V0041JobInfoMsgInner) GetTresAllocStrOk() (*string, bool)`

GetTresAllocStrOk returns a tuple with the TresAllocStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresAllocStr

`func (o *V0041JobInfoMsgInner) SetTresAllocStr(v string)`

SetTresAllocStr sets TresAllocStr field to given value.

### HasTresAllocStr

`func (o *V0041JobInfoMsgInner) HasTresAllocStr() bool`

HasTresAllocStr returns a boolean if a field has been set.

### GetUserId

`func (o *V0041JobInfoMsgInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0041JobInfoMsgInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0041JobInfoMsgInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0041JobInfoMsgInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *V0041JobInfoMsgInner) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *V0041JobInfoMsgInner) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *V0041JobInfoMsgInner) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *V0041JobInfoMsgInner) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetMaximumSwitchWaitTime

`func (o *V0041JobInfoMsgInner) GetMaximumSwitchWaitTime() int32`

GetMaximumSwitchWaitTime returns the MaximumSwitchWaitTime field if non-nil, zero value otherwise.

### GetMaximumSwitchWaitTimeOk

`func (o *V0041JobInfoMsgInner) GetMaximumSwitchWaitTimeOk() (*int32, bool)`

GetMaximumSwitchWaitTimeOk returns a tuple with the MaximumSwitchWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSwitchWaitTime

`func (o *V0041JobInfoMsgInner) SetMaximumSwitchWaitTime(v int32)`

SetMaximumSwitchWaitTime sets MaximumSwitchWaitTime field to given value.

### HasMaximumSwitchWaitTime

`func (o *V0041JobInfoMsgInner) HasMaximumSwitchWaitTime() bool`

HasMaximumSwitchWaitTime returns a boolean if a field has been set.

### GetWckey

`func (o *V0041JobInfoMsgInner) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0041JobInfoMsgInner) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0041JobInfoMsgInner) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0041JobInfoMsgInner) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0041JobInfoMsgInner) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0041JobInfoMsgInner) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0041JobInfoMsgInner) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0041JobInfoMsgInner) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


