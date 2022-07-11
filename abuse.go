package leaseweb

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const ABUSE_API_DEFAULT_VERSION = "v1"

type AbuseApi struct {
	version string
}

func (aba *AbuseApi) SetVersion(version string) {
	aba.version = version
}

func (aba AbuseApi) GetVersion() string {
	if aba.version == "" {
		return ABUSE_API_DEFAULT_VERSION
	}
	return aba.version
}

func (aba AbuseApi) GetPath() string {
	return "/abuse"
}

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

func (aba AbuseApi) ListAbuseReports(args ...interface{}) (AbuseReports, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	abuseReports := &AbuseReports{}
	r := leasewebRequest{
		response: abuseReports,
		method:   GET,
		endpoint: aba.GetPath() + "/" + aba.GetVersion() + "/reports" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return AbuseReports{}, err
	}

	return *abuseReports, nil
}

func (aba AbuseApi) GetAbuseReport(abuseReportId string) (AbuseReport, error) {

	abuseReport := &AbuseReport{}
	r := leasewebRequest{
		response: abuseReport,
		method:   GET,
		endpoint: aba.GetPath() + "/" + aba.GetVersion() + "/reports/" + abuseReportId,
	}
	err := doRequest(r)
	if err != nil {
		return AbuseReport{}, err
	}

	return *abuseReport, nil
}

func (aba AbuseApi) GetAbuseReportMessages(abuseReportId string, args ...int) (AbuseMessages, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	abuseMessages := &AbuseMessages{}
	r := leasewebRequest{
		response: abuseMessages,
		method:   GET,
		endpoint: aba.GetPath() + "/" + aba.GetVersion() + "/reports/" + abuseReportId + "/messages" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return AbuseMessages{}, err
	}

	return *abuseMessages, nil
}

func (aba AbuseApi) CreateNewAbuseReportMessage(abuseReportId string, body string) ([]string, error) {
	var resp []string
	r := leasewebRequest{
		payload:  map[string]string{body: body},
		response: &resp,
		method:   POST,
		endpoint: aba.GetPath() + "/" + aba.GetVersion() + "/reports/" + abuseReportId + "/messages",
	}
	err := doRequest(r)
	if err != nil {
		return []string{}, err
	}

	return resp, nil
}

func (aba AbuseApi) ListResolutionOptions(abuseReportId string) (Resolutions, error) {
	resolutions := &Resolutions{}
	r := leasewebRequest{
		response: resolutions,
		method:   GET,
		endpoint: aba.GetPath() + "/" + aba.GetVersion() + "/reports/" + abuseReportId + "/resolutions",
	}
	err := doRequest(r)
	if err != nil {
		return Resolutions{}, err
	}

	return *resolutions, nil
}

func (aba AbuseApi) ResolveAbuseReport(abuseReportId string, resolutions []string) error {
	r := leasewebRequest{
		payload:  map[string][]string{"resolutions": resolutions},
		method:   POST,
		endpoint: aba.GetPath() + "/" + aba.GetVersion() + "/reports/" + abuseReportId + "/resolve",
	}
	err := doRequest(r)
	if err != nil {
		return err
	}

	return nil
}

// TODO
// func (aba AbuseApi) GetAbuseReportMessageAttachments() {}

// TODO
// func (aba AbuseApi) GetAbuseReportAttachments() {}
