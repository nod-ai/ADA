/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.08.0-0rc1&openapi/slurmctld&openapi/v0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestdapi

import (
	"encoding/json"
)

// checks if the V0041OpenapiMetaPlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiMetaPlugin{}

// V0041OpenapiMetaPlugin struct for V0041OpenapiMetaPlugin
type V0041OpenapiMetaPlugin struct {
	// Slurm plugin type (if applicable)
	Type *string `json:"type,omitempty"`
	// Slurm plugin name (if applicable)
	Name *string `json:"name,omitempty"`
	// Slurm data_parser plugin
	DataParser *string `json:"data_parser,omitempty"`
	// Slurm accounting plugin
	AccountingStorage *string `json:"accounting_storage,omitempty"`
}

// NewV0041OpenapiMetaPlugin instantiates a new V0041OpenapiMetaPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiMetaPlugin() *V0041OpenapiMetaPlugin {
	this := V0041OpenapiMetaPlugin{}
	return &this
}

// NewV0041OpenapiMetaPluginWithDefaults instantiates a new V0041OpenapiMetaPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiMetaPluginWithDefaults() *V0041OpenapiMetaPlugin {
	this := V0041OpenapiMetaPlugin{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V0041OpenapiMetaPlugin) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiMetaPlugin) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V0041OpenapiMetaPlugin) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V0041OpenapiMetaPlugin) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0041OpenapiMetaPlugin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiMetaPlugin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0041OpenapiMetaPlugin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0041OpenapiMetaPlugin) SetName(v string) {
	o.Name = &v
}

// GetDataParser returns the DataParser field value if set, zero value otherwise.
func (o *V0041OpenapiMetaPlugin) GetDataParser() string {
	if o == nil || IsNil(o.DataParser) {
		var ret string
		return ret
	}
	return *o.DataParser
}

// GetDataParserOk returns a tuple with the DataParser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiMetaPlugin) GetDataParserOk() (*string, bool) {
	if o == nil || IsNil(o.DataParser) {
		return nil, false
	}
	return o.DataParser, true
}

// HasDataParser returns a boolean if a field has been set.
func (o *V0041OpenapiMetaPlugin) HasDataParser() bool {
	if o != nil && !IsNil(o.DataParser) {
		return true
	}

	return false
}

// SetDataParser gets a reference to the given string and assigns it to the DataParser field.
func (o *V0041OpenapiMetaPlugin) SetDataParser(v string) {
	o.DataParser = &v
}

// GetAccountingStorage returns the AccountingStorage field value if set, zero value otherwise.
func (o *V0041OpenapiMetaPlugin) GetAccountingStorage() string {
	if o == nil || IsNil(o.AccountingStorage) {
		var ret string
		return ret
	}
	return *o.AccountingStorage
}

// GetAccountingStorageOk returns a tuple with the AccountingStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiMetaPlugin) GetAccountingStorageOk() (*string, bool) {
	if o == nil || IsNil(o.AccountingStorage) {
		return nil, false
	}
	return o.AccountingStorage, true
}

// HasAccountingStorage returns a boolean if a field has been set.
func (o *V0041OpenapiMetaPlugin) HasAccountingStorage() bool {
	if o != nil && !IsNil(o.AccountingStorage) {
		return true
	}

	return false
}

// SetAccountingStorage gets a reference to the given string and assigns it to the AccountingStorage field.
func (o *V0041OpenapiMetaPlugin) SetAccountingStorage(v string) {
	o.AccountingStorage = &v
}

func (o V0041OpenapiMetaPlugin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiMetaPlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DataParser) {
		toSerialize["data_parser"] = o.DataParser
	}
	if !IsNil(o.AccountingStorage) {
		toSerialize["accounting_storage"] = o.AccountingStorage
	}
	return toSerialize, nil
}

type NullableV0041OpenapiMetaPlugin struct {
	value *V0041OpenapiMetaPlugin
	isSet bool
}

func (v NullableV0041OpenapiMetaPlugin) Get() *V0041OpenapiMetaPlugin {
	return v.value
}

func (v *NullableV0041OpenapiMetaPlugin) Set(val *V0041OpenapiMetaPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiMetaPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiMetaPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiMetaPlugin(val *V0041OpenapiMetaPlugin) *NullableV0041OpenapiMetaPlugin {
	return &NullableV0041OpenapiMetaPlugin{value: val, isSet: true}
}

func (v NullableV0041OpenapiMetaPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiMetaPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}