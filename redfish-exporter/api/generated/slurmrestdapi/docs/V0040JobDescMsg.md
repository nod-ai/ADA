# V0040JobDescMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AccountGatherFrequency** | Pointer to **string** |  | [optional] 
**AdminComment** | Pointer to **string** |  | [optional] 
**AllocationNodeList** | Pointer to **string** |  | [optional] 
**AllocationNodePort** | Pointer to **int32** |  | [optional] 
**Argv** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to **string** |  | [optional] 
**BatchFeatures** | Pointer to **string** |  | [optional] 
**BeginTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**BurstBuffer** | Pointer to **string** |  | [optional] 
**Clusters** | Pointer to **string** |  | [optional] 
**ClusterConstraint** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Contiguous** | Pointer to **bool** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 
**CoreSpecification** | Pointer to **int32** |  | [optional] 
**ThreadSpecification** | Pointer to **int32** |  | [optional] 
**CpuBinding** | Pointer to **string** |  | [optional] 
**CpuBindingFlags** | Pointer to **[]string** |  | [optional] 
**CpuFrequency** | Pointer to **string** |  | [optional] 
**CpusPerTres** | Pointer to **string** |  | [optional] 
**Crontab** | Pointer to [**V0040CronEntry**](V0040CronEntry.md) |  | [optional] 
**Deadline** | Pointer to **int64** |  | [optional] 
**DelayBoot** | Pointer to **int32** |  | [optional] 
**Dependency** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
**Rlimits** | Pointer to [**V0040JobDescMsgRlimits**](V0040JobDescMsgRlimits.md) |  | [optional] 
**ExcludedNodes** | Pointer to **[]string** |  | [optional] 
**Extra** | Pointer to **string** |  | [optional] 
**Constraints** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**HetjobGroup** | Pointer to **int32** |  | [optional] 
**Immediate** | Pointer to **bool** |  | [optional] 
**JobId** | Pointer to **int32** |  | [optional] 
**KillOnNodeFail** | Pointer to **bool** |  | [optional] 
**Licenses** | Pointer to **string** |  | [optional] 
**MailType** | Pointer to **[]string** |  | [optional] 
**MailUser** | Pointer to **string** |  | [optional] 
**McsLabel** | Pointer to **string** |  | [optional] 
**MemoryBinding** | Pointer to **string** |  | [optional] 
**MemoryBindingType** | Pointer to **[]string** |  | [optional] 
**MemoryPerTres** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Nice** | Pointer to **int32** |  | [optional] 
**Tasks** | Pointer to **int32** |  | [optional] 
**OpenMode** | Pointer to **[]string** |  | [optional] 
**ReservePorts** | Pointer to **int32** |  | [optional] 
**Overcommit** | Pointer to **bool** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**DistributionPlaneSize** | Pointer to **int32** |  | [optional] 
**PowerFlags** | Pointer to **[]string** |  | [optional] 
**Prefer** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Profile** | Pointer to **[]string** |  | [optional] 
**Qos** | Pointer to **string** |  | [optional] 
**Reboot** | Pointer to **bool** |  | [optional] 
**RequiredNodes** | Pointer to **[]string** |  | [optional] 
**Requeue** | Pointer to **bool** |  | [optional] 
**Reservation** | Pointer to **string** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **[]string** |  | [optional] 
**Exclusive** | Pointer to **[]string** |  | [optional] 
**Oversubscribe** | Pointer to **bool** |  | [optional] 
**SiteFactor** | Pointer to **int32** |  | [optional] 
**SpankEnvironment** | Pointer to **[]string** |  | [optional] 
**Distribution** | Pointer to **string** |  | [optional] 
**TimeLimit** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**TimeMinimum** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**TresBind** | Pointer to **string** |  | [optional] 
**TresFreq** | Pointer to **string** |  | [optional] 
**TresPerJob** | Pointer to **string** |  | [optional] 
**TresPerNode** | Pointer to **string** |  | [optional] 
**TresPerSocket** | Pointer to **string** |  | [optional] 
**TresPerTask** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**WaitAllNodes** | Pointer to **bool** |  | [optional] 
**KillWarningFlags** | Pointer to **[]string** |  | [optional] 
**KillWarningSignal** | Pointer to **string** |  | [optional] 
**KillWarningDelay** | Pointer to [**V0040Uint16NoVal**](V0040Uint16NoVal.md) |  | [optional] 
**CurrentWorkingDirectory** | Pointer to **string** |  | [optional] 
**CpusPerTask** | Pointer to **int32** |  | [optional] 
**MinimumCpus** | Pointer to **int32** |  | [optional] 
**MaximumCpus** | Pointer to **int32** |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**MinimumNodes** | Pointer to **int32** |  | [optional] 
**MaximumNodes** | Pointer to **int32** |  | [optional] 
**MinimumBoardsPerNode** | Pointer to **int32** |  | [optional] 
**MinimumSocketsPerBoard** | Pointer to **int32** |  | [optional] 
**SocketsPerNode** | Pointer to **int32** |  | [optional] 
**ThreadsPerCore** | Pointer to **int32** |  | [optional] 
**TasksPerNode** | Pointer to **int32** |  | [optional] 
**TasksPerSocket** | Pointer to **int32** |  | [optional] 
**TasksPerCore** | Pointer to **int32** |  | [optional] 
**TasksPerBoard** | Pointer to **int32** |  | [optional] 
**NtasksPerTres** | Pointer to **int32** |  | [optional] 
**MinimumCpusPerNode** | Pointer to **int32** |  | [optional] 
**MemoryPerCpu** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**TemporaryDiskPerNode** | Pointer to **int32** |  | [optional] 
**SelinuxContext** | Pointer to **string** |  | [optional] 
**RequiredSwitches** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**StandardError** | Pointer to **string** |  | [optional] 
**StandardInput** | Pointer to **string** |  | [optional] 
**StandardOutput** | Pointer to **string** |  | [optional] 
**WaitForSwitch** | Pointer to **int32** |  | [optional] 
**Wckey** | Pointer to **string** |  | [optional] 
**X11** | Pointer to **[]string** |  | [optional] 
**X11MagicCookie** | Pointer to **string** |  | [optional] 
**X11TargetHost** | Pointer to **string** |  | [optional] 
**X11TargetPort** | Pointer to **int32** |  | [optional] 

## Methods

### NewV0040JobDescMsg

`func NewV0040JobDescMsg() *V0040JobDescMsg`

NewV0040JobDescMsg instantiates a new V0040JobDescMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobDescMsgWithDefaults

`func NewV0040JobDescMsgWithDefaults() *V0040JobDescMsg`

NewV0040JobDescMsgWithDefaults instantiates a new V0040JobDescMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0040JobDescMsg) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0040JobDescMsg) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0040JobDescMsg) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0040JobDescMsg) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountGatherFrequency

`func (o *V0040JobDescMsg) GetAccountGatherFrequency() string`

GetAccountGatherFrequency returns the AccountGatherFrequency field if non-nil, zero value otherwise.

### GetAccountGatherFrequencyOk

`func (o *V0040JobDescMsg) GetAccountGatherFrequencyOk() (*string, bool)`

GetAccountGatherFrequencyOk returns a tuple with the AccountGatherFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGatherFrequency

`func (o *V0040JobDescMsg) SetAccountGatherFrequency(v string)`

SetAccountGatherFrequency sets AccountGatherFrequency field to given value.

### HasAccountGatherFrequency

`func (o *V0040JobDescMsg) HasAccountGatherFrequency() bool`

HasAccountGatherFrequency returns a boolean if a field has been set.

### GetAdminComment

`func (o *V0040JobDescMsg) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *V0040JobDescMsg) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *V0040JobDescMsg) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *V0040JobDescMsg) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAllocationNodeList

`func (o *V0040JobDescMsg) GetAllocationNodeList() string`

GetAllocationNodeList returns the AllocationNodeList field if non-nil, zero value otherwise.

### GetAllocationNodeListOk

`func (o *V0040JobDescMsg) GetAllocationNodeListOk() (*string, bool)`

GetAllocationNodeListOk returns a tuple with the AllocationNodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodeList

`func (o *V0040JobDescMsg) SetAllocationNodeList(v string)`

SetAllocationNodeList sets AllocationNodeList field to given value.

### HasAllocationNodeList

`func (o *V0040JobDescMsg) HasAllocationNodeList() bool`

HasAllocationNodeList returns a boolean if a field has been set.

### GetAllocationNodePort

`func (o *V0040JobDescMsg) GetAllocationNodePort() int32`

GetAllocationNodePort returns the AllocationNodePort field if non-nil, zero value otherwise.

### GetAllocationNodePortOk

`func (o *V0040JobDescMsg) GetAllocationNodePortOk() (*int32, bool)`

GetAllocationNodePortOk returns a tuple with the AllocationNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodePort

`func (o *V0040JobDescMsg) SetAllocationNodePort(v int32)`

SetAllocationNodePort sets AllocationNodePort field to given value.

### HasAllocationNodePort

`func (o *V0040JobDescMsg) HasAllocationNodePort() bool`

HasAllocationNodePort returns a boolean if a field has been set.

### GetArgv

`func (o *V0040JobDescMsg) GetArgv() []string`

GetArgv returns the Argv field if non-nil, zero value otherwise.

### GetArgvOk

`func (o *V0040JobDescMsg) GetArgvOk() (*[]string, bool)`

GetArgvOk returns a tuple with the Argv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgv

`func (o *V0040JobDescMsg) SetArgv(v []string)`

SetArgv sets Argv field to given value.

### HasArgv

`func (o *V0040JobDescMsg) HasArgv() bool`

HasArgv returns a boolean if a field has been set.

### GetArray

`func (o *V0040JobDescMsg) GetArray() string`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *V0040JobDescMsg) GetArrayOk() (*string, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *V0040JobDescMsg) SetArray(v string)`

SetArray sets Array field to given value.

### HasArray

`func (o *V0040JobDescMsg) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0040JobDescMsg) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0040JobDescMsg) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0040JobDescMsg) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0040JobDescMsg) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBeginTime

`func (o *V0040JobDescMsg) GetBeginTime() V0040Uint64NoVal`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *V0040JobDescMsg) GetBeginTimeOk() (*V0040Uint64NoVal, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *V0040JobDescMsg) SetBeginTime(v V0040Uint64NoVal)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *V0040JobDescMsg) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.

### GetFlags

`func (o *V0040JobDescMsg) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0040JobDescMsg) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0040JobDescMsg) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0040JobDescMsg) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0040JobDescMsg) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0040JobDescMsg) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0040JobDescMsg) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0040JobDescMsg) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetClusters

`func (o *V0040JobDescMsg) GetClusters() string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0040JobDescMsg) GetClustersOk() (*string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0040JobDescMsg) SetClusters(v string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0040JobDescMsg) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetClusterConstraint

`func (o *V0040JobDescMsg) GetClusterConstraint() string`

GetClusterConstraint returns the ClusterConstraint field if non-nil, zero value otherwise.

### GetClusterConstraintOk

`func (o *V0040JobDescMsg) GetClusterConstraintOk() (*string, bool)`

GetClusterConstraintOk returns a tuple with the ClusterConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConstraint

`func (o *V0040JobDescMsg) SetClusterConstraint(v string)`

SetClusterConstraint sets ClusterConstraint field to given value.

### HasClusterConstraint

`func (o *V0040JobDescMsg) HasClusterConstraint() bool`

HasClusterConstraint returns a boolean if a field has been set.

### GetComment

`func (o *V0040JobDescMsg) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0040JobDescMsg) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0040JobDescMsg) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0040JobDescMsg) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContiguous

`func (o *V0040JobDescMsg) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *V0040JobDescMsg) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *V0040JobDescMsg) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *V0040JobDescMsg) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetContainer

`func (o *V0040JobDescMsg) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0040JobDescMsg) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0040JobDescMsg) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0040JobDescMsg) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *V0040JobDescMsg) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *V0040JobDescMsg) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *V0040JobDescMsg) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *V0040JobDescMsg) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetCoreSpecification

`func (o *V0040JobDescMsg) GetCoreSpecification() int32`

GetCoreSpecification returns the CoreSpecification field if non-nil, zero value otherwise.

### GetCoreSpecificationOk

`func (o *V0040JobDescMsg) GetCoreSpecificationOk() (*int32, bool)`

GetCoreSpecificationOk returns a tuple with the CoreSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecification

`func (o *V0040JobDescMsg) SetCoreSpecification(v int32)`

SetCoreSpecification sets CoreSpecification field to given value.

### HasCoreSpecification

`func (o *V0040JobDescMsg) HasCoreSpecification() bool`

HasCoreSpecification returns a boolean if a field has been set.

### GetThreadSpecification

`func (o *V0040JobDescMsg) GetThreadSpecification() int32`

GetThreadSpecification returns the ThreadSpecification field if non-nil, zero value otherwise.

### GetThreadSpecificationOk

`func (o *V0040JobDescMsg) GetThreadSpecificationOk() (*int32, bool)`

GetThreadSpecificationOk returns a tuple with the ThreadSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpecification

`func (o *V0040JobDescMsg) SetThreadSpecification(v int32)`

SetThreadSpecification sets ThreadSpecification field to given value.

### HasThreadSpecification

`func (o *V0040JobDescMsg) HasThreadSpecification() bool`

HasThreadSpecification returns a boolean if a field has been set.

### GetCpuBinding

`func (o *V0040JobDescMsg) GetCpuBinding() string`

GetCpuBinding returns the CpuBinding field if non-nil, zero value otherwise.

### GetCpuBindingOk

`func (o *V0040JobDescMsg) GetCpuBindingOk() (*string, bool)`

GetCpuBindingOk returns a tuple with the CpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBinding

`func (o *V0040JobDescMsg) SetCpuBinding(v string)`

SetCpuBinding sets CpuBinding field to given value.

### HasCpuBinding

`func (o *V0040JobDescMsg) HasCpuBinding() bool`

HasCpuBinding returns a boolean if a field has been set.

### GetCpuBindingFlags

`func (o *V0040JobDescMsg) GetCpuBindingFlags() []string`

GetCpuBindingFlags returns the CpuBindingFlags field if non-nil, zero value otherwise.

### GetCpuBindingFlagsOk

`func (o *V0040JobDescMsg) GetCpuBindingFlagsOk() (*[]string, bool)`

GetCpuBindingFlagsOk returns a tuple with the CpuBindingFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBindingFlags

`func (o *V0040JobDescMsg) SetCpuBindingFlags(v []string)`

SetCpuBindingFlags sets CpuBindingFlags field to given value.

### HasCpuBindingFlags

`func (o *V0040JobDescMsg) HasCpuBindingFlags() bool`

HasCpuBindingFlags returns a boolean if a field has been set.

### GetCpuFrequency

`func (o *V0040JobDescMsg) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *V0040JobDescMsg) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *V0040JobDescMsg) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *V0040JobDescMsg) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *V0040JobDescMsg) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *V0040JobDescMsg) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *V0040JobDescMsg) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *V0040JobDescMsg) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetCrontab

`func (o *V0040JobDescMsg) GetCrontab() V0040CronEntry`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *V0040JobDescMsg) GetCrontabOk() (*V0040CronEntry, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *V0040JobDescMsg) SetCrontab(v V0040CronEntry)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *V0040JobDescMsg) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### GetDeadline

`func (o *V0040JobDescMsg) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0040JobDescMsg) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0040JobDescMsg) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0040JobDescMsg) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0040JobDescMsg) GetDelayBoot() int32`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0040JobDescMsg) GetDelayBootOk() (*int32, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0040JobDescMsg) SetDelayBoot(v int32)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0040JobDescMsg) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0040JobDescMsg) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0040JobDescMsg) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0040JobDescMsg) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0040JobDescMsg) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetEndTime

`func (o *V0040JobDescMsg) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0040JobDescMsg) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0040JobDescMsg) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0040JobDescMsg) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnvironment

`func (o *V0040JobDescMsg) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V0040JobDescMsg) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V0040JobDescMsg) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *V0040JobDescMsg) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRlimits

`func (o *V0040JobDescMsg) GetRlimits() V0040JobDescMsgRlimits`

GetRlimits returns the Rlimits field if non-nil, zero value otherwise.

### GetRlimitsOk

`func (o *V0040JobDescMsg) GetRlimitsOk() (*V0040JobDescMsgRlimits, bool)`

GetRlimitsOk returns a tuple with the Rlimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlimits

`func (o *V0040JobDescMsg) SetRlimits(v V0040JobDescMsgRlimits)`

SetRlimits sets Rlimits field to given value.

### HasRlimits

`func (o *V0040JobDescMsg) HasRlimits() bool`

HasRlimits returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *V0040JobDescMsg) GetExcludedNodes() []string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *V0040JobDescMsg) GetExcludedNodesOk() (*[]string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *V0040JobDescMsg) SetExcludedNodes(v []string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *V0040JobDescMsg) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExtra

`func (o *V0040JobDescMsg) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0040JobDescMsg) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0040JobDescMsg) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0040JobDescMsg) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetConstraints

`func (o *V0040JobDescMsg) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *V0040JobDescMsg) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *V0040JobDescMsg) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *V0040JobDescMsg) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetGroupId

`func (o *V0040JobDescMsg) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V0040JobDescMsg) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V0040JobDescMsg) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V0040JobDescMsg) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHetjobGroup

`func (o *V0040JobDescMsg) GetHetjobGroup() int32`

GetHetjobGroup returns the HetjobGroup field if non-nil, zero value otherwise.

### GetHetjobGroupOk

`func (o *V0040JobDescMsg) GetHetjobGroupOk() (*int32, bool)`

GetHetjobGroupOk returns a tuple with the HetjobGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjobGroup

`func (o *V0040JobDescMsg) SetHetjobGroup(v int32)`

SetHetjobGroup sets HetjobGroup field to given value.

### HasHetjobGroup

`func (o *V0040JobDescMsg) HasHetjobGroup() bool`

HasHetjobGroup returns a boolean if a field has been set.

### GetImmediate

`func (o *V0040JobDescMsg) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *V0040JobDescMsg) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *V0040JobDescMsg) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *V0040JobDescMsg) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.

### GetJobId

`func (o *V0040JobDescMsg) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0040JobDescMsg) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0040JobDescMsg) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0040JobDescMsg) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetKillOnNodeFail

`func (o *V0040JobDescMsg) GetKillOnNodeFail() bool`

GetKillOnNodeFail returns the KillOnNodeFail field if non-nil, zero value otherwise.

### GetKillOnNodeFailOk

`func (o *V0040JobDescMsg) GetKillOnNodeFailOk() (*bool, bool)`

GetKillOnNodeFailOk returns a tuple with the KillOnNodeFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillOnNodeFail

`func (o *V0040JobDescMsg) SetKillOnNodeFail(v bool)`

SetKillOnNodeFail sets KillOnNodeFail field to given value.

### HasKillOnNodeFail

`func (o *V0040JobDescMsg) HasKillOnNodeFail() bool`

HasKillOnNodeFail returns a boolean if a field has been set.

### GetLicenses

`func (o *V0040JobDescMsg) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0040JobDescMsg) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0040JobDescMsg) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0040JobDescMsg) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMailType

`func (o *V0040JobDescMsg) GetMailType() []string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *V0040JobDescMsg) GetMailTypeOk() (*[]string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *V0040JobDescMsg) SetMailType(v []string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *V0040JobDescMsg) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *V0040JobDescMsg) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *V0040JobDescMsg) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *V0040JobDescMsg) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *V0040JobDescMsg) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0040JobDescMsg) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0040JobDescMsg) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0040JobDescMsg) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0040JobDescMsg) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryBinding

`func (o *V0040JobDescMsg) GetMemoryBinding() string`

GetMemoryBinding returns the MemoryBinding field if non-nil, zero value otherwise.

### GetMemoryBindingOk

`func (o *V0040JobDescMsg) GetMemoryBindingOk() (*string, bool)`

GetMemoryBindingOk returns a tuple with the MemoryBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBinding

`func (o *V0040JobDescMsg) SetMemoryBinding(v string)`

SetMemoryBinding sets MemoryBinding field to given value.

### HasMemoryBinding

`func (o *V0040JobDescMsg) HasMemoryBinding() bool`

HasMemoryBinding returns a boolean if a field has been set.

### GetMemoryBindingType

`func (o *V0040JobDescMsg) GetMemoryBindingType() []string`

GetMemoryBindingType returns the MemoryBindingType field if non-nil, zero value otherwise.

### GetMemoryBindingTypeOk

`func (o *V0040JobDescMsg) GetMemoryBindingTypeOk() (*[]string, bool)`

GetMemoryBindingTypeOk returns a tuple with the MemoryBindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBindingType

`func (o *V0040JobDescMsg) SetMemoryBindingType(v []string)`

SetMemoryBindingType sets MemoryBindingType field to given value.

### HasMemoryBindingType

`func (o *V0040JobDescMsg) HasMemoryBindingType() bool`

HasMemoryBindingType returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *V0040JobDescMsg) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *V0040JobDescMsg) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *V0040JobDescMsg) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *V0040JobDescMsg) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *V0040JobDescMsg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040JobDescMsg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040JobDescMsg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040JobDescMsg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *V0040JobDescMsg) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *V0040JobDescMsg) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *V0040JobDescMsg) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *V0040JobDescMsg) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNice

`func (o *V0040JobDescMsg) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0040JobDescMsg) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0040JobDescMsg) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0040JobDescMsg) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasks

`func (o *V0040JobDescMsg) GetTasks() int32`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0040JobDescMsg) GetTasksOk() (*int32, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0040JobDescMsg) SetTasks(v int32)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0040JobDescMsg) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOpenMode

`func (o *V0040JobDescMsg) GetOpenMode() []string`

GetOpenMode returns the OpenMode field if non-nil, zero value otherwise.

### GetOpenModeOk

`func (o *V0040JobDescMsg) GetOpenModeOk() (*[]string, bool)`

GetOpenModeOk returns a tuple with the OpenMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMode

`func (o *V0040JobDescMsg) SetOpenMode(v []string)`

SetOpenMode sets OpenMode field to given value.

### HasOpenMode

`func (o *V0040JobDescMsg) HasOpenMode() bool`

HasOpenMode returns a boolean if a field has been set.

### GetReservePorts

`func (o *V0040JobDescMsg) GetReservePorts() int32`

GetReservePorts returns the ReservePorts field if non-nil, zero value otherwise.

### GetReservePortsOk

`func (o *V0040JobDescMsg) GetReservePortsOk() (*int32, bool)`

GetReservePortsOk returns a tuple with the ReservePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservePorts

`func (o *V0040JobDescMsg) SetReservePorts(v int32)`

SetReservePorts sets ReservePorts field to given value.

### HasReservePorts

`func (o *V0040JobDescMsg) HasReservePorts() bool`

HasReservePorts returns a boolean if a field has been set.

### GetOvercommit

`func (o *V0040JobDescMsg) GetOvercommit() bool`

GetOvercommit returns the Overcommit field if non-nil, zero value otherwise.

### GetOvercommitOk

`func (o *V0040JobDescMsg) GetOvercommitOk() (*bool, bool)`

GetOvercommitOk returns a tuple with the Overcommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommit

`func (o *V0040JobDescMsg) SetOvercommit(v bool)`

SetOvercommit sets Overcommit field to given value.

### HasOvercommit

`func (o *V0040JobDescMsg) HasOvercommit() bool`

HasOvercommit returns a boolean if a field has been set.

### GetPartition

`func (o *V0040JobDescMsg) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0040JobDescMsg) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0040JobDescMsg) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0040JobDescMsg) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetDistributionPlaneSize

`func (o *V0040JobDescMsg) GetDistributionPlaneSize() int32`

GetDistributionPlaneSize returns the DistributionPlaneSize field if non-nil, zero value otherwise.

### GetDistributionPlaneSizeOk

`func (o *V0040JobDescMsg) GetDistributionPlaneSizeOk() (*int32, bool)`

GetDistributionPlaneSizeOk returns a tuple with the DistributionPlaneSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPlaneSize

`func (o *V0040JobDescMsg) SetDistributionPlaneSize(v int32)`

SetDistributionPlaneSize sets DistributionPlaneSize field to given value.

### HasDistributionPlaneSize

`func (o *V0040JobDescMsg) HasDistributionPlaneSize() bool`

HasDistributionPlaneSize returns a boolean if a field has been set.

### GetPowerFlags

`func (o *V0040JobDescMsg) GetPowerFlags() []string`

GetPowerFlags returns the PowerFlags field if non-nil, zero value otherwise.

### GetPowerFlagsOk

`func (o *V0040JobDescMsg) GetPowerFlagsOk() (*[]string, bool)`

GetPowerFlagsOk returns a tuple with the PowerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFlags

`func (o *V0040JobDescMsg) SetPowerFlags(v []string)`

SetPowerFlags sets PowerFlags field to given value.

### HasPowerFlags

`func (o *V0040JobDescMsg) HasPowerFlags() bool`

HasPowerFlags returns a boolean if a field has been set.

### GetPrefer

`func (o *V0040JobDescMsg) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *V0040JobDescMsg) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *V0040JobDescMsg) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *V0040JobDescMsg) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetPriority

`func (o *V0040JobDescMsg) GetPriority() V0040Uint32NoVal`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0040JobDescMsg) GetPriorityOk() (*V0040Uint32NoVal, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0040JobDescMsg) SetPriority(v V0040Uint32NoVal)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0040JobDescMsg) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *V0040JobDescMsg) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V0040JobDescMsg) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V0040JobDescMsg) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V0040JobDescMsg) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *V0040JobDescMsg) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0040JobDescMsg) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0040JobDescMsg) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0040JobDescMsg) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *V0040JobDescMsg) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *V0040JobDescMsg) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *V0040JobDescMsg) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *V0040JobDescMsg) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *V0040JobDescMsg) GetRequiredNodes() []string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *V0040JobDescMsg) GetRequiredNodesOk() (*[]string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *V0040JobDescMsg) SetRequiredNodes(v []string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *V0040JobDescMsg) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetRequeue

`func (o *V0040JobDescMsg) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0040JobDescMsg) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0040JobDescMsg) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0040JobDescMsg) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetReservation

`func (o *V0040JobDescMsg) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *V0040JobDescMsg) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *V0040JobDescMsg) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *V0040JobDescMsg) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetScript

`func (o *V0040JobDescMsg) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0040JobDescMsg) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0040JobDescMsg) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0040JobDescMsg) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetShared

`func (o *V0040JobDescMsg) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *V0040JobDescMsg) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *V0040JobDescMsg) SetShared(v []string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *V0040JobDescMsg) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetExclusive

`func (o *V0040JobDescMsg) GetExclusive() []string`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *V0040JobDescMsg) GetExclusiveOk() (*[]string, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *V0040JobDescMsg) SetExclusive(v []string)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *V0040JobDescMsg) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0040JobDescMsg) GetOversubscribe() bool`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0040JobDescMsg) GetOversubscribeOk() (*bool, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0040JobDescMsg) SetOversubscribe(v bool)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0040JobDescMsg) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetSiteFactor

`func (o *V0040JobDescMsg) GetSiteFactor() int32`

GetSiteFactor returns the SiteFactor field if non-nil, zero value otherwise.

### GetSiteFactorOk

`func (o *V0040JobDescMsg) GetSiteFactorOk() (*int32, bool)`

GetSiteFactorOk returns a tuple with the SiteFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteFactor

`func (o *V0040JobDescMsg) SetSiteFactor(v int32)`

SetSiteFactor sets SiteFactor field to given value.

### HasSiteFactor

`func (o *V0040JobDescMsg) HasSiteFactor() bool`

HasSiteFactor returns a boolean if a field has been set.

### GetSpankEnvironment

`func (o *V0040JobDescMsg) GetSpankEnvironment() []string`

GetSpankEnvironment returns the SpankEnvironment field if non-nil, zero value otherwise.

### GetSpankEnvironmentOk

`func (o *V0040JobDescMsg) GetSpankEnvironmentOk() (*[]string, bool)`

GetSpankEnvironmentOk returns a tuple with the SpankEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpankEnvironment

`func (o *V0040JobDescMsg) SetSpankEnvironment(v []string)`

SetSpankEnvironment sets SpankEnvironment field to given value.

### HasSpankEnvironment

`func (o *V0040JobDescMsg) HasSpankEnvironment() bool`

HasSpankEnvironment returns a boolean if a field has been set.

### GetDistribution

`func (o *V0040JobDescMsg) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *V0040JobDescMsg) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *V0040JobDescMsg) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *V0040JobDescMsg) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0040JobDescMsg) GetTimeLimit() V0040Uint32NoVal`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0040JobDescMsg) GetTimeLimitOk() (*V0040Uint32NoVal, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0040JobDescMsg) SetTimeLimit(v V0040Uint32NoVal)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0040JobDescMsg) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0040JobDescMsg) GetTimeMinimum() V0040Uint32NoVal`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0040JobDescMsg) GetTimeMinimumOk() (*V0040Uint32NoVal, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0040JobDescMsg) SetTimeMinimum(v V0040Uint32NoVal)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0040JobDescMsg) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetTresBind

`func (o *V0040JobDescMsg) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *V0040JobDescMsg) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *V0040JobDescMsg) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *V0040JobDescMsg) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *V0040JobDescMsg) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *V0040JobDescMsg) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *V0040JobDescMsg) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *V0040JobDescMsg) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *V0040JobDescMsg) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *V0040JobDescMsg) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *V0040JobDescMsg) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *V0040JobDescMsg) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *V0040JobDescMsg) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *V0040JobDescMsg) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *V0040JobDescMsg) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *V0040JobDescMsg) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *V0040JobDescMsg) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *V0040JobDescMsg) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *V0040JobDescMsg) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *V0040JobDescMsg) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *V0040JobDescMsg) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *V0040JobDescMsg) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *V0040JobDescMsg) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *V0040JobDescMsg) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetUserId

`func (o *V0040JobDescMsg) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0040JobDescMsg) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0040JobDescMsg) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0040JobDescMsg) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWaitAllNodes

`func (o *V0040JobDescMsg) GetWaitAllNodes() bool`

GetWaitAllNodes returns the WaitAllNodes field if non-nil, zero value otherwise.

### GetWaitAllNodesOk

`func (o *V0040JobDescMsg) GetWaitAllNodesOk() (*bool, bool)`

GetWaitAllNodesOk returns a tuple with the WaitAllNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitAllNodes

`func (o *V0040JobDescMsg) SetWaitAllNodes(v bool)`

SetWaitAllNodes sets WaitAllNodes field to given value.

### HasWaitAllNodes

`func (o *V0040JobDescMsg) HasWaitAllNodes() bool`

HasWaitAllNodes returns a boolean if a field has been set.

### GetKillWarningFlags

`func (o *V0040JobDescMsg) GetKillWarningFlags() []string`

GetKillWarningFlags returns the KillWarningFlags field if non-nil, zero value otherwise.

### GetKillWarningFlagsOk

`func (o *V0040JobDescMsg) GetKillWarningFlagsOk() (*[]string, bool)`

GetKillWarningFlagsOk returns a tuple with the KillWarningFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningFlags

`func (o *V0040JobDescMsg) SetKillWarningFlags(v []string)`

SetKillWarningFlags sets KillWarningFlags field to given value.

### HasKillWarningFlags

`func (o *V0040JobDescMsg) HasKillWarningFlags() bool`

HasKillWarningFlags returns a boolean if a field has been set.

### GetKillWarningSignal

`func (o *V0040JobDescMsg) GetKillWarningSignal() string`

GetKillWarningSignal returns the KillWarningSignal field if non-nil, zero value otherwise.

### GetKillWarningSignalOk

`func (o *V0040JobDescMsg) GetKillWarningSignalOk() (*string, bool)`

GetKillWarningSignalOk returns a tuple with the KillWarningSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningSignal

`func (o *V0040JobDescMsg) SetKillWarningSignal(v string)`

SetKillWarningSignal sets KillWarningSignal field to given value.

### HasKillWarningSignal

`func (o *V0040JobDescMsg) HasKillWarningSignal() bool`

HasKillWarningSignal returns a boolean if a field has been set.

### GetKillWarningDelay

`func (o *V0040JobDescMsg) GetKillWarningDelay() V0040Uint16NoVal`

GetKillWarningDelay returns the KillWarningDelay field if non-nil, zero value otherwise.

### GetKillWarningDelayOk

`func (o *V0040JobDescMsg) GetKillWarningDelayOk() (*V0040Uint16NoVal, bool)`

GetKillWarningDelayOk returns a tuple with the KillWarningDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillWarningDelay

`func (o *V0040JobDescMsg) SetKillWarningDelay(v V0040Uint16NoVal)`

SetKillWarningDelay sets KillWarningDelay field to given value.

### HasKillWarningDelay

`func (o *V0040JobDescMsg) HasKillWarningDelay() bool`

HasKillWarningDelay returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0040JobDescMsg) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0040JobDescMsg) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0040JobDescMsg) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0040JobDescMsg) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0040JobDescMsg) GetCpusPerTask() int32`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0040JobDescMsg) GetCpusPerTaskOk() (*int32, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0040JobDescMsg) SetCpusPerTask(v int32)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0040JobDescMsg) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetMinimumCpus

`func (o *V0040JobDescMsg) GetMinimumCpus() int32`

GetMinimumCpus returns the MinimumCpus field if non-nil, zero value otherwise.

### GetMinimumCpusOk

`func (o *V0040JobDescMsg) GetMinimumCpusOk() (*int32, bool)`

GetMinimumCpusOk returns a tuple with the MinimumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpus

`func (o *V0040JobDescMsg) SetMinimumCpus(v int32)`

SetMinimumCpus sets MinimumCpus field to given value.

### HasMinimumCpus

`func (o *V0040JobDescMsg) HasMinimumCpus() bool`

HasMinimumCpus returns a boolean if a field has been set.

### GetMaximumCpus

`func (o *V0040JobDescMsg) GetMaximumCpus() int32`

GetMaximumCpus returns the MaximumCpus field if non-nil, zero value otherwise.

### GetMaximumCpusOk

`func (o *V0040JobDescMsg) GetMaximumCpusOk() (*int32, bool)`

GetMaximumCpusOk returns a tuple with the MaximumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpus

`func (o *V0040JobDescMsg) SetMaximumCpus(v int32)`

SetMaximumCpus sets MaximumCpus field to given value.

### HasMaximumCpus

`func (o *V0040JobDescMsg) HasMaximumCpus() bool`

HasMaximumCpus returns a boolean if a field has been set.

### GetNodes

`func (o *V0040JobDescMsg) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0040JobDescMsg) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0040JobDescMsg) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0040JobDescMsg) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetMinimumNodes

`func (o *V0040JobDescMsg) GetMinimumNodes() int32`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *V0040JobDescMsg) GetMinimumNodesOk() (*int32, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *V0040JobDescMsg) SetMinimumNodes(v int32)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *V0040JobDescMsg) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetMaximumNodes

`func (o *V0040JobDescMsg) GetMaximumNodes() int32`

GetMaximumNodes returns the MaximumNodes field if non-nil, zero value otherwise.

### GetMaximumNodesOk

`func (o *V0040JobDescMsg) GetMaximumNodesOk() (*int32, bool)`

GetMaximumNodesOk returns a tuple with the MaximumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodes

`func (o *V0040JobDescMsg) SetMaximumNodes(v int32)`

SetMaximumNodes sets MaximumNodes field to given value.

### HasMaximumNodes

`func (o *V0040JobDescMsg) HasMaximumNodes() bool`

HasMaximumNodes returns a boolean if a field has been set.

### GetMinimumBoardsPerNode

`func (o *V0040JobDescMsg) GetMinimumBoardsPerNode() int32`

GetMinimumBoardsPerNode returns the MinimumBoardsPerNode field if non-nil, zero value otherwise.

### GetMinimumBoardsPerNodeOk

`func (o *V0040JobDescMsg) GetMinimumBoardsPerNodeOk() (*int32, bool)`

GetMinimumBoardsPerNodeOk returns a tuple with the MinimumBoardsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBoardsPerNode

`func (o *V0040JobDescMsg) SetMinimumBoardsPerNode(v int32)`

SetMinimumBoardsPerNode sets MinimumBoardsPerNode field to given value.

### HasMinimumBoardsPerNode

`func (o *V0040JobDescMsg) HasMinimumBoardsPerNode() bool`

HasMinimumBoardsPerNode returns a boolean if a field has been set.

### GetMinimumSocketsPerBoard

`func (o *V0040JobDescMsg) GetMinimumSocketsPerBoard() int32`

GetMinimumSocketsPerBoard returns the MinimumSocketsPerBoard field if non-nil, zero value otherwise.

### GetMinimumSocketsPerBoardOk

`func (o *V0040JobDescMsg) GetMinimumSocketsPerBoardOk() (*int32, bool)`

GetMinimumSocketsPerBoardOk returns a tuple with the MinimumSocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSocketsPerBoard

`func (o *V0040JobDescMsg) SetMinimumSocketsPerBoard(v int32)`

SetMinimumSocketsPerBoard sets MinimumSocketsPerBoard field to given value.

### HasMinimumSocketsPerBoard

`func (o *V0040JobDescMsg) HasMinimumSocketsPerBoard() bool`

HasMinimumSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0040JobDescMsg) GetSocketsPerNode() int32`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0040JobDescMsg) GetSocketsPerNodeOk() (*int32, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0040JobDescMsg) SetSocketsPerNode(v int32)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0040JobDescMsg) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0040JobDescMsg) GetThreadsPerCore() int32`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0040JobDescMsg) GetThreadsPerCoreOk() (*int32, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0040JobDescMsg) SetThreadsPerCore(v int32)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0040JobDescMsg) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *V0040JobDescMsg) GetTasksPerNode() int32`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *V0040JobDescMsg) GetTasksPerNodeOk() (*int32, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *V0040JobDescMsg) SetTasksPerNode(v int32)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *V0040JobDescMsg) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0040JobDescMsg) GetTasksPerSocket() int32`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0040JobDescMsg) GetTasksPerSocketOk() (*int32, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0040JobDescMsg) SetTasksPerSocket(v int32)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0040JobDescMsg) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0040JobDescMsg) GetTasksPerCore() int32`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0040JobDescMsg) GetTasksPerCoreOk() (*int32, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0040JobDescMsg) SetTasksPerCore(v int32)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0040JobDescMsg) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *V0040JobDescMsg) GetTasksPerBoard() int32`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *V0040JobDescMsg) GetTasksPerBoardOk() (*int32, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *V0040JobDescMsg) SetTasksPerBoard(v int32)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *V0040JobDescMsg) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetNtasksPerTres

`func (o *V0040JobDescMsg) GetNtasksPerTres() int32`

GetNtasksPerTres returns the NtasksPerTres field if non-nil, zero value otherwise.

### GetNtasksPerTresOk

`func (o *V0040JobDescMsg) GetNtasksPerTresOk() (*int32, bool)`

GetNtasksPerTresOk returns a tuple with the NtasksPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtasksPerTres

`func (o *V0040JobDescMsg) SetNtasksPerTres(v int32)`

SetNtasksPerTres sets NtasksPerTres field to given value.

### HasNtasksPerTres

`func (o *V0040JobDescMsg) HasNtasksPerTres() bool`

HasNtasksPerTres returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0040JobDescMsg) GetMinimumCpusPerNode() int32`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0040JobDescMsg) GetMinimumCpusPerNodeOk() (*int32, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0040JobDescMsg) SetMinimumCpusPerNode(v int32)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0040JobDescMsg) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0040JobDescMsg) GetMemoryPerCpu() V0040Uint64NoVal`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0040JobDescMsg) GetMemoryPerCpuOk() (*V0040Uint64NoVal, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0040JobDescMsg) SetMemoryPerCpu(v V0040Uint64NoVal)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0040JobDescMsg) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0040JobDescMsg) GetMemoryPerNode() V0040Uint64NoVal`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0040JobDescMsg) GetMemoryPerNodeOk() (*V0040Uint64NoVal, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0040JobDescMsg) SetMemoryPerNode(v V0040Uint64NoVal)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0040JobDescMsg) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetTemporaryDiskPerNode

`func (o *V0040JobDescMsg) GetTemporaryDiskPerNode() int32`

GetTemporaryDiskPerNode returns the TemporaryDiskPerNode field if non-nil, zero value otherwise.

### GetTemporaryDiskPerNodeOk

`func (o *V0040JobDescMsg) GetTemporaryDiskPerNodeOk() (*int32, bool)`

GetTemporaryDiskPerNodeOk returns a tuple with the TemporaryDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDiskPerNode

`func (o *V0040JobDescMsg) SetTemporaryDiskPerNode(v int32)`

SetTemporaryDiskPerNode sets TemporaryDiskPerNode field to given value.

### HasTemporaryDiskPerNode

`func (o *V0040JobDescMsg) HasTemporaryDiskPerNode() bool`

HasTemporaryDiskPerNode returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *V0040JobDescMsg) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *V0040JobDescMsg) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *V0040JobDescMsg) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *V0040JobDescMsg) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetRequiredSwitches

`func (o *V0040JobDescMsg) GetRequiredSwitches() V0040Uint32NoVal`

GetRequiredSwitches returns the RequiredSwitches field if non-nil, zero value otherwise.

### GetRequiredSwitchesOk

`func (o *V0040JobDescMsg) GetRequiredSwitchesOk() (*V0040Uint32NoVal, bool)`

GetRequiredSwitchesOk returns a tuple with the RequiredSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSwitches

`func (o *V0040JobDescMsg) SetRequiredSwitches(v V0040Uint32NoVal)`

SetRequiredSwitches sets RequiredSwitches field to given value.

### HasRequiredSwitches

`func (o *V0040JobDescMsg) HasRequiredSwitches() bool`

HasRequiredSwitches returns a boolean if a field has been set.

### GetStandardError

`func (o *V0040JobDescMsg) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0040JobDescMsg) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0040JobDescMsg) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0040JobDescMsg) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0040JobDescMsg) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0040JobDescMsg) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0040JobDescMsg) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0040JobDescMsg) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0040JobDescMsg) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0040JobDescMsg) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0040JobDescMsg) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0040JobDescMsg) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetWaitForSwitch

`func (o *V0040JobDescMsg) GetWaitForSwitch() int32`

GetWaitForSwitch returns the WaitForSwitch field if non-nil, zero value otherwise.

### GetWaitForSwitchOk

`func (o *V0040JobDescMsg) GetWaitForSwitchOk() (*int32, bool)`

GetWaitForSwitchOk returns a tuple with the WaitForSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForSwitch

`func (o *V0040JobDescMsg) SetWaitForSwitch(v int32)`

SetWaitForSwitch sets WaitForSwitch field to given value.

### HasWaitForSwitch

`func (o *V0040JobDescMsg) HasWaitForSwitch() bool`

HasWaitForSwitch returns a boolean if a field has been set.

### GetWckey

`func (o *V0040JobDescMsg) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0040JobDescMsg) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0040JobDescMsg) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0040JobDescMsg) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetX11

`func (o *V0040JobDescMsg) GetX11() []string`

GetX11 returns the X11 field if non-nil, zero value otherwise.

### GetX11Ok

`func (o *V0040JobDescMsg) GetX11Ok() (*[]string, bool)`

GetX11Ok returns a tuple with the X11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11

`func (o *V0040JobDescMsg) SetX11(v []string)`

SetX11 sets X11 field to given value.

### HasX11

`func (o *V0040JobDescMsg) HasX11() bool`

HasX11 returns a boolean if a field has been set.

### GetX11MagicCookie

`func (o *V0040JobDescMsg) GetX11MagicCookie() string`

GetX11MagicCookie returns the X11MagicCookie field if non-nil, zero value otherwise.

### GetX11MagicCookieOk

`func (o *V0040JobDescMsg) GetX11MagicCookieOk() (*string, bool)`

GetX11MagicCookieOk returns a tuple with the X11MagicCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11MagicCookie

`func (o *V0040JobDescMsg) SetX11MagicCookie(v string)`

SetX11MagicCookie sets X11MagicCookie field to given value.

### HasX11MagicCookie

`func (o *V0040JobDescMsg) HasX11MagicCookie() bool`

HasX11MagicCookie returns a boolean if a field has been set.

### GetX11TargetHost

`func (o *V0040JobDescMsg) GetX11TargetHost() string`

GetX11TargetHost returns the X11TargetHost field if non-nil, zero value otherwise.

### GetX11TargetHostOk

`func (o *V0040JobDescMsg) GetX11TargetHostOk() (*string, bool)`

GetX11TargetHostOk returns a tuple with the X11TargetHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11TargetHost

`func (o *V0040JobDescMsg) SetX11TargetHost(v string)`

SetX11TargetHost sets X11TargetHost field to given value.

### HasX11TargetHost

`func (o *V0040JobDescMsg) HasX11TargetHost() bool`

HasX11TargetHost returns a boolean if a field has been set.

### GetX11TargetPort

`func (o *V0040JobDescMsg) GetX11TargetPort() int32`

GetX11TargetPort returns the X11TargetPort field if non-nil, zero value otherwise.

### GetX11TargetPortOk

`func (o *V0040JobDescMsg) GetX11TargetPortOk() (*int32, bool)`

GetX11TargetPortOk returns a tuple with the X11TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX11TargetPort

`func (o *V0040JobDescMsg) SetX11TargetPort(v int32)`

SetX11TargetPort sets X11TargetPort field to given value.

### HasX11TargetPort

`func (o *V0040JobDescMsg) HasX11TargetPort() bool`

HasX11TargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


