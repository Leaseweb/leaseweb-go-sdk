# LaunchInstanceCreatedResult

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
**PublicIpV4** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | Instance type | [optional] 
**PrivateNetwork** | Pointer to **bool** |  | [optional] 
**RootDiskSize** | Pointer to **int32** | The root disk size as specified during its launch or update, in GB | [optional] 
**RootDiskStorageType** | Pointer to **string** | The root disk&#39;s storage type | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**BillingFrequency** | Pointer to **int32** |  | [optional] 
**ContractTerm** | Pointer to **int32** | The contract commitment (in months) | [optional] 
**ContractType** | Pointer to [**ContractType**](ContractType.md) |  | [optional] 
**ContractEndsAt** | Pointer to **NullableTime** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** | Date and time when the instance was started for the first time, right after launching it | [optional] 
**ContractRenewalsAt** | Pointer to **time.Time** | Date when the contract will be automatically renewed | [optional] 
**ContractCreatedAt** | Pointer to **time.Time** | Date when the contract was created | [optional] 
**Iso** | Pointer to [**NullableIso**](Iso.md) |  | [optional] 
**MarketAppId** | Pointer to **NullableString** | The market app ID which was deployed in the instance, if any | [optional] 

## Methods

### NewLaunchInstanceCreatedResult

`func NewLaunchInstanceCreatedResult() *LaunchInstanceCreatedResult`

NewLaunchInstanceCreatedResult instantiates a new LaunchInstanceCreatedResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchInstanceCreatedResultWithDefaults

`func NewLaunchInstanceCreatedResultWithDefaults() *LaunchInstanceCreatedResult`

NewLaunchInstanceCreatedResultWithDefaults instantiates a new LaunchInstanceCreatedResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *LaunchInstanceCreatedResult) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *LaunchInstanceCreatedResult) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *LaunchInstanceCreatedResult) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *LaunchInstanceCreatedResult) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *LaunchInstanceCreatedResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LaunchInstanceCreatedResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LaunchInstanceCreatedResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LaunchInstanceCreatedResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResources

`func (o *LaunchInstanceCreatedResult) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LaunchInstanceCreatedResult) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LaunchInstanceCreatedResult) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *LaunchInstanceCreatedResult) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRegion

`func (o *LaunchInstanceCreatedResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LaunchInstanceCreatedResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LaunchInstanceCreatedResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LaunchInstanceCreatedResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *LaunchInstanceCreatedResult) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LaunchInstanceCreatedResult) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LaunchInstanceCreatedResult) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LaunchInstanceCreatedResult) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *LaunchInstanceCreatedResult) GetOperatingSystem() OperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *LaunchInstanceCreatedResult) GetOperatingSystemOk() (*OperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *LaunchInstanceCreatedResult) SetOperatingSystem(v OperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *LaunchInstanceCreatedResult) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetState

`func (o *LaunchInstanceCreatedResult) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LaunchInstanceCreatedResult) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LaunchInstanceCreatedResult) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *LaunchInstanceCreatedResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPublicIpV4

`func (o *LaunchInstanceCreatedResult) GetPublicIpV4() bool`

GetPublicIpV4 returns the PublicIpV4 field if non-nil, zero value otherwise.

### GetPublicIpV4Ok

`func (o *LaunchInstanceCreatedResult) GetPublicIpV4Ok() (*bool, bool)`

GetPublicIpV4Ok returns a tuple with the PublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpV4

`func (o *LaunchInstanceCreatedResult) SetPublicIpV4(v bool)`

SetPublicIpV4 sets PublicIpV4 field to given value.

### HasPublicIpV4

`func (o *LaunchInstanceCreatedResult) HasPublicIpV4() bool`

HasPublicIpV4 returns a boolean if a field has been set.

### GetType

`func (o *LaunchInstanceCreatedResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LaunchInstanceCreatedResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LaunchInstanceCreatedResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LaunchInstanceCreatedResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *LaunchInstanceCreatedResult) GetPrivateNetwork() bool`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *LaunchInstanceCreatedResult) GetPrivateNetworkOk() (*bool, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *LaunchInstanceCreatedResult) SetPrivateNetwork(v bool)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *LaunchInstanceCreatedResult) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### GetRootDiskSize

`func (o *LaunchInstanceCreatedResult) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *LaunchInstanceCreatedResult) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *LaunchInstanceCreatedResult) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *LaunchInstanceCreatedResult) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetRootDiskStorageType

`func (o *LaunchInstanceCreatedResult) GetRootDiskStorageType() string`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *LaunchInstanceCreatedResult) GetRootDiskStorageTypeOk() (*string, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *LaunchInstanceCreatedResult) SetRootDiskStorageType(v string)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.

### HasRootDiskStorageType

`func (o *LaunchInstanceCreatedResult) HasRootDiskStorageType() bool`

HasRootDiskStorageType returns a boolean if a field has been set.

### GetIps

`func (o *LaunchInstanceCreatedResult) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *LaunchInstanceCreatedResult) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *LaunchInstanceCreatedResult) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *LaunchInstanceCreatedResult) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetBillingFrequency

`func (o *LaunchInstanceCreatedResult) GetBillingFrequency() int32`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *LaunchInstanceCreatedResult) GetBillingFrequencyOk() (*int32, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *LaunchInstanceCreatedResult) SetBillingFrequency(v int32)`

SetBillingFrequency sets BillingFrequency field to given value.

### HasBillingFrequency

`func (o *LaunchInstanceCreatedResult) HasBillingFrequency() bool`

HasBillingFrequency returns a boolean if a field has been set.

### GetContractTerm

`func (o *LaunchInstanceCreatedResult) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *LaunchInstanceCreatedResult) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *LaunchInstanceCreatedResult) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *LaunchInstanceCreatedResult) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetContractType

`func (o *LaunchInstanceCreatedResult) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *LaunchInstanceCreatedResult) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *LaunchInstanceCreatedResult) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *LaunchInstanceCreatedResult) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetContractEndsAt

`func (o *LaunchInstanceCreatedResult) GetContractEndsAt() time.Time`

GetContractEndsAt returns the ContractEndsAt field if non-nil, zero value otherwise.

### GetContractEndsAtOk

`func (o *LaunchInstanceCreatedResult) GetContractEndsAtOk() (*time.Time, bool)`

GetContractEndsAtOk returns a tuple with the ContractEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractEndsAt

`func (o *LaunchInstanceCreatedResult) SetContractEndsAt(v time.Time)`

SetContractEndsAt sets ContractEndsAt field to given value.

### HasContractEndsAt

`func (o *LaunchInstanceCreatedResult) HasContractEndsAt() bool`

HasContractEndsAt returns a boolean if a field has been set.

### SetContractEndsAtNil

`func (o *LaunchInstanceCreatedResult) SetContractEndsAtNil(b bool)`

 SetContractEndsAtNil sets the value for ContractEndsAt to be an explicit nil

### UnsetContractEndsAt
`func (o *LaunchInstanceCreatedResult) UnsetContractEndsAt()`

UnsetContractEndsAt ensures that no value is present for ContractEndsAt, not even an explicit nil
### GetStartedAt

`func (o *LaunchInstanceCreatedResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *LaunchInstanceCreatedResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *LaunchInstanceCreatedResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *LaunchInstanceCreatedResult) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *LaunchInstanceCreatedResult) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *LaunchInstanceCreatedResult) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetContractRenewalsAt

`func (o *LaunchInstanceCreatedResult) GetContractRenewalsAt() time.Time`

GetContractRenewalsAt returns the ContractRenewalsAt field if non-nil, zero value otherwise.

### GetContractRenewalsAtOk

`func (o *LaunchInstanceCreatedResult) GetContractRenewalsAtOk() (*time.Time, bool)`

GetContractRenewalsAtOk returns a tuple with the ContractRenewalsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractRenewalsAt

`func (o *LaunchInstanceCreatedResult) SetContractRenewalsAt(v time.Time)`

SetContractRenewalsAt sets ContractRenewalsAt field to given value.

### HasContractRenewalsAt

`func (o *LaunchInstanceCreatedResult) HasContractRenewalsAt() bool`

HasContractRenewalsAt returns a boolean if a field has been set.

### GetContractCreatedAt

`func (o *LaunchInstanceCreatedResult) GetContractCreatedAt() time.Time`

GetContractCreatedAt returns the ContractCreatedAt field if non-nil, zero value otherwise.

### GetContractCreatedAtOk

`func (o *LaunchInstanceCreatedResult) GetContractCreatedAtOk() (*time.Time, bool)`

GetContractCreatedAtOk returns a tuple with the ContractCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCreatedAt

`func (o *LaunchInstanceCreatedResult) SetContractCreatedAt(v time.Time)`

SetContractCreatedAt sets ContractCreatedAt field to given value.

### HasContractCreatedAt

`func (o *LaunchInstanceCreatedResult) HasContractCreatedAt() bool`

HasContractCreatedAt returns a boolean if a field has been set.

### GetIso

`func (o *LaunchInstanceCreatedResult) GetIso() Iso`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *LaunchInstanceCreatedResult) GetIsoOk() (*Iso, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *LaunchInstanceCreatedResult) SetIso(v Iso)`

SetIso sets Iso field to given value.

### HasIso

`func (o *LaunchInstanceCreatedResult) HasIso() bool`

HasIso returns a boolean if a field has been set.

### SetIsoNil

`func (o *LaunchInstanceCreatedResult) SetIsoNil(b bool)`

 SetIsoNil sets the value for Iso to be an explicit nil

### UnsetIso
`func (o *LaunchInstanceCreatedResult) UnsetIso()`

UnsetIso ensures that no value is present for Iso, not even an explicit nil
### GetMarketAppId

`func (o *LaunchInstanceCreatedResult) GetMarketAppId() string`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *LaunchInstanceCreatedResult) GetMarketAppIdOk() (*string, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *LaunchInstanceCreatedResult) SetMarketAppId(v string)`

SetMarketAppId sets MarketAppId field to given value.

### HasMarketAppId

`func (o *LaunchInstanceCreatedResult) HasMarketAppId() bool`

HasMarketAppId returns a boolean if a field has been set.

### SetMarketAppIdNil

`func (o *LaunchInstanceCreatedResult) SetMarketAppIdNil(b bool)`

 SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil

### UnsetMarketAppId
`func (o *LaunchInstanceCreatedResult) UnsetMarketAppId()`

UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


