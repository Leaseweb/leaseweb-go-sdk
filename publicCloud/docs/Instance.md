# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentId** | Pointer to **string** | The customer&#39;s equipment ID | [optional] 
**SalesOrgId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** | The customer ID who owns the instance | [optional] 
**Id** | Pointer to **string** | The unique identifier of the instance | [optional] 
**Resources** | Pointer to [**InstanceResources**](InstanceResources.md) |  | [optional] 
**Region** | Pointer to **string** | The region in which the instance was launched | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the instance | [optional] 
**OperatingSystem** | Pointer to [**OperatingSystem**](OperatingSystem.md) |  | [optional] 
**State** | Pointer to [**InstanceState**](InstanceState.md) |  | [optional] 
**ProductType** | Pointer to **string** | The product type | [optional] 
**HasPublicIpV4** | Pointer to **bool** |  | [optional] 
**includesPrivateNetwork** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**InstanceTypeName**](InstanceTypeName.md) |  | [optional] 
**RootDiskSize** | Pointer to **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | [optional] 
**RootDiskStorageType** | Pointer to **string** | The root disk&#39;s storage type | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 
**StartedAt** | Pointer to **NullableTime** | Date and time when the instance was started for the first time, right after launching it | [optional] 
**Contract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**Iso** | Pointer to [**NullableIso**](Iso.md) |  | [optional] 
**MarketAppId** | Pointer to **NullableString** | Market App ID that must be installed into the instance | [optional] 
**PrivateNetwork** | Pointer to [**NullablePrivateNetwork**](PrivateNetwork.md) |  | [optional] 

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

### GetEquipmentId

`func (o *Instance) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *Instance) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *Instance) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *Instance) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetSalesOrgId

`func (o *Instance) GetSalesOrgId() string`

GetSalesOrgId returns the SalesOrgId field if non-nil, zero value otherwise.

### GetSalesOrgIdOk

`func (o *Instance) GetSalesOrgIdOk() (*string, bool)`

GetSalesOrgIdOk returns a tuple with the SalesOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrgId

`func (o *Instance) SetSalesOrgId(v string)`

SetSalesOrgId sets SalesOrgId field to given value.

### HasSalesOrgId

`func (o *Instance) HasSalesOrgId() bool`

HasSalesOrgId returns a boolean if a field has been set.

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

### GetProductType

`func (o *Instance) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Instance) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Instance) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Instance) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

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

### GetincludesPrivateNetwork

`func (o *Instance) GetincludesPrivateNetwork() bool`

GetincludesPrivateNetwork returns the includesPrivateNetwork field if non-nil, zero value otherwise.

### GetincludesPrivateNetworkOk

`func (o *Instance) GetincludesPrivateNetworkOk() (*bool, bool)`

GetincludesPrivateNetworkOk returns a tuple with the includesPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetincludesPrivateNetwork

`func (o *Instance) SetincludesPrivateNetwork(v bool)`

SetincludesPrivateNetwork sets includesPrivateNetwork field to given value.

### HasincludesPrivateNetwork

`func (o *Instance) HasincludesPrivateNetwork() bool`

HasincludesPrivateNetwork returns a boolean if a field has been set.

### GetType

`func (o *Instance) GetType() InstanceTypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*InstanceTypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v InstanceTypeName)`

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

### GetRootDiskStorageType

`func (o *Instance) GetRootDiskStorageType() string`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *Instance) GetRootDiskStorageTypeOk() (*string, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *Instance) SetRootDiskStorageType(v string)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.

### HasRootDiskStorageType

`func (o *Instance) HasRootDiskStorageType() bool`

HasRootDiskStorageType returns a boolean if a field has been set.

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

### SetStartedAtNil

`func (o *Instance) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Instance) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetContract

`func (o *Instance) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *Instance) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *Instance) SetContract(v Contract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *Instance) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetIso

`func (o *Instance) GetIso() Iso`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *Instance) GetIsoOk() (*Iso, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *Instance) SetIso(v Iso)`

SetIso sets Iso field to given value.

### HasIso

`func (o *Instance) HasIso() bool`

HasIso returns a boolean if a field has been set.

### SetIsoNil

`func (o *Instance) SetIsoNil(b bool)`

 SetIsoNil sets the value for Iso to be an explicit nil

### UnsetIso
`func (o *Instance) UnsetIso()`

UnsetIso ensures that no value is present for Iso, not even an explicit nil
### GetMarketAppId

`func (o *Instance) GetMarketAppId() string`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *Instance) GetMarketAppIdOk() (*string, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *Instance) SetMarketAppId(v string)`

SetMarketAppId sets MarketAppId field to given value.

### HasMarketAppId

`func (o *Instance) HasMarketAppId() bool`

HasMarketAppId returns a boolean if a field has been set.

### SetMarketAppIdNil

`func (o *Instance) SetMarketAppIdNil(b bool)`

 SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil

### UnsetMarketAppId
`func (o *Instance) UnsetMarketAppId()`

UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil
### GetPrivateNetwork

`func (o *Instance) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *Instance) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *Instance) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *Instance) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### SetPrivateNetworkNil

`func (o *Instance) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *Instance) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


