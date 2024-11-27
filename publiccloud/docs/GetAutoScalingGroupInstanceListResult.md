# GetAutoScalingGroupInstanceListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]AutoScaledInstance**](AutoScaledInstance.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetAutoScalingGroupInstanceListResult

`func NewGetAutoScalingGroupInstanceListResult() *GetAutoScalingGroupInstanceListResult`

NewGetAutoScalingGroupInstanceListResult instantiates a new GetAutoScalingGroupInstanceListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutoScalingGroupInstanceListResultWithDefaults

`func NewGetAutoScalingGroupInstanceListResultWithDefaults() *GetAutoScalingGroupInstanceListResult`

NewGetAutoScalingGroupInstanceListResultWithDefaults instantiates a new GetAutoScalingGroupInstanceListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *GetAutoScalingGroupInstanceListResult) GetInstances() []AutoScaledInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetAutoScalingGroupInstanceListResult) GetInstancesOk() (*[]AutoScaledInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetAutoScalingGroupInstanceListResult) SetInstances(v []AutoScaledInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetAutoScalingGroupInstanceListResult) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMetadata

`func (o *GetAutoScalingGroupInstanceListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAutoScalingGroupInstanceListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAutoScalingGroupInstanceListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAutoScalingGroupInstanceListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


