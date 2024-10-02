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

// checks if the V0041JobResNodesInnerCpus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041JobResNodesInnerCpus{}

// V0041JobResNodesInnerCpus struct for V0041JobResNodesInnerCpus
type V0041JobResNodesInnerCpus struct {
	// Total number of CPUs assigned to job
	Count *int32 `json:"count,omitempty"`
	// Total number of CPUs used by job
	Used *int32 `json:"used,omitempty"`
}

// NewV0041JobResNodesInnerCpus instantiates a new V0041JobResNodesInnerCpus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041JobResNodesInnerCpus() *V0041JobResNodesInnerCpus {
	this := V0041JobResNodesInnerCpus{}
	return &this
}

// NewV0041JobResNodesInnerCpusWithDefaults instantiates a new V0041JobResNodesInnerCpus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041JobResNodesInnerCpusWithDefaults() *V0041JobResNodesInnerCpus {
	this := V0041JobResNodesInnerCpus{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0041JobResNodesInnerCpus) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobResNodesInnerCpus) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0041JobResNodesInnerCpus) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0041JobResNodesInnerCpus) SetCount(v int32) {
	o.Count = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *V0041JobResNodesInnerCpus) GetUsed() int32 {
	if o == nil || IsNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobResNodesInnerCpus) GetUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *V0041JobResNodesInnerCpus) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *V0041JobResNodesInnerCpus) SetUsed(v int32) {
	o.Used = &v
}

func (o V0041JobResNodesInnerCpus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041JobResNodesInnerCpus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return toSerialize, nil
}

type NullableV0041JobResNodesInnerCpus struct {
	value *V0041JobResNodesInnerCpus
	isSet bool
}

func (v NullableV0041JobResNodesInnerCpus) Get() *V0041JobResNodesInnerCpus {
	return v.value
}

func (v *NullableV0041JobResNodesInnerCpus) Set(val *V0041JobResNodesInnerCpus) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041JobResNodesInnerCpus) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041JobResNodesInnerCpus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041JobResNodesInnerCpus(val *V0041JobResNodesInnerCpus) *NullableV0041JobResNodesInnerCpus {
	return &NullableV0041JobResNodesInnerCpus{value: val, isSet: true}
}

func (v NullableV0041JobResNodesInnerCpus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041JobResNodesInnerCpus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}