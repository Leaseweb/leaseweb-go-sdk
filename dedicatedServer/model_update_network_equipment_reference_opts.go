/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedServer

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateNetworkEquipmentReferenceOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkEquipmentReferenceOpts{}

// UpdateNetworkEquipmentReferenceOpts struct for UpdateNetworkEquipmentReferenceOpts
type UpdateNetworkEquipmentReferenceOpts struct {
	// The reference for this network equipment
	Reference string `json:"reference"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNetworkEquipmentReferenceOpts UpdateNetworkEquipmentReferenceOpts

// NewUpdateNetworkEquipmentReferenceOpts instantiates a new UpdateNetworkEquipmentReferenceOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkEquipmentReferenceOpts(reference string) *UpdateNetworkEquipmentReferenceOpts {
	this := UpdateNetworkEquipmentReferenceOpts{}
	this.Reference = reference
	return &this
}

// NewUpdateNetworkEquipmentReferenceOptsWithDefaults instantiates a new UpdateNetworkEquipmentReferenceOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkEquipmentReferenceOptsWithDefaults() *UpdateNetworkEquipmentReferenceOpts {
	this := UpdateNetworkEquipmentReferenceOpts{}
	return &this
}

// GetReference returns the Reference field value
func (o *UpdateNetworkEquipmentReferenceOpts) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkEquipmentReferenceOpts) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *UpdateNetworkEquipmentReferenceOpts) SetReference(v string) {
	o.Reference = v
}

func (o UpdateNetworkEquipmentReferenceOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkEquipmentReferenceOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reference"] = o.Reference

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNetworkEquipmentReferenceOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reference",
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

	varUpdateNetworkEquipmentReferenceOpts := _UpdateNetworkEquipmentReferenceOpts{}

	err = json.Unmarshal(data, &varUpdateNetworkEquipmentReferenceOpts)

	if err != nil {
		return err
	}

	*o = UpdateNetworkEquipmentReferenceOpts(varUpdateNetworkEquipmentReferenceOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNetworkEquipmentReferenceOpts struct {
	value *UpdateNetworkEquipmentReferenceOpts
	isSet bool
}

func (v NullableUpdateNetworkEquipmentReferenceOpts) Get() *UpdateNetworkEquipmentReferenceOpts {
	return v.value
}

func (v *NullableUpdateNetworkEquipmentReferenceOpts) Set(val *UpdateNetworkEquipmentReferenceOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkEquipmentReferenceOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkEquipmentReferenceOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkEquipmentReferenceOpts(val *UpdateNetworkEquipmentReferenceOpts) *NullableUpdateNetworkEquipmentReferenceOpts {
	return &NullableUpdateNetworkEquipmentReferenceOpts{value: val, isSet: true}
}

func (v NullableUpdateNetworkEquipmentReferenceOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkEquipmentReferenceOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


