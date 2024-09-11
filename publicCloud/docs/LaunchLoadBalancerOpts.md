# LaunchLoadBalancerOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [**RegionName**](RegionName.md) |  | 
**Type** | [**TypeName**](TypeName.md) |  | 
**Reference** | Pointer to **string** | An identifying name you can refer to the load balancer | [optional] 
**ContractType** | **string** | The contract applicable for the load balancer | 
**BillingFrequency** | **int32** | How often you wish to be charged. Used only when contract type is MONTHLY. &#39;1&#39; means every month, &#39;3&#39; every three months and so on. | 
**RootDiskStorageType** | [**StorageType**](StorageType.md) |  | 
**TargetPort** | **int32** | The port that the registered instances listen to | 

## Methods

### NewLaunchLoadBalancerOpts

`func NewLaunchLoadBalancerOpts(region RegionName, type_ TypeName, contractType string, billingFrequency int32, rootDiskStorageType StorageType, targetPort int32, ) *LaunchLoadBalancerOpts`

NewLaunchLoadBalancerOpts instantiates a new LaunchLoadBalancerOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchLoadBalancerOptsWithDefaults

`func NewLaunchLoadBalancerOptsWithDefaults() *LaunchLoadBalancerOpts`

NewLaunchLoadBalancerOptsWithDefaults instantiates a new LaunchLoadBalancerOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *LaunchLoadBalancerOpts) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LaunchLoadBalancerOpts) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LaunchLoadBalancerOpts) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetType

`func (o *LaunchLoadBalancerOpts) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LaunchLoadBalancerOpts) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LaunchLoadBalancerOpts) SetType(v TypeName)`

SetType sets Type field to given value.


### GetReference

`func (o *LaunchLoadBalancerOpts) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LaunchLoadBalancerOpts) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LaunchLoadBalancerOpts) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LaunchLoadBalancerOpts) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetContractType

`func (o *LaunchLoadBalancerOpts) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *LaunchLoadBalancerOpts) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *LaunchLoadBalancerOpts) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetBillingFrequency

`func (o *LaunchLoadBalancerOpts) GetBillingFrequency() int32`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *LaunchLoadBalancerOpts) GetBillingFrequencyOk() (*int32, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *LaunchLoadBalancerOpts) SetBillingFrequency(v int32)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetRootDiskStorageType

`func (o *LaunchLoadBalancerOpts) GetRootDiskStorageType() StorageType`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *LaunchLoadBalancerOpts) GetRootDiskStorageTypeOk() (*StorageType, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *LaunchLoadBalancerOpts) SetRootDiskStorageType(v StorageType)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.


### GetTargetPort

`func (o *LaunchLoadBalancerOpts) GetTargetPort() int32`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *LaunchLoadBalancerOpts) GetTargetPortOk() (*int32, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *LaunchLoadBalancerOpts) SetTargetPort(v int32)`

SetTargetPort sets TargetPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


