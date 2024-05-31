# GetNetworkEquipmentNullRouteHistoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**NullRoutes** | Pointer to [**[]NullRoute**](NullRoute.md) | An array of network equipment null route events | [optional] 

## Methods

### NewGetNetworkEquipmentNullRouteHistoryResult

`func NewGetNetworkEquipmentNullRouteHistoryResult() *GetNetworkEquipmentNullRouteHistoryResult`

NewGetNetworkEquipmentNullRouteHistoryResult instantiates a new GetNetworkEquipmentNullRouteHistoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkEquipmentNullRouteHistoryResultWithDefaults

`func NewGetNetworkEquipmentNullRouteHistoryResultWithDefaults() *GetNetworkEquipmentNullRouteHistoryResult`

NewGetNetworkEquipmentNullRouteHistoryResultWithDefaults instantiates a new GetNetworkEquipmentNullRouteHistoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetNetworkEquipmentNullRouteHistoryResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNetworkEquipmentNullRouteHistoryResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNetworkEquipmentNullRouteHistoryResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNetworkEquipmentNullRouteHistoryResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNullRoutes

`func (o *GetNetworkEquipmentNullRouteHistoryResult) GetNullRoutes() []NullRoute`

GetNullRoutes returns the NullRoutes field if non-nil, zero value otherwise.

### GetNullRoutesOk

`func (o *GetNetworkEquipmentNullRouteHistoryResult) GetNullRoutesOk() (*[]NullRoute, bool)`

GetNullRoutesOk returns a tuple with the NullRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullRoutes

`func (o *GetNetworkEquipmentNullRouteHistoryResult) SetNullRoutes(v []NullRoute)`

SetNullRoutes sets NullRoutes field to given value.

### HasNullRoutes

`func (o *GetNetworkEquipmentNullRouteHistoryResult) HasNullRoutes() bool`

HasNullRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


