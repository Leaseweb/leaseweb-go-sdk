# SchemasHealthCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**HealthCheckStatus**](HealthCheckStatus.md) |  | 
**Description** | **string** | Description of the health check status | 

## Methods

### NewSchemasHealthCheckStatus

`func NewSchemasHealthCheckStatus(state HealthCheckStatus, description string, ) *SchemasHealthCheckStatus`

NewSchemasHealthCheckStatus instantiates a new SchemasHealthCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasHealthCheckStatusWithDefaults

`func NewSchemasHealthCheckStatusWithDefaults() *SchemasHealthCheckStatus`

NewSchemasHealthCheckStatusWithDefaults instantiates a new SchemasHealthCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SchemasHealthCheckStatus) GetState() HealthCheckStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SchemasHealthCheckStatus) GetStateOk() (*HealthCheckStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SchemasHealthCheckStatus) SetState(v HealthCheckStatus)`

SetState sets State field to given value.


### GetDescription

`func (o *SchemasHealthCheckStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemasHealthCheckStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemasHealthCheckStatus) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


