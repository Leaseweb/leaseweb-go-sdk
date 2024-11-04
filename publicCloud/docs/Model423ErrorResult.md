# Model423ErrorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | The error code. | [optional] 
**ErrorMessage** | Pointer to **string** | A human friendly description of the error. | [optional] 
**ErrorDetails** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewModel423ErrorResult

`func NewModel423ErrorResult() *Model423ErrorResult`

NewModel423ErrorResult instantiates a new Model423ErrorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel423ErrorResultWithDefaults

`func NewModel423ErrorResultWithDefaults() *Model423ErrorResult`

NewModel423ErrorResultWithDefaults instantiates a new Model423ErrorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Model423ErrorResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Model423ErrorResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Model423ErrorResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Model423ErrorResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Model423ErrorResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Model423ErrorResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Model423ErrorResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Model423ErrorResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorDetails

`func (o *Model423ErrorResult) GetErrorDetails() map[string][]string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *Model423ErrorResult) GetErrorDetailsOk() (*map[string][]string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *Model423ErrorResult) SetErrorDetails(v map[string][]string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *Model423ErrorResult) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


