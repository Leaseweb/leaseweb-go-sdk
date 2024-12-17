# BaseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingFrequency** | [**BillingFrequency**](BillingFrequency.md) |  | 
**Term** | [**ContractTerm**](ContractTerm.md) |  | 
**Type** | [**ContractType**](ContractType.md) |  | 
**EndsAt** | **NullableTime** |  | 
**CreatedAt** | **time.Time** | Date when the contract was created | 
**State** | [**ContractState**](ContractState.md) |  | 

## Methods

### NewBaseContract

`func NewBaseContract(billingFrequency BillingFrequency, term ContractTerm, type_ ContractType, endsAt NullableTime, createdAt time.Time, state ContractState, ) *BaseContract`

NewBaseContract instantiates a new BaseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContractWithDefaults

`func NewBaseContractWithDefaults() *BaseContract`

NewBaseContractWithDefaults instantiates a new BaseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingFrequency

`func (o *BaseContract) GetBillingFrequency() BillingFrequency`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *BaseContract) GetBillingFrequencyOk() (*BillingFrequency, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *BaseContract) SetBillingFrequency(v BillingFrequency)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetTerm

`func (o *BaseContract) GetTerm() ContractTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *BaseContract) GetTermOk() (*ContractTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *BaseContract) SetTerm(v ContractTerm)`

SetTerm sets Term field to given value.


### GetType

`func (o *BaseContract) GetType() ContractType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseContract) GetTypeOk() (*ContractType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseContract) SetType(v ContractType)`

SetType sets Type field to given value.


### GetEndsAt

`func (o *BaseContract) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *BaseContract) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *BaseContract) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### SetEndsAtNil

`func (o *BaseContract) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *BaseContract) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetCreatedAt

`func (o *BaseContract) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseContract) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseContract) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetState

`func (o *BaseContract) GetState() ContractState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BaseContract) GetStateOk() (*ContractState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BaseContract) SetState(v ContractState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


