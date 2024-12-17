# VpsBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Pack** | [**VpsPackType**](VpsPackType.md) |  | 
**Resources** | [**BaseResources**](BaseResources.md) |  | 
**Region** | [**RegionName**](RegionName.md) |  | 
**Datacenter** | [**Datacenter**](Datacenter.md) |  | 
**Reference** | **string** | The identifying name set to the instance | 
**Image** | [**Image**](Image.md) |  | 
**MarketAppId** | Pointer to [**NullableMarketAppId**](MarketAppId.md) |  | [optional] 
**State** | [**VpsState**](VpsState.md) |  | 
**HasPublicIpV4** | **bool** |  | 
**RootDiskSize** | **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | 
**StartedAt** | **NullableTime** | Date and time when the VPS was started for the first time, right after launching it | 
**Contract** | [**VpsContract**](VpsContract.md) |  | 

## Methods

### NewVpsBase

`func NewVpsBase(id string, pack VpsPackType, resources BaseResources, region RegionName, datacenter Datacenter, reference string, image Image, state VpsState, hasPublicIpV4 bool, rootDiskSize int32, startedAt NullableTime, contract VpsContract, ) *VpsBase`

NewVpsBase instantiates a new VpsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsBaseWithDefaults

`func NewVpsBaseWithDefaults() *VpsBase`

NewVpsBaseWithDefaults instantiates a new VpsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpsBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpsBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpsBase) SetId(v string)`

SetId sets Id field to given value.


### GetPack

`func (o *VpsBase) GetPack() VpsPackType`

GetPack returns the Pack field if non-nil, zero value otherwise.

### GetPackOk

`func (o *VpsBase) GetPackOk() (*VpsPackType, bool)`

GetPackOk returns a tuple with the Pack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPack

`func (o *VpsBase) SetPack(v VpsPackType)`

SetPack sets Pack field to given value.


### GetResources

`func (o *VpsBase) GetResources() BaseResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VpsBase) GetResourcesOk() (*BaseResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VpsBase) SetResources(v BaseResources)`

SetResources sets Resources field to given value.


### GetRegion

`func (o *VpsBase) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VpsBase) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VpsBase) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetDatacenter

`func (o *VpsBase) GetDatacenter() Datacenter`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VpsBase) GetDatacenterOk() (*Datacenter, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VpsBase) SetDatacenter(v Datacenter)`

SetDatacenter sets Datacenter field to given value.


### GetReference

`func (o *VpsBase) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *VpsBase) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *VpsBase) SetReference(v string)`

SetReference sets Reference field to given value.


### GetImage

`func (o *VpsBase) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VpsBase) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VpsBase) SetImage(v Image)`

SetImage sets Image field to given value.


### GetMarketAppId

`func (o *VpsBase) GetMarketAppId() MarketAppId`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *VpsBase) GetMarketAppIdOk() (*MarketAppId, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *VpsBase) SetMarketAppId(v MarketAppId)`

SetMarketAppId sets MarketAppId field to given value.

### HasMarketAppId

`func (o *VpsBase) HasMarketAppId() bool`

HasMarketAppId returns a boolean if a field has been set.

### SetMarketAppIdNil

`func (o *VpsBase) SetMarketAppIdNil(b bool)`

 SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil

### UnsetMarketAppId
`func (o *VpsBase) UnsetMarketAppId()`

UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil
### GetState

`func (o *VpsBase) GetState() VpsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpsBase) GetStateOk() (*VpsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpsBase) SetState(v VpsState)`

SetState sets State field to given value.


### GetHasPublicIpV4

`func (o *VpsBase) GetHasPublicIpV4() bool`

GetHasPublicIpV4 returns the HasPublicIpV4 field if non-nil, zero value otherwise.

### GetHasPublicIpV4Ok

`func (o *VpsBase) GetHasPublicIpV4Ok() (*bool, bool)`

GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIpV4

`func (o *VpsBase) SetHasPublicIpV4(v bool)`

SetHasPublicIpV4 sets HasPublicIpV4 field to given value.


### GetRootDiskSize

`func (o *VpsBase) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *VpsBase) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *VpsBase) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.


### GetStartedAt

`func (o *VpsBase) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *VpsBase) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *VpsBase) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *VpsBase) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *VpsBase) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetContract

`func (o *VpsBase) GetContract() VpsContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *VpsBase) GetContractOk() (*VpsContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *VpsBase) SetContract(v VpsContract)`

SetContract sets Contract field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


