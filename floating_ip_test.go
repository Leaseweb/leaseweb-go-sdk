package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatingIpListRanges(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "ranges": [
			{
				"id": "85.17.0.0_17",
				"range": "85.17.0.0/17",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"location": "AMS-01",
				"type": "SITE"
			},
			{
				"id": "86.17.0.1_17",
				"range": "86.17.0.1/17",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"location": "AMS",
				"type": "METRO"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	response, err := floatingIpApi.ListRanges(ctx, FloatingIpListRangesOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ranges), 2)

	range1 := response.Ranges[0]
	assert.Equal(range1.Id, "85.17.0.0_17")
	assert.Equal(range1.Range, "85.17.0.0/17")
	assert.Equal(range1.CustomerId, "10001234")
	assert.Equal(range1.SalesOrgId, "2000")
	assert.Equal(range1.Location, "AMS-01")
	assert.Equal(range1.Type, "SITE")

	range2 := response.Ranges[1]
	assert.Equal(range2.Id, "86.17.0.1_17")
	assert.Equal(range2.Range, "86.17.0.1/17")
	assert.Equal(range2.CustomerId, "10001234")
	assert.Equal(range2.SalesOrgId, "2000")
	assert.Equal(range2.Location, "AMS")
	assert.Equal(range2.Type, "METRO")
}

func TestFloatingIpListRangesPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "ranges": [
			{
				"id": "85.17.0.0_17",
				"range": "85.17.0.0/17",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"location": "AMS-01",
				"type": "SITE"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	opts := FloatingIpListRangesOptions{
		PaginationOptions: PaginationOptions{
			Limit:  Int(1),
			Offset: Int(10),
		},
		Type:     String("SITE, METRO"),
		Location: String("AMS-01"),
	}
	response, err := floatingIpApi.ListRanges(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ranges), 1)

	range1 := response.Ranges[0]
	assert.Equal(range1.Id, "85.17.0.0_17")
	assert.Equal(range1.Range, "85.17.0.0/17")
	assert.Equal(range1.CustomerId, "10001234")
	assert.Equal(range1.SalesOrgId, "2000")
	assert.Equal(range1.Location, "AMS-01")
	assert.Equal(range1.Type, "SITE")
}

func TestFloatingIpListRangesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.ListRanges(ctx, FloatingIpListRangesOptions{})
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
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.ListRanges(ctx, FloatingIpListRangesOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.ListRanges(ctx, FloatingIpListRangesOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestFloatingIpGetRange(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "85.17.0.0_17",
			"range": "88.17.0.0/17",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"location": "AMS-01",
			"type": "SITE"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	response, err := floatingIpApi.GetRange(ctx, "123456789")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "85.17.0.0_17")
	assert.Equal(response.Range, "88.17.0.0/17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
}

func TestFloatingIpGetRangeServerErrors(t *testing.T) {
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
				return FloatingIpApi{}.GetRange(ctx, "123456789")
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
				fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Range with id 88.17.0.0_17 does not exist"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.GetRange(ctx, "123456789")
			},
			ExpectedError: ApiError{
				CorrelationId: "39e010ed-0e93-42c3-c28f-3ffc373553d5",
				Code:          "404",
				Message:       "Range with id 88.17.0.0_17 does not exist",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.GetRange(ctx, "123456789")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.GetRange(ctx, "123456789")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestFloatingIpListRangeDefinitions(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "floatingIpDefinitions": [
			{
				"id": "88.17.34.108_32",
				"rangeId": "88.17.0.0_17",
				"location": "AMS-01",
				"type": "SITE",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"floatingIp": "88.17.34.108/32",
				"anchorIp": "95.10.126.1",
				"status": "ACTIVE",
				"createdAt": "2019-03-13T09:10:02+0000",
				"updatedAt": "2019-03-13T09:10:02+0000"
			},
			{
				"id": "88.17.34.109_32",
				"rangeId": "88.17.0.0_17",
				"location": "AMS-01",
				"type": "SITE",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"floatingIp": "88.17.34.109/32",
				"anchorIp": "95.10.126.12",
				"status": "ACTIVE",
				"createdAt": "2019-03-13T09:10:02+0000",
				"updatedAt": "2019-03-13T09:10:02+0000"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	response, err := floatingIpApi.ListRangeDefinitions(ctx, "123456789", FloatingIpListRangeDefinitionsOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Definitions), 2)

	floatingIpDefinition1 := response.Definitions[0]
	assert.Equal(floatingIpDefinition1.Id, "88.17.34.108_32")
	assert.Equal(floatingIpDefinition1.RangeId, "88.17.0.0_17")
	assert.Equal(floatingIpDefinition1.CustomerId, "10001234")
	assert.Equal(floatingIpDefinition1.SalesOrgId, "2000")
	assert.Equal(floatingIpDefinition1.Location, "AMS-01")
	assert.Equal(floatingIpDefinition1.Type, "SITE")
	assert.Equal(floatingIpDefinition1.FloatingIp, "88.17.34.108/32")
	assert.Equal(floatingIpDefinition1.AnchorIp, "95.10.126.1")
	assert.Equal(floatingIpDefinition1.Status, "ACTIVE")
	assert.Equal(floatingIpDefinition1.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(floatingIpDefinition1.UpdatedAt, "2019-03-13T09:10:02+0000")

	floatingIpDefinition2 := response.Definitions[1]
	assert.Equal(floatingIpDefinition2.Id, "88.17.34.109_32")
	assert.Equal(floatingIpDefinition2.RangeId, "88.17.0.0_17")
	assert.Equal(floatingIpDefinition2.CustomerId, "10001234")
	assert.Equal(floatingIpDefinition2.SalesOrgId, "2000")
	assert.Equal(floatingIpDefinition2.Location, "AMS-01")
	assert.Equal(floatingIpDefinition2.Type, "SITE")
	assert.Equal(floatingIpDefinition2.FloatingIp, "88.17.34.109/32")
	assert.Equal(floatingIpDefinition2.AnchorIp, "95.10.126.12")
	assert.Equal(floatingIpDefinition2.Status, "ACTIVE")
	assert.Equal(floatingIpDefinition2.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(floatingIpDefinition2.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestFloatingIpListRangeDefinitionsPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "floatingIpDefinitions": [
			{
				"id": "88.17.34.108_32",
				"rangeId": "88.17.0.0_17",
				"location": "AMS-01",
				"type": "SITE",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"floatingIp": "88.17.34.108/32",
				"anchorIp": "95.10.126.1",
				"status": "ACTIVE",
				"createdAt": "2019-03-13T09:10:02+0000",
				"updatedAt": "2019-03-13T09:10:02+0000"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	opts := FloatingIpListRangeDefinitionsOptions{
		PaginationOptions: PaginationOptions{
			Limit:  Int(1),
			Offset: Int(10),
		},
		Location: String("AMS-01"),
		Type:     []string{"SITE", "METRO"},
	}
	response, err := floatingIpApi.ListRangeDefinitions(ctx, "123456789", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)

	floatingIpDefinition1 := response.Definitions[0]
	assert.Equal(floatingIpDefinition1.Id, "88.17.34.108_32")
	assert.Equal(floatingIpDefinition1.RangeId, "88.17.0.0_17")
	assert.Equal(floatingIpDefinition1.CustomerId, "10001234")
	assert.Equal(floatingIpDefinition1.SalesOrgId, "2000")
	assert.Equal(floatingIpDefinition1.Location, "AMS-01")
	assert.Equal(floatingIpDefinition1.Type, "SITE")
	assert.Equal(floatingIpDefinition1.FloatingIp, "88.17.34.108/32")
	assert.Equal(floatingIpDefinition1.AnchorIp, "95.10.126.1")
	assert.Equal(floatingIpDefinition1.Status, "ACTIVE")
	assert.Equal(floatingIpDefinition1.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(floatingIpDefinition1.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestFloatingIpListRangeDefinitionsServerErrors(t *testing.T) {
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
				return FloatingIpApi{}.ListRangeDefinitions(ctx, "123456789", FloatingIpListRangeDefinitionsOptions{})
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
				fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Range with id 88.17.0.0_17 does not exist"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.ListRangeDefinitions(ctx, "123456789", FloatingIpListRangeDefinitionsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "39e010ed-0e93-42c3-c28f-3ffc373553d5",
				Code:          "404",
				Message:       "Range with id 88.17.0.0_17 does not exist",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.ListRangeDefinitions(ctx, "123456789", FloatingIpListRangeDefinitionsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.ListRangeDefinitions(ctx, "123456789", FloatingIpListRangeDefinitionsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestFloatingIpCreateRangeDefinition(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	response, err := floatingIpApi.CreateRangeDefinition(ctx, "10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestFloatingIpCreateRangeDefinitionServerError(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 400",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, `{"errorCode": "400", "errorMessage": "Validation Failed"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.CreateRangeDefinition(ctx, "10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				Code:    "400",
				Message: "Validation Failed",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.CreateRangeDefinition(ctx, "10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.CreateRangeDefinition(ctx, "10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.CreateRangeDefinition(ctx, "10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestFloatingIpGetRangeDefinition(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	response, err := floatingIpApi.GetRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestFloatingIpGetRangeDefinitionServerErrors(t *testing.T) {
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
				return FloatingIpApi{}.GetRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")
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
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.GetRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.GetRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestFloatingIpUpdateRangeDefinition(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	response, err := floatingIpApi.UpdateRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32", "95.10.126.1")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestFloatingIpUpdateRangeDefinitionServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 400",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, `{"correlationId": "945bef2e-1caf-4027-bd0a-8976848f3dee", "errorCode": "400", "errorMessage": "Validation Failed"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.UpdateRangeDefinition(ctx, "wrong 1", "88.17.34.108_32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "945bef2e-1caf-4027-bd0a-8976848f3dee",
				Code:          "400",
				Message:       "Validation Failed",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.UpdateRangeDefinition(ctx, "wrong 1", "88.17.34.108_32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.UpdateRangeDefinition(ctx, "wrong 1", "88.17.34.108_32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.UpdateRangeDefinition(ctx, "wrong 1", "88.17.34.108_32", "95.10.126.1")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestFloatingIpRemoveRangeDefinition(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	ctx := context.Background()
	response, err := floatingIpApi.RemoveRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestFloatingIpRemoveRangeDefinitionServerErrors(t *testing.T) {

	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.RemoveRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.RemoveRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return FloatingIpApi{}.RemoveRangeDefinition(ctx, "88.17.0.0_17", "88.17.34.108_32")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
