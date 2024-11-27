# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The load balancer unique identifier | 
**Type** | [**TypeName**](TypeName.md) |  | 
**Resources** | [**Resources**](Resources.md) |  | 
**Reference** | **NullableString** | The identifying name set to the load balancer | 
**State** | [**State**](State.md) |  | 
**StartedAt** | **NullableTime** | Date and time when the load balancer was started for the first time, right after launching it | 
**Region** | [**RegionName**](RegionName.md) |  | 
**Configuration** | [**NullableLoadBalancerConfiguration**](LoadBalancerConfiguration.md) |  | 
**AutoScalingGroup** | [**NullableAutoScalingGroup**](AutoScalingGroup.md) |  | 
**PrivateNetwork** | [**NullablePrivateNetwork**](PrivateNetwork.md) |  | 
**Contract** | [**Contract**](Contract.md) |  | 

## Methods

### NewLoadBalancer

`func NewLoadBalancer(id string, type_ TypeName, resources Resources, reference NullableString, state State, startedAt NullableTime, region RegionName, configuration NullableLoadBalancerConfiguration, autoScalingGroup NullableAutoScalingGroup, privateNetwork NullablePrivateNetwork, contract Contract, ) *LoadBalancer`

NewLoadBalancer instantiates a new LoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerWithDefaults

`func NewLoadBalancerWithDefaults() *LoadBalancer`

NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancer) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *LoadBalancer) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancer) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancer) SetType(v TypeName)`

SetType sets Type field to given value.


### GetResources

`func (o *LoadBalancer) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LoadBalancer) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LoadBalancer) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetReference

`func (o *LoadBalancer) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LoadBalancer) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LoadBalancer) SetReference(v string)`

SetReference sets Reference field to given value.


### SetReferenceNil

`func (o *LoadBalancer) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *LoadBalancer) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetState

`func (o *LoadBalancer) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoadBalancer) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoadBalancer) SetState(v State)`

SetState sets State field to given value.


### GetStartedAt

`func (o *LoadBalancer) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *LoadBalancer) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *LoadBalancer) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *LoadBalancer) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *LoadBalancer) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetRegion

`func (o *LoadBalancer) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoadBalancer) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoadBalancer) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetConfiguration

`func (o *LoadBalancer) GetConfiguration() LoadBalancerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *LoadBalancer) GetConfigurationOk() (*LoadBalancerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *LoadBalancer) SetConfiguration(v LoadBalancerConfiguration)`

SetConfiguration sets Configuration field to given value.


### SetConfigurationNil

`func (o *LoadBalancer) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *LoadBalancer) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetAutoScalingGroup

`func (o *LoadBalancer) GetAutoScalingGroup() AutoScalingGroup`

GetAutoScalingGroup returns the AutoScalingGroup field if non-nil, zero value otherwise.

### GetAutoScalingGroupOk

`func (o *LoadBalancer) GetAutoScalingGroupOk() (*AutoScalingGroup, bool)`

GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroup

`func (o *LoadBalancer) SetAutoScalingGroup(v AutoScalingGroup)`

SetAutoScalingGroup sets AutoScalingGroup field to given value.


### SetAutoScalingGroupNil

`func (o *LoadBalancer) SetAutoScalingGroupNil(b bool)`

 SetAutoScalingGroupNil sets the value for AutoScalingGroup to be an explicit nil

### UnsetAutoScalingGroup
`func (o *LoadBalancer) UnsetAutoScalingGroup()`

UnsetAutoScalingGroup ensures that no value is present for AutoScalingGroup, not even an explicit nil
### GetPrivateNetwork

`func (o *LoadBalancer) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *LoadBalancer) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *LoadBalancer) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.


### SetPrivateNetworkNil

`func (o *LoadBalancer) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *LoadBalancer) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil
### GetContract

`func (o *LoadBalancer) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *LoadBalancer) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *LoadBalancer) SetContract(v Contract)`

SetContract sets Contract field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


