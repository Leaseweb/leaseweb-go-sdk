# Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]ExpenseResultInstance**](ExpenseResultInstance.md) | List of instances to be billed in the period | [optional] 
**Traffic** | Pointer to [**Traffic**](Traffic.md) |  | [optional] 

## Methods

### NewBilling

`func NewBilling() *Billing`

NewBilling instantiates a new Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWithDefaults

`func NewBillingWithDefaults() *Billing`

NewBillingWithDefaults instantiates a new Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *Billing) GetInstances() []ExpenseResultInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Billing) GetInstancesOk() (*[]ExpenseResultInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Billing) SetInstances(v []ExpenseResultInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *Billing) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetTraffic

`func (o *Billing) GetTraffic() Traffic`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *Billing) GetTrafficOk() (*Traffic, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *Billing) SetTraffic(v Traffic)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *Billing) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


