# LoadBalancerListenerUpdateOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **NullableString** |  | 
**Port** | **NullableInt32** | Port that the listener listens to | 
**Certificate** | [**NullableCertificate1**](Certificate1.md) |  | 
**DefaultRule** | [**NullableListenerDefaultRule**](ListenerDefaultRule.md) |  | 

## Methods

### NewLoadBalancerListenerUpdateOpts

`func NewLoadBalancerListenerUpdateOpts(protocol NullableString, port NullableInt32, certificate NullableCertificate1, defaultRule NullableListenerDefaultRule, ) *LoadBalancerListenerUpdateOpts`

NewLoadBalancerListenerUpdateOpts instantiates a new LoadBalancerListenerUpdateOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListenerUpdateOptsWithDefaults

`func NewLoadBalancerListenerUpdateOptsWithDefaults() *LoadBalancerListenerUpdateOpts`

NewLoadBalancerListenerUpdateOptsWithDefaults instantiates a new LoadBalancerListenerUpdateOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *LoadBalancerListenerUpdateOpts) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerListenerUpdateOpts) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerListenerUpdateOpts) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *LoadBalancerListenerUpdateOpts) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *LoadBalancerListenerUpdateOpts) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPort

`func (o *LoadBalancerListenerUpdateOpts) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerListenerUpdateOpts) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerListenerUpdateOpts) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *LoadBalancerListenerUpdateOpts) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *LoadBalancerListenerUpdateOpts) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetCertificate

`func (o *LoadBalancerListenerUpdateOpts) GetCertificate() Certificate1`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LoadBalancerListenerUpdateOpts) GetCertificateOk() (*Certificate1, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LoadBalancerListenerUpdateOpts) SetCertificate(v Certificate1)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *LoadBalancerListenerUpdateOpts) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *LoadBalancerListenerUpdateOpts) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDefaultRule

`func (o *LoadBalancerListenerUpdateOpts) GetDefaultRule() ListenerDefaultRule`

GetDefaultRule returns the DefaultRule field if non-nil, zero value otherwise.

### GetDefaultRuleOk

`func (o *LoadBalancerListenerUpdateOpts) GetDefaultRuleOk() (*ListenerDefaultRule, bool)`

GetDefaultRuleOk returns a tuple with the DefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRule

`func (o *LoadBalancerListenerUpdateOpts) SetDefaultRule(v ListenerDefaultRule)`

SetDefaultRule sets DefaultRule field to given value.


### SetDefaultRuleNil

`func (o *LoadBalancerListenerUpdateOpts) SetDefaultRuleNil(b bool)`

 SetDefaultRuleNil sets the value for DefaultRule to be an explicit nil

### UnsetDefaultRule
`func (o *LoadBalancerListenerUpdateOpts) UnsetDefaultRule()`

UnsetDefaultRule ensures that no value is present for DefaultRule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


