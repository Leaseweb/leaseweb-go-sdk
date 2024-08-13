# \DedicatedServerAPI

All URIs are relative to *https://api.leaseweb.com/bareMetals/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServerToPrivateNetwork**](DedicatedServerAPI.md#AddServerToPrivateNetwork) | **Put** /servers/{serverId}/privateNetworks/{privateNetworkId} | Add a server to private network
[**CancelActiveJob**](DedicatedServerAPI.md#CancelActiveJob) | **Post** /servers/{serverId}/cancelActiveJob | Cancel active job
[**CloseNetworkInterface**](DedicatedServerAPI.md#CloseNetworkInterface) | **Post** /servers/{serverId}/networkInterfaces/{networkType}/close | Close network interface
[**CloseNetworkInterfaces**](DedicatedServerAPI.md#CloseNetworkInterfaces) | **Post** /servers/{serverId}/networkInterfaces/close | Close all network interfaces
[**CreateNetworkEquipmentCredential**](DedicatedServerAPI.md#CreateNetworkEquipmentCredential) | **Post** /networkEquipments/{networkEquipmentId}/credentials | Create new network equipment credentials
[**CreateServerBandwidthNotificationSetting**](DedicatedServerAPI.md#CreateServerBandwidthNotificationSetting) | **Post** /servers/{serverId}/notificationSettings/bandwidth | Create a bandwidth notification setting
[**CreateServerCredential**](DedicatedServerAPI.md#CreateServerCredential) | **Post** /servers/{serverId}/credentials | Create new server credentials
[**CreateServerDataTrafficNotificationSetting**](DedicatedServerAPI.md#CreateServerDataTrafficNotificationSetting) | **Post** /servers/{serverId}/notificationSettings/datatraffic | Create a data traffic notification setting
[**CreateServerDhcpReservation**](DedicatedServerAPI.md#CreateServerDhcpReservation) | **Post** /servers/{serverId}/leases | Create a DHCP reservation
[**DeleteNetworkEquipmentCredential**](DedicatedServerAPI.md#DeleteNetworkEquipmentCredential) | **Delete** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Delete network equipment credentials
[**DeleteServerBandwidthNotificationSetting**](DedicatedServerAPI.md#DeleteServerBandwidthNotificationSetting) | **Delete** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Delete a bandwidth notification setting
[**DeleteServerCredential**](DedicatedServerAPI.md#DeleteServerCredential) | **Delete** /servers/{serverId}/credentials/{type}/{username} | Delete server credentials
[**DeleteServerDataTrafficNotificationSetting**](DedicatedServerAPI.md#DeleteServerDataTrafficNotificationSetting) | **Delete** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Delete a data traffic notification setting
[**DeleteServerDhcpReservation**](DedicatedServerAPI.md#DeleteServerDhcpReservation) | **Delete** /servers/{serverId}/leases | Delete a DHCP reservation
[**DeleteServerFromPrivateNetwork**](DedicatedServerAPI.md#DeleteServerFromPrivateNetwork) | **Delete** /servers/{serverId}/privateNetworks/{privateNetworkId} | Delete a server from a private network
[**EnableServerRescueMode**](DedicatedServerAPI.md#EnableServerRescueMode) | **Post** /servers/{serverId}/rescueMode | Launch rescue mode
[**ExpireActiveJob**](DedicatedServerAPI.md#ExpireActiveJob) | **Post** /servers/{serverId}/expireActiveJob | Expire active job
[**GetControlPanelList**](DedicatedServerAPI.md#GetControlPanelList) | **Get** /controlPanels | List control panels
[**GetControlPanelListByOperatingSystemId**](DedicatedServerAPI.md#GetControlPanelListByOperatingSystemId) | **Get** /operatingSystems/{operatingSystemId}/controlPanels | List control panels by Operating System
[**GetDdosNotificationSetting**](DedicatedServerAPI.md#GetDdosNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/ddos | Inspect DDoS notification settings
[**GetNetworkEquipment**](DedicatedServerAPI.md#GetNetworkEquipment) | **Get** /networkEquipments/{networkEquipmentId} | Get network equipment
[**GetNetworkEquipmentCredential**](DedicatedServerAPI.md#GetNetworkEquipmentCredential) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Show network equipment credentials
[**GetNetworkEquipmentCredentialList**](DedicatedServerAPI.md#GetNetworkEquipmentCredentialList) | **Get** /networkEquipments/{networkEquipmentId}/credentials | List network equipment credentials
[**GetNetworkEquipmentCredentialListByType**](DedicatedServerAPI.md#GetNetworkEquipmentCredentialListByType) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type} | List network equipment credentials by type
[**GetNetworkEquipmentIp**](DedicatedServerAPI.md#GetNetworkEquipmentIp) | **Get** /networkEquipments/{networkEquipmentId}/ips/{ip} | Show a network equipment IP
[**GetNetworkEquipmentIpList**](DedicatedServerAPI.md#GetNetworkEquipmentIpList) | **Get** /networkEquipments/{networkEquipmentId}/ips | List IPs
[**GetNetworkEquipmentList**](DedicatedServerAPI.md#GetNetworkEquipmentList) | **Get** /networkEquipments | List network equipment
[**GetNetworkEquipmentNullRouteHistory**](DedicatedServerAPI.md#GetNetworkEquipmentNullRouteHistory) | **Get** /networkEquipments/{networkEquipmentId}/nullRouteHistory | Show null route history
[**GetNetworkEquipmentPowerStatus**](DedicatedServerAPI.md#GetNetworkEquipmentPowerStatus) | **Get** /networkEquipments/{networkEquipmentId}/powerInfo | Show power status
[**GetNetworkInterface**](DedicatedServerAPI.md#GetNetworkInterface) | **Get** /servers/{serverId}/networkInterfaces/{networkType} | Show a network interface
[**GetNetworkInterfaceList**](DedicatedServerAPI.md#GetNetworkInterfaceList) | **Get** /servers/{serverId}/networkInterfaces | List network interfaces
[**GetOperatingSystem**](DedicatedServerAPI.md#GetOperatingSystem) | **Get** /operatingSystems/{operatingSystemId} | Show an operating system
[**GetOperatingSystemList**](DedicatedServerAPI.md#GetOperatingSystemList) | **Get** /operatingSystems | List Operating Systems
[**GetRescueImageList**](DedicatedServerAPI.md#GetRescueImageList) | **Get** /rescueImages | Rescue Images
[**GetServer**](DedicatedServerAPI.md#GetServer) | **Get** /servers/{serverId} | Get server
[**GetServerBandwidthMetrics**](DedicatedServerAPI.md#GetServerBandwidthMetrics) | **Get** /servers/{serverId}/metrics/bandwidth | Show bandwidth metrics
[**GetServerBandwidthNotificationSetting**](DedicatedServerAPI.md#GetServerBandwidthNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Show a bandwidth notification setting
[**GetServerBandwidthNotificationSettingList**](DedicatedServerAPI.md#GetServerBandwidthNotificationSettingList) | **Get** /servers/{serverId}/notificationSettings/bandwidth | List bandwidth notification settings
[**GetServerCredential**](DedicatedServerAPI.md#GetServerCredential) | **Get** /servers/{serverId}/credentials/{type}/{username} | Show server credentials
[**GetServerCredentialList**](DedicatedServerAPI.md#GetServerCredentialList) | **Get** /servers/{serverId}/credentials | List server credentials
[**GetServerCredentialListByType**](DedicatedServerAPI.md#GetServerCredentialListByType) | **Get** /servers/{serverId}/credentials/{type} | List server credentials by type
[**GetServerDataTrafficMetrics**](DedicatedServerAPI.md#GetServerDataTrafficMetrics) | **Get** /servers/{serverId}/metrics/datatraffic | Show datatraffic metrics
[**GetServerDataTrafficNotificationSetting**](DedicatedServerAPI.md#GetServerDataTrafficNotificationSetting) | **Get** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Show a data traffic notification setting
[**GetServerDataTrafficNotificationSettingList**](DedicatedServerAPI.md#GetServerDataTrafficNotificationSettingList) | **Get** /servers/{serverId}/notificationSettings/datatraffic | List data traffic notification settings
[**GetServerDhcpReservationList**](DedicatedServerAPI.md#GetServerDhcpReservationList) | **Get** /servers/{serverId}/leases | List DHCP reservations
[**GetServerHardware**](DedicatedServerAPI.md#GetServerHardware) | **Get** /servers/{serverId}/hardwareInfo | Show hardware information
[**GetServerIp**](DedicatedServerAPI.md#GetServerIp) | **Get** /servers/{serverId}/ips/{ip} | Show a server IP
[**GetServerIpList**](DedicatedServerAPI.md#GetServerIpList) | **Get** /servers/{serverId}/ips | List IPs
[**GetServerJob**](DedicatedServerAPI.md#GetServerJob) | **Get** /servers/{serverId}/jobs/{jobId} | Show a job
[**GetServerJobList**](DedicatedServerAPI.md#GetServerJobList) | **Get** /servers/{serverId}/jobs | List jobs
[**GetServerList**](DedicatedServerAPI.md#GetServerList) | **Get** /servers | List servers
[**GetServerNullRouteHistory**](DedicatedServerAPI.md#GetServerNullRouteHistory) | **Get** /servers/{serverId}/nullRouteHistory | Show null route history
[**GetServerPowerStatus**](DedicatedServerAPI.md#GetServerPowerStatus) | **Get** /servers/{serverId}/powerInfo | Show power status
[**InstallOperatingSystem**](DedicatedServerAPI.md#InstallOperatingSystem) | **Post** /servers/{serverId}/install | Launch installation
[**IpmiResetServer**](DedicatedServerAPI.md#IpmiResetServer) | **Post** /servers/{serverId}/ipmiReset | Launch IPMI reset
[**NullIpRoute**](DedicatedServerAPI.md#NullIpRoute) | **Post** /servers/{serverId}/ips/{ip}/null | Null route an IP
[**NullNetworkEquipmentIpRoute**](DedicatedServerAPI.md#NullNetworkEquipmentIpRoute) | **Post** /networkEquipments/{networkEquipmentId}/ips/{ip}/null | Null route an IP
[**OpenNetworkInterface**](DedicatedServerAPI.md#OpenNetworkInterface) | **Post** /servers/{serverId}/networkInterfaces/{networkType}/open | Open network interface
[**OpenNetworkInterfaces**](DedicatedServerAPI.md#OpenNetworkInterfaces) | **Post** /servers/{serverId}/networkInterfaces/open | Open all network interfaces
[**PowerCycleNetworkEquipment**](DedicatedServerAPI.md#PowerCycleNetworkEquipment) | **Post** /networkEquipments/{networkEquipmentId}/powerCycle | Power cycle a network equipment
[**PowerCycleServer**](DedicatedServerAPI.md#PowerCycleServer) | **Post** /servers/{serverId}/powerCycle | Power cycle a server
[**PowerNetworkEquipmentOff**](DedicatedServerAPI.md#PowerNetworkEquipmentOff) | **Post** /networkEquipments/{networkEquipmentId}/powerOff | Power off network equipment
[**PowerNetworkEquipmentOn**](DedicatedServerAPI.md#PowerNetworkEquipmentOn) | **Post** /networkEquipments/{networkEquipmentId}/powerOn | Power on network equipment
[**PowerServerOff**](DedicatedServerAPI.md#PowerServerOff) | **Post** /servers/{serverId}/powerOff | Power off server
[**PowerServerOn**](DedicatedServerAPI.md#PowerServerOn) | **Post** /servers/{serverId}/powerOn | Power on server
[**RemoveNullIpRoute**](DedicatedServerAPI.md#RemoveNullIpRoute) | **Post** /servers/{serverId}/ips/{ip}/unnull | Remove a null route
[**RetryServerJob**](DedicatedServerAPI.md#RetryServerJob) | **Post** /servers/{serverId}/jobs/{jobId}/retry | Retry a job
[**ScanHardware**](DedicatedServerAPI.md#ScanHardware) | **Post** /servers/{serverId}/hardwareScan | Launch hardware scan
[**UnNullNetworkEquipmentIpRoute**](DedicatedServerAPI.md#UnNullNetworkEquipmentIpRoute) | **Post** /networkEquipments/{networkEquipmentId}/ips/{ip}/unnull | Remove a null route
[**UpdateDdosNotificationSetting**](DedicatedServerAPI.md#UpdateDdosNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/ddos | Update DDoS notification settings
[**UpdateIpProfile**](DedicatedServerAPI.md#UpdateIpProfile) | **Put** /servers/{serverId}/ips/{ip} | Update an IP
[**UpdateNetworkEquipmentCredential**](DedicatedServerAPI.md#UpdateNetworkEquipmentCredential) | **Put** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Update network equipment credentials
[**UpdateNetworkEquipmentIp**](DedicatedServerAPI.md#UpdateNetworkEquipmentIp) | **Put** /networkEquipments/{networkEquipmentId}/ips/{ip} | Update an IP
[**UpdateNetworkEquipmentReference**](DedicatedServerAPI.md#UpdateNetworkEquipmentReference) | **Put** /networkEquipments/{networkEquipmentId} | Update network equipment
[**UpdateServerBandwidthNotificationSetting**](DedicatedServerAPI.md#UpdateServerBandwidthNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/bandwidth/{notificationSettingId} | Update a bandwidth notification setting
[**UpdateServerCredential**](DedicatedServerAPI.md#UpdateServerCredential) | **Put** /servers/{serverId}/credentials/{type}/{username} | Update server credentials
[**UpdateServerDataTrafficNotificationSetting**](DedicatedServerAPI.md#UpdateServerDataTrafficNotificationSetting) | **Put** /servers/{serverId}/notificationSettings/datatraffic/{notificationSettingId} | Update a data traffic notification setting
[**UpdateServerReference**](DedicatedServerAPI.md#UpdateServerReference) | **Put** /servers/{serverId} | Update server



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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	privateNetworkId := "892" // string | The ID of a Private Network
	addServerToPrivateNetworkOpts := *openapiclient.NewAddServerToPrivateNetworkOpts(openapiclient.linkSpeed(100)) // AddServerToPrivateNetworkOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.AddServerToPrivateNetwork(context.Background(), serverId, privateNetworkId).AddServerToPrivateNetworkOpts(addServerToPrivateNetworkOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.AddServerToPrivateNetwork``: %v\n", err)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelActiveJob

> ServerJob CancelActiveJob(ctx, serverId).Execute()

Cancel active job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.CancelActiveJob(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CancelActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelActiveJob`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.CancelActiveJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelActiveJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := openapiclient.networkType("INTERNAL") // NetworkType | The network type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.CloseNetworkInterface(context.Background(), serverId, networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CloseNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkType** | [**NetworkType**](.md) | The network type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseNetworkInterfaceRequest struct via the builder pattern


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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.CloseNetworkInterfaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CloseNetworkInterfaces``: %v\n", err)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkEquipmentCredential

> Credential CreateNetworkEquipmentCredential(ctx, networkEquipmentId).CreateNetworkEquipmentCredentialOpts(createNetworkEquipmentCredentialOpts).Execute()

Create new network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	createNetworkEquipmentCredentialOpts := *openapiclient.NewCreateNetworkEquipmentCredentialOpts("Password_example", "Type_example", "Username_example") // CreateNetworkEquipmentCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.CreateNetworkEquipmentCredential(context.Background(), networkEquipmentId).CreateNetworkEquipmentCredentialOpts(createNetworkEquipmentCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CreateNetworkEquipmentCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkEquipmentCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.CreateNetworkEquipmentCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkEquipmentCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkEquipmentCredentialOpts** | [**CreateNetworkEquipmentCredentialOpts**](CreateNetworkEquipmentCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerBandwidthNotificationSetting

> BandwidthNotificationSetting CreateServerBandwidthNotificationSetting(ctx, serverId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()

Create a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	bandwidthNotificationSettingOpts := *openapiclient.NewBandwidthNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // BandwidthNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.CreateServerBandwidthNotificationSetting(context.Background(), serverId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CreateServerBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.CreateServerBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bandwidthNotificationSettingOpts** | [**BandwidthNotificationSettingOpts**](BandwidthNotificationSettingOpts.md) |  | 

### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerCredential

> Credential CreateServerCredential(ctx, serverId).CreateServerCredentialOpts(createServerCredentialOpts).Execute()

Create new server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	createServerCredentialOpts := *openapiclient.NewCreateServerCredentialOpts("Password_example", "Type_example", "Username_example") // CreateServerCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.CreateServerCredential(context.Background(), serverId).CreateServerCredentialOpts(createServerCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CreateServerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.CreateServerCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServerCredentialOpts** | [**CreateServerCredentialOpts**](CreateServerCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerDataTrafficNotificationSetting

> DataTrafficNotificationSetting CreateServerDataTrafficNotificationSetting(ctx, serverId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()

Create a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	dataTrafficNotificationSettingOpts := *openapiclient.NewDataTrafficNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // DataTrafficNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.CreateServerDataTrafficNotificationSetting(context.Background(), serverId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CreateServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerDataTrafficNotificationSetting`: DataTrafficNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.CreateServerDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataTrafficNotificationSettingOpts** | [**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md) |  | 

### Return type

[**DataTrafficNotificationSetting**](DataTrafficNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerDhcpReservation

> CreateServerDhcpReservation(ctx, serverId).CreateServerDhcpReservationOpts(createServerDhcpReservationOpts).Execute()

Create a DHCP reservation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	createServerDhcpReservationOpts := *openapiclient.NewCreateServerDhcpReservationOpts("Bootfile_example") // CreateServerDhcpReservationOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.CreateServerDhcpReservation(context.Background(), serverId).CreateServerDhcpReservationOpts(createServerDhcpReservationOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.CreateServerDhcpReservation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateServerDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServerDhcpReservationOpts** | [**CreateServerDhcpReservationOpts**](CreateServerDhcpReservationOpts.md) |  | 

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


## DeleteNetworkEquipmentCredential

> DeleteNetworkEquipmentCredential(ctx, networkEquipmentId, type_, username).Execute()

Delete network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.DeleteNetworkEquipmentCredential(context.Background(), networkEquipmentId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.DeleteNetworkEquipmentCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkEquipmentCredentialRequest struct via the builder pattern


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


## DeleteServerBandwidthNotificationSetting

> DeleteServerBandwidthNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Delete a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.DeleteServerBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.DeleteServerBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerBandwidthNotificationSettingRequest struct via the builder pattern


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


## DeleteServerCredential

> DeleteServerCredential(ctx, serverId, type_, username).Execute()

Delete server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.DeleteServerCredential(context.Background(), serverId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.DeleteServerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerCredentialRequest struct via the builder pattern


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


## DeleteServerDataTrafficNotificationSetting

> DeleteServerDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Delete a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.DeleteServerDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.DeleteServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerDataTrafficNotificationSettingRequest struct via the builder pattern


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


## DeleteServerDhcpReservation

> DeleteServerDhcpReservation(ctx, serverId).Execute()

Delete a DHCP reservation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.DeleteServerDhcpReservation(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.DeleteServerDhcpReservation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteServerDhcpReservationRequest struct via the builder pattern


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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	privateNetworkId := "892" // string | The ID of a Private Network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.DeleteServerFromPrivateNetwork(context.Background(), serverId, privateNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.DeleteServerFromPrivateNetwork``: %v\n", err)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableServerRescueMode

> ServerJob EnableServerRescueMode(ctx, serverId).EnableServerRescueModeOpts(enableServerRescueModeOpts).Execute()

Launch rescue mode



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	enableServerRescueModeOpts := *openapiclient.NewEnableServerRescueModeOpts("RescueImageId_example") // EnableServerRescueModeOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.EnableServerRescueMode(context.Background(), serverId).EnableServerRescueModeOpts(enableServerRescueModeOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.EnableServerRescueMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableServerRescueMode`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.EnableServerRescueMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableServerRescueModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableServerRescueModeOpts** | [**EnableServerRescueModeOpts**](EnableServerRescueModeOpts.md) |  | 

### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireActiveJob

> ServerJob ExpireActiveJob(ctx, serverId).Execute()

Expire active job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.ExpireActiveJob(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.ExpireActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpireActiveJob`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.ExpireActiveJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireActiveJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControlPanelList

> ControlPanelList GetControlPanelList(ctx).Limit(limit).Offset(offset).Execute()

List control panels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetControlPanelList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetControlPanelList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetControlPanelList`: ControlPanelList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetControlPanelList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetControlPanelListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**ControlPanelList**](ControlPanelList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControlPanelListByOperatingSystemId

> ControlPanelList GetControlPanelListByOperatingSystemId(ctx, operatingSystemId).Limit(limit).Offset(offset).Execute()

List control panels by Operating System



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	operatingSystemId := "UBUNTU_22_04_64BIT" // string | Operating system ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetControlPanelListByOperatingSystemId(context.Background(), operatingSystemId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetControlPanelListByOperatingSystemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetControlPanelListByOperatingSystemId`: ControlPanelList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetControlPanelListByOperatingSystemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatingSystemId** | **string** | Operating system ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControlPanelListByOperatingSystemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**ControlPanelList**](ControlPanelList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDdosNotificationSetting

> GetDdosNotificationSettingResult GetDdosNotificationSetting(ctx, serverId).Execute()

Inspect DDoS notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetDdosNotificationSetting(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetDdosNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDdosNotificationSetting`: GetDdosNotificationSettingResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetDdosNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDdosNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDdosNotificationSettingResult**](GetDdosNotificationSettingResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipment(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipment`: NetworkEquipment
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipment`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentCredential

> Credential GetNetworkEquipmentCredential(ctx, networkEquipmentId, type_, username).Execute()

Show network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentCredential(context.Background(), networkEquipmentId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentCredentialList

> CredentialList GetNetworkEquipmentCredentialList(ctx, networkEquipmentId).Limit(limit).Offset(offset).Execute()

List network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentCredentialList(context.Background(), networkEquipmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentCredentialList`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentCredentialListByType

> CredentialList GetNetworkEquipmentCredentialListByType(ctx, networkEquipmentId, type_).Limit(limit).Offset(offset).Execute()

List network equipment credentials by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentCredentialListByType(context.Background(), networkEquipmentId, type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentCredentialListByType`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentCredentialListByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**type_** | **string** | Credential type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentCredentialListByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentIp

> Ip GetNetworkEquipmentIp(ctx, networkEquipmentId, ip).Execute()

Show a network equipment IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentIp(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentIp`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	networkType := openapiclient.networkType("INTERNAL") // NetworkType | Filter the collection of ip addresses by network type (optional)
	version := "version_example" // string | Filter the collection by ip version (optional)
	nullRouted := "nullRouted_example" // string | Filter Ips by Nulled-Status (optional)
	ips := "ips_example" // string | Filter the collection of Ips for the comma separated list of Ips (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentIpList(context.Background(), networkEquipmentId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentIpList`: GetNetworkEquipmentIpListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentIpList`: %v\n", resp)
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

 **networkType** | [**NetworkType**](NetworkType.md) | Filter the collection of ip addresses by network type | 
 **version** | **string** | Filter the collection by ip version | 
 **nullRouted** | **string** | Filter Ips by Nulled-Status | 
 **ips** | **string** | Filter the collection of Ips for the comma separated list of Ips | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetNetworkEquipmentIpListResult**](GetNetworkEquipmentIpListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
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
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentList(context.Background()).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentList`: GetNetworkEquipmentListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentList`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentNullRouteHistory(context.Background(), networkEquipmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentNullRouteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentNullRouteHistory`: GetNetworkEquipmentNullRouteHistoryResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentNullRouteHistory`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentPowerStatus

> GetNetworkEquipmentPowerStatusResult GetNetworkEquipmentPowerStatus(ctx, networkEquipmentId).Execute()

Show power status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkEquipmentPowerStatus(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkEquipmentPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentPowerStatus`: GetNetworkEquipmentPowerStatusResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkEquipmentPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEquipmentPowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkEquipmentPowerStatusResult**](GetNetworkEquipmentPowerStatusResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterface

> OperationNetworkInterface GetNetworkInterface(ctx, serverId, networkType).Execute()

Show a network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := openapiclient.networkType("INTERNAL") // NetworkType | The network type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkInterface(context.Background(), serverId, networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterface`: OperationNetworkInterface
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkType** | [**NetworkType**](.md) | The network type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationNetworkInterface**](OperationNetworkInterface.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetNetworkInterfaceList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetNetworkInterfaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterfaceList`: GetNetworkInterfaceListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetNetworkInterfaceList`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatingSystem

> GetOperatingSystemResult GetOperatingSystem(ctx, operatingSystemId).ControlPanelId(controlPanelId).Execute()

Show an operating system



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	operatingSystemId := "UBUNTU_22_04_64BIT" // string | Operating system ID
	controlPanelId := "controlPanelId_example" // string | The Control Panel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetOperatingSystem(context.Background(), operatingSystemId).ControlPanelId(controlPanelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystem`: GetOperatingSystemResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetOperatingSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatingSystemId** | **string** | Operating system ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatingSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **controlPanelId** | **string** | The Control Panel ID | 

### Return type

[**GetOperatingSystemResult**](GetOperatingSystemResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatingSystemList

> GetOperatingSystemListResult GetOperatingSystemList(ctx).Limit(limit).Offset(offset).ControlPanelId(controlPanelId).Execute()

List Operating Systems



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	controlPanelId := "controlPanelId_example" // string | Filter operating systems by control panel id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetOperatingSystemList(context.Background()).Limit(limit).Offset(offset).ControlPanelId(controlPanelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetOperatingSystemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystemList`: GetOperatingSystemListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetOperatingSystemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatingSystemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **controlPanelId** | **string** | Filter operating systems by control panel id | 

### Return type

[**GetOperatingSystemListResult**](GetOperatingSystemListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRescueImageList

> GetRescueImageListResult GetRescueImageList(ctx).Limit(limit).Offset(offset).Execute()

Rescue Images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetRescueImageList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetRescueImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRescueImageList`: GetRescueImageListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetRescueImageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRescueImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetRescueImageListResult**](GetRescueImageListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServer`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerBandwidthMetrics

> Metrics GetServerBandwidthMetrics(ctx, serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()

Show bandwidth metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	from := time.Now() // time.Time | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time.
	to := time.Now() // time.Time | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time.
	aggregation := "aggregation_example" // string | Aggregate each metric using the given aggregation function. When the aggregation type `95TH` is specified the granularity parameter should be omitted from the request.
	granularity := "granularity_example" // string | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerBandwidthMetrics(context.Background(), serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerBandwidthMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerBandwidthMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerBandwidthMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerBandwidthMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time. | 
 **to** | **time.Time** | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time. | 
 **aggregation** | **string** | Aggregate each metric using the given aggregation function. When the aggregation type &#x60;95TH&#x60; is specified the granularity parameter should be omitted from the request. | 
 **granularity** | **string** | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerBandwidthNotificationSetting

> BandwidthNotificationSetting GetServerBandwidthNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Show a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerBandwidthNotificationSettingList

> GetServerBandwidthNotificationSettingListResult GetServerBandwidthNotificationSettingList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List bandwidth notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerBandwidthNotificationSettingList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerBandwidthNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerBandwidthNotificationSettingList`: GetServerBandwidthNotificationSettingListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerBandwidthNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerBandwidthNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetServerBandwidthNotificationSettingListResult**](GetServerBandwidthNotificationSettingListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerCredential

> Credential GetServerCredential(ctx, serverId, type_, username).Execute()

Show server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerCredential(context.Background(), serverId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerCredentialList

> CredentialList GetServerCredentialList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerCredentialList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCredentialList`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerCredentialListByType

> CredentialList GetServerCredentialListByType(ctx, serverId, type_).Limit(limit).Offset(offset).Execute()

List server credentials by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerCredentialListByType(context.Background(), serverId, type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCredentialListByType`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerCredentialListByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | **string** | Credential type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCredentialListByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDataTrafficMetrics

> Metrics GetServerDataTrafficMetrics(ctx, serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()

Show datatraffic metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	from := time.Now() // time.Time | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time.
	to := time.Now() // time.Time | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time.
	aggregation := "aggregation_example" // string | Aggregate each metric using the given aggregation function.
	granularity := "granularity_example" // string | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerDataTrafficMetrics(context.Background(), serverId).From(from).To(to).Aggregation(aggregation).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerDataTrafficMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDataTrafficMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerDataTrafficMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDataTrafficMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of date interval in ISO-8601 format. The returned data will include everything up from - and including - the specified date time. | 
 **to** | **time.Time** | End of date interval in ISO-8601 format. The returned data will include everything up until - but not including - the specified date time. | 
 **aggregation** | **string** | Aggregate each metric using the given aggregation function. | 
 **granularity** | **string** | Specify the preferred interval for each metric. If granularity is omitted from the request, only one metric is returned. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDataTrafficNotificationSetting

> DataTrafficNotificationSetting GetServerDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).Execute()

Show a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDataTrafficNotificationSetting`: DataTrafficNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataTrafficNotificationSetting**](DataTrafficNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDataTrafficNotificationSettingList

> GetServerDataTrafficNotificationSettingListResult GetServerDataTrafficNotificationSettingList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List data traffic notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerDataTrafficNotificationSettingList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerDataTrafficNotificationSettingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDataTrafficNotificationSettingList`: GetServerDataTrafficNotificationSettingListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerDataTrafficNotificationSettingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDataTrafficNotificationSettingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**GetServerDataTrafficNotificationSettingListResult**](GetServerDataTrafficNotificationSettingListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDhcpReservationList

> GetServerDhcpReservationListResult GetServerDhcpReservationList(ctx, serverId).Execute()

List DHCP reservations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerDhcpReservationList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerDhcpReservationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDhcpReservationList`: GetServerDhcpReservationListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerDhcpReservationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDhcpReservationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServerDhcpReservationListResult**](GetServerDhcpReservationListResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerHardware(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerHardware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerHardware`: GetServerHardwareResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerHardware`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerIp

> Ip GetServerIp(ctx, serverId, ip).Execute()

Show a server IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerIp(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerIp`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := openapiclient.networkType("INTERNAL") // NetworkType | Filter the collection of ip addresses by network type (optional)
	version := "version_example" // string | Filter the collection by ip version (optional)
	nullRouted := "nullRouted_example" // string | Filter Ips by Nulled-Status (optional)
	ips := "ips_example" // string | Filter the collection of Ips for the comma separated list of Ips (optional)
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerIpList(context.Background(), serverId).NetworkType(networkType).Version(version).NullRouted(nullRouted).Ips(ips).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerIpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerIpList`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerIpList`: %v\n", resp)
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

 **networkType** | [**NetworkType**](NetworkType.md) | Filter the collection of ip addresses by network type | 
 **version** | **string** | Filter the collection by ip version | 
 **nullRouted** | **string** | Filter Ips by Nulled-Status | 
 **ips** | **string** | Filter the collection of Ips for the comma separated list of Ips | 
 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

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


## GetServerJob

> CurrentServerJob GetServerJob(ctx, serverId, jobId).Execute()

Show a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	jobId := "3a867358-5b4b-44ee-88ac-4274603ef641" // string | The ID of a Job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerJob(context.Background(), serverId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerJob`: CurrentServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**jobId** | **string** | The ID of a Job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CurrentServerJob**](CurrentServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerJobList

> ServerJobList GetServerJobList(ctx, serverId).Limit(limit).Offset(offset).Type_(type_).Status(status).IsRunning(isRunning).Execute()

List jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	type_ := "install" // string | Filter the list of jobs by type. (optional)
	status := "CANCELED" // string | Filter the list of jobs by status. (optional)
	isRunning := "true" // string | Filter the list for running jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerJobList(context.Background(), serverId).Limit(limit).Offset(offset).Type_(type_).Status(status).IsRunning(isRunning).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerJobList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerJobList`: ServerJobList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerJobList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerJobListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 
 **type_** | **string** | Filter the list of jobs by type. | 
 **status** | **string** | Filter the list of jobs by status. | 
 **isRunning** | **string** | Filter the list for running jobs | 

### Return type

[**ServerJobList**](ServerJobList.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
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
	resp, r, err := apiClient.DedicatedServerAPI.GetServerList(context.Background()).Limit(limit).Offset(offset).Reference(reference).Ip(ip).MacAddress(macAddress).Site(site).PrivateRackId(privateRackId).PrivateNetworkCapable(privateNetworkCapable).PrivateNetworkEnabled(privateNetworkEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerList`: GetServerListResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerList`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerNullRouteHistory(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerNullRouteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerNullRouteHistory`: GetServerNullRouteHistoryResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerNullRouteHistory`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerPowerStatus

> GetServerPowerStatusResult GetServerPowerStatus(ctx, serverId).Execute()

Show power status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.GetServerPowerStatus(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.GetServerPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerPowerStatus`: GetServerPowerStatusResult
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.GetServerPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerPowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServerPowerStatusResult**](GetServerPowerStatusResult.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallOperatingSystem

> ServerJob InstallOperatingSystem(ctx, serverId).InstallOperatingSystemOpts(installOperatingSystemOpts).Execute()

Launch installation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	installOperatingSystemOpts := *openapiclient.NewInstallOperatingSystemOpts("OperatingSystemId_example") // InstallOperatingSystemOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.InstallOperatingSystem(context.Background(), serverId).InstallOperatingSystemOpts(installOperatingSystemOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.InstallOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallOperatingSystem`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.InstallOperatingSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallOperatingSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **installOperatingSystemOpts** | [**InstallOperatingSystemOpts**](InstallOperatingSystemOpts.md) |  | 

### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpmiResetServer

> ServerJob IpmiResetServer(ctx, serverId).IpmiResetServerOpts(ipmiResetServerOpts).Execute()

Launch IPMI reset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ipmiResetServerOpts := *openapiclient.NewIpmiResetServerOpts() // IpmiResetServerOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.IpmiResetServer(context.Background(), serverId).IpmiResetServerOpts(ipmiResetServerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.IpmiResetServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IpmiResetServer`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.IpmiResetServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpmiResetServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipmiResetServerOpts** | [**IpmiResetServerOpts**](IpmiResetServerOpts.md) |  | 

### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.NullIpRoute(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.NullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.NullIpRoute`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NullNetworkEquipmentIpRoute

> Ip NullNetworkEquipmentIpRoute(ctx, networkEquipmentId, ip).Execute()

Null route an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.NullNetworkEquipmentIpRoute(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.NullNetworkEquipmentIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NullNetworkEquipmentIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.NullNetworkEquipmentIpRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiNullNetworkEquipmentIpRouteRequest struct via the builder pattern


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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	networkType := openapiclient.networkType("INTERNAL") // NetworkType | The network type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.OpenNetworkInterface(context.Background(), serverId, networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.OpenNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**networkType** | [**NetworkType**](.md) | The network type | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenNetworkInterfaceRequest struct via the builder pattern


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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.OpenNetworkInterfaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.OpenNetworkInterfaces``: %v\n", err)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PowerCycleNetworkEquipment

> PowerCycleNetworkEquipment(ctx, networkEquipmentId).Execute()

Power cycle a network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.PowerCycleNetworkEquipment(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.PowerCycleNetworkEquipment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPowerCycleNetworkEquipmentRequest struct via the builder pattern


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


## PowerCycleServer

> PowerCycleServer(ctx, serverId).Execute()

Power cycle a server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.PowerCycleServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.PowerCycleServer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPowerCycleServerRequest struct via the builder pattern


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


## PowerNetworkEquipmentOff

> PowerNetworkEquipmentOff(ctx, networkEquipmentId).Execute()

Power off network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.PowerNetworkEquipmentOff(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.PowerNetworkEquipmentOff``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPowerNetworkEquipmentOffRequest struct via the builder pattern


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


## PowerNetworkEquipmentOn

> PowerNetworkEquipmentOn(ctx, networkEquipmentId).Execute()

Power on network equipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.PowerNetworkEquipmentOn(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.PowerNetworkEquipmentOn``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPowerNetworkEquipmentOnRequest struct via the builder pattern


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


## PowerServerOff

> PowerServerOff(ctx, serverId).Execute()

Power off server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.PowerServerOff(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.PowerServerOff``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPowerServerOffRequest struct via the builder pattern


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


## PowerServerOn

> PowerServerOn(ctx, serverId).Execute()

Power on server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.PowerServerOn(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.PowerServerOn``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPowerServerOnRequest struct via the builder pattern


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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.RemoveNullIpRoute(context.Background(), serverId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.RemoveNullIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNullIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.RemoveNullIpRoute`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryServerJob

> CurrentServerJob RetryServerJob(ctx, serverId, jobId).Execute()

Retry a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	jobId := "3a867358-5b4b-44ee-88ac-4274603ef641" // string | The ID of a Job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.RetryServerJob(context.Background(), serverId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.RetryServerJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryServerJob`: CurrentServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.RetryServerJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**jobId** | **string** | The ID of a Job | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryServerJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CurrentServerJob**](CurrentServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanHardware

> ServerJob ScanHardware(ctx, serverId).ScanHardwareOpts(scanHardwareOpts).Execute()

Launch hardware scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	scanHardwareOpts := *openapiclient.NewScanHardwareOpts() // ScanHardwareOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.ScanHardware(context.Background(), serverId).ScanHardwareOpts(scanHardwareOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.ScanHardware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanHardware`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.ScanHardware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanHardwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanHardwareOpts** | [**ScanHardwareOpts**](ScanHardwareOpts.md) |  | 

### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.UnNullNetworkEquipmentIpRoute(context.Background(), networkEquipmentId, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UnNullNetworkEquipmentIpRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnNullNetworkEquipmentIpRoute`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.UnNullNetworkEquipmentIpRoute`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDdosNotificationSetting

> UpdateDdosNotificationSetting(ctx, serverId).UpdateDdosNotificationSettingOpts(updateDdosNotificationSettingOpts).Execute()

Update DDoS notification settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	updateDdosNotificationSettingOpts := *openapiclient.NewUpdateDdosNotificationSettingOpts() // UpdateDdosNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.UpdateDdosNotificationSetting(context.Background(), serverId).UpdateDdosNotificationSettingOpts(updateDdosNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateDdosNotificationSetting``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateDdosNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDdosNotificationSettingOpts** | [**UpdateDdosNotificationSettingOpts**](UpdateDdosNotificationSettingOpts.md) |  | 

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


## UpdateIpProfile

> Ip UpdateIpProfile(ctx, serverId, ip).UpdateIpProfileOpts(updateIpProfileOpts).Execute()

Update an IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ip := "127.0.0.6" // string | The IP Address
	updateIpProfileOpts := *openapiclient.NewUpdateIpProfileOpts() // UpdateIpProfileOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.UpdateIpProfile(context.Background(), serverId, ip).UpdateIpProfileOpts(updateIpProfileOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateIpProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIpProfile`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.UpdateIpProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**ip** | **string** | The IP Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIpProfileOpts** | [**UpdateIpProfileOpts**](UpdateIpProfileOpts.md) |  | 

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


## UpdateNetworkEquipmentCredential

> Credential UpdateNetworkEquipmentCredential(ctx, networkEquipmentId, type_, username).UpdateNetworkEquipmentCredentialOpts(updateNetworkEquipmentCredentialOpts).Execute()

Update network equipment credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username
	updateNetworkEquipmentCredentialOpts := *openapiclient.NewUpdateNetworkEquipmentCredentialOpts("Password_example") // UpdateNetworkEquipmentCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.UpdateNetworkEquipmentCredential(context.Background(), networkEquipmentId, type_, username).UpdateNetworkEquipmentCredentialOpts(updateNetworkEquipmentCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateNetworkEquipmentCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEquipmentCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.UpdateNetworkEquipmentCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkEquipmentCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkEquipmentCredentialOpts** | [**UpdateNetworkEquipmentCredentialOpts**](UpdateNetworkEquipmentCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	ip := "127.0.0.6" // string | The IP Address
	updateNetworkEquipmentIpOpts := *openapiclient.NewUpdateNetworkEquipmentIpOpts() // UpdateNetworkEquipmentIpOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.UpdateNetworkEquipmentIp(context.Background(), networkEquipmentId, ip).UpdateNetworkEquipmentIpOpts(updateNetworkEquipmentIpOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateNetworkEquipmentIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEquipmentIp`: Ip
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.UpdateNetworkEquipmentIp`: %v\n", resp)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	updateNetworkEquipmentReferenceOpts := *openapiclient.NewUpdateNetworkEquipmentReferenceOpts("Reference_example") // UpdateNetworkEquipmentReferenceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.UpdateNetworkEquipmentReference(context.Background(), networkEquipmentId).UpdateNetworkEquipmentReferenceOpts(updateNetworkEquipmentReferenceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateNetworkEquipmentReference``: %v\n", err)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerBandwidthNotificationSetting

> BandwidthNotificationSetting UpdateServerBandwidthNotificationSetting(ctx, serverId, notificationSettingId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()

Update a bandwidth notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting
	bandwidthNotificationSettingOpts := *openapiclient.NewBandwidthNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // BandwidthNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.UpdateServerBandwidthNotificationSetting(context.Background(), serverId, notificationSettingId).BandwidthNotificationSettingOpts(bandwidthNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateServerBandwidthNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerBandwidthNotificationSetting`: BandwidthNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.UpdateServerBandwidthNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerBandwidthNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandwidthNotificationSettingOpts** | [**BandwidthNotificationSettingOpts**](BandwidthNotificationSettingOpts.md) |  | 

### Return type

[**BandwidthNotificationSetting**](BandwidthNotificationSetting.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerCredential

> Credential UpdateServerCredential(ctx, serverId, type_, username).UpdateServerCredentialOpts(updateServerCredentialOpts).Execute()

Update server credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username
	updateServerCredentialOpts := *openapiclient.NewUpdateServerCredentialOpts("Password_example") // UpdateServerCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.UpdateServerCredential(context.Background(), serverId, type_, username).UpdateServerCredentialOpts(updateServerCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateServerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.UpdateServerCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | **string** | Credential type | 
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateServerCredentialOpts** | [**UpdateServerCredentialOpts**](UpdateServerCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDataTrafficNotificationSetting

> DataTrafficNotificationSettingOpts UpdateServerDataTrafficNotificationSetting(ctx, serverId, notificationSettingId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()

Update a data traffic notification setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	notificationSettingId := "839111" // string | The ID of a notification setting
	dataTrafficNotificationSettingOpts := *openapiclient.NewDataTrafficNotificationSettingOpts("Frequency_example", "Threshold_example", "Unit_example") // DataTrafficNotificationSettingOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedServerAPI.UpdateServerDataTrafficNotificationSetting(context.Background(), serverId, notificationSettingId).DataTrafficNotificationSettingOpts(dataTrafficNotificationSettingOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateServerDataTrafficNotificationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerDataTrafficNotificationSetting`: DataTrafficNotificationSettingOpts
	fmt.Fprintf(os.Stdout, "Response from `DedicatedServerAPI.UpdateServerDataTrafficNotificationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**notificationSettingId** | **string** | The ID of a notification setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDataTrafficNotificationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataTrafficNotificationSettingOpts** | [**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md) |  | 

### Return type

[**DataTrafficNotificationSettingOpts**](DataTrafficNotificationSettingOpts.md)

### Authorization

[X-LSW-Auth](../README.md#X-LSW-Auth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dedicatedServer"
)

func main() {
	serverId := "12345" // string | The ID of a server
	updateServerReferenceOpts := *openapiclient.NewUpdateServerReferenceOpts("Reference_example") // UpdateServerReferenceOpts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DedicatedServerAPI.UpdateServerReference(context.Background(), serverId).UpdateServerReferenceOpts(updateServerReferenceOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServerAPI.UpdateServerReference``: %v\n", err)
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

[X-LSW-Auth](../README.md#X-LSW-Auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

