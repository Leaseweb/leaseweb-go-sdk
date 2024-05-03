# CpuCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpufreq** | Pointer to **string** |  | [optional] 
**Ht** | Pointer to **string** |  | [optional] 
**Vmx** | Pointer to **bool** |  | [optional] 
**X8664** | Pointer to **string** |  | [optional] 

## Methods

### NewCpuCapabilities

`func NewCpuCapabilities() *CpuCapabilities`

NewCpuCapabilities instantiates a new CpuCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuCapabilitiesWithDefaults

`func NewCpuCapabilitiesWithDefaults() *CpuCapabilities`

NewCpuCapabilitiesWithDefaults instantiates a new CpuCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpufreq

`func (o *CpuCapabilities) GetCpufreq() string`

GetCpufreq returns the Cpufreq field if non-nil, zero value otherwise.

### GetCpufreqOk

`func (o *CpuCapabilities) GetCpufreqOk() (*string, bool)`

GetCpufreqOk returns a tuple with the Cpufreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpufreq

`func (o *CpuCapabilities) SetCpufreq(v string)`

SetCpufreq sets Cpufreq field to given value.

### HasCpufreq

`func (o *CpuCapabilities) HasCpufreq() bool`

HasCpufreq returns a boolean if a field has been set.

### GetHt

`func (o *CpuCapabilities) GetHt() string`

GetHt returns the Ht field if non-nil, zero value otherwise.

### GetHtOk

`func (o *CpuCapabilities) GetHtOk() (*string, bool)`

GetHtOk returns a tuple with the Ht field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHt

`func (o *CpuCapabilities) SetHt(v string)`

SetHt sets Ht field to given value.

### HasHt

`func (o *CpuCapabilities) HasHt() bool`

HasHt returns a boolean if a field has been set.

### GetVmx

`func (o *CpuCapabilities) GetVmx() bool`

GetVmx returns the Vmx field if non-nil, zero value otherwise.

### GetVmxOk

`func (o *CpuCapabilities) GetVmxOk() (*bool, bool)`

GetVmxOk returns a tuple with the Vmx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmx

`func (o *CpuCapabilities) SetVmx(v bool)`

SetVmx sets Vmx field to given value.

### HasVmx

`func (o *CpuCapabilities) HasVmx() bool`

HasVmx returns a boolean if a field has been set.

### GetX8664

`func (o *CpuCapabilities) GetX8664() string`

GetX8664 returns the X8664 field if non-nil, zero value otherwise.

### GetX8664Ok

`func (o *CpuCapabilities) GetX8664Ok() (*string, bool)`

GetX8664Ok returns a tuple with the X8664 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX8664

`func (o *CpuCapabilities) SetX8664(v string)`

SetX8664 sets X8664 field to given value.

### HasX8664

`func (o *CpuCapabilities) HasX8664() bool`

HasX8664 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


