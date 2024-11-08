# GetLoadBalancerListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | Pointer to [**[]LoadBalancerListItem**](LoadBalancerListItem.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetLoadBalancerListResult

`func NewGetLoadBalancerListResult() *GetLoadBalancerListResult`

NewGetLoadBalancerListResult instantiates a new GetLoadBalancerListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancerListResultWithDefaults

`func NewGetLoadBalancerListResultWithDefaults() *GetLoadBalancerListResult`

NewGetLoadBalancerListResultWithDefaults instantiates a new GetLoadBalancerListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *GetLoadBalancerListResult) GetLoadBalancers() []LoadBalancerListItem`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *GetLoadBalancerListResult) GetLoadBalancersOk() (*[]LoadBalancerListItem, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *GetLoadBalancerListResult) SetLoadBalancers(v []LoadBalancerListItem)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *GetLoadBalancerListResult) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetMetadata

`func (o *GetLoadBalancerListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetLoadBalancerListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetLoadBalancerListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetLoadBalancerListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


