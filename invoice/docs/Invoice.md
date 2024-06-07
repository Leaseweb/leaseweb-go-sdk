# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of the invoice. | [optional] 
**Date** | Pointer to **string** | The date the invoice was issued | [optional] 
**DueDate** | Pointer to **string** | The date the invoice is due for payment | [optional] 
**Id** | Pointer to **string** | The unique id of the invoice | [optional] 
**IsPartialPaymentAllowed** | Pointer to **bool** | The invoice can be paid partially | [optional] 
**OpenAmount** | Pointer to **float32** | The open amount of the invoice | [optional] 
**Status** | Pointer to **string** | The status of the invoice. | [optional] 
**TaxAmount** | Pointer to **float32** | The tax amount of the invoice | [optional] 
**Total** | Pointer to **float32** | The total amount of the invoice | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *Invoice) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoice) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoice) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Invoice) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDueDate

`func (o *Invoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Invoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPartialPaymentAllowed

`func (o *Invoice) GetIsPartialPaymentAllowed() bool`

GetIsPartialPaymentAllowed returns the IsPartialPaymentAllowed field if non-nil, zero value otherwise.

### GetIsPartialPaymentAllowedOk

`func (o *Invoice) GetIsPartialPaymentAllowedOk() (*bool, bool)`

GetIsPartialPaymentAllowedOk returns a tuple with the IsPartialPaymentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartialPaymentAllowed

`func (o *Invoice) SetIsPartialPaymentAllowed(v bool)`

SetIsPartialPaymentAllowed sets IsPartialPaymentAllowed field to given value.

### HasIsPartialPaymentAllowed

`func (o *Invoice) HasIsPartialPaymentAllowed() bool`

HasIsPartialPaymentAllowed returns a boolean if a field has been set.

### GetOpenAmount

`func (o *Invoice) GetOpenAmount() float32`

GetOpenAmount returns the OpenAmount field if non-nil, zero value otherwise.

### GetOpenAmountOk

`func (o *Invoice) GetOpenAmountOk() (*float32, bool)`

GetOpenAmountOk returns a tuple with the OpenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmount

`func (o *Invoice) SetOpenAmount(v float32)`

SetOpenAmount sets OpenAmount field to given value.

### HasOpenAmount

`func (o *Invoice) HasOpenAmount() bool`

HasOpenAmount returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxAmount

`func (o *Invoice) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *Invoice) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *Invoice) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *Invoice) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotal

`func (o *Invoice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Invoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


