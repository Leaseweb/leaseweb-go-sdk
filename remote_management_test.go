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

func TestTestChangeCredentialsError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	err := remoteManagementApi.ChangeCredentials("new password")

	assert := assert.New(t)

	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestTestChangeCredentialsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	err := remoteManagementApi.ChangeCredentials("new password")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestTestChangeCredentialsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	err := remoteManagementApi.ChangeCredentials("new password")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestTestChangeCredentialsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	err := remoteManagementApi.ChangeCredentials("new password")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
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

func TestTestListProfilesError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	resp, err := remoteManagementApi.ListProfiles()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestTestListProfilesError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	resp, err := remoteManagementApi.ListProfiles()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestTestListProfilesError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	resp, err := remoteManagementApi.ListProfiles()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestTestListProfilesError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	remoteManagementApi := RemoteManagementApi{}
	resp, err := remoteManagementApi.ListProfiles()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}
