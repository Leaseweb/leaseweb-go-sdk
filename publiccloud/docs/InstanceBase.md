# InstanceBase

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

## Methods

### NewInstanceBase

`func NewInstanceBase(id string, type_ TypeName, resources Resources, region RegionName, reference NullableString, startedAt NullableTime, marketAppId NullableString, state State, productType string, hasPublicIpV4 bool, includesPrivateNetwork bool, hasUserData bool, rootDiskSize int32, rootDiskStorageType StorageType, contract Contract, autoScalingGroup NullableAutoScalingGroup, image Image, ) *InstanceBase`

NewInstanceBase instantiates a new InstanceBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceBaseWithDefaults

`func NewInstanceBaseWithDefaults() *InstanceBase`

NewInstanceBaseWithDefaults instantiates a new InstanceBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceBase) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InstanceBase) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceBase) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceBase) SetType(v TypeName)`

SetType sets Type field to given value.


### GetResources

`func (o *InstanceBase) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *InstanceBase) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *InstanceBase) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetRegion

`func (o *InstanceBase) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceBase) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceBase) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetReference

`func (o *InstanceBase) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InstanceBase) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InstanceBase) SetReference(v string)`

SetReference sets Reference field to given value.


### SetReferenceNil

`func (o *InstanceBase) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *InstanceBase) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetStartedAt

`func (o *InstanceBase) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InstanceBase) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InstanceBase) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *InstanceBase) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *InstanceBase) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetMarketAppId

`func (o *InstanceBase) GetMarketAppId() string`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *InstanceBase) GetMarketAppIdOk() (*string, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *InstanceBase) SetMarketAppId(v string)`

SetMarketAppId sets MarketAppId field to given value.


### SetMarketAppIdNil

`func (o *InstanceBase) SetMarketAppIdNil(b bool)`

 SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil

### UnsetMarketAppId
`func (o *InstanceBase) UnsetMarketAppId()`

UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil
### GetState

`func (o *InstanceBase) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceBase) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceBase) SetState(v State)`

SetState sets State field to given value.


### GetProductType

`func (o *InstanceBase) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InstanceBase) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InstanceBase) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetHasPublicIpV4

`func (o *InstanceBase) GetHasPublicIpV4() bool`

GetHasPublicIpV4 returns the HasPublicIpV4 field if non-nil, zero value otherwise.

### GetHasPublicIpV4Ok

`func (o *InstanceBase) GetHasPublicIpV4Ok() (*bool, bool)`

GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIpV4

`func (o *InstanceBase) SetHasPublicIpV4(v bool)`

SetHasPublicIpV4 sets HasPublicIpV4 field to given value.


### GetIncludesPrivateNetwork

`func (o *InstanceBase) GetIncludesPrivateNetwork() bool`

GetIncludesPrivateNetwork returns the IncludesPrivateNetwork field if non-nil, zero value otherwise.

### GetIncludesPrivateNetworkOk

`func (o *InstanceBase) GetIncludesPrivateNetworkOk() (*bool, bool)`

GetIncludesPrivateNetworkOk returns a tuple with the IncludesPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPrivateNetwork

`func (o *InstanceBase) SetIncludesPrivateNetwork(v bool)`

SetIncludesPrivateNetwork sets IncludesPrivateNetwork field to given value.


### GetHasUserData

`func (o *InstanceBase) GetHasUserData() bool`

GetHasUserData returns the HasUserData field if non-nil, zero value otherwise.

### GetHasUserDataOk

`func (o *InstanceBase) GetHasUserDataOk() (*bool, bool)`

GetHasUserDataOk returns a tuple with the HasUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUserData

`func (o *InstanceBase) SetHasUserData(v bool)`

SetHasUserData sets HasUserData field to given value.


### GetRootDiskSize

`func (o *InstanceBase) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *InstanceBase) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *InstanceBase) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.


### GetRootDiskStorageType

`func (o *InstanceBase) GetRootDiskStorageType() StorageType`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *InstanceBase) GetRootDiskStorageTypeOk() (*StorageType, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *InstanceBase) SetRootDiskStorageType(v StorageType)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.


### GetContract

`func (o *InstanceBase) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *InstanceBase) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *InstanceBase) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetAutoScalingGroup

`func (o *InstanceBase) GetAutoScalingGroup() AutoScalingGroup`

GetAutoScalingGroup returns the AutoScalingGroup field if non-nil, zero value otherwise.

### GetAutoScalingGroupOk

`func (o *InstanceBase) GetAutoScalingGroupOk() (*AutoScalingGroup, bool)`

GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroup

`func (o *InstanceBase) SetAutoScalingGroup(v AutoScalingGroup)`

SetAutoScalingGroup sets AutoScalingGroup field to given value.


### SetAutoScalingGroupNil

`func (o *InstanceBase) SetAutoScalingGroupNil(b bool)`

 SetAutoScalingGroupNil sets the value for AutoScalingGroup to be an explicit nil

### UnsetAutoScalingGroup
`func (o *InstanceBase) UnsetAutoScalingGroup()`

UnsetAutoScalingGroup ensures that no value is present for AutoScalingGroup, not even an explicit nil
### GetImage

`func (o *InstanceBase) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InstanceBase) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InstanceBase) SetImage(v Image)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


