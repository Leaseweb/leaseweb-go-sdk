# DeleteFirewallRulesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to [**[]FirewallRule**](FirewallRule.md) | List of rules that have been successfully deleted | [optional] 
**Failed** | Pointer to [**[]FirewallRule**](FirewallRule.md) | List of rules that have failed to be deleted | [optional] 

## Methods

### NewDeleteFirewallRulesResult

`func NewDeleteFirewallRulesResult() *DeleteFirewallRulesResult`

NewDeleteFirewallRulesResult instantiates a new DeleteFirewallRulesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFirewallRulesResultWithDefaults

`func NewDeleteFirewallRulesResultWithDefaults() *DeleteFirewallRulesResult`

NewDeleteFirewallRulesResultWithDefaults instantiates a new DeleteFirewallRulesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *DeleteFirewallRulesResult) GetDeleted() []FirewallRule`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteFirewallRulesResult) GetDeletedOk() (*[]FirewallRule, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteFirewallRulesResult) SetDeleted(v []FirewallRule)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteFirewallRulesResult) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFailed

`func (o *DeleteFirewallRulesResult) GetFailed() []FirewallRule`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *DeleteFirewallRulesResult) GetFailedOk() (*[]FirewallRule, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *DeleteFirewallRulesResult) SetFailed(v []FirewallRule)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *DeleteFirewallRulesResult) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


