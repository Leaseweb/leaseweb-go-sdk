# \DedicatedserverAPI

All URIs are relative to *https://api.leaseweb.com/internal/dedicatedserverapi/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServerToPrivateNetwork**](DedicatedserverAPI.md#AddServerToPrivateNetwork) | **Put** /servers/{serverId}/privateNetworks/{privateNetworkId} | Add a server to private network
[**DeleteServerFromPrivateNetwork**](DedicatedserverAPI.md#DeleteServerFromPrivateNetwork) | **Delete** /servers/{serverId}/privateNetworks/{privateNetworkId} | Delete a server from a private network
[**GetNetworkEquipment**](DedicatedserverAPI.md#GetNetworkEquipment) | **Get** /networkEquipments/{networkEquipmentId} | Get network equipment
[**GetNetworkEquipmentIp**](DedicatedserverAPI.md#GetNetworkEquipmentIp) | **Get** /networkEquipments/{networkEquipmentId}/ips/{ip} | Show an IP
[**GetNetworkEquipmentIpList**](DedicatedserverAPI.md#GetNetworkEquipmentIpList) | **Get** /networkEquipments/{networkEquipmentId}/ips | List IPs
[**GetNetworkEquipmentList**](DedicatedserverAPI.md#GetNetworkEquipmentList) | **Get** /networkEquipments | List network equipment
[**GetNetworkEquipmentNullRouteHistory**](DedicatedserverAPI.md#GetNetworkEquipmentNullRouteHistory) | **Get** /networkEquipments/{networkEquipmentId}/nullRouteHistory | Show null route history
[**GetServer**](DedicatedserverAPI.md#GetServer) | **Get** /servers/{serverId} | Get server
[**GetServerIp**](DedicatedserverAPI.md#GetServerIp) | **Get** /servers/{serverId}/ips/{ip} | Show an IP
[**GetServerIpList**](DedicatedserverAPI.md#GetServerIpList) | **Get** /servers/{serverId}/ips | List IPs
[**GetServerList**](DedicatedserverAPI.md#GetServerList) | **Get** /servers | List servers
[**GetServerNullRouteHistory**](DedicatedserverAPI.md#GetServerNullRouteHistory) | **Get** /servers/{serverId}/nullRouteHistory | Show null route history
[**NetworkEquipmentsIpsNull**](DedicatedserverAPI.md#NetworkEquipmentsIpsNull) | **Post** /networkEquipments/{networkEquipmentId}/ips/{ip}/null | Null route an IP
[**NullIpRoute**](DedicatedserverAPI.md#NullIpRoute) | **Post** /servers/{serverId}/ips/{ip}/null | Null route an IP
[**PutServerIp**](DedicatedserverAPI.md#PutServerIp) | **Put** /servers/{serverId}/ips/{ip} | Update an IP
[**RemoveNullIpRoute**](DedicatedserverAPI.md#RemoveNullIpRoute) | **Post** /servers/{serverId}/ips/{ip}/unnull | Remove a null route
[**UnNullNetworkEquipmentIpRoute**](DedicatedserverAPI.md#UnNullNetworkEquipmentIpRoute) | **Post** /networkEquipments/{networkEquipmentId}/ips/{ip}/unnull | Remove a null route
[**UpdateNetworkEquipmentIp**](DedicatedserverAPI.md#UpdateNetworkEquipmentIp) | **Put** /networkEquipments/{networkEquipmentId}/ips/{ip} | Update an IP
[**UpdateNetworkEquipmentReference**](DedicatedserverAPI.md#UpdateNetworkEquipmentReference) | **Put** /networkEquipments/{networkEquipmentId} | Update network equipment
[**UpdateServerReference**](DedicatedserverAPI.md#UpdateServerReference) | **Put** /servers/{serverId} | Update server



## AddServerToPrivateNetwork

> AddServerToPrivateNetwork(ctx, serverId, privateNetworkId).AddServerToPrivateNetworkOpts(addServerToPrivateNetworkOpts).Execute()

Add a server to private network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	privateNetworkId := "892" // string | The ID of a Private Network
	addServerToPrivateNetworkOpts := *openapiclient.NewAddServerToPrivateNetworkOpts(openapiclient.linkSpeed(100)) // AddServerToPrivateNetworkOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.AddServerToPrivateNetwork(context.Background(), serverId, privateNetworkId).AddServerToPrivateNetworkOpts(addServerToPrivateNetworkOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.AddServerToPrivateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**privateNetworkId** | **string** | The ID of a Private Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServerToPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addServerToPrivateNetworkOpts** | [**AddServerToPrivateNetworkOpts**](AddServerToPrivateNetworkOpts.md) |  | 

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


## DeleteServerFromPrivateNetwork

> DeleteServerFromPrivateNetwork(ctx, serverId, privateNetworkId).Execute()

Delete a server from a private network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	privateNetworkId := "892" // string | The ID of a Private Network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.DeleteServerFromPrivateNetwork(context.Background(), serverId, privateNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.DeleteServerFromPrivateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**privateNetworkId** | **string** | The ID of a Private Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerFromPrivateNetworkRequest struct via the builder pattern


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


## GetNetworkEquipment

> NetworkEquipment GetNetworkEquipment(ctx, networkEquipmentId).Execute()

Get network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNetworkEquipment(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNetworkEquipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipment`: NetworkEquipment
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNetworkEquipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkEquipment**](NetworkEquipment.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentIp

> Ip GetNetworkEquipmentIp(ctx, networkEquipmentId, ip).Execute()

Show an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNetworkEquipmentIp(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNetworkEquipmentIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNetworkEquipmentIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentIpList

> GetNetworkEquipmentIpListResult GetNetworkEquipmentIpList(ctx, networkEquipmentId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()

List IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	networkType := "networkType_example" // string | Filter the collection of ip addresses by network type (optional)
	version := "version_example" // string | Filter the collection by ip version (optional)
	nullRouted := "nullRouted_example" // string | Filter Ips by Nulled-Status (optional)
	ips := "ips_example" // string | Filter the collection of Ips for the comma separated list of Ips (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNetworkEquipmentIpList(context.Background(), networkEquipmentId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNetworkEquipmentIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentIpList`: GetNetworkEquipmentIpListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNetworkEquipmentIpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentIpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkType** | **string** | Filter the collection of ip addresses by network type | 
 **version** | **string** | Filter the collection by ip version | 
 **nullRouted** | **string** | Filter Ips by Nulled-Status | 
 **ips** | **string** | Filter the collection of Ips for the comma separated list of Ips | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetNetworkEquipmentIpListResult**](GetNetworkEquipmentIpListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentList

> GetNetworkEquipmentListResult GetNetworkEquipmentList(ctx).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()

List network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	reference := "my-switch" // string | Filter the list of network equipment by reference. (optional)
	ip := "127.0.0.4" // string | Filter the list of network equipment by ip address. (optional)
	macAddress := "aa:bb:cc:dd:ee:ff" // string | Filter the list of network equipment by mac address. (optional)
	site := "FRA-10" // string | Filter the list of network equipment by site (location). (optional)
	privateRackId := "123" // string | Filter the list of network equipment by dedicated rack id. (optional)
	privateNetworkCapable := "privateNetworkCapable_example" // string | Filter the list for private network capable network equipment (optional)
	privateNetworkEnabled := "privateNetworkEnabled_example" // string | Filter the list for private network enabled network equipment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNetworkEquipmentList(context.Background()).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNetworkEquipmentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentList`: GetNetworkEquipmentListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNetworkEquipmentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **reference** | **string** | Filter the list of network equipment by reference. | 
 **ip** | **string** | Filter the list of network equipment by ip address. | 
 **macAddress** | **string** | Filter the list of network equipment by mac address. | 
 **site** | **string** | Filter the list of network equipment by site (location). | 
 **privateRackId** | **string** | Filter the list of network equipment by dedicated rack id. | 
 **privateNetworkCapable** | **string** | Filter the list for private network capable network equipment | 
 **privateNetworkEnabled** | **string** | Filter the list for private network enabled network equipment | 

### Return type

[**GetNetworkEquipmentListResult**](GetNetworkEquipmentListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentNullRouteHistory

> GetNetworkEquipmentNullRouteHistoryResult GetNetworkEquipmentNullRouteHistory(ctx, networkEquipmentId).Limit(limit).Offset(offset).Execute()

Show null route history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetNetworkEquipmentNullRouteHistory(context.Background(), networkEquipmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetNetworkEquipmentNullRouteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentNullRouteHistory`: GetNetworkEquipmentNullRouteHistoryResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetNetworkEquipmentNullRouteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentNullRouteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetNetworkEquipmentNullRouteHistoryResult**](GetNetworkEquipmentNullRouteHistoryResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServer

> Server GetServer(ctx, serverId).Execute()

Get server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerIp

> Ip GetServerIp(ctx, serverId, ip).Execute()

Show an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetServerIp(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServerIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServerIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerIpList

> Ip GetServerIpList(ctx, serverId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()

List IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := "networkType_example" // string | Filter the collection of ip addresses by network type (optional)
	version := "version_example" // string | Filter the collection by ip version (optional)
	nullRouted := "nullRouted_example" // string | Filter Ips by Nulled-Status (optional)
	ips := "ips_example" // string | Filter the collection of Ips for the comma separated list of Ips (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetServerIpList(context.Background(), serverId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServerIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerIpList`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServerIpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerIpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkType** | **string** | Filter the collection of ip addresses by network type | 
 **version** | **string** | Filter the collection by ip version | 
 **nullRouted** | **string** | Filter Ips by Nulled-Status | 
 **ips** | **string** | Filter the collection of Ips for the comma separated list of Ips | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerList

> GetServerListResult GetServerList(ctx).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()

List servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	reference := "my-db" // string | Filter the list of servers by reference. (optional)
	ip := "127.0.0.4" // string | Filter the list of servers by ip address. (optional)
	macAddress := "aa:bb:cc:dd:ee:ff" // string | Filter the list of servers by mac address. (optional)
	site := "FRA-10" // string | Filter the list of servers by site (location). (optional)
	privateRackId := "123" // string | Filter the list of servers by dedicated rack id. (optional)
	privateNetworkCapable := "privateNetworkCapable_example" // string | Filter the list for private network capable servers (optional)
	privateNetworkEnabled := "privateNetworkEnabled_example" // string | Filter the list for private network enabled servers (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetServerList(context.Background()).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerList`: GetServerListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **reference** | **string** | Filter the list of servers by reference. | 
 **ip** | **string** | Filter the list of servers by ip address. | 
 **macAddress** | **string** | Filter the list of servers by mac address. | 
 **site** | **string** | Filter the list of servers by site (location). | 
 **privateRackId** | **string** | Filter the list of servers by dedicated rack id. | 
 **privateNetworkCapable** | **string** | Filter the list for private network capable servers | 
 **privateNetworkEnabled** | **string** | Filter the list for private network enabled servers | 

### Return type

[**GetServerListResult**](GetServerListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerNullRouteHistory

> GetServerNullRouteHistoryResult GetServerNullRouteHistory(ctx, serverId).Limit(limit).Offset(offset).Execute()

Show null route history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.GetServerNullRouteHistory(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServerNullRouteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerNullRouteHistory`: GetServerNullRouteHistoryResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServerNullRouteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerNullRouteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetServerNullRouteHistoryResult**](GetServerNullRouteHistoryResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEquipmentsIpsNull

> Ip NetworkEquipmentsIpsNull(ctx, networkEquipmentId, ip).Execute()

Null route an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.NetworkEquipmentsIpsNull(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.NetworkEquipmentsIpsNull``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkEquipmentsIpsNull`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.NetworkEquipmentsIpsNull`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkEquipmentsIpsNullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NullIpRoute

> Ip NullIpRoute(ctx, serverId, ip).Execute()

Null route an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.NullIpRoute(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.NullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.NullIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiNullIpRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerIp

> Ip PutServerIp(ctx, serverId, ip).UpdateIpProfileOpts(updateIpProfileOpts).Execute()

Update an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address
	updateIpProfileOpts := *openapiclient.NewUpdateIpProfileOpts() // UpdateIpProfileOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.PutServerIp(context.Background(), serverId, ip).UpdateIpProfileOpts(updateIpProfileOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.PutServerIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServerIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.PutServerIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIpProfileOpts** | [**UpdateIpProfileOpts**](UpdateIpProfileOpts.md) |  | 

### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNullIpRoute

> Ip RemoveNullIpRoute(ctx, serverId, ip).Execute()

Remove a null route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.RemoveNullIpRoute(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.RemoveNullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.RemoveNullIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNullIpRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnNullNetworkEquipmentIpRoute

> Ip UnNullNetworkEquipmentIpRoute(ctx, networkEquipmentId, ip).Execute()

Remove a null route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.UnNullNetworkEquipmentIpRoute(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UnNullNetworkEquipmentIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnNullNetworkEquipmentIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.UnNullNetworkEquipmentIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnNullNetworkEquipmentIpRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkEquipmentIp

> Ip UpdateNetworkEquipmentIp(ctx, networkEquipmentId, ip).UpdateNetworkEquipmentIpOpts(updateNetworkEquipmentIpOpts).Execute()

Update an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address
	updateNetworkEquipmentIpOpts := *openapiclient.NewUpdateNetworkEquipmentIpOpts() // UpdateNetworkEquipmentIpOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedserverAPI.UpdateNetworkEquipmentIp(context.Background(), networkEquipmentId, ip).UpdateNetworkEquipmentIpOpts(updateNetworkEquipmentIpOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateNetworkEquipmentIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEquipmentIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.UpdateNetworkEquipmentIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkEquipmentIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkEquipmentIpOpts** | [**UpdateNetworkEquipmentIpOpts**](UpdateNetworkEquipmentIpOpts.md) |  | 

### Return type

[**Ip**](Ip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkEquipmentReference

> UpdateNetworkEquipmentReference(ctx, networkEquipmentId).UpdateNetworkEquipmentReferenceOpts(updateNetworkEquipmentReferenceOpts).Execute()

Update network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	updateNetworkEquipmentReferenceOpts := *openapiclient.NewUpdateNetworkEquipmentReferenceOpts("Reference_example") // UpdateNetworkEquipmentReferenceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.UpdateNetworkEquipmentReference(context.Background(), networkEquipmentId).UpdateNetworkEquipmentReferenceOpts(updateNetworkEquipmentReferenceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateNetworkEquipmentReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkEquipmentReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkEquipmentReferenceOpts** | [**UpdateNetworkEquipmentReferenceOpts**](UpdateNetworkEquipmentReferenceOpts.md) |  | 

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


## UpdateServerReference

> UpdateServerReference(ctx, serverId).UpdateServerReferenceOpts(updateServerReferenceOpts).Execute()

Update server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

func main() {
	serverId := "12345" // string | The ID of a server
	updateServerReferenceOpts := *openapiclient.NewUpdateServerReferenceOpts("Reference_example") // UpdateServerReferenceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedserverAPI.UpdateServerReference(context.Background(), serverId).UpdateServerReferenceOpts(updateServerReferenceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.UpdateServerReference``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateServerReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerReferenceOpts** | [**UpdateServerReferenceOpts**](UpdateServerReferenceOpts.md) |  | 

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

