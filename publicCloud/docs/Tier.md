# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | Pointer to **float32** | Traffic sent, in GB | [optional] 
**Price** | Pointer to **float32** | Total price of the tier, based on the usage. The first tier is free, so this will be 0 for the first tier. From tier 1 onwards, the usage has costs. Each tier has it own price. | [optional] 

## Methods

### NewTier

`func NewTier() *Tier`

NewTier instantiates a new Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWithDefaults

`func NewTierWithDefaults() *Tier`

NewTierWithDefaults instantiates a new Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *Tier) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Tier) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Tier) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Tier) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetPrice

`func (o *Tier) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Tier) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Tier) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Tier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


