package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDedicatedServerList(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 1}, "servers": [
			{
				"assetId": "627293",
				"contract": {
					"customerId": "32923828192",
					"deliveryStatus": "ACTIVE",
					"id": "674382",
					"reference": "database.server",
					"salesOrgId": "2300"
				},
				"featureAvailability": {
					"automation": true,
					"ipmiReboot": false,
					"powerCycle": true,
					"privateNetwork": true,
					"remoteManagement": false
				},
				"id": "12345",
				"location": {
					"rack": "A83",
					"site": "AMS-01",
					"suite": "99",
					"unit": "16-17"
				},
				"networkInterfaces": {
					"internal": {
						"gateway": "10.22.192.12",
						"ip": "10.22.192.3",
						"mac": "AA:BB:CC:DD:EE:FF",
						"ports": [
							{
								"name": "EVO-AABB-01",
								"port": "30"
							}
						]
					},
					"public": {
						"gateway": "95.211.162.62",
						"ip": "95.211.162.0",
						"mac": "AA:AC:CC:88:EE:E4",
						"ports": []
					},
					"remoteManagement": {
						"gateway": "10.22.192.126",
						"ip": "10.22.192.1",
						"mac": "AA:AC:CC:88:EE:E4",
						"ports": []
					}
				},
				"powerPorts": [
					{
						"name": "EVO-JV12-APC02",
						"port": "10"
					}
				],
				"privateNetworks": [
					{
						"id": "1",
						"linkSpeed": 1000,
						"status": "CONFIGURED",
						"subnet": "127.0.0.80/24",
						"vlanId": "2120"
					}
				],
				"rack": {
					"type": "SHARED"
				}
			}
		]}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.List(ctx, DedicatedServerListOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Servers), 1)

	Server := response.Servers[0]
	assert.Equal(Server.AssetId, "627293")
	assert.Equal(Server.Id, "12345")
	assert.Equal(Server.Contract.CustomerId, "32923828192")
	assert.Equal(Server.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(Server.Contract.Id, "674382")
	assert.Equal(Server.Contract.Reference, "database.server")
	assert.Equal(Server.Contract.SalesOrgId, "2300")
	assert.Equal(Server.FeatureAvailability.Automation, true)
	assert.Equal(Server.FeatureAvailability.IpmiReboot, false)
	assert.Equal(Server.FeatureAvailability.PowerCycle, true)
	assert.Equal(Server.FeatureAvailability.PrivateNetwork, true)
	assert.Equal(Server.FeatureAvailability.RemoteManagement, false)
	assert.Equal(Server.Location.Rack, "A83")
	assert.Equal(Server.Location.Site, "AMS-01")
	assert.Equal(Server.Location.Suite, "99")
	assert.Equal(Server.Location.Unit, "16-17")

	Internal := Server.NetworkInterfaces.Internal
	assert.Equal(Internal.Gateway, "10.22.192.12")
	assert.Equal(Internal.Ip, "10.22.192.3")
	assert.Equal(Internal.Mac, "AA:BB:CC:DD:EE:FF")
	assert.Equal(len(Internal.Ports), 1)
	assert.Equal(Internal.Ports[0].Name, "EVO-AABB-01")
	assert.Equal(Internal.Ports[0].Port, "30")

	Public := Server.NetworkInterfaces.Public
	assert.Equal(Public.Gateway, "95.211.162.62")
	assert.Equal(Public.Ip, "95.211.162.0")
	assert.Equal(Public.Mac, "AA:AC:CC:88:EE:E4")
	assert.Equal(len(Public.Ports), 0)

	RemoteManagement := Server.NetworkInterfaces.RemoteManagement
	assert.Equal(RemoteManagement.Gateway, "10.22.192.126")
	assert.Equal(RemoteManagement.Ip, "10.22.192.1")
	assert.Equal(RemoteManagement.Mac, "AA:AC:CC:88:EE:E4")

	assert.Equal(len(Public.Ports), 0)
	assert.Equal(len(Server.PowerPorts), 1)
	assert.Equal(Server.PowerPorts[0].Name, "EVO-JV12-APC02")
	assert.Equal(Server.PowerPorts[0].Port, "10")
	assert.Equal(len(Server.PrivateNetworks), 1)
	assert.Equal(Server.PrivateNetworks[0].Id, "1")
	assert.Equal(Server.PrivateNetworks[0].LinkSpeed.String(), "1000")
	assert.Equal(Server.PrivateNetworks[0].Status, "CONFIGURED")
	assert.Equal(Server.PrivateNetworks[0].Subnet, "127.0.0.80/24")
	assert.Equal(Server.PrivateNetworks[0].VLanId, "2120")
	assert.Equal(Server.Rack.Type, "SHARED")
}

func TestDedicatedServerListBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "servers": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedServerApi{}.List(ctx, DedicatedServerListOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Servers), 0)
}

func TestDedicatedServerListPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "servers": [
			{
				"assetId": "627293",
				"contract": {
					"customerId": "32923828192",
					"deliveryStatus": "ACTIVE",
					"id": "674382",
					"reference": "database.server",
					"salesOrgId": "2300"
				},
				"featureAvailability": {
					"automation": true,
					"ipmiReboot": false,
					"powerCycle": true,
					"privateNetwork": true,
					"remoteManagement": false
				},
				"id": "12345",
				"location": {
					"rack": "A83",
					"site": "AMS-01",
					"suite": "99",
					"unit": "16-17"
				},
				"networkInterfaces": {
					"internal": {
						"gateway": "10.22.192.12",
						"ip": "10.22.192.3",
						"mac": "AA:BB:CC:DD:EE:FF",
						"ports": [
							{
								"name": "EVO-AABB-01",
								"port": "30"
							}
						]
					},
					"public": {
						"gateway": "95.211.162.62",
						"ip": "95.211.162.0",
						"mac": "AA:AC:CC:88:EE:E4",
						"ports": []
					},
					"remoteManagement": {
						"gateway": "10.22.192.126",
						"ip": "10.22.192.1",
						"mac": "AA:AC:CC:88:EE:E4",
						"ports": []
					}
				},
				"powerPorts": [
					{
						"name": "EVO-JV12-APC02",
						"port": "10"
					}
				],
				"privateNetworks": [
					{
						"id": "1",
						"linkSpeed": 1000,
						"status": "CONFIGURED",
						"subnet": "127.0.0.80/24",
						"vlanId": "2120"
					}
				],
				"rack": {
					"type": "SHARED"
				}
			}
		]}`)
	})
	defer teardown()

	ctx := context.Background()

	opts := DedicatedServerListOptions{
		PaginationOptions: PaginationOptions{
			Offset: Int(0),
			Limit:  Int(10),
		},
		Site:       String("AMS-01"),
		IP:         String("10.22.192.3"),
		MacAddress: String("AA:BB:CC:DD:EE:FF"),
	}

	response, err := DedicatedServerApi{}.List(ctx, opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Servers), 1)

	Server := response.Servers[0]
	assert.Equal(Server.AssetId, "627293")
	assert.Equal(Server.Id, "12345")
	assert.Equal(Server.Contract.CustomerId, "32923828192")
	assert.Equal(Server.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(Server.Contract.Id, "674382")
	assert.Equal(Server.Contract.Reference, "database.server")
	assert.Equal(Server.Contract.SalesOrgId, "2300")
	assert.Equal(Server.FeatureAvailability.Automation, true)
	assert.Equal(Server.FeatureAvailability.IpmiReboot, false)
	assert.Equal(Server.FeatureAvailability.PowerCycle, true)
	assert.Equal(Server.FeatureAvailability.PrivateNetwork, true)
	assert.Equal(Server.FeatureAvailability.RemoteManagement, false)
	assert.Equal(Server.Location.Rack, "A83")
	assert.Equal(Server.Location.Site, "AMS-01")
	assert.Equal(Server.Location.Suite, "99")
	assert.Equal(Server.Location.Unit, "16-17")

	Internal := Server.NetworkInterfaces.Internal
	assert.Equal(Internal.Gateway, "10.22.192.12")
	assert.Equal(Internal.Ip, "10.22.192.3")
	assert.Equal(Internal.Mac, "AA:BB:CC:DD:EE:FF")
	assert.Equal(len(Internal.Ports), 1)
	assert.Equal(Internal.Ports[0].Name, "EVO-AABB-01")
	assert.Equal(Internal.Ports[0].Port, "30")

	Public := Server.NetworkInterfaces.Public
	assert.Equal(Public.Gateway, "95.211.162.62")
	assert.Equal(Public.Ip, "95.211.162.0")
	assert.Equal(Public.Mac, "AA:AC:CC:88:EE:E4")
	assert.Equal(len(Public.Ports), 0)

	RemoteManagement := Server.NetworkInterfaces.RemoteManagement
	assert.Equal(RemoteManagement.Gateway, "10.22.192.126")
	assert.Equal(RemoteManagement.Ip, "10.22.192.1")
	assert.Equal(RemoteManagement.Mac, "AA:AC:CC:88:EE:E4")

	assert.Equal(len(Public.Ports), 0)
	assert.Equal(len(Server.PowerPorts), 1)
	assert.Equal(Server.PowerPorts[0].Name, "EVO-JV12-APC02")
	assert.Equal(Server.PowerPorts[0].Port, "10")
	assert.Equal(len(Server.PrivateNetworks), 1)
	assert.Equal(Server.PrivateNetworks[0].Id, "1")
	assert.Equal(Server.PrivateNetworks[0].LinkSpeed.String(), "1000")
	assert.Equal(Server.PrivateNetworks[0].Status, "CONFIGURED")
	assert.Equal(Server.PrivateNetworks[0].Subnet, "127.0.0.80/24")
	assert.Equal(Server.PrivateNetworks[0].VLanId, "2120")
	assert.Equal(Server.Rack.Type, "SHARED")
}

func TestDedicatedServerListServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.List(ctx, DedicatedServerListOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.List(ctx, DedicatedServerListOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.List(ctx, DedicatedServerListOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.List(ctx, DedicatedServerListOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "12345",
			"assetId": "627294",
			"serialNumber": "JDK18291JK",
			"contract": {
				"billingCycle": 12,
				"billingFrequency": "MONTH",
				"contractTerm": 12,
				"currency": "EUR",
				"customerId": "32923828192",
				"deliveryStatus": "ACTIVE",
				"endsAt": "2017-10-01T01:00:00+0100",
				"id": "674382",
				"networkTraffic": {
					"datatrafficLimit": 100,
					"datatrafficUnit": "TB",
					"trafficType": "PREMIUM",
					"type": "FLATFEE"
				},
				"pricePerFrequency": 49,
				"privateNetworks": [
					{
						"id": "1",
						"linkSpeed": 1000,
						"status": "CONFIGURED",
						"subnet": "127.0.0.80/24",
						"vlanId": "2120"
					}
				],
				"reference": "database.server",
				"salesOrgId": "2300",
				"sla": "BRONZE",
				"softwareLicenses": [
					{
						"currency": "EUR",
						"name": "WINDOWS_2012_R2_SERVER",
						"price": 12.12
					}
				],
				"startsAt": "2014-01-01T01:00:00+0100"
			},
			"featureAvailability": {
				"automation": true,
				"ipmiReboot": false,
				"powerCycle": true,
				"privateNetwork": true,
				"remoteManagement": false
			},
			"location": {
				"rack": "13",
				"site": "AMS-01",
				"suite": "A6",
				"unit": "16-17"
			},
			"networkInterfaces": {
				"internal": {
					"gateway": "123.123.123.126",
					"ip": "123.123.123.123/27",
					"mac": "AA:BB:CC:DD:EE:FF",
					"ports": [ ]
				},
				"public": {
					"gateway": "123.123.123.126",
					"ip": "123.123.123.123/27",
					"mac": "AA:BB:CC:DD:EE:FF",
					"ports": [
						{
							"name": "EVO-JV12-1",
							"port": "33"
						}
					]
				},
				"remoteManagement": null
			},
			"powerPorts": [
				{
					"name": "EVO-JV12-APC02",
					"port": "10"
				}
			],
			"rack": {
				"type": "PRIVATE"
			},
			"specs": {
				"chassis": "Dell R210 II",
				"cpu": {
					"quantity": 4,
					"type": "Intel Xeon E3-1220"
				},
				"hardwareRaidCapable": true,
				"hdd": [
					{
						"amount": 2,
						"id": "SATA2TB",
						"performanceType": null,
						"size": 2,
						"type": "SATA",
						"unit": "TB"
					}
				],
				"pciCards": [
					{
						"description": "2x10GE UTP card"
					},
					{
						"description": "2x30GE UTP card"
					}
				],
				"ram": {
					"size": 32,
					"unit": "GB"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Server, err := DedicatedServerApi{}.Get(ctx, "12345")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Server.Id, "12345")
	assert.Equal(Server.AssetId, "627294")
	assert.Equal(Server.SerialNumber, "JDK18291JK")
	assert.Equal(Server.Contract.CustomerId, "32923828192")
	assert.Equal(Server.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(Server.Contract.Id, "674382")
	assert.Equal(Server.Contract.Reference, "database.server")
	assert.Equal(Server.Contract.SalesOrgId, "2300")
	assert.Equal(Server.Contract.StartsAt, "2014-01-01T01:00:00+0100")
	assert.Equal(Server.Contract.Sla, "BRONZE")

	SoftwareLicense := Server.Contract.SoftwareLicenses[0]
	assert.Equal(len(Server.Contract.SoftwareLicenses), 1)
	assert.Equal(SoftwareLicense.Currency, "EUR")
	assert.Equal(SoftwareLicense.Name, "WINDOWS_2012_R2_SERVER")
	assert.Equal(SoftwareLicense.Price.String(), "12.12")
	assert.Equal(Server.Contract.BillingCycle.String(), "12")
	assert.Equal(Server.Contract.BillingFrequency, "MONTH")
	assert.Equal(Server.Contract.ContractTerm.String(), "12")
	assert.Equal(Server.Contract.Currency, "EUR")
	assert.Equal(Server.Contract.EndsAt, "2017-10-01T01:00:00+0100")
	assert.Equal(Server.Contract.PricePerFrequency.String(), "49")
	assert.Equal(Server.Contract.NetworkTraffic.DataTrafficLimit.String(), "100")
	assert.Equal(Server.Contract.NetworkTraffic.DataTrafficUnit, "TB")
	assert.Equal(Server.Contract.NetworkTraffic.TrafficType, "PREMIUM")
	assert.Equal(Server.Contract.NetworkTraffic.Type, "FLATFEE")

	PrivateNetwork := Server.Contract.PrivateNetworks[0]
	assert.Equal(len(Server.Contract.PrivateNetworks), 1)
	assert.Equal(PrivateNetwork.Id, "1")
	assert.Equal(PrivateNetwork.LinkSpeed.String(), "1000")
	assert.Equal(PrivateNetwork.Status, "CONFIGURED")
	assert.Equal(PrivateNetwork.Subnet, "127.0.0.80/24")
	assert.Equal(PrivateNetwork.VLanId, "2120")

	assert.Equal(Server.FeatureAvailability.Automation, true)
	assert.Equal(Server.FeatureAvailability.IpmiReboot, false)
	assert.Equal(Server.FeatureAvailability.PowerCycle, true)
	assert.Equal(Server.FeatureAvailability.PrivateNetwork, true)
	assert.Equal(Server.FeatureAvailability.RemoteManagement, false)
	assert.Equal(Server.Location.Rack, "13")
	assert.Equal(Server.Location.Site, "AMS-01")
	assert.Equal(Server.Location.Suite, "A6")
	assert.Equal(Server.Location.Unit, "16-17")

	Internal := Server.NetworkInterfaces.Internal
	assert.Equal(Internal.Gateway, "123.123.123.126")
	assert.Equal(Internal.Ip, "123.123.123.123/27")
	assert.Equal(Internal.Mac, "AA:BB:CC:DD:EE:FF")
	assert.Equal(len(Internal.Ports), 0)

	Public := Server.NetworkInterfaces.Public
	assert.Equal(Public.Gateway, "123.123.123.126")
	assert.Equal(Public.Ip, "123.123.123.123/27")
	assert.Equal(Public.Mac, "AA:BB:CC:DD:EE:FF")
	assert.Equal(len(Public.Ports), 1)
	assert.Equal(Public.Ports[0].Name, "EVO-JV12-1")
	assert.Equal(Public.Ports[0].Port, "33")

	assert.Empty(Server.NetworkInterfaces.RemoteManagement)

	assert.Equal(len(Server.PowerPorts), 1)
	assert.Equal(Server.PowerPorts[0].Name, "EVO-JV12-APC02")
	assert.Equal(Server.PowerPorts[0].Port, "10")

	assert.Equal(Server.Rack.Type, "PRIVATE")

	assert.Equal(Server.Specs.Chassis, "Dell R210 II")
	assert.Equal(Server.Specs.HardwareRaidCapable, true)
	assert.Equal(Server.Specs.Cpu.Quantity.String(), "4")
	assert.Equal(Server.Specs.Cpu.Type, "Intel Xeon E3-1220")
	assert.Equal(Server.Specs.Hdd[0].Amount.String(), "2")
	assert.Equal(Server.Specs.Hdd[0].Id, "SATA2TB")
	assert.Empty(Server.Specs.Hdd[0].PerformanceType)
	assert.Equal(Server.Specs.Hdd[0].Size.String(), "2")
	assert.Equal(Server.Specs.Hdd[0].Type, "SATA")
	assert.Equal(Server.Specs.Hdd[0].Unit, "TB")
	assert.Equal(Server.Specs.Ram.Size.String(), "32")
	assert.Equal(Server.Specs.Ram.Unit, "GB")
	assert.Equal(Server.Specs.PciCards[0].Description, "2x10GE UTP card")
	assert.Equal(Server.Specs.PciCards[1].Description, "2x30GE UTP card")
}

func TestDedicatedServerGetServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.Get(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.Get(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.Get(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.Get(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.Get(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerUpdate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	payload := map[string]interface{}{"reference": "new reference"}
	ctx := context.Background()
	err := DedicatedServerApi{}.Update(ctx, "12345", payload)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerUpdateServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"reference": "new reference"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.Update(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"reference": "new reference"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.Update(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"reference": "new reference"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.Update(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"reference": "new reference"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.Update(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"reference": "new reference"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.Update(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetHardwareInformation(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "2378237",
			"parserVersion": "3.6",
			"scannedAt": "2017-09-27T14:21:01Z",
			"serverId": "62264",
			"result": {
				"chassis": {
					"description": "Rack Mount Chassis",
					"firmware": {
						"date": "07/01/2013",
						"description": "BIOS",
						"vendor": "HP",
						"version": "J01"
					},
					"motherboard": {
						"product": "111",
						"serial": "222",
						"vendor": "333"
					},
					"product": "ProLiant DL120 G7 (647339-B21)",
					"serial": "CZ33109CHV",
					"vendor": "HP"
				},
				"cpu": [
					{
						"capabilities": {
							"cpufreq": "CPU Frequency scaling",
							"ht": "HyperThreading",
							"vmx": false,
							"x86-64": "64bits extensions (x86-64)"
						},
						"description": "Intel(R) Xeon(R) CPU E31230",
						"hz": "2792640000",
						"serial_number": "123456",
						"settings": {
							"cores": "4",
							"enabledcores": "4",
							"threads": "8"
						},
						"slot": "Proc 1",
						"vendor": "Intel Corp."
					}
				],
				"disks": [
					{
						"description": "ATA Disk",
						"id": "disk:0",
						"product": "Hitachi HDS72302",
						"serial_number": "MS77215W07S6SA",
						"size": "2000398934016",
						"smartctl": {
							"ata_version": "ATA8-ACS T13/1699-D revision 4",
							"attributes": {
								"Power_On_Hours": {
									"flag": "0x0012",
									"id": "9",
									"raw_value": "39832",
									"thresh": "000",
									"type": "Old_age",
									"updated": "Always",
									"value": "095",
									"when_failed": "-",
									"worst": "095"
								},
								"Reallocated_Sector_Ct": {
									"flag": "0x0033",
									"id": "5",
									"raw_value": "0",
									"thresh": "005",
									"type": "Pre-fail",
									"updated": "Always",
									"value": "100",
									"when_failed": "-",
									"worst": "100"
								}
							},
							"device_model": "Hitachi HDS723020BLE640",
							"execution_status": "0",
							"firmware_version": "MX4OAAB0",
							"is_sas": false,
							"overall_health": "PASSED",
							"rpm": "7200 rpm",
							"sata_version": "SATA 3.0, 6.0 Gb/s (current: 6.0 Gb/s)",
							"sector_size": "512 bytes logical, 4096 bytes physical",
							"serial_number": "MS77215W07S6SA",
							"smart_error_log": "No Errors Logged",
							"smart_support": {
								"available": true,
								"enabled": true
							},
							"smartctl_version": "6.2",
							"user_capacity": "2,000,398,934,016 bytes [2.00 TB]"
						},
						"vendor": "Hitachi"
					}
				],
				"ipmi": {
					"defgateway": "10.19.79.126",
					"firmware": "1.88",
					"ipaddress": "10.19.79.67",
					"ipsource": "DHCP Address",
					"macaddress": "28:92:4a:33:48:e8",
					"subnetmask": "255.255.255.192",
					"vendor": "Hewlett-Packard"
				},
				"memory": [
					{
						"clock_hz": "1333000000",
						"description": "DIMM DDR3 Synchronous 1333 MHz (0.8 ns)",
						"id": "memory/bank:0",
						"serial_number": "8369AF58",
						"size_bytes": "4294967296"
					},
					{
						"clock_hz": "1333000000",
						"description": "DIMM DDR3 Synchronous 1333 MHz (0.8 ns)",
						"id": "memory/bank:1",
						"serial_number": "8369B174",
						"size_bytes": "4294967296"
					}
				],
				"network": [
					{
						"capabilities": {
							"autonegotiation": "Auto-negotiation",
							"bus_master": "bus mastering",
							"cap_list": "PCI capabilities listing",
							"ethernet": "123",
							"link_speeds": {
								"1000bt-fd": "1Gbit/s (full duplex)",
								"100bt": "100Mbit/s",
								"100bt-fd": "100Mbit/s (full duplex)",
								"10bt": "10Mbit/s",
								"10bt-fd": "10Mbit/s (full duplex)"
							},
							"msi": "Message Signalled Interrupts",
							"msix": "MSI-X",
							"pciexpress": "PCI Express",
							"physical": "Physical interface",
							"pm": "Power Management",
							"tp": "twisted pair"
						},
						"lldp": {
							"chassis": {
								"description": "Juniper Networks, Inc. ex3300-48t Ethernet Switch, kernel JUNOS 15.1R5.5, Build date: 2016-11-25 16:02:59 UTC Copyright (c) 1996-2016 Juniper Networks, Inc.",
								"mac_address": "4c:16:fc:3a:84:c0",
								"name": "EVO-NS19-1"
							},
							"port": {
								"auto_negotiation": {
									"enabled": "yes",
									"supported": "yes"
								},
								"description": "ge-0/0/2.0"
							},
							"vlan": {
								"id": "0",
								"label": "VLAN",
								"name": "default"
							}
						},
						"logical_name": "eth0",
						"mac_address": "28:92:4a:33:48:e6",
						"product": "82574L Gigabit Network Connection",
						"settings": {
							"autonegotiation": "on",
							"broadcast": "yes",
							"driver": "e1000e",
							"driverversion": "3.2.6-k",
							"duplex": "full",
							"firmware": "2.1-2",
							"ip": "212.32.230.67",
							"latency": "0",
							"link": "yes",
							"multicast": "yes",
							"port": "twisted pair",
							"speed": "1Gbit/s"
						},
						"vendor": "Intel Corporation"
					}
				]
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Server, err := DedicatedServerApi{}.GetHardwareInformation(ctx, "2378237")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Server.Id, "2378237")
	assert.Equal(Server.ParserVersion, "3.6")
	assert.Equal(Server.ScannedAt, "2017-09-27T14:21:01Z")
	assert.Equal(Server.ServerId, "62264")

	assert.Equal(Server.Result.Chassis.Description, "Rack Mount Chassis")
	assert.Equal(Server.Result.Chassis.Product, "ProLiant DL120 G7 (647339-B21)")
	assert.Equal(Server.Result.Chassis.Serial, "CZ33109CHV")
	assert.Equal(Server.Result.Chassis.Vendor, "HP")
	assert.Equal(Server.Result.Chassis.Motherboard.Product, "111")
	assert.Equal(Server.Result.Chassis.Motherboard.Serial, "222")
	assert.Equal(Server.Result.Chassis.Motherboard.Vendor, "333")
	assert.Equal(Server.Result.Chassis.Firmware.Date, "07/01/2013")
	assert.Equal(Server.Result.Chassis.Firmware.Description, "BIOS")
	assert.Equal(Server.Result.Chassis.Firmware.Vendor, "HP")
	assert.Equal(Server.Result.Chassis.Firmware.Version, "J01")

	assert.Equal(len(Server.Result.Cpu), 1)
	Cpu := Server.Result.Cpu[0]
	assert.Equal(Cpu.Capabilities.CpuFreq, "CPU Frequency scaling")
	assert.Equal(Cpu.Capabilities.HT, "HyperThreading")
	assert.Equal(Cpu.Capabilities.VMX, false)
	assert.Equal(Cpu.Capabilities.X8664, "64bits extensions (x86-64)")
	assert.Equal(Cpu.Settings.Cores, "4")
	assert.Equal(Cpu.Settings.EnabledCores, "4")
	assert.Equal(Cpu.Settings.Threads, "8")
	assert.Equal(Cpu.Description, "Intel(R) Xeon(R) CPU E31230")
	assert.Equal(Cpu.HZ, "2792640000")
	assert.Equal(Cpu.SerialNumber, "123456")
	assert.Equal(Cpu.Slot, "Proc 1")
	assert.Equal(Cpu.Vendor, "Intel Corp.")

	assert.Equal(len(Server.Result.Disks), 1)
	Disk := Server.Result.Disks[0]
	assert.Equal(Disk.Description, "ATA Disk")
	assert.Equal(Disk.Id, "disk:0")
	assert.Equal(Disk.Product, "Hitachi HDS72302")
	assert.Equal(Disk.SerialNumber, "MS77215W07S6SA")
	assert.Equal(Disk.Size, "2000398934016")
	assert.Equal(Disk.SmartCTL.ATAVersion, "ATA8-ACS T13/1699-D revision 4")
	assert.Equal(Disk.SmartCTL.DeviceModel, "Hitachi HDS723020BLE640")
	assert.Equal(Disk.SmartCTL.ExecutionStatus, "0")
	assert.Equal(Disk.SmartCTL.FirmwareVersion, "MX4OAAB0")
	assert.Equal(Disk.SmartCTL.ISSAS, false)
	assert.Equal(Disk.SmartCTL.OverallHealth, "PASSED")
	assert.Equal(Disk.SmartCTL.RPM, "7200 rpm")
	assert.Equal(Disk.SmartCTL.SATAVersion, "SATA 3.0, 6.0 Gb/s (current: 6.0 Gb/s)")
	assert.Equal(Disk.SmartCTL.SectorSize, "512 bytes logical, 4096 bytes physical")
	assert.Equal(Disk.SmartCTL.SerialNumber, "MS77215W07S6SA")
	assert.Equal(Disk.SmartCTL.SmartErrorLog, "No Errors Logged")
	assert.Equal(Disk.SmartCTL.SmartctlVersion, "6.2")
	assert.Equal(Disk.SmartCTL.UserCapacity, "2,000,398,934,016 bytes [2.00 TB]")
	assert.Equal(Disk.SmartCTL.SmartSupport.Available, true)
	assert.Equal(Disk.SmartCTL.SmartSupport.Enabled, true)
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.Flag, "0x0033")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.Id, "5")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.RawValue, "0")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.Thresh, "005")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.Type, "Pre-fail")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.Updated, "Always")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.Value, "100")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.WhenFailed, "-")
	assert.Equal(Disk.SmartCTL.Attributes.ReallocatedSectorCT.Worst, "100")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.Flag, "0x0012")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.Id, "9")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.RawValue, "39832")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.Thresh, "000")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.Type, "Old_age")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.Updated, "Always")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.Value, "095")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.WhenFailed, "-")
	assert.Equal(Disk.SmartCTL.Attributes.PowerOnHours.Worst, "095")
	assert.Equal(Disk.Vendor, "Hitachi")
	assert.Equal(Server.Result.Ipmi.Defgateway, "10.19.79.126")
	assert.Equal(Server.Result.Ipmi.Firmware, "1.88")
	assert.Equal(Server.Result.Ipmi.IpAddress, "10.19.79.67")
	assert.Equal(Server.Result.Ipmi.IpSource, "DHCP Address")
	assert.Equal(Server.Result.Ipmi.MacAddress, "28:92:4a:33:48:e8")
	assert.Equal(Server.Result.Ipmi.SubnetMask, "255.255.255.192")
	assert.Equal(Server.Result.Ipmi.Vendor, "Hewlett-Packard")

	assert.Equal(len(Server.Result.Memories), 2)
	Memory1 := Server.Result.Memories[0]
	assert.Equal(Memory1.ClockHZ, "1333000000")
	assert.Equal(Memory1.Description, "DIMM DDR3 Synchronous 1333 MHz (0.8 ns)")
	assert.Equal(Memory1.Id, "memory/bank:0")
	assert.Equal(Memory1.SerialNumber, "8369AF58")
	assert.Equal(Memory1.SizeBytes, "4294967296")

	Memory2 := Server.Result.Memories[1]
	assert.Equal(Memory2.ClockHZ, "1333000000")
	assert.Equal(Memory2.Description, "DIMM DDR3 Synchronous 1333 MHz (0.8 ns)")
	assert.Equal(Memory2.Id, "memory/bank:1")
	assert.Equal(Memory2.SerialNumber, "8369B174")
	assert.Equal(Memory2.SizeBytes, "4294967296")

	assert.Equal(len(Server.Result.Networks), 1)
	Network := Server.Result.Networks[0]
	assert.Equal(Network.Vendor, "Intel Corporation")
	assert.Equal(Network.LogicalName, "eth0")
	assert.Equal(Network.MacAddress, "28:92:4a:33:48:e6")
	assert.Equal(Network.Product, "82574L Gigabit Network Connection")

	assert.Equal(Network.Capabilities.AutoNegotiation, "Auto-negotiation")
	assert.Equal(Network.Capabilities.BusMaster, "bus mastering")
	assert.Equal(Network.Capabilities.CapList, "PCI capabilities listing")
	assert.Equal(Network.Capabilities.Ethernet, "123")
	assert.Equal(Network.Capabilities.LinkSpeeds.BT10, "10Mbit/s")
	assert.Equal(Network.Capabilities.LinkSpeeds.BT100, "100Mbit/s")
	assert.Equal(Network.Capabilities.LinkSpeeds.BTFD10, "10Mbit/s (full duplex)")
	assert.Equal(Network.Capabilities.LinkSpeeds.BTFD100, "100Mbit/s (full duplex)")
	assert.Equal(Network.Capabilities.LinkSpeeds.BTFD1000, "1Gbit/s (full duplex)")
	assert.Equal(Network.Capabilities.MSI, "Message Signalled Interrupts")
	assert.Equal(Network.Capabilities.MSIX, "MSI-X")
	assert.Equal(Network.Capabilities.PciExpress, "PCI Express")
	assert.Equal(Network.Capabilities.Physical, "Physical interface")
	assert.Equal(Network.Capabilities.PM, "Power Management")
	assert.Equal(Network.Capabilities.TP, "twisted pair")

	assert.Equal(Network.LLDP.Chassis.Description, "Juniper Networks, Inc. ex3300-48t Ethernet Switch, kernel JUNOS 15.1R5.5, Build date: 2016-11-25 16:02:59 UTC Copyright (c) 1996-2016 Juniper Networks, Inc.")
	assert.Equal(Network.LLDP.Chassis.MacAddress, "4c:16:fc:3a:84:c0")
	assert.Equal(Network.LLDP.Chassis.Name, "EVO-NS19-1")
	assert.Equal(Network.LLDP.Port.AutoNegotiation.Enabled, "yes")
	assert.Equal(Network.LLDP.Port.AutoNegotiation.Supported, "yes")
	assert.Equal(Network.LLDP.Port.Description, "ge-0/0/2.0")
	assert.Equal(Network.LLDP.Vlan.Id, "0")
	assert.Equal(Network.LLDP.Vlan.Label, "VLAN")
	assert.Equal(Network.LLDP.Vlan.Name, "default")
	assert.Equal(Network.Settings.AutoNegotiation, "on")
	assert.Equal(Network.Settings.Broadcast, "yes")
	assert.Equal(Network.Settings.Driver, "e1000e")
	assert.Equal(Network.Settings.DriverVersion, "3.2.6-k")
	assert.Equal(Network.Settings.Duplex, "full")
	assert.Equal(Network.Settings.Firmware, "2.1-2")
	assert.Equal(Network.Settings.Ip, "212.32.230.67")
	assert.Equal(Network.Settings.Latency, "0")
	assert.Equal(Network.Settings.Link, "yes")
	assert.Equal(Network.Settings.Multicast, "yes")
	assert.Equal(Network.Settings.Port, "twisted pair")
	assert.Equal(Network.Settings.Speed, "1Gbit/s")
}

func TestDedicatedServerGetHardwareInformationServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetHardwareInformation(ctx, "2378237")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetHardwareInformation(ctx, "2378237")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetHardwareInformation(ctx, "2378237")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetHardwareInformation(ctx, "2378237")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetHardwareInformation(ctx, "2378237")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListIps(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"ips": [
				{
					"ddos": {
						"detectionProfile": "ADVANCED_LOW_UDP",
						"protectionType": "ADVANCED"
					},
					"floatingIp": false,
					"gateway": "12.123.123.254",
					"ip": "12.123.123.1/24",
					"mainIp": true,
					"networkType": "PUBLIC",
					"nullRouted": true,
					"reverseLookup": "domain.example.com",
					"version": 4
				},
				{
					"ddos": {
						"detectionProfile": "STANDARD_DEFAULT",
						"protectionType": "STANDARD"
					},
					"floatingIp": false,
					"gateway": "2001:db8:85a3::8a2e:370:1",
					"ip": "2001:db8:85a3::8a2e:370:7334/64",
					"mainIp": false,
					"networkType": "REMOTE_MANAGEMENT",
					"nullRouted": false,
					"reverseLookup": "domain.example.com",
					"version": 6
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListIps(ctx, "server-id", DedicatedServerListIpsOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 2)

	Ip1 := response.Ips[0]
	assert.Equal(Ip1.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip1.DDOS.ProtectionType, "ADVANCED")
	assert.Equal(Ip1.FloatingIp, false)
	assert.Equal(Ip1.Gateway, "12.123.123.254")
	assert.Equal(Ip1.Ip, "12.123.123.1/24")
	assert.Equal(Ip1.MainIp, true)
	assert.Equal(Ip1.NetworkType, "PUBLIC")
	assert.Equal(Ip1.NullRouted, true)
	assert.Equal(Ip1.ReverseLookup, "domain.example.com")
	assert.Equal(Ip1.Version.String(), "4")

	Ip2 := response.Ips[1]
	assert.Equal(Ip2.DDOS.DetectionProfile, "STANDARD_DEFAULT")
	assert.Equal(Ip2.DDOS.ProtectionType, "STANDARD")
	assert.Equal(Ip2.FloatingIp, false)
	assert.Equal(Ip2.Gateway, "2001:db8:85a3::8a2e:370:1")
	assert.Equal(Ip2.Ip, "2001:db8:85a3::8a2e:370:7334/64")
	assert.Equal(Ip2.MainIp, false)
	assert.Equal(Ip2.NetworkType, "REMOTE_MANAGEMENT")
	assert.Equal(Ip2.NullRouted, false)
	assert.Equal(Ip2.ReverseLookup, "domain.example.com")
	assert.Equal(Ip2.Version.String(), "6")
}

func TestDedicatedServerListIpsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "ips": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListIps(ctx, "server-id", DedicatedServerListIpsOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 0)
}

func TestDedicatedServerListIpsFilterAndPagination(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"ips": [
				{
					"ddos": {
						"detectionProfile": "ADVANCED_LOW_UDP",
						"protectionType": "ADVANCED"
					},
					"floatingIp": false,
					"gateway": "12.123.123.254",
					"ip": "12.123.123.1/24",
					"mainIp": true,
					"networkType": "PUBLIC",
					"nullRouted": true,
					"reverseLookup": "domain.example.com",
					"version": 4
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := DedicatedServerListIpsOptions{
		PaginationOptions: PaginationOptions{
			Offset: Int(0),
			Limit:  Int(10),
		},
	}
	response, err := DedicatedServerApi{}.ListIps(ctx, "server-id", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 1)

	Ip1 := response.Ips[0]
	assert.Equal(Ip1.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip1.DDOS.ProtectionType, "ADVANCED")
	assert.Equal(Ip1.FloatingIp, false)
	assert.Equal(Ip1.Gateway, "12.123.123.254")
	assert.Equal(Ip1.Ip, "12.123.123.1/24")
	assert.Equal(Ip1.MainIp, true)
	assert.Equal(Ip1.NetworkType, "PUBLIC")
	assert.Equal(Ip1.NullRouted, true)
	assert.Equal(Ip1.ReverseLookup, "domain.example.com")
	assert.Equal(Ip1.Version.String(), "4")
}

func TestDedicatedServerListIpsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListIps(ctx, "server-id", DedicatedServerListIpsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListIps(ctx, "server-id", DedicatedServerListIpsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListIps(ctx, "server-id", DedicatedServerListIpsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListIps(ctx, "server-id", DedicatedServerListIpsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ddos": {
				"detectionProfile": "ADVANCED_LOW_UDP",
				"protectionType": "ADVANCED"
			},
			"floatingIp": false,
			"gateway": "12.123.123.254",
			"ip": "12.123.123.1/24",
			"mainIp": true,
			"networkType": "PUBLIC",
			"nullRouted": true,
			"reverseLookup": "domain.example.com",
			"version": 4
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := DedicatedServerApi{}.GetIp(ctx, "12345", "127.0.0.6")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.NullRouted, true)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.Version.String(), "4")
}

func TestDedicatedServerGetIpServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerUpdateIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ddos": {
				"detectionProfile": "ADVANCED_DEFAULT",
				"protectionType": "ADVANCED"
			},
			"floatingIp": false,
			"gateway": "12.123.123.254",
			"ip": "12.123.123.1/24",
			"mainIp": true,
			"networkType": "PUBLIC",
			"nullRouted": true,
			"reverseLookup": "domain.example.com",
			"version": 4
		}`)
	})
	defer teardown()

	payload := map[string]string{"detectionProfile": "ADVANCED_DEFAULT", "reverseLookup": "domain.example.com"}
	ctx := context.Background()
	Ip, err := DedicatedServerApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_DEFAULT")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.NullRouted, true)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.Version.String(), "4")
}

func TestDedicatedServerUpdateIpServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"detectionProfile": "ADVANCED_DEFAULT", "reverseLookup": "domain.example.com"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"detectionProfile": "ADVANCED_DEFAULT", "reverseLookup": "domain.example.com"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"detectionProfile": "ADVANCED_DEFAULT", "reverseLookup": "domain.example.com"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"detectionProfile": "ADVANCED_DEFAULT", "reverseLookup": "domain.example.com"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"detectionProfile": "ADVANCED_DEFAULT", "reverseLookup": "domain.example.com"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerNullRouteAnIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ddos": {
				"detectionProfile": "ADVANCED_DEFAULT",
				"protectionType": "ADVANCED"
			},
			"floatingIp": false,
			"gateway": "12.123.123.254",
			"ip": "12.123.123.1/24",
			"mainIp": true,
			"networkType": "PUBLIC",
			"nullRouted": false,
			"reverseLookup": "domain.example.com",
			"version": 4
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := DedicatedServerApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_DEFAULT")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.Version.String(), "4")
}

func TestDedicatedServerNullRouteAnIpServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerRemoveNullRouteAnIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ddos": {
				"detectionProfile": "ADVANCED_DEFAULT",
				"protectionType": "ADVANCED"
			},
			"floatingIp": false,
			"gateway": "12.123.123.254",
			"ip": "12.123.123.1/24",
			"mainIp": true,
			"networkType": "PUBLIC",
			"nullRouted": false,
			"reverseLookup": "domain.example.com",
			"version": 4
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := DedicatedServerApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_DEFAULT")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.Version.String(), "4")
}

func TestDedicatedServerRemoveNullRouteAnIpServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListNullRoutes(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 1
			},
			"nullRoutes": [
				{
					"automatedUnnullingAt": "2016-08-12T07:45:33+00:00",
					"comment": "Device Null Route related to DDoS Mitigation",
					"ip": "1.1.1.1/32",
					"nullLevel": 3,
					"nulledAt": "2016-08-12T07:40:27+00:00",
					"ticketId": "282912"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NullRoutes), 1)

	NullRoute := response.NullRoutes[0]
	assert.Equal(NullRoute.AutomatedUnnullingAt, "2016-08-12T07:45:33+00:00")
	assert.Equal(NullRoute.Comment, "Device Null Route related to DDoS Mitigation")
	assert.Equal(NullRoute.Ip, "1.1.1.1/32")
	assert.Equal(NullRoute.NullLevel.String(), "3")
	assert.Equal(NullRoute.NulledAt, "2016-08-12T07:40:27+00:00")
	assert.Equal(NullRoute.TicketId, "282912")
}

func TestDedicatedServerListNullRoutesBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "nullRoutes": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NullRoutes), 0)
}

func TestDedicatedServerListNullRoutesFilterAndPagination(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"nullRoutes": [
				{
					"automatedUnnullingAt": "2016-08-12T07:45:33+00:00",
					"comment": "Device Null Route related to DDoS Mitigation",
					"ip": "1.1.1.1/32",
					"nullLevel": 3,
					"nulledAt": "2016-08-12T07:40:27+00:00",
					"ticketId": "282912"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := DedicatedServerApi{}.ListNullRoutes(ctx, "server-id", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NullRoutes), 1)

	NullRoute := response.NullRoutes[0]
	assert.Equal(NullRoute.AutomatedUnnullingAt, "2016-08-12T07:45:33+00:00")
	assert.Equal(NullRoute.Comment, "Device Null Route related to DDoS Mitigation")
	assert.Equal(NullRoute.Ip, "1.1.1.1/32")
	assert.Equal(NullRoute.NullLevel.String(), "3")
	assert.Equal(NullRoute.NulledAt, "2016-08-12T07:40:27+00:00")
	assert.Equal(NullRoute.TicketId, "282912")
}

func TestDedicatedServerListNullRoutesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListNetworkInterfaces(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 1
			},
			"networkInterfaces": [
				{
					"linkSpeed": "100Mbps",
					"operStatus": "OPEN",
					"status": "OPEN",
					"switchInterface": "33",
					"switchName": "EVO-AA11-1",
					"type": "PUBLIC"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := DedicatedServerApi{}.ListNetworkInterfaces(ctx, "server-id", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NetworkInterfaces), 1)

	NetworkInterface := response.NetworkInterfaces[0]
	assert.Equal(NetworkInterface.LinkSpeed, "100Mbps")
	assert.Equal(NetworkInterface.OperStatus, "OPEN")
	assert.Equal(NetworkInterface.Status, "OPEN")
	assert.Equal(NetworkInterface.SwitchInterface, "33")
	assert.Equal(NetworkInterface.SwitchName, "EVO-AA11-1")
	assert.Equal(NetworkInterface.Type, "PUBLIC")
}

func TestDedicatedServerListNetworkInterfacesBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "networkInterfaces": []}`)
	})
	defer teardown()
	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := DedicatedServerApi{}.ListNetworkInterfaces(ctx, "server-id", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NetworkInterfaces), 0)
}

func TestDedicatedServerListNetworkInterfacesPagination(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"networkInterfaces": [
				{
					"linkSpeed": "100Mbps",
					"operStatus": "OPEN",
					"status": "OPEN",
					"switchInterface": "33",
					"switchName": "EVO-AA11-1",
					"type": "PUBLIC"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := DedicatedServerApi{}.ListNetworkInterfaces(ctx, "server-id", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NetworkInterfaces), 1)

	NetworkInterface := response.NetworkInterfaces[0]
	assert.Equal(NetworkInterface.LinkSpeed, "100Mbps")
	assert.Equal(NetworkInterface.OperStatus, "OPEN")
	assert.Equal(NetworkInterface.Status, "OPEN")
	assert.Equal(NetworkInterface.SwitchInterface, "33")
	assert.Equal(NetworkInterface.SwitchName, "EVO-AA11-1")
	assert.Equal(NetworkInterface.Type, "PUBLIC")
}

func TestDedicatedServerListNetworkInterfacesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return DedicatedServerApi{}.ListNetworkInterfaces(ctx, "server-id", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return DedicatedServerApi{}.ListNetworkInterfaces(ctx, "server-id", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return DedicatedServerApi{}.ListNetworkInterfaces(ctx, "server-id", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return DedicatedServerApi{}.ListNetworkInterfaces(ctx, "server-id", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerCloseAllNetworkInterfaces(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.CloseAllNetworkInterfaces(ctx, "12345")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerCloseAllNetworkInterfacesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerOpenAllNetworkInterfaces(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.OpenAllNetworkInterfaces(ctx, "12345")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerOpenAllNetworkInterfacesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenAllNetworkInterfaces(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetNetworkInterface(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
				"linkSpeed": "100Mbps",
				"operStatus": "OPEN",
				"status": "OPEN",
				"switchInterface": "33",
				"switchName": "EVO-AA11-1",
				"type": "PUBLIC"
			}
		`)
	})
	defer teardown()

	ctx := context.Background()
	NetworkInterface, err := DedicatedServerApi{}.GetNetworkInterface(ctx, "server-id", "public")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(NetworkInterface.LinkSpeed, "100Mbps")
	assert.Equal(NetworkInterface.OperStatus, "OPEN")
	assert.Equal(NetworkInterface.Status, "OPEN")
	assert.Equal(NetworkInterface.SwitchInterface, "33")
	assert.Equal(NetworkInterface.SwitchName, "EVO-AA11-1")
	assert.Equal(NetworkInterface.Type, "PUBLIC")
}

func TestDedicatedServerGetNetworkInterfaceServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetNetworkInterface(ctx, "server-id", "public")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetNetworkInterface(ctx, "server-id", "public")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetNetworkInterface(ctx, "server-id", "public")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetNetworkInterface(ctx, "server-id", "public")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerCloseNetworkInterface(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.CloseNetworkInterface(ctx, "12345", "PUBLIC")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerCloseNetworkInterfaceServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CloseNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerOpenNetworkInterface(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.OpenNetworkInterface(ctx, "12345", "PUBLIC")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerOpenNetworkInterfaceServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.OpenNetworkInterface(ctx, "12345", "PUBLIC")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerDeleteServerFromPrivateNetwork(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.DeleteServerFromPrivateNetwork(ctx, "12345", "892")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerDeleteServerFromPrivateNetworkServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteServerFromPrivateNetwork(ctx, "12345", "892")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteServerFromPrivateNetwork(ctx, "12345", "892")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteServerFromPrivateNetwork(ctx, "12345", "892")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteServerFromPrivateNetwork(ctx, "12345", "892")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteServerFromPrivateNetwork(ctx, "12345", "892")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerAddServerToPrivateNetwork(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.AddServerToPrivateNetwork(ctx, "12345", "892", 1000)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerAddServerToPrivateNetworkServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.AddServerToPrivateNetwork(ctx, "12345", "892", 1000)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.AddServerToPrivateNetwork(ctx, "12345", "892", 1000)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.AddServerToPrivateNetwork(ctx, "12345", "892", 1000)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.AddServerToPrivateNetwork(ctx, "12345", "892", 1000)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.AddServerToPrivateNetwork(ctx, "12345", "892", 1000)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerDeleteDhcpReservation(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.DeleteDhcpReservation(ctx, "12345")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerDeleteDhcpReservationServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDhcpReservation(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDhcpReservation(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDhcpReservation(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDhcpReservation(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDhcpReservation(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerCreateDhcpReservation(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	payload := map[string]string{"bootfile": "http://example.com/bootme.ipxe", "hostname": "my-server"}
	ctx := context.Background()
	err := DedicatedServerApi{}.CreateDhcpReservation(ctx, "12345", payload)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerCreateDhcpReservationServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"bootfile": "http://example.com/bootme.ipxe", "hostname": "my-server"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CreateDhcpReservation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"bootfile": "http://example.com/bootme.ipxe", "hostname": "my-server"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CreateDhcpReservation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"bootfile": "http://example.com/bootme.ipxe", "hostname": "my-server"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CreateDhcpReservation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"bootfile": "http://example.com/bootme.ipxe", "hostname": "my-server"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CreateDhcpReservation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"bootfile": "http://example.com/bootme.ipxe", "hostname": "my-server"}
				ctx := context.Background()
				return nil, DedicatedServerApi{}.CreateDhcpReservation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListDhcpReservation(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 1
			},
			"leases": [
				{
					"bootfile": "http://mirror.leaseweb.com/ipxe-files/ubuntu-18.04.ipxe",
					"createdAt": "2019-10-18T17:31:01+00:00",
					"gateway": "192.168.0.254",
					"hostname": "my-server",
					"ip": "192.168.0.100",
					"lastClientRequest": {
						"relayAgent": null,
						"type": "DHCP_REQUEST",
						"userAgent": "Ubuntu 18.04 dhcpc"
					},
					"mac": "AA:BB:CC:DD:EE:FF",
					"netmask": "255.255.255.0",
					"site": "AMS-01",
					"updatedAt": "2019-11-18T19:29:01+00:00"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListDhcpReservation(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Leases), 1)

	Lease1 := response.Leases[0]
	assert.Equal(Lease1.BootFile, "http://mirror.leaseweb.com/ipxe-files/ubuntu-18.04.ipxe")
	assert.Equal(Lease1.CreatedAt, "2019-10-18T17:31:01+00:00")
	assert.Equal(Lease1.Gateway, "192.168.0.254")
	assert.Equal(Lease1.Hostname, "my-server")
	assert.Equal(Lease1.Ip, "192.168.0.100")
	assert.Empty(Lease1.LastClientRequest.RelayAgent)
	assert.Equal(Lease1.LastClientRequest.Type, "DHCP_REQUEST")
	assert.Equal(Lease1.LastClientRequest.UserAgent, "Ubuntu 18.04 dhcpc")
	assert.Equal(Lease1.Mac, "AA:BB:CC:DD:EE:FF")
	assert.Equal(Lease1.Netmask, "255.255.255.0")
	assert.Equal(Lease1.Site, "AMS-01")
	assert.Equal(Lease1.UpdatedAt, "2019-11-18T19:29:01+00:00")
}

func TestDedicatedServerListDhcpReservationBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "leases": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListDhcpReservation(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Leases), 0)
}

func TestDedicatedServerListDhcpReservationPagination(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"leases": [
				{
					"bootfile": "http://mirror.leaseweb.com/ipxe-files/ubuntu-18.04.ipxe",
					"createdAt": "2019-10-18T17:31:01+00:00",
					"gateway": "192.168.0.254",
					"hostname": "my-server",
					"ip": "192.168.0.100",
					"lastClientRequest": {
						"relayAgent": null,
						"type": "DHCP_REQUEST",
						"userAgent": "Ubuntu 18.04 dhcpc"
					},
					"mac": "AA:BB:CC:DD:EE:FF",
					"netmask": "255.255.255.0",
					"site": "AMS-01",
					"updatedAt": "2019-11-18T19:29:01+00:00"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListDhcpReservation(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Leases), 1)

	Lease1 := response.Leases[0]
	assert.Equal(Lease1.BootFile, "http://mirror.leaseweb.com/ipxe-files/ubuntu-18.04.ipxe")
	assert.Equal(Lease1.CreatedAt, "2019-10-18T17:31:01+00:00")
	assert.Equal(Lease1.Gateway, "192.168.0.254")
	assert.Equal(Lease1.Hostname, "my-server")
	assert.Equal(Lease1.Ip, "192.168.0.100")
	assert.Empty(Lease1.LastClientRequest.RelayAgent)
	assert.Equal(Lease1.LastClientRequest.Type, "DHCP_REQUEST")
	assert.Equal(Lease1.LastClientRequest.UserAgent, "Ubuntu 18.04 dhcpc")
	assert.Equal(Lease1.Mac, "AA:BB:CC:DD:EE:FF")
	assert.Equal(Lease1.Netmask, "255.255.255.0")
	assert.Equal(Lease1.Site, "AMS-01")
	assert.Equal(Lease1.UpdatedAt, "2019-11-18T19:29:01+00:00")
}

func TestDedicatedServerListDhcpReservationServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDhcpReservation(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDhcpReservation(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDhcpReservation(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDhcpReservation(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerCancelActiveJob(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"createdAt": "2021-01-09T08:54:06+0000",
			"flow": "#stop",
			"isRunning": false,
			"node": "80:18:44:E0:AF:C4!JGNTQ92",
			"payload": {
				"configurable": true,
				"device": "SATA2TB",
				"fileserverBaseUrl": "http://asd.asd",
				"jobType": "install",
				"numberOfDisks": null,
				"operatingSystemId": "UBUNTU_20_04_64BIT",
				"os": {
					"architecture": "64bit",
					"family": "ubuntu",
					"name": "Ubuntu 20.04 LTS (Focal Fossa) (amd64)",
					"type": "linux",
					"version": "20.04"
				},
				"partitions": [
					{
						"filesystem": "swap",
						"size": "4096"
					}
				],
				"pop": "AMS-01",
				"powerCycle": true,
				"raidLevel": null,
				"serverId": "99944",
				"timezone": "UTC",
				"x": 1
			},
			"progress": {
				"canceled": 1,
				"expired": 0,
				"failed": 0,
				"finished": 0,
				"inprogress": 0,
				"pending": 0,
				"percentage": 0,
				"total": 1,
				"waiting": 0
			},
			"serverId": "99944",
			"status": "CANCELED",
			"tasks": [
				{
					"description": "dummy",
					"errorMessage": "The job was canceled by the api consumer",
					"flow": "tasks",
					"onError": "break",
					"status": "CANCELED",
					"statusTimestamps": {
						"CANCELED": "2021-01-09T08:54:15+00:00",
						"PENDING": "2021-01-09T08:54:06+00:00",
						"WAITING": "2021-01-09T08:54:06+00:00"
					},
					"uuid": "085ce145-39bd-4cb3-8e2b-53f17a97a463"
				}
			],
			"type": "install",
			"updatedAt": "2021-01-09T08:54:15+0000",
			"uuid": "c77d8a6b-d255-4744-8b95-8bf4af6f8b48"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Job, err := DedicatedServerApi{}.CancelActiveJob(ctx, "12345")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Job.CreatedAt, "2021-01-09T08:54:06+0000")
	assert.Equal(Job.Flow, "#stop")
	assert.Equal(Job.IsRunning, false)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.ServerId, "99944")
	assert.Equal(Job.Status, "CANCELED")
	assert.Equal(Job.Type, "install")
	assert.Equal(Job.UpdatedAt, "2021-01-09T08:54:15+0000")
	assert.Equal(Job.Uuid, "c77d8a6b-d255-4744-8b95-8bf4af6f8b48")

	assert.Equal(len(Job.Tasks), 1)
	Task := Job.Tasks[0]
	assert.Equal(Task.Description, "dummy")
	assert.Equal(Task.Message, "The job was canceled by the api consumer")
	assert.Equal(Task.Flow, "tasks")
	assert.Equal(Task.OnError, "break")
	assert.Equal(Task.Status, "CANCELED")
	assert.Equal(Task.Uuid, "085ce145-39bd-4cb3-8e2b-53f17a97a463")
	assert.Equal(Task.StatusTimestamps.Canceled, "2021-01-09T08:54:15+00:00")
	assert.Equal(Task.StatusTimestamps.Pending, "2021-01-09T08:54:06+00:00")
	assert.Equal(Task.StatusTimestamps.Waiting, "2021-01-09T08:54:06+00:00")
	assert.Equal(Job.Progress.Canceled.String(), "1")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "0")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "1")
	assert.Equal(Job.Progress.Waiting.String(), "0")
	assert.Equal(Job.Payload["configurable"], true)
	assert.Equal(Job.Payload["device"], "SATA2TB")
	assert.Equal(Job.Payload["fileserverBaseUrl"], "http://asd.asd")
	assert.Equal(Job.Payload["jobType"], "install")
	assert.Empty(Job.Payload["numberOfDisks"])
	assert.Equal(Job.Payload["operatingSystemId"], "UBUNTU_20_04_64BIT")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["architecture"], "64bit")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["family"], "ubuntu")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["name"], "Ubuntu 20.04 LTS (Focal Fossa) (amd64)")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["type"], "linux")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["version"], "20.04")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["filesystem"], "swap")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["size"], "4096")
	assert.Equal(Job.Payload["pop"], "AMS-01")
	assert.Equal(Job.Payload["powerCycle"], true)
	assert.Empty(Job.Payload["raidLevel"])
	assert.Equal(Job.Payload["serverId"], "99944")
	assert.Equal(Job.Payload["timezone"], "UTC")
	assert.Equal(Job.Payload["x"].(float64), float64(1))
}

func TestDedicatedServerCancelActiveJobServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CancelActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CancelActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CancelActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CancelActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CancelActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerExpireActiveJob(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"createdAt": "2021-01-09T08:54:06+0000",
			"flow": "#stop",
			"isRunning": false,
			"node": "80:18:44:E0:AF:C4!JGNTQ92",
			"payload": {
				"configurable": true,
				"device": "SATA2TB",
				"fileserverBaseUrl": "http://asd.asd",
				"jobType": "install",
				"numberOfDisks": null,
				"operatingSystemId": "UBUNTU_20_04_64BIT",
				"os": {
					"architecture": "64bit",
					"family": "ubuntu",
					"name": "Ubuntu 20.04 LTS (Focal Fossa) (amd64)",
					"type": "linux",
					"version": "20.04"
				},
				"partitions": [
					{
						"filesystem": "swap",
						"size": "4096"
					}
				],
				"pop": "AMS-01",
				"powerCycle": true,
				"raidLevel": null,
				"serverId": "99944",
				"timezone": "UTC",
				"x": 1
			},
			"progress": {
				"canceled": 1,
				"expired": 0,
				"failed": 0,
				"finished": 0,
				"inprogress": 0,
				"pending": 0,
				"percentage": 0,
				"total": 1,
				"waiting": 0
			},
			"serverId": "99944",
			"status": "CANCELED",
			"tasks": [
				{
					"description": "dummy",
					"errorMessage": "The job was canceled by the api consumer",
					"flow": "tasks",
					"onError": "break",
					"status": "CANCELED",
					"statusTimestamps": {
						"CANCELED": "2021-01-09T08:54:15+00:00",
						"PENDING": "2021-01-09T08:54:06+00:00",
						"WAITING": "2021-01-09T08:54:06+00:00"
					},
					"uuid": "085ce145-39bd-4cb3-8e2b-53f17a97a463"
				}
			],
			"type": "install",
			"updatedAt": "2021-01-09T08:54:15+0000",
			"uuid": "c77d8a6b-d255-4744-8b95-8bf4af6f8b48"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Job, err := DedicatedServerApi{}.ExpireActiveJob(ctx, "12345")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Job.CreatedAt, "2021-01-09T08:54:06+0000")
	assert.Equal(Job.Flow, "#stop")
	assert.Equal(Job.IsRunning, false)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.ServerId, "99944")
	assert.Equal(Job.Status, "CANCELED")
	assert.Equal(Job.Type, "install")
	assert.Equal(Job.UpdatedAt, "2021-01-09T08:54:15+0000")
	assert.Equal(Job.Uuid, "c77d8a6b-d255-4744-8b95-8bf4af6f8b48")

	assert.Equal(len(Job.Tasks), 1)
	Task := Job.Tasks[0]
	assert.Equal(Task.Description, "dummy")
	assert.Equal(Task.Message, "The job was canceled by the api consumer")
	assert.Equal(Task.Flow, "tasks")
	assert.Equal(Task.OnError, "break")
	assert.Equal(Task.Status, "CANCELED")
	assert.Equal(Task.Uuid, "085ce145-39bd-4cb3-8e2b-53f17a97a463")
	assert.Equal(Task.StatusTimestamps.Canceled, "2021-01-09T08:54:15+00:00")
	assert.Equal(Task.StatusTimestamps.Pending, "2021-01-09T08:54:06+00:00")
	assert.Equal(Task.StatusTimestamps.Waiting, "2021-01-09T08:54:06+00:00")
	assert.Equal(Job.Progress.Canceled.String(), "1")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "0")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "1")
	assert.Equal(Job.Progress.Waiting.String(), "0")
	assert.Equal(Job.Payload["configurable"], true)
	assert.Equal(Job.Payload["device"], "SATA2TB")
	assert.Equal(Job.Payload["fileserverBaseUrl"], "http://asd.asd")
	assert.Equal(Job.Payload["jobType"], "install")
	assert.Empty(Job.Payload["numberOfDisks"])
	assert.Equal(Job.Payload["operatingSystemId"], "UBUNTU_20_04_64BIT")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["architecture"], "64bit")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["family"], "ubuntu")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["name"], "Ubuntu 20.04 LTS (Focal Fossa) (amd64)")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["type"], "linux")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["version"], "20.04")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["filesystem"], "swap")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["size"], "4096")
	assert.Equal(Job.Payload["pop"], "AMS-01")
	assert.Equal(Job.Payload["powerCycle"], true)
	assert.Empty(Job.Payload["raidLevel"])
	assert.Equal(Job.Payload["serverId"], "99944")
	assert.Equal(Job.Payload["timezone"], "UTC")
	assert.Equal(Job.Payload["x"].(float64), float64(1))
}

func TestDedicatedServerExpireActiveJobServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ExpireActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ExpireActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ExpireActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ExpireActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ExpireActiveJob(ctx, "12345")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerLaunchHardwareScan(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"createdAt": "2021-01-09T08:54:06+0000",
			"flow": "tasks",
			"isRunning": true,
			"node": "80:18:44:E0:AF:C4!JGNTQ92",
			"payload": {
				"fileserverBaseUrl": "http://asd.asd",
				"jobType": "hardwareScan",
				"pop": "AMS-01",
				"powerCycle": true,
				"serverId": "99944"
			},
			"progress": {
				"canceled": 0,
				"expired": 0,
				"failed": 0,
				"finished": 0,
				"inprogress": 0,
				"pending": 1,
				"percentage": 0,
				"total": 5,
				"waiting": 4
			},
			"serverId": "99944",
			"status": "ACTIVE",
			"tasks": [
				{
					"description": "dummy",
					"errorMessage": null,
					"flow": "tasks",
					"onError": "break",
					"status": "PENDING",
					"statusTimestamps": {
						"PENDING": "2021-01-09T08:54:06+00:00",
						"WAITING": "2021-01-09T08:54:06+00:00"
					},
					"uuid": "085ce145-39bd-4cb3-8e2b-53f17a97a463"
				}
			],
			"type": "hardwareScan",
			"updatedAt": "2021-01-09T08:54:15+0000",
			"uuid": "c77d8a6b-d255-4744-8b95-8bf4af6f8b48"
		}`)
	})
	defer teardown()

	payload := map[string]interface{}{"callbackUrl": "http://callback.url", "powerCycle": true}
	ctx := context.Background()
	Job, err := DedicatedServerApi{}.LaunchHardwareScan(ctx, "12345", payload)
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Job.CreatedAt, "2021-01-09T08:54:06+0000")
	assert.Equal(Job.Flow, "tasks")
	assert.Equal(Job.IsRunning, true)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.ServerId, "99944")
	assert.Equal(Job.Status, "ACTIVE")
	assert.Equal(Job.Type, "hardwareScan")
	assert.Equal(Job.UpdatedAt, "2021-01-09T08:54:15+0000")
	assert.Equal(Job.Uuid, "c77d8a6b-d255-4744-8b95-8bf4af6f8b48")

	assert.Equal(len(Job.Tasks), 1)
	Task := Job.Tasks[0]
	assert.Equal(Task.Description, "dummy")
	assert.Empty(Task.Message)
	assert.Equal(Task.Flow, "tasks")
	assert.Equal(Task.OnError, "break")
	assert.Equal(Task.Status, "PENDING")
	assert.Equal(Task.Uuid, "085ce145-39bd-4cb3-8e2b-53f17a97a463")
	assert.Equal(Task.StatusTimestamps.Pending, "2021-01-09T08:54:06+00:00")
	assert.Equal(Task.StatusTimestamps.Waiting, "2021-01-09T08:54:06+00:00")
	assert.Equal(Job.Progress.Canceled.String(), "0")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "1")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "5")
	assert.Equal(Job.Progress.Waiting.String(), "4")
	assert.Equal(Job.Payload["fileserverBaseUrl"], "http://asd.asd")
	assert.Equal(Job.Payload["jobType"], "hardwareScan")
	assert.Equal(Job.Payload["pop"], "AMS-01")
	assert.Equal(Job.Payload["powerCycle"], true)
	assert.Equal(Job.Payload["serverId"], "99944")
}

func TestDedicatedServerLaunchHardwareScanServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "http://callback.url", "powerCycle": true}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchHardwareScan(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "http://callback.url", "powerCycle": true}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchHardwareScan(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "http://callback.url", "powerCycle": true}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchHardwareScan(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "http://callback.url", "powerCycle": true}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchHardwareScan(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "http://callback.url", "powerCycle": true}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchHardwareScan(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerLaunchInstallation(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"createdAt": "2021-03-06T21:55:32+0000",
			"flow": "tasks",
			"isRunning": true,
			"node": "AA:BB:CC:DD:EE:FF!DKFJKD8989",
			"payload": {
				"configurable": true,
				"device": "SATA2TB",
				"fileserverBaseUrl": "10.11.20.2",
				"isUnassignedServer": false,
				"jobType": "install",
				"numberOfDisks": null,
				"operatingSystemId": "UBUNTU_20_04_64BIT",
				"os": {
					"architecture": "64bit",
					"family": "ubuntu",
					"name": "Ubuntu 20.04 LTS (Focal Fossa) (amd64)",
					"type": "linux",
					"version": "20.04"
				},
				"partitions": [
					{
						"bootable": true,
						"filesystem": "ext2",
						"mountpoint": "/boot",
						"primary": true,
						"size": "500"
					},
					{
						"filesystem": "swap",
						"size": "4096"
					},
					{
						"filesystem": "ext4",
						"mountpoint": "/tmp",
						"size": "4096"
					},
					{
						"filesystem": "ext4",
						"mountpoint": "/",
						"primary": true,
						"size": "*"
					}
				],
				"pop": "AMS-01",
				"powerCycle": true,
				"raidLevel": null,
				"serverId": "12345",
				"timezone": "UTC"
			},
			"progress": {
				"canceled": 0,
				"expired": 0,
				"failed": 0,
				"finished": 0,
				"inprogress": 0,
				"pending": 1,
				"percentage": 0,
				"total": 26,
				"waiting": 25
			},
			"serverId": "12345",
			"status": "ACTIVE",
			"tasks": [
				{
					"description": "dummy",
					"errorMessage": null,
					"flow": "tasks",
					"onError": "break",
					"status": "PENDING",
					"statusTimestamps": {
						"PENDING": "2021-01-09T09:18:06+00:00",
						"WAITING": "2021-01-09T09:18:06+00:00"
					},
					"uuid": "c1a56a5a-ae38-4b12-acb6-0cba9ceb1016"
				}
			],
			"type": "install",
			"updatedAt": "2021-03-06T21:55:32+0000",
			"uuid": "bcf2bedf-8450-4b22-86a8-f30aeb3a38f9"
		}`)
	})
	defer teardown()

	payload := map[string]interface{}{
		"controlPanelId":    "PLESK_12",
		"device":            "SATA2TB",
		"hostname":          "ubuntu20.local",
		"operatingSystemId": "UBUNTU_20_04_64BIT",
		"partitions": []map[string]interface{}{
			{
				"bootable":   true,
				"filesystem": "ext2",
				"mountpoint": "/boot",
				"primary":    true,
				"size":       "512",
			},
			{
				"filesystem": "swap",
				"size":       "4096",
			},
			{
				"filesystem": "ext4",
				"mountpoint": "/tmp",
				"size":       "2048",
			},
			{
				"filesystem": "ext4",
				"mountpoint": "/",
				"primary":    true,
				"size":       "*",
			},
		},
		"sshKeys": "ssh-rsa AAAAB3NzaC1y... user@domain.com",
	}
	ctx := context.Background()
	Job, err := DedicatedServerApi{}.LaunchInstallation(ctx, "12345", payload)
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Job.CreatedAt, "2021-03-06T21:55:32+0000")
	assert.Equal(Job.Flow, "tasks")
	assert.Equal(Job.IsRunning, true)
	assert.Equal(Job.Node, "AA:BB:CC:DD:EE:FF!DKFJKD8989")
	assert.Equal(Job.ServerId, "12345")
	assert.Equal(Job.Status, "ACTIVE")
	assert.Equal(Job.Type, "install")
	assert.Equal(Job.UpdatedAt, "2021-03-06T21:55:32+0000")
	assert.Equal(Job.Uuid, "bcf2bedf-8450-4b22-86a8-f30aeb3a38f9")

	assert.Equal(len(Job.Tasks), 1)
	Task := Job.Tasks[0]

	assert.Equal(Task.Description, "dummy")
	assert.Empty(Task.Message)
	assert.Equal(Task.Flow, "tasks")
	assert.Equal(Task.OnError, "break")
	assert.Equal(Task.Status, "PENDING")
	assert.Equal(Task.Uuid, "c1a56a5a-ae38-4b12-acb6-0cba9ceb1016")
	assert.Equal(Task.StatusTimestamps.Pending, "2021-01-09T09:18:06+00:00")
	assert.Equal(Task.StatusTimestamps.Waiting, "2021-01-09T09:18:06+00:00")
	assert.Equal(Job.Progress.Canceled.String(), "0")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "1")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "26")
	assert.Equal(Job.Progress.Waiting.String(), "25")
	assert.Equal(Job.Payload["fileserverBaseUrl"], "10.11.20.2")
	assert.Equal(Job.Payload["configurable"], true)
	assert.Equal(Job.Payload["pop"], "AMS-01")
	assert.Equal(Job.Payload["powerCycle"], true)
	assert.Equal(Job.Payload["serverId"], "12345")
	assert.Equal(Job.Payload["timezone"], "UTC")
	assert.Empty(Job.Payload["raidLevel"])
	assert.Equal(Job.Payload["device"], "SATA2TB")
	assert.Equal(Job.Payload["isUnassignedServer"], false)
	assert.Equal(Job.Payload["jobType"], "install")
	assert.Empty(Job.Payload["numberOfDisks"])
	assert.Equal(Job.Payload["operatingSystemId"], "UBUNTU_20_04_64BIT")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["architecture"], "64bit")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["family"], "ubuntu")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["name"], "Ubuntu 20.04 LTS (Focal Fossa) (amd64)")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["type"], "linux")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["version"], "20.04")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["filesystem"], "ext2")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["size"], "500")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["bootable"], true)
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["mountpoint"], "/boot")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["primary"], true)
	assert.Equal(Job.Payload["partitions"].([]interface{})[1].(map[string]interface{})["filesystem"], "swap")
	assert.Equal(Job.Payload["partitions"].([]interface{})[1].(map[string]interface{})["size"], "4096")
	assert.Equal(Job.Payload["partitions"].([]interface{})[2].(map[string]interface{})["filesystem"], "ext4")
	assert.Equal(Job.Payload["partitions"].([]interface{})[2].(map[string]interface{})["mountpoint"], "/tmp")
	assert.Equal(Job.Payload["partitions"].([]interface{})[2].(map[string]interface{})["size"], "4096")
	assert.Equal(Job.Payload["partitions"].([]interface{})[3].(map[string]interface{})["filesystem"], "ext4")
	assert.Equal(Job.Payload["partitions"].([]interface{})[3].(map[string]interface{})["size"], "*")
	assert.Equal(Job.Payload["partitions"].([]interface{})[3].(map[string]interface{})["mountpoint"], "/")
	assert.Equal(Job.Payload["partitions"].([]interface{})[3].(map[string]interface{})["primary"], true)
}

func TestDedicatedServerLaunchInstallationServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"controlPanelId":    "PLESK_12",
					"device":            "SATA2TB",
					"hostname":          "ubuntu20.local",
					"operatingSystemId": "UBUNTU_20_04_64BIT",
					"partitions": []map[string]interface{}{
						{
							"bootable":   true,
							"filesystem": "ext2",
							"mountpoint": "/boot",
							"primary":    true,
							"size":       "512",
						},
						{
							"filesystem": "swap",
							"size":       "4096",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/tmp",
							"size":       "2048",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/",
							"primary":    true,
							"size":       "*",
						},
					},
					"sshKeys": "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchInstallation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"controlPanelId":    "PLESK_12",
					"device":            "SATA2TB",
					"hostname":          "ubuntu20.local",
					"operatingSystemId": "UBUNTU_20_04_64BIT",
					"partitions": []map[string]interface{}{
						{
							"bootable":   true,
							"filesystem": "ext2",
							"mountpoint": "/boot",
							"primary":    true,
							"size":       "512",
						},
						{
							"filesystem": "swap",
							"size":       "4096",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/tmp",
							"size":       "2048",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/",
							"primary":    true,
							"size":       "*",
						},
					},
					"sshKeys": "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchInstallation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"controlPanelId":    "PLESK_12",
					"device":            "SATA2TB",
					"hostname":          "ubuntu20.local",
					"operatingSystemId": "UBUNTU_20_04_64BIT",
					"partitions": []map[string]interface{}{
						{
							"bootable":   true,
							"filesystem": "ext2",
							"mountpoint": "/boot",
							"primary":    true,
							"size":       "512",
						},
						{
							"filesystem": "swap",
							"size":       "4096",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/tmp",
							"size":       "2048",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/",
							"primary":    true,
							"size":       "*",
						},
					},
					"sshKeys": "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchInstallation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"controlPanelId":    "PLESK_12",
					"device":            "SATA2TB",
					"hostname":          "ubuntu20.local",
					"operatingSystemId": "UBUNTU_20_04_64BIT",
					"partitions": []map[string]interface{}{
						{
							"bootable":   true,
							"filesystem": "ext2",
							"mountpoint": "/boot",
							"primary":    true,
							"size":       "512",
						},
						{
							"filesystem": "swap",
							"size":       "4096",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/tmp",
							"size":       "2048",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/",
							"primary":    true,
							"size":       "*",
						},
					},
					"sshKeys": "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchInstallation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"controlPanelId":    "PLESK_12",
					"device":            "SATA2TB",
					"hostname":          "ubuntu20.local",
					"operatingSystemId": "UBUNTU_20_04_64BIT",
					"partitions": []map[string]interface{}{
						{
							"bootable":   true,
							"filesystem": "ext2",
							"mountpoint": "/boot",
							"primary":    true,
							"size":       "512",
						},
						{
							"filesystem": "swap",
							"size":       "4096",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/tmp",
							"size":       "2048",
						},
						{
							"filesystem": "ext4",
							"mountpoint": "/",
							"primary":    true,
							"size":       "*",
						},
					},
					"sshKeys": "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchInstallation(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerLaunchIpmiRest(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"createdAt": "2018-01-09T09:18:06+0000",
			"flow": "tasks",
			"isRunning": true,
			"node": "80:18:44:E0:AF:C4!JGNTQ92",
			"payload": {
				"fileserverBaseUrl": "",
				"hasPublicIpmiIp": false,
				"jobType": "ipmiReset",
				"pop": "AMS-01",
				"powerCycle": true,
				"serverId": "99944"
			},
			"progress": {
				"canceled": 0,
				"expired": 0,
				"failed": 0,
				"finished": 0,
				"inprogress": 0,
				"pending": 1,
				"percentage": 0,
				"total": 1,
				"waiting": 0
			},
			"serverId": "99944",
			"status": "ACTIVE",
			"tasks": [
				{
					"description": "dummy",
					"errorMessage": null,
					"flow": "tasks",
					"onError": "break",
					"status": "PENDING",
					"statusTimestamps": {
						"PENDING": "2018-01-09T09:18:06+00:00",
						"WAITING": "2018-01-09T09:18:06+00:00"
					},
					"uuid": "c1a56a5a-ae38-4b12-acb6-0cba9ceb1016"
				}
			],
			"type": "ipmiReset",
			"updatedAt": "2018-01-09T09:18:06+0000",
			"uuid": "754154c2-cc7f-4d5f-b8bf-b654084ba4a9"
		}`)
	})
	defer teardown()

	payload := map[string]interface{}{"callbackUrl": "https://callbacks.example.org"}
	ctx := context.Background()
	Job, err := DedicatedServerApi{}.LaunchIpmiRest(ctx, "12345", payload)
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Job.CreatedAt, "2018-01-09T09:18:06+0000")
	assert.Equal(Job.Flow, "tasks")
	assert.Equal(Job.IsRunning, true)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.ServerId, "99944")
	assert.Equal(Job.Status, "ACTIVE")
	assert.Equal(Job.Type, "ipmiReset")
	assert.Equal(Job.UpdatedAt, "2018-01-09T09:18:06+0000")
	assert.Equal(Job.Uuid, "754154c2-cc7f-4d5f-b8bf-b654084ba4a9")

	assert.Equal(len(Job.Tasks), 1)
	Task := Job.Tasks[0]

	assert.Equal(Task.Description, "dummy")
	assert.Empty(Task.Message)
	assert.Equal(Task.Flow, "tasks")
	assert.Equal(Task.OnError, "break")
	assert.Equal(Task.Status, "PENDING")
	assert.Equal(Task.Uuid, "c1a56a5a-ae38-4b12-acb6-0cba9ceb1016")
	assert.Equal(Task.StatusTimestamps.Pending, "2018-01-09T09:18:06+00:00")
	assert.Equal(Task.StatusTimestamps.Waiting, "2018-01-09T09:18:06+00:00")
	assert.Equal(Job.Progress.Canceled.String(), "0")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "1")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "1")
	assert.Equal(Job.Progress.Waiting.String(), "0")

	assert.Equal(Job.Payload["fileserverBaseUrl"], "")
	assert.Equal(Job.Payload["pop"], "AMS-01")
	assert.Equal(Job.Payload["powerCycle"], true)
	assert.Equal(Job.Payload["serverId"], "99944")
	assert.Equal(Job.Payload["jobType"], "ipmiReset")
	assert.Equal(Job.Payload["hasPublicIpmiIp"], false)
}

func TestDedicatedServerLaunchIpmiRestServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "https://callbacks.example.org"}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchIpmiRest(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "https://callbacks.example.org"}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchIpmiRest(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "https://callbacks.example.org"}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchIpmiRest(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "https://callbacks.example.org"}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchIpmiRest(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{"callbackUrl": "https://callbacks.example.org"}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchIpmiRest(ctx, "12345", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListJobs(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 1
			},
			"jobs": [
				{
					"createdAt": "2018-01-09T10:38:12+0000",
					"flow": "tasks",
					"isRunning": true,
					"node": "80:18:44:E0:AF:C4!JGNTQ92",
					"progress": {
						"canceled": 0,
						"expired": 0,
						"failed": 0,
						"finished": 0,
						"inprogress": 0,
						"pending": 1,
						"percentage": 0,
						"total": 1,
						"waiting": 0
					},
					"serverId": "99944",
					"status": "ACTIVE",
					"type": "install",
					"updatedAt": "2018-01-09T10:38:12+0000",
					"uuid": "3a867358-5b4b-44ee-88ac-4274603ef641"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListJobs(ctx, "99944", DedicatedServerListJobOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Jobs), 1)

	Job := response.Jobs[0]
	assert.Equal(Job.Progress.Canceled.String(), "0")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "1")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "1")
	assert.Equal(Job.Progress.Waiting.String(), "0")
	assert.Equal(Job.Flow, "tasks")
	assert.Equal(Job.IsRunning, true)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.CreatedAt, "2018-01-09T10:38:12+0000")
	assert.Equal(Job.ServerId, "99944")
	assert.Equal(Job.Status, "ACTIVE")
	assert.Equal(Job.Type, "install")
	assert.Equal(Job.UpdatedAt, "2018-01-09T10:38:12+0000")
	assert.Equal(Job.Uuid, "3a867358-5b4b-44ee-88ac-4274603ef641")
}

func TestDedicatedServerListJobsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "jobs": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListJobs(ctx, "99944", DedicatedServerListJobOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Jobs), 0)
}

func TestDedicatedServerListJobsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"jobs": [
				{
					"createdAt": "2018-01-09T10:38:12+0000",
					"flow": "tasks",
					"isRunning": true,
					"node": "80:18:44:E0:AF:C4!JGNTQ92",
					"progress": {
						"canceled": 0,
						"expired": 0,
						"failed": 0,
						"finished": 0,
						"inprogress": 0,
						"pending": 1,
						"percentage": 0,
						"total": 1,
						"waiting": 0
					},
					"serverId": "99944",
					"status": "ACTIVE",
					"type": "install",
					"updatedAt": "2018-01-09T10:38:12+0000",
					"uuid": "3a867358-5b4b-44ee-88ac-4274603ef641"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := DedicatedServerListJobOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(10),
		},
	}
	response, err := DedicatedServerApi{}.ListJobs(ctx, "99944", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Jobs), 1)

	Job := response.Jobs[0]
	assert.Equal(Job.Progress.Canceled.String(), "0")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "1")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "1")
	assert.Equal(Job.Progress.Waiting.String(), "0")
	assert.Equal(Job.Flow, "tasks")
	assert.Equal(Job.IsRunning, true)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.CreatedAt, "2018-01-09T10:38:12+0000")
	assert.Equal(Job.ServerId, "99944")
	assert.Equal(Job.Status, "ACTIVE")
	assert.Equal(Job.Type, "install")
	assert.Equal(Job.UpdatedAt, "2018-01-09T10:38:12+0000")
	assert.Equal(Job.Uuid, "3a867358-5b4b-44ee-88ac-4274603ef641")
}

func TestDedicatedServerListJobsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListJobs(ctx, "99944", DedicatedServerListJobOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListJobs(ctx, "99944", DedicatedServerListJobOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListJobs(ctx, "99944", DedicatedServerListJobOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListJobs(ctx, "99944", DedicatedServerListJobOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetJob(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"createdAt": "2021-01-09T10:38:12+0000",
			"flow": "tasks",
			"isRunning": true,
			"metadata": {
				"BATCH_ID": "biannual-os-upgrade-installs"
			},
			"node": "80:18:44:E0:AF:C4!JGNTQ92",
			"payload": {
				"configurable": true,
				"device": "SATA2TB",
				"fileserverBaseUrl": "",
				"jobType": "install",
				"numberOfDisks": null,
				"operatingSystemId": "UBUNTU_20_04_64BIT",
				"os": {
					"architecture": "64bit",
					"family": "ubuntu",
					"name": "Ubuntu 20.04 LTS (Focal Fossa) (amd64)",
					"type": "linux",
					"version": "20.04"
				},
				"partitions": [
					{
						"filesystem": "swap",
						"size": "4096"
					}
				],
				"pop": "AMS-01",
				"powerCycle": true,
				"raidLevel": null,
				"serverId": "99944",
				"timezone": "UTC",
				"x": 1
			},
			"progress": {
				"canceled": 0,
				"expired": 0,
				"failed": 0,
				"finished": 0,
				"inprogress": 0,
				"pending": 1,
				"percentage": 0,
				"total": 1,
				"waiting": 0
			},
			"serverId": "99944",
			"status": "ACTIVE",
			"tasks": [
				{
					"description": "dummy",
					"errorMessage": null,
					"flow": "tasks",
					"onError": "break",
					"status": "PENDING",
					"statusTimestamps": {
						"PENDING": "2021-01-09T10:38:12+00:00",
						"WAITING": "2021-01-09T10:38:12+00:00"
					},
					"uuid": "8a10b74b-2a94-4a3b-88da-b9c07faa240d"
				}
			],
			"type": "install",
			"updatedAt": "2021-01-09T10:38:12+0000",
			"uuid": "3a867358-5b4b-44ee-88ac-4274603ef641"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Job, err := DedicatedServerApi{}.GetJob(ctx, "99944", "job id")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Job.Progress.Canceled.String(), "0")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "1")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "1")
	assert.Equal(Job.Progress.Waiting.String(), "0")
	assert.Equal(Job.Flow, "tasks")
	assert.Equal(Job.IsRunning, true)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.CreatedAt, "2021-01-09T10:38:12+0000")
	assert.Equal(Job.ServerId, "99944")
	assert.Equal(Job.Status, "ACTIVE")
	assert.Equal(Job.Type, "install")
	assert.Equal(Job.UpdatedAt, "2021-01-09T10:38:12+0000")
	assert.Equal(Job.Uuid, "3a867358-5b4b-44ee-88ac-4274603ef641")
	assert.Equal(Job.Metadata.BatchId, "biannual-os-upgrade-installs")
	assert.Equal(len(Job.Tasks), 1)
	assert.Equal(Job.Tasks[0].Description, "dummy")
	assert.Empty(Job.Tasks[0].Message)
	assert.Equal(Job.Tasks[0].Flow, "tasks")
	assert.Equal(Job.Tasks[0].OnError, "break")
	assert.Equal(Job.Tasks[0].Status, "PENDING")
	assert.Equal(Job.Tasks[0].StatusTimestamps.Pending, "2021-01-09T10:38:12+00:00")
	assert.Equal(Job.Tasks[0].StatusTimestamps.Waiting, "2021-01-09T10:38:12+00:00")
	assert.Equal(Job.Tasks[0].Uuid, "8a10b74b-2a94-4a3b-88da-b9c07faa240d")

	assert.Equal(Job.Payload["configurable"], true)
	assert.Equal(Job.Payload["device"], "SATA2TB")
	assert.Equal(Job.Payload["fileserverBaseUrl"], "")
	assert.Equal(Job.Payload["jobType"], "install")
	assert.Empty(Job.Payload["numberOfDisks"])
	assert.Equal(Job.Payload["operatingSystemId"], "UBUNTU_20_04_64BIT")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["architecture"], "64bit")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["family"], "ubuntu")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["name"], "Ubuntu 20.04 LTS (Focal Fossa) (amd64)")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["type"], "linux")
	assert.Equal(Job.Payload["os"].(map[string]interface{})["version"], "20.04")
	assert.Equal(Job.Payload["pop"], "AMS-01")
	assert.Equal(Job.Payload["powerCycle"], true)
	assert.Empty(Job.Payload["raidLevel"])
	assert.Equal(Job.Payload["serverId"], "99944")
	assert.Equal(Job.Payload["timezone"], "UTC")
	assert.Equal(Job.Payload["x"].(float64), float64(1))
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["filesystem"], "swap")
	assert.Equal(Job.Payload["partitions"].([]interface{})[0].(map[string]interface{})["size"], "4096")
}

func TestDedicatedServerGetJobServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetJob(ctx, "99944", "job id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetJob(ctx, "99944", "job id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetJob(ctx, "99944", "job id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetJob(ctx, "99944", "job id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetJob(ctx, "99944", "job id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerLaunchRescueMode(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"createdAt": "2018-01-09T09:18:06+0000",
			"flow": "tasks",
			"isRunning": true,
			"node": "80:18:44:E0:AF:C4!JGNTQ92",
			"payload": {
				"fileserverBaseUrl": "10.11.20.2",
				"isUnassignedServer": false,
				"jobType": "rescueMode",
				"pop": "AMS-01",
				"powerCycle": true,
				"serverId": "2349839"
			},
			"progress": {
				"canceled": 0,
				"expired": 0,
				"failed": 0,
				"finished": 0,
				"inprogress": 0,
				"pending": 1,
				"percentage": 0,
				"total": 12,
				"waiting": 11
			},
			"serverId": "2349839",
			"status": "ACTIVE",
			"tasks": [
				{
					"description": "dummy",
					"errorMessage": null,
					"flow": "tasks",
					"onError": "break",
					"status": "PENDING",
					"statusTimestamps": {
						"PENDING": "2018-01-09T09:18:06+00:00",
						"WAITING": "2018-01-09T09:18:06+00:00"
					},
					"uuid": "c1a56a5a-ae38-4b12-acb6-0cba9ceb1016"
				}
			],
			"type": "rescueMode",
			"updatedAt": "2018-01-09T09:18:06+0000",
			"uuid": "754154c2-cc7f-4d5f-b8bf-b654084ba4a9"
		}`)
	})
	defer teardown()

	payload := map[string]interface{}{
		"callbackUrl":   "https://example.com/urlExecutedOnCallback",
		"powerCycle":    true,
		"rescueImageId": "GRML",
		"sshKeys":       "ssh-rsa AAAAB3NzaC1y... user@domain.com",
	}
	ctx := context.Background()
	Job, err := DedicatedServerApi{}.LaunchRescueMode(ctx, "2349839", payload)
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Job.CreatedAt, "2018-01-09T09:18:06+0000")
	assert.Equal(Job.Flow, "tasks")
	assert.Equal(Job.IsRunning, true)
	assert.Equal(Job.Node, "80:18:44:E0:AF:C4!JGNTQ92")
	assert.Equal(Job.ServerId, "2349839")
	assert.Equal(Job.Status, "ACTIVE")
	assert.Equal(Job.Type, "rescueMode")
	assert.Equal(Job.UpdatedAt, "2018-01-09T09:18:06+0000")
	assert.Equal(Job.Uuid, "754154c2-cc7f-4d5f-b8bf-b654084ba4a9")

	assert.Equal(len(Job.Tasks), 1)
	Task := Job.Tasks[0]

	assert.Equal(Task.Description, "dummy")
	assert.Empty(Task.Message)
	assert.Equal(Task.Flow, "tasks")
	assert.Equal(Task.OnError, "break")
	assert.Equal(Task.Status, "PENDING")
	assert.Equal(Task.Uuid, "c1a56a5a-ae38-4b12-acb6-0cba9ceb1016")
	assert.Equal(Task.StatusTimestamps.Pending, "2018-01-09T09:18:06+00:00")
	assert.Equal(Task.StatusTimestamps.Waiting, "2018-01-09T09:18:06+00:00")
	assert.Equal(Job.Progress.Canceled.String(), "0")
	assert.Equal(Job.Progress.Expired.String(), "0")
	assert.Equal(Job.Progress.Failed.String(), "0")
	assert.Equal(Job.Progress.Finished.String(), "0")
	assert.Equal(Job.Progress.InProgress.String(), "0")
	assert.Equal(Job.Progress.Pending.String(), "1")
	assert.Equal(Job.Progress.Percentage.String(), "0")
	assert.Equal(Job.Progress.Total.String(), "12")
	assert.Equal(Job.Progress.Waiting.String(), "11")

	assert.Equal(Job.Payload["fileserverBaseUrl"], "10.11.20.2")
	assert.Equal(Job.Payload["pop"], "AMS-01")
	assert.Equal(Job.Payload["powerCycle"], true)
	assert.Equal(Job.Payload["serverId"], "2349839")
	assert.Equal(Job.Payload["jobType"], "rescueMode")
	assert.Equal(Job.Payload["isUnassignedServer"], false)
}

func TestDedicatedServerLaunchRescueModeServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"callbackUrl":   "https://example.com/urlExecutedOnCallback",
					"powerCycle":    true,
					"rescueImageId": "GRML",
					"sshKeys":       "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchRescueMode(ctx, "2349839", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"callbackUrl":   "https://example.com/urlExecutedOnCallback",
					"powerCycle":    true,
					"rescueImageId": "GRML",
					"sshKeys":       "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchRescueMode(ctx, "2349839", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"callbackUrl":   "https://example.com/urlExecutedOnCallback",
					"powerCycle":    true,
					"rescueImageId": "GRML",
					"sshKeys":       "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchRescueMode(ctx, "2349839", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"callbackUrl":   "https://example.com/urlExecutedOnCallback",
					"powerCycle":    true,
					"rescueImageId": "GRML",
					"sshKeys":       "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchRescueMode(ctx, "2349839", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]interface{}{
					"callbackUrl":   "https://example.com/urlExecutedOnCallback",
					"powerCycle":    true,
					"rescueImageId": "GRML",
					"sshKeys":       "ssh-rsa AAAAB3NzaC1y... user@domain.com",
				}
				ctx := context.Background()
				return DedicatedServerApi{}.LaunchRescueMode(ctx, "2349839", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 4
			},
			"credentials": [
				{
					"type": "REMOTE_MANAGEMENT",
					"username": "admin"
				},
				{
					"type": "REMOTE_MANAGEMENT",
					"username": "root"
				},
				{
					"type": "OPERATING_SYSTEM",
					"username": "root"
				},
				{
					"type": "OPERATING_SYSTEM",
					"username": "user"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 4)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 4)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
	assert.Equal(response.Credentials[1].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[1].Username, "root")
	assert.Equal(response.Credentials[2].Type, "OPERATING_SYSTEM")
	assert.Equal(response.Credentials[2].Username, "root")
	assert.Equal(response.Credentials[3].Type, "OPERATING_SYSTEM")
	assert.Equal(response.Credentials[3].Username, "user")
}

func TestDedicatedServerListCredentialsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestDedicatedServerListCredentialsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"credentials": [
				{
					"type": "REMOTE_MANAGEMENT",
					"username": "admin"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
}

func TestDedicatedServerListCredentialsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerCreateCredential(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"password": "mys3cr3tp@ssw0rd",
			"type": "OPERATING_SYSTEM",
			"username": "root"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Type, "OPERATING_SYSTEM")
	assert.Equal(resp.Username, "root")
	assert.Equal(resp.Password, "mys3cr3tp@ssw0rd")
}

func TestDedicatedServerCreateCredentialServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 400",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"430495b4-c052-451a-a016-30d8f72ac59b","errorCode":"400","errorMessage":"Validation Failed","errorDetails":{"username":["Usernames can only contain alphanumeric values and @.-_ characters"],"password":["This field is missing."]}}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "ro=ot", "")
			},
			ExpectedError: ApiError{
				CorrelationId: "430495b4-c052-451a-a016-30d8f72ac59b",
				Code:          "400",
				Message:       "Validation Failed",
				Details: map[string][]string{
					"username": []string{"Usernames can only contain alphanumeric values and @.-_ characters"},
					"password": []string{"This field is missing."},
				},
			},
		},
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListCredentialsByType(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"credentials": [
				{
					"type": "REMOTE_MANAGEMENT",
					"username": "admin"
				},
				{
					"type": "REMOTE_MANAGEMENT",
					"username": "root"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 2)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
	assert.Equal(response.Credentials[1].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[1].Username, "root")
}

func TestDedicatedServerListCredentialsByTypeBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestDedicatedServerListCredentialsByTypePaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"credentials": [
				{
					"type": "REMOTE_MANAGEMENT",
					"username": "admin"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
}

func TestDedicatedServerListCredentialsByTypeServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"password": "mys3cr3tp@ssw0rd",
			"type": "OPERATING_SYSTEM",
			"username": "root"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Credential, err := DedicatedServerApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Credential.Password, "mys3cr3tp@ssw0rd")
	assert.Equal(Credential.Username, "root")
	assert.Equal(Credential.Type, "OPERATING_SYSTEM")
}

func TestDedicatedServerGetCredentialsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerDeleteCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerDeleteCredentialsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerUpdateCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"password": "new password",
			"type": "OPERATING_SYSTEM",
			"username": "admin"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Password, "new password")
	assert.Equal(resp.Type, "OPERATING_SYSTEM")
	assert.Equal(resp.Username, "admin")
}

func TestDedicatedServerUpdateCredentialsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetBandWidthMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"aggregation": "AVG",
				"from": "2016-10-20T09:00:00Z",
				"granularity": "HOUR",
				"to": "2016-10-20T11:00:00Z"
			},
			"metrics": {
				"DOWN_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2016-10-20T09:00:00Z",
							"value": 202499
						},
						{
							"timestamp": "2016-10-20T10:00:00Z",
							"value": 29900
						}
					]
				},
				"UP_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2016-10-20T09:00:00Z",
							"value": 43212393
						},
						{
							"timestamp": "2016-10-20T10:00:00Z",
							"value": 12342929
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := MetricsOptions{
		From:        String("2016-10-20T09:00:00Z"),
		To:          String("2016-10-20T11:00:00Z"),
		Granularity: String("HOUR"),
		Aggregation: String("AVG"),
	}
	Metric, err := DedicatedServerApi{}.GetBandWidthMetrics(ctx, "99944", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Metric.Metadata.Aggregation, "AVG")
	assert.Equal(Metric.Metadata.From, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metadata.To, "2016-10-20T11:00:00Z")
	assert.Equal(Metric.Metadata.Granularity, "HOUR")
	assert.Equal(Metric.Metric.DownPublic.Unit, "bps")
	assert.Equal(Metric.Metric.DownPublic.Values[0].Value.String(), "202499")
	assert.Equal(Metric.Metric.DownPublic.Values[0].Timestamp, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metric.DownPublic.Values[1].Value.String(), "29900")
	assert.Equal(Metric.Metric.DownPublic.Values[1].Timestamp, "2016-10-20T10:00:00Z")

	assert.Equal(Metric.Metric.UpPublic.Unit, "bps")
	assert.Equal(Metric.Metric.UpPublic.Values[0].Value.String(), "43212393")
	assert.Equal(Metric.Metric.UpPublic.Values[0].Timestamp, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metric.UpPublic.Values[1].Value.String(), "12342929")
	assert.Equal(Metric.Metric.UpPublic.Values[1].Timestamp, "2016-10-20T10:00:00Z")
}

func TestDedicatedServerGetBandWidthMetricsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("AVG"),
				}
				return DedicatedServerApi{}.GetBandWidthMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("AVG"),
				}
				return DedicatedServerApi{}.GetBandWidthMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("AVG"),
				}
				return DedicatedServerApi{}.GetBandWidthMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("AVG"),
				}
				return DedicatedServerApi{}.GetBandWidthMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("AVG"),
				}
				return DedicatedServerApi{}.GetBandWidthMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetDataTrafficMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"aggregation": "SUM",
				"from": "2016-10-20T09:00:00Z",
				"granularity": "HOUR",
				"to": "2016-10-20T11:00:00Z"
			},
			"metrics": {
				"DOWN_PUBLIC": {
					"unit": "B",
					"values": [
						{
							"timestamp": "2016-10-20T09:00:00Z",
							"value": 202499
						},
						{
							"timestamp": "2016-10-20T10:00:00Z",
							"value": 29900
						}
					]
				},
				"UP_PUBLIC": {
					"unit": "B",
					"values": [
						{
							"timestamp": "2016-10-20T09:00:00Z",
							"value": 43212393
						},
						{
							"timestamp": "2016-10-20T10:00:00Z",
							"value": 12342929
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := MetricsOptions{
		From:        String("2016-10-20T09:00:00Z"),
		To:          String("2016-10-20T11:00:00Z"),
		Granularity: String("HOUR"),
		Aggregation: String("SUM"),
	}
	Metric, err := DedicatedServerApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Metric.Metadata.Aggregation, "SUM")
	assert.Equal(Metric.Metadata.From, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metadata.To, "2016-10-20T11:00:00Z")
	assert.Equal(Metric.Metadata.Granularity, "HOUR")
	assert.Equal(Metric.Metric.DownPublic.Unit, "B")
	assert.Equal(Metric.Metric.DownPublic.Values[0].Value.String(), "202499")
	assert.Equal(Metric.Metric.DownPublic.Values[0].Timestamp, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metric.DownPublic.Values[1].Value.String(), "29900")
	assert.Equal(Metric.Metric.DownPublic.Values[1].Timestamp, "2016-10-20T10:00:00Z")

	assert.Equal(Metric.Metric.UpPublic.Unit, "B")
	assert.Equal(Metric.Metric.UpPublic.Values[0].Value.String(), "43212393")
	assert.Equal(Metric.Metric.UpPublic.Values[0].Timestamp, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metric.UpPublic.Values[1].Value.String(), "12342929")
	assert.Equal(Metric.Metric.UpPublic.Values[1].Timestamp, "2016-10-20T10:00:00Z")
}

func TestDedicatedServerGetDataTrafficMetricsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("SUM"),
				}
				return DedicatedServerApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("SUM"),
				}
				return DedicatedServerApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("SUM"),
				}
				return DedicatedServerApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("SUM"),
				}
				return DedicatedServerApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := MetricsOptions{
					From:        String("2016-10-20T09:00:00Z"),
					To:          String("2016-10-20T11:00:00Z"),
					Granularity: String("HOUR"),
					Aggregation: String("SUM"),
				}
				return DedicatedServerApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListBandWidthNotificationSettings(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"bandwidthNotificationSettings": [
				{
					"actions": [
						{
							"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
							"type": "EMAIL"
						}
					],
					"frequency": "WEEKLY",
					"id": "12345",
					"lastCheckedAt": "2021-03-16T01:01:41+00:00",
					"threshold": "1",
					"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
					"unit": "Gbps"
				},
				{
					"actions": [
						{
							"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
							"type": "EMAIL"
						}
					],
					"frequency": "DAILY",
					"id": "123456",
					"lastCheckedAt": "2021-03-16T01:01:41+00:00",
					"threshold": "1",
					"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
					"unit": "Mbps"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Metadata.TotalCount, 2)
	assert.Equal(resp.Metadata.Offset, 0)
	assert.Equal(resp.Metadata.Limit, 10)
	assert.Equal(len(resp.Settings), 2)

	assert.Equal(resp.Settings[0].Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Settings[0].Actions[0].Type, "EMAIL")
	assert.Equal(resp.Settings[0].Frequency, "WEEKLY")
	assert.Equal(resp.Settings[0].Id, "12345")
	assert.Equal(resp.Settings[0].LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Settings[0].Threshold.String(), "1")
	assert.Equal(resp.Settings[0].ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Settings[0].Unit, "Gbps")

	assert.Equal(resp.Settings[1].Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Settings[1].Actions[0].Type, "EMAIL")
	assert.Equal(resp.Settings[1].Frequency, "DAILY")
	assert.Equal(resp.Settings[1].Id, "123456")
	assert.Equal(resp.Settings[1].LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Settings[1].Threshold.String(), "1")
	assert.Equal(resp.Settings[1].ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Settings[1].Unit, "Mbps")
}

func TestDedicatedServerListBandWidthNotificationSettingsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"bandwidthNotificationSettings": [
				{
					"actions": [
						{
							"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
							"type": "EMAIL"
						}
					],
					"frequency": "WEEKLY",
					"id": "12345",
					"lastCheckedAt": "2021-03-16T01:01:41+00:00",
					"threshold": "1",
					"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
					"unit": "Gbps"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	resp, err := DedicatedServerApi{}.ListBandWidthNotificationSettings(ctx, "server-id", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Metadata.TotalCount, 11)
	assert.Equal(resp.Metadata.Offset, 1)
	assert.Equal(resp.Metadata.Limit, 10)
	assert.Equal(len(resp.Settings), 1)

	assert.Equal(resp.Settings[0].Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Settings[0].Actions[0].Type, "EMAIL")
	assert.Equal(resp.Settings[0].Frequency, "WEEKLY")
	assert.Equal(resp.Settings[0].Id, "12345")
	assert.Equal(resp.Settings[0].LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Settings[0].Threshold.String(), "1")
	assert.Equal(resp.Settings[0].ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Settings[0].Unit, "Gbps")
}

func TestDedicatedServerListBandWidthNotificationSettingsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerCreateBandWidthNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"actions": [
				{
					"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
					"type": "EMAIL"
				}
			],
			"frequency": "WEEKLY",
			"id": "12345",
			"lastCheckedAt": "2021-03-16T01:01:41+00:00",
			"threshold": "1",
			"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
			"unit": "Gbps"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", 1, "Gbps")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Actions[0].Type, "EMAIL")
	assert.Equal(resp.Frequency, "WEEKLY")
	assert.Equal(resp.Id, "12345")
	assert.Equal(resp.LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Threshold.String(), "1")
	assert.Equal(resp.ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Unit, "Gbps")
}

func TestDedicatedServerCreateBandWidthNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", 1, "Gbps")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", 1, "Gbps")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", 1, "Gbps")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", 1, "Gbps")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", 1, "Gbps")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerDeleteBandWidthNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerDeleteBandWidthNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetBandWidthNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"actions": [
				{
					"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
					"type": "EMAIL"
				}
			],
			"frequency": "WEEKLY",
			"id": "12345",
			"lastCheckedAt": "2021-03-16T01:01:41+00:00",
			"threshold": "1",
			"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
			"unit": "Gbps"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Actions[0].Type, "EMAIL")
	assert.Equal(resp.Frequency, "WEEKLY")
	assert.Equal(resp.Id, "12345")
	assert.Equal(resp.LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Threshold.String(), "1")
	assert.Equal(resp.ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Unit, "Gbps")
}

func TestDedicatedServerGetBandWidthNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerUpdateBandWidthNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"actions": [
				{
					"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
					"type": "EMAIL"
				}
			],
			"frequency": "MONTHLY",
			"id": "12345",
			"lastCheckedAt": "2021-03-16T01:01:41+00:00",
			"threshold": "2",
			"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
			"unit": "Mbps"
		}`)
	})
	defer teardown()

	payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "Mbps"}
	ctx := context.Background()
	resp, err := DedicatedServerApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Actions[0].Type, "EMAIL")
	assert.Equal(resp.Frequency, "MONTHLY")
	assert.Equal(resp.Id, "12345")
	assert.Equal(resp.LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Threshold.String(), "2")
	assert.Equal(resp.ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Unit, "Mbps")
}

func TestDedicatedServerUpdateBandWidthNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "Mbps"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "Mbps"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "Mbps"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "Mbps"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "Mbps"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListDataTrafficNotificationSettings(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"datatrafficNotificationSettings": [
				{
					"actions": [
						{
							"lastTriggeredAt": null,
							"type": "EMAIL"
						}
					],
					"frequency": "WEEKLY",
					"id": "12345",
					"lastCheckedAt": null,
					"threshold": "1",
					"thresholdExceededAt": null,
					"unit": "MB"
				},
				{
					"actions": [
						{
							"lastTriggeredAt": null,
							"type": "EMAIL"
						}
					],
					"frequency": "DAILY",
					"id": "123456",
					"lastCheckedAt": null,
					"threshold": "1",
					"thresholdExceededAt": null,
					"unit": "GB"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Metadata.TotalCount, 2)
	assert.Equal(resp.Metadata.Offset, 0)
	assert.Equal(resp.Metadata.Limit, 10)
	assert.Equal(len(resp.Settings), 2)

	assert.Empty(resp.Settings[0].Actions[0].LastTriggeredAt)
	assert.Equal(resp.Settings[0].Actions[0].Type, "EMAIL")
	assert.Equal(resp.Settings[0].Frequency, "WEEKLY")
	assert.Equal(resp.Settings[0].Id, "12345")
	assert.Empty(resp.Settings[0].LastCheckedAt)
	assert.Equal(resp.Settings[0].Threshold.String(), "1")
	assert.Empty(resp.Settings[0].ThresholdExceededAt)
	assert.Equal(resp.Settings[0].Unit, "MB")

	assert.Empty(resp.Settings[1].Actions[0].LastTriggeredAt)
	assert.Equal(resp.Settings[1].Actions[0].Type, "EMAIL")
	assert.Equal(resp.Settings[1].Frequency, "DAILY")
	assert.Equal(resp.Settings[1].Id, "123456")
	assert.Empty(resp.Settings[1].LastCheckedAt)
	assert.Equal(resp.Settings[1].Threshold.String(), "1")
	assert.Empty(resp.Settings[1].ThresholdExceededAt)
	assert.Equal(resp.Settings[1].Unit, "GB")
}

func TestDedicatedServerListDataTrafficNotificationSettingsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"datatrafficNotificationSettings": [
				{
					"actions": [
						{
							"lastTriggeredAt": null,
							"type": "EMAIL"
						}
					],
					"frequency": "WEEKLY",
					"id": "12345",
					"lastCheckedAt": null,
					"threshold": "1",
					"thresholdExceededAt": null,
					"unit": "MB"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	resp, err := DedicatedServerApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Metadata.TotalCount, 11)
	assert.Equal(resp.Metadata.Offset, 1)
	assert.Equal(resp.Metadata.Limit, 10)
	assert.Equal(len(resp.Settings), 1)

	assert.Empty(resp.Settings[0].Actions[0].LastTriggeredAt)
	assert.Equal(resp.Settings[0].Actions[0].Type, "EMAIL")
	assert.Equal(resp.Settings[0].Frequency, "WEEKLY")
	assert.Equal(resp.Settings[0].Id, "12345")
	assert.Empty(resp.Settings[0].LastCheckedAt)
	assert.Equal(resp.Settings[0].Threshold.String(), "1")
	assert.Empty(resp.Settings[0].ThresholdExceededAt)
	assert.Equal(resp.Settings[0].Unit, "MB")
}

func TestDedicatedServerListDataTrafficNotificationSettingsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerCreateDataTrafficNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"actions": [
				{
					"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
					"type": "EMAIL"
				}
			],
			"frequency": "WEEKLY",
			"id": "12345",
			"lastCheckedAt": "2021-03-16T01:01:41+00:00",
			"threshold": "1",
			"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
			"unit": "GB"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", 1, "GB")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Actions[0].Type, "EMAIL")
	assert.Equal(resp.Frequency, "WEEKLY")
	assert.Equal(resp.Id, "12345")
	assert.Equal(resp.LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Threshold.String(), "1")
	assert.Equal(resp.ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Unit, "GB")
}

func TestDedicatedServerCreateDataTrafficNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", 1, "GB")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", 1, "GB")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", 1, "GB")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", 1, "GB")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", 1, "GB")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerDeleteDataTrafficNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerDeleteDataTrafficNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetDataTrafficNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"actions": [
				{
					"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
					"type": "EMAIL"
				}
			],
			"frequency": "WEEKLY",
			"id": "12345",
			"lastCheckedAt": "2021-03-16T01:01:41+00:00",
			"threshold": "1",
			"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
			"unit": "GB"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Actions[0].Type, "EMAIL")
	assert.Equal(resp.Frequency, "WEEKLY")
	assert.Equal(resp.Id, "12345")
	assert.Equal(resp.LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Threshold.String(), "1")
	assert.Equal(resp.ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Unit, "GB")
}

func TestDedicatedServerGetDataTrafficNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerUpdateDataTrafficNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"actions": [
				{
					"lastTriggeredAt": "2021-03-16T01:01:44+00:00",
					"type": "EMAIL"
				}
			],
			"frequency": "MONTHLY",
			"id": "12345",
			"lastCheckedAt": "2021-03-16T01:01:41+00:00",
			"threshold": "2",
			"thresholdExceededAt": "2021-03-16T01:01:41+00:00",
			"unit": "GB"
		}`)
	})
	defer teardown()

	payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "GB"}
	ctx := context.Background()
	resp, err := DedicatedServerApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Actions[0].LastTriggeredAt, "2021-03-16T01:01:44+00:00")
	assert.Equal(resp.Actions[0].Type, "EMAIL")
	assert.Equal(resp.Frequency, "MONTHLY")
	assert.Equal(resp.Id, "12345")
	assert.Equal(resp.LastCheckedAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Threshold.String(), "2")
	assert.Equal(resp.ThresholdExceededAt, "2021-03-16T01:01:41+00:00")
	assert.Equal(resp.Unit, "GB")
}

func TestDedicatedServerUpdateDataTrafficNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "GB"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "GB"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "GB"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "GB"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{"frequency": "MONTHLY", "threshold": "2", "unit": "GB"}
				ctx := context.Background()
				return DedicatedServerApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetDdosNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"nulling": "ENABLED",
			"scrubbing": "DISABLED"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.GetDdosNotificationSetting(ctx, "server-id")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Nulling, "ENABLED")
	assert.Equal(resp.Scrubbing, "DISABLED")
}

func TestDedicatedServerGetDdosNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDdosNotificationSetting(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDdosNotificationSetting(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDdosNotificationSetting(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDdosNotificationSetting(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetDdosNotificationSetting(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerUpdateDdosNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerUpdateDdosNotificationSettingServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerPowerCycleServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.PowerCycleServer(ctx, "server id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerPowerCycleServerServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerCycleServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerCycleServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerCycleServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerCycleServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerCycleServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerPowerOffServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.PowerOffServer(ctx, "server id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerPowerOffServerServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOffServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOffServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOffServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOffServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOffServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerPowerOnServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedServerApi{}.PowerOnServer(ctx, "server id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedServerPowerOnServerServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOnServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOnServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOnServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOnServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedServerApi{}.PowerOnServer(ctx, "server id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetPowerStatus(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ipmi": {
				"status": "off"
			},
			"pdu": {
				"status": "on"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := DedicatedServerApi{}.GetPowerStatus(ctx, "server-id")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Ipmi.Status, "off")
	assert.Equal(resp.Pdu.Status, "on")
}

func TestDedicatedServerGetPowerStatusServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetPowerStatus(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetPowerStatus(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetPowerStatus(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetPowerStatus(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetPowerStatus(ctx, "server-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListOperatingSystems(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"operatingSystems": [
				{
					"id": "ALMALINUX_8_64BIT",
					"name": "AlmaLinux 8 (x86_64)"
				},
				{
					"id": "CENTOS_7_64BIT",
					"name": "CentOS 7 (x86_64)"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListOperatingSystems(ctx, DedicatedServerListOperatingSystemsOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.OperatingSystems), 2)

	assert.Equal(response.OperatingSystems[0].Id, "ALMALINUX_8_64BIT")
	assert.Equal(response.OperatingSystems[0].Name, "AlmaLinux 8 (x86_64)")
	assert.Equal(response.OperatingSystems[1].Id, "CENTOS_7_64BIT")
	assert.Equal(response.OperatingSystems[1].Name, "CentOS 7 (x86_64)")
}

func TestDedicatedServerListOperatingSystemsPagination(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"operatingSystems": [
				{
					"id": "ALMALINUX_8_64BIT",
					"name": "AlmaLinux 8 (x86_64)"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListOperatingSystems(ctx, DedicatedServerListOperatingSystemsOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.OperatingSystems), 1)

	assert.Equal(response.OperatingSystems[0].Id, "ALMALINUX_8_64BIT")
	assert.Equal(response.OperatingSystems[0].Name, "AlmaLinux 8 (x86_64)")
}

func TestDedicatedServerListOperatingSystemsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListOperatingSystems(ctx, DedicatedServerListOperatingSystemsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListOperatingSystems(ctx, DedicatedServerListOperatingSystemsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListOperatingSystems(ctx, DedicatedServerListOperatingSystemsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListOperatingSystems(ctx, DedicatedServerListOperatingSystemsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerGetOperatingSystem(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"architecture": "64bit",
			"configurable": true,
			"defaults": {
				"device": "SATA_SAS",
				"partitions": [
					{
						"bootable": true,
						"filesystem": "ext2",
						"mountpoint": "/boot",
						"primary": true,
						"size": "1024"
					},
					{
						"filesystem": "swap",
						"size": "4096"
					},
					{
						"filesystem": "ext4",
						"mountpoint": "/tmp",
						"size": "4096"
					},
					{
						"filesystem": "ext4",
						"mountpoint": "/",
						"primary": true,
						"size": "*"
					}
				]
			},
			"family": "ubuntu",
			"features": [
				"PARTITIONING",
				"SW_RAID",
				"TIMEZONE",
				"HOSTNAME",
				"SSH_KEYS",
				"POST_INSTALL_SCRIPTS"
			],
			"id": "UBUNTU_20_04_64BIT",
			"name": "Ubuntu 20.04 LTS (Focal Fossa) (amd64)",
			"supportedBootDevices": [
				"SATA_SAS",
				"NVME"
			],
			"supportedFileSystems": [
				"ext2",
				"ext3",
				"ext4",
				"xfs",
				"swap"
			],
			"type": "linux",
			"version": "20.04"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	OperatingSystem, err := DedicatedServerApi{}.GetOperatingSystem(ctx, "UBUNTU_20_04_64BIT", "controll panel id")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(OperatingSystem.Architecture, "64bit")
	assert.Equal(OperatingSystem.Configurable, true)
	assert.Equal(OperatingSystem.Defaults.Device, "SATA_SAS")

	assert.Equal(OperatingSystem.Defaults.Partitions[0].Bootable, true)
	assert.Equal(OperatingSystem.Defaults.Partitions[0].Filesystem, "ext2")
	assert.Equal(OperatingSystem.Defaults.Partitions[0].Mountpoint, "/boot")
	assert.Equal(OperatingSystem.Defaults.Partitions[0].Primary, true)
	assert.Equal(OperatingSystem.Defaults.Partitions[0].Size, "1024")
	assert.Equal(OperatingSystem.Defaults.Partitions[1].Filesystem, "swap")
	assert.Equal(OperatingSystem.Defaults.Partitions[1].Size, "4096")
	assert.Equal(OperatingSystem.Defaults.Partitions[2].Filesystem, "ext4")
	assert.Equal(OperatingSystem.Defaults.Partitions[2].Mountpoint, "/tmp")
	assert.Equal(OperatingSystem.Defaults.Partitions[2].Size, "4096")
	assert.Equal(OperatingSystem.Defaults.Partitions[3].Filesystem, "ext4")
	assert.Equal(OperatingSystem.Defaults.Partitions[3].Mountpoint, "/")
	assert.Equal(OperatingSystem.Defaults.Partitions[3].Primary, true)
	assert.Equal(OperatingSystem.Defaults.Partitions[3].Size, "*")
	assert.Equal(OperatingSystem.Family, "ubuntu")
	assert.Equal(OperatingSystem.Features[0], "PARTITIONING")
	assert.Equal(OperatingSystem.Features[1], "SW_RAID")
	assert.Equal(OperatingSystem.Features[2], "TIMEZONE")
	assert.Equal(OperatingSystem.Features[3], "HOSTNAME")
	assert.Equal(OperatingSystem.Features[4], "SSH_KEYS")
	assert.Equal(OperatingSystem.Features[5], "POST_INSTALL_SCRIPTS")
	assert.Equal(OperatingSystem.Id, "UBUNTU_20_04_64BIT")
	assert.Equal(OperatingSystem.Name, "Ubuntu 20.04 LTS (Focal Fossa) (amd64)")
	assert.Equal(OperatingSystem.SupportedBootDevices[0], "SATA_SAS")
	assert.Equal(OperatingSystem.SupportedBootDevices[1], "NVME")
	assert.Equal(OperatingSystem.SupportedFileSystems[0], "ext2")
	assert.Equal(OperatingSystem.SupportedFileSystems[1], "ext3")
	assert.Equal(OperatingSystem.SupportedFileSystems[2], "ext4")
	assert.Equal(OperatingSystem.SupportedFileSystems[3], "xfs")
	assert.Equal(OperatingSystem.SupportedFileSystems[4], "swap")
	assert.Equal(OperatingSystem.Type, "linux")
	assert.Equal(OperatingSystem.Version, "20.04")
}

func TestDedicatedServerGetOperatingSystemServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetOperatingSystem(ctx, "UBUNTU_20_04_64BIT", "controll panel id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetOperatingSystem(ctx, "UBUNTU_20_04_64BIT", "controll panel id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetOperatingSystem(ctx, "UBUNTU_20_04_64BIT", "controll panel id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.GetOperatingSystem(ctx, "UBUNTU_20_04_64BIT", "controll panel id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListControlPanels(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"controlPanels": [
				{
					"id": "CPANEL_PREMIER_100",
					"name": "cPanel Premier 100"
				},
				{
					"id": "CPANEL_PREMIER_150",
					"name": "cPanel Premier 150"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListControlPanels(ctx, DedicatedServerListControlPanelsOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.ControlPanels), 2)

	assert.Equal(response.ControlPanels[0].Id, "CPANEL_PREMIER_100")
	assert.Equal(response.ControlPanels[0].Name, "cPanel Premier 100")
	assert.Equal(response.ControlPanels[1].Id, "CPANEL_PREMIER_150")
	assert.Equal(response.ControlPanels[1].Name, "cPanel Premier 150")
}

func TestDedicatedServerListControlPanelsPagination(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"controlPanels": [
				{
					"id": "CPANEL_PREMIER_100",
					"name": "cPanel Premier 100"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := DedicatedServerListControlPanelsOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(10),
		},
	}
	response, err := DedicatedServerApi{}.ListControlPanels(ctx, opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.ControlPanels), 1)

	assert.Equal(response.ControlPanels[0].Id, "CPANEL_PREMIER_100")
	assert.Equal(response.ControlPanels[0].Name, "cPanel Premier 100")
}

func TestDedicatedServerListControlPanelsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListControlPanels(ctx, DedicatedServerListControlPanelsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListControlPanels(ctx, DedicatedServerListControlPanelsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListControlPanels(ctx, DedicatedServerListControlPanelsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListControlPanels(ctx, DedicatedServerListControlPanelsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestDedicatedServerListRescueImages(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"rescueImages": [
				{
					"id": "GRML",
					"name": "GRML Linux Rescue Image (amd64)"
				},
				{
					"id": "CENTOS_7",
					"name": "CentOS 7 Linux Rescue Image (amd64)"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedServerApi{}.ListRescueImages(ctx, PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.RescueImages), 2)

	assert.Equal(response.RescueImages[0].Id, "GRML")
	assert.Equal(response.RescueImages[0].Name, "GRML Linux Rescue Image (amd64)")
	assert.Equal(response.RescueImages[1].Id, "CENTOS_7")
	assert.Equal(response.RescueImages[1].Name, "CentOS 7 Linux Rescue Image (amd64)")
}

func TestDedicatedServerListRescueImagesPagination(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"rescueImages": [
				{
					"id": "GRML",
					"name": "GRML Linux Rescue Image (amd64)"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := DedicatedServerApi{}.ListRescueImages(ctx, opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.RescueImages), 1)

	assert.Equal(response.RescueImages[0].Id, "GRML")
	assert.Equal(response.RescueImages[0].Name, "GRML Linux Rescue Image (amd64)")
}

func TestDedicatedServerListRescueImagesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListRescueImages(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListRescueImages(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListRescueImages(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedServerApi{}.ListRescueImages(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
