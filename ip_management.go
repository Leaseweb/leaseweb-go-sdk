package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const IP_MANAGEMENT_API_VERSION = "v2"

type IpManagementApi struct{}

func (ima IpManagementApi) getPath(endpoint string) string {
	return "/ipMgmt/" + IP_MANAGEMENT_API_VERSION + endpoint
}

func (ima IpManagementApi) List(ctx context.Context, params ...map[string]interface{}) (*Ips, error) {
	v := url.Values{}
	if len(params) != 0 {
		for key, value := range params[0] {
			v.Add(key, fmt.Sprint(value))
		}
	}
	path := ima.getPath("/ips")
	query := v.Encode()
	result := &Ips{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) Get(ctx context.Context, ip string) (*Ip, error) {
	path := ima.getPath("/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) Update(ctx context.Context, ip, reverseLookup string) (*Ip, error) {
	payload := map[string]string{"reverseLookup": reverseLookup}
	path := ima.getPath("/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) NullRouteAnIp(ctx context.Context, ip string, params ...map[string]string) (*NullRoute, error) {
	payload := make(map[string]string)
	if len(params) != 0 {
		for key, value := range params[0] {
			payload[key] = value
		}
	}
	path := ima.getPath("/ips/" + ip + "/nullRoute")
	result := &NullRoute{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) RemoveNullRouteAnIp(ctx context.Context, ip string) error {
	path := ima.getPath("/ips/" + ip + "/nullRoute")
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (ima IpManagementApi) ListNullRoutes(ctx context.Context, params ...map[string]interface{}) (*NullRoutes, error) {
	v := url.Values{}
	if len(params) != 0 {
		for key, value := range params[0] {
			v.Add(key, fmt.Sprint(value))
		}
	}
	path := ima.getPath("/nullRoutes")
	query := v.Encode()
	result := &NullRoutes{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) GetNullRoute(ctx context.Context, id string) (*NullRoute, error) {
	path := ima.getPath("/nullRoutes/" + id)
	result := &NullRoute{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ima IpManagementApi) UpdateNullRoute(ctx context.Context, id string, params ...map[string]string) (*NullRoute, error) {
	payload := make(map[string]string)
	if len(params) != 0 {
		for key, value := range params[0] {
			payload[key] = value
		}
	}
	path := ima.getPath("/nullRoutes/" + id)
	result := &NullRoute{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}
