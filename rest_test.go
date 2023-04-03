package leaseweb

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
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
	InitLeasewebClient(testApiKey)
	os.Exit(m.Run())
}

func TestStructToURLValues(t *testing.T) {
	testCases := []struct {
		name     string
		opts     DedicatedServerListOptions
		expected string
	}{
		{
			name: "Test Case 1: Full options",
			opts: DedicatedServerListOptions{
				Offset:                Int(0),
				Limit:                 Int(10),
				IP:                    String("10.22.192.3"),
				MacAddress:            String("AA:BB:CC:DD:EE:FF"),
				Site:                  String("AMS-01"),
				PrivateRackID:         Int(1),
				Reference:             String("test-reference"),
				PrivateNetworkCapable: Bool(true),
				PrivateNetworkEnabled: Bool(false),
			},
			expected: "offset=0&limit=10&ip=10.22.192.3&macAddress=AA%3ABB%3ACC%3ADD%3AEE%3AFF&site=AMS-01&privateRackId=1&reference=test-reference&privateNetworkCapable=true&privateNetworkEnabled=false",
		},
		{
			name: "Test Case 2: Limited options",
			opts: DedicatedServerListOptions{
				Offset:     Int(0),
				Limit:      Int(10),
				MacAddress: String(""),
				Site:       String("AMS-01"),
			},
			expected: "offset=0&limit=10&macAddress=&site=AMS-01",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := structToURLValues(tc.opts)

			expectedValues, _ := url.ParseQuery(tc.expected)
			resultValues, _ := url.ParseQuery(result)

			if !reflect.DeepEqual(expectedValues, resultValues) {
				t.Errorf("Expected '%v', but got '%v'", expectedValues, resultValues)
			}
		})
	}
}
