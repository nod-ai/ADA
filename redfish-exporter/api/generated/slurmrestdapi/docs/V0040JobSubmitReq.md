# V0040JobSubmitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** | batch job script | [optional] 
**Jobs** | Pointer to [**[]V0040JobDescMsg**](V0040JobDescMsg.md) |  | [optional] 
**Job** | Pointer to [**V0040JobDescMsg**](V0040JobDescMsg.md) |  | [optional] 

## Methods

### NewV0040JobSubmitReq

`func NewV0040JobSubmitReq() *V0040JobSubmitReq`

NewV0040JobSubmitReq instantiates a new V0040JobSubmitReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobSubmitReqWithDefaults

`func NewV0040JobSubmitReqWithDefaults() *V0040JobSubmitReq`

NewV0040JobSubmitReqWithDefaults instantiates a new V0040JobSubmitReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *V0040JobSubmitReq) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0040JobSubmitReq) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0040JobSubmitReq) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0040JobSubmitReq) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetJobs

`func (o *V0040JobSubmitReq) GetJobs() []V0040JobDescMsg`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0040JobSubmitReq) GetJobsOk() (*[]V0040JobDescMsg, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0040JobSubmitReq) SetJobs(v []V0040JobDescMsg)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0040JobSubmitReq) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetJob

`func (o *V0040JobSubmitReq) GetJob() V0040JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0040JobSubmitReq) GetJobOk() (*V0040JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0040JobSubmitReq) SetJob(v V0040JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0040JobSubmitReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


