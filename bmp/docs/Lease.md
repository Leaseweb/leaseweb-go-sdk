# Lease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootfile** | Pointer to **string** | The PXE bootfile the server will network boot from | [optional] 
**CreatedAt** | Pointer to **string** | The time when the DHCP reservation was created | [optional] 
**Gateway** | Pointer to **string** | The gateway for this DHCP reservation | [optional] 
**Hostname** | Pointer to **string** | The hostname for the server | [optional] 
**Ip** | Pointer to **string** | The IP address this DHCP reservation is for | [optional] 
**LastClientRequest** | Pointer to [**LastClientRequest**](LastClientRequest.md) |  | [optional] 
**Mac** | Pointer to **string** | Represents a MAC Address in the standard colon delimited format. Eg. &#x60;01:23:45:67:89:0A&#x60; | [optional] 
**Netmask** | Pointer to **string** | The netmask for this DHCP reservation | [optional] 
**Site** | Pointer to **string** | The site serving this DHCP reservation | [optional] 
**UpdatedAt** | Pointer to **string** | The time when the DHCP reservation was last updated or used by a client | [optional] 

## Methods

### NewLease

`func NewLease() *Lease`

NewLease instantiates a new Lease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaseWithDefaults

`func NewLeaseWithDefaults() *Lease`

NewLeaseWithDefaults instantiates a new Lease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootfile

`func (o *Lease) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *Lease) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *Lease) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *Lease) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Lease) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Lease) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Lease) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Lease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGateway

`func (o *Lease) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Lease) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Lease) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Lease) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostname

`func (o *Lease) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Lease) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Lease) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Lease) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *Lease) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Lease) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Lease) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Lease) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastClientRequest

`func (o *Lease) GetLastClientRequest() LastClientRequest`

GetLastClientRequest returns the LastClientRequest field if non-nil, zero value otherwise.

### GetLastClientRequestOk

`func (o *Lease) GetLastClientRequestOk() (*LastClientRequest, bool)`

GetLastClientRequestOk returns a tuple with the LastClientRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastClientRequest

`func (o *Lease) SetLastClientRequest(v LastClientRequest)`

SetLastClientRequest sets LastClientRequest field to given value.

### HasLastClientRequest

`func (o *Lease) HasLastClientRequest() bool`

HasLastClientRequest returns a boolean if a field has been set.

### GetMac

`func (o *Lease) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Lease) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Lease) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Lease) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetmask

`func (o *Lease) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *Lease) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *Lease) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *Lease) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSite

`func (o *Lease) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Lease) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Lease) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *Lease) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Lease) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Lease) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Lease) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Lease) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


