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

// LoadBalancerGranularity Defines the time interval for data aggregation
type LoadBalancerGranularity string

// List of loadBalancerGranularity
const (
	LOADBALANCERGRANULARITY__5M LoadBalancerGranularity = "5m"
	LOADBALANCERGRANULARITY__10M LoadBalancerGranularity = "10m"
	LOADBALANCERGRANULARITY__30M LoadBalancerGranularity = "30m"
	LOADBALANCERGRANULARITY__1H LoadBalancerGranularity = "1h"
	LOADBALANCERGRANULARITY__1D LoadBalancerGranularity = "1d"
	LOADBALANCERGRANULARITY__1W LoadBalancerGranularity = "1w"
)

// All allowed values of LoadBalancerGranularity enum
var AllowedLoadBalancerGranularityEnumValues = []LoadBalancerGranularity{
	"5m",
	"10m",
	"30m",
	"1h",
	"1d",
	"1w",
}

func (v *LoadBalancerGranularity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoadBalancerGranularity(value)
	for _, existing := range AllowedLoadBalancerGranularityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoadBalancerGranularity", value)
}

// NewLoadBalancerGranularityFromValue returns a pointer to a valid LoadBalancerGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoadBalancerGranularityFromValue(v string) (*LoadBalancerGranularity, error) {
	ev := LoadBalancerGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoadBalancerGranularity: valid values are %v", v, AllowedLoadBalancerGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoadBalancerGranularity) IsValid() bool {
	for _, existing := range AllowedLoadBalancerGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to loadBalancerGranularity value
func (v LoadBalancerGranularity) Ptr() *LoadBalancerGranularity {
	return &v
}

type NullableLoadBalancerGranularity struct {
	value *LoadBalancerGranularity
	isSet bool
}

func (v NullableLoadBalancerGranularity) Get() *LoadBalancerGranularity {
	return v.value
}

func (v *NullableLoadBalancerGranularity) Set(val *LoadBalancerGranularity) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerGranularity) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerGranularity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerGranularity(val *LoadBalancerGranularity) *NullableLoadBalancerGranularity {
	return &NullableLoadBalancerGranularity{value: val, isSet: true}
}

func (v NullableLoadBalancerGranularity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerGranularity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
