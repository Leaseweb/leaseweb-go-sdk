# GetProformaResultContractItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | Pointer to **string** | The unique id of the contract | [optional] 
**Currency** | Pointer to **string** | The currency of the pro forma contract item | [optional] 
**EndDate** | Pointer to **string** | The end date of the contract | [optional] 
**EquipmentId** | Pointer to **string** | The unique id of the equipment | [optional] 
**PoNumber** | Pointer to **string** | The purchase order number. | [optional] 
**Price** | Pointer to **float32** | The price of the contract item. | [optional] 
**Product** | Pointer to **string** | The identifier of a product. | [optional] 
**Reference** | Pointer to **string** | The reference a customer gave to the service | [optional] 
**StartDate** | Pointer to **string** | The start date of the contract | [optional] 

## Methods

### NewGetProformaResultContractItemsInner

`func NewGetProformaResultContractItemsInner() *GetProformaResultContractItemsInner`

NewGetProformaResultContractItemsInner instantiates a new GetProformaResultContractItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProformaResultContractItemsInnerWithDefaults

`func NewGetProformaResultContractItemsInnerWithDefaults() *GetProformaResultContractItemsInner`

NewGetProformaResultContractItemsInnerWithDefaults instantiates a new GetProformaResultContractItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *GetProformaResultContractItemsInner) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *GetProformaResultContractItemsInner) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *GetProformaResultContractItemsInner) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *GetProformaResultContractItemsInner) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCurrency

`func (o *GetProformaResultContractItemsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetProformaResultContractItemsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetProformaResultContractItemsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetProformaResultContractItemsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndDate

`func (o *GetProformaResultContractItemsInner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetProformaResultContractItemsInner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetProformaResultContractItemsInner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetProformaResultContractItemsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEquipmentId

`func (o *GetProformaResultContractItemsInner) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *GetProformaResultContractItemsInner) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *GetProformaResultContractItemsInner) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *GetProformaResultContractItemsInner) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetPoNumber

`func (o *GetProformaResultContractItemsInner) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *GetProformaResultContractItemsInner) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *GetProformaResultContractItemsInner) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *GetProformaResultContractItemsInner) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetPrice

`func (o *GetProformaResultContractItemsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetProformaResultContractItemsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetProformaResultContractItemsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetProformaResultContractItemsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProduct

`func (o *GetProformaResultContractItemsInner) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *GetProformaResultContractItemsInner) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *GetProformaResultContractItemsInner) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *GetProformaResultContractItemsInner) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetReference

`func (o *GetProformaResultContractItemsInner) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetProformaResultContractItemsInner) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetProformaResultContractItemsInner) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetProformaResultContractItemsInner) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStartDate

`func (o *GetProformaResultContractItemsInner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetProformaResultContractItemsInner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetProformaResultContractItemsInner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetProformaResultContractItemsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


