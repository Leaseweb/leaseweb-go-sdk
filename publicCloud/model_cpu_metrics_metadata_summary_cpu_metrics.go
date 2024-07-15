/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
)

// checks if the CpuMetricsMetadataSummaryCpuMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpuMetricsMetadataSummaryCpuMetrics{}

// CpuMetricsMetadataSummaryCpuMetrics struct for CpuMetricsMetadataSummaryCpuMetrics
type CpuMetricsMetadataSummaryCpuMetrics struct {
	// Average CPU based on the amount of grouped data points, in percentage
	Average *string `json:"average,omitempty"`
	// Expected CPU given the average times the amount of days between the `from` and `to` dates, in percentage
	Expected *string `json:"expected,omitempty"`
	Peak *CpuMetricsMetadataSummaryPeak `json:"peak,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CpuMetricsMetadataSummaryCpuMetrics CpuMetricsMetadataSummaryCpuMetrics

// NewCpuMetricsMetadataSummaryCpuMetrics instantiates a new CpuMetricsMetadataSummaryCpuMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuMetricsMetadataSummaryCpuMetrics() *CpuMetricsMetadataSummaryCpuMetrics {
	this := CpuMetricsMetadataSummaryCpuMetrics{}
	return &this
}

// NewCpuMetricsMetadataSummaryCpuMetricsWithDefaults instantiates a new CpuMetricsMetadataSummaryCpuMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuMetricsMetadataSummaryCpuMetricsWithDefaults() *CpuMetricsMetadataSummaryCpuMetrics {
	this := CpuMetricsMetadataSummaryCpuMetrics{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *CpuMetricsMetadataSummaryCpuMetrics) GetAverage() string {
	if o == nil || IsNil(o.Average) {
		var ret string
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadataSummaryCpuMetrics) GetAverageOk() (*string, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *CpuMetricsMetadataSummaryCpuMetrics) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given string and assigns it to the Average field.
func (o *CpuMetricsMetadataSummaryCpuMetrics) SetAverage(v string) {
	o.Average = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *CpuMetricsMetadataSummaryCpuMetrics) GetExpected() string {
	if o == nil || IsNil(o.Expected) {
		var ret string
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadataSummaryCpuMetrics) GetExpectedOk() (*string, bool) {
	if o == nil || IsNil(o.Expected) {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *CpuMetricsMetadataSummaryCpuMetrics) HasExpected() bool {
	if o != nil && !IsNil(o.Expected) {
		return true
	}

	return false
}

// SetExpected gets a reference to the given string and assigns it to the Expected field.
func (o *CpuMetricsMetadataSummaryCpuMetrics) SetExpected(v string) {
	o.Expected = &v
}

// GetPeak returns the Peak field value if set, zero value otherwise.
func (o *CpuMetricsMetadataSummaryCpuMetrics) GetPeak() CpuMetricsMetadataSummaryPeak {
	if o == nil || IsNil(o.Peak) {
		var ret CpuMetricsMetadataSummaryPeak
		return ret
	}
	return *o.Peak
}

// GetPeakOk returns a tuple with the Peak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadataSummaryCpuMetrics) GetPeakOk() (*CpuMetricsMetadataSummaryPeak, bool) {
	if o == nil || IsNil(o.Peak) {
		return nil, false
	}
	return o.Peak, true
}

// HasPeak returns a boolean if a field has been set.
func (o *CpuMetricsMetadataSummaryCpuMetrics) HasPeak() bool {
	if o != nil && !IsNil(o.Peak) {
		return true
	}

	return false
}

// SetPeak gets a reference to the given CpuMetricsMetadataSummaryPeak and assigns it to the Peak field.
func (o *CpuMetricsMetadataSummaryCpuMetrics) SetPeak(v CpuMetricsMetadataSummaryPeak) {
	o.Peak = &v
}

func (o CpuMetricsMetadataSummaryCpuMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuMetricsMetadataSummaryCpuMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !IsNil(o.Expected) {
		toSerialize["expected"] = o.Expected
	}
	if !IsNil(o.Peak) {
		toSerialize["peak"] = o.Peak
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CpuMetricsMetadataSummaryCpuMetrics) UnmarshalJSON(data []byte) (err error) {
	varCpuMetricsMetadataSummaryCpuMetrics := _CpuMetricsMetadataSummaryCpuMetrics{}

	err = json.Unmarshal(data, &varCpuMetricsMetadataSummaryCpuMetrics)

	if err != nil {
		return err
	}

	*o = CpuMetricsMetadataSummaryCpuMetrics(varCpuMetricsMetadataSummaryCpuMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "average")
		delete(additionalProperties, "expected")
		delete(additionalProperties, "peak")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCpuMetricsMetadataSummaryCpuMetrics struct {
	value *CpuMetricsMetadataSummaryCpuMetrics
	isSet bool
}

func (v NullableCpuMetricsMetadataSummaryCpuMetrics) Get() *CpuMetricsMetadataSummaryCpuMetrics {
	return v.value
}

func (v *NullableCpuMetricsMetadataSummaryCpuMetrics) Set(val *CpuMetricsMetadataSummaryCpuMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuMetricsMetadataSummaryCpuMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuMetricsMetadataSummaryCpuMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuMetricsMetadataSummaryCpuMetrics(val *CpuMetricsMetadataSummaryCpuMetrics) *NullableCpuMetricsMetadataSummaryCpuMetrics {
	return &NullableCpuMetricsMetadataSummaryCpuMetrics{value: val, isSet: true}
}

func (v NullableCpuMetricsMetadataSummaryCpuMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuMetricsMetadataSummaryCpuMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


