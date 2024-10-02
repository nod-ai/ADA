# V0040OpenapiJobSubmitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**V0040JobSubmitResponseMsg**](V0040JobSubmitResponseMsg.md) |  | [optional] 
**JobId** | Pointer to **int32** | submited JobId | [optional] 
**StepId** | Pointer to **string** | submited StepID | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | job submision user message | [optional] 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiJobSubmitResponse

`func NewV0040OpenapiJobSubmitResponse() *V0040OpenapiJobSubmitResponse`

NewV0040OpenapiJobSubmitResponse instantiates a new V0040OpenapiJobSubmitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiJobSubmitResponseWithDefaults

`func NewV0040OpenapiJobSubmitResponseWithDefaults() *V0040OpenapiJobSubmitResponse`

NewV0040OpenapiJobSubmitResponseWithDefaults instantiates a new V0040OpenapiJobSubmitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *V0040OpenapiJobSubmitResponse) GetResult() V0040JobSubmitResponseMsg`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V0040OpenapiJobSubmitResponse) GetResultOk() (*V0040JobSubmitResponseMsg, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V0040OpenapiJobSubmitResponse) SetResult(v V0040JobSubmitResponseMsg)`

SetResult sets Result field to given value.

### HasResult

`func (o *V0040OpenapiJobSubmitResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetJobId

`func (o *V0040OpenapiJobSubmitResponse) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0040OpenapiJobSubmitResponse) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0040OpenapiJobSubmitResponse) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0040OpenapiJobSubmitResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0040OpenapiJobSubmitResponse) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0040OpenapiJobSubmitResponse) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0040OpenapiJobSubmitResponse) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0040OpenapiJobSubmitResponse) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *V0040OpenapiJobSubmitResponse) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *V0040OpenapiJobSubmitResponse) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *V0040OpenapiJobSubmitResponse) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *V0040OpenapiJobSubmitResponse) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.

### GetMeta

`func (o *V0040OpenapiJobSubmitResponse) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiJobSubmitResponse) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiJobSubmitResponse) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiJobSubmitResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiJobSubmitResponse) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiJobSubmitResponse) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiJobSubmitResponse) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiJobSubmitResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiJobSubmitResponse) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiJobSubmitResponse) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiJobSubmitResponse) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiJobSubmitResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


