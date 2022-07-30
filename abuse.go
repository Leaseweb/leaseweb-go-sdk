package leaseweb

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const ABUSE_API_VERSION = "v1"

type AbuseApi struct{}

type AbuseReport struct {
	Id                  string               `json:"id"`
	Subject             string               `json:"subject"`
	Status              string               `json:"status"`
	Reopened            bool                 `json:"reopened"`
	ReportedAt          string               `json:"reportedAt"`
	UpdatedAt           string               `json:"updatedAt"`
	Notifier            string               `json:"notifier"`
	CustomerId          string               `json:"customerId"`
	LegalEntityId       string               `json:"legalEntityId"`
	Deadline            string               `json:"deadline"`
	Body                string               `json:"body"`
	DetectedIpAddresses []string             `json:"detectedIpAddresses"`
	DetectedDomainNames []DetectedDomainName `json:"detectedDomainNames"`
	Attachments         []Attachment         `json:"attachments"`
	TotalMessagesCount  int                  `json:"totalMessagesCount"`
	LatestMessages      []AbuseMessage       `json:"latestMessages"`
}

type Attachment struct {
	Id       string `json:"id"`
	MimeType string `json:"mimeType"`
	Filename string `json:"filename"`
}

type DetectedDomainName struct {
	Name        string   `json:"name"`
	IpAddresses []string `json:"ipAddresses"`
}

type AbuseReports struct {
	Metadata     Metadata      `json:"_metadata"`
	AbuseReports []AbuseReport `json:"reports"`
}

type AbuseMessage struct {
	PostedBy   string     `json:"postedBy"`
	PostedAt   string     `json:"postedAt"`
	Body       string     `json:"body"`
	Attachment Attachment `json:"attachment"`
}

type AbuseMessages struct {
	Messages []AbuseMessage `json:"messages"`
	Metadata Metadata       `json:"_metadata"`
}

type Resolution struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type Resolutions struct {
	Resolutions []Resolution `json:"resolutions"`
}

func (aba AbuseApi) getPath(endpoint string) string {
	return "/abuse/" + ABUSE_API_VERSION + endpoint
}

func (aba AbuseApi) ListAbuseReports(args ...interface{}) (*AbuseReports, error) {
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

	path := aba.getPath("/reports?" + v.Encode())
	result := &AbuseReports{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) GetAbuseReport(abuseReportId string) (*AbuseReport, error) {
	path := aba.getPath("/reports/" + abuseReportId)
	result := &AbuseReport{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) GetAbuseReportMessages(abuseReportId string, args ...int) (*AbuseMessages, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := aba.getPath("/reports/" + abuseReportId + "/messages" + v.Encode())
	result := &AbuseMessages{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) CreateNewAbuseReportMessage(abuseReportId string, body string) ([]string, error) {
	var result []string
	payload := map[string]string{body: body}
	path := aba.getPath("/reports/" + abuseReportId + "/messages")
	if err := doRequest(POST, path, &result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) ListResolutionOptions(abuseReportId string) (*Resolutions, error) {
	path := aba.getPath("/reports/" + abuseReportId + "/resolutions")
	result := &Resolutions{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (aba AbuseApi) ResolveAbuseReport(abuseReportId string, resolutions []string) error {
	payload := map[string][]string{"resolutions": resolutions}
	path := aba.getPath("/reports/" + abuseReportId + "/resolve")
	if err := doRequest(POST, path, nil, payload); err != nil {
		return err
	}
	return nil
}

// TODO
// func (aba AbuseApi) GetAbuseReportMessageAttachments() {}

// TODO
// func (aba AbuseApi) GetAbuseReportAttachments() {}
