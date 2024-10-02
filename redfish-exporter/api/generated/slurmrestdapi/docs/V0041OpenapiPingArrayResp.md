# V0041OpenapiPingArrayResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pings** | [**[]V0041ControllerPingArrayInner**](V0041ControllerPingArrayInner.md) |  | 
**Meta** | Pointer to [**V0041OpenapiMeta**](V0041OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiErrorsInner**](V0041OpenapiErrorsInner.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiWarningsInner**](V0041OpenapiWarningsInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiPingArrayResp

`func NewV0041OpenapiPingArrayResp(pings []V0041ControllerPingArrayInner, ) *V0041OpenapiPingArrayResp`

NewV0041OpenapiPingArrayResp instantiates a new V0041OpenapiPingArrayResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiPingArrayRespWithDefaults

`func NewV0041OpenapiPingArrayRespWithDefaults() *V0041OpenapiPingArrayResp`

NewV0041OpenapiPingArrayRespWithDefaults instantiates a new V0041OpenapiPingArrayResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPings

`func (o *V0041OpenapiPingArrayResp) GetPings() []V0041ControllerPingArrayInner`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0041OpenapiPingArrayResp) GetPingsOk() (*[]V0041ControllerPingArrayInner, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0041OpenapiPingArrayResp) SetPings(v []V0041ControllerPingArrayInner)`

SetPings sets Pings field to given value.


### GetMeta

`func (o *V0041OpenapiPingArrayResp) GetMeta() V0041OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiPingArrayResp) GetMetaOk() (*V0041OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiPingArrayResp) SetMeta(v V0041OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiPingArrayResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiPingArrayResp) GetErrors() []V0041OpenapiErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiPingArrayResp) GetErrorsOk() (*[]V0041OpenapiErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiPingArrayResp) SetErrors(v []V0041OpenapiErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiPingArrayResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiPingArrayResp) GetWarnings() []V0041OpenapiWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiPingArrayResp) GetWarningsOk() (*[]V0041OpenapiWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiPingArrayResp) SetWarnings(v []V0041OpenapiWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiPingArrayResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


