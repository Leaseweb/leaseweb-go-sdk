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

// checks if the CreateTargetGroupOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTargetGroupOpts{}

// CreateTargetGroupOpts struct for CreateTargetGroupOpts
type CreateTargetGroupOpts struct {
	// The name of the target group
	Name string `json:"name"`
	Protocol Protocol `json:"protocol"`
	// The port of the target group
	Port int32 `json:"port"`
	Region RegionName `json:"region"`
	HealthCheck *HealthCheckOpts `json:"healthCheck,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateTargetGroupOpts CreateTargetGroupOpts

// NewCreateTargetGroupOpts instantiates a new CreateTargetGroupOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTargetGroupOpts(name string, protocol Protocol, port int32, region RegionName) *CreateTargetGroupOpts {
	this := CreateTargetGroupOpts{}
	this.Name = name
	this.Protocol = protocol
	this.Port = port
	this.Region = region
	return &this
}

// NewCreateTargetGroupOptsWithDefaults instantiates a new CreateTargetGroupOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTargetGroupOptsWithDefaults() *CreateTargetGroupOpts {
	this := CreateTargetGroupOpts{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTargetGroupOpts) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTargetGroupOpts) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTargetGroupOpts) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value
func (o *CreateTargetGroupOpts) GetProtocol() Protocol {
	if o == nil {
		var ret Protocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *CreateTargetGroupOpts) GetProtocolOk() (*Protocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *CreateTargetGroupOpts) SetProtocol(v Protocol) {
	o.Protocol = v
}

// GetPort returns the Port field value
func (o *CreateTargetGroupOpts) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateTargetGroupOpts) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateTargetGroupOpts) SetPort(v int32) {
	o.Port = v
}

// GetRegion returns the Region field value
func (o *CreateTargetGroupOpts) GetRegion() RegionName {
	if o == nil {
		var ret RegionName
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateTargetGroupOpts) GetRegionOk() (*RegionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateTargetGroupOpts) SetRegion(v RegionName) {
	o.Region = v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *CreateTargetGroupOpts) GetHealthCheck() HealthCheckOpts {
	if o == nil || IsNil(o.HealthCheck) {
		var ret HealthCheckOpts
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTargetGroupOpts) GetHealthCheckOk() (*HealthCheckOpts, bool) {
	if o == nil || IsNil(o.HealthCheck) {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *CreateTargetGroupOpts) HasHealthCheck() bool {
	if o != nil && !IsNil(o.HealthCheck) {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given HealthCheckOpts and assigns it to the HealthCheck field.
func (o *CreateTargetGroupOpts) SetHealthCheck(v HealthCheckOpts) {
	o.HealthCheck = &v
}

func (o CreateTargetGroupOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTargetGroupOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["protocol"] = o.Protocol
	toSerialize["port"] = o.Port
	toSerialize["region"] = o.Region
	if !IsNil(o.HealthCheck) {
		toSerialize["healthCheck"] = o.HealthCheck
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateTargetGroupOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"protocol",
		"port",
		"region",
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

	varCreateTargetGroupOpts := _CreateTargetGroupOpts{}

	err = json.Unmarshal(data, &varCreateTargetGroupOpts)

	if err != nil {
		return err
	}

	*o = CreateTargetGroupOpts(varCreateTargetGroupOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "port")
		delete(additionalProperties, "region")
		delete(additionalProperties, "healthCheck")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateTargetGroupOpts struct {
	value *CreateTargetGroupOpts
	isSet bool
}

func (v NullableCreateTargetGroupOpts) Get() *CreateTargetGroupOpts {
	return v.value
}

func (v *NullableCreateTargetGroupOpts) Set(val *CreateTargetGroupOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTargetGroupOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTargetGroupOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTargetGroupOpts(val *CreateTargetGroupOpts) *NullableCreateTargetGroupOpts {
	return &NullableCreateTargetGroupOpts{value: val, isSet: true}
}

func (v NullableCreateTargetGroupOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTargetGroupOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

