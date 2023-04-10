package leaseweb

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ctx = ctxT{}
var testApiKey = "test-api-key"

type serverErrorTest struct {
	Title         string
	MockServer    func(http.ResponseWriter, *http.Request)
	FunctionCall  func() (interface{}, error)
	ExpectedError ApiError
}

type ctxT struct {
	ts            *httptest.Server
	oldHost       string
	oldHttpClient *http.Client
}

func assertServerErrorTests(t *testing.T, serverErrorTests []serverErrorTest) {
	for _, serverErrorTest := range serverErrorTests {
		setup(serverErrorTest.MockServer)
		defer teardown()
		resp, err := serverErrorTest.FunctionCall()
		assert := assert.New(t)
		assert.Empty(resp)
		assert.NotNil(err)
		assert.Equal(err.Error(), "leaseweb: "+serverErrorTest.ExpectedError.Message)
		lswErr, ok := err.(*ApiError)
		assert.Equal(true, ok)
		assert.Equal(lswErr.Message, serverErrorTest.ExpectedError.Message)
		assert.Equal(lswErr.CorrelationId, serverErrorTest.ExpectedError.CorrelationId)
		assert.Equal(lswErr.Code, serverErrorTest.ExpectedError.Code)
		assert.Equal(lswErr.Reference, serverErrorTest.ExpectedError.Reference)
		assert.Equal(lswErr.UserMessage, serverErrorTest.ExpectedError.UserMessage)
		assert.Equal(lswErr.Details, serverErrorTest.ExpectedError.Details)
	}
}

func setup(handler http.HandlerFunc) *httptest.Server {
	ts := httptest.NewTLSServer(handler)
	ctx.ts = ts
	_url, err := url.Parse(ts.URL)
	if err != nil {
		panic(err)
	}

	ctx.oldHost = lswClient.baseUrl
	ctx.oldHttpClient = lswClient.client

	lswClient.baseUrl = "https://" + _url.Host
	lswClient.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return ts
}

func teardown() {
	ctx.ts.Close()
	lswClient.baseUrl = ctx.oldHost
	lswClient.client = ctx.oldHttpClient
}

func TestMain(m *testing.M) {
	SetDefaultAccount(Account{ApiKey: testApiKey})
	os.Exit(m.Run())
}

func TestSetContextWithAccount(t *testing.T) {
	newApiKey := "new account api key"
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, newApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	newCtx := SetContextWithAccount(ctx, Account{ApiKey: newApiKey})
	_, err := customerAccountApi.Get(newCtx)

	assert := assert.New(t)
	assert.Nil(err)
}

func TestSetContextWithAccountWithOldCtx(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	SetContextWithAccount(ctx, Account{ApiKey: "api key"})
	_, err := customerAccountApi.Get(ctx)

	assert := assert.New(t)
	assert.Nil(err)
}
