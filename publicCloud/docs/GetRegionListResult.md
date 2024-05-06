# GetRegionListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to [**[]Region**](Region.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetRegionListResult

`func NewGetRegionListResult() *GetRegionListResult`

NewGetRegionListResult instantiates a new GetRegionListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegionListResultWithDefaults

`func NewGetRegionListResultWithDefaults() *GetRegionListResult`

NewGetRegionListResultWithDefaults instantiates a new GetRegionListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *GetRegionListResult) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetRegionListResult) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetRegionListResult) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GetRegionListResult) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetMetadata

`func (o *GetRegionListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetRegionListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetRegionListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetRegionListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


