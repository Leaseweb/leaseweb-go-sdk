# NullRouteIpOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The reason why the IP is being null-routed | [optional] 
**AutomatedUnnulingAt** | Pointer to **int32** | If provided, reverts the operation automatically in the specified value, in hours | [optional] 

## Methods

### NewNullRouteIpOpts

`func NewNullRouteIpOpts() *NullRouteIpOpts`

NewNullRouteIpOpts instantiates a new NullRouteIpOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullRouteIpOptsWithDefaults

`func NewNullRouteIpOptsWithDefaults() *NullRouteIpOpts`

NewNullRouteIpOptsWithDefaults instantiates a new NullRouteIpOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *NullRouteIpOpts) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NullRouteIpOpts) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NullRouteIpOpts) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NullRouteIpOpts) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAutomatedUnnulingAt

`func (o *NullRouteIpOpts) GetAutomatedUnnulingAt() int32`

GetAutomatedUnnulingAt returns the AutomatedUnnulingAt field if non-nil, zero value otherwise.

### GetAutomatedUnnulingAtOk

`func (o *NullRouteIpOpts) GetAutomatedUnnulingAtOk() (*int32, bool)`

GetAutomatedUnnulingAtOk returns a tuple with the AutomatedUnnulingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedUnnulingAt

`func (o *NullRouteIpOpts) SetAutomatedUnnulingAt(v int32)`

SetAutomatedUnnulingAt sets AutomatedUnnulingAt field to given value.

### HasAutomatedUnnulingAt

`func (o *NullRouteIpOpts) HasAutomatedUnnulingAt() bool`

HasAutomatedUnnulingAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


