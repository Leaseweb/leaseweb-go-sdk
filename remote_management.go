package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const REMOTE_MANAGEMENT_API_VERSION = "v2"

type RemoteManagementApi struct{}

type RemoteManagementProfiles struct {
	Metadata Metadata                  `json:"_metadata"`
	Profiles []RemoteManagementProfile `json:"profiles"`
}

type RemoteManagementProfile struct {
	DataCenter           string   `json:"datacenter"`
	File                 string   `json:"file"`
	SatelliteDataCenters []string `json:"satelliteDatacenters"`
}

func (rma RemoteManagementApi) getPath(endpoint string) string {
	return "/bareMetals/" + REMOTE_MANAGEMENT_API_VERSION + endpoint
}

func (rma RemoteManagementApi) ChangeCredentials(password string) error {
	payload := map[string]string{password: password}
	path := rma.getPath("/remoteManagement/changeCredentials")
	return doRequest(http.MethodPost, path, nil, payload)
}

func (rma RemoteManagementApi) ListProfiles(args ...int) (*RemoteManagementProfiles, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := rma.getPath("/remoteManagement/profiles" + v.Encode())
	result := &RemoteManagementProfiles{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}
