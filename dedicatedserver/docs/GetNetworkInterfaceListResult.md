# GetNetworkInterfaceListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]OperationNetworkInterface**](OperationNetworkInterface.md) | An array of network interfaces | [optional] 

## Methods

### NewGetNetworkInterfaceListResult

`func NewGetNetworkInterfaceListResult() *GetNetworkInterfaceListResult`

NewGetNetworkInterfaceListResult instantiates a new GetNetworkInterfaceListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkInterfaceListResultWithDefaults

`func NewGetNetworkInterfaceListResultWithDefaults() *GetNetworkInterfaceListResult`

NewGetNetworkInterfaceListResultWithDefaults instantiates a new GetNetworkInterfaceListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetNetworkInterfaceListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNetworkInterfaceListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNetworkInterfaceListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNetworkInterfaceListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *GetNetworkInterfaceListResult) GetNetworkInterfaces() []OperationNetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *GetNetworkInterfaceListResult) GetNetworkInterfacesOk() (*[]OperationNetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *GetNetworkInterfaceListResult) SetNetworkInterfaces(v []OperationNetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *GetNetworkInterfaceListResult) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


