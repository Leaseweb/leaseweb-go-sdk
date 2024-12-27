/*
LeaseWeb API for IP address management

> The base URL for this API is: **https://api.leaseweb.com/ipMgmt/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ips

import (
	"encoding/json"
	"fmt"
)

// checks if the GetReverseLookupRecordListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReverseLookupRecordListResult{}

// GetReverseLookupRecordListResult Object containing a list of reverse lookup values set for ips in a specific range
type GetReverseLookupRecordListResult struct {
	ReverseLookups []ReverseLookup `json:"reverseLookups"`
	Metadata Metadata `json:"_metadata"`
	AdditionalProperties map[string]interface{}
}

type _GetReverseLookupRecordListResult GetReverseLookupRecordListResult

// NewGetReverseLookupRecordListResult instantiates a new GetReverseLookupRecordListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReverseLookupRecordListResult(reverseLookups []ReverseLookup, metadata Metadata) *GetReverseLookupRecordListResult {
	this := GetReverseLookupRecordListResult{}
	this.ReverseLookups = reverseLookups
	this.Metadata = metadata
	return &this
}

// NewGetReverseLookupRecordListResultWithDefaults instantiates a new GetReverseLookupRecordListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReverseLookupRecordListResultWithDefaults() *GetReverseLookupRecordListResult {
	this := GetReverseLookupRecordListResult{}
	return &this
}

// GetReverseLookups returns the ReverseLookups field value
func (o *GetReverseLookupRecordListResult) GetReverseLookups() []ReverseLookup {
	if o == nil {
		var ret []ReverseLookup
		return ret
	}

	return o.ReverseLookups
}

// GetReverseLookupsOk returns a tuple with the ReverseLookups field value
// and a boolean to check if the value has been set.
func (o *GetReverseLookupRecordListResult) GetReverseLookupsOk() ([]ReverseLookup, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseLookups, true
}

// SetReverseLookups sets field value
func (o *GetReverseLookupRecordListResult) SetReverseLookups(v []ReverseLookup) {
	o.ReverseLookups = v
}

// GetMetadata returns the Metadata field value
func (o *GetReverseLookupRecordListResult) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *GetReverseLookupRecordListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *GetReverseLookupRecordListResult) SetMetadata(v Metadata) {
	o.Metadata = v
}

func (o GetReverseLookupRecordListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReverseLookupRecordListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reverseLookups"] = o.ReverseLookups
	toSerialize["_metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetReverseLookupRecordListResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reverseLookups",
		"_metadata",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetReverseLookupRecordListResult := _GetReverseLookupRecordListResult{}

	err = json.Unmarshal(data, &varGetReverseLookupRecordListResult)

	if err != nil {
		return err
	}

	*o = GetReverseLookupRecordListResult(varGetReverseLookupRecordListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reverseLookups")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetReverseLookupRecordListResult struct {
	value *GetReverseLookupRecordListResult
	isSet bool
}

func (v NullableGetReverseLookupRecordListResult) Get() *GetReverseLookupRecordListResult {
	return v.value
}

func (v *NullableGetReverseLookupRecordListResult) Set(val *GetReverseLookupRecordListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReverseLookupRecordListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReverseLookupRecordListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReverseLookupRecordListResult(val *GetReverseLookupRecordListResult) *NullableGetReverseLookupRecordListResult {
	return &NullableGetReverseLookupRecordListResult{value: val, isSet: true}
}

func (v NullableGetReverseLookupRecordListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReverseLookupRecordListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

