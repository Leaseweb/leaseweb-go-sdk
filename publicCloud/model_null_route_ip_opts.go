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

// checks if the NullRouteIpOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NullRouteIpOpts{}

// NullRouteIpOpts struct for NullRouteIpOpts
type NullRouteIpOpts struct {
	// The reason why the IP is being null-routed
	Comment *string `json:"comment,omitempty"`
	// If provided, reverts the operation automatically in the specified value, in hours
	AutomatedUnnulingAt *int32 `json:"automatedUnnulingAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NullRouteIpOpts NullRouteIpOpts

// NewNullRouteIpOpts instantiates a new NullRouteIpOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullRouteIpOpts() *NullRouteIpOpts {
	this := NullRouteIpOpts{}
	return &this
}

// NewNullRouteIpOptsWithDefaults instantiates a new NullRouteIpOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullRouteIpOptsWithDefaults() *NullRouteIpOpts {
	this := NullRouteIpOpts{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NullRouteIpOpts) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRouteIpOpts) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NullRouteIpOpts) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NullRouteIpOpts) SetComment(v string) {
	o.Comment = &v
}

// GetAutomatedUnnulingAt returns the AutomatedUnnulingAt field value if set, zero value otherwise.
func (o *NullRouteIpOpts) GetAutomatedUnnulingAt() int32 {
	if o == nil || IsNil(o.AutomatedUnnulingAt) {
		var ret int32
		return ret
	}
	return *o.AutomatedUnnulingAt
}

// GetAutomatedUnnulingAtOk returns a tuple with the AutomatedUnnulingAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullRouteIpOpts) GetAutomatedUnnulingAtOk() (*int32, bool) {
	if o == nil || IsNil(o.AutomatedUnnulingAt) {
		return nil, false
	}
	return o.AutomatedUnnulingAt, true
}

// HasAutomatedUnnulingAt returns a boolean if a field has been set.
func (o *NullRouteIpOpts) HasAutomatedUnnulingAt() bool {
	if o != nil && !IsNil(o.AutomatedUnnulingAt) {
		return true
	}

	return false
}

// SetAutomatedUnnulingAt gets a reference to the given int32 and assigns it to the AutomatedUnnulingAt field.
func (o *NullRouteIpOpts) SetAutomatedUnnulingAt(v int32) {
	o.AutomatedUnnulingAt = &v
}

func (o NullRouteIpOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NullRouteIpOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.AutomatedUnnulingAt) {
		toSerialize["automatedUnnulingAt"] = o.AutomatedUnnulingAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NullRouteIpOpts) UnmarshalJSON(data []byte) (err error) {
	varNullRouteIpOpts := _NullRouteIpOpts{}

	err = json.Unmarshal(data, &varNullRouteIpOpts)

	if err != nil {
		return err
	}

	*o = NullRouteIpOpts(varNullRouteIpOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "automatedUnnulingAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNullRouteIpOpts struct {
	value *NullRouteIpOpts
	isSet bool
}

func (v NullableNullRouteIpOpts) Get() *NullRouteIpOpts {
	return v.value
}

func (v *NullableNullRouteIpOpts) Set(val *NullRouteIpOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableNullRouteIpOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableNullRouteIpOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullRouteIpOpts(val *NullRouteIpOpts) *NullableNullRouteIpOpts {
	return &NullableNullRouteIpOpts{value: val, isSet: true}
}

func (v NullableNullRouteIpOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullRouteIpOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


