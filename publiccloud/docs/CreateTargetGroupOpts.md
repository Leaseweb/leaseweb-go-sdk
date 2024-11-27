# CreateTargetGroupOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the target group | 
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Port** | **int32** | The port of the target group | 
**Region** | [**RegionName**](RegionName.md) |  | 
**HealthCheck** | Pointer to [**HealthCheckOpts**](HealthCheckOpts.md) |  | [optional] 

## Methods

### NewCreateTargetGroupOpts

`func NewCreateTargetGroupOpts(name string, protocol Protocol, port int32, region RegionName, ) *CreateTargetGroupOpts`

NewCreateTargetGroupOpts instantiates a new CreateTargetGroupOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTargetGroupOptsWithDefaults

`func NewCreateTargetGroupOptsWithDefaults() *CreateTargetGroupOpts`

NewCreateTargetGroupOptsWithDefaults instantiates a new CreateTargetGroupOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTargetGroupOpts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTargetGroupOpts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTargetGroupOpts) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *CreateTargetGroupOpts) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateTargetGroupOpts) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateTargetGroupOpts) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *CreateTargetGroupOpts) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateTargetGroupOpts) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateTargetGroupOpts) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRegion

`func (o *CreateTargetGroupOpts) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateTargetGroupOpts) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateTargetGroupOpts) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetHealthCheck

`func (o *CreateTargetGroupOpts) GetHealthCheck() HealthCheckOpts`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *CreateTargetGroupOpts) GetHealthCheckOk() (*HealthCheckOpts, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *CreateTargetGroupOpts) SetHealthCheck(v HealthCheckOpts)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *CreateTargetGroupOpts) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


