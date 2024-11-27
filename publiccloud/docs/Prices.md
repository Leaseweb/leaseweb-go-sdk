# Prices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**CurrencySymbol** | **string** |  | 
**Compute** | [**Price**](Price.md) |  | 
**Storage** | [**Storage**](Storage.md) |  | 

## Methods

### NewPrices

`func NewPrices(currency string, currencySymbol string, compute Price, storage Storage, ) *Prices`

NewPrices instantiates a new Prices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricesWithDefaults

`func NewPricesWithDefaults() *Prices`

NewPricesWithDefaults instantiates a new Prices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Prices) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Prices) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Prices) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrencySymbol

`func (o *Prices) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *Prices) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *Prices) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.


### GetCompute

`func (o *Prices) GetCompute() Price`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *Prices) GetComputeOk() (*Price, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *Prices) SetCompute(v Price)`

SetCompute sets Compute field to given value.


### GetStorage

`func (o *Prices) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Prices) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Prices) SetStorage(v Storage)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


