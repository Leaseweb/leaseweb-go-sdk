# GetServerPowerStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipmi** | Pointer to [**Ipmi**](Ipmi.md) |  | [optional] 
**Pdu** | Pointer to [**Pdu**](Pdu.md) |  | [optional] 

## Methods

### NewGetServerPowerStatusResult

`func NewGetServerPowerStatusResult() *GetServerPowerStatusResult`

NewGetServerPowerStatusResult instantiates a new GetServerPowerStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerPowerStatusResultWithDefaults

`func NewGetServerPowerStatusResultWithDefaults() *GetServerPowerStatusResult`

NewGetServerPowerStatusResultWithDefaults instantiates a new GetServerPowerStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpmi

`func (o *GetServerPowerStatusResult) GetIpmi() Ipmi`

GetIpmi returns the Ipmi field if non-nil, zero value otherwise.

### GetIpmiOk

`func (o *GetServerPowerStatusResult) GetIpmiOk() (*Ipmi, bool)`

GetIpmiOk returns a tuple with the Ipmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmi

`func (o *GetServerPowerStatusResult) SetIpmi(v Ipmi)`

SetIpmi sets Ipmi field to given value.

### HasIpmi

`func (o *GetServerPowerStatusResult) HasIpmi() bool`

HasIpmi returns a boolean if a field has been set.

### GetPdu

`func (o *GetServerPowerStatusResult) GetPdu() Pdu`

GetPdu returns the Pdu field if non-nil, zero value otherwise.

### GetPduOk

`func (o *GetServerPowerStatusResult) GetPduOk() (*Pdu, bool)`

GetPduOk returns a tuple with the Pdu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdu

`func (o *GetServerPowerStatusResult) SetPdu(v Pdu)`

SetPdu sets Pdu field to given value.

### HasPdu

`func (o *GetServerPowerStatusResult) HasPdu() bool`

HasPdu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


