# UpdateInstanceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentId** | Pointer to **string** | The customer&#39;s equipment ID | [optional] 
**CustomerId** | Pointer to **string** | The customer&#39;s ID | [optional] 
**SalesOrgId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | The instance&#39;s identifier | [optional] 
**Type** | Pointer to **string** | Instance type | [optional] 
**Region** | Pointer to **string** | The region where the instance has been launched into | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the instance | [optional] 
**OperatingSystem** | Pointer to [**OperatingSystem**](OperatingSystem.md) |  | [optional] 
**Resources** | Pointer to [**InstanceResources**](InstanceResources.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**PublicIpV4** | Pointer to **bool** |  | [optional] 
**PrivateNetwork** | Pointer to **bool** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**RootDiskSize** | Pointer to **int32** |  | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**BillingFrequency** | Pointer to **int32** |  | [optional] 
**ContractTerm** | Pointer to **int32** |  | [optional] 
**ContractType** | Pointer to [**ContractType**](ContractType.md) |  | [optional] 
**ContractEndsAt** | Pointer to **NullableTime** |  | [optional] 
**ContractRenewalsAt** | Pointer to **time.Time** |  | [optional] 
**ContractCreatedAt** | Pointer to **time.Time** |  | [optional] 
**Iso** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateInstanceResult

`func NewUpdateInstanceResult() *UpdateInstanceResult`

NewUpdateInstanceResult instantiates a new UpdateInstanceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceResultWithDefaults

`func NewUpdateInstanceResultWithDefaults() *UpdateInstanceResult`

NewUpdateInstanceResultWithDefaults instantiates a new UpdateInstanceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentId

`func (o *UpdateInstanceResult) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *UpdateInstanceResult) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *UpdateInstanceResult) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *UpdateInstanceResult) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetCustomerId

`func (o *UpdateInstanceResult) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UpdateInstanceResult) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UpdateInstanceResult) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UpdateInstanceResult) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetSalesOrgId

`func (o *UpdateInstanceResult) GetSalesOrgId() string`

GetSalesOrgId returns the SalesOrgId field if non-nil, zero value otherwise.

### GetSalesOrgIdOk

`func (o *UpdateInstanceResult) GetSalesOrgIdOk() (*string, bool)`

GetSalesOrgIdOk returns a tuple with the SalesOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrgId

`func (o *UpdateInstanceResult) SetSalesOrgId(v string)`

SetSalesOrgId sets SalesOrgId field to given value.

### HasSalesOrgId

`func (o *UpdateInstanceResult) HasSalesOrgId() bool`

HasSalesOrgId returns a boolean if a field has been set.

### GetId

`func (o *UpdateInstanceResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInstanceResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateInstanceResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateInstanceResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateInstanceResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateInstanceResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegion

`func (o *UpdateInstanceResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateInstanceResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateInstanceResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UpdateInstanceResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *UpdateInstanceResult) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UpdateInstanceResult) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UpdateInstanceResult) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UpdateInstanceResult) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *UpdateInstanceResult) GetOperatingSystem() OperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *UpdateInstanceResult) GetOperatingSystemOk() (*OperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *UpdateInstanceResult) SetOperatingSystem(v OperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *UpdateInstanceResult) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetResources

`func (o *UpdateInstanceResult) GetResources() InstanceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UpdateInstanceResult) GetResourcesOk() (*InstanceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UpdateInstanceResult) SetResources(v InstanceResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *UpdateInstanceResult) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetState

`func (o *UpdateInstanceResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateInstanceResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateInstanceResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateInstanceResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPublicIpV4

`func (o *UpdateInstanceResult) GetPublicIpV4() bool`

GetPublicIpV4 returns the PublicIpV4 field if non-nil, zero value otherwise.

### GetPublicIpV4Ok

`func (o *UpdateInstanceResult) GetPublicIpV4Ok() (*bool, bool)`

GetPublicIpV4Ok returns a tuple with the PublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpV4

`func (o *UpdateInstanceResult) SetPublicIpV4(v bool)`

SetPublicIpV4 sets PublicIpV4 field to given value.

### HasPublicIpV4

`func (o *UpdateInstanceResult) HasPublicIpV4() bool`

HasPublicIpV4 returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *UpdateInstanceResult) GetPrivateNetwork() bool`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *UpdateInstanceResult) GetPrivateNetworkOk() (*bool, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *UpdateInstanceResult) SetPrivateNetwork(v bool)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *UpdateInstanceResult) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### GetStartedAt

`func (o *UpdateInstanceResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *UpdateInstanceResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *UpdateInstanceResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *UpdateInstanceResult) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *UpdateInstanceResult) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *UpdateInstanceResult) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetRootDiskSize

`func (o *UpdateInstanceResult) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *UpdateInstanceResult) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *UpdateInstanceResult) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *UpdateInstanceResult) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetIps

`func (o *UpdateInstanceResult) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *UpdateInstanceResult) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *UpdateInstanceResult) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *UpdateInstanceResult) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetBillingFrequency

`func (o *UpdateInstanceResult) GetBillingFrequency() int32`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *UpdateInstanceResult) GetBillingFrequencyOk() (*int32, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *UpdateInstanceResult) SetBillingFrequency(v int32)`

SetBillingFrequency sets BillingFrequency field to given value.

### HasBillingFrequency

`func (o *UpdateInstanceResult) HasBillingFrequency() bool`

HasBillingFrequency returns a boolean if a field has been set.

### GetContractTerm

`func (o *UpdateInstanceResult) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *UpdateInstanceResult) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *UpdateInstanceResult) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *UpdateInstanceResult) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetContractType

`func (o *UpdateInstanceResult) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *UpdateInstanceResult) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *UpdateInstanceResult) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *UpdateInstanceResult) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetContractEndsAt

`func (o *UpdateInstanceResult) GetContractEndsAt() time.Time`

GetContractEndsAt returns the ContractEndsAt field if non-nil, zero value otherwise.

### GetContractEndsAtOk

`func (o *UpdateInstanceResult) GetContractEndsAtOk() (*time.Time, bool)`

GetContractEndsAtOk returns a tuple with the ContractEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractEndsAt

`func (o *UpdateInstanceResult) SetContractEndsAt(v time.Time)`

SetContractEndsAt sets ContractEndsAt field to given value.

### HasContractEndsAt

`func (o *UpdateInstanceResult) HasContractEndsAt() bool`

HasContractEndsAt returns a boolean if a field has been set.

### SetContractEndsAtNil

`func (o *UpdateInstanceResult) SetContractEndsAtNil(b bool)`

 SetContractEndsAtNil sets the value for ContractEndsAt to be an explicit nil

### UnsetContractEndsAt
`func (o *UpdateInstanceResult) UnsetContractEndsAt()`

UnsetContractEndsAt ensures that no value is present for ContractEndsAt, not even an explicit nil
### GetContractRenewalsAt

`func (o *UpdateInstanceResult) GetContractRenewalsAt() time.Time`

GetContractRenewalsAt returns the ContractRenewalsAt field if non-nil, zero value otherwise.

### GetContractRenewalsAtOk

`func (o *UpdateInstanceResult) GetContractRenewalsAtOk() (*time.Time, bool)`

GetContractRenewalsAtOk returns a tuple with the ContractRenewalsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractRenewalsAt

`func (o *UpdateInstanceResult) SetContractRenewalsAt(v time.Time)`

SetContractRenewalsAt sets ContractRenewalsAt field to given value.

### HasContractRenewalsAt

`func (o *UpdateInstanceResult) HasContractRenewalsAt() bool`

HasContractRenewalsAt returns a boolean if a field has been set.

### GetContractCreatedAt

`func (o *UpdateInstanceResult) GetContractCreatedAt() time.Time`

GetContractCreatedAt returns the ContractCreatedAt field if non-nil, zero value otherwise.

### GetContractCreatedAtOk

`func (o *UpdateInstanceResult) GetContractCreatedAtOk() (*time.Time, bool)`

GetContractCreatedAtOk returns a tuple with the ContractCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCreatedAt

`func (o *UpdateInstanceResult) SetContractCreatedAt(v time.Time)`

SetContractCreatedAt sets ContractCreatedAt field to given value.

### HasContractCreatedAt

`func (o *UpdateInstanceResult) HasContractCreatedAt() bool`

HasContractCreatedAt returns a boolean if a field has been set.

### GetIso

`func (o *UpdateInstanceResult) GetIso() string`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *UpdateInstanceResult) GetIsoOk() (*string, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *UpdateInstanceResult) SetIso(v string)`

SetIso sets Iso field to given value.

### HasIso

`func (o *UpdateInstanceResult) HasIso() bool`

HasIso returns a boolean if a field has been set.

### SetIsoNil

`func (o *UpdateInstanceResult) SetIsoNil(b bool)`

 SetIsoNil sets the value for Iso to be an explicit nil

### UnsetIso
`func (o *UpdateInstanceResult) UnsetIso()`

UnsetIso ensures that no value is present for Iso, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


