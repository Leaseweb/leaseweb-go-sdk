# LoadBalancer1

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
**Configuration** | Pointer to [**NullableLoadBalancerConfiguration1**](LoadBalancerConfiguration1.md) |  | [optional] 
**PrivateNetwork** | Pointer to [**NullablePrivateNetwork**](PrivateNetwork.md) |  | [optional] 

## Methods

### NewLoadBalancer1

`func NewLoadBalancer1() *LoadBalancer1`

NewLoadBalancer1 instantiates a new LoadBalancer1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancer1WithDefaults

`func NewLoadBalancer1WithDefaults() *LoadBalancer1`

NewLoadBalancer1WithDefaults instantiates a new LoadBalancer1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *LoadBalancer1) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *LoadBalancer1) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *LoadBalancer1) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *LoadBalancer1) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *LoadBalancer1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancer1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancer1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancer1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LoadBalancer1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancer1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancer1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LoadBalancer1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResources

`func (o *LoadBalancer1) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LoadBalancer1) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LoadBalancer1) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *LoadBalancer1) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRegion

`func (o *LoadBalancer1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoadBalancer1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoadBalancer1) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoadBalancer1) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *LoadBalancer1) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LoadBalancer1) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LoadBalancer1) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LoadBalancer1) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetState

`func (o *LoadBalancer1) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoadBalancer1) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoadBalancer1) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *LoadBalancer1) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIps

`func (o *LoadBalancer1) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *LoadBalancer1) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *LoadBalancer1) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *LoadBalancer1) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetContract

`func (o *LoadBalancer1) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *LoadBalancer1) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *LoadBalancer1) SetContract(v Contract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *LoadBalancer1) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetStartedAt

`func (o *LoadBalancer1) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *LoadBalancer1) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *LoadBalancer1) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *LoadBalancer1) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetConfiguration

`func (o *LoadBalancer1) GetConfiguration() LoadBalancerConfiguration1`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *LoadBalancer1) GetConfigurationOk() (*LoadBalancerConfiguration1, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *LoadBalancer1) SetConfiguration(v LoadBalancerConfiguration1)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *LoadBalancer1) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *LoadBalancer1) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *LoadBalancer1) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetPrivateNetwork

`func (o *LoadBalancer1) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *LoadBalancer1) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *LoadBalancer1) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *LoadBalancer1) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### SetPrivateNetworkNil

`func (o *LoadBalancer1) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *LoadBalancer1) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


