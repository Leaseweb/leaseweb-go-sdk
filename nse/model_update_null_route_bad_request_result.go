/*
nse

This documents the rest api nse api provides.

API version: v2
Contact: development-networkautomation@leaseweb.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nse

import (
	"encoding/json"
)

// checks if the UpdateNullRouteBadRequestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNullRouteBadRequestResult{}

// UpdateNullRouteBadRequestResult struct for UpdateNullRouteBadRequestResult
type UpdateNullRouteBadRequestResult struct {
	// The correlation ID of the current request.
	CorrelationId *string `json:"correlationId,omitempty"`
	// The error code.
	ErrorCode *string `json:"errorCode,omitempty"`
	// An object describing the errors for the current request.
	ErrorDetails map[string]interface{} `json:"errorDetails,omitempty"`
	// A human friendly description of the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewUpdateNullRouteBadRequestResult instantiates a new UpdateNullRouteBadRequestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNullRouteBadRequestResult() *UpdateNullRouteBadRequestResult {
	this := UpdateNullRouteBadRequestResult{}
	return &this
}

// NewUpdateNullRouteBadRequestResultWithDefaults instantiates a new UpdateNullRouteBadRequestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNullRouteBadRequestResultWithDefaults() *UpdateNullRouteBadRequestResult {
	this := UpdateNullRouteBadRequestResult{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *UpdateNullRouteBadRequestResult) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNullRouteBadRequestResult) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *UpdateNullRouteBadRequestResult) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *UpdateNullRouteBadRequestResult) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *UpdateNullRouteBadRequestResult) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNullRouteBadRequestResult) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *UpdateNullRouteBadRequestResult) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *UpdateNullRouteBadRequestResult) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *UpdateNullRouteBadRequestResult) GetErrorDetails() map[string]interface{} {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNullRouteBadRequestResult) GetErrorDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return map[string]interface{}{}, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *UpdateNullRouteBadRequestResult) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given map[string]interface{} and assigns it to the ErrorDetails field.
func (o *UpdateNullRouteBadRequestResult) SetErrorDetails(v map[string]interface{}) {
	o.ErrorDetails = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *UpdateNullRouteBadRequestResult) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNullRouteBadRequestResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *UpdateNullRouteBadRequestResult) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *UpdateNullRouteBadRequestResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o UpdateNullRouteBadRequestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNullRouteBadRequestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableUpdateNullRouteBadRequestResult struct {
	value *UpdateNullRouteBadRequestResult
	isSet bool
}

func (v NullableUpdateNullRouteBadRequestResult) Get() *UpdateNullRouteBadRequestResult {
	return v.value
}

func (v *NullableUpdateNullRouteBadRequestResult) Set(val *UpdateNullRouteBadRequestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNullRouteBadRequestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNullRouteBadRequestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNullRouteBadRequestResult(val *UpdateNullRouteBadRequestResult) *NullableUpdateNullRouteBadRequestResult {
	return &NullableUpdateNullRouteBadRequestResult{value: val, isSet: true}
}

func (v NullableUpdateNullRouteBadRequestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNullRouteBadRequestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


