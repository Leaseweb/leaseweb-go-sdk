# BaseFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseFirewallRule

`func NewBaseFirewallRule() *BaseFirewallRule`

NewBaseFirewallRule instantiates a new BaseFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFirewallRuleWithDefaults

`func NewBaseFirewallRuleWithDefaults() *BaseFirewallRule`

NewBaseFirewallRuleWithDefaults instantiates a new BaseFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseFirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseFirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseFirewallRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseFirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseFirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseFirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseFirewallRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseFirewallRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BaseFirewallRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BaseFirewallRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidr

`func (o *BaseFirewallRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BaseFirewallRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BaseFirewallRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *BaseFirewallRule) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


