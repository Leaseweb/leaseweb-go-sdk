# TargetGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | The name of the target group | 
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Port** | **int32** | The port of the target group | 
**Region** | [**RegionName**](RegionName.md) |  | 
**HealthCheck** | [**NullableHealthCheck**](HealthCheck.md) |  | 

## Methods

### NewTargetGroup

`func NewTargetGroup(id string, name string, protocol Protocol, port int32, region RegionName, healthCheck NullableHealthCheck, ) *TargetGroup`

NewTargetGroup instantiates a new TargetGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupWithDefaults

`func NewTargetGroupWithDefaults() *TargetGroup`

NewTargetGroupWithDefaults instantiates a new TargetGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TargetGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TargetGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *TargetGroup) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *TargetGroup) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *TargetGroup) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *TargetGroup) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetGroup) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetGroup) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRegion

`func (o *TargetGroup) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TargetGroup) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TargetGroup) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetHealthCheck

`func (o *TargetGroup) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *TargetGroup) GetHealthCheckOk() (*HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *TargetGroup) SetHealthCheck(v HealthCheck)`

SetHealthCheck sets HealthCheck field to given value.


### SetHealthCheckNil

`func (o *TargetGroup) SetHealthCheckNil(b bool)`

 SetHealthCheckNil sets the value for HealthCheck to be an explicit nil

### UnsetHealthCheck
`func (o *TargetGroup) UnsetHealthCheck()`

UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


