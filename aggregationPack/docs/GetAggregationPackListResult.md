# GetAggregationPackListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**ConnectContractItems** | Pointer to [**[]AggregationPack**](AggregationPack.md) | An array of aggregation packs | [optional] 

## Methods

### NewGetAggregationPackListResult

`func NewGetAggregationPackListResult() *GetAggregationPackListResult`

NewGetAggregationPackListResult instantiates a new GetAggregationPackListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAggregationPackListResultWithDefaults

`func NewGetAggregationPackListResultWithDefaults() *GetAggregationPackListResult`

NewGetAggregationPackListResultWithDefaults instantiates a new GetAggregationPackListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetAggregationPackListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAggregationPackListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAggregationPackListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAggregationPackListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConnectContractItems

`func (o *GetAggregationPackListResult) GetConnectContractItems() []AggregationPack`

GetConnectContractItems returns the ConnectContractItems field if non-nil, zero value otherwise.

### GetConnectContractItemsOk

`func (o *GetAggregationPackListResult) GetConnectContractItemsOk() (*[]AggregationPack, bool)`

GetConnectContractItemsOk returns a tuple with the ConnectContractItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectContractItems

`func (o *GetAggregationPackListResult) SetConnectContractItems(v []AggregationPack)`

SetConnectContractItems sets ConnectContractItems field to given value.

### HasConnectContractItems

`func (o *GetAggregationPackListResult) HasConnectContractItems() bool`

HasConnectContractItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


