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

// checks if the Metric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metric{}

// Metric struct for Metric
type Metric struct {
	// The metric unit that's used
	Unit *string `json:"unit,omitempty"`
	Values []MetricValue `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Metric Metric

// NewMetric instantiates a new Metric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetric() *Metric {
	this := Metric{}
	return &this
}

// NewMetricWithDefaults instantiates a new Metric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricWithDefaults() *Metric {
	this := Metric{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Metric) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Metric) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Metric) SetUnit(v string) {
	o.Unit = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Metric) GetValues() []MetricValue {
	if o == nil || IsNil(o.Values) {
		var ret []MetricValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetValuesOk() ([]MetricValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Metric) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []MetricValue and assigns it to the Values field.
func (o *Metric) SetValues(v []MetricValue) {
	o.Values = v
}

func (o Metric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Metric) UnmarshalJSON(data []byte) (err error) {
	varMetric := _Metric{}

	err = json.Unmarshal(data, &varMetric)

	if err != nil {
		return err
	}

	*o = Metric(varMetric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "unit")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetric struct {
	value *Metric
	isSet bool
}

func (v NullableMetric) Get() *Metric {
	return v.value
}

func (v *NullableMetric) Set(val *Metric) {
	v.value = val
	v.isSet = true
}

func (v NullableMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetric(val *Metric) *NullableMetric {
	return &NullableMetric{value: val, isSet: true}
}

func (v NullableMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


