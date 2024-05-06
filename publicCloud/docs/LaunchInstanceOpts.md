# LaunchInstanceOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region to launch the instance into | 
**Type** | Pointer to **string** | Instance type | [optional] 
**OperatingSystemId** | **string** | Operating System ID | 
**MarketAppId** | Pointer to **NullableString** | Market App ID that must be installed into the instance | [optional] 
**Reference** | Pointer to **string** | An identifying name you can refer to the instance | [optional] 
**ContractType** | **string** |  | 
**ContractTerm** | **int32** | Contract commitment. Used only when contract type is MONTHLY | 
**BillingFrequency** | **int32** | How often you wish to be charged. Used only when contract type is MONTHLY. &#39;1&#39; means every month, &#39;3&#39; every three months and so on. | 
**RootDiskSize** | Pointer to **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | [optional] 
**RootDiskStorageType** | **string** | The root disk&#39;s storage type | 
**SshKey** | Pointer to **string** | Public SSH key to be installed into the instance. Must be used only on Linux/FreeBSD instances | [optional] 

## Methods

### NewLaunchInstanceOpts

`func NewLaunchInstanceOpts(region string, operatingSystemId string, contractType string, contractTerm int32, billingFrequency int32, rootDiskStorageType string, ) *LaunchInstanceOpts`

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

`func (o *LaunchInstanceOpts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LaunchInstanceOpts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LaunchInstanceOpts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LaunchInstanceOpts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOperatingSystemId

`func (o *LaunchInstanceOpts) GetOperatingSystemId() string`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *LaunchInstanceOpts) GetOperatingSystemIdOk() (*string, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *LaunchInstanceOpts) SetOperatingSystemId(v string)`

SetOperatingSystemId sets OperatingSystemId field to given value.


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

### SetMarketAppIdNil

`func (o *LaunchInstanceOpts) SetMarketAppIdNil(b bool)`

 SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil

### UnsetMarketAppId
`func (o *LaunchInstanceOpts) UnsetMarketAppId()`

UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil
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

`func (o *LaunchInstanceOpts) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *LaunchInstanceOpts) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *LaunchInstanceOpts) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetContractTerm

`func (o *LaunchInstanceOpts) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *LaunchInstanceOpts) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *LaunchInstanceOpts) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.


### GetBillingFrequency

`func (o *LaunchInstanceOpts) GetBillingFrequency() int32`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *LaunchInstanceOpts) GetBillingFrequencyOk() (*int32, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *LaunchInstanceOpts) SetBillingFrequency(v int32)`

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

`func (o *LaunchInstanceOpts) GetRootDiskStorageType() string`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *LaunchInstanceOpts) GetRootDiskStorageTypeOk() (*string, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *LaunchInstanceOpts) SetRootDiskStorageType(v string)`

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


