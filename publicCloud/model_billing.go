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

// checks if the Billing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Billing{}

// Billing struct for Billing
type Billing struct {
	// List of instances to be billed in the period
	Instances []ExpenseResultInstance `json:"instances,omitempty"`
	Traffic *Traffic `json:"traffic,omitempty"`
}

// NewBilling instantiates a new Billing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBilling() *Billing {
	this := Billing{}
	return &this
}

// NewBillingWithDefaults instantiates a new Billing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingWithDefaults() *Billing {
	this := Billing{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *Billing) GetInstances() []ExpenseResultInstance {
	if o == nil || IsNil(o.Instances) {
		var ret []ExpenseResultInstance
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetInstancesOk() ([]ExpenseResultInstance, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *Billing) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []ExpenseResultInstance and assigns it to the Instances field.
func (o *Billing) SetInstances(v []ExpenseResultInstance) {
	o.Instances = v
}

// GetTraffic returns the Traffic field value if set, zero value otherwise.
func (o *Billing) GetTraffic() Traffic {
	if o == nil || IsNil(o.Traffic) {
		var ret Traffic
		return ret
	}
	return *o.Traffic
}

// GetTrafficOk returns a tuple with the Traffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetTrafficOk() (*Traffic, bool) {
	if o == nil || IsNil(o.Traffic) {
		return nil, false
	}
	return o.Traffic, true
}

// HasTraffic returns a boolean if a field has been set.
func (o *Billing) HasTraffic() bool {
	if o != nil && !IsNil(o.Traffic) {
		return true
	}

	return false
}

// SetTraffic gets a reference to the given Traffic and assigns it to the Traffic field.
func (o *Billing) SetTraffic(v Traffic) {
	o.Traffic = &v
}

func (o Billing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Billing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.Traffic) {
		toSerialize["traffic"] = o.Traffic
	}
	return toSerialize, nil
}

type NullableBilling struct {
	value *Billing
	isSet bool
}

func (v NullableBilling) Get() *Billing {
	return v.value
}

func (v *NullableBilling) Set(val *Billing) {
	v.value = val
	v.isSet = true
}

func (v NullableBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBilling(val *Billing) *NullableBilling {
	return &NullableBilling{value: val, isSet: true}
}

func (v NullableBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


