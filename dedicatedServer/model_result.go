/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedServer

import (
	"encoding/json"
)

// checks if the Result type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Result{}

// Result Hardware info e.g. chassis, cpu, memory and etc.
type Result struct {
	Chassis *Chassis `json:"chassis,omitempty"`
	Cpu []Cpu1 `json:"cpu,omitempty"`
	Disks []Disk `json:"disks,omitempty"`
	Ipmi *Ipmi1 `json:"ipmi,omitempty"`
	Memory []MemoryBank `json:"memory,omitempty"`
	Network []Network `json:"network,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Result Result

// NewResult instantiates a new Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResult() *Result {
	this := Result{}
	return &this
}

// NewResultWithDefaults instantiates a new Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultWithDefaults() *Result {
	this := Result{}
	return &this
}

// GetChassis returns the Chassis field value if set, zero value otherwise.
func (o *Result) GetChassis() Chassis {
	if o == nil || IsNil(o.Chassis) {
		var ret Chassis
		return ret
	}
	return *o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetChassisOk() (*Chassis, bool) {
	if o == nil || IsNil(o.Chassis) {
		return nil, false
	}
	return o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *Result) HasChassis() bool {
	if o != nil && !IsNil(o.Chassis) {
		return true
	}

	return false
}

// SetChassis gets a reference to the given Chassis and assigns it to the Chassis field.
func (o *Result) SetChassis(v Chassis) {
	o.Chassis = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Result) GetCpu() []Cpu1 {
	if o == nil || IsNil(o.Cpu) {
		var ret []Cpu1
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetCpuOk() ([]Cpu1, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Result) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given []Cpu1 and assigns it to the Cpu field.
func (o *Result) SetCpu(v []Cpu1) {
	o.Cpu = v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *Result) GetDisks() []Disk {
	if o == nil || IsNil(o.Disks) {
		var ret []Disk
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetDisksOk() ([]Disk, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *Result) HasDisks() bool {
	if o != nil && !IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []Disk and assigns it to the Disks field.
func (o *Result) SetDisks(v []Disk) {
	o.Disks = v
}

// GetIpmi returns the Ipmi field value if set, zero value otherwise.
func (o *Result) GetIpmi() Ipmi1 {
	if o == nil || IsNil(o.Ipmi) {
		var ret Ipmi1
		return ret
	}
	return *o.Ipmi
}

// GetIpmiOk returns a tuple with the Ipmi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetIpmiOk() (*Ipmi1, bool) {
	if o == nil || IsNil(o.Ipmi) {
		return nil, false
	}
	return o.Ipmi, true
}

// HasIpmi returns a boolean if a field has been set.
func (o *Result) HasIpmi() bool {
	if o != nil && !IsNil(o.Ipmi) {
		return true
	}

	return false
}

// SetIpmi gets a reference to the given Ipmi1 and assigns it to the Ipmi field.
func (o *Result) SetIpmi(v Ipmi1) {
	o.Ipmi = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Result) GetMemory() []MemoryBank {
	if o == nil || IsNil(o.Memory) {
		var ret []MemoryBank
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetMemoryOk() ([]MemoryBank, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Result) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given []MemoryBank and assigns it to the Memory field.
func (o *Result) SetMemory(v []MemoryBank) {
	o.Memory = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Result) GetNetwork() []Network {
	if o == nil || IsNil(o.Network) {
		var ret []Network
		return ret
	}
	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetNetworkOk() ([]Network, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Result) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given []Network and assigns it to the Network field.
func (o *Result) SetNetwork(v []Network) {
	o.Network = v
}

func (o Result) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Chassis) {
		toSerialize["chassis"] = o.Chassis
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Disks) {
		toSerialize["disks"] = o.Disks
	}
	if !IsNil(o.Ipmi) {
		toSerialize["ipmi"] = o.Ipmi
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Result) UnmarshalJSON(data []byte) (err error) {
	varResult := _Result{}

	err = json.Unmarshal(data, &varResult)

	if err != nil {
		return err
	}

	*o = Result(varResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chassis")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "disks")
		delete(additionalProperties, "ipmi")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "network")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResult struct {
	value *Result
	isSet bool
}

func (v NullableResult) Get() *Result {
	return v.value
}

func (v *NullableResult) Set(val *Result) {
	v.value = val
	v.isSet = true
}

func (v NullableResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResult(val *Result) *NullableResult {
	return &NullableResult{value: val, isSet: true}
}

func (v NullableResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


