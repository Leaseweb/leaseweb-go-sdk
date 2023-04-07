package leaseweb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const ABUSE_API_VERSION = "v1"

type AbuseApi struct {
	Client *LeasewebClient
}

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

func (aba AbuseApi) getPath(endpoint string) string {
	return "/abuse/" + ABUSE_API_VERSION + endpoint
}

func (aba AbuseApi) List(ctx context.Context, args ...interface{}) (*AbuseReports, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		s := reflect.ValueOf(args[1])
		var statuses []string
		for i := 0; i < s.Len(); i++ {
			statuses = append(statuses, s.Index(i).String())
		}
		v.Add("status", strings.Join(statuses, ","))
	}
	if len(args) >= 3 {
		v.Add("limit", fmt.Sprint(args[2]))
	}

	path := aba.getPath("/reports")
	query := v.Encode()
	result := &AbuseReports{}
	if err := getClient(aba.Client).
		doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) Get(ctx context.Context, abuseReportId string) (*AbuseReport, error) {
	path := aba.getPath("/reports/" + abuseReportId)
	result := &AbuseReport{}
	if err := getClient(aba.Client).
		doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) ListMessages(ctx context.Context, abuseReportId string, args ...int) (*AbuseMessages, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := aba.getPath("/reports/" + abuseReportId + "/messages")
	query := v.Encode()
	result := &AbuseMessages{}
	if err := getClient(aba.Client).
		doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) CreateMessage(ctx context.Context, abuseReportId string, body string) ([]string, error) {
	var result []string
	payload := map[string]string{body: body}
	path := aba.getPath("/reports/" + abuseReportId + "/messages")
	if err := getClient(aba.Client).
		doRequest(ctx, http.MethodPost, path, "", &result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) ListResolutionOptions(ctx context.Context, abuseReportId string) (*AbuseResolutions, error) {
	path := aba.getPath("/reports/" + abuseReportId + "/resolutions")
	result := &AbuseResolutions{}
	if err := getClient(aba.Client).
		doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) Resolve(ctx context.Context, abuseReportId string, resolutions []string) error {
	payload := map[string][]string{"resolutions": resolutions}
	path := aba.getPath("/reports/" + abuseReportId + "/resolve")
	return getClient(aba.Client).
		doRequest(ctx, http.MethodPost, path, "", nil, payload)
}

// TODO
// func (aba AbuseApi) GetAbuseReportMessageAttachments() {}

// TODO
// func (aba AbuseApi) GetAbuseReportAttachments() {}
