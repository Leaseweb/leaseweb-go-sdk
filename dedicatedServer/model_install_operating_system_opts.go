/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedServer

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InstallOperatingSystemOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallOperatingSystemOpts{}

// InstallOperatingSystemOpts struct for InstallOperatingSystemOpts
type InstallOperatingSystemOpts struct {
	// Url which will receive callbacks when the installation is finished or failed
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// Control panel identifier
	ControlPanelId *string `json:"controlPanelId,omitempty"`
	// Block devices in a disk set in which the partitions will be installed. Supported values are any disk set id, SATA_SAS or NVME
	Device *string `json:"device,omitempty"`
	// Hostname to be used in your installation
	Hostname *string `json:"hostname,omitempty"`
	// Operating system identifier
	OperatingSystemId string `json:"operatingSystemId"`
	// Array of partition objects that should be installed per partition
	Partitions []Partition `json:"partitions,omitempty"`
	// Server root password. If not provided, it would be automatically generated
	Password *string `json:"password,omitempty"`
	// Base64 Encoded string containing a valid bash script to be run right after the installation
	PostInstallScript *string `json:"postInstallScript,omitempty"`
	// If true, allows system reboots to happen automatically within the process. Otherwise, you should do them manually
	PowerCycle *bool `json:"powerCycle,omitempty"`
	Raid *Raid `json:"raid,omitempty"`
	// List of public sshKeys to be setup in your installation, separated by new lines
	SshKeys *string `json:"sshKeys,omitempty"`
	// Timezone represented as Geographical_Area/City
	Timezone *string `json:"timezone,omitempty"`
}

type _InstallOperatingSystemOpts InstallOperatingSystemOpts

// NewInstallOperatingSystemOpts instantiates a new InstallOperatingSystemOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallOperatingSystemOpts(operatingSystemId string) *InstallOperatingSystemOpts {
	this := InstallOperatingSystemOpts{}
	this.OperatingSystemId = operatingSystemId
	return &this
}

// NewInstallOperatingSystemOptsWithDefaults instantiates a new InstallOperatingSystemOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallOperatingSystemOptsWithDefaults() *InstallOperatingSystemOpts {
	this := InstallOperatingSystemOpts{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *InstallOperatingSystemOpts) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetControlPanelId returns the ControlPanelId field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetControlPanelId() string {
	if o == nil || IsNil(o.ControlPanelId) {
		var ret string
		return ret
	}
	return *o.ControlPanelId
}

// GetControlPanelIdOk returns a tuple with the ControlPanelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetControlPanelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ControlPanelId) {
		return nil, false
	}
	return o.ControlPanelId, true
}

// HasControlPanelId returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasControlPanelId() bool {
	if o != nil && !IsNil(o.ControlPanelId) {
		return true
	}

	return false
}

// SetControlPanelId gets a reference to the given string and assigns it to the ControlPanelId field.
func (o *InstallOperatingSystemOpts) SetControlPanelId(v string) {
	o.ControlPanelId = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *InstallOperatingSystemOpts) SetDevice(v string) {
	o.Device = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *InstallOperatingSystemOpts) SetHostname(v string) {
	o.Hostname = &v
}

// GetOperatingSystemId returns the OperatingSystemId field value
func (o *InstallOperatingSystemOpts) GetOperatingSystemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystemId
}

// GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field value
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetOperatingSystemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystemId, true
}

// SetOperatingSystemId sets field value
func (o *InstallOperatingSystemOpts) SetOperatingSystemId(v string) {
	o.OperatingSystemId = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetPartitions() []Partition {
	if o == nil || IsNil(o.Partitions) {
		var ret []Partition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetPartitionsOk() ([]Partition, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *InstallOperatingSystemOpts) SetPartitions(v []Partition) {
	o.Partitions = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InstallOperatingSystemOpts) SetPassword(v string) {
	o.Password = &v
}

// GetPostInstallScript returns the PostInstallScript field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetPostInstallScript() string {
	if o == nil || IsNil(o.PostInstallScript) {
		var ret string
		return ret
	}
	return *o.PostInstallScript
}

// GetPostInstallScriptOk returns a tuple with the PostInstallScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetPostInstallScriptOk() (*string, bool) {
	if o == nil || IsNil(o.PostInstallScript) {
		return nil, false
	}
	return o.PostInstallScript, true
}

// HasPostInstallScript returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasPostInstallScript() bool {
	if o != nil && !IsNil(o.PostInstallScript) {
		return true
	}

	return false
}

// SetPostInstallScript gets a reference to the given string and assigns it to the PostInstallScript field.
func (o *InstallOperatingSystemOpts) SetPostInstallScript(v string) {
	o.PostInstallScript = &v
}

// GetPowerCycle returns the PowerCycle field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetPowerCycle() bool {
	if o == nil || IsNil(o.PowerCycle) {
		var ret bool
		return ret
	}
	return *o.PowerCycle
}

// GetPowerCycleOk returns a tuple with the PowerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetPowerCycleOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerCycle) {
		return nil, false
	}
	return o.PowerCycle, true
}

// HasPowerCycle returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasPowerCycle() bool {
	if o != nil && !IsNil(o.PowerCycle) {
		return true
	}

	return false
}

// SetPowerCycle gets a reference to the given bool and assigns it to the PowerCycle field.
func (o *InstallOperatingSystemOpts) SetPowerCycle(v bool) {
	o.PowerCycle = &v
}

// GetRaid returns the Raid field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetRaid() Raid {
	if o == nil || IsNil(o.Raid) {
		var ret Raid
		return ret
	}
	return *o.Raid
}

// GetRaidOk returns a tuple with the Raid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetRaidOk() (*Raid, bool) {
	if o == nil || IsNil(o.Raid) {
		return nil, false
	}
	return o.Raid, true
}

// HasRaid returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasRaid() bool {
	if o != nil && !IsNil(o.Raid) {
		return true
	}

	return false
}

// SetRaid gets a reference to the given Raid and assigns it to the Raid field.
func (o *InstallOperatingSystemOpts) SetRaid(v Raid) {
	o.Raid = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetSshKeys() string {
	if o == nil || IsNil(o.SshKeys) {
		var ret string
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetSshKeysOk() (*string, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given string and assigns it to the SshKeys field.
func (o *InstallOperatingSystemOpts) SetSshKeys(v string) {
	o.SshKeys = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InstallOperatingSystemOpts) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallOperatingSystemOpts) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InstallOperatingSystemOpts) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InstallOperatingSystemOpts) SetTimezone(v string) {
	o.Timezone = &v
}

func (o InstallOperatingSystemOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallOperatingSystemOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallbackUrl) {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if !IsNil(o.ControlPanelId) {
		toSerialize["controlPanelId"] = o.ControlPanelId
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	toSerialize["operatingSystemId"] = o.OperatingSystemId
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PostInstallScript) {
		toSerialize["postInstallScript"] = o.PostInstallScript
	}
	if !IsNil(o.PowerCycle) {
		toSerialize["powerCycle"] = o.PowerCycle
	}
	if !IsNil(o.Raid) {
		toSerialize["raid"] = o.Raid
	}
	if !IsNil(o.SshKeys) {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

func (o *InstallOperatingSystemOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operatingSystemId",
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

	varInstallOperatingSystemOpts := _InstallOperatingSystemOpts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstallOperatingSystemOpts)

	if err != nil {
		return err
	}

	*o = InstallOperatingSystemOpts(varInstallOperatingSystemOpts)

	return err
}

type NullableInstallOperatingSystemOpts struct {
	value *InstallOperatingSystemOpts
	isSet bool
}

func (v NullableInstallOperatingSystemOpts) Get() *InstallOperatingSystemOpts {
	return v.value
}

func (v *NullableInstallOperatingSystemOpts) Set(val *InstallOperatingSystemOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallOperatingSystemOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallOperatingSystemOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallOperatingSystemOpts(val *InstallOperatingSystemOpts) *NullableInstallOperatingSystemOpts {
	return &NullableInstallOperatingSystemOpts{value: val, isSet: true}
}

func (v NullableInstallOperatingSystemOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallOperatingSystemOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


