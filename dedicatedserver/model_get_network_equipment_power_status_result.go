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

// checks if the GetNetworkEquipmentPowerStatusResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkEquipmentPowerStatusResult{}

// GetNetworkEquipmentPowerStatusResult struct for GetNetworkEquipmentPowerStatusResult
type GetNetworkEquipmentPowerStatusResult struct {
	Pdu *Pdu `json:"pdu,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNetworkEquipmentPowerStatusResult GetNetworkEquipmentPowerStatusResult

// NewGetNetworkEquipmentPowerStatusResult instantiates a new GetNetworkEquipmentPowerStatusResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkEquipmentPowerStatusResult() *GetNetworkEquipmentPowerStatusResult {
	this := GetNetworkEquipmentPowerStatusResult{}
	return &this
}

// NewGetNetworkEquipmentPowerStatusResultWithDefaults instantiates a new GetNetworkEquipmentPowerStatusResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkEquipmentPowerStatusResultWithDefaults() *GetNetworkEquipmentPowerStatusResult {
	this := GetNetworkEquipmentPowerStatusResult{}
	return &this
}

// GetPdu returns the Pdu field value if set, zero value otherwise.
func (o *GetNetworkEquipmentPowerStatusResult) GetPdu() Pdu {
	if o == nil || IsNil(o.Pdu) {
		var ret Pdu
		return ret
	}
	return *o.Pdu
}

// GetPduOk returns a tuple with the Pdu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEquipmentPowerStatusResult) GetPduOk() (*Pdu, bool) {
	if o == nil || IsNil(o.Pdu) {
		return nil, false
	}
	return o.Pdu, true
}

// HasPdu returns a boolean if a field has been set.
func (o *GetNetworkEquipmentPowerStatusResult) HasPdu() bool {
	if o != nil && !IsNil(o.Pdu) {
		return true
	}

	return false
}

// SetPdu gets a reference to the given Pdu and assigns it to the Pdu field.
func (o *GetNetworkEquipmentPowerStatusResult) SetPdu(v Pdu) {
	o.Pdu = &v
}

func (o GetNetworkEquipmentPowerStatusResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkEquipmentPowerStatusResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pdu) {
		toSerialize["pdu"] = o.Pdu
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNetworkEquipmentPowerStatusResult) UnmarshalJSON(data []byte) (err error) {
	varGetNetworkEquipmentPowerStatusResult := _GetNetworkEquipmentPowerStatusResult{}

	err = json.Unmarshal(data, &varGetNetworkEquipmentPowerStatusResult)

	if err != nil {
		return err
	}

	*o = GetNetworkEquipmentPowerStatusResult(varGetNetworkEquipmentPowerStatusResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pdu")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNetworkEquipmentPowerStatusResult struct {
	value *GetNetworkEquipmentPowerStatusResult
	isSet bool
}

func (v NullableGetNetworkEquipmentPowerStatusResult) Get() *GetNetworkEquipmentPowerStatusResult {
	return v.value
}

func (v *NullableGetNetworkEquipmentPowerStatusResult) Set(val *GetNetworkEquipmentPowerStatusResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkEquipmentPowerStatusResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkEquipmentPowerStatusResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkEquipmentPowerStatusResult(val *GetNetworkEquipmentPowerStatusResult) *NullableGetNetworkEquipmentPowerStatusResult {
	return &NullableGetNetworkEquipmentPowerStatusResult{value: val, isSet: true}
}

func (v NullableGetNetworkEquipmentPowerStatusResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkEquipmentPowerStatusResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


