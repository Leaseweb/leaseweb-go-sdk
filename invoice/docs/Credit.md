# Credit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date of the credit connected to the invoice | [optional] 
**Id** | Pointer to **string** | The unique id of the credit | [optional] 
**TaxAmount** | Pointer to **float32** | The total tax amount of the credit connected to the invoice | [optional] 
**Total** | Pointer to **float32** | The total amount of the credit connected to the invoice | [optional] 

## Methods

### NewCredit

`func NewCredit() *Credit`

NewCredit instantiates a new Credit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditWithDefaults

`func NewCreditWithDefaults() *Credit`

NewCreditWithDefaults instantiates a new Credit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *Credit) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Credit) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Credit) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Credit) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *Credit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Credit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *Credit) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *Credit) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *Credit) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *Credit) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotal

`func (o *Credit) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Credit) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Credit) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Credit) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


