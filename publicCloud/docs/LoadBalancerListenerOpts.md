# LoadBalancerListenerOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**Port** | **int32** | Port that the listener listens to | 
**Certificate** | [**Certificate**](Certificate.md) |  | 
**DefaultRule** | [**NullableListenerDefaultRule**](ListenerDefaultRule.md) |  | 

## Methods

### NewLoadBalancerListenerOpts

`func NewLoadBalancerListenerOpts(protocol string, port int32, certificate Certificate, defaultRule NullableListenerDefaultRule, ) *LoadBalancerListenerOpts`

NewLoadBalancerListenerOpts instantiates a new LoadBalancerListenerOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListenerOptsWithDefaults

`func NewLoadBalancerListenerOptsWithDefaults() *LoadBalancerListenerOpts`

NewLoadBalancerListenerOptsWithDefaults instantiates a new LoadBalancerListenerOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *LoadBalancerListenerOpts) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerListenerOpts) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerListenerOpts) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *LoadBalancerListenerOpts) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerListenerOpts) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerListenerOpts) SetPort(v int32)`

SetPort sets Port field to given value.


### GetCertificate

`func (o *LoadBalancerListenerOpts) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LoadBalancerListenerOpts) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LoadBalancerListenerOpts) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.


### GetDefaultRule

`func (o *LoadBalancerListenerOpts) GetDefaultRule() ListenerDefaultRule`

GetDefaultRule returns the DefaultRule field if non-nil, zero value otherwise.

### GetDefaultRuleOk

`func (o *LoadBalancerListenerOpts) GetDefaultRuleOk() (*ListenerDefaultRule, bool)`

GetDefaultRuleOk returns a tuple with the DefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRule

`func (o *LoadBalancerListenerOpts) SetDefaultRule(v ListenerDefaultRule)`

SetDefaultRule sets DefaultRule field to given value.


### SetDefaultRuleNil

`func (o *LoadBalancerListenerOpts) SetDefaultRuleNil(b bool)`

 SetDefaultRuleNil sets the value for DefaultRule to be an explicit nil

### UnsetDefaultRule
`func (o *LoadBalancerListenerOpts) UnsetDefaultRule()`

UnsetDefaultRule ensures that no value is present for DefaultRule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


