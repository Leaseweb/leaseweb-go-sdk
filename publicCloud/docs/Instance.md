# Instance

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
**Ips** | [**[]Ip**](Ip.md) |  | 

## Methods

### NewInstance

`func NewInstance(id string, type_ TypeName, resources Resources, region RegionName, reference NullableString, startedAt NullableTime, marketAppId NullableString, state State, productType string, hasPublicIpV4 bool, includesPrivateNetwork bool, hasUserData bool, rootDiskSize int32, rootDiskStorageType StorageType, contract Contract, autoScalingGroup NullableAutoScalingGroup, image Image, ips []Ip, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetType

`func (o *Instance) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v TypeName)`

SetType sets Type field to given value.


### GetResources

`func (o *Instance) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Instance) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Instance) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetRegion

`func (o *Instance) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Instance) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Instance) SetRegion(v RegionName)`

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
### GetState

`func (o *Instance) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v State)`

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


### GetIncludesPrivateNetwork

`func (o *Instance) GetIncludesPrivateNetwork() bool`

GetIncludesPrivateNetwork returns the IncludesPrivateNetwork field if non-nil, zero value otherwise.

### GetIncludesPrivateNetworkOk

`func (o *Instance) GetIncludesPrivateNetworkOk() (*bool, bool)`

GetIncludesPrivateNetworkOk returns a tuple with the IncludesPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPrivateNetwork

`func (o *Instance) SetIncludesPrivateNetwork(v bool)`

SetIncludesPrivateNetwork sets IncludesPrivateNetwork field to given value.


### GetHasUserData

`func (o *Instance) GetHasUserData() bool`

GetHasUserData returns the HasUserData field if non-nil, zero value otherwise.

### GetHasUserDataOk

`func (o *Instance) GetHasUserDataOk() (*bool, bool)`

GetHasUserDataOk returns a tuple with the HasUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUserData

`func (o *Instance) SetHasUserData(v bool)`

SetHasUserData sets HasUserData field to given value.


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

`func (o *Instance) GetRootDiskStorageType() StorageType`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *Instance) GetRootDiskStorageTypeOk() (*StorageType, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *Instance) SetRootDiskStorageType(v StorageType)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.


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
### GetImage

`func (o *Instance) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Instance) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Instance) SetImage(v Image)`

SetImage sets Image field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


