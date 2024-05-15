# \AbuseAPI

All URIs are relative to *https://api.leaseweb.com/abuse/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReportMessage**](AbuseAPI.md#CreateReportMessage) | **Post** /reports/{reportId}/messages | Create new message
[**GetReport**](AbuseAPI.md#GetReport) | **Get** /reports/{reportId} | Inspect a report
[**GetReportAttachmentList**](AbuseAPI.md#GetReportAttachmentList) | **Get** /reports/{reportId}/reportAttachments/{fileId} | Inspect a report attachment
[**GetReportList**](AbuseAPI.md#GetReportList) | **Get** /reports | List reports
[**GetReportMessageAttachmentList**](AbuseAPI.md#GetReportMessageAttachmentList) | **Get** /reports/{reportId}/messageAttachments/{fileId} | Inspect a message attachment
[**GetReportMessageList**](AbuseAPI.md#GetReportMessageList) | **Get** /reports/{reportId}/messages | List report messages
[**GetReportResolutionList**](AbuseAPI.md#GetReportResolutionList) | **Get** /reports/{reportId}/resolutions | List resolution options
[**ResolveReport**](AbuseAPI.md#ResolveReport) | **Post** /reports/{reportId}/resolve | Resolve a report



## CreateReportMessage

> []string CreateReportMessage(ctx, reportId).CreateReportMessageOpts(createReportMessageOpts).Execute()

Create new message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	reportId := "abc123" // string | Report Id
	createReportMessageOpts := *openapiclient.NewCreateReportMessageOpts("I have resolved this issue.") // CreateReportMessageOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbuseAPI.CreateReportMessage(context.Background(), reportId).CreateReportMessageOpts(createReportMessageOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.CreateReportMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReportMessage`: []string
	fmt.Fprintf(os.Stdout, "Response from `AbuseAPI.CreateReportMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Report Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createReportMessageOpts** | [**CreateReportMessageOpts**](CreateReportMessageOpts.md) |  | 

### Return type

**[]string**

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReport

> GetReportResult GetReport(ctx, reportId).Execute()

Inspect a report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	reportId := "abc123" // string | Report Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbuseAPI.GetReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReport`: GetReportResult
	fmt.Fprintf(os.Stdout, "Response from `AbuseAPI.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Report Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReportResult**](GetReportResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportAttachmentList

> GetReportAttachmentList(ctx, reportId, fileId).Execute()

Inspect a report attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	reportId := "abc123" // string | Report Id
	fileId := "sdfa73-adsfs-4fadf-sdfasdfa" // string | File Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AbuseAPI.GetReportAttachmentList(context.Background(), reportId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetReportAttachmentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Report Id | 
**fileId** | **string** | File Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportAttachmentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportList

> GetReportListResult GetReportList(ctx).Limit(limit).Offset(offset).Status(status).Execute()

List reports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	status := "OPEN,WAITING,CLOSED" // string | Comma separated list of report statuses to filter on.  (optional) (default to "OPEN,WAITING,CLOSED")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbuseAPI.GetReportList(context.Background()).Limit(limit).Offset(offset).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetReportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportList`: GetReportListResult
	fmt.Fprintf(os.Stdout, "Response from `AbuseAPI.GetReportList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **status** | **string** | Comma separated list of report statuses to filter on.  | [default to &quot;OPEN,WAITING,CLOSED&quot;]

### Return type

[**GetReportListResult**](GetReportListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportMessageAttachmentList

> GetReportMessageAttachmentList(ctx, reportId, fileId).Execute()

Inspect a message attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	reportId := "abc123" // string | Report Id
	fileId := "sdfa73-adsfs-4fadf-sdfasdfa" // string | File Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AbuseAPI.GetReportMessageAttachmentList(context.Background(), reportId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetReportMessageAttachmentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Report Id | 
**fileId** | **string** | File Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportMessageAttachmentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportMessageList

> GetReportMessageListResult GetReportMessageList(ctx, reportId).Limit(limit).Offset(offset).Execute()

List report messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	reportId := "abc123" // string | Report Id
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbuseAPI.GetReportMessageList(context.Background(), reportId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetReportMessageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportMessageList`: GetReportMessageListResult
	fmt.Fprintf(os.Stdout, "Response from `AbuseAPI.GetReportMessageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Report Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportMessageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetReportMessageListResult**](GetReportMessageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportResolutionList

> GetReportResolutionListResult GetReportResolutionList(ctx, reportId).Execute()

List resolution options



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	reportId := "abc123" // string | Report Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbuseAPI.GetReportResolutionList(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetReportResolutionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportResolutionList`: GetReportResolutionListResult
	fmt.Fprintf(os.Stdout, "Response from `AbuseAPI.GetReportResolutionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Report Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportResolutionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReportResolutionListResult**](GetReportResolutionListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReport

> ResolveReport(ctx, reportId).ResolveReportResult(resolveReportResult).Execute()

Resolve a report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	reportId := "abc123" // string | Report Id
	resolveReportResult := *openapiclient.NewResolveReportResult([]string{"["CONTENT_REMOVED","SUSPENDED"]"}) // ResolveReportResult |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AbuseAPI.ResolveReport(context.Background(), reportId).ResolveReportResult(resolveReportResult).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.ResolveReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Report Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolveReportResult** | [**ResolveReportResult**](ResolveReportResult.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

