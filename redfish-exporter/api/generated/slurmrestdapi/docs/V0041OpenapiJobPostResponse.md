# V0041OpenapiJobPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]V0041JobArrayResponseArrayInner**](V0041JobArrayResponseArrayInner.md) |  | [optional] 
**JobId** | Pointer to **string** | First updated JobId - Use results instead | [optional] 
**StepId** | Pointer to **string** | First updated StepID - Use results instead | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | First updated Job submision user message - Use results instead | [optional] 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiJobPostResponse

`func NewV0041OpenapiJobPostResponse() *V0041OpenapiJobPostResponse`

NewV0041OpenapiJobPostResponse instantiates a new V0041OpenapiJobPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiJobPostResponseWithDefaults

`func NewV0041OpenapiJobPostResponseWithDefaults() *V0041OpenapiJobPostResponse`

NewV0041OpenapiJobPostResponseWithDefaults instantiates a new V0041OpenapiJobPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V0041OpenapiJobPostResponse) GetResults() []V0041JobArrayResponseArrayInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V0041OpenapiJobPostResponse) GetResultsOk() (*[]V0041JobArrayResponseArrayInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V0041OpenapiJobPostResponse) SetResults(v []V0041JobArrayResponseArrayInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *V0041OpenapiJobPostResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetJobId

`func (o *V0041OpenapiJobPostResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0041OpenapiJobPostResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0041OpenapiJobPostResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0041OpenapiJobPostResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0041OpenapiJobPostResponse) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0041OpenapiJobPostResponse) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0041OpenapiJobPostResponse) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0041OpenapiJobPostResponse) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *V0041OpenapiJobPostResponse) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *V0041OpenapiJobPostResponse) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *V0041OpenapiJobPostResponse) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *V0041OpenapiJobPostResponse) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.

### GetMeta

`func (o *V0041OpenapiJobPostResponse) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiJobPostResponse) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiJobPostResponse) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiJobPostResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiJobPostResponse) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiJobPostResponse) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiJobPostResponse) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiJobPostResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiJobPostResponse) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiJobPostResponse) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiJobPostResponse) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiJobPostResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


