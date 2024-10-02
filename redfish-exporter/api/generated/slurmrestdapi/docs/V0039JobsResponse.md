# V0039JobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**Jobs** | Pointer to [**[]V0039JobInfo**](V0039JobInfo.md) |  | [optional] 

## Methods

### NewV0039JobsResponse

`func NewV0039JobsResponse() *V0039JobsResponse`

NewV0039JobsResponse instantiates a new V0039JobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobsResponseWithDefaults

`func NewV0039JobsResponseWithDefaults() *V0039JobsResponse`

NewV0039JobsResponseWithDefaults instantiates a new V0039JobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039JobsResponse) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039JobsResponse) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039JobsResponse) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039JobsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039JobsResponse) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039JobsResponse) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039JobsResponse) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039JobsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039JobsResponse) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039JobsResponse) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039JobsResponse) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039JobsResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetJobs

`func (o *V0039JobsResponse) GetJobs() []V0039JobInfo`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0039JobsResponse) GetJobsOk() (*[]V0039JobInfo, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0039JobsResponse) SetJobs(v []V0039JobInfo)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0039JobsResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


