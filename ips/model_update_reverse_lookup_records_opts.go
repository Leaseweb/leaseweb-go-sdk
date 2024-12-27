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

// checks if the UpdateReverseLookupRecordsOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateReverseLookupRecordsOpts{}

// UpdateReverseLookupRecordsOpts struct for UpdateReverseLookupRecordsOpts
type UpdateReverseLookupRecordsOpts struct {
	ReverseLookups []ReverseLookup `json:"reverseLookups"`
	AdditionalProperties map[string]interface{}
}

type _UpdateReverseLookupRecordsOpts UpdateReverseLookupRecordsOpts

// NewUpdateReverseLookupRecordsOpts instantiates a new UpdateReverseLookupRecordsOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReverseLookupRecordsOpts(reverseLookups []ReverseLookup) *UpdateReverseLookupRecordsOpts {
	this := UpdateReverseLookupRecordsOpts{}
	this.ReverseLookups = reverseLookups
	return &this
}

// NewUpdateReverseLookupRecordsOptsWithDefaults instantiates a new UpdateReverseLookupRecordsOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReverseLookupRecordsOptsWithDefaults() *UpdateReverseLookupRecordsOpts {
	this := UpdateReverseLookupRecordsOpts{}
	return &this
}

// GetReverseLookups returns the ReverseLookups field value
func (o *UpdateReverseLookupRecordsOpts) GetReverseLookups() []ReverseLookup {
	if o == nil {
		var ret []ReverseLookup
		return ret
	}

	return o.ReverseLookups
}

// GetReverseLookupsOk returns a tuple with the ReverseLookups field value
// and a boolean to check if the value has been set.
func (o *UpdateReverseLookupRecordsOpts) GetReverseLookupsOk() ([]ReverseLookup, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseLookups, true
}

// SetReverseLookups sets field value
func (o *UpdateReverseLookupRecordsOpts) SetReverseLookups(v []ReverseLookup) {
	o.ReverseLookups = v
}

func (o UpdateReverseLookupRecordsOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateReverseLookupRecordsOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reverseLookups"] = o.ReverseLookups

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateReverseLookupRecordsOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reverseLookups",
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

	varUpdateReverseLookupRecordsOpts := _UpdateReverseLookupRecordsOpts{}

	err = json.Unmarshal(data, &varUpdateReverseLookupRecordsOpts)

	if err != nil {
		return err
	}

	*o = UpdateReverseLookupRecordsOpts(varUpdateReverseLookupRecordsOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reverseLookups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateReverseLookupRecordsOpts struct {
	value *UpdateReverseLookupRecordsOpts
	isSet bool
}

func (v NullableUpdateReverseLookupRecordsOpts) Get() *UpdateReverseLookupRecordsOpts {
	return v.value
}

func (v *NullableUpdateReverseLookupRecordsOpts) Set(val *UpdateReverseLookupRecordsOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReverseLookupRecordsOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReverseLookupRecordsOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReverseLookupRecordsOpts(val *UpdateReverseLookupRecordsOpts) *NullableUpdateReverseLookupRecordsOpts {
	return &NullableUpdateReverseLookupRecordsOpts{value: val, isSet: true}
}

func (v NullableUpdateReverseLookupRecordsOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReverseLookupRecordsOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


