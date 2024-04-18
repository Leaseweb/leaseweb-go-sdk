# GetServerHardwareScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Hardware Scan ID | [optional] 
**ParserVersion** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**Results**](Results.md) |  | [optional] 
**ScannedAt** | Pointer to **string** | Date time using the ISO-8601 format | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetServerHardwareScanResult

`func NewGetServerHardwareScanResult() *GetServerHardwareScanResult`

NewGetServerHardwareScanResult instantiates a new GetServerHardwareScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerHardwareScanResultWithDefaults

`func NewGetServerHardwareScanResultWithDefaults() *GetServerHardwareScanResult`

NewGetServerHardwareScanResultWithDefaults instantiates a new GetServerHardwareScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetServerHardwareScanResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServerHardwareScanResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServerHardwareScanResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetServerHardwareScanResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParserVersion

`func (o *GetServerHardwareScanResult) GetParserVersion() string`

GetParserVersion returns the ParserVersion field if non-nil, zero value otherwise.

### GetParserVersionOk

`func (o *GetServerHardwareScanResult) GetParserVersionOk() (*string, bool)`

GetParserVersionOk returns a tuple with the ParserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserVersion

`func (o *GetServerHardwareScanResult) SetParserVersion(v string)`

SetParserVersion sets ParserVersion field to given value.

### HasParserVersion

`func (o *GetServerHardwareScanResult) HasParserVersion() bool`

HasParserVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetServerHardwareScanResult) GetResult() Results`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetServerHardwareScanResult) GetResultOk() (*Results, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetServerHardwareScanResult) SetResult(v Results)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetServerHardwareScanResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScannedAt

`func (o *GetServerHardwareScanResult) GetScannedAt() string`

GetScannedAt returns the ScannedAt field if non-nil, zero value otherwise.

### GetScannedAtOk

`func (o *GetServerHardwareScanResult) GetScannedAtOk() (*string, bool)`

GetScannedAtOk returns a tuple with the ScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedAt

`func (o *GetServerHardwareScanResult) SetScannedAt(v string)`

SetScannedAt sets ScannedAt field to given value.

### HasScannedAt

`func (o *GetServerHardwareScanResult) HasScannedAt() bool`

HasScannedAt returns a boolean if a field has been set.

### GetServerId

`func (o *GetServerHardwareScanResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetServerHardwareScanResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetServerHardwareScanResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetServerHardwareScanResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


