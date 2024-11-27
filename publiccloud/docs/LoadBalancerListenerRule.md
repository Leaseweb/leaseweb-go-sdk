# LoadBalancerListenerRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The listener unique identifier | 
**Default** | **bool** | Condition of the rule | 
**TargetGroupId** | **string** | The target group unique identifier | 

## Methods

### NewLoadBalancerListenerRule

`func NewLoadBalancerListenerRule(id string, default_ bool, targetGroupId string, ) *LoadBalancerListenerRule`

NewLoadBalancerListenerRule instantiates a new LoadBalancerListenerRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListenerRuleWithDefaults

`func NewLoadBalancerListenerRuleWithDefaults() *LoadBalancerListenerRule`

NewLoadBalancerListenerRuleWithDefaults instantiates a new LoadBalancerListenerRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerListenerRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerListenerRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerListenerRule) SetId(v string)`

SetId sets Id field to given value.


### GetDefault

`func (o *LoadBalancerListenerRule) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *LoadBalancerListenerRule) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *LoadBalancerListenerRule) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetTargetGroupId

`func (o *LoadBalancerListenerRule) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *LoadBalancerListenerRule) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *LoadBalancerListenerRule) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


