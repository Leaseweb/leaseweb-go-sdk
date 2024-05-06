# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**Compute** | Pointer to [**Compute**](Compute.md) |  | [optional] 
**Storage** | Pointer to [**Storage**](Storage.md) |  | [optional] 

## Methods

### NewPrice

`func NewPrice() *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Price) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *Price) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *Price) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *Price) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *Price) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCompute

`func (o *Price) GetCompute() Compute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *Price) GetComputeOk() (*Compute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *Price) SetCompute(v Compute)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *Price) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetStorage

`func (o *Price) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Price) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Price) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Price) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


