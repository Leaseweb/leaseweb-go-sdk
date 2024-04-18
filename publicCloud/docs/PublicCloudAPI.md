# \PublicCloudAPI

All URIs are relative to *https://api.leaseweb.com/publicCloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachIso**](PublicCloudAPI.md#AttachIso) | **Post** /instances/{instanceId}/attachIso | Attach ISO to instance
[**CancelInstanceTermination**](PublicCloudAPI.md#CancelInstanceTermination) | **Post** /instances/{instanceId}/cancelTermination | Cancel instance termination
[**CreateFirewallRules**](PublicCloudAPI.md#CreateFirewallRules) | **Post** /instances/{instanceId}/firewall/batchCreate | Add firewall rules in batch
[**CredentialsDelete**](PublicCloudAPI.md#CredentialsDelete) | **Delete** /instances/{instanceId}/credentials | Delete all instance credentials
[**DeleteCredential**](PublicCloudAPI.md#DeleteCredential) | **Delete** /instances/{instanceId}/credentials/{type}/{username} | Delete credentials
[**DeleteFirewallRules**](PublicCloudAPI.md#DeleteFirewallRules) | **Post** /instances/{instanceId}/firewall/batchDelete | Delete firewall rules in batch
[**DetachIso**](PublicCloudAPI.md#DetachIso) | **Post** /instances/{instanceId}/detachIso | Detach ISO from instance
[**EditFirewallRules**](PublicCloudAPI.md#EditFirewallRules) | **Post** /instances/{instanceId}/firewall/batchUpdate | Edit firewall rules in batch
[**GetConsoleAccessToInstance**](PublicCloudAPI.md#GetConsoleAccessToInstance) | **Get** /instances/{instanceId}/console | Get console access
[**GetCredential**](PublicCloudAPI.md#GetCredential) | **Get** /instances/{instanceId}/credentials/{type}/{username} | Get credentials by type and username
[**GetCredentialList**](PublicCloudAPI.md#GetCredentialList) | **Get** /instances/{instanceId}/credentials | List credentials stored for instance
[**GetCredentialListByType**](PublicCloudAPI.md#GetCredentialListByType) | **Get** /instances/{instanceId}/credentials/{type} | Get credentials by type
[**GetExpenses**](PublicCloudAPI.md#GetExpenses) | **Get** /equipments/{equipmentId}/expenses | Get costs for a given month.
[**GetFirewallRuleList**](PublicCloudAPI.md#GetFirewallRuleList) | **Get** /instances/{instanceId}/firewall | List firewall rules
[**GetInstance**](PublicCloudAPI.md#GetInstance) | **Get** /instances/{instanceId} | Get instance details
[**GetInstanceTypeList**](PublicCloudAPI.md#GetInstanceTypeList) | **Get** /instanceTypes | List instance types
[**GetIsoList**](PublicCloudAPI.md#GetIsoList) | **Get** /isos | List available ISOs
[**GetMarketAppList**](PublicCloudAPI.md#GetMarketAppList) | **Get** /marketApps | Get marketplace apps
[**GetOperatingSystemList**](PublicCloudAPI.md#GetOperatingSystemList) | **Get** /operatingSystems | List all available Operating Systems
[**GetReinstallOsList**](PublicCloudAPI.md#GetReinstallOsList) | **Get** /instances/{instanceId}/reinstall/operatingSystems | List OSes available for reinstall
[**GetUpdateInstanceTypeList**](PublicCloudAPI.md#GetUpdateInstanceTypeList) | **Get** /instances/{instanceId}/instanceTypesUpdate | List available instance types for update
[**InstanceList**](PublicCloudAPI.md#InstanceList) | **Get** /instances | Get instance list
[**LaunchInstance**](PublicCloudAPI.md#LaunchInstance) | **Post** /instances | Launch instance
[**RebootInstance**](PublicCloudAPI.md#RebootInstance) | **Post** /instances/{instanceId}/reboot | Reboot instance
[**RegionsList**](PublicCloudAPI.md#RegionsList) | **Get** /regions | List regions
[**ResetPassword**](PublicCloudAPI.md#ResetPassword) | **Post** /instances/{instanceId}/resetPassword | Reset instance password
[**StartInstance**](PublicCloudAPI.md#StartInstance) | **Post** /instances/{instanceId}/start | Start instance
[**StopInstance**](PublicCloudAPI.md#StopInstance) | **Post** /instances/{instanceId}/stop | Stop instance
[**StoreCredential**](PublicCloudAPI.md#StoreCredential) | **Post** /instances/{instanceId}/credentials | Store credentials
[**TerminateInstance**](PublicCloudAPI.md#TerminateInstance) | **Delete** /instances/{instanceId} | Terminate instance
[**UpdateCredential**](PublicCloudAPI.md#UpdateCredential) | **Put** /instances/{instanceId}/credentials/{type}/{username} | Update credentials
[**UpdateInstance**](PublicCloudAPI.md#UpdateInstance) | **Put** /instances/{instanceId} | Update instance



## AttachIso

> AttachIso(ctx, instanceId).AttachIsoOpts(attachIsoOpts).Execute()

Attach ISO to instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	attachIsoOpts := *openapiclient.NewAttachIsoOpts("IsoId_example") // AttachIsoOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.AttachIso(context.Background(), instanceId).AttachIsoOpts(attachIsoOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.AttachIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachIsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachIsoOpts** | [**AttachIsoOpts**](AttachIsoOpts.md) |  | 

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


## CancelInstanceTermination

> CancelInstanceTermination(ctx, instanceId).Execute()

Cancel instance termination



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.CancelInstanceTermination(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.CancelInstanceTermination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelInstanceTerminationRequest struct via the builder pattern


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


## CreateFirewallRules

> CreateFirewallRulesCreatedResult CreateFirewallRules(ctx, instanceId).CreateFirewallRulesOpts(createFirewallRulesOpts).Execute()

Add firewall rules in batch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	createFirewallRulesOpts := *openapiclient.NewCreateFirewallRulesOpts([]openapiclient.CreateRule{*openapiclient.NewCreateRule("Protocol_example", NullableInt32(123), NullableInt32(123), int32(-1), int32(-1), "HTTP rule", "0.0.0.0/0")}) // CreateFirewallRulesOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.CreateFirewallRules(context.Background(), instanceId).CreateFirewallRulesOpts(createFirewallRulesOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.CreateFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirewallRules`: CreateFirewallRulesCreatedResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.CreateFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFirewallRulesOpts** | [**CreateFirewallRulesOpts**](CreateFirewallRulesOpts.md) |  | 

### Return type

[**CreateFirewallRulesCreatedResult**](CreateFirewallRulesCreatedResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialsDelete

> CredentialsDelete(ctx, instanceId).Execute()

Delete all instance credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.CredentialsDelete(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.CredentialsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCredentialsDeleteRequest struct via the builder pattern


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


## DeleteCredential

> DeleteCredential(ctx, instanceId, type_, username).Execute()

Delete credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.DeleteCredential(context.Background(), instanceId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeleteCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialRequest struct via the builder pattern


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


## DeleteFirewallRules

> DeleteFirewallRulesResult DeleteFirewallRules(ctx, instanceId).DeleteFirewallRulesOpts(deleteFirewallRulesOpts).Execute()

Delete firewall rules in batch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	deleteFirewallRulesOpts := *openapiclient.NewDeleteFirewallRulesOpts([]openapiclient.DeletionRule{*openapiclient.NewDeletionRule("7e59b33d-05f3-4078-b251-c7831ae8fe14")}) // DeleteFirewallRulesOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.DeleteFirewallRules(context.Background(), instanceId).DeleteFirewallRulesOpts(deleteFirewallRulesOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeleteFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRules`: DeleteFirewallRulesResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.DeleteFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteFirewallRulesOpts** | [**DeleteFirewallRulesOpts**](DeleteFirewallRulesOpts.md) |  | 

### Return type

[**DeleteFirewallRulesResult**](DeleteFirewallRulesResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachIso

> DetachIso(ctx, instanceId).Execute()

Detach ISO from instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.DetachIso(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DetachIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachIsoRequest struct via the builder pattern


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


## EditFirewallRules

> EditFirewallRulesResult EditFirewallRules(ctx, instanceId).EditFirewallRulesOpts(editFirewallRulesOpts).Execute()

Edit firewall rules in batch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	editFirewallRulesOpts := *openapiclient.NewEditFirewallRulesOpts([]openapiclient.EditRule{*openapiclient.NewEditRule("Protocol_example", NullableInt32(123), NullableInt32(123), int32(-1), int32(-1), "7e59b33d-05f3-4078-b251-c7831ae8fe14", "HTTP rule", "0.0.0.0/0")}) // EditFirewallRulesOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.EditFirewallRules(context.Background(), instanceId).EditFirewallRulesOpts(editFirewallRulesOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.EditFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditFirewallRules`: EditFirewallRulesResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.EditFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editFirewallRulesOpts** | [**EditFirewallRulesOpts**](EditFirewallRulesOpts.md) |  | 

### Return type

[**EditFirewallRulesResult**](EditFirewallRulesResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsoleAccessToInstance

> GetConsoleAccessToInstanceResult GetConsoleAccessToInstance(ctx, instanceId).Execute()

Get console access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetConsoleAccessToInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetConsoleAccessToInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsoleAccessToInstance`: GetConsoleAccessToInstanceResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetConsoleAccessToInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsoleAccessToInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConsoleAccessToInstanceResult**](GetConsoleAccessToInstanceResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredential

> GetCredentialResult GetCredential(ctx, instanceId, type_, username).Execute()

Get credentials by type and username



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetCredential(context.Background(), instanceId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: GetCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetCredentialResult**](GetCredentialResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentialList

> GetCredentialListResult GetCredentialList(ctx, instanceId).Execute()

List credentials stored for instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetCredentialList(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialList`: GetCredentialListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCredentialListResult**](GetCredentialListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentialListByType

> GetCredentialListByTypeResult GetCredentialListByType(ctx, instanceId, type_).Execute()

Get credentials by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetCredentialListByType(context.Background(), instanceId, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialListByType`: GetCredentialListByTypeResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetCredentialListByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialListByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCredentialListByTypeResult**](GetCredentialListByTypeResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenses

> GetExpensesResult GetExpenses(ctx, equipmentId).From(from).To(to).Execute()

Get costs for a given month.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	equipmentId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Equipment's UUID
	from := time.Now() // string | Start date of the period to get costs for. It must be the first day of the month
	to := time.Now() // string | End date of the period to get costs for. This date needs to be exactly one month after the 'from' date. If this value is not passed, it will be calculated based on 'from' parameter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetExpenses(context.Background(), equipmentId).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetExpenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenses`: GetExpensesResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetExpenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **string** | Equipment&#39;s UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | Start date of the period to get costs for. It must be the first day of the month | 
 **to** | **string** | End date of the period to get costs for. This date needs to be exactly one month after the &#39;from&#39; date. If this value is not passed, it will be calculated based on &#39;from&#39; parameter | 

### Return type

[**GetExpensesResult**](GetExpensesResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRuleList

> GetFirewallRuleListResult GetFirewallRuleList(ctx, instanceId).Limit(limit).Offset(offset).Execute()

List firewall rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetFirewallRuleList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetFirewallRuleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallRuleList`: GetFirewallRuleListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetFirewallRuleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetFirewallRuleListResult**](GetFirewallRuleListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> GetInstanceResult GetInstance(ctx, instanceId).Execute()

Get instance details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: GetInstanceResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInstanceResult**](GetInstanceResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeList

> GetInstanceTypeListResult GetInstanceTypeList(ctx).Region(region).Limit(limit).Offset(offset).Execute()

List instance types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	region := "region_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetInstanceTypeList(context.Background()).Region(region).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetInstanceTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeList`: GetInstanceTypeListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetInstanceTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**GetInstanceTypeListResult**](GetInstanceTypeListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsoList

> GetIsoListResult GetIsoList(ctx).Limit(limit).Offset(offset).Execute()

List available ISOs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetIsoList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetIsoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsoList`: GetIsoListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetIsoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIsoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetIsoListResult**](GetIsoListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketAppList

> GetMarketAppListResult GetMarketAppList(ctx).Execute()

Get marketplace apps



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetMarketAppList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetMarketAppList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketAppList`: GetMarketAppListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetMarketAppList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketAppListRequest struct via the builder pattern


### Return type

[**GetMarketAppListResult**](GetMarketAppListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatingSystemList

> GetOperatingSystemListResult GetOperatingSystemList(ctx).Limit(limit).Offset(offset).Execute()

List all available Operating Systems

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetOperatingSystemList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetOperatingSystemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystemList`: GetOperatingSystemListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetOperatingSystemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatingSystemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetOperatingSystemListResult**](GetOperatingSystemListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReinstallOsList

> []OperatingSystem GetReinstallOsList(ctx, instanceId).Execute()

List OSes available for reinstall



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetReinstallOsList(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetReinstallOsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReinstallOsList`: []OperatingSystem
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetReinstallOsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReinstallOsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OperatingSystem**](OperatingSystem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateInstanceTypeList

> GetUpdateInstanceTypeListResult GetUpdateInstanceTypeList(ctx, instanceId).Limit(limit).Offset(offset).Execute()

List available instance types for update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetUpdateInstanceTypeList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetUpdateInstanceTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateInstanceTypeList`: GetUpdateInstanceTypeListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetUpdateInstanceTypeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateInstanceTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**GetUpdateInstanceTypeListResult**](GetUpdateInstanceTypeListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceList

> GetInstanceListResult InstanceList(ctx).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).Filter(filter).ContractType(contractType).State(state).Region(region).Type_(type_).Execute()

Get instance list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	ip := "10.0.0.1" // string |  (optional)
	reference := "my-webserver" // string |  (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	filter := "my-webserver" // string | Using this parameter, you can filter by id, reference or IP. (optional)
	contractType := "HOURLY" // string |  (optional)
	state := "RUNNING" // string |  (optional)
	region := "eu-west-3" // string | Available regions can be found using the List Regions endpoint. (optional)
	type_ := "lsw.c3.xlarge" // string | Available instance types can be found using the List Instance Types endpoint. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.InstanceList(context.Background()).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).Filter(filter).ContractType(contractType).State(state).Region(region).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.InstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceList`: GetInstanceListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.InstanceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **ip** | **string** |  | 
 **reference** | **string** |  | 
 **id** | **string** |  | 
 **filter** | **string** | Using this parameter, you can filter by id, reference or IP. | 
 **contractType** | **string** |  | 
 **state** | **string** |  | 
 **region** | **string** | Available regions can be found using the List Regions endpoint. | 
 **type_** | **string** | Available instance types can be found using the List Instance Types endpoint. | 

### Return type

[**GetInstanceListResult**](GetInstanceListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchInstance

> LaunchInstanceCreatedResult LaunchInstance(ctx).LaunchInstanceOpts(launchInstanceOpts).Execute()

Launch instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	launchInstanceOpts := *openapiclient.NewLaunchInstanceOpts("Region_example", "UBUNTU_22_04_64BIT", "ContractType_example", int32(123), int32(123), "RootDiskStorageType_example") // LaunchInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.LaunchInstance(context.Background()).LaunchInstanceOpts(launchInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.LaunchInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchInstance`: LaunchInstanceCreatedResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.LaunchInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLaunchInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **launchInstanceOpts** | [**LaunchInstanceOpts**](LaunchInstanceOpts.md) |  | 

### Return type

[**LaunchInstanceCreatedResult**](LaunchInstanceCreatedResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootInstance

> RebootInstance(ctx, instanceId).Execute()

Reboot instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.RebootInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.RebootInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootInstanceRequest struct via the builder pattern


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


## RegionsList

> GetRegionListResult RegionsList(ctx).Limit(limit).Offset(offset).Execute()

List regions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.RegionsList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.RegionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsList`: GetRegionListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.RegionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetRegionListResult**](GetRegionListResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword(ctx, instanceId).Execute()

Reset instance password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.ResetPassword(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


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


## StartInstance

> StartInstance(ctx, instanceId).Execute()

Start instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.StartInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.StartInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstanceRequest struct via the builder pattern


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


## StopInstance

> StopInstance(ctx, instanceId).Execute()

Stop instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.StopInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.StopInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstanceRequest struct via the builder pattern


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


## StoreCredential

> StoreCredentialResult StoreCredential(ctx, instanceId).StoreCredentialOpts(storeCredentialOpts).Execute()

Store credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	storeCredentialOpts := *openapiclient.NewStoreCredentialOpts(openapiclient.credentialType("OPERATING_SYSTEM"), "Username_example", "Password_example") // StoreCredentialOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.StoreCredential(context.Background(), instanceId).StoreCredentialOpts(storeCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.StoreCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreCredential`: StoreCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.StoreCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeCredentialOpts** | [**StoreCredentialOpts**](StoreCredentialOpts.md) |  | 

### Return type

[**StoreCredentialResult**](StoreCredentialResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateInstance

> TerminateInstance(ctx, instanceId).Execute()

Terminate instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.TerminateInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.TerminateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateInstanceRequest struct via the builder pattern


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


## UpdateCredential

> UpdateCredentialResult UpdateCredential(ctx, instanceId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()

Update credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username
	updateCredentialOpts := *openapiclient.NewUpdateCredentialOpts("Password_example") // UpdateCredentialOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.UpdateCredential(context.Background(), instanceId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.UpdateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredential`: UpdateCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateCredentialOpts** | [**UpdateCredentialOpts**](UpdateCredentialOpts.md) |  | 

### Return type

[**UpdateCredentialResult**](UpdateCredentialResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstance

> UpdateInstanceResult UpdateInstance(ctx, instanceId).UpdateInstanceOpts(updateInstanceOpts).Execute()

Update instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	updateInstanceOpts := *openapiclient.NewUpdateInstanceOpts() // UpdateInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.UpdateInstance(context.Background(), instanceId).UpdateInstanceOpts(updateInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstance`: UpdateInstanceResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInstanceOpts** | [**UpdateInstanceOpts**](UpdateInstanceOpts.md) |  | 

### Return type

[**UpdateInstanceResult**](UpdateInstanceResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

