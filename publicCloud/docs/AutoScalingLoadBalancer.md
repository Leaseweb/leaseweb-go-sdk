# AutoScalingLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The load balancer unique identifier | 
**Type** | **string** | Load balancer type | 
**Resources** | [**InstanceResources**](InstanceResources.md) |  | 
**Region** | **string** | The region where the load balancer was launched into | 
**Reference** | **string** | The identifying name set to the load balancer | 
**State** | [**InstanceState**](InstanceState.md) |  | 
**Ips** | [**[]Ip**](Ip.md) |  | 
**Contract** | [**Contract**](Contract.md) |  | 
**StartedAt** | **time.Time** | Date and time when the instance was started for the first time, right after launching it | 

## Methods

### NewAutoScalingLoadBalancer

`func NewAutoScalingLoadBalancer(id string, type_ string, resources InstanceResources, region string, reference string, state InstanceState, ips []Ip, contract Contract, startedAt time.Time, ) *AutoScalingLoadBalancer`

NewAutoScalingLoadBalancer instantiates a new AutoScalingLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingLoadBalancerWithDefaults

`func NewAutoScalingLoadBalancerWithDefaults() *AutoScalingLoadBalancer`

NewAutoScalingLoadBalancerWithDefaults instantiates a new AutoScalingLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


