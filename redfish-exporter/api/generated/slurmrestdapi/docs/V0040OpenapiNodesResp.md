# V0040OpenapiNodesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]V0040Node**](V0040Node.md) |  | 
**LastUpdate** | [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiNodesResp

`func NewV0040OpenapiNodesResp(nodes []V0040Node, lastUpdate V0040Uint64NoVal, ) *V0040OpenapiNodesResp`

NewV0040OpenapiNodesResp instantiates a new V0040OpenapiNodesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiNodesRespWithDefaults

`func NewV0040OpenapiNodesRespWithDefaults() *V0040OpenapiNodesResp`

NewV0040OpenapiNodesRespWithDefaults instantiates a new V0040OpenapiNodesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0040OpenapiNodesResp) GetNodes() []V0040Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0040OpenapiNodesResp) GetNodesOk() (*[]V0040Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0040OpenapiNodesResp) SetNodes(v []V0040Node)`

SetNodes sets Nodes field to given value.


### GetLastUpdate

`func (o *V0040OpenapiNodesResp) GetLastUpdate() V0040Uint64NoVal`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0040OpenapiNodesResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0040OpenapiNodesResp) SetLastUpdate(v V0040Uint64NoVal)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0040OpenapiNodesResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiNodesResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiNodesResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiNodesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiNodesResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiNodesResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiNodesResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiNodesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiNodesResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiNodesResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiNodesResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiNodesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


