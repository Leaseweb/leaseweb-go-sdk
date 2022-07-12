package leaseweb

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const FLOATING_IP_API_DEFAULT_VERSION = "v2"

type FloatingIpApi struct {
	version string
}

func (fia *FloatingIpApi) SetVersion(version string) {
	fia.version = version
}

func (fia FloatingIpApi) GetVersion() string {
	if fia.version == "" {
		return FLOATING_IP_API_DEFAULT_VERSION
	}
	return fia.version
}

func (fia FloatingIpApi) GetPath() string {
	return "/floatingIps"
}

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

func (fia FloatingIpApi) ListRanges(args ...interface{}) (FloatingIpRanges, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	floatingIpRanges := &FloatingIpRanges{}
	r := leasewebRequest{
		response: floatingIpRanges,
		method:   GET,
		endpoint: fia.GetPath() + "/" + fia.GetVersion() + "/ranges" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return FloatingIpRanges{}, err
	}

	return *floatingIpRanges, nil
}

func (fia FloatingIpApi) GetRange(rangeId string) (FloatingIpRange, error) {

	floatingIpRange := &FloatingIpRange{}
	r := leasewebRequest{
		response: floatingIpRange,
		method:   GET,
		endpoint: fia.GetPath() + "/" + fia.GetVersion() + "/ranges/" + rangeId,
	}
	err := doRequest(r)
	if err != nil {
		return FloatingIpRange{}, err
	}

	return *floatingIpRange, nil
}

func (fia FloatingIpApi) ListRangeDefinitions(rangeId string, args ...interface{}) (FloatingIpDefinitions, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	floatingIpDefinitions := &FloatingIpDefinitions{}
	r := leasewebRequest{
		response: floatingIpDefinitions,
		method:   GET,
		endpoint: fia.GetPath() + "/" + fia.GetVersion() + "/ranges/" + rangeId + "/floatingIpDefinitions" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return FloatingIpDefinitions{}, err
	}

	return *floatingIpDefinitions, nil
}

func (fia FloatingIpApi) CreateRangeDefinition(rangeId string, floatingIp string, anchorIp string) (FloatingIpDefinition, error) {

	floatingIpDefinition := &FloatingIpDefinition{}
	r := leasewebRequest{
		response: floatingIpDefinition,
		payload:  map[string]string{"floatingIp": floatingIp, "anchorIp": anchorIp},
		method:   POST,
		endpoint: fia.GetPath() + "/" + fia.GetVersion() + "/ranges/" + rangeId + "/floatingIpDefinitions",
	}
	err := doRequest(r)
	if err != nil {
		return FloatingIpDefinition{}, err
	}

	return *floatingIpDefinition, nil
}

func (fia FloatingIpApi) GetRangeDefinition(rangeId string, floatingIpDefinitionId string) (FloatingIpDefinition, error) {

	floatingIpDefinition := &FloatingIpDefinition{}
	r := leasewebRequest{
		response: floatingIpDefinition,
		method:   GET,
		endpoint: fia.GetPath() + "/" + fia.GetVersion() + "/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId,
	}
	err := doRequest(r)
	if err != nil {
		return FloatingIpDefinition{}, err
	}

	return *floatingIpDefinition, nil
}

func (fia FloatingIpApi) UpdateRangeDefinition(rangeId string, floatingIpDefinitionId string, anchorIp string) (FloatingIpDefinition, error) {

	floatingIpDefinition := &FloatingIpDefinition{}
	r := leasewebRequest{
		response: floatingIpDefinition,
		payload:  map[string]string{"anchorIp": anchorIp},
		method:   PUT,
		endpoint: fia.GetPath() + "/" + fia.GetVersion() + "/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId,
	}
	err := doRequest(r)
	if err != nil {
		return FloatingIpDefinition{}, err
	}

	return *floatingIpDefinition, nil
}

func (fia FloatingIpApi) RemoveRangeDefinition(rangeId string, floatingIpDefinitionId string) (FloatingIpDefinition, error) {

	floatingIpDefinition := &FloatingIpDefinition{}
	r := leasewebRequest{
		response: floatingIpDefinition,
		method:   DELETE,
		endpoint: fia.GetPath() + "/" + fia.GetVersion() + "/ranges/" + rangeId + "/floatingIpDefinitions/" + floatingIpDefinitionId,
	}
	err := doRequest(r)
	if err != nil {
		return FloatingIpDefinition{}, err
	}

	return *floatingIpDefinition, nil
}
