package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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
	FloatingIpDefinitions []FloatingIpDefinition `json:"floatingIpDefinitions"`
	Metadata              Metadata               `json:"_metadata"`
}

func (fia FloatingIpApi) getPath(endpoint string) string {
	return "/floatingIps/" + FLOATING_IP_API_VERSION + endpoint
}

func (fia FloatingIpApi) ListRanges(args ...interface{}) (*FloatingIpRanges, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		s := reflect.ValueOf(args[3])
		var types []string
		for i := 0; i < s.Len(); i++ {
			types = append(types, s.Index(i).String())
		}
		v.Add("type", strings.Join(types, ","))
	}
	if len(args) >= 4 {
		v.Add("location", fmt.Sprint(args[1]))
	}

	path := fia.getPath("/ranges?" + v.Encode())
	result := &FloatingIpRanges{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) GetRange(rangeId string) (*FloatingIpRange, error) {
	path := fia.getPath("/ranges/" + rangeId)
	result := &FloatingIpRange{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) ListRangeDefinitions(rangeId string, args ...interface{}) (*FloatingIpDefinitions, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		s := reflect.ValueOf(args[3])
		var types []string
		for i := 0; i < s.Len(); i++ {
			types = append(types, s.Index(i).String())
		}
		v.Add("type", strings.Join(types, ","))
	}
	if len(args) >= 4 {
		v.Add("location", fmt.Sprint(args[1]))
	}

	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions?" + v.Encode())
	result := &FloatingIpDefinitions{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) CreateRangeDefinition(rangeId string, floatingIp string, anchorIp string) (*FloatingIpDefinition, error) {
	payload := map[string]string{"floatingIp": floatingIp, "anchorIp": anchorIp}
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions")
	result := &FloatingIpDefinition{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) GetRangeDefinition(rangeId string, floatingIpDefinitionId string) (*FloatingIpDefinition, error) {
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId)
	result := &FloatingIpDefinition{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) UpdateRangeDefinition(rangeId string, floatingIpDefinitionId string, anchorIp string) (*FloatingIpDefinition, error) {
	payload := map[string]string{"anchorIp": anchorIp}
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId)
	result := &FloatingIpDefinition{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (fia FloatingIpApi) RemoveRangeDefinition(rangeId string, floatingIpDefinitionId string) (*FloatingIpDefinition, error) {
	path := fia.getPath("/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId)
	result := &FloatingIpDefinition{}
	if err := doRequest(http.MethodDelete, path, result); err != nil {
		return nil, err
	}
	return result, nil
}
