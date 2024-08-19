# AutoScaledInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the instance | [optional] 
**Type** | Pointer to **string** | The instance type, which determines the amount of resources | [optional] 
**Resources** | Pointer to [**Resources**](Resources.md) |  | [optional] 
**Region** | Pointer to [**RegionName**](RegionName.md) |  | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the instance | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**IncludesPrivateNetwork** | Pointer to **bool** |  | [optional] 
**Ips** | Pointer to [**[]IpDetails**](IpDetails.md) |  | [optional] 

## Methods

### NewAutoScaledInstance

`func NewAutoScaledInstance() *AutoScaledInstance`

NewAutoScaledInstance instantiates a new AutoScaledInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScaledInstanceWithDefaults

`func NewAutoScaledInstanceWithDefaults() *AutoScaledInstance`

NewAutoScaledInstanceWithDefaults instantiates a new AutoScaledInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoScaledInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScaledInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScaledInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoScaledInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AutoScaledInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoScaledInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoScaledInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoScaledInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResources

`func (o *AutoScaledInstance) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AutoScaledInstance) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AutoScaledInstance) SetResources(v Resources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AutoScaledInstance) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRegion

`func (o *AutoScaledInstance) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AutoScaledInstance) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AutoScaledInstance) SetRegion(v RegionName)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AutoScaledInstance) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *AutoScaledInstance) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AutoScaledInstance) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AutoScaledInstance) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AutoScaledInstance) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetState

`func (o *AutoScaledInstance) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScaledInstance) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScaledInstance) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *AutoScaledInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProductType

`func (o *AutoScaledInstance) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AutoScaledInstance) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AutoScaledInstance) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AutoScaledInstance) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetIncludesPrivateNetwork

`func (o *AutoScaledInstance) GetIncludesPrivateNetwork() bool`

GetIncludesPrivateNetwork returns the IncludesPrivateNetwork field if non-nil, zero value otherwise.

### GetIncludesPrivateNetworkOk

`func (o *AutoScaledInstance) GetIncludesPrivateNetworkOk() (*bool, bool)`

GetIncludesPrivateNetworkOk returns a tuple with the IncludesPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPrivateNetwork

`func (o *AutoScaledInstance) SetIncludesPrivateNetwork(v bool)`

SetIncludesPrivateNetwork sets IncludesPrivateNetwork field to given value.

### HasIncludesPrivateNetwork

`func (o *AutoScaledInstance) HasIncludesPrivateNetwork() bool`

HasIncludesPrivateNetwork returns a boolean if a field has been set.

### GetIps

`func (o *AutoScaledInstance) GetIps() []IpDetails`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *AutoScaledInstance) GetIpsOk() (*[]IpDetails, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *AutoScaledInstance) SetIps(v []IpDetails)`

SetIps sets Ips field to given value.

### HasIps

`func (o *AutoScaledInstance) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


