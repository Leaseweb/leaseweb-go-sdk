# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HourlyPrice** | **string** |  | 
**MonthlyPrice** | **string** |  | 

## Methods

### NewPrice

`func NewPrice(hourlyPrice string, monthlyPrice string, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHourlyPrice

`func (o *Price) GetHourlyPrice() string`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *Price) GetHourlyPriceOk() (*string, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *Price) SetHourlyPrice(v string)`

SetHourlyPrice sets HourlyPrice field to given value.


### GetMonthlyPrice

`func (o *Price) GetMonthlyPrice() string`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *Price) GetMonthlyPriceOk() (*string, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *Price) SetMonthlyPrice(v string)`

SetMonthlyPrice sets MonthlyPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


