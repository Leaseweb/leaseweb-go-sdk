# IcmpFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**StartPort** | Pointer to **NullableInt32** |  | [optional] 
**EndPort** | Pointer to **NullableInt32** |  | [optional] 
**IcmpType** | **int32** |  | 
**IcmpCode** | **int32** |  | 

## Methods

### NewIcmpFirewallRule

`func NewIcmpFirewallRule(icmpType int32, icmpCode int32, ) *IcmpFirewallRule`

NewIcmpFirewallRule instantiates a new IcmpFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIcmpFirewallRuleWithDefaults

`func NewIcmpFirewallRuleWithDefaults() *IcmpFirewallRule`

NewIcmpFirewallRuleWithDefaults instantiates a new IcmpFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IcmpFirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IcmpFirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IcmpFirewallRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IcmpFirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IcmpFirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IcmpFirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IcmpFirewallRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IcmpFirewallRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IcmpFirewallRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IcmpFirewallRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidr

`func (o *IcmpFirewallRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IcmpFirewallRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IcmpFirewallRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IcmpFirewallRule) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetProtocol

`func (o *IcmpFirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IcmpFirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IcmpFirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IcmpFirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStartPort

`func (o *IcmpFirewallRule) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *IcmpFirewallRule) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *IcmpFirewallRule) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *IcmpFirewallRule) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.

### SetStartPortNil

`func (o *IcmpFirewallRule) SetStartPortNil(b bool)`

 SetStartPortNil sets the value for StartPort to be an explicit nil

### UnsetStartPort
`func (o *IcmpFirewallRule) UnsetStartPort()`

UnsetStartPort ensures that no value is present for StartPort, not even an explicit nil
### GetEndPort

`func (o *IcmpFirewallRule) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *IcmpFirewallRule) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *IcmpFirewallRule) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *IcmpFirewallRule) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### SetEndPortNil

`func (o *IcmpFirewallRule) SetEndPortNil(b bool)`

 SetEndPortNil sets the value for EndPort to be an explicit nil

### UnsetEndPort
`func (o *IcmpFirewallRule) UnsetEndPort()`

UnsetEndPort ensures that no value is present for EndPort, not even an explicit nil
### GetIcmpType

`func (o *IcmpFirewallRule) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *IcmpFirewallRule) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *IcmpFirewallRule) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.


### GetIcmpCode

`func (o *IcmpFirewallRule) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *IcmpFirewallRule) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *IcmpFirewallRule) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


