# \InvoiceAPI

All URIs are relative to *https://api.leaseweb.com/invoices/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoice**](InvoiceAPI.md#GetInvoice) | **Get** /invoices/{invoiceId} | Inspect an invoice
[**GetInvoiceList**](InvoiceAPI.md#GetInvoiceList) | **Get** /invoices | List invoices
[**GetInvoicePdf**](InvoiceAPI.md#GetInvoicePdf) | **Get** /invoices/{invoiceId}/pdf | Get invoice PDF
[**GetProforma**](InvoiceAPI.md#GetProforma) | **Get** /invoices/proforma | Pro Forma



## GetInvoice

> GetInvoiceResult GetInvoice(ctx, invoiceId).Execute()

Inspect an invoice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/invoice"
)

func main() {
	invoiceId := "00000001" // string | Invoice Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.GetInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoice`: GetInvoiceResult
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Invoice Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInvoiceResult**](GetInvoiceResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceList

> GetInvoiceListResult GetInvoiceList(ctx).Limit(limit).Offset(offset).Execute()

List invoices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/invoice"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.GetInvoiceList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetInvoiceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceList`: GetInvoiceListResult
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.GetInvoiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetInvoiceListResult**](GetInvoiceListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoicePdf

> GetInvoicePdf(ctx, invoiceId).Execute()

Get invoice PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/invoice"
)

func main() {
	invoiceId := "00000001" // string | Invoice Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceAPI.GetInvoicePdf(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetInvoicePdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Invoice Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePdfRequest struct via the builder pattern


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


## GetProforma

> GetProformaResult GetProforma(ctx).Limit(limit).Offset(offset).Execute()

Pro Forma



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/invoice"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.GetProforma(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetProforma``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProforma`: GetProformaResult
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.GetProforma`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProformaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetProformaResult**](GetProformaResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

