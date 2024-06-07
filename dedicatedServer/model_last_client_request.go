/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedServer

import (
	"encoding/json"
)

// checks if the LastClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastClientRequest{}

// LastClientRequest If the server acquired this DHCP reservation, this object shows information about the client
type LastClientRequest struct {
	// The relay agent that forwarded the DHCP traffic
	RelayAgent NullableString `json:"relayAgent,omitempty"`
	// The type of DHCP packet requested by the client
	Type *string `json:"type,omitempty"`
	// The user agent of the client making the DHCP request
	UserAgent *string `json:"userAgent,omitempty"`
}

// NewLastClientRequest instantiates a new LastClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastClientRequest() *LastClientRequest {
	this := LastClientRequest{}
	return &this
}

// NewLastClientRequestWithDefaults instantiates a new LastClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastClientRequestWithDefaults() *LastClientRequest {
	this := LastClientRequest{}
	return &this
}

// GetRelayAgent returns the RelayAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LastClientRequest) GetRelayAgent() string {
	if o == nil || IsNil(o.RelayAgent.Get()) {
		var ret string
		return ret
	}
	return *o.RelayAgent.Get()
}

// GetRelayAgentOk returns a tuple with the RelayAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LastClientRequest) GetRelayAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelayAgent.Get(), o.RelayAgent.IsSet()
}

// HasRelayAgent returns a boolean if a field has been set.
func (o *LastClientRequest) HasRelayAgent() bool {
	if o != nil && o.RelayAgent.IsSet() {
		return true
	}

	return false
}

// SetRelayAgent gets a reference to the given NullableString and assigns it to the RelayAgent field.
func (o *LastClientRequest) SetRelayAgent(v string) {
	o.RelayAgent.Set(&v)
}
// SetRelayAgentNil sets the value for RelayAgent to be an explicit nil
func (o *LastClientRequest) SetRelayAgentNil() {
	o.RelayAgent.Set(nil)
}

// UnsetRelayAgent ensures that no value is present for RelayAgent, not even an explicit nil
func (o *LastClientRequest) UnsetRelayAgent() {
	o.RelayAgent.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LastClientRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastClientRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LastClientRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LastClientRequest) SetType(v string) {
	o.Type = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *LastClientRequest) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastClientRequest) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *LastClientRequest) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *LastClientRequest) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o LastClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RelayAgent.IsSet() {
		toSerialize["relayAgent"] = o.RelayAgent.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	return toSerialize, nil
}

type NullableLastClientRequest struct {
	value *LastClientRequest
	isSet bool
}

func (v NullableLastClientRequest) Get() *LastClientRequest {
	return v.value
}

func (v *NullableLastClientRequest) Set(val *LastClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLastClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLastClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastClientRequest(val *LastClientRequest) *NullableLastClientRequest {
	return &NullableLastClientRequest{value: val, isSet: true}
}

func (v NullableLastClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


