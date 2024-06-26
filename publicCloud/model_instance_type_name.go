/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// InstanceTypeName Instance type
type InstanceTypeName string

// List of instanceTypeName
const (
	M3_LARGE InstanceTypeName = "lsw.m3.large"
	M3_XLARGE InstanceTypeName = "lsw.m3.xlarge"
	M3_2XLARGE InstanceTypeName = "lsw.m3.2xlarge"
	M4_LARGE InstanceTypeName = "lsw.m4.large"
	M4_XLARGE InstanceTypeName = "lsw.m4.xlarge"
	M4_2XLARGE InstanceTypeName = "lsw.m4.2xlarge"
	M4_4XLARGE InstanceTypeName = "lsw.m4.4xlarge"
	M5_LARGE InstanceTypeName = "lsw.m5.large"
	M5_XLARGE InstanceTypeName = "lsw.m5.xlarge"
	M5_2XLARGE InstanceTypeName = "lsw.m5.2xlarge"
	M5_4XLARGE InstanceTypeName = "lsw.m5.4xlarge"
	M5A_LARGE InstanceTypeName = "lsw.m5a.large"
	M5A_XLARGE InstanceTypeName = "lsw.m5a.xlarge"
	M5A_2XLARGE InstanceTypeName = "lsw.m5a.2xlarge"
	M5A_4XLARGE InstanceTypeName = "lsw.m5a.4xlarge"
	M5A_8XLARGE InstanceTypeName = "lsw.m5a.8xlarge"
	M5A_12XLARGE InstanceTypeName = "lsw.m5a.12xlarge"
	M6A_LARGE InstanceTypeName = "lsw.m6a.large"
	M6A_XLARGE InstanceTypeName = "lsw.m6a.xlarge"
	M6A_2XLARGE InstanceTypeName = "lsw.m6a.2xlarge"
	M6A_4XLARGE InstanceTypeName = "lsw.m6a.4xlarge"
	M6A_8XLARGE InstanceTypeName = "lsw.m6a.8xlarge"
	M6A_12XLARGE InstanceTypeName = "lsw.m6a.12xlarge"
	M6A_16XLARGE InstanceTypeName = "lsw.m6a.16xlarge"
	M6A_24XLARGE InstanceTypeName = "lsw.m6a.24xlarge"
	C3_LARGE InstanceTypeName = "lsw.c3.large"
	C3_XLARGE InstanceTypeName = "lsw.c3.xlarge"
	C3_2XLARGE InstanceTypeName = "lsw.c3.2xlarge"
	C3_4XLARGE InstanceTypeName = "lsw.c3.4xlarge"
	C4_LARGE InstanceTypeName = "lsw.c4.large"
	C4_XLARGE InstanceTypeName = "lsw.c4.xlarge"
	C4_2XLARGE InstanceTypeName = "lsw.c4.2xlarge"
	C4_4XLARGE InstanceTypeName = "lsw.c4.4xlarge"
	C5_LARGE InstanceTypeName = "lsw.c5.large"
	C5_XLARGE InstanceTypeName = "lsw.c5.xlarge"
	C5_2XLARGE InstanceTypeName = "lsw.c5.2xlarge"
	C5_4XLARGE InstanceTypeName = "lsw.c5.4xlarge"
	C5A_LARGE InstanceTypeName = "lsw.c5a.large"
	C5A_XLARGE InstanceTypeName = "lsw.c5a.xlarge"
	C5A_2XLARGE InstanceTypeName = "lsw.c5a.2xlarge"
	C5A_4XLARGE InstanceTypeName = "lsw.c5a.4xlarge"
	C5A_9XLARGE InstanceTypeName = "lsw.c5a.9xlarge"
	C5A_12XLARGE InstanceTypeName = "lsw.c5a.12xlarge"
	C6A_LARGE InstanceTypeName = "lsw.c6a.large"
	C6A_XLARGE InstanceTypeName = "lsw.c6a.xlarge"
	C6A_2XLARGE InstanceTypeName = "lsw.c6a.2xlarge"
	C6A_4XLARGE InstanceTypeName = "lsw.c6a.4xlarge"
	C6A_8XLARGE InstanceTypeName = "lsw.c6a.8xlarge"
	C6A_12XLARGE InstanceTypeName = "lsw.c6a.12xlarge"
	C6A_16XLARGE InstanceTypeName = "lsw.c6a.16xlarge"
	C6A_24XLARGE InstanceTypeName = "lsw.c6a.24xlarge"
	R3_LARGE InstanceTypeName = "lsw.r3.large"
	R3_XLARGE InstanceTypeName = "lsw.r3.xlarge"
	R3_2XLARGE InstanceTypeName = "lsw.r3.2xlarge"
	R4_LARGE InstanceTypeName = "lsw.r4.large"
	R4_XLARGE InstanceTypeName = "lsw.r4.xlarge"
	R4_2XLARGE InstanceTypeName = "lsw.r4.2xlarge"
	R5_LARGE InstanceTypeName = "lsw.r5.large"
	R5_XLARGE InstanceTypeName = "lsw.r5.xlarge"
	R5_2XLARGE InstanceTypeName = "lsw.r5.2xlarge"
	R5A_LARGE InstanceTypeName = "lsw.r5a.large"
	R5A_XLARGE InstanceTypeName = "lsw.r5a.xlarge"
	R5A_2XLARGE InstanceTypeName = "lsw.r5a.2xlarge"
	R5A_4XLARGE InstanceTypeName = "lsw.r5a.4xlarge"
	R5A_8XLARGE InstanceTypeName = "lsw.r5a.8xlarge"
	R5A_12XLARGE InstanceTypeName = "lsw.r5a.12xlarge"
	R6A_LARGE InstanceTypeName = "lsw.r6a.large"
	R6A_XLARGE InstanceTypeName = "lsw.r6a.xlarge"
	R6A_2XLARGE InstanceTypeName = "lsw.r6a.2xlarge"
	R6A_4XLARGE InstanceTypeName = "lsw.r6a.4xlarge"
	R6A_8XLARGE InstanceTypeName = "lsw.r6a.8xlarge"
	R6A_12XLARGE InstanceTypeName = "lsw.r6a.12xlarge"
	R6A_16XLARGE InstanceTypeName = "lsw.r6a.16xlarge"
	R6A_24XLARGE InstanceTypeName = "lsw.r6a.24xlarge"
)

// All allowed values of InstanceTypeName enum
var AllowedInstanceTypeNameEnumValues = []InstanceTypeName{
	"lsw.m3.large",
	"lsw.m3.xlarge",
	"lsw.m3.2xlarge",
	"lsw.m4.large",
	"lsw.m4.xlarge",
	"lsw.m4.2xlarge",
	"lsw.m4.4xlarge",
	"lsw.m5.large",
	"lsw.m5.xlarge",
	"lsw.m5.2xlarge",
	"lsw.m5.4xlarge",
	"lsw.m5a.large",
	"lsw.m5a.xlarge",
	"lsw.m5a.2xlarge",
	"lsw.m5a.4xlarge",
	"lsw.m5a.8xlarge",
	"lsw.m5a.12xlarge",
	"lsw.m6a.large",
	"lsw.m6a.xlarge",
	"lsw.m6a.2xlarge",
	"lsw.m6a.4xlarge",
	"lsw.m6a.8xlarge",
	"lsw.m6a.12xlarge",
	"lsw.m6a.16xlarge",
	"lsw.m6a.24xlarge",
	"lsw.c3.large",
	"lsw.c3.xlarge",
	"lsw.c3.2xlarge",
	"lsw.c3.4xlarge",
	"lsw.c4.large",
	"lsw.c4.xlarge",
	"lsw.c4.2xlarge",
	"lsw.c4.4xlarge",
	"lsw.c5.large",
	"lsw.c5.xlarge",
	"lsw.c5.2xlarge",
	"lsw.c5.4xlarge",
	"lsw.c5a.large",
	"lsw.c5a.xlarge",
	"lsw.c5a.2xlarge",
	"lsw.c5a.4xlarge",
	"lsw.c5a.9xlarge",
	"lsw.c5a.12xlarge",
	"lsw.c6a.large",
	"lsw.c6a.xlarge",
	"lsw.c6a.2xlarge",
	"lsw.c6a.4xlarge",
	"lsw.c6a.8xlarge",
	"lsw.c6a.12xlarge",
	"lsw.c6a.16xlarge",
	"lsw.c6a.24xlarge",
	"lsw.r3.large",
	"lsw.r3.xlarge",
	"lsw.r3.2xlarge",
	"lsw.r4.large",
	"lsw.r4.xlarge",
	"lsw.r4.2xlarge",
	"lsw.r5.large",
	"lsw.r5.xlarge",
	"lsw.r5.2xlarge",
	"lsw.r5a.large",
	"lsw.r5a.xlarge",
	"lsw.r5a.2xlarge",
	"lsw.r5a.4xlarge",
	"lsw.r5a.8xlarge",
	"lsw.r5a.12xlarge",
	"lsw.r6a.large",
	"lsw.r6a.xlarge",
	"lsw.r6a.2xlarge",
	"lsw.r6a.4xlarge",
	"lsw.r6a.8xlarge",
	"lsw.r6a.12xlarge",
	"lsw.r6a.16xlarge",
	"lsw.r6a.24xlarge",
}

func (v *InstanceTypeName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceTypeName(value)
	for _, existing := range AllowedInstanceTypeNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceTypeName", value)
}

// NewInstanceTypeNameFromValue returns a pointer to a valid InstanceTypeName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceTypeNameFromValue(v string) (*InstanceTypeName, error) {
	ev := InstanceTypeName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceTypeName: valid values are %v", v, AllowedInstanceTypeNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceTypeName) IsValid() bool {
	for _, existing := range AllowedInstanceTypeNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to instanceTypeName value
func (v InstanceTypeName) Ptr() *InstanceTypeName {
	return &v
}

type NullableInstanceTypeName struct {
	value *InstanceTypeName
	isSet bool
}

func (v NullableInstanceTypeName) Get() *InstanceTypeName {
	return v.value
}

func (v *NullableInstanceTypeName) Set(val *InstanceTypeName) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeName) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeName(val *InstanceTypeName) *NullableInstanceTypeName {
	return &NullableInstanceTypeName{value: val, isSet: true}
}

func (v NullableInstanceTypeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

