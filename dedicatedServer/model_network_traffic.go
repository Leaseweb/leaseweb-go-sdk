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

// checks if the NetworkTraffic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkTraffic{}

// NetworkTraffic Network traffic associated to the server
type NetworkTraffic struct {
	// Type of network traffic
	Type NullableString `json:"type,omitempty"`
	ConnectivityType NullableString `json:"connectivityType,omitempty"`
	// Type of traffic
	TrafficType NullableString `json:"trafficType,omitempty"`
	// Unit in which the data traffic limit is represented
	DatatrafficUnit NullableString `json:"datatrafficUnit,omitempty"`
	// Data traffic limit
	DatatrafficLimit NullableFloat32 `json:"datatrafficLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkTraffic NetworkTraffic

// NewNetworkTraffic instantiates a new NetworkTraffic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkTraffic() *NetworkTraffic {
	this := NetworkTraffic{}
	return &this
}

// NewNetworkTrafficWithDefaults instantiates a new NetworkTraffic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkTrafficWithDefaults() *NetworkTraffic {
	this := NetworkTraffic{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkTraffic) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkTraffic) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *NetworkTraffic) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *NetworkTraffic) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *NetworkTraffic) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *NetworkTraffic) UnsetType() {
	o.Type.Unset()
}

// GetConnectivityType returns the ConnectivityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkTraffic) GetConnectivityType() string {
	if o == nil || IsNil(o.ConnectivityType.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectivityType.Get()
}

// GetConnectivityTypeOk returns a tuple with the ConnectivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkTraffic) GetConnectivityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectivityType.Get(), o.ConnectivityType.IsSet()
}

// HasConnectivityType returns a boolean if a field has been set.
func (o *NetworkTraffic) HasConnectivityType() bool {
	if o != nil && o.ConnectivityType.IsSet() {
		return true
	}

	return false
}

// SetConnectivityType gets a reference to the given NullableString and assigns it to the ConnectivityType field.
func (o *NetworkTraffic) SetConnectivityType(v string) {
	o.ConnectivityType.Set(&v)
}
// SetConnectivityTypeNil sets the value for ConnectivityType to be an explicit nil
func (o *NetworkTraffic) SetConnectivityTypeNil() {
	o.ConnectivityType.Set(nil)
}

// UnsetConnectivityType ensures that no value is present for ConnectivityType, not even an explicit nil
func (o *NetworkTraffic) UnsetConnectivityType() {
	o.ConnectivityType.Unset()
}

// GetTrafficType returns the TrafficType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkTraffic) GetTrafficType() string {
	if o == nil || IsNil(o.TrafficType.Get()) {
		var ret string
		return ret
	}
	return *o.TrafficType.Get()
}

// GetTrafficTypeOk returns a tuple with the TrafficType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkTraffic) GetTrafficTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafficType.Get(), o.TrafficType.IsSet()
}

// HasTrafficType returns a boolean if a field has been set.
func (o *NetworkTraffic) HasTrafficType() bool {
	if o != nil && o.TrafficType.IsSet() {
		return true
	}

	return false
}

// SetTrafficType gets a reference to the given NullableString and assigns it to the TrafficType field.
func (o *NetworkTraffic) SetTrafficType(v string) {
	o.TrafficType.Set(&v)
}
// SetTrafficTypeNil sets the value for TrafficType to be an explicit nil
func (o *NetworkTraffic) SetTrafficTypeNil() {
	o.TrafficType.Set(nil)
}

// UnsetTrafficType ensures that no value is present for TrafficType, not even an explicit nil
func (o *NetworkTraffic) UnsetTrafficType() {
	o.TrafficType.Unset()
}

// GetDatatrafficUnit returns the DatatrafficUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkTraffic) GetDatatrafficUnit() string {
	if o == nil || IsNil(o.DatatrafficUnit.Get()) {
		var ret string
		return ret
	}
	return *o.DatatrafficUnit.Get()
}

// GetDatatrafficUnitOk returns a tuple with the DatatrafficUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkTraffic) GetDatatrafficUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatatrafficUnit.Get(), o.DatatrafficUnit.IsSet()
}

// HasDatatrafficUnit returns a boolean if a field has been set.
func (o *NetworkTraffic) HasDatatrafficUnit() bool {
	if o != nil && o.DatatrafficUnit.IsSet() {
		return true
	}

	return false
}

// SetDatatrafficUnit gets a reference to the given NullableString and assigns it to the DatatrafficUnit field.
func (o *NetworkTraffic) SetDatatrafficUnit(v string) {
	o.DatatrafficUnit.Set(&v)
}
// SetDatatrafficUnitNil sets the value for DatatrafficUnit to be an explicit nil
func (o *NetworkTraffic) SetDatatrafficUnitNil() {
	o.DatatrafficUnit.Set(nil)
}

// UnsetDatatrafficUnit ensures that no value is present for DatatrafficUnit, not even an explicit nil
func (o *NetworkTraffic) UnsetDatatrafficUnit() {
	o.DatatrafficUnit.Unset()
}

// GetDatatrafficLimit returns the DatatrafficLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkTraffic) GetDatatrafficLimit() float32 {
	if o == nil || IsNil(o.DatatrafficLimit.Get()) {
		var ret float32
		return ret
	}
	return *o.DatatrafficLimit.Get()
}

// GetDatatrafficLimitOk returns a tuple with the DatatrafficLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkTraffic) GetDatatrafficLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatatrafficLimit.Get(), o.DatatrafficLimit.IsSet()
}

// HasDatatrafficLimit returns a boolean if a field has been set.
func (o *NetworkTraffic) HasDatatrafficLimit() bool {
	if o != nil && o.DatatrafficLimit.IsSet() {
		return true
	}

	return false
}

// SetDatatrafficLimit gets a reference to the given NullableFloat32 and assigns it to the DatatrafficLimit field.
func (o *NetworkTraffic) SetDatatrafficLimit(v float32) {
	o.DatatrafficLimit.Set(&v)
}
// SetDatatrafficLimitNil sets the value for DatatrafficLimit to be an explicit nil
func (o *NetworkTraffic) SetDatatrafficLimitNil() {
	o.DatatrafficLimit.Set(nil)
}

// UnsetDatatrafficLimit ensures that no value is present for DatatrafficLimit, not even an explicit nil
func (o *NetworkTraffic) UnsetDatatrafficLimit() {
	o.DatatrafficLimit.Unset()
}

func (o NetworkTraffic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkTraffic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.ConnectivityType.IsSet() {
		toSerialize["connectivityType"] = o.ConnectivityType.Get()
	}
	if o.TrafficType.IsSet() {
		toSerialize["trafficType"] = o.TrafficType.Get()
	}
	if o.DatatrafficUnit.IsSet() {
		toSerialize["datatrafficUnit"] = o.DatatrafficUnit.Get()
	}
	if o.DatatrafficLimit.IsSet() {
		toSerialize["datatrafficLimit"] = o.DatatrafficLimit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkTraffic) UnmarshalJSON(data []byte) (err error) {
	varNetworkTraffic := _NetworkTraffic{}

	err = json.Unmarshal(data, &varNetworkTraffic)

	if err != nil {
		return err
	}

	*o = NetworkTraffic(varNetworkTraffic)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "connectivityType")
		delete(additionalProperties, "trafficType")
		delete(additionalProperties, "datatrafficUnit")
		delete(additionalProperties, "datatrafficLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkTraffic struct {
	value *NetworkTraffic
	isSet bool
}

func (v NullableNetworkTraffic) Get() *NetworkTraffic {
	return v.value
}

func (v *NullableNetworkTraffic) Set(val *NetworkTraffic) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkTraffic) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkTraffic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkTraffic(val *NetworkTraffic) *NullableNetworkTraffic {
	return &NullableNetworkTraffic{value: val, isSet: true}
}

func (v NullableNetworkTraffic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkTraffic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


