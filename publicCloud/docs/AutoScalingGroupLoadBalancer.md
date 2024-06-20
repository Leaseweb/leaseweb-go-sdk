# AutoScalingGroupLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customers ID who owns the load balancer | [optional] 
**Id** | Pointer to **string** | The load balancer unique identifier | [optional] 
**Type** | Pointer to **string** | Load balancer type | [optional] 
**Resources** | Pointer to [**InstanceResources**](InstanceResources.md) |  | [optional] 
**Region** | Pointer to **string** | The region where the load balancer was launched into | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the load balancer | [optional] 
**State** | Pointer to [**InstanceState**](InstanceState.md) |  | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**Contract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** | Date and time when the instance was started for the first time, right after launching it | [optional] 
**Configuration** | Pointer to [**NullableLoadBalancerConfiguration**](LoadBalancerConfiguration.md) |  | [optional] 
**PrivateNetwork** | Pointer to [**NullablePrivateNetwork**](PrivateNetwork.md) |  | [optional] 

## Methods

### NewAutoScalingGroupLoadBalancer

`func NewAutoScalingGroupLoadBalancer() *AutoScalingGroupLoadBalancer`

NewAutoScalingGroupLoadBalancer instantiates a new AutoScalingGroupLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupLoadBalancerWithDefaults

`func NewAutoScalingGroupLoadBalancerWithDefaults() *AutoScalingGroupLoadBalancer`

NewAutoScalingGroupLoadBalancerWithDefaults instantiates a new AutoScalingGroupLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *AutoScalingGroupLoadBalancer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AutoScalingGroupLoadBalancer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AutoScalingGroupLoadBalancer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AutoScalingGroupLoadBalancer) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *AutoScalingGroupLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupLoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoScalingGroupLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AutoScalingGroupLoadBalancer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoScalingGroupLoadBalancer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoScalingGroupLoadBalancer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoScalingGroupLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResources

`func (o *AutoScalingGroupLoadBalancer) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AutoScalingGroupLoadBalancer) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AutoScalingGroupLoadBalancer) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AutoScalingGroupLoadBalancer) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRegion

`func (o *AutoScalingGroupLoadBalancer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AutoScalingGroupLoadBalancer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AutoScalingGroupLoadBalancer) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AutoScalingGroupLoadBalancer) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *AutoScalingGroupLoadBalancer) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AutoScalingGroupLoadBalancer) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AutoScalingGroupLoadBalancer) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AutoScalingGroupLoadBalancer) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetState

`func (o *AutoScalingGroupLoadBalancer) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupLoadBalancer) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupLoadBalancer) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *AutoScalingGroupLoadBalancer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIps

`func (o *AutoScalingGroupLoadBalancer) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *AutoScalingGroupLoadBalancer) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *AutoScalingGroupLoadBalancer) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *AutoScalingGroupLoadBalancer) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetContract

`func (o *AutoScalingGroupLoadBalancer) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AutoScalingGroupLoadBalancer) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AutoScalingGroupLoadBalancer) SetContract(v Contract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *AutoScalingGroupLoadBalancer) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetStartedAt

`func (o *AutoScalingGroupLoadBalancer) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AutoScalingGroupLoadBalancer) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AutoScalingGroupLoadBalancer) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AutoScalingGroupLoadBalancer) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetConfiguration

`func (o *AutoScalingGroupLoadBalancer) GetConfiguration() LoadBalancerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AutoScalingGroupLoadBalancer) GetConfigurationOk() (*LoadBalancerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AutoScalingGroupLoadBalancer) SetConfiguration(v LoadBalancerConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AutoScalingGroupLoadBalancer) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *AutoScalingGroupLoadBalancer) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *AutoScalingGroupLoadBalancer) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetPrivateNetwork

`func (o *AutoScalingGroupLoadBalancer) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *AutoScalingGroupLoadBalancer) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *AutoScalingGroupLoadBalancer) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *AutoScalingGroupLoadBalancer) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### SetPrivateNetworkNil

`func (o *AutoScalingGroupLoadBalancer) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *AutoScalingGroupLoadBalancer) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


