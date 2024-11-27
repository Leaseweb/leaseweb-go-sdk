# LoadBalancerListenerOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**Protocol**](Protocol.md) |  | [optional] 
**Port** | Pointer to **int32** | Port that the listener listens to | [optional] 
**Certificate** | Pointer to [**SslCertificate**](SslCertificate.md) |  | [optional] 
**DefaultRule** | Pointer to [**LoadBalancerListenerDefaultRule**](LoadBalancerListenerDefaultRule.md) |  | [optional] 

## Methods

### NewLoadBalancerListenerOpts

`func NewLoadBalancerListenerOpts() *LoadBalancerListenerOpts`

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

`func (o *LoadBalancerListenerOpts) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerListenerOpts) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerListenerOpts) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LoadBalancerListenerOpts) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

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

### HasPort

`func (o *LoadBalancerListenerOpts) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCertificate

`func (o *LoadBalancerListenerOpts) GetCertificate() SslCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LoadBalancerListenerOpts) GetCertificateOk() (*SslCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LoadBalancerListenerOpts) SetCertificate(v SslCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LoadBalancerListenerOpts) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDefaultRule

`func (o *LoadBalancerListenerOpts) GetDefaultRule() LoadBalancerListenerDefaultRule`

GetDefaultRule returns the DefaultRule field if non-nil, zero value otherwise.

### GetDefaultRuleOk

`func (o *LoadBalancerListenerOpts) GetDefaultRuleOk() (*LoadBalancerListenerDefaultRule, bool)`

GetDefaultRuleOk returns a tuple with the DefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRule

`func (o *LoadBalancerListenerOpts) SetDefaultRule(v LoadBalancerListenerDefaultRule)`

SetDefaultRule sets DefaultRule field to given value.

### HasDefaultRule

`func (o *LoadBalancerListenerOpts) HasDefaultRule() bool`

HasDefaultRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


