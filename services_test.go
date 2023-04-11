package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServicesList(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 2
			},
			"services": [
				{
					"billingCycle": "1 MONTH",
					"cancellable": false,
					"contractId": "00000110",
					"contractTerm": "1 YEAR",
					"contractTermEndDate": "2020-01-31T00:00:00+00:00",
					"currency": "EUR",
					"deliveryDate": "2019-01-01T00:00:00+00:00",
					"deliveryEstimate": "5 - 7 business days",
					"endDate": "2020-01-31T00:00:00+00:00",
					"equipmentId": "12345678",
					"id": "10000000000010",
					"orderDate": "2019-01-01T00:00:00+00:00",
					"pricePerFrequency": 396.01,
					"productId": "DEDICATED_SERVER",
					"reference": "this is a reference",
					"startDate": "2019-01-01T00:00:00+00:00",
					"status": "ACTIVE",
					"uncancellable": true
				},
				{
					"billingCycle": "1 YEAR",
					"cancellable": true,
					"contractId": "00000110",
					"contractTerm": "1 YEAR",
					"contractTermEndDate": "2020-01-31T00:00:00+00:00",
					"currency": "EUR",
					"deliveryDate": "2019-01-01T00:00:00+00:00",
					"deliveryEstimate": "5 - 7 business days",
					"id": "10000000000011",
					"orderDate": "2019-01-01T00:00:00+00:00",
					"pricePerFrequency": 139.99,
					"productId": "DOMAIN",
					"startDate": "2019-01-01T00:00:00+00:00",
					"status": "ACTIVE",
					"uncancellable": false
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := ServicesApi{}.List(ctx, ServicesListOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Services), 2)

	service1 := response.Services[0]
	assert.Equal(service1.BillingCycle, "1 MONTH")
	assert.Equal(service1.Cancellable, false)
	assert.Equal(service1.ContractId, "00000110")
	assert.Equal(service1.ContractTerm, "1 YEAR")
	assert.Equal(service1.ContractTermEndDate, "2020-01-31T00:00:00+00:00")
	assert.Equal(service1.Currency, "EUR")
	assert.Equal(service1.DeliveryDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service1.DeliveryEstimate, "5 - 7 business days")
	assert.Equal(service1.EndDate, "2020-01-31T00:00:00+00:00")
	assert.Equal(service1.EquipmentId, "12345678")
	assert.Equal(service1.Id, "10000000000010")
	assert.Equal(service1.OrderDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service1.PricePerFrequency.String(), "396.01")
	assert.Equal(service1.ProductId, "DEDICATED_SERVER")
	assert.Equal(service1.Reference, "this is a reference")
	assert.Equal(service1.StartDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service1.Status, "ACTIVE")
	assert.Equal(service1.Uncancellable, true)

	service2 := response.Services[1]
	assert.Equal(service2.BillingCycle, "1 YEAR")
	assert.Equal(service2.Cancellable, true)
	assert.Equal(service2.ContractId, "00000110")
	assert.Equal(service2.ContractTerm, "1 YEAR")
	assert.Equal(service2.ContractTermEndDate, "2020-01-31T00:00:00+00:00")
	assert.Equal(service2.Currency, "EUR")
	assert.Equal(service2.DeliveryDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service2.DeliveryEstimate, "5 - 7 business days")
	assert.Equal(service2.Id, "10000000000011")
	assert.Equal(service2.OrderDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service2.PricePerFrequency.String(), "139.99")
	assert.Equal(service2.ProductId, "DOMAIN")
	assert.Equal(service2.StartDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service2.Status, "ACTIVE")
	assert.Equal(service2.Uncancellable, false)
}

func TestServicesListPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"services": [
				{
					"billingCycle": "1 MONTH",
					"cancellable": false,
					"contractId": "00000110",
					"contractTerm": "1 YEAR",
					"contractTermEndDate": "2020-01-31T00:00:00+00:00",
					"currency": "EUR",
					"deliveryDate": "2019-01-01T00:00:00+00:00",
					"deliveryEstimate": "5 - 7 business days",
					"endDate": "2020-01-31T00:00:00+00:00",
					"equipmentId": "12345678",
					"id": "10000000000010",
					"orderDate": "2019-01-01T00:00:00+00:00",
					"pricePerFrequency": 396.01,
					"productId": "DEDICATED_SERVER",
					"reference": "this is a reference",
					"startDate": "2019-01-01T00:00:00+00:00",
					"status": "ACTIVE",
					"uncancellable": true
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := ServicesListOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(1),
		},
	}
	response, err := ServicesApi{}.List(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Services), 1)

	service1 := response.Services[0]
	assert.Equal(service1.BillingCycle, "1 MONTH")
	assert.Equal(service1.Cancellable, false)
	assert.Equal(service1.ContractId, "00000110")
	assert.Equal(service1.ContractTerm, "1 YEAR")
	assert.Equal(service1.ContractTermEndDate, "2020-01-31T00:00:00+00:00")
	assert.Equal(service1.Currency, "EUR")
	assert.Equal(service1.DeliveryDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service1.DeliveryEstimate, "5 - 7 business days")
	assert.Equal(service1.EndDate, "2020-01-31T00:00:00+00:00")
	assert.Equal(service1.EquipmentId, "12345678")
	assert.Equal(service1.Id, "10000000000010")
	assert.Equal(service1.OrderDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service1.PricePerFrequency.String(), "396.01")
	assert.Equal(service1.ProductId, "DEDICATED_SERVER")
	assert.Equal(service1.Reference, "this is a reference")
	assert.Equal(service1.StartDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(service1.Status, "ACTIVE")
	assert.Equal(service1.Uncancellable, true)
}

func TestServicesListServerErrors(t *testing.T) {
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
				return ServicesApi{}.List(ctx, ServicesListOptions{})
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
				return ServicesApi{}.List(ctx, ServicesListOptions{})
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
				return ServicesApi{}.List(ctx, ServicesListOptions{})
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
				return ServicesApi{}.List(ctx, ServicesListOptions{})
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

func TestServicesListCancellationReasons(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"cancellationReasons": [
				{
					"reason": "I no longer need it",
					"reasonCode": "CANCEL_NO_NEED"
				},
				{
					"reason": "I was using it for trial only",
					"reasonCode": "CANCEL_TRIAL_ONLY"
				},
				{
					"reason": "I purchased another service at Leaseweb",
					"reasonCode": "CANCEL_PURCHASED_OTHER"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := ServicesApi{}.ListCancellationReasons(ctx)

	assert := assert.New(t)
	assert.Nil(err)
	Reason1 := response.CancellationReasons[0]
	assert.Equal(Reason1.Reason, "I no longer need it")
	assert.Equal(Reason1.ReasonCode, "CANCEL_NO_NEED")
	Reason2 := response.CancellationReasons[1]
	assert.Equal(Reason2.Reason, "I was using it for trial only")
	assert.Equal(Reason2.ReasonCode, "CANCEL_TRIAL_ONLY")
	Reason3 := response.CancellationReasons[2]
	assert.Equal(Reason3.Reason, "I purchased another service at Leaseweb")
	assert.Equal(Reason3.ReasonCode, "CANCEL_PURCHASED_OTHER")
}

func TestServicesListCancellationReasonsServerErrors(t *testing.T) {
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
				return ServicesApi{}.ListCancellationReasons(ctx)
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
				return ServicesApi{}.ListCancellationReasons(ctx)
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
				return ServicesApi{}.ListCancellationReasons(ctx)
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
				return ServicesApi{}.ListCancellationReasons(ctx)
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

func TestServicesGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"billingCycle": "1 MONTH",
			"cancellable": false,
			"contractId": "00000110",
			"contractTerm": "1 YEAR",
			"contractTermEndDate": "2020-01-31T00:00:00+00:00",
			"currency": "EUR",
			"deliveryDate": "2019-01-01T00:00:00+00:00",
			"deliveryEstimate": "5 - 7 business days",
			"endDate": "2020-01-31T00:00:00+00:00",
			"equipmentId": "12345678",
			"id": "10000000000010",
			"orderDate": "2019-01-01T00:00:00+00:00",
			"pricePerFrequency": 396.01,
			"productId": "DEDICATED_SERVER",
			"reference": "this is a reference",
			"startDate": "2019-01-01T00:00:00+00:00",
			"status": "ACTIVE",
			"uncancellable": true
		}`)
	})
	defer teardown()

	ctx := context.Background()
	Service, err := ServicesApi{}.Get(ctx, "12345")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(Service.BillingCycle, "1 MONTH")
	assert.Equal(Service.Cancellable, false)
	assert.Equal(Service.ContractId, "00000110")
	assert.Equal(Service.ContractTerm, "1 YEAR")
	assert.Equal(Service.ContractTermEndDate, "2020-01-31T00:00:00+00:00")
	assert.Equal(Service.Currency, "EUR")
	assert.Equal(Service.DeliveryDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(Service.DeliveryEstimate, "5 - 7 business days")
	assert.Equal(Service.EndDate, "2020-01-31T00:00:00+00:00")
	assert.Equal(Service.EquipmentId, "12345678")
	assert.Equal(Service.Id, "10000000000010")
	assert.Equal(Service.OrderDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(Service.PricePerFrequency.String(), "396.01")
	assert.Equal(Service.ProductId, "DEDICATED_SERVER")
	assert.Equal(Service.Reference, "this is a reference")
	assert.Equal(Service.StartDate, "2019-01-01T00:00:00+00:00")
	assert.Equal(Service.Status, "ACTIVE")
	assert.Equal(Service.Uncancellable, true)
}

func TestServicesGetServerErrors(t *testing.T) {
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
				return ServicesApi{}.Get(ctx, "12345")
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
				return ServicesApi{}.Get(ctx, "12345")
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
				return ServicesApi{}.Get(ctx, "12345")
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
				return ServicesApi{}.Get(ctx, "12345")
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

func TestServicesCancel(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := ServicesApi{}.Cancel(ctx, "12345", "reason", "reason code")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestServicesCancelServerErrors(t *testing.T) {
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
				return nil, ServicesApi{}.Cancel(ctx, "12345", "reason", "reason code")
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
				ctx := context.Background()
				return nil, ServicesApi{}.Cancel(ctx, "12345", "reason", "reason code")
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
				ctx := context.Background()
				return nil, ServicesApi{}.Cancel(ctx, "12345", "reason", "reason code")
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
				ctx := context.Background()
				return nil, ServicesApi{}.Cancel(ctx, "12345", "reason", "reason code")
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

func TestServicesUncancel(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := ServicesApi{}.Uncancel(ctx, "12345")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestServicesUncancelServerErrors(t *testing.T) {
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
				return nil, ServicesApi{}.Uncancel(ctx, "12345")
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
				ctx := context.Background()
				return nil, ServicesApi{}.Uncancel(ctx, "12345")
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
				ctx := context.Background()
				return nil, ServicesApi{}.Uncancel(ctx, "12345")
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
				ctx := context.Background()
				return nil, ServicesApi{}.Uncancel(ctx, "12345")
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
