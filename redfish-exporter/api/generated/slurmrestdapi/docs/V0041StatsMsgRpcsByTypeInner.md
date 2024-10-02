# V0041StatsMsgRpcsByTypeInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int32** | Message type as integer | 
**MessageType** | **string** | Message type as string | 
**Count** | **int32** | Number of RPCs received | 
**TotalTime** | **int64** | Total time spent processing RPC in seconds | 
**AverageTime** | [**V0041Uint64NoValStruct**](V0041Uint64NoValStruct.md) |  | 

## Methods

### NewV0041StatsMsgRpcsByTypeInner

`func NewV0041StatsMsgRpcsByTypeInner(typeId int32, messageType string, count int32, totalTime int64, averageTime V0041Uint64NoValStruct, ) *V0041StatsMsgRpcsByTypeInner`

NewV0041StatsMsgRpcsByTypeInner instantiates a new V0041StatsMsgRpcsByTypeInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041StatsMsgRpcsByTypeInnerWithDefaults

`func NewV0041StatsMsgRpcsByTypeInnerWithDefaults() *V0041StatsMsgRpcsByTypeInner`

NewV0041StatsMsgRpcsByTypeInnerWithDefaults instantiates a new V0041StatsMsgRpcsByTypeInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *V0041StatsMsgRpcsByTypeInner) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *V0041StatsMsgRpcsByTypeInner) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *V0041StatsMsgRpcsByTypeInner) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetMessageType

`func (o *V0041StatsMsgRpcsByTypeInner) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *V0041StatsMsgRpcsByTypeInner) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *V0041StatsMsgRpcsByTypeInner) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetCount

`func (o *V0041StatsMsgRpcsByTypeInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0041StatsMsgRpcsByTypeInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0041StatsMsgRpcsByTypeInner) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotalTime

`func (o *V0041StatsMsgRpcsByTypeInner) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *V0041StatsMsgRpcsByTypeInner) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *V0041StatsMsgRpcsByTypeInner) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetAverageTime

`func (o *V0041StatsMsgRpcsByTypeInner) GetAverageTime() V0041Uint64NoValStruct`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *V0041StatsMsgRpcsByTypeInner) GetAverageTimeOk() (*V0041Uint64NoValStruct, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *V0041StatsMsgRpcsByTypeInner) SetAverageTime(v V0041Uint64NoValStruct)`

SetAverageTime sets AverageTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


