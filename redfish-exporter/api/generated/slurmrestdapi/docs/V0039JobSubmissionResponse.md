# V0039JobSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**JobId** | Pointer to **int32** | new job ID | [optional] 
**StepId** | Pointer to **string** | new job step ID | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | Message to user from job_submit plugin | [optional] 

## Methods

### NewV0039JobSubmissionResponse

`func NewV0039JobSubmissionResponse() *V0039JobSubmissionResponse`

NewV0039JobSubmissionResponse instantiates a new V0039JobSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobSubmissionResponseWithDefaults

`func NewV0039JobSubmissionResponseWithDefaults() *V0039JobSubmissionResponse`

NewV0039JobSubmissionResponseWithDefaults instantiates a new V0039JobSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039JobSubmissionResponse) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039JobSubmissionResponse) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039JobSubmissionResponse) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039JobSubmissionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039JobSubmissionResponse) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039JobSubmissionResponse) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039JobSubmissionResponse) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039JobSubmissionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039JobSubmissionResponse) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039JobSubmissionResponse) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039JobSubmissionResponse) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039JobSubmissionResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetJobId

`func (o *V0039JobSubmissionResponse) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0039JobSubmissionResponse) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0039JobSubmissionResponse) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0039JobSubmissionResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0039JobSubmissionResponse) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0039JobSubmissionResponse) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0039JobSubmissionResponse) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0039JobSubmissionResponse) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *V0039JobSubmissionResponse) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *V0039JobSubmissionResponse) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *V0039JobSubmissionResponse) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *V0039JobSubmissionResponse) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


