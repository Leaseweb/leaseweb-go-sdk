# \BmsdbAPI

All URIs are relative to *https://api.leaseweb.com/internal/bmsdb/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerHardware**](BmsdbAPI.md#GetServerHardware) | **Get** /servers/{serverId}/hardwareInfo | Show hardware information



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

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

