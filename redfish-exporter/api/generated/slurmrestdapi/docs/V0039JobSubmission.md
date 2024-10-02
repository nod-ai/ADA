# V0039JobSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** | Executable script (full contents) to run in batch step for all job components | [optional] 
**Job** | Pointer to [**V0039JobDescMsg**](V0039JobDescMsg.md) |  | [optional] 
**Jobs** | Pointer to [**[]V0039JobDescMsg**](V0039JobDescMsg.md) |  | [optional] 

## Methods

### NewV0039JobSubmission

`func NewV0039JobSubmission() *V0039JobSubmission`

NewV0039JobSubmission instantiates a new V0039JobSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobSubmissionWithDefaults

`func NewV0039JobSubmissionWithDefaults() *V0039JobSubmission`

NewV0039JobSubmissionWithDefaults instantiates a new V0039JobSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *V0039JobSubmission) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0039JobSubmission) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0039JobSubmission) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0039JobSubmission) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetJob

`func (o *V0039JobSubmission) GetJob() V0039JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0039JobSubmission) GetJobOk() (*V0039JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0039JobSubmission) SetJob(v V0039JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0039JobSubmission) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobs

`func (o *V0039JobSubmission) GetJobs() []V0039JobDescMsg`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0039JobSubmission) GetJobsOk() (*[]V0039JobDescMsg, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0039JobSubmission) SetJobs(v []V0039JobDescMsg)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0039JobSubmission) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


