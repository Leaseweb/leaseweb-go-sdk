# GetInvoiceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credits** | Pointer to [**[]Credit**](Credit.md) | All the credits attached to the invoice | [optional] 
**Currency** | Pointer to **string** | The currency of the invoice. | [optional] 
**Date** | Pointer to **string** | The date the invoice was issued | [optional] 
**DueDate** | Pointer to **string** | The date the invoice is due for payment | [optional] 
**Id** | Pointer to **string** | The unique id of the invoice | [optional] 
**IsPartialPaymentAllowed** | Pointer to **bool** | The invoice can be paid partially | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | All the line items attached to the invoice | [optional] 
**OpenAmount** | Pointer to **float32** | The open amount of the invoice | [optional] 
**Status** | Pointer to **string** | The status of the invoice. | [optional] 
**TaxAmount** | Pointer to **float32** | The tax amount of the invoice | [optional] 
**Total** | Pointer to **float32** | The total amount of the invoice | [optional] 

## Methods

### NewGetInvoiceResult

`func NewGetInvoiceResult() *GetInvoiceResult`

NewGetInvoiceResult instantiates a new GetInvoiceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceResultWithDefaults

`func NewGetInvoiceResultWithDefaults() *GetInvoiceResult`

NewGetInvoiceResultWithDefaults instantiates a new GetInvoiceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredits

`func (o *GetInvoiceResult) GetCredits() []Credit`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *GetInvoiceResult) GetCreditsOk() (*[]Credit, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *GetInvoiceResult) SetCredits(v []Credit)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *GetInvoiceResult) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetCurrency

`func (o *GetInvoiceResult) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetInvoiceResult) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetInvoiceResult) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetInvoiceResult) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *GetInvoiceResult) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetInvoiceResult) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetInvoiceResult) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetInvoiceResult) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDueDate

`func (o *GetInvoiceResult) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *GetInvoiceResult) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *GetInvoiceResult) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *GetInvoiceResult) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *GetInvoiceResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInvoiceResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInvoiceResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInvoiceResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPartialPaymentAllowed

`func (o *GetInvoiceResult) GetIsPartialPaymentAllowed() bool`

GetIsPartialPaymentAllowed returns the IsPartialPaymentAllowed field if non-nil, zero value otherwise.

### GetIsPartialPaymentAllowedOk

`func (o *GetInvoiceResult) GetIsPartialPaymentAllowedOk() (*bool, bool)`

GetIsPartialPaymentAllowedOk returns a tuple with the IsPartialPaymentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartialPaymentAllowed

`func (o *GetInvoiceResult) SetIsPartialPaymentAllowed(v bool)`

SetIsPartialPaymentAllowed sets IsPartialPaymentAllowed field to given value.

### HasIsPartialPaymentAllowed

`func (o *GetInvoiceResult) HasIsPartialPaymentAllowed() bool`

HasIsPartialPaymentAllowed returns a boolean if a field has been set.

### GetLineItems

`func (o *GetInvoiceResult) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *GetInvoiceResult) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *GetInvoiceResult) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *GetInvoiceResult) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetOpenAmount

`func (o *GetInvoiceResult) GetOpenAmount() float32`

GetOpenAmount returns the OpenAmount field if non-nil, zero value otherwise.

### GetOpenAmountOk

`func (o *GetInvoiceResult) GetOpenAmountOk() (*float32, bool)`

GetOpenAmountOk returns a tuple with the OpenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmount

`func (o *GetInvoiceResult) SetOpenAmount(v float32)`

SetOpenAmount sets OpenAmount field to given value.

### HasOpenAmount

`func (o *GetInvoiceResult) HasOpenAmount() bool`

HasOpenAmount returns a boolean if a field has been set.

### GetStatus

`func (o *GetInvoiceResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInvoiceResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInvoiceResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInvoiceResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxAmount

`func (o *GetInvoiceResult) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *GetInvoiceResult) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *GetInvoiceResult) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *GetInvoiceResult) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotal

`func (o *GetInvoiceResult) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInvoiceResult) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInvoiceResult) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetInvoiceResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


