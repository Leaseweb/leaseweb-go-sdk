# EnableServerRescueModeOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | Url which will receive callbacks | [optional] 
**Password** | Pointer to **string** | Rescue mode password. If not provided, it would be automatically generated | [optional] 
**PostInstallScript** | Pointer to **string** | Base64 Encoded string containing a valid bash script to be run right after rescue mode is launched | [optional] 
**PowerCycle** | Pointer to **bool** | If set to &#x60;true&#x60;, server will be power cycled in order to complete the operation | [optional] [default to true]
**RescueImageId** | **string** | Rescue image identifier | [default to "GRML"]
**SshKeys** | Pointer to **string** | User ssh keys | [optional] 

## Methods

### NewEnableServerRescueModeOpts

`func NewEnableServerRescueModeOpts(rescueImageId string, ) *EnableServerRescueModeOpts`

NewEnableServerRescueModeOpts instantiates a new EnableServerRescueModeOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableServerRescueModeOptsWithDefaults

`func NewEnableServerRescueModeOptsWithDefaults() *EnableServerRescueModeOpts`

NewEnableServerRescueModeOptsWithDefaults instantiates a new EnableServerRescueModeOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *EnableServerRescueModeOpts) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *EnableServerRescueModeOpts) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *EnableServerRescueModeOpts) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *EnableServerRescueModeOpts) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetPassword

`func (o *EnableServerRescueModeOpts) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnableServerRescueModeOpts) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnableServerRescueModeOpts) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnableServerRescueModeOpts) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPostInstallScript

`func (o *EnableServerRescueModeOpts) GetPostInstallScript() string`

GetPostInstallScript returns the PostInstallScript field if non-nil, zero value otherwise.

### GetPostInstallScriptOk

`func (o *EnableServerRescueModeOpts) GetPostInstallScriptOk() (*string, bool)`

GetPostInstallScriptOk returns a tuple with the PostInstallScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostInstallScript

`func (o *EnableServerRescueModeOpts) SetPostInstallScript(v string)`

SetPostInstallScript sets PostInstallScript field to given value.

### HasPostInstallScript

`func (o *EnableServerRescueModeOpts) HasPostInstallScript() bool`

HasPostInstallScript returns a boolean if a field has been set.

### GetPowerCycle

`func (o *EnableServerRescueModeOpts) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *EnableServerRescueModeOpts) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *EnableServerRescueModeOpts) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *EnableServerRescueModeOpts) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetRescueImageId

`func (o *EnableServerRescueModeOpts) GetRescueImageId() string`

GetRescueImageId returns the RescueImageId field if non-nil, zero value otherwise.

### GetRescueImageIdOk

`func (o *EnableServerRescueModeOpts) GetRescueImageIdOk() (*string, bool)`

GetRescueImageIdOk returns a tuple with the RescueImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueImageId

`func (o *EnableServerRescueModeOpts) SetRescueImageId(v string)`

SetRescueImageId sets RescueImageId field to given value.


### GetSshKeys

`func (o *EnableServerRescueModeOpts) GetSshKeys() string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *EnableServerRescueModeOpts) GetSshKeysOk() (*string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *EnableServerRescueModeOpts) SetSshKeys(v string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *EnableServerRescueModeOpts) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


