/*
DNS

>The base URL for this API is: **https://api.leaseweb.com/hosting/v2**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// checks if the ImportResourceRecordSetsOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportResourceRecordSetsOpts{}

// ImportResourceRecordSetsOpts struct for ImportResourceRecordSetsOpts
type ImportResourceRecordSetsOpts struct {
	// Dns bind file content as string
	Content string `json:"content"`
	AdditionalProperties map[string]interface{}
}

type _ImportResourceRecordSetsOpts ImportResourceRecordSetsOpts

// NewImportResourceRecordSetsOpts instantiates a new ImportResourceRecordSetsOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportResourceRecordSetsOpts(content string) *ImportResourceRecordSetsOpts {
	this := ImportResourceRecordSetsOpts{}
	this.Content = content
	return &this
}

// NewImportResourceRecordSetsOptsWithDefaults instantiates a new ImportResourceRecordSetsOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportResourceRecordSetsOptsWithDefaults() *ImportResourceRecordSetsOpts {
	this := ImportResourceRecordSetsOpts{}
	return &this
}

// GetContent returns the Content field value
func (o *ImportResourceRecordSetsOpts) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ImportResourceRecordSetsOpts) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ImportResourceRecordSetsOpts) SetContent(v string) {
	o.Content = v
}

func (o ImportResourceRecordSetsOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportResourceRecordSetsOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportResourceRecordSetsOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
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

	varImportResourceRecordSetsOpts := _ImportResourceRecordSetsOpts{}

	err = json.Unmarshal(data, &varImportResourceRecordSetsOpts)

	if err != nil {
		return err
	}

	*o = ImportResourceRecordSetsOpts(varImportResourceRecordSetsOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportResourceRecordSetsOpts struct {
	value *ImportResourceRecordSetsOpts
	isSet bool
}

func (v NullableImportResourceRecordSetsOpts) Get() *ImportResourceRecordSetsOpts {
	return v.value
}

func (v *NullableImportResourceRecordSetsOpts) Set(val *ImportResourceRecordSetsOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableImportResourceRecordSetsOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableImportResourceRecordSetsOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportResourceRecordSetsOpts(val *ImportResourceRecordSetsOpts) *NullableImportResourceRecordSetsOpts {
	return &NullableImportResourceRecordSetsOpts{value: val, isSet: true}
}

func (v NullableImportResourceRecordSetsOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportResourceRecordSetsOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


