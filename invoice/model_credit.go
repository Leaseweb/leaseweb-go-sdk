/*
Invoices

> The base URL for this API is: **https://api.leaseweb.com/invoices/v1/_**

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoice

import (
	"encoding/json"
)

// checks if the Credit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credit{}

// Credit struct for Credit
type Credit struct {
	// The date of the credit connected to the invoice
	Date *string `json:"date,omitempty"`
	// The unique id of the credit
	Id *string `json:"id,omitempty"`
	// The total tax amount of the credit connected to the invoice
	TaxAmount *float32 `json:"taxAmount,omitempty"`
	// The total amount of the credit connected to the invoice
	Total *float32 `json:"total,omitempty"`
}

// NewCredit instantiates a new Credit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredit() *Credit {
	this := Credit{}
	return &this
}

// NewCreditWithDefaults instantiates a new Credit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditWithDefaults() *Credit {
	this := Credit{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Credit) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Credit) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Credit) SetDate(v string) {
	o.Date = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Credit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Credit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Credit) SetId(v string) {
	o.Id = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *Credit) GetTaxAmount() float32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *Credit) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float32 and assigns it to the TaxAmount field.
func (o *Credit) SetTaxAmount(v float32) {
	o.TaxAmount = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Credit) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Credit) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *Credit) SetTotal(v float32) {
	o.Total = &v
}

func (o Credit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Credit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableCredit struct {
	value *Credit
	isSet bool
}

func (v NullableCredit) Get() *Credit {
	return v.value
}

func (v *NullableCredit) Set(val *Credit) {
	v.value = val
	v.isSet = true
}

func (v NullableCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredit(val *Credit) *NullableCredit {
	return &NullableCredit{value: val, isSet: true}
}

func (v NullableCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

