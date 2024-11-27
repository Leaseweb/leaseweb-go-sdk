# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**NetworkCapabilities**](NetworkCapabilities.md) |  | [optional] 
**Lldp** | Pointer to [**Lldp**](Lldp.md) |  | [optional] 
**LogicalName** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** | Represents a MAC Address in the standard colon delimited format. Eg. &#x60;01:23:45:67:89:0A&#x60; | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**NetworkSettings**](NetworkSettings.md) |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *Network) GetCapabilities() NetworkCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Network) GetCapabilitiesOk() (*NetworkCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Network) SetCapabilities(v NetworkCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Network) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetLldp

`func (o *Network) GetLldp() Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *Network) GetLldpOk() (*Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *Network) SetLldp(v Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *Network) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetLogicalName

`func (o *Network) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *Network) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *Network) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *Network) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### GetMacAddress

`func (o *Network) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Network) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Network) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Network) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetProduct

`func (o *Network) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Network) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Network) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Network) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSettings

`func (o *Network) GetSettings() NetworkSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Network) GetSettingsOk() (*NetworkSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Network) SetSettings(v NetworkSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Network) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetVendor

`func (o *Network) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Network) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Network) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Network) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


