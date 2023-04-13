package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteManagementChangeCredentials(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	ctx := context.Background()
	err := remoteManagementApi.ChangeCredentials(ctx, "new password")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestRemoteManagementTestChangeCredentialsServerErrors(t *testing.T) {
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
				return nil, RemoteManagementApi{}.ChangeCredentials(ctx, "new password")
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
				return nil, RemoteManagementApi{}.ChangeCredentials(ctx, "new password")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, RemoteManagementApi{}.ChangeCredentials(ctx, "new password")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, RemoteManagementApi{}.ChangeCredentials(ctx, "new password")
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

func TestRemoteManagementListProfiles(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "profiles": [
			{
				"datacenter": "AMS-01",
				"satelliteDatacenters": [
					"AMS-10",
					"AMS-11",
					"AMS-12"
				],
				"file": "https://api.leaseweb.com/bareMetals/v2/remoteManagement/profiles/lsw-rmvpn-AMS-01.ovpn"
			},
			{
				"datacenter": "AMS-02",
				"satelliteDatacenters": [
					"AMS-13",
					"AMS-15"
				],
				"file": "https://api.leaseweb.com/bareMetals/v2/remoteManagement/profiles/lsw-rmvpn-AMS-02.ovpn"
			}
		]}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	ctx := context.Background()
	response, err := remoteManagementApi.ListProfiles(ctx, PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Profiles), 2)

	profile1 := response.Profiles[0]
	assert.Equal(profile1.DataCenter, "AMS-01")
	assert.Equal(profile1.File, "https://api.leaseweb.com/bareMetals/v2/remoteManagement/profiles/lsw-rmvpn-AMS-01.ovpn")
	assert.Equal(len(profile1.SatelliteDataCenters), 3)
	assert.Equal(profile1.SatelliteDataCenters[0], "AMS-10")
	assert.Equal(profile1.SatelliteDataCenters[1], "AMS-11")
	assert.Equal(profile1.SatelliteDataCenters[2], "AMS-12")

	profile2 := response.Profiles[1]
	assert.Equal(profile2.DataCenter, "AMS-02")
	assert.Equal(profile2.File, "https://api.leaseweb.com/bareMetals/v2/remoteManagement/profiles/lsw-rmvpn-AMS-02.ovpn")
	assert.Equal(len(profile2.SatelliteDataCenters), 2)
	assert.Equal(profile2.SatelliteDataCenters[0], "AMS-13")
	assert.Equal(profile2.SatelliteDataCenters[1], "AMS-15")
}

func TestRemoteManagementListProfilesPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "profiles": [
			{
				"datacenter": "AMS-02",
				"satelliteDatacenters": [
					"AMS-13",
					"AMS-15"
				],
				"file": "https://api.leaseweb.com/bareMetals/v2/remoteManagement/profiles/lsw-rmvpn-AMS-02.ovpn"
			}
		]}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := remoteManagementApi.ListProfiles(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Profiles), 1)

	profile1 := response.Profiles[0]
	assert.Equal(profile1.DataCenter, "AMS-02")
	assert.Equal(profile1.File, "https://api.leaseweb.com/bareMetals/v2/remoteManagement/profiles/lsw-rmvpn-AMS-02.ovpn")
	assert.Equal(len(profile1.SatelliteDataCenters), 2)
	assert.Equal(profile1.SatelliteDataCenters[0], "AMS-13")
	assert.Equal(profile1.SatelliteDataCenters[1], "AMS-15")
}

func TestRemoteManagementTestListProfilesServerErrors(t *testing.T) {
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
				return RemoteManagementApi{}.ListProfiles(ctx, PaginationOptions{})
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
				return RemoteManagementApi{}.ListProfiles(ctx, PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return RemoteManagementApi{}.ListProfiles(ctx, PaginationOptions{})
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return RemoteManagementApi{}.ListProfiles(ctx, PaginationOptions{})
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
