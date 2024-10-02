# V0041CronEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **[]string** |  | [optional] 
**Minute** | Pointer to **string** |  | [optional] 
**Hour** | Pointer to **string** |  | [optional] 
**DayOfMonth** | Pointer to **string** |  | [optional] 
**Month** | Pointer to **string** |  | [optional] 
**DayOfWeek** | Pointer to **string** |  | [optional] 
**Specification** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Line** | Pointer to [**V0041CronEntryLine**](V0041CronEntryLine.md) |  | [optional] 

## Methods

### NewV0041CronEntry

`func NewV0041CronEntry() *V0041CronEntry`

NewV0041CronEntry instantiates a new V0041CronEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041CronEntryWithDefaults

`func NewV0041CronEntryWithDefaults() *V0041CronEntry`

NewV0041CronEntryWithDefaults instantiates a new V0041CronEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *V0041CronEntry) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041CronEntry) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041CronEntry) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041CronEntry) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMinute

`func (o *V0041CronEntry) GetMinute() string`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *V0041CronEntry) GetMinuteOk() (*string, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *V0041CronEntry) SetMinute(v string)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *V0041CronEntry) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetHour

`func (o *V0041CronEntry) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *V0041CronEntry) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *V0041CronEntry) SetHour(v string)`

SetHour sets Hour field to given value.

### HasHour

`func (o *V0041CronEntry) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *V0041CronEntry) GetDayOfMonth() string`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *V0041CronEntry) GetDayOfMonthOk() (*string, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *V0041CronEntry) SetDayOfMonth(v string)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *V0041CronEntry) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetMonth

`func (o *V0041CronEntry) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *V0041CronEntry) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *V0041CronEntry) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *V0041CronEntry) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *V0041CronEntry) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *V0041CronEntry) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *V0041CronEntry) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *V0041CronEntry) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetSpecification

`func (o *V0041CronEntry) GetSpecification() string`

GetSpecification returns the Specification field if non-nil, zero value otherwise.

### GetSpecificationOk

`func (o *V0041CronEntry) GetSpecificationOk() (*string, bool)`

GetSpecificationOk returns a tuple with the Specification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecification

`func (o *V0041CronEntry) SetSpecification(v string)`

SetSpecification sets Specification field to given value.

### HasSpecification

`func (o *V0041CronEntry) HasSpecification() bool`

HasSpecification returns a boolean if a field has been set.

### GetCommand

`func (o *V0041CronEntry) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0041CronEntry) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0041CronEntry) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0041CronEntry) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetLine

`func (o *V0041CronEntry) GetLine() V0041CronEntryLine`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *V0041CronEntry) GetLineOk() (*V0041CronEntryLine, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *V0041CronEntry) SetLine(v V0041CronEntryLine)`

SetLine sets Line field to given value.

### HasLine

`func (o *V0041CronEntry) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


