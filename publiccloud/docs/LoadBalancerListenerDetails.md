# LoadBalancerListenerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Port** | **int32** | Port for the listener | 
**Id** | **string** | The listener unique identifier | 
**Rules** | [**[]LoadBalancerListenerRule**](LoadBalancerListenerRule.md) |  | 
**SslCertificates** | [**[]SslCertificate**](SslCertificate.md) |  | 

## Methods

### NewLoadBalancerListenerDetails

`func NewLoadBalancerListenerDetails(protocol Protocol, port int32, id string, rules []LoadBalancerListenerRule, sslCertificates []SslCertificate, ) *LoadBalancerListenerDetails`

NewLoadBalancerListenerDetails instantiates a new LoadBalancerListenerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListenerDetailsWithDefaults

`func NewLoadBalancerListenerDetailsWithDefaults() *LoadBalancerListenerDetails`

NewLoadBalancerListenerDetailsWithDefaults instantiates a new LoadBalancerListenerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *LoadBalancerListenerDetails) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerListenerDetails) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerListenerDetails) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *LoadBalancerListenerDetails) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerListenerDetails) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerListenerDetails) SetPort(v int32)`

SetPort sets Port field to given value.


### GetId

`func (o *LoadBalancerListenerDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerListenerDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerListenerDetails) SetId(v string)`

SetId sets Id field to given value.


### GetRules

`func (o *LoadBalancerListenerDetails) GetRules() []LoadBalancerListenerRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *LoadBalancerListenerDetails) GetRulesOk() (*[]LoadBalancerListenerRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *LoadBalancerListenerDetails) SetRules(v []LoadBalancerListenerRule)`

SetRules sets Rules field to given value.


### GetSslCertificates

`func (o *LoadBalancerListenerDetails) GetSslCertificates() []SslCertificate`

GetSslCertificates returns the SslCertificates field if non-nil, zero value otherwise.

### GetSslCertificatesOk

`func (o *LoadBalancerListenerDetails) GetSslCertificatesOk() (*[]SslCertificate, bool)`

GetSslCertificatesOk returns a tuple with the SslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificates

`func (o *LoadBalancerListenerDetails) SetSslCertificates(v []SslCertificate)`

SetSslCertificates sets SslCertificates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


