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

// checks if the Metrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metrics{}

// Metrics struct for Metrics
type Metrics struct {
	Metadata *MetricsMetadata `json:"_metadata,omitempty"`
	Metrics *MetricValues `json:"metrics,omitempty"`
}

// NewMetrics instantiates a new Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetrics() *Metrics {
	this := Metrics{}
	return &this
}

// NewMetricsWithDefaults instantiates a new Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsWithDefaults() *Metrics {
	this := Metrics{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Metrics) GetMetadata() MetricsMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret MetricsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetMetadataOk() (*MetricsMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Metrics) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MetricsMetadata and assigns it to the Metadata field.
func (o *Metrics) SetMetadata(v MetricsMetadata) {
	o.Metadata = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Metrics) GetMetrics() MetricValues {
	if o == nil || IsNil(o.Metrics) {
		var ret MetricValues
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetMetricsOk() (*MetricValues, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Metrics) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given MetricValues and assigns it to the Metrics field.
func (o *Metrics) SetMetrics(v MetricValues) {
	o.Metrics = &v
}

func (o Metrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableMetrics struct {
	value *Metrics
	isSet bool
}

func (v NullableMetrics) Get() *Metrics {
	return v.value
}

func (v *NullableMetrics) Set(val *Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetrics(val *Metrics) *NullableMetrics {
	return &NullableMetrics{value: val, isSet: true}
}

func (v NullableMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

