# GetInvoiceListResultMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The limit used to generate this response. | [optional] [default to 10]
**Offset** | Pointer to **int32** | The offset used to generate this response. | [optional] [default to 0]
**TotalCount** | Pointer to **int32** | Total amount of invoices in this collection. | [optional] 

## Methods

### NewGetInvoiceListResultMetadata

`func NewGetInvoiceListResultMetadata() *GetInvoiceListResultMetadata`

NewGetInvoiceListResultMetadata instantiates a new GetInvoiceListResultMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceListResultMetadataWithDefaults

`func NewGetInvoiceListResultMetadataWithDefaults() *GetInvoiceListResultMetadata`

NewGetInvoiceListResultMetadataWithDefaults instantiates a new GetInvoiceListResultMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetInvoiceListResultMetadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetInvoiceListResultMetadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetInvoiceListResultMetadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetInvoiceListResultMetadata) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetInvoiceListResultMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetInvoiceListResultMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetInvoiceListResultMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetInvoiceListResultMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetInvoiceListResultMetadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetInvoiceListResultMetadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetInvoiceListResultMetadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetInvoiceListResultMetadata) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


