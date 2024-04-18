# DeleteFirewallRulesOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]DeletionRule**](DeletionRule.md) |  | 

## Methods

### NewDeleteFirewallRulesOpts

`func NewDeleteFirewallRulesOpts(rules []DeletionRule, ) *DeleteFirewallRulesOpts`

NewDeleteFirewallRulesOpts instantiates a new DeleteFirewallRulesOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFirewallRulesOptsWithDefaults

`func NewDeleteFirewallRulesOptsWithDefaults() *DeleteFirewallRulesOpts`

NewDeleteFirewallRulesOptsWithDefaults instantiates a new DeleteFirewallRulesOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *DeleteFirewallRulesOpts) GetRules() []DeletionRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DeleteFirewallRulesOpts) GetRulesOk() (*[]DeletionRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DeleteFirewallRulesOpts) SetRules(v []DeletionRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


