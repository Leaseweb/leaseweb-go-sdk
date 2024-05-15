# GetServerHardwareResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of this hardware scan result | [optional] 
**ParserVersion** | Pointer to **string** | Version of the parser used for this hardware info | [optional] 
**Result** | Pointer to [**Result**](Result.md) |  | [optional] 
**ScannedAt** | Pointer to **time.Time** | Timestamp of hardware scan, the ISO-8601 format | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetServerHardwareResult

`func NewGetServerHardwareResult() *GetServerHardwareResult`

NewGetServerHardwareResult instantiates a new GetServerHardwareResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerHardwareResultWithDefaults

`func NewGetServerHardwareResultWithDefaults() *GetServerHardwareResult`

NewGetServerHardwareResultWithDefaults instantiates a new GetServerHardwareResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetServerHardwareResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServerHardwareResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServerHardwareResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetServerHardwareResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParserVersion

`func (o *GetServerHardwareResult) GetParserVersion() string`

GetParserVersion returns the ParserVersion field if non-nil, zero value otherwise.

### GetParserVersionOk

`func (o *GetServerHardwareResult) GetParserVersionOk() (*string, bool)`

GetParserVersionOk returns a tuple with the ParserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserVersion

`func (o *GetServerHardwareResult) SetParserVersion(v string)`

SetParserVersion sets ParserVersion field to given value.

### HasParserVersion

`func (o *GetServerHardwareResult) HasParserVersion() bool`

HasParserVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetServerHardwareResult) GetResult() Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetServerHardwareResult) GetResultOk() (*Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetServerHardwareResult) SetResult(v Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetServerHardwareResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScannedAt

`func (o *GetServerHardwareResult) GetScannedAt() time.Time`

GetScannedAt returns the ScannedAt field if non-nil, zero value otherwise.

### GetScannedAtOk

`func (o *GetServerHardwareResult) GetScannedAtOk() (*time.Time, bool)`

GetScannedAtOk returns a tuple with the ScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedAt

`func (o *GetServerHardwareResult) SetScannedAt(v time.Time)`

SetScannedAt sets ScannedAt field to given value.

### HasScannedAt

`func (o *GetServerHardwareResult) HasScannedAt() bool`

HasScannedAt returns a boolean if a field has been set.

### GetServerId

`func (o *GetServerHardwareResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetServerHardwareResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetServerHardwareResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetServerHardwareResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


