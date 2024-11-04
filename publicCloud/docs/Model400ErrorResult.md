# Model400ErrorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | The correlation ID of the current request. | [optional] 
**ErrorCode** | Pointer to **string** | The error code. | [optional] 
**ErrorMessage** | Pointer to **string** | A human friendly description of the error. | [optional] 
**ErrorDetails** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewModel400ErrorResult

`func NewModel400ErrorResult() *Model400ErrorResult`

NewModel400ErrorResult instantiates a new Model400ErrorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel400ErrorResultWithDefaults

`func NewModel400ErrorResultWithDefaults() *Model400ErrorResult`

NewModel400ErrorResultWithDefaults instantiates a new Model400ErrorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *Model400ErrorResult) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *Model400ErrorResult) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *Model400ErrorResult) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *Model400ErrorResult) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetErrorCode

`func (o *Model400ErrorResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Model400ErrorResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Model400ErrorResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Model400ErrorResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Model400ErrorResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Model400ErrorResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Model400ErrorResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Model400ErrorResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorDetails

`func (o *Model400ErrorResult) GetErrorDetails() map[string][]string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *Model400ErrorResult) GetErrorDetailsOk() (*map[string][]string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *Model400ErrorResult) SetErrorDetails(v map[string][]string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *Model400ErrorResult) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


