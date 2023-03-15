package leaseweb

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var lswClient *leasewebClient

var defaultErrors = map[int]LeasewebError{
	http.StatusBadRequest:          {ErrorCode: strconv.Itoa(http.StatusBadRequest), ErrorMessage: "Bad Request"},
	http.StatusUnauthorized:        {ErrorCode: strconv.Itoa(http.StatusUnauthorized), ErrorMessage: "Unauthorized"},
	http.StatusForbidden:           {ErrorCode: strconv.Itoa(http.StatusForbidden), ErrorMessage: "Forbidden"},
	http.StatusNotFound:            {ErrorCode: strconv.Itoa(http.StatusNotFound), ErrorMessage: "Not Found"},
	http.StatusServiceUnavailable:  {ErrorCode: strconv.Itoa(http.StatusServiceUnavailable), ErrorMessage: "Service Unavailable"},
	http.StatusInternalServerError: {ErrorCode: strconv.Itoa(http.StatusInternalServerError), ErrorMessage: "Internal Server Error"},
}

const DEFAULT_BASE_URL = "https://api.leaseweb.com"

type leasewebClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

type LeasewebError struct {
	ErrorCode     string              `json:"errorCode"`
	ErrorMessage  string              `json:"errorMessage"`
	ErrorDetails  map[string][]string `json:"errorDetails"`
	CorrelationId string              `json:"correlationId"`
	UserMessage   string              `json:"userMessage"`
	Reference     string              `json:"reference"`
}

func (le *LeasewebError) Error() string {
	return le.ErrorMessage
}

type DecodingError struct {
	Method string
	Url    string
	Err    error
}

func (errd *DecodingError) Error() string {
	return "Decoding JSON response body failed (" + errd.Method + " " + errd.Url + ")"
}

type EncodingError struct {
	Method string
	Url    string
	Err    error
}

func (erre *EncodingError) Error() string {
	return "Encoding JSON request body failed (" + erre.Method + " " + erre.Url + ")"
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
	url := getBaseUrl() + endpoint

	var tmpPayload io.Reader
	if method == http.MethodPost || method == http.MethodPut {
		if len(args) > 1 {
			b, err := json.Marshal(args[1])
			if err != nil {
				return &EncodingError{Method: method, Url: url, Err: err}
			}
			tmpPayload = strings.NewReader(string(b))
		}
	}

	req, err := http.NewRequest(method, url, tmpPayload)
	if err != nil {
		return &LeasewebError{ErrorCode: "0", ErrorMessage: err.Error()}
	}

	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Add("x-lsw-auth", lswClient.apiKey)
	resp, err := lswClient.client.Do(req)
	if err != nil {
		return &LeasewebError{ErrorCode: "0", ErrorMessage: err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &LeasewebError{ErrorCode: "0", ErrorMessage: err.Error()}
	}

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		lswErr := &LeasewebError{}
		if err = json.Unmarshal(respBody, lswErr); err != nil {
			if defaultError, ok := defaultErrors[resp.StatusCode]; ok {
				return &defaultError
			}
			return &DecodingError{Method: method, Url: url, Err: err}
		}
		return lswErr
	}

	if len(args) > 0 {
		if err = json.Unmarshal(respBody, &args[0]); err != nil {
			return &DecodingError{Method: method, Url: url, Err: err}
		}
	}
	return nil
}
