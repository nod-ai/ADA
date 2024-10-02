# V0041OpenapiMetaSlurm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**V0041OpenapiMetaSlurmVersion**](V0041OpenapiMetaSlurmVersion.md) |  | [optional] 
**Release** | Pointer to **string** | Slurm release string | [optional] 
**Cluster** | Pointer to **string** | Slurm cluster name | [optional] 

## Methods

### NewV0041OpenapiMetaSlurm

`func NewV0041OpenapiMetaSlurm() *V0041OpenapiMetaSlurm`

NewV0041OpenapiMetaSlurm instantiates a new V0041OpenapiMetaSlurm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiMetaSlurmWithDefaults

`func NewV0041OpenapiMetaSlurmWithDefaults() *V0041OpenapiMetaSlurm`

NewV0041OpenapiMetaSlurmWithDefaults instantiates a new V0041OpenapiMetaSlurm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *V0041OpenapiMetaSlurm) GetVersion() V0041OpenapiMetaSlurmVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V0041OpenapiMetaSlurm) GetVersionOk() (*V0041OpenapiMetaSlurmVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V0041OpenapiMetaSlurm) SetVersion(v V0041OpenapiMetaSlurmVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V0041OpenapiMetaSlurm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *V0041OpenapiMetaSlurm) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *V0041OpenapiMetaSlurm) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *V0041OpenapiMetaSlurm) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *V0041OpenapiMetaSlurm) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetCluster

`func (o *V0041OpenapiMetaSlurm) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041OpenapiMetaSlurm) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041OpenapiMetaSlurm) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041OpenapiMetaSlurm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


