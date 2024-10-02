# V0040Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**BurstbufferNetworkAddress** | Pointer to **string** |  | [optional] 
**Boards** | Pointer to **int32** |  | [optional] 
**BootTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**Cores** | Pointer to **int32** |  | [optional] 
**SpecializedCores** | Pointer to **int32** |  | [optional] 
**CpuBinding** | Pointer to **int32** |  | [optional] 
**CpuLoad** | Pointer to **int32** |  | [optional] 
**FreeMem** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Cpus** | Pointer to **int32** |  | [optional] 
**EffectiveCpus** | Pointer to **int32** |  | [optional] 
**SpecializedCpus** | Pointer to **string** |  | [optional] 
**Energy** | Pointer to [**V0040AcctGatherEnergy**](V0040AcctGatherEnergy.md) |  | [optional] 
**ExternalSensors** | Pointer to [**V0040ExtSensorsData**](V0040ExtSensorsData.md) |  | [optional] 
**Extra** | Pointer to **string** |  | [optional] 
**Power** | Pointer to [**V0040PowerMgmtData**](V0040PowerMgmtData.md) |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**ActiveFeatures** | Pointer to **[]string** |  | [optional] 
**Gres** | Pointer to **string** |  | [optional] 
**GresDrained** | Pointer to **string** |  | [optional] 
**GresUsed** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**LastBusy** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**McsLabel** | Pointer to **string** |  | [optional] 
**SpecializedMemory** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextStateAfterReboot** | Pointer to **[]string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**State** | Pointer to **[]string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Partitions** | Pointer to **[]string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RealMemory** | Pointer to **int64** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ReasonChangedAt** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**ReasonSetByUser** | Pointer to **string** |  | [optional] 
**ResumeAfter** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Reservation** | Pointer to **string** |  | [optional] 
**AllocMemory** | Pointer to **int64** |  | [optional] 
**AllocCpus** | Pointer to **int32** |  | [optional] 
**AllocIdleCpus** | Pointer to **int32** |  | [optional] 
**TresUsed** | Pointer to **string** |  | [optional] 
**TresWeighted** | Pointer to **float64** |  | [optional] 
**SlurmdStartTime** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Sockets** | Pointer to **int32** |  | [optional] 
**Threads** | Pointer to **int32** |  | [optional] 
**TemporaryDisk** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Tres** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV0040Node

`func NewV0040Node() *V0040Node`

NewV0040Node instantiates a new V0040Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040NodeWithDefaults

`func NewV0040NodeWithDefaults() *V0040Node`

NewV0040NodeWithDefaults instantiates a new V0040Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *V0040Node) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *V0040Node) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *V0040Node) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *V0040Node) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBurstbufferNetworkAddress

`func (o *V0040Node) GetBurstbufferNetworkAddress() string`

GetBurstbufferNetworkAddress returns the BurstbufferNetworkAddress field if non-nil, zero value otherwise.

### GetBurstbufferNetworkAddressOk

`func (o *V0040Node) GetBurstbufferNetworkAddressOk() (*string, bool)`

GetBurstbufferNetworkAddressOk returns a tuple with the BurstbufferNetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstbufferNetworkAddress

`func (o *V0040Node) SetBurstbufferNetworkAddress(v string)`

SetBurstbufferNetworkAddress sets BurstbufferNetworkAddress field to given value.

### HasBurstbufferNetworkAddress

`func (o *V0040Node) HasBurstbufferNetworkAddress() bool`

HasBurstbufferNetworkAddress returns a boolean if a field has been set.

### GetBoards

`func (o *V0040Node) GetBoards() int32`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *V0040Node) GetBoardsOk() (*int32, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *V0040Node) SetBoards(v int32)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *V0040Node) HasBoards() bool`

HasBoards returns a boolean if a field has been set.

### GetBootTime

`func (o *V0040Node) GetBootTime() V0040Uint64NoVal`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *V0040Node) GetBootTimeOk() (*V0040Uint64NoVal, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *V0040Node) SetBootTime(v V0040Uint64NoVal)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *V0040Node) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetClusterName

`func (o *V0040Node) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *V0040Node) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *V0040Node) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *V0040Node) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCores

`func (o *V0040Node) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *V0040Node) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *V0040Node) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *V0040Node) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetSpecializedCores

`func (o *V0040Node) GetSpecializedCores() int32`

GetSpecializedCores returns the SpecializedCores field if non-nil, zero value otherwise.

### GetSpecializedCoresOk

`func (o *V0040Node) GetSpecializedCoresOk() (*int32, bool)`

GetSpecializedCoresOk returns a tuple with the SpecializedCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializedCores

`func (o *V0040Node) SetSpecializedCores(v int32)`

SetSpecializedCores sets SpecializedCores field to given value.

### HasSpecializedCores

`func (o *V0040Node) HasSpecializedCores() bool`

HasSpecializedCores returns a boolean if a field has been set.

### GetCpuBinding

`func (o *V0040Node) GetCpuBinding() int32`

GetCpuBinding returns the CpuBinding field if non-nil, zero value otherwise.

### GetCpuBindingOk

`func (o *V0040Node) GetCpuBindingOk() (*int32, bool)`

GetCpuBindingOk returns a tuple with the CpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBinding

`func (o *V0040Node) SetCpuBinding(v int32)`

SetCpuBinding sets CpuBinding field to given value.

### HasCpuBinding

`func (o *V0040Node) HasCpuBinding() bool`

HasCpuBinding returns a boolean if a field has been set.

### GetCpuLoad

`func (o *V0040Node) GetCpuLoad() int32`

GetCpuLoad returns the CpuLoad field if non-nil, zero value otherwise.

### GetCpuLoadOk

`func (o *V0040Node) GetCpuLoadOk() (*int32, bool)`

GetCpuLoadOk returns a tuple with the CpuLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad

`func (o *V0040Node) SetCpuLoad(v int32)`

SetCpuLoad sets CpuLoad field to given value.

### HasCpuLoad

`func (o *V0040Node) HasCpuLoad() bool`

HasCpuLoad returns a boolean if a field has been set.

### GetFreeMem

`func (o *V0040Node) GetFreeMem() V0040Uint64NoVal`

GetFreeMem returns the FreeMem field if non-nil, zero value otherwise.

### GetFreeMemOk

`func (o *V0040Node) GetFreeMemOk() (*V0040Uint64NoVal, bool)`

GetFreeMemOk returns a tuple with the FreeMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMem

`func (o *V0040Node) SetFreeMem(v V0040Uint64NoVal)`

SetFreeMem sets FreeMem field to given value.

### HasFreeMem

`func (o *V0040Node) HasFreeMem() bool`

HasFreeMem returns a boolean if a field has been set.

### GetCpus

`func (o *V0040Node) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0040Node) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0040Node) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0040Node) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetEffectiveCpus

`func (o *V0040Node) GetEffectiveCpus() int32`

GetEffectiveCpus returns the EffectiveCpus field if non-nil, zero value otherwise.

### GetEffectiveCpusOk

`func (o *V0040Node) GetEffectiveCpusOk() (*int32, bool)`

GetEffectiveCpusOk returns a tuple with the EffectiveCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCpus

`func (o *V0040Node) SetEffectiveCpus(v int32)`

SetEffectiveCpus sets EffectiveCpus field to given value.

### HasEffectiveCpus

`func (o *V0040Node) HasEffectiveCpus() bool`

HasEffectiveCpus returns a boolean if a field has been set.

### GetSpecializedCpus

`func (o *V0040Node) GetSpecializedCpus() string`

GetSpecializedCpus returns the SpecializedCpus field if non-nil, zero value otherwise.

### GetSpecializedCpusOk

`func (o *V0040Node) GetSpecializedCpusOk() (*string, bool)`

GetSpecializedCpusOk returns a tuple with the SpecializedCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializedCpus

`func (o *V0040Node) SetSpecializedCpus(v string)`

SetSpecializedCpus sets SpecializedCpus field to given value.

### HasSpecializedCpus

`func (o *V0040Node) HasSpecializedCpus() bool`

HasSpecializedCpus returns a boolean if a field has been set.

### GetEnergy

`func (o *V0040Node) GetEnergy() V0040AcctGatherEnergy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *V0040Node) GetEnergyOk() (*V0040AcctGatherEnergy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *V0040Node) SetEnergy(v V0040AcctGatherEnergy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *V0040Node) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.

### GetExternalSensors

`func (o *V0040Node) GetExternalSensors() V0040ExtSensorsData`

GetExternalSensors returns the ExternalSensors field if non-nil, zero value otherwise.

### GetExternalSensorsOk

`func (o *V0040Node) GetExternalSensorsOk() (*V0040ExtSensorsData, bool)`

GetExternalSensorsOk returns a tuple with the ExternalSensors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSensors

`func (o *V0040Node) SetExternalSensors(v V0040ExtSensorsData)`

SetExternalSensors sets ExternalSensors field to given value.

### HasExternalSensors

`func (o *V0040Node) HasExternalSensors() bool`

HasExternalSensors returns a boolean if a field has been set.

### GetExtra

`func (o *V0040Node) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0040Node) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0040Node) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0040Node) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetPower

`func (o *V0040Node) GetPower() V0040PowerMgmtData`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *V0040Node) GetPowerOk() (*V0040PowerMgmtData, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *V0040Node) SetPower(v V0040PowerMgmtData)`

SetPower sets Power field to given value.

### HasPower

`func (o *V0040Node) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetFeatures

`func (o *V0040Node) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0040Node) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0040Node) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0040Node) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetActiveFeatures

`func (o *V0040Node) GetActiveFeatures() []string`

GetActiveFeatures returns the ActiveFeatures field if non-nil, zero value otherwise.

### GetActiveFeaturesOk

`func (o *V0040Node) GetActiveFeaturesOk() (*[]string, bool)`

GetActiveFeaturesOk returns a tuple with the ActiveFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFeatures

`func (o *V0040Node) SetActiveFeatures(v []string)`

SetActiveFeatures sets ActiveFeatures field to given value.

### HasActiveFeatures

`func (o *V0040Node) HasActiveFeatures() bool`

HasActiveFeatures returns a boolean if a field has been set.

### GetGres

`func (o *V0040Node) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0040Node) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0040Node) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0040Node) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetGresDrained

`func (o *V0040Node) GetGresDrained() string`

GetGresDrained returns the GresDrained field if non-nil, zero value otherwise.

### GetGresDrainedOk

`func (o *V0040Node) GetGresDrainedOk() (*string, bool)`

GetGresDrainedOk returns a tuple with the GresDrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDrained

`func (o *V0040Node) SetGresDrained(v string)`

SetGresDrained sets GresDrained field to given value.

### HasGresDrained

`func (o *V0040Node) HasGresDrained() bool`

HasGresDrained returns a boolean if a field has been set.

### GetGresUsed

`func (o *V0040Node) GetGresUsed() string`

GetGresUsed returns the GresUsed field if non-nil, zero value otherwise.

### GetGresUsedOk

`func (o *V0040Node) GetGresUsedOk() (*string, bool)`

GetGresUsedOk returns a tuple with the GresUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresUsed

`func (o *V0040Node) SetGresUsed(v string)`

SetGresUsed sets GresUsed field to given value.

### HasGresUsed

`func (o *V0040Node) HasGresUsed() bool`

HasGresUsed returns a boolean if a field has been set.

### GetInstanceId

`func (o *V0040Node) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *V0040Node) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *V0040Node) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *V0040Node) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceType

`func (o *V0040Node) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *V0040Node) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *V0040Node) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *V0040Node) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetLastBusy

`func (o *V0040Node) GetLastBusy() V0040Uint64NoVal`

GetLastBusy returns the LastBusy field if non-nil, zero value otherwise.

### GetLastBusyOk

`func (o *V0040Node) GetLastBusyOk() (*V0040Uint64NoVal, bool)`

GetLastBusyOk returns a tuple with the LastBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBusy

`func (o *V0040Node) SetLastBusy(v V0040Uint64NoVal)`

SetLastBusy sets LastBusy field to given value.

### HasLastBusy

`func (o *V0040Node) HasLastBusy() bool`

HasLastBusy returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0040Node) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0040Node) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0040Node) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0040Node) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetSpecializedMemory

`func (o *V0040Node) GetSpecializedMemory() int64`

GetSpecializedMemory returns the SpecializedMemory field if non-nil, zero value otherwise.

### GetSpecializedMemoryOk

`func (o *V0040Node) GetSpecializedMemoryOk() (*int64, bool)`

GetSpecializedMemoryOk returns a tuple with the SpecializedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializedMemory

`func (o *V0040Node) SetSpecializedMemory(v int64)`

SetSpecializedMemory sets SpecializedMemory field to given value.

### HasSpecializedMemory

`func (o *V0040Node) HasSpecializedMemory() bool`

HasSpecializedMemory returns a boolean if a field has been set.

### GetName

`func (o *V0040Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextStateAfterReboot

`func (o *V0040Node) GetNextStateAfterReboot() []string`

GetNextStateAfterReboot returns the NextStateAfterReboot field if non-nil, zero value otherwise.

### GetNextStateAfterRebootOk

`func (o *V0040Node) GetNextStateAfterRebootOk() (*[]string, bool)`

GetNextStateAfterRebootOk returns a tuple with the NextStateAfterReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStateAfterReboot

`func (o *V0040Node) SetNextStateAfterReboot(v []string)`

SetNextStateAfterReboot sets NextStateAfterReboot field to given value.

### HasNextStateAfterReboot

`func (o *V0040Node) HasNextStateAfterReboot() bool`

HasNextStateAfterReboot returns a boolean if a field has been set.

### GetAddress

`func (o *V0040Node) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V0040Node) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V0040Node) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V0040Node) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *V0040Node) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0040Node) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0040Node) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0040Node) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetState

`func (o *V0040Node) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0040Node) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0040Node) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0040Node) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *V0040Node) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *V0040Node) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *V0040Node) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *V0040Node) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOwner

`func (o *V0040Node) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V0040Node) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V0040Node) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V0040Node) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPartitions

`func (o *V0040Node) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0040Node) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0040Node) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0040Node) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetPort

`func (o *V0040Node) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V0040Node) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V0040Node) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V0040Node) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRealMemory

`func (o *V0040Node) GetRealMemory() int64`

GetRealMemory returns the RealMemory field if non-nil, zero value otherwise.

### GetRealMemoryOk

`func (o *V0040Node) GetRealMemoryOk() (*int64, bool)`

GetRealMemoryOk returns a tuple with the RealMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealMemory

`func (o *V0040Node) SetRealMemory(v int64)`

SetRealMemory sets RealMemory field to given value.

### HasRealMemory

`func (o *V0040Node) HasRealMemory() bool`

HasRealMemory returns a boolean if a field has been set.

### GetComment

`func (o *V0040Node) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0040Node) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0040Node) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0040Node) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *V0040Node) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V0040Node) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V0040Node) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V0040Node) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonChangedAt

`func (o *V0040Node) GetReasonChangedAt() V0040Uint64NoVal`

GetReasonChangedAt returns the ReasonChangedAt field if non-nil, zero value otherwise.

### GetReasonChangedAtOk

`func (o *V0040Node) GetReasonChangedAtOk() (*V0040Uint64NoVal, bool)`

GetReasonChangedAtOk returns a tuple with the ReasonChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonChangedAt

`func (o *V0040Node) SetReasonChangedAt(v V0040Uint64NoVal)`

SetReasonChangedAt sets ReasonChangedAt field to given value.

### HasReasonChangedAt

`func (o *V0040Node) HasReasonChangedAt() bool`

HasReasonChangedAt returns a boolean if a field has been set.

### GetReasonSetByUser

`func (o *V0040Node) GetReasonSetByUser() string`

GetReasonSetByUser returns the ReasonSetByUser field if non-nil, zero value otherwise.

### GetReasonSetByUserOk

`func (o *V0040Node) GetReasonSetByUserOk() (*string, bool)`

GetReasonSetByUserOk returns a tuple with the ReasonSetByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSetByUser

`func (o *V0040Node) SetReasonSetByUser(v string)`

SetReasonSetByUser sets ReasonSetByUser field to given value.

### HasReasonSetByUser

`func (o *V0040Node) HasReasonSetByUser() bool`

HasReasonSetByUser returns a boolean if a field has been set.

### GetResumeAfter

`func (o *V0040Node) GetResumeAfter() V0040Uint64NoVal`

GetResumeAfter returns the ResumeAfter field if non-nil, zero value otherwise.

### GetResumeAfterOk

`func (o *V0040Node) GetResumeAfterOk() (*V0040Uint64NoVal, bool)`

GetResumeAfterOk returns a tuple with the ResumeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAfter

`func (o *V0040Node) SetResumeAfter(v V0040Uint64NoVal)`

SetResumeAfter sets ResumeAfter field to given value.

### HasResumeAfter

`func (o *V0040Node) HasResumeAfter() bool`

HasResumeAfter returns a boolean if a field has been set.

### GetReservation

`func (o *V0040Node) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *V0040Node) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *V0040Node) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *V0040Node) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetAllocMemory

`func (o *V0040Node) GetAllocMemory() int64`

GetAllocMemory returns the AllocMemory field if non-nil, zero value otherwise.

### GetAllocMemoryOk

`func (o *V0040Node) GetAllocMemoryOk() (*int64, bool)`

GetAllocMemoryOk returns a tuple with the AllocMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocMemory

`func (o *V0040Node) SetAllocMemory(v int64)`

SetAllocMemory sets AllocMemory field to given value.

### HasAllocMemory

`func (o *V0040Node) HasAllocMemory() bool`

HasAllocMemory returns a boolean if a field has been set.

### GetAllocCpus

`func (o *V0040Node) GetAllocCpus() int32`

GetAllocCpus returns the AllocCpus field if non-nil, zero value otherwise.

### GetAllocCpusOk

`func (o *V0040Node) GetAllocCpusOk() (*int32, bool)`

GetAllocCpusOk returns a tuple with the AllocCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocCpus

`func (o *V0040Node) SetAllocCpus(v int32)`

SetAllocCpus sets AllocCpus field to given value.

### HasAllocCpus

`func (o *V0040Node) HasAllocCpus() bool`

HasAllocCpus returns a boolean if a field has been set.

### GetAllocIdleCpus

`func (o *V0040Node) GetAllocIdleCpus() int32`

GetAllocIdleCpus returns the AllocIdleCpus field if non-nil, zero value otherwise.

### GetAllocIdleCpusOk

`func (o *V0040Node) GetAllocIdleCpusOk() (*int32, bool)`

GetAllocIdleCpusOk returns a tuple with the AllocIdleCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocIdleCpus

`func (o *V0040Node) SetAllocIdleCpus(v int32)`

SetAllocIdleCpus sets AllocIdleCpus field to given value.

### HasAllocIdleCpus

`func (o *V0040Node) HasAllocIdleCpus() bool`

HasAllocIdleCpus returns a boolean if a field has been set.

### GetTresUsed

`func (o *V0040Node) GetTresUsed() string`

GetTresUsed returns the TresUsed field if non-nil, zero value otherwise.

### GetTresUsedOk

`func (o *V0040Node) GetTresUsedOk() (*string, bool)`

GetTresUsedOk returns a tuple with the TresUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresUsed

`func (o *V0040Node) SetTresUsed(v string)`

SetTresUsed sets TresUsed field to given value.

### HasTresUsed

`func (o *V0040Node) HasTresUsed() bool`

HasTresUsed returns a boolean if a field has been set.

### GetTresWeighted

`func (o *V0040Node) GetTresWeighted() float64`

GetTresWeighted returns the TresWeighted field if non-nil, zero value otherwise.

### GetTresWeightedOk

`func (o *V0040Node) GetTresWeightedOk() (*float64, bool)`

GetTresWeightedOk returns a tuple with the TresWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresWeighted

`func (o *V0040Node) SetTresWeighted(v float64)`

SetTresWeighted sets TresWeighted field to given value.

### HasTresWeighted

`func (o *V0040Node) HasTresWeighted() bool`

HasTresWeighted returns a boolean if a field has been set.

### GetSlurmdStartTime

`func (o *V0040Node) GetSlurmdStartTime() V0040Uint64NoVal`

GetSlurmdStartTime returns the SlurmdStartTime field if non-nil, zero value otherwise.

### GetSlurmdStartTimeOk

`func (o *V0040Node) GetSlurmdStartTimeOk() (*V0040Uint64NoVal, bool)`

GetSlurmdStartTimeOk returns a tuple with the SlurmdStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurmdStartTime

`func (o *V0040Node) SetSlurmdStartTime(v V0040Uint64NoVal)`

SetSlurmdStartTime sets SlurmdStartTime field to given value.

### HasSlurmdStartTime

`func (o *V0040Node) HasSlurmdStartTime() bool`

HasSlurmdStartTime returns a boolean if a field has been set.

### GetSockets

`func (o *V0040Node) GetSockets() int32`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0040Node) GetSocketsOk() (*int32, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0040Node) SetSockets(v int32)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *V0040Node) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetThreads

`func (o *V0040Node) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *V0040Node) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *V0040Node) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *V0040Node) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetTemporaryDisk

`func (o *V0040Node) GetTemporaryDisk() int32`

GetTemporaryDisk returns the TemporaryDisk field if non-nil, zero value otherwise.

### GetTemporaryDiskOk

`func (o *V0040Node) GetTemporaryDiskOk() (*int32, bool)`

GetTemporaryDiskOk returns a tuple with the TemporaryDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDisk

`func (o *V0040Node) SetTemporaryDisk(v int32)`

SetTemporaryDisk sets TemporaryDisk field to given value.

### HasTemporaryDisk

`func (o *V0040Node) HasTemporaryDisk() bool`

HasTemporaryDisk returns a boolean if a field has been set.

### GetWeight

`func (o *V0040Node) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V0040Node) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V0040Node) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V0040Node) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetTres

`func (o *V0040Node) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0040Node) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0040Node) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0040Node) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetVersion

`func (o *V0040Node) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V0040Node) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V0040Node) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V0040Node) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


