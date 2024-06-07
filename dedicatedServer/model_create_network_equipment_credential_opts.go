/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedServer

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateNetworkEquipmentCredentialOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkEquipmentCredentialOpts{}

// CreateNetworkEquipmentCredentialOpts struct for CreateNetworkEquipmentCredentialOpts
type CreateNetworkEquipmentCredentialOpts struct {
	// The password for the credentials
	Password string `json:"password"`
	// Enum: `OPERATING_SYSTEM`, `CONTROL_PANEL`, `REMOTE_MANAGEMENT`, `RESCUE_MODE`, `SWITCH`, `PDU`, `FIREWALL`, `LOAD_BALANCER`  The type of the credential. 
	Type string `json:"type"`
	// The username for the credentials
	Username string `json:"username"`
}

type _CreateNetworkEquipmentCredentialOpts CreateNetworkEquipmentCredentialOpts

// NewCreateNetworkEquipmentCredentialOpts instantiates a new CreateNetworkEquipmentCredentialOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkEquipmentCredentialOpts(password string, type_ string, username string) *CreateNetworkEquipmentCredentialOpts {
	this := CreateNetworkEquipmentCredentialOpts{}
	this.Password = password
	this.Type = type_
	this.Username = username
	return &this
}

// NewCreateNetworkEquipmentCredentialOptsWithDefaults instantiates a new CreateNetworkEquipmentCredentialOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkEquipmentCredentialOptsWithDefaults() *CreateNetworkEquipmentCredentialOpts {
	this := CreateNetworkEquipmentCredentialOpts{}
	return &this
}

// GetPassword returns the Password field value
func (o *CreateNetworkEquipmentCredentialOpts) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkEquipmentCredentialOpts) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateNetworkEquipmentCredentialOpts) SetPassword(v string) {
	o.Password = v
}

// GetType returns the Type field value
func (o *CreateNetworkEquipmentCredentialOpts) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkEquipmentCredentialOpts) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateNetworkEquipmentCredentialOpts) SetType(v string) {
	o.Type = v
}

// GetUsername returns the Username field value
func (o *CreateNetworkEquipmentCredentialOpts) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkEquipmentCredentialOpts) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateNetworkEquipmentCredentialOpts) SetUsername(v string) {
	o.Username = v
}

func (o CreateNetworkEquipmentCredentialOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkEquipmentCredentialOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["type"] = o.Type
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *CreateNetworkEquipmentCredentialOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"type",
		"username",
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

	varCreateNetworkEquipmentCredentialOpts := _CreateNetworkEquipmentCredentialOpts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkEquipmentCredentialOpts)

	if err != nil {
		return err
	}

	*o = CreateNetworkEquipmentCredentialOpts(varCreateNetworkEquipmentCredentialOpts)

	return err
}

type NullableCreateNetworkEquipmentCredentialOpts struct {
	value *CreateNetworkEquipmentCredentialOpts
	isSet bool
}

func (v NullableCreateNetworkEquipmentCredentialOpts) Get() *CreateNetworkEquipmentCredentialOpts {
	return v.value
}

func (v *NullableCreateNetworkEquipmentCredentialOpts) Set(val *CreateNetworkEquipmentCredentialOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkEquipmentCredentialOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkEquipmentCredentialOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkEquipmentCredentialOpts(val *CreateNetworkEquipmentCredentialOpts) *NullableCreateNetworkEquipmentCredentialOpts {
	return &NullableCreateNetworkEquipmentCredentialOpts{value: val, isSet: true}
}

func (v NullableCreateNetworkEquipmentCredentialOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkEquipmentCredentialOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


