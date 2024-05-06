# Chassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Firmware** | Pointer to [**Firmware**](Firmware.md) |  | [optional] 
**Motherboard** | Pointer to [**Motherboard**](Motherboard.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewChassis

`func NewChassis() *Chassis`

NewChassis instantiates a new Chassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisWithDefaults

`func NewChassisWithDefaults() *Chassis`

NewChassisWithDefaults instantiates a new Chassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Chassis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Chassis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Chassis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Chassis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirmware

`func (o *Chassis) GetFirmware() Firmware`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *Chassis) GetFirmwareOk() (*Firmware, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *Chassis) SetFirmware(v Firmware)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *Chassis) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetMotherboard

`func (o *Chassis) GetMotherboard() Motherboard`

GetMotherboard returns the Motherboard field if non-nil, zero value otherwise.

### GetMotherboardOk

`func (o *Chassis) GetMotherboardOk() (*Motherboard, bool)`

GetMotherboardOk returns a tuple with the Motherboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotherboard

`func (o *Chassis) SetMotherboard(v Motherboard)`

SetMotherboard sets Motherboard field to given value.

### HasMotherboard

`func (o *Chassis) HasMotherboard() bool`

HasMotherboard returns a boolean if a field has been set.

### GetProduct

`func (o *Chassis) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Chassis) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Chassis) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Chassis) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSerial

`func (o *Chassis) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Chassis) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Chassis) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Chassis) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *Chassis) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Chassis) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Chassis) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Chassis) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


