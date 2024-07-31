# LaunchInstanceOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region to launch the instance into | 
**Type** | [**TypeName**](TypeName.md) |  | 
**ImageId** | **string** | imageId can be either an Operating System or a UUID in case of a Custom Image | 
**MarketAppId** | Pointer to **string** | Market App ID that must be installed into the instance | [optional] 
**Reference** | Pointer to **string** | An identifying name you can refer to the instance | [optional] 
**ContractType** | [**ContractType**](ContractType.md) |  | 
**ContractTerm** | [**ContractTerm**](ContractTerm.md) |  | 
**BillingFrequency** | [**BillingFrequency**](BillingFrequency.md) |  | 
**RootDiskSize** | Pointer to **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | [optional] 
**RootDiskStorageType** | [**RootDiskStorageType**](RootDiskStorageType.md) |  | 
**SshKey** | Pointer to **string** | Public SSH key to be installed into the instance. Must be used only on Linux/FreeBSD instances | [optional] 

## Methods

### NewLaunchInstanceOpts

`func NewLaunchInstanceOpts(region string, type_ TypeName, imageId string, contractType ContractType, contractTerm ContractTerm, billingFrequency BillingFrequency, rootDiskStorageType RootDiskStorageType, ) *LaunchInstanceOpts`

NewLaunchInstanceOpts instantiates a new LaunchInstanceOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchInstanceOptsWithDefaults

`func NewLaunchInstanceOptsWithDefaults() *LaunchInstanceOpts`

NewLaunchInstanceOptsWithDefaults instantiates a new LaunchInstanceOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *LaunchInstanceOpts) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LaunchInstanceOpts) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LaunchInstanceOpts) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetType

`func (o *LaunchInstanceOpts) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LaunchInstanceOpts) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LaunchInstanceOpts) SetType(v TypeName)`

SetType sets Type field to given value.


### GetImageId

`func (o *LaunchInstanceOpts) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *LaunchInstanceOpts) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *LaunchInstanceOpts) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetMarketAppId

`func (o *LaunchInstanceOpts) GetMarketAppId() string`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *LaunchInstanceOpts) GetMarketAppIdOk() (*string, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *LaunchInstanceOpts) SetMarketAppId(v string)`

SetMarketAppId sets MarketAppId field to given value.

### HasMarketAppId

`func (o *LaunchInstanceOpts) HasMarketAppId() bool`

HasMarketAppId returns a boolean if a field has been set.

### GetReference

`func (o *LaunchInstanceOpts) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LaunchInstanceOpts) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LaunchInstanceOpts) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LaunchInstanceOpts) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetContractType

`func (o *LaunchInstanceOpts) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *LaunchInstanceOpts) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *LaunchInstanceOpts) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.


### GetContractTerm

`func (o *LaunchInstanceOpts) GetContractTerm() ContractTerm`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *LaunchInstanceOpts) GetContractTermOk() (*ContractTerm, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *LaunchInstanceOpts) SetContractTerm(v ContractTerm)`

SetContractTerm sets ContractTerm field to given value.


### GetBillingFrequency

`func (o *LaunchInstanceOpts) GetBillingFrequency() BillingFrequency`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *LaunchInstanceOpts) GetBillingFrequencyOk() (*BillingFrequency, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *LaunchInstanceOpts) SetBillingFrequency(v BillingFrequency)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetRootDiskSize

`func (o *LaunchInstanceOpts) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *LaunchInstanceOpts) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *LaunchInstanceOpts) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *LaunchInstanceOpts) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetRootDiskStorageType

`func (o *LaunchInstanceOpts) GetRootDiskStorageType() RootDiskStorageType`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *LaunchInstanceOpts) GetRootDiskStorageTypeOk() (*RootDiskStorageType, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *LaunchInstanceOpts) SetRootDiskStorageType(v RootDiskStorageType)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.


### GetSshKey

`func (o *LaunchInstanceOpts) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *LaunchInstanceOpts) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *LaunchInstanceOpts) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *LaunchInstanceOpts) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


