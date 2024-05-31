# LastClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayAgent** | Pointer to **NullableString** | The relay agent that forwarded the DHCP traffic | [optional] 
**Type** | Pointer to **string** | The type of DHCP packet requested by the client | [optional] 
**UserAgent** | Pointer to **string** | The user agent of the client making the DHCP request | [optional] 

## Methods

### NewLastClientRequest

`func NewLastClientRequest() *LastClientRequest`

NewLastClientRequest instantiates a new LastClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastClientRequestWithDefaults

`func NewLastClientRequestWithDefaults() *LastClientRequest`

NewLastClientRequestWithDefaults instantiates a new LastClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayAgent

`func (o *LastClientRequest) GetRelayAgent() string`

GetRelayAgent returns the RelayAgent field if non-nil, zero value otherwise.

### GetRelayAgentOk

`func (o *LastClientRequest) GetRelayAgentOk() (*string, bool)`

GetRelayAgentOk returns a tuple with the RelayAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAgent

`func (o *LastClientRequest) SetRelayAgent(v string)`

SetRelayAgent sets RelayAgent field to given value.

### HasRelayAgent

`func (o *LastClientRequest) HasRelayAgent() bool`

HasRelayAgent returns a boolean if a field has been set.

### SetRelayAgentNil

`func (o *LastClientRequest) SetRelayAgentNil(b bool)`

 SetRelayAgentNil sets the value for RelayAgent to be an explicit nil

### UnsetRelayAgent
`func (o *LastClientRequest) UnsetRelayAgent()`

UnsetRelayAgent ensures that no value is present for RelayAgent, not even an explicit nil
### GetType

`func (o *LastClientRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LastClientRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LastClientRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LastClientRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserAgent

`func (o *LastClientRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *LastClientRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *LastClientRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *LastClientRequest) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


