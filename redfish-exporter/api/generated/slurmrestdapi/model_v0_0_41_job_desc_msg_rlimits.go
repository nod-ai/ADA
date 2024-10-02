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

// checks if the V0041JobDescMsgRlimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041JobDescMsgRlimits{}

// V0041JobDescMsgRlimits struct for V0041JobDescMsgRlimits
type V0041JobDescMsgRlimits struct {
	Cpu     *V0041Uint64NoValStruct `json:"cpu,omitempty"`
	Fsize   *V0041Uint64NoValStruct `json:"fsize,omitempty"`
	Data    *V0041Uint64NoValStruct `json:"data,omitempty"`
	Stack   *V0041Uint64NoValStruct `json:"stack,omitempty"`
	Core    *V0041Uint64NoValStruct `json:"core,omitempty"`
	Rss     *V0041Uint64NoValStruct `json:"rss,omitempty"`
	Nproc   *V0041Uint64NoValStruct `json:"nproc,omitempty"`
	Nofile  *V0041Uint64NoValStruct `json:"nofile,omitempty"`
	Memlock *V0041Uint64NoValStruct `json:"memlock,omitempty"`
	As      *V0041Uint64NoValStruct `json:"as,omitempty"`
}

// NewV0041JobDescMsgRlimits instantiates a new V0041JobDescMsgRlimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041JobDescMsgRlimits() *V0041JobDescMsgRlimits {
	this := V0041JobDescMsgRlimits{}
	return &this
}

// NewV0041JobDescMsgRlimitsWithDefaults instantiates a new V0041JobDescMsgRlimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041JobDescMsgRlimitsWithDefaults() *V0041JobDescMsgRlimits {
	this := V0041JobDescMsgRlimits{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetCpu() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Cpu) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetCpuOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given V0041Uint64NoValStruct and assigns it to the Cpu field.
func (o *V0041JobDescMsgRlimits) SetCpu(v V0041Uint64NoValStruct) {
	o.Cpu = &v
}

// GetFsize returns the Fsize field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetFsize() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Fsize) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Fsize
}

// GetFsizeOk returns a tuple with the Fsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetFsizeOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Fsize) {
		return nil, false
	}
	return o.Fsize, true
}

// HasFsize returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasFsize() bool {
	if o != nil && !IsNil(o.Fsize) {
		return true
	}

	return false
}

// SetFsize gets a reference to the given V0041Uint64NoValStruct and assigns it to the Fsize field.
func (o *V0041JobDescMsgRlimits) SetFsize(v V0041Uint64NoValStruct) {
	o.Fsize = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetData() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Data) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetDataOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given V0041Uint64NoValStruct and assigns it to the Data field.
func (o *V0041JobDescMsgRlimits) SetData(v V0041Uint64NoValStruct) {
	o.Data = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetStack() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Stack) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetStackOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasStack() bool {
	if o != nil && !IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given V0041Uint64NoValStruct and assigns it to the Stack field.
func (o *V0041JobDescMsgRlimits) SetStack(v V0041Uint64NoValStruct) {
	o.Stack = &v
}

// GetCore returns the Core field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetCore() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Core) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Core
}

// GetCoreOk returns a tuple with the Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetCoreOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Core) {
		return nil, false
	}
	return o.Core, true
}

// HasCore returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasCore() bool {
	if o != nil && !IsNil(o.Core) {
		return true
	}

	return false
}

// SetCore gets a reference to the given V0041Uint64NoValStruct and assigns it to the Core field.
func (o *V0041JobDescMsgRlimits) SetCore(v V0041Uint64NoValStruct) {
	o.Core = &v
}

// GetRss returns the Rss field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetRss() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Rss) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Rss
}

// GetRssOk returns a tuple with the Rss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetRssOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Rss) {
		return nil, false
	}
	return o.Rss, true
}

// HasRss returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasRss() bool {
	if o != nil && !IsNil(o.Rss) {
		return true
	}

	return false
}

// SetRss gets a reference to the given V0041Uint64NoValStruct and assigns it to the Rss field.
func (o *V0041JobDescMsgRlimits) SetRss(v V0041Uint64NoValStruct) {
	o.Rss = &v
}

// GetNproc returns the Nproc field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetNproc() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Nproc) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Nproc
}

// GetNprocOk returns a tuple with the Nproc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetNprocOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Nproc) {
		return nil, false
	}
	return o.Nproc, true
}

// HasNproc returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasNproc() bool {
	if o != nil && !IsNil(o.Nproc) {
		return true
	}

	return false
}

// SetNproc gets a reference to the given V0041Uint64NoValStruct and assigns it to the Nproc field.
func (o *V0041JobDescMsgRlimits) SetNproc(v V0041Uint64NoValStruct) {
	o.Nproc = &v
}

// GetNofile returns the Nofile field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetNofile() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Nofile) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Nofile
}

// GetNofileOk returns a tuple with the Nofile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetNofileOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Nofile) {
		return nil, false
	}
	return o.Nofile, true
}

// HasNofile returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasNofile() bool {
	if o != nil && !IsNil(o.Nofile) {
		return true
	}

	return false
}

// SetNofile gets a reference to the given V0041Uint64NoValStruct and assigns it to the Nofile field.
func (o *V0041JobDescMsgRlimits) SetNofile(v V0041Uint64NoValStruct) {
	o.Nofile = &v
}

// GetMemlock returns the Memlock field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetMemlock() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.Memlock) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.Memlock
}

// GetMemlockOk returns a tuple with the Memlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetMemlockOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.Memlock) {
		return nil, false
	}
	return o.Memlock, true
}

// HasMemlock returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasMemlock() bool {
	if o != nil && !IsNil(o.Memlock) {
		return true
	}

	return false
}

// SetMemlock gets a reference to the given V0041Uint64NoValStruct and assigns it to the Memlock field.
func (o *V0041JobDescMsgRlimits) SetMemlock(v V0041Uint64NoValStruct) {
	o.Memlock = &v
}

// GetAs returns the As field value if set, zero value otherwise.
func (o *V0041JobDescMsgRlimits) GetAs() V0041Uint64NoValStruct {
	if o == nil || IsNil(o.As) {
		var ret V0041Uint64NoValStruct
		return ret
	}
	return *o.As
}

// GetAsOk returns a tuple with the As field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgRlimits) GetAsOk() (*V0041Uint64NoValStruct, bool) {
	if o == nil || IsNil(o.As) {
		return nil, false
	}
	return o.As, true
}

// HasAs returns a boolean if a field has been set.
func (o *V0041JobDescMsgRlimits) HasAs() bool {
	if o != nil && !IsNil(o.As) {
		return true
	}

	return false
}

// SetAs gets a reference to the given V0041Uint64NoValStruct and assigns it to the As field.
func (o *V0041JobDescMsgRlimits) SetAs(v V0041Uint64NoValStruct) {
	o.As = &v
}

func (o V0041JobDescMsgRlimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041JobDescMsgRlimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Fsize) {
		toSerialize["fsize"] = o.Fsize
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	if !IsNil(o.Core) {
		toSerialize["core"] = o.Core
	}
	if !IsNil(o.Rss) {
		toSerialize["rss"] = o.Rss
	}
	if !IsNil(o.Nproc) {
		toSerialize["nproc"] = o.Nproc
	}
	if !IsNil(o.Nofile) {
		toSerialize["nofile"] = o.Nofile
	}
	if !IsNil(o.Memlock) {
		toSerialize["memlock"] = o.Memlock
	}
	if !IsNil(o.As) {
		toSerialize["as"] = o.As
	}
	return toSerialize, nil
}

type NullableV0041JobDescMsgRlimits struct {
	value *V0041JobDescMsgRlimits
	isSet bool
}

func (v NullableV0041JobDescMsgRlimits) Get() *V0041JobDescMsgRlimits {
	return v.value
}

func (v *NullableV0041JobDescMsgRlimits) Set(val *V0041JobDescMsgRlimits) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041JobDescMsgRlimits) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041JobDescMsgRlimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041JobDescMsgRlimits(val *V0041JobDescMsgRlimits) *NullableV0041JobDescMsgRlimits {
	return &NullableV0041JobDescMsgRlimits{value: val, isSet: true}
}

func (v NullableV0041JobDescMsgRlimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041JobDescMsgRlimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}