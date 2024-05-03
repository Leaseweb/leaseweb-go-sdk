# GetNetworkEquipmentListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**NetworkEquipments** | Pointer to [**[]NetworkEquipment**](NetworkEquipment.md) | An array of network equipment | [optional] 

## Methods

### NewGetNetworkEquipmentListResult

`func NewGetNetworkEquipmentListResult() *GetNetworkEquipmentListResult`

NewGetNetworkEquipmentListResult instantiates a new GetNetworkEquipmentListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkEquipmentListResultWithDefaults

`func NewGetNetworkEquipmentListResultWithDefaults() *GetNetworkEquipmentListResult`

NewGetNetworkEquipmentListResultWithDefaults instantiates a new GetNetworkEquipmentListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetNetworkEquipmentListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNetworkEquipmentListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNetworkEquipmentListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNetworkEquipmentListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNetworkEquipments

`func (o *GetNetworkEquipmentListResult) GetNetworkEquipments() []NetworkEquipment`

GetNetworkEquipments returns the NetworkEquipments field if non-nil, zero value otherwise.

### GetNetworkEquipmentsOk

`func (o *GetNetworkEquipmentListResult) GetNetworkEquipmentsOk() (*[]NetworkEquipment, bool)`

GetNetworkEquipmentsOk returns a tuple with the NetworkEquipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipments

`func (o *GetNetworkEquipmentListResult) SetNetworkEquipments(v []NetworkEquipment)`

SetNetworkEquipments sets NetworkEquipments field to given value.

### HasNetworkEquipments

`func (o *GetNetworkEquipmentListResult) HasNetworkEquipments() bool`

HasNetworkEquipments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


