# Go API client for bmp

This documents the rest api bmp api provides.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2
- Package version: 1.0.0
- Generator version: 7.4.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import bmp "github.com/leaseweb/leaseweb-go-sdk/bmp"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `bmp.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), bmp.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `bmp.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), bmp.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `bmp.ContextOperationServerIndices` and `bmp.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), bmp.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), bmp.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.leaseweb.com/internal/bmpapi/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BmpAPI* | [**AddServerCredential**](docs/BmpAPI.md#addservercredential) | **Post** /servers/{serverId}/credentials | Create new credentials
*BmpAPI* | [**CancelActiveJob**](docs/BmpAPI.md#cancelactivejob) | **Post** /servers/{serverId}/cancelActiveJob | Cancel active job
*BmpAPI* | [**ConfigureHardwareRaid**](docs/BmpAPI.md#configurehardwareraid) | **Post** /servers/{serverId}/configureHardwareRaid | Configure Hardware Raid
*BmpAPI* | [**CreateServerDhcpReservation**](docs/BmpAPI.md#createserverdhcpreservation) | **Post** /servers/{serverId}/leases | Create a DHCP reservation
*BmpAPI* | [**DeleteNetworkEquipmentCredential**](docs/BmpAPI.md#deletenetworkequipmentcredential) | **Delete** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Delete user credentials
*BmpAPI* | [**DeleteServerCredential**](docs/BmpAPI.md#deleteservercredential) | **Delete** /servers/{serverId}/credentials/{type}/{username} | Delete server credential
*BmpAPI* | [**EnableServerRescueMode**](docs/BmpAPI.md#enableserverrescuemode) | **Post** /servers/{serverId}/rescueMode | Launch rescue mode
*BmpAPI* | [**ExpireActiveJob**](docs/BmpAPI.md#expireactivejob) | **Post** /servers/{serverId}/expireActiveJob | Expire active job
*BmpAPI* | [**GetControlPanelList**](docs/BmpAPI.md#getcontrolpanellist) | **Get** /controlPanels | List control panels
*BmpAPI* | [**GetControlPanelListByOperatingSystem**](docs/BmpAPI.md#getcontrolpanellistbyoperatingsystem) | **Get** /operatingSystems/{operatingSystemId}/controlPanels | List control panels
*BmpAPI* | [**GetNetworkEquipmentCredentialListByType**](docs/BmpAPI.md#getnetworkequipmentcredentiallistbytype) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type} | List credentials by type
*BmpAPI* | [**GetNetworkEquipmentPowerStatus**](docs/BmpAPI.md#getnetworkequipmentpowerstatus) | **Get** /networkEquipments/{networkEquipmentId}/powerInfo | Show power status
*BmpAPI* | [**GetOperatingSystem**](docs/BmpAPI.md#getoperatingsystem) | **Get** /operatingSystems/{operatingSystemId} | Show an operating system
*BmpAPI* | [**GetOperatingSystemList**](docs/BmpAPI.md#getoperatingsystemlist) | **Get** /operatingSystems | List Operating Systems
*BmpAPI* | [**GetRescueImageList**](docs/BmpAPI.md#getrescueimagelist) | **Get** /rescueImages | Rescue Images
*BmpAPI* | [**GetServerCredentialList**](docs/BmpAPI.md#getservercredentiallist) | **Get** /servers/{serverId}/credentials | List credentials
*BmpAPI* | [**GetServerDhcpReservationList**](docs/BmpAPI.md#getserverdhcpreservationlist) | **Get** /servers/{serverId}/leases | List DHCP reservations
*BmpAPI* | [**GetServerJobList**](docs/BmpAPI.md#getserverjoblist) | **Get** /servers/{serverId}/jobs | List jobs
*BmpAPI* | [**GetServerPowerStatus**](docs/BmpAPI.md#getserverpowerstatus) | **Get** /servers/{serverId}/powerInfo | Show power status
*BmpAPI* | [**InstallOperatingSystem**](docs/BmpAPI.md#installoperatingsystem) | **Post** /servers/{serverId}/install | Launch installation
*BmpAPI* | [**IpmiResetServer**](docs/BmpAPI.md#ipmiresetserver) | **Post** /servers/{serverId}/ipmiReset | Launch IPMI reset
*BmpAPI* | [**NetworkEquipmentsCredentialsGet**](docs/BmpAPI.md#networkequipmentscredentialsget) | **Get** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Show user credentials
*BmpAPI* | [**NetworkEquipmentsCredentialsList**](docs/BmpAPI.md#networkequipmentscredentialslist) | **Get** /networkEquipments/{networkEquipmentId}/credentials | List credentials
*BmpAPI* | [**NetworkEquipmentsCredentialsPost**](docs/BmpAPI.md#networkequipmentscredentialspost) | **Post** /networkEquipments/{networkEquipmentId}/credentials | Create new credentials
*BmpAPI* | [**PowerCycleNetworkEquipment**](docs/BmpAPI.md#powercyclenetworkequipment) | **Post** /networkEquipments/{networkEquipmentId}/powerCycle | Power cycle a network equipment
*BmpAPI* | [**PowerCycleServer**](docs/BmpAPI.md#powercycleserver) | **Post** /servers/{serverId}/powerCycle | Power cycle a server
*BmpAPI* | [**PowerNetworkEquipmentOff**](docs/BmpAPI.md#powernetworkequipmentoff) | **Post** /networkEquipments/{networkEquipmentId}/powerOff | Power off network equipment
*BmpAPI* | [**PowerNetworkEquipmentOn**](docs/BmpAPI.md#powernetworkequipmenton) | **Post** /networkEquipments/{networkEquipmentId}/powerOn | Power on network equipment
*BmpAPI* | [**PowerServerOff**](docs/BmpAPI.md#powerserveroff) | **Post** /servers/{serverId}/powerOff | Power off server
*BmpAPI* | [**PowerServerOn**](docs/BmpAPI.md#powerserveron) | **Post** /servers/{serverId}/powerOn | Power on server
*BmpAPI* | [**RetryServerJob**](docs/BmpAPI.md#retryserverjob) | **Post** /servers/{serverId}/jobs/{jobId}/retry | Retry a job
*BmpAPI* | [**ScanHardware**](docs/BmpAPI.md#scanhardware) | **Post** /servers/{serverId}/hardwareScan | Launch hardware scan
*BmpAPI* | [**ServersCredentialsGet**](docs/BmpAPI.md#serverscredentialsget) | **Get** /servers/{serverId}/credentials/{type}/{username} | Show user credentials
*BmpAPI* | [**ServersCredentialsListType**](docs/BmpAPI.md#serverscredentialslisttype) | **Get** /servers/{serverId}/credentials/{type} | List credentials by type
*BmpAPI* | [**ServersCredentialsPut**](docs/BmpAPI.md#serverscredentialsput) | **Put** /servers/{serverId}/credentials/{type}/{username} | Update user credentials
*BmpAPI* | [**ServersJobsGet**](docs/BmpAPI.md#serversjobsget) | **Get** /servers/{serverId}/jobs/{jobId} | Show a job
*BmpAPI* | [**ServersLeasesDelete**](docs/BmpAPI.md#serversleasesdelete) | **Delete** /servers/{serverId}/leases | Delete a DHCP reservation
*BmpAPI* | [**UpdateNetworkEquipmentCredential**](docs/BmpAPI.md#updatenetworkequipmentcredential) | **Put** /networkEquipments/{networkEquipmentId}/credentials/{type}/{username} | Update user credentials


## Documentation For Models

 - [ConfigureHardwareRaidOpts](docs/ConfigureHardwareRaidOpts.md)
 - [ControlPanel](docs/ControlPanel.md)
 - [ControlPanelList](docs/ControlPanelList.md)
 - [CreateNetworkEquipmentCredentialOpts](docs/CreateNetworkEquipmentCredentialOpts.md)
 - [CreateServerCredentialOpts](docs/CreateServerCredentialOpts.md)
 - [CreateServerDhcpReservationOpts](docs/CreateServerDhcpReservationOpts.md)
 - [Credential](docs/Credential.md)
 - [CredentialList](docs/CredentialList.md)
 - [CredentialType](docs/CredentialType.md)
 - [Defaults](docs/Defaults.md)
 - [EnableServerRescueModeOpts](docs/EnableServerRescueModeOpts.md)
 - [ErrorResult](docs/ErrorResult.md)
 - [GetNetworkEquipmentPowerStatusResult](docs/GetNetworkEquipmentPowerStatusResult.md)
 - [GetOperatingSystemListResult](docs/GetOperatingSystemListResult.md)
 - [GetOperatingSystemResult](docs/GetOperatingSystemResult.md)
 - [GetRescueImageListResult](docs/GetRescueImageListResult.md)
 - [GetServerDhcpReservationListResult](docs/GetServerDhcpReservationListResult.md)
 - [GetServerPowerStatusResult](docs/GetServerPowerStatusResult.md)
 - [InstallOperatingSystemOpts](docs/InstallOperatingSystemOpts.md)
 - [Ipmi](docs/Ipmi.md)
 - [IpmiResetServerOpts](docs/IpmiResetServerOpts.md)
 - [LastClientRequest](docs/LastClientRequest.md)
 - [Lease](docs/Lease.md)
 - [Metadata](docs/Metadata.md)
 - [OperatingSystem](docs/OperatingSystem.md)
 - [Partition](docs/Partition.md)
 - [Pdu](docs/Pdu.md)
 - [Raid](docs/Raid.md)
 - [RaidType](docs/RaidType.md)
 - [RescueImage](docs/RescueImage.md)
 - [ScanHardwareOpts](docs/ScanHardwareOpts.md)
 - [ServerJob](docs/ServerJob.md)
 - [ServerJobList](docs/ServerJobList.md)
 - [Task](docs/Task.md)
 - [UpdateNetworkEquipmentCredentialOpts](docs/UpdateNetworkEquipmentCredentialOpts.md)
 - [UpdateServerCredentialOpts](docs/UpdateServerCredentialOpts.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-LSW-Auth
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-LSW-Auth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		bmp.ContextAPIKeys,
		map[string]bmp.APIKey{
			"X-LSW-Auth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### OAuth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), bmp.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, bmp.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), bmp.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

development-networkautomation@leaseweb.com
