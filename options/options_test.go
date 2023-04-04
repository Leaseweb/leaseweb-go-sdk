package options

import (
	"testing"
)

type ListOptions struct {
	Offset                *int    `param:"offset"`
	MacAddress            *string `param:"macAddress"`
	PrivateNetworkCapable *bool   `param:"privateNetworkCapable"`
}

func TestEncode(t *testing.T) {
	testCases := []struct {
		name     string
		opts     ListOptions
		expected string
	}{
		{
			name: "Test Case 1: Full options",
			opts: func() ListOptions {
				macAddress := "AA:BB:CC:DD:EE:FF"
				offset := 0
				privateNetworkCapable := false
				return ListOptions{
					Offset:                &offset,
					MacAddress:            &macAddress,
					PrivateNetworkCapable: &privateNetworkCapable,
				}
			}(),
			expected: "macAddress=AA%3ABB%3ACC%3ADD%3AEE%3AFF&offset=0&privateNetworkCapable=false",
		},
		{
			name: "Limited options",
			opts: func() ListOptions {
				macAddress := ""
				offset := 0
				return ListOptions{
					Offset:     &offset,
					MacAddress: &macAddress,
				}
			}(),
			expected: "macAddress=&offset=0",
		},
		{
			name:     "Empty struct",
			opts:     ListOptions{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Encode(tc.opts)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
