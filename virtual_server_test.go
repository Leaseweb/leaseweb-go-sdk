package leaseweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListVirtualServers(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"virtualServers": [
				{
					"id": "222903",
					"reference": "Web server",
					"customerId": "1301178860",
					"dataCenter": "AMS-01",
					"cloudServerId": null,
					"state": "STOPPED",
					"firewallState": "DISABLED",
					"template": "Ubuntu 14.04 64 40 20140707T1340",
					"serviceOffering": "S",
					"sla": "Bronze",
					"contract": {
						"id": "30000778",
						"startsAt": "2016-02-01T00:00:00+0200",
						"endsAt": "2017-01-31T00:00:00+0200",
						"billingCycle": 12,
						"billingFrequency": "MONTH",
						"pricePerFrequency": 4.7,
						"currency": "EUR"
					},
					"hardware": {
						"cpu": {
							"cores": 1
						},
						"memory": {
							"unit": "MB",
							"amount": 1024
						},
						"storage": {
							"unit": "GB",
							"amount": 40
						}
					},
					"iso": null,
					"ips": [
						{
							"ip": "10.11.116.130",
							"version": 4,
							"type": "PUBLIC"
						}
					]
				},
				{
					"id": "301708",
					"reference": null,
					"customerId": "1301178860",
					"dataCenter": "AMS-01",
					"cloudServerId": null,
					"state": "STOPPED",
					"firewallState": "ENABLED",
					"template": "CentOS 7.0 64 60 20140711T1039",
					"serviceOffering": "M",
					"sla": "Bronze",
					"contract": {
						"id": "30000779",
						"startsAt": "2016-02-01T00:00:00+0200",
						"endsAt": "2017-01-31T00:00:00+0200",
						"billingCycle": 12,
						"billingFrequency": "MONTH",
						"pricePerFrequency": 4.7,
						"currency": "EUR"
					},
					"hardware": {
						"cpu": {
							"cores": 2
						},
						"memory": {
							"unit": "MB",
							"amount": 2048
						},
						"storage": {
							"unit": "GB",
							"amount": 60
						}
					},
					"iso": {
						"id": "9eadbe14-69be-4dee-8f56-5ebb23bb3c33",
						"name": "Knoppix",
						"displayName": "Knoppix"
					},
					"ips": [
						{
							"ip": "10.11.116.132",
							"version": 4,
							"type": "PUBLIC"
						}
					]
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

	response, err := VirtualServerApi{}.ListVirtualServers()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.VirtualServers), 2)

	virtualServer1 := response.VirtualServers[0]
	assert.Empty(virtualServer1.CloudServerId)
	assert.Equal(virtualServer1.CustomerId, "1301178860")
	assert.Equal(virtualServer1.DataCenter, "AMS-01")
	assert.Equal(virtualServer1.FirewallState, "DISABLED")
	assert.Equal(virtualServer1.Id, "222903")
	assert.Empty(virtualServer1.Iso)
	assert.Equal(virtualServer1.Ips[0].Ip, "10.11.116.130")
	assert.Equal(virtualServer1.Ips[0].Version, 4)
	assert.Equal(virtualServer1.Ips[0].Type, "PUBLIC")
	assert.Equal(virtualServer1.Contract.BillingCycle, 12)
	assert.Equal(virtualServer1.Contract.BillingFrequency, "MONTH")
	assert.Equal(virtualServer1.Contract.Currency, "EUR")
	assert.Equal(virtualServer1.Contract.EndsAt, "2017-01-31T00:00:00+0200")
	assert.Equal(virtualServer1.Contract.Id, "30000778")
	assert.Equal(virtualServer1.Contract.PricePerFrequency, float32(4.7))
	assert.Equal(virtualServer1.Contract.StartsAt, "2016-02-01T00:00:00+0200")
	assert.Equal(virtualServer1.Hardware.Cpu.Cores, 1)
	assert.Equal(virtualServer1.Hardware.Memory.Amount, 1024)
	assert.Equal(virtualServer1.Hardware.Memory.Unit, "MB")
	assert.Equal(virtualServer1.Hardware.Storage.Amount, 40)
	assert.Equal(virtualServer1.Hardware.Storage.Unit, "GB")
	assert.Equal(virtualServer1.Reference, "Web server")
	assert.Equal(virtualServer1.ServiceOffering, "S")
	assert.Equal(virtualServer1.Sla, "Bronze")
	assert.Equal(virtualServer1.State, "STOPPED")
	assert.Equal(virtualServer1.Template, "Ubuntu 14.04 64 40 20140707T1340")

	virtualServer2 := response.VirtualServers[1]
	assert.Empty(virtualServer2.CloudServerId)
	assert.Equal(virtualServer2.CustomerId, "1301178860")
	assert.Equal(virtualServer2.DataCenter, "AMS-01")
	assert.Equal(virtualServer2.FirewallState, "ENABLED")
	assert.Equal(virtualServer2.Id, "301708")
	assert.Empty(virtualServer2.Reference)
	assert.Equal(virtualServer2.ServiceOffering, "M")
	assert.Equal(virtualServer2.Sla, "Bronze")
	assert.Equal(virtualServer2.State, "STOPPED")
	assert.Equal(virtualServer2.Template, "CentOS 7.0 64 60 20140711T1039")
	assert.Equal(virtualServer2.Iso.Id, "9eadbe14-69be-4dee-8f56-5ebb23bb3c33")
	assert.Equal(virtualServer2.Iso.Name, "Knoppix")
	assert.Equal(virtualServer2.Iso.DisplayName, "Knoppix")
	assert.Equal(virtualServer2.Ips[0].Ip, "10.11.116.132")
	assert.Equal(virtualServer2.Ips[0].Version, 4)
	assert.Equal(virtualServer2.Ips[0].Type, "PUBLIC")
	assert.Equal(virtualServer2.Contract.BillingCycle, 12)
	assert.Equal(virtualServer2.Contract.BillingFrequency, "MONTH")
	assert.Equal(virtualServer2.Contract.Currency, "EUR")
	assert.Equal(virtualServer2.Contract.EndsAt, "2017-01-31T00:00:00+0200")
	assert.Equal(virtualServer2.Contract.Id, "30000779")
	assert.Equal(virtualServer2.Contract.PricePerFrequency, float32(4.7))
	assert.Equal(virtualServer2.Contract.StartsAt, "2016-02-01T00:00:00+0200")
	assert.Equal(virtualServer2.Hardware.Cpu.Cores, 2)
	assert.Equal(virtualServer2.Hardware.Memory.Amount, 2048)
	assert.Equal(virtualServer2.Hardware.Memory.Unit, "MB")
	assert.Equal(virtualServer2.Hardware.Storage.Amount, 60)
	assert.Equal(virtualServer2.Hardware.Storage.Unit, "GB")
}

func TestListVirtualServersPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"virtualServers": [
				{
					"id": "222903",
					"reference": "Web server",
					"customerId": "1301178860",
					"dataCenter": "AMS-01",
					"cloudServerId": null,
					"state": "STOPPED",
					"firewallState": "DISABLED",
					"template": "Ubuntu 14.04 64 40 20140707T1340",
					"serviceOffering": "S",
					"sla": "Bronze",
					"contract": {
						"id": "30000778",
						"startsAt": "2016-02-01T00:00:00+0200",
						"endsAt": "2017-01-31T00:00:00+0200",
						"billingCycle": 12,
						"billingFrequency": "MONTH",
						"pricePerFrequency": 4.7,
						"currency": "EUR"
					},
					"hardware": {
						"cpu": {
							"cores": 1
						},
						"memory": {
							"unit": "MB",
							"amount": 1024
						},
						"storage": {
							"unit": "GB",
							"amount": 40
						}
					},
					"iso": null,
					"ips": [
						{
							"ip": "10.11.116.130",
							"version": 4,
							"type": "PUBLIC"
						}
					]
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

	response, err := VirtualServerApi{}.ListVirtualServers(1)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.VirtualServers), 1)

	virtualServer1 := response.VirtualServers[0]
	assert.Empty(virtualServer1.CloudServerId)
	assert.Equal(virtualServer1.CustomerId, "1301178860")
	assert.Equal(virtualServer1.DataCenter, "AMS-01")
	assert.Equal(virtualServer1.FirewallState, "DISABLED")
	assert.Equal(virtualServer1.Id, "222903")
	assert.Empty(virtualServer1.Iso)
	assert.Equal(virtualServer1.Ips[0].Ip, "10.11.116.130")
	assert.Equal(virtualServer1.Ips[0].Version, 4)
	assert.Equal(virtualServer1.Ips[0].Type, "PUBLIC")
	assert.Equal(virtualServer1.Contract.BillingCycle, 12)
	assert.Equal(virtualServer1.Contract.BillingFrequency, "MONTH")
	assert.Equal(virtualServer1.Contract.Currency, "EUR")
	assert.Equal(virtualServer1.Contract.EndsAt, "2017-01-31T00:00:00+0200")
	assert.Equal(virtualServer1.Contract.Id, "30000778")
	assert.Equal(virtualServer1.Contract.PricePerFrequency, float32(4.7))
	assert.Equal(virtualServer1.Contract.StartsAt, "2016-02-01T00:00:00+0200")
	assert.Equal(virtualServer1.Hardware.Cpu.Cores, 1)
	assert.Equal(virtualServer1.Hardware.Memory.Amount, 1024)
	assert.Equal(virtualServer1.Hardware.Memory.Unit, "MB")
	assert.Equal(virtualServer1.Hardware.Storage.Amount, 40)
	assert.Equal(virtualServer1.Hardware.Storage.Unit, "GB")
	assert.Equal(virtualServer1.Reference, "Web server")
	assert.Equal(virtualServer1.ServiceOffering, "S")
	assert.Equal(virtualServer1.Sla, "Bronze")
	assert.Equal(virtualServer1.State, "STOPPED")
	assert.Equal(virtualServer1.Template, "Ubuntu 14.04 64 40 20140707T1340")
}

func TestListVirtualServersServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.ListVirtualServers()
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.ListVirtualServers()
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
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
				return VirtualServerApi{}.ListVirtualServers()
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.ListVirtualServers()
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestGetVirtualServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "222903",
			"reference": "Web server",
			"customerId": "1301178860",
			"dataCenter": "AMS-01",
			"cloudServerId": null,
			"state": "STOPPED",
			"firewallState": "DISABLED",
			"template": "Ubuntu 14.04 64 40 20140707T1340",
			"serviceOffering": "S",
			"sla": "Bronze",
			"contract": {
				"id": "30000778",
				"startsAt": "2016-02-01T00:00:00+0200",
				"endsAt": "2017-01-31T00:00:00+0200",
				"billingCycle": 12,
				"billingFrequency": "MONTH",
				"pricePerFrequency": 4.7,
				"currency": "EUR"
			},
			"hardware": {
				"cpu": {
					"cores": 1
				},
				"memory": {
					"unit": "MB",
					"amount": 1024
				},
				"storage": {
					"unit": "GB",
					"amount": 40
				}
			},
			"iso": null,
			"ips": [
				{
					"ip": "10.11.116.130",
					"version": 4,
					"type": "PUBLIC"
				}
			]
		}`)
	})
	defer teardown()

	virtualServer, err := VirtualServerApi{}.GetVirtualServer("123456")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Empty(virtualServer.CloudServerId)
	assert.Equal(virtualServer.CustomerId, "1301178860")
	assert.Equal(virtualServer.DataCenter, "AMS-01")
	assert.Equal(virtualServer.FirewallState, "DISABLED")
	assert.Equal(virtualServer.Id, "222903")
	assert.Empty(virtualServer.Iso)
	assert.Equal(virtualServer.Ips[0].Ip, "10.11.116.130")
	assert.Equal(virtualServer.Ips[0].Version, 4)
	assert.Equal(virtualServer.Ips[0].Type, "PUBLIC")
	assert.Equal(virtualServer.Contract.BillingCycle, 12)
	assert.Equal(virtualServer.Contract.BillingFrequency, "MONTH")
	assert.Equal(virtualServer.Contract.Currency, "EUR")
	assert.Equal(virtualServer.Contract.EndsAt, "2017-01-31T00:00:00+0200")
	assert.Equal(virtualServer.Contract.Id, "30000778")
	assert.Equal(virtualServer.Contract.PricePerFrequency, float32(4.7))
	assert.Equal(virtualServer.Contract.StartsAt, "2016-02-01T00:00:00+0200")
	assert.Equal(virtualServer.Hardware.Cpu.Cores, 1)
	assert.Equal(virtualServer.Hardware.Memory.Amount, 1024)
	assert.Equal(virtualServer.Hardware.Memory.Unit, "MB")
	assert.Equal(virtualServer.Hardware.Storage.Amount, 40)
	assert.Equal(virtualServer.Hardware.Storage.Unit, "GB")
	assert.Equal(virtualServer.Reference, "Web server")
	assert.Equal(virtualServer.ServiceOffering, "S")
	assert.Equal(virtualServer.Sla, "Bronze")
	assert.Equal(virtualServer.State, "STOPPED")
	assert.Equal(virtualServer.Template, "Ubuntu 14.04 64 40 20140707T1340")
}

func TestGetVirtualServerServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.GetVirtualServer("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.GetVirtualServer("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return VirtualServerApi{}.GetVirtualServer("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return VirtualServerApi{}.GetVirtualServer("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.GetVirtualServer("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestUpdateVirtualServer(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "222903",
			"reference": "Web server",
			"customerId": "1301178860",
			"dataCenter": "AMS-01",
			"cloudServerId": null,
			"state": "STOPPED",
			"firewallState": "DISABLED",
			"template": "Ubuntu 14.04 64 40 20140707T1340",
			"serviceOffering": "S",
			"sla": "Bronze",
			"contract": {
				"id": "30000778",
				"startsAt": "2016-02-01T00:00:00+0200",
				"endsAt": "2017-01-31T00:00:00+0200",
				"billingCycle": 12,
				"billingFrequency": "MONTH",
				"pricePerFrequency": 4.7,
				"currency": "EUR"
			},
			"hardware": {
				"cpu": {
					"cores": 1
				},
				"memory": {
					"unit": "MB",
					"amount": 1024
				},
				"storage": {
					"unit": "GB",
					"amount": 40
				}
			},
			"iso": null,
			"ips": [
				{
					"ip": "10.11.116.130",
					"version": 4,
					"type": "PUBLIC"
				}
			]
		}`)
	})
	defer teardown()

	virtualServer, err := VirtualServerApi{}.UpdateVirtualServer("123456", "Web server")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Empty(virtualServer.CloudServerId)
	assert.Equal(virtualServer.CustomerId, "1301178860")
	assert.Equal(virtualServer.DataCenter, "AMS-01")
	assert.Equal(virtualServer.FirewallState, "DISABLED")
	assert.Equal(virtualServer.Id, "222903")
	assert.Empty(virtualServer.Iso)
	assert.Equal(virtualServer.Ips[0].Ip, "10.11.116.130")
	assert.Equal(virtualServer.Ips[0].Version, 4)
	assert.Equal(virtualServer.Ips[0].Type, "PUBLIC")
	assert.Equal(virtualServer.Contract.BillingCycle, 12)
	assert.Equal(virtualServer.Contract.BillingFrequency, "MONTH")
	assert.Equal(virtualServer.Contract.Currency, "EUR")
	assert.Equal(virtualServer.Contract.EndsAt, "2017-01-31T00:00:00+0200")
	assert.Equal(virtualServer.Contract.Id, "30000778")
	assert.Equal(virtualServer.Contract.PricePerFrequency, float32(4.7))
	assert.Equal(virtualServer.Contract.StartsAt, "2016-02-01T00:00:00+0200")
	assert.Equal(virtualServer.Hardware.Cpu.Cores, 1)
	assert.Equal(virtualServer.Hardware.Memory.Amount, 1024)
	assert.Equal(virtualServer.Hardware.Memory.Unit, "MB")
	assert.Equal(virtualServer.Hardware.Storage.Amount, 40)
	assert.Equal(virtualServer.Hardware.Storage.Unit, "GB")
	assert.Equal(virtualServer.Reference, "Web server")
	assert.Equal(virtualServer.ServiceOffering, "S")
	assert.Equal(virtualServer.Sla, "Bronze")
	assert.Equal(virtualServer.State, "STOPPED")
	assert.Equal(virtualServer.Template, "Ubuntu 14.04 64 40 20140707T1340")
}

func TestUpdateVirtualServerServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.UpdateVirtualServer("123456", "Web server")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.UpdateVirtualServer("123456", "Web server")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return VirtualServerApi{}.UpdateVirtualServer("123456", "Web server")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return VirtualServerApi{}.UpdateVirtualServer("123456", "Web server")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.UpdateVirtualServer("123456", "Web server")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPowerOff(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{			
			"id": "cs01.237daad0-2aed-4260-b0e4-488d9cd55607",
			"name": "virtualServers.powerOff",
			"status": "PENDING",
			"createdAt": "2016-12-31T01:00:59+00:00"
		}`)
	})
	defer teardown()

	resp, err := VirtualServerApi{}.PowerOff("123456")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Id, "cs01.237daad0-2aed-4260-b0e4-488d9cd55607")
	assert.Equal(resp.Name, "virtualServers.powerOff")
	assert.Equal(resp.Status, "PENDING")
	assert.Equal(resp.CreatedAt, "2016-12-31T01:00:59+00:00")
}

func TestPowerOffServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.PowerOff("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.PowerOff("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return VirtualServerApi{}.PowerOff("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return VirtualServerApi{}.PowerOff("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.PowerOff("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPowerOn(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{			
			"id": "cs01.237daad0-2aed-4260-b0e4-488d9cd55607",
			"name": "virtualServers.powerOn",
			"status": "PENDING",
			"createdAt": "2016-12-31T01:00:59+00:00"
		}`)
	})
	defer teardown()

	resp, err := VirtualServerApi{}.PowerOn("123456")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Id, "cs01.237daad0-2aed-4260-b0e4-488d9cd55607")
	assert.Equal(resp.Name, "virtualServers.powerOn")
	assert.Equal(resp.Status, "PENDING")
	assert.Equal(resp.CreatedAt, "2016-12-31T01:00:59+00:00")
}

func TestPowerOnServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.PowerOn("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.PowerOn("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return VirtualServerApi{}.PowerOn("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return VirtualServerApi{}.PowerOn("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.PowerOn("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestReboot(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{			
			"id": "cs01.237daad0-2aed-4260-b0e4-488d9cd55607",
			"name": "virtualServers.reboot",
			"status": "PENDING",
			"createdAt": "2016-12-31T01:00:59+00:00"
		}`)
	})
	defer teardown()

	resp, err := VirtualServerApi{}.Reboot("123456")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Id, "cs01.237daad0-2aed-4260-b0e4-488d9cd55607")
	assert.Equal(resp.Name, "virtualServers.reboot")
	assert.Equal(resp.Status, "PENDING")
	assert.Equal(resp.CreatedAt, "2016-12-31T01:00:59+00:00")
}

func TestRebootServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.Reboot("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.Reboot("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return VirtualServerApi{}.Reboot("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return VirtualServerApi{}.Reboot("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.Reboot("123456")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestReinstall(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{			
			"id": "cs01.237daad0-2aed-4260-b0e4-488d9cd55607",
			"name": "virtualServers.reinstall",
			"status": "PENDING",
			"createdAt": "2016-12-31T01:00:59+00:00"
		}`)
	})
	defer teardown()

	resp, err := VirtualServerApi{}.Reinstall("123456", "CENTOS_7_64_PLESK")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Id, "cs01.237daad0-2aed-4260-b0e4-488d9cd55607")
	assert.Equal(resp.Name, "virtualServers.reinstall")
	assert.Equal(resp.Status, "PENDING")
	assert.Equal(resp.CreatedAt, "2016-12-31T01:00:59+00:00")
}

func TestReinstallServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.Reinstall("123456", "CENTOS_7_64_PLESK")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.Reinstall("123456", "CENTOS_7_64_PLESK")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return VirtualServerApi{}.Reinstall("123456", "CENTOS_7_64_PLESK")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return VirtualServerApi{}.Reinstall("123456", "CENTOS_7_64_PLESK")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.Reinstall("123456", "CENTOS_7_64_PLESK")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestVirtualServerUpdateCredential(t *testing.T) {
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

	err := VirtualServerApi{}.UpdateCredential("12345", "OPERATING_SYSTEM", "admin", "new password")
	assert := assert.New(t)
	assert.Nil(err)
}

func TestVirtualServerUpdateCredentialServerErrors(t *testing.T) {
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
				return nil, VirtualServerApi{}.UpdateCredential("12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return nil, VirtualServerApi{}.UpdateCredential("12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return nil, VirtualServerApi{}.UpdateCredential("12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return nil, VirtualServerApi{}.UpdateCredential("12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return nil, VirtualServerApi{}.UpdateCredential("12345", "OPERATING_SYSTEM", "admin", "new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestVirtualServerListCredentials(t *testing.T) {
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

	response, err := VirtualServerApi{}.ListCredentials("99944", "OPERATING_SYSTEM")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 2)

	assert.Equal(response.Credentials[0].Type, "OPERATING_SYSTEM")
	assert.Equal(response.Credentials[0].Username, "root")
	assert.Equal(response.Credentials[1].Type, "OPERATING_SYSTEM")
	assert.Equal(response.Credentials[1].Username, "user")
}

func TestVirtualServerListCredentialsPaginate(t *testing.T) {
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
					"type": "OPERATING_SYSTEM",
					"username": "root"
				}
			]
		}`)
	})
	defer teardown()

	response, err := VirtualServerApi{}.ListCredentials("99944", "OPERATING_SYSTEM", 1)
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	assert.Equal(response.Credentials[0].Type, "OPERATING_SYSTEM")
	assert.Equal(response.Credentials[0].Username, "root")
}

func TestVirtualServerListCredentialsServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.ListCredentials("99944", "OPERATING_SYSTEM")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.ListCredentials("99944", "OPERATING_SYSTEM")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
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
				return VirtualServerApi{}.ListCredentials("99944", "OPERATING_SYSTEM")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.ListCredentials("99944", "OPERATING_SYSTEM")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestVirtualServerGetCredential(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"type": "OPERATING_SYSTEM",
			"username": "root",
			"password": "password123"
		}`)
	})
	defer teardown()

	response, err := VirtualServerApi{}.GetCredential("99944", "OPERATING_SYSTEM", "root")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Type, "OPERATING_SYSTEM")
	assert.Equal(response.Username, "root")
	assert.Equal(response.Password, "password123")
}

func TestVirtualServerGetCredentialServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.GetCredential("99944", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.GetCredential("99944", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
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
				return VirtualServerApi{}.GetCredential("99944", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.GetCredential("99944", "OPERATING_SYSTEM", "root")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestVirtualServerGetDataTrafficMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"aggregation": "SUM",
				"from": "2016-10-20T09:00:00Z",
				"granularity": "DAY",
				"to": "2016-10-20T11:00:00Z"
			},
			"metrics": {
				"DATATRAFFIC_DOWN": {
					"unit": "B",
					"values": [
						{
							"timestamp": "2016-10-20T09:00:00Z",
							"value": 900
						},
						{
							"timestamp": "2016-10-20T10:00:00Z",
							"value": 2500
						}
					]
				},
				"DATATRAFFIC_UP": {
					"unit": "B",
					"values": [
						{
							"timestamp": "2016-10-20T09:00:00Z",
							"value": 90
						},
						{
							"timestamp": "2016-10-20T10:00:00Z",
							"value": 250
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	Metric, err := VirtualServerApi{}.GetDataTrafficMetrics("12345", "DAY", "SUM", "2016-10-20T09:00:00Z", "2016-10-20T11:00:00Z")
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Metric.Metadata.Aggregation, "SUM")
	assert.Equal(Metric.Metadata.From, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metadata.To, "2016-10-20T11:00:00Z")
	assert.Equal(Metric.Metadata.Granularity, "DAY")
	assert.Equal(Metric.Metric.DownPublic.Unit, "B")
	assert.Equal(Metric.Metric.DownPublic.Values[0].Value, 900)
	assert.Equal(Metric.Metric.DownPublic.Values[0].Timestamp, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metric.DownPublic.Values[1].Value, 2500)
	assert.Equal(Metric.Metric.DownPublic.Values[1].Timestamp, "2016-10-20T10:00:00Z")

	assert.Equal(Metric.Metric.UpPublic.Unit, "B")
	assert.Equal(Metric.Metric.UpPublic.Values[0].Value, 90)
	assert.Equal(Metric.Metric.UpPublic.Values[0].Timestamp, "2016-10-20T09:00:00Z")
	assert.Equal(Metric.Metric.UpPublic.Values[1].Value, 250)
	assert.Equal(Metric.Metric.UpPublic.Values[1].Timestamp, "2016-10-20T10:00:00Z")
}

func TestVirtualServerGetDataTrafficMetricsServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.GetDataTrafficMetrics("12345", "DAY", "SUM", "2016-10-20T09:00:00Z", "2016-10-20T11:00:00Z")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.GetDataTrafficMetrics("12345", "DAY", "SUM", "2016-10-20T09:00:00Z", "2016-10-20T11:00:00Z")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource '218030' was not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				return VirtualServerApi{}.GetDataTrafficMetrics("12345", "DAY", "SUM", "2016-10-20T09:00:00Z", "2016-10-20T11:00:00Z")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "404",
				ErrorMessage:  "Resource '218030' was not found",
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
				return VirtualServerApi{}.GetDataTrafficMetrics("12345", "DAY", "SUM", "2016-10-20T09:00:00Z", "2016-10-20T11:00:00Z")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.GetDataTrafficMetrics("12345", "DAY", "SUM", "2016-10-20T09:00:00Z", "2016-10-20T11:00:00Z")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestListTemplates(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			},
			"templates": [
				{
					"id": "WINDOWS_SERVER_2012_R2_STANDARD_64",
					"name": "Windows Server 2012 R2 Standard (64-bit)"
				},
				{
					"id": "CENTOS_7_64_PLESK",
					"name": "CentOS 7 (64-bit) Plesk"
				}
			]
		}`)
	})
	defer teardown()

	response, err := VirtualServerApi{}.ListTemplates("12345")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Templates), 2)

	assert.Equal(response.Templates[0].Id, "WINDOWS_SERVER_2012_R2_STANDARD_64")
	assert.Equal(response.Templates[0].Name, "Windows Server 2012 R2 Standard (64-bit)")
	assert.Equal(response.Templates[1].Id, "CENTOS_7_64_PLESK")
	assert.Equal(response.Templates[1].Name, "CentOS 7 (64-bit) Plesk")
}

func TestListTemplatesPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			},
			"templates": [
				{
					"id": "WINDOWS_SERVER_2012_R2_STANDARD_64",
					"name": "Windows Server 2012 R2 Standard (64-bit)"
				}
			]
		}`)
	})
	defer teardown()

	response, err := VirtualServerApi{}.ListTemplates("12345", 1)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Templates), 1)

	assert.Equal(response.Templates[0].Id, "WINDOWS_SERVER_2012_R2_STANDARD_64")
	assert.Equal(response.Templates[0].Name, "Windows Server 2012 R2 Standard (64-bit)")
}

func TestListTemplatesServerErrors(t *testing.T) {
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
				return VirtualServerApi{}.ListTemplates("12345")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "401",
				ErrorMessage:  "You are not authorized to view this resource.",
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
				return VirtualServerApi{}.ListTemplates("12345")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "ACCESS_DENIED",
				ErrorMessage:  "The access token is expired or invalid.",
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
				return VirtualServerApi{}.ListTemplates("12345")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "SERVER_ERROR",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
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
				return VirtualServerApi{}.ListTemplates("12345")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "TEMPORARILY_UNAVAILABLE",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
