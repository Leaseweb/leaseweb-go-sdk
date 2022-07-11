package leaseweb

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type LeasewebApi interface {
	SetVersion(version string)
	GetVersion() string
	GetPath() string
}

type leasewebRequest struct {
	response interface{}
	payload  interface{}
	method   string
	endpoint string
}

type LeasewebError struct {
	ErrorCode     string `json:"errorCode"`
	ErrorMessage  string `json:"errorMessage"`
	CorrelationId string `json:"correlationId"`
	UserMessage   string `json:"userMessage"`
	Reference     string `json:"reference"`
}

func (le *LeasewebError) Error() string {
	return le.ErrorMessage
}

type leasewebClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

var lswClient *leasewebClient

const (
	DEFAULT_BASE_URL = "https://api.staging.devleaseweb.com"
	GET              = "GET"
	POST             = "POST"
	PUT              = "PUT"
	DELETE           = "DELETE"
)

func InitLeasewebClient(key string) {
	lswClient = &leasewebClient{
		client: &http.Client{},
		apiKey: key,
	}
}

func getBaseUrl() string {
	if lswClient.baseUrl != "" {
		return lswClient.baseUrl
	}
	return DEFAULT_BASE_URL
}

func doRequest(r leasewebRequest) error {
	var payload io.Reader
	if r.method == POST || r.method == PUT {
		b, err := json.Marshal(r.payload)
		if err != nil {
			return err
		}
		payload = strings.NewReader(string(b))
	}

	req, err := http.NewRequest(r.method, getBaseUrl()+r.endpoint, payload)
	if err != nil {
		return err
	}

	req.Header.Add("x-lsw-auth", lswClient.apiKey)
	resp, err := lswClient.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		lswErr := &LeasewebError{}
		err = json.Unmarshal(respBody, lswErr)
		if err != nil {
			return err
		}
		return lswErr
	}

	err = json.Unmarshal(respBody, r.response)
	if err != nil {
		return err
	}

	return nil
}
