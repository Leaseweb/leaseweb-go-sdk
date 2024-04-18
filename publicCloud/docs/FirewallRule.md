# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** |  | [optional] 
**StartPort** | **NullableInt32** |  | 
**EndPort** | **NullableInt32** |  | 
**IcmpType** | **int32** |  | 
**IcmpCode** | **int32** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallRule

`func NewFirewallRule(startPort NullableInt32, endPort NullableInt32, icmpType int32, icmpCode int32, ) *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *FirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStartPort

`func (o *FirewallRule) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *FirewallRule) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *FirewallRule) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.


### SetStartPortNil

`func (o *FirewallRule) SetStartPortNil(b bool)`

 SetStartPortNil sets the value for StartPort to be an explicit nil

### UnsetStartPort
`func (o *FirewallRule) UnsetStartPort()`

UnsetStartPort ensures that no value is present for StartPort, not even an explicit nil
### GetEndPort

`func (o *FirewallRule) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *FirewallRule) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *FirewallRule) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.


### SetEndPortNil

`func (o *FirewallRule) SetEndPortNil(b bool)`

 SetEndPortNil sets the value for EndPort to be an explicit nil

### UnsetEndPort
`func (o *FirewallRule) UnsetEndPort()`

UnsetEndPort ensures that no value is present for EndPort, not even an explicit nil
### GetIcmpType

`func (o *FirewallRule) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *FirewallRule) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *FirewallRule) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.


### GetIcmpCode

`func (o *FirewallRule) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *FirewallRule) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *FirewallRule) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.


### GetId

`func (o *FirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FirewallRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FirewallRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidr

`func (o *FirewallRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FirewallRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FirewallRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FirewallRule) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


