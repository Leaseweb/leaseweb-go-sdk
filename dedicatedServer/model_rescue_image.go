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

// checks if the RescueImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RescueImage{}

// RescueImage A single rescue image
type RescueImage struct {
	// The unique ID of this rescue image
	Id *string `json:"id,omitempty"`
	// A human readable name describing the rescue image
	Name *string `json:"name,omitempty"`
}

// NewRescueImage instantiates a new RescueImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRescueImage() *RescueImage {
	this := RescueImage{}
	return &this
}

// NewRescueImageWithDefaults instantiates a new RescueImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRescueImageWithDefaults() *RescueImage {
	this := RescueImage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RescueImage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RescueImage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RescueImage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RescueImage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RescueImage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RescueImage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RescueImage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RescueImage) SetName(v string) {
	o.Name = &v
}

func (o RescueImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RescueImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableRescueImage struct {
	value *RescueImage
	isSet bool
}

func (v NullableRescueImage) Get() *RescueImage {
	return v.value
}

func (v *NullableRescueImage) Set(val *RescueImage) {
	v.value = val
	v.isSet = true
}

func (v NullableRescueImage) IsSet() bool {
	return v.isSet
}

func (v *NullableRescueImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRescueImage(val *RescueImage) *NullableRescueImage {
	return &NullableRescueImage{value: val, isSet: true}
}

func (v NullableRescueImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRescueImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


