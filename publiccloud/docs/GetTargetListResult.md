# GetTargetListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]Target**](Target.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetTargetListResult

`func NewGetTargetListResult() *GetTargetListResult`

NewGetTargetListResult instantiates a new GetTargetListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTargetListResultWithDefaults

`func NewGetTargetListResultWithDefaults() *GetTargetListResult`

NewGetTargetListResultWithDefaults instantiates a new GetTargetListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *GetTargetListResult) GetTargets() []Target`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GetTargetListResult) GetTargetsOk() (*[]Target, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GetTargetListResult) SetTargets(v []Target)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *GetTargetListResult) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetMetadata

`func (o *GetTargetListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetTargetListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetTargetListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetTargetListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


