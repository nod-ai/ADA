# V0041ReservationInfoMsgInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **string** |  | [optional] 
**BurstBuffer** | Pointer to **string** |  | [optional] 
**CoreCount** | Pointer to **int32** |  | [optional] 
**CoreSpecializations** | Pointer to [**[]V0041ReservationInfoCoreSpecInner**](V0041ReservationInfoCoreSpecInner.md) |  | [optional] 
**EndTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **string** |  | [optional] 
**Licenses** | Pointer to **string** |  | [optional] 
**MaxStartDelay** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**NodeList** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**PurgeCompleted** | Pointer to [**V0041ReservationInfoMsgInnerPurgeCompleted**](V0041ReservationInfoMsgInnerPurgeCompleted.md) |  | [optional] 
**StartTime** | Pointer to [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | [optional] 
**Watts** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Tres** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **string** |  | [optional] 

## Methods

### NewV0041ReservationInfoMsgInner

`func NewV0041ReservationInfoMsgInner() *V0041ReservationInfoMsgInner`

NewV0041ReservationInfoMsgInner instantiates a new V0041ReservationInfoMsgInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041ReservationInfoMsgInnerWithDefaults

`func NewV0041ReservationInfoMsgInnerWithDefaults() *V0041ReservationInfoMsgInner`

NewV0041ReservationInfoMsgInnerWithDefaults instantiates a new V0041ReservationInfoMsgInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0041ReservationInfoMsgInner) GetAccounts() string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041ReservationInfoMsgInner) GetAccountsOk() (*string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041ReservationInfoMsgInner) SetAccounts(v string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0041ReservationInfoMsgInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0041ReservationInfoMsgInner) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0041ReservationInfoMsgInner) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0041ReservationInfoMsgInner) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0041ReservationInfoMsgInner) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetCoreCount

`func (o *V0041ReservationInfoMsgInner) GetCoreCount() int32`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *V0041ReservationInfoMsgInner) GetCoreCountOk() (*int32, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *V0041ReservationInfoMsgInner) SetCoreCount(v int32)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *V0041ReservationInfoMsgInner) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetCoreSpecializations

`func (o *V0041ReservationInfoMsgInner) GetCoreSpecializations() []V0041ReservationInfoCoreSpecInner`

GetCoreSpecializations returns the CoreSpecializations field if non-nil, zero value otherwise.

### GetCoreSpecializationsOk

`func (o *V0041ReservationInfoMsgInner) GetCoreSpecializationsOk() (*[]V0041ReservationInfoCoreSpecInner, bool)`

GetCoreSpecializationsOk returns a tuple with the CoreSpecializations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecializations

`func (o *V0041ReservationInfoMsgInner) SetCoreSpecializations(v []V0041ReservationInfoCoreSpecInner)`

SetCoreSpecializations sets CoreSpecializations field to given value.

### HasCoreSpecializations

`func (o *V0041ReservationInfoMsgInner) HasCoreSpecializations() bool`

HasCoreSpecializations returns a boolean if a field has been set.

### GetEndTime

`func (o *V0041ReservationInfoMsgInner) GetEndTime() V0041Uint64NoValStruct`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0041ReservationInfoMsgInner) GetEndTimeOk() (*V0041Uint64NoValStruct, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0041ReservationInfoMsgInner) SetEndTime(v V0041Uint64NoValStruct)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0041ReservationInfoMsgInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFeatures

`func (o *V0041ReservationInfoMsgInner) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0041ReservationInfoMsgInner) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0041ReservationInfoMsgInner) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0041ReservationInfoMsgInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFlags

`func (o *V0041ReservationInfoMsgInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041ReservationInfoMsgInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041ReservationInfoMsgInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041ReservationInfoMsgInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroups

`func (o *V0041ReservationInfoMsgInner) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0041ReservationInfoMsgInner) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0041ReservationInfoMsgInner) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0041ReservationInfoMsgInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLicenses

`func (o *V0041ReservationInfoMsgInner) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0041ReservationInfoMsgInner) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0041ReservationInfoMsgInner) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0041ReservationInfoMsgInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMaxStartDelay

`func (o *V0041ReservationInfoMsgInner) GetMaxStartDelay() int32`

GetMaxStartDelay returns the MaxStartDelay field if non-nil, zero value otherwise.

### GetMaxStartDelayOk

`func (o *V0041ReservationInfoMsgInner) GetMaxStartDelayOk() (*int32, bool)`

GetMaxStartDelayOk returns a tuple with the MaxStartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDelay

`func (o *V0041ReservationInfoMsgInner) SetMaxStartDelay(v int32)`

SetMaxStartDelay sets MaxStartDelay field to given value.

### HasMaxStartDelay

`func (o *V0041ReservationInfoMsgInner) HasMaxStartDelay() bool`

HasMaxStartDelay returns a boolean if a field has been set.

### GetName

`func (o *V0041ReservationInfoMsgInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041ReservationInfoMsgInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041ReservationInfoMsgInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041ReservationInfoMsgInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0041ReservationInfoMsgInner) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0041ReservationInfoMsgInner) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0041ReservationInfoMsgInner) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0041ReservationInfoMsgInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeList

`func (o *V0041ReservationInfoMsgInner) GetNodeList() string`

GetNodeList returns the NodeList field if non-nil, zero value otherwise.

### GetNodeListOk

`func (o *V0041ReservationInfoMsgInner) GetNodeListOk() (*string, bool)`

GetNodeListOk returns a tuple with the NodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeList

`func (o *V0041ReservationInfoMsgInner) SetNodeList(v string)`

SetNodeList sets NodeList field to given value.

### HasNodeList

`func (o *V0041ReservationInfoMsgInner) HasNodeList() bool`

HasNodeList returns a boolean if a field has been set.

### GetPartition

`func (o *V0041ReservationInfoMsgInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041ReservationInfoMsgInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041ReservationInfoMsgInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041ReservationInfoMsgInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPurgeCompleted

`func (o *V0041ReservationInfoMsgInner) GetPurgeCompleted() V0041ReservationInfoMsgInnerPurgeCompleted`

GetPurgeCompleted returns the PurgeCompleted field if non-nil, zero value otherwise.

### GetPurgeCompletedOk

`func (o *V0041ReservationInfoMsgInner) GetPurgeCompletedOk() (*V0041ReservationInfoMsgInnerPurgeCompleted, bool)`

GetPurgeCompletedOk returns a tuple with the PurgeCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeCompleted

`func (o *V0041ReservationInfoMsgInner) SetPurgeCompleted(v V0041ReservationInfoMsgInnerPurgeCompleted)`

SetPurgeCompleted sets PurgeCompleted field to given value.

### HasPurgeCompleted

`func (o *V0041ReservationInfoMsgInner) HasPurgeCompleted() bool`

HasPurgeCompleted returns a boolean if a field has been set.

### GetStartTime

`func (o *V0041ReservationInfoMsgInner) GetStartTime() V0041Uint64NoValStruct`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0041ReservationInfoMsgInner) GetStartTimeOk() (*V0041Uint64NoValStruct, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0041ReservationInfoMsgInner) SetStartTime(v V0041Uint64NoValStruct)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0041ReservationInfoMsgInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWatts

`func (o *V0041ReservationInfoMsgInner) GetWatts() V0041Uint32NoValStruct`

GetWatts returns the Watts field if non-nil, zero value otherwise.

### GetWattsOk

`func (o *V0041ReservationInfoMsgInner) GetWattsOk() (*V0041Uint32NoValStruct, bool)`

GetWattsOk returns a tuple with the Watts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatts

`func (o *V0041ReservationInfoMsgInner) SetWatts(v V0041Uint32NoValStruct)`

SetWatts sets Watts field to given value.

### HasWatts

`func (o *V0041ReservationInfoMsgInner) HasWatts() bool`

HasWatts returns a boolean if a field has been set.

### GetTres

`func (o *V0041ReservationInfoMsgInner) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041ReservationInfoMsgInner) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041ReservationInfoMsgInner) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041ReservationInfoMsgInner) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUsers

`func (o *V0041ReservationInfoMsgInner) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0041ReservationInfoMsgInner) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0041ReservationInfoMsgInner) SetUsers(v string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0041ReservationInfoMsgInner) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


