# GetTargetGroupListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetGroups** | Pointer to [**[]TargetGroup**](TargetGroup.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetTargetGroupListResult

`func NewGetTargetGroupListResult() *GetTargetGroupListResult`

NewGetTargetGroupListResult instantiates a new GetTargetGroupListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTargetGroupListResultWithDefaults

`func NewGetTargetGroupListResultWithDefaults() *GetTargetGroupListResult`

NewGetTargetGroupListResultWithDefaults instantiates a new GetTargetGroupListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGroups

`func (o *GetTargetGroupListResult) GetTargetGroups() []TargetGroup`

GetTargetGroups returns the TargetGroups field if non-nil, zero value otherwise.

### GetTargetGroupsOk

`func (o *GetTargetGroupListResult) GetTargetGroupsOk() (*[]TargetGroup, bool)`

GetTargetGroupsOk returns a tuple with the TargetGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroups

`func (o *GetTargetGroupListResult) SetTargetGroups(v []TargetGroup)`

SetTargetGroups sets TargetGroups field to given value.

### HasTargetGroups

`func (o *GetTargetGroupListResult) HasTargetGroups() bool`

HasTargetGroups returns a boolean if a field has been set.

### GetMetadata

`func (o *GetTargetGroupListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetTargetGroupListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetTargetGroupListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetTargetGroupListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


