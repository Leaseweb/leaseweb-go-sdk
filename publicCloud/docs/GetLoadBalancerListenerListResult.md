# GetLoadBalancerListenerListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listeners** | Pointer to [**[]LoadBalancerListener**](LoadBalancerListener.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetLoadBalancerListenerListResult

`func NewGetLoadBalancerListenerListResult() *GetLoadBalancerListenerListResult`

NewGetLoadBalancerListenerListResult instantiates a new GetLoadBalancerListenerListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancerListenerListResultWithDefaults

`func NewGetLoadBalancerListenerListResultWithDefaults() *GetLoadBalancerListenerListResult`

NewGetLoadBalancerListenerListResultWithDefaults instantiates a new GetLoadBalancerListenerListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListeners

`func (o *GetLoadBalancerListenerListResult) GetListeners() []LoadBalancerListener`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *GetLoadBalancerListenerListResult) GetListenersOk() (*[]LoadBalancerListener, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *GetLoadBalancerListenerListResult) SetListeners(v []LoadBalancerListener)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *GetLoadBalancerListenerListResult) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### GetMetadata

`func (o *GetLoadBalancerListenerListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetLoadBalancerListenerListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetLoadBalancerListenerListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetLoadBalancerListenerListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


