package leaseweb

import (
	"context"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const FLOATING_IP_API_VERSION = "v2"

type FloatingIpApi struct{}

type FloatingIpRange struct {
	Id         string `json:"id"`
	Range      string `json:"range"`
	CustomerId string `json:"customerId"`
	SalesOrgId string `json:"salesOrgId"`
	Location   string `json:"location"`
	Type       string `json:"type"`
}

type FloatingIpRanges struct {
	Ranges   []FloatingIpRange `json:"ranges"`
	Metadata Metadata          `json:"_metadata"`
}

type FloatingIpDefinition struct {
	Id         string `json:"id"`
	RangeId    string `json:"rangeId"`
	Location   string `json:"location"`
	Type       string `json:"type"`
	CustomerId string `json:"customerId"`
	SalesOrgId string `json:"salesOrgId"`
	FloatingIp string `json:"floatingIp"`
	AnchorIp   string `json:"anchorIp"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type FloatingIpDefinitions struct {
	Definitions []FloatingIpDefinition `json:"floatingIpDefinitions"`
	Metadata    Metadata               `json:"_metadata"`
}

type FloatingIpListRangesOptions struct {
	PaginationOptions
	Type     *string `param:"type"`
	Location *string `param:"location"`
}

type FloatingIpListRangeDefinitionsOptions struct {
	PaginationOptions
	Location *string  `param:"location"`
	Type     []string `param:"type"`
}

func (fia FloatingIpApi) getPath(endpoint string) string {
	return "/floatingIps/" + FLOATING_IP_API_VERSION + endpoint
}

func (fia FloatingIpApi) ListRanges(ctx context.Context, opts FloatingIpListRangesOptions) (*FloatingIpRanges, error) {
	path := fia.getPath("/ranges")
	query := options.Encode(opts)
	result := &FloatingIpRanges{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) GetRange(ctx context.Context, rangeId string) (*FloatingIpRange, error) {
	path := fia.getPath("/ranges/" + rangeId)
	result := &FloatingIpRange{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) ListRangeDefinitions(ctx context.Context, rangeId string, opts FloatingIpListRangeDefinitionsOptions) (*FloatingIpDefinitions, error) {
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions")
	query := options.Encode(opts)
	result := &FloatingIpDefinitions{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) CreateRangeDefinition(ctx context.Context, rangeId string, floatingIp string, anchorIp string) (*FloatingIpDefinition, error) {
	payload := map[string]string{"floatingIp": floatingIp, "anchorIp": anchorIp}
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions")
	result := &FloatingIpDefinition{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) GetRangeDefinition(ctx context.Context, rangeId string, floatingIpDefinitionId string) (*FloatingIpDefinition, error) {
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId)
	result := &FloatingIpDefinition{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) UpdateRangeDefinition(ctx context.Context, rangeId string, floatingIpDefinitionId string, anchorIp string) (*FloatingIpDefinition, error) {
	payload := map[string]string{"anchorIp": anchorIp}
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId)
	result := &FloatingIpDefinition{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) RemoveRangeDefinition(ctx context.Context, rangeId string, floatingIpDefinitionId string) (*FloatingIpDefinition, error) {
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId)
	result := &FloatingIpDefinition{}
	if err := doRequest(ctx, http.MethodDelete, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}
