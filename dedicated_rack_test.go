package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDedicatedRackList(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"privateRacks": [
				{
					"contract": {
						"customerId": "2738283",
						"deliveryStatus": "ACTIVE",
						"id": "123456",
						"reference": "AAAA - Private rack 001",
						"salesOrgId": "2000"
					},
					"featureAvailability": {
						"powerCycle": false,
						"privateNetwork": false
					},
					"id": "123456",
					"location": {
						"rack": "22",
						"site": "AMS-01",
						"suite": "8.24"
					},
					"networkInterfaces": {
						"public": {
							"ports": [
								{
									"name": "EVO-BB99-1",
									"port": "0-9"
								}
							]
						}
					}
				},
				{
					"contract": {
						"customerId": "2738283",
						"deliveryStatus": "ACTIVE",
						"id": "267940",
						"reference": "AAAA - Private rack 002",
						"salesOrgId": "2000"
					},
					"featureAvailability": {
						"powerCycle": false,
						"privateNetwork": false
					},
					"id": "267940",
					"location": {
						"rack": "MX66",
						"site": "AMS-01",
						"suite": "Hall3"
					},
					"networkInterfaces": {
						"public": {
							"ports": [
								{
									"name": "ce99.ams-01",
									"port": "0-1"
								}
							]
						}
					}
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedRackApi{}.List(ctx, DedicatedRackListOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateRacks), 2)

	Rack1 := response.PrivateRacks[0]
	assert.Equal(Rack1.Id, "123456")
	assert.Equal(Rack1.Contract.CustomerId, "2738283")
	assert.Equal(Rack1.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(Rack1.Contract.Id, "123456")
	assert.Equal(Rack1.Contract.Reference, "AAAA - Private rack 001")
	assert.Equal(Rack1.Contract.SalesOrgId, "2000")
	assert.Equal(Rack1.FeatureAvailability.PowerCycle, false)
	assert.Equal(Rack1.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(Rack1.Location.Rack, "22")
	assert.Equal(Rack1.Location.Site, "AMS-01")
	assert.Equal(Rack1.Location.Suite, "8.24")
	assert.Equal(Rack1.NetworkInterfaces.Public.Ports[0].Name, "EVO-BB99-1")
	assert.Equal(Rack1.NetworkInterfaces.Public.Ports[0].Port, "0-9")

	Rack2 := response.PrivateRacks[1]
	assert.Equal(Rack2.Id, "267940")
	assert.Equal(Rack2.Contract.CustomerId, "2738283")
	assert.Equal(Rack2.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(Rack2.Contract.Id, "267940")
	assert.Equal(Rack2.Contract.Reference, "AAAA - Private rack 002")
	assert.Equal(Rack2.Contract.SalesOrgId, "2000")
	assert.Equal(Rack2.FeatureAvailability.PowerCycle, false)
	assert.Equal(Rack2.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(Rack2.Location.Rack, "MX66")
	assert.Equal(Rack2.Location.Site, "AMS-01")
	assert.Equal(Rack2.Location.Suite, "Hall3")
	assert.Equal(Rack2.NetworkInterfaces.Public.Ports[0].Name, "ce99.ams-01")
	assert.Equal(Rack2.NetworkInterfaces.Public.Ports[0].Port, "0-1")
}

func TestDedicatedRackListBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "privateRacks": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedRackApi{}.List(ctx, DedicatedRackListOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateRacks), 0)
}

func TestDedicatedRackListPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"privateRacks": [
				{
					"contract": {
						"customerId": "2738283",
						"deliveryStatus": "ACTIVE",
						"id": "123456",
						"reference": "AAAA - Private rack 001",
						"salesOrgId": "2000"
					},
					"featureAvailability": {
						"powerCycle": false,
						"privateNetwork": false
					},
					"id": "123456",
					"location": {
						"rack": "22",
						"site": "AMS-01",
						"suite": "8.24"
					},
					"networkInterfaces": {
						"public": {
							"ports": [
								{
									"name": "EVO-BB99-1",
									"port": "0-9"
								}
							]
						}
					}
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := DedicatedRackListOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(1),
		},
	}
	response, err := DedicatedRackApi{}.List(ctx, opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateRacks), 1)

	Rack1 := response.PrivateRacks[0]
	assert.Equal(Rack1.Id, "123456")
	assert.Equal(Rack1.Contract.CustomerId, "2738283")
	assert.Equal(Rack1.Contract.DeliveryStatus, "ACTIVE")
	assert.Equal(Rack1.Contract.Id, "123456")
	assert.Equal(Rack1.Contract.Reference, "AAAA - Private rack 001")
	assert.Equal(Rack1.Contract.SalesOrgId, "2000")
	assert.Equal(Rack1.FeatureAvailability.PowerCycle, false)
	assert.Equal(Rack1.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(Rack1.Location.Rack, "22")
	assert.Equal(Rack1.Location.Site, "AMS-01")
	assert.Equal(Rack1.Location.Suite, "8.24")
	assert.Equal(Rack1.NetworkInterfaces.Public.Ports[0].Name, "EVO-BB99-1")
	assert.Equal(Rack1.NetworkInterfaces.Public.Ports[0].Port, "0-9")
}

func TestDedicatedRackListServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.List(ctx, DedicatedRackListOptions{})
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
				return DedicatedRackApi{}.List(ctx, DedicatedRackListOptions{})
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
				return DedicatedRackApi{}.List(ctx, DedicatedRackListOptions{})
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
				return DedicatedRackApi{}.List(ctx, DedicatedRackListOptions{})
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

func TestDedicatedRackGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"contract": {
				"customerId": "2738283",
				"deliveryStatus": "ACTIVE",
				"endsAt": null,
				"id": "2893829",
				"networkTraffic": {
					"datatrafficLimit": 0,
					"datatrafficUnit": null,
					"trafficType": "CUSTOM",
					"type": "CONNECTIVITY"
				},
				"reference": "AAAA - Private rack 002",
				"salesOrgId": "2000",
				"sla": "Platinum - 24x7x½",
				"startsAt": "2017-08-01T00:00:00"
			},
			"featureAvailability": {
				"powerCycle": false,
				"privateNetwork": false
			},
			"id": "2893829",
			"location": {
				"rack": "MI15",
				"site": "AMS-01",
				"suite": "Hall3"
			},
			"networkInterfaces": {
				"public": {
					"ports": [
						{
							"name": "ce05.ams-01",
							"port": "0-26"
						}
					]
				}
			},
			"powerPorts": [],
			"units": [
				{
					"unit": "1",
					"status": "FREE",
					"connectedUnits": ["1"]
				},
				{
					"unit": "13",
					"status": "OCCUPIED",
					"connectedUnits": ["13", "14"]
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Rack, err := DedicatedRackApi{}.Get(ctx, "2893829")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Rack.Id, "2893829")
	assert.Equal(Rack.Contract.CustomerId, "2738283")
	assert.Equal(Rack.Contract.DeliveryStatus, "ACTIVE")
	assert.Empty(Rack.Contract.EndsAt)
	assert.Equal(Rack.Contract.Id, "2893829")
	assert.Equal(Rack.Contract.Reference, "AAAA - Private rack 002")
	assert.Equal(Rack.Contract.SalesOrgId, "2000")
	assert.Equal(Rack.Contract.NetworkTraffic.DataTrafficLimit.String(), "0")
	assert.Empty(Rack.Contract.NetworkTraffic.DataTrafficUnit)
	assert.Equal(Rack.Contract.NetworkTraffic.TrafficType, "CUSTOM")
	assert.Equal(Rack.Contract.NetworkTraffic.Type, "CONNECTIVITY")
	assert.Equal(Rack.Contract.Sla, "Platinum - 24x7x½")
	assert.Equal(Rack.Contract.StartsAt, "2017-08-01T00:00:00")
	assert.Equal(Rack.FeatureAvailability.PowerCycle, false)
	assert.Equal(Rack.FeatureAvailability.PrivateNetwork, false)
	assert.Equal(Rack.Location.Rack, "MI15")
	assert.Equal(Rack.Location.Site, "AMS-01")
	assert.Equal(Rack.Location.Suite, "Hall3")
	assert.Equal(Rack.NetworkInterfaces.Public.Ports[0].Name, "ce05.ams-01")
	assert.Equal(Rack.NetworkInterfaces.Public.Ports[0].Port, "0-26")
	assert.Equal(len(Rack.PowerPorts), 0)
	assert.Equal(Rack.Units[0].Unit, "1")
	assert.Equal(Rack.Units[0].Status, "FREE")
	assert.Equal(Rack.Units[0].ConnectedUnits[0], "1")
	assert.Equal(Rack.Units[1].Unit, "13")
	assert.Equal(Rack.Units[1].Status, "OCCUPIED")
	assert.Equal(Rack.Units[1].ConnectedUnits[0], "13")
	assert.Equal(Rack.Units[1].ConnectedUnits[1], "14")
}

func TestDedicatedRackGetServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.Get(ctx, "2893829")
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
				return DedicatedRackApi{}.Get(ctx, "2893829")
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
				return DedicatedRackApi{}.Get(ctx, "2893829")
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
				return DedicatedRackApi{}.Get(ctx, "2893829")
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

func TestDedicatedRackUpdate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedRackApi{}.Update(ctx, "2893829", "new reference")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedRackUpdateServerErrors(t *testing.T) {
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
				return nil, DedicatedRackApi{}.Update(ctx, "2893829", "new reference")
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
				return nil, DedicatedRackApi{}.Update(ctx, "2893829", "new reference")
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
				return nil, DedicatedRackApi{}.Update(ctx, "2893829", "new reference")
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
				return nil, DedicatedRackApi{}.Update(ctx, "2893829", "new reference")
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

func TestDedicatedRackListNullRoutes(t *testing.T) {
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
	response, err := DedicatedRackApi{}.ListNullRoutes(ctx, "123456", PaginationOptions{})
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

func TestDedicatedRackListNullRoutesPaginateAndFilter(t *testing.T) {
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
	response, err := DedicatedRackApi{}.ListNullRoutes(ctx, "123456", opts)
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

func TestDedicatedRackListNullRoutesServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.ListNullRoutes(ctx, "123456", PaginationOptions{})
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
				return DedicatedRackApi{}.ListNullRoutes(ctx, "123456", PaginationOptions{})
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
				return DedicatedRackApi{}.ListNullRoutes(ctx, "123456", PaginationOptions{})
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
				return DedicatedRackApi{}.ListNullRoutes(ctx, "123456", PaginationOptions{})
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

func TestDedicatedRackListIps(t *testing.T) {
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
					"ip": "12.123.123.1/24",
					"gateway": "12.123.123.254",
					"floatingIp": false,
					"version": 4,
					"nullRouted": true,
					"reverseLookup": "domain.example.com",
					"mainIp": true,
					"networkType": "PUBLIC",
					"ddos": {
						"detectionProfile": "ADVANCED_LOW_UDP",
						"protectionType": "ADVANCED"
					}
				},
				{
					"ip": "2001:db8:85a3::8a2e:370:7334/64",
					"gateway": "2001:db8:85a3::8a2e:370:1",
					"floatingIp": false,
					"version": 6,
					"nullRouted": false,
					"reverseLookup": "domain.example.com",
					"mainIp": false,
					"networkType": "REMOTE_MANAGEMENT",
					"ddos": {
						"detectionProfile": "STANDARD_DEFAULT",
						"protectionType": "STANDARD"
					}
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := DedicatedRackApi{}.ListIps(ctx, "123456", DedicatedRackListIpsOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 2)

	Ip1 := response.Ips[0]
	assert.Equal(Ip1.Ip, "12.123.123.1/24")
	assert.Equal(Ip1.Gateway, "12.123.123.254")
	assert.Equal(Ip1.FloatingIp, false)
	assert.Equal(Ip1.Version.String(), "4")
	assert.Equal(Ip1.NullRouted, true)
	assert.Equal(Ip1.ReverseLookup, "domain.example.com")
	assert.Equal(Ip1.MainIp, true)
	assert.Equal(Ip1.NetworkType, "PUBLIC")
	assert.Equal(Ip1.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip1.DDOS.ProtectionType, "ADVANCED")

	Ip2 := response.Ips[1]
	assert.Equal(Ip2.Ip, "2001:db8:85a3::8a2e:370:7334/64")
	assert.Equal(Ip2.Gateway, "2001:db8:85a3::8a2e:370:1")
	assert.Equal(Ip2.FloatingIp, false)
	assert.Equal(Ip2.Version.String(), "6")
	assert.Equal(Ip2.NullRouted, false)
	assert.Equal(Ip2.ReverseLookup, "domain.example.com")
	assert.Equal(Ip2.MainIp, false)
	assert.Equal(Ip2.NetworkType, "REMOTE_MANAGEMENT")
	assert.Equal(Ip2.DDOS.DetectionProfile, "STANDARD_DEFAULT")
	assert.Equal(Ip2.DDOS.ProtectionType, "STANDARD")
}

func TestDedicatedRackListIpsPaginateAndFilter(t *testing.T) {
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
					"ip": "12.123.123.1/24",
					"gateway": "12.123.123.254",
					"floatingIp": false,
					"version": 4,
					"nullRouted": true,
					"reverseLookup": "domain.example.com",
					"mainIp": true,
					"networkType": "PUBLIC",
					"ddos": {
						"detectionProfile": "ADVANCED_LOW_UDP",
						"protectionType": "ADVANCED"
					}
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := DedicatedRackListIpsOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(1),
		},
	}
	response, err := DedicatedRackApi{}.ListIps(ctx, "123456", opts)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 1)

	Ip1 := response.Ips[0]
	assert.Equal(Ip1.Ip, "12.123.123.1/24")
	assert.Equal(Ip1.Gateway, "12.123.123.254")
	assert.Equal(Ip1.FloatingIp, false)
	assert.Equal(Ip1.Version.String(), "4")
	assert.Equal(Ip1.NullRouted, true)
	assert.Equal(Ip1.ReverseLookup, "domain.example.com")
	assert.Equal(Ip1.MainIp, true)
	assert.Equal(Ip1.NetworkType, "PUBLIC")
	assert.Equal(Ip1.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip1.DDOS.ProtectionType, "ADVANCED")
}

func TestDedicatedRackListIpsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.ListIps(ctx, "123456", DedicatedRackListIpsOptions{})
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
				return DedicatedRackApi{}.ListIps(ctx, "123456", DedicatedRackListIpsOptions{})
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
				return DedicatedRackApi{}.ListIps(ctx, "123456", DedicatedRackListIpsOptions{})
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
				return DedicatedRackApi{}.ListIps(ctx, "123456", DedicatedRackListIpsOptions{})
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

func TestDedicatedRackGetIps(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ip": "12.123.123.1/24",
			"gateway": "12.123.123.254",
			"floatingIp": false,
			"version": 4,
			"nullRouted": false,
			"reverseLookup": "domain.example.com",
			"mainIp": true,
			"networkType": "PUBLIC",
			"ddos": {
				"detectionProfile": "ADVANCED_LOW_UDP",
				"protectionType": "ADVANCED"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := DedicatedRackApi{}.GetIp(ctx, "123456", "12.123.123.1")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Version.String(), "4")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
}

func TestDedicatedRackGetIpsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.GetIp(ctx, "123456", "12.123.123.1")
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
				return DedicatedRackApi{}.GetIp(ctx, "123456", "12.123.123.1")
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
				return DedicatedRackApi{}.GetIp(ctx, "123456", "12.123.123.1")
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
				return DedicatedRackApi{}.GetIp(ctx, "123456", "12.123.123.1")
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

func TestDedicatedRackUpdateIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ip": "12.123.123.1/24",
			"gateway": "12.123.123.254",
			"floatingIp": false,
			"version": 4,
			"nullRouted": false,
			"reverseLookup": "domain.example.com",
			"mainIp": true,
			"networkType": "PUBLIC",
			"ddos": {
				"detectionProfile": "ADVANCED_LOW_UDP",
				"protectionType": "ADVANCED"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := DedicatedRackApi{}.UpdateIp(ctx, "123456", "12.123.123.1", map[string]string{"reverseLookup": "domain.example.com", "detectionProfile": "ADVANCED_LOW_UDP"})
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Version.String(), "4")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
}

func TestDedicatedRackUpdateIpServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.UpdateIp(ctx, "123456", "12.123.123.1", map[string]string{"reverseLookup": "domain.example.com", "detectionProfile": "ADVANCED_LOW_UDP"})
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
				return DedicatedRackApi{}.UpdateIp(ctx, "123456", "12.123.123.1", map[string]string{"reverseLookup": "domain.example.com", "detectionProfile": "ADVANCED_LOW_UDP"})
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
				return DedicatedRackApi{}.UpdateIp(ctx, "123456", "12.123.123.1", map[string]string{"reverseLookup": "domain.example.com", "detectionProfile": "ADVANCED_LOW_UDP"})
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
				return DedicatedRackApi{}.UpdateIp(ctx, "123456", "12.123.123.1", map[string]string{"reverseLookup": "domain.example.com", "detectionProfile": "ADVANCED_LOW_UDP"})
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

func TestDedicatedRackNullRouteAnIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ip": "12.123.123.1/24",
			"gateway": "12.123.123.254",
			"floatingIp": false,
			"version": 4,
			"nullRouted": false,
			"reverseLookup": "domain.example.com",
			"mainIp": true,
			"networkType": "PUBLIC",
			"ddos": {
				"detectionProfile": "ADVANCED_LOW_UDP",
				"protectionType": "ADVANCED"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := DedicatedRackApi{}.NullRouteAnIp(ctx, "123456", "12.123.123.1")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Version.String(), "4")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
}

func TestDedicatedRackNullRouteAnIpServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.NullRouteAnIp(ctx, "123456", "12.123.123.1")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedRackApi{}.NullRouteAnIp(ctx, "123456", "12.123.123.1")
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
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedRackApi{}.NullRouteAnIp(ctx, "123456", "12.123.123.1")
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
				return DedicatedRackApi{}.NullRouteAnIp(ctx, "123456", "12.123.123.1")
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

func TestDedicatedRackRemoveNullRouteAnIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ip": "12.123.123.1/24",
			"gateway": "12.123.123.254",
			"floatingIp": false,
			"version": 4,
			"nullRouted": false,
			"reverseLookup": "domain.example.com",
			"mainIp": true,
			"networkType": "PUBLIC",
			"ddos": {
				"detectionProfile": "ADVANCED_LOW_UDP",
				"protectionType": "ADVANCED"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := DedicatedRackApi{}.RemoveNullRouteAnIp(ctx, "123456", "12.123.123.1")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Ip.Ip, "12.123.123.1/24")
	assert.Equal(Ip.Gateway, "12.123.123.254")
	assert.Equal(Ip.FloatingIp, false)
	assert.Equal(Ip.Version.String(), "4")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.ReverseLookup, "domain.example.com")
	assert.Equal(Ip.MainIp, true)
	assert.Equal(Ip.NetworkType, "PUBLIC")
	assert.Equal(Ip.DDOS.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(Ip.DDOS.ProtectionType, "ADVANCED")
}

func TestDedicatedRackRemoveNullRouteAnIpServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.RemoveNullRouteAnIp(ctx, "123456", "12.123.123.1")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedRackApi{}.RemoveNullRouteAnIp(ctx, "123456", "12.123.123.1")
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
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return DedicatedRackApi{}.RemoveNullRouteAnIp(ctx, "123456", "12.123.123.1")
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
				return DedicatedRackApi{}.RemoveNullRouteAnIp(ctx, "123456", "12.123.123.1")
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

func TestDedicatedRackListCredentials(t *testing.T) {
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
	response, err := DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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

func TestDedicatedRackListCredentialsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestDedicatedRackListCredentialsPaginate(t *testing.T) {
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
	response, err := DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
}

func TestDedicatedRackListCredentialsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentials(ctx, "99944", PaginationOptions{})
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

func TestDedicatedRackCreateCredential(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Type, "OPERATING_SYSTEM")
	assert.Equal(resp.Username, "root")
	assert.Equal(resp.Password, "mys3cr3tp@ssw0rd")
}

func TestDedicatedRackCreateCredentialServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedRackApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedRackApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedRackApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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
				return DedicatedRackApi{}.CreateCredential(ctx, "12345", "OPERATING_SYSTEM", "root", "mys3cr3tp@ssw0rd")
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

func TestDedicatedRackListCredentialsByType(t *testing.T) {
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
	response, err := DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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

func TestDedicatedRackListCredentialsByTypeBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()
	ctx := context.Background()
	response, err := DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestDedicatedRackListCredentialsByTypePaginate(t *testing.T) {
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
	response, err := DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	assert.Equal(response.Credentials[0].Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Credentials[0].Username, "admin")
}

func TestDedicatedRackListCredentialsByTypeServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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
				return DedicatedRackApi{}.ListCredentialsByType(ctx, "99944", "REMOTE_MANAGEMENT", PaginationOptions{})
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

func TestDedicatedRackGetCredentials(t *testing.T) {
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
	Credential, err := DedicatedRackApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Credential.Password, "mys3cr3tp@ssw0rd")
	assert.Equal(Credential.Username, "root")
	assert.Equal(Credential.Type, "OPERATING_SYSTEM")
}

func TestDedicatedRackGetCredentialsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedRackApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedRackApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedRackApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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
				return DedicatedRackApi{}.GetCredential(ctx, "12345", "OPERATING_SYSTEM", "root")
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

func TestDedicatedRackDeleteCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedRackApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedRackDeleteCredentialsServerErrors(t *testing.T) {
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
				return nil, DedicatedRackApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedRackApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedRackApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedRackApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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
				return nil, DedicatedRackApi{}.DeleteCredential(ctx, "12345", "OPERATING_SYSTEM", "admin")
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

func TestDedicatedRackUpdateCredentials(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Password, "new password")
	assert.Equal(resp.Type, "OPERATING_SYSTEM")
	assert.Equal(resp.Username, "admin")
}

func TestDedicatedRackUpdateCredentialsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedRackApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedRackApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedRackApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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
				return DedicatedRackApi{}.UpdateCredential(ctx, "12345", "OPERATING_SYSTEM", "admin", "new password")
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

func TestDedicatedRackGetBandWidthMetrics(t *testing.T) {
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
	Metric, err := DedicatedRackApi{}.GetBandWidthMetrics(ctx, "99944", opts)
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

func TestDedicatedRackGetBandWidthMetricsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.GetBandWidthMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetBandWidthMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetBandWidthMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetBandWidthMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetBandWidthMetrics(ctx, "99944", opts)
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

func TestDedicatedRackGetDataTrafficMetrics(t *testing.T) {
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
	Metric, err := DedicatedRackApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
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

func TestDedicatedRackGetDataTrafficMetricsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
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
				return DedicatedRackApi{}.GetDataTrafficMetrics(ctx, "99944", opts)
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

func TestDedicatedRackGetDdosNotificationSetting(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.GetDdosNotificationSetting(ctx, "server-id")
	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Nulling, "ENABLED")
	assert.Equal(resp.Scrubbing, "DISABLED")
}

func TestDedicatedRackGetDdosNotificationSettingServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.GetDdosNotificationSetting(ctx, "server-id")
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
				return DedicatedRackApi{}.GetDdosNotificationSetting(ctx, "server-id")
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
				return DedicatedRackApi{}.GetDdosNotificationSetting(ctx, "server-id")
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
				return DedicatedRackApi{}.GetDdosNotificationSetting(ctx, "server-id")
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
				return DedicatedRackApi{}.GetDdosNotificationSetting(ctx, "server-id")
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

func TestDedicatedRackUpdateDdosNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedRackApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedRackUpdateDdosNotificationSettingServerErrors(t *testing.T) {
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
				return nil, DedicatedRackApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
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
				return nil, DedicatedRackApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
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
				return nil, DedicatedRackApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
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
				return nil, DedicatedRackApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
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
				return nil, DedicatedRackApi{}.UpdateDdosNotificationSetting(ctx, "server id", map[string]string{"nulling": "DISABLED", "scrubbing": "ENABLED"})
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

func TestDedicatedRackListDataTrafficNotificationSettings(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
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

func TestDedicatedRackListDataTrafficNotificationSettingsPaginate(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", opts)
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

func TestDedicatedRackListDataTrafficNotificationSettingsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListDataTrafficNotificationSettings(ctx, "server-id", PaginationOptions{})
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

func TestDedicatedRackCreateDataTrafficNotificationSetting(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", "1", "GB")
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

func TestDedicatedRackCreateDataTrafficNotificationSettingServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", "1", "GB")
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
				return DedicatedRackApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", "1", "GB")
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
				return DedicatedRackApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", "1", "GB")
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
				return DedicatedRackApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", "1", "GB")
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
				return DedicatedRackApi{}.CreateDataTrafficNotificationSetting(ctx, "server-id", "WEEKLY", "1", "GB")
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

func TestDedicatedRackDeleteDataTrafficNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedRackApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedRackDeleteDataTrafficNotificationSettingServerErrors(t *testing.T) {
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
				return nil, DedicatedRackApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteDataTrafficNotificationSetting(ctx, "server id", "notification id")
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

func TestDedicatedRackGetDataTrafficNotificationSetting(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
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

func TestDedicatedRackGetDataTrafficNotificationSettingServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetDataTrafficNotificationSetting(ctx, "server-id", "notification id")
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

func TestDedicatedRackUpdateDataTrafficNotificationSetting(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
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

func TestDedicatedRackUpdateDataTrafficNotificationSettingServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateDataTrafficNotificationSetting(ctx, "server-id", "notification id", payload)
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

func TestDedicatedRackListBandWidthNotificationSettings(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
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

func TestDedicatedRackListBandWidthNotificationSettingsPaginate(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.ListBandWidthNotificationSettings(ctx, "server-id", opts)
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

func TestDedicatedRackListBandWidthNotificationSettingsServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
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
				return DedicatedRackApi{}.ListBandWidthNotificationSettings(ctx, "server-id", PaginationOptions{})
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

func TestDedicatedRackCreateBandWidthNotificationSetting(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", "1", "Gbps")
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

func TestDedicatedRackCreateBandWidthNotificationSettingServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", "1", "Gbps")
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
				return DedicatedRackApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", "1", "Gbps")
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
				return DedicatedRackApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", "1", "Gbps")
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
				return DedicatedRackApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", "1", "Gbps")
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
				return DedicatedRackApi{}.CreateBandWidthNotificationSetting(ctx, "server-id", "WEEKLY", "1", "Gbps")
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

func TestDedicatedRackDeleteBandWidthNotificationSetting(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := DedicatedRackApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestDedicatedRackDeleteBandWidthNotificationSettingServerErrors(t *testing.T) {
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
				return nil, DedicatedRackApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
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
				return nil, DedicatedRackApi{}.DeleteBandWidthNotificationSetting(ctx, "server id", "notification id")
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

func TestDedicatedRackGetBandWidthNotificationSetting(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
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

func TestDedicatedRackGetBandWidthNotificationSettingServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
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
				return DedicatedRackApi{}.GetBandWidthNotificationSetting(ctx, "server-id", "notification id")
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

func TestDedicatedRackUpdateBandWidthNotificationSetting(t *testing.T) {
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
	resp, err := DedicatedRackApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
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

func TestDedicatedRackUpdateBandWidthNotificationSettingServerErrors(t *testing.T) {
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
				return DedicatedRackApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
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
				return DedicatedRackApi{}.UpdateBandWidthNotificationSetting(ctx, "server-id", "notification id", payload)
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
