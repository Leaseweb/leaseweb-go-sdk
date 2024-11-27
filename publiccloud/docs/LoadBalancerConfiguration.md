# LoadBalancerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StickySession** | [**NullableStickySession**](StickySession.md) |  | 
**Balance** | [**Balance**](Balance.md) |  | 
**XForwardedFor** | **bool** | Is xForwardedFor header enabled or not | 
**IdleTimeOut** | **int32** | Time to close the connection if load balancer is idle | 

## Methods

### NewLoadBalancerConfiguration

`func NewLoadBalancerConfiguration(stickySession NullableStickySession, balance Balance, xForwardedFor bool, idleTimeOut int32, ) *LoadBalancerConfiguration`

NewLoadBalancerConfiguration instantiates a new LoadBalancerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerConfigurationWithDefaults

`func NewLoadBalancerConfigurationWithDefaults() *LoadBalancerConfiguration`

NewLoadBalancerConfigurationWithDefaults instantiates a new LoadBalancerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStickySession

`func (o *LoadBalancerConfiguration) GetStickySession() StickySession`

GetStickySession returns the StickySession field if non-nil, zero value otherwise.

### GetStickySessionOk

`func (o *LoadBalancerConfiguration) GetStickySessionOk() (*StickySession, bool)`

GetStickySessionOk returns a tuple with the StickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySession

`func (o *LoadBalancerConfiguration) SetStickySession(v StickySession)`

SetStickySession sets StickySession field to given value.


### SetStickySessionNil

`func (o *LoadBalancerConfiguration) SetStickySessionNil(b bool)`

 SetStickySessionNil sets the value for StickySession to be an explicit nil

### UnsetStickySession
`func (o *LoadBalancerConfiguration) UnsetStickySession()`

UnsetStickySession ensures that no value is present for StickySession, not even an explicit nil
### GetBalance

`func (o *LoadBalancerConfiguration) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *LoadBalancerConfiguration) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *LoadBalancerConfiguration) SetBalance(v Balance)`

SetBalance sets Balance field to given value.


### GetXForwardedFor

`func (o *LoadBalancerConfiguration) GetXForwardedFor() bool`

GetXForwardedFor returns the XForwardedFor field if non-nil, zero value otherwise.

### GetXForwardedForOk

`func (o *LoadBalancerConfiguration) GetXForwardedForOk() (*bool, bool)`

GetXForwardedForOk returns a tuple with the XForwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedFor

`func (o *LoadBalancerConfiguration) SetXForwardedFor(v bool)`

SetXForwardedFor sets XForwardedFor field to given value.


### GetIdleTimeOut

`func (o *LoadBalancerConfiguration) GetIdleTimeOut() int32`

GetIdleTimeOut returns the IdleTimeOut field if non-nil, zero value otherwise.

### GetIdleTimeOutOk

`func (o *LoadBalancerConfiguration) GetIdleTimeOutOk() (*int32, bool)`

GetIdleTimeOutOk returns a tuple with the IdleTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeOut

`func (o *LoadBalancerConfiguration) SetIdleTimeOut(v int32)`

SetIdleTimeOut sets IdleTimeOut field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


