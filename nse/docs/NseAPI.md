# \NseAPI

All URIs are relative to *https://api.leaseweb.com/internal/nseapi/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseNetworkInterface**](NseAPI.md#CloseNetworkInterface) | **Post** /servers/{serverId}/networkInterfaces/{networkType}/close | Close network interface
[**CloseNetworkInterfaces**](NseAPI.md#CloseNetworkInterfaces) | **Post** /servers/{serverId}/networkInterfaces/close | Close all network interfaces
[**GetDdosNotificationSettingList**](NseAPI.md#GetDdosNotificationSettingList) | **Get** /servers/{serverId}/notificationSettings/ddos | Inspect DDoS notification settings
[**GetNetworkInterface**](NseAPI.md#GetNetworkInterface) | **Get** /servers/{serverId}/networkInterfaces/{networkType} | Show a network interface
[**GetNetworkInterfaceList**](NseAPI.md#GetNetworkInterfaceList) | **Get** /servers/{serverId}/networkInterfaces | List network interfaces
[**OpenNetworkInterface**](NseAPI.md#OpenNetworkInterface) | **Post** /servers/{serverId}/networkInterfaces/{networkType}/open | Open network interface
[**OpenNetworkInterfaces**](NseAPI.md#OpenNetworkInterfaces) | **Post** /servers/{serverId}/networkInterfaces/open | Open all network interfaces
[**UpdateNullRoute**](NseAPI.md#UpdateNullRoute) | **Post** /actions/nullRoute | Announce/Withdraw Null Routes
[**UpdateServerDdosNotificationSetting**](NseAPI.md#UpdateServerDdosNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/ddos | Update DDoS notification settings



## CloseNetworkInterface

> CloseNetworkInterface(ctx, serverId, networkType).Execute()

Close network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := "public" // string | The network type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NseAPI.CloseNetworkInterface(context.Background(), serverId, networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.CloseNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkType** | **string** | The network type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseNetworkInterfaceRequest struct via the builder pattern


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


## CloseNetworkInterfaces

> CloseNetworkInterfaces(ctx, serverId).Execute()

Close all network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NseAPI.CloseNetworkInterfaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.CloseNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseNetworkInterfacesRequest struct via the builder pattern


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


## GetDdosNotificationSettingList

> GetDdosNotificationSettingResult GetDdosNotificationSettingList(ctx, serverId).Execute()

Inspect DDoS notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NseAPI.GetDdosNotificationSettingList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.GetDdosNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDdosNotificationSettingList`: GetDdosNotificationSettingResult
	fmt.Fprintf(os.Stdout, "Response from `NseAPI.GetDdosNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDdosNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDdosNotificationSettingResult**](GetDdosNotificationSettingResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterface

> NetworkInterface GetNetworkInterface(ctx, serverId, networkType).Execute()

Show a network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := "public" // string | The network type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NseAPI.GetNetworkInterface(context.Background(), serverId, networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.GetNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterface`: NetworkInterface
	fmt.Fprintf(os.Stdout, "Response from `NseAPI.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkType** | **string** | The network type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkInterface**](NetworkInterface.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterfaceList

> GetNetworkInterfaceListResult GetNetworkInterfaceList(ctx, serverId).Execute()

List network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NseAPI.GetNetworkInterfaceList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.GetNetworkInterfaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterfaceList`: GetNetworkInterfaceListResult
	fmt.Fprintf(os.Stdout, "Response from `NseAPI.GetNetworkInterfaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkInterfaceListResult**](GetNetworkInterfaceListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenNetworkInterface

> OpenNetworkInterface(ctx, serverId, networkType).Execute()

Open network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := "public" // string | The network type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NseAPI.OpenNetworkInterface(context.Background(), serverId, networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.OpenNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkType** | **string** | The network type | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenNetworkInterfaceRequest struct via the builder pattern


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


## OpenNetworkInterfaces

> OpenNetworkInterfaces(ctx, serverId).Execute()

Open all network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NseAPI.OpenNetworkInterfaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.OpenNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenNetworkInterfacesRequest struct via the builder pattern


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


## UpdateNullRoute

> UpdateNullRouteAcceptedResult UpdateNullRoute(ctx).UpdateNullRouteOpts(updateNullRouteOpts).Execute()

Announce/Withdraw Null Routes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	updateNullRouteOpts := *openapiclient.NewUpdateNullRouteOpts("WITHDRAW", "IpAddress_example") // UpdateNullRouteOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NseAPI.UpdateNullRoute(context.Background()).UpdateNullRouteOpts(updateNullRouteOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.UpdateNullRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNullRoute`: UpdateNullRouteAcceptedResult
	fmt.Fprintf(os.Stdout, "Response from `NseAPI.UpdateNullRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNullRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNullRouteOpts** | [**UpdateNullRouteOpts**](UpdateNullRouteOpts.md) |  | 

### Return type

[**UpdateNullRouteAcceptedResult**](UpdateNullRouteAcceptedResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDdosNotificationSetting

> UpdateServerDdosNotificationSetting(ctx, serverId).UpdateDdosNotificationSettingOpts(updateDdosNotificationSettingOpts).Execute()

Update DDoS notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/nse"
)

func main() {
	serverId := "12345" // string | The ID of a server
	updateDdosNotificationSettingOpts := *openapiclient.NewUpdateDdosNotificationSettingOpts() // UpdateDdosNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NseAPI.UpdateServerDdosNotificationSetting(context.Background(), serverId).UpdateDdosNotificationSettingOpts(updateDdosNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NseAPI.UpdateServerDdosNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDdosNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDdosNotificationSettingOpts** | [**UpdateDdosNotificationSettingOpts**](UpdateDdosNotificationSettingOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

