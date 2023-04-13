package leaseweb

import (
	"context"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
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

func (pna PrivateNetworkingApi) List(ctx context.Context, opts PaginationOptions) (*PrivateNetworks, error) {
	path := pna.getPath("/privateNetworks")
	query := options.Encode(opts)
	result := &PrivateNetworks{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Create(ctx context.Context, name string) (*PrivateNetwork, error) {
	payload := map[string]string{
		"name": name,
	}
	path := pna.getPath("/privateNetworks")
	result := &PrivateNetwork{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Get(ctx context.Context, id string) (*PrivateNetwork, error) {
	path := pna.getPath("/privateNetworks/" + id)
	result := &PrivateNetwork{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Update(ctx context.Context, id, name string) (*PrivateNetwork, error) {
	payload := map[string]string{
		"name": name,
	}
	path := pna.getPath("/privateNetworks")
	result := &PrivateNetwork{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) Delete(ctx context.Context, id string) error {
	path := pna.getPath("/privateNetworks/" + id)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (pna PrivateNetworkingApi) ListDhcpReservations(ctx context.Context, id string, opts PaginationOptions) (*PrivateNetworkingDhcpReservations, error) {
	path := pna.getPath("/privateNetworks/" + id + "/reservations")
	result := &PrivateNetworkingDhcpReservations{}
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) CreateDhcpReservation(ctx context.Context, id, ip, mac string, sticky bool) (*PrivateNetworkingDhcpReservation, error) {
	payload := map[string]interface{}{
		"ip":     ip,
		"mac":    mac,
		"sticky": sticky,
	}
	path := pna.getPath("/privateNetworks/" + id + "/reservations")
	result := &PrivateNetworkingDhcpReservation{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (pna PrivateNetworkingApi) DeleteDhcpReservation(ctx context.Context, id, ip string) error {
	path := pna.getPath("/privateNetworks/" + id + "/reservations/" + ip)
	return doRequest(ctx, http.MethodDelete, path, "")
}
