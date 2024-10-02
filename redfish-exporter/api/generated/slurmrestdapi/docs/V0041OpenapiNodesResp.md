# V0041OpenapiNodesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]V0041NodesInner**](V0041NodesInner.md) |  | 
**LastUpdate** | [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiNodesResp

`func NewV0041OpenapiNodesResp(nodes []V0041NodesInner, lastUpdate V0041Uint64NoValStruct, ) *V0041OpenapiNodesResp`

NewV0041OpenapiNodesResp instantiates a new V0041OpenapiNodesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiNodesRespWithDefaults

`func NewV0041OpenapiNodesRespWithDefaults() *V0041OpenapiNodesResp`

NewV0041OpenapiNodesRespWithDefaults instantiates a new V0041OpenapiNodesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0041OpenapiNodesResp) GetNodes() []V0041NodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiNodesResp) GetNodesOk() (*[]V0041NodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiNodesResp) SetNodes(v []V0041NodesInner)`

SetNodes sets Nodes field to given value.


### GetLastUpdate

`func (o *V0041OpenapiNodesResp) GetLastUpdate() V0041Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0041OpenapiNodesResp) GetLastUpdateOk() (*V0041Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0041OpenapiNodesResp) SetLastUpdate(v V0041Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0041OpenapiNodesResp) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiNodesResp) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiNodesResp) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiNodesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiNodesResp) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiNodesResp) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiNodesResp) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiNodesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiNodesResp) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiNodesResp) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiNodesResp) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiNodesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


