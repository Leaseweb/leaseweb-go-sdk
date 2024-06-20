# AutoScalingLoadBalancer

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
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**Contract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** | Date and time when the instance was started for the first time, right after launching it | [optional] 
**Configuration** | Pointer to [**NullableLoadBalancerConfiguration**](LoadBalancerConfiguration.md) |  | [optional] 
**PrivateNetwork** | Pointer to [**NullablePrivateNetwork**](PrivateNetwork.md) |  | [optional] 

## Methods

### NewAutoScalingLoadBalancer

`func NewAutoScalingLoadBalancer() *AutoScalingLoadBalancer`

NewAutoScalingLoadBalancer instantiates a new AutoScalingLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingLoadBalancerWithDefaults

`func NewAutoScalingLoadBalancerWithDefaults() *AutoScalingLoadBalancer`

NewAutoScalingLoadBalancerWithDefaults instantiates a new AutoScalingLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *AutoScalingLoadBalancer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AutoScalingLoadBalancer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AutoScalingLoadBalancer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AutoScalingLoadBalancer) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *AutoScalingLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingLoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoScalingLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AutoScalingLoadBalancer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoScalingLoadBalancer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoScalingLoadBalancer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoScalingLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResources

`func (o *AutoScalingLoadBalancer) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AutoScalingLoadBalancer) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AutoScalingLoadBalancer) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AutoScalingLoadBalancer) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRegion

`func (o *AutoScalingLoadBalancer) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AutoScalingLoadBalancer) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AutoScalingLoadBalancer) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AutoScalingLoadBalancer) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *AutoScalingLoadBalancer) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AutoScalingLoadBalancer) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AutoScalingLoadBalancer) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AutoScalingLoadBalancer) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetState

`func (o *AutoScalingLoadBalancer) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingLoadBalancer) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingLoadBalancer) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *AutoScalingLoadBalancer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIps

`func (o *AutoScalingLoadBalancer) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *AutoScalingLoadBalancer) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *AutoScalingLoadBalancer) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *AutoScalingLoadBalancer) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetContract

`func (o *AutoScalingLoadBalancer) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AutoScalingLoadBalancer) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AutoScalingLoadBalancer) SetContract(v Contract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *AutoScalingLoadBalancer) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetStartedAt

`func (o *AutoScalingLoadBalancer) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AutoScalingLoadBalancer) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AutoScalingLoadBalancer) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AutoScalingLoadBalancer) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetConfiguration

`func (o *AutoScalingLoadBalancer) GetConfiguration() LoadBalancerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AutoScalingLoadBalancer) GetConfigurationOk() (*LoadBalancerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AutoScalingLoadBalancer) SetConfiguration(v LoadBalancerConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AutoScalingLoadBalancer) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *AutoScalingLoadBalancer) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *AutoScalingLoadBalancer) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetPrivateNetwork

`func (o *AutoScalingLoadBalancer) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *AutoScalingLoadBalancer) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *AutoScalingLoadBalancer) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *AutoScalingLoadBalancer) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### SetPrivateNetworkNil

`func (o *AutoScalingLoadBalancer) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *AutoScalingLoadBalancer) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


