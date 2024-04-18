# \BmusageAPI

All URIs are relative to *https://api.leaseweb.com/internal/bmusageapi/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerBandwidthNotificationSettingJson**](BmusageAPI.md#CreateServerBandwidthNotificationSettingJson) | **Post** /servers/{serverId}/notificationSettings/bandwidth | Create a bandwidth notification setting
[**CreateServerDataTrafficNotificationSetting**](BmusageAPI.md#CreateServerDataTrafficNotificationSetting) | **Post** /servers/{serverId}/notificationSettings/datatraffic | Create a data traffic notification setting
[**DeleteServerBandwidthNotificationSetting**](BmusageAPI.md#DeleteServerBandwidthNotificationSetting) | **Delete** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Delete a bandwidth notification setting
[**DeleteServerDataTrafficNotificationSetting**](BmusageAPI.md#DeleteServerDataTrafficNotificationSetting) | **Delete** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Delete a data traffic notification setting
[**GetServerBandwidthNotificationSetting**](BmusageAPI.md#GetServerBandwidthNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Show a bandwidth notification setting
[**GetServerBandwidthNotificationSettingList**](BmusageAPI.md#GetServerBandwidthNotificationSettingList) | **Get** /servers/{serverId}/notificationSettings/bandwidth | List bandwidth notification settings
[**GetServerDataTrafficMetrics**](BmusageAPI.md#GetServerDataTrafficMetrics) | **Get** /servers/{serverId}/metrics/datatraffic | Show datatraffic metrics
[**GetServerDataTrafficNotificationSetting**](BmusageAPI.md#GetServerDataTrafficNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Show a data traffic notification setting
[**GetServerDataTrafficNotificationSettingList**](BmusageAPI.md#GetServerDataTrafficNotificationSettingList) | **Get** /servers/{serverId}/notificationSettings/datatraffic | List data traffic notification settings
[**ServersMetricsBandwidth**](BmusageAPI.md#ServersMetricsBandwidth) | **Get** /servers/{serverId}/metrics/bandwidth | Show bandwidth metrics
[**UpdateServerBandwidthNotificationSetting**](BmusageAPI.md#UpdateServerBandwidthNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Update a bandwidth notification setting
[**UpdateServerDataTrafficNotificationSetting**](BmusageAPI.md#UpdateServerDataTrafficNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Update a data traffic notification setting



## CreateServerBandwidthNotificationSettingJson

> BandwidthNotificationSetting CreateServerBandwidthNotificationSettingJson(ctx, serverId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()

Create a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	bandwidthNotificationSettingOpts := *openapiclient.NewBandwidthNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // BandwidthNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.CreateServerBandwidthNotificationSettingJson(context.Background(), serverId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.CreateServerBandwidthNotificationSettingJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerBandwidthNotificationSettingJson`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.CreateServerBandwidthNotificationSettingJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerBandwidthNotificationSettingJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bandwidthNotificationSettingOpts** | [**BandwidthNotificationSettingOpts**](BandwidthNotificationSettingOpts.md) |  | 

### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerDataTrafficNotificationSetting

> DataTrafficNotificationSetting CreateServerDataTrafficNotificationSetting(ctx, serverId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()

Create a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	dataTrafficNotificationSettingOpts := *openapiclient.NewDataTrafficNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // DataTrafficNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.CreateServerDataTrafficNotificationSetting(context.Background(), serverId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.CreateServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerDataTrafficNotificationSetting`: DataTrafficNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.CreateServerDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataTrafficNotificationSettingOpts** | [**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md) |  | 

### Return type

[**DataTrafficNotificationSetting**](DataTrafficNotificationSetting.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerBandwidthNotificationSetting

> DeleteServerBandwidthNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Delete a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmusageAPI.DeleteServerBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.DeleteServerBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerDataTrafficNotificationSetting

> DeleteServerDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Delete a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmusageAPI.DeleteServerDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.DeleteServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerBandwidthNotificationSetting

> BandwidthNotificationSetting GetServerBandwidthNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Show a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.GetServerBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.GetServerBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.GetServerBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerBandwidthNotificationSettingList

> GetServerBandwidthNotificationSettingListResult GetServerBandwidthNotificationSettingList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List bandwidth notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.GetServerBandwidthNotificationSettingList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.GetServerBandwidthNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerBandwidthNotificationSettingList`: GetServerBandwidthNotificationSettingListResult
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.GetServerBandwidthNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerBandwidthNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetServerBandwidthNotificationSettingListResult**](GetServerBandwidthNotificationSettingListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDataTrafficMetrics

> Metrics GetServerDataTrafficMetrics(ctx, serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()

Show datatraffic metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	from := time.Now() // time.Time | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time.
	to := time.Now() // time.Time | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time.
	aggregation := "aggregation_example" // string | Aggregate each metric using the given aggregation function.
	granularity := "granularity_example" // string | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.GetServerDataTrafficMetrics(context.Background(), serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.GetServerDataTrafficMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDataTrafficMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.GetServerDataTrafficMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDataTrafficMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time. | 
 **to** | **time.Time** | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time. | 
 **aggregation** | **string** | Aggregate each metric using the given aggregation function. | 
 **granularity** | **string** | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDataTrafficNotificationSetting

> DataTrafficNotificationSetting GetServerDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Show a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.GetServerDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.GetServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDataTrafficNotificationSetting`: DataTrafficNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.GetServerDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataTrafficNotificationSetting**](DataTrafficNotificationSetting.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDataTrafficNotificationSettingList

> GetServerDataTrafficNotificationSettingListResult GetServerDataTrafficNotificationSettingList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List data traffic notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.GetServerDataTrafficNotificationSettingList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.GetServerDataTrafficNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDataTrafficNotificationSettingList`: GetServerDataTrafficNotificationSettingListResult
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.GetServerDataTrafficNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDataTrafficNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetServerDataTrafficNotificationSettingListResult**](GetServerDataTrafficNotificationSettingListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersMetricsBandwidth

> Metrics ServersMetricsBandwidth(ctx, serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()

Show bandwidth metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	from := time.Now() // time.Time | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time.
	to := time.Now() // time.Time | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time.
	aggregation := "aggregation_example" // string | Aggregate each metric using the given aggregation function. When the aggregation type `95TH` is specified the granularity parameter should be omitted from the request.
	granularity := "granularity_example" // string | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.ServersMetricsBandwidth(context.Background(), serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.ServersMetricsBandwidth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersMetricsBandwidth`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.ServersMetricsBandwidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersMetricsBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time. | 
 **to** | **time.Time** | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time. | 
 **aggregation** | **string** | Aggregate each metric using the given aggregation function. When the aggregation type &#x60;95TH&#x60; is specified the granularity parameter should be omitted from the request. | 
 **granularity** | **string** | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerBandwidthNotificationSetting

> BandwidthNotificationSetting UpdateServerBandwidthNotificationSetting(ctx, serverId, notificationSettingId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()

Update a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting
	bandwidthNotificationSettingOpts := *openapiclient.NewBandwidthNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // BandwidthNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.UpdateServerBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.UpdateServerBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.UpdateServerBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandwidthNotificationSettingOpts** | [**BandwidthNotificationSettingOpts**](BandwidthNotificationSettingOpts.md) |  | 

### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDataTrafficNotificationSetting

> DataTrafficNotificationSettingOpts UpdateServerDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()

Update a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmusage"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting
	dataTrafficNotificationSettingOpts := *openapiclient.NewDataTrafficNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // DataTrafficNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmusageAPI.UpdateServerDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmusageAPI.UpdateServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerDataTrafficNotificationSetting`: DataTrafficNotificationSettingOpts
	fmt.Fprintf(os.Stdout, "Response from `BmusageAPI.UpdateServerDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataTrafficNotificationSettingOpts** | [**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md) |  | 

### Return type

[**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

