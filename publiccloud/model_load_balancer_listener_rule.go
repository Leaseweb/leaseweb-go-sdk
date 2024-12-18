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

// checks if the LoadBalancerListenerRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerListenerRule{}

// LoadBalancerListenerRule struct for LoadBalancerListenerRule
type LoadBalancerListenerRule struct {
	// The listener unique identifier
	Id string `json:"id"`
	// Condition of the rule
	Default bool `json:"default"`
	// The target group unique identifier
	TargetGroupId string `json:"targetGroupId"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerListenerRule LoadBalancerListenerRule

// NewLoadBalancerListenerRule instantiates a new LoadBalancerListenerRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerListenerRule(id string, default_ bool, targetGroupId string) *LoadBalancerListenerRule {
	this := LoadBalancerListenerRule{}
	this.Id = id
	this.Default = default_
	this.TargetGroupId = targetGroupId
	return &this
}

// NewLoadBalancerListenerRuleWithDefaults instantiates a new LoadBalancerListenerRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerListenerRuleWithDefaults() *LoadBalancerListenerRule {
	this := LoadBalancerListenerRule{}
	return &this
}

// GetId returns the Id field value
func (o *LoadBalancerListenerRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListenerRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancerListenerRule) SetId(v string) {
	o.Id = v
}

// GetDefault returns the Default field value
func (o *LoadBalancerListenerRule) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListenerRule) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *LoadBalancerListenerRule) SetDefault(v bool) {
	o.Default = v
}

// GetTargetGroupId returns the TargetGroupId field value
func (o *LoadBalancerListenerRule) GetTargetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListenerRule) GetTargetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetGroupId, true
}

// SetTargetGroupId sets field value
func (o *LoadBalancerListenerRule) SetTargetGroupId(v string) {
	o.TargetGroupId = v
}

func (o LoadBalancerListenerRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerListenerRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["default"] = o.Default
	toSerialize["targetGroupId"] = o.TargetGroupId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerListenerRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"default",
		"targetGroupId",
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

	varLoadBalancerListenerRule := _LoadBalancerListenerRule{}

	err = json.Unmarshal(data, &varLoadBalancerListenerRule)

	if err != nil {
		return err
	}

	*o = LoadBalancerListenerRule(varLoadBalancerListenerRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "default")
		delete(additionalProperties, "targetGroupId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerListenerRule struct {
	value *LoadBalancerListenerRule
	isSet bool
}

func (v NullableLoadBalancerListenerRule) Get() *LoadBalancerListenerRule {
	return v.value
}

func (v *NullableLoadBalancerListenerRule) Set(val *LoadBalancerListenerRule) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerListenerRule) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerListenerRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerListenerRule(val *LoadBalancerListenerRule) *NullableLoadBalancerListenerRule {
	return &NullableLoadBalancerListenerRule{value: val, isSet: true}
}

func (v NullableLoadBalancerListenerRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerListenerRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

