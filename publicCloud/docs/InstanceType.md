# InstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**TypeName**](TypeName.md) |  | 
**Resources** | [**Resources**](Resources.md) |  | 
**StorageTypes** | [**[]StorageType**](StorageType.md) | The supported storage types for the instance type | 
**Prices** | [**Prices**](Prices.md) |  | 

## Methods

### NewInstanceType

`func NewInstanceType(name TypeName, resources Resources, storageTypes []StorageType, prices Prices, ) *InstanceType`

NewInstanceType instantiates a new InstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeWithDefaults

`func NewInstanceTypeWithDefaults() *InstanceType`

NewInstanceTypeWithDefaults instantiates a new InstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceType) GetName() TypeName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceType) GetNameOk() (*TypeName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceType) SetName(v TypeName)`

SetName sets Name field to given value.


### GetResources

`func (o *InstanceType) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *InstanceType) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *InstanceType) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetStorageTypes

`func (o *InstanceType) GetStorageTypes() []StorageType`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *InstanceType) GetStorageTypesOk() (*[]StorageType, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *InstanceType) SetStorageTypes(v []StorageType)`

SetStorageTypes sets StorageTypes field to given value.


### SetStorageTypesNil

`func (o *InstanceType) SetStorageTypesNil(b bool)`

 SetStorageTypesNil sets the value for StorageTypes to be an explicit nil

### UnsetStorageTypes
`func (o *InstanceType) UnsetStorageTypes()`

UnsetStorageTypes ensures that no value is present for StorageTypes, not even an explicit nil
### GetPrices

`func (o *InstanceType) GetPrices() Prices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *InstanceType) GetPricesOk() (*Prices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *InstanceType) SetPrices(v Prices)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


