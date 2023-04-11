package options

import (
	"testing"
)

type PaginationOptions struct {
	Limit  *int `param:"limit"`
	Offset *int `param:"offset"`
}

type ListOptions struct {
	PaginationOptions
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
			name: "Full options",
			opts: func() ListOptions {
				macAddress := "AA:BB:CC:DD:EE:FF"
				offset := 0
				limit := 10
				privateNetworkCapable := false
				return ListOptions{
					PaginationOptions: PaginationOptions{
						Limit:  &limit,
						Offset: &offset,
					},
					MacAddress:            &macAddress,
					PrivateNetworkCapable: &privateNetworkCapable,
				}
			}(),
			expected: "limit=10&macAddress=AA%3ABB%3ACC%3ADD%3AEE%3AFF&offset=0&privateNetworkCapable=false",
		},
		{
			name: "Limited options",
			opts: func() ListOptions {
				offset := 0
				limit := 10
				return ListOptions{
					PaginationOptions: PaginationOptions{
						Limit:  &limit,
						Offset: &offset,
					},
				}
			}(),
			expected: "limit=10&offset=0",
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
