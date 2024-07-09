# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingFrequency** | [**BillingFrequency**](BillingFrequency.md) |  | 
**Term** | [**ContractTerm**](ContractTerm.md) |  | 
**Type** | [**ContractType**](ContractType.md) |  | 
**EndsAt** | **NullableTime** |  | 
**RenewalsAt** | **time.Time** | Date when the contract will be automatically renewed | 
**CreatedAt** | **time.Time** | Date when the contract was created | 
**State** | [**ContractState**](ContractState.md) |  | 

## Methods

### NewContract

`func NewContract(billingFrequency BillingFrequency, term ContractTerm, type_ ContractType, endsAt NullableTime, renewalsAt time.Time, createdAt time.Time, state ContractState, ) *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingFrequency

`func (o *Contract) GetBillingFrequency() BillingFrequency`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *Contract) GetBillingFrequencyOk() (*BillingFrequency, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *Contract) SetBillingFrequency(v BillingFrequency)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetTerm

`func (o *Contract) GetTerm() ContractTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *Contract) GetTermOk() (*ContractTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *Contract) SetTerm(v ContractTerm)`

SetTerm sets Term field to given value.


### GetType

`func (o *Contract) GetType() ContractType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Contract) GetTypeOk() (*ContractType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Contract) SetType(v ContractType)`

SetType sets Type field to given value.


### GetEndsAt

`func (o *Contract) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *Contract) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *Contract) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### SetEndsAtNil

`func (o *Contract) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *Contract) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetRenewalsAt

`func (o *Contract) GetRenewalsAt() time.Time`

GetRenewalsAt returns the RenewalsAt field if non-nil, zero value otherwise.

### GetRenewalsAtOk

`func (o *Contract) GetRenewalsAtOk() (*time.Time, bool)`

GetRenewalsAtOk returns a tuple with the RenewalsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalsAt

`func (o *Contract) SetRenewalsAt(v time.Time)`

SetRenewalsAt sets RenewalsAt field to given value.


### GetCreatedAt

`func (o *Contract) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Contract) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Contract) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetState

`func (o *Contract) GetState() ContractState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Contract) GetStateOk() (*ContractState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Contract) SetState(v ContractState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


