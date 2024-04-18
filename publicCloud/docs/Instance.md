# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customer ID who owns the instance | [optional] 
**Id** | Pointer to **string** | The instance unique identifier | [optional] 
**Resources** | Pointer to [**InstanceResources**](InstanceResources.md) |  | [optional] 
**Region** | Pointer to **string** | The region where the instance was launched into | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the instance | [optional] 
**OperatingSystem** | Pointer to [**OperatingSystem**](OperatingSystem.md) |  | [optional] 
**State** | Pointer to [**InstanceState**](InstanceState.md) |  | [optional] 
**HasPublicIpV4** | Pointer to **bool** |  | [optional] 
**HasPrivateNetwork** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | Instance type | [optional] 
**RootDiskSize** | Pointer to **int32** | The root disk size as specified during its launch or update, in GB | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**BillingFrequency** | Pointer to **int32** |  | [optional] 
**ContractTerm** | Pointer to **int32** | The contract commitment (in months) | [optional] 
**ContractType** | Pointer to [**ContractType**](ContractType.md) |  | [optional] 
**ContractEndsAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** | Date and time when the instance was started for the first time, right after launching it | [optional] 
**ContractRenewalsAt** | Pointer to **time.Time** | Date when the contract will be automatically renewed | [optional] 
**ContractCreatedAt** | Pointer to **time.Time** | Date when the contract was created | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *Instance) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Instance) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Instance) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Instance) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResources

`func (o *Instance) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Instance) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Instance) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Instance) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRegion

`func (o *Instance) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Instance) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Instance) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Instance) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *Instance) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Instance) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Instance) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Instance) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Instance) GetOperatingSystem() OperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Instance) GetOperatingSystemOk() (*OperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Instance) SetOperatingSystem(v OperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Instance) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetState

`func (o *Instance) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *Instance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHasPublicIpV4

`func (o *Instance) GetHasPublicIpV4() bool`

GetHasPublicIpV4 returns the HasPublicIpV4 field if non-nil, zero value otherwise.

### GetHasPublicIpV4Ok

`func (o *Instance) GetHasPublicIpV4Ok() (*bool, bool)`

GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIpV4

`func (o *Instance) SetHasPublicIpV4(v bool)`

SetHasPublicIpV4 sets HasPublicIpV4 field to given value.

### HasHasPublicIpV4

`func (o *Instance) HasHasPublicIpV4() bool`

HasHasPublicIpV4 returns a boolean if a field has been set.

### GetHasPrivateNetwork

`func (o *Instance) GetHasPrivateNetwork() bool`

GetHasPrivateNetwork returns the HasPrivateNetwork field if non-nil, zero value otherwise.

### GetHasPrivateNetworkOk

`func (o *Instance) GetHasPrivateNetworkOk() (*bool, bool)`

GetHasPrivateNetworkOk returns a tuple with the HasPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateNetwork

`func (o *Instance) SetHasPrivateNetwork(v bool)`

SetHasPrivateNetwork sets HasPrivateNetwork field to given value.

### HasHasPrivateNetwork

`func (o *Instance) HasHasPrivateNetwork() bool`

HasHasPrivateNetwork returns a boolean if a field has been set.

### GetType

`func (o *Instance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Instance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRootDiskSize

`func (o *Instance) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *Instance) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *Instance) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *Instance) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetIps

`func (o *Instance) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Instance) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Instance) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *Instance) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetBillingFrequency

`func (o *Instance) GetBillingFrequency() int32`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *Instance) GetBillingFrequencyOk() (*int32, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *Instance) SetBillingFrequency(v int32)`

SetBillingFrequency sets BillingFrequency field to given value.

### HasBillingFrequency

`func (o *Instance) HasBillingFrequency() bool`

HasBillingFrequency returns a boolean if a field has been set.

### GetContractTerm

`func (o *Instance) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *Instance) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *Instance) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *Instance) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetContractType

`func (o *Instance) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *Instance) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *Instance) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *Instance) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetContractEndsAt

`func (o *Instance) GetContractEndsAt() time.Time`

GetContractEndsAt returns the ContractEndsAt field if non-nil, zero value otherwise.

### GetContractEndsAtOk

`func (o *Instance) GetContractEndsAtOk() (*time.Time, bool)`

GetContractEndsAtOk returns a tuple with the ContractEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractEndsAt

`func (o *Instance) SetContractEndsAt(v time.Time)`

SetContractEndsAt sets ContractEndsAt field to given value.

### HasContractEndsAt

`func (o *Instance) HasContractEndsAt() bool`

HasContractEndsAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Instance) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Instance) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Instance) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Instance) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetContractRenewalsAt

`func (o *Instance) GetContractRenewalsAt() time.Time`

GetContractRenewalsAt returns the ContractRenewalsAt field if non-nil, zero value otherwise.

### GetContractRenewalsAtOk

`func (o *Instance) GetContractRenewalsAtOk() (*time.Time, bool)`

GetContractRenewalsAtOk returns a tuple with the ContractRenewalsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractRenewalsAt

`func (o *Instance) SetContractRenewalsAt(v time.Time)`

SetContractRenewalsAt sets ContractRenewalsAt field to given value.

### HasContractRenewalsAt

`func (o *Instance) HasContractRenewalsAt() bool`

HasContractRenewalsAt returns a boolean if a field has been set.

### GetContractCreatedAt

`func (o *Instance) GetContractCreatedAt() time.Time`

GetContractCreatedAt returns the ContractCreatedAt field if non-nil, zero value otherwise.

### GetContractCreatedAtOk

`func (o *Instance) GetContractCreatedAtOk() (*time.Time, bool)`

GetContractCreatedAtOk returns a tuple with the ContractCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCreatedAt

`func (o *Instance) SetContractCreatedAt(v time.Time)`

SetContractCreatedAt sets ContractCreatedAt field to given value.

### HasContractCreatedAt

`func (o *Instance) HasContractCreatedAt() bool`

HasContractCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


