# IpmiResetServerOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | Url which will receive callbacks | [optional] 
**PowerCycle** | Pointer to **bool** | If set to &#x60;true&#x60;, server will be power cycled in order to complete the operation | [optional] [default to true]

## Methods

### NewIpmiResetServerOpts

`func NewIpmiResetServerOpts() *IpmiResetServerOpts`

NewIpmiResetServerOpts instantiates a new IpmiResetServerOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiResetServerOptsWithDefaults

`func NewIpmiResetServerOptsWithDefaults() *IpmiResetServerOpts`

NewIpmiResetServerOptsWithDefaults instantiates a new IpmiResetServerOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *IpmiResetServerOpts) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *IpmiResetServerOpts) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *IpmiResetServerOpts) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *IpmiResetServerOpts) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetPowerCycle

`func (o *IpmiResetServerOpts) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *IpmiResetServerOpts) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *IpmiResetServerOpts) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *IpmiResetServerOpts) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


