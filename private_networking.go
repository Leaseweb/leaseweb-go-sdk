package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const PRIVATE_NETWORKING_API_VERSION = "v2"

type PrivateNetworkingApi struct{}

type PrivateNetworks struct {
	PrivateNetworks []PrivateNetwork `json:"privateNetworks"`
	Metadata        Metadata         `json:"_metadata"`
}

type PrivateNetworkingDhcpReservations struct {
	DhcpReservations []PrivateNetworkingDhcpReservation `json:"reservations"`
	Metadata         Metadata                           `json:"_metadata"`
}

type PrivateNetworkingDhcpReservation struct {
	Ip     string `json:"ip"`
	Mac    string `json:"mac"`
	Sticky bool   `json:"sticky"`
}

func (pna PrivateNetworkingApi) getPath(endpoint string) string {
	return "/bareMetals/" + PRIVATE_NETWORKING_API_VERSION + endpoint
}

func (pna PrivateNetworkingApi) List(args ...int) (*PrivateNetworks, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := pna.getPath("/privateNetworks?" + v.Encode())
	result := &PrivateNetworks{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Create(name string) (*PrivateNetwork, error) {
	payload := map[string]string{
		"name": name,
	}
	path := pna.getPath("/privateNetworks")
	result := &PrivateNetwork{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Get(id string) (*PrivateNetwork, error) {
	path := pna.getPath("/privateNetworks/" + id)
	result := &PrivateNetwork{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Update(id, name string) (*PrivateNetwork, error) {
	payload := map[string]string{
		"name": name,
	}
	path := pna.getPath("/privateNetworks")
	result := &PrivateNetwork{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Delete(id string) error {
	path := pna.getPath("/privateNetworks/" + id)
	return doRequest(http.MethodDelete, path)
}

func (pna PrivateNetworkingApi) ListDhcpReservations(id string, args ...int) (*PrivateNetworkingDhcpReservations, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := pna.getPath("/privateNetworks/" + id + "/reservations?" + v.Encode())
	result := &PrivateNetworkingDhcpReservations{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) CreateDhcpReservation(id, ip, mac string, sticky bool) (*PrivateNetworkingDhcpReservation, error) {
	payload := map[string]interface{}{
		"ip":     ip,
		"mac":    mac,
		"sticky": sticky,
	}
	path := pna.getPath("/privateNetworks/" + id + "/reservations")
	result := &PrivateNetworkingDhcpReservation{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) DeleteDhcpReservation(id, ip string) error {
	path := pna.getPath("/privateNetworks/" + id + "/reservations/" + ip)
	return doRequest(http.MethodDelete, path)
}
