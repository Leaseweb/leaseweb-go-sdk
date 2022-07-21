package leaseweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeCredentials(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	err := remoteManagementApi.ChangeCredentials("new password")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestTestChangeCredentialsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return nil, RemoteManagementApi{}.ChangeCredentials("new password")
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
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return nil, RemoteManagementApi{}.ChangeCredentials("new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "403",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return nil, RemoteManagementApi{}.ChangeCredentials("new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "500",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return nil, RemoteManagementApi{}.ChangeCredentials("new password")
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "503",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestListProfiles(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
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
	response, err := remoteManagementApi.ListProfiles()

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

func TestListProfilesPaginate(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
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
	response, err := remoteManagementApi.ListProfiles(1)

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

func TestTestListProfilesServerErrors(t *testing.T) {

	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return RemoteManagementApi{}.ListProfiles()
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
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return RemoteManagementApi{}.ListProfiles()
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "403",
				ErrorMessage:  "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return RemoteManagementApi{}.ListProfiles()
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "500",
				ErrorMessage:  "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				return RemoteManagementApi{}.ListProfiles()
			},
			ExpectedError: LeasewebError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				ErrorCode:     "503",
				ErrorMessage:  "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
