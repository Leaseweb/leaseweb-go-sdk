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

// checks if the OperationNetworkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationNetworkInterface{}

// OperationNetworkInterface struct for OperationNetworkInterface
type OperationNetworkInterface struct {
	// The link speed
	LinkSpeed *string `json:"linkSpeed,omitempty"`
	// The operational status
	OperStatus *string `json:"operStatus,omitempty"`
	// The administrative status
	Status *string `json:"status,omitempty"`
	// The switch port number
	SwitchInterface *string `json:"switchInterface,omitempty"`
	// The switch name
	SwitchName *string `json:"switchName,omitempty"`
	// The network type
	Type *string `json:"type,omitempty"`
}

// NewOperationNetworkInterface instantiates a new OperationNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationNetworkInterface() *OperationNetworkInterface {
	this := OperationNetworkInterface{}
	return &this
}

// NewOperationNetworkInterfaceWithDefaults instantiates a new OperationNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationNetworkInterfaceWithDefaults() *OperationNetworkInterface {
	this := OperationNetworkInterface{}
	return &this
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *OperationNetworkInterface) GetLinkSpeed() string {
	if o == nil || IsNil(o.LinkSpeed) {
		var ret string
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationNetworkInterface) GetLinkSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.LinkSpeed) {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *OperationNetworkInterface) HasLinkSpeed() bool {
	if o != nil && !IsNil(o.LinkSpeed) {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given string and assigns it to the LinkSpeed field.
func (o *OperationNetworkInterface) SetLinkSpeed(v string) {
	o.LinkSpeed = &v
}

// GetOperStatus returns the OperStatus field value if set, zero value otherwise.
func (o *OperationNetworkInterface) GetOperStatus() string {
	if o == nil || IsNil(o.OperStatus) {
		var ret string
		return ret
	}
	return *o.OperStatus
}

// GetOperStatusOk returns a tuple with the OperStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationNetworkInterface) GetOperStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OperStatus) {
		return nil, false
	}
	return o.OperStatus, true
}

// HasOperStatus returns a boolean if a field has been set.
func (o *OperationNetworkInterface) HasOperStatus() bool {
	if o != nil && !IsNil(o.OperStatus) {
		return true
	}

	return false
}

// SetOperStatus gets a reference to the given string and assigns it to the OperStatus field.
func (o *OperationNetworkInterface) SetOperStatus(v string) {
	o.OperStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OperationNetworkInterface) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationNetworkInterface) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OperationNetworkInterface) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OperationNetworkInterface) SetStatus(v string) {
	o.Status = &v
}

// GetSwitchInterface returns the SwitchInterface field value if set, zero value otherwise.
func (o *OperationNetworkInterface) GetSwitchInterface() string {
	if o == nil || IsNil(o.SwitchInterface) {
		var ret string
		return ret
	}
	return *o.SwitchInterface
}

// GetSwitchInterfaceOk returns a tuple with the SwitchInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationNetworkInterface) GetSwitchInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchInterface) {
		return nil, false
	}
	return o.SwitchInterface, true
}

// HasSwitchInterface returns a boolean if a field has been set.
func (o *OperationNetworkInterface) HasSwitchInterface() bool {
	if o != nil && !IsNil(o.SwitchInterface) {
		return true
	}

	return false
}

// SetSwitchInterface gets a reference to the given string and assigns it to the SwitchInterface field.
func (o *OperationNetworkInterface) SetSwitchInterface(v string) {
	o.SwitchInterface = &v
}

// GetSwitchName returns the SwitchName field value if set, zero value otherwise.
func (o *OperationNetworkInterface) GetSwitchName() string {
	if o == nil || IsNil(o.SwitchName) {
		var ret string
		return ret
	}
	return *o.SwitchName
}

// GetSwitchNameOk returns a tuple with the SwitchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationNetworkInterface) GetSwitchNameOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchName) {
		return nil, false
	}
	return o.SwitchName, true
}

// HasSwitchName returns a boolean if a field has been set.
func (o *OperationNetworkInterface) HasSwitchName() bool {
	if o != nil && !IsNil(o.SwitchName) {
		return true
	}

	return false
}

// SetSwitchName gets a reference to the given string and assigns it to the SwitchName field.
func (o *OperationNetworkInterface) SetSwitchName(v string) {
	o.SwitchName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OperationNetworkInterface) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationNetworkInterface) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OperationNetworkInterface) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OperationNetworkInterface) SetType(v string) {
	o.Type = &v
}

func (o OperationNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationNetworkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkSpeed) {
		toSerialize["linkSpeed"] = o.LinkSpeed
	}
	if !IsNil(o.OperStatus) {
		toSerialize["operStatus"] = o.OperStatus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SwitchInterface) {
		toSerialize["switchInterface"] = o.SwitchInterface
	}
	if !IsNil(o.SwitchName) {
		toSerialize["switchName"] = o.SwitchName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableOperationNetworkInterface struct {
	value *OperationNetworkInterface
	isSet bool
}

func (v NullableOperationNetworkInterface) Get() *OperationNetworkInterface {
	return v.value
}

func (v *NullableOperationNetworkInterface) Set(val *OperationNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationNetworkInterface(val *OperationNetworkInterface) *NullableOperationNetworkInterface {
	return &NullableOperationNetworkInterface{value: val, isSet: true}
}

func (v NullableOperationNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


