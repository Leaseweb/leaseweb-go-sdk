package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const VIRTUAL_SERVER_API_VERSION = "v2"

type VirtualServerApi struct{}

type VirtualServers struct {
	VirtualServers []VirtualServer `json:"virtualServers"`
	Metadata       Metadata        `json:"_metadata"`
}

type VirtualServer struct {
	Id              string           `json:"id"`
	Reference       string           `json:"reference"`
	CustomerId      string           `json:"customerId"`
	DataCenter      string           `json:"dataCenter"`
	CloudServerId   string           `json:"cloudServerId"`
	State           string           `json:"state"`
	FirewallState   string           `json:"firewallState"`
	Template        string           `json:"template"`
	ServiceOffering string           `json:"serviceOffering"`
	Sla             string           `json:"sla"`
	Iso             VirtualServerIso `json:"iso"`
	Contract        Contract         `json:"contract"`
	Ips             []Ip             `json:"ips"`
	Hardware        Hardware         `json:"hardware"`
}

type VirtualServerIso struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type VirtualServerResult struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

type VirtualServerTemplates struct {
	Templates []VirtualServerTemplate `json:"templates"`
	Metadata  Metadata                `json:"_metadata"`
}

type VirtualServerTemplate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (vsa VirtualServerApi) getPath(endpoint string) string {
	return "/cloud/" + VIRTUAL_SERVER_API_VERSION + endpoint
}

func (vsa VirtualServerApi) List(ctx context.Context, args ...int) (*VirtualServers, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := vsa.getPath("/virtualServers")
	query := v.Encode()
	result := &VirtualServers{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) Get(ctx context.Context, virtualServerId string) (*VirtualServer, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId)
	result := &VirtualServer{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) Update(ctx context.Context, virtualServerId, reference string) (*VirtualServer, error) {
	payload := map[string]string{"reference": reference}
	path := vsa.getPath("/virtualServers/" + virtualServerId)
	result := &VirtualServer{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) PowerOn(ctx context.Context, virtualServerId string) (*VirtualServerResult, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/powerOn")
	result := &VirtualServerResult{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) PowerOff(ctx context.Context, virtualServerId string) (*VirtualServerResult, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/powerOff")
	result := &VirtualServerResult{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) Reboot(ctx context.Context, virtualServerId string) (*VirtualServerResult, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/reboot")
	result := &VirtualServerResult{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) Reinstall(ctx context.Context, virtualServerId, operatingSystemId string) (*VirtualServerResult, error) {
	payload := map[string]string{"operatingSystemId": operatingSystemId}
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/reinstall")
	result := &VirtualServerResult{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) UpdateCredential(ctx context.Context, virtualServerId, username, credentialType, password string) error {
	payload := map[string]string{"username": username, "type": credentialType, "password": password}
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/credentials")
	return doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (vsa VirtualServerApi) ListCredentials(ctx context.Context, virtualServerId, credentialType string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := vsa.getPath("/virtualServers/" + virtualServerId + "/credentials/" + credentialType)
	query := v.Encode()
	result := &Credentials{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) GetCredential(ctx context.Context, virtualServerId, username, credentialType string) (*Credential, error) {
	path := vsa.getPath("/virtualServers/" + virtualServerId + "/credentials/" + credentialType + "/" + username)
	result := &Credential{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) GetDataTrafficMetrics(ctx context.Context, virtualServerId string, args ...interface{}) (*DataTrafficMetricsV2, error) {
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

	path := vsa.getPath("/virtualServers/" + virtualServerId + "/metrics/datatraffic")
	query := v.Encode()
	result := &DataTrafficMetricsV2{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (vsa VirtualServerApi) ListTemplates(ctx context.Context, virtualServerId string, args ...int) (*VirtualServerTemplates, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := vsa.getPath("/virtualServers/" + virtualServerId + "/templates")
	result := &VirtualServerTemplates{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}
