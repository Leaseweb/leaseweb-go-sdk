package options

import (
	"testing"
)

type BasicOptions struct {
	Count                 *int    `param:"count"`
	MacAddress            *string `param:"macAddress"`
	PrivateNetworkCapable *bool   `param:"privateNetworkCapable"`
}

type ExtraOptions struct {
	Limit  *int `param:"limit"`
	Offset *int `param:"offset"`
}

type AdvancedOptions struct {
	ExtraOptions
	MacAddress            *string  `param:"macAddress"`
	PrivateNetworkCapable *bool    `param:"privateNetworkCapable"`
	Status                []string `param:"status"`
}

func TestEncode(t *testing.T) {
	testCases := []struct {
		name     string
		opts     BasicOptions
		expected string
	}{
		{
			name: "BasicOptions Full options",
			opts: func() BasicOptions {
				macAddress := "AA:BB:CC:DD:EE:FF"
				count := 0
				privateNetworkCapable := false
				return BasicOptions{
					Count:                 &count,
					MacAddress:            &macAddress,
					PrivateNetworkCapable: &privateNetworkCapable,
				}
			}(),
			expected: "count=0&macAddress=AA%3ABB%3ACC%3ADD%3AEE%3AFF&privateNetworkCapable=false",
		},
		{
			name: "BasicOptions Limited options",
			opts: func() BasicOptions {
				macAddress := ""
				count := 0
				return BasicOptions{
					Count:      &count,
					MacAddress: &macAddress,
				}
			}(),
			expected: "count=0&macAddress=",
		},
		{
			name:     "BasicOptions Empty struct",
			opts:     BasicOptions{},
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

func TestEncodeWithExtraOptions(t *testing.T) {
	testCases := []struct {
		name     string
		opts     AdvancedOptions
		expected string
	}{
		{
			name: "ExtraOptions Full options",
			opts: func() AdvancedOptions {
				macAddress := "AA:BB:CC:DD:EE:FF"
				offset := 0
				limit := 10
				privateNetworkCapable := false
				return AdvancedOptions{
					ExtraOptions: ExtraOptions{
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
			name: "ExtraOptions Limited options",
			opts: func() AdvancedOptions {
				offset := 0
				limit := 10
				return AdvancedOptions{
					ExtraOptions: ExtraOptions{
						Limit:  &limit,
						Offset: &offset,
					},
				}
			}(),
			expected: "limit=10&offset=0",
		},
		{
			name: "Empty ExtraOptions",
			opts: func() AdvancedOptions {
				macAddress := "AA:BB:CC:DD:EE:FF"
				privateNetworkCapable := false
				return AdvancedOptions{
					MacAddress:            &macAddress,
					PrivateNetworkCapable: &privateNetworkCapable,
				}
			}(),
			expected: "macAddress=AA%3ABB%3ACC%3ADD%3AEE%3AFF&privateNetworkCapable=false",
		},
		{
			name: "CommaSeparatedValues",
			opts: func() AdvancedOptions {
				status := []string{"OPEN", "WAITING"}
				limit := 1
				return AdvancedOptions{
					ExtraOptions: ExtraOptions{
						Limit: &limit,
					},
					Status: status,
				}
			}(),
			expected: "limit=1&status=OPEN%2CWAITING",
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
