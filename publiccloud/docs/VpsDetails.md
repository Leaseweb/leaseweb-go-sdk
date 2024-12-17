# VpsDetails

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
**MarketAppId** | Pointer to [**MarketAppId**](MarketAppId.md) |  | [optional] 
**State** | [**VpsState**](VpsState.md) |  | 
**HasPublicIpV4** | **bool** |  | 
**RootDiskSize** | **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | 
**StartedAt** | **NullableTime** | Date and time when the VPS was started for the first time, right after launching it | 
**Contract** | [**VpsContract**](VpsContract.md) |  | 
**Iso** | Pointer to [**NullableIso**](Iso.md) |  | [optional] 
**Ips** | [**[]IpDetails**](IpDetails.md) |  | 

## Methods

### NewVpsDetails

`func NewVpsDetails(id string, pack VpsPackType, resources BaseResources, region RegionName, datacenter Datacenter, reference string, image Image, state VpsState, hasPublicIpV4 bool, rootDiskSize int32, startedAt NullableTime, contract VpsContract, ips []IpDetails, ) *VpsDetails`

NewVpsDetails instantiates a new VpsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsDetailsWithDefaults

`func NewVpsDetailsWithDefaults() *VpsDetails`

NewVpsDetailsWithDefaults instantiates a new VpsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpsDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpsDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpsDetails) SetId(v string)`

SetId sets Id field to given value.


### GetPack

`func (o *VpsDetails) GetPack() VpsPackType`

GetPack returns the Pack field if non-nil, zero value otherwise.

### GetPackOk

`func (o *VpsDetails) GetPackOk() (*VpsPackType, bool)`

GetPackOk returns a tuple with the Pack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPack

`func (o *VpsDetails) SetPack(v VpsPackType)`

SetPack sets Pack field to given value.


### GetResources

`func (o *VpsDetails) GetResources() BaseResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VpsDetails) GetResourcesOk() (*BaseResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VpsDetails) SetResources(v BaseResources)`

SetResources sets Resources field to given value.


### GetRegion

`func (o *VpsDetails) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VpsDetails) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VpsDetails) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetDatacenter

`func (o *VpsDetails) GetDatacenter() Datacenter`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VpsDetails) GetDatacenterOk() (*Datacenter, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VpsDetails) SetDatacenter(v Datacenter)`

SetDatacenter sets Datacenter field to given value.


### GetReference

`func (o *VpsDetails) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *VpsDetails) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *VpsDetails) SetReference(v string)`

SetReference sets Reference field to given value.


### GetImage

`func (o *VpsDetails) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VpsDetails) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VpsDetails) SetImage(v Image)`

SetImage sets Image field to given value.


### GetMarketAppId

`func (o *VpsDetails) GetMarketAppId() MarketAppId`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *VpsDetails) GetMarketAppIdOk() (*MarketAppId, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *VpsDetails) SetMarketAppId(v MarketAppId)`

SetMarketAppId sets MarketAppId field to given value.

### HasMarketAppId

`func (o *VpsDetails) HasMarketAppId() bool`

HasMarketAppId returns a boolean if a field has been set.

### GetState

`func (o *VpsDetails) GetState() VpsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpsDetails) GetStateOk() (*VpsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpsDetails) SetState(v VpsState)`

SetState sets State field to given value.


### GetHasPublicIpV4

`func (o *VpsDetails) GetHasPublicIpV4() bool`

GetHasPublicIpV4 returns the HasPublicIpV4 field if non-nil, zero value otherwise.

### GetHasPublicIpV4Ok

`func (o *VpsDetails) GetHasPublicIpV4Ok() (*bool, bool)`

GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIpV4

`func (o *VpsDetails) SetHasPublicIpV4(v bool)`

SetHasPublicIpV4 sets HasPublicIpV4 field to given value.


### GetRootDiskSize

`func (o *VpsDetails) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *VpsDetails) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *VpsDetails) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.


### GetStartedAt

`func (o *VpsDetails) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *VpsDetails) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *VpsDetails) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *VpsDetails) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *VpsDetails) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetContract

`func (o *VpsDetails) GetContract() VpsContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *VpsDetails) GetContractOk() (*VpsContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *VpsDetails) SetContract(v VpsContract)`

SetContract sets Contract field to given value.


### GetIso

`func (o *VpsDetails) GetIso() Iso`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *VpsDetails) GetIsoOk() (*Iso, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *VpsDetails) SetIso(v Iso)`

SetIso sets Iso field to given value.

### HasIso

`func (o *VpsDetails) HasIso() bool`

HasIso returns a boolean if a field has been set.

### SetIsoNil

`func (o *VpsDetails) SetIsoNil(b bool)`

 SetIsoNil sets the value for Iso to be an explicit nil

### UnsetIso
`func (o *VpsDetails) UnsetIso()`

UnsetIso ensures that no value is present for Iso, not even an explicit nil
### GetIps

`func (o *VpsDetails) GetIps() []IpDetails`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *VpsDetails) GetIpsOk() (*[]IpDetails, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *VpsDetails) SetIps(v []IpDetails)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


