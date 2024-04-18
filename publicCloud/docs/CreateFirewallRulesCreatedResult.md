# CreateFirewallRulesCreatedResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**[]FirewallRule**](FirewallRule.md) | List of rules that have been successfully added | [optional] 
**Failed** | Pointer to [**[]FirewallRule**](FirewallRule.md) | List of rules that have failed to be added | [optional] 

## Methods

### NewCreateFirewallRulesCreatedResult

`func NewCreateFirewallRulesCreatedResult() *CreateFirewallRulesCreatedResult`

NewCreateFirewallRulesCreatedResult instantiates a new CreateFirewallRulesCreatedResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRulesCreatedResultWithDefaults

`func NewCreateFirewallRulesCreatedResultWithDefaults() *CreateFirewallRulesCreatedResult`

NewCreateFirewallRulesCreatedResultWithDefaults instantiates a new CreateFirewallRulesCreatedResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CreateFirewallRulesCreatedResult) GetCreated() []FirewallRule`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateFirewallRulesCreatedResult) GetCreatedOk() (*[]FirewallRule, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateFirewallRulesCreatedResult) SetCreated(v []FirewallRule)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateFirewallRulesCreatedResult) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFailed

`func (o *CreateFirewallRulesCreatedResult) GetFailed() []FirewallRule`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CreateFirewallRulesCreatedResult) GetFailedOk() (*[]FirewallRule, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CreateFirewallRulesCreatedResult) SetFailed(v []FirewallRule)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CreateFirewallRulesCreatedResult) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


