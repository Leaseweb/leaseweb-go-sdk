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

// checks if the ListenerDefaultRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListenerDefaultRule{}

// ListenerDefaultRule struct for ListenerDefaultRule
type ListenerDefaultRule struct {
	// The target group unique identifier
	TargetGroupId *string `json:"targetGroupId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListenerDefaultRule ListenerDefaultRule

// NewListenerDefaultRule instantiates a new ListenerDefaultRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListenerDefaultRule() *ListenerDefaultRule {
	this := ListenerDefaultRule{}
	return &this
}

// NewListenerDefaultRuleWithDefaults instantiates a new ListenerDefaultRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenerDefaultRuleWithDefaults() *ListenerDefaultRule {
	this := ListenerDefaultRule{}
	return &this
}

// GetTargetGroupId returns the TargetGroupId field value if set, zero value otherwise.
func (o *ListenerDefaultRule) GetTargetGroupId() string {
	if o == nil || IsNil(o.TargetGroupId) {
		var ret string
		return ret
	}
	return *o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerDefaultRule) GetTargetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetGroupId) {
		return nil, false
	}
	return o.TargetGroupId, true
}

// HasTargetGroupId returns a boolean if a field has been set.
func (o *ListenerDefaultRule) HasTargetGroupId() bool {
	if o != nil && !IsNil(o.TargetGroupId) {
		return true
	}

	return false
}

// SetTargetGroupId gets a reference to the given string and assigns it to the TargetGroupId field.
func (o *ListenerDefaultRule) SetTargetGroupId(v string) {
	o.TargetGroupId = &v
}

func (o ListenerDefaultRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListenerDefaultRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetGroupId) {
		toSerialize["targetGroupId"] = o.TargetGroupId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListenerDefaultRule) UnmarshalJSON(data []byte) (err error) {
	varListenerDefaultRule := _ListenerDefaultRule{}

	err = json.Unmarshal(data, &varListenerDefaultRule)

	if err != nil {
		return err
	}

	*o = ListenerDefaultRule(varListenerDefaultRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "targetGroupId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListenerDefaultRule struct {
	value *ListenerDefaultRule
	isSet bool
}

func (v NullableListenerDefaultRule) Get() *ListenerDefaultRule {
	return v.value
}

func (v *NullableListenerDefaultRule) Set(val *ListenerDefaultRule) {
	v.value = val
	v.isSet = true
}

func (v NullableListenerDefaultRule) IsSet() bool {
	return v.isSet
}

func (v *NullableListenerDefaultRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListenerDefaultRule(val *ListenerDefaultRule) *NullableListenerDefaultRule {
	return &NullableListenerDefaultRule{value: val, isSet: true}
}

func (v NullableListenerDefaultRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListenerDefaultRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

