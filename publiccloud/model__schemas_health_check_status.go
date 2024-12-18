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

// checks if the SchemasHealthCheckStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasHealthCheckStatus{}

// SchemasHealthCheckStatus struct for SchemasHealthCheckStatus
type SchemasHealthCheckStatus struct {
	State HealthCheckStatus `json:"state"`
	// Description of the health check status
	Description string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _SchemasHealthCheckStatus SchemasHealthCheckStatus

// NewSchemasHealthCheckStatus instantiates a new SchemasHealthCheckStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasHealthCheckStatus(state HealthCheckStatus, description string) *SchemasHealthCheckStatus {
	this := SchemasHealthCheckStatus{}
	this.State = state
	this.Description = description
	return &this
}

// NewSchemasHealthCheckStatusWithDefaults instantiates a new SchemasHealthCheckStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasHealthCheckStatusWithDefaults() *SchemasHealthCheckStatus {
	this := SchemasHealthCheckStatus{}
	return &this
}

// GetState returns the State field value
func (o *SchemasHealthCheckStatus) GetState() HealthCheckStatus {
	if o == nil {
		var ret HealthCheckStatus
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SchemasHealthCheckStatus) GetStateOk() (*HealthCheckStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SchemasHealthCheckStatus) SetState(v HealthCheckStatus) {
	o.State = v
}

// GetDescription returns the Description field value
func (o *SchemasHealthCheckStatus) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemasHealthCheckStatus) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemasHealthCheckStatus) SetDescription(v string) {
	o.Description = v
}

func (o SchemasHealthCheckStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasHealthCheckStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchemasHealthCheckStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
		"description",
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

	varSchemasHealthCheckStatus := _SchemasHealthCheckStatus{}

	err = json.Unmarshal(data, &varSchemasHealthCheckStatus)

	if err != nil {
		return err
	}

	*o = SchemasHealthCheckStatus(varSchemasHealthCheckStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchemasHealthCheckStatus struct {
	value *SchemasHealthCheckStatus
	isSet bool
}

func (v NullableSchemasHealthCheckStatus) Get() *SchemasHealthCheckStatus {
	return v.value
}

func (v *NullableSchemasHealthCheckStatus) Set(val *SchemasHealthCheckStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasHealthCheckStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasHealthCheckStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasHealthCheckStatus(val *SchemasHealthCheckStatus) *NullableSchemasHealthCheckStatus {
	return &NullableSchemasHealthCheckStatus{value: val, isSet: true}
}

func (v NullableSchemasHealthCheckStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasHealthCheckStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

