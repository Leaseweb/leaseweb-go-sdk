package leaseweb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const ABUSE_API_VERSION = "v1"

type AbuseApi struct{}

type AbuseReport struct {
	Id                  string                    `json:"id"`
	Subject             string                    `json:"subject"`
	Status              string                    `json:"status"`
	Reopened            bool                      `json:"reopened"`
	ReportedAt          string                    `json:"reportedAt"`
	UpdatedAt           string                    `json:"updatedAt"`
	Notifier            string                    `json:"notifier"`
	CustomerId          string                    `json:"customerId"`
	LegalEntityId       string                    `json:"legalEntityId"`
	Deadline            string                    `json:"deadline"`
	Body                string                    `json:"body"`
	DetectedIpAddresses []string                  `json:"detectedIpAddresses"`
	DetectedDomainNames []AbuseDetectedDomainName `json:"detectedDomainNames"`
	Attachments         []AbuseAttachment         `json:"attachments"`
	TotalMessagesCount  json.Number               `json:"totalMessagesCount"`
	LatestMessages      []AbuseMessage            `json:"latestMessages"`
}

type AbuseAttachment struct {
	Id       string `json:"id"`
	MimeType string `json:"mimeType"`
	Filename string `json:"filename"`
}

type AbuseDetectedDomainName struct {
	Name        string   `json:"name"`
	IpAddresses []string `json:"ipAddresses"`
}

type AbuseReports struct {
	Metadata Metadata      `json:"_metadata"`
	Reports  []AbuseReport `json:"reports"`
}

type AbuseMessage struct {
	PostedBy   string          `json:"postedBy"`
	PostedAt   string          `json:"postedAt"`
	Body       string          `json:"body"`
	Attachment AbuseAttachment `json:"attachment"`
}

type AbuseMessages struct {
	Messages []AbuseMessage `json:"messages"`
	Metadata Metadata       `json:"_metadata"`
}

type AbuseResolution struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type AbuseResolutions struct {
	Resolutions []AbuseResolution `json:"resolutions"`
}

type AbuseListOptions struct {
	PaginationOptions
	Status []string `param:"status"`
}

func (aba AbuseApi) getPath(endpoint string) string {
	return "/abuse/" + ABUSE_API_VERSION + endpoint
}

func (aba AbuseApi) List(ctx context.Context, opts AbuseListOptions) (*AbuseReports, error) {
	path := aba.getPath("/reports")
	query := options.Encode(opts)
	result := &AbuseReports{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) Get(ctx context.Context, abuseReportId string) (*AbuseReport, error) {
	path := aba.getPath("/reports/" + abuseReportId)
	result := &AbuseReport{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) ListMessages(ctx context.Context, abuseReportId string, opts PaginationOptions) (*AbuseMessages, error) {
	path := aba.getPath("/reports/" + abuseReportId + "/messages")
	query := options.Encode(opts)
	result := &AbuseMessages{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) CreateMessage(ctx context.Context, abuseReportId string, body string) ([]string, error) {
	var result []string
	payload := map[string]string{body: body}
	path := aba.getPath("/reports/" + abuseReportId + "/messages")
	if err := doRequest(ctx, http.MethodPost, path, "", &result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) ListResolutionOptions(ctx context.Context, abuseReportId string) (*AbuseResolutions, error) {
	path := aba.getPath("/reports/" + abuseReportId + "/resolutions")
	result := &AbuseResolutions{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) Resolve(ctx context.Context, abuseReportId string, resolutions []string) error {
	payload := map[string][]string{"resolutions": resolutions}
	path := aba.getPath("/reports/" + abuseReportId + "/resolve")
	return doRequest(ctx, http.MethodPost, path, "", nil, payload)
}

// TODO
// func (aba AbuseApi) GetAbuseReportMessageAttachments() {}

// TODO
// func (aba AbuseApi) GetAbuseReportAttachments() {}
