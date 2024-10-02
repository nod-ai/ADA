# V0039PartitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**Partitions** | Pointer to [**[]V0039PartitionInfo**](V0039PartitionInfo.md) |  | [optional] 

## Methods

### NewV0039PartitionsResponse

`func NewV0039PartitionsResponse() *V0039PartitionsResponse`

NewV0039PartitionsResponse instantiates a new V0039PartitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039PartitionsResponseWithDefaults

`func NewV0039PartitionsResponseWithDefaults() *V0039PartitionsResponse`

NewV0039PartitionsResponseWithDefaults instantiates a new V0039PartitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039PartitionsResponse) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039PartitionsResponse) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039PartitionsResponse) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039PartitionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039PartitionsResponse) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039PartitionsResponse) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039PartitionsResponse) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039PartitionsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039PartitionsResponse) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039PartitionsResponse) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039PartitionsResponse) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039PartitionsResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetPartitions

`func (o *V0039PartitionsResponse) GetPartitions() []V0039PartitionInfo`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0039PartitionsResponse) GetPartitionsOk() (*[]V0039PartitionInfo, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0039PartitionsResponse) SetPartitions(v []V0039PartitionInfo)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0039PartitionsResponse) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


