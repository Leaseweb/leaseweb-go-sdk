# FailedRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** | The reason which the edition has failed | [optional] 

## Methods

### NewFailedRule

`func NewFailedRule() *FailedRule`

NewFailedRule instantiates a new FailedRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedRuleWithDefaults

`func NewFailedRuleWithDefaults() *FailedRule`

NewFailedRuleWithDefaults instantiates a new FailedRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FailedRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailedRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailedRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FailedRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *FailedRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FailedRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FailedRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FailedRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


