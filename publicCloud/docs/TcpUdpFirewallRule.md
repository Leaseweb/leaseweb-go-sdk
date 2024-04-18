# TcpUdpFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**StartPort** | **int32** |  | 
**EndPort** | **int32** | Value will be equals to the startPort when not provided | 
**IcmpType** | Pointer to **NullableInt32** |  | [optional] 
**IcmpCode** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTcpUdpFirewallRule

`func NewTcpUdpFirewallRule(startPort int32, endPort int32, ) *TcpUdpFirewallRule`

NewTcpUdpFirewallRule instantiates a new TcpUdpFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcpUdpFirewallRuleWithDefaults

`func NewTcpUdpFirewallRuleWithDefaults() *TcpUdpFirewallRule`

NewTcpUdpFirewallRuleWithDefaults instantiates a new TcpUdpFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TcpUdpFirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TcpUdpFirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TcpUdpFirewallRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TcpUdpFirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TcpUdpFirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TcpUdpFirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TcpUdpFirewallRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TcpUdpFirewallRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TcpUdpFirewallRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TcpUdpFirewallRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidr

`func (o *TcpUdpFirewallRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *TcpUdpFirewallRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *TcpUdpFirewallRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *TcpUdpFirewallRule) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetProtocol

`func (o *TcpUdpFirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *TcpUdpFirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *TcpUdpFirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *TcpUdpFirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStartPort

`func (o *TcpUdpFirewallRule) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *TcpUdpFirewallRule) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *TcpUdpFirewallRule) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.


### GetEndPort

`func (o *TcpUdpFirewallRule) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *TcpUdpFirewallRule) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *TcpUdpFirewallRule) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.


### GetIcmpType

`func (o *TcpUdpFirewallRule) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *TcpUdpFirewallRule) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *TcpUdpFirewallRule) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *TcpUdpFirewallRule) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### SetIcmpTypeNil

`func (o *TcpUdpFirewallRule) SetIcmpTypeNil(b bool)`

 SetIcmpTypeNil sets the value for IcmpType to be an explicit nil

### UnsetIcmpType
`func (o *TcpUdpFirewallRule) UnsetIcmpType()`

UnsetIcmpType ensures that no value is present for IcmpType, not even an explicit nil
### GetIcmpCode

`func (o *TcpUdpFirewallRule) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *TcpUdpFirewallRule) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *TcpUdpFirewallRule) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.

### HasIcmpCode

`func (o *TcpUdpFirewallRule) HasIcmpCode() bool`

HasIcmpCode returns a boolean if a field has been set.

### SetIcmpCodeNil

`func (o *TcpUdpFirewallRule) SetIcmpCodeNil(b bool)`

 SetIcmpCodeNil sets the value for IcmpCode to be an explicit nil

### UnsetIcmpCode
`func (o *TcpUdpFirewallRule) UnsetIcmpCode()`

UnsetIcmpCode ensures that no value is present for IcmpCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


