# \BmpAPI

All URIs are relative to *https://api.leaseweb.com/internal/bmpapi/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServerCredential**](BmpAPI.md#AddServerCredential) | **Post** /servers/{serverId}/credentials | Create new credentials
[**CancelActiveJob**](BmpAPI.md#CancelActiveJob) | **Post** /servers/{serverId}/cancelActiveJob | Cancel active job
[**ConfigureHardwareRaid**](BmpAPI.md#ConfigureHardwareRaid) | **Post** /servers/{serverId}/configureHardwareRaid | Configure Hardware Raid
[**CreateServerDhcpReservation**](BmpAPI.md#CreateServerDhcpReservation) | **Post** /servers/{serverId}/leases | Create a DHCP reservation
[**DeleteNetworkEquipmentCredential**](BmpAPI.md#DeleteNetworkEquipmentCredential) | **Delete** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Delete user credentials
[**DeleteServerCredential**](BmpAPI.md#DeleteServerCredential) | **Delete** /servers/{serverId}/credentials/{type}/{username} | Delete server credential
[**EnableServerRescueMode**](BmpAPI.md#EnableServerRescueMode) | **Post** /servers/{serverId}/rescueMode | Launch rescue mode
[**ExpireActiveJob**](BmpAPI.md#ExpireActiveJob) | **Post** /servers/{serverId}/expireActiveJob | Expire active job
[**GetControlPanelList**](BmpAPI.md#GetControlPanelList) | **Get** /controlPanels | List control panels
[**GetControlPanelListByOperatingSystem**](BmpAPI.md#GetControlPanelListByOperatingSystem) | **Get** /operatingSystems/{operatingSystemId}/controlPanels | List control panels
[**GetNetworkEquipmentCredentialListByType**](BmpAPI.md#GetNetworkEquipmentCredentialListByType) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type} | List credentials by type
[**GetNetworkEquipmentPowerStatus**](BmpAPI.md#GetNetworkEquipmentPowerStatus) | **Get** /networkEquipments/{networkEquipmentId}/powerInfo | Show power status
[**GetOperatingSystem**](BmpAPI.md#GetOperatingSystem) | **Get** /operatingSystems/{operatingSystemId} | Show an operating system
[**GetOperatingSystemList**](BmpAPI.md#GetOperatingSystemList) | **Get** /operatingSystems | List Operating Systems
[**GetRescueImageList**](BmpAPI.md#GetRescueImageList) | **Get** /rescueImages | Rescue Images
[**GetServerCredentialList**](BmpAPI.md#GetServerCredentialList) | **Get** /servers/{serverId}/credentials | List credentials
[**GetServerDhcpReservationList**](BmpAPI.md#GetServerDhcpReservationList) | **Get** /servers/{serverId}/leases | List DHCP reservations
[**GetServerJobList**](BmpAPI.md#GetServerJobList) | **Get** /servers/{serverId}/jobs | List jobs
[**GetServerPowerStatus**](BmpAPI.md#GetServerPowerStatus) | **Get** /servers/{serverId}/powerInfo | Show power status
[**InstallOperatingSystem**](BmpAPI.md#InstallOperatingSystem) | **Post** /servers/{serverId}/install | Launch installation
[**IpmiResetServer**](BmpAPI.md#IpmiResetServer) | **Post** /servers/{serverId}/ipmiReset | Launch IPMI reset
[**NetworkEquipmentsCredentialsGet**](BmpAPI.md#NetworkEquipmentsCredentialsGet) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Show user credentials
[**NetworkEquipmentsCredentialsList**](BmpAPI.md#NetworkEquipmentsCredentialsList) | **Get** /networkEquipments/{networkEquipmentId}/credentials | List credentials
[**NetworkEquipmentsCredentialsPost**](BmpAPI.md#NetworkEquipmentsCredentialsPost) | **Post** /networkEquipments/{networkEquipmentId}/credentials | Create new credentials
[**PowerCycleNetworkEquipment**](BmpAPI.md#PowerCycleNetworkEquipment) | **Post** /networkEquipments/{networkEquipmentId}/powerCycle | Power cycle a network equipment
[**PowerCycleServer**](BmpAPI.md#PowerCycleServer) | **Post** /servers/{serverId}/powerCycle | Power cycle a server
[**PowerNetworkEquipmentOff**](BmpAPI.md#PowerNetworkEquipmentOff) | **Post** /networkEquipments/{networkEquipmentId}/powerOff | Power off network equipment
[**PowerNetworkEquipmentOn**](BmpAPI.md#PowerNetworkEquipmentOn) | **Post** /networkEquipments/{networkEquipmentId}/powerOn | Power on network equipment
[**PowerServerOff**](BmpAPI.md#PowerServerOff) | **Post** /servers/{serverId}/powerOff | Power off server
[**PowerServerOn**](BmpAPI.md#PowerServerOn) | **Post** /servers/{serverId}/powerOn | Power on server
[**RetryServerJob**](BmpAPI.md#RetryServerJob) | **Post** /servers/{serverId}/jobs/{jobId}/retry | Retry a job
[**ScanHardware**](BmpAPI.md#ScanHardware) | **Post** /servers/{serverId}/hardwareScan | Launch hardware scan
[**ServersCredentialsGet**](BmpAPI.md#ServersCredentialsGet) | **Get** /servers/{serverId}/credentials/{type}/{username} | Show user credentials
[**ServersCredentialsListType**](BmpAPI.md#ServersCredentialsListType) | **Get** /servers/{serverId}/credentials/{type} | List credentials by type
[**ServersCredentialsPut**](BmpAPI.md#ServersCredentialsPut) | **Put** /servers/{serverId}/credentials/{type}/{username} | Update user credentials
[**ServersJobsGet**](BmpAPI.md#ServersJobsGet) | **Get** /servers/{serverId}/jobs/{jobId} | Show a job
[**ServersLeasesDelete**](BmpAPI.md#ServersLeasesDelete) | **Delete** /servers/{serverId}/leases | Delete a DHCP reservation
[**UpdateNetworkEquipmentCredential**](BmpAPI.md#UpdateNetworkEquipmentCredential) | **Put** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Update user credentials



## AddServerCredential

> Credential AddServerCredential(ctx, serverId).CreateServerCredentialOpts(createServerCredentialOpts).Execute()

Create new credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	createServerCredentialOpts := *openapiclient.NewCreateServerCredentialOpts("Password_example", openapiclient.credentialType("OPERATING_SYSTEM"), "Username_example") // CreateServerCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.AddServerCredential(context.Background(), serverId).CreateServerCredentialOpts(createServerCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.AddServerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddServerCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.AddServerCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServerCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServerCredentialOpts** | [**CreateServerCredentialOpts**](CreateServerCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.CancelActiveJob(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.CancelActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelActiveJob`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.CancelActiveJob`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigureHardwareRaid

> ServerJobList ConfigureHardwareRaid(ctx, serverId).ConfigureHardwareRaidOpts(configureHardwareRaidOpts).Execute()

Configure Hardware Raid



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	configureHardwareRaidOpts := *openapiclient.NewConfigureHardwareRaidOpts(openapiclient.raidType("HW")) // ConfigureHardwareRaidOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.ConfigureHardwareRaid(context.Background(), serverId).ConfigureHardwareRaidOpts(configureHardwareRaidOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ConfigureHardwareRaid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureHardwareRaid`: ServerJobList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.ConfigureHardwareRaid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureHardwareRaidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configureHardwareRaidOpts** | [**ConfigureHardwareRaidOpts**](ConfigureHardwareRaidOpts.md) |  | 

### Return type

[**ServerJobList**](ServerJobList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	createServerDhcpReservationOpts := *openapiclient.NewCreateServerDhcpReservationOpts("Bootfile_example") // CreateServerDhcpReservationOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.CreateServerDhcpReservation(context.Background(), serverId).CreateServerDhcpReservationOpts(createServerDhcpReservationOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.CreateServerDhcpReservation``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkEquipmentCredential

> DeleteNetworkEquipmentCredential(ctx, networkEquipmentId, type_, username).Execute()

Delete user credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.DeleteNetworkEquipmentCredential(context.Background(), networkEquipmentId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.DeleteNetworkEquipmentCredential``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerCredential

> DeleteServerCredential(ctx, serverId, type_, username).Execute()

Delete server credential



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.DeleteServerCredential(context.Background(), serverId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.DeleteServerCredential``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	enableServerRescueModeOpts := *openapiclient.NewEnableServerRescueModeOpts("RescueImageId_example") // EnableServerRescueModeOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.EnableServerRescueMode(context.Background(), serverId).EnableServerRescueModeOpts(enableServerRescueModeOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.EnableServerRescueMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableServerRescueMode`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.EnableServerRescueMode`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.ExpireActiveJob(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ExpireActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpireActiveJob`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.ExpireActiveJob`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetControlPanelList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetControlPanelList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetControlPanelList`: ControlPanelList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetControlPanelList`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControlPanelListByOperatingSystem

> ControlPanelList GetControlPanelListByOperatingSystem(ctx, operatingSystemId).Limit(limit).Offset(offset).Execute()

List control panels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	operatingSystemId := "UBUNTU_22_04_64BIT" // string | Operating system ID
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetControlPanelListByOperatingSystem(context.Background(), operatingSystemId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetControlPanelListByOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetControlPanelListByOperatingSystem`: ControlPanelList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetControlPanelListByOperatingSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatingSystemId** | **string** | Operating system ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControlPanelListByOperatingSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**ControlPanelList**](ControlPanelList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEquipmentCredentialListByType

> CredentialList GetNetworkEquipmentCredentialListByType(ctx, networkEquipmentId, type_).Limit(limit).Offset(offset).Execute()

List credentials by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetNetworkEquipmentCredentialListByType(context.Background(), networkEquipmentId, type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetNetworkEquipmentCredentialListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentCredentialListByType`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetNetworkEquipmentCredentialListByType`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetNetworkEquipmentPowerStatus(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetNetworkEquipmentPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEquipmentPowerStatus`: GetNetworkEquipmentPowerStatusResult
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetNetworkEquipmentPowerStatus`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	operatingSystemId := "UBUNTU_22_04_64BIT" // string | Operating system ID
	controlPanelId := "controlPanelId_example" // string | The Control Panel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetOperatingSystem(context.Background(), operatingSystemId).ControlPanelId(controlPanelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystem`: GetOperatingSystemResult
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetOperatingSystem`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)
	controlPanelId := "controlPanelId_example" // string | Filter operating systems by control panel id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetOperatingSystemList(context.Background()).Limit(limit).Offset(offset).ControlPanelId(controlPanelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetOperatingSystemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperatingSystemList`: GetOperatingSystemListResult
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetOperatingSystemList`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetRescueImageList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetRescueImageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRescueImageList`: GetRescueImageListResult
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetRescueImageList`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerCredentialList

> CredentialList GetServerCredentialList(ctx, serverId).Limit(limit).Offset(offset).Execute()

List credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetServerCredentialList(context.Background(), serverId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetServerCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCredentialList`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetServerCredentialList`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetServerDhcpReservationList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetServerDhcpReservationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDhcpReservationList`: GetServerDhcpReservationListResult
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetServerDhcpReservationList`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
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
	resp, r, err := apiClient.BmpAPI.GetServerJobList(context.Background(), serverId).Limit(limit).Offset(offset).Type_(type_).Status(status).IsRunning(isRunning).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetServerJobList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerJobList`: ServerJobList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetServerJobList`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.GetServerPowerStatus(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.GetServerPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerPowerStatus`: GetServerPowerStatusResult
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.GetServerPowerStatus`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	installOperatingSystemOpts := *openapiclient.NewInstallOperatingSystemOpts("OperatingSystemId_example") // InstallOperatingSystemOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.InstallOperatingSystem(context.Background(), serverId).InstallOperatingSystemOpts(installOperatingSystemOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.InstallOperatingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallOperatingSystem`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.InstallOperatingSystem`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	ipmiResetServerOpts := *openapiclient.NewIpmiResetServerOpts() // IpmiResetServerOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.IpmiResetServer(context.Background(), serverId).IpmiResetServerOpts(ipmiResetServerOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.IpmiResetServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IpmiResetServer`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.IpmiResetServer`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEquipmentsCredentialsGet

> Credential NetworkEquipmentsCredentialsGet(ctx, networkEquipmentId, type_, username).Execute()

Show user credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.NetworkEquipmentsCredentialsGet(context.Background(), networkEquipmentId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.NetworkEquipmentsCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkEquipmentsCredentialsGet`: Credential
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.NetworkEquipmentsCredentialsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiNetworkEquipmentsCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Credential**](Credential.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEquipmentsCredentialsList

> CredentialList NetworkEquipmentsCredentialsList(ctx, networkEquipmentId).Limit(limit).Offset(offset).Execute()

List credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.NetworkEquipmentsCredentialsList(context.Background(), networkEquipmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.NetworkEquipmentsCredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkEquipmentsCredentialsList`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.NetworkEquipmentsCredentialsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkEquipmentsCredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEquipmentsCredentialsPost

> Credential NetworkEquipmentsCredentialsPost(ctx, networkEquipmentId).CreateNetworkEquipmentCredentialOpts(createNetworkEquipmentCredentialOpts).Execute()

Create new credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	createNetworkEquipmentCredentialOpts := *openapiclient.NewCreateNetworkEquipmentCredentialOpts("Password_example", openapiclient.credentialType("OPERATING_SYSTEM"), "Username_example") // CreateNetworkEquipmentCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.NetworkEquipmentsCredentialsPost(context.Background(), networkEquipmentId).CreateNetworkEquipmentCredentialOpts(createNetworkEquipmentCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.NetworkEquipmentsCredentialsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkEquipmentsCredentialsPost`: Credential
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.NetworkEquipmentsCredentialsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEquipmentId** | **string** | The ID of a dedicated network equipment | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkEquipmentsCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkEquipmentCredentialOpts** | [**CreateNetworkEquipmentCredentialOpts**](CreateNetworkEquipmentCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.PowerCycleNetworkEquipment(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.PowerCycleNetworkEquipment``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.PowerCycleServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.PowerCycleServer``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.PowerNetworkEquipmentOff(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.PowerNetworkEquipmentOff``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.PowerNetworkEquipmentOn(context.Background(), networkEquipmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.PowerNetworkEquipmentOn``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.PowerServerOff(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.PowerServerOff``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.PowerServerOn(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.PowerServerOn``: %v\n", err)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryServerJob

> ServerJob RetryServerJob(ctx, serverId, jobId).Execute()

Retry a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	jobId := "3a867358-5b4b-44ee-88ac-4274603ef641" // string | The ID of a Job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.RetryServerJob(context.Background(), serverId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.RetryServerJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryServerJob`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.RetryServerJob`: %v\n", resp)
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

[**ServerJob**](ServerJob.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	scanHardwareOpts := *openapiclient.NewScanHardwareOpts() // ScanHardwareOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.ScanHardware(context.Background(), serverId).ScanHardwareOpts(scanHardwareOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ScanHardware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanHardware`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.ScanHardware`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersCredentialsGet

> Credential ServersCredentialsGet(ctx, serverId, type_, username).Execute()

Show user credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.ServersCredentialsGet(context.Background(), serverId, type_, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ServersCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersCredentialsGet`: Credential
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.ServersCredentialsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiServersCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Credential**](Credential.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersCredentialsListType

> CredentialList ServersCredentialsListType(ctx, serverId, type_).Limit(limit).Offset(offset).Execute()

List credentials by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	limit := int32(20) // int32 | Limit the number of results returned. (optional)
	offset := int32(10) // int32 | Return results starting from the given offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.ServersCredentialsListType(context.Background(), serverId, type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ServersCredentialsListType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersCredentialsListType`: CredentialList
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.ServersCredentialsListType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**type_** | **string** | Credential type | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersCredentialsListTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limit the number of results returned. | 
 **offset** | **int32** | Return results starting from the given offset. | 

### Return type

[**CredentialList**](CredentialList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersCredentialsPut

> Credential ServersCredentialsPut(ctx, serverId, type_, username).UpdateServerCredentialOpts(updateServerCredentialOpts).Execute()

Update user credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username
	updateServerCredentialOpts := *openapiclient.NewUpdateServerCredentialOpts("Password_example") // UpdateServerCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.ServersCredentialsPut(context.Background(), serverId, type_, username).UpdateServerCredentialOpts(updateServerCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ServersCredentialsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersCredentialsPut`: Credential
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.ServersCredentialsPut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiServersCredentialsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateServerCredentialOpts** | [**UpdateServerCredentialOpts**](UpdateServerCredentialOpts.md) |  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersJobsGet

> ServerJob ServersJobsGet(ctx, serverId, jobId).Execute()

Show a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server
	jobId := "3a867358-5b4b-44ee-88ac-4274603ef641" // string | The ID of a Job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.ServersJobsGet(context.Background(), serverId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ServersJobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersJobsGet`: ServerJob
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.ServersJobsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The ID of a server | 
**jobId** | **string** | The ID of a Job | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerJob**](ServerJob.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersLeasesDelete

> ServersLeasesDelete(ctx, serverId).Execute()

Delete a DHCP reservation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	serverId := "12345" // string | The ID of a server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BmpAPI.ServersLeasesDelete(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.ServersLeasesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServersLeasesDeleteRequest struct via the builder pattern


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


## UpdateNetworkEquipmentCredential

> Credential UpdateNetworkEquipmentCredential(ctx, networkEquipmentId, type_, username).UpdateNetworkEquipmentCredentialOpts(updateNetworkEquipmentCredentialOpts).Execute()

Update user credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/bmp"
)

func main() {
	networkEquipmentId := "12345" // string | The ID of a dedicated network equipment
	type_ := "OPERATING_SYSTEM" // string | Credential type
	username := "root" // string | Username
	updateNetworkEquipmentCredentialOpts := *openapiclient.NewUpdateNetworkEquipmentCredentialOpts("Password_example") // UpdateNetworkEquipmentCredentialOpts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BmpAPI.UpdateNetworkEquipmentCredential(context.Background(), networkEquipmentId, type_, username).UpdateNetworkEquipmentCredentialOpts(updateNetworkEquipmentCredentialOpts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BmpAPI.UpdateNetworkEquipmentCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEquipmentCredential`: Credential
	fmt.Fprintf(os.Stdout, "Response from `BmpAPI.UpdateNetworkEquipmentCredential`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

