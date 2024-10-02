# V0039Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorNumber** | Pointer to **int32** | Slurm internal error number | [optional] 
**Error** | Pointer to **string** | Error message | [optional] 
**Source** | Pointer to **string** | Where error occurred in the source | [optional] 
**Description** | Pointer to **string** | Explanation of cause of error | [optional] 

## Methods

### NewV0039Error

`func NewV0039Error() *V0039Error`

NewV0039Error instantiates a new V0039Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039ErrorWithDefaults

`func NewV0039ErrorWithDefaults() *V0039Error`

NewV0039ErrorWithDefaults instantiates a new V0039Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorNumber

`func (o *V0039Error) GetErrorNumber() int32`

GetErrorNumber returns the ErrorNumber field if non-nil, zero value otherwise.

### GetErrorNumberOk

`func (o *V0039Error) GetErrorNumberOk() (*int32, bool)`

GetErrorNumberOk returns a tuple with the ErrorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNumber

`func (o *V0039Error) SetErrorNumber(v int32)`

SetErrorNumber sets ErrorNumber field to given value.

### HasErrorNumber

`func (o *V0039Error) HasErrorNumber() bool`

HasErrorNumber returns a boolean if a field has been set.

### GetError

`func (o *V0039Error) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0039Error) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0039Error) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V0039Error) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSource

`func (o *V0039Error) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V0039Error) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V0039Error) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *V0039Error) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDescription

`func (o *V0039Error) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0039Error) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0039Error) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V0039Error) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


