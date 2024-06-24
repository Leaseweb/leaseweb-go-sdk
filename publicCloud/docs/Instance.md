# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentId** | Pointer to **string** | The customer&#39;s equipment ID | [optional] 
**SalesOrgId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** | The customer ID who owns the instance | [optional] 
**Id** | **string** | The unique identifier of the instance | 
**Resources** | [**InstanceResources**](InstanceResources.md) |  | 
**Region** | **string** | The region in which the instance was launched | 
**Reference** | **NullableString** | The identifying name set to the instance | 
**OperatingSystem** | [**OperatingSystem**](OperatingSystem.md) |  | 
**State** | [**InstanceState**](InstanceState.md) |  | 
**ProductType** | **string** | The product type | 
**HasPublicIpV4** | **bool** |  | 
**includesPrivateNetwork** | **bool** |  | 
**Type** | [**InstanceTypeName**](InstanceTypeName.md) |  | 
**RootDiskSize** | **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | 
**RootDiskStorageType** | **string** | The root disk&#39;s storage type | 
**Ips** | [**[]Ip**](Ip.md) |  | 
**StartedAt** | **NullableTime** | Date and time when the instance was started for the first time, right after launching it | 
**Contract** | [**Contract**](Contract.md) |  | 
**Iso** | [**NullableIso**](Iso.md) |  | 
**MarketAppId** | **NullableString** | Market App ID that must be installed into the instance | 
**PrivateNetwork** | [**NullablePrivateNetwork**](PrivateNetwork.md) |  | 
**AutoScalingGroup** | [**NullableAutoScalingGroup**](AutoScalingGroup.md) |  | 

## Methods

### NewInstance

`func NewInstance(id string, resources InstanceResources, region string, reference NullableString, operatingSystem OperatingSystem, state InstanceState, productType string, hasPublicIpV4 bool, includesPrivateNetwork bool, type_ InstanceTypeName, rootDiskSize int32, rootDiskStorageType string, ips []Ip, startedAt NullableTime, contract Contract, iso NullableIso, marketAppId NullableString, privateNetwork NullablePrivateNetwork, autoScalingGroup NullableAutoScalingGroup, ) *Instance`

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


### SetReferenceNil

`func (o *Instance) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *Instance) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
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


### SetPrivateNetworkNil

`func (o *Instance) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *Instance) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil
### GetAutoScalingGroup

`func (o *Instance) GetAutoScalingGroup() AutoScalingGroup`

GetAutoScalingGroup returns the AutoScalingGroup field if non-nil, zero value otherwise.

### GetAutoScalingGroupOk

`func (o *Instance) GetAutoScalingGroupOk() (*AutoScalingGroup, bool)`

GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroup

`func (o *Instance) SetAutoScalingGroup(v AutoScalingGroup)`

SetAutoScalingGroup sets AutoScalingGroup field to given value.


### SetAutoScalingGroupNil

`func (o *Instance) SetAutoScalingGroupNil(b bool)`

 SetAutoScalingGroupNil sets the value for AutoScalingGroup to be an explicit nil

### UnsetAutoScalingGroup
`func (o *Instance) UnsetAutoScalingGroup()`

UnsetAutoScalingGroup ensures that no value is present for AutoScalingGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


