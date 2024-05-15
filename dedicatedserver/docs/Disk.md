# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Smartctl** | Pointer to [**Smartctl**](Smartctl.md) |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewDisk

`func NewDisk() *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Disk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Disk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Disk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Disk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Disk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Disk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProduct

`func (o *Disk) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Disk) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Disk) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Disk) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Disk) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Disk) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Disk) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Disk) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSize

`func (o *Disk) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Disk) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Disk) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Disk) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSmartctl

`func (o *Disk) GetSmartctl() Smartctl`

GetSmartctl returns the Smartctl field if non-nil, zero value otherwise.

### GetSmartctlOk

`func (o *Disk) GetSmartctlOk() (*Smartctl, bool)`

GetSmartctlOk returns a tuple with the Smartctl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartctl

`func (o *Disk) SetSmartctl(v Smartctl)`

SetSmartctl sets Smartctl field to given value.

### HasSmartctl

`func (o *Disk) HasSmartctl() bool`

HasSmartctl returns a boolean if a field has been set.

### GetVendor

`func (o *Disk) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Disk) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Disk) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Disk) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


