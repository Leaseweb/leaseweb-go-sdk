# GetInvoiceResultLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | Pointer to **string** | The id of the contract | [optional] 
**EquipmentId** | Pointer to **string** | The id of the equipment | [optional] 
**Product** | Pointer to **string** | A string to indicate what kind of product this is. | [optional] 
**Quantity** | Pointer to **int32** | The quantity of a product | [optional] 
**Reference** | Pointer to **string** | The reference a customer gave to the service | [optional] 
**Total** | Pointer to **float32** | The total amount of the product | [optional] 
**UnitAmount** | Pointer to **float32** | The amount it cost for a single product unit. | [optional] 
**FromDate** | Pointer to **NullableTime** | The product start date (UTC) | [optional] 
**ToDate** | Pointer to **NullableTime** | The product end date (UTC) | [optional] 

## Methods

### NewGetInvoiceResultLineItemsInner

`func NewGetInvoiceResultLineItemsInner() *GetInvoiceResultLineItemsInner`

NewGetInvoiceResultLineItemsInner instantiates a new GetInvoiceResultLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceResultLineItemsInnerWithDefaults

`func NewGetInvoiceResultLineItemsInnerWithDefaults() *GetInvoiceResultLineItemsInner`

NewGetInvoiceResultLineItemsInnerWithDefaults instantiates a new GetInvoiceResultLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *GetInvoiceResultLineItemsInner) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *GetInvoiceResultLineItemsInner) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *GetInvoiceResultLineItemsInner) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *GetInvoiceResultLineItemsInner) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetEquipmentId

`func (o *GetInvoiceResultLineItemsInner) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *GetInvoiceResultLineItemsInner) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *GetInvoiceResultLineItemsInner) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *GetInvoiceResultLineItemsInner) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetProduct

`func (o *GetInvoiceResultLineItemsInner) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *GetInvoiceResultLineItemsInner) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *GetInvoiceResultLineItemsInner) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *GetInvoiceResultLineItemsInner) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetQuantity

`func (o *GetInvoiceResultLineItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetInvoiceResultLineItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetInvoiceResultLineItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GetInvoiceResultLineItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReference

`func (o *GetInvoiceResultLineItemsInner) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetInvoiceResultLineItemsInner) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetInvoiceResultLineItemsInner) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetInvoiceResultLineItemsInner) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTotal

`func (o *GetInvoiceResultLineItemsInner) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInvoiceResultLineItemsInner) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInvoiceResultLineItemsInner) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetInvoiceResultLineItemsInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnitAmount

`func (o *GetInvoiceResultLineItemsInner) GetUnitAmount() float32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *GetInvoiceResultLineItemsInner) GetUnitAmountOk() (*float32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *GetInvoiceResultLineItemsInner) SetUnitAmount(v float32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *GetInvoiceResultLineItemsInner) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetFromDate

`func (o *GetInvoiceResultLineItemsInner) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *GetInvoiceResultLineItemsInner) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *GetInvoiceResultLineItemsInner) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *GetInvoiceResultLineItemsInner) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDateNil

`func (o *GetInvoiceResultLineItemsInner) SetFromDateNil(b bool)`

 SetFromDateNil sets the value for FromDate to be an explicit nil

### UnsetFromDate
`func (o *GetInvoiceResultLineItemsInner) UnsetFromDate()`

UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
### GetToDate

`func (o *GetInvoiceResultLineItemsInner) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *GetInvoiceResultLineItemsInner) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *GetInvoiceResultLineItemsInner) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *GetInvoiceResultLineItemsInner) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDateNil

`func (o *GetInvoiceResultLineItemsInner) SetToDateNil(b bool)`

 SetToDateNil sets the value for ToDate to be an explicit nil

### UnsetToDate
`func (o *GetInvoiceResultLineItemsInner) UnsetToDate()`

UnsetToDate ensures that no value is present for ToDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


