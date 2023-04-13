package leaseweb

import (
	"context"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
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

func (rma RemoteManagementApi) ChangeCredentials(ctx context.Context, password string) error {
	payload := map[string]string{password: password}
	path := rma.getPath("/remoteManagement/changeCredentials")
	return doRequest(ctx, http.MethodPost, path, "", nil, payload)
}

func (rma RemoteManagementApi) ListProfiles(ctx context.Context, opts PaginationOptions) (*RemoteManagementProfiles, error) {
	path := rma.getPath("/remoteManagement/profiles")
	result := &RemoteManagementProfiles{}
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}
