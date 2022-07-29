package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const VIRTUAL_SERVER_API_DEFAULT_VERSION = "v2"

type VirtualServerApi struct {
	version string
}

type VirtualServers struct {
	VirtualServers []VirtualServer `json:"virtualServers"`
	Metadata       Metadata        `json:"_metadata"`
}

type VirtualServer struct {
	Id              string                `json:"id"`
	Reference       string                `json:"reference"`
	CustomerId      string                `json:"customerId"`
	DataCenter      string                `json:"dataCenter"`
	CloudServerId   string                `json:"cloudServerId"`
	State           string                `json:"state"`
	FirewallState   string                `json:"firewallState"`
	Template        string                `json:"template"`
	ServiceOffering string                `json:"serviceOffering"`
	Sla             string                `json:"sla"`
	Iso             VirtualServerIso      `json:"iso"`
	Contract        VirtualServerContract `json:"contract"`
	Ips             []VirtualServerIp     `json:"ips"`
	Hardware        VirtualServerHardware `json:"hardware"`
}

type VirtualServerContract struct {
	Id                string  `json:"id"`
	StartsAt          string  `json:"startsAt"`
	EndsAt            string  `json:"endsAt"`
	BillingCycle      int     `json:"billingCycle"`
	BillingFrequency  string  `json:"billingFrequency"`
	Currency          string  `json:"currency"`
	PricePerFrequency float32 `json:"pricePerFrequency"`
}

type VirtualServerIp struct {
	Ip      string `json:"ip"`
	Version int    `json:"version"`
	Type    string `json:"type"`
}

type VirtualServerIso struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type VirtualServerHardware struct {
	Cpu struct {
		Cores int `json:"cores"`
	} `json:"cpu"`

	Memory struct {
		Unit   string `json:"unit"`
		Amount int    `json:"amount"`
	} `json:"memory"`

	Storage struct {
		Unit   string `json:"unit"`
		Amount int    `json:"amount"`
	} `json:"storage"`
}

type VirtualServerResult struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

type Templates struct {
	Templates []Template `json:"templates"`
	Metadata  Metadata   `json:"_metadata"`
}

type Template struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (vsa *VirtualServerApi) SetVersion(version string) {
	vsa.version = version
}

func (vsa VirtualServerApi) getPath(endpoint string) string {
	version := vsa.version
	if version == "" {
		version = VIRTUAL_SERVER_API_DEFAULT_VERSION
	}
	return "/cloud/" + version + endpoint
}

func (vsa VirtualServerApi) ListVirtualServers(args ...int) (*VirtualServers, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := vsa.getPath("/virtualServers?" + v.Encode())
	result := &VirtualServers{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) GetVirtualServer(virtualServerId string) (*VirtualServer, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId)
	result := &VirtualServer{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) UpdateVirtualServer(virtualServerId, reference string) (*VirtualServer, error) {
	payload := map[string]string{"reference": reference}
	path := vsa.getPath("/virtualServers/" + virtualServerId)
	result := &VirtualServer{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) PowerOn(virtualServerId string) (*VirtualServerResult, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/powerOn")
	result := &VirtualServerResult{}
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) PowerOff(virtualServerId string) (*VirtualServerResult, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/powerOff")
	result := &VirtualServerResult{}
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) Reboot(virtualServerId string) (*VirtualServerResult, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/reboot")
	result := &VirtualServerResult{}
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) Reinstall(virtualServerId, operatingSystemId string) (*VirtualServerResult, error) {
	payload := map[string]string{"operatingSystemId": operatingSystemId}
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/reinstall")
	result := &VirtualServerResult{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) UpdateCredential(virtualServerId, username, credentialType, password string) error {
	payload := map[string]string{"username": username, "type": credentialType, "password": password}
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/credentials")
	if err := doRequest(http.MethodPut, path, nil, payload); err != nil {
		return err
	}
	return nil
}

func (vsa VirtualServerApi) ListCredentials(virtualServerId, credentialType string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := vsa.getPath("/virtualServers/" + virtualServerId + "/credentials/" + credentialType + "?" + v.Encode())
	result := &Credentials{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) GetCredential(virtualServerId, username, credentialType string) (*Credential, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/credentials/" + credentialType + "/" + username)
	result := &Credential{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) GetDataTrafficMetrics(virtualServerId string, args ...interface{}) (*DataTrafficMetrics, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("granularity", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("aggregation", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("from", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("to", fmt.Sprint(args[3]))
	}

	path := vsa.getPath("/virtualServers/" + virtualServerId + "/metrics/datatraffic?" + v.Encode())
	result := &DataTrafficMetrics{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) ListTemplates(virtualServerId string, args ...int) (*Templates, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := vsa.getPath("/virtualServers/" + virtualServerId + "/templates")
	result := &Templates{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}
