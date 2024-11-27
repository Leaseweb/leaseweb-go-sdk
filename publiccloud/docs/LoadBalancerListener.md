# LoadBalancerListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Port** | **int32** | Port for the listener | 
**Id** | **string** | The listener unique identifier | 
**Rules** | [**[]LoadBalancerListenerRule**](LoadBalancerListenerRule.md) |  | 

## Methods

### NewLoadBalancerListener

`func NewLoadBalancerListener(protocol Protocol, port int32, id string, rules []LoadBalancerListenerRule, ) *LoadBalancerListener`

NewLoadBalancerListener instantiates a new LoadBalancerListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListenerWithDefaults

`func NewLoadBalancerListenerWithDefaults() *LoadBalancerListener`

NewLoadBalancerListenerWithDefaults instantiates a new LoadBalancerListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *LoadBalancerListener) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerListener) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerListener) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *LoadBalancerListener) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerListener) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerListener) SetPort(v int32)`

SetPort sets Port field to given value.


### GetId

`func (o *LoadBalancerListener) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerListener) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerListener) SetId(v string)`

SetId sets Id field to given value.


### GetRules

`func (o *LoadBalancerListener) GetRules() []LoadBalancerListenerRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *LoadBalancerListener) GetRulesOk() (*[]LoadBalancerListenerRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *LoadBalancerListener) SetRules(v []LoadBalancerListenerRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


