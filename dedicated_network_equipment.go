package leaseweb

import (
	"context"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const DEDICATED_NETWORK_EQUIPMENT_API_VERSION = "v2"

type DedicatedNetworkEquipmentApi struct{}

type DedicatedNetworkEquipments struct {
	NetworkEquipments []DedicatedNetworkEquipment `json:"networkEquipments"`
	Metadata          Metadata                    `json:"_metadata"`
}

type DedicatedNetworkEquipment struct {
	Contract            Contract                      `json:"contract"`
	FeatureAvailability FeatureAvailability           `json:"featureAvailability"`
	Id                  string                        `json:"id"`
	Name                string                        `json:"name"`
	Location            Location                      `json:"location"`
	NetworkInterfaces   NetworkInterfaces             `json:"networkInterfaces"`
	PowerPorts          []Port                        `json:"powerPorts"`
	Type                string                        `json:"type"`
	SerialNumber        string                        `json:"serialNumber"`
	Rack                DedicatedNetworkEquipmentRack `json:"rack"`
	Specs               DedicatedNetworkEquipmentSpec `json:"specs"`
}

type DedicatedNetworkEquipmentRack struct {
	Capacity string `json:"capacity"`
	Id       string `json:"id"`
}

type DedicatedNetworkEquipmentSpec struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type DedicatedNetworkEquipmentListOptions struct {
	PaginationOptions
	Reference             *string `param:"reference"`
	IP                    *string `param:"ip"`
	MacAddress            *string `param:"macAddress"`
	Site                  *string `param:"site"`
	PrivateRackID         *string `param:"privateRackId"`
	PrivateNetworkCapable *bool   `param:"privateNetworkCapable"`
	PrivateNetworkEnabled *bool   `param:"privateNetworkEnabled"`
}

type DedicatedNetworkEquipmentListIpsOptions struct {
	PaginationOptions
	NetworkType *string  `param:"networkType"`
	Version     *string  `param:"version"`
	NullRouted  *string  `param:"nullRouted"`
	IPs         []string `param:"ips"`
}

func (dnea DedicatedNetworkEquipmentApi) getPath(endpoint string) string {
	return "/bareMetals/" + DEDICATED_NETWORK_EQUIPMENT_API_VERSION + endpoint
}

func (dnea DedicatedNetworkEquipmentApi) List(ctx context.Context, opts DedicatedNetworkEquipmentListOptions) (*DedicatedNetworkEquipments, error) {
	path := dnea.getPath("/networkEquipments")
	query := options.Encode(opts)
	result := &DedicatedNetworkEquipments{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) Get(ctx context.Context, networkEquipmentId string) (*DedicatedNetworkEquipment, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId)
	result := &DedicatedNetworkEquipment{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) Update(ctx context.Context, networkEquipmentId, reference string) error {
	payload := map[string]string{"reference": reference}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId)
	return doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (dnea DedicatedNetworkEquipmentApi) ListIps(ctx context.Context, networkEquipmentId string, opts DedicatedNetworkEquipmentListIpsOptions) (*Ips, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips")
	query := options.Encode(opts)
	result := &Ips{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) GetIp(ctx context.Context, networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) UpdateIp(ctx context.Context, networkEquipmentId, ip string, payload map[string]string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) NullRouteAnIp(ctx context.Context, networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip + "/null")
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) RemoveNullRouteAnIp(ctx context.Context, networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip + "/unnull")
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListNullRoutes(ctx context.Context, networkEquipmentId string, opts PaginationOptions) (*NullRoutes, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/nullRouteHistory")
	query := options.Encode(opts)
	result := &NullRoutes{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListCredentials(ctx context.Context, networkEquipmentId string, opts PaginationOptions) (*Credentials, error) {
	result := &Credentials{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) CreateCredential(ctx context.Context, networkEquipmentId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := make(map[string]string)
	payload["type"] = credentialType
	payload["username"] = username
	payload["password"] = password
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListCredentialsByType(ctx context.Context, networkEquipmentId, credentialType string, opts PaginationOptions) (*Credentials, error) {
	result := &Credentials{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType)
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) GetCredential(ctx context.Context, networkEquipmentId, credentialType, username string) (*Credential, error) {
	result := &Credential{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) DeleteCredential(ctx context.Context, networkEquipmentId, credentialType, username string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dnea DedicatedNetworkEquipmentApi) UpdateCredential(ctx context.Context, networkEquipmentId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := map[string]string{"password": password}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) PowerCycleServer(ctx context.Context, networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerCycle")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dnea DedicatedNetworkEquipmentApi) GetPowerStatus(ctx context.Context, networkEquipmentId string) (*PowerStatus, error) {
	result := &PowerStatus{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerInfo")
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) PowerOffServer(ctx context.Context, networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerOff")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dnea DedicatedNetworkEquipmentApi) PowerOnServer(ctx context.Context, networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerOn")
	return doRequest(ctx, http.MethodPost, path, "")
}
