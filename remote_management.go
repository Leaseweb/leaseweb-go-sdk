package leaseweb

import (
	"fmt"
	"net/url"
)

const REMOTE_MANAGEMENT_API_DEFAULT_VERSION = "v2"

type RemoteManagementApi struct {
	version string
}

func (rma *RemoteManagementApi) SetVersion(version string) {
	rma.version = version
}

func (rma RemoteManagementApi) GetVersion() string {
	if rma.version == "" {
		return REMOTE_MANAGEMENT_API_DEFAULT_VERSION
	}
	return rma.version
}

func (rma RemoteManagementApi) GetPath() string {
	return "/bareMetals"
}

type Profiles struct {
	Metadata Metadata  `json:"_metadata"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	DataCenter           string   `json:"datacenter"`
	File                 string   `json:"file"`
	SatelliteDataCenters []string `json:"satelliteDatacenters"`
}

func (rma RemoteManagementApi) ChangeCredentials(password string) error {

	r := leasewebRequest{
		payload:  map[string]string{password: password},
		method:   POST,
		endpoint: rma.GetPath() + "/" + rma.GetVersion() + "/remoteManagement/changeCredentials",
	}
	err := doRequest(r)
	if err != nil {
		return err
	}

	return nil
}

func (rma RemoteManagementApi) ListProfiles(args ...int) (Profiles, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	profiles := &Profiles{}
	r := leasewebRequest{
		response: profiles,
		method:   GET,
		endpoint: rma.GetPath() + "/" + rma.GetVersion() + "/remoteManagement/profiles" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return Profiles{}, err
	}

	return *profiles, nil
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
