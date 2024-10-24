# UpdateLoadBalancerOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeName**](TypeName.md) |  | [optional] 
**Reference** | Pointer to **string** | An identifying name you can refer to the load balancer | [optional] 
**ContractType** | Pointer to [**ContractType**](ContractType.md) |  | [optional] 
**StickySession** | Pointer to [**NullableStickySession**](StickySession.md) |  | [optional] 
**Balance** | Pointer to [**Balance**](Balance.md) |  | [optional] 
**XForwardedFor** | Pointer to **bool** | Is xForwardedFor enabled or not | [optional] 
**IdleTimeOut** | Pointer to **int32** | Time to close the connection if load balancer is idle | [optional] 

## Methods

### NewUpdateLoadBalancerOpts

`func NewUpdateLoadBalancerOpts() *UpdateLoadBalancerOpts`

NewUpdateLoadBalancerOpts instantiates a new UpdateLoadBalancerOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerOptsWithDefaults

`func NewUpdateLoadBalancerOptsWithDefaults() *UpdateLoadBalancerOpts`

NewUpdateLoadBalancerOptsWithDefaults instantiates a new UpdateLoadBalancerOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateLoadBalancerOpts) GetType() TypeName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateLoadBalancerOpts) GetTypeOk() (*TypeName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateLoadBalancerOpts) SetType(v TypeName)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateLoadBalancerOpts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReference

`func (o *UpdateLoadBalancerOpts) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UpdateLoadBalancerOpts) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UpdateLoadBalancerOpts) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UpdateLoadBalancerOpts) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetContractType

`func (o *UpdateLoadBalancerOpts) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *UpdateLoadBalancerOpts) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *UpdateLoadBalancerOpts) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *UpdateLoadBalancerOpts) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetStickySession

`func (o *UpdateLoadBalancerOpts) GetStickySession() StickySession`

GetStickySession returns the StickySession field if non-nil, zero value otherwise.

### GetStickySessionOk

`func (o *UpdateLoadBalancerOpts) GetStickySessionOk() (*StickySession, bool)`

GetStickySessionOk returns a tuple with the StickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySession

`func (o *UpdateLoadBalancerOpts) SetStickySession(v StickySession)`

SetStickySession sets StickySession field to given value.

### HasStickySession

`func (o *UpdateLoadBalancerOpts) HasStickySession() bool`

HasStickySession returns a boolean if a field has been set.

### SetStickySessionNil

`func (o *UpdateLoadBalancerOpts) SetStickySessionNil(b bool)`

 SetStickySessionNil sets the value for StickySession to be an explicit nil

### UnsetStickySession
`func (o *UpdateLoadBalancerOpts) UnsetStickySession()`

UnsetStickySession ensures that no value is present for StickySession, not even an explicit nil
### GetBalance

`func (o *UpdateLoadBalancerOpts) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *UpdateLoadBalancerOpts) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *UpdateLoadBalancerOpts) SetBalance(v Balance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *UpdateLoadBalancerOpts) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetXForwardedFor

`func (o *UpdateLoadBalancerOpts) GetXForwardedFor() bool`

GetXForwardedFor returns the XForwardedFor field if non-nil, zero value otherwise.

### GetXForwardedForOk

`func (o *UpdateLoadBalancerOpts) GetXForwardedForOk() (*bool, bool)`

GetXForwardedForOk returns a tuple with the XForwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedFor

`func (o *UpdateLoadBalancerOpts) SetXForwardedFor(v bool)`

SetXForwardedFor sets XForwardedFor field to given value.

### HasXForwardedFor

`func (o *UpdateLoadBalancerOpts) HasXForwardedFor() bool`

HasXForwardedFor returns a boolean if a field has been set.

### GetIdleTimeOut

`func (o *UpdateLoadBalancerOpts) GetIdleTimeOut() int32`

GetIdleTimeOut returns the IdleTimeOut field if non-nil, zero value otherwise.

### GetIdleTimeOutOk

`func (o *UpdateLoadBalancerOpts) GetIdleTimeOutOk() (*int32, bool)`

GetIdleTimeOutOk returns a tuple with the IdleTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeOut

`func (o *UpdateLoadBalancerOpts) SetIdleTimeOut(v int32)`

SetIdleTimeOut sets IdleTimeOut field to given value.

### HasIdleTimeOut

`func (o *UpdateLoadBalancerOpts) HasIdleTimeOut() bool`

HasIdleTimeOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


