# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The load balancer unique identifier | 
**Type** | **string** | Load balancer type | 
**Resources** | [**Resources**](Resources.md) |  | 
**Region** | **string** | The region where the load balancer was launched into | 
**Reference** | **NullableString** | The identifying name set to the load balancer | 
**State** | [**State**](State.md) |  | 
**Contract** | [**Contract**](Contract.md) |  | 
**StartedAt** | **NullableTime** | Date and time when the load balancer was started for the first time, right after launching it | 

## Methods

### NewLoadBalancer

`func NewLoadBalancer(id string, type_ string, resources Resources, region string, reference NullableString, state State, contract Contract, startedAt NullableTime, ) *LoadBalancer`

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

`func (o *LoadBalancer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancer) SetType(v string)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


