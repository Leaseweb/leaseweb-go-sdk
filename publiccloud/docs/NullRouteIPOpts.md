# NullRouteIPOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The reason why the IP is being null-routed | [optional] 
**AutomatedUnnulingAt** | Pointer to **int32** | If provided, reverts the operation automatically in the specified value, in hours | [optional] 

## Methods

### NewNullRouteIPOpts

`func NewNullRouteIPOpts() *NullRouteIPOpts`

NewNullRouteIPOpts instantiates a new NullRouteIPOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullRouteIPOptsWithDefaults

`func NewNullRouteIPOptsWithDefaults() *NullRouteIPOpts`

NewNullRouteIPOptsWithDefaults instantiates a new NullRouteIPOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *NullRouteIPOpts) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NullRouteIPOpts) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NullRouteIPOpts) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NullRouteIPOpts) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAutomatedUnnulingAt

`func (o *NullRouteIPOpts) GetAutomatedUnnulingAt() int32`

GetAutomatedUnnulingAt returns the AutomatedUnnulingAt field if non-nil, zero value otherwise.

### GetAutomatedUnnulingAtOk

`func (o *NullRouteIPOpts) GetAutomatedUnnulingAtOk() (*int32, bool)`

GetAutomatedUnnulingAtOk returns a tuple with the AutomatedUnnulingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedUnnulingAt

`func (o *NullRouteIPOpts) SetAutomatedUnnulingAt(v int32)`

SetAutomatedUnnulingAt sets AutomatedUnnulingAt field to given value.

### HasAutomatedUnnulingAt

`func (o *NullRouteIPOpts) HasAutomatedUnnulingAt() bool`

HasAutomatedUnnulingAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


