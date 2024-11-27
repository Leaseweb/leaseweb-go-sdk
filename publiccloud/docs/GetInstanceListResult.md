# GetInstanceListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]Instance**](Instance.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetInstanceListResult

`func NewGetInstanceListResult() *GetInstanceListResult`

NewGetInstanceListResult instantiates a new GetInstanceListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceListResultWithDefaults

`func NewGetInstanceListResultWithDefaults() *GetInstanceListResult`

NewGetInstanceListResultWithDefaults instantiates a new GetInstanceListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *GetInstanceListResult) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetInstanceListResult) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetInstanceListResult) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetInstanceListResult) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMetadata

`func (o *GetInstanceListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetInstanceListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetInstanceListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetInstanceListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


