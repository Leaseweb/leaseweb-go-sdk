# \DnsAPI

All URIs are relative to *https://api.leaseweb.com/hosting/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceRecordSet**](DnsAPI.md#CreateResourceRecordSet) | **Post** /domains/{domainName}/resourceRecordSets | Create a resource record set
[**DeleteResourceRecordSet**](DnsAPI.md#DeleteResourceRecordSet) | **Delete** /domains/{domainName}/resourceRecordSets/{name}/{type} | Delete a resource record set
[**DeleteResourceRecordSets**](DnsAPI.md#DeleteResourceRecordSets) | **Delete** /domains/{domainName}/resourceRecordSets | Delete resource record sets
[**ExportResourceRecordSets**](DnsAPI.md#ExportResourceRecordSets) | **Get** /domains/{domainName}/resourceRecordSets/import | Export dns records as a bind file content
[**GetResourceRecordSet**](DnsAPI.md#GetResourceRecordSet) | **Get** /domains/{domainName}/resourceRecordSets/{name}/{type} | Inspect resource record set
[**GetResourceRecordSetList**](DnsAPI.md#GetResourceRecordSetList) | **Get** /domains/{domainName}/resourceRecordSets | List resource record sets
[**ImportResourceRecordSets**](DnsAPI.md#ImportResourceRecordSets) | **Post** /domains/{domainName}/resourceRecordSets/import | Import dns records from bind file content
[**UpdateResourceRecordSet**](DnsAPI.md#UpdateResourceRecordSet) | **Put** /domains/{domainName}/resourceRecordSets/{name}/{type} | Update a resource record set
[**UpdateResourceRecordSets**](DnsAPI.md#UpdateResourceRecordSets) | **Put** /domains/{domainName}/resourceRecordSets | Update resource record sets
[**ValidateResourceRecordSet**](DnsAPI.md#ValidateResourceRecordSet) | **Post** /domains/{domainName}/resourceRecordSets/validateSet | Validate a resource record set
[**ValidateZone**](DnsAPI.md#ValidateZone) | **Post** /domains/{domainName}/validateZone | Validate zone



## CreateResourceRecordSet

> ResourceRecordSetDetails CreateResourceRecordSet(ctx, domainName).ResourceRecordSet(resourceRecordSet).Execute()

Create a resource record set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	resourceRecordSet := *openapiclient.NewResourceRecordSet("example.com.", openapiclient.resourceRecordSetType("A"), []string{"85.17.150.50"}, openapiclient.ttl(60)) // ResourceRecordSet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.CreateResourceRecordSet(context.Background(), domainName).ResourceRecordSet(resourceRecordSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.CreateResourceRecordSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResourceRecordSet`: ResourceRecordSetDetails
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.CreateResourceRecordSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRecordSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceRecordSet** | [**ResourceRecordSet**](ResourceRecordSet.md) |  | 

### Return type

[**ResourceRecordSetDetails**](ResourceRecordSetDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json
- **Accept**: application/json, application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceRecordSet

> DeleteResourceRecordSet(ctx, domainName, name, type_).Execute()

Delete a resource record set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	name := "example.com." // string | Record name
	type_ := "A" // string | Record type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsAPI.DeleteResourceRecordSet(context.Background(), domainName, name, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.DeleteResourceRecordSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 
**name** | **string** | Record name | 
**type_** | **string** | Record type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRecordSetRequest struct via the builder pattern


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


## DeleteResourceRecordSets

> DeleteResourceRecordSets(ctx, domainName).Execute()

Delete resource record sets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsAPI.DeleteResourceRecordSets(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.DeleteResourceRecordSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRecordSetsRequest struct via the builder pattern


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


## ExportResourceRecordSets

> ExportResourceRecordSets(ctx, domainName).Execute()

Export dns records as a bind file content

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsAPI.ExportResourceRecordSets(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.ExportResourceRecordSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportResourceRecordSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceRecordSet

> ResourceRecordSetDetails GetResourceRecordSet(ctx, domainName, name, type_).Execute()

Inspect resource record set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	name := "example.com." // string | Record name
	type_ := "A" // string | Record type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.GetResourceRecordSet(context.Background(), domainName, name, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.GetResourceRecordSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceRecordSet`: ResourceRecordSetDetails
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.GetResourceRecordSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 
**name** | **string** | Record name | 
**type_** | **string** | Record type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRecordSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceRecordSetDetails**](ResourceRecordSetDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceRecordSetList

> GetResourceRecordSetListResult GetResourceRecordSetList(ctx, domainName).Execute()

List resource record sets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.GetResourceRecordSetList(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.GetResourceRecordSetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceRecordSetList`: GetResourceRecordSetListResult
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.GetResourceRecordSetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRecordSetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetResourceRecordSetListResult**](GetResourceRecordSetListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportResourceRecordSets

> ImportResourceRecordSetsResult ImportResourceRecordSets(ctx, domainName).ImportResourceRecordSetsOpts(importResourceRecordSetsOpts).Execute()

Import dns records from bind file content

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	importResourceRecordSetsOpts := *openapiclient.NewImportResourceRecordSetsOpts("["$ORIGIN example.com. \\r\\n$TTL 86400 \\r\\n\\tIN\\tMX\\t10\\tmail.example.com.       \\r\\n\\r\\n\\t\\r\\ndns1\\tIN\\tA\\t10.0.1.1"]") // ImportResourceRecordSetsOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.ImportResourceRecordSets(context.Background(), domainName).ImportResourceRecordSetsOpts(importResourceRecordSetsOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.ImportResourceRecordSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportResourceRecordSets`: ImportResourceRecordSetsResult
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.ImportResourceRecordSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportResourceRecordSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importResourceRecordSetsOpts** | [**ImportResourceRecordSetsOpts**](ImportResourceRecordSetsOpts.md) |  | 

### Return type

[**ImportResourceRecordSetsResult**](ImportResourceRecordSetsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceRecordSet

> ResourceRecordSetDetails UpdateResourceRecordSet(ctx, domainName, name, type_).UpdateResourceRecordSetOpts(updateResourceRecordSetOpts).Execute()

Update a resource record set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	name := "example.com." // string | Record name
	type_ := "A" // string | Record type
	updateResourceRecordSetOpts := *openapiclient.NewUpdateResourceRecordSetOpts([]string{"85.17.150.50"}, openapiclient.ttl(60)) // UpdateResourceRecordSetOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.UpdateResourceRecordSet(context.Background(), domainName, name, type_).UpdateResourceRecordSetOpts(updateResourceRecordSetOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.UpdateResourceRecordSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceRecordSet`: ResourceRecordSetDetails
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.UpdateResourceRecordSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 
**name** | **string** | Record name | 
**type_** | **string** | Record type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRecordSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateResourceRecordSetOpts** | [**UpdateResourceRecordSetOpts**](UpdateResourceRecordSetOpts.md) |  | 

### Return type

[**ResourceRecordSetDetails**](ResourceRecordSetDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json
- **Accept**: application/json, application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceRecordSets

> ResourceRecordSetDetails UpdateResourceRecordSets(ctx, domainName).ResultRecordSets(resultRecordSets).Execute()

Update resource record sets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	resultRecordSets := *openapiclient.NewResultRecordSets([]openapiclient.ResourceRecordSet{*openapiclient.NewResourceRecordSet("example.com.", openapiclient.resourceRecordSetType("A"), []string{"85.17.150.50"}, openapiclient.ttl(60))}) // ResultRecordSets |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.UpdateResourceRecordSets(context.Background(), domainName).ResultRecordSets(resultRecordSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.UpdateResourceRecordSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceRecordSets`: ResourceRecordSetDetails
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.UpdateResourceRecordSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRecordSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resultRecordSets** | [**ResultRecordSets**](ResultRecordSets.md) |  | 

### Return type

[**ResourceRecordSetDetails**](ResourceRecordSetDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json
- **Accept**: application/json, application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateResourceRecordSet

> InfoMessageResult ValidateResourceRecordSet(ctx, domainName).ResourceRecordSet(resourceRecordSet).Execute()

Validate a resource record set



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	resourceRecordSet := *openapiclient.NewResourceRecordSet("example.com.", openapiclient.resourceRecordSetType("A"), []string{"85.17.150.50"}, openapiclient.ttl(60)) // ResourceRecordSet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.ValidateResourceRecordSet(context.Background(), domainName).ResourceRecordSet(resourceRecordSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.ValidateResourceRecordSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateResourceRecordSet`: InfoMessageResult
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.ValidateResourceRecordSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateResourceRecordSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceRecordSet** | [**ResourceRecordSet**](ResourceRecordSet.md) |  | 

### Return type

[**InfoMessageResult**](InfoMessageResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json - Validate resource with content, application/json - Validate resource with geo content
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateZone

> ValidateZoneResult ValidateZone(ctx, domainName).ResultRecordSets(resultRecordSets).Execute()

Validate zone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func main() {
	domainName := "example.com" // string | Domain name
	resultRecordSets := *openapiclient.NewResultRecordSets([]openapiclient.ResourceRecordSet{*openapiclient.NewResourceRecordSet("example.com.", openapiclient.resourceRecordSetType("A"), []string{"85.17.150.50"}, openapiclient.ttl(60))}) // ResultRecordSets |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsAPI.ValidateZone(context.Background(), domainName).ResultRecordSets(resultRecordSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsAPI.ValidateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateZone`: ValidateZoneResult
	fmt.Fprintf(os.Stdout, "Response from `DnsAPI.ValidateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resultRecordSets** | [**ResultRecordSets**](ResultRecordSets.md) |  | 

### Return type

[**ValidateZoneResult**](ValidateZoneResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

