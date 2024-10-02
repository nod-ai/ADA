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

// checks if the V0040License type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040License{}

// V0040License struct for V0040License
type V0040License struct {
	LicenseName  *string `json:"LicenseName,omitempty"`
	Total        *int32  `json:"Total,omitempty"`
	Used         *int32  `json:"Used,omitempty"`
	Free         *int32  `json:"Free,omitempty"`
	Remote       *bool   `json:"Remote,omitempty"`
	Reserved     *int32  `json:"Reserved,omitempty"`
	LastConsumed *int32  `json:"LastConsumed,omitempty"`
	LastDeficit  *int32  `json:"LastDeficit,omitempty"`
	LastUpdate   *int64  `json:"LastUpdate,omitempty"`
}

// NewV0040License instantiates a new V0040License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040License() *V0040License {
	this := V0040License{}
	return &this
}

// NewV0040LicenseWithDefaults instantiates a new V0040License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040LicenseWithDefaults() *V0040License {
	this := V0040License{}
	return &this
}

// GetLicenseName returns the LicenseName field value if set, zero value otherwise.
func (o *V0040License) GetLicenseName() string {
	if o == nil || IsNil(o.LicenseName) {
		var ret string
		return ret
	}
	return *o.LicenseName
}

// GetLicenseNameOk returns a tuple with the LicenseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetLicenseNameOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseName) {
		return nil, false
	}
	return o.LicenseName, true
}

// HasLicenseName returns a boolean if a field has been set.
func (o *V0040License) HasLicenseName() bool {
	if o != nil && !IsNil(o.LicenseName) {
		return true
	}

	return false
}

// SetLicenseName gets a reference to the given string and assigns it to the LicenseName field.
func (o *V0040License) SetLicenseName(v string) {
	o.LicenseName = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0040License) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0040License) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *V0040License) SetTotal(v int32) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *V0040License) GetUsed() int32 {
	if o == nil || IsNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *V0040License) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *V0040License) SetUsed(v int32) {
	o.Used = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *V0040License) GetFree() int32 {
	if o == nil || IsNil(o.Free) {
		var ret int32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetFreeOk() (*int32, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *V0040License) HasFree() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given int32 and assigns it to the Free field.
func (o *V0040License) SetFree(v int32) {
	o.Free = &v
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *V0040License) GetRemote() bool {
	if o == nil || IsNil(o.Remote) {
		var ret bool
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetRemoteOk() (*bool, bool) {
	if o == nil || IsNil(o.Remote) {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *V0040License) HasRemote() bool {
	if o != nil && !IsNil(o.Remote) {
		return true
	}

	return false
}

// SetRemote gets a reference to the given bool and assigns it to the Remote field.
func (o *V0040License) SetRemote(v bool) {
	o.Remote = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *V0040License) GetReserved() int32 {
	if o == nil || IsNil(o.Reserved) {
		var ret int32
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetReservedOk() (*int32, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *V0040License) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int32 and assigns it to the Reserved field.
func (o *V0040License) SetReserved(v int32) {
	o.Reserved = &v
}

// GetLastConsumed returns the LastConsumed field value if set, zero value otherwise.
func (o *V0040License) GetLastConsumed() int32 {
	if o == nil || IsNil(o.LastConsumed) {
		var ret int32
		return ret
	}
	return *o.LastConsumed
}

// GetLastConsumedOk returns a tuple with the LastConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetLastConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.LastConsumed) {
		return nil, false
	}
	return o.LastConsumed, true
}

// HasLastConsumed returns a boolean if a field has been set.
func (o *V0040License) HasLastConsumed() bool {
	if o != nil && !IsNil(o.LastConsumed) {
		return true
	}

	return false
}

// SetLastConsumed gets a reference to the given int32 and assigns it to the LastConsumed field.
func (o *V0040License) SetLastConsumed(v int32) {
	o.LastConsumed = &v
}

// GetLastDeficit returns the LastDeficit field value if set, zero value otherwise.
func (o *V0040License) GetLastDeficit() int32 {
	if o == nil || IsNil(o.LastDeficit) {
		var ret int32
		return ret
	}
	return *o.LastDeficit
}

// GetLastDeficitOk returns a tuple with the LastDeficit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetLastDeficitOk() (*int32, bool) {
	if o == nil || IsNil(o.LastDeficit) {
		return nil, false
	}
	return o.LastDeficit, true
}

// HasLastDeficit returns a boolean if a field has been set.
func (o *V0040License) HasLastDeficit() bool {
	if o != nil && !IsNil(o.LastDeficit) {
		return true
	}

	return false
}

// SetLastDeficit gets a reference to the given int32 and assigns it to the LastDeficit field.
func (o *V0040License) SetLastDeficit(v int32) {
	o.LastDeficit = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *V0040License) GetLastUpdate() int64 {
	if o == nil || IsNil(o.LastUpdate) {
		var ret int64
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040License) GetLastUpdateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *V0040License) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given int64 and assigns it to the LastUpdate field.
func (o *V0040License) SetLastUpdate(v int64) {
	o.LastUpdate = &v
}

func (o V0040License) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040License) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicenseName) {
		toSerialize["LicenseName"] = o.LicenseName
	}
	if !IsNil(o.Total) {
		toSerialize["Total"] = o.Total
	}
	if !IsNil(o.Used) {
		toSerialize["Used"] = o.Used
	}
	if !IsNil(o.Free) {
		toSerialize["Free"] = o.Free
	}
	if !IsNil(o.Remote) {
		toSerialize["Remote"] = o.Remote
	}
	if !IsNil(o.Reserved) {
		toSerialize["Reserved"] = o.Reserved
	}
	if !IsNil(o.LastConsumed) {
		toSerialize["LastConsumed"] = o.LastConsumed
	}
	if !IsNil(o.LastDeficit) {
		toSerialize["LastDeficit"] = o.LastDeficit
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["LastUpdate"] = o.LastUpdate
	}
	return toSerialize, nil
}

type NullableV0040License struct {
	value *V0040License
	isSet bool
}

func (v NullableV0040License) Get() *V0040License {
	return v.value
}

func (v *NullableV0040License) Set(val *V0040License) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040License) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040License) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040License(val *V0040License) *NullableV0040License {
	return &NullableV0040License{value: val, isSet: true}
}

func (v NullableV0040License) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040License) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}