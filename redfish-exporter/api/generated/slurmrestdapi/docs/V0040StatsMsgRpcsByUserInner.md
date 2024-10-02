# V0040StatsMsgRpcsByUserInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | user name | [optional] 
**UserId** | Pointer to **int32** | user id (numeric) | [optional] 
**Count** | Pointer to **int64** | Number of RPCs received | [optional] 
**AverageTime** | Pointer to **int64** | Average time spent processing RPC in seconds | [optional] 
**TotalTime** | Pointer to **int64** | Total time spent processing RPC in seconds | [optional] 

## Methods

### NewV0040StatsMsgRpcsByUserInner

`func NewV0040StatsMsgRpcsByUserInner() *V0040StatsMsgRpcsByUserInner`

NewV0040StatsMsgRpcsByUserInner instantiates a new V0040StatsMsgRpcsByUserInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040StatsMsgRpcsByUserInnerWithDefaults

`func NewV0040StatsMsgRpcsByUserInnerWithDefaults() *V0040StatsMsgRpcsByUserInner`

NewV0040StatsMsgRpcsByUserInnerWithDefaults instantiates a new V0040StatsMsgRpcsByUserInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *V0040StatsMsgRpcsByUserInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0040StatsMsgRpcsByUserInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0040StatsMsgRpcsByUserInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V0040StatsMsgRpcsByUserInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *V0040StatsMsgRpcsByUserInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0040StatsMsgRpcsByUserInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0040StatsMsgRpcsByUserInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0040StatsMsgRpcsByUserInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCount

`func (o *V0040StatsMsgRpcsByUserInner) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0040StatsMsgRpcsByUserInner) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0040StatsMsgRpcsByUserInner) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0040StatsMsgRpcsByUserInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAverageTime

`func (o *V0040StatsMsgRpcsByUserInner) GetAverageTime() int64`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *V0040StatsMsgRpcsByUserInner) GetAverageTimeOk() (*int64, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *V0040StatsMsgRpcsByUserInner) SetAverageTime(v int64)`

SetAverageTime sets AverageTime field to given value.

### HasAverageTime

`func (o *V0040StatsMsgRpcsByUserInner) HasAverageTime() bool`

HasAverageTime returns a boolean if a field has been set.

### GetTotalTime

`func (o *V0040StatsMsgRpcsByUserInner) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *V0040StatsMsgRpcsByUserInner) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *V0040StatsMsgRpcsByUserInner) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *V0040StatsMsgRpcsByUserInner) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


