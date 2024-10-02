/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.08.0-0rc1&openapi/slurmctld&openapi/v0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestdapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the V0041JobResNodesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041JobResNodesInner{}

// V0041JobResNodesInner Job resources for a node
type V0041JobResNodesInner struct {
	// Node index
	Index int32 `json:"index"`
	// Node name
	Name    string                        `json:"name"`
	Cpus    *V0041JobResNodesInnerCpus    `json:"cpus,omitempty"`
	Memory  *V0041JobResNodesInnerMemory  `json:"memory,omitempty"`
	Sockets []V0041JobResSocketArrayInner `json:"sockets"`
}

type _V0041JobResNodesInner V0041JobResNodesInner

// NewV0041JobResNodesInner instantiates a new V0041JobResNodesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041JobResNodesInner(index int32, name string, sockets []V0041JobResSocketArrayInner) *V0041JobResNodesInner {
	this := V0041JobResNodesInner{}
	this.Index = index
	this.Name = name
	this.Sockets = sockets
	return &this
}

// NewV0041JobResNodesInnerWithDefaults instantiates a new V0041JobResNodesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041JobResNodesInnerWithDefaults() *V0041JobResNodesInner {
	this := V0041JobResNodesInner{}
	return &this
}

// GetIndex returns the Index field value
func (o *V0041JobResNodesInner) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *V0041JobResNodesInner) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *V0041JobResNodesInner) SetIndex(v int32) {
	o.Index = v
}

// GetName returns the Name field value
func (o *V0041JobResNodesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0041JobResNodesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0041JobResNodesInner) SetName(v string) {
	o.Name = v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *V0041JobResNodesInner) GetCpus() V0041JobResNodesInnerCpus {
	if o == nil || IsNil(o.Cpus) {
		var ret V0041JobResNodesInnerCpus
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobResNodesInner) GetCpusOk() (*V0041JobResNodesInnerCpus, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *V0041JobResNodesInner) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given V0041JobResNodesInnerCpus and assigns it to the Cpus field.
func (o *V0041JobResNodesInner) SetCpus(v V0041JobResNodesInnerCpus) {
	o.Cpus = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *V0041JobResNodesInner) GetMemory() V0041JobResNodesInnerMemory {
	if o == nil || IsNil(o.Memory) {
		var ret V0041JobResNodesInnerMemory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobResNodesInner) GetMemoryOk() (*V0041JobResNodesInnerMemory, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *V0041JobResNodesInner) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given V0041JobResNodesInnerMemory and assigns it to the Memory field.
func (o *V0041JobResNodesInner) SetMemory(v V0041JobResNodesInnerMemory) {
	o.Memory = &v
}

// GetSockets returns the Sockets field value
func (o *V0041JobResNodesInner) GetSockets() []V0041JobResSocketArrayInner {
	if o == nil {
		var ret []V0041JobResSocketArrayInner
		return ret
	}

	return o.Sockets
}

// GetSocketsOk returns a tuple with the Sockets field value
// and a boolean to check if the value has been set.
func (o *V0041JobResNodesInner) GetSocketsOk() ([]V0041JobResSocketArrayInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sockets, true
}

// SetSockets sets field value
func (o *V0041JobResNodesInner) SetSockets(v []V0041JobResSocketArrayInner) {
	o.Sockets = v
}

func (o V0041JobResNodesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041JobResNodesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["name"] = o.Name
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	toSerialize["sockets"] = o.Sockets
	return toSerialize, nil
}

func (o *V0041JobResNodesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"name",
		"sockets",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0041JobResNodesInner := _V0041JobResNodesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041JobResNodesInner)

	if err != nil {
		return err
	}

	*o = V0041JobResNodesInner(varV0041JobResNodesInner)

	return err
}

type NullableV0041JobResNodesInner struct {
	value *V0041JobResNodesInner
	isSet bool
}

func (v NullableV0041JobResNodesInner) Get() *V0041JobResNodesInner {
	return v.value
}

func (v *NullableV0041JobResNodesInner) Set(val *V0041JobResNodesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041JobResNodesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041JobResNodesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041JobResNodesInner(val *V0041JobResNodesInner) *NullableV0041JobResNodesInner {
	return &NullableV0041JobResNodesInner{value: val, isSet: true}
}

func (v NullableV0041JobResNodesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041JobResNodesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}