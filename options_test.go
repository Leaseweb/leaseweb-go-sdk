package leaseweb

import (
	"testing"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

func TestEncode(t *testing.T) {
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
			expected: "ip=10.22.192.3&limit=10&macAddress=AA%3ABB%3ACC%3ADD%3AEE%3AFF&offset=0&privateNetworkCapable=true&privateNetworkEnabled=false&privateRackId=1&reference=test-reference&site=AMS-01",
		},
		{
			name: "Test Case 2: Limited options",
			opts: DedicatedServerListOptions{
				Offset:     Int(0),
				Limit:      Int(10),
				MacAddress: String(""),
				Site:       String("AMS-01"),
			},
			expected: "limit=10&macAddress=&offset=0&site=AMS-01",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := options.Encode(tc.opts)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
