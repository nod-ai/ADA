# V0039JobUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**Results** | Pointer to [**[]V0039JobArrayResponseMsgInner**](V0039JobArrayResponseMsgInner.md) | Result per ArrayJob | [optional] 

## Methods

### NewV0039JobUpdateResponse

`func NewV0039JobUpdateResponse() *V0039JobUpdateResponse`

NewV0039JobUpdateResponse instantiates a new V0039JobUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobUpdateResponseWithDefaults

`func NewV0039JobUpdateResponseWithDefaults() *V0039JobUpdateResponse`

NewV0039JobUpdateResponseWithDefaults instantiates a new V0039JobUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039JobUpdateResponse) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039JobUpdateResponse) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039JobUpdateResponse) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039JobUpdateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039JobUpdateResponse) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039JobUpdateResponse) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039JobUpdateResponse) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039JobUpdateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039JobUpdateResponse) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039JobUpdateResponse) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039JobUpdateResponse) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039JobUpdateResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetResults

`func (o *V0039JobUpdateResponse) GetResults() []V0039JobArrayResponseMsgInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V0039JobUpdateResponse) GetResultsOk() (*[]V0039JobArrayResponseMsgInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V0039JobUpdateResponse) SetResults(v []V0039JobArrayResponseMsgInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *V0039JobUpdateResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


