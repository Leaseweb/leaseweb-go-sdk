# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the target | 
**Reference** | **string** | The reference of the target | 
**Image** | [**Image**](Image.md) |  | 
**State** | **string** | The state of the target | 
**Ips** | [**[]Ip**](Ip.md) | The IP addresses of the target | 
**HealthCheck** | [**NullableSchemasHealthCheckStatus**](SchemasHealthCheckStatus.md) |  | 

## Methods

### NewTarget

`func NewTarget(id string, reference string, image Image, state string, ips []Ip, healthCheck NullableSchemasHealthCheckStatus, ) *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Target) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Target) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Target) SetId(v string)`

SetId sets Id field to given value.


### GetReference

`func (o *Target) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Target) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Target) SetReference(v string)`

SetReference sets Reference field to given value.


### GetImage

`func (o *Target) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Target) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Target) SetImage(v Image)`

SetImage sets Image field to given value.


### GetState

`func (o *Target) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Target) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Target) SetState(v string)`

SetState sets State field to given value.


### GetIps

`func (o *Target) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Target) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Target) SetIps(v []Ip)`

SetIps sets Ips field to given value.


### GetHealthCheck

`func (o *Target) GetHealthCheck() SchemasHealthCheckStatus`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *Target) GetHealthCheckOk() (*SchemasHealthCheckStatus, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *Target) SetHealthCheck(v SchemasHealthCheckStatus)`

SetHealthCheck sets HealthCheck field to given value.


### SetHealthCheckNil

`func (o *Target) SetHealthCheckNil(b bool)`

 SetHealthCheckNil sets the value for HealthCheck to be an explicit nil

### UnsetHealthCheck
`func (o *Target) UnsetHealthCheck()`

UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


