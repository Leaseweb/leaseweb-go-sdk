# LoadBalancerListenerCreateOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Port** | **int32** | Port that the listener listens to | 
**Certificate** | Pointer to [**SslCertificate**](SslCertificate.md) |  | [optional] 
**DefaultRule** | [**LoadBalancerListenerDefaultRule**](LoadBalancerListenerDefaultRule.md) |  | 

## Methods

### NewLoadBalancerListenerCreateOpts

`func NewLoadBalancerListenerCreateOpts(protocol Protocol, port int32, defaultRule LoadBalancerListenerDefaultRule, ) *LoadBalancerListenerCreateOpts`

NewLoadBalancerListenerCreateOpts instantiates a new LoadBalancerListenerCreateOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListenerCreateOptsWithDefaults

`func NewLoadBalancerListenerCreateOptsWithDefaults() *LoadBalancerListenerCreateOpts`

NewLoadBalancerListenerCreateOptsWithDefaults instantiates a new LoadBalancerListenerCreateOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *LoadBalancerListenerCreateOpts) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerListenerCreateOpts) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerListenerCreateOpts) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *LoadBalancerListenerCreateOpts) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerListenerCreateOpts) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerListenerCreateOpts) SetPort(v int32)`

SetPort sets Port field to given value.


### GetCertificate

`func (o *LoadBalancerListenerCreateOpts) GetCertificate() SslCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LoadBalancerListenerCreateOpts) GetCertificateOk() (*SslCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LoadBalancerListenerCreateOpts) SetCertificate(v SslCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LoadBalancerListenerCreateOpts) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDefaultRule

`func (o *LoadBalancerListenerCreateOpts) GetDefaultRule() LoadBalancerListenerDefaultRule`

GetDefaultRule returns the DefaultRule field if non-nil, zero value otherwise.

### GetDefaultRuleOk

`func (o *LoadBalancerListenerCreateOpts) GetDefaultRuleOk() (*LoadBalancerListenerDefaultRule, bool)`

GetDefaultRuleOk returns a tuple with the DefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRule

`func (o *LoadBalancerListenerCreateOpts) SetDefaultRule(v LoadBalancerListenerDefaultRule)`

SetDefaultRule sets DefaultRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


