/*
bmusage

This documents the rest api bmusage api provides.

API version: v2
Contact: development-networkautomation@leaseweb.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmusage

import (
	"encoding/json"
)

// checks if the MetricValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricValues{}

// MetricValues Object containing all metrics
type MetricValues struct {
	UP_PUBLIC *Metric `json:"UP_PUBLIC,omitempty"`
	DOWN_PUBLIC *Metric `json:"DOWN_PUBLIC,omitempty"`
}

// NewMetricValues instantiates a new MetricValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricValues() *MetricValues {
	this := MetricValues{}
	return &this
}

// NewMetricValuesWithDefaults instantiates a new MetricValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricValuesWithDefaults() *MetricValues {
	this := MetricValues{}
	return &this
}

// GetUP_PUBLIC returns the UP_PUBLIC field value if set, zero value otherwise.
func (o *MetricValues) GetUP_PUBLIC() Metric {
	if o == nil || IsNil(o.UP_PUBLIC) {
		var ret Metric
		return ret
	}
	return *o.UP_PUBLIC
}

// GetUP_PUBLICOk returns a tuple with the UP_PUBLIC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricValues) GetUP_PUBLICOk() (*Metric, bool) {
	if o == nil || IsNil(o.UP_PUBLIC) {
		return nil, false
	}
	return o.UP_PUBLIC, true
}

// HasUP_PUBLIC returns a boolean if a field has been set.
func (o *MetricValues) HasUP_PUBLIC() bool {
	if o != nil && !IsNil(o.UP_PUBLIC) {
		return true
	}

	return false
}

// SetUP_PUBLIC gets a reference to the given Metric and assigns it to the UP_PUBLIC field.
func (o *MetricValues) SetUP_PUBLIC(v Metric) {
	o.UP_PUBLIC = &v
}

// GetDOWN_PUBLIC returns the DOWN_PUBLIC field value if set, zero value otherwise.
func (o *MetricValues) GetDOWN_PUBLIC() Metric {
	if o == nil || IsNil(o.DOWN_PUBLIC) {
		var ret Metric
		return ret
	}
	return *o.DOWN_PUBLIC
}

// GetDOWN_PUBLICOk returns a tuple with the DOWN_PUBLIC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricValues) GetDOWN_PUBLICOk() (*Metric, bool) {
	if o == nil || IsNil(o.DOWN_PUBLIC) {
		return nil, false
	}
	return o.DOWN_PUBLIC, true
}

// HasDOWN_PUBLIC returns a boolean if a field has been set.
func (o *MetricValues) HasDOWN_PUBLIC() bool {
	if o != nil && !IsNil(o.DOWN_PUBLIC) {
		return true
	}

	return false
}

// SetDOWN_PUBLIC gets a reference to the given Metric and assigns it to the DOWN_PUBLIC field.
func (o *MetricValues) SetDOWN_PUBLIC(v Metric) {
	o.DOWN_PUBLIC = &v
}

func (o MetricValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UP_PUBLIC) {
		toSerialize["UP_PUBLIC"] = o.UP_PUBLIC
	}
	if !IsNil(o.DOWN_PUBLIC) {
		toSerialize["DOWN_PUBLIC"] = o.DOWN_PUBLIC
	}
	return toSerialize, nil
}

type NullableMetricValues struct {
	value *MetricValues
	isSet bool
}

func (v NullableMetricValues) Get() *MetricValues {
	return v.value
}

func (v *NullableMetricValues) Set(val *MetricValues) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricValues) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricValues(val *MetricValues) *NullableMetricValues {
	return &NullableMetricValues{value: val, isSet: true}
}

func (v NullableMetricValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


