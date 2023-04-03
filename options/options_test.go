package options

import (
	"net/url"
	"reflect"
	"testing"
)

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
