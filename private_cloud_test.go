package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateCloudList(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 1}, "privateClouds": [
			{
				"id": "218030",
				"customerId": "1301178860",
				"dataCenter": "AMS-01",
				"serviceOffering": "FLAT_FEE",
				"sla": "Bronze",
				"contract": {
					"id": "30000775",
					"startsAt": "2015-11-01T00:00:00+02:00",
					"endsAt": "2016-12-30T10:39:27+01:00",
					"billingCycle": 12,
					"billingFrequency": "MONTH",
					"pricePerFrequency": 0,
					"currency": "EUR"
				},
				"hardware": {
					"cpu": {
						"cores": 25
					},
					"memory": {
						"unit": "GB",
						"amount": 50
					},
					"storage": {
						"unit": "GB",
						"amount": 1
					}
				},
				"ips": [
					{
						"ip": "10.12.144.32",
						"version": 4,
						"type": "PUBLIC"
					}
				],
				"networkTraffic": {
					"type": "DATATRAFFIC",
					"trafficType": "PREMIUM",
					"datatrafficUnit": "TB",
					"datatrafficLimit": 6
				}
			}
		]}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.List(ctx, PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateClouds), 1)

	privateCloud1 := response.PrivateClouds[0]
	assert.Equal(privateCloud1.Id, "218030")
	assert.Equal(privateCloud1.CustomerId, "1301178860")
	assert.Equal(privateCloud1.DataCenter, "AMS-01")
	assert.Equal(privateCloud1.ServiceOffering, "FLAT_FEE")
	assert.Equal(privateCloud1.Sla, "Bronze")
	assert.Equal(privateCloud1.Contract.Id, "30000775")
	assert.Equal(privateCloud1.Contract.StartsAt, "2015-11-01T00:00:00+02:00")
	assert.Equal(privateCloud1.Contract.EndsAt, "2016-12-30T10:39:27+01:00")
	assert.Equal(privateCloud1.Contract.BillingCycle.String(), "12")
	assert.Equal(privateCloud1.Contract.BillingFrequency, "MONTH")
	assert.Equal(privateCloud1.Contract.PricePerFrequency.String(), "0")
	assert.Equal(privateCloud1.Contract.Currency, "EUR")
	assert.Equal(privateCloud1.Hardware.Cpu.Cores.String(), "25")
	assert.Equal(privateCloud1.Hardware.Memory.Amount.String(), "50")
	assert.Equal(privateCloud1.Hardware.Memory.Unit, "GB")
	assert.Equal(privateCloud1.Hardware.Storage.Amount.String(), "1")
	assert.Equal(privateCloud1.Hardware.Storage.Unit, "GB")
	assert.Equal(privateCloud1.Ips[0].Ip, "10.12.144.32")
	assert.Equal(privateCloud1.Ips[0].Type, "PUBLIC")
	assert.Equal(privateCloud1.Ips[0].Version.String(), "4")
	assert.Equal(privateCloud1.NetworkTraffic.DataTrafficLimit.String(), "6")
	assert.Equal(privateCloud1.NetworkTraffic.DataTrafficUnit, "TB")
	assert.Equal(privateCloud1.NetworkTraffic.TrafficType, "PREMIUM")
	assert.Equal(privateCloud1.NetworkTraffic.Type, "DATATRAFFIC")
}

func TestPrivateCloudListBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "privateClouds": []}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.List(ctx, PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateClouds), 0)
}

func TestPrivateCloudListPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 20, "offset": 10, "totalCount": 2}, "privateClouds": []}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	opts := PaginationOptions{
		Limit:  Int(10),
		Offset: Int(20),
	}
	response, err := privateCloudApi.List(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 10)
	assert.Equal(response.Metadata.Limit, 20)
	assert.Equal(len(response.PrivateClouds), 0)
}

func TestPrivateCloudListServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.List(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.List(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.List(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "218030",
			"customerId": "1301178860",
			"dataCenter": "AMS-01",
			"serviceOffering": "FLAT_FEE",
			"sla": "Bronze",
			"contract": {
				"id": "30000775",
				"startsAt": "2015-11-01T00:00:00+02:00",
				"endsAt": "2016-12-30T10:39:27+01:00",
				"billingCycle": 12,
				"billingFrequency": "MONTH",
				"pricePerFrequency": 0,
				"currency": "EUR"
			},
			"hardware": {
				"cpu": {
					"cores": 25
				},
				"memory": {
					"unit": "GB",
					"amount": 50
				},
				"storage": {
					"unit": "GB",
					"amount": 1
				}
			},
			"ips": [
				{
					"ip": "10.12.144.32",
					"version": 4,
					"type": "PUBLIC"
				}
			],
			"networkTraffic": {
				"type": "DATATRAFFIC",
				"trafficType": "PREMIUM",
				"datatrafficUnit": "TB",
				"datatrafficLimit": 6
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.Get(ctx, "218030")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "218030")
	assert.Equal(response.CustomerId, "1301178860")
	assert.Equal(response.DataCenter, "AMS-01")
	assert.Equal(response.ServiceOffering, "FLAT_FEE")
	assert.Equal(response.Sla, "Bronze")
	assert.Equal(response.Contract.Id, "30000775")
	assert.Equal(response.Contract.StartsAt, "2015-11-01T00:00:00+02:00")
	assert.Equal(response.Contract.EndsAt, "2016-12-30T10:39:27+01:00")
	assert.Equal(response.Contract.BillingCycle.String(), "12")
	assert.Equal(response.Contract.BillingFrequency, "MONTH")
	assert.Equal(response.Contract.PricePerFrequency.String(), "0")
	assert.Equal(response.Contract.Currency, "EUR")
	assert.Equal(response.Hardware.Cpu.Cores.String(), "25")
	assert.Equal(response.Hardware.Memory.Amount.String(), "50")
	assert.Equal(response.Hardware.Memory.Unit, "GB")
	assert.Equal(response.Hardware.Storage.Amount.String(), "1")
	assert.Equal(response.Hardware.Storage.Unit, "GB")
	assert.Equal(response.Ips[0].Ip, "10.12.144.32")
	assert.Equal(response.Ips[0].Type, "PUBLIC")
	assert.Equal(response.Ips[0].Version.String(), "4")
	assert.Equal(response.NetworkTraffic.DataTrafficLimit.String(), "6")
	assert.Equal(response.NetworkTraffic.DataTrafficUnit, "TB")
	assert.Equal(response.NetworkTraffic.TrafficType, "PREMIUM")
	assert.Equal(response.NetworkTraffic.Type, "DATATRAFFIC")
}

func TestPrivateCloudGetServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.Get(ctx, "218030")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.Get(ctx, "218030")
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.Get(ctx, "218030")
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.Get(ctx, "218030")
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudListCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 1}, "credentials": [
			{
				"type": "REMOTE_MANAGEMENT",
				"username": "root",
				"domain": "123456"
			}
		]}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.ListCredentials(ctx, "12345678", "REMOTE_MANAGEMENT", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	credential := response.Credentials[0]
	assert.Equal(credential.Type, "REMOTE_MANAGEMENT")
	assert.Equal(credential.Username, "root")
	assert.Equal(credential.Domain, "123456")
}

func TestPrivateCloudListCredentialsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.ListCredentials(ctx, "12345678", "REMOTE_MANAGEMENT", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestPrivateCloudListCredentialsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "credentials": [
			{
				"type": "REMOTE_MANAGEMENT",
				"username": "root",
				"password": "password123",
				"domain": "123456"
			}
		]}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	opts := PaginationOptions{
		Limit:  Int(1),
		Offset: Int(10),
	}
	response, err := privateCloudApi.ListCredentials(ctx, "12345678", "REMOTE_MANAGEMENT", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	credential := response.Credentials[0]
	assert.Equal(credential.Type, "REMOTE_MANAGEMENT")
	assert.Equal(credential.Username, "root")
	assert.Equal(credential.Password, "password123")
	assert.Equal(credential.Domain, "123456")
}

func TestPrivateCloudListCredentialsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.ListCredentials(ctx, "12345678", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.ListCredentials(ctx, "12345678", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.ListCredentials(ctx, "12345678", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.ListCredentials(ctx, "12345678", "REMOTE_MANAGEMENT", PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudGetCredential(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"type": "REMOTE_MANAGEMENT",
			"username": "root",
			"password": "password123",
			"domain": "123456"
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.GetCredential(ctx, "218030", "REMOTE_MANAGEMENT", "root")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Username, "root")
	assert.Equal(response.Password, "password123")
	assert.Equal(response.Domain, "123456")
}

func TestPrivateCloudGetCredentialServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCredential(ctx, "218030", "REMOTE_MANAGEMENT", "root")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCredential(ctx, "218030", "REMOTE_MANAGEMENT", "root")
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCredential(ctx, "218030", "REMOTE_MANAGEMENT", "root")
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCredential(ctx, "218030", "REMOTE_MANAGEMENT", "root")
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudGetDataTrafficMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "SUM"
			},
			"metrics": {
				"DATATRAFFIC_UP": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				},
				"DATATRAFFIC_DOWN": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 90
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 250
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.GetDataTrafficMetrics(ctx, "218030", MetricsOptions{})

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "SUM")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DownPublic.Unit, "GB")
	assert.Equal(response.Metric.DownPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[0].Value.String(), "90")
	assert.Equal(response.Metric.DownPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[1].Value.String(), "250")
	assert.Equal(response.Metric.UpPublic.Unit, "GB")
	assert.Equal(response.Metric.UpPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[0].Value.String(), "900")
	assert.Equal(response.Metric.UpPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[1].Value.String(), "2500")
}

func TestPrivateCloudGetDataTrafficMetricsWithFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "SUM"
			},
			"metrics": {
				"DATATRAFFIC_UP": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				},
				"DATATRAFFIC_DOWN": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 90
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 250
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	opts := MetricsOptions{
		Granularity: String("MONTH"),
		Aggregation: String("SUM"),
		From:        String("2017-07-01T00:00:00+00:00"),
		To:          String("2017-07-02T00:00:00+00:00"),
	}
	response, err := privateCloudApi.GetDataTrafficMetrics(ctx, "218030", opts)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "SUM")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DownPublic.Unit, "GB")
	assert.Equal(response.Metric.DownPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[0].Value.String(), "90")
	assert.Equal(response.Metric.DownPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[1].Value.String(), "250")
	assert.Equal(response.Metric.UpPublic.Unit, "GB")
	assert.Equal(response.Metric.UpPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[0].Value.String(), "900")
	assert.Equal(response.Metric.UpPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[1].Value.String(), "2500")
}

func TestPrivateCloudGetDataTrafficMetricsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetDataTrafficMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetDataTrafficMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetDataTrafficMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetDataTrafficMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudGetBandWidthMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "AVG"
			},
			"metrics": {
				"DOWN_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 28202556
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 28202557
						}
					]
				},
				"UP_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 158317518
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 158317519
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.GetBandWidthMetrics(ctx, "218030", MetricsOptions{})

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "AVG")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DownPublic.Unit, "bps")
	assert.Equal(response.Metric.DownPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[0].Value.String(), "28202556")
	assert.Equal(response.Metric.DownPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[1].Value.String(), "28202557")
	assert.Equal(response.Metric.UpPublic.Unit, "bps")
	assert.Equal(response.Metric.UpPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[0].Value.String(), "158317518")
	assert.Equal(response.Metric.UpPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[1].Value.String(), "158317519")
}

func TestPrivateCloudGetBandWidthMetricsWithFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "AVG"
			},
			"metrics": {
				"DOWN_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 28202556
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 28202557
						}
					]
				},
				"UP_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 158317518
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 158317519
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	opts := MetricsOptions{
		Granularity: String("MONTH"),
		Aggregation: String("AVG"),
		From:        String("2017-07-01T00:00:00+00:00"),
		To:          String("2017-07-02T00:00:00+00:00"),
	}
	response, err := privateCloudApi.GetBandWidthMetrics(ctx, "218030", opts)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "AVG")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DownPublic.Unit, "bps")
	assert.Equal(response.Metric.DownPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[0].Value.String(), "28202556")
	assert.Equal(response.Metric.DownPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[1].Value.String(), "28202557")
	assert.Equal(response.Metric.UpPublic.Unit, "bps")
	assert.Equal(response.Metric.UpPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[0].Value.String(), "158317518")
	assert.Equal(response.Metric.UpPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[1].Value.String(), "158317519")
}

func TestPrivateCloudGetBandWidthMetricsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetBandWidthMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetBandWidthMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetBandWidthMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetBandWidthMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudGetCpuMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-01T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"CPU": {
					"unit": "CORES",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 24
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 24
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.GetCpuMetrics(ctx, "218030", MetricsOptions{})

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Cpu.Unit, "CORES")
	assert.Equal(response.Metric.Cpu.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[0].Value.String(), "24")
	assert.Equal(response.Metric.Cpu.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[1].Value.String(), "24")
}

func TestPrivateCloudGetCpuMetricsWithFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"CPU": {
					"unit": "CORES",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 24
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 24
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	opts := MetricsOptions{
		Granularity: String("MONTH"),
		Aggregation: String("MAX"),
		From:        String("2017-07-01T00:00:00+00:00"),
		To:          String("2017-07-02T00:00:00+00:00"),
	}
	response, err := privateCloudApi.GetCpuMetrics(ctx, "218030", opts)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Cpu.Unit, "CORES")
	assert.Equal(response.Metric.Cpu.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[0].Value.String(), "24")
	assert.Equal(response.Metric.Cpu.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[1].Value.String(), "24")
}

func TestPrivateCloudGetCpuMetricsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCpuMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCpuMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCpuMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetCpuMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudGetMemoryMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"MEMORY": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 8
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 16
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.GetMemoryMetrics(ctx, "218030", MetricsOptions{})

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Memory.Unit, "GB")
	assert.Equal(response.Metric.Memory.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[0].Value.String(), "8")
	assert.Equal(response.Metric.Memory.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[1].Value.String(), "16")
}

func TestPrivateCloudGetMemoryMetricsWithFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"MEMORY": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 8
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 16
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	opts := MetricsOptions{
		Granularity: String("MONTH"),
		Aggregation: String("MAX"),
		From:        String("2017-07-01T00:00:00+00:00"),
		To:          String("2017-07-02T00:00:00+00:00"),
	}
	response, err := privateCloudApi.GetMemoryMetrics(ctx, "218030", opts)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Memory.Unit, "GB")
	assert.Equal(response.Metric.Memory.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[0].Value.String(), "8")
	assert.Equal(response.Metric.Memory.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[1].Value.String(), "16")
}

func TestPrivateCloudGetMemoryMetricsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetMemoryMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetMemoryMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetMemoryMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetMemoryMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestPrivateCloudGetStorageMetrics(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"STORAGE": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	response, err := privateCloudApi.GetStorageMetrics(ctx, "218030", MetricsOptions{})

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Storage.Unit, "GB")
	assert.Equal(response.Metric.Storage.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[0].Value.String(), "900")
	assert.Equal(response.Metric.Storage.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[1].Value.String(), "2500")
}

func TestPrivateCloudGetStorageMetricsWithFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"STORAGE": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	ctx := context.Background()
	opts := MetricsOptions{
		Granularity: String("MONTH"),
		Aggregation: String("MAX"),
		From:        String("2017-07-01T00:00:00+00:00"),
		To:          String("2017-07-02T00:00:00+00:00"),
	}
	response, err := privateCloudApi.GetStorageMetrics(ctx, "218030", opts)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Storage.Unit, "GB")
	assert.Equal(response.Metric.Storage.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[0].Value.String(), "900")
	assert.Equal(response.Metric.Storage.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[1].Value.String(), "2500")
}

func TestPrivateCloudGetStorageMetricsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetStorageMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetStorageMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:        "404",
				Message:     "Resource 218030 was not found",
				UserMessage: "Resource with id 218030 not found.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetStorageMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return PrivateCloudApi{}.GetStorageMetrics(ctx, "218030", MetricsOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
