# V0041JobResNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Node index | 
**Name** | **string** | Node name | 
**Cpus** | Pointer to [**V0041JobResNodesInnerCpus**](V0041JobResNodesInnerCpus.md) |  | [optional] 
**Memory** | Pointer to [**V0041JobResNodesInnerMemory**](V0041JobResNodesInnerMemory.md) |  | [optional] 
**Sockets** | [**[]V0041JobResSocketArrayInner**](V0041JobResSocketArrayInner.md) |  | 

## Methods

### NewV0041JobResNodesInner

`func NewV0041JobResNodesInner(index int32, name string, sockets []V0041JobResSocketArrayInner, ) *V0041JobResNodesInner`

NewV0041JobResNodesInner instantiates a new V0041JobResNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041JobResNodesInnerWithDefaults

`func NewV0041JobResNodesInnerWithDefaults() *V0041JobResNodesInner`

NewV0041JobResNodesInnerWithDefaults instantiates a new V0041JobResNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *V0041JobResNodesInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *V0041JobResNodesInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *V0041JobResNodesInner) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetName

`func (o *V0041JobResNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041JobResNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041JobResNodesInner) SetName(v string)`

SetName sets Name field to given value.


### GetCpus

`func (o *V0041JobResNodesInner) GetCpus() V0041JobResNodesInnerCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041JobResNodesInner) GetCpusOk() (*V0041JobResNodesInnerCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041JobResNodesInner) SetCpus(v V0041JobResNodesInnerCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0041JobResNodesInner) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetMemory

`func (o *V0041JobResNodesInner) GetMemory() V0041JobResNodesInnerMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V0041JobResNodesInner) GetMemoryOk() (*V0041JobResNodesInnerMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V0041JobResNodesInner) SetMemory(v V0041JobResNodesInnerMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V0041JobResNodesInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSockets

`func (o *V0041JobResNodesInner) GetSockets() []V0041JobResSocketArrayInner`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0041JobResNodesInner) GetSocketsOk() (*[]V0041JobResSocketArrayInner, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0041JobResNodesInner) SetSockets(v []V0041JobResSocketArrayInner)`

SetSockets sets Sockets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


