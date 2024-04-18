/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
)

// checks if the PrivateNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateNetwork{}

// PrivateNetwork struct for PrivateNetwork
type PrivateNetwork struct {
	PrivateNetworkId *string `json:"privateNetworkId,omitempty"`
	Status *string `json:"status,omitempty"`
	Subnet *string `json:"subnet,omitempty"`
}

// NewPrivateNetwork instantiates a new PrivateNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetwork() *PrivateNetwork {
	this := PrivateNetwork{}
	return &this
}

// NewPrivateNetworkWithDefaults instantiates a new PrivateNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkWithDefaults() *PrivateNetwork {
	this := PrivateNetwork{}
	return &this
}

// GetPrivateNetworkId returns the PrivateNetworkId field value if set, zero value otherwise.
func (o *PrivateNetwork) GetPrivateNetworkId() string {
	if o == nil || IsNil(o.PrivateNetworkId) {
		var ret string
		return ret
	}
	return *o.PrivateNetworkId
}

// GetPrivateNetworkIdOk returns a tuple with the PrivateNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetPrivateNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNetworkId) {
		return nil, false
	}
	return o.PrivateNetworkId, true
}

// HasPrivateNetworkId returns a boolean if a field has been set.
func (o *PrivateNetwork) HasPrivateNetworkId() bool {
	if o != nil && !IsNil(o.PrivateNetworkId) {
		return true
	}

	return false
}

// SetPrivateNetworkId gets a reference to the given string and assigns it to the PrivateNetworkId field.
func (o *PrivateNetwork) SetPrivateNetworkId(v string) {
	o.PrivateNetworkId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PrivateNetwork) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PrivateNetwork) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PrivateNetwork) SetStatus(v string) {
	o.Status = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *PrivateNetwork) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *PrivateNetwork) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *PrivateNetwork) SetSubnet(v string) {
	o.Subnet = &v
}

func (o PrivateNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrivateNetworkId) {
		toSerialize["privateNetworkId"] = o.PrivateNetworkId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	return toSerialize, nil
}

type NullablePrivateNetwork struct {
	value *PrivateNetwork
	isSet bool
}

func (v NullablePrivateNetwork) Get() *PrivateNetwork {
	return v.value
}

func (v *NullablePrivateNetwork) Set(val *PrivateNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetwork(val *PrivateNetwork) *NullablePrivateNetwork {
	return &NullablePrivateNetwork{value: val, isSet: true}
}

func (v NullablePrivateNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


