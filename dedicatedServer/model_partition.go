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

// checks if the Partition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Partition{}

// Partition struct for Partition
type Partition struct {
	// File system in which partition would be mounted
	Filesystem *string `json:"filesystem,omitempty"`
	// The partition mount point (eg `/home`). Mandatory for the root partition (`/`) and not intended to be used in swap partition
	Mountpoint *string `json:"mountpoint,omitempty"`
	// Size of the partition (Normally in MB, but this is OS-specific)
	Size *string `json:"size,omitempty"`
}

// NewPartition instantiates a new Partition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartition() *Partition {
	this := Partition{}
	return &this
}

// NewPartitionWithDefaults instantiates a new Partition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionWithDefaults() *Partition {
	this := Partition{}
	return &this
}

// GetFilesystem returns the Filesystem field value if set, zero value otherwise.
func (o *Partition) GetFilesystem() string {
	if o == nil || IsNil(o.Filesystem) {
		var ret string
		return ret
	}
	return *o.Filesystem
}

// GetFilesystemOk returns a tuple with the Filesystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetFilesystemOk() (*string, bool) {
	if o == nil || IsNil(o.Filesystem) {
		return nil, false
	}
	return o.Filesystem, true
}

// HasFilesystem returns a boolean if a field has been set.
func (o *Partition) HasFilesystem() bool {
	if o != nil && !IsNil(o.Filesystem) {
		return true
	}

	return false
}

// SetFilesystem gets a reference to the given string and assigns it to the Filesystem field.
func (o *Partition) SetFilesystem(v string) {
	o.Filesystem = &v
}

// GetMountpoint returns the Mountpoint field value if set, zero value otherwise.
func (o *Partition) GetMountpoint() string {
	if o == nil || IsNil(o.Mountpoint) {
		var ret string
		return ret
	}
	return *o.Mountpoint
}

// GetMountpointOk returns a tuple with the Mountpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetMountpointOk() (*string, bool) {
	if o == nil || IsNil(o.Mountpoint) {
		return nil, false
	}
	return o.Mountpoint, true
}

// HasMountpoint returns a boolean if a field has been set.
func (o *Partition) HasMountpoint() bool {
	if o != nil && !IsNil(o.Mountpoint) {
		return true
	}

	return false
}

// SetMountpoint gets a reference to the given string and assigns it to the Mountpoint field.
func (o *Partition) SetMountpoint(v string) {
	o.Mountpoint = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Partition) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Partition) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Partition) SetSize(v string) {
	o.Size = &v
}

func (o Partition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Partition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filesystem) {
		toSerialize["filesystem"] = o.Filesystem
	}
	if !IsNil(o.Mountpoint) {
		toSerialize["mountpoint"] = o.Mountpoint
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullablePartition struct {
	value *Partition
	isSet bool
}

func (v NullablePartition) Get() *Partition {
	return v.value
}

func (v *NullablePartition) Set(val *Partition) {
	v.value = val
	v.isSet = true
}

func (v NullablePartition) IsSet() bool {
	return v.isSet
}

func (v *NullablePartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartition(val *Partition) *NullablePartition {
	return &NullablePartition{value: val, isSet: true}
}

func (v NullablePartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

