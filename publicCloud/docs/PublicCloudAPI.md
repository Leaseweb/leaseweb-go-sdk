# \PublicCloudAPI

All URIs are relative to *https://api.leaseweb.com/publicCloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToPrivateNetwork**](PublicCloudAPI.md#AddToPrivateNetwork) | **Put** /instances/{instanceId}/addToPrivateNetwork | Add instance to Private Network
[**AttachIso**](PublicCloudAPI.md#AttachIso) | **Post** /instances/{instanceId}/attachIso | Attach ISO to instance
[**CancelInstanceTermination**](PublicCloudAPI.md#CancelInstanceTermination) | **Post** /instances/{instanceId}/cancelTermination | Cancel instance termination
[**CreateAutoScalingGroup**](PublicCloudAPI.md#CreateAutoScalingGroup) | **Post** /autoScalingGroups | Create Auto Scaling Group
[**CreateLoadBalancerListener**](PublicCloudAPI.md#CreateLoadBalancerListener) | **Post** /loadBalancers/{loadBalancerId}/listeners | Create listener
[**CreateSnapshot**](PublicCloudAPI.md#CreateSnapshot) | **Post** /instances/{instanceId}/snapshots | Create instance snapshot
[**DeleteAutoScalingGroup**](PublicCloudAPI.md#DeleteAutoScalingGroup) | **Delete** /autoScalingGroups/{autoScalingGroupId} | Delete Auto Scaling Group
[**DeleteCredential**](PublicCloudAPI.md#DeleteCredential) | **Delete** /instances/{instanceId}/credentials/{type}/{username} | Delete credentials
[**DeleteCredentials**](PublicCloudAPI.md#DeleteCredentials) | **Delete** /instances/{instanceId}/credentials | Delete all instance credentials
[**DeleteLoadBalancerListener**](PublicCloudAPI.md#DeleteLoadBalancerListener) | **Delete** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Delete load balancer listener
[**DeleteSnapshot**](PublicCloudAPI.md#DeleteSnapshot) | **Delete** /instances/{instanceId}/snapshots/{snapshotId} | Delete instance snapshot
[**DeregisterAutoScalingGroupLoadBalancer**](PublicCloudAPI.md#DeregisterAutoScalingGroupLoadBalancer) | **Post** /autoScalingGroups/{autoScalingGroupId}/deregisterLoadBalancer | Deregister Load balancer
[**DeregisterLoadBalancerTargets**](PublicCloudAPI.md#DeregisterLoadBalancerTargets) | **Post** /loadBalancers/{loadBalancerId}/deregisterTargets | Deregister targets
[**DetachIso**](PublicCloudAPI.md#DetachIso) | **Post** /instances/{instanceId}/detachIso | Detach ISO from instance
[**GetAutoScalingGroup**](PublicCloudAPI.md#GetAutoScalingGroup) | **Get** /autoScalingGroups/{autoScalingGroupId} | Get Auto Scaling Group details
[**GetAutoScalingGroupInstanceList**](PublicCloudAPI.md#GetAutoScalingGroupInstanceList) | **Get** /autoScalingGroups/{autoScalingGroupId}/instances | Get list of instances belonging to an Auto Scaling Group
[**GetAutoScalingGroupList**](PublicCloudAPI.md#GetAutoScalingGroupList) | **Get** /autoScalingGroups | Get Auto Scaling Group list
[**GetConsoleAccessToInstance**](PublicCloudAPI.md#GetConsoleAccessToInstance) | **Get** /instances/{instanceId}/console | Get console access
[**GetCpuMetrics**](PublicCloudAPI.md#GetCpuMetrics) | **Get** /instances/{instanceId}/metrics/cpu | Get instance CPU metrics
[**GetCredential**](PublicCloudAPI.md#GetCredential) | **Get** /instances/{instanceId}/credentials/{type}/{username} | Get credentials by type and username
[**GetCredentialList**](PublicCloudAPI.md#GetCredentialList) | **Get** /instances/{instanceId}/credentials | List credentials stored for instance
[**GetCredentialListByType**](PublicCloudAPI.md#GetCredentialListByType) | **Get** /instances/{instanceId}/credentials/{type} | Get credentials by type
[**GetDataTrafficMetrics**](PublicCloudAPI.md#GetDataTrafficMetrics) | **Get** /instances/{instanceId}/metrics/datatraffic | Get instance data traffic metrics
[**GetExpenses**](PublicCloudAPI.md#GetExpenses) | **Get** /equipments/{equipmentId}/expenses | Get costs for a given month.
[**GetImageList**](PublicCloudAPI.md#GetImageList) | **Get** /images | List all available Images
[**GetInstance**](PublicCloudAPI.md#GetInstance) | **Get** /instances/{instanceId} | Get instance details
[**GetInstanceList**](PublicCloudAPI.md#GetInstanceList) | **Get** /instances | Get instance list
[**GetInstanceTypeList**](PublicCloudAPI.md#GetInstanceTypeList) | **Get** /instanceTypes | List instance types
[**GetIp**](PublicCloudAPI.md#GetIp) | **Get** /instances/{instanceId}/ips/{ip} | Get details about an instance&#39;s IP
[**GetIpList**](PublicCloudAPI.md#GetIpList) | **Get** /instances/{instanceId}/ips | List instance&#39;s IPs
[**GetIsoList**](PublicCloudAPI.md#GetIsoList) | **Get** /isos | List available ISOs
[**GetLoadBalancer**](PublicCloudAPI.md#GetLoadBalancer) | **Get** /loadBalancers/{loadBalancerId} | Get load balancer details
[**GetLoadBalancerList**](PublicCloudAPI.md#GetLoadBalancerList) | **Get** /loadBalancers | Get load balancer list
[**GetLoadBalancerListener**](PublicCloudAPI.md#GetLoadBalancerListener) | **Get** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Get listener details
[**GetLoadBalancerTargetList**](PublicCloudAPI.md#GetLoadBalancerTargetList) | **Get** /loadBalancers/{loadBalancerId}/targets | List registered targets
[**GetMarketAppList**](PublicCloudAPI.md#GetMarketAppList) | **Get** /marketApps | Get marketplace apps
[**GetRegionList**](PublicCloudAPI.md#GetRegionList) | **Get** /regions | List regions
[**GetReinstallOsList**](PublicCloudAPI.md#GetReinstallOsList) | **Get** /instances/{instanceId}/reinstall/images | List OSes available for reinstall
[**GetSnapshot**](PublicCloudAPI.md#GetSnapshot) | **Get** /instances/{instanceId}/snapshots/{snapshotId} | Get snapshot detail
[**GetSnapshotList**](PublicCloudAPI.md#GetSnapshotList) | **Get** /instances/{instanceId}/snapshots | List snapshots
[**GetUpdateInstanceTypeList**](PublicCloudAPI.md#GetUpdateInstanceTypeList) | **Get** /instances/{instanceId}/instanceTypesUpdate | List available instance types for update
[**LaunchInstance**](PublicCloudAPI.md#LaunchInstance) | **Post** /instances | Launch instance
[**LaunchLoadBalancer**](PublicCloudAPI.md#LaunchLoadBalancer) | **Post** /loadBalancers | Launch Load balancer
[**NullRouteIp**](PublicCloudAPI.md#NullRouteIp) | **Post** /instances/{instanceId}/ips/{ip}/null | Null route IP
[**RebootInstance**](PublicCloudAPI.md#RebootInstance) | **Post** /instances/{instanceId}/reboot | Reboot instance
[**RegisterAutoScalingGroupLoadBalancer**](PublicCloudAPI.md#RegisterAutoScalingGroupLoadBalancer) | **Post** /autoScalingGroups/{autoScalingGroupId}/registerLoadBalancer | Register Load balancer
[**RegisterLoadBalancerTargets**](PublicCloudAPI.md#RegisterLoadBalancerTargets) | **Post** /loadBalancers/{loadBalancerId}/registerTargets | Register targets
[**ReinstallInstance**](PublicCloudAPI.md#ReinstallInstance) | **Put** /instances/{instanceId}/reinstall | Reinstall instance
[**RemoveFromPrivateNetwork**](PublicCloudAPI.md#RemoveFromPrivateNetwork) | **Delete** /instances/{instanceId}/removeFromPrivateNetwork | Remove instance from Private Network
[**RemoveIpNullRoute**](PublicCloudAPI.md#RemoveIpNullRoute) | **Post** /instances/{instanceId}/ips/{ip}/unnull | Remove an IP null route
[**ResetPassword**](PublicCloudAPI.md#ResetPassword) | **Post** /instances/{instanceId}/resetPassword | Reset instance password
[**RestoreSnapshot**](PublicCloudAPI.md#RestoreSnapshot) | **Put** /instances/{instanceId}/snapshots/{snapshotId} | Restore instance snapshot
[**StartInstance**](PublicCloudAPI.md#StartInstance) | **Post** /instances/{instanceId}/start | Start instance
[**StopInstance**](PublicCloudAPI.md#StopInstance) | **Post** /instances/{instanceId}/stop | Stop instance
[**StoreCredential**](PublicCloudAPI.md#StoreCredential) | **Post** /instances/{instanceId}/credentials | Store credentials
[**TerminateInstance**](PublicCloudAPI.md#TerminateInstance) | **Delete** /instances/{instanceId} | Terminate instance
[**TerminateLoadBalancer**](PublicCloudAPI.md#TerminateLoadBalancer) | **Delete** /loadBalancers/{loadBalancerId} | Delete load balancer
[**UpdateAutoScalingGroup**](PublicCloudAPI.md#UpdateAutoScalingGroup) | **Put** /autoScalingGroups/{autoScalingGroupId} | Update Auto Scaling Group
[**UpdateCredential**](PublicCloudAPI.md#UpdateCredential) | **Put** /instances/{instanceId}/credentials/{type}/{username} | Update credentials
[**UpdateInstance**](PublicCloudAPI.md#UpdateInstance) | **Put** /instances/{instanceId} | Update instance
[**UpdateIp**](PublicCloudAPI.md#UpdateIp) | **Put** /instances/{instanceId}/ips/{ip} | Update IP
[**UpdateLoadBalancer**](PublicCloudAPI.md#UpdateLoadBalancer) | **Put** /loadBalancers/{loadBalancerId} | Update load balancer
[**UpdateLoadBalancerListener**](PublicCloudAPI.md#UpdateLoadBalancerListener) | **Put** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Update a listener



## AddToPrivateNetwork

> AddToPrivateNetwork(ctx, instanceId).Execute()

Add instance to Private Network



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
	r, err := apiClient.PublicCloudAPI.AddToPrivateNetwork(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.AddToPrivateNetwork``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddToPrivateNetworkRequest struct via the builder pattern


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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutoScalingGroup

> AutoScalingGroupDetails CreateAutoScalingGroup(ctx).CreateAutoScalingGroupOpts(createAutoScalingGroupOpts).Execute()

Create Auto Scaling Group



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
	createAutoScalingGroupOpts := *openapiclient.NewCreateAutoScalingGroupOpts("InstanceId_example", "Reference_example", "Type_example") // CreateAutoScalingGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.CreateAutoScalingGroup(context.Background()).CreateAutoScalingGroupOpts(createAutoScalingGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.CreateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.CreateAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutoScalingGroupOpts** | [**CreateAutoScalingGroupOpts**](CreateAutoScalingGroupOpts.md) |  | 

### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancerListener

> LoadBalancerListener CreateLoadBalancerListener(ctx, loadBalancerId).LoadBalancerListenerOpts(loadBalancerListenerOpts).Execute()

Create listener



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	loadBalancerListenerOpts := *openapiclient.NewLoadBalancerListenerOpts("Protocol_example", int32(123)) // LoadBalancerListenerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.CreateLoadBalancerListener(context.Background(), loadBalancerId).LoadBalancerListenerOpts(loadBalancerListenerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.CreateLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoadBalancerListener`: LoadBalancerListener
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.CreateLoadBalancerListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadBalancerListenerOpts** | [**LoadBalancerListenerOpts**](LoadBalancerListenerOpts.md) |  | 

### Return type

[**LoadBalancerListener**](LoadBalancerListener.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshot

> CreateSnapshot(ctx, instanceId).Execute()

Create instance snapshot



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
	r, err := apiClient.PublicCloudAPI.CreateSnapshot(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.CreateSnapshot``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


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


## DeleteAutoScalingGroup

> DeleteAutoScalingGroup(ctx, autoScalingGroupId).Execute()

Delete Auto Scaling Group



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
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.DeleteAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeleteAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoScalingGroupRequest struct via the builder pattern


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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentials

> DeleteCredentials(ctx, instanceId).Execute()

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
	r, err := apiClient.PublicCloudAPI.DeleteCredentials(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeleteCredentials``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCredentialsRequest struct via the builder pattern


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


## DeleteLoadBalancerListener

> DeleteLoadBalancerListener(ctx, loadBalancerId, listenerId).Execute()

Delete load balancer listener



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.DeleteLoadBalancerListener(context.Background(), loadBalancerId, listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeleteLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**listenerId** | **string** | Listener ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerListenerRequest struct via the builder pattern


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


## DeleteSnapshot

> DeleteSnapshot(ctx, instanceId, snapshotId).Execute()

Delete instance snapshot

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
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.DeleteSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**snapshotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


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


## DeregisterAutoScalingGroupLoadBalancer

> AutoScalingGroupDetails DeregisterAutoScalingGroupLoadBalancer(ctx, autoScalingGroupId).Execute()

Deregister Load balancer



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
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.DeregisterAutoScalingGroupLoadBalancer(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeregisterAutoScalingGroupLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeregisterAutoScalingGroupLoadBalancer`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.DeregisterAutoScalingGroupLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterAutoScalingGroupLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeregisterLoadBalancerTargets

> DeregisterLoadBalancerTargets(ctx, loadBalancerId).LoadBalancerTargetOpts(loadBalancerTargetOpts).Execute()

Deregister targets



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	loadBalancerTargetOpts := *openapiclient.NewLoadBalancerTargetOpts() // LoadBalancerTargetOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.DeregisterLoadBalancerTargets(context.Background(), loadBalancerId).LoadBalancerTargetOpts(loadBalancerTargetOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.DeregisterLoadBalancerTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterLoadBalancerTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadBalancerTargetOpts** | [**LoadBalancerTargetOpts**](LoadBalancerTargetOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoScalingGroup

> AutoScalingGroupDetails GetAutoScalingGroup(ctx, autoScalingGroupId).Execute()

Get Auto Scaling Group details



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
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoScalingGroupInstanceList

> GetAutoScalingGroupInstanceListResult GetAutoScalingGroupInstanceList(ctx, autoScalingGroupId).Limit(limit).Offset(offset).Execute()

Get list of instances belonging to an Auto Scaling Group



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
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetAutoScalingGroupInstanceList(context.Background(), autoScalingGroupId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetAutoScalingGroupInstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroupInstanceList`: GetAutoScalingGroupInstanceListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetAutoScalingGroupInstanceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoScalingGroupInstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetAutoScalingGroupInstanceListResult**](GetAutoScalingGroupInstanceListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoScalingGroupList

> GetAutoScalingGroupListResult GetAutoScalingGroupList(ctx).Limit(limit).Offset(offset).Id(id).InstanceId(instanceId).Type_(type_).Region(region).Reference(reference).State(state).Execute()

Get Auto Scaling Group list



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
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	instanceId := "6762182e-7ae9-4d0b-b3b7-bea5a49b3f94" // string |  (optional)
	type_ := "type__example" // string | The Auto Scaling Group's type (optional)
	region := "eu-west-3" // string | The region in which the Auto Scaling Group was created (optional)
	reference := "reference_example" // string | The reference used to identify identifies the Auto Scaling Group (optional)
	state := "state_example" // string | The Auto Scaling Group's current state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetAutoScalingGroupList(context.Background()).Limit(limit).Offset(offset).Id(id).InstanceId(instanceId).Type_(type_).Region(region).Reference(reference).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetAutoScalingGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroupList`: GetAutoScalingGroupListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetAutoScalingGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoScalingGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **id** | **string** |  | 
 **instanceId** | **string** |  | 
 **type_** | **string** | The Auto Scaling Group&#39;s type | 
 **region** | **string** | The region in which the Auto Scaling Group was created | 
 **reference** | **string** | The reference used to identify identifies the Auto Scaling Group | 
 **state** | **string** | The Auto Scaling Group&#39;s current state | 

### Return type

[**GetAutoScalingGroupListResult**](GetAutoScalingGroupListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCpuMetrics

> GetCpuMetricsResult GetCpuMetrics(ctx, instanceId).From(from).To(to).Granularity(granularity).Execute()

Get instance CPU metrics

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
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := "granularity_example" // string | How the metrics are grouped by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetCpuMetrics(context.Background(), instanceId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetCpuMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCpuMetrics`: GetCpuMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetCpuMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCpuMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | **string** | How the metrics are grouped by | 

### Return type

[**GetCpuMetricsResult**](GetCpuMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTrafficMetrics

> GetDataTrafficMetricsResult GetDataTrafficMetrics(ctx, instanceId).From(from).To(to).Granularity(granularity).Aggregation(aggregation).Execute()

Get instance data traffic metrics

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
	instanceId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Instance's ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := "granularity_example" // string | How the metrics are grouped by (optional)
	aggregation := "aggregation_example" // string | How the metrics are aggregated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetDataTrafficMetrics(context.Background(), instanceId).From(from).To(to).Granularity(granularity).Aggregation(aggregation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetDataTrafficMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrafficMetrics`: GetDataTrafficMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetDataTrafficMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTrafficMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | **string** | How the metrics are grouped by | 
 **aggregation** | **string** | How the metrics are aggregated | 

### Return type

[**GetDataTrafficMetricsResult**](GetDataTrafficMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageList

> GetImageListResult GetImageList(ctx).Limit(limit).Offset(offset).Execute()

List all available Images

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
	resp, r, err := apiClient.PublicCloudAPI.GetImageList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageList`: GetImageListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetImageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetImageListResult**](GetImageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> InstanceDetails GetInstance(ctx, instanceId).Execute()

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
	// response from `GetInstance`: InstanceDetails
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

[**InstanceDetails**](InstanceDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceList

> GetInstanceListResult GetInstanceList(ctx).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractType(contractType).ContractState(contractState).ImageId(imageId).State(state).Region(region).Type_(type_).Execute()

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
	contractType := openapiclient.contractType("HOURLY") // ContractType |  (optional)
	contractState := openapiclient.contractState("ACTIVE") // ContractState |  (optional)
	imageId := openapiclient.imageId("ALMALINUX_8_64BIT") // ImageId | Available Images can be obtained using `/v1/images`. (optional)
	state := openapiclient.state("CREATING") // State | The instance's current state(s), separated by commas. (optional)
	region := "eu-west-3" // string | Available regions can be obtained using `/v1/regions` (optional)
	type_ := *openapiclient.NewInstanceType("Name_example", *openapiclient.NewResources(*openapiclient.NewCpu(int32(2), "vCPU"), *openapiclient.NewMemory(float32(3.75), "GiB"), *openapiclient.NewNetworkSpeed(int32(10), "Gbps"), *openapiclient.NewNetworkSpeed(int32(10), "Gbps")), []openapiclient.RootDiskStorageType{openapiclient.rootDiskStorageType("LOCAL")}, *openapiclient.NewPrices("USD", "$", *openapiclient.NewPrice("0.00004", "0.03000"), *openapiclient.NewStorage(*openapiclient.NewPrice("0.00004", "0.03000"), ))) // InstanceType | Available instance types for your region can be obtained using `/v1/instanceTypes`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetInstanceList(context.Background()).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractType(contractType).ContractState(contractState).ImageId(imageId).State(state).Region(region).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetInstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceList`: GetInstanceListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetInstanceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **ip** | **string** |  | 
 **reference** | **string** |  | 
 **id** | **string** |  | 
 **contractType** | [**ContractType**](ContractType.md) |  | 
 **contractState** | [**ContractState**](ContractState.md) |  | 
 **imageId** | [**ImageId**](ImageId.md) | Available Images can be obtained using &#x60;/v1/images&#x60;. | 
 **state** | [**State**](State.md) | The instance&#39;s current state(s), separated by commas. | 
 **region** | **string** | Available regions can be obtained using &#x60;/v1/regions&#x60; | 
 **type_** | [**InstanceType**](InstanceType.md) | Available instance types for your region can be obtained using &#x60;/v1/instanceTypes&#x60;. | 

### Return type

[**GetInstanceListResult**](GetInstanceListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeList

> InstanceTypes GetInstanceTypeList(ctx).Region(region).Limit(limit).Offset(offset).Execute()

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
	region := "eu-west-3" // string | 
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetInstanceTypeList(context.Background()).Region(region).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetInstanceTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeList`: InstanceTypes
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

[**InstanceTypes**](InstanceTypes.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIp

> IpDetails GetIp(ctx, instanceId, ip).Execute()

Get details about an instance's IP

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
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetIp(context.Background(), instanceId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIp`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpList

> GetIpListResult GetIpList(ctx, instanceId).Version(version).NullRouted(nullRouted).Ips(ips).Execute()

List instance's IPs



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
	version := int32(4) // int32 |  (optional)
	nullRouted := true // bool |  (optional)
	ips := "ips_example" // string | A list of IPs separated by `|` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetIpList(context.Background(), instanceId).Version(version).NullRouted(nullRouted).Ips(ips).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpList`: GetIpListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetIpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 
 **nullRouted** | **bool** |  | 
 **ips** | **string** | A list of IPs separated by &#x60;|&#x60; | 

### Return type

[**GetIpListResult**](GetIpListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancer

> LoadBalancerDetails GetLoadBalancer(ctx, loadBalancerId).Execute()

Get load balancer details



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoadBalancerDetails**](LoadBalancerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerList

> GetLoadBalancerListResult GetLoadBalancerList(ctx).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractState(contractState).ContractType(contractType).State(state).Region(region).Type_(type_).Execute()

Get load balancer list



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
	reference := "my-lb" // string |  (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	contractState := openapiclient.contractState("ACTIVE") // ContractState |  (optional)
	contractType := "HOURLY" // string |  (optional)
	state := "RUNNING" // string |  (optional)
	region := "eu-west-3" // string | Available regions can be found using the List Regions endpoint. (optional)
	type_ := "lsw.c3.xlarge" // string | Available load balancer types can be found using the List Load Balancer Types endpoint. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetLoadBalancerList(context.Background()).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractState(contractState).ContractType(contractType).State(state).Region(region).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetLoadBalancerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerList`: GetLoadBalancerListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetLoadBalancerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **ip** | **string** |  | 
 **reference** | **string** |  | 
 **id** | **string** |  | 
 **contractState** | [**ContractState**](ContractState.md) |  | 
 **contractType** | **string** |  | 
 **state** | **string** |  | 
 **region** | **string** | Available regions can be found using the List Regions endpoint. | 
 **type_** | **string** | Available load balancer types can be found using the List Load Balancer Types endpoint. | 

### Return type

[**GetLoadBalancerListResult**](GetLoadBalancerListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerListener

> LoadBalancerListener GetLoadBalancerListener(ctx, loadBalancerId, listenerId).Execute()

Get listener details



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetLoadBalancerListener(context.Background(), loadBalancerId, listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerListener`: LoadBalancerListener
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetLoadBalancerListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**listenerId** | **string** | Listener ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoadBalancerListener**](LoadBalancerListener.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerTargetList

> GetLoadBalancerTargetListResult GetLoadBalancerTargetList(ctx, loadBalancerId).Execute()

List registered targets



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetLoadBalancerTargetList(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetLoadBalancerTargetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerTargetList`: GetLoadBalancerTargetListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetLoadBalancerTargetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerTargetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLoadBalancerTargetListResult**](GetLoadBalancerTargetListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionList

> GetRegionListResult GetRegionList(ctx).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PublicCloudAPI.GetRegionList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetRegionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegionList`: GetRegionListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetRegionListResult**](GetRegionListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReinstallOsList

> []ImageDetails GetReinstallOsList(ctx, instanceId).Execute()

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
	// response from `GetReinstallOsList`: []ImageDetails
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

[**[]ImageDetails**](ImageDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> Snapshot GetSnapshot(ctx, instanceId, snapshotId).Execute()

Get snapshot detail

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
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: Snapshot
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**snapshotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshotList

> GetSnapshotListResult GetSnapshotList(ctx, instanceId).Limit(limit).Offset(offset).Execute()

List snapshots



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
	resp, r, err := apiClient.PublicCloudAPI.GetSnapshotList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetSnapshotList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotList`: GetSnapshotListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetSnapshotList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetSnapshotListResult**](GetSnapshotListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateInstanceTypeList

> InstanceTypes GetUpdateInstanceTypeList(ctx, instanceId).Limit(limit).Offset(offset).Execute()

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
	// response from `GetUpdateInstanceTypeList`: InstanceTypes
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

[**InstanceTypes**](InstanceTypes.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchInstance

> Instance LaunchInstance(ctx).LaunchInstanceOpts(launchInstanceOpts).Execute()

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
	launchInstanceOpts := *openapiclient.NewLaunchInstanceOpts("eu-west-3", openapiclient.typeName("lsw.m3.large"), openapiclient.imageId("ALMALINUX_8_64BIT"), openapiclient.contractType("HOURLY"), openapiclient.contractTerm(0), openapiclient.billingFrequency(1), openapiclient.rootDiskStorageType("LOCAL")) // LaunchInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.LaunchInstance(context.Background()).LaunchInstanceOpts(launchInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.LaunchInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchInstance`: Instance
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

[**Instance**](Instance.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchLoadBalancer

> LoadBalancerDetails LaunchLoadBalancer(ctx).LaunchLoadBalancerOpts(launchLoadBalancerOpts).Execute()

Launch Load balancer



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
	launchLoadBalancerOpts := *openapiclient.NewLaunchLoadBalancerOpts("eu-west-3", openapiclient.typeName("lsw.m3.large"), "ContractType_example", int32(123), openapiclient.rootDiskStorageType("LOCAL"), int32(123)) // LaunchLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.LaunchLoadBalancer(context.Background()).LaunchLoadBalancerOpts(launchLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.LaunchLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.LaunchLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLaunchLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **launchLoadBalancerOpts** | [**LaunchLoadBalancerOpts**](LaunchLoadBalancerOpts.md) |  | 

### Return type

[**LoadBalancerDetails**](LoadBalancerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NullRouteIp

> IpDetails NullRouteIp(ctx, instanceId, ip).NullRouteIpOpts(nullRouteIpOpts).Execute()

Null route IP



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
	ip := "10.0.0.1" // string | 
	nullRouteIpOpts := *openapiclient.NewNullRouteIpOpts() // NullRouteIpOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.NullRouteIp(context.Background(), instanceId, ip).NullRouteIpOpts(nullRouteIpOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.NullRouteIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullRouteIp`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.NullRouteIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNullRouteIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nullRouteIpOpts** | [**NullRouteIpOpts**](NullRouteIpOpts.md) |  | 

### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAutoScalingGroupLoadBalancer

> AutoScalingGroupDetails RegisterAutoScalingGroupLoadBalancer(ctx, autoScalingGroupId).RegisterAutoScalingGroupLoadBalancerOpts(registerAutoScalingGroupLoadBalancerOpts).Execute()

Register Load balancer



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
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	registerAutoScalingGroupLoadBalancerOpts := *openapiclient.NewRegisterAutoScalingGroupLoadBalancerOpts("32082a93-d1e2-4bc0-8f5e-8fe4312b0844") // RegisterAutoScalingGroupLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.RegisterAutoScalingGroupLoadBalancer(context.Background(), autoScalingGroupId).RegisterAutoScalingGroupLoadBalancerOpts(registerAutoScalingGroupLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.RegisterAutoScalingGroupLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAutoScalingGroupLoadBalancer`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.RegisterAutoScalingGroupLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAutoScalingGroupLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerAutoScalingGroupLoadBalancerOpts** | [**RegisterAutoScalingGroupLoadBalancerOpts**](RegisterAutoScalingGroupLoadBalancerOpts.md) |  | 

### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterLoadBalancerTargets

> RegisterLoadBalancerTargets(ctx, loadBalancerId).LoadBalancerTargetOpts(loadBalancerTargetOpts).Execute()

Register targets



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	loadBalancerTargetOpts := *openapiclient.NewLoadBalancerTargetOpts() // LoadBalancerTargetOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.RegisterLoadBalancerTargets(context.Background(), loadBalancerId).LoadBalancerTargetOpts(loadBalancerTargetOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.RegisterLoadBalancerTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterLoadBalancerTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadBalancerTargetOpts** | [**LoadBalancerTargetOpts**](LoadBalancerTargetOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReinstallInstance

> ReinstallInstance(ctx, instanceId).ReinstallInstanceOpts(reinstallInstanceOpts).Execute()

Reinstall instance



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
	reinstallInstanceOpts := *openapiclient.NewReinstallInstanceOpts("ImageId_example") // ReinstallInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.ReinstallInstance(context.Background(), instanceId).ReinstallInstanceOpts(reinstallInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.ReinstallInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReinstallInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reinstallInstanceOpts** | [**ReinstallInstanceOpts**](ReinstallInstanceOpts.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFromPrivateNetwork

> RemoveFromPrivateNetwork(ctx, instanceId).Execute()

Remove instance from Private Network



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
	r, err := apiClient.PublicCloudAPI.RemoveFromPrivateNetwork(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.RemoveFromPrivateNetwork``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveFromPrivateNetworkRequest struct via the builder pattern


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


## RemoveIpNullRoute

> IpDetails RemoveIpNullRoute(ctx, instanceId, ip).Execute()

Remove an IP null route



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
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.RemoveIpNullRoute(context.Background(), instanceId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.RemoveIpNullRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveIpNullRoute`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.RemoveIpNullRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIpNullRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreSnapshot

> RestoreSnapshot(ctx, instanceId, snapshotId).Execute()

Restore instance snapshot

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
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.RestoreSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.RestoreSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**snapshotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSnapshotRequest struct via the builder pattern


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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateLoadBalancer

> TerminateLoadBalancer(ctx, loadBalancerId).Execute()

Delete load balancer



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicCloudAPI.TerminateLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.TerminateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateLoadBalancerRequest struct via the builder pattern


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


## UpdateAutoScalingGroup

> AutoScalingGroupDetails UpdateAutoScalingGroup(ctx, autoScalingGroupId).UpdateAutoScalingGroupOpts(updateAutoScalingGroupOpts).Execute()

Update Auto Scaling Group



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
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	updateAutoScalingGroupOpts := *openapiclient.NewUpdateAutoScalingGroupOpts() // UpdateAutoScalingGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.UpdateAutoScalingGroup(context.Background(), autoScalingGroupId).UpdateAutoScalingGroupOpts(updateAutoScalingGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.UpdateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.UpdateAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAutoScalingGroupOpts** | [**UpdateAutoScalingGroupOpts**](UpdateAutoScalingGroupOpts.md) |  | 

### Return type

[**AutoScalingGroupDetails**](AutoScalingGroupDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstance

> InstanceDetails UpdateInstance(ctx, instanceId).UpdateInstanceOpts(updateInstanceOpts).Execute()

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
	// response from `UpdateInstance`: InstanceDetails
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

[**InstanceDetails**](InstanceDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIp

> IpDetails UpdateIp(ctx, instanceId, ip).UpdateIpOpts(updateIpOpts).Execute()

Update IP



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
	ip := "10.0.0.1" // string | 
	updateIpOpts := *openapiclient.NewUpdateIpOpts("ReverseLookup_example") // UpdateIpOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.UpdateIp(context.Background(), instanceId, ip).UpdateIpOpts(updateIpOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.UpdateIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIp`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.UpdateIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIpOpts** | [**UpdateIpOpts**](UpdateIpOpts.md) |  | 

### Return type

[**IpDetails**](IpDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancer

> LoadBalancerDetails UpdateLoadBalancer(ctx, loadBalancerId).UpdateLoadBalancerOpts(updateLoadBalancerOpts).Execute()

Update load balancer



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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	updateLoadBalancerOpts := *openapiclient.NewUpdateLoadBalancerOpts() // UpdateLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.UpdateLoadBalancer(context.Background(), loadBalancerId).UpdateLoadBalancerOpts(updateLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.UpdateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.UpdateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLoadBalancerOpts** | [**UpdateLoadBalancerOpts**](UpdateLoadBalancerOpts.md) |  | 

### Return type

[**LoadBalancerDetails**](LoadBalancerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancerListener

> LoadBalancerListener UpdateLoadBalancerListener(ctx, loadBalancerId, listenerId).LoadBalancerListenerOpts(loadBalancerListenerOpts).Execute()

Update a listener

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
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID
	loadBalancerListenerOpts := *openapiclient.NewLoadBalancerListenerOpts("Protocol_example", int32(123)) // LoadBalancerListenerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.UpdateLoadBalancerListener(context.Background(), loadBalancerId, listenerId).LoadBalancerListenerOpts(loadBalancerListenerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.UpdateLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancerListener`: LoadBalancerListener
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.UpdateLoadBalancerListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 
**listenerId** | **string** | Listener ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loadBalancerListenerOpts** | [**LoadBalancerListenerOpts**](LoadBalancerListenerOpts.md) |  | 

### Return type

[**LoadBalancerListener**](LoadBalancerListener.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

