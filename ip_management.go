package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const IP_MANAGEMENT_API_VERSION = "v2"

type IpManagementApi struct{}

func (ima IpManagementApi) getPath(endpoint string) string {
	return "/ipMgmt/" + IP_MANAGEMENT_API_VERSION + endpoint
}

func (ima IpManagementApi) List(params ...map[string]interface{}) (*Ips, error) {
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

func (ima IpManagementApi) Get(ip string) (*Ip, error) {
	path := ima.getPath("/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) Update(ip, reverseLookup string) (*Ip, error) {
	payload := map[string]string{"reverseLookup": reverseLookup}
	path := ima.getPath("/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) NullRouteAnIp(ip string, params ...map[string]string) (*NullRoute, error) {
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

func (ima IpManagementApi) RemoveNullRouteAnIp(ip string) error {
	path := ima.getPath("/ips/" + ip + "/nullRoute")
	return doRequest(http.MethodDelete, path)
}

func (ima IpManagementApi) ListNullRoutes(params ...map[string]interface{}) (*NullRoutes, error) {
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

func (ima IpManagementApi) GetNullRoute(id string) (*NullRoute, error) {
	path := ima.getPath("/nullRoutes/" + id)
	result := &NullRoute{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) UpdateNullRoute(id string, params ...map[string]string) (*NullRoute, error) {
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
