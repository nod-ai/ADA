# V0041JobSubmitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** | batch job script | [optional] 
**Jobs** | Pointer to [**[]V0041JobDescMsg**](V0041JobDescMsg.md) |  | [optional] 
**Job** | Pointer to [**V0041JobDescMsg**](V0041JobDescMsg.md) |  | [optional] 

## Methods

### NewV0041JobSubmitReq

`func NewV0041JobSubmitReq() *V0041JobSubmitReq`

NewV0041JobSubmitReq instantiates a new V0041JobSubmitReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041JobSubmitReqWithDefaults

`func NewV0041JobSubmitReqWithDefaults() *V0041JobSubmitReq`

NewV0041JobSubmitReqWithDefaults instantiates a new V0041JobSubmitReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *V0041JobSubmitReq) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0041JobSubmitReq) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0041JobSubmitReq) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0041JobSubmitReq) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetJobs

`func (o *V0041JobSubmitReq) GetJobs() []V0041JobDescMsg`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0041JobSubmitReq) GetJobsOk() (*[]V0041JobDescMsg, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0041JobSubmitReq) SetJobs(v []V0041JobDescMsg)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0041JobSubmitReq) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetJob

`func (o *V0041JobSubmitReq) GetJob() V0041JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0041JobSubmitReq) GetJobOk() (*V0041JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0041JobSubmitReq) SetJob(v V0041JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0041JobSubmitReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


