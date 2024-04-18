# CreateRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**StartPort** | **NullableInt32** |  | 
**EndPort** | **NullableInt32** |  | 
**IcmpType** | **int32** |  | 
**IcmpCode** | **int32** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **NullableString** |  | 
**Cidr** | **string** |  | 

## Methods

### NewCreateRule

`func NewCreateRule(protocol string, startPort NullableInt32, endPort NullableInt32, icmpType int32, icmpCode int32, name NullableString, cidr string, ) *CreateRule`

NewCreateRule instantiates a new CreateRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRuleWithDefaults

`func NewCreateRuleWithDefaults() *CreateRule`

NewCreateRuleWithDefaults instantiates a new CreateRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *CreateRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetStartPort

`func (o *CreateRule) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *CreateRule) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *CreateRule) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.


### SetStartPortNil

`func (o *CreateRule) SetStartPortNil(b bool)`

 SetStartPortNil sets the value for StartPort to be an explicit nil

### UnsetStartPort
`func (o *CreateRule) UnsetStartPort()`

UnsetStartPort ensures that no value is present for StartPort, not even an explicit nil
### GetEndPort

`func (o *CreateRule) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *CreateRule) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *CreateRule) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.


### SetEndPortNil

`func (o *CreateRule) SetEndPortNil(b bool)`

 SetEndPortNil sets the value for EndPort to be an explicit nil

### UnsetEndPort
`func (o *CreateRule) UnsetEndPort()`

UnsetEndPort ensures that no value is present for EndPort, not even an explicit nil
### GetIcmpType

`func (o *CreateRule) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *CreateRule) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *CreateRule) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.


### GetIcmpCode

`func (o *CreateRule) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *CreateRule) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *CreateRule) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.


### GetId

`func (o *CreateRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRule) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidr

`func (o *CreateRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


