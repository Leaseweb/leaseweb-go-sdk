package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const DEDICATED_NETWORK_EQUIPMENT_API_VERSION = "v2"

type DedicatedNetworkEquipmentApi struct{}

type NetworkEquipments struct {
	NetworkEquipments []NetworkEquipment `json:"networkEquipments"`
	Metadata          Metadata           `json:"_metadata"`
}

type NetworkEquipment struct {
	Contract            Contract             `json:"contract"`
	FeatureAvailability FeatureAvailability  `json:"featureAvailability"`
	Id                  string               `json:"id"`
	Name                string               `json:"name"`
	Location            Location             `json:"location"`
	NetworkInterfaces   NetworkInterfaces    `json:"networkInterfaces"`
	PowerPorts          []Port               `json:"powerPorts"`
	Type                string               `json:"type"`
	SerialNumber        string               `json:"serialNumber"`
	Rack                NetworkEquipmentRack `json:"rack"`
	Specs               NetworkEquipmentSpec `json:"specs"`
}

type NetworkEquipmentRack struct {
	Capacity string `json:"capacity"`
	Id       string `json:"id"`
}

type NetworkEquipmentSpec struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func (dnea DedicatedNetworkEquipmentApi) getPath(endpoint string) string {
	return "/bareMetals/" + DEDICATED_NETWORK_EQUIPMENT_API_VERSION + endpoint
}

func (dnea DedicatedNetworkEquipmentApi) List(args ...interface{}) (*NetworkEquipments, error) {
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

	path := dnea.getPath("/networkEquipments?" + v.Encode())
	result := &NetworkEquipments{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) Get(networkEquipmentId string) (*NetworkEquipment, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId)
	result := &NetworkEquipment{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) Update(networkEquipmentId, reference string) error {
	payload := map[string]string{"reference": reference}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId)
	return doRequest(http.MethodPut, path, nil, payload)
}

func (dnea DedicatedNetworkEquipmentApi) ListIps(networkEquipmentId string, args ...interface{}) (*Ips, error) {
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

	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips" + v.Encode())
	result := &Ips{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) GetIp(networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) UpdateIp(networkEquipmentId, ip string, payload map[string]string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) NullRouteAnIp(networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip + "/null")
	result := &Ip{}
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) RemoveNullRouteAnIp(networkEquipmentId, ip string) (*Ip, error) {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/ips/" + ip + "/unnull")
	result := &Ip{}
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListNullRoutes(networkEquipmentId string, args ...interface{}) (*NullRoutes, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/nullRouteHistory?" + v.Encode())
	result := &NullRoutes{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListCredentials(networkEquipmentId string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &Credentials{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials?" + v.Encode())
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) CreateCredential(networkEquipmentId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := make(map[string]string)
	payload["type"] = credentialType
	payload["username"] = username
	payload["password"] = password
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials")
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) ListCredentialsByType(networkEquipmentId, credentialType string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &Credentials{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "?" + v.Encode())
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) GetCredential(networkEquipmentId, credentialType, username string) (*Credential, error) {
	result := &Credential{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) DeleteCredential(networkEquipmentId, credentialType, username string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	return doRequest(http.MethodDelete, path)
}

func (dnea DedicatedNetworkEquipmentApi) UpdateCredential(networkEquipmentId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := map[string]string{"password": password}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) PowerCycleServer(networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerCycle")
	return doRequest(http.MethodPost, path)
}

func (dnea DedicatedNetworkEquipmentApi) GetPowerStatus(networkEquipmentId string) (*PowerStatus, error) {
	result := &PowerStatus{}
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerInfo")
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dnea DedicatedNetworkEquipmentApi) PowerOffServer(networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerOff")
	return doRequest(http.MethodPost, path)
}

func (dnea DedicatedNetworkEquipmentApi) PowerOnServer(networkEquipmentId string) error {
	path := dnea.getPath("/networkEquipments/" + networkEquipmentId + "/powerOn")
	return doRequest(http.MethodPost, path)
}
