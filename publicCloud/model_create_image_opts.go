/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateImageOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateImageOpts{}

// CreateImageOpts struct for CreateImageOpts
type CreateImageOpts struct {
	// The name of the custom image to be created.
	Name string `json:"name"`
	// The id of the instance from which the custom image will be created.
	InstanceId string `json:"instanceId"`
	AdditionalProperties map[string]interface{}
}

type _CreateImageOpts CreateImageOpts

// NewCreateImageOpts instantiates a new CreateImageOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageOpts(name string, instanceId string) *CreateImageOpts {
	this := CreateImageOpts{}
	this.Name = name
	this.InstanceId = instanceId
	return &this
}

// NewCreateImageOptsWithDefaults instantiates a new CreateImageOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageOptsWithDefaults() *CreateImageOpts {
	this := CreateImageOpts{}
	return &this
}

// GetName returns the Name field value
func (o *CreateImageOpts) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateImageOpts) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateImageOpts) SetName(v string) {
	o.Name = v
}

// GetInstanceId returns the InstanceId field value
func (o *CreateImageOpts) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *CreateImageOpts) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *CreateImageOpts) SetInstanceId(v string) {
	o.InstanceId = v
}

func (o CreateImageOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateImageOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["instanceId"] = o.InstanceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateImageOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"instanceId",
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

	varCreateImageOpts := _CreateImageOpts{}

	err = json.Unmarshal(data, &varCreateImageOpts)

	if err != nil {
		return err
	}

	*o = CreateImageOpts(varCreateImageOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "instanceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateImageOpts struct {
	value *CreateImageOpts
	isSet bool
}

func (v NullableCreateImageOpts) Get() *CreateImageOpts {
	return v.value
}

func (v *NullableCreateImageOpts) Set(val *CreateImageOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageOpts(val *CreateImageOpts) *NullableCreateImageOpts {
	return &NullableCreateImageOpts{value: val, isSet: true}
}

func (v NullableCreateImageOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


