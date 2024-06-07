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

// checks if the ScanHardwareOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScanHardwareOpts{}

// ScanHardwareOpts struct for ScanHardwareOpts
type ScanHardwareOpts struct {
	// Url which will receive callbacks
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// If set to `true`, server will be power cycled in order to complete the operation
	PowerCycle *bool `json:"powerCycle,omitempty"`
}

// NewScanHardwareOpts instantiates a new ScanHardwareOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScanHardwareOpts() *ScanHardwareOpts {
	this := ScanHardwareOpts{}
	var powerCycle bool = true
	this.PowerCycle = &powerCycle
	return &this
}

// NewScanHardwareOptsWithDefaults instantiates a new ScanHardwareOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScanHardwareOptsWithDefaults() *ScanHardwareOpts {
	this := ScanHardwareOpts{}
	var powerCycle bool = true
	this.PowerCycle = &powerCycle
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *ScanHardwareOpts) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanHardwareOpts) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *ScanHardwareOpts) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *ScanHardwareOpts) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetPowerCycle returns the PowerCycle field value if set, zero value otherwise.
func (o *ScanHardwareOpts) GetPowerCycle() bool {
	if o == nil || IsNil(o.PowerCycle) {
		var ret bool
		return ret
	}
	return *o.PowerCycle
}

// GetPowerCycleOk returns a tuple with the PowerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanHardwareOpts) GetPowerCycleOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerCycle) {
		return nil, false
	}
	return o.PowerCycle, true
}

// HasPowerCycle returns a boolean if a field has been set.
func (o *ScanHardwareOpts) HasPowerCycle() bool {
	if o != nil && !IsNil(o.PowerCycle) {
		return true
	}

	return false
}

// SetPowerCycle gets a reference to the given bool and assigns it to the PowerCycle field.
func (o *ScanHardwareOpts) SetPowerCycle(v bool) {
	o.PowerCycle = &v
}

func (o ScanHardwareOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScanHardwareOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallbackUrl) {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if !IsNil(o.PowerCycle) {
		toSerialize["powerCycle"] = o.PowerCycle
	}
	return toSerialize, nil
}

type NullableScanHardwareOpts struct {
	value *ScanHardwareOpts
	isSet bool
}

func (v NullableScanHardwareOpts) Get() *ScanHardwareOpts {
	return v.value
}

func (v *NullableScanHardwareOpts) Set(val *ScanHardwareOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableScanHardwareOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableScanHardwareOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScanHardwareOpts(val *ScanHardwareOpts) *NullableScanHardwareOpts {
	return &NullableScanHardwareOpts{value: val, isSet: true}
}

func (v NullableScanHardwareOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScanHardwareOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

