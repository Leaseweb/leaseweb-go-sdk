package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const PRIVATE_NETWORKING_API_DEFAULT_VERSION = "v2"

type PrivateNetworkingApi struct {
	version string
}

type PrivateNetworks struct {
	PrivateNetworks []PrivateNetwork `json:"privateNetworks"`
	Metadata        Metadata         `json:"_metadata"`
}

type PrivateNetwork struct {
	EquipmentCount int      `json:"equipmentCount"`
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
	Servers        []string `json:"servers"`
}

type DhcpReservations struct {
	DhcpReservations []DhcpReservation `json:"reservations"`
	Metadata         Metadata          `json:"_metadata"`
}

type DhcpReservation struct {
	Ip     string `json:"ip"`
	Mac    string `json:"mac"`
	Sticky bool   `json:"sticky"`
}

func (pna *PrivateNetworkingApi) SetVersion(version string) {
	pna.version = version
}

func (pna PrivateNetworkingApi) getPath(endpoint string) string {
	version := pna.version
	if version == "" {
		version = PRIVATE_NETWORKING_API_DEFAULT_VERSION
	}
	return "/bareMetals/" + version + endpoint
}

func (pna PrivateNetworkingApi) ListPrivateNetworks(args ...int) (*PrivateNetworks, error) {
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

func (pna PrivateNetworkingApi) CreatePrivateNetwork(name string) (*PrivateNetwork, error) {
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

func (pna PrivateNetworkingApi) GetPrivateNetwork(id string) (*PrivateNetwork, error) {
	path := pna.getPath("/privateNetworks/" + id)
	result := &PrivateNetwork{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) UpdatePrivateNetwork(id, name string) (*PrivateNetwork, error) {
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

func (pna PrivateNetworkingApi) DeletePrivateNetwork(id string) error {
	path := pna.getPath("/privateNetworks/" + id)
	return doRequest(http.MethodDelete, path)
}

func (pna PrivateNetworkingApi) ListDhcpReservations(id string, args ...int) (*DhcpReservations, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := pna.getPath("/privateNetworks/" + id + "/reservations?" + v.Encode())
	result := &DhcpReservations{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) CreateDhcpReservation(id, ip, mac string, sticky bool) (*DhcpReservation, error) {
	payload := map[string]interface{}{
		"ip":     ip,
		"mac":    mac,
		"sticky": sticky,
	}
	path := pna.getPath("/privateNetworks/" + id + "/reservations")
	result := &DhcpReservation{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) DeleteDhcpReservation(id, ip string) error {
	path := pna.getPath("/privateNetworks/" + id + "/reservations/" + ip)
	return doRequest(http.MethodDelete, path)
}
