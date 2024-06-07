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

// checks if the GetServerNullRouteHistoryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerNullRouteHistoryResult{}

// GetServerNullRouteHistoryResult struct for GetServerNullRouteHistoryResult
type GetServerNullRouteHistoryResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	// An array of server null route events
	NullRoutes []NullRoute `json:"nullRoutes,omitempty"`
}

// NewGetServerNullRouteHistoryResult instantiates a new GetServerNullRouteHistoryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerNullRouteHistoryResult() *GetServerNullRouteHistoryResult {
	this := GetServerNullRouteHistoryResult{}
	return &this
}

// NewGetServerNullRouteHistoryResultWithDefaults instantiates a new GetServerNullRouteHistoryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerNullRouteHistoryResultWithDefaults() *GetServerNullRouteHistoryResult {
	this := GetServerNullRouteHistoryResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetServerNullRouteHistoryResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerNullRouteHistoryResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetServerNullRouteHistoryResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetServerNullRouteHistoryResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetNullRoutes returns the NullRoutes field value if set, zero value otherwise.
func (o *GetServerNullRouteHistoryResult) GetNullRoutes() []NullRoute {
	if o == nil || IsNil(o.NullRoutes) {
		var ret []NullRoute
		return ret
	}
	return o.NullRoutes
}

// GetNullRoutesOk returns a tuple with the NullRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerNullRouteHistoryResult) GetNullRoutesOk() ([]NullRoute, bool) {
	if o == nil || IsNil(o.NullRoutes) {
		return nil, false
	}
	return o.NullRoutes, true
}

// HasNullRoutes returns a boolean if a field has been set.
func (o *GetServerNullRouteHistoryResult) HasNullRoutes() bool {
	if o != nil && !IsNil(o.NullRoutes) {
		return true
	}

	return false
}

// SetNullRoutes gets a reference to the given []NullRoute and assigns it to the NullRoutes field.
func (o *GetServerNullRouteHistoryResult) SetNullRoutes(v []NullRoute) {
	o.NullRoutes = v
}

func (o GetServerNullRouteHistoryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerNullRouteHistoryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.NullRoutes) {
		toSerialize["nullRoutes"] = o.NullRoutes
	}
	return toSerialize, nil
}

type NullableGetServerNullRouteHistoryResult struct {
	value *GetServerNullRouteHistoryResult
	isSet bool
}

func (v NullableGetServerNullRouteHistoryResult) Get() *GetServerNullRouteHistoryResult {
	return v.value
}

func (v *NullableGetServerNullRouteHistoryResult) Set(val *GetServerNullRouteHistoryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerNullRouteHistoryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerNullRouteHistoryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerNullRouteHistoryResult(val *GetServerNullRouteHistoryResult) *NullableGetServerNullRouteHistoryResult {
	return &NullableGetServerNullRouteHistoryResult{value: val, isSet: true}
}

func (v NullableGetServerNullRouteHistoryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerNullRouteHistoryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


