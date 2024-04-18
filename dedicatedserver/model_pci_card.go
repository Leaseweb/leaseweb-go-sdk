/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
Contact: development-networkautomation@leaseweb.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the PciCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PciCard{}

// PciCard A single object of the PCI card
type PciCard struct {
	// The description of the PCI card of the server
	Description *string `json:"description,omitempty"`
}

// NewPciCard instantiates a new PciCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciCard() *PciCard {
	this := PciCard{}
	return &this
}

// NewPciCardWithDefaults instantiates a new PciCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciCardWithDefaults() *PciCard {
	this := PciCard{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PciCard) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciCard) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PciCard) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PciCard) SetDescription(v string) {
	o.Description = &v
}

func (o PciCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullablePciCard struct {
	value *PciCard
	isSet bool
}

func (v NullablePciCard) Get() *PciCard {
	return v.value
}

func (v *NullablePciCard) Set(val *PciCard) {
	v.value = val
	v.isSet = true
}

func (v NullablePciCard) IsSet() bool {
	return v.isSet
}

func (v *NullablePciCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciCard(val *PciCard) *NullablePciCard {
	return &NullablePciCard{value: val, isSet: true}
}

func (v NullablePciCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

