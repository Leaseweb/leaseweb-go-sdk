# EditFirewallRulesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to [**[]FirewallRule**](FirewallRule.md) | List of rules that have been successfully edited | [optional] 
**Failed** | Pointer to [**[]FailedRule**](FailedRule.md) | List of rules that have failed to be edited along with their error messages | [optional] 

## Methods

### NewEditFirewallRulesResult

`func NewEditFirewallRulesResult() *EditFirewallRulesResult`

NewEditFirewallRulesResult instantiates a new EditFirewallRulesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditFirewallRulesResultWithDefaults

`func NewEditFirewallRulesResultWithDefaults() *EditFirewallRulesResult`

NewEditFirewallRulesResultWithDefaults instantiates a new EditFirewallRulesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *EditFirewallRulesResult) GetUpdated() []FirewallRule`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *EditFirewallRulesResult) GetUpdatedOk() (*[]FirewallRule, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *EditFirewallRulesResult) SetUpdated(v []FirewallRule)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *EditFirewallRulesResult) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetFailed

`func (o *EditFirewallRulesResult) GetFailed() []FailedRule`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *EditFirewallRulesResult) GetFailedOk() (*[]FailedRule, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *EditFirewallRulesResult) SetFailed(v []FailedRule)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *EditFirewallRulesResult) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


