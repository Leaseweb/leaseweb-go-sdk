# GetInvoiceListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Invoices** | Pointer to [**[]Invoice**](Invoice.md) | An array of invoices. | [optional] 

## Methods

### NewGetInvoiceListResult

`func NewGetInvoiceListResult() *GetInvoiceListResult`

NewGetInvoiceListResult instantiates a new GetInvoiceListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceListResultWithDefaults

`func NewGetInvoiceListResultWithDefaults() *GetInvoiceListResult`

NewGetInvoiceListResultWithDefaults instantiates a new GetInvoiceListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetInvoiceListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetInvoiceListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetInvoiceListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetInvoiceListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInvoices

`func (o *GetInvoiceListResult) GetInvoices() []Invoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *GetInvoiceListResult) GetInvoicesOk() (*[]Invoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *GetInvoiceListResult) SetInvoices(v []Invoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *GetInvoiceListResult) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


