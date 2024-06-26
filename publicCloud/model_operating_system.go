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

// checks if the OperatingSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatingSystem{}

// OperatingSystem struct for OperatingSystem
type OperatingSystem struct {
	Id OperatingSystemId `json:"id"`
	Name string `json:"name"`
	Version string `json:"version"`
	Family string `json:"family"`
	Flavour string `json:"flavour"`
	Architecture string `json:"architecture"`
}

type _OperatingSystem OperatingSystem

// NewOperatingSystem instantiates a new OperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingSystem(id OperatingSystemId, name string, version string, family string, flavour string, architecture string) *OperatingSystem {
	this := OperatingSystem{}
	this.Id = id
	this.Name = name
	this.Version = version
	this.Family = family
	this.Flavour = flavour
	this.Architecture = architecture
	return &this
}

// NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingSystemWithDefaults() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// GetId returns the Id field value
func (o *OperatingSystem) GetId() OperatingSystemId {
	if o == nil {
		var ret OperatingSystemId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetIdOk() (*OperatingSystemId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OperatingSystem) SetId(v OperatingSystemId) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OperatingSystem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OperatingSystem) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *OperatingSystem) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *OperatingSystem) SetVersion(v string) {
	o.Version = v
}

// GetFamily returns the Family field value
func (o *OperatingSystem) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *OperatingSystem) SetFamily(v string) {
	o.Family = v
}

// GetFlavour returns the Flavour field value
func (o *OperatingSystem) GetFlavour() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetFlavourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flavour, true
}

// SetFlavour sets field value
func (o *OperatingSystem) SetFlavour(v string) {
	o.Flavour = v
}

// GetArchitecture returns the Architecture field value
func (o *OperatingSystem) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *OperatingSystem) SetArchitecture(v string) {
	o.Architecture = v
}

func (o OperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatingSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	toSerialize["family"] = o.Family
	toSerialize["flavour"] = o.Flavour
	toSerialize["architecture"] = o.Architecture
	return toSerialize, nil
}

func (o *OperatingSystem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"version",
		"family",
		"flavour",
		"architecture",
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

	varOperatingSystem := _OperatingSystem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperatingSystem)

	if err != nil {
		return err
	}

	*o = OperatingSystem(varOperatingSystem)

	return err
}

type NullableOperatingSystem struct {
	value *OperatingSystem
	isSet bool
}

func (v NullableOperatingSystem) Get() *OperatingSystem {
	return v.value
}

func (v *NullableOperatingSystem) Set(val *OperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingSystem(val *OperatingSystem) *NullableOperatingSystem {
	return &NullableOperatingSystem{value: val, isSet: true}
}

func (v NullableOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


