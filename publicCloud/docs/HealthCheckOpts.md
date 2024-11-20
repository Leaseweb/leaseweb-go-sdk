# HealthCheckOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Method** | Pointer to [**HttpMethodOpt**](HttpMethodOpt.md) |  | [optional] 
**Uri** | **string** | URI to check in the target instances | 
**Host** | Pointer to **string** | Host for the health check if any | [optional] 
**Port** | **int32** | Port number | 

## Methods

### NewHealthCheckOpts

`func NewHealthCheckOpts(protocol Protocol, uri string, port int32, ) *HealthCheckOpts`

NewHealthCheckOpts instantiates a new HealthCheckOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckOptsWithDefaults

`func NewHealthCheckOptsWithDefaults() *HealthCheckOpts`

NewHealthCheckOptsWithDefaults instantiates a new HealthCheckOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *HealthCheckOpts) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheckOpts) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HealthCheckOpts) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetMethod

`func (o *HealthCheckOpts) GetMethod() HttpMethodOpt`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HealthCheckOpts) GetMethodOk() (*HttpMethodOpt, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HealthCheckOpts) SetMethod(v HttpMethodOpt)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HealthCheckOpts) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *HealthCheckOpts) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *HealthCheckOpts) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *HealthCheckOpts) SetUri(v string)`

SetUri sets Uri field to given value.


### GetHost

`func (o *HealthCheckOpts) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HealthCheckOpts) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HealthCheckOpts) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HealthCheckOpts) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *HealthCheckOpts) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheckOpts) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HealthCheckOpts) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


