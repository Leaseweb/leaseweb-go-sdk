# LoadBalancerTargetOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]LoadBalancerTargetOpt**](LoadBalancerTargetOpt.md) |  | [optional] 

## Methods

### NewLoadBalancerTargetOpts

`func NewLoadBalancerTargetOpts() *LoadBalancerTargetOpts`

NewLoadBalancerTargetOpts instantiates a new LoadBalancerTargetOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerTargetOptsWithDefaults

`func NewLoadBalancerTargetOptsWithDefaults() *LoadBalancerTargetOpts`

NewLoadBalancerTargetOptsWithDefaults instantiates a new LoadBalancerTargetOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *LoadBalancerTargetOpts) GetTargets() []LoadBalancerTargetOpt`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *LoadBalancerTargetOpts) GetTargetsOk() (*[]LoadBalancerTargetOpt, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *LoadBalancerTargetOpts) SetTargets(v []LoadBalancerTargetOpt)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *LoadBalancerTargetOpts) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


