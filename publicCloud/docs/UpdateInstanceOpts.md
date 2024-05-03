# UpdateInstanceOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Instance type | [optional] 
**Reference** | Pointer to **string** | An identifying name you can refer to the instance | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**ContractTerm** | Pointer to **int32** | Contract commitment. Can only be used when updating the contract type from HOURLY and MONTHLY. | [optional] 
**BillingFrequency** | Pointer to **int32** | How often you wish to be charged. Can only be used when updating the contract type from HOURLY to MONTHLY. &#39;1&#39; means every month, &#39;3&#39; every three months and so on. | [optional] 
**RootDiskSize** | Pointer to **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | [optional] 

## Methods

### NewUpdateInstanceOpts

`func NewUpdateInstanceOpts() *UpdateInstanceOpts`

NewUpdateInstanceOpts instantiates a new UpdateInstanceOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceOptsWithDefaults

`func NewUpdateInstanceOptsWithDefaults() *UpdateInstanceOpts`

NewUpdateInstanceOptsWithDefaults instantiates a new UpdateInstanceOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateInstanceOpts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateInstanceOpts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateInstanceOpts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateInstanceOpts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReference

`func (o *UpdateInstanceOpts) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UpdateInstanceOpts) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UpdateInstanceOpts) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UpdateInstanceOpts) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetContractType

`func (o *UpdateInstanceOpts) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *UpdateInstanceOpts) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *UpdateInstanceOpts) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *UpdateInstanceOpts) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetContractTerm

`func (o *UpdateInstanceOpts) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *UpdateInstanceOpts) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *UpdateInstanceOpts) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *UpdateInstanceOpts) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetBillingFrequency

`func (o *UpdateInstanceOpts) GetBillingFrequency() int32`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *UpdateInstanceOpts) GetBillingFrequencyOk() (*int32, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *UpdateInstanceOpts) SetBillingFrequency(v int32)`

SetBillingFrequency sets BillingFrequency field to given value.

### HasBillingFrequency

`func (o *UpdateInstanceOpts) HasBillingFrequency() bool`

HasBillingFrequency returns a boolean if a field has been set.

### GetRootDiskSize

`func (o *UpdateInstanceOpts) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *UpdateInstanceOpts) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *UpdateInstanceOpts) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *UpdateInstanceOpts) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


