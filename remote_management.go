package leaseweb

import (
	"fmt"
	"net/url"
)

const REMOTE_MANAGEMENT_API_VERSION = "v2"

type RemoteManagementApi struct{}

type Profiles struct {
	Metadata Metadata  `json:"_metadata"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
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
	if err := doRequest(POST, path, nil, payload); err != nil {
		return err
	}
	return nil
}

func (rma RemoteManagementApi) ListProfiles(args ...int) (*Profiles, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := rma.getPath("/remoteManagement/profiles" + v.Encode())
	result := &Profiles{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TODO: GetProfile should be tested
// func (rma RemoteManagementApi) GetProfile(datacenter string) (Profile, error) {
// 	profile := &Profile{}
// 	r := leasewebRequest{
// 		response: profile,
// 		method:   GET,
// 		endpoint: rma.GetPath() + "/" + rma.GetVersion() + "/remoteManagement/profiles/lsw-rmvpn-" + datacenter + ".ovpn",
// 	}
// 	err := doRequest(r)
// 	if err != nil {
// 		return Profile{}, err
// 	}
// 	return *profile, nil
// }
