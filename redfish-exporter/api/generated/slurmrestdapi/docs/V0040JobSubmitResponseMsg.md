# V0040JobSubmitResponseMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** |  | [optional] 
**StepId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**JobSubmitUserMsg** | Pointer to **string** |  | [optional] 

## Methods

### NewV0040JobSubmitResponseMsg

`func NewV0040JobSubmitResponseMsg() *V0040JobSubmitResponseMsg`

NewV0040JobSubmitResponseMsg instantiates a new V0040JobSubmitResponseMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobSubmitResponseMsgWithDefaults

`func NewV0040JobSubmitResponseMsgWithDefaults() *V0040JobSubmitResponseMsg`

NewV0040JobSubmitResponseMsgWithDefaults instantiates a new V0040JobSubmitResponseMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *V0040JobSubmitResponseMsg) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0040JobSubmitResponseMsg) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0040JobSubmitResponseMsg) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0040JobSubmitResponseMsg) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0040JobSubmitResponseMsg) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0040JobSubmitResponseMsg) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0040JobSubmitResponseMsg) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0040JobSubmitResponseMsg) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetErrorCode

`func (o *V0040JobSubmitResponseMsg) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *V0040JobSubmitResponseMsg) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *V0040JobSubmitResponseMsg) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *V0040JobSubmitResponseMsg) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetError

`func (o *V0040JobSubmitResponseMsg) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0040JobSubmitResponseMsg) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0040JobSubmitResponseMsg) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V0040JobSubmitResponseMsg) HasError() bool`

HasError returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *V0040JobSubmitResponseMsg) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *V0040JobSubmitResponseMsg) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *V0040JobSubmitResponseMsg) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *V0040JobSubmitResponseMsg) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


