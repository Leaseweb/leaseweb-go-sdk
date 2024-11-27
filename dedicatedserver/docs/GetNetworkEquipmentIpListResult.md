# GetNetworkEquipmentIpListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) | An array of IP addresses | [optional] 

## Methods

### NewGetNetworkEquipmentIpListResult

`func NewGetNetworkEquipmentIpListResult() *GetNetworkEquipmentIpListResult`

NewGetNetworkEquipmentIpListResult instantiates a new GetNetworkEquipmentIpListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkEquipmentIpListResultWithDefaults

`func NewGetNetworkEquipmentIpListResultWithDefaults() *GetNetworkEquipmentIpListResult`

NewGetNetworkEquipmentIpListResultWithDefaults instantiates a new GetNetworkEquipmentIpListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetNetworkEquipmentIpListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNetworkEquipmentIpListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNetworkEquipmentIpListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNetworkEquipmentIpListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIps

`func (o *GetNetworkEquipmentIpListResult) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *GetNetworkEquipmentIpListResult) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *GetNetworkEquipmentIpListResult) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *GetNetworkEquipmentIpListResult) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


