# InstanceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the instance | 
**Type** | [**TypeName**](TypeName.md) |  | 
**Resources** | [**Resources**](Resources.md) |  | 
**Region** | [**RegionName**](RegionName.md) |  | 
**Reference** | **NullableString** | The identifying name set to the instance | 
**StartedAt** | **NullableTime** | Date and time when the instance was started for the first time, right after launching it | 
**MarketAppId** | **NullableString** | Market App ID that must be installed into the instance | 
**State** | [**State**](State.md) |  | 
**ProductType** | **string** | The product type | 
**HasPublicIpV4** | **bool** |  | 
**IncludesPrivateNetwork** | **bool** |  | 
**HasUserData** | **bool** |  | 
**RootDiskSize** | **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | 
**RootDiskStorageType** | [**StorageType**](StorageType.md) |  | 
**Contract** | [**Contract**](Contract.md) |  | 
**AutoScalingGroup** | [**NullableAutoScalingGroup**](AutoScalingGroup.md) |  | 
**Image** | [**Image**](Image.md) |  | 
**Iso** | [**NullableIso**](Iso.md) |  | 
**PrivateNetwork** | [**NullablePrivateNetwork**](PrivateNetwork.md) |  | 
**Ips** | [**[]IpDetails**](IpDetails.md) |  | 

## Methods

### NewInstanceDetails

`func NewInstanceDetails(id string, type_ TypeName, resources Resources, region RegionName, reference NullableString, startedAt NullableTime, marketAppId NullableString, state State, productType string, hasPublicIpV4 bool, includesPrivateNetwork bool, hasUserData bool, rootDiskSize int32, rootDiskStorageType StorageType, contract Contract, autoScalingGroup NullableAutoScalingGroup, image Image, iso NullableIso, privateNetwork NullablePrivateNetwork, ips []IpDetails, ) *InstanceDetails`

NewInstanceDetails instantiates a new InstanceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDetailsWithDefaults

`func NewInstanceDetailsWithDefaults() *InstanceDetails`

NewInstanceDetailsWithDefaults instantiates a new InstanceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceDetails) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InstanceDetails) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceDetails) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceDetails) SetType(v TypeName)`

SetType sets Type field to given value.


### GetResources

`func (o *InstanceDetails) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *InstanceDetails) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *InstanceDetails) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetRegion

`func (o *InstanceDetails) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceDetails) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceDetails) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetReference

`func (o *InstanceDetails) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InstanceDetails) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InstanceDetails) SetReference(v string)`

SetReference sets Reference field to given value.


### SetReferenceNil

`func (o *InstanceDetails) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *InstanceDetails) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetStartedAt

`func (o *InstanceDetails) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InstanceDetails) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InstanceDetails) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *InstanceDetails) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *InstanceDetails) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetMarketAppId

`func (o *InstanceDetails) GetMarketAppId() string`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *InstanceDetails) GetMarketAppIdOk() (*string, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *InstanceDetails) SetMarketAppId(v string)`

SetMarketAppId sets MarketAppId field to given value.


### SetMarketAppIdNil

`func (o *InstanceDetails) SetMarketAppIdNil(b bool)`

 SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil

### UnsetMarketAppId
`func (o *InstanceDetails) UnsetMarketAppId()`

UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil
### GetState

`func (o *InstanceDetails) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceDetails) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceDetails) SetState(v State)`

SetState sets State field to given value.


### GetProductType

`func (o *InstanceDetails) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InstanceDetails) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InstanceDetails) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetHasPublicIpV4

`func (o *InstanceDetails) GetHasPublicIpV4() bool`

GetHasPublicIpV4 returns the HasPublicIpV4 field if non-nil, zero value otherwise.

### GetHasPublicIpV4Ok

`func (o *InstanceDetails) GetHasPublicIpV4Ok() (*bool, bool)`

GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIpV4

`func (o *InstanceDetails) SetHasPublicIpV4(v bool)`

SetHasPublicIpV4 sets HasPublicIpV4 field to given value.


### GetIncludesPrivateNetwork

`func (o *InstanceDetails) GetIncludesPrivateNetwork() bool`

GetIncludesPrivateNetwork returns the IncludesPrivateNetwork field if non-nil, zero value otherwise.

### GetIncludesPrivateNetworkOk

`func (o *InstanceDetails) GetIncludesPrivateNetworkOk() (*bool, bool)`

GetIncludesPrivateNetworkOk returns a tuple with the IncludesPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPrivateNetwork

`func (o *InstanceDetails) SetIncludesPrivateNetwork(v bool)`

SetIncludesPrivateNetwork sets IncludesPrivateNetwork field to given value.


### GetHasUserData

`func (o *InstanceDetails) GetHasUserData() bool`

GetHasUserData returns the HasUserData field if non-nil, zero value otherwise.

### GetHasUserDataOk

`func (o *InstanceDetails) GetHasUserDataOk() (*bool, bool)`

GetHasUserDataOk returns a tuple with the HasUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUserData

`func (o *InstanceDetails) SetHasUserData(v bool)`

SetHasUserData sets HasUserData field to given value.


### GetRootDiskSize

`func (o *InstanceDetails) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *InstanceDetails) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *InstanceDetails) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.


### GetRootDiskStorageType

`func (o *InstanceDetails) GetRootDiskStorageType() StorageType`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *InstanceDetails) GetRootDiskStorageTypeOk() (*StorageType, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *InstanceDetails) SetRootDiskStorageType(v StorageType)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.


### GetContract

`func (o *InstanceDetails) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *InstanceDetails) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *InstanceDetails) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetAutoScalingGroup

`func (o *InstanceDetails) GetAutoScalingGroup() AutoScalingGroup`

GetAutoScalingGroup returns the AutoScalingGroup field if non-nil, zero value otherwise.

### GetAutoScalingGroupOk

`func (o *InstanceDetails) GetAutoScalingGroupOk() (*AutoScalingGroup, bool)`

GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroup

`func (o *InstanceDetails) SetAutoScalingGroup(v AutoScalingGroup)`

SetAutoScalingGroup sets AutoScalingGroup field to given value.


### SetAutoScalingGroupNil

`func (o *InstanceDetails) SetAutoScalingGroupNil(b bool)`

 SetAutoScalingGroupNil sets the value for AutoScalingGroup to be an explicit nil

### UnsetAutoScalingGroup
`func (o *InstanceDetails) UnsetAutoScalingGroup()`

UnsetAutoScalingGroup ensures that no value is present for AutoScalingGroup, not even an explicit nil
### GetImage

`func (o *InstanceDetails) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InstanceDetails) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InstanceDetails) SetImage(v Image)`

SetImage sets Image field to given value.


### GetIso

`func (o *InstanceDetails) GetIso() Iso`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *InstanceDetails) GetIsoOk() (*Iso, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *InstanceDetails) SetIso(v Iso)`

SetIso sets Iso field to given value.


### SetIsoNil

`func (o *InstanceDetails) SetIsoNil(b bool)`

 SetIsoNil sets the value for Iso to be an explicit nil

### UnsetIso
`func (o *InstanceDetails) UnsetIso()`

UnsetIso ensures that no value is present for Iso, not even an explicit nil
### GetPrivateNetwork

`func (o *InstanceDetails) GetPrivateNetwork() PrivateNetwork`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *InstanceDetails) GetPrivateNetworkOk() (*PrivateNetwork, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *InstanceDetails) SetPrivateNetwork(v PrivateNetwork)`

SetPrivateNetwork sets PrivateNetwork field to given value.


### SetPrivateNetworkNil

`func (o *InstanceDetails) SetPrivateNetworkNil(b bool)`

 SetPrivateNetworkNil sets the value for PrivateNetwork to be an explicit nil

### UnsetPrivateNetwork
`func (o *InstanceDetails) UnsetPrivateNetwork()`

UnsetPrivateNetwork ensures that no value is present for PrivateNetwork, not even an explicit nil
### GetIps

`func (o *InstanceDetails) GetIps() []IpDetails`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *InstanceDetails) GetIpsOk() (*[]IpDetails, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *InstanceDetails) SetIps(v []IpDetails)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


