/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the AutoNegotiation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoNegotiation{}

// AutoNegotiation struct for AutoNegotiation
type AutoNegotiation struct {
	Enabled *string `json:"enabled,omitempty"`
	Supported *string `json:"supported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoNegotiation AutoNegotiation

// NewAutoNegotiation instantiates a new AutoNegotiation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoNegotiation() *AutoNegotiation {
	this := AutoNegotiation{}
	return &this
}

// NewAutoNegotiationWithDefaults instantiates a new AutoNegotiation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoNegotiationWithDefaults() *AutoNegotiation {
	this := AutoNegotiation{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AutoNegotiation) GetEnabled() string {
	if o == nil || IsNil(o.Enabled) {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoNegotiation) GetEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AutoNegotiation) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *AutoNegotiation) SetEnabled(v string) {
	o.Enabled = &v
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *AutoNegotiation) GetSupported() string {
	if o == nil || IsNil(o.Supported) {
		var ret string
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoNegotiation) GetSupportedOk() (*string, bool) {
	if o == nil || IsNil(o.Supported) {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *AutoNegotiation) HasSupported() bool {
	if o != nil && !IsNil(o.Supported) {
		return true
	}

	return false
}

// SetSupported gets a reference to the given string and assigns it to the Supported field.
func (o *AutoNegotiation) SetSupported(v string) {
	o.Supported = &v
}

func (o AutoNegotiation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoNegotiation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Supported) {
		toSerialize["supported"] = o.Supported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoNegotiation) UnmarshalJSON(data []byte) (err error) {
	varAutoNegotiation := _AutoNegotiation{}

	err = json.Unmarshal(data, &varAutoNegotiation)

	if err != nil {
		return err
	}

	*o = AutoNegotiation(varAutoNegotiation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "supported")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoNegotiation struct {
	value *AutoNegotiation
	isSet bool
}

func (v NullableAutoNegotiation) Get() *AutoNegotiation {
	return v.value
}

func (v *NullableAutoNegotiation) Set(val *AutoNegotiation) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoNegotiation) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoNegotiation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoNegotiation(val *AutoNegotiation) *NullableAutoNegotiation {
	return &NullableAutoNegotiation{value: val, isSet: true}
}

func (v NullableAutoNegotiation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoNegotiation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

