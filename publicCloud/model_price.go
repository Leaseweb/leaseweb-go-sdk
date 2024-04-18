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

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Price{}

// Price struct for Price
type Price struct {
	Currency *string `json:"currency,omitempty"`
	CurrencySymbol *string `json:"currencySymbol,omitempty"`
	Compute *Compute `json:"compute,omitempty"`
	Storage *Storage `json:"storage,omitempty"`
}

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice() *Price {
	this := Price{}
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Price) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Price) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Price) SetCurrency(v string) {
	o.Currency = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *Price) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *Price) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *Price) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *Price) GetCompute() Compute {
	if o == nil || IsNil(o.Compute) {
		var ret Compute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetComputeOk() (*Compute, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *Price) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given Compute and assigns it to the Compute field.
func (o *Price) SetCompute(v Compute) {
	o.Compute = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *Price) GetStorage() Storage {
	if o == nil || IsNil(o.Storage) {
		var ret Storage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetStorageOk() (*Storage, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *Price) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given Storage and assigns it to the Storage field.
func (o *Price) SetStorage(v Storage) {
	o.Storage = &v
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CurrencySymbol) {
		toSerialize["currencySymbol"] = o.CurrencySymbol
	}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	return toSerialize, nil
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


