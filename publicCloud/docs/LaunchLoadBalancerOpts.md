# LaunchLoadBalancerOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [**RegionName**](RegionName.md) |  | 
**Type** | [**TypeName**](TypeName.md) |  | 
**Reference** | Pointer to **string** | An identifying name you can refer to the load balancer | [optional] 
**ContractType** | [**ContractType**](ContractType.md) |  | 
**ContractTerm** | [**ContractTerm**](ContractTerm.md) |  | 
**BillingFrequency** | [**BillingFrequency**](BillingFrequency.md) |  | 

## Methods

### NewLaunchLoadBalancerOpts

`func NewLaunchLoadBalancerOpts(region RegionName, type_ TypeName, contractType ContractType, contractTerm ContractTerm, billingFrequency BillingFrequency, ) *LaunchLoadBalancerOpts`

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

`func (o *LaunchLoadBalancerOpts) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *LaunchLoadBalancerOpts) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *LaunchLoadBalancerOpts) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.


### GetContractTerm

`func (o *LaunchLoadBalancerOpts) GetContractTerm() ContractTerm`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *LaunchLoadBalancerOpts) GetContractTermOk() (*ContractTerm, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *LaunchLoadBalancerOpts) SetContractTerm(v ContractTerm)`

SetContractTerm sets ContractTerm field to given value.


### GetBillingFrequency

`func (o *LaunchLoadBalancerOpts) GetBillingFrequency() BillingFrequency`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *LaunchLoadBalancerOpts) GetBillingFrequencyOk() (*BillingFrequency, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *LaunchLoadBalancerOpts) SetBillingFrequency(v BillingFrequency)`

SetBillingFrequency sets BillingFrequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


