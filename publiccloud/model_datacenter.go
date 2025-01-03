/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// Datacenter the model 'Datacenter'
type Datacenter string

// List of datacenter
const (
	DATACENTER_LON_01 Datacenter = "LON-01"
	DATACENTER_MTL_02 Datacenter = "MTL-02"
	DATACENTER_AMS_01 Datacenter = "AMS-01"
	DATACENTER_FRA_01 Datacenter = "FRA-01"
	DATACENTER_WDC_02 Datacenter = "WDC-02"
	DATACENTER_SFO_12 Datacenter = "SFO-12"
	DATACENTER_SIN_01 Datacenter = "SIN-01"
)

// All allowed values of Datacenter enum
var AllowedDatacenterEnumValues = []Datacenter{
	"LON-01",
	"MTL-02",
	"AMS-01",
	"FRA-01",
	"WDC-02",
	"SFO-12",
	"SIN-01",
}

func (v *Datacenter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Datacenter(value)
	for _, existing := range AllowedDatacenterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Datacenter", value)
}

// NewDatacenterFromValue returns a pointer to a valid Datacenter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatacenterFromValue(v string) (*Datacenter, error) {
	ev := Datacenter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Datacenter: valid values are %v", v, AllowedDatacenterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Datacenter) IsValid() bool {
	for _, existing := range AllowedDatacenterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to datacenter value
func (v Datacenter) Ptr() *Datacenter {
	return &v
}

type NullableDatacenter struct {
	value *Datacenter
	isSet bool
}

func (v NullableDatacenter) Get() *Datacenter {
	return v.value
}

func (v *NullableDatacenter) Set(val *Datacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenter(val *Datacenter) *NullableDatacenter {
	return &NullableDatacenter{value: val, isSet: true}
}

func (v NullableDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

