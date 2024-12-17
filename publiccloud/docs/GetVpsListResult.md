# GetVpsListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vps** | Pointer to [**[]VpsList**](VpsList.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetVpsListResult

`func NewGetVpsListResult() *GetVpsListResult`

NewGetVpsListResult instantiates a new GetVpsListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVpsListResultWithDefaults

`func NewGetVpsListResultWithDefaults() *GetVpsListResult`

NewGetVpsListResultWithDefaults instantiates a new GetVpsListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVps

`func (o *GetVpsListResult) GetVps() []VpsList`

GetVps returns the Vps field if non-nil, zero value otherwise.

### GetVpsOk

`func (o *GetVpsListResult) GetVpsOk() (*[]VpsList, bool)`

GetVpsOk returns a tuple with the Vps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVps

`func (o *GetVpsListResult) SetVps(v []VpsList)`

SetVps sets Vps field to given value.

### HasVps

`func (o *GetVpsListResult) HasVps() bool`

HasVps returns a boolean if a field has been set.

### GetMetadata

`func (o *GetVpsListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetVpsListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetVpsListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetVpsListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


