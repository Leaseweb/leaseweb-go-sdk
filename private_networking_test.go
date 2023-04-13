package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateNetworkList(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"totalCount": 1,
				"limit": 10,
				"offset": 0
			},
			"privateNetworks": [
				{
					"equipmentCount": 4,
					"id": "811",
					"name": "default",
					"createdAt": "2015-07-16T13:06:45+0200",
					"updatedAt": "2015-07-16T13:06:45+0200"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := PrivateNetworkingApi{}.List(ctx, PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateNetworks), 1)

	PrivateNetwork := response.PrivateNetworks[0]
	assert.Equal(PrivateNetwork.EquipmentCount.String(), "4")
	assert.Equal(PrivateNetwork.Id, "811")
	assert.Equal(PrivateNetwork.Name, "default")
	assert.Equal(PrivateNetwork.CreatedAt, "2015-07-16T13:06:45+0200")
	assert.Equal(PrivateNetwork.UpdatedAt, "2015-07-16T13:06:45+0200")
}

func TestPrivateNetworkListPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			},
			"privateNetworks": [
				{
					"equipmentCount": 4,
					"id": "811",
					"name": "default",
					"createdAt": "2015-07-16T13:06:45+0200",
					"updatedAt": "2015-07-16T13:06:45+0200"
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := PrivateNetworkingApi{}.List(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateNetworks), 1)

	PrivateNetwork := response.PrivateNetworks[0]
	assert.Equal(PrivateNetwork.EquipmentCount.String(), "4")
	assert.Equal(PrivateNetwork.Id, "811")
	assert.Equal(PrivateNetwork.Name, "default")
	assert.Equal(PrivateNetwork.CreatedAt, "2015-07-16T13:06:45+0200")
	assert.Equal(PrivateNetwork.UpdatedAt, "2015-07-16T13:06:45+0200")
}

func TestPrivateNetworkListServerErrors(t *testing.T) {
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
				return PrivateNetworkingApi{}.List(ctx, PaginationOptions{})
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
				return PrivateNetworkingApi{}.List(ctx, PaginationOptions{})
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
				return PrivateNetworkingApi{}.List(ctx, PaginationOptions{})
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
				return PrivateNetworkingApi{}.List(ctx, PaginationOptions{})
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

func TestPrivateNetworkCreate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "12345",
			"name": "production",
			"createdAt": "2015-01-21T14:34:12+0000",
			"updatedAt": "2015-01-21T14:34:12+0000",
			"equipmentCount": 0,
			"servers": []
		}`)
	})
	defer teardown()

	ctx := context.Background()
	PrivateNetwork, err := PrivateNetworkingApi{}.Create(ctx, "production")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(PrivateNetwork.EquipmentCount.String(), "0")
	assert.Equal(PrivateNetwork.Id, "12345")
	assert.Equal(PrivateNetwork.Name, "production")
	assert.Equal(PrivateNetwork.CreatedAt, "2015-01-21T14:34:12+0000")
	assert.Equal(PrivateNetwork.UpdatedAt, "2015-01-21T14:34:12+0000")
	assert.Equal(len(PrivateNetwork.Servers), 0)
}

func TestPrivateNetworkCreateServerErrors(t *testing.T) {
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
				return PrivateNetworkingApi{}.Create(ctx, "production")
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
				return PrivateNetworkingApi{}.Create(ctx, "production")
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
				return PrivateNetworkingApi{}.Create(ctx, "production")
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
				return PrivateNetworkingApi{}.Create(ctx, "production")
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

func TestPrivateNetworkGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "12345",
			"name": "default",
			"createdAt": "2015-01-21T14:34:12+0000",
			"updatedAt": "2015-01-21T14:34:12+0000"
		}`)
	})
	defer teardown()

	ctx := context.Background()
	PrivateNetwork, err := PrivateNetworkingApi{}.Get(ctx, "12345")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(PrivateNetwork.Id, "12345")
	assert.Equal(PrivateNetwork.Name, "default")
	assert.Equal(PrivateNetwork.CreatedAt, "2015-01-21T14:34:12+0000")
	assert.Equal(PrivateNetwork.UpdatedAt, "2015-01-21T14:34:12+0000")
}

func TestPrivateNetworkGetServerErrors(t *testing.T) {
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
				return PrivateNetworkingApi{}.Get(ctx, "12345")
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
				return PrivateNetworkingApi{}.Get(ctx, "12345")
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
				return PrivateNetworkingApi{}.Get(ctx, "12345")
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
				return PrivateNetworkingApi{}.Get(ctx, "12345")
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

func TestPrivateNetworkUpdate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "12345",
			"name": "production",
			"createdAt": "2015-01-21T14:34:12+0000",
			"updatedAt": "2015-01-21T14:34:12+0000",
			"equipmentCount": 0,
			"servers": []
		}`)
	})
	defer teardown()

	ctx := context.Background()
	PrivateNetwork, err := PrivateNetworkingApi{}.Update(ctx, "12345", "production")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(PrivateNetwork.EquipmentCount.String(), "0")
	assert.Equal(PrivateNetwork.Id, "12345")
	assert.Equal(PrivateNetwork.Name, "production")
	assert.Equal(PrivateNetwork.CreatedAt, "2015-01-21T14:34:12+0000")
	assert.Equal(PrivateNetwork.UpdatedAt, "2015-01-21T14:34:12+0000")
	assert.Equal(len(PrivateNetwork.Servers), 0)
}

func TestPrivateNetworkUpdateServerErrors(t *testing.T) {
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
				return PrivateNetworkingApi{}.Update(ctx, "12345", "production")
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
				return PrivateNetworkingApi{}.Update(ctx, "12345", "production")
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
				return PrivateNetworkingApi{}.Update(ctx, "12345", "production")
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
				return PrivateNetworkingApi{}.Update(ctx, "12345", "production")
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

func TestPrivateNetworkDelete(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := PrivateNetworkingApi{}.Delete(ctx, "12345")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestPrivateNetworkDeleteServerErrors(t *testing.T) {
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
				return nil, PrivateNetworkingApi{}.Delete(ctx, "12345")
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
				return nil, PrivateNetworkingApi{}.Delete(ctx, "12345")
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
				return nil, PrivateNetworkingApi{}.Delete(ctx, "12345")
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
				return nil, PrivateNetworkingApi{}.Delete(ctx, "12345")
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

func TestPrivateNetworkListDhcpReservations(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 0,
				"totalCount": 1
			},
			"reservations": [
				{
					"ip": "127.0.0.1",
					"mac": "d8:87:03:52:0a:0f",
					"sticky": true
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := PrivateNetworkingApi{}.ListDhcpReservations(ctx, "12345", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.DhcpReservations), 1)

	DhcpReservation := response.DhcpReservations[0]
	assert.Equal(DhcpReservation.Ip, "127.0.0.1")
	assert.Equal(DhcpReservation.Mac, "d8:87:03:52:0a:0f")
	assert.Equal(DhcpReservation.Sticky, true)
}

func TestPrivateNetworkListDhcpReservationsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"_metadata": {
				"limit": 10,
				"offset": 1,
				"totalCount": 11
			},
			"reservations": [
				{
					"ip": "127.0.0.1",
					"mac": "d8:87:03:52:0a:0f",
					"sticky": true
				}
			]
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := PrivateNetworkingApi{}.ListDhcpReservations(ctx, "12345", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.DhcpReservations), 1)

	DhcpReservation := response.DhcpReservations[0]
	assert.Equal(DhcpReservation.Ip, "127.0.0.1")
	assert.Equal(DhcpReservation.Mac, "d8:87:03:52:0a:0f")
	assert.Equal(DhcpReservation.Sticky, true)
}

func TestPrivateNetworkListDhcpReservationsServerErrors(t *testing.T) {
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
				return PrivateNetworkingApi{}.ListDhcpReservations(ctx, "12345", PaginationOptions{})
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
				return PrivateNetworkingApi{}.ListDhcpReservations(ctx, "12345", PaginationOptions{})
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
				return PrivateNetworkingApi{}.ListDhcpReservations(ctx, "12345", PaginationOptions{})
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
				return PrivateNetworkingApi{}.ListDhcpReservations(ctx, "12345", PaginationOptions{})
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

func TestPrivateNetworkCreateDhcpReservation(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"ip": "127.0.0.1",
			"mac": "d8:87:03:52:0a:0f",
			"sticky": true
		}`)
	})
	defer teardown()

	ctx := context.Background()
	DhcpReservation, err := PrivateNetworkingApi{}.CreateDhcpReservation(ctx, "12345", "127.0.0.1", "d8:87:03:52:0a:0f", true)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(DhcpReservation.Ip, "127.0.0.1")
	assert.Equal(DhcpReservation.Mac, "d8:87:03:52:0a:0f")
	assert.Equal(DhcpReservation.Sticky, true)
}

func TestPrivateNetworkCreateDhcpReservationServerErrors(t *testing.T) {
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
				return PrivateNetworkingApi{}.CreateDhcpReservation(ctx, "12345", "127.0.0.1", "d8:87:03:52:0a:0f", true)
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
				return PrivateNetworkingApi{}.CreateDhcpReservation(ctx, "12345", "127.0.0.1", "d8:87:03:52:0a:0f", true)
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
				return PrivateNetworkingApi{}.CreateDhcpReservation(ctx, "12345", "127.0.0.1", "d8:87:03:52:0a:0f", true)
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
				return PrivateNetworkingApi{}.CreateDhcpReservation(ctx, "12345", "127.0.0.1", "d8:87:03:52:0a:0f", true)
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

func TestPrivateNetworkDeleteDhcpReservation(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := PrivateNetworkingApi{}.DeleteDhcpReservation(ctx, "12345", "127.0.0.1")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestPrivateNetworkDeleteDhcpReservationServerErrors(t *testing.T) {
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
				return nil, PrivateNetworkingApi{}.DeleteDhcpReservation(ctx, "12345", "127.0.0.1")
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
				return nil, PrivateNetworkingApi{}.DeleteDhcpReservation(ctx, "12345", "127.0.0.1")
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
				return nil, PrivateNetworkingApi{}.DeleteDhcpReservation(ctx, "12345", "127.0.0.1")
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
				return nil, PrivateNetworkingApi{}.DeleteDhcpReservation(ctx, "12345", "127.0.0.1")
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
