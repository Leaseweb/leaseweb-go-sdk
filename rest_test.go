package leaseweb

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

type ctxT struct {
	ts            *httptest.Server
	oldHost       string
	oldHttpClient *http.Client
}

var ctx = ctxT{}

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
	InitLeasewebClient("test-api-key")
	os.Exit(m.Run())
}
