# VpsContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingFrequency** | [**BillingFrequency**](BillingFrequency.md) |  | 
**Term** | [**ContractTerm**](ContractTerm.md) |  | 
**Type** | [**ContractType**](ContractType.md) |  | 
**EndsAt** | **NullableTime** |  | 
**CreatedAt** | **time.Time** | Date when the contract was created | 
**State** | [**ContractState**](ContractState.md) |  | 
**Id** | **string** |  | 
**Sla** | **string** |  | 
**ControlPanel** | **NullableString** |  | 
**InModification** | **bool** |  | 

## Methods

### NewVpsContract

`func NewVpsContract(billingFrequency BillingFrequency, term ContractTerm, type_ ContractType, endsAt NullableTime, createdAt time.Time, state ContractState, id string, sla string, controlPanel NullableString, inModification bool, ) *VpsContract`

NewVpsContract instantiates a new VpsContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsContractWithDefaults

`func NewVpsContractWithDefaults() *VpsContract`

NewVpsContractWithDefaults instantiates a new VpsContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingFrequency

`func (o *VpsContract) GetBillingFrequency() BillingFrequency`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *VpsContract) GetBillingFrequencyOk() (*BillingFrequency, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *VpsContract) SetBillingFrequency(v BillingFrequency)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetTerm

`func (o *VpsContract) GetTerm() ContractTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *VpsContract) GetTermOk() (*ContractTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *VpsContract) SetTerm(v ContractTerm)`

SetTerm sets Term field to given value.


### GetType

`func (o *VpsContract) GetType() ContractType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VpsContract) GetTypeOk() (*ContractType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VpsContract) SetType(v ContractType)`

SetType sets Type field to given value.


### GetEndsAt

`func (o *VpsContract) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *VpsContract) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *VpsContract) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### SetEndsAtNil

`func (o *VpsContract) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *VpsContract) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetCreatedAt

`func (o *VpsContract) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpsContract) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpsContract) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetState

`func (o *VpsContract) GetState() ContractState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpsContract) GetStateOk() (*ContractState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpsContract) SetState(v ContractState)`

SetState sets State field to given value.


### GetId

`func (o *VpsContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpsContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpsContract) SetId(v string)`

SetId sets Id field to given value.


### GetSla

`func (o *VpsContract) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *VpsContract) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *VpsContract) SetSla(v string)`

SetSla sets Sla field to given value.


### GetControlPanel

`func (o *VpsContract) GetControlPanel() string`

GetControlPanel returns the ControlPanel field if non-nil, zero value otherwise.

### GetControlPanelOk

`func (o *VpsContract) GetControlPanelOk() (*string, bool)`

GetControlPanelOk returns a tuple with the ControlPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanel

`func (o *VpsContract) SetControlPanel(v string)`

SetControlPanel sets ControlPanel field to given value.


### SetControlPanelNil

`func (o *VpsContract) SetControlPanelNil(b bool)`

 SetControlPanelNil sets the value for ControlPanel to be an explicit nil

### UnsetControlPanel
`func (o *VpsContract) UnsetControlPanel()`

UnsetControlPanel ensures that no value is present for ControlPanel, not even an explicit nil
### GetInModification

`func (o *VpsContract) GetInModification() bool`

GetInModification returns the InModification field if non-nil, zero value otherwise.

### GetInModificationOk

`func (o *VpsContract) GetInModificationOk() (*bool, bool)`

GetInModificationOk returns a tuple with the InModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInModification

`func (o *VpsContract) SetInModification(v bool)`

SetInModification sets InModification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


