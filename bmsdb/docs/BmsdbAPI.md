# \BmsdbAPI

All URIs are relative to *https://api.leaseweb.com/internal/bmsdb/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerHardware**](BmsdbAPI.md#GetServerHardware) | **Get** /servers/{serverId}/hardwareInfo | Show hardware information
[**GetServerHardwareScan**](BmsdbAPI.md#GetServerHardwareScan) | **Get** /servers/{serverId}/hardwareScans/{hardwareScanId} | Show a hardware scan
[**ServersHardwarescansList**](BmsdbAPI.md#ServersHardwarescansList) | **Get** /servers/{serverId}/hardwareScans | List hardware scans



## GetServerHardware

> GetServerHardwareResult GetServerHardware(ctx, serverId).Execute()

Show hardware information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmsdb"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmsdbAPI.GetServerHardware(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmsdbAPI.GetServerHardware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerHardware`: GetServerHardwareResult
	fmt.Fprintf(os.Stdout, "Response from `BmsdbAPI.GetServerHardware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerHardwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServerHardwareResult**](GetServerHardwareResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerHardwareScan

> GetServerHardwareScanResult GetServerHardwareScan(ctx, serverId, hardwareScanId).Format(format).Execute()

Show a hardware scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmsdb"
)

func main() {
	serverId := "12345" // string | The ID of a server
	hardwareScanId := "hardwareScanId_example" // string | The ID that identifies a hardware scan
	format := "json" // string | Defines the format of the response (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmsdbAPI.GetServerHardwareScan(context.Background(), serverId, hardwareScanId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmsdbAPI.GetServerHardwareScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerHardwareScan`: GetServerHardwareScanResult
	fmt.Fprintf(os.Stdout, "Response from `BmsdbAPI.GetServerHardwareScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**hardwareScanId** | **string** | The ID that identifies a hardware scan | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerHardwareScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Defines the format of the response | [default to &quot;json&quot;]

### Return type

[**GetServerHardwareScanResult**](GetServerHardwareScanResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersHardwarescansList

> GetServerHardwareScanListResult ServersHardwarescansList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List hardware scans



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmsdb"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmsdbAPI.ServersHardwarescansList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmsdbAPI.ServersHardwarescansList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersHardwarescansList`: GetServerHardwareScanListResult
	fmt.Fprintf(os.Stdout, "Response from `BmsdbAPI.ServersHardwarescansList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersHardwarescansListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetServerHardwareScanListResult**](GetServerHardwareScanListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

