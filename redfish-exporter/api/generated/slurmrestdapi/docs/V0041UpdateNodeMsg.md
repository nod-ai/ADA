# V0041UpdateNodeMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | arbitrary comment | [optional] 
**CpuBind** | Pointer to **int32** | default CPU binding type | [optional] 
**Extra** | Pointer to **string** | arbitrary string | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**FeaturesAct** | Pointer to **[]string** |  | [optional] 
**Gres** | Pointer to **string** | new generic resources for node | [optional] 
**Address** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **[]string** | assign new node state | [optional] 
**Reason** | Pointer to **string** | reason for node being DOWN or DRAINING | [optional] 
**ReasonUid** | Pointer to **string** | user ID of sending (needed if user root is sending message) | [optional] 
**ResumeAfter** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Weight** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 

## Methods

### NewV0041UpdateNodeMsg

`func NewV0041UpdateNodeMsg() *V0041UpdateNodeMsg`

NewV0041UpdateNodeMsg instantiates a new V0041UpdateNodeMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041UpdateNodeMsgWithDefaults

`func NewV0041UpdateNodeMsgWithDefaults() *V0041UpdateNodeMsg`

NewV0041UpdateNodeMsgWithDefaults instantiates a new V0041UpdateNodeMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0041UpdateNodeMsg) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0041UpdateNodeMsg) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0041UpdateNodeMsg) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0041UpdateNodeMsg) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCpuBind

`func (o *V0041UpdateNodeMsg) GetCpuBind() int32`

GetCpuBind returns the CpuBind field if non-nil, zero value otherwise.

### GetCpuBindOk

`func (o *V0041UpdateNodeMsg) GetCpuBindOk() (*int32, bool)`

GetCpuBindOk returns a tuple with the CpuBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBind

`func (o *V0041UpdateNodeMsg) SetCpuBind(v int32)`

SetCpuBind sets CpuBind field to given value.

### HasCpuBind

`func (o *V0041UpdateNodeMsg) HasCpuBind() bool`

HasCpuBind returns a boolean if a field has been set.

### GetExtra

`func (o *V0041UpdateNodeMsg) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0041UpdateNodeMsg) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0041UpdateNodeMsg) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0041UpdateNodeMsg) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFeatures

`func (o *V0041UpdateNodeMsg) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0041UpdateNodeMsg) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0041UpdateNodeMsg) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0041UpdateNodeMsg) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesAct

`func (o *V0041UpdateNodeMsg) GetFeaturesAct() []string`

GetFeaturesAct returns the FeaturesAct field if non-nil, zero value otherwise.

### GetFeaturesActOk

`func (o *V0041UpdateNodeMsg) GetFeaturesActOk() (*[]string, bool)`

GetFeaturesActOk returns a tuple with the FeaturesAct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesAct

`func (o *V0041UpdateNodeMsg) SetFeaturesAct(v []string)`

SetFeaturesAct sets FeaturesAct field to given value.

### HasFeaturesAct

`func (o *V0041UpdateNodeMsg) HasFeaturesAct() bool`

HasFeaturesAct returns a boolean if a field has been set.

### GetGres

`func (o *V0041UpdateNodeMsg) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0041UpdateNodeMsg) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0041UpdateNodeMsg) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0041UpdateNodeMsg) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetAddress

`func (o *V0041UpdateNodeMsg) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V0041UpdateNodeMsg) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V0041UpdateNodeMsg) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V0041UpdateNodeMsg) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *V0041UpdateNodeMsg) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0041UpdateNodeMsg) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0041UpdateNodeMsg) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0041UpdateNodeMsg) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetName

`func (o *V0041UpdateNodeMsg) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041UpdateNodeMsg) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041UpdateNodeMsg) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041UpdateNodeMsg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *V0041UpdateNodeMsg) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0041UpdateNodeMsg) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0041UpdateNodeMsg) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0041UpdateNodeMsg) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *V0041UpdateNodeMsg) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V0041UpdateNodeMsg) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V0041UpdateNodeMsg) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V0041UpdateNodeMsg) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonUid

`func (o *V0041UpdateNodeMsg) GetReasonUid() string`

GetReasonUid returns the ReasonUid field if non-nil, zero value otherwise.

### GetReasonUidOk

`func (o *V0041UpdateNodeMsg) GetReasonUidOk() (*string, bool)`

GetReasonUidOk returns a tuple with the ReasonUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUid

`func (o *V0041UpdateNodeMsg) SetReasonUid(v string)`

SetReasonUid sets ReasonUid field to given value.

### HasReasonUid

`func (o *V0041UpdateNodeMsg) HasReasonUid() bool`

HasReasonUid returns a boolean if a field has been set.

### GetResumeAfter

`func (o *V0041UpdateNodeMsg) GetResumeAfter() V0041Uint32NoValStruct`

GetResumeAfter returns the ResumeAfter field if non-nil, zero value otherwise.

### GetResumeAfterOk

`func (o *V0041UpdateNodeMsg) GetResumeAfterOk() (*V0041Uint32NoValStruct, bool)`

GetResumeAfterOk returns a tuple with the ResumeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAfter

`func (o *V0041UpdateNodeMsg) SetResumeAfter(v V0041Uint32NoValStruct)`

SetResumeAfter sets ResumeAfter field to given value.

### HasResumeAfter

`func (o *V0041UpdateNodeMsg) HasResumeAfter() bool`

HasResumeAfter returns a boolean if a field has been set.

### GetWeight

`func (o *V0041UpdateNodeMsg) GetWeight() V0041Uint32NoValStruct`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V0041UpdateNodeMsg) GetWeightOk() (*V0041Uint32NoValStruct, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V0041UpdateNodeMsg) SetWeight(v V0041Uint32NoValStruct)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V0041UpdateNodeMsg) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


