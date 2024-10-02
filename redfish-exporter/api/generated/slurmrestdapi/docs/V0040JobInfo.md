# V0040JobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AccrueTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**AdminComment** | Pointer to **string** |  | [optional] 
**AllocatingNode** | Pointer to **string** |  | [optional] 
**ArrayJobId** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**ArrayTaskId** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**ArrayMaxTasks** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
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
**CoresPerSocket** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**BillableTres** | Pointer to [**V0040Float64NoVal**](V0040Float64NoVal.md) |  | [optional] 
**CpusPerTask** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**CpuFrequencyMinimum** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**CpuFrequencyMaximum** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**CpuFrequencyGovernor** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**CpusPerTres** | Pointer to **string** |  | [optional] 
**Cron** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**DelayBoot** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Dependency** | Pointer to **string** |  | [optional] 
**DerivedExitCode** | Pointer to [**V0040ProcessExitCodeVerbose**](V0040ProcessExitCodeVerbose.md) |  | [optional] 
**EligibleTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**EndTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**ExcludedNodes** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to [**V0040ProcessExitCodeVerbose**](V0040ProcessExitCodeVerbose.md) |  | [optional] 
**Extra** | Pointer to **string** |  | [optional] 
**FailedNode** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**FederationOrigin** | Pointer to **string** |  | [optional] 
**FederationSiblingsActive** | Pointer to **string** |  | [optional] 
**FederationSiblingsViable** | Pointer to **string** |  | [optional] 
**GresDetail** | Pointer to **[]string** |  | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**HetJobId** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**HetJobIdSet** | Pointer to **string** |  | [optional] 
**HetJobOffset** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**JobId** | Pointer to **int32** |  | [optional] 
**JobResources** | Pointer to [**V0040JobRes**](V0040JobRes.md) |  | [optional] 
**JobSizeStr** | Pointer to **[]string** |  | [optional] 
**JobState** | Pointer to **[]string** |  | [optional] 
**LastSchedEvaluation** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Licenses** | Pointer to **string** |  | [optional] 
**MailType** | Pointer to **[]string** |  | [optional] 
**MailUser** | Pointer to **string** |  | [optional] 
**MaxCpus** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**MaxNodes** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**McsLabel** | Pointer to **string** |  | [optional] 
**MemoryPerTres** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**Nice** | Pointer to **int32** |  | [optional] 
**TasksPerCore** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**TasksPerTres** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**TasksPerNode** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**TasksPerSocket** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**TasksPerBoard** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**Cpus** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**NodeCount** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Tasks** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Prefer** | Pointer to **string** |  | [optional] 
**MemoryPerCpu** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**MinimumCpusPerNode** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**MinimumTmpDiskPerNode** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Power** | Pointer to [**V0041JobInfoMsgInnerPower**](V0041JobInfoMsgInnerPower.md) |  | [optional] 
**PreemptTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**PreemptableTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**PreSusTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Priority** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Profile** | Pointer to **[]string** |  | [optional] 
**Qos** | Pointer to **string** |  | [optional] 
**Reboot** | Pointer to **bool** |  | [optional] 
**RequiredNodes** | Pointer to **string** |  | [optional] 
**MinimumSwitches** | Pointer to **int32** |  | [optional] 
**Requeue** | Pointer to **bool** |  | [optional] 
**ResizeTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**RestartCnt** | Pointer to **int32** |  | [optional] 
**ResvName** | Pointer to **string** |  | [optional] 
**ScheduledNodes** | Pointer to **string** |  | [optional] 
**SelinuxContext** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **[]string** |  | [optional] 
**Exclusive** | Pointer to **[]string** |  | [optional] 
**Oversubscribe** | Pointer to **bool** |  | [optional] 
**ShowFlags** | Pointer to **[]string** |  | [optional] 
**SocketsPerBoard** | Pointer to **int32** |  | [optional] 
**SocketsPerNode** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**StartTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**StateDescription** | Pointer to **string** |  | [optional] 
**StateReason** | Pointer to **string** |  | [optional] 
**StandardError** | Pointer to **string** |  | [optional] 
**StandardInput** | Pointer to **string** |  | [optional] 
**StandardOutput** | Pointer to **string** |  | [optional] 
**SubmitTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**SystemComment** | Pointer to **string** |  | [optional] 
**TimeLimit** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**TimeMinimum** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**ThreadsPerCore** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
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

### NewV0040JobInfo

`func NewV0040JobInfo() *V0040JobInfo`

NewV0040JobInfo instantiates a new V0040JobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobInfoWithDefaults

`func NewV0040JobInfoWithDefaults() *V0040JobInfo`

NewV0040JobInfoWithDefaults instantiates a new V0040JobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0040JobInfo) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0040JobInfo) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0040JobInfo) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0040JobInfo) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccrueTime

`func (o *V0040JobInfo) GetAccrueTime() V0040Uint64NoVal`

GetAccrueTime returns the AccrueTime field if non-nil, zero value otherwise.

### GetAccrueTimeOk

`func (o *V0040JobInfo) GetAccrueTimeOk() (*V0040Uint64NoVal, bool)`

GetAccrueTimeOk returns a tuple with the AccrueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrueTime

`func (o *V0040JobInfo) SetAccrueTime(v V0040Uint64NoVal)`

SetAccrueTime sets AccrueTime field to given value.

### HasAccrueTime

`func (o *V0040JobInfo) HasAccrueTime() bool`

HasAccrueTime returns a boolean if a field has been set.

### GetAdminComment

`func (o *V0040JobInfo) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *V0040JobInfo) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *V0040JobInfo) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *V0040JobInfo) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAllocatingNode

`func (o *V0040JobInfo) GetAllocatingNode() string`

GetAllocatingNode returns the AllocatingNode field if non-nil, zero value otherwise.

### GetAllocatingNodeOk

`func (o *V0040JobInfo) GetAllocatingNodeOk() (*string, bool)`

GetAllocatingNodeOk returns a tuple with the AllocatingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatingNode

`func (o *V0040JobInfo) SetAllocatingNode(v string)`

SetAllocatingNode sets AllocatingNode field to given value.

### HasAllocatingNode

`func (o *V0040JobInfo) HasAllocatingNode() bool`

HasAllocatingNode returns a boolean if a field has been set.

### GetArrayJobId

`func (o *V0040JobInfo) GetArrayJobId() V0040Uint32NoVal`

GetArrayJobId returns the ArrayJobId field if non-nil, zero value otherwise.

### GetArrayJobIdOk

`func (o *V0040JobInfo) GetArrayJobIdOk() (*V0040Uint32NoVal, bool)`

GetArrayJobIdOk returns a tuple with the ArrayJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayJobId

`func (o *V0040JobInfo) SetArrayJobId(v V0040Uint32NoVal)`

SetArrayJobId sets ArrayJobId field to given value.

### HasArrayJobId

`func (o *V0040JobInfo) HasArrayJobId() bool`

HasArrayJobId returns a boolean if a field has been set.

### GetArrayTaskId

`func (o *V0040JobInfo) GetArrayTaskId() V0040Uint32NoVal`

GetArrayTaskId returns the ArrayTaskId field if non-nil, zero value otherwise.

### GetArrayTaskIdOk

`func (o *V0040JobInfo) GetArrayTaskIdOk() (*V0040Uint32NoVal, bool)`

GetArrayTaskIdOk returns a tuple with the ArrayTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskId

`func (o *V0040JobInfo) SetArrayTaskId(v V0040Uint32NoVal)`

SetArrayTaskId sets ArrayTaskId field to given value.

### HasArrayTaskId

`func (o *V0040JobInfo) HasArrayTaskId() bool`

HasArrayTaskId returns a boolean if a field has been set.

### GetArrayMaxTasks

`func (o *V0040JobInfo) GetArrayMaxTasks() V0040Uint32NoVal`

GetArrayMaxTasks returns the ArrayMaxTasks field if non-nil, zero value otherwise.

### GetArrayMaxTasksOk

`func (o *V0040JobInfo) GetArrayMaxTasksOk() (*V0040Uint32NoVal, bool)`

GetArrayMaxTasksOk returns a tuple with the ArrayMaxTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayMaxTasks

`func (o *V0040JobInfo) SetArrayMaxTasks(v V0040Uint32NoVal)`

SetArrayMaxTasks sets ArrayMaxTasks field to given value.

### HasArrayMaxTasks

`func (o *V0040JobInfo) HasArrayMaxTasks() bool`

HasArrayMaxTasks returns a boolean if a field has been set.

### GetArrayTaskString

`func (o *V0040JobInfo) GetArrayTaskString() string`

GetArrayTaskString returns the ArrayTaskString field if non-nil, zero value otherwise.

### GetArrayTaskStringOk

`func (o *V0040JobInfo) GetArrayTaskStringOk() (*string, bool)`

GetArrayTaskStringOk returns a tuple with the ArrayTaskString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskString

`func (o *V0040JobInfo) SetArrayTaskString(v string)`

SetArrayTaskString sets ArrayTaskString field to given value.

### HasArrayTaskString

`func (o *V0040JobInfo) HasArrayTaskString() bool`

HasArrayTaskString returns a boolean if a field has been set.

### GetAssociationId

`func (o *V0040JobInfo) GetAssociationId() int32`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *V0040JobInfo) GetAssociationIdOk() (*int32, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *V0040JobInfo) SetAssociationId(v int32)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *V0040JobInfo) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0040JobInfo) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0040JobInfo) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0040JobInfo) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0040JobInfo) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBatchFlag

`func (o *V0040JobInfo) GetBatchFlag() bool`

GetBatchFlag returns the BatchFlag field if non-nil, zero value otherwise.

### GetBatchFlagOk

`func (o *V0040JobInfo) GetBatchFlagOk() (*bool, bool)`

GetBatchFlagOk returns a tuple with the BatchFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFlag

`func (o *V0040JobInfo) SetBatchFlag(v bool)`

SetBatchFlag sets BatchFlag field to given value.

### HasBatchFlag

`func (o *V0040JobInfo) HasBatchFlag() bool`

HasBatchFlag returns a boolean if a field has been set.

### GetBatchHost

`func (o *V0040JobInfo) GetBatchHost() string`

GetBatchHost returns the BatchHost field if non-nil, zero value otherwise.

### GetBatchHostOk

`func (o *V0040JobInfo) GetBatchHostOk() (*string, bool)`

GetBatchHostOk returns a tuple with the BatchHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchHost

`func (o *V0040JobInfo) SetBatchHost(v string)`

SetBatchHost sets BatchHost field to given value.

### HasBatchHost

`func (o *V0040JobInfo) HasBatchHost() bool`

HasBatchHost returns a boolean if a field has been set.

### GetFlags

`func (o *V0040JobInfo) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0040JobInfo) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0040JobInfo) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0040JobInfo) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0040JobInfo) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0040JobInfo) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0040JobInfo) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0040JobInfo) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetBurstBufferState

`func (o *V0040JobInfo) GetBurstBufferState() string`

GetBurstBufferState returns the BurstBufferState field if non-nil, zero value otherwise.

### GetBurstBufferStateOk

`func (o *V0040JobInfo) GetBurstBufferStateOk() (*string, bool)`

GetBurstBufferStateOk returns a tuple with the BurstBufferState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBufferState

`func (o *V0040JobInfo) SetBurstBufferState(v string)`

SetBurstBufferState sets BurstBufferState field to given value.

### HasBurstBufferState

`func (o *V0040JobInfo) HasBurstBufferState() bool`

HasBurstBufferState returns a boolean if a field has been set.

### GetCluster

`func (o *V0040JobInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0040JobInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0040JobInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0040JobInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterFeatures

`func (o *V0040JobInfo) GetClusterFeatures() string`

GetClusterFeatures returns the ClusterFeatures field if non-nil, zero value otherwise.

### GetClusterFeaturesOk

`func (o *V0040JobInfo) GetClusterFeaturesOk() (*string, bool)`

GetClusterFeaturesOk returns a tuple with the ClusterFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFeatures

`func (o *V0040JobInfo) SetClusterFeatures(v string)`

SetClusterFeatures sets ClusterFeatures field to given value.

### HasClusterFeatures

`func (o *V0040JobInfo) HasClusterFeatures() bool`

HasClusterFeatures returns a boolean if a field has been set.

### GetCommand

`func (o *V0040JobInfo) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0040JobInfo) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0040JobInfo) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0040JobInfo) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetComment

`func (o *V0040JobInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0040JobInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0040JobInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0040JobInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContainer

`func (o *V0040JobInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0040JobInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0040JobInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0040JobInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *V0040JobInfo) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *V0040JobInfo) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *V0040JobInfo) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *V0040JobInfo) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetContiguous

`func (o *V0040JobInfo) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *V0040JobInfo) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *V0040JobInfo) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *V0040JobInfo) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetCoreSpec

`func (o *V0040JobInfo) GetCoreSpec() int32`

GetCoreSpec returns the CoreSpec field if non-nil, zero value otherwise.

### GetCoreSpecOk

`func (o *V0040JobInfo) GetCoreSpecOk() (*int32, bool)`

GetCoreSpecOk returns a tuple with the CoreSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpec

`func (o *V0040JobInfo) SetCoreSpec(v int32)`

SetCoreSpec sets CoreSpec field to given value.

### HasCoreSpec

`func (o *V0040JobInfo) HasCoreSpec() bool`

HasCoreSpec returns a boolean if a field has been set.

### GetThreadSpec

`func (o *V0040JobInfo) GetThreadSpec() int32`

GetThreadSpec returns the ThreadSpec field if non-nil, zero value otherwise.

### GetThreadSpecOk

`func (o *V0040JobInfo) GetThreadSpecOk() (*int32, bool)`

GetThreadSpecOk returns a tuple with the ThreadSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpec

`func (o *V0040JobInfo) SetThreadSpec(v int32)`

SetThreadSpec sets ThreadSpec field to given value.

### HasThreadSpec

`func (o *V0040JobInfo) HasThreadSpec() bool`

HasThreadSpec returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *V0040JobInfo) GetCoresPerSocket() V0040Uint16NoVal`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *V0040JobInfo) GetCoresPerSocketOk() (*V0040Uint16NoVal, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *V0040JobInfo) SetCoresPerSocket(v V0040Uint16NoVal)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *V0040JobInfo) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetBillableTres

`func (o *V0040JobInfo) GetBillableTres() V0040Float64NoVal`

GetBillableTres returns the BillableTres field if non-nil, zero value otherwise.

### GetBillableTresOk

`func (o *V0040JobInfo) GetBillableTresOk() (*V0040Float64NoVal, bool)`

GetBillableTresOk returns a tuple with the BillableTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTres

`func (o *V0040JobInfo) SetBillableTres(v V0040Float64NoVal)`

SetBillableTres sets BillableTres field to given value.

### HasBillableTres

`func (o *V0040JobInfo) HasBillableTres() bool`

HasBillableTres returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0040JobInfo) GetCpusPerTask() V0040Uint16NoVal`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0040JobInfo) GetCpusPerTaskOk() (*V0040Uint16NoVal, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0040JobInfo) SetCpusPerTask(v V0040Uint16NoVal)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0040JobInfo) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetCpuFrequencyMinimum

`func (o *V0040JobInfo) GetCpuFrequencyMinimum() V0040Uint32NoVal`

GetCpuFrequencyMinimum returns the CpuFrequencyMinimum field if non-nil, zero value otherwise.

### GetCpuFrequencyMinimumOk

`func (o *V0040JobInfo) GetCpuFrequencyMinimumOk() (*V0040Uint32NoVal, bool)`

GetCpuFrequencyMinimumOk returns a tuple with the CpuFrequencyMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMinimum

`func (o *V0040JobInfo) SetCpuFrequencyMinimum(v V0040Uint32NoVal)`

SetCpuFrequencyMinimum sets CpuFrequencyMinimum field to given value.

### HasCpuFrequencyMinimum

`func (o *V0040JobInfo) HasCpuFrequencyMinimum() bool`

HasCpuFrequencyMinimum returns a boolean if a field has been set.

### GetCpuFrequencyMaximum

`func (o *V0040JobInfo) GetCpuFrequencyMaximum() V0040Uint32NoVal`

GetCpuFrequencyMaximum returns the CpuFrequencyMaximum field if non-nil, zero value otherwise.

### GetCpuFrequencyMaximumOk

`func (o *V0040JobInfo) GetCpuFrequencyMaximumOk() (*V0040Uint32NoVal, bool)`

GetCpuFrequencyMaximumOk returns a tuple with the CpuFrequencyMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMaximum

`func (o *V0040JobInfo) SetCpuFrequencyMaximum(v V0040Uint32NoVal)`

SetCpuFrequencyMaximum sets CpuFrequencyMaximum field to given value.

### HasCpuFrequencyMaximum

`func (o *V0040JobInfo) HasCpuFrequencyMaximum() bool`

HasCpuFrequencyMaximum returns a boolean if a field has been set.

### GetCpuFrequencyGovernor

`func (o *V0040JobInfo) GetCpuFrequencyGovernor() V0040Uint32NoVal`

GetCpuFrequencyGovernor returns the CpuFrequencyGovernor field if non-nil, zero value otherwise.

### GetCpuFrequencyGovernorOk

`func (o *V0040JobInfo) GetCpuFrequencyGovernorOk() (*V0040Uint32NoVal, bool)`

GetCpuFrequencyGovernorOk returns a tuple with the CpuFrequencyGovernor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyGovernor

`func (o *V0040JobInfo) SetCpuFrequencyGovernor(v V0040Uint32NoVal)`

SetCpuFrequencyGovernor sets CpuFrequencyGovernor field to given value.

### HasCpuFrequencyGovernor

`func (o *V0040JobInfo) HasCpuFrequencyGovernor() bool`

HasCpuFrequencyGovernor returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *V0040JobInfo) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *V0040JobInfo) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *V0040JobInfo) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *V0040JobInfo) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetCron

`func (o *V0040JobInfo) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *V0040JobInfo) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *V0040JobInfo) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *V0040JobInfo) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDeadline

`func (o *V0040JobInfo) GetDeadline() V0040Uint64NoVal`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0040JobInfo) GetDeadlineOk() (*V0040Uint64NoVal, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0040JobInfo) SetDeadline(v V0040Uint64NoVal)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0040JobInfo) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0040JobInfo) GetDelayBoot() V0040Uint32NoVal`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0040JobInfo) GetDelayBootOk() (*V0040Uint32NoVal, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0040JobInfo) SetDelayBoot(v V0040Uint32NoVal)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0040JobInfo) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0040JobInfo) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0040JobInfo) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0040JobInfo) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0040JobInfo) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *V0040JobInfo) GetDerivedExitCode() V0040ProcessExitCodeVerbose`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *V0040JobInfo) GetDerivedExitCodeOk() (*V0040ProcessExitCodeVerbose, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *V0040JobInfo) SetDerivedExitCode(v V0040ProcessExitCodeVerbose)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *V0040JobInfo) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetEligibleTime

`func (o *V0040JobInfo) GetEligibleTime() V0040Uint64NoVal`

GetEligibleTime returns the EligibleTime field if non-nil, zero value otherwise.

### GetEligibleTimeOk

`func (o *V0040JobInfo) GetEligibleTimeOk() (*V0040Uint64NoVal, bool)`

GetEligibleTimeOk returns a tuple with the EligibleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleTime

`func (o *V0040JobInfo) SetEligibleTime(v V0040Uint64NoVal)`

SetEligibleTime sets EligibleTime field to given value.

### HasEligibleTime

`func (o *V0040JobInfo) HasEligibleTime() bool`

HasEligibleTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V0040JobInfo) GetEndTime() V0040Uint64NoVal`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0040JobInfo) GetEndTimeOk() (*V0040Uint64NoVal, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0040JobInfo) SetEndTime(v V0040Uint64NoVal)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0040JobInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *V0040JobInfo) GetExcludedNodes() string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *V0040JobInfo) GetExcludedNodesOk() (*string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *V0040JobInfo) SetExcludedNodes(v string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *V0040JobInfo) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExitCode

`func (o *V0040JobInfo) GetExitCode() V0040ProcessExitCodeVerbose`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0040JobInfo) GetExitCodeOk() (*V0040ProcessExitCodeVerbose, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0040JobInfo) SetExitCode(v V0040ProcessExitCodeVerbose)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0040JobInfo) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExtra

`func (o *V0040JobInfo) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0040JobInfo) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0040JobInfo) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0040JobInfo) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFailedNode

`func (o *V0040JobInfo) GetFailedNode() string`

GetFailedNode returns the FailedNode field if non-nil, zero value otherwise.

### GetFailedNodeOk

`func (o *V0040JobInfo) GetFailedNodeOk() (*string, bool)`

GetFailedNodeOk returns a tuple with the FailedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNode

`func (o *V0040JobInfo) SetFailedNode(v string)`

SetFailedNode sets FailedNode field to given value.

### HasFailedNode

`func (o *V0040JobInfo) HasFailedNode() bool`

HasFailedNode returns a boolean if a field has been set.

### GetFeatures

`func (o *V0040JobInfo) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0040JobInfo) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0040JobInfo) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0040JobInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFederationOrigin

`func (o *V0040JobInfo) GetFederationOrigin() string`

GetFederationOrigin returns the FederationOrigin field if non-nil, zero value otherwise.

### GetFederationOriginOk

`func (o *V0040JobInfo) GetFederationOriginOk() (*string, bool)`

GetFederationOriginOk returns a tuple with the FederationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationOrigin

`func (o *V0040JobInfo) SetFederationOrigin(v string)`

SetFederationOrigin sets FederationOrigin field to given value.

### HasFederationOrigin

`func (o *V0040JobInfo) HasFederationOrigin() bool`

HasFederationOrigin returns a boolean if a field has been set.

### GetFederationSiblingsActive

`func (o *V0040JobInfo) GetFederationSiblingsActive() string`

GetFederationSiblingsActive returns the FederationSiblingsActive field if non-nil, zero value otherwise.

### GetFederationSiblingsActiveOk

`func (o *V0040JobInfo) GetFederationSiblingsActiveOk() (*string, bool)`

GetFederationSiblingsActiveOk returns a tuple with the FederationSiblingsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsActive

`func (o *V0040JobInfo) SetFederationSiblingsActive(v string)`

SetFederationSiblingsActive sets FederationSiblingsActive field to given value.

### HasFederationSiblingsActive

`func (o *V0040JobInfo) HasFederationSiblingsActive() bool`

HasFederationSiblingsActive returns a boolean if a field has been set.

### GetFederationSiblingsViable

`func (o *V0040JobInfo) GetFederationSiblingsViable() string`

GetFederationSiblingsViable returns the FederationSiblingsViable field if non-nil, zero value otherwise.

### GetFederationSiblingsViableOk

`func (o *V0040JobInfo) GetFederationSiblingsViableOk() (*string, bool)`

GetFederationSiblingsViableOk returns a tuple with the FederationSiblingsViable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsViable

`func (o *V0040JobInfo) SetFederationSiblingsViable(v string)`

SetFederationSiblingsViable sets FederationSiblingsViable field to given value.

### HasFederationSiblingsViable

`func (o *V0040JobInfo) HasFederationSiblingsViable() bool`

HasFederationSiblingsViable returns a boolean if a field has been set.

### GetGresDetail

`func (o *V0040JobInfo) GetGresDetail() []string`

GetGresDetail returns the GresDetail field if non-nil, zero value otherwise.

### GetGresDetailOk

`func (o *V0040JobInfo) GetGresDetailOk() (*[]string, bool)`

GetGresDetailOk returns a tuple with the GresDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDetail

`func (o *V0040JobInfo) SetGresDetail(v []string)`

SetGresDetail sets GresDetail field to given value.

### HasGresDetail

`func (o *V0040JobInfo) HasGresDetail() bool`

HasGresDetail returns a boolean if a field has been set.

### GetGroupId

`func (o *V0040JobInfo) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V0040JobInfo) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V0040JobInfo) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V0040JobInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *V0040JobInfo) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *V0040JobInfo) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *V0040JobInfo) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *V0040JobInfo) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetHetJobId

`func (o *V0040JobInfo) GetHetJobId() V0040Uint32NoVal`

GetHetJobId returns the HetJobId field if non-nil, zero value otherwise.

### GetHetJobIdOk

`func (o *V0040JobInfo) GetHetJobIdOk() (*V0040Uint32NoVal, bool)`

GetHetJobIdOk returns a tuple with the HetJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobId

`func (o *V0040JobInfo) SetHetJobId(v V0040Uint32NoVal)`

SetHetJobId sets HetJobId field to given value.

### HasHetJobId

`func (o *V0040JobInfo) HasHetJobId() bool`

HasHetJobId returns a boolean if a field has been set.

### GetHetJobIdSet

`func (o *V0040JobInfo) GetHetJobIdSet() string`

GetHetJobIdSet returns the HetJobIdSet field if non-nil, zero value otherwise.

### GetHetJobIdSetOk

`func (o *V0040JobInfo) GetHetJobIdSetOk() (*string, bool)`

GetHetJobIdSetOk returns a tuple with the HetJobIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobIdSet

`func (o *V0040JobInfo) SetHetJobIdSet(v string)`

SetHetJobIdSet sets HetJobIdSet field to given value.

### HasHetJobIdSet

`func (o *V0040JobInfo) HasHetJobIdSet() bool`

HasHetJobIdSet returns a boolean if a field has been set.

### GetHetJobOffset

`func (o *V0040JobInfo) GetHetJobOffset() V0040Uint32NoVal`

GetHetJobOffset returns the HetJobOffset field if non-nil, zero value otherwise.

### GetHetJobOffsetOk

`func (o *V0040JobInfo) GetHetJobOffsetOk() (*V0040Uint32NoVal, bool)`

GetHetJobOffsetOk returns a tuple with the HetJobOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobOffset

`func (o *V0040JobInfo) SetHetJobOffset(v V0040Uint32NoVal)`

SetHetJobOffset sets HetJobOffset field to given value.

### HasHetJobOffset

`func (o *V0040JobInfo) HasHetJobOffset() bool`

HasHetJobOffset returns a boolean if a field has been set.

### GetJobId

`func (o *V0040JobInfo) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0040JobInfo) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0040JobInfo) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0040JobInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobResources

`func (o *V0040JobInfo) GetJobResources() V0040JobRes`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *V0040JobInfo) GetJobResourcesOk() (*V0040JobRes, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *V0040JobInfo) SetJobResources(v V0040JobRes)`

SetJobResources sets JobResources field to given value.

### HasJobResources

`func (o *V0040JobInfo) HasJobResources() bool`

HasJobResources returns a boolean if a field has been set.

### GetJobSizeStr

`func (o *V0040JobInfo) GetJobSizeStr() []string`

GetJobSizeStr returns the JobSizeStr field if non-nil, zero value otherwise.

### GetJobSizeStrOk

`func (o *V0040JobInfo) GetJobSizeStrOk() (*[]string, bool)`

GetJobSizeStrOk returns a tuple with the JobSizeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSizeStr

`func (o *V0040JobInfo) SetJobSizeStr(v []string)`

SetJobSizeStr sets JobSizeStr field to given value.

### HasJobSizeStr

`func (o *V0040JobInfo) HasJobSizeStr() bool`

HasJobSizeStr returns a boolean if a field has been set.

### GetJobState

`func (o *V0040JobInfo) GetJobState() []string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *V0040JobInfo) GetJobStateOk() (*[]string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *V0040JobInfo) SetJobState(v []string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *V0040JobInfo) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetLastSchedEvaluation

`func (o *V0040JobInfo) GetLastSchedEvaluation() V0040Uint64NoVal`

GetLastSchedEvaluation returns the LastSchedEvaluation field if non-nil, zero value otherwise.

### GetLastSchedEvaluationOk

`func (o *V0040JobInfo) GetLastSchedEvaluationOk() (*V0040Uint64NoVal, bool)`

GetLastSchedEvaluationOk returns a tuple with the LastSchedEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedEvaluation

`func (o *V0040JobInfo) SetLastSchedEvaluation(v V0040Uint64NoVal)`

SetLastSchedEvaluation sets LastSchedEvaluation field to given value.

### HasLastSchedEvaluation

`func (o *V0040JobInfo) HasLastSchedEvaluation() bool`

HasLastSchedEvaluation returns a boolean if a field has been set.

### GetLicenses

`func (o *V0040JobInfo) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0040JobInfo) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0040JobInfo) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0040JobInfo) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMailType

`func (o *V0040JobInfo) GetMailType() []string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *V0040JobInfo) GetMailTypeOk() (*[]string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *V0040JobInfo) SetMailType(v []string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *V0040JobInfo) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *V0040JobInfo) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *V0040JobInfo) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *V0040JobInfo) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *V0040JobInfo) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMaxCpus

`func (o *V0040JobInfo) GetMaxCpus() V0040Uint32NoVal`

GetMaxCpus returns the MaxCpus field if non-nil, zero value otherwise.

### GetMaxCpusOk

`func (o *V0040JobInfo) GetMaxCpusOk() (*V0040Uint32NoVal, bool)`

GetMaxCpusOk returns a tuple with the MaxCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpus

`func (o *V0040JobInfo) SetMaxCpus(v V0040Uint32NoVal)`

SetMaxCpus sets MaxCpus field to given value.

### HasMaxCpus

`func (o *V0040JobInfo) HasMaxCpus() bool`

HasMaxCpus returns a boolean if a field has been set.

### GetMaxNodes

`func (o *V0040JobInfo) GetMaxNodes() V0040Uint32NoVal`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *V0040JobInfo) GetMaxNodesOk() (*V0040Uint32NoVal, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *V0040JobInfo) SetMaxNodes(v V0040Uint32NoVal)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *V0040JobInfo) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0040JobInfo) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0040JobInfo) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0040JobInfo) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0040JobInfo) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *V0040JobInfo) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *V0040JobInfo) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *V0040JobInfo) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *V0040JobInfo) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *V0040JobInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040JobInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040JobInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040JobInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *V0040JobInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *V0040JobInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *V0040JobInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *V0040JobInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodes

`func (o *V0040JobInfo) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0040JobInfo) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0040JobInfo) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0040JobInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNice

`func (o *V0040JobInfo) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0040JobInfo) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0040JobInfo) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0040JobInfo) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0040JobInfo) GetTasksPerCore() V0040Uint16NoVal`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0040JobInfo) GetTasksPerCoreOk() (*V0040Uint16NoVal, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0040JobInfo) SetTasksPerCore(v V0040Uint16NoVal)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0040JobInfo) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerTres

`func (o *V0040JobInfo) GetTasksPerTres() V0040Uint16NoVal`

GetTasksPerTres returns the TasksPerTres field if non-nil, zero value otherwise.

### GetTasksPerTresOk

`func (o *V0040JobInfo) GetTasksPerTresOk() (*V0040Uint16NoVal, bool)`

GetTasksPerTresOk returns a tuple with the TasksPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerTres

`func (o *V0040JobInfo) SetTasksPerTres(v V0040Uint16NoVal)`

SetTasksPerTres sets TasksPerTres field to given value.

### HasTasksPerTres

`func (o *V0040JobInfo) HasTasksPerTres() bool`

HasTasksPerTres returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *V0040JobInfo) GetTasksPerNode() V0040Uint16NoVal`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *V0040JobInfo) GetTasksPerNodeOk() (*V0040Uint16NoVal, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *V0040JobInfo) SetTasksPerNode(v V0040Uint16NoVal)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *V0040JobInfo) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0040JobInfo) GetTasksPerSocket() V0040Uint16NoVal`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0040JobInfo) GetTasksPerSocketOk() (*V0040Uint16NoVal, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0040JobInfo) SetTasksPerSocket(v V0040Uint16NoVal)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0040JobInfo) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *V0040JobInfo) GetTasksPerBoard() V0040Uint16NoVal`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *V0040JobInfo) GetTasksPerBoardOk() (*V0040Uint16NoVal, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *V0040JobInfo) SetTasksPerBoard(v V0040Uint16NoVal)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *V0040JobInfo) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetCpus

`func (o *V0040JobInfo) GetCpus() V0040Uint32NoVal`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0040JobInfo) GetCpusOk() (*V0040Uint32NoVal, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0040JobInfo) SetCpus(v V0040Uint32NoVal)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0040JobInfo) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0040JobInfo) GetNodeCount() V0040Uint32NoVal`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0040JobInfo) GetNodeCountOk() (*V0040Uint32NoVal, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0040JobInfo) SetNodeCount(v V0040Uint32NoVal)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0040JobInfo) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTasks

`func (o *V0040JobInfo) GetTasks() V0040Uint32NoVal`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0040JobInfo) GetTasksOk() (*V0040Uint32NoVal, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0040JobInfo) SetTasks(v V0040Uint32NoVal)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0040JobInfo) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPartition

`func (o *V0040JobInfo) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0040JobInfo) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0040JobInfo) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0040JobInfo) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPrefer

`func (o *V0040JobInfo) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *V0040JobInfo) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *V0040JobInfo) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *V0040JobInfo) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0040JobInfo) GetMemoryPerCpu() V0040Uint64NoVal`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0040JobInfo) GetMemoryPerCpuOk() (*V0040Uint64NoVal, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0040JobInfo) SetMemoryPerCpu(v V0040Uint64NoVal)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0040JobInfo) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0040JobInfo) GetMemoryPerNode() V0040Uint64NoVal`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0040JobInfo) GetMemoryPerNodeOk() (*V0040Uint64NoVal, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0040JobInfo) SetMemoryPerNode(v V0040Uint64NoVal)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0040JobInfo) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0040JobInfo) GetMinimumCpusPerNode() V0040Uint16NoVal`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0040JobInfo) GetMinimumCpusPerNodeOk() (*V0040Uint16NoVal, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0040JobInfo) SetMinimumCpusPerNode(v V0040Uint16NoVal)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0040JobInfo) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMinimumTmpDiskPerNode

`func (o *V0040JobInfo) GetMinimumTmpDiskPerNode() V0040Uint32NoVal`

GetMinimumTmpDiskPerNode returns the MinimumTmpDiskPerNode field if non-nil, zero value otherwise.

### GetMinimumTmpDiskPerNodeOk

`func (o *V0040JobInfo) GetMinimumTmpDiskPerNodeOk() (*V0040Uint32NoVal, bool)`

GetMinimumTmpDiskPerNodeOk returns a tuple with the MinimumTmpDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTmpDiskPerNode

`func (o *V0040JobInfo) SetMinimumTmpDiskPerNode(v V0040Uint32NoVal)`

SetMinimumTmpDiskPerNode sets MinimumTmpDiskPerNode field to given value.

### HasMinimumTmpDiskPerNode

`func (o *V0040JobInfo) HasMinimumTmpDiskPerNode() bool`

HasMinimumTmpDiskPerNode returns a boolean if a field has been set.

### GetPower

`func (o *V0040JobInfo) GetPower() V0041JobInfoMsgInnerPower`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *V0040JobInfo) GetPowerOk() (*V0041JobInfoMsgInnerPower, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *V0040JobInfo) SetPower(v V0041JobInfoMsgInnerPower)`

SetPower sets Power field to given value.

### HasPower

`func (o *V0040JobInfo) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPreemptTime

`func (o *V0040JobInfo) GetPreemptTime() V0040Uint64NoVal`

GetPreemptTime returns the PreemptTime field if non-nil, zero value otherwise.

### GetPreemptTimeOk

`func (o *V0040JobInfo) GetPreemptTimeOk() (*V0040Uint64NoVal, bool)`

GetPreemptTimeOk returns a tuple with the PreemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptTime

`func (o *V0040JobInfo) SetPreemptTime(v V0040Uint64NoVal)`

SetPreemptTime sets PreemptTime field to given value.

### HasPreemptTime

`func (o *V0040JobInfo) HasPreemptTime() bool`

HasPreemptTime returns a boolean if a field has been set.

### GetPreemptableTime

`func (o *V0040JobInfo) GetPreemptableTime() V0040Uint64NoVal`

GetPreemptableTime returns the PreemptableTime field if non-nil, zero value otherwise.

### GetPreemptableTimeOk

`func (o *V0040JobInfo) GetPreemptableTimeOk() (*V0040Uint64NoVal, bool)`

GetPreemptableTimeOk returns a tuple with the PreemptableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptableTime

`func (o *V0040JobInfo) SetPreemptableTime(v V0040Uint64NoVal)`

SetPreemptableTime sets PreemptableTime field to given value.

### HasPreemptableTime

`func (o *V0040JobInfo) HasPreemptableTime() bool`

HasPreemptableTime returns a boolean if a field has been set.

### GetPreSusTime

`func (o *V0040JobInfo) GetPreSusTime() V0040Uint64NoVal`

GetPreSusTime returns the PreSusTime field if non-nil, zero value otherwise.

### GetPreSusTimeOk

`func (o *V0040JobInfo) GetPreSusTimeOk() (*V0040Uint64NoVal, bool)`

GetPreSusTimeOk returns a tuple with the PreSusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSusTime

`func (o *V0040JobInfo) SetPreSusTime(v V0040Uint64NoVal)`

SetPreSusTime sets PreSusTime field to given value.

### HasPreSusTime

`func (o *V0040JobInfo) HasPreSusTime() bool`

HasPreSusTime returns a boolean if a field has been set.

### GetPriority

`func (o *V0040JobInfo) GetPriority() V0040Uint32NoVal`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0040JobInfo) GetPriorityOk() (*V0040Uint32NoVal, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0040JobInfo) SetPriority(v V0040Uint32NoVal)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0040JobInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *V0040JobInfo) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V0040JobInfo) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V0040JobInfo) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V0040JobInfo) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *V0040JobInfo) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0040JobInfo) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0040JobInfo) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0040JobInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *V0040JobInfo) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *V0040JobInfo) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *V0040JobInfo) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *V0040JobInfo) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *V0040JobInfo) GetRequiredNodes() string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *V0040JobInfo) GetRequiredNodesOk() (*string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *V0040JobInfo) SetRequiredNodes(v string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *V0040JobInfo) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetMinimumSwitches

`func (o *V0040JobInfo) GetMinimumSwitches() int32`

GetMinimumSwitches returns the MinimumSwitches field if non-nil, zero value otherwise.

### GetMinimumSwitchesOk

`func (o *V0040JobInfo) GetMinimumSwitchesOk() (*int32, bool)`

GetMinimumSwitchesOk returns a tuple with the MinimumSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSwitches

`func (o *V0040JobInfo) SetMinimumSwitches(v int32)`

SetMinimumSwitches sets MinimumSwitches field to given value.

### HasMinimumSwitches

`func (o *V0040JobInfo) HasMinimumSwitches() bool`

HasMinimumSwitches returns a boolean if a field has been set.

### GetRequeue

`func (o *V0040JobInfo) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0040JobInfo) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0040JobInfo) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0040JobInfo) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetResizeTime

`func (o *V0040JobInfo) GetResizeTime() V0040Uint64NoVal`

GetResizeTime returns the ResizeTime field if non-nil, zero value otherwise.

### GetResizeTimeOk

`func (o *V0040JobInfo) GetResizeTimeOk() (*V0040Uint64NoVal, bool)`

GetResizeTimeOk returns a tuple with the ResizeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeTime

`func (o *V0040JobInfo) SetResizeTime(v V0040Uint64NoVal)`

SetResizeTime sets ResizeTime field to given value.

### HasResizeTime

`func (o *V0040JobInfo) HasResizeTime() bool`

HasResizeTime returns a boolean if a field has been set.

### GetRestartCnt

`func (o *V0040JobInfo) GetRestartCnt() int32`

GetRestartCnt returns the RestartCnt field if non-nil, zero value otherwise.

### GetRestartCntOk

`func (o *V0040JobInfo) GetRestartCntOk() (*int32, bool)`

GetRestartCntOk returns a tuple with the RestartCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCnt

`func (o *V0040JobInfo) SetRestartCnt(v int32)`

SetRestartCnt sets RestartCnt field to given value.

### HasRestartCnt

`func (o *V0040JobInfo) HasRestartCnt() bool`

HasRestartCnt returns a boolean if a field has been set.

### GetResvName

`func (o *V0040JobInfo) GetResvName() string`

GetResvName returns the ResvName field if non-nil, zero value otherwise.

### GetResvNameOk

`func (o *V0040JobInfo) GetResvNameOk() (*string, bool)`

GetResvNameOk returns a tuple with the ResvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResvName

`func (o *V0040JobInfo) SetResvName(v string)`

SetResvName sets ResvName field to given value.

### HasResvName

`func (o *V0040JobInfo) HasResvName() bool`

HasResvName returns a boolean if a field has been set.

### GetScheduledNodes

`func (o *V0040JobInfo) GetScheduledNodes() string`

GetScheduledNodes returns the ScheduledNodes field if non-nil, zero value otherwise.

### GetScheduledNodesOk

`func (o *V0040JobInfo) GetScheduledNodesOk() (*string, bool)`

GetScheduledNodesOk returns a tuple with the ScheduledNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledNodes

`func (o *V0040JobInfo) SetScheduledNodes(v string)`

SetScheduledNodes sets ScheduledNodes field to given value.

### HasScheduledNodes

`func (o *V0040JobInfo) HasScheduledNodes() bool`

HasScheduledNodes returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *V0040JobInfo) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *V0040JobInfo) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *V0040JobInfo) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *V0040JobInfo) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetShared

`func (o *V0040JobInfo) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *V0040JobInfo) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *V0040JobInfo) SetShared(v []string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *V0040JobInfo) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetExclusive

`func (o *V0040JobInfo) GetExclusive() []string`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *V0040JobInfo) GetExclusiveOk() (*[]string, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *V0040JobInfo) SetExclusive(v []string)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *V0040JobInfo) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0040JobInfo) GetOversubscribe() bool`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0040JobInfo) GetOversubscribeOk() (*bool, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0040JobInfo) SetOversubscribe(v bool)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0040JobInfo) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetShowFlags

`func (o *V0040JobInfo) GetShowFlags() []string`

GetShowFlags returns the ShowFlags field if non-nil, zero value otherwise.

### GetShowFlagsOk

`func (o *V0040JobInfo) GetShowFlagsOk() (*[]string, bool)`

GetShowFlagsOk returns a tuple with the ShowFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFlags

`func (o *V0040JobInfo) SetShowFlags(v []string)`

SetShowFlags sets ShowFlags field to given value.

### HasShowFlags

`func (o *V0040JobInfo) HasShowFlags() bool`

HasShowFlags returns a boolean if a field has been set.

### GetSocketsPerBoard

`func (o *V0040JobInfo) GetSocketsPerBoard() int32`

GetSocketsPerBoard returns the SocketsPerBoard field if non-nil, zero value otherwise.

### GetSocketsPerBoardOk

`func (o *V0040JobInfo) GetSocketsPerBoardOk() (*int32, bool)`

GetSocketsPerBoardOk returns a tuple with the SocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerBoard

`func (o *V0040JobInfo) SetSocketsPerBoard(v int32)`

SetSocketsPerBoard sets SocketsPerBoard field to given value.

### HasSocketsPerBoard

`func (o *V0040JobInfo) HasSocketsPerBoard() bool`

HasSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0040JobInfo) GetSocketsPerNode() V0040Uint16NoVal`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0040JobInfo) GetSocketsPerNodeOk() (*V0040Uint16NoVal, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0040JobInfo) SetSocketsPerNode(v V0040Uint16NoVal)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0040JobInfo) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetStartTime

`func (o *V0040JobInfo) GetStartTime() V0040Uint64NoVal`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0040JobInfo) GetStartTimeOk() (*V0040Uint64NoVal, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0040JobInfo) SetStartTime(v V0040Uint64NoVal)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0040JobInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStateDescription

`func (o *V0040JobInfo) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *V0040JobInfo) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *V0040JobInfo) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *V0040JobInfo) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetStateReason

`func (o *V0040JobInfo) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *V0040JobInfo) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *V0040JobInfo) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *V0040JobInfo) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStandardError

`func (o *V0040JobInfo) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0040JobInfo) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0040JobInfo) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0040JobInfo) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0040JobInfo) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0040JobInfo) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0040JobInfo) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0040JobInfo) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0040JobInfo) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0040JobInfo) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0040JobInfo) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0040JobInfo) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetSubmitTime

`func (o *V0040JobInfo) GetSubmitTime() V0040Uint64NoVal`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *V0040JobInfo) GetSubmitTimeOk() (*V0040Uint64NoVal, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *V0040JobInfo) SetSubmitTime(v V0040Uint64NoVal)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *V0040JobInfo) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0040JobInfo) GetSuspendTime() V0040Uint64NoVal`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0040JobInfo) GetSuspendTimeOk() (*V0040Uint64NoVal, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0040JobInfo) SetSuspendTime(v V0040Uint64NoVal)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0040JobInfo) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.

### GetSystemComment

`func (o *V0040JobInfo) GetSystemComment() string`

GetSystemComment returns the SystemComment field if non-nil, zero value otherwise.

### GetSystemCommentOk

`func (o *V0040JobInfo) GetSystemCommentOk() (*string, bool)`

GetSystemCommentOk returns a tuple with the SystemComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemComment

`func (o *V0040JobInfo) SetSystemComment(v string)`

SetSystemComment sets SystemComment field to given value.

### HasSystemComment

`func (o *V0040JobInfo) HasSystemComment() bool`

HasSystemComment returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0040JobInfo) GetTimeLimit() V0040Uint32NoVal`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0040JobInfo) GetTimeLimitOk() (*V0040Uint32NoVal, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0040JobInfo) SetTimeLimit(v V0040Uint32NoVal)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0040JobInfo) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0040JobInfo) GetTimeMinimum() V0040Uint32NoVal`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0040JobInfo) GetTimeMinimumOk() (*V0040Uint32NoVal, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0040JobInfo) SetTimeMinimum(v V0040Uint32NoVal)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0040JobInfo) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0040JobInfo) GetThreadsPerCore() V0040Uint16NoVal`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0040JobInfo) GetThreadsPerCoreOk() (*V0040Uint16NoVal, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0040JobInfo) SetThreadsPerCore(v V0040Uint16NoVal)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0040JobInfo) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTresBind

`func (o *V0040JobInfo) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *V0040JobInfo) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *V0040JobInfo) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *V0040JobInfo) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *V0040JobInfo) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *V0040JobInfo) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *V0040JobInfo) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *V0040JobInfo) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *V0040JobInfo) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *V0040JobInfo) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *V0040JobInfo) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *V0040JobInfo) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *V0040JobInfo) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *V0040JobInfo) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *V0040JobInfo) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *V0040JobInfo) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *V0040JobInfo) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *V0040JobInfo) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *V0040JobInfo) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *V0040JobInfo) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *V0040JobInfo) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *V0040JobInfo) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *V0040JobInfo) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *V0040JobInfo) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetTresReqStr

`func (o *V0040JobInfo) GetTresReqStr() string`

GetTresReqStr returns the TresReqStr field if non-nil, zero value otherwise.

### GetTresReqStrOk

`func (o *V0040JobInfo) GetTresReqStrOk() (*string, bool)`

GetTresReqStrOk returns a tuple with the TresReqStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresReqStr

`func (o *V0040JobInfo) SetTresReqStr(v string)`

SetTresReqStr sets TresReqStr field to given value.

### HasTresReqStr

`func (o *V0040JobInfo) HasTresReqStr() bool`

HasTresReqStr returns a boolean if a field has been set.

### GetTresAllocStr

`func (o *V0040JobInfo) GetTresAllocStr() string`

GetTresAllocStr returns the TresAllocStr field if non-nil, zero value otherwise.

### GetTresAllocStrOk

`func (o *V0040JobInfo) GetTresAllocStrOk() (*string, bool)`

GetTresAllocStrOk returns a tuple with the TresAllocStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresAllocStr

`func (o *V0040JobInfo) SetTresAllocStr(v string)`

SetTresAllocStr sets TresAllocStr field to given value.

### HasTresAllocStr

`func (o *V0040JobInfo) HasTresAllocStr() bool`

HasTresAllocStr returns a boolean if a field has been set.

### GetUserId

`func (o *V0040JobInfo) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0040JobInfo) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0040JobInfo) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0040JobInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *V0040JobInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *V0040JobInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *V0040JobInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *V0040JobInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetMaximumSwitchWaitTime

`func (o *V0040JobInfo) GetMaximumSwitchWaitTime() int32`

GetMaximumSwitchWaitTime returns the MaximumSwitchWaitTime field if non-nil, zero value otherwise.

### GetMaximumSwitchWaitTimeOk

`func (o *V0040JobInfo) GetMaximumSwitchWaitTimeOk() (*int32, bool)`

GetMaximumSwitchWaitTimeOk returns a tuple with the MaximumSwitchWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSwitchWaitTime

`func (o *V0040JobInfo) SetMaximumSwitchWaitTime(v int32)`

SetMaximumSwitchWaitTime sets MaximumSwitchWaitTime field to given value.

### HasMaximumSwitchWaitTime

`func (o *V0040JobInfo) HasMaximumSwitchWaitTime() bool`

HasMaximumSwitchWaitTime returns a boolean if a field has been set.

### GetWckey

`func (o *V0040JobInfo) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0040JobInfo) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0040JobInfo) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0040JobInfo) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0040JobInfo) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0040JobInfo) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0040JobInfo) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0040JobInfo) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


