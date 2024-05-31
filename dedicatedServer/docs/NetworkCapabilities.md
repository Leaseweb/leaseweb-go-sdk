# NetworkCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autonegotiation** | Pointer to **string** |  | [optional] 
**BusMaster** | Pointer to **string** |  | [optional] 
**CapList** | Pointer to **string** |  | [optional] 
**Ethernet** | Pointer to **string** |  | [optional] 
**LinkSpeeds** | Pointer to [**LinkSpeeds**](LinkSpeeds.md) |  | [optional] 
**Msi** | Pointer to **string** |  | [optional] 
**Msix** | Pointer to **string** |  | [optional] 
**Pciexpress** | Pointer to **string** |  | [optional] 
**Physical** | Pointer to **string** |  | [optional] 
**Pm** | Pointer to **string** |  | [optional] 
**Tp** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkCapabilities

`func NewNetworkCapabilities() *NetworkCapabilities`

NewNetworkCapabilities instantiates a new NetworkCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkCapabilitiesWithDefaults

`func NewNetworkCapabilitiesWithDefaults() *NetworkCapabilities`

NewNetworkCapabilitiesWithDefaults instantiates a new NetworkCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonegotiation

`func (o *NetworkCapabilities) GetAutonegotiation() string`

GetAutonegotiation returns the Autonegotiation field if non-nil, zero value otherwise.

### GetAutonegotiationOk

`func (o *NetworkCapabilities) GetAutonegotiationOk() (*string, bool)`

GetAutonegotiationOk returns a tuple with the Autonegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonegotiation

`func (o *NetworkCapabilities) SetAutonegotiation(v string)`

SetAutonegotiation sets Autonegotiation field to given value.

### HasAutonegotiation

`func (o *NetworkCapabilities) HasAutonegotiation() bool`

HasAutonegotiation returns a boolean if a field has been set.

### GetBusMaster

`func (o *NetworkCapabilities) GetBusMaster() string`

GetBusMaster returns the BusMaster field if non-nil, zero value otherwise.

### GetBusMasterOk

`func (o *NetworkCapabilities) GetBusMasterOk() (*string, bool)`

GetBusMasterOk returns a tuple with the BusMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusMaster

`func (o *NetworkCapabilities) SetBusMaster(v string)`

SetBusMaster sets BusMaster field to given value.

### HasBusMaster

`func (o *NetworkCapabilities) HasBusMaster() bool`

HasBusMaster returns a boolean if a field has been set.

### GetCapList

`func (o *NetworkCapabilities) GetCapList() string`

GetCapList returns the CapList field if non-nil, zero value otherwise.

### GetCapListOk

`func (o *NetworkCapabilities) GetCapListOk() (*string, bool)`

GetCapListOk returns a tuple with the CapList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapList

`func (o *NetworkCapabilities) SetCapList(v string)`

SetCapList sets CapList field to given value.

### HasCapList

`func (o *NetworkCapabilities) HasCapList() bool`

HasCapList returns a boolean if a field has been set.

### GetEthernet

`func (o *NetworkCapabilities) GetEthernet() string`

GetEthernet returns the Ethernet field if non-nil, zero value otherwise.

### GetEthernetOk

`func (o *NetworkCapabilities) GetEthernetOk() (*string, bool)`

GetEthernetOk returns a tuple with the Ethernet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernet

`func (o *NetworkCapabilities) SetEthernet(v string)`

SetEthernet sets Ethernet field to given value.

### HasEthernet

`func (o *NetworkCapabilities) HasEthernet() bool`

HasEthernet returns a boolean if a field has been set.

### GetLinkSpeeds

`func (o *NetworkCapabilities) GetLinkSpeeds() LinkSpeeds`

GetLinkSpeeds returns the LinkSpeeds field if non-nil, zero value otherwise.

### GetLinkSpeedsOk

`func (o *NetworkCapabilities) GetLinkSpeedsOk() (*LinkSpeeds, bool)`

GetLinkSpeedsOk returns a tuple with the LinkSpeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeeds

`func (o *NetworkCapabilities) SetLinkSpeeds(v LinkSpeeds)`

SetLinkSpeeds sets LinkSpeeds field to given value.

### HasLinkSpeeds

`func (o *NetworkCapabilities) HasLinkSpeeds() bool`

HasLinkSpeeds returns a boolean if a field has been set.

### GetMsi

`func (o *NetworkCapabilities) GetMsi() string`

GetMsi returns the Msi field if non-nil, zero value otherwise.

### GetMsiOk

`func (o *NetworkCapabilities) GetMsiOk() (*string, bool)`

GetMsiOk returns a tuple with the Msi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsi

`func (o *NetworkCapabilities) SetMsi(v string)`

SetMsi sets Msi field to given value.

### HasMsi

`func (o *NetworkCapabilities) HasMsi() bool`

HasMsi returns a boolean if a field has been set.

### GetMsix

`func (o *NetworkCapabilities) GetMsix() string`

GetMsix returns the Msix field if non-nil, zero value otherwise.

### GetMsixOk

`func (o *NetworkCapabilities) GetMsixOk() (*string, bool)`

GetMsixOk returns a tuple with the Msix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsix

`func (o *NetworkCapabilities) SetMsix(v string)`

SetMsix sets Msix field to given value.

### HasMsix

`func (o *NetworkCapabilities) HasMsix() bool`

HasMsix returns a boolean if a field has been set.

### GetPciexpress

`func (o *NetworkCapabilities) GetPciexpress() string`

GetPciexpress returns the Pciexpress field if non-nil, zero value otherwise.

### GetPciexpressOk

`func (o *NetworkCapabilities) GetPciexpressOk() (*string, bool)`

GetPciexpressOk returns a tuple with the Pciexpress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciexpress

`func (o *NetworkCapabilities) SetPciexpress(v string)`

SetPciexpress sets Pciexpress field to given value.

### HasPciexpress

`func (o *NetworkCapabilities) HasPciexpress() bool`

HasPciexpress returns a boolean if a field has been set.

### GetPhysical

`func (o *NetworkCapabilities) GetPhysical() string`

GetPhysical returns the Physical field if non-nil, zero value otherwise.

### GetPhysicalOk

`func (o *NetworkCapabilities) GetPhysicalOk() (*string, bool)`

GetPhysicalOk returns a tuple with the Physical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysical

`func (o *NetworkCapabilities) SetPhysical(v string)`

SetPhysical sets Physical field to given value.

### HasPhysical

`func (o *NetworkCapabilities) HasPhysical() bool`

HasPhysical returns a boolean if a field has been set.

### GetPm

`func (o *NetworkCapabilities) GetPm() string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *NetworkCapabilities) GetPmOk() (*string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *NetworkCapabilities) SetPm(v string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *NetworkCapabilities) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetTp

`func (o *NetworkCapabilities) GetTp() string`

GetTp returns the Tp field if non-nil, zero value otherwise.

### GetTpOk

`func (o *NetworkCapabilities) GetTpOk() (*string, bool)`

GetTpOk returns a tuple with the Tp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTp

`func (o *NetworkCapabilities) SetTp(v string)`

SetTp sets Tp field to given value.

### HasTp

`func (o *NetworkCapabilities) HasTp() bool`

HasTp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


