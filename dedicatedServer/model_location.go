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

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location Device location
type Location struct {
	Site *string `json:"site,omitempty"`
	Suite *string `json:"suite,omitempty"`
	Rack *string `json:"rack,omitempty"`
	Unit *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Location Location

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *Location) GetSite() string {
	if o == nil || IsNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetSiteOk() (*string, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *Location) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *Location) SetSite(v string) {
	o.Site = &v
}

// GetSuite returns the Suite field value if set, zero value otherwise.
func (o *Location) GetSuite() string {
	if o == nil || IsNil(o.Suite) {
		var ret string
		return ret
	}
	return *o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetSuiteOk() (*string, bool) {
	if o == nil || IsNil(o.Suite) {
		return nil, false
	}
	return o.Suite, true
}

// HasSuite returns a boolean if a field has been set.
func (o *Location) HasSuite() bool {
	if o != nil && !IsNil(o.Suite) {
		return true
	}

	return false
}

// SetSuite gets a reference to the given string and assigns it to the Suite field.
func (o *Location) SetSuite(v string) {
	o.Suite = &v
}

// GetRack returns the Rack field value if set, zero value otherwise.
func (o *Location) GetRack() string {
	if o == nil || IsNil(o.Rack) {
		var ret string
		return ret
	}
	return *o.Rack
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetRackOk() (*string, bool) {
	if o == nil || IsNil(o.Rack) {
		return nil, false
	}
	return o.Rack, true
}

// HasRack returns a boolean if a field has been set.
func (o *Location) HasRack() bool {
	if o != nil && !IsNil(o.Rack) {
		return true
	}

	return false
}

// SetRack gets a reference to the given string and assigns it to the Rack field.
func (o *Location) SetRack(v string) {
	o.Rack = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Location) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Location) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Location) SetUnit(v string) {
	o.Unit = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.Suite) {
		toSerialize["suite"] = o.Suite
	}
	if !IsNil(o.Rack) {
		toSerialize["rack"] = o.Rack
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Location) UnmarshalJSON(data []byte) (err error) {
	varLocation := _Location{}

	err = json.Unmarshal(data, &varLocation)

	if err != nil {
		return err
	}

	*o = Location(varLocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site")
		delete(additionalProperties, "suite")
		delete(additionalProperties, "rack")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


