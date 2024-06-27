# OperatingSystemDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**OperatingSystemId**](OperatingSystemId.md) |  | 
**Name** | **string** |  | 
**Version** | **string** |  | 
**Family** | **string** |  | 
**Flavour** | **string** |  | 
**Architecture** | **string** |  | 
**MarketApps** | **[]string** |  | 
**StorageTypes** | **[]string** | The supported storage types for the instance type | 

## Methods

### NewOperatingSystemDetail

`func NewOperatingSystemDetail(id OperatingSystemId, name string, version string, family string, flavour string, architecture string, marketApps []string, storageTypes []string, ) *OperatingSystemDetail`

NewOperatingSystemDetail instantiates a new OperatingSystemDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemDetailWithDefaults

`func NewOperatingSystemDetailWithDefaults() *OperatingSystemDetail`

NewOperatingSystemDetailWithDefaults instantiates a new OperatingSystemDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatingSystemDetail) GetId() OperatingSystemId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatingSystemDetail) GetIdOk() (*OperatingSystemId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatingSystemDetail) SetId(v OperatingSystemId)`

SetId sets Id field to given value.


### GetName

`func (o *OperatingSystemDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatingSystemDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatingSystemDetail) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *OperatingSystemDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OperatingSystemDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OperatingSystemDetail) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFamily

`func (o *OperatingSystemDetail) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *OperatingSystemDetail) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *OperatingSystemDetail) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetFlavour

`func (o *OperatingSystemDetail) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *OperatingSystemDetail) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *OperatingSystemDetail) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.


### GetArchitecture

`func (o *OperatingSystemDetail) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *OperatingSystemDetail) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *OperatingSystemDetail) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetMarketApps

`func (o *OperatingSystemDetail) GetMarketApps() []string`

GetMarketApps returns the MarketApps field if non-nil, zero value otherwise.

### GetMarketAppsOk

`func (o *OperatingSystemDetail) GetMarketAppsOk() (*[]string, bool)`

GetMarketAppsOk returns a tuple with the MarketApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketApps

`func (o *OperatingSystemDetail) SetMarketApps(v []string)`

SetMarketApps sets MarketApps field to given value.


### GetStorageTypes

`func (o *OperatingSystemDetail) GetStorageTypes() []string`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *OperatingSystemDetail) GetStorageTypesOk() (*[]string, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *OperatingSystemDetail) SetStorageTypes(v []string)`

SetStorageTypes sets StorageTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


