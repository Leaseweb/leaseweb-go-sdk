package leaseweb

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var lswClient *leasewebClient

const DEFAULT_BASE_URL = "https://api.leaseweb.com"

type leasewebClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
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

func InitLeasewebClient(key string) {
	lswClient = &leasewebClient{
		client: &http.Client{},
		apiKey: key,
	}
}

func SetBaseUrl(baseUrl string) {
	lswClient.baseUrl = baseUrl
}

func getBaseUrl() string {
	if lswClient.baseUrl != "" {
		return lswClient.baseUrl
	}
	return DEFAULT_BASE_URL
}

func doRequest(method string, endpoint string, args ...interface{}) error {
	var tmpPayload io.Reader
	if method == http.MethodPost || method == http.MethodPut {
		if len(args) > 1 {
			b, err := json.Marshal(args[1])
			if err != nil {
				return err
			}
			tmpPayload = strings.NewReader(string(b))
		}
	}

	req, err := http.NewRequest(method, getBaseUrl()+endpoint, tmpPayload)
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
		if err = json.Unmarshal(respBody, lswErr); err != nil {
			return err
		}
		return lswErr
	}

	if len(args) > 0 {
		if err = json.Unmarshal(respBody, &args[0]); err != nil {
			return err
		}
	}
	return nil
}
