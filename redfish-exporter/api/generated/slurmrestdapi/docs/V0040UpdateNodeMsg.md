# V0040UpdateNodeMsg

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
**ResumeAfter** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Weight** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 

## Methods

### NewV0040UpdateNodeMsg

`func NewV0040UpdateNodeMsg() *V0040UpdateNodeMsg`

NewV0040UpdateNodeMsg instantiates a new V0040UpdateNodeMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040UpdateNodeMsgWithDefaults

`func NewV0040UpdateNodeMsgWithDefaults() *V0040UpdateNodeMsg`

NewV0040UpdateNodeMsgWithDefaults instantiates a new V0040UpdateNodeMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0040UpdateNodeMsg) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0040UpdateNodeMsg) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0040UpdateNodeMsg) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0040UpdateNodeMsg) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCpuBind

`func (o *V0040UpdateNodeMsg) GetCpuBind() int32`

GetCpuBind returns the CpuBind field if non-nil, zero value otherwise.

### GetCpuBindOk

`func (o *V0040UpdateNodeMsg) GetCpuBindOk() (*int32, bool)`

GetCpuBindOk returns a tuple with the CpuBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBind

`func (o *V0040UpdateNodeMsg) SetCpuBind(v int32)`

SetCpuBind sets CpuBind field to given value.

### HasCpuBind

`func (o *V0040UpdateNodeMsg) HasCpuBind() bool`

HasCpuBind returns a boolean if a field has been set.

### GetExtra

`func (o *V0040UpdateNodeMsg) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0040UpdateNodeMsg) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0040UpdateNodeMsg) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0040UpdateNodeMsg) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFeatures

`func (o *V0040UpdateNodeMsg) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0040UpdateNodeMsg) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0040UpdateNodeMsg) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0040UpdateNodeMsg) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesAct

`func (o *V0040UpdateNodeMsg) GetFeaturesAct() []string`

GetFeaturesAct returns the FeaturesAct field if non-nil, zero value otherwise.

### GetFeaturesActOk

`func (o *V0040UpdateNodeMsg) GetFeaturesActOk() (*[]string, bool)`

GetFeaturesActOk returns a tuple with the FeaturesAct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesAct

`func (o *V0040UpdateNodeMsg) SetFeaturesAct(v []string)`

SetFeaturesAct sets FeaturesAct field to given value.

### HasFeaturesAct

`func (o *V0040UpdateNodeMsg) HasFeaturesAct() bool`

HasFeaturesAct returns a boolean if a field has been set.

### GetGres

`func (o *V0040UpdateNodeMsg) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0040UpdateNodeMsg) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0040UpdateNodeMsg) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0040UpdateNodeMsg) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetAddress

`func (o *V0040UpdateNodeMsg) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V0040UpdateNodeMsg) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V0040UpdateNodeMsg) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V0040UpdateNodeMsg) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *V0040UpdateNodeMsg) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0040UpdateNodeMsg) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0040UpdateNodeMsg) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0040UpdateNodeMsg) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetName

`func (o *V0040UpdateNodeMsg) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040UpdateNodeMsg) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040UpdateNodeMsg) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040UpdateNodeMsg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *V0040UpdateNodeMsg) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0040UpdateNodeMsg) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0040UpdateNodeMsg) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0040UpdateNodeMsg) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *V0040UpdateNodeMsg) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V0040UpdateNodeMsg) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V0040UpdateNodeMsg) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V0040UpdateNodeMsg) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonUid

`func (o *V0040UpdateNodeMsg) GetReasonUid() string`

GetReasonUid returns the ReasonUid field if non-nil, zero value otherwise.

### GetReasonUidOk

`func (o *V0040UpdateNodeMsg) GetReasonUidOk() (*string, bool)`

GetReasonUidOk returns a tuple with the ReasonUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUid

`func (o *V0040UpdateNodeMsg) SetReasonUid(v string)`

SetReasonUid sets ReasonUid field to given value.

### HasReasonUid

`func (o *V0040UpdateNodeMsg) HasReasonUid() bool`

HasReasonUid returns a boolean if a field has been set.

### GetResumeAfter

`func (o *V0040UpdateNodeMsg) GetResumeAfter() V0040Uint32NoVal`

GetResumeAfter returns the ResumeAfter field if non-nil, zero value otherwise.

### GetResumeAfterOk

`func (o *V0040UpdateNodeMsg) GetResumeAfterOk() (*V0040Uint32NoVal, bool)`

GetResumeAfterOk returns a tuple with the ResumeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAfter

`func (o *V0040UpdateNodeMsg) SetResumeAfter(v V0040Uint32NoVal)`

SetResumeAfter sets ResumeAfter field to given value.

### HasResumeAfter

`func (o *V0040UpdateNodeMsg) HasResumeAfter() bool`

HasResumeAfter returns a boolean if a field has been set.

### GetWeight

`func (o *V0040UpdateNodeMsg) GetWeight() V0040Uint32NoVal`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V0040UpdateNodeMsg) GetWeightOk() (*V0040Uint32NoVal, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V0040UpdateNodeMsg) SetWeight(v V0040Uint32NoVal)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V0040UpdateNodeMsg) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


