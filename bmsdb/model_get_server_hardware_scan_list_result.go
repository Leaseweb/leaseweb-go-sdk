/*
bmsdb

This documents the rest api bmsdb provides.

API version: v2
Contact: development-networkautomation@leaseweb.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmsdb

import (
	"encoding/json"
)

// checks if the GetServerHardwareScanListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerHardwareScanListResult{}

// GetServerHardwareScanListResult List of Hardware Scans executed in a server
type GetServerHardwareScanListResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	// Simple list of Hardware Scans
	HardwareScans []HardwareScan `json:"hardwareScans,omitempty"`
}

// NewGetServerHardwareScanListResult instantiates a new GetServerHardwareScanListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerHardwareScanListResult() *GetServerHardwareScanListResult {
	this := GetServerHardwareScanListResult{}
	return &this
}

// NewGetServerHardwareScanListResultWithDefaults instantiates a new GetServerHardwareScanListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerHardwareScanListResultWithDefaults() *GetServerHardwareScanListResult {
	this := GetServerHardwareScanListResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetServerHardwareScanListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerHardwareScanListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetServerHardwareScanListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetServerHardwareScanListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetHardwareScans returns the HardwareScans field value if set, zero value otherwise.
func (o *GetServerHardwareScanListResult) GetHardwareScans() []HardwareScan {
	if o == nil || IsNil(o.HardwareScans) {
		var ret []HardwareScan
		return ret
	}
	return o.HardwareScans
}

// GetHardwareScansOk returns a tuple with the HardwareScans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerHardwareScanListResult) GetHardwareScansOk() ([]HardwareScan, bool) {
	if o == nil || IsNil(o.HardwareScans) {
		return nil, false
	}
	return o.HardwareScans, true
}

// HasHardwareScans returns a boolean if a field has been set.
func (o *GetServerHardwareScanListResult) HasHardwareScans() bool {
	if o != nil && !IsNil(o.HardwareScans) {
		return true
	}

	return false
}

// SetHardwareScans gets a reference to the given []HardwareScan and assigns it to the HardwareScans field.
func (o *GetServerHardwareScanListResult) SetHardwareScans(v []HardwareScan) {
	o.HardwareScans = v
}

func (o GetServerHardwareScanListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerHardwareScanListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.HardwareScans) {
		toSerialize["hardwareScans"] = o.HardwareScans
	}
	return toSerialize, nil
}

type NullableGetServerHardwareScanListResult struct {
	value *GetServerHardwareScanListResult
	isSet bool
}

func (v NullableGetServerHardwareScanListResult) Get() *GetServerHardwareScanListResult {
	return v.value
}

func (v *NullableGetServerHardwareScanListResult) Set(val *GetServerHardwareScanListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerHardwareScanListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerHardwareScanListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerHardwareScanListResult(val *GetServerHardwareScanListResult) *NullableGetServerHardwareScanListResult {
	return &NullableGetServerHardwareScanListResult{value: val, isSet: true}
}

func (v NullableGetServerHardwareScanListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerHardwareScanListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

