# \AggregationPackAPI

All URIs are relative to *https://api.leaseweb.com/bareMetals/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregationPack**](AggregationPackAPI.md#GetAggregationPack) | **Get** /aggregationPacks/{aggregationPackId} | Get aggregation pack
[**GetAggregationPackList**](AggregationPackAPI.md#GetAggregationPackList) | **Get** /aggregationPacks | List aggregation packs



## GetAggregationPack

> AggregationPack GetAggregationPack(ctx, aggregationPackId).Execute()

Get aggregation pack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/aggregationPack"
)

func main() {
	aggregationPackId := "aggregationPackId_example" // string | The ID of aggregation pack

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationPackAPI.GetAggregationPack(context.Background(), aggregationPackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationPackAPI.GetAggregationPack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAggregationPack`: AggregationPack
	fmt.Fprintf(os.Stdout, "Response from `AggregationPackAPI.GetAggregationPack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregationPackId** | **string** | The ID of aggregation pack | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AggregationPack**](AggregationPack.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAggregationPackList

> GetAggregationPackListResult GetAggregationPackList(ctx).Limit(limit).Offset(offset).Execute()

List aggregation packs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/aggregationPack"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationPackAPI.GetAggregationPackList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationPackAPI.GetAggregationPackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAggregationPackList`: GetAggregationPackListResult
	fmt.Fprintf(os.Stdout, "Response from `AggregationPackAPI.GetAggregationPackList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationPackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetAggregationPackListResult**](GetAggregationPackListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

