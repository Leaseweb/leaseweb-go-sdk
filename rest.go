package leaseweb

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var lswClient *LeasewebClient

const DEFAULT_BASE_URL = "https://api.leaseweb.com"

type LeasewebClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

type ApiContext struct {
	Method string
	Url    string
}

type ApiError struct {
	ApiContext
	Code          string              `json:"errorCode"`
	Message       string              `json:"errorMessage"`
	Details       map[string][]string `json:"errorDetails"`
	CorrelationId string              `json:"correlationId"`
	UserMessage   string              `json:"userMessage"`
	Reference     string              `json:"reference"`
}

func (erra *ApiError) Error() string {
	return "leaseweb: " + erra.Message

}

type DecodingError struct {
	ApiContext
	Err error
}

func (errd *DecodingError) Error() string {
	return "leaseweb: decoding JSON response body failed (" + errd.Err.Error() + ")"
}

type EncodingError struct {
	ApiContext
	Err error
}

func (erre *EncodingError) Error() string {
	return "leaseweb: encoding JSON request body failed (" + erre.Err.Error() + ")"
}

func InitLeasewebClient(key string) {
	lswClient = NewLeaseWebClient(key)
}

type option func(client *LeasewebClient)

func WithHTTPClient(client *http.Client) option {
	return func(c *LeasewebClient) {
		c.client = client
	}
}

func NewLeaseWebClient(key string, opts ...option) *LeasewebClient {
	client := &LeasewebClient{
		client: &http.Client{},
		apiKey: key,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func getClient(c *LeasewebClient) *LeasewebClient {
	if c != nil {
		return c
	}
	return lswClient
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

func (lsw *LeasewebClient) doRequest(ctx context.Context, method, path, query string, args ...interface{}) error {
	url := getBaseUrl() + path
	if query != "" {
		url += "?" + query
	}

	apiContext := ApiContext{method, url}

	var tmpPayload io.Reader
	if method == http.MethodPost || method == http.MethodPut {
		if len(args) > 1 {
			b, err := json.Marshal(args[1])
			if err != nil {
				return &EncodingError{apiContext, err}
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

	req.Header.Add("x-lsw-auth", lsw.apiKey)
	resp, err := lsw.client.Do(req)
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
		statusCode, statusMessage, ok := strings.Cut(resp.Status, " ")
		if !ok {
			statusCode = strconv.Itoa(resp.StatusCode)
			statusMessage = "An error occurred"
		}

		apiErr := &ApiError{
			ApiContext: apiContext,
			Code:       statusCode,
			Message:    statusMessage,
		}
		json.Unmarshal(respBody, apiErr)
		return apiErr
	}

	if len(args) > 0 {
		if err = json.Unmarshal(respBody, &args[0]); err != nil {
			return &DecodingError{apiContext, err}
		}
	}
	return nil
}
