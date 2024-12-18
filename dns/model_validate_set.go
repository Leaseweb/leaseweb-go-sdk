/*
DNS

>The base URL for this API is: **https://api.leaseweb.com/hosting/v2**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ValidateSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateSet{}

// ValidateSet Link to validate a single resource record set
type ValidateSet struct {
	// Hyperlink to validate a single resource record set for a domain
	Href *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidateSet ValidateSet

// NewValidateSet instantiates a new ValidateSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateSet() *ValidateSet {
	this := ValidateSet{}
	return &this
}

// NewValidateSetWithDefaults instantiates a new ValidateSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateSetWithDefaults() *ValidateSet {
	this := ValidateSet{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ValidateSet) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSet) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ValidateSet) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ValidateSet) SetHref(v string) {
	o.Href = &v
}

func (o ValidateSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateSet) UnmarshalJSON(data []byte) (err error) {
	varValidateSet := _ValidateSet{}

	err = json.Unmarshal(data, &varValidateSet)

	if err != nil {
		return err
	}

	*o = ValidateSet(varValidateSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateSet struct {
	value *ValidateSet
	isSet bool
}

func (v NullableValidateSet) Get() *ValidateSet {
	return v.value
}

func (v *NullableValidateSet) Set(val *ValidateSet) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateSet) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateSet(val *ValidateSet) *NullableValidateSet {
	return &NullableValidateSet{value: val, isSet: true}
}

func (v NullableValidateSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


