# V0041ProcessExitCodeVerbose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **[]string** | Status given by return code | [optional] 
**ReturnCode** | Pointer to [**V0041Uint32NoValStruct**](V0041Uint32NoValStruct.md) |  | [optional] 
**Signal** | Pointer to [**V0041ProcessExitCodeVerboseSignal**](V0041ProcessExitCodeVerboseSignal.md) |  | [optional] 

## Methods

### NewV0041ProcessExitCodeVerbose

`func NewV0041ProcessExitCodeVerbose() *V0041ProcessExitCodeVerbose`

NewV0041ProcessExitCodeVerbose instantiates a new V0041ProcessExitCodeVerbose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041ProcessExitCodeVerboseWithDefaults

`func NewV0041ProcessExitCodeVerboseWithDefaults() *V0041ProcessExitCodeVerbose`

NewV0041ProcessExitCodeVerboseWithDefaults instantiates a new V0041ProcessExitCodeVerbose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V0041ProcessExitCodeVerbose) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0041ProcessExitCodeVerbose) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0041ProcessExitCodeVerbose) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0041ProcessExitCodeVerbose) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnCode

`func (o *V0041ProcessExitCodeVerbose) GetReturnCode() V0041Uint32NoValStruct`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *V0041ProcessExitCodeVerbose) GetReturnCodeOk() (*V0041Uint32NoValStruct, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *V0041ProcessExitCodeVerbose) SetReturnCode(v V0041Uint32NoValStruct)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *V0041ProcessExitCodeVerbose) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetSignal

`func (o *V0041ProcessExitCodeVerbose) GetSignal() V0041ProcessExitCodeVerboseSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V0041ProcessExitCodeVerbose) GetSignalOk() (*V0041ProcessExitCodeVerboseSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V0041ProcessExitCodeVerbose) SetSignal(v V0041ProcessExitCodeVerboseSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V0041ProcessExitCodeVerbose) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


