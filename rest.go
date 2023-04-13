package leaseweb

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const DEFAULT_BASE_URL = "https://api.leaseweb.com"
const CONTEXT_ACCOUNT_KEY = "CONTEXT_ACCOUNT_KEY"

var lswClient *client

type Account struct {
	ApiKey string
}

func getSelectedApiKey(ctx context.Context) string {

	account := Account{}
	ctxValue := ctx.Value(CONTEXT_ACCOUNT_KEY)
	if ctxValue != nil {
		account = ctxValue.(Account)
	}

	if account.ApiKey != "" {
		return account.ApiKey
	}
	return lswClient.defaultAccount.ApiKey
}

func SetContextWithAccount(ctx context.Context, account Account) context.Context {
	return context.WithValue(ctx, CONTEXT_ACCOUNT_KEY, account)
}

func SetDefaultAccount(account Account) {
	lswClient = &client{
		client:         &http.Client{},
		baseUrl:        DEFAULT_BASE_URL,
		defaultAccount: account,
	}
}

type client struct {
	client         *http.Client
	baseUrl        string
	defaultAccount Account
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

// Deprecated: Use SetDefaultAccount instead
func InitLeasewebClient(key string) {
	SetDefaultAccount(Account{ApiKey: key})
}

// Deprecated: This function will remove!
func SetBaseUrl(baseUrl string) {
	lswClient.baseUrl = baseUrl
}

func doRequest(ctx context.Context, method, path, query string, args ...interface{}) error {

	if lswClient == nil {
		return errors.New("leaseweb account not found! Please use `SetDefaultAccount` first")
	}

	url := lswClient.baseUrl + path
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

	req.Header.Add("x-lsw-auth", getSelectedApiKey(ctx))
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
