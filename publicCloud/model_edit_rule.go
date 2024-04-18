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

// checks if the EditRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditRule{}

// EditRule struct for EditRule
type EditRule struct {
	Protocol string `json:"protocol"`
	StartPort NullableInt32 `json:"startPort"`
	EndPort NullableInt32 `json:"endPort"`
	IcmpType int32 `json:"icmpType"`
	IcmpCode int32 `json:"icmpCode"`
	Id string `json:"id"`
	Name NullableString `json:"name"`
	Cidr string `json:"cidr"`
}

type _EditRule EditRule

// NewEditRule instantiates a new EditRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditRule(protocol string, startPort NullableInt32, endPort NullableInt32, icmpType int32, icmpCode int32, id string, name NullableString, cidr string) *EditRule {
	this := EditRule{}
	this.Protocol = protocol
	this.StartPort = startPort
	this.EndPort = endPort
	this.IcmpType = icmpType
	this.IcmpCode = icmpCode
	this.Id = id
	this.Name = name
	this.Cidr = cidr
	return &this
}

// NewEditRuleWithDefaults instantiates a new EditRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditRuleWithDefaults() *EditRule {
	this := EditRule{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *EditRule) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *EditRule) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *EditRule) SetProtocol(v string) {
	o.Protocol = v
}

// GetStartPort returns the StartPort field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EditRule) GetStartPort() int32 {
	if o == nil || o.StartPort.Get() == nil {
		var ret int32
		return ret
	}

	return *o.StartPort.Get()
}

// GetStartPortOk returns a tuple with the StartPort field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditRule) GetStartPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartPort.Get(), o.StartPort.IsSet()
}

// SetStartPort sets field value
func (o *EditRule) SetStartPort(v int32) {
	o.StartPort.Set(&v)
}

// GetEndPort returns the EndPort field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EditRule) GetEndPort() int32 {
	if o == nil || o.EndPort.Get() == nil {
		var ret int32
		return ret
	}

	return *o.EndPort.Get()
}

// GetEndPortOk returns a tuple with the EndPort field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditRule) GetEndPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndPort.Get(), o.EndPort.IsSet()
}

// SetEndPort sets field value
func (o *EditRule) SetEndPort(v int32) {
	o.EndPort.Set(&v)
}

// GetIcmpType returns the IcmpType field value
func (o *EditRule) GetIcmpType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IcmpType
}

// GetIcmpTypeOk returns a tuple with the IcmpType field value
// and a boolean to check if the value has been set.
func (o *EditRule) GetIcmpTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IcmpType, true
}

// SetIcmpType sets field value
func (o *EditRule) SetIcmpType(v int32) {
	o.IcmpType = v
}

// GetIcmpCode returns the IcmpCode field value
func (o *EditRule) GetIcmpCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IcmpCode
}

// GetIcmpCodeOk returns a tuple with the IcmpCode field value
// and a boolean to check if the value has been set.
func (o *EditRule) GetIcmpCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IcmpCode, true
}

// SetIcmpCode sets field value
func (o *EditRule) SetIcmpCode(v int32) {
	o.IcmpCode = v
}

// GetId returns the Id field value
func (o *EditRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EditRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EditRule) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditRule) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *EditRule) SetName(v string) {
	o.Name.Set(&v)
}

// GetCidr returns the Cidr field value
func (o *EditRule) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *EditRule) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *EditRule) SetCidr(v string) {
	o.Cidr = v
}

func (o EditRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["startPort"] = o.StartPort.Get()
	toSerialize["endPort"] = o.EndPort.Get()
	toSerialize["icmpType"] = o.IcmpType
	toSerialize["icmpCode"] = o.IcmpCode
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["cidr"] = o.Cidr
	return toSerialize, nil
}

func (o *EditRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protocol",
		"startPort",
		"endPort",
		"icmpType",
		"icmpCode",
		"id",
		"name",
		"cidr",
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

	varEditRule := _EditRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditRule)

	if err != nil {
		return err
	}

	*o = EditRule(varEditRule)

	return err
}

type NullableEditRule struct {
	value *EditRule
	isSet bool
}

func (v NullableEditRule) Get() *EditRule {
	return v.value
}

func (v *NullableEditRule) Set(val *EditRule) {
	v.value = val
	v.isSet = true
}

func (v NullableEditRule) IsSet() bool {
	return v.isSet
}

func (v *NullableEditRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditRule(val *EditRule) *NullableEditRule {
	return &NullableEditRule{value: val, isSet: true}
}

func (v NullableEditRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


