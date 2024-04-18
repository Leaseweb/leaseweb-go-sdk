# InstallOperatingSystemOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | Url which will receive callbacks when the installation is finished or failed | [optional] 
**ControlPanelId** | Pointer to **string** | Control panel identifier | [optional] 
**Device** | Pointer to **string** | Block devices in a disk set in which the partitions will be installed. Supported values are any disk set id, SATA_SAS or NVME | [optional] 
**Hostname** | Pointer to **string** | Hostname to be used in your installation | [optional] 
**OperatingSystemId** | **string** | Operating system identifier | 
**Partitions** | Pointer to [**[]Partition**](Partition.md) | Array of partition objects that should be installed per partition | [optional] 
**Password** | Pointer to **string** | Server root password. If not provided, it would be automatically generated | [optional] 
**PostInstallScript** | Pointer to **string** | Base64 Encoded string containing a valid bash script to be run right after the installation | [optional] 
**PowerCycle** | Pointer to **bool** | If true, allows system reboots to happen automatically within the process. Otherwise, you should do them manually | [optional] 
**Raid** | Pointer to [**Raid**](Raid.md) |  | [optional] 
**SshKeys** | Pointer to **string** | List of public sshKeys to be setup in your installation, separated by new lines | [optional] 
**Timezone** | Pointer to **string** | Timezone represented as Geographical_Area/City | [optional] 

## Methods

### NewInstallOperatingSystemOpts

`func NewInstallOperatingSystemOpts(operatingSystemId string, ) *InstallOperatingSystemOpts`

NewInstallOperatingSystemOpts instantiates a new InstallOperatingSystemOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallOperatingSystemOptsWithDefaults

`func NewInstallOperatingSystemOptsWithDefaults() *InstallOperatingSystemOpts`

NewInstallOperatingSystemOptsWithDefaults instantiates a new InstallOperatingSystemOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *InstallOperatingSystemOpts) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *InstallOperatingSystemOpts) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *InstallOperatingSystemOpts) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *InstallOperatingSystemOpts) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetControlPanelId

`func (o *InstallOperatingSystemOpts) GetControlPanelId() string`

GetControlPanelId returns the ControlPanelId field if non-nil, zero value otherwise.

### GetControlPanelIdOk

`func (o *InstallOperatingSystemOpts) GetControlPanelIdOk() (*string, bool)`

GetControlPanelIdOk returns a tuple with the ControlPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanelId

`func (o *InstallOperatingSystemOpts) SetControlPanelId(v string)`

SetControlPanelId sets ControlPanelId field to given value.

### HasControlPanelId

`func (o *InstallOperatingSystemOpts) HasControlPanelId() bool`

HasControlPanelId returns a boolean if a field has been set.

### GetDevice

`func (o *InstallOperatingSystemOpts) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InstallOperatingSystemOpts) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InstallOperatingSystemOpts) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InstallOperatingSystemOpts) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHostname

`func (o *InstallOperatingSystemOpts) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstallOperatingSystemOpts) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstallOperatingSystemOpts) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstallOperatingSystemOpts) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOperatingSystemId

`func (o *InstallOperatingSystemOpts) GetOperatingSystemId() string`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *InstallOperatingSystemOpts) GetOperatingSystemIdOk() (*string, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *InstallOperatingSystemOpts) SetOperatingSystemId(v string)`

SetOperatingSystemId sets OperatingSystemId field to given value.


### GetPartitions

`func (o *InstallOperatingSystemOpts) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *InstallOperatingSystemOpts) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *InstallOperatingSystemOpts) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *InstallOperatingSystemOpts) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetPassword

`func (o *InstallOperatingSystemOpts) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InstallOperatingSystemOpts) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InstallOperatingSystemOpts) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InstallOperatingSystemOpts) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPostInstallScript

`func (o *InstallOperatingSystemOpts) GetPostInstallScript() string`

GetPostInstallScript returns the PostInstallScript field if non-nil, zero value otherwise.

### GetPostInstallScriptOk

`func (o *InstallOperatingSystemOpts) GetPostInstallScriptOk() (*string, bool)`

GetPostInstallScriptOk returns a tuple with the PostInstallScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostInstallScript

`func (o *InstallOperatingSystemOpts) SetPostInstallScript(v string)`

SetPostInstallScript sets PostInstallScript field to given value.

### HasPostInstallScript

`func (o *InstallOperatingSystemOpts) HasPostInstallScript() bool`

HasPostInstallScript returns a boolean if a field has been set.

### GetPowerCycle

`func (o *InstallOperatingSystemOpts) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *InstallOperatingSystemOpts) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *InstallOperatingSystemOpts) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *InstallOperatingSystemOpts) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetRaid

`func (o *InstallOperatingSystemOpts) GetRaid() Raid`

GetRaid returns the Raid field if non-nil, zero value otherwise.

### GetRaidOk

`func (o *InstallOperatingSystemOpts) GetRaidOk() (*Raid, bool)`

GetRaidOk returns a tuple with the Raid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid

`func (o *InstallOperatingSystemOpts) SetRaid(v Raid)`

SetRaid sets Raid field to given value.

### HasRaid

`func (o *InstallOperatingSystemOpts) HasRaid() bool`

HasRaid returns a boolean if a field has been set.

### GetSshKeys

`func (o *InstallOperatingSystemOpts) GetSshKeys() string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *InstallOperatingSystemOpts) GetSshKeysOk() (*string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *InstallOperatingSystemOpts) SetSshKeys(v string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *InstallOperatingSystemOpts) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetTimezone

`func (o *InstallOperatingSystemOpts) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InstallOperatingSystemOpts) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InstallOperatingSystemOpts) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InstallOperatingSystemOpts) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


