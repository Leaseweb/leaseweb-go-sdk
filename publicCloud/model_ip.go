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

// checks if the Ip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ip{}

// Ip struct for Ip
type Ip struct {
	// Ip Address
	Ip *string `json:"ip,omitempty"`
	// The number of leading bits in the IP address
	PrefixLength *string `json:"prefixLength,omitempty"`
	// Ip version
	Version *int32 `json:"version,omitempty"`
	// Whether or not the IP has been nulled
	NullRouted *bool `json:"nullRouted,omitempty"`
	MainIp *bool `json:"mainIp,omitempty"`
	ReverseLookup NullableString `json:"reverseLookup,omitempty"`
	Ddos *Ddos `json:"ddos,omitempty"`
}

// NewIp instantiates a new Ip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIp() *Ip {
	this := Ip{}
	return &this
}

// NewIpWithDefaults instantiates a new Ip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpWithDefaults() *Ip {
	this := Ip{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Ip) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *Ip) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Ip) SetIp(v string) {
	o.Ip = &v
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *Ip) GetPrefixLength() string {
	if o == nil || IsNil(o.PrefixLength) {
		var ret string
		return ret
	}
	return *o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetPrefixLengthOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixLength) {
		return nil, false
	}
	return o.PrefixLength, true
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *Ip) HasPrefixLength() bool {
	if o != nil && !IsNil(o.PrefixLength) {
		return true
	}

	return false
}

// SetPrefixLength gets a reference to the given string and assigns it to the PrefixLength field.
func (o *Ip) SetPrefixLength(v string) {
	o.PrefixLength = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Ip) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Ip) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Ip) SetVersion(v int32) {
	o.Version = &v
}

// GetNullRouted returns the NullRouted field value if set, zero value otherwise.
func (o *Ip) GetNullRouted() bool {
	if o == nil || IsNil(o.NullRouted) {
		var ret bool
		return ret
	}
	return *o.NullRouted
}

// GetNullRoutedOk returns a tuple with the NullRouted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetNullRoutedOk() (*bool, bool) {
	if o == nil || IsNil(o.NullRouted) {
		return nil, false
	}
	return o.NullRouted, true
}

// HasNullRouted returns a boolean if a field has been set.
func (o *Ip) HasNullRouted() bool {
	if o != nil && !IsNil(o.NullRouted) {
		return true
	}

	return false
}

// SetNullRouted gets a reference to the given bool and assigns it to the NullRouted field.
func (o *Ip) SetNullRouted(v bool) {
	o.NullRouted = &v
}

// GetMainIp returns the MainIp field value if set, zero value otherwise.
func (o *Ip) GetMainIp() bool {
	if o == nil || IsNil(o.MainIp) {
		var ret bool
		return ret
	}
	return *o.MainIp
}

// GetMainIpOk returns a tuple with the MainIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetMainIpOk() (*bool, bool) {
	if o == nil || IsNil(o.MainIp) {
		return nil, false
	}
	return o.MainIp, true
}

// HasMainIp returns a boolean if a field has been set.
func (o *Ip) HasMainIp() bool {
	if o != nil && !IsNil(o.MainIp) {
		return true
	}

	return false
}

// SetMainIp gets a reference to the given bool and assigns it to the MainIp field.
func (o *Ip) SetMainIp(v bool) {
	o.MainIp = &v
}

// GetReverseLookup returns the ReverseLookup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ip) GetReverseLookup() string {
	if o == nil || IsNil(o.ReverseLookup.Get()) {
		var ret string
		return ret
	}
	return *o.ReverseLookup.Get()
}

// GetReverseLookupOk returns a tuple with the ReverseLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ip) GetReverseLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseLookup.Get(), o.ReverseLookup.IsSet()
}

// HasReverseLookup returns a boolean if a field has been set.
func (o *Ip) HasReverseLookup() bool {
	if o != nil && o.ReverseLookup.IsSet() {
		return true
	}

	return false
}

// SetReverseLookup gets a reference to the given NullableString and assigns it to the ReverseLookup field.
func (o *Ip) SetReverseLookup(v string) {
	o.ReverseLookup.Set(&v)
}
// SetReverseLookupNil sets the value for ReverseLookup to be an explicit nil
func (o *Ip) SetReverseLookupNil() {
	o.ReverseLookup.Set(nil)
}

// UnsetReverseLookup ensures that no value is present for ReverseLookup, not even an explicit nil
func (o *Ip) UnsetReverseLookup() {
	o.ReverseLookup.Unset()
}

// GetDdos returns the Ddos field value if set, zero value otherwise.
func (o *Ip) GetDdos() Ddos {
	if o == nil || IsNil(o.Ddos) {
		var ret Ddos
		return ret
	}
	return *o.Ddos
}

// GetDdosOk returns a tuple with the Ddos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetDdosOk() (*Ddos, bool) {
	if o == nil || IsNil(o.Ddos) {
		return nil, false
	}
	return o.Ddos, true
}

// HasDdos returns a boolean if a field has been set.
func (o *Ip) HasDdos() bool {
	if o != nil && !IsNil(o.Ddos) {
		return true
	}

	return false
}

// SetDdos gets a reference to the given Ddos and assigns it to the Ddos field.
func (o *Ip) SetDdos(v Ddos) {
	o.Ddos = &v
}

func (o Ip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.PrefixLength) {
		toSerialize["prefixLength"] = o.PrefixLength
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.NullRouted) {
		toSerialize["nullRouted"] = o.NullRouted
	}
	if !IsNil(o.MainIp) {
		toSerialize["mainIp"] = o.MainIp
	}
	if o.ReverseLookup.IsSet() {
		toSerialize["reverseLookup"] = o.ReverseLookup.Get()
	}
	if !IsNil(o.Ddos) {
		toSerialize["ddos"] = o.Ddos
	}
	return toSerialize, nil
}

type NullableIp struct {
	value *Ip
	isSet bool
}

func (v NullableIp) Get() *Ip {
	return v.value
}

func (v *NullableIp) Set(val *Ip) {
	v.value = val
	v.isSet = true
}

func (v NullableIp) IsSet() bool {
	return v.isSet
}

func (v *NullableIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIp(val *Ip) *NullableIp {
	return &NullableIp{value: val, isSet: true}
}

func (v NullableIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

