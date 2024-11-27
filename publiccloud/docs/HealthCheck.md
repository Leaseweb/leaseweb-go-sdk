# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Method** | [**NullableHttpMethod**](HttpMethod.md) |  | 
**Uri** | **string** | URI to check in the target instances | 
**Host** | **NullableString** | Host for the health check if any | 
**Port** | **int32** | Port number | 

## Methods

### NewHealthCheck

`func NewHealthCheck(protocol Protocol, method NullableHttpMethod, uri string, host NullableString, port int32, ) *HealthCheck`

NewHealthCheck instantiates a new HealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckWithDefaults

`func NewHealthCheckWithDefaults() *HealthCheck`

NewHealthCheckWithDefaults instantiates a new HealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *HealthCheck) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheck) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HealthCheck) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetMethod

`func (o *HealthCheck) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HealthCheck) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HealthCheck) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.


### SetMethodNil

`func (o *HealthCheck) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *HealthCheck) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetUri

`func (o *HealthCheck) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *HealthCheck) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *HealthCheck) SetUri(v string)`

SetUri sets Uri field to given value.


### GetHost

`func (o *HealthCheck) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HealthCheck) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HealthCheck) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *HealthCheck) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HealthCheck) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HealthCheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HealthCheck) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


