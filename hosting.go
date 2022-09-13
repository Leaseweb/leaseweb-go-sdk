package leaseweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const HOSTING_API_VERSION = "v2"

type HostingApi struct{}

type HostingDomains struct {
	Domains  []HostingDomain `json:"domains"`
	Link     HostingLink     `json:"_links"`
	Metadata Metadata        `json:"_metadata"`
}

type HostingDomain struct {
	DomainName        string      `json:"domainName"`
	Suspended         bool        `json:"suspended"`
	DnsOnly           bool        `json:"dnsOnly"`
	Status            string      `json:"status"`
	ContractStartDate string      `json:"contractStartDate"`
	ContractEndDate   string      `json:"contractEndDate"`
	Link              HostingLink `json:"_links"`
}

type HostingLink struct {
	Self               HostingHref `json:"self"`
	Dns                HostingHref `json:"dns"`
	Email              HostingHref `json:"email"`
	Collection         HostingHref `json:"collection"`
	CatchAll           HostingHref `json:"catchAll"`
	Contacts           HostingHref `json:"contacts"`
	Nameservers        HostingHref `json:"nameservers"`
	EmailAliases       HostingHref `json:"emailAliases"`
	Forwards           HostingHref `json:"forwards"`
	Mailboxes          HostingHref `json:"mailboxes"`
	ResourceRecordSets HostingHref `json:"resourceRecordSets"`
	ValidateZone       HostingHref `json:"validateZone"`
	ValidateSet        HostingHref `json:"validateSet"`
	Parent             HostingHref `json:"parent"`
	AutoResponder      HostingHref `json:"autoResponder"`
}

type HostingHref struct {
	Href string `json:"href"`
}

type HostingDomainAvailability struct {
	DomainName string      `json:"domainName"`
	Available  bool        `json:"available"`
	Link       HostingLink `json:"_links"`
}

type HostingNameservers struct {
	Nameservers []HostingNameserver `json:"nameservers"`
	Status      string              `json:"status"`
	Link        HostingLink         `json:"_links"`
	Metadata    Metadata            `json:"_metadata"`
}

type HostingNameserver struct {
	Name string `json:"name"`
}

type HostingDnsSecurity struct {
	Status      string                        `json:"status"`
	Nsec        HostingDnsSecurityNsec        `json:"nsec3"`
	KeyRollover HostingDnsSecurityKeyRollover `json:"keyRollover"`
	Link        HostingLink                   `json:"_links"`
}

type HostingDnsSecurityNsec struct {
	Active bool `json:"active"`
}

type HostingDnsSecurityKeyRollover struct {
	Status  string `json:"status"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type HostingInfoMessageResponse struct {
	InfoMessage string `json:"infoMessage"`
}

type HostingResourceRecordSets struct {
	InfoMessage        string                     `json:"infoMessage"`
	ResourceRecordSets []HostingResourceRecordSet `json:"resourceRecordSets"`
	Link               HostingLink                `json:"_links"`
	Metadata           Metadata                   `json:"_metadata"`
}

type HostingResourceRecordSet struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Ttl      json.Number `json:"ttl"`
	Content  []string    `json:"content"`
	Editable bool        `json:"editable"`
	Link     HostingLink `json:"_links"`
}

type HostingEmailAliases struct {
	EmailAliases []HostingEmail `json:"emailAliases"`
	Link         HostingLink    `json:"_links"`
	Metadata     Metadata       `json:"_metadata"`
}

type HostingEmailForwards struct {
	Forwards []HostingEmail `json:"forwards"`
	Link     HostingLink    `json:"_links"`
	Metadata Metadata       `json:"_metadata"`
}

type HostingEmailMailboxes struct {
	Mailboxes []HostingEmail `json:"mailboxes"`
	Link      HostingLink    `json:"_links"`
	Metadata  Metadata       `json:"_metadata"`
}

type HostingEmail struct {
	Destination        string      `json:"destination"`
	Source             string      `json:"source"`
	Active             bool        `json:"active"`
	SpamChecksEnabled  bool        `json:"spamChecksEnabled"`
	VirusChecksEnabled bool        `json:"virusChecksEnabled"`
	CurrentSize        json.Number `json:"currentSize"`
	MaximumSize        json.Number `json:"maximumSize"`
	Suspended          bool        `json:"suspended"`
	LocalDelivery      bool        `json:"localDelivery"`
	EmailAddress       string      `json:"emailAddress"`
	Body               string      `json:"body"`
	Subject            string      `json:"subject"`
	Embedded           struct {
		AutoResponder struct {
			Link HostingLink `json:"_links"`
		} `json:"autoResponder"`
	} `json:"_embedded"`
	Link HostingLink `json:"_links"`
}

func (ha HostingApi) getPath(endpoint string) string {
	return "/hosting/" + FLOATING_IP_API_VERSION + endpoint
}

func (ha HostingApi) ListDomains(args ...interface{}) (*HostingDomains, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("type", fmt.Sprint(args[1]))
	}

	path := ha.getPath("/domains?" + v.Encode())
	result := &HostingDomains{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetDomain(domainName string) (*HostingDomain, error) {
	path := ha.getPath("/domains/" + domainName)
	result := &HostingDomain{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetAvailability(domainName string) (*HostingDomainAvailability, error) {
	path := ha.getPath("/domains/" + domainName + "/available")
	result := &HostingDomainAvailability{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListNameservers(domainName string, args ...interface{}) (*HostingNameservers, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ha.getPath("/domains/" + domainName + "/nameservers?" + v.Encode())
	result := &HostingNameservers{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateNameservers(domainName string, nameservers []string) (*HostingNameservers, error) {
	payload := make(map[string][]string)
	payload["nameservers"] = nameservers
	path := ha.getPath("/domains/" + domainName + "/nameservers")
	result := &HostingNameservers{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetDnsSecurity(domainName string) (*HostingDnsSecurity, error) {
	path := ha.getPath("/domains/" + domainName + "/dnssec")
	result := &HostingDnsSecurity{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateDnsSecurity(domainName string, payload map[string]interface{}) (*HostingInfoMessageResponse, error) {
	path := ha.getPath("/domains/" + domainName + "/dnssec")
	result := &HostingInfoMessageResponse{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListResourceRecordSets(domainName string, args ...interface{}) (*HostingResourceRecordSets, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets?" + v.Encode())
	result := &HostingResourceRecordSets{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateResourceRecordSet(domainName string, payload map[string]interface{}) (*HostingResourceRecordSet, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets")
	result := &HostingResourceRecordSet{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteResourceRecordSets(domainName string) error {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets")
	return doRequest(http.MethodDelete, path)
}

func (ha HostingApi) GetResourceRecordSet(domainName, name, recordType string) (*HostingResourceRecordSet, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/" + name + "/" + recordType)
	result := &HostingResourceRecordSet{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateResourceRecordSet(domainName, name, recordType string, content []string, ttl int) (*HostingResourceRecordSet, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/" + name + "/" + recordType)
	result := &HostingResourceRecordSet{}
	payload := make(map[string]interface{})
	payload["content"] = content
	payload["ttl"] = ttl
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteResourceRecordSet(domainName, name, recordType string) error {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/" + name + "/" + recordType)
	return doRequest(http.MethodDelete, path)
}

func (ha HostingApi) ValidateResourceRecordSet(domainName, name, recordType string, content []string, ttl int) (*HostingInfoMessageResponse, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/validateSet")
	result := &HostingInfoMessageResponse{}
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["type"] = recordType
	payload["content"] = content
	payload["ttl"] = ttl
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListCatchAll(domainName string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/catchAll")
	result := &HostingEmail{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateCatchAll(domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/catchAll")
	result := &HostingEmail{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateCatchAll(domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	return ha.UpdateCatchAll(domainName, payload)
}

func (ha HostingApi) DeleteCatchAll(domainName string) error {
	path := ha.getPath("/domains/" + domainName + "/catchAll")
	return doRequest(http.MethodDelete, path)
}

func (ha HostingApi) ListEmailAliases(domainName string, args ...interface{}) (*HostingEmailAliases, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ha.getPath("/domains/" + domainName + "/emailAliases?" + v.Encode())
	result := &HostingEmailAliases{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateEmailAlias(domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/emailAliases")
	result := &HostingEmail{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetEmailAlias(domainName, source, destination string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/emailAliases/" + source + "/" + destination)
	result := &HostingEmail{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateEmailAlias(domainName, source, destination string, payload map[string]bool) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/emailAliases/" + source + "/" + destination)
	result := &HostingEmail{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteEmailAlias(domainName, source, destination string) error {
	path := ha.getPath("/domains/" + domainName + "/emailAliases/" + source + "/" + destination)
	return doRequest(http.MethodDelete, path)
}

func (ha HostingApi) ListDomainForwards(domainName string, args ...interface{}) (*HostingEmailForwards, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ha.getPath("/domains/" + domainName + "/forwards?" + v.Encode())
	result := &HostingEmailForwards{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListMailBoxes(domainName string, args ...interface{}) (*HostingEmailMailboxes, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ha.getPath("/domains/" + domainName + "/mailboxes?" + v.Encode())
	result := &HostingEmailMailboxes{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateMailBox(domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes")
	result := &HostingEmail{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetMailBox(domainName, emailAddress string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress)
	result := &HostingEmail{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateMailBox(domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress)
	result := &HostingEmail{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteMailBox(domainName, emailAddress string) error {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress)
	return doRequest(http.MethodDelete, path)
}

func (ha HostingApi) GetAutoResponder(domainName, emailAddress string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/autoResponder")
	result := &HostingEmail{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateAutoResponder(domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/autoResponder")
	result := &HostingEmail{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateAutoResponder(domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	return ha.UpdateAutoResponder(domainName, emailAddress, payload)
}

func (ha HostingApi) DeleteAutoResponder(domainName, emailAddress string) error {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/autoResponder")
	return doRequest(http.MethodDelete, path)
}

func (ha HostingApi) ListForwards(domainName, emailAddress string, args ...interface{}) (*HostingEmailForwards, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards?" + v.Encode())
	result := &HostingEmailForwards{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateForward(domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards")
	result := &HostingEmail{}
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetForward(domainName, emailAddress, destination string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards/" + destination)
	result := &HostingEmail{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateForward(domainName, emailAddress, destination string, payload map[string]bool) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards/" + destination)
	result := &HostingEmail{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteForward(domainName, emailAddress, destination string) error {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards/" + destination)
	return doRequest(http.MethodDelete, path)
}
