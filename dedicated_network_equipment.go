package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const DEDICATED_NETWORK_EQUIPMENT_API_VERSION = "v2"

type DedicatedNetworkEquipmentApi struct {
	Client *LeasewebClient
}

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

func (dnea DedicatedNetworkEquipmentApi) getPath(endpoint string) string {
	return "/bareMetals/" + DEDICATED_NETWORK_EQUIPMENT_API_VERSION + endpoint
}

func (dnea DedicatedNetworkEquipmentApi) List(ctx context.Context, args ...interface{}) (*DedicatedNetworkEquipments, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("ip", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("macAddress", fmt.Sprint(args[3]))
	}
	if len(args) >= 5 {
		v.Add("site", fmt.Sprint(args[4]))
	}
	if len(args) >= 6 {
		v.Add("privateRackId", fmt.Sprint(args[5]))
	}
	if len(args) >= 7 {
		v.Add("privateNetworkCapable", fmt.Sprint(args[6]))
	}
	if len(args) >= 8 {
		v.Add("privateNetworkEnabled", fmt.Sprint(args[7]))
	}

	path := dnea.getPath("/networkEquipments")
	query := v.Encode()
	result := &DedicatedNetworkEquipments{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) Get(ctx context.Context, networkEquipmentId string) (*DedicatedNetworkEquipment, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId)
	result := &DedicatedNetworkEquipment{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) Update(ctx context.Context, networkEquipmentId, reference string) error {
	payload := map[string]string{"reference": reference}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId)
	return getClient(dnea.Client).doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (dnea DedicatedNetworkEquipmentApi) ListIps(ctx context.Context, networkEquipmentId string, args ...interface{}) (*Ips, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("networkType", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("version", fmt.Sprint(args[3]))
	}
	if len(args) >= 5 {
		v.Add("nullRouted", fmt.Sprint(args[4]))
	}
	if len(args) >= 6 {
		v.Add("ips", fmt.Sprint(args[5]))
	}

	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips")
	query := v.Encode()
	result := &Ips{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) GetIp(ctx context.Context, networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip)
	result := &Ip{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) UpdateIp(ctx context.Context, networkEquipmentId, ip string, payload map[string]string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip)
	result := &Ip{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) NullRouteAnIp(ctx context.Context, networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip + "/null")
	result := &Ip{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) RemoveNullRouteAnIp(ctx context.Context, networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip + "/unnull")
	result := &Ip{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListNullRoutes(ctx context.Context, networkEquipmentId string, args ...interface{}) (*NullRoutes, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/nullRouteHistory")
	query := v.Encode()
	result := &NullRoutes{}
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListCredentials(ctx context.Context, networkEquipmentId string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &Credentials{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials")
	query := v.Encode()
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, query, result); err != nil {
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
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListCredentialsByType(ctx context.Context, networkEquipmentId, credentialType string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &Credentials{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType)
	query := v.Encode()
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) GetCredential(ctx context.Context, networkEquipmentId, credentialType, username string) (*Credential, error) {
	result := &Credential{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) DeleteCredential(ctx context.Context, networkEquipmentId, credentialType, username string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	return getClient(dnea.Client).doRequest(ctx, http.MethodDelete, path, "")
}

func (dnea DedicatedNetworkEquipmentApi) UpdateCredential(ctx context.Context, networkEquipmentId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := map[string]string{"password": password}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) PowerCycleServer(ctx context.Context, networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerCycle")
	return getClient(dnea.Client).doRequest(ctx, http.MethodPost, path, "")
}

func (dnea DedicatedNetworkEquipmentApi) GetPowerStatus(ctx context.Context, networkEquipmentId string) (*PowerStatus, error) {
	result := &PowerStatus{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerInfo")
	if err := getClient(dnea.Client).doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) PowerOffServer(ctx context.Context, networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerOff")
	return getClient(dnea.Client).doRequest(ctx, http.MethodPost, path, "")
}

func (dnea DedicatedNetworkEquipmentApi) PowerOnServer(ctx context.Context, networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerOn")
	return getClient(dnea.Client).doRequest(ctx, http.MethodPost, path, "")
}
