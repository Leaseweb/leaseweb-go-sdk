# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customer ID who owns the load balancer | [optional] 
**Id** | Pointer to **string** | The load balancer unique identifier | [optional] 
**Type** | Pointer to **string** | Load balancer type | [optional] 
**Resources** | Pointer to [**InstanceResources**](InstanceResources.md) |  | [optional] 
**Region** | Pointer to **string** | The region where the load balancer was launched into | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the load balancer | [optional] 
**State** | Pointer to [**InstanceState**](InstanceState.md) |  | [optional] 
**RootDiskSize** | Pointer to **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** | Date and time when the instance was started for the first time, right after launching it | [optional] 
**Contract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**Configuration** | Pointer to [**LoadBalancerConfiguration**](LoadBalancerConfiguration.md) |  | [optional] 
**AutoScalingGroup** | Pointer to [**NullableLoadBalancerAutoScalingGroup**](LoadBalancerAutoScalingGroup.md) |  | [optional] 

## Methods

### NewLoadBalancer

`func NewLoadBalancer() *LoadBalancer`

NewLoadBalancer instantiates a new LoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerWithDefaults

`func NewLoadBalancerWithDefaults() *LoadBalancer`

NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *LoadBalancer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *LoadBalancer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *LoadBalancer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *LoadBalancer) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

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

### HasId

`func (o *LoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LoadBalancer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResources

`func (o *LoadBalancer) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LoadBalancer) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LoadBalancer) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *LoadBalancer) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRegion

`func (o *LoadBalancer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoadBalancer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoadBalancer) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoadBalancer) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

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

### HasReference

`func (o *LoadBalancer) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetState

`func (o *LoadBalancer) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoadBalancer) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoadBalancer) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *LoadBalancer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRootDiskSize

`func (o *LoadBalancer) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *LoadBalancer) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *LoadBalancer) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *LoadBalancer) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetIps

`func (o *LoadBalancer) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *LoadBalancer) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *LoadBalancer) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *LoadBalancer) HasIps() bool`

HasIps returns a boolean if a field has been set.

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

### HasStartedAt

`func (o *LoadBalancer) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

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

### HasContract

`func (o *LoadBalancer) HasContract() bool`

HasContract returns a boolean if a field has been set.

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

### HasConfiguration

`func (o *LoadBalancer) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAutoScalingGroup

`func (o *LoadBalancer) GetAutoScalingGroup() LoadBalancerAutoScalingGroup`

GetAutoScalingGroup returns the AutoScalingGroup field if non-nil, zero value otherwise.

### GetAutoScalingGroupOk

`func (o *LoadBalancer) GetAutoScalingGroupOk() (*LoadBalancerAutoScalingGroup, bool)`

GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroup

`func (o *LoadBalancer) SetAutoScalingGroup(v LoadBalancerAutoScalingGroup)`

SetAutoScalingGroup sets AutoScalingGroup field to given value.

### HasAutoScalingGroup

`func (o *LoadBalancer) HasAutoScalingGroup() bool`

HasAutoScalingGroup returns a boolean if a field has been set.

### SetAutoScalingGroupNil

`func (o *LoadBalancer) SetAutoScalingGroupNil(b bool)`

 SetAutoScalingGroupNil sets the value for AutoScalingGroup to be an explicit nil

### UnsetAutoScalingGroup
`func (o *LoadBalancer) UnsetAutoScalingGroup()`

UnsetAutoScalingGroup ensures that no value is present for AutoScalingGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


