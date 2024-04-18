/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
Contact: development-networkautomation@leaseweb.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the GetNetworkEquipmentIpListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkEquipmentIpListResult{}

// GetNetworkEquipmentIpListResult struct for GetNetworkEquipmentIpListResult
type GetNetworkEquipmentIpListResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	// An array of IP addresses
	Ips []Ip `json:"ips,omitempty"`
}

// NewGetNetworkEquipmentIpListResult instantiates a new GetNetworkEquipmentIpListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkEquipmentIpListResult() *GetNetworkEquipmentIpListResult {
	this := GetNetworkEquipmentIpListResult{}
	return &this
}

// NewGetNetworkEquipmentIpListResultWithDefaults instantiates a new GetNetworkEquipmentIpListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkEquipmentIpListResultWithDefaults() *GetNetworkEquipmentIpListResult {
	this := GetNetworkEquipmentIpListResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetNetworkEquipmentIpListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEquipmentIpListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetNetworkEquipmentIpListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetNetworkEquipmentIpListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *GetNetworkEquipmentIpListResult) GetIps() []Ip {
	if o == nil || IsNil(o.Ips) {
		var ret []Ip
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEquipmentIpListResult) GetIpsOk() ([]Ip, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *GetNetworkEquipmentIpListResult) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []Ip and assigns it to the Ips field.
func (o *GetNetworkEquipmentIpListResult) SetIps(v []Ip) {
	o.Ips = v
}

func (o GetNetworkEquipmentIpListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkEquipmentIpListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	return toSerialize, nil
}

type NullableGetNetworkEquipmentIpListResult struct {
	value *GetNetworkEquipmentIpListResult
	isSet bool
}

func (v NullableGetNetworkEquipmentIpListResult) Get() *GetNetworkEquipmentIpListResult {
	return v.value
}

func (v *NullableGetNetworkEquipmentIpListResult) Set(val *GetNetworkEquipmentIpListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkEquipmentIpListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkEquipmentIpListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkEquipmentIpListResult(val *GetNetworkEquipmentIpListResult) *NullableGetNetworkEquipmentIpListResult {
	return &NullableGetNetworkEquipmentIpListResult{value: val, isSet: true}
}

func (v NullableGetNetworkEquipmentIpListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkEquipmentIpListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


