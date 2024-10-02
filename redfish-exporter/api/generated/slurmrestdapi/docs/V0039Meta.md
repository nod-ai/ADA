# V0039Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**V0039MetaPlugin**](V0039MetaPlugin.md) |  | [optional] 
**Slurm** | Pointer to [**V0039MetaSlurm**](V0039MetaSlurm.md) |  | [optional] 

## Methods

### NewV0039Meta

`func NewV0039Meta() *V0039Meta`

NewV0039Meta instantiates a new V0039Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039MetaWithDefaults

`func NewV0039MetaWithDefaults() *V0039Meta`

NewV0039MetaWithDefaults instantiates a new V0039Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *V0039Meta) GetPlugin() V0039MetaPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *V0039Meta) GetPluginOk() (*V0039MetaPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *V0039Meta) SetPlugin(v V0039MetaPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *V0039Meta) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetSlurm

`func (o *V0039Meta) GetSlurm() V0039MetaSlurm`

GetSlurm returns the Slurm field if non-nil, zero value otherwise.

### GetSlurmOk

`func (o *V0039Meta) GetSlurmOk() (*V0039MetaSlurm, bool)`

GetSlurmOk returns a tuple with the Slurm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurm

`func (o *V0039Meta) SetSlurm(v V0039MetaSlurm)`

SetSlurm sets Slurm field to given value.

### HasSlurm

`func (o *V0039Meta) HasSlurm() bool`

HasSlurm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


