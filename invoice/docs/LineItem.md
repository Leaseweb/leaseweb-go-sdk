# LineItem

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

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *LineItem) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *LineItem) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *LineItem) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *LineItem) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetEquipmentId

`func (o *LineItem) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *LineItem) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *LineItem) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *LineItem) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetProduct

`func (o *LineItem) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LineItem) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LineItem) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LineItem) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetQuantity

`func (o *LineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReference

`func (o *LineItem) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LineItem) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LineItem) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LineItem) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTotal

`func (o *LineItem) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LineItem) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LineItem) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LineItem) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnitAmount

`func (o *LineItem) GetUnitAmount() float32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *LineItem) GetUnitAmountOk() (*float32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *LineItem) SetUnitAmount(v float32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *LineItem) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetFromDate

`func (o *LineItem) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *LineItem) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *LineItem) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *LineItem) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDateNil

`func (o *LineItem) SetFromDateNil(b bool)`

 SetFromDateNil sets the value for FromDate to be an explicit nil

### UnsetFromDate
`func (o *LineItem) UnsetFromDate()`

UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
### GetToDate

`func (o *LineItem) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *LineItem) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *LineItem) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *LineItem) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDateNil

`func (o *LineItem) SetToDateNil(b bool)`

 SetToDateNil sets the value for ToDate to be an explicit nil

### UnsetToDate
`func (o *LineItem) UnsetToDate()`

UnsetToDate ensures that no value is present for ToDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


