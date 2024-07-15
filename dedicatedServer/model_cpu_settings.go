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

// checks if the CpuSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpuSettings{}

// CpuSettings struct for CpuSettings
type CpuSettings struct {
	Cores *string `json:"cores,omitempty"`
	Enabledcores *string `json:"enabledcores,omitempty"`
	Threads *string `json:"threads,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CpuSettings CpuSettings

// NewCpuSettings instantiates a new CpuSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuSettings() *CpuSettings {
	this := CpuSettings{}
	return &this
}

// NewCpuSettingsWithDefaults instantiates a new CpuSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuSettingsWithDefaults() *CpuSettings {
	this := CpuSettings{}
	return &this
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *CpuSettings) GetCores() string {
	if o == nil || IsNil(o.Cores) {
		var ret string
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuSettings) GetCoresOk() (*string, bool) {
	if o == nil || IsNil(o.Cores) {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *CpuSettings) HasCores() bool {
	if o != nil && !IsNil(o.Cores) {
		return true
	}

	return false
}

// SetCores gets a reference to the given string and assigns it to the Cores field.
func (o *CpuSettings) SetCores(v string) {
	o.Cores = &v
}

// GetEnabledcores returns the Enabledcores field value if set, zero value otherwise.
func (o *CpuSettings) GetEnabledcores() string {
	if o == nil || IsNil(o.Enabledcores) {
		var ret string
		return ret
	}
	return *o.Enabledcores
}

// GetEnabledcoresOk returns a tuple with the Enabledcores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuSettings) GetEnabledcoresOk() (*string, bool) {
	if o == nil || IsNil(o.Enabledcores) {
		return nil, false
	}
	return o.Enabledcores, true
}

// HasEnabledcores returns a boolean if a field has been set.
func (o *CpuSettings) HasEnabledcores() bool {
	if o != nil && !IsNil(o.Enabledcores) {
		return true
	}

	return false
}

// SetEnabledcores gets a reference to the given string and assigns it to the Enabledcores field.
func (o *CpuSettings) SetEnabledcores(v string) {
	o.Enabledcores = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *CpuSettings) GetThreads() string {
	if o == nil || IsNil(o.Threads) {
		var ret string
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuSettings) GetThreadsOk() (*string, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *CpuSettings) HasThreads() bool {
	if o != nil && !IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given string and assigns it to the Threads field.
func (o *CpuSettings) SetThreads(v string) {
	o.Threads = &v
}

func (o CpuSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	if !IsNil(o.Enabledcores) {
		toSerialize["enabledcores"] = o.Enabledcores
	}
	if !IsNil(o.Threads) {
		toSerialize["threads"] = o.Threads
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CpuSettings) UnmarshalJSON(data []byte) (err error) {
	varCpuSettings := _CpuSettings{}

	err = json.Unmarshal(data, &varCpuSettings)

	if err != nil {
		return err
	}

	*o = CpuSettings(varCpuSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cores")
		delete(additionalProperties, "enabledcores")
		delete(additionalProperties, "threads")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCpuSettings struct {
	value *CpuSettings
	isSet bool
}

func (v NullableCpuSettings) Get() *CpuSettings {
	return v.value
}

func (v *NullableCpuSettings) Set(val *CpuSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuSettings(val *CpuSettings) *NullableCpuSettings {
	return &NullableCpuSettings{value: val, isSet: true}
}

func (v NullableCpuSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


