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

// checks if the Cpu1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cpu1{}

// Cpu1 struct for Cpu1
type Cpu1 struct {
	Capabilities *CpuCapabilities `json:"capabilities,omitempty"`
	Description *string `json:"description,omitempty"`
	Hz *string `json:"hz,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Settings *CpuSettings `json:"settings,omitempty"`
	Slot *string `json:"slot,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewCpu1 instantiates a new Cpu1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpu1() *Cpu1 {
	this := Cpu1{}
	return &this
}

// NewCpu1WithDefaults instantiates a new Cpu1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpu1WithDefaults() *Cpu1 {
	this := Cpu1{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *Cpu1) GetCapabilities() CpuCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret CpuCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu1) GetCapabilitiesOk() (*CpuCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *Cpu1) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given CpuCapabilities and assigns it to the Capabilities field.
func (o *Cpu1) SetCapabilities(v CpuCapabilities) {
	o.Capabilities = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Cpu1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Cpu1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Cpu1) SetDescription(v string) {
	o.Description = &v
}

// GetHz returns the Hz field value if set, zero value otherwise.
func (o *Cpu1) GetHz() string {
	if o == nil || IsNil(o.Hz) {
		var ret string
		return ret
	}
	return *o.Hz
}

// GetHzOk returns a tuple with the Hz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu1) GetHzOk() (*string, bool) {
	if o == nil || IsNil(o.Hz) {
		return nil, false
	}
	return o.Hz, true
}

// HasHz returns a boolean if a field has been set.
func (o *Cpu1) HasHz() bool {
	if o != nil && !IsNil(o.Hz) {
		return true
	}

	return false
}

// SetHz gets a reference to the given string and assigns it to the Hz field.
func (o *Cpu1) SetHz(v string) {
	o.Hz = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Cpu1) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu1) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Cpu1) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Cpu1) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Cpu1) GetSettings() CpuSettings {
	if o == nil || IsNil(o.Settings) {
		var ret CpuSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu1) GetSettingsOk() (*CpuSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Cpu1) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given CpuSettings and assigns it to the Settings field.
func (o *Cpu1) SetSettings(v CpuSettings) {
	o.Settings = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *Cpu1) GetSlot() string {
	if o == nil || IsNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu1) GetSlotOk() (*string, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *Cpu1) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *Cpu1) SetSlot(v string) {
	o.Slot = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *Cpu1) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu1) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *Cpu1) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *Cpu1) SetVendor(v string) {
	o.Vendor = &v
}

func (o Cpu1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cpu1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Hz) {
		toSerialize["hz"] = o.Hz
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	return toSerialize, nil
}

type NullableCpu1 struct {
	value *Cpu1
	isSet bool
}

func (v NullableCpu1) Get() *Cpu1 {
	return v.value
}

func (v *NullableCpu1) Set(val *Cpu1) {
	v.value = val
	v.isSet = true
}

func (v NullableCpu1) IsSet() bool {
	return v.isSet
}

func (v *NullableCpu1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpu1(val *Cpu1) *NullableCpu1 {
	return &NullableCpu1{value: val, isSet: true}
}

func (v NullableCpu1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpu1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


