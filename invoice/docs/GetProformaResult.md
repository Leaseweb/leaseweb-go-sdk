# GetProformaResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**ContractItems** | Pointer to [**[]ContractItem**](ContractItem.md) |  | [optional] 
**Currency** | Pointer to **string** | The currency of the invoice. Based on ISO 4217 | [optional] 
**ProformaDate** | Pointer to **string** | The date of the next invoice on which this proforma is based on. | [optional] 
**SubTotal** | Pointer to **float32** | Total amount without vat which will be invoiced the upcoming month. | [optional] 
**Total** | Pointer to **float32** | Total amount which will be invoiced the upcoming month. | [optional] 
**VatAmount** | Pointer to **float32** | The total amount of vat. | [optional] 

## Methods

### NewGetProformaResult

`func NewGetProformaResult() *GetProformaResult`

NewGetProformaResult instantiates a new GetProformaResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProformaResultWithDefaults

`func NewGetProformaResultWithDefaults() *GetProformaResult`

NewGetProformaResultWithDefaults instantiates a new GetProformaResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetProformaResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetProformaResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetProformaResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetProformaResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetContractItems

`func (o *GetProformaResult) GetContractItems() []ContractItem`

GetContractItems returns the ContractItems field if non-nil, zero value otherwise.

### GetContractItemsOk

`func (o *GetProformaResult) GetContractItemsOk() (*[]ContractItem, bool)`

GetContractItemsOk returns a tuple with the ContractItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractItems

`func (o *GetProformaResult) SetContractItems(v []ContractItem)`

SetContractItems sets ContractItems field to given value.

### HasContractItems

`func (o *GetProformaResult) HasContractItems() bool`

HasContractItems returns a boolean if a field has been set.

### GetCurrency

`func (o *GetProformaResult) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetProformaResult) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetProformaResult) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetProformaResult) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetProformaDate

`func (o *GetProformaResult) GetProformaDate() string`

GetProformaDate returns the ProformaDate field if non-nil, zero value otherwise.

### GetProformaDateOk

`func (o *GetProformaResult) GetProformaDateOk() (*string, bool)`

GetProformaDateOk returns a tuple with the ProformaDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProformaDate

`func (o *GetProformaResult) SetProformaDate(v string)`

SetProformaDate sets ProformaDate field to given value.

### HasProformaDate

`func (o *GetProformaResult) HasProformaDate() bool`

HasProformaDate returns a boolean if a field has been set.

### GetSubTotal

`func (o *GetProformaResult) GetSubTotal() float32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *GetProformaResult) GetSubTotalOk() (*float32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *GetProformaResult) SetSubTotal(v float32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *GetProformaResult) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetTotal

`func (o *GetProformaResult) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetProformaResult) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetProformaResult) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetProformaResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetVatAmount

`func (o *GetProformaResult) GetVatAmount() float32`

GetVatAmount returns the VatAmount field if non-nil, zero value otherwise.

### GetVatAmountOk

`func (o *GetProformaResult) GetVatAmountOk() (*float32, bool)`

GetVatAmountOk returns a tuple with the VatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAmount

`func (o *GetProformaResult) SetVatAmount(v float32)`

SetVatAmount sets VatAmount field to given value.

### HasVatAmount

`func (o *GetProformaResult) HasVatAmount() bool`

HasVatAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


