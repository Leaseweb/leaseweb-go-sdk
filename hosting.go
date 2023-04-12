package leaseweb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
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

type HostingListDomainsOptions struct {
	PaginationOptions
	Type *string `param:"type"`
}

func (ha HostingApi) getPath(endpoint string) string {
	return "/hosting/" + FLOATING_IP_API_VERSION + endpoint
}

func (ha HostingApi) ListDomains(ctx context.Context, opts HostingListDomainsOptions) (*HostingDomains, error) {
	path := ha.getPath("/domains")
	query := options.Encode(opts)
	result := &HostingDomains{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetDomain(ctx context.Context, domainName string) (*HostingDomain, error) {
	path := ha.getPath("/domains/" + domainName)
	result := &HostingDomain{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetAvailability(ctx context.Context, domainName string) (*HostingDomainAvailability, error) {
	path := ha.getPath("/domains/" + domainName + "/available")
	result := &HostingDomainAvailability{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListNameservers(ctx context.Context, domainName string, opts PaginationOptions) (*HostingNameservers, error) {
	path := ha.getPath("/domains/" + domainName + "/nameservers")
	result := &HostingNameservers{}
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateNameservers(ctx context.Context, domainName string, nameservers []string) (*HostingNameservers, error) {
	payload := make(map[string][]string)
	payload["nameservers"] = nameservers
	path := ha.getPath("/domains/" + domainName + "/nameservers")
	result := &HostingNameservers{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetDnsSecurity(ctx context.Context, domainName string) (*HostingDnsSecurity, error) {
	path := ha.getPath("/domains/" + domainName + "/dnssec")
	result := &HostingDnsSecurity{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateDnsSecurity(ctx context.Context, domainName string, payload map[string]interface{}) (*HostingInfoMessageResponse, error) {
	path := ha.getPath("/domains/" + domainName + "/dnssec")
	result := &HostingInfoMessageResponse{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListResourceRecordSets(ctx context.Context, domainName string, opts PaginationOptions) (*HostingResourceRecordSets, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets")
	result := &HostingResourceRecordSets{}
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateResourceRecordSet(ctx context.Context, domainName string, payload map[string]interface{}) (*HostingResourceRecordSet, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets")
	result := &HostingResourceRecordSet{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteResourceRecordSets(ctx context.Context, domainName string) error {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets")
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (ha HostingApi) GetResourceRecordSet(ctx context.Context, domainName, name, recordType string) (*HostingResourceRecordSet, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/" + name + "/" + recordType)
	result := &HostingResourceRecordSet{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateResourceRecordSet(ctx context.Context, domainName, name, recordType string, content []string, ttl int) (*HostingResourceRecordSet, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/" + name + "/" + recordType)
	result := &HostingResourceRecordSet{}
	payload := make(map[string]interface{})
	payload["content"] = content
	payload["ttl"] = ttl
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteResourceRecordSet(ctx context.Context, domainName, name, recordType string) error {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/" + name + "/" + recordType)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (ha HostingApi) ValidateResourceRecordSet(ctx context.Context, domainName, name, recordType string, content []string, ttl int) (*HostingInfoMessageResponse, error) {
	path := ha.getPath("/domains/" + domainName + "/resourceRecordSets/validateSet")
	result := &HostingInfoMessageResponse{}
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["type"] = recordType
	payload["content"] = content
	payload["ttl"] = ttl
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListCatchAll(ctx context.Context, domainName string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/catchAll")
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateCatchAll(ctx context.Context, domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/catchAll")
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateCatchAll(ctx context.Context, domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	return ha.UpdateCatchAll(ctx, domainName, payload)
}

func (ha HostingApi) DeleteCatchAll(ctx context.Context, domainName string) error {
	path := ha.getPath("/domains/" + domainName + "/catchAll")
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (ha HostingApi) ListEmailAliases(ctx context.Context, domainName string, opts PaginationOptions) (*HostingEmailAliases, error) {
	path := ha.getPath("/domains/" + domainName + "/emailAliases")
	query := options.Encode(opts)
	result := &HostingEmailAliases{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateEmailAlias(ctx context.Context, domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/emailAliases")
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetEmailAlias(ctx context.Context, domainName, source, destination string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/emailAliases/" + source + "/" + destination)
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateEmailAlias(ctx context.Context, domainName, source, destination string, payload map[string]bool) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/emailAliases/" + source + "/" + destination)
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteEmailAlias(ctx context.Context, domainName, source, destination string) error {
	path := ha.getPath("/domains/" + domainName + "/emailAliases/" + source + "/" + destination)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (ha HostingApi) ListDomainForwards(ctx context.Context, domainName string, opts PaginationOptions) (*HostingEmailForwards, error) {
	path := ha.getPath("/domains/" + domainName + "/forwards")
	query := options.Encode(opts)
	result := &HostingEmailForwards{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) ListMailBoxes(ctx context.Context, domainName string, opts PaginationOptions) (*HostingEmailMailboxes, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes")
	query := options.Encode(opts)
	result := &HostingEmailMailboxes{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateMailBox(ctx context.Context, domainName string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes")
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetMailBox(ctx context.Context, domainName, emailAddress string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress)
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateMailBox(ctx context.Context, domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress)
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteMailBox(ctx context.Context, domainName, emailAddress string) error {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (ha HostingApi) GetAutoResponder(ctx context.Context, domainName, emailAddress string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/autoResponder")
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateAutoResponder(ctx context.Context, domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/autoResponder")
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateAutoResponder(ctx context.Context, domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	return ha.UpdateAutoResponder(ctx, domainName, emailAddress, payload)
}

func (ha HostingApi) DeleteAutoResponder(ctx context.Context, domainName, emailAddress string) error {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/autoResponder")
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (ha HostingApi) ListForwards(ctx context.Context, domainName, emailAddress string, opts PaginationOptions) (*HostingEmailForwards, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards")
	result := &HostingEmailForwards{}
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) CreateForward(ctx context.Context, domainName, emailAddress string, payload map[string]interface{}) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards")
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) GetForward(ctx context.Context, domainName, emailAddress, destination string) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards/" + destination)
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) UpdateForward(ctx context.Context, domainName, emailAddress, destination string, payload map[string]bool) (*HostingEmail, error) {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards/" + destination)
	result := &HostingEmail{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (ha HostingApi) DeleteForward(ctx context.Context, domainName, emailAddress, destination string) error {
	path := ha.getPath("/domains/" + domainName + "/mailboxes/" + emailAddress + "/forwards/" + destination)
	return doRequest(ctx, http.MethodDelete, path, "")
}
