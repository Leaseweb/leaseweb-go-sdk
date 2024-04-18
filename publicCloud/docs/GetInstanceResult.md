# GetInstanceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customer ID who owns the instance | [optional] 
**Id** | Pointer to **string** | The instance unique identifier | [optional] 
**Resources** | Pointer to [**InstanceResources**](InstanceResources.md) |  | [optional] 
**Type** | Pointer to **string** | Instance type | [optional] 
**Region** | Pointer to **string** | The region where the instance was launched into | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the instance | [optional] 
**OperatingSystem** | Pointer to [**OperatingSystem**](OperatingSystem.md) |  | [optional] 
**State** | Pointer to [**InstanceState**](InstanceState.md) |  | [optional] 
**HasPublicIpV4** | Pointer to **bool** |  | [optional] 
**HasPrivateNetwork** | Pointer to **bool** |  | [optional] 
**RootDiskSize** | Pointer to **int32** | The root disk size as specified during its launch or update, in GB | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**BillingFrequency** | Pointer to **int32** |  | [optional] 
**ContractTerm** | Pointer to **int32** | The contract commitment (in months) | [optional] 
**ContractType** | Pointer to [**ContractType**](ContractType.md) |  | [optional] 
**ContractEndsAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** | Date and time when the instance was started for the first time, right after launching it | [optional] 
**ContractRenewalsAt** | Pointer to **time.Time** | Date when the contract will be automatically renewed | [optional] 
**ContractCreatedAt** | Pointer to **time.Time** | Date when the contract was created | [optional] 
**Iso** | Pointer to [**NullableIso**](Iso.md) |  | [optional] 
**PrivateNetwork** | Pointer to [**PrivateNetwork**](PrivateNetwork.md) |  | [optional] 

## Methods

### NewGetInstanceResult

`func NewGetInstanceResult() *GetInstanceResult`

NewGetInstanceResult instantiates a new GetInstanceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceResultWithDefaults

`func NewGetInstanceResultWithDefaults() *GetInstanceResult`

NewGetInstanceResultWithDefaults instantiates a new GetInstanceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GetInstanceResult) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GetInstanceResult) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GetInstanceResult) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *GetInstanceResult) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *GetInstanceResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResources

`func (o *GetInstanceResult) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetInstanceResult) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetInstanceResult) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GetInstanceResult) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetType

`func (o *GetInstanceResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetInstanceResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetInstanceResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetInstanceResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegion

`func (o *GetInstanceResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetInstanceResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetInstanceResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetInstanceResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *GetInstanceResult) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetInstanceResult) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetInstanceResult) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetInstanceResult) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *GetInstanceResult) GetOperatingSystem() OperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *GetInstanceResult) GetOperatingSystemOk() (*OperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *GetInstanceResult) SetOperatingSystem(v OperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *GetInstanceResult) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetState

`func (o *GetInstanceResult) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetInstanceResult) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetInstanceResult) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *GetInstanceResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHasPublicIpV4

`func (o *GetInstanceResult) GetHasPublicIpV4() bool`

GetHasPublicIpV4 returns the HasPublicIpV4 field if non-nil, zero value otherwise.

### GetHasPublicIpV4Ok

`func (o *GetInstanceResult) GetHasPublicIpV4Ok() (*bool, bool)`

GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIpV4

`func (o *GetInstanceResult) SetHasPublicIpV4(v bool)`

SetHasPublicIpV4 sets HasPublicIpV4 field to given value.

### HasHasPublicIpV4

`func (o *GetInstanceResult) HasHasPublicIpV4() bool`

HasHasPublicIpV4 returns a boolean if a field has been set.

### GetHasPrivateNetwork

`func (o *GetInstanceResult) GetHasPrivateNetwork() bool`

GetHasPrivateNetwork returns the HasPrivateNetwork field if non-nil, zero value otherwise.

### GetHasPrivateNetworkOk

`func (o *GetInstanceResult) GetHasPrivateNetworkOk() (*bool, bool)`

GetHasPrivateNetworkOk returns a tuple with the HasPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateNetwork

`func (o *GetInstanceResult) SetHasPrivateNetwork(v bool)`

SetHasPrivateNetwork sets HasPrivateNetwork field to given value.

### HasHasPrivateNetwork

`func (o *GetInstanceResult) HasHasPrivateNetwork() bool`

HasHasPrivateNetwork returns a boolean if a field has been set.

### GetRootDiskSize

`func (o *GetInstanceResult) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *GetInstanceResult) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *GetInstanceResult) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *GetInstanceResult) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetIps

`func (o *GetInstanceResult) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *GetInstanceResult) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *GetInstanceResult) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *GetInstanceResult) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetBillingFrequency

`func (o *GetInstanceResult) GetBillingFrequency() int32`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *GetInstanceResult) GetBillingFrequencyOk() (*int32, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *GetInstanceResult) SetBillingFrequency(v int32)`

SetBillingFrequency sets BillingFrequency field to given value.

### HasBillingFrequency

`func (o *GetInstanceResult) HasBillingFrequency() bool`

HasBillingFrequency returns a boolean if a field has been set.

### GetContractTerm

`func (o *GetInstanceResult) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *GetInstanceResult) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *GetInstanceResult) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *GetInstanceResult) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetContractType

`func (o *GetInstanceResult) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *GetInstanceResult) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *GetInstanceResult) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *GetInstanceResult) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetContractEndsAt

`func (o *GetInstanceResult) GetContractEndsAt() time.Time`

GetContractEndsAt returns the ContractEndsAt field if non-nil, zero value otherwise.

### GetContractEndsAtOk

`func (o *GetInstanceResult) GetContractEndsAtOk() (*time.Time, bool)`

GetContractEndsAtOk returns a tuple with the ContractEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractEndsAt

`func (o *GetInstanceResult) SetContractEndsAt(v time.Time)`

SetContractEndsAt sets ContractEndsAt field to given value.

### HasContractEndsAt

`func (o *GetInstanceResult) HasContractEndsAt() bool`

HasContractEndsAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetInstanceResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetInstanceResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetInstanceResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetInstanceResult) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetContractRenewalsAt

`func (o *GetInstanceResult) GetContractRenewalsAt() time.Time`

GetContractRenewalsAt returns the ContractRenewalsAt field if non-nil, zero value otherwise.

### GetContractRenewalsAtOk

`func (o *GetInstanceResult) GetContractRenewalsAtOk() (*time.Time, bool)`

GetContractRenewalsAtOk returns a tuple with the ContractRenewalsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractRenewalsAt

`func (o *GetInstanceResult) SetContractRenewalsAt(v time.Time)`

SetContractRenewalsAt sets ContractRenewalsAt field to given value.

### HasContractRenewalsAt

`func (o *GetInstanceResult) HasContractRenewalsAt() bool`

HasContractRenewalsAt returns a boolean if a field has been set.

### GetContractCreatedAt

`func (o *GetInstanceResult) GetContractCreatedAt() time.Time`

GetContractCreatedAt returns the ContractCreatedAt field if non-nil, zero value otherwise.

### GetContractCreatedAtOk

`func (o *GetInstanceResult) GetContractCreatedAtOk() (*time.Time, bool)`

GetContractCreatedAtOk returns a tuple with the ContractCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCreatedAt

`func (o *GetInstanceResult) SetContractCreatedAt(v time.Time)`

SetContractCreatedAt sets ContractCreatedAt field to given value.

### HasContractCreatedAt

`func (o *GetInstanceResult) HasContractCreatedAt() bool`

HasContractCreatedAt returns a boolean if a field has been set.

### GetIso

`func (o *GetInstanceResult) GetIso() Iso`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *GetInstanceResult) GetIsoOk() (*Iso, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *GetInstanceResult) SetIso(v Iso)`

SetIso sets Iso field to given value.

### HasIso

`func (o *GetInstanceResult) HasIso() bool`

HasIso returns a boolean if a field has been set.

### SetIsoNil

`func (o *GetInstanceResult) SetIsoNil(b bool)`

 SetIsoNil sets the value for Iso to be an explicit nil

### UnsetIso
`func (o *GetInstanceResult) UnsetIso()`

UnsetIso ensures that no value is present for Iso, not even an explicit nil
### GetPrivateNetwork

`func (o *GetInstanceResult) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *GetInstanceResult) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *GetInstanceResult) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *GetInstanceResult) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


