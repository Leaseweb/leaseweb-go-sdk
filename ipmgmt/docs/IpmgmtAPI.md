# \IpmgmtAPI

All URIs are relative to *https://api.leaseweb.com/ipMgmt/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIPList**](IpmgmtAPI.md#GetIPList) | **Get** /ips | List IPs
[**GetNullRouteHistoryList**](IpmgmtAPI.md#GetNullRouteHistoryList) | **Get** /nullRoutes | Inspect null route history
[**GetReverseLookupRecordList**](IpmgmtAPI.md#GetReverseLookupRecordList) | **Get** /ips/{ip}/reverseLookup | List reverse lookup records for an IPv6 range
[**InspectIP**](IpmgmtAPI.md#InspectIP) | **Get** /ips/{ip} | Inspect an IP
[**InspectNullRouteHistory**](IpmgmtAPI.md#InspectNullRouteHistory) | **Get** /nullRoutes/{id} | Inspect null route history
[**NullRouteIP**](IpmgmtAPI.md#NullRouteIP) | **Post** /ips/{ip}/nullRoute | Null route an IP
[**RemoveIPNullRoute**](IpmgmtAPI.md#RemoveIPNullRoute) | **Delete** /ips/{ip}/nullRoute | Remove a null route
[**UpdateIP**](IpmgmtAPI.md#UpdateIP) | **Put** /ips/{ip} | Update an IP
[**UpdateNullRoute**](IpmgmtAPI.md#UpdateNullRoute) | **Put** /nullRoutes/{id} | Update a null route
[**UpdateReverseLookupRecords**](IpmgmtAPI.md#UpdateReverseLookupRecords) | **Put** /ips/{ip}/reverseLookup | Set or remove reverse lookup records for an IPv6 range



## GetIPList

> GetIPListResult GetIPList(ctx).Limit(limit).Offset(offset).SubnetId(subnetId).Version(version).Type_(type_).NullRouted(nullRouted).Primary(primary).FromIp(fromIp).ToIp(toIp).Ips(ips).EquipmentIds(equipmentIds).AssignedContractIds(assignedContractIds).Sort(sort).ReverseLookup(reverseLookup).Execute()

List IPs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	subnetId := "subnetId_example" // string | Filter by subnet (optional)
	version := openapiclient.protocolVersion(4) // ProtocolVersion | Filter by protocol version (optional)
	type_ := openapiclient.ipType("NORMAL_IP") // IpType | Filter by IP type (optional)
	nullRouted := true // bool | Filter by whether or not the IP has an active null route (true or false) (optional)
	primary := true // bool | Filter by whether or not the IP is primary (true or false) (optional)
	fromIp := "fromIp_example" // string | Return only IPs greater or equal to the specified address (optional)
	toIp := "toIp_example" // string | Return only IPs lower or equal to the specified address (optional)
	ips := "ips_example" // string | Return only IPs specified as a comma-separated list (optional)
	equipmentIds := "equipmentIds_example" // string | Return only IPs assigned to equipment items specified as a comma-separated list of IDs (optional)
	assignedContractIds := "assignedContractIds_example" // string | Return only IPs assigned to contracts specified as a comma-separated list of IDs (optional)
	sort := "sort_example" // string | Comma-separated list of sort field names. Prepend the field name with '-' for descending order. E.g. `sort=ip,-nullrouted`. Sortable field names are ip, nullRouted, reverseLookup. (optional)
	reverseLookup := "mydomain1.example.com" // string | Filter by reverse lookup. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.GetIPList(context.Background()).Limit(limit).Offset(offset).SubnetId(subnetId).Version(version).Type_(type_).NullRouted(nullRouted).Primary(primary).FromIp(fromIp).ToIp(toIp).Ips(ips).EquipmentIds(equipmentIds).AssignedContractIds(assignedContractIds).Sort(sort).ReverseLookup(reverseLookup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.GetIPList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPList`: GetIPListResult
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.GetIPList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIPListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **subnetId** | **string** | Filter by subnet | 
 **version** | [**ProtocolVersion**](ProtocolVersion.md) | Filter by protocol version | 
 **type_** | [**IpType**](IpType.md) | Filter by IP type | 
 **nullRouted** | **bool** | Filter by whether or not the IP has an active null route (true or false) | 
 **primary** | **bool** | Filter by whether or not the IP is primary (true or false) | 
 **fromIp** | **string** | Return only IPs greater or equal to the specified address | 
 **toIp** | **string** | Return only IPs lower or equal to the specified address | 
 **ips** | **string** | Return only IPs specified as a comma-separated list | 
 **equipmentIds** | **string** | Return only IPs assigned to equipment items specified as a comma-separated list of IDs | 
 **assignedContractIds** | **string** | Return only IPs assigned to contracts specified as a comma-separated list of IDs | 
 **sort** | **string** | Comma-separated list of sort field names. Prepend the field name with &#39;-&#39; for descending order. E.g. &#x60;sort&#x3D;ip,-nullrouted&#x60;. Sortable field names are ip, nullRouted, reverseLookup. | 
 **reverseLookup** | **string** | Filter by reverse lookup. | 

### Return type

[**GetIPListResult**](GetIPListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNullRouteHistoryList

> GetNullRouteHistoryListResult GetNullRouteHistoryList(ctx).Limit(limit).Offset(offset).FromIp(fromIp).ToIp(toIp).FromDate(fromDate).ToDate(toDate).NulledBy(nulledBy).UnnulledBy(unnulledBy).TicketId(ticketId).ContractId(contractId).EquipmentId(equipmentId).Sort(sort).Execute()

Inspect null route history

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	fromIp := "fromIp_example" // string | Return only IPs greater or equal to the specified address (optional)
	toIp := "toIp_example" // string | Return only IPs lower or equal to the specified address (optional)
	fromDate := time.Now() // string | Return only null routes active after the specified date and time (optional)
	toDate := time.Now() // string | Return only null routes active before the specified date and time (optional)
	nulledBy := "nulledBy_example" // string | Filter by the email address of the user who created the null route (optional)
	unnulledBy := "unnulledBy_example" // string | Filter by the email address of the user who removed the null route (optional)
	ticketId := "ticketId_example" // string | Filter by the reference stored with the null route (optional)
	contractId := "contractId_example" // string | Filter by ID of the contract assigned to the IP at the time of null route creation (optional)
	equipmentId := "equipmentId_example" // string | Filter by ID of the server assigned to the IP at the time of null route creation (optional)
	sort := "sort_example" // string | Comma-separated list of sort field names. Prepend the field name with '-' for descending order. E.g. `sort=ip,-nullrouted`. Sortable field names are ip, nullRouted, reverseLookup. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.GetNullRouteHistoryList(context.Background()).Limit(limit).Offset(offset).FromIp(fromIp).ToIp(toIp).FromDate(fromDate).ToDate(toDate).NulledBy(nulledBy).UnnulledBy(unnulledBy).TicketId(ticketId).ContractId(contractId).EquipmentId(equipmentId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.GetNullRouteHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNullRouteHistoryList`: GetNullRouteHistoryListResult
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.GetNullRouteHistoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNullRouteHistoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **fromIp** | **string** | Return only IPs greater or equal to the specified address | 
 **toIp** | **string** | Return only IPs lower or equal to the specified address | 
 **fromDate** | **string** | Return only null routes active after the specified date and time | 
 **toDate** | **string** | Return only null routes active before the specified date and time | 
 **nulledBy** | **string** | Filter by the email address of the user who created the null route | 
 **unnulledBy** | **string** | Filter by the email address of the user who removed the null route | 
 **ticketId** | **string** | Filter by the reference stored with the null route | 
 **contractId** | **string** | Filter by ID of the contract assigned to the IP at the time of null route creation | 
 **equipmentId** | **string** | Filter by ID of the server assigned to the IP at the time of null route creation | 
 **sort** | **string** | Comma-separated list of sort field names. Prepend the field name with &#39;-&#39; for descending order. E.g. &#x60;sort&#x3D;ip,-nullrouted&#x60;. Sortable field names are ip, nullRouted, reverseLookup. | 

### Return type

[**GetNullRouteHistoryListResult**](GetNullRouteHistoryListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReverseLookupRecordList

> GetReverseLookupRecordListResult GetReverseLookupRecordList(ctx, ip).ReverseLookup(reverseLookup).Limit(limit).Offset(offset).Execute()

List reverse lookup records for an IPv6 range



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	ip := "192.0.2.1" // string | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength = 112
	reverseLookup := "mydomain1.example.com" // string | Filter by reverse lookup. (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.GetReverseLookupRecordList(context.Background(), ip).ReverseLookup(reverseLookup).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.GetReverseLookupRecordList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReverseLookupRecordList`: GetReverseLookupRecordListResult
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.GetReverseLookupRecordList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength &#x3D; 112 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReverseLookupRecordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseLookup** | **string** | Filter by reverse lookup. | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetReverseLookupRecordListResult**](GetReverseLookupRecordListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InspectIP

> Ip InspectIP(ctx, ip).Execute()

Inspect an IP

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	ip := "192.0.2.1" // string | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength = 112

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.InspectIP(context.Background(), ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.InspectIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InspectIP`: Ip
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.InspectIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength &#x3D; 112 | 

### Other Parameters

Other parameters are passed through a pointer to a apiInspectIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ip**](Ip.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InspectNullRouteHistory

> NullRoutedIP InspectNullRouteHistory(ctx, id).Execute()

Inspect null route history

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	id := "23234" // string | Null route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.InspectNullRouteHistory(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.InspectNullRouteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InspectNullRouteHistory`: NullRoutedIP
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.InspectNullRouteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Null route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInspectNullRouteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NullRoutedIP**](NullRoutedIP.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NullRouteIP

> NullRoutedIP NullRouteIP(ctx, ip).NullRouteIPOpts(nullRouteIPOpts).Execute()

Null route an IP

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	ip := "192.0.2.1" // string | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength = 112
	nullRouteIPOpts := *openapiclient.NewNullRouteIPOpts() // NullRouteIPOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.NullRouteIP(context.Background(), ip).NullRouteIPOpts(nullRouteIPOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.NullRouteIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullRouteIP`: NullRoutedIP
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.NullRouteIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength &#x3D; 112 | 

### Other Parameters

Other parameters are passed through a pointer to a apiNullRouteIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nullRouteIPOpts** | [**NullRouteIPOpts**](NullRouteIPOpts.md) |  | 

### Return type

[**NullRoutedIP**](NullRoutedIP.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveIPNullRoute

> RemoveIPNullRoute(ctx, ip).Execute()

Remove a null route

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	ip := "192.0.2.1" // string | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength = 112

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IpmgmtAPI.RemoveIPNullRoute(context.Background(), ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.RemoveIPNullRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength &#x3D; 112 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIPNullRouteRequest struct via the builder pattern


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


## UpdateIP

> Ip UpdateIP(ctx, ip).UpdateIPOpts(updateIPOpts).Execute()

Update an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	ip := "ip_example" // string | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32
	updateIPOpts := *openapiclient.NewUpdateIPOpts("mydomain1.example.com") // UpdateIPOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.UpdateIP(context.Background(), ip).UpdateIPOpts(updateIPOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.UpdateIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIP`: Ip
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.UpdateIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIPOpts** | [**UpdateIPOpts**](UpdateIPOpts.md) |  | 

### Return type

[**Ip**](Ip.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNullRoute

> NullRoutedIP UpdateNullRoute(ctx, id).UpdateNullRouteOpts(updateNullRouteOpts).Execute()

Update a null route

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	id := "23234" // string | Null route ID
	updateNullRouteOpts := *openapiclient.NewUpdateNullRouteOpts() // UpdateNullRouteOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.UpdateNullRoute(context.Background(), id).UpdateNullRouteOpts(updateNullRouteOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.UpdateNullRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNullRoute`: NullRoutedIP
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.UpdateNullRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Null route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNullRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNullRouteOpts** | [**UpdateNullRouteOpts**](UpdateNullRouteOpts.md) |  | 

### Return type

[**NullRoutedIP**](NullRoutedIP.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReverseLookupRecords

> UpdateReverseLookupRecordsResult UpdateReverseLookupRecords(ctx, ip).UpdateReverseLookupRecordsOpts(updateReverseLookupRecordsOpts).Execute()

Set or remove reverse lookup records for an IPv6 range



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func main() {
	ip := "192.0.2.1" // string | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength = 112
	updateReverseLookupRecordsOpts := *openapiclient.NewUpdateReverseLookupRecordsOpts([]openapiclient.ReverseLookup{*openapiclient.NewReverseLookup("Ip_example", "ReverseLookup_example")}) // UpdateReverseLookupRecordsOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpmgmtAPI.UpdateReverseLookupRecords(context.Background(), ip).UpdateReverseLookupRecordsOpts(updateReverseLookupRecordsOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpmgmtAPI.UpdateReverseLookupRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReverseLookupRecords`: UpdateReverseLookupRecordsResult
	fmt.Fprintf(os.Stdout, "Response from `IpmgmtAPI.UpdateReverseLookupRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | IP address or IP address with prefixLength {ip}_{prefix}. If prefixLength is not given, then we assume 32 (for IPv4) or 128 (for IPv6). PrefixLength is mandatory for IP range, for example, the IPv6 address range with prefixLength &#x3D; 112 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReverseLookupRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReverseLookupRecordsOpts** | [**UpdateReverseLookupRecordsOpts**](UpdateReverseLookupRecordsOpts.md) |  | 

### Return type

[**UpdateReverseLookupRecordsResult**](UpdateReverseLookupRecordsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

