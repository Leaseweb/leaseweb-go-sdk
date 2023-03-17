package leaseweb

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var lswClient *leasewebClient

var defaultErrors = map[int]ApiError{
	http.StatusBadRequest:          {Code: strconv.Itoa(http.StatusBadRequest), Message: "Bad Request"},
	http.StatusUnauthorized:        {Code: strconv.Itoa(http.StatusUnauthorized), Message: "Unauthorized"},
	http.StatusForbidden:           {Code: strconv.Itoa(http.StatusForbidden), Message: "Forbidden"},
	http.StatusNotFound:            {Code: strconv.Itoa(http.StatusNotFound), Message: "Not Found"},
	http.StatusServiceUnavailable:  {Code: strconv.Itoa(http.StatusServiceUnavailable), Message: "Service Unavailable"},
	http.StatusInternalServerError: {Code: strconv.Itoa(http.StatusInternalServerError), Message: "Internal Server Error"},
}

const DEFAULT_BASE_URL = "https://api.leaseweb.com"

type leasewebClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

type ApiError struct {
	Code          string              `json:"errorCode"`
	Message       string              `json:"errorMessage"`
	Details       map[string][]string `json:"errorDetails"`
	CorrelationId string              `json:"correlationId"`
	UserMessage   string              `json:"userMessage"`
	Reference     string              `json:"reference"`
}

func (le *ApiError) Error() string {
	return "leaseweb: " + le.Message
}

type DecodingError struct {
	Method string
	Url    string
	Err    error
}

func (errd *DecodingError) Error() string {
	return "leaseweb: decoding JSON response body failed (" + errd.Method + " " + errd.Url + ")"
}

type EncodingError struct {
	Method string
	Url    string
	Err    error
}

func (erre *EncodingError) Error() string {
	return "leaseweb: encoding JSON request body failed (" + erre.Method + " " + erre.Url + ")"
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

func doRequest(ctx context.Context, method string, path string, args ...interface{}) error {
	url := getBaseUrl() + path

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

	req, err := http.NewRequestWithContext(ctx, method, url, tmpPayload)
	if err != nil {
		return err
	}

	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		lswErr := &ApiError{}
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
