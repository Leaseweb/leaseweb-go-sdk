/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TcpUdpFirewallRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TcpUdpFirewallRule{}

// TcpUdpFirewallRule struct for TcpUdpFirewallRule
type TcpUdpFirewallRule struct {
	Id *string `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Cidr *string `json:"cidr,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	StartPort int32 `json:"startPort"`
	// Value will be equals to the startPort when not provided
	EndPort int32 `json:"endPort"`
	IcmpType NullableInt32 `json:"icmpType,omitempty"`
	IcmpCode NullableInt32 `json:"icmpCode,omitempty"`
}

type _TcpUdpFirewallRule TcpUdpFirewallRule

// NewTcpUdpFirewallRule instantiates a new TcpUdpFirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpUdpFirewallRule(startPort int32, endPort int32) *TcpUdpFirewallRule {
	this := TcpUdpFirewallRule{}
	this.StartPort = startPort
	this.EndPort = endPort
	return &this
}

// NewTcpUdpFirewallRuleWithDefaults instantiates a new TcpUdpFirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpUdpFirewallRuleWithDefaults() *TcpUdpFirewallRule {
	this := TcpUdpFirewallRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TcpUdpFirewallRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpFirewallRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TcpUdpFirewallRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TcpUdpFirewallRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TcpUdpFirewallRule) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TcpUdpFirewallRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TcpUdpFirewallRule) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TcpUdpFirewallRule) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TcpUdpFirewallRule) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TcpUdpFirewallRule) UnsetName() {
	o.Name.Unset()
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *TcpUdpFirewallRule) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpFirewallRule) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *TcpUdpFirewallRule) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *TcpUdpFirewallRule) SetCidr(v string) {
	o.Cidr = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *TcpUdpFirewallRule) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpUdpFirewallRule) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *TcpUdpFirewallRule) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *TcpUdpFirewallRule) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStartPort returns the StartPort field value
func (o *TcpUdpFirewallRule) GetStartPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartPort
}

// GetStartPortOk returns a tuple with the StartPort field value
// and a boolean to check if the value has been set.
func (o *TcpUdpFirewallRule) GetStartPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartPort, true
}

// SetStartPort sets field value
func (o *TcpUdpFirewallRule) SetStartPort(v int32) {
	o.StartPort = v
}

// GetEndPort returns the EndPort field value
func (o *TcpUdpFirewallRule) GetEndPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndPort
}

// GetEndPortOk returns a tuple with the EndPort field value
// and a boolean to check if the value has been set.
func (o *TcpUdpFirewallRule) GetEndPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPort, true
}

// SetEndPort sets field value
func (o *TcpUdpFirewallRule) SetEndPort(v int32) {
	o.EndPort = v
}

// GetIcmpType returns the IcmpType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TcpUdpFirewallRule) GetIcmpType() int32 {
	if o == nil || IsNil(o.IcmpType.Get()) {
		var ret int32
		return ret
	}
	return *o.IcmpType.Get()
}

// GetIcmpTypeOk returns a tuple with the IcmpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TcpUdpFirewallRule) GetIcmpTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IcmpType.Get(), o.IcmpType.IsSet()
}

// HasIcmpType returns a boolean if a field has been set.
func (o *TcpUdpFirewallRule) HasIcmpType() bool {
	if o != nil && o.IcmpType.IsSet() {
		return true
	}

	return false
}

// SetIcmpType gets a reference to the given NullableInt32 and assigns it to the IcmpType field.
func (o *TcpUdpFirewallRule) SetIcmpType(v int32) {
	o.IcmpType.Set(&v)
}
// SetIcmpTypeNil sets the value for IcmpType to be an explicit nil
func (o *TcpUdpFirewallRule) SetIcmpTypeNil() {
	o.IcmpType.Set(nil)
}

// UnsetIcmpType ensures that no value is present for IcmpType, not even an explicit nil
func (o *TcpUdpFirewallRule) UnsetIcmpType() {
	o.IcmpType.Unset()
}

// GetIcmpCode returns the IcmpCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TcpUdpFirewallRule) GetIcmpCode() int32 {
	if o == nil || IsNil(o.IcmpCode.Get()) {
		var ret int32
		return ret
	}
	return *o.IcmpCode.Get()
}

// GetIcmpCodeOk returns a tuple with the IcmpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TcpUdpFirewallRule) GetIcmpCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IcmpCode.Get(), o.IcmpCode.IsSet()
}

// HasIcmpCode returns a boolean if a field has been set.
func (o *TcpUdpFirewallRule) HasIcmpCode() bool {
	if o != nil && o.IcmpCode.IsSet() {
		return true
	}

	return false
}

// SetIcmpCode gets a reference to the given NullableInt32 and assigns it to the IcmpCode field.
func (o *TcpUdpFirewallRule) SetIcmpCode(v int32) {
	o.IcmpCode.Set(&v)
}
// SetIcmpCodeNil sets the value for IcmpCode to be an explicit nil
func (o *TcpUdpFirewallRule) SetIcmpCodeNil() {
	o.IcmpCode.Set(nil)
}

// UnsetIcmpCode ensures that no value is present for IcmpCode, not even an explicit nil
func (o *TcpUdpFirewallRule) UnsetIcmpCode() {
	o.IcmpCode.Unset()
}

func (o TcpUdpFirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpUdpFirewallRule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	toSerialize["startPort"] = o.StartPort
	toSerialize["endPort"] = o.EndPort
	if o.IcmpType.IsSet() {
		toSerialize["icmpType"] = o.IcmpType.Get()
	}
	if o.IcmpCode.IsSet() {
		toSerialize["icmpCode"] = o.IcmpCode.Get()
	}
	return toSerialize, nil
}

func (o *TcpUdpFirewallRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startPort",
		"endPort",
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

	varTcpUdpFirewallRule := _TcpUdpFirewallRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTcpUdpFirewallRule)

	if err != nil {
		return err
	}

	*o = TcpUdpFirewallRule(varTcpUdpFirewallRule)

	return err
}

type NullableTcpUdpFirewallRule struct {
	value *TcpUdpFirewallRule
	isSet bool
}

func (v NullableTcpUdpFirewallRule) Get() *TcpUdpFirewallRule {
	return v.value
}

func (v *NullableTcpUdpFirewallRule) Set(val *TcpUdpFirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpUdpFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpUdpFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpUdpFirewallRule(val *TcpUdpFirewallRule) *NullableTcpUdpFirewallRule {
	return &NullableTcpUdpFirewallRule{value: val, isSet: true}
}

func (v NullableTcpUdpFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpUdpFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

