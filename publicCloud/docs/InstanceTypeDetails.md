# InstanceTypeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Instance type&#39;s name | [optional] 
**Resources** | Pointer to [**InstanceResources**](InstanceResources.md) |  | [optional] 
**StorageTypes** | Pointer to **[]string** | The supported storage types for the instance type | [optional] 
**Prices** | Pointer to [**Price**](Price.md) |  | [optional] 

## Methods

### NewInstanceTypeDetails

`func NewInstanceTypeDetails() *InstanceTypeDetails`

NewInstanceTypeDetails instantiates a new InstanceTypeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeDetailsWithDefaults

`func NewInstanceTypeDetailsWithDefaults() *InstanceTypeDetails`

NewInstanceTypeDetailsWithDefaults instantiates a new InstanceTypeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceTypeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *InstanceTypeDetails) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *InstanceTypeDetails) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *InstanceTypeDetails) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *InstanceTypeDetails) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetStorageTypes

`func (o *InstanceTypeDetails) GetStorageTypes() []string`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *InstanceTypeDetails) GetStorageTypesOk() (*[]string, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *InstanceTypeDetails) SetStorageTypes(v []string)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *InstanceTypeDetails) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### SetStorageTypesNil

`func (o *InstanceTypeDetails) SetStorageTypesNil(b bool)`

 SetStorageTypesNil sets the value for StorageTypes to be an explicit nil

### UnsetStorageTypes
`func (o *InstanceTypeDetails) UnsetStorageTypes()`

UnsetStorageTypes ensures that no value is present for StorageTypes, not even an explicit nil
### GetPrices

`func (o *InstanceTypeDetails) GetPrices() Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *InstanceTypeDetails) GetPricesOk() (*Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *InstanceTypeDetails) SetPrices(v Price)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *InstanceTypeDetails) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


