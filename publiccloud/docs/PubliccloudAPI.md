# \PubliccloudAPI

All URIs are relative to *https://api.leaseweb.com/publicCloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToPrivateNetwork**](PubliccloudAPI.md#AddToPrivateNetwork) | **Put** /instances/{instanceId}/addToPrivateNetwork | Add instance to Private Network
[**AttachIso**](PubliccloudAPI.md#AttachIso) | **Post** /instances/{instanceId}/attachIso | Attach ISO to instance
[**CancelInstanceTermination**](PubliccloudAPI.md#CancelInstanceTermination) | **Post** /instances/{instanceId}/cancelTermination | Cancel instance termination
[**CreateAutoScalingGroup**](PubliccloudAPI.md#CreateAutoScalingGroup) | **Post** /autoScalingGroups | Create Auto Scaling Group
[**CreateImage**](PubliccloudAPI.md#CreateImage) | **Post** /images | Create Custom Image
[**CreateLoadBalancerListener**](PubliccloudAPI.md#CreateLoadBalancerListener) | **Post** /loadBalancers/{loadBalancerId}/listeners | Create listener
[**CreateSnapshot**](PubliccloudAPI.md#CreateSnapshot) | **Post** /instances/{instanceId}/snapshots | Create instance snapshot
[**CreateTargetGroup**](PubliccloudAPI.md#CreateTargetGroup) | **Post** /targetGroups | Create Target Group
[**DeleteAutoScalingGroup**](PubliccloudAPI.md#DeleteAutoScalingGroup) | **Delete** /autoScalingGroups/{autoScalingGroupId} | Delete Auto Scaling Group
[**DeleteCredential**](PubliccloudAPI.md#DeleteCredential) | **Delete** /instances/{instanceId}/credentials/{type}/{username} | Delete credentials
[**DeleteCredentials**](PubliccloudAPI.md#DeleteCredentials) | **Delete** /instances/{instanceId}/credentials | Delete all instance credentials
[**DeleteLoadBalancerListener**](PubliccloudAPI.md#DeleteLoadBalancerListener) | **Delete** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Delete load balancer listener
[**DeleteSnapshot**](PubliccloudAPI.md#DeleteSnapshot) | **Delete** /instances/{instanceId}/snapshots/{snapshotId} | Delete instance snapshot
[**DeleteTargetGroup**](PubliccloudAPI.md#DeleteTargetGroup) | **Delete** /targetGroups/{targetGroupId} | Delete Target Group
[**DeregisterAutoScalingGroupTargetGroup**](PubliccloudAPI.md#DeregisterAutoScalingGroupTargetGroup) | **Post** /autoScalingGroups/{autoScalingGroupId}/deregisterTargetGroup | Deregister Target Group
[**DeregisterTargets**](PubliccloudAPI.md#DeregisterTargets) | **Post** /targetGroups/{targetGroupId}/deregisterTargets | Deregister Targets
[**DetachIso**](PubliccloudAPI.md#DetachIso) | **Post** /instances/{instanceId}/detachIso | Detach ISO from instance
[**GetAutoScalingGroup**](PubliccloudAPI.md#GetAutoScalingGroup) | **Get** /autoScalingGroups/{autoScalingGroupId} | Get Auto Scaling Group details
[**GetAutoScalingGroupInstanceList**](PubliccloudAPI.md#GetAutoScalingGroupInstanceList) | **Get** /autoScalingGroups/{autoScalingGroupId}/instances | Get list of instances belonging to an Auto Scaling Group
[**GetAutoScalingGroupList**](PubliccloudAPI.md#GetAutoScalingGroupList) | **Get** /autoScalingGroups | Get Auto Scaling Group list
[**GetConnectionsMetrics**](PubliccloudAPI.md#GetConnectionsMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/connections | Get connections metrics
[**GetConnectionsPerSecondMetrics**](PubliccloudAPI.md#GetConnectionsPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/connectionsPerSecond | Get connections per second metrics
[**GetConsoleAccessToInstance**](PubliccloudAPI.md#GetConsoleAccessToInstance) | **Get** /instances/{instanceId}/console | Get console access
[**GetCpuMetrics**](PubliccloudAPI.md#GetCpuMetrics) | **Get** /instances/{instanceId}/metrics/cpu | Get instance CPU metrics
[**GetCredential**](PubliccloudAPI.md#GetCredential) | **Get** /instances/{instanceId}/credentials/{type}/{username} | Get credentials by type and username
[**GetCredentialList**](PubliccloudAPI.md#GetCredentialList) | **Get** /instances/{instanceId}/credentials | List credentials stored for instance
[**GetCredentialListByType**](PubliccloudAPI.md#GetCredentialListByType) | **Get** /instances/{instanceId}/credentials/{type} | Get credentials by type
[**GetDataTrafficMetrics**](PubliccloudAPI.md#GetDataTrafficMetrics) | **Get** /instances/{instanceId}/metrics/datatraffic | Get instance data traffic metrics
[**GetDataTransferredMetrics**](PubliccloudAPI.md#GetDataTransferredMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/dataTransferred | Get load balancer data transferred metrics
[**GetDataTransferredPerSecondMetrics**](PubliccloudAPI.md#GetDataTransferredPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/dataTransferredPerSecond | Get load balancer data transferred per second metrics
[**GetExpenses**](PubliccloudAPI.md#GetExpenses) | **Get** /equipments/{equipmentId}/expenses | Get costs for a given month.
[**GetImageList**](PubliccloudAPI.md#GetImageList) | **Get** /images | List all available Images
[**GetInstance**](PubliccloudAPI.md#GetInstance) | **Get** /instances/{instanceId} | Get instance details
[**GetInstanceList**](PubliccloudAPI.md#GetInstanceList) | **Get** /instances | Get instance list
[**GetInstanceTypeList**](PubliccloudAPI.md#GetInstanceTypeList) | **Get** /instanceTypes | List instance types
[**GetIp**](PubliccloudAPI.md#GetIp) | **Get** /instances/{instanceId}/ips/{ip} | Get details about an instance&#39;s IP
[**GetIpList**](PubliccloudAPI.md#GetIpList) | **Get** /instances/{instanceId}/ips | List instance&#39;s IPs
[**GetIsoList**](PubliccloudAPI.md#GetIsoList) | **Get** /isos | List available ISOs
[**GetLoadBalancer**](PubliccloudAPI.md#GetLoadBalancer) | **Get** /loadBalancers/{loadBalancerId} | Get load balancer details
[**GetLoadBalancerList**](PubliccloudAPI.md#GetLoadBalancerList) | **Get** /loadBalancers | Get load balancer list
[**GetLoadBalancerListener**](PubliccloudAPI.md#GetLoadBalancerListener) | **Get** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Get listener details
[**GetLoadBalancerListenerList**](PubliccloudAPI.md#GetLoadBalancerListenerList) | **Get** /loadBalancers/{loadBalancerId}/listeners | Get listener list
[**GetMarketAppList**](PubliccloudAPI.md#GetMarketAppList) | **Get** /marketApps | Get marketplace apps
[**GetRegionList**](PubliccloudAPI.md#GetRegionList) | **Get** /regions | List regions
[**GetReinstallImageList**](PubliccloudAPI.md#GetReinstallImageList) | **Get** /instances/{instanceId}/reinstall/images | List images available for reinstall
[**GetRequestsMetrics**](PubliccloudAPI.md#GetRequestsMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/requests | Get load balancer requests metrics
[**GetRequestsPerSecondMetrics**](PubliccloudAPI.md#GetRequestsPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/requestsPerSecond | Get load balancer requests per second metrics
[**GetResponseCodesMetrics**](PubliccloudAPI.md#GetResponseCodesMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/responseCodes | Get response codes metrics
[**GetResponseCodesPerSecondMetrics**](PubliccloudAPI.md#GetResponseCodesPerSecondMetrics) | **Get** /loadBalancers/{loadBalancerId}/metrics/responseCodesPerSecond | Get response codes per second metrics
[**GetSnapshot**](PubliccloudAPI.md#GetSnapshot) | **Get** /instances/{instanceId}/snapshots/{snapshotId} | Get snapshot detail
[**GetSnapshotList**](PubliccloudAPI.md#GetSnapshotList) | **Get** /instances/{instanceId}/snapshots | List snapshots
[**GetTargetGroup**](PubliccloudAPI.md#GetTargetGroup) | **Get** /targetGroups/{targetGroupId} | Get Target Group details
[**GetTargetGroupList**](PubliccloudAPI.md#GetTargetGroupList) | **Get** /targetGroups | Get Target Group list
[**GetTargetList**](PubliccloudAPI.md#GetTargetList) | **Get** /targetGroups/{targetGroupId}/targets | Get Targets
[**GetUpdateInstanceTypeList**](PubliccloudAPI.md#GetUpdateInstanceTypeList) | **Get** /instances/{instanceId}/instanceTypesUpdate | List available instance types for update
[**LaunchInstance**](PubliccloudAPI.md#LaunchInstance) | **Post** /instances | Launch instance
[**LaunchLoadBalancer**](PubliccloudAPI.md#LaunchLoadBalancer) | **Post** /loadBalancers | Launch Load balancer
[**NullRouteIp**](PubliccloudAPI.md#NullRouteIp) | **Post** /instances/{instanceId}/ips/{ip}/null | Null route IP
[**RebootInstance**](PubliccloudAPI.md#RebootInstance) | **Post** /instances/{instanceId}/reboot | Reboot instance
[**RegisterAutoScalingGroupTargetGroup**](PubliccloudAPI.md#RegisterAutoScalingGroupTargetGroup) | **Post** /autoScalingGroups/{autoScalingGroupId}/registerTargetGroup | Register Target Group
[**RegisterTargets**](PubliccloudAPI.md#RegisterTargets) | **Post** /targetGroups/{targetGroupId}/registerTargets | Register Targets
[**ReinstallInstance**](PubliccloudAPI.md#ReinstallInstance) | **Put** /instances/{instanceId}/reinstall | Reinstall instance
[**RemoveFromPrivateNetwork**](PubliccloudAPI.md#RemoveFromPrivateNetwork) | **Delete** /instances/{instanceId}/removeFromPrivateNetwork | Remove instance from Private Network
[**RemoveIpNullRoute**](PubliccloudAPI.md#RemoveIpNullRoute) | **Post** /instances/{instanceId}/ips/{ip}/unnull | Remove an IP null route
[**ResetPassword**](PubliccloudAPI.md#ResetPassword) | **Post** /instances/{instanceId}/resetPassword | Reset instance password
[**RestoreSnapshot**](PubliccloudAPI.md#RestoreSnapshot) | **Put** /instances/{instanceId}/snapshots/{snapshotId} | Restore instance snapshot
[**StartInstance**](PubliccloudAPI.md#StartInstance) | **Post** /instances/{instanceId}/start | Start instance
[**StopInstance**](PubliccloudAPI.md#StopInstance) | **Post** /instances/{instanceId}/stop | Stop instance
[**StoreCredential**](PubliccloudAPI.md#StoreCredential) | **Post** /instances/{instanceId}/credentials | Store credentials
[**TerminateInstance**](PubliccloudAPI.md#TerminateInstance) | **Delete** /instances/{instanceId} | Terminate instance
[**TerminateLoadBalancer**](PubliccloudAPI.md#TerminateLoadBalancer) | **Delete** /loadBalancers/{loadBalancerId} | Delete load balancer
[**UpdateAutoScalingGroup**](PubliccloudAPI.md#UpdateAutoScalingGroup) | **Put** /autoScalingGroups/{autoScalingGroupId} | Update Auto Scaling Group
[**UpdateCredential**](PubliccloudAPI.md#UpdateCredential) | **Put** /instances/{instanceId}/credentials/{type}/{username} | Update credentials
[**UpdateImage**](PubliccloudAPI.md#UpdateImage) | **Put** /images/{imageId} | Update Custom Image
[**UpdateInstance**](PubliccloudAPI.md#UpdateInstance) | **Put** /instances/{instanceId} | Update instance
[**UpdateIp**](PubliccloudAPI.md#UpdateIp) | **Put** /instances/{instanceId}/ips/{ip} | Update IP
[**UpdateLoadBalancer**](PubliccloudAPI.md#UpdateLoadBalancer) | **Put** /loadBalancers/{loadBalancerId} | Update load balancer
[**UpdateLoadBalancerListener**](PubliccloudAPI.md#UpdateLoadBalancerListener) | **Put** /loadBalancers/{loadBalancerId}/listeners/{listenerId} | Update a listener
[**UpdateTargetGroup**](PubliccloudAPI.md#UpdateTargetGroup) | **Put** /targetGroups/{targetGroupId} | Update Target Group



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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.AddToPrivateNetwork(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.AddToPrivateNetwork``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	attachIsoOpts := *openapiclient.NewAttachIsoOpts("IsoId_example") // AttachIsoOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.AttachIso(context.Background(), instanceId).AttachIsoOpts(attachIsoOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.AttachIso``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.CancelInstanceTermination(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CancelInstanceTermination``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	createAutoScalingGroupOpts := *openapiclient.NewCreateAutoScalingGroupOpts("InstanceId_example", "Reference_example", "Type_example") // CreateAutoScalingGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateAutoScalingGroup(context.Background()).CreateAutoScalingGroupOpts(createAutoScalingGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateAutoScalingGroup`: %v\n", resp)
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


## CreateImage

> ImageDetails CreateImage(ctx).CreateImageOpts(createImageOpts).Execute()

Create Custom Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	createImageOpts := *openapiclient.NewCreateImageOpts("Name_example", "InstanceId_example") // CreateImageOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateImage(context.Background()).CreateImageOpts(createImageOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImage`: ImageDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageOpts** | [**CreateImageOpts**](CreateImageOpts.md) |  | 

### Return type

[**ImageDetails**](ImageDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancerListener

> LoadBalancerListener CreateLoadBalancerListener(ctx, loadBalancerId).LoadBalancerListenerCreateOpts(loadBalancerListenerCreateOpts).Execute()

Create listener



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	loadBalancerListenerCreateOpts := *openapiclient.NewLoadBalancerListenerCreateOpts(openapiclient.protocol("HTTP"), int32(123), *openapiclient.NewLoadBalancerListenerDefaultRule("TargetGroupId_example")) // LoadBalancerListenerCreateOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateLoadBalancerListener(context.Background(), loadBalancerId).LoadBalancerListenerCreateOpts(loadBalancerListenerCreateOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoadBalancerListener`: LoadBalancerListener
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateLoadBalancerListener`: %v\n", resp)
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

 **loadBalancerListenerCreateOpts** | [**LoadBalancerListenerCreateOpts**](LoadBalancerListenerCreateOpts.md) |  | 

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.CreateSnapshot(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateSnapshot``: %v\n", err)
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


## CreateTargetGroup

> TargetGroup CreateTargetGroup(ctx).CreateTargetGroupOpts(createTargetGroupOpts).Execute()

Create Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	createTargetGroupOpts := *openapiclient.NewCreateTargetGroupOpts("Name_example", openapiclient.protocol("HTTP"), int32(123), openapiclient.regionName("eu-west-3")) // CreateTargetGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.CreateTargetGroup(context.Background()).CreateTargetGroupOpts(createTargetGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.CreateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTargetGroup`: TargetGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.CreateTargetGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTargetGroupOpts** | [**CreateTargetGroupOpts**](CreateTargetGroupOpts.md) |  | 

### Return type

[**TargetGroup**](TargetGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteAutoScalingGroup``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteCredential(context.Background(), instanceId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteCredential``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteCredentials(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteCredentials``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteLoadBalancerListener(context.Background(), loadBalancerId, listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteLoadBalancerListener``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteSnapshot``: %v\n", err)
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


## DeleteTargetGroup

> DeleteTargetGroup(ctx, targetGroupId).Execute()

Delete Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeleteTargetGroup(context.Background(), targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeleteTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetGroupRequest struct via the builder pattern


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


## DeregisterAutoScalingGroupTargetGroup

> AutoScalingGroupDetails DeregisterAutoScalingGroupTargetGroup(ctx, autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()

Deregister Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	targetGroupIdOpts := *openapiclient.NewTargetGroupIdOpts("TargetGroupId_example") // TargetGroupIdOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.DeregisterAutoScalingGroupTargetGroup(context.Background(), autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeregisterAutoScalingGroupTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeregisterAutoScalingGroupTargetGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.DeregisterAutoScalingGroupTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterAutoScalingGroupTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetGroupIdOpts** | [**TargetGroupIdOpts**](TargetGroupIdOpts.md) |  | 

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


## DeregisterTargets

> DeregisterTargets(ctx, targetGroupId).RequestBody(requestBody).Execute()

Deregister Targets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DeregisterTargets(context.Background(), targetGroupId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DeregisterTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.DetachIso(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.DetachIso``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetAutoScalingGroup`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetAutoScalingGroupInstanceList(context.Background(), autoScalingGroupId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetAutoScalingGroupInstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroupInstanceList`: GetAutoScalingGroupInstanceListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetAutoScalingGroupInstanceList`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	instanceId := "6762182e-7ae9-4d0b-b3b7-bea5a49b3f94" // string |  (optional)
	type_ := "type__example" // string | The Auto Scaling Group's type (optional)
	region := openapiclient.regionName("eu-west-3") // RegionName | The region in which the Auto Scaling Group was created (optional)
	reference := "reference_example" // string | The reference used to identify identifies the Auto Scaling Group (optional)
	state := "state_example" // string | The Auto Scaling Group's current state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetAutoScalingGroupList(context.Background()).Limit(limit).Offset(offset).Id(id).InstanceId(instanceId).Type_(type_).Region(region).Reference(reference).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetAutoScalingGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoScalingGroupList`: GetAutoScalingGroupListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetAutoScalingGroupList`: %v\n", resp)
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
 **region** | [**RegionName**](RegionName.md) | The region in which the Auto Scaling Group was created | 
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


## GetConnectionsMetrics

> GetConnectionsMetricsResult GetConnectionsMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get connections metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := "granularity_example" // string | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetConnectionsMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetConnectionsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsMetrics`: GetConnectionsMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetConnectionsMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | **string** | The interval for each metric | 

### Return type

[**GetConnectionsMetricsResult**](GetConnectionsMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsPerSecondMetrics

> GetConnectionsPerSecondMetricsResult GetConnectionsPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get connections per second metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := "granularity_example" // string | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetConnectionsPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetConnectionsPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsPerSecondMetrics`: GetConnectionsPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetConnectionsPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | **string** | The interval for each metric | 

### Return type

[**GetConnectionsPerSecondMetricsResult**](GetConnectionsPerSecondMetricsResult.md)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetConsoleAccessToInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetConsoleAccessToInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsoleAccessToInstance`: GetConsoleAccessToInstanceResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetConsoleAccessToInstance`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := openapiclient.metricsGranularity("5m") // MetricsGranularity | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetCpuMetrics(context.Background(), instanceId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetCpuMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCpuMetrics`: GetCpuMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetCpuMetrics`: %v\n", resp)
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
 **granularity** | [**MetricsGranularity**](MetricsGranularity.md) | The interval for each metric | 

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetCredential(context.Background(), instanceId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: GetCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetCredential`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetCredentialList(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialList`: GetCredentialListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetCredentialList`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetCredentialListByType(context.Background(), instanceId, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialListByType`: GetCredentialListByTypeResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetCredentialListByType`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := "granularity_example" // string | How the metrics are grouped by (optional)
	aggregation := "aggregation_example" // string | How the metrics are aggregated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetDataTrafficMetrics(context.Background(), instanceId).From(from).To(to).Granularity(granularity).Aggregation(aggregation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetDataTrafficMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrafficMetrics`: GetDataTrafficMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetDataTrafficMetrics`: %v\n", resp)
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


## GetDataTransferredMetrics

> GetDataTransferredMetricsResult GetDataTransferredMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer data transferred metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := openapiclient.metricsGranularity("5m") // MetricsGranularity | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetDataTransferredMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetDataTransferredMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTransferredMetrics`: GetDataTransferredMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetDataTransferredMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTransferredMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricsGranularity**](MetricsGranularity.md) | The interval for each metric | 

### Return type

[**GetDataTransferredMetricsResult**](GetDataTransferredMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTransferredPerSecondMetrics

> GetDataTransferredPerSecondMetricsResult GetDataTransferredPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer data transferred per second metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := openapiclient.metricsGranularity("5m") // MetricsGranularity | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetDataTransferredPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetDataTransferredPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTransferredPerSecondMetrics`: GetDataTransferredPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetDataTransferredPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTransferredPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricsGranularity**](MetricsGranularity.md) | The interval for each metric | 

### Return type

[**GetDataTransferredPerSecondMetricsResult**](GetDataTransferredPerSecondMetricsResult.md)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	equipmentId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Equipment's UUID
	from := time.Now() // string | Start date of the period to get costs for. It must be the first day of the month
	to := time.Now() // string | End date of the period to get costs for. This date needs to be exactly one month after the 'from' date. If this value is not passed, it will be calculated based on 'from' parameter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetExpenses(context.Background(), equipmentId).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetExpenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenses`: GetExpensesResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetExpenses`: %v\n", resp)
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

> GetImageListResult GetImageList(ctx).Limit(limit).Offset(offset).Custom(custom).Standard(standard).State(state).MarketAppId(marketAppId).StorageType(storageType).Name(name).Flavour(flavour).Region(region).Execute()

List all available Images

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	custom := true // bool | Filters the list to include only custom images. (optional)
	standard := true // bool | Filters the list to include only standard images. (optional)
	state := openapiclient.imageState("CREATING") // ImageState | Filters the list by the state of custom images. (optional)
	marketAppId := openapiclient.marketAppId("CPANEL_30") // MarketAppId | Filters the list by the market app of standard images. (optional)
	storageType := openapiclient.storageType("LOCAL") // StorageType | Filters the list by the market app of standard images. (optional)
	name := "Ubuntu 20.04 LTS (x86_64)" // string | Filters the list by the name of images. (optional)
	flavour := openapiclient.flavour("ubuntu") // Flavour | Filters the list by the flavour of standard images. (optional)
	region := "eu-west-3" // string | Available regions can be found using the List Regions endpoint. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetImageList(context.Background()).Limit(limit).Offset(offset).Custom(custom).Standard(standard).State(state).MarketAppId(marketAppId).StorageType(storageType).Name(name).Flavour(flavour).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageList`: GetImageListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetImageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **custom** | **bool** | Filters the list to include only custom images. | 
 **standard** | **bool** | Filters the list to include only standard images. | 
 **state** | [**ImageState**](ImageState.md) | Filters the list by the state of custom images. | 
 **marketAppId** | [**MarketAppId**](MarketAppId.md) | Filters the list by the market app of standard images. | 
 **storageType** | [**StorageType**](StorageType.md) | Filters the list by the market app of standard images. | 
 **name** | **string** | Filters the list by the name of images. | 
 **flavour** | [**Flavour**](Flavour.md) | Filters the list by the flavour of standard images. | 
 **region** | **string** | Available regions can be found using the List Regions endpoint. | 

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: InstanceDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstance`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	ip := "10.0.0.1" // string |  (optional)
	reference := "my-webserver" // string |  (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	contractType := openapiclient.contractType("HOURLY") // ContractType |  (optional)
	contractState := openapiclient.contractState("ACTIVE") // ContractState |  (optional)
	imageId := "UBUNTU_22_04_64BIT" // string | Available Images can be obtained using `/v1/images`. (optional)
	state := openapiclient.state("CREATING") // State | The instance's current state(s), separated by commas. (optional)
	region := openapiclient.regionName("eu-west-3") // RegionName | Available regions can be obtained using `/v1/regions` (optional)
	type_ := openapiclient.typeName("lsw.m3.large") // TypeName | Available instance types for your region can be obtained using `/v1/instanceTypes`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceList(context.Background()).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractType(contractType).ContractState(contractState).ImageId(imageId).State(state).Region(region).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceList`: GetInstanceListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceList`: %v\n", resp)
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
 **imageId** | **string** | Available Images can be obtained using &#x60;/v1/images&#x60;. | 
 **state** | [**State**](State.md) | The instance&#39;s current state(s), separated by commas. | 
 **region** | [**RegionName**](RegionName.md) | Available regions can be obtained using &#x60;/v1/regions&#x60; | 
 **type_** | [**TypeName**](TypeName.md) | Available instance types for your region can be obtained using &#x60;/v1/instanceTypes&#x60;. | 

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	region := openapiclient.regionName("eu-west-3") // RegionName | 
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceTypeList(context.Background()).Region(region).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeList`: InstanceTypes
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**RegionName**](RegionName.md) |  | 
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetIp(context.Background(), instanceId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIp`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetIp`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	version := int32(4) // int32 |  (optional)
	nullRouted := true // bool |  (optional)
	ips := "ips_example" // string | A list of IPs separated by `|` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetIpList(context.Background(), instanceId).Version(version).NullRouted(nullRouted).Ips(ips).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpList`: GetIpListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetIpList`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetIsoList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetIsoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsoList`: GetIsoListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetIsoList`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancer`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
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
	region := openapiclient.regionName("eu-west-3") // RegionName | Available regions can be found using the List Regions endpoint. (optional)
	type_ := "lsw.c3.xlarge" // string | Available load balancer types can be found using the List Load Balancer Types endpoint. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerList(context.Background()).Limit(limit).Offset(offset).Ip(ip).Reference(reference).Id(id).ContractState(contractState).ContractType(contractType).State(state).Region(region).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerList`: GetLoadBalancerListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerList`: %v\n", resp)
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
 **region** | [**RegionName**](RegionName.md) | Available regions can be found using the List Regions endpoint. | 
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

> LoadBalancerListenerDetails GetLoadBalancerListener(ctx, loadBalancerId, listenerId).Execute()

Get listener details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerListener(context.Background(), loadBalancerId, listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerListener`: LoadBalancerListenerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerListener`: %v\n", resp)
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

[**LoadBalancerListenerDetails**](LoadBalancerListenerDetails.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerListenerList

> GetLoadBalancerListenerListResult GetLoadBalancerListenerList(ctx, loadBalancerId).Limit(limit).Offset(offset).Execute()

Get listener list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetLoadBalancerListenerList(context.Background(), loadBalancerId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetLoadBalancerListenerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancerListenerList`: GetLoadBalancerListenerListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetLoadBalancerListenerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerListenerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetLoadBalancerListenerListResult**](GetLoadBalancerListenerListResult.md)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetMarketAppList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetMarketAppList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketAppList`: GetMarketAppListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetMarketAppList`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetRegionList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetRegionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegionList`: GetRegionListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetRegionList`: %v\n", resp)
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


## GetReinstallImageList

> GetReinstallImageListResult GetReinstallImageList(ctx, instanceId).Limit(limit).Offset(offset).Custom(custom).Standard(standard).Execute()

List images available for reinstall



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	custom := true // bool | Filters the list to include only custom images. (optional)
	standard := true // bool | Filters the list to include only standard images. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetReinstallImageList(context.Background(), instanceId).Limit(limit).Offset(offset).Custom(custom).Standard(standard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetReinstallImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReinstallImageList`: GetReinstallImageListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetReinstallImageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReinstallImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **custom** | **bool** | Filters the list to include only custom images. | 
 **standard** | **bool** | Filters the list to include only standard images. | 

### Return type

[**GetReinstallImageListResult**](GetReinstallImageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestsMetrics

> GetRequestsMetricsResult GetRequestsMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer requests metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := openapiclient.metricsGranularity("5m") // MetricsGranularity | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetRequestsMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetRequestsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestsMetrics`: GetRequestsMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetRequestsMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricsGranularity**](MetricsGranularity.md) | The interval for each metric | 

### Return type

[**GetRequestsMetricsResult**](GetRequestsMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestsPerSecondMetrics

> GetRequestsPerSecondMetricsResult GetRequestsPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get load balancer requests per second metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := openapiclient.metricsGranularity("5m") // MetricsGranularity | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetRequestsPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetRequestsPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestsPerSecondMetrics`: GetRequestsPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetRequestsPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | [**MetricsGranularity**](MetricsGranularity.md) | The interval for each metric | 

### Return type

[**GetRequestsPerSecondMetricsResult**](GetRequestsPerSecondMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResponseCodesMetrics

> GetResponseCodesMetricsResult GetResponseCodesMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get response codes metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := "granularity_example" // string | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetResponseCodesMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetResponseCodesMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResponseCodesMetrics`: GetResponseCodesMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetResponseCodesMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponseCodesMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | **string** | The interval for each metric | 

### Return type

[**GetResponseCodesMetricsResult**](GetResponseCodesMetricsResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResponseCodesPerSecondMetrics

> GetResponseCodesPerSecondMetricsResult GetResponseCodesPerSecondMetrics(ctx, loadBalancerId).From(from).To(to).Granularity(granularity).Execute()

Get response codes per second metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	from := time.Now() // string | The start of the interval to get the metric (optional)
	to := time.Now() // string | The end of the interval to get the metric. Must be greater than the date provided in `from` (optional)
	granularity := "granularity_example" // string | The interval for each metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetResponseCodesPerSecondMetrics(context.Background(), loadBalancerId).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetResponseCodesPerSecondMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResponseCodesPerSecondMetrics`: GetResponseCodesPerSecondMetricsResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetResponseCodesPerSecondMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponseCodesPerSecondMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the interval to get the metric | 
 **to** | **string** | The end of the interval to get the metric. Must be greater than the date provided in &#x60;from&#x60; | 
 **granularity** | **string** | The interval for each metric | 

### Return type

[**GetResponseCodesPerSecondMetricsResult**](GetResponseCodesPerSecondMetricsResult.md)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: Snapshot
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetSnapshot`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetSnapshotList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetSnapshotList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotList`: GetSnapshotListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetSnapshotList`: %v\n", resp)
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


## GetTargetGroup

> TargetGroup GetTargetGroup(ctx, targetGroupId).Execute()

Get Target Group details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetTargetGroup(context.Background(), targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroup`: TargetGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TargetGroup**](TargetGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetGroupList

> GetTargetGroupListResult GetTargetGroupList(ctx).Limit(limit).Offset(offset).Id(id).Name(name).Protocol(protocol).Port(port).Region(region).Execute()

Get Target Group list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	id := "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" // string |  (optional)
	name := "Foo bar" // string |  (optional)
	protocol := openapiclient.protocol("HTTP") // Protocol |  (optional)
	port := int32(80) // int32 |  (optional)
	region := openapiclient.regionName("eu-west-3") // RegionName |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetTargetGroupList(context.Background()).Limit(limit).Offset(offset).Id(id).Name(name).Protocol(protocol).Port(port).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetTargetGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroupList`: GetTargetGroupListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetTargetGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **id** | **string** |  | 
 **name** | **string** |  | 
 **protocol** | [**Protocol**](Protocol.md) |  | 
 **port** | **int32** |  | 
 **region** | [**RegionName**](RegionName.md) |  | 

### Return type

[**GetTargetGroupListResult**](GetTargetGroupListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetList

> GetTargetListResult GetTargetList(ctx, targetGroupId).Limit(limit).Offset(offset).Execute()

Get Targets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetTargetList(context.Background(), targetGroupId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetTargetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetList`: GetTargetListResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetTargetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetTargetListResult**](GetTargetListResult.md)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetUpdateInstanceTypeList(context.Background(), instanceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetUpdateInstanceTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpdateInstanceTypeList`: InstanceTypes
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetUpdateInstanceTypeList`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	launchInstanceOpts := *openapiclient.NewLaunchInstanceOpts(openapiclient.regionName("eu-west-3"), openapiclient.typeName("lsw.m3.large"), "UBUNTU_22_04_64BIT", openapiclient.contractType("HOURLY"), openapiclient.contractTerm(0), openapiclient.billingFrequency(1), openapiclient.storageType("LOCAL")) // LaunchInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.LaunchInstance(context.Background()).LaunchInstanceOpts(launchInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.LaunchInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.LaunchInstance`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	launchLoadBalancerOpts := *openapiclient.NewLaunchLoadBalancerOpts(openapiclient.regionName("eu-west-3"), openapiclient.typeName("lsw.m3.large"), openapiclient.contractType("HOURLY"), openapiclient.contractTerm(0), openapiclient.billingFrequency(1)) // LaunchLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.LaunchLoadBalancer(context.Background()).LaunchLoadBalancerOpts(launchLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.LaunchLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.LaunchLoadBalancer`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 
	nullRouteIpOpts := *openapiclient.NewNullRouteIpOpts() // NullRouteIpOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.NullRouteIp(context.Background(), instanceId, ip).NullRouteIpOpts(nullRouteIpOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.NullRouteIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullRouteIp`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.NullRouteIp`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RebootInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RebootInstance``: %v\n", err)
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


## RegisterAutoScalingGroupTargetGroup

> AutoScalingGroupDetails RegisterAutoScalingGroupTargetGroup(ctx, autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()

Register Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	targetGroupIdOpts := *openapiclient.NewTargetGroupIdOpts("TargetGroupId_example") // TargetGroupIdOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.RegisterAutoScalingGroupTargetGroup(context.Background(), autoScalingGroupId).TargetGroupIdOpts(targetGroupIdOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RegisterAutoScalingGroupTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAutoScalingGroupTargetGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.RegisterAutoScalingGroupTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAutoScalingGroupTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetGroupIdOpts** | [**TargetGroupIdOpts**](TargetGroupIdOpts.md) |  | 

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


## RegisterTargets

> RegisterTargets(ctx, targetGroupId).RequestBody(requestBody).Execute()

Register Targets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RegisterTargets(context.Background(), targetGroupId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RegisterTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	reinstallInstanceOpts := *openapiclient.NewReinstallInstanceOpts("ImageId_example") // ReinstallInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.ReinstallInstance(context.Background(), instanceId).ReinstallInstanceOpts(reinstallInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.ReinstallInstance``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RemoveFromPrivateNetwork(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RemoveFromPrivateNetwork``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.RemoveIpNullRoute(context.Background(), instanceId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RemoveIpNullRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveIpNullRoute`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.RemoveIpNullRoute`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.ResetPassword(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.ResetPassword``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.RestoreSnapshot(context.Background(), instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.RestoreSnapshot``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.StartInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StartInstance``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.StopInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StopInstance``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	storeCredentialOpts := *openapiclient.NewStoreCredentialOpts(openapiclient.credentialType("OPERATING_SYSTEM"), "Username_example", "Password_example") // StoreCredentialOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.StoreCredential(context.Background(), instanceId).StoreCredentialOpts(storeCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.StoreCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreCredential`: StoreCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.StoreCredential`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.TerminateInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.TerminateInstance``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubliccloudAPI.TerminateLoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.TerminateLoadBalancer``: %v\n", err)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	autoScalingGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Auto Scaling Group ID
	updateAutoScalingGroupOpts := *openapiclient.NewUpdateAutoScalingGroupOpts() // UpdateAutoScalingGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateAutoScalingGroup(context.Background(), autoScalingGroupId).UpdateAutoScalingGroupOpts(updateAutoScalingGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroup`: AutoScalingGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateAutoScalingGroup`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username
	updateCredentialOpts := *openapiclient.NewUpdateCredentialOpts("Password_example") // UpdateCredentialOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateCredential(context.Background(), instanceId, type_, username).UpdateCredentialOpts(updateCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredential`: UpdateCredentialResult
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateCredential`: %v\n", resp)
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


## UpdateImage

> ImageDetails UpdateImage(ctx, imageId).UpdateImageOpts(updateImageOpts).Execute()

Update Custom Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	imageId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Image's ID
	updateImageOpts := *openapiclient.NewUpdateImageOpts("Name_example") // UpdateImageOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateImage(context.Background(), imageId).UpdateImageOpts(updateImageOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImage`: ImageDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateImageOpts** | [**UpdateImageOpts**](UpdateImageOpts.md) |  | 

### Return type

[**ImageDetails**](ImageDetails.md)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	updateInstanceOpts := *openapiclient.NewUpdateInstanceOpts() // UpdateInstanceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateInstance(context.Background(), instanceId).UpdateInstanceOpts(updateInstanceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstance`: InstanceDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateInstance`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	instanceId := "ace712e9-a166-47f1-9065-4af0f7e7fce1" // string | Instance's ID
	ip := "10.0.0.1" // string | 
	updateIpOpts := *openapiclient.NewUpdateIpOpts("ReverseLookup_example") // UpdateIpOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateIp(context.Background(), instanceId, ip).UpdateIpOpts(updateIpOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIp`: IpDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateIp`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	updateLoadBalancerOpts := *openapiclient.NewUpdateLoadBalancerOpts() // UpdateLoadBalancerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateLoadBalancer(context.Background(), loadBalancerId).UpdateLoadBalancerOpts(updateLoadBalancerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancer`: LoadBalancerDetails
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateLoadBalancer`: %v\n", resp)
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	loadBalancerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Load balancer ID
	listenerId := "695ddd91-051f-4dd6-9120-938a927a47d0" // string | Listener ID
	loadBalancerListenerOpts := *openapiclient.NewLoadBalancerListenerOpts() // LoadBalancerListenerOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateLoadBalancerListener(context.Background(), loadBalancerId, listenerId).LoadBalancerListenerOpts(loadBalancerListenerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateLoadBalancerListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancerListener`: LoadBalancerListener
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateLoadBalancerListener`: %v\n", resp)
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


## UpdateTargetGroup

> TargetGroup UpdateTargetGroup(ctx, targetGroupId).UpdateTargetGroupOpts(updateTargetGroupOpts).Execute()

Update Target Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	targetGroupId := "fb769dab-3daa-47e4-89ed-06a4b6499176" // string | Target Group ID
	updateTargetGroupOpts := *openapiclient.NewUpdateTargetGroupOpts() // UpdateTargetGroupOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.UpdateTargetGroup(context.Background(), targetGroupId).UpdateTargetGroupOpts(updateTargetGroupOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.UpdateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTargetGroup`: TargetGroup
	fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.UpdateTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | Target Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTargetGroupOpts** | [**UpdateTargetGroupOpts**](UpdateTargetGroupOpts.md) |  | 

### Return type

[**TargetGroup**](TargetGroup.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

