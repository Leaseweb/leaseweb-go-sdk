# GetInvoiceListResultInvoicesInner

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

### NewGetInvoiceListResultInvoicesInner

`func NewGetInvoiceListResultInvoicesInner() *GetInvoiceListResultInvoicesInner`

NewGetInvoiceListResultInvoicesInner instantiates a new GetInvoiceListResultInvoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceListResultInvoicesInnerWithDefaults

`func NewGetInvoiceListResultInvoicesInnerWithDefaults() *GetInvoiceListResultInvoicesInner`

NewGetInvoiceListResultInvoicesInnerWithDefaults instantiates a new GetInvoiceListResultInvoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *GetInvoiceListResultInvoicesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetInvoiceListResultInvoicesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetInvoiceListResultInvoicesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetInvoiceListResultInvoicesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *GetInvoiceListResultInvoicesInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetInvoiceListResultInvoicesInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetInvoiceListResultInvoicesInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetInvoiceListResultInvoicesInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDueDate

`func (o *GetInvoiceListResultInvoicesInner) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *GetInvoiceListResultInvoicesInner) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *GetInvoiceListResultInvoicesInner) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *GetInvoiceListResultInvoicesInner) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *GetInvoiceListResultInvoicesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInvoiceListResultInvoicesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInvoiceListResultInvoicesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInvoiceListResultInvoicesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPartialPaymentAllowed

`func (o *GetInvoiceListResultInvoicesInner) GetIsPartialPaymentAllowed() bool`

GetIsPartialPaymentAllowed returns the IsPartialPaymentAllowed field if non-nil, zero value otherwise.

### GetIsPartialPaymentAllowedOk

`func (o *GetInvoiceListResultInvoicesInner) GetIsPartialPaymentAllowedOk() (*bool, bool)`

GetIsPartialPaymentAllowedOk returns a tuple with the IsPartialPaymentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartialPaymentAllowed

`func (o *GetInvoiceListResultInvoicesInner) SetIsPartialPaymentAllowed(v bool)`

SetIsPartialPaymentAllowed sets IsPartialPaymentAllowed field to given value.

### HasIsPartialPaymentAllowed

`func (o *GetInvoiceListResultInvoicesInner) HasIsPartialPaymentAllowed() bool`

HasIsPartialPaymentAllowed returns a boolean if a field has been set.

### GetOpenAmount

`func (o *GetInvoiceListResultInvoicesInner) GetOpenAmount() float32`

GetOpenAmount returns the OpenAmount field if non-nil, zero value otherwise.

### GetOpenAmountOk

`func (o *GetInvoiceListResultInvoicesInner) GetOpenAmountOk() (*float32, bool)`

GetOpenAmountOk returns a tuple with the OpenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmount

`func (o *GetInvoiceListResultInvoicesInner) SetOpenAmount(v float32)`

SetOpenAmount sets OpenAmount field to given value.

### HasOpenAmount

`func (o *GetInvoiceListResultInvoicesInner) HasOpenAmount() bool`

HasOpenAmount returns a boolean if a field has been set.

### GetStatus

`func (o *GetInvoiceListResultInvoicesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInvoiceListResultInvoicesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInvoiceListResultInvoicesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInvoiceListResultInvoicesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxAmount

`func (o *GetInvoiceListResultInvoicesInner) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *GetInvoiceListResultInvoicesInner) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *GetInvoiceListResultInvoicesInner) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *GetInvoiceListResultInvoicesInner) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotal

`func (o *GetInvoiceListResultInvoicesInner) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInvoiceListResultInvoicesInner) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInvoiceListResultInvoicesInner) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetInvoiceListResultInvoicesInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


