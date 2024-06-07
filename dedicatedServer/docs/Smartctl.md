# Smartctl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtaVersion** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**DeviceModel** | Pointer to **string** |  | [optional] 
**ExecutionStatus** | Pointer to **string** |  | [optional] 
**FirmwareVersion** | Pointer to **string** |  | [optional] 
**IsSas** | Pointer to **bool** |  | [optional] 
**OverallHealth** | Pointer to **string** |  | [optional] 
**Rpm** | Pointer to **string** |  | [optional] 
**SataVersion** | Pointer to **string** |  | [optional] 
**SectorSize** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SmartErrorLog** | Pointer to **string** |  | [optional] 
**SmartSupport** | Pointer to [**SmartSupport**](SmartSupport.md) |  | [optional] 
**SmartctlVersion** | Pointer to **string** |  | [optional] 
**UserCapacity** | Pointer to **string** |  | [optional] 

## Methods

### NewSmartctl

`func NewSmartctl() *Smartctl`

NewSmartctl instantiates a new Smartctl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartctlWithDefaults

`func NewSmartctlWithDefaults() *Smartctl`

NewSmartctlWithDefaults instantiates a new Smartctl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtaVersion

`func (o *Smartctl) GetAtaVersion() string`

GetAtaVersion returns the AtaVersion field if non-nil, zero value otherwise.

### GetAtaVersionOk

`func (o *Smartctl) GetAtaVersionOk() (*string, bool)`

GetAtaVersionOk returns a tuple with the AtaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtaVersion

`func (o *Smartctl) SetAtaVersion(v string)`

SetAtaVersion sets AtaVersion field to given value.

### HasAtaVersion

`func (o *Smartctl) HasAtaVersion() bool`

HasAtaVersion returns a boolean if a field has been set.

### GetAttributes

`func (o *Smartctl) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Smartctl) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Smartctl) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Smartctl) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDeviceModel

`func (o *Smartctl) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *Smartctl) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *Smartctl) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *Smartctl) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetExecutionStatus

`func (o *Smartctl) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *Smartctl) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *Smartctl) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *Smartctl) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *Smartctl) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *Smartctl) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *Smartctl) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *Smartctl) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetIsSas

`func (o *Smartctl) GetIsSas() bool`

GetIsSas returns the IsSas field if non-nil, zero value otherwise.

### GetIsSasOk

`func (o *Smartctl) GetIsSasOk() (*bool, bool)`

GetIsSasOk returns a tuple with the IsSas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSas

`func (o *Smartctl) SetIsSas(v bool)`

SetIsSas sets IsSas field to given value.

### HasIsSas

`func (o *Smartctl) HasIsSas() bool`

HasIsSas returns a boolean if a field has been set.

### GetOverallHealth

`func (o *Smartctl) GetOverallHealth() string`

GetOverallHealth returns the OverallHealth field if non-nil, zero value otherwise.

### GetOverallHealthOk

`func (o *Smartctl) GetOverallHealthOk() (*string, bool)`

GetOverallHealthOk returns a tuple with the OverallHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallHealth

`func (o *Smartctl) SetOverallHealth(v string)`

SetOverallHealth sets OverallHealth field to given value.

### HasOverallHealth

`func (o *Smartctl) HasOverallHealth() bool`

HasOverallHealth returns a boolean if a field has been set.

### GetRpm

`func (o *Smartctl) GetRpm() string`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *Smartctl) GetRpmOk() (*string, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *Smartctl) SetRpm(v string)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *Smartctl) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### GetSataVersion

`func (o *Smartctl) GetSataVersion() string`

GetSataVersion returns the SataVersion field if non-nil, zero value otherwise.

### GetSataVersionOk

`func (o *Smartctl) GetSataVersionOk() (*string, bool)`

GetSataVersionOk returns a tuple with the SataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSataVersion

`func (o *Smartctl) SetSataVersion(v string)`

SetSataVersion sets SataVersion field to given value.

### HasSataVersion

`func (o *Smartctl) HasSataVersion() bool`

HasSataVersion returns a boolean if a field has been set.

### GetSectorSize

`func (o *Smartctl) GetSectorSize() string`

GetSectorSize returns the SectorSize field if non-nil, zero value otherwise.

### GetSectorSizeOk

`func (o *Smartctl) GetSectorSizeOk() (*string, bool)`

GetSectorSizeOk returns a tuple with the SectorSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSize

`func (o *Smartctl) SetSectorSize(v string)`

SetSectorSize sets SectorSize field to given value.

### HasSectorSize

`func (o *Smartctl) HasSectorSize() bool`

HasSectorSize returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Smartctl) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Smartctl) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Smartctl) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Smartctl) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSmartErrorLog

`func (o *Smartctl) GetSmartErrorLog() string`

GetSmartErrorLog returns the SmartErrorLog field if non-nil, zero value otherwise.

### GetSmartErrorLogOk

`func (o *Smartctl) GetSmartErrorLogOk() (*string, bool)`

GetSmartErrorLogOk returns a tuple with the SmartErrorLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartErrorLog

`func (o *Smartctl) SetSmartErrorLog(v string)`

SetSmartErrorLog sets SmartErrorLog field to given value.

### HasSmartErrorLog

`func (o *Smartctl) HasSmartErrorLog() bool`

HasSmartErrorLog returns a boolean if a field has been set.

### GetSmartSupport

`func (o *Smartctl) GetSmartSupport() SmartSupport`

GetSmartSupport returns the SmartSupport field if non-nil, zero value otherwise.

### GetSmartSupportOk

`func (o *Smartctl) GetSmartSupportOk() (*SmartSupport, bool)`

GetSmartSupportOk returns a tuple with the SmartSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSupport

`func (o *Smartctl) SetSmartSupport(v SmartSupport)`

SetSmartSupport sets SmartSupport field to given value.

### HasSmartSupport

`func (o *Smartctl) HasSmartSupport() bool`

HasSmartSupport returns a boolean if a field has been set.

### GetSmartctlVersion

`func (o *Smartctl) GetSmartctlVersion() string`

GetSmartctlVersion returns the SmartctlVersion field if non-nil, zero value otherwise.

### GetSmartctlVersionOk

`func (o *Smartctl) GetSmartctlVersionOk() (*string, bool)`

GetSmartctlVersionOk returns a tuple with the SmartctlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartctlVersion

`func (o *Smartctl) SetSmartctlVersion(v string)`

SetSmartctlVersion sets SmartctlVersion field to given value.

### HasSmartctlVersion

`func (o *Smartctl) HasSmartctlVersion() bool`

HasSmartctlVersion returns a boolean if a field has been set.

### GetUserCapacity

`func (o *Smartctl) GetUserCapacity() string`

GetUserCapacity returns the UserCapacity field if non-nil, zero value otherwise.

### GetUserCapacityOk

`func (o *Smartctl) GetUserCapacityOk() (*string, bool)`

GetUserCapacityOk returns a tuple with the UserCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCapacity

`func (o *Smartctl) SetUserCapacity(v string)`

SetUserCapacity sets UserCapacity field to given value.

### HasUserCapacity

`func (o *Smartctl) HasUserCapacity() bool`

HasUserCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


