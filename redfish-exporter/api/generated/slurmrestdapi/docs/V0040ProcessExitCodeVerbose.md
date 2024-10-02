# V0040ProcessExitCodeVerbose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **[]string** | Status given by return code | [optional] 
**ReturnCode** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Signal** | Pointer to [**V0040ProcessExitCodeVerboseSignal**](V0040ProcessExitCodeVerboseSignal.md) |  | [optional] 

## Methods

### NewV0040ProcessExitCodeVerbose

`func NewV0040ProcessExitCodeVerbose() *V0040ProcessExitCodeVerbose`

NewV0040ProcessExitCodeVerbose instantiates a new V0040ProcessExitCodeVerbose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040ProcessExitCodeVerboseWithDefaults

`func NewV0040ProcessExitCodeVerboseWithDefaults() *V0040ProcessExitCodeVerbose`

NewV0040ProcessExitCodeVerboseWithDefaults instantiates a new V0040ProcessExitCodeVerbose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V0040ProcessExitCodeVerbose) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0040ProcessExitCodeVerbose) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0040ProcessExitCodeVerbose) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0040ProcessExitCodeVerbose) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnCode

`func (o *V0040ProcessExitCodeVerbose) GetReturnCode() V0040Uint32NoVal`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *V0040ProcessExitCodeVerbose) GetReturnCodeOk() (*V0040Uint32NoVal, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *V0040ProcessExitCodeVerbose) SetReturnCode(v V0040Uint32NoVal)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *V0040ProcessExitCodeVerbose) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetSignal

`func (o *V0040ProcessExitCodeVerbose) GetSignal() V0040ProcessExitCodeVerboseSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V0040ProcessExitCodeVerbose) GetSignalOk() (*V0040ProcessExitCodeVerboseSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V0040ProcessExitCodeVerbose) SetSignal(v V0040ProcessExitCodeVerboseSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V0040ProcessExitCodeVerbose) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


