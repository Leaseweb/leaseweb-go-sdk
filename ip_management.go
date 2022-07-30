package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const IP_MANAGEMENT_API_VERSION = "v2"

type IpManagementApi struct{}

type Ips struct {
	Ips      []Ip     `json:"ips"`
	Metadata Metadata `json:"_metadata"`
}

type Ip struct {
	Ip               string             `json:"ip"`
	Version          int                `json:"version"`
	Type             string             `json:"type"`
	PrefixLength     int                `json:"prefixLength"`
	Primary          bool               `json:"primary"`
	ReverseLookup    string             `json:"reverseLookup"`
	NullRouted       bool               `json:"nullRouted"`
	UnnullingAllowed bool               `json:"unnullingAllowed"`
	EquipmentId      string             `json:"equipmentId"`
	AssignedContract IpAssignedContract `json:"assignedContract"`
	Subnet           IpSubnet           `json:"subnet"`
}

type IpAssignedContract struct {
	Id string `json:"id"`
}

type IpSubnet struct {
	Id           string `json:"id"`
	NetworkIp    string `json:"networkIp"`
	PrefixLength int    `json:"prefixLength"`
	Gateway      string `json:"gateway"`
}

type NullRoutes struct {
	NullRoutes []NullRoute `json:"nullroutes"`
	Metadata   Metadata    `json:"_metadata"`
}

type NullRoute struct {
	Id                   string             `json:"id"`
	Ip                   string             `json:"ip"`
	NulledAt             string             `json:"nulledAt"`
	NulledBy             string             `json:"nulledBy"`
	NullLevel            int                `json:"nullLevel"`
	AutomatedUnnullingAt string             `json:"automatedUnnullingAt"`
	UnnulledAt           string             `json:"unnulledAt"`
	UnnulledBy           string             `json:"unnulledBy"`
	TicketId             string             `json:"ticketId"`
	Comment              string             `json:"comment"`
	EquipmentId          string             `json:"equipmentId"`
	AssignedContract     IpAssignedContract `json:"assignedContract"`
}

func (ima IpManagementApi) getPath(endpoint string) string {
	return "/ipMgmt/" + IP_MANAGEMENT_API_VERSION + endpoint
}

func (ima IpManagementApi) ListIps(params ...map[string]interface{}) (*Ips, error) {
	v := url.Values{}
	if len(params) != 0 {
		for key, value := range params[0] {
			v.Add(key, fmt.Sprint(value))
		}
	}
	path := ima.getPath("/ips?" + v.Encode())
	result := &Ips{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) GetIp(ip string) (*Ip, error) {
	path := ima.getPath("/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) UpdateIp(ip, reverseLookup string) (*Ip, error) {
	payload := map[string]string{"reverseLookup": reverseLookup}
	path := ima.getPath("/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) NullRouteIp(ip string, params ...map[string]string) (*NullRoute, error) {
	payload := make(map[string]string)
	if len(params) != 0 {
		for key, value := range params[0] {
			payload[key] = value
		}
	}
	path := ima.getPath("/ips/" + ip + "/nullRoute")
	result := &NullRoute{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) RemoveNullRouteIp(ip string) error {
	path := ima.getPath("/ips/" + ip + "/nullRoute")
	if err := doRequest(http.MethodDelete, path); err != nil {
		return err
	}
	return nil
}

func (ima IpManagementApi) ListNullRouteHistory(params ...map[string]interface{}) (*NullRoutes, error) {
	v := url.Values{}
	if len(params) != 0 {
		for key, value := range params[0] {
			v.Add(key, fmt.Sprint(value))
		}
	}
	path := ima.getPath("/nullRoutes?" + v.Encode())
	result := &NullRoutes{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) GetNullRouteHistory(id string) (*NullRoute, error) {
	path := ima.getPath("/nullRoutes/" + id)
	result := &NullRoute{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) UpdateNullRouteIp(id string, params ...map[string]string) (*NullRoute, error) {
	payload := make(map[string]string)
	if len(params) != 0 {
		for key, value := range params[0] {
			payload[key] = value
		}
	}
	path := ima.getPath("/nullRoutes/" + id)
	result := &NullRoute{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}
