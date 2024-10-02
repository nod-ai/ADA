# V0040OpenapiJobPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]V0040JobArrayResponseMsgEntry**](V0040JobArrayResponseMsgEntry.md) |  | [optional] 
**JobId** | Pointer to **string** | First updated JobId - Use results instead | [optional] 
**StepId** | Pointer to **string** | First updated StepID - Use results instead | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | First updated Job submision user message - Use results instead | [optional] 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiJobPostResponse

`func NewV0040OpenapiJobPostResponse() *V0040OpenapiJobPostResponse`

NewV0040OpenapiJobPostResponse instantiates a new V0040OpenapiJobPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiJobPostResponseWithDefaults

`func NewV0040OpenapiJobPostResponseWithDefaults() *V0040OpenapiJobPostResponse`

NewV0040OpenapiJobPostResponseWithDefaults instantiates a new V0040OpenapiJobPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V0040OpenapiJobPostResponse) GetResults() []V0040JobArrayResponseMsgEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V0040OpenapiJobPostResponse) GetResultsOk() (*[]V0040JobArrayResponseMsgEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V0040OpenapiJobPostResponse) SetResults(v []V0040JobArrayResponseMsgEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *V0040OpenapiJobPostResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetJobId

`func (o *V0040OpenapiJobPostResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0040OpenapiJobPostResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0040OpenapiJobPostResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0040OpenapiJobPostResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0040OpenapiJobPostResponse) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0040OpenapiJobPostResponse) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0040OpenapiJobPostResponse) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0040OpenapiJobPostResponse) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *V0040OpenapiJobPostResponse) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *V0040OpenapiJobPostResponse) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *V0040OpenapiJobPostResponse) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *V0040OpenapiJobPostResponse) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.

### GetMeta

`func (o *V0040OpenapiJobPostResponse) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiJobPostResponse) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiJobPostResponse) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiJobPostResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiJobPostResponse) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiJobPostResponse) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiJobPostResponse) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiJobPostResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiJobPostResponse) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiJobPostResponse) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiJobPostResponse) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiJobPostResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


