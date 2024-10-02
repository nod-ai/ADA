# V0039NodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0039Meta**](V0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0039Error**](V0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]V0039Warning**](V0039Warning.md) | Slurm warnings | [optional] 
**Nodes** | Pointer to [**[]V0039Node**](V0039Node.md) |  | [optional] 

## Methods

### NewV0039NodesResponse

`func NewV0039NodesResponse() *V0039NodesResponse`

NewV0039NodesResponse instantiates a new V0039NodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039NodesResponseWithDefaults

`func NewV0039NodesResponseWithDefaults() *V0039NodesResponse`

NewV0039NodesResponseWithDefaults instantiates a new V0039NodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0039NodesResponse) GetMeta() V0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0039NodesResponse) GetMetaOk() (*V0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0039NodesResponse) SetMeta(v V0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0039NodesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0039NodesResponse) GetErrors() []V0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0039NodesResponse) GetErrorsOk() (*[]V0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0039NodesResponse) SetErrors(v []V0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0039NodesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0039NodesResponse) GetWarnings() []V0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0039NodesResponse) GetWarningsOk() (*[]V0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0039NodesResponse) SetWarnings(v []V0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0039NodesResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetNodes

`func (o *V0039NodesResponse) GetNodes() []V0039Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0039NodesResponse) GetNodesOk() (*[]V0039Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0039NodesResponse) SetNodes(v []V0039Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0039NodesResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


