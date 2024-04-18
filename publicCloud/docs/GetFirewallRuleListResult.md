# GetFirewallRuleListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetFirewallRuleListResult

`func NewGetFirewallRuleListResult() *GetFirewallRuleListResult`

NewGetFirewallRuleListResult instantiates a new GetFirewallRuleListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFirewallRuleListResultWithDefaults

`func NewGetFirewallRuleListResultWithDefaults() *GetFirewallRuleListResult`

NewGetFirewallRuleListResultWithDefaults instantiates a new GetFirewallRuleListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallRules

`func (o *GetFirewallRuleListResult) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *GetFirewallRuleListResult) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *GetFirewallRuleListResult) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *GetFirewallRuleListResult) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.

### GetMetadata

`func (o *GetFirewallRuleListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetFirewallRuleListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetFirewallRuleListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetFirewallRuleListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


