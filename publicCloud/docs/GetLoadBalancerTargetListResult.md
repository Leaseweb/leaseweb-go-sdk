# GetLoadBalancerTargetListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]Target**](Target.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetLoadBalancerTargetListResult

`func NewGetLoadBalancerTargetListResult() *GetLoadBalancerTargetListResult`

NewGetLoadBalancerTargetListResult instantiates a new GetLoadBalancerTargetListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancerTargetListResultWithDefaults

`func NewGetLoadBalancerTargetListResultWithDefaults() *GetLoadBalancerTargetListResult`

NewGetLoadBalancerTargetListResultWithDefaults instantiates a new GetLoadBalancerTargetListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *GetLoadBalancerTargetListResult) GetTargets() []Target`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GetLoadBalancerTargetListResult) GetTargetsOk() (*[]Target, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GetLoadBalancerTargetListResult) SetTargets(v []Target)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *GetLoadBalancerTargetListResult) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetMetadata

`func (o *GetLoadBalancerTargetListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetLoadBalancerTargetListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetLoadBalancerTargetListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetLoadBalancerTargetListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


