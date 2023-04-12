package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDedicatedNetworkEquipmentList(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"networkEquipments": [
				{
					"contract": {
						"customerId": "10085996",
						"deliveryStatus": "ACTIVE",
						"id": "49031391001170",
						"reference": "My Private Switch",
						"salesOrgId": "2000"
					},
					"featureAvailability": {
						"automation": true,
						"ipmiReboot": false,
						"powerCycle": true,
						"privateNetwork": false,
						"remoteManagement": false
					},
					"id": "12345",
					"location": {
						"rack": "YY11",
						"site": "AMS-01",
						"suite": "HALL3",
						"unit": "21"
					},
					"networkInterfaces": {
						"internal": {
							"gateway": null,
							"ip": null,
							"ports": []
						},
						"public": {
							"gateway": "127.0.2.254",
							"ip": "127.0.2.1/24",
							"locationId": "",
							"nullRouted": false,
							"ports": []
						},
						"remoteManagement": {
							"gateway": null,
							"ip": null,
							"locationId": null,
							"ports": []
						}
					},
					"type": "SWITCH"
				},
				{
					"contract": {
						"customerId": "10085996",
						"deliveryStatus": "ACTIVE",
						"id": "49031513001110",
						"reference": "My Other Private Switch",
						"salesOrgId": "2000"
					},
					"featureAvailability": {
						"automation": false,
						"ipmiReboot": false,
						"powerCycle": false,
						"privateNetwork": false,
						"remoteManagement": false
					},
					"id": "45678",
					"location": {
						"rack": "XX00",
						"site": "AMS-01",
						"suite": "HALL3",
						"unit": "21"
					},
					"networkInterfaces": {
						"internal": {
							"gateway": null,
							"ip": null,
							"ports": []
						},
						"public": {
							"gateway": "127.1.1.254",
							"ip": "127.1.1.68/24",
							"locationId": "",
							"nullRouted": false,
							"ports": []
						},
						"remoteManagement": {
							"gateway": null,
							"ip": null,
							"locationId": null,
							"ports": []
						}
					},
					"type": "SWITCH"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedNetworkEquipmentApi{}.List(ctx, DedicatedNetworkEquipmentListOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NetworkEquipments), 2)

	NetworkEquipment1 := response.NetworkEquipments[0]
	assert.Equal(NetworkEquipment1.Id, "12345")
	assert.Equal(NetworkEquipment1.Contract.CustomerId, "10085996")
	assert.Equal(NetworkEquipment1.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(NetworkEquipment1.Contract.Id, "49031391001170")
	assert.Equal(NetworkEquipment1.Contract.Reference, "My Private Switch")
	assert.Equal(NetworkEquipment1.Contract.SalesOrgId, "2000")
	assert.Equal(NetworkEquipment1.FeatureAvailability.Automation, true)
	assert.Equal(NetworkEquipment1.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(NetworkEquipment1.FeatureAvailability.IpmiReboot, false)
	assert.Equal(NetworkEquipment1.FeatureAvailability.PowerCycle, true)
	assert.Equal(NetworkEquipment1.FeatureAvailability.RemoteManagement, false)
	assert.Equal(NetworkEquipment1.Location.Rack, "YY11")
	assert.Equal(NetworkEquipment1.Location.Site, "AMS-01")
	assert.Equal(NetworkEquipment1.Location.Suite, "HALL3")
	assert.Equal(NetworkEquipment1.Location.Unit, "21")
	assert.Empty(NetworkEquipment1.NetworkInterfaces.Internal.Gateway)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.Internal.Ip)
	assert.Equal(len(NetworkEquipment1.NetworkInterfaces.Internal.Ports), 0)
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.Gateway, "127.0.2.254")
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.Ip, "127.0.2.1/24")
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.LocationId, "")
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.NullRouted, false)
	assert.Equal(len(NetworkEquipment1.NetworkInterfaces.Public.Ports), 0)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.RemoteManagement.Gateway)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.RemoteManagement.Ip)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.RemoteManagement.LocationId)
	assert.Equal(len(NetworkEquipment1.NetworkInterfaces.RemoteManagement.Ports), 0)
	assert.Equal(NetworkEquipment1.Type, "SWITCH")

	NetworkEquipment2 := response.NetworkEquipments[1]
	assert.Equal(NetworkEquipment2.Id, "45678")
	assert.Equal(NetworkEquipment2.Contract.CustomerId, "10085996")
	assert.Equal(NetworkEquipment2.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(NetworkEquipment2.Contract.Id, "49031513001110")
	assert.Equal(NetworkEquipment2.Contract.Reference, "My Other Private Switch")
	assert.Equal(NetworkEquipment2.Contract.SalesOrgId, "2000")
	assert.Equal(NetworkEquipment2.FeatureAvailability.Automation, false)
	assert.Equal(NetworkEquipment2.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(NetworkEquipment2.FeatureAvailability.IpmiReboot, false)
	assert.Equal(NetworkEquipment2.FeatureAvailability.PowerCycle, false)
	assert.Equal(NetworkEquipment2.FeatureAvailability.RemoteManagement, false)
	assert.Equal(NetworkEquipment2.Location.Rack, "XX00")
	assert.Equal(NetworkEquipment2.Location.Site, "AMS-01")
	assert.Equal(NetworkEquipment2.Location.Suite, "HALL3")
	assert.Equal(NetworkEquipment2.Location.Unit, "21")
	assert.Empty(NetworkEquipment2.NetworkInterfaces.Internal.Gateway)
	assert.Empty(NetworkEquipment2.NetworkInterfaces.Internal.Ip)
	assert.Equal(len(NetworkEquipment2.NetworkInterfaces.Internal.Ports), 0)
	assert.Equal(NetworkEquipment2.NetworkInterfaces.Public.Gateway, "127.1.1.254")
	assert.Equal(NetworkEquipment2.NetworkInterfaces.Public.Ip, "127.1.1.68/24")
	assert.Equal(NetworkEquipment2.NetworkInterfaces.Public.LocationId, "")
	assert.Equal(NetworkEquipment2.NetworkInterfaces.Public.NullRouted, false)
	assert.Equal(len(NetworkEquipment2.NetworkInterfaces.Public.Ports), 0)
	assert.Empty(NetworkEquipment2.NetworkInterfaces.RemoteManagement.Gateway)
	assert.Empty(NetworkEquipment2.NetworkInterfaces.RemoteManagement.Ip)
	assert.Empty(NetworkEquipment2.NetworkInterfaces.RemoteManagement.LocationId)
	assert.Equal(len(NetworkEquipment2.NetworkInterfaces.RemoteManagement.Ports), 0)
	assert.Equal(NetworkEquipment2.Type, "SWITCH")
}

func TestDedicatedNetworkEquipmentListBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "networkEquipments": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedNetworkEquipmentApi{}.List(ctx, DedicatedNetworkEquipmentListOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NetworkEquipments), 0)
}

func TestDedicatedNetworkEquipmentListPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"networkEquipments": [
				{
					"contract": {
						"customerId": "10085996",
						"deliveryStatus": "ACTIVE",
						"id": "49031391001170",
						"reference": "My Private Switch",
						"salesOrgId": "2000"
					},
					"featureAvailability": {
						"automation": true,
						"ipmiReboot": false,
						"powerCycle": true,
						"privateNetwork": false,
						"remoteManagement": false
					},
					"id": "12345",
					"location": {
						"rack": "YY11",
						"site": "AMS-01",
						"suite": "HALL3",
						"unit": "21"
					},
					"networkInterfaces": {
						"internal": {
							"gateway": null,
							"ip": null,
							"ports": []
						},
						"public": {
							"gateway": "127.0.2.254",
							"ip": "127.0.2.1/24",
							"locationId": "",
							"nullRouted": false,
							"ports": []
						},
						"remoteManagement": {
							"gateway": null,
							"ip": null,
							"locationId": null,
							"ports": []
						}
					},
					"type": "SWITCH"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := DedicatedNetworkEquipmentListOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(1),
		},
	}
	response, err := DedicatedNetworkEquipmentApi{}.List(ctx, opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NetworkEquipments), 1)

	NetworkEquipment1 := response.NetworkEquipments[0]
	assert.Equal(NetworkEquipment1.Id, "12345")
	assert.Equal(NetworkEquipment1.Contract.CustomerId, "10085996")
	assert.Equal(NetworkEquipment1.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(NetworkEquipment1.Contract.Id, "49031391001170")
	assert.Equal(NetworkEquipment1.Contract.Reference, "My Private Switch")
	assert.Equal(NetworkEquipment1.Contract.SalesOrgId, "2000")
	assert.Equal(NetworkEquipment1.FeatureAvailability.Automation, true)
	assert.Equal(NetworkEquipment1.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(NetworkEquipment1.FeatureAvailability.IpmiReboot, false)
	assert.Equal(NetworkEquipment1.FeatureAvailability.PowerCycle, true)
	assert.Equal(NetworkEquipment1.FeatureAvailability.RemoteManagement, false)
	assert.Equal(NetworkEquipment1.Location.Rack, "YY11")
	assert.Equal(NetworkEquipment1.Location.Site, "AMS-01")
	assert.Equal(NetworkEquipment1.Location.Suite, "HALL3")
	assert.Equal(NetworkEquipment1.Location.Unit, "21")
	assert.Empty(NetworkEquipment1.NetworkInterfaces.Internal.Gateway)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.Internal.Ip)
	assert.Equal(len(NetworkEquipment1.NetworkInterfaces.Internal.Ports), 0)
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.Gateway, "127.0.2.254")
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.Ip, "127.0.2.1/24")
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.LocationId, "")
	assert.Equal(NetworkEquipment1.NetworkInterfaces.Public.NullRouted, false)
	assert.Equal(len(NetworkEquipment1.NetworkInterfaces.Public.Ports), 0)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.RemoteManagement.Gateway)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.RemoteManagement.Ip)
	assert.Empty(NetworkEquipment1.NetworkInterfaces.RemoteManagement.LocationId)
	assert.Equal(len(NetworkEquipment1.NetworkInterfaces.RemoteManagement.Ports), 0)
	assert.Equal(NetworkEquipment1.Type, "SWITCH")
}

func TestDedicatedNetworkEquipmentListServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.List(ctx, DedicatedNetworkEquipmentListOptions{})
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
				return DedicatedNetworkEquipmentApi{}.List(ctx, DedicatedNetworkEquipmentListOptions{})
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
				return DedicatedNetworkEquipmentApi{}.List(ctx, DedicatedNetworkEquipmentListOptions{})
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
				return DedicatedNetworkEquipmentApi{}.List(ctx, DedicatedNetworkEquipmentListOptions{})
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

func TestDedicatedNetworkEquipmentGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"contract": {
				"aggregationPackId": null,
				"billingCycle": 1,
				"billingFrequency": "MONTH",
				"contractTerm": 12,
				"contractType": "NORMAL",
				"currency": "EUR",
				"customerId": "10085996",
				"deliveryStatus": "ACTIVE",
				"endsAt": null,
				"id": "49031391001170",
				"networkTraffic": {
					"connectivityType": "INTERCONNECTED",
					"datatrafficLimit": null,
					"datatrafficUnit": null,
					"trafficType": null,
					"type": null
				},
				"pricePerFrequency": 0.00,
				"reference": "My Switch",
				"salesOrgId": "2000",
				"sla": "Platinum",
				"startsAt": "2017-08-01T00:00:00Z",
				"status": "ACTIVE",
				"subnets": []
			},
			"featureAvailability": {
				"automation": true,
				"ipmiReboot": false,
				"powerCycle": true,
				"privateNetwork": false,
				"remoteManagement": false
			},
			"id": "12345",
			"location": {
				"rack": "YY11",
				"site": "AMS-01",
				"suite": "HALL3",
				"unit": "21"
			},
			"name": "ABC-DE-001",
			"networkInterfaces": {
				"internal": {
					"gateway": null,
					"ip": null,
					"ports": []
				},
				"public": {
					"gateway": "127.0.0.254",
					"ip": "127.0.0.124/24",
					"locationId": "",
					"nullRouted": false,
					"ports": []
				},
				"remoteManagement": {
					"gateway": null,
					"ip": null,
					"locationId": null,
					"ports": []
				}
			},
			"powerPorts": [
				{
					"name": "AMS-01-HALL3-YY11-PDU01",
					"port": "7"
				}
			],
			"rack": {
				"capacity": "",
				"id": "11111"
			},
			"serialNumber": "XN51FPD0QX",
			"specs": {
				"brand": "HP",
				"model": "PC 2530-48 J9781A"
			},
			"type": "SWITCH"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	NetworkEquipment, err := DedicatedNetworkEquipmentApi{}.Get(ctx, "12345")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(NetworkEquipment.Contract.CustomerId, "10085996")
	assert.Equal(NetworkEquipment.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(NetworkEquipment.Contract.Id, "49031391001170")
	assert.Equal(NetworkEquipment.Contract.Reference, "My Switch")
	assert.Equal(NetworkEquipment.Contract.SalesOrgId, "2000")
	assert.Equal(NetworkEquipment.Contract.BillingCycle.String(), "1")
	assert.Equal(NetworkEquipment.Contract.BillingFrequency, "MONTH")
	assert.Equal(NetworkEquipment.Contract.ContractTerm.String(), "12")
	assert.Equal(NetworkEquipment.Contract.Currency, "EUR")
	assert.Empty(NetworkEquipment.Contract.EndsAt)
	assert.Empty(NetworkEquipment.Contract.NetworkTraffic.DataTrafficLimit)
	assert.Empty(NetworkEquipment.Contract.NetworkTraffic.DataTrafficUnit)
	assert.Empty(NetworkEquipment.Contract.NetworkTraffic.TrafficType)
	assert.Empty(NetworkEquipment.Contract.NetworkTraffic.Type)
	assert.Equal(NetworkEquipment.Contract.NetworkTraffic.ConnectivityType, "INTERCONNECTED")
	assert.Equal(NetworkEquipment.Contract.PricePerFrequency.String(), "0.00")
	assert.Equal(NetworkEquipment.Contract.Sla, "Platinum")
	assert.Equal(NetworkEquipment.Contract.StartsAt, "2017-08-01T00:00:00Z")
	assert.Empty(NetworkEquipment.Contract.AggregationPackId)
	assert.Equal(NetworkEquipment.Contract.ContractType, "NORMAL")
	assert.Equal(NetworkEquipment.Contract.Status, "ACTIVE")
	assert.Equal(len(NetworkEquipment.Contract.Subnets), 0)
	assert.Equal(NetworkEquipment.Id, "12345")
	assert.Equal(NetworkEquipment.FeatureAvailability.Automation, true)
	assert.Equal(NetworkEquipment.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(NetworkEquipment.FeatureAvailability.IpmiReboot, false)
	assert.Equal(NetworkEquipment.FeatureAvailability.PowerCycle, true)
	assert.Equal(NetworkEquipment.FeatureAvailability.RemoteManagement, false)
	assert.Equal(NetworkEquipment.Location.Rack, "YY11")
	assert.Equal(NetworkEquipment.Location.Site, "AMS-01")
	assert.Equal(NetworkEquipment.Location.Suite, "HALL3")
	assert.Equal(NetworkEquipment.Location.Unit, "21")
	assert.Empty(NetworkEquipment.NetworkInterfaces.Internal.Gateway)
	assert.Empty(NetworkEquipment.NetworkInterfaces.Internal.Ip)
	assert.Equal(len(NetworkEquipment.NetworkInterfaces.Internal.Ports), 0)
	assert.Equal(NetworkEquipment.NetworkInterfaces.Public.Gateway, "127.0.0.254")
	assert.Equal(NetworkEquipment.NetworkInterfaces.Public.Ip, "127.0.0.124/24")
	assert.Equal(NetworkEquipment.NetworkInterfaces.Public.LocationId, "")
	assert.Equal(NetworkEquipment.NetworkInterfaces.Public.NullRouted, false)
	assert.Equal(len(NetworkEquipment.NetworkInterfaces.Public.Ports), 0)
	assert.Empty(NetworkEquipment.NetworkInterfaces.RemoteManagement.Gateway)
	assert.Empty(NetworkEquipment.NetworkInterfaces.RemoteManagement.Ip)
	assert.Empty(NetworkEquipment.NetworkInterfaces.RemoteManagement.LocationId)
	assert.Equal(len(NetworkEquipment.NetworkInterfaces.RemoteManagement.Ports), 0)
	assert.Equal(NetworkEquipment.PowerPorts[0].Name, "AMS-01-HALL3-YY11-PDU01")
	assert.Equal(NetworkEquipment.PowerPorts[0].Port, "7")
	assert.Equal(NetworkEquipment.Specs.Brand, "HP")
	assert.Equal(NetworkEquipment.Specs.Model, "PC 2530-48 J9781A")
	assert.Equal(NetworkEquipment.Rack.Capacity, "")
	assert.Equal(NetworkEquipment.Rack.Id, "11111")
	assert.Equal(NetworkEquipment.SerialNumber, "XN51FPD0QX")
	assert.Equal(NetworkEquipment.Type, "SWITCH")
}

func TestDedicatedNetworkEquipmentGetServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.Get(ctx, "12345")
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
				return DedicatedNetworkEquipmentApi{}.Get(ctx, "12345")
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
				return DedicatedNetworkEquipmentApi{}.Get(ctx, "12345")
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
				return DedicatedNetworkEquipmentApi{}.Get(ctx, "12345")
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

func TestDedicatedNetworkEquipmentUpdate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedNetworkEquipmentApi{}.Update(ctx, "2893829", "new reference")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedNetworkEquipmentUpdateServerErrors(t *testing.T) {
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
				return nil, DedicatedNetworkEquipmentApi{}.Update(ctx, "2893829", "new reference")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedNetworkEquipmentApi{}.Update(ctx, "2893829", "new reference")
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
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, DedicatedNetworkEquipmentApi{}.Update(ctx, "2893829", "new reference")
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
				return nil, DedicatedNetworkEquipmentApi{}.Update(ctx, "2893829", "new reference")
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

func TestDedicatedNetworkEquipmentListIps(t *testing.T) {
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
	response, err := DedicatedNetworkEquipmentApi{}.ListIps(ctx, "server-id", DedicatedNetworkEquipmentListIpsOptions{})
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

func TestDedicatedNetworkEquipmentListIpsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "ips": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedNetworkEquipmentApi{}.ListIps(ctx, "server-id", DedicatedNetworkEquipmentListIpsOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 0)
}

func TestDedicatedNetworkEquipmentListIpsFilterAndPagination(t *testing.T) {
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
	opts := DedicatedNetworkEquipmentListIpsOptions{
		PaginationOptions: PaginationOptions{
			Limit:  Int(1),
			Offset: Int(10),
		},
	}
	response, err := DedicatedNetworkEquipmentApi{}.ListIps(ctx, "server-id", opts)
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

func TestDedicatedNetworkEquipmentListIpsServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.ListIps(ctx, "server-id", DedicatedNetworkEquipmentListIpsOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListIps(ctx, "server-id", DedicatedNetworkEquipmentListIpsOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListIps(ctx, "server-id", DedicatedNetworkEquipmentListIpsOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListIps(ctx, "server-id", DedicatedNetworkEquipmentListIpsOptions{})
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

func TestDedicatedNetworkEquipmentGetIp(t *testing.T) {
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
	Ip, err := DedicatedNetworkEquipmentApi{}.GetIp(ctx, "12345", "127.0.0.6")
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

func TestDedicatedNetworkEquipmentGetIpServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.GetIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.GetIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.GetIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.GetIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.GetIp(ctx, "12345", "127.0.0.6")
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

func TestDedicatedNetworkEquipmentUpdateIp(t *testing.T) {
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
	Ip, err := DedicatedNetworkEquipmentApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
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

func TestDedicatedNetworkEquipmentUpdateIpServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
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
				return DedicatedNetworkEquipmentApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
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
				return DedicatedNetworkEquipmentApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
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
				return DedicatedNetworkEquipmentApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
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
				return DedicatedNetworkEquipmentApi{}.UpdateIp(ctx, "12345", "127.0.0.6", payload)
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

func TestDedicatedNetworkEquipmentNullRouteAnIp(t *testing.T) {
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
	Ip, err := DedicatedNetworkEquipmentApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
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

func TestDedicatedNetworkEquipmentNullRouteAnIpServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.NullRouteAnIp(ctx, "12345", "127.0.0.6")
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

func TestDedicatedNetworkEquipmentRemoveNullRouteAnIp(t *testing.T) {
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
	Ip, err := DedicatedNetworkEquipmentApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
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

func TestDedicatedNetworkEquipmentRemoveNullRouteAnIpServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
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
				return DedicatedNetworkEquipmentApi{}.RemoveNullRouteAnIp(ctx, "12345", "127.0.0.6")
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

func TestDedicatedNetworkEquipmentListNullRoutes(t *testing.T) {
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
	response, err := DedicatedNetworkEquipmentApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
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

func TestDedicatedNetworkEquipmentListNullRoutesBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "nullRoutes": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedNetworkEquipmentApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NullRoutes), 0)
}

func TestDedicatedNetworkEquipmentListNullRoutesFilterAndPagination(t *testing.T) {
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
	response, err := DedicatedNetworkEquipmentApi{}.ListNullRoutes(ctx, "server-id", opts)
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

func TestDedicatedNetworkEquipmentListNullRoutesServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListNullRoutes(ctx, "server-id", PaginationOptions{})
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

func TestDedicatedNetworkEquipmentListCredentials(t *testing.T) {
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
	response, err := DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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

func TestDedicatedNetworkEquipmentListCredentialsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestDedicatedNetworkEquipmentListCredentialsPaginate(t *testing.T) {
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
	response, err := DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
}

func TestDedicatedNetworkEquipmentListCredentialsServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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

func TestDedicatedNetworkEquipmentCreateCredential(t *testing.T) {
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
	resp, err := DedicatedNetworkEquipmentApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Type, "OPERATING_SYSTEM")
	assert.Equal(resp.Username, "root")
	assert.Equal(resp.Password, "mys3cr3tp@ssw0rd")
}

func TestDedicatedNetworkEquipmentCreateCredentialServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedNetworkEquipmentApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedNetworkEquipmentApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedNetworkEquipmentApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedNetworkEquipmentApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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

func TestDedicatedNetworkEquipmentListCredentialsByType(t *testing.T) {
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
	response, err := DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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

func TestDedicatedNetworkEquipmentListCredentialsByTypeBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestDedicatedNetworkEquipmentListCredentialsByTypePaginate(t *testing.T) {
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
	response, err := DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
}

func TestDedicatedNetworkEquipmentListCredentialsByTypeServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedNetworkEquipmentApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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

func TestDedicatedNetworkEquipmentGetCredentials(t *testing.T) {
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
	Credential, err := DedicatedNetworkEquipmentApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Credential.Password, "mys3cr3tp@ssw0rd")
	assert.Equal(Credential.Username, "root")
	assert.Equal(Credential.Type, "OPERATING_SYSTEM")
}

func TestDedicatedNetworkEquipmentGetCredentialsServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedNetworkEquipmentApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedNetworkEquipmentApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedNetworkEquipmentApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedNetworkEquipmentApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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

func TestDedicatedNetworkEquipmentCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedNetworkEquipmentApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedNetworkEquipmentCredentialsServerErrors(t *testing.T) {
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
				return nil, DedicatedNetworkEquipmentApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedNetworkEquipmentApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedNetworkEquipmentApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedNetworkEquipmentApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedNetworkEquipmentApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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

func TestDedicatedNetworkEquipmentUpdateCredentials(t *testing.T) {
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
	resp, err := DedicatedNetworkEquipmentApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Password, "new password")
	assert.Equal(resp.Type, "OPERATING_SYSTEM")
	assert.Equal(resp.Username, "admin")
}

func TestDedicatedNetworkEquipmentUpdateCredentialsServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedNetworkEquipmentApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedNetworkEquipmentApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedNetworkEquipmentApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedNetworkEquipmentApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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

func TestDedicatedNetworkEquipmentPowerCycleServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedNetworkEquipmentApi{}.PowerCycleServer(ctx, "server id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedNetworkEquipmentPowerCycleServerServerErrors(t *testing.T) {
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerCycleServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerCycleServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerCycleServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerCycleServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerCycleServer(ctx, "server id")
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

func TestDedicatedNetworkEquipmentPowerOffServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedNetworkEquipmentApi{}.PowerOffServer(ctx, "server id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedNetworkEquipmentPowerOffServerServerErrors(t *testing.T) {
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOffServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOffServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOffServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOffServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOffServer(ctx, "server id")
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

func TestDedicatedNetworkEquipmentPowerOnServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedNetworkEquipmentApi{}.PowerOnServer(ctx, "server id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedNetworkEquipmentPowerOnServerServerErrors(t *testing.T) {
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOnServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOnServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOnServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOnServer(ctx, "server id")
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
				return nil, DedicatedNetworkEquipmentApi{}.PowerOnServer(ctx, "server id")
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

func TestDedicatedNetworkEquipmentGetPowerStatus(t *testing.T) {
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
	resp, err := DedicatedNetworkEquipmentApi{}.GetPowerStatus(ctx, "server-id")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Ipmi.Status, "off")
	assert.Equal(resp.Pdu.Status, "on")
}

func TestDedicatedNetworkEquipmentGetPowerStatusServerErrors(t *testing.T) {
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
				return DedicatedNetworkEquipmentApi{}.GetPowerStatus(ctx, "server-id")
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
				return DedicatedNetworkEquipmentApi{}.GetPowerStatus(ctx, "server-id")
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
				return DedicatedNetworkEquipmentApi{}.GetPowerStatus(ctx, "server-id")
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
				return DedicatedNetworkEquipmentApi{}.GetPowerStatus(ctx, "server-id")
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
				return DedicatedNetworkEquipmentApi{}.GetPowerStatus(ctx, "server-id")
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
