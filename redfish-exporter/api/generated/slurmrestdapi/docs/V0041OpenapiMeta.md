# V0041OpenapiMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**V0041OpenapiMetaPlugin**](V0041OpenapiMetaPlugin.md) |  | [optional] 
**Client** | Pointer to [**V0041OpenapiMetaClient**](V0041OpenapiMetaClient.md) |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Slurm** | Pointer to [**V0041OpenapiMetaSlurm**](V0041OpenapiMetaSlurm.md) |  | [optional] 

## Methods

### NewV0041OpenapiMeta

`func NewV0041OpenapiMeta() *V0041OpenapiMeta`

NewV0041OpenapiMeta instantiates a new V0041OpenapiMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiMetaWithDefaults

`func NewV0041OpenapiMetaWithDefaults() *V0041OpenapiMeta`

NewV0041OpenapiMetaWithDefaults instantiates a new V0041OpenapiMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *V0041OpenapiMeta) GetPlugin() V0041OpenapiMetaPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *V0041OpenapiMeta) GetPluginOk() (*V0041OpenapiMetaPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *V0041OpenapiMeta) SetPlugin(v V0041OpenapiMetaPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *V0041OpenapiMeta) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetClient

`func (o *V0041OpenapiMeta) GetClient() V0041OpenapiMetaClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *V0041OpenapiMeta) GetClientOk() (*V0041OpenapiMetaClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *V0041OpenapiMeta) SetClient(v V0041OpenapiMetaClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *V0041OpenapiMeta) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCommand

`func (o *V0041OpenapiMeta) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0041OpenapiMeta) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0041OpenapiMeta) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0041OpenapiMeta) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSlurm

`func (o *V0041OpenapiMeta) GetSlurm() V0041OpenapiMetaSlurm`

GetSlurm returns the Slurm field if non-nil, zero value otherwise.

### GetSlurmOk

`func (o *V0041OpenapiMeta) GetSlurmOk() (*V0041OpenapiMetaSlurm, bool)`

GetSlurmOk returns a tuple with the Slurm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurm

`func (o *V0041OpenapiMeta) SetSlurm(v V0041OpenapiMetaSlurm)`

SetSlurm sets Slurm field to given value.

### HasSlurm

`func (o *V0041OpenapiMeta) HasSlurm() bool`

HasSlurm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


