package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpManagementListIps(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ips": [
				{
					"ip": "192.0.2.1",
					"version": 4,
					"type": "NORMAL_IP",
					"prefixLength": 32,
					"primary": true,
					"reverseLookup": "mydomain1.example.com",
					"nullRouted": false,
					"unnullingAllowed": false,
					"equipmentId": "1234",
					"assignedContract": {
						"id": "5643634"
					},
					"subnet": {
						"id": "192.0.2.0_24",
						"networkIp": "192.0.2.0",
						"prefixLength": 24,
						"gateway": "192.0.2.254"
					}
				},
				{
					"ip": "192.0.2.2",
					"version": 4,
					"type": "NORMAL_IP",
					"prefixLength": 32,
					"primary": false,
					"reverseLookup": "mydomain2.example.com",
					"nullRouted": false,
					"unnullingAllowed": false,
					"equipmentId": "1235",
					"assignedContract": {
						"id": "4363465"
					},
					"subnet": {
						"id": "192.0.2.0_24",
						"networkIp": "192.0.2.0",
						"prefixLength": 24,
						"gateway": "192.0.2.254"
					}
				}
			],
			"_metadata": {
				"totalCount": 2,
				"offset": 0,
				"limit": 10
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := IpManagementApi{}.List(ctx)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 2)

	Ip1 := response.Ips[0]
	assert.Equal(Ip1.Ip, "192.0.2.1")
	assert.Equal(Ip1.Version.String(), "4")
	assert.Equal(Ip1.Type, "NORMAL_IP")
	assert.Equal(Ip1.PrefixLength.String(), "32")
	assert.Equal(Ip1.Primary, true)
	assert.Equal(Ip1.ReverseLookup, "mydomain1.example.com")
	assert.Equal(Ip1.NullRouted, false)
	assert.Equal(Ip1.UnnullingAllowed, false)
	assert.Equal(Ip1.EquipmentId, "1234")
	assert.Equal(Ip1.Subnet.Gateway, "192.0.2.254")
	assert.Equal(Ip1.Subnet.Id, "192.0.2.0_24")
	assert.Equal(Ip1.Subnet.NetworkIp, "192.0.2.0")
	assert.Equal(Ip1.Subnet.PrefixLength.String(), "24")
	assert.Equal(Ip1.AssignedContract.Id, "5643634")

	Ip2 := response.Ips[1]
	assert.Equal(Ip2.Ip, "192.0.2.2")
	assert.Equal(Ip2.Version.String(), "4")
	assert.Equal(Ip2.Type, "NORMAL_IP")
	assert.Equal(Ip2.PrefixLength.String(), "32")
	assert.Equal(Ip2.Primary, false)
	assert.Equal(Ip2.ReverseLookup, "mydomain2.example.com")
	assert.Equal(Ip2.NullRouted, false)
	assert.Equal(Ip2.UnnullingAllowed, false)
	assert.Equal(Ip2.EquipmentId, "1235")
	assert.Equal(Ip2.Subnet.Gateway, "192.0.2.254")
	assert.Equal(Ip2.Subnet.Id, "192.0.2.0_24")
	assert.Equal(Ip2.Subnet.NetworkIp, "192.0.2.0")
	assert.Equal(Ip2.Subnet.PrefixLength.String(), "24")
	assert.Equal(Ip2.AssignedContract.Id, "4363465")
}

func TestIpManagementListIpsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ips": [
				{
					"ip": "192.0.2.1",
					"version": 4,
					"type": "NORMAL_IP",
					"prefixLength": 32,
					"primary": true,
					"reverseLookup": "mydomain1.example.com",
					"nullRouted": false,
					"unnullingAllowed": false,
					"equipmentId": "1234",
					"assignedContract": {
						"id": "5643634"
					},
					"subnet": {
						"id": "192.0.2.0_24",
						"networkIp": "192.0.2.0",
						"prefixLength": 24,
						"gateway": "192.0.2.254"
					}
				}
			],
			"_metadata": {
				"totalCount": 11,
				"offset": 1,
				"limit": 10
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := IpManagementApi{}.List(ctx, map[string]interface{}{"offset": 1})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ips), 1)

	Ip1 := response.Ips[0]
	assert.Equal(Ip1.Ip, "192.0.2.1")
	assert.Equal(Ip1.Version.String(), "4")
	assert.Equal(Ip1.Type, "NORMAL_IP")
	assert.Equal(Ip1.PrefixLength.String(), "32")
	assert.Equal(Ip1.Primary, true)
	assert.Equal(Ip1.ReverseLookup, "mydomain1.example.com")
	assert.Equal(Ip1.NullRouted, false)
	assert.Equal(Ip1.UnnullingAllowed, false)
	assert.Equal(Ip1.EquipmentId, "1234")
	assert.Equal(Ip1.Subnet.Gateway, "192.0.2.254")
	assert.Equal(Ip1.Subnet.Id, "192.0.2.0_24")
	assert.Equal(Ip1.Subnet.NetworkIp, "192.0.2.0")
	assert.Equal(Ip1.Subnet.PrefixLength.String(), "24")
	assert.Equal(Ip1.AssignedContract.Id, "5643634")
}

func TestIpManagementListIpsServerErrors(t *testing.T) {
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
				return IpManagementApi{}.List(ctx, map[string]interface{}{"offset": 1})
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.List(ctx, map[string]interface{}{"offset": 1})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.List(ctx, map[string]interface{}{"offset": 1})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.List(ctx, map[string]interface{}{"offset": 1})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestIpManagementGetIps(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ip": "192.0.2.1",
			"version": 4,
			"type": "NORMAL_IP",
			"prefixLength": 32,
			"primary": true,
			"reverseLookup": "mydomain1.example.com",
			"nullRouted": false,
			"unnullingAllowed": false,
			"equipmentId": "1234",
			"assignedContract": {
				"id": "5643634"
			},
			"subnet": {
				"id": "192.0.2.0_24",
				"networkIp": "192.0.2.0",
				"prefixLength": 24,
				"gateway": "192.0.2.254"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := IpManagementApi{}.Get(ctx, "192.0.2.1")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Ip.Ip, "192.0.2.1")
	assert.Equal(Ip.Version.String(), "4")
	assert.Equal(Ip.Type, "NORMAL_IP")
	assert.Equal(Ip.PrefixLength.String(), "32")
	assert.Equal(Ip.Primary, true)
	assert.Equal(Ip.ReverseLookup, "mydomain1.example.com")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.UnnullingAllowed, false)
	assert.Equal(Ip.EquipmentId, "1234")
	assert.Equal(Ip.Subnet.Gateway, "192.0.2.254")
	assert.Equal(Ip.Subnet.Id, "192.0.2.0_24")
	assert.Equal(Ip.Subnet.NetworkIp, "192.0.2.0")
	assert.Equal(Ip.Subnet.PrefixLength.String(), "24")
	assert.Equal(Ip.AssignedContract.Id, "5643634")
}

func TestIpManagementGetIpsServerErrors(t *testing.T) {
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
				return IpManagementApi{}.Get(ctx, "192.0.2.1")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.Get(ctx, "192.0.2.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.Get(ctx, "192.0.2.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.Get(ctx, "192.0.2.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestIpManagementUpdate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ip": "192.0.2.1",
			"version": 4,
			"type": "NORMAL_IP",
			"prefixLength": 32,
			"primary": true,
			"reverseLookup": "mydomain1.example.com",
			"nullRouted": false,
			"unnullingAllowed": false,
			"equipmentId": "1234",
			"assignedContract": {
				"id": "5643634"
			},
			"subnet": {
				"id": "192.0.2.0_24",
				"networkIp": "192.0.2.0",
				"prefixLength": 24,
				"gateway": "192.0.2.254"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Ip, err := IpManagementApi{}.Update(ctx, "192.0.2.1", "mydomain1.example.com")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Ip.Ip, "192.0.2.1")
	assert.Equal(Ip.Version.String(), "4")
	assert.Equal(Ip.Type, "NORMAL_IP")
	assert.Equal(Ip.PrefixLength.String(), "32")
	assert.Equal(Ip.Primary, true)
	assert.Equal(Ip.ReverseLookup, "mydomain1.example.com")
	assert.Equal(Ip.NullRouted, false)
	assert.Equal(Ip.UnnullingAllowed, false)
	assert.Equal(Ip.EquipmentId, "1234")
	assert.Equal(Ip.Subnet.Gateway, "192.0.2.254")
	assert.Equal(Ip.Subnet.Id, "192.0.2.0_24")
	assert.Equal(Ip.Subnet.NetworkIp, "192.0.2.0")
	assert.Equal(Ip.Subnet.PrefixLength.String(), "24")
	assert.Equal(Ip.AssignedContract.Id, "5643634")
}

func TestIpManagementUpdateServerErrors(t *testing.T) {
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
				return IpManagementApi{}.Update(ctx, "192.0.2.1", "mydomain1.example.com")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.Update(ctx, "192.0.2.1", "mydomain1.example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.Update(ctx, "192.0.2.1", "mydomain1.example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.Update(ctx, "192.0.2.1", "mydomain1.example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestIpManagementNullRouteAnIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "4534536",
			"ip": "192.0.2.1",
			"nulledAt": "2015-06-28T12:00:00Z",
			"nulledBy": "john.doe@example.com",
			"nullLevel": 1,
			"automatedUnnullingAt": "2015-06-25T11:13:00Z",
			"unnulledAt": null,
			"unnulledBy": null,
			"ticketId": "188612",
			"comment": "This IP is evil",
			"equipmentId": "456",
			"assignedContract": {
				"id": "123456"
			}
		}`)
	})
	defer teardown()

	payload := map[string]string{
		"automatedUnnullingAt": "2015-06-25T11:13:00Z",
		"ticketId":             "188612",
		"comment":              "This IP is evil",
	}
	ctx := context.Background()
	NullRoute, err := IpManagementApi{}.NullRouteAnIp(ctx, "192.0.2.1", payload)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(NullRoute.Id, "4534536")
	assert.Equal(NullRoute.AssignedContract.Id, "123456")
	assert.Equal(NullRoute.AutomatedUnnullingAt, "2015-06-25T11:13:00Z")
	assert.Equal(NullRoute.Comment, "This IP is evil")
	assert.Equal(NullRoute.EquipmentId, "456")
	assert.Equal(NullRoute.Ip, "192.0.2.1")
	assert.Equal(NullRoute.NullLevel.String(), "1")
	assert.Equal(NullRoute.NulledAt, "2015-06-28T12:00:00Z")
	assert.Equal(NullRoute.NulledBy, "john.doe@example.com")
	assert.Equal(NullRoute.TicketId, "188612")
	assert.Empty(NullRoute.UnnulledAt)
	assert.Empty(NullRoute.UnnulledBy)
}

func TestIpManagementNullRouteAnIpServerErrors(t *testing.T) {
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
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.NullRouteAnIp(ctx, "192.0.2.1", payload)
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.NullRouteAnIp(ctx, "192.0.2.1", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.NullRouteAnIp(ctx, "192.0.2.1", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.NullRouteAnIp(ctx, "192.0.2.1", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestIpManagementRemoveNullRouteAnIp(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := IpManagementApi{}.RemoveNullRouteAnIp(ctx, "192.0.2.1")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestIpManagementRemoveNullRouteAnIpServerErrors(t *testing.T) {
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
				return nil, IpManagementApi{}.RemoveNullRouteAnIp(ctx, "192.0.2.1")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, IpManagementApi{}.RemoveNullRouteAnIp(ctx, "192.0.2.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, IpManagementApi{}.RemoveNullRouteAnIp(ctx, "192.0.2.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, IpManagementApi{}.RemoveNullRouteAnIp(ctx, "192.0.2.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestIpManagementListNullRoutes(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"nullroutes": [
				{
					"id": "4534536",
					"ip": "192.0.2.1",
					"nulledAt": "2015-06-28T12:00:00Z",
					"nulledBy": "john.doe@example.com",
					"nullLevel": 1,
					"automatedUnnullingAt": "2015-06-28T13:00:00Z",
					"unnulledAt": null,
					"unnulledBy": null,
					"ticketId": "188612",
					"comment": "This IP is evil",
					"equipmentid": "456",
					"assignedContract": {
						"id": "123456"
					}
				},
				{
					"id": "4534535",
					"ip": "192.0.2.1",
					"nulledAt": "2015-06-27T12:00:00Z",
					"nulledBy": "john.doe@example.com",
					"nullLevel": 1,
					"automatedUnnullingAt": "2015-06-27T13:00:00Z",
					"unnulledAt": "2015-06-27T13:00:05Z",
					"unnulledBy": "UnnullRunner",
					"ticketId": "188612",
					"comment": "This IP is evil",
					"equipmentId": "456",
					"assignedContract": {
						"id": "123456"
					}
				}
			],
			"_metadata": {
				"totalCount": 2,
				"offset": 0,
				"limit": 10
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := IpManagementApi{}.ListNullRoutes(ctx)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NullRoutes), 2)

	NullRoute1 := response.NullRoutes[0]
	assert.Equal(NullRoute1.Id, "4534536")
	assert.Equal(NullRoute1.AssignedContract.Id, "123456")
	assert.Equal(NullRoute1.AutomatedUnnullingAt, "2015-06-28T13:00:00Z")
	assert.Equal(NullRoute1.Comment, "This IP is evil")
	assert.Equal(NullRoute1.EquipmentId, "456")
	assert.Equal(NullRoute1.Ip, "192.0.2.1")
	assert.Equal(NullRoute1.NullLevel.String(), "1")
	assert.Equal(NullRoute1.NulledAt, "2015-06-28T12:00:00Z")
	assert.Equal(NullRoute1.NulledBy, "john.doe@example.com")
	assert.Equal(NullRoute1.TicketId, "188612")
	assert.Empty(NullRoute1.UnnulledAt)
	assert.Empty(NullRoute1.UnnulledBy)
	NullRoute2 := response.NullRoutes[1]

	assert.Equal(NullRoute2.Id, "4534535")
	assert.Equal(NullRoute2.AssignedContract.Id, "123456")
	assert.Equal(NullRoute2.AutomatedUnnullingAt, "2015-06-27T13:00:00Z")
	assert.Equal(NullRoute2.Comment, "This IP is evil")
	assert.Equal(NullRoute2.EquipmentId, "456")
	assert.Equal(NullRoute2.Ip, "192.0.2.1")
	assert.Equal(NullRoute2.NullLevel.String(), "1")
	assert.Equal(NullRoute2.NulledAt, "2015-06-27T12:00:00Z")
	assert.Equal(NullRoute2.NulledBy, "john.doe@example.com")
	assert.Equal(NullRoute2.TicketId, "188612")
	assert.Equal(NullRoute2.UnnulledAt, "2015-06-27T13:00:05Z")
	assert.Equal(NullRoute2.UnnulledBy, "UnnullRunner")
}

func TestIpManagementListNullRoutesPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"nullroutes": [
				{
					"id": "4534536",
					"ip": "192.0.2.1",
					"nulledAt": "2015-06-28T12:00:00Z",
					"nulledBy": "john.doe@example.com",
					"nullLevel": 1,
					"automatedUnnullingAt": "2015-06-28T13:00:00Z",
					"unnulledAt": null,
					"unnulledBy": null,
					"ticketId": "188612",
					"comment": "This IP is evil",
					"equipmentid": "456",
					"assignedContract": {
						"id": "123456"
					}
				}
			],
			"_metadata": {
				"totalCount": 11,
				"offset": 1,
				"limit": 10
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := IpManagementApi{}.ListNullRoutes(ctx, map[string]interface{}{"offset": 1})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.NullRoutes), 1)

	NullRoute1 := response.NullRoutes[0]
	assert.Equal(NullRoute1.Id, "4534536")
	assert.Equal(NullRoute1.AssignedContract.Id, "123456")
	assert.Equal(NullRoute1.AutomatedUnnullingAt, "2015-06-28T13:00:00Z")
	assert.Equal(NullRoute1.Comment, "This IP is evil")
	assert.Equal(NullRoute1.EquipmentId, "456")
	assert.Equal(NullRoute1.Ip, "192.0.2.1")
	assert.Equal(NullRoute1.NullLevel.String(), "1")
	assert.Equal(NullRoute1.NulledAt, "2015-06-28T12:00:00Z")
	assert.Equal(NullRoute1.NulledBy, "john.doe@example.com")
	assert.Equal(NullRoute1.TicketId, "188612")
	assert.Empty(NullRoute1.UnnulledAt)
	assert.Empty(NullRoute1.UnnulledBy)
}

func TestIpManagementListNullRoutesServerErrors(t *testing.T) {
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
				return IpManagementApi{}.ListNullRoutes(ctx, map[string]interface{}{"offset": 1})
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.ListNullRoutes(ctx, map[string]interface{}{"offset": 1})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.ListNullRoutes(ctx, map[string]interface{}{"offset": 1})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.ListNullRoutes(ctx, map[string]interface{}{"offset": 1})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestIpManagementGetNullRoute(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "4534536",
			"ip": "192.0.2.1",
			"nulledAt": "2015-06-28T12:00:00Z",
			"nulledBy": "john.doe@example.com",
			"nullLevel": 1,
			"automatedUnnullingAt": "2015-06-28T13:00:00Z",
			"unnulledAt": null,
			"unnulledBy": null,
			"ticketId": "188612",
			"comment": "This IP is evil",
			"equipmentid": "456",
			"assignedContract": {
				"id": "123456"
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	NullRoute, err := IpManagementApi{}.GetNullRoute(ctx, "123456")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(NullRoute.Id, "4534536")
	assert.Equal(NullRoute.AssignedContract.Id, "123456")
	assert.Equal(NullRoute.AutomatedUnnullingAt, "2015-06-28T13:00:00Z")
	assert.Equal(NullRoute.Comment, "This IP is evil")
	assert.Equal(NullRoute.EquipmentId, "456")
	assert.Equal(NullRoute.Ip, "192.0.2.1")
	assert.Equal(NullRoute.NullLevel.String(), "1")
	assert.Equal(NullRoute.NulledAt, "2015-06-28T12:00:00Z")
	assert.Equal(NullRoute.NulledBy, "john.doe@example.com")
	assert.Equal(NullRoute.TicketId, "188612")
	assert.Empty(NullRoute.UnnulledAt)
	assert.Empty(NullRoute.UnnulledBy)
}

func TestIpManagementGetNullRouteServerErrors(t *testing.T) {
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
				return IpManagementApi{}.GetNullRoute(ctx, "123456")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.GetNullRoute(ctx, "123456")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.GetNullRoute(ctx, "123456")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return IpManagementApi{}.GetNullRoute(ctx, "123456")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestIpManagementUpdateNullRouteHistory(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "4534536",
			"ip": "192.0.2.1",
			"nulledAt": "2015-06-28T12:00:00Z",
			"nulledBy": "john.doe@example.com",
			"nullLevel": 1,
			"automatedUnnullingAt": "2015-06-28T13:00:00Z",
			"unnulledAt": null,
			"unnulledBy": null,
			"ticketId": "188612",
			"comment": "This IP is evil",
			"equipmentid": "456",
			"assignedContract": {
				"id": "123456"
			}
		}`)
	})
	defer teardown()

	payload := map[string]string{
		"automatedUnnullingAt": "2015-06-25T11:13:00Z",
		"ticketId":             "188612",
		"comment":              "This IP is evil",
	}
	ctx := context.Background()
	NullRoute, err := IpManagementApi{}.UpdateNullRoute(ctx, "123456", payload)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(NullRoute.Id, "4534536")
	assert.Equal(NullRoute.AssignedContract.Id, "123456")
	assert.Equal(NullRoute.AutomatedUnnullingAt, "2015-06-28T13:00:00Z")
	assert.Equal(NullRoute.Comment, "This IP is evil")
	assert.Equal(NullRoute.EquipmentId, "456")
	assert.Equal(NullRoute.Ip, "192.0.2.1")
	assert.Equal(NullRoute.NullLevel.String(), "1")
	assert.Equal(NullRoute.NulledAt, "2015-06-28T12:00:00Z")
	assert.Equal(NullRoute.NulledBy, "john.doe@example.com")
	assert.Equal(NullRoute.TicketId, "188612")
	assert.Empty(NullRoute.UnnulledAt)
	assert.Empty(NullRoute.UnnulledBy)
}

func TestIpManagementUpdateNullRouteHistoryServerErrors(t *testing.T) {
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
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.UpdateNullRoute(ctx, "123456", payload)
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.UpdateNullRoute(ctx, "123456", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.UpdateNullRoute(ctx, "123456", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				payload := map[string]string{
					"automatedUnnullingAt": "2015-06-25T11:13:00Z",
					"ticketId":             "188612",
					"comment":              "This IP is evil",
				}
				ctx := context.Background()
				return IpManagementApi{}.UpdateNullRoute(ctx, "123456", payload)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
