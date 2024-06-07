# HealthCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthCheckStatus

`func NewHealthCheckStatus() *HealthCheckStatus`

NewHealthCheckStatus instantiates a new HealthCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckStatusWithDefaults

`func NewHealthCheckStatusWithDefaults() *HealthCheckStatus`

NewHealthCheckStatusWithDefaults instantiates a new HealthCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckStatus

`func (o *HealthCheckStatus) GetCheckStatus() string`

GetCheckStatus returns the CheckStatus field if non-nil, zero value otherwise.

### GetCheckStatusOk

`func (o *HealthCheckStatus) GetCheckStatusOk() (*string, bool)`

GetCheckStatusOk returns a tuple with the CheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatus

`func (o *HealthCheckStatus) SetCheckStatus(v string)`

SetCheckStatus sets CheckStatus field to given value.

### HasCheckStatus

`func (o *HealthCheckStatus) HasCheckStatus() bool`

HasCheckStatus returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheckStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthCheckStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


