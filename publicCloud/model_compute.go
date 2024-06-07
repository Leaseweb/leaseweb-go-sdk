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

// checks if the Compute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Compute{}

// Compute struct for Compute
type Compute struct {
	HourlyPrice *string `json:"hourlyPrice,omitempty"`
	MonthlyPrice *string `json:"monthlyPrice,omitempty"`
}

// NewCompute instantiates a new Compute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompute() *Compute {
	this := Compute{}
	return &this
}

// NewComputeWithDefaults instantiates a new Compute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeWithDefaults() *Compute {
	this := Compute{}
	return &this
}

// GetHourlyPrice returns the HourlyPrice field value if set, zero value otherwise.
func (o *Compute) GetHourlyPrice() string {
	if o == nil || IsNil(o.HourlyPrice) {
		var ret string
		return ret
	}
	return *o.HourlyPrice
}

// GetHourlyPriceOk returns a tuple with the HourlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Compute) GetHourlyPriceOk() (*string, bool) {
	if o == nil || IsNil(o.HourlyPrice) {
		return nil, false
	}
	return o.HourlyPrice, true
}

// HasHourlyPrice returns a boolean if a field has been set.
func (o *Compute) HasHourlyPrice() bool {
	if o != nil && !IsNil(o.HourlyPrice) {
		return true
	}

	return false
}

// SetHourlyPrice gets a reference to the given string and assigns it to the HourlyPrice field.
func (o *Compute) SetHourlyPrice(v string) {
	o.HourlyPrice = &v
}

// GetMonthlyPrice returns the MonthlyPrice field value if set, zero value otherwise.
func (o *Compute) GetMonthlyPrice() string {
	if o == nil || IsNil(o.MonthlyPrice) {
		var ret string
		return ret
	}
	return *o.MonthlyPrice
}

// GetMonthlyPriceOk returns a tuple with the MonthlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Compute) GetMonthlyPriceOk() (*string, bool) {
	if o == nil || IsNil(o.MonthlyPrice) {
		return nil, false
	}
	return o.MonthlyPrice, true
}

// HasMonthlyPrice returns a boolean if a field has been set.
func (o *Compute) HasMonthlyPrice() bool {
	if o != nil && !IsNil(o.MonthlyPrice) {
		return true
	}

	return false
}

// SetMonthlyPrice gets a reference to the given string and assigns it to the MonthlyPrice field.
func (o *Compute) SetMonthlyPrice(v string) {
	o.MonthlyPrice = &v
}

func (o Compute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Compute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HourlyPrice) {
		toSerialize["hourlyPrice"] = o.HourlyPrice
	}
	if !IsNil(o.MonthlyPrice) {
		toSerialize["monthlyPrice"] = o.MonthlyPrice
	}
	return toSerialize, nil
}

type NullableCompute struct {
	value *Compute
	isSet bool
}

func (v NullableCompute) Get() *Compute {
	return v.value
}

func (v *NullableCompute) Set(val *Compute) {
	v.value = val
	v.isSet = true
}

func (v NullableCompute) IsSet() bool {
	return v.isSet
}

func (v *NullableCompute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompute(val *Compute) *NullableCompute {
	return &NullableCompute{value: val, isSet: true}
}

func (v NullableCompute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


