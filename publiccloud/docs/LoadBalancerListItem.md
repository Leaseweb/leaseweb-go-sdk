# LoadBalancerListItem

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
**Ips** | [**[]Ip**](Ip.md) |  | 

## Methods

### NewLoadBalancerListItem

`func NewLoadBalancerListItem(id string, type_ TypeName, resources Resources, reference NullableString, state State, startedAt NullableTime, region RegionName, configuration NullableLoadBalancerConfiguration, autoScalingGroup NullableAutoScalingGroup, privateNetwork NullablePrivateNetwork, contract Contract, ips []Ip, ) *LoadBalancerListItem`

NewLoadBalancerListItem instantiates a new LoadBalancerListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListItemWithDefaults

`func NewLoadBalancerListItemWithDefaults() *LoadBalancerListItem`

NewLoadBalancerListItemWithDefaults instantiates a new LoadBalancerListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerListItem) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *LoadBalancerListItem) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancerListItem) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancerListItem) SetType(v TypeName)`

SetType sets Type field to given value.


### GetResources

`func (o *LoadBalancerListItem) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LoadBalancerListItem) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LoadBalancerListItem) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetReference

`func (o *LoadBalancerListItem) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LoadBalancerListItem) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LoadBalancerListItem) SetReference(v string)`

SetReference sets Reference field to given value.


### SetReferenceNil

`func (o *LoadBalancerListItem) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *LoadBalancerListItem) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetState

`func (o *LoadBalancerListItem) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoadBalancerListItem) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoadBalancerListItem) SetState(v State)`

SetState sets State field to given value.


### GetStartedAt

`func (o *LoadBalancerListItem) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *LoadBalancerListItem) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *LoadBalancerListItem) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *LoadBalancerListItem) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *LoadBalancerListItem) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetRegion

`func (o *LoadBalancerListItem) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoadBalancerListItem) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoadBalancerListItem) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetConfiguration

`func (o *LoadBalancerListItem) GetConfiguration() LoadBalancerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *LoadBalancerListItem) GetConfigurationOk() (*LoadBalancerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *LoadBalancerListItem) SetConfiguration(v LoadBalancerConfiguration)`

SetConfiguration sets Configuration field to given value.


### SetConfigurationNil

`func (o *LoadBalancerListItem) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *LoadBalancerListItem) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetAutoScalingGroup

`func (o *LoadBalancerListItem) GetAutoScalingGroup() AutoScalingGroup`

GetAutoScalingGroup returns the AutoScalingGroup field if non-nil, zero value otherwise.

### GetAutoScalingGroupOk

`func (o *LoadBalancerListItem) GetAutoScalingGroupOk() (*AutoScalingGroup, bool)`

GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroup

`func (o *LoadBalancerListItem) SetAutoScalingGroup(v AutoScalingGroup)`

SetAutoScalingGroup sets AutoScalingGroup field to given value.


### SetAutoScalingGroupNil

`func (o *LoadBalancerListItem) SetAutoScalingGroupNil(b bool)`

 SetAutoScalingGroupNil sets the value for AutoScalingGroup to be an explicit nil

### UnsetAutoScalingGroup
`func (o *LoadBalancerListItem) UnsetAutoScalingGroup()`

UnsetAutoScalingGroup ensures that no value is present for AutoScalingGroup, not even an explicit nil
### GetPrivateNetwork

`func (o *LoadBalancerListItem) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *LoadBalancerListItem) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *LoadBalancerListItem) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.


### SetPrivateNetworkNil

`func (o *LoadBalancerListItem) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *LoadBalancerListItem) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil
### GetContract

`func (o *LoadBalancerListItem) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *LoadBalancerListItem) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *LoadBalancerListItem) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetIps

`func (o *LoadBalancerListItem) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *LoadBalancerListItem) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *LoadBalancerListItem) SetIps(v []Ip)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


