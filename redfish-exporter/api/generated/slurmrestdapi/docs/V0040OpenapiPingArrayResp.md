# V0040OpenapiPingArrayResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pings** | [**[]V0040ControllerPing**](V0040ControllerPing.md) |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiPingArrayResp

`func NewV0040OpenapiPingArrayResp(pings []V0040ControllerPing, ) *V0040OpenapiPingArrayResp`

NewV0040OpenapiPingArrayResp instantiates a new V0040OpenapiPingArrayResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiPingArrayRespWithDefaults

`func NewV0040OpenapiPingArrayRespWithDefaults() *V0040OpenapiPingArrayResp`

NewV0040OpenapiPingArrayRespWithDefaults instantiates a new V0040OpenapiPingArrayResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPings

`func (o *V0040OpenapiPingArrayResp) GetPings() []V0040ControllerPing`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0040OpenapiPingArrayResp) GetPingsOk() (*[]V0040ControllerPing, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0040OpenapiPingArrayResp) SetPings(v []V0040ControllerPing)`

SetPings sets Pings field to given value.


### GetMeta

`func (o *V0040OpenapiPingArrayResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiPingArrayResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiPingArrayResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiPingArrayResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiPingArrayResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiPingArrayResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiPingArrayResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiPingArrayResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiPingArrayResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiPingArrayResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiPingArrayResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiPingArrayResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


