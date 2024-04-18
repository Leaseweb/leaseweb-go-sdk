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

// checks if the EditFirewallRulesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditFirewallRulesResult{}

// EditFirewallRulesResult struct for EditFirewallRulesResult
type EditFirewallRulesResult struct {
	// List of rules that have been successfully edited
	Updated []FirewallRule `json:"updated,omitempty"`
	// List of rules that have failed to be edited along with their error messages
	Failed []FailedRule `json:"failed,omitempty"`
}

// NewEditFirewallRulesResult instantiates a new EditFirewallRulesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditFirewallRulesResult() *EditFirewallRulesResult {
	this := EditFirewallRulesResult{}
	return &this
}

// NewEditFirewallRulesResultWithDefaults instantiates a new EditFirewallRulesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditFirewallRulesResultWithDefaults() *EditFirewallRulesResult {
	this := EditFirewallRulesResult{}
	return &this
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *EditFirewallRulesResult) GetUpdated() []FirewallRule {
	if o == nil || IsNil(o.Updated) {
		var ret []FirewallRule
		return ret
	}
	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFirewallRulesResult) GetUpdatedOk() ([]FirewallRule, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *EditFirewallRulesResult) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given []FirewallRule and assigns it to the Updated field.
func (o *EditFirewallRulesResult) SetUpdated(v []FirewallRule) {
	o.Updated = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *EditFirewallRulesResult) GetFailed() []FailedRule {
	if o == nil || IsNil(o.Failed) {
		var ret []FailedRule
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFirewallRulesResult) GetFailedOk() ([]FailedRule, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *EditFirewallRulesResult) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given []FailedRule and assigns it to the Failed field.
func (o *EditFirewallRulesResult) SetFailed(v []FailedRule) {
	o.Failed = v
}

func (o EditFirewallRulesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditFirewallRulesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	return toSerialize, nil
}

type NullableEditFirewallRulesResult struct {
	value *EditFirewallRulesResult
	isSet bool
}

func (v NullableEditFirewallRulesResult) Get() *EditFirewallRulesResult {
	return v.value
}

func (v *NullableEditFirewallRulesResult) Set(val *EditFirewallRulesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEditFirewallRulesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEditFirewallRulesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditFirewallRulesResult(val *EditFirewallRulesResult) *NullableEditFirewallRulesResult {
	return &NullableEditFirewallRulesResult{value: val, isSet: true}
}

func (v NullableEditFirewallRulesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditFirewallRulesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

