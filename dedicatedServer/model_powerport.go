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

// checks if the Powerport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Powerport{}

// Powerport struct for Powerport
type Powerport struct {
	Name *string `json:"name,omitempty"`
	Port *string `json:"port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Powerport Powerport

// NewPowerport instantiates a new Powerport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerport() *Powerport {
	this := Powerport{}
	return &this
}

// NewPowerportWithDefaults instantiates a new Powerport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerportWithDefaults() *Powerport {
	this := Powerport{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Powerport) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Powerport) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Powerport) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Powerport) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Powerport) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Powerport) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Powerport) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *Powerport) SetPort(v string) {
	o.Port = &v
}

func (o Powerport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Powerport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Powerport) UnmarshalJSON(data []byte) (err error) {
	varPowerport := _Powerport{}

	err = json.Unmarshal(data, &varPowerport)

	if err != nil {
		return err
	}

	*o = Powerport(varPowerport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerport struct {
	value *Powerport
	isSet bool
}

func (v NullablePowerport) Get() *Powerport {
	return v.value
}

func (v *NullablePowerport) Set(val *Powerport) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerport) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerport(val *Powerport) *NullablePowerport {
	return &NullablePowerport{value: val, isSet: true}
}

func (v NullablePowerport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


