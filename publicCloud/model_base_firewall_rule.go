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

// checks if the BaseFirewallRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseFirewallRule{}

// BaseFirewallRule struct for BaseFirewallRule
type BaseFirewallRule struct {
	Id *string `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Cidr *string `json:"cidr,omitempty"`
}

// NewBaseFirewallRule instantiates a new BaseFirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseFirewallRule() *BaseFirewallRule {
	this := BaseFirewallRule{}
	return &this
}

// NewBaseFirewallRuleWithDefaults instantiates a new BaseFirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseFirewallRuleWithDefaults() *BaseFirewallRule {
	this := BaseFirewallRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseFirewallRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseFirewallRule) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseFirewallRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BaseFirewallRule) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *BaseFirewallRule) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BaseFirewallRule) UnsetName() {
	o.Name.Unset()
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *BaseFirewallRule) SetCidr(v string) {
	o.Cidr = &v
}

func (o BaseFirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseFirewallRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return toSerialize, nil
}

type NullableBaseFirewallRule struct {
	value *BaseFirewallRule
	isSet bool
}

func (v NullableBaseFirewallRule) Get() *BaseFirewallRule {
	return v.value
}

func (v *NullableBaseFirewallRule) Set(val *BaseFirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseFirewallRule(val *BaseFirewallRule) *NullableBaseFirewallRule {
	return &NullableBaseFirewallRule{value: val, isSet: true}
}

func (v NullableBaseFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

