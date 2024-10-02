# V0040Float64NoVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **bool** | True if number has been set. False if number is unset | [optional] [default to false]
**Infinite** | Pointer to **bool** | True if number has been set to infinite. \&quot;set\&quot; and \&quot;number\&quot; will be ignored. | [optional] 
**Number** | Pointer to **float64** | If set is True the number will be set with value. Otherwise ignore number contents. | [optional] 

## Methods

### NewV0040Float64NoVal

`func NewV0040Float64NoVal() *V0040Float64NoVal`

NewV0040Float64NoVal instantiates a new V0040Float64NoVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040Float64NoValWithDefaults

`func NewV0040Float64NoValWithDefaults() *V0040Float64NoVal`

NewV0040Float64NoValWithDefaults instantiates a new V0040Float64NoVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *V0040Float64NoVal) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *V0040Float64NoVal) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *V0040Float64NoVal) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *V0040Float64NoVal) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetInfinite

`func (o *V0040Float64NoVal) GetInfinite() bool`

GetInfinite returns the Infinite field if non-nil, zero value otherwise.

### GetInfiniteOk

`func (o *V0040Float64NoVal) GetInfiniteOk() (*bool, bool)`

GetInfiniteOk returns a tuple with the Infinite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfinite

`func (o *V0040Float64NoVal) SetInfinite(v bool)`

SetInfinite sets Infinite field to given value.

### HasInfinite

`func (o *V0040Float64NoVal) HasInfinite() bool`

HasInfinite returns a boolean if a field has been set.

### GetNumber

`func (o *V0040Float64NoVal) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *V0040Float64NoVal) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *V0040Float64NoVal) SetNumber(v float64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *V0040Float64NoVal) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


