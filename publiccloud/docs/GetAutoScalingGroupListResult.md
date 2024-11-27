# GetAutoScalingGroupListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScalingGroups** | Pointer to [**[]AutoScalingGroup**](AutoScalingGroup.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetAutoScalingGroupListResult

`func NewGetAutoScalingGroupListResult() *GetAutoScalingGroupListResult`

NewGetAutoScalingGroupListResult instantiates a new GetAutoScalingGroupListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutoScalingGroupListResultWithDefaults

`func NewGetAutoScalingGroupListResultWithDefaults() *GetAutoScalingGroupListResult`

NewGetAutoScalingGroupListResultWithDefaults instantiates a new GetAutoScalingGroupListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScalingGroups

`func (o *GetAutoScalingGroupListResult) GetAutoScalingGroups() []AutoScalingGroup`

GetAutoScalingGroups returns the AutoScalingGroups field if non-nil, zero value otherwise.

### GetAutoScalingGroupsOk

`func (o *GetAutoScalingGroupListResult) GetAutoScalingGroupsOk() (*[]AutoScalingGroup, bool)`

GetAutoScalingGroupsOk returns a tuple with the AutoScalingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroups

`func (o *GetAutoScalingGroupListResult) SetAutoScalingGroups(v []AutoScalingGroup)`

SetAutoScalingGroups sets AutoScalingGroups field to given value.

### HasAutoScalingGroups

`func (o *GetAutoScalingGroupListResult) HasAutoScalingGroups() bool`

HasAutoScalingGroups returns a boolean if a field has been set.

### GetMetadata

`func (o *GetAutoScalingGroupListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAutoScalingGroupListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAutoScalingGroupListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAutoScalingGroupListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


