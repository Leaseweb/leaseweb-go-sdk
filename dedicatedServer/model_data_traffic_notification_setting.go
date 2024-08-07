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

// checks if the DataTrafficNotificationSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataTrafficNotificationSetting{}

// DataTrafficNotificationSetting struct for DataTrafficNotificationSetting
type DataTrafficNotificationSetting struct {
	// An array of notification setting actions
	Actions []Actions `json:"actions,omitempty"`
	// Frequency
	Frequency *string `json:"frequency,omitempty"`
	// Identifier of the notification setting
	Id *string `json:"id,omitempty"`
	// Date timestamp when the system last checked the server for threshold limit
	LastCheckedAt NullableString `json:"lastCheckedAt,omitempty"`
	// Threshold Value
	Threshold *string `json:"threshold,omitempty"`
	// Date timestamp when the threshold exceeded the limit
	ThresholdExceededAt NullableString `json:"thresholdExceededAt,omitempty"`
	// Unit
	Unit *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataTrafficNotificationSetting DataTrafficNotificationSetting

// NewDataTrafficNotificationSetting instantiates a new DataTrafficNotificationSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTrafficNotificationSetting() *DataTrafficNotificationSetting {
	this := DataTrafficNotificationSetting{}
	return &this
}

// NewDataTrafficNotificationSettingWithDefaults instantiates a new DataTrafficNotificationSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTrafficNotificationSettingWithDefaults() *DataTrafficNotificationSetting {
	this := DataTrafficNotificationSetting{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *DataTrafficNotificationSetting) GetActions() []Actions {
	if o == nil || IsNil(o.Actions) {
		var ret []Actions
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSetting) GetActionsOk() ([]Actions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DataTrafficNotificationSetting) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []Actions and assigns it to the Actions field.
func (o *DataTrafficNotificationSetting) SetActions(v []Actions) {
	o.Actions = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *DataTrafficNotificationSetting) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSetting) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *DataTrafficNotificationSetting) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *DataTrafficNotificationSetting) SetFrequency(v string) {
	o.Frequency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataTrafficNotificationSetting) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSetting) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataTrafficNotificationSetting) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataTrafficNotificationSetting) SetId(v string) {
	o.Id = &v
}

// GetLastCheckedAt returns the LastCheckedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTrafficNotificationSetting) GetLastCheckedAt() string {
	if o == nil || IsNil(o.LastCheckedAt.Get()) {
		var ret string
		return ret
	}
	return *o.LastCheckedAt.Get()
}

// GetLastCheckedAtOk returns a tuple with the LastCheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTrafficNotificationSetting) GetLastCheckedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastCheckedAt.Get(), o.LastCheckedAt.IsSet()
}

// HasLastCheckedAt returns a boolean if a field has been set.
func (o *DataTrafficNotificationSetting) HasLastCheckedAt() bool {
	if o != nil && o.LastCheckedAt.IsSet() {
		return true
	}

	return false
}

// SetLastCheckedAt gets a reference to the given NullableString and assigns it to the LastCheckedAt field.
func (o *DataTrafficNotificationSetting) SetLastCheckedAt(v string) {
	o.LastCheckedAt.Set(&v)
}
// SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil
func (o *DataTrafficNotificationSetting) SetLastCheckedAtNil() {
	o.LastCheckedAt.Set(nil)
}

// UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
func (o *DataTrafficNotificationSetting) UnsetLastCheckedAt() {
	o.LastCheckedAt.Unset()
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *DataTrafficNotificationSetting) GetThreshold() string {
	if o == nil || IsNil(o.Threshold) {
		var ret string
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSetting) GetThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *DataTrafficNotificationSetting) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given string and assigns it to the Threshold field.
func (o *DataTrafficNotificationSetting) SetThreshold(v string) {
	o.Threshold = &v
}

// GetThresholdExceededAt returns the ThresholdExceededAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTrafficNotificationSetting) GetThresholdExceededAt() string {
	if o == nil || IsNil(o.ThresholdExceededAt.Get()) {
		var ret string
		return ret
	}
	return *o.ThresholdExceededAt.Get()
}

// GetThresholdExceededAtOk returns a tuple with the ThresholdExceededAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTrafficNotificationSetting) GetThresholdExceededAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThresholdExceededAt.Get(), o.ThresholdExceededAt.IsSet()
}

// HasThresholdExceededAt returns a boolean if a field has been set.
func (o *DataTrafficNotificationSetting) HasThresholdExceededAt() bool {
	if o != nil && o.ThresholdExceededAt.IsSet() {
		return true
	}

	return false
}

// SetThresholdExceededAt gets a reference to the given NullableString and assigns it to the ThresholdExceededAt field.
func (o *DataTrafficNotificationSetting) SetThresholdExceededAt(v string) {
	o.ThresholdExceededAt.Set(&v)
}
// SetThresholdExceededAtNil sets the value for ThresholdExceededAt to be an explicit nil
func (o *DataTrafficNotificationSetting) SetThresholdExceededAtNil() {
	o.ThresholdExceededAt.Set(nil)
}

// UnsetThresholdExceededAt ensures that no value is present for ThresholdExceededAt, not even an explicit nil
func (o *DataTrafficNotificationSetting) UnsetThresholdExceededAt() {
	o.ThresholdExceededAt.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DataTrafficNotificationSetting) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSetting) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DataTrafficNotificationSetting) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DataTrafficNotificationSetting) SetUnit(v string) {
	o.Unit = &v
}

func (o DataTrafficNotificationSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataTrafficNotificationSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.LastCheckedAt.IsSet() {
		toSerialize["lastCheckedAt"] = o.LastCheckedAt.Get()
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if o.ThresholdExceededAt.IsSet() {
		toSerialize["thresholdExceededAt"] = o.ThresholdExceededAt.Get()
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataTrafficNotificationSetting) UnmarshalJSON(data []byte) (err error) {
	varDataTrafficNotificationSetting := _DataTrafficNotificationSetting{}

	err = json.Unmarshal(data, &varDataTrafficNotificationSetting)

	if err != nil {
		return err
	}

	*o = DataTrafficNotificationSetting(varDataTrafficNotificationSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "frequency")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastCheckedAt")
		delete(additionalProperties, "threshold")
		delete(additionalProperties, "thresholdExceededAt")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataTrafficNotificationSetting struct {
	value *DataTrafficNotificationSetting
	isSet bool
}

func (v NullableDataTrafficNotificationSetting) Get() *DataTrafficNotificationSetting {
	return v.value
}

func (v *NullableDataTrafficNotificationSetting) Set(val *DataTrafficNotificationSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTrafficNotificationSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTrafficNotificationSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTrafficNotificationSetting(val *DataTrafficNotificationSetting) *NullableDataTrafficNotificationSetting {
	return &NullableDataTrafficNotificationSetting{value: val, isSet: true}
}

func (v NullableDataTrafficNotificationSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTrafficNotificationSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


