package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostingListDomains(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"domains": [
				{
					"domainName": "example.com",
					"suspended": true,
					"dnsOnly": false,
					"status": "SUSPENDED",
					"contractStartDate": "2017-04-01T00:00:00+02:00",
					"contractEndDate": "2018-04-01T00:00:00+02:00",
					"_links": {
						"self": {
							"href": "/domains/example.com"
						},
						"collection": {
							"href": "/domains"
						},
						"catchAll": {
							"href": "/domains/example.com/catchAll"
						},
						"contacts": {
							"href": "/domains/example.com/contacts"
						},
						"nameservers": {
							"href": "/domains/example.com/nameservers"
						},
						"emailAliases": {
							"href": "/domains/example.com/emailAliases"
						},
						"forwards": {
							"href": "/domains/example.com/forwards"
						},
						"mailboxes": {
							"href": "/domains/example.com/mailboxes"
						},
						"resourceRecordSets": {
							"href": "/domains/example.com/resourceRecordSets"
						},
						"validateZone": {
							"href": "/domains/example.com/validateZone"
						}
					}
				},
				{
					"domainName": "example.org",
					"suspended": false,
					"dnsOnly": true,
					"status": "ACTIVE",
					"contractStartDate": "2015-01-01T00:00:00+00:00",
					"_links": {
						"self": {
							"href": "/domains/example.org"
						},
						"collection": {
							"href": "/domains"
						},
						"catchAll": {
							"href": "/domains/example.org/catchAll"
						},
						"contacts": {
							"href": "/domains/example.com/contacts"
						},
						"nameservers": {
							"href": "/domains/example.com/nameservers"
						},
						"emailAliases": {
							"href": "/domains/example.org/emailAliases"
						},
						"forwards": {
							"href": "/domains/example.org/forwards"
						},
						"mailboxes": {
							"href": "/domains/example.org/mailboxes"
						},
						"resourceRecordSets": {
							"href": "/domains/example.org/resourceRecordSets"
						},
						"validateZone": {
							"href": "/domains/example.org/validateZone"
						}
					}
				}
			],
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	response, err := hostingApi.ListDomains(ctx, HostingListDomainsOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Domains), 2)

	domain1 := response.Domains[0]
	assert.Equal(domain1.DomainName, "example.com")
	assert.Equal(domain1.Suspended, true)
	assert.Equal(domain1.DnsOnly, false)
	assert.Equal(domain1.Status, "SUSPENDED")
	assert.Equal(domain1.ContractStartDate, "2017-04-01T00:00:00+02:00")
	assert.Equal(domain1.ContractEndDate, "2018-04-01T00:00:00+02:00")
	assert.Equal(domain1.Link.Self.Href, "/domains/example.com")
	assert.Equal(domain1.Link.Collection.Href, "/domains")
	assert.Equal(domain1.Link.CatchAll.Href, "/domains/example.com/catchAll")
	assert.Equal(domain1.Link.Contacts.Href, "/domains/example.com/contacts")
	assert.Equal(domain1.Link.Nameservers.Href, "/domains/example.com/nameservers")
	assert.Equal(domain1.Link.EmailAliases.Href, "/domains/example.com/emailAliases")
	assert.Equal(domain1.Link.Forwards.Href, "/domains/example.com/forwards")
	assert.Equal(domain1.Link.Mailboxes.Href, "/domains/example.com/mailboxes")
	assert.Equal(domain1.Link.ResourceRecordSets.Href, "/domains/example.com/resourceRecordSets")
	assert.Equal(domain1.Link.ValidateZone.Href, "/domains/example.com/validateZone")

	domain2 := response.Domains[1]
	assert.Equal(domain2.DomainName, "example.org")
	assert.Equal(domain2.Suspended, false)
	assert.Equal(domain2.DnsOnly, true)
	assert.Equal(domain2.Status, "ACTIVE")
	assert.Equal(domain2.ContractStartDate, "2015-01-01T00:00:00+00:00")
	assert.Equal(domain1.Link.Self.Href, "/domains/example.com")
	assert.Equal(domain1.Link.Collection.Href, "/domains")
	assert.Equal(domain1.Link.CatchAll.Href, "/domains/example.com/catchAll")
	assert.Equal(domain1.Link.Contacts.Href, "/domains/example.com/contacts")
	assert.Equal(domain1.Link.Nameservers.Href, "/domains/example.com/nameservers")
	assert.Equal(domain1.Link.EmailAliases.Href, "/domains/example.com/emailAliases")
	assert.Equal(domain1.Link.Forwards.Href, "/domains/example.com/forwards")
	assert.Equal(domain1.Link.Mailboxes.Href, "/domains/example.com/mailboxes")
	assert.Equal(domain1.Link.ResourceRecordSets.Href, "/domains/example.com/resourceRecordSets")
	assert.Equal(domain1.Link.ValidateZone.Href, "/domains/example.com/validateZone")
}

func TestHostingListDomainsPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"domains": [
				{
					"domainName": "example.com",
					"suspended": true,
					"dnsOnly": false,
					"status": "SUSPENDED",
					"contractStartDate": "2017-04-01T00:00:00+02:00",
					"contractEndDate": "2018-04-01T00:00:00+02:00",
					"_links": {
						"self": {
							"href": "/domains/example.com"
						},
						"collection": {
							"href": "/domains"
						},
						"catchAll": {
							"href": "/domains/example.com/catchAll"
						},
						"contacts": {
							"href": "/domains/example.com/contacts"
						},
						"nameservers": {
							"href": "/domains/example.com/nameservers"
						},
						"emailAliases": {
							"href": "/domains/example.com/emailAliases"
						},
						"forwards": {
							"href": "/domains/example.com/forwards"
						},
						"mailboxes": {
							"href": "/domains/example.com/mailboxes"
						},
						"resourceRecordSets": {
							"href": "/domains/example.com/resourceRecordSets"
						},
						"validateZone": {
							"href": "/domains/example.com/validateZone"
						}
					}
				}
			],
			"_metadata": {
				"totalCount": 1,
				"limit": 10,
				"offset": 1
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	opts := HostingListDomainsOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(1),
		},
	}
	response, err := hostingApi.ListDomains(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Domains), 1)

	domain1 := response.Domains[0]
	assert.Equal(domain1.DomainName, "example.com")
	assert.Equal(domain1.Suspended, true)
	assert.Equal(domain1.DnsOnly, false)
	assert.Equal(domain1.Status, "SUSPENDED")
	assert.Equal(domain1.ContractStartDate, "2017-04-01T00:00:00+02:00")
	assert.Equal(domain1.ContractEndDate, "2018-04-01T00:00:00+02:00")
	assert.Equal(domain1.Link.Self.Href, "/domains/example.com")
	assert.Equal(domain1.Link.Collection.Href, "/domains")
	assert.Equal(domain1.Link.CatchAll.Href, "/domains/example.com/catchAll")
	assert.Equal(domain1.Link.Contacts.Href, "/domains/example.com/contacts")
	assert.Equal(domain1.Link.Nameservers.Href, "/domains/example.com/nameservers")
	assert.Equal(domain1.Link.EmailAliases.Href, "/domains/example.com/emailAliases")
	assert.Equal(domain1.Link.Forwards.Href, "/domains/example.com/forwards")
	assert.Equal(domain1.Link.Mailboxes.Href, "/domains/example.com/mailboxes")
	assert.Equal(domain1.Link.ResourceRecordSets.Href, "/domains/example.com/resourceRecordSets")
	assert.Equal(domain1.Link.ValidateZone.Href, "/domains/example.com/validateZone")
}

func TestHostingListDomainsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListDomains(ctx, HostingListDomainsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListDomains(ctx, HostingListDomainsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListDomains(ctx, HostingListDomainsOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetDomain(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"domainName": "example.com",
			"suspended": true,
			"dnsOnly": false,
			"status": "SUSPENDED",
			"contractStartDate": "2017-04-01T00:00:00+02:00",
			"contractEndDate": "2018-04-01T00:00:00+02:00",
			"_links": {
				"self": {
					"href": "/domains/example.com"
				},
				"collection": {
					"href": "/domains"
				},
				"catchAll": {
					"href": "/domains/example.com/catchAll"
				},
				"contacts": {
					"href": "/domains/example.com/contacts"
				},
				"nameservers": {
					"href": "/domains/example.com/nameservers"
				},
				"emailAliases": {
					"href": "/domains/example.com/emailAliases"
				},
				"forwards": {
					"href": "/domains/example.com/forwards"
				},
				"mailboxes": {
					"href": "/domains/example.com/mailboxes"
				},
				"resourceRecordSets": {
					"href": "/domains/example.com/resourceRecordSets"
				},
				"validateZone": {
					"href": "/domains/example.com/validateZone"
				}
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	domain, err := hostingApi.GetDomain(ctx, "exmple.com")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(domain.DomainName, "example.com")
	assert.Equal(domain.Suspended, true)
	assert.Equal(domain.DnsOnly, false)
	assert.Equal(domain.Status, "SUSPENDED")
	assert.Equal(domain.ContractStartDate, "2017-04-01T00:00:00+02:00")
	assert.Equal(domain.ContractEndDate, "2018-04-01T00:00:00+02:00")
	assert.Equal(domain.Link.Self.Href, "/domains/example.com")
	assert.Equal(domain.Link.Collection.Href, "/domains")
	assert.Equal(domain.Link.CatchAll.Href, "/domains/example.com/catchAll")
	assert.Equal(domain.Link.Contacts.Href, "/domains/example.com/contacts")
	assert.Equal(domain.Link.Nameservers.Href, "/domains/example.com/nameservers")
	assert.Equal(domain.Link.EmailAliases.Href, "/domains/example.com/emailAliases")
	assert.Equal(domain.Link.Forwards.Href, "/domains/example.com/forwards")
	assert.Equal(domain.Link.Mailboxes.Href, "/domains/example.com/mailboxes")
	assert.Equal(domain.Link.ResourceRecordSets.Href, "/domains/example.com/resourceRecordSets")
	assert.Equal(domain.Link.ValidateZone.Href, "/domains/example.com/validateZone")
}

func TestHostingGetDomainServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDomain(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Range with id 88.17.0.0_17 does not exist"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDomain(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "39e010ed-0e93-42c3-c28f-3ffc373553d5",
				Code:          "404",
				Message:       "Range with id 88.17.0.0_17 does not exist",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDomain(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDomain(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetAvailability(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"domainName": "example.com",
			"available": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/available"
				}
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	domain, err := hostingApi.GetAvailability(ctx, "exmple.com")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(domain.DomainName, "example.com")
	assert.Equal(domain.Available, true)
	assert.Equal(domain.Link.Self.Href, "/domains/example.com/available")
}

func TestHostingGetAvailabilityServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAvailability(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Range with id 88.17.0.0_17 does not exist"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAvailability(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "39e010ed-0e93-42c3-c28f-3ffc373553d5",
				Code:          "404",
				Message:       "Range with id 88.17.0.0_17 does not exist",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAvailability(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAvailability(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingListNameservers(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"nameservers": [
				{
					"name": "ns1.example.com"
				},
				{
					"name": "ns2.example.net"
				}
			],
			"status": "Nameservers update request is accepted. This request will be processed after 24 hours.",
			"_links": {
				"self": {
					"href": "/domains/{domainName}/nameservers"
				}
			},
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	response, err := hostingApi.ListNameservers(ctx, "example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Nameservers), 2)

	assert.Equal(response.Nameservers[0].Name, "ns1.example.com")
	assert.Equal(response.Nameservers[1].Name, "ns2.example.net")
	assert.Equal(response.Status, "Nameservers update request is accepted. This request will be processed after 24 hours.")
	assert.Equal(response.Link.Self.Href, "/domains/{domainName}/nameservers")
}

func TestHostingListNameserversPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"nameservers": [
				{
					"name": "ns1.example.com"
				}
			],
			"status": "Nameservers update request is accepted. This request will be processed after 24 hours.",
			"_links": {
				"self": {
					"href": "/domains/{domainName}/nameservers"
				}
			},
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	response, err := hostingApi.ListNameservers(ctx, "example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Nameservers), 1)

	assert.Equal(response.Nameservers[0].Name, "ns1.example.com")
	assert.Equal(response.Status, "Nameservers update request is accepted. This request will be processed after 24 hours.")
	assert.Equal(response.Link.Self.Href, "/domains/{domainName}/nameservers")
}

func TestHostingListNameserversServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return HostingApi{}.ListNameservers(ctx, "example.com", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return HostingApi{}.ListNameservers(ctx, "example.com", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return HostingApi{}.ListNameservers(ctx, "example.com", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateNameservers(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"nameservers": [
				{
					"name": "ns1.example.com"
				},
				{
					"name": "ns2.example.com"
				}
			],
			"status": "Nameservers update request is accepted. This request will be processed after 24 hours.",
			"_links": {
				"self": {
					"href": "/domains/{domainName}/nameservers"
				}
			},
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := HostingApi{}.UpdateNameservers(ctx, "example.com", []string{"ns1.example.com", "ns2.example.com"})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Nameservers), 2)

	assert.Equal(response.Nameservers[0].Name, "ns1.example.com")
	assert.Equal(response.Nameservers[1].Name, "ns2.example.com")
	assert.Equal(response.Status, "Nameservers update request is accepted. This request will be processed after 24 hours.")
	assert.Equal(response.Link.Self.Href, "/domains/{domainName}/nameservers")
}

func TestHostingUpdateNameserversServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateNameservers(ctx, "example.com", []string{"ns1.example.com", "ns2.example.com"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateNameservers(ctx, "example.com", []string{"ns1.example.com", "ns2.example.com"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateNameservers(ctx, "example.com", []string{"ns1.example.com", "ns2.example.com"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateNameservers(ctx, "example.com", []string{"ns1.example.com", "ns2.example.com"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetGetDnsSecurity(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"status": "ENABLED",
			"nsec3": {
			  	"active": true
			},
			"keyRollover": {
			  	"status": "INITIALIZED",
			  	"type": "KSK",
			  	"message": "It would take 25 hours to complete the process from the time its been initialized"
			},
			"_links": {
			  	"self": {
					"href": "/domains/example.com/dnssec"
			  	}
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	resp, err := hostingApi.GetDnsSecurity(ctx, "exmple.com")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Status, "ENABLED")
	assert.Equal(resp.Nsec.Active, true)
	assert.Equal(resp.KeyRollover.Message, "It would take 25 hours to complete the process from the time its been initialized")
	assert.Equal(resp.KeyRollover.Type, "KSK")
	assert.Equal(resp.KeyRollover.Status, "INITIALIZED")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/dnssec")
}

func TestHostingGetGetDnsSecurityServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDnsSecurity(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Range with id 88.17.0.0_17 does not exist"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDnsSecurity(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "39e010ed-0e93-42c3-c28f-3ffc373553d5",
				Code:          "404",
				Message:       "Range with id 88.17.0.0_17 does not exist",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDnsSecurity(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetDnsSecurity(ctx, "exmple.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateDnsSecurity(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"infoMessage": "DNSSEC enabled successfully. It may take up to 25 hours to reflect."
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := HostingApi{}.UpdateDnsSecurity(ctx, "example.com", map[string]interface{}{"status": "ENABLED"})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.InfoMessage, "DNSSEC enabled successfully. It may take up to 25 hours to reflect.")
}

func TestHostingUpdateDnsSecurityServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateDnsSecurity(ctx, "example.com", map[string]interface{}{"status": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateDnsSecurity(ctx, "example.com", map[string]interface{}{"status": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateDnsSecurity(ctx, "example.com", map[string]interface{}{"status": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateDnsSecurity(ctx, "example.com", map[string]interface{}{"status": "ENABLED"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingListResourceRecordSets(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"infoMessage": "Sorry, but your domain is not yet listed on our nameservers. Please create a DNS record first so we can add your domain.",
			"resourceRecordSets": [
				{
					"name": "example.com.",
					"type": "A",
					"content": [
						"85.17.150.50",
						"85.17.150.50",
						"85.17.150.53"
					],
					"ttl": 300,
					"editable": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/resourceRecordSets/example.com./A"
						},
						"collection": {
							"href": "/domains/example.com/resourceRecordSets"
						}
					}
				},
				{
					"name": "subdomain.example.com.",
					"type": "A",
					"content": [
						"85.17.150.54"
					],
					"ttl": 86400,
					"editable": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/resourceRecordSets/subdomain.example.com./A"
						},
						"collection": {
							"href": "/domains/example.com/resourceRecordSets"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/resourceRecordSets"
				},
				"parent": {
					"href": "/domains/example.com"
				},
				"validateSet": {
					"href": "/domains/example.com/resourceRecordSets/validateSet"
				}
			},
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	response, err := hostingApi.ListResourceRecordSets(ctx, "example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.ResourceRecordSets), 2)

	assert.Equal(response.InfoMessage, "Sorry, but your domain is not yet listed on our nameservers. Please create a DNS record first so we can add your domain.")
	assert.Equal(response.Link.Self.Href, "/domains/example.com/resourceRecordSets")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")
	assert.Equal(response.Link.ValidateSet.Href, "/domains/example.com/resourceRecordSets/validateSet")

	recordSet1 := response.ResourceRecordSets[0]
	assert.Equal(recordSet1.Name, "example.com.")
	assert.Equal(recordSet1.Type, "A")
	assert.Equal(recordSet1.Content[0], "85.17.150.50")
	assert.Equal(recordSet1.Content[1], "85.17.150.50")
	assert.Equal(recordSet1.Content[2], "85.17.150.53")
	assert.Equal(recordSet1.Ttl.String(), "300")
	assert.Equal(recordSet1.Editable, true)
	assert.Equal(recordSet1.Link.Self.Href, "/domains/example.com/resourceRecordSets/example.com./A")
	assert.Equal(recordSet1.Link.Collection.Href, "/domains/example.com/resourceRecordSets")

	recordSet2 := response.ResourceRecordSets[1]
	assert.Equal(recordSet2.Name, "subdomain.example.com.")
	assert.Equal(recordSet2.Type, "A")
	assert.Equal(recordSet2.Content[0], "85.17.150.54")
	assert.Equal(recordSet2.Ttl.String(), "86400")
	assert.Equal(recordSet2.Editable, true)
	assert.Equal(recordSet2.Link.Self.Href, "/domains/example.com/resourceRecordSets/subdomain.example.com./A")
	assert.Equal(recordSet2.Link.Collection.Href, "/domains/example.com/resourceRecordSets")
}

func TestHostingListResourceRecordSetsPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"infoMessage": "Sorry, but your domain is not yet listed on our nameservers. Please create a DNS record first so we can add your domain.",
			"resourceRecordSets": [
				{
					"name": "example.com.",
					"type": "A",
					"content": [
						"85.17.150.50",
						"85.17.150.50",
						"85.17.150.53"
					],
					"ttl": 300,
					"editable": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/resourceRecordSets/example.com./A"
						},
						"collection": {
							"href": "/domains/example.com/resourceRecordSets"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/resourceRecordSets"
				},
				"parent": {
					"href": "/domains/example.com"
				},
				"validateSet": {
					"href": "/domains/example.com/resourceRecordSets/validateSet"
				}
			},
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			}
		}`)
	})
	defer teardown()

	hostingApi := HostingApi{}
	ctx := context.Background()
	response, err := hostingApi.ListResourceRecordSets(ctx, "example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.ResourceRecordSets), 1)

	assert.Equal(response.InfoMessage, "Sorry, but your domain is not yet listed on our nameservers. Please create a DNS record first so we can add your domain.")
	assert.Equal(response.Link.Self.Href, "/domains/example.com/resourceRecordSets")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")
	assert.Equal(response.Link.ValidateSet.Href, "/domains/example.com/resourceRecordSets/validateSet")

	recordSet1 := response.ResourceRecordSets[0]
	assert.Equal(recordSet1.Name, "example.com.")
	assert.Equal(recordSet1.Type, "A")
	assert.Equal(recordSet1.Content[0], "85.17.150.50")
	assert.Equal(recordSet1.Content[1], "85.17.150.50")
	assert.Equal(recordSet1.Content[2], "85.17.150.53")
	assert.Equal(recordSet1.Ttl.String(), "300")
	assert.Equal(recordSet1.Editable, true)
	assert.Equal(recordSet1.Link.Self.Href, "/domains/example.com/resourceRecordSets/example.com./A")
	assert.Equal(recordSet1.Link.Collection.Href, "/domains/example.com/resourceRecordSets")
}

func TestHostingListResourceRecordSetsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListResourceRecordSets(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return HostingApi{}.ListResourceRecordSets(ctx, "example.com", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListResourceRecordSets(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingCreateResourceRecordSet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"name": "example.com.",
			"type": "A",
			"content": [
				"85.17.150.51",
				"85.17.150.52",
				"85.17.150.53"
			],
			"ttl": 3600,
			"editable": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/resourceRecordSets/example.com./A"
				},
				"collection": {
					"href": "/domains/example.com/resourceRecordSets"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.CreateResourceRecordSet(ctx, "example.com", map[string]interface{}{
		"name":    "example.com.",
		"type":    "A",
		"ttl":     3600,
		"content": []string{"85.17.150.51", "85.17.150.52", "85.17.150.53"},
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Name, "example.com.")
	assert.Equal(resp.Type, "A")
	assert.Equal(resp.Content[0], "85.17.150.51")
	assert.Equal(resp.Content[1], "85.17.150.52")
	assert.Equal(resp.Content[2], "85.17.150.53")
	assert.Equal(resp.Ttl.String(), "3600")
	assert.Equal(resp.Editable, true)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/resourceRecordSets/example.com./A")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/resourceRecordSets")
}

func TestHostingCreateResourceRecordSetServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateResourceRecordSet(ctx, "example.com", map[string]interface{}{
					"name":    "example.com.",
					"type":    "A",
					"ttl":     3600,
					"content": []string{"85.17.150.51", "85.17.150.52", "85.17.150.53"},
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateResourceRecordSet(ctx, "example.com", map[string]interface{}{
					"name":    "example.com.",
					"type":    "A",
					"ttl":     3600,
					"content": []string{"85.17.150.51", "85.17.150.52", "85.17.150.53"},
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateResourceRecordSet(ctx, "example.com", map[string]interface{}{
					"name":    "example.com.",
					"type":    "A",
					"ttl":     3600,
					"content": []string{"85.17.150.51", "85.17.150.52", "85.17.150.53"},
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateResourceRecordSet(ctx, "example.com", map[string]interface{}{
					"name":    "example.com.",
					"type":    "A",
					"ttl":     3600,
					"content": []string{"85.17.150.51", "85.17.150.52", "85.17.150.53"},
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingDeleteResourceRecordSets(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := HostingApi{}.DeleteResourceRecordSets(ctx, "example.com")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestHostingDeleteResourceRecordSetsServerError(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSet(ctx, "example.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSet(ctx, "example.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSet(ctx, "example.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSet(ctx, "example.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetResourceRecordSet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"name": "example.com.",
			"type": "A",
			"content": [
				"85.17.150.51",
				"85.17.150.52",
				"85.17.150.53"
			],
			"ttl": 3600,
			"editable": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/resourceRecordSets/example.com./A"
				},
				"collection": {
					"href": "/domains/example.com/resourceRecordSets"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.GetResourceRecordSet(ctx, "exmple.com", "example.com.", "A")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Name, "example.com.")
	assert.Equal(resp.Type, "A")
	assert.Equal(resp.Ttl.String(), "3600")
	assert.Equal(resp.Editable, true)
	assert.Equal(resp.Content[0], "85.17.150.51")
	assert.Equal(resp.Content[1], "85.17.150.52")
	assert.Equal(resp.Content[2], "85.17.150.53")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/resourceRecordSets/example.com./A")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/resourceRecordSets")
}

func TestHostingGetResourceRecordSetServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetResourceRecordSet(ctx, "exmple.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Range with id 88.17.0.0_17 does not exist"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetResourceRecordSet(ctx, "exmple.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				CorrelationId: "39e010ed-0e93-42c3-c28f-3ffc373553d5",
				Code:          "404",
				Message:       "Range with id 88.17.0.0_17 does not exist",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetResourceRecordSet(ctx, "exmple.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetResourceRecordSet(ctx, "exmple.com", "example.com.", "A")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateResourceRecordSet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"name": "example.com.",
			"type": "A",
			"content": [
				"85.17.150.54"
			],
			"ttl": 4200,
			"editable": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/resourceRecordSets/example.com./A"
				},
				"collection": {
					"href": "/domains/example.com/resourceRecordSets"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.UpdateResourceRecordSet(ctx, "exmple.com", "example.com.", "A", []string{"85.17.150.54"}, 4200)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Name, "example.com.")
	assert.Equal(resp.Type, "A")
	assert.Equal(resp.Ttl.String(), "4200")
	assert.Equal(resp.Editable, true)
	assert.Equal(resp.Content[0], "85.17.150.54")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/resourceRecordSets/example.com./A")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/resourceRecordSets")
}

func TestHostingUpdateResourceRecordSetServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateResourceRecordSet(ctx, "exmple.com", "example.com.", "A", []string{"85.17.150.54"}, 4200)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateResourceRecordSet(ctx, "exmple.com", "example.com.", "A", []string{"85.17.150.54"}, 4200)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateResourceRecordSet(ctx, "exmple.com", "example.com.", "A", []string{"85.17.150.54"}, 4200)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateResourceRecordSet(ctx, "exmple.com", "example.com.", "A", []string{"85.17.150.54"}, 4200)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingDeleteResourceRecordSet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := HostingApi{}.DeleteResourceRecordSet(ctx, "example.com", "example.com.", "A")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestHostingDeleteResourceRecordSetsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSets(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSets(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSets(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteResourceRecordSets(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingValidateResourceRecordSet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"infoMessage": "Resource record set is valid."
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.ValidateResourceRecordSet(ctx, "example.com", "example.com.", "A", []string{"127.0.0.1"}, 3600)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.InfoMessage, "Resource record set is valid.")
}

func TestHostingValidateResourceRecordSetServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ValidateResourceRecordSet(ctx, "example.com", "example.com.", "A", []string{"127.0.0.1"}, 3600)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ValidateResourceRecordSet(ctx, "example.com", "example.com.", "A", []string{"127.0.0.1"}, 3600)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ValidateResourceRecordSet(ctx, "example.com", "example.com.", "A", []string{"127.0.0.1"}, 3600)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ValidateResourceRecordSet(ctx, "example.com", "example.com.", "A", []string{"127.0.0.1"}, 3600)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingListCatchAll(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"destination": "destination@example.org",
			"spamChecksEnabled": true,
			"virusChecksEnabled": false,
			"_links": {
				"self": {
					"href": "/domains/example.com/catchAll"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.ListCatchAll(ctx, "example.com")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Destination, "destination@example.org")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, false)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/catchAll")
	assert.Equal(resp.Link.Parent.Href, "/domains/example.com")
}

func TestHostingListCatchAllServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListCatchAll(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListCatchAll(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListCatchAll(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateOrCreateCatchAll(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"destination": "destination@example.org",
			"spamChecksEnabled": true,
			"virusChecksEnabled": false,
			"_links": {
				"self": {
					"href": "/domains/example.com/catchAll"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.UpdateCatchAll(ctx, "exmple.com", map[string]interface{}{
		"destination":        "destination@example.org",
		"spamChecksEnabled":  true,
		"virusChecksEnabled": false,
	})

	assert := assert.New(t)

	assert.Nil(err)
	assert.Equal(resp.Destination, "destination@example.org")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, false)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/catchAll")
	assert.Equal(resp.Link.Parent.Href, "/domains/example.com")
}

func TestHostingUpdateOrCreateCatchAllServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateCatchAll(ctx, "exmple.com", map[string]interface{}{
					"destination":        "destination@example.org",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateCatchAll(ctx, "exmple.com", map[string]interface{}{
					"destination":        "destination@example.org",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateCatchAll(ctx, "exmple.com", map[string]interface{}{
					"destination":        "destination@example.org",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateCatchAll(ctx, "exmple.com", map[string]interface{}{
					"destination":        "destination@example.org",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingDeleteCatchAll(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := HostingApi{}.DeleteCatchAll(ctx, "example.com")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestHostingDeleteCatchAllsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteCatchAll(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteCatchAll(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteCatchAll(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteCatchAll(ctx, "example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingListEmailAliases(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"emailAliases": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				},
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": false,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.org"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := HostingApi{}.ListEmailAliases(ctx, "example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.EmailAliases), 2)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/emailAliases")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	email1 := response.EmailAliases[0]
	assert.Equal(email1.Active, true)
	assert.Equal(email1.Destination, "destination@example.com")
	assert.Equal(email1.Source, "source@example.com")
	assert.Equal(email1.SpamChecksEnabled, true)
	assert.Equal(email1.VirusChecksEnabled, true)
	assert.Equal(email1.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(email1.Link.Collection.Href, "/domains/example.com/emailAliases")

	email2 := response.EmailAliases[1]
	assert.Equal(email2.Active, true)
	assert.Equal(email2.Destination, "destination@example.com")
	assert.Equal(email2.Source, "source@example.com")
	assert.Equal(email2.SpamChecksEnabled, false)
	assert.Equal(email2.VirusChecksEnabled, true)
	assert.Equal(email2.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.org")
	assert.Equal(email2.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingListEmailAliasesPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"emailAliases": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := HostingApi{}.ListEmailAliases(ctx, "example.com", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.EmailAliases), 1)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/emailAliases")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	email1 := response.EmailAliases[0]
	assert.Equal(email1.Active, true)
	assert.Equal(email1.Destination, "destination@example.com")
	assert.Equal(email1.Source, "source@example.com")
	assert.Equal(email1.SpamChecksEnabled, true)
	assert.Equal(email1.VirusChecksEnabled, true)
	assert.Equal(email1.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(email1.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingListEmailAliasesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				opts := PaginationOptions{
					Limit: Int(1),
				}
				return HostingApi{}.ListEmailAliases(ctx, "example.com", opts)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListEmailAliases(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListEmailAliases(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingCreateEmailAlias(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"destination": "destination@example.com",
			"source": "source@example.com",
			"spamChecksEnabled": true,
			"virusChecksEnabled": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
				},
				"collection": {
					"href": "/domains/example.com/emailAliases"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.CreateEmailAlias(ctx, "example.com", map[string]interface{}{
		"active":             true,
		"destination":        "destination@example.com",
		"source":             "source@example.com",
		"spamChecksEnabled":  true,
		"virusChecksEnabled": true,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Destination, "destination@example.com")
	assert.Equal(resp.Source, "source@example.com")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, true)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingCreateEmailAliasServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateEmailAlias(ctx, "example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateEmailAlias(ctx, "example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateEmailAlias(ctx, "example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateEmailAlias(ctx, "example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetEmailAlias(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"destination": "destination@example.com",
			"source": "source@example.com",
			"spamChecksEnabled": true,
			"virusChecksEnabled": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
				},
				"collection": {
					"href": "/domains/example.com/emailAliases"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.GetEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Destination, "destination@example.com")
	assert.Equal(resp.Source, "source@example.com")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, true)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingGetEmailAliasServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateEmailAlias(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"destination": "destination@example.com",
			"source": "source@example.com",
			"spamChecksEnabled": true,
			"virusChecksEnabled": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
				},
				"collection": {
					"href": "/domains/example.com/emailAliases"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.UpdateEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com", map[string]bool{
		"active":             true,
		"spamChecksEnabled":  true,
		"virusChecksEnabled": true,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Destination, "destination@example.com")
	assert.Equal(resp.Source, "source@example.com")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, true)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingUpdateEmailAliasServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateEmailAlias(ctx, "example.com", "source@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingDeleteEmailAlias(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := HostingApi{}.DeleteEmailAlias(ctx, "example.com", "source", "dest")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestHostingDeleteEmailAliasServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteEmailAlias(ctx, "example.com", "source", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteEmailAlias(ctx, "example.com", "source", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteEmailAlias(ctx, "example.com", "source", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteEmailAlias(ctx, "example.com", "source", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingListDomainForwards(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"forwards": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/forwards/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/forwards"
						}
					}
				},
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": false,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/forwards/source.example.com/destination.example.org"
						},
						"collection": {
							"href": "/domains/example.com/forwards"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/forwards"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := HostingApi{}.ListDomainForwards(ctx, "example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Forwards), 2)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/forwards")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	forward1 := response.Forwards[0]
	assert.Equal(forward1.Active, true)
	assert.Equal(forward1.Destination, "destination@example.com")
	assert.Equal(forward1.Source, "source@example.com")
	assert.Equal(forward1.SpamChecksEnabled, true)
	assert.Equal(forward1.VirusChecksEnabled, true)
	assert.Equal(forward1.Link.Self.Href, "/domains/example.com/forwards/source.example.com/destination.example.com")
	assert.Equal(forward1.Link.Collection.Href, "/domains/example.com/forwards")

	forward2 := response.Forwards[1]
	assert.Equal(forward2.Active, true)
	assert.Equal(forward2.Destination, "destination@example.com")
	assert.Equal(forward2.Source, "source@example.com")
	assert.Equal(forward2.SpamChecksEnabled, false)
	assert.Equal(forward2.VirusChecksEnabled, true)
	assert.Equal(forward2.Link.Self.Href, "/domains/example.com/forwards/source.example.com/destination.example.org")
	assert.Equal(forward2.Link.Collection.Href, "/domains/example.com/forwards")
}

func TestHostingListDomainForwardsPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"forwards": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/forwards/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/forwards"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/forwards"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := HostingApi{}.ListDomainForwards(ctx, "example.com", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Forwards), 1)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/forwards")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	forward1 := response.Forwards[0]
	assert.Equal(forward1.Active, true)
	assert.Equal(forward1.Destination, "destination@example.com")
	assert.Equal(forward1.Source, "source@example.com")
	assert.Equal(forward1.SpamChecksEnabled, true)
	assert.Equal(forward1.VirusChecksEnabled, true)
	assert.Equal(forward1.Link.Self.Href, "/domains/example.com/forwards/source.example.com/destination.example.com")
	assert.Equal(forward1.Link.Collection.Href, "/domains/example.com/forwards")
}

func TestHostingListDomainForwardsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListDomainForwards(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListDomainForwards(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListDomainForwards(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingListMailBoxes(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"mailboxes": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"currentSize": 1234,
					"maximumSize": 4321,
					"suspended": true,
					"emailAddress": "mailbox@example.com",
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				},
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": false,
					"virusChecksEnabled": true,
					"currentSize": 6789,
					"maximumSize": 9876,
					"suspended": false,
					"emailAddress": "info@example.com",
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.org"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := HostingApi{}.ListMailBoxes(ctx, "example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Mailboxes), 2)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/emailAliases")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	mailBox1 := response.Mailboxes[0]
	assert.Equal(mailBox1.Active, true)
	assert.Equal(mailBox1.Destination, "destination@example.com")
	assert.Equal(mailBox1.Source, "source@example.com")
	assert.Equal(mailBox1.SpamChecksEnabled, true)
	assert.Equal(mailBox1.VirusChecksEnabled, true)
	assert.Equal(mailBox1.CurrentSize.String(), "1234")
	assert.Equal(mailBox1.EmailAddress, "mailbox@example.com")
	assert.Equal(mailBox1.MaximumSize.String(), "4321")
	assert.Equal(mailBox1.Suspended, true)
	assert.Equal(mailBox1.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(mailBox1.Link.Collection.Href, "/domains/example.com/emailAliases")

	mailBox2 := response.Mailboxes[1]
	assert.Equal(mailBox2.Active, true)
	assert.Equal(mailBox2.Destination, "destination@example.com")
	assert.Equal(mailBox2.Source, "source@example.com")
	assert.Equal(mailBox2.SpamChecksEnabled, false)
	assert.Equal(mailBox2.VirusChecksEnabled, true)
	assert.Equal(mailBox2.CurrentSize.String(), "6789")
	assert.Equal(mailBox2.EmailAddress, "info@example.com")
	assert.Equal(mailBox2.MaximumSize.String(), "9876")
	assert.Equal(mailBox2.Suspended, false)
	assert.Equal(mailBox2.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.org")
	assert.Equal(mailBox2.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingListMailBoxesPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"mailboxes": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"currentSize": 1234,
					"maximumSize": 4321,
					"suspended": true,
					"emailAddress": "mailbox@example.com",
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := HostingApi{}.ListMailBoxes(ctx, "example.com", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Mailboxes), 1)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/emailAliases")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	mailBox1 := response.Mailboxes[0]
	assert.Equal(mailBox1.Active, true)
	assert.Equal(mailBox1.Destination, "destination@example.com")
	assert.Equal(mailBox1.Source, "source@example.com")
	assert.Equal(mailBox1.SpamChecksEnabled, true)
	assert.Equal(mailBox1.VirusChecksEnabled, true)
	assert.Equal(mailBox1.CurrentSize.String(), "1234")
	assert.Equal(mailBox1.EmailAddress, "mailbox@example.com")
	assert.Equal(mailBox1.MaximumSize.String(), "4321")
	assert.Equal(mailBox1.Suspended, true)
	assert.Equal(mailBox1.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(mailBox1.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingListMailBoxesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListMailBoxes(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListMailBoxes(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListMailBoxes(ctx, "example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingCreateMailBox(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"currentSize": 123456789,
			"emailAddress": "mailbox@example.com",
			"maximumSize": 5368709120,
			"spamChecksEnabled": true,
			"suspended": false,
			"virusChecksEnabled": false,
			"localDelivery": true,
			"_embedded": {
				"autoResponder": {
					"_links": {
						"self": {
							"href": "/domains/example.com/mailboxes/mailbox.example.com/autoResponder"
						}
					}
				}
			},
			"_links": {
				"self": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com"
				},
				"collection": {
					"href": "/domains/example.com/mailboxes"
				},
				"autoResponder": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com/autoResponder"
				},
				"forwards": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com/forwards"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.CreateMailBox(ctx, "example.com", map[string]interface{}{
		"emailAddress":       "mailbox@example.com",
		"active":             true,
		"password":           "CHANGETHIS",
		"spamChecksEnabled":  true,
		"virusChecksEnabled": false,
		"localDelivery":      true,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.CurrentSize.String(), "123456789")
	assert.Equal(resp.EmailAddress, "mailbox@example.com")
	assert.Equal(resp.MaximumSize.String(), "5368709120")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.Suspended, false)
	assert.Equal(resp.VirusChecksEnabled, false)
	assert.Equal(resp.LocalDelivery, true)
	assert.Equal(resp.Embedded.AutoResponder.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com/autoResponder")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/mailboxes")
	assert.Equal(resp.Link.AutoResponder.Href, "/domains/example.com/mailboxes/mailbox.example.com/autoResponder")
	assert.Equal(resp.Link.Forwards.Href, "/domains/example.com/mailboxes/mailbox.example.com/forwards")
}

func TestHostingCreateMailBoxServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateMailBox(ctx, "example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateMailBox(ctx, "example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateMailBox(ctx, "example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateMailBox(ctx, "example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetMailBox(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"currentSize": 123456789,
			"emailAddress": "mailbox@example.com",
			"maximumSize": 5368709120,
			"spamChecksEnabled": true,
			"suspended": false,
			"virusChecksEnabled": false,
			"localDelivery": true,
			"_embedded": {
				"autoResponder": {
					"_links": {
						"self": {
							"href": "/domains/example.com/mailboxes/mailbox.example.com/autoResponder"
						}
					}
				}
			},
			"_links": {
				"self": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com"
				},
				"collection": {
					"href": "/domains/example.com/mailboxes"
				},
				"autoResponder": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com/autoResponder"
				},
				"forwards": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com/forwards"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.GetMailBox(ctx, "example.com", "mailbox@example.com")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.CurrentSize.String(), "123456789")
	assert.Equal(resp.EmailAddress, "mailbox@example.com")
	assert.Equal(resp.MaximumSize.String(), "5368709120")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.Suspended, false)
	assert.Equal(resp.VirusChecksEnabled, false)
	assert.Equal(resp.LocalDelivery, true)
	assert.Equal(resp.Embedded.AutoResponder.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com/autoResponder")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/mailboxes")
	assert.Equal(resp.Link.AutoResponder.Href, "/domains/example.com/mailboxes/mailbox.example.com/autoResponder")
	assert.Equal(resp.Link.Forwards.Href, "/domains/example.com/mailboxes/mailbox.example.com/forwards")
}

func TestHostingGetMailBoxServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetMailBox(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetMailBox(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetMailBox(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetMailBox(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateMailBox(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"currentSize": 123456789,
			"emailAddress": "mailbox@example.com",
			"maximumSize": 5368709120,
			"spamChecksEnabled": true,
			"suspended": false,
			"virusChecksEnabled": false,
			"localDelivery": true,
			"_embedded": {
				"autoResponder": {
					"_links": {
						"self": {
							"href": "/domains/example.com/mailboxes/mailbox.example.com/autoResponder"
						}
					}
				}
			},
			"_links": {
				"self": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com"
				},
				"collection": {
					"href": "/domains/example.com/mailboxes"
				},
				"autoResponder": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com/autoResponder"
				},
				"forwards": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com/forwards"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.UpdateMailBox(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
		"emailAddress":       "mailbox@example.com",
		"active":             true,
		"password":           "CHANGETHIS",
		"spamChecksEnabled":  true,
		"virusChecksEnabled": false,
		"localDelivery":      true,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.CurrentSize.String(), "123456789")
	assert.Equal(resp.EmailAddress, "mailbox@example.com")
	assert.Equal(resp.MaximumSize.String(), "5368709120")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.Suspended, false)
	assert.Equal(resp.VirusChecksEnabled, false)
	assert.Equal(resp.LocalDelivery, true)
	assert.Equal(resp.Embedded.AutoResponder.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com/autoResponder")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/mailboxes")
	assert.Equal(resp.Link.AutoResponder.Href, "/domains/example.com/mailboxes/mailbox.example.com/autoResponder")
	assert.Equal(resp.Link.Forwards.Href, "/domains/example.com/mailboxes/mailbox.example.com/forwards")
}

func TestHostingUpdateMailBoxServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateMailBox(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateMailBox(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateMailBox(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateMailBox(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"emailAddress":       "mailbox@example.com",
					"active":             true,
					"password":           "CHANGETHIS",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": false,
					"localDelivery":      true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingDeleteMailBox(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := HostingApi{}.DeleteMailBox(ctx, "example.com", "info@example.com")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestHostingDeleteMailBoxServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteMailBox(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteMailBox(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteMailBox(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteMailBox(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetAutoResponder(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"body": "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.",
			"subject": "Out of office",
			"_links": {
				"self": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com"
				},
				"parent": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.GetAutoResponder(ctx, "example.com", "mailbox@example.com")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Body, "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.")
	assert.Equal(resp.Subject, "Out of office")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com")
	assert.Equal(resp.Link.Parent.Href, "/domains/example.com/mailboxes/mailbox.example.com")
}

func TestHostingGetAutoResponderServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAutoResponder(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAutoResponder(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAutoResponder(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetAutoResponder(ctx, "example.com", "mailbox@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateOrCreateAutoResponder(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"body": "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.",
			"subject": "Out of office",
			"_links": {
				"self": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com"
				},
				"parent": {
					"href": "/domains/example.com/mailboxes/mailbox.example.com"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.UpdateAutoResponder(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
		"active":  true,
		"body":    "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.",
		"subject": "Out of office",
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Body, "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.")
	assert.Equal(resp.Subject, "Out of office")
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/mailboxes/mailbox.example.com")
	assert.Equal(resp.Link.Parent.Href, "/domains/example.com/mailboxes/mailbox.example.com")
}

func TestHostingUpdateOrCreateAutoResponderServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateAutoResponder(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"active":  true,
					"body":    "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.",
					"subject": "Out of office",
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateAutoResponder(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"active":  true,
					"body":    "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.",
					"subject": "Out of office",
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateAutoResponder(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"active":  true,
					"body":    "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.",
					"subject": "Out of office",
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateAutoResponder(ctx, "example.com", "mailbox@example.com", map[string]interface{}{
					"active":  true,
					"body":    "I will be out of office until 14-12-2020. Please contact the reception desk for urgent matters.",
					"subject": "Out of office",
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingDeleteAutoResponder(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := HostingApi{}.DeleteAutoResponder(ctx, "example.com", "info@example.com")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestHostingDeleteAutoResponderServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteAutoResponder(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteAutoResponder(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteAutoResponder(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteAutoResponder(ctx, "example.com", "info@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingListForwards(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"forwards": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				},
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": false,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.org"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 2,
				"limit": 10,
				"offset": 0
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	response, err := HostingApi{}.ListForwards(ctx, "example.com", "random@example.com", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Forwards), 2)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/emailAliases")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	forward1 := response.Forwards[0]
	assert.Equal(forward1.Active, true)
	assert.Equal(forward1.Destination, "destination@example.com")
	assert.Equal(forward1.Source, "source@example.com")
	assert.Equal(forward1.SpamChecksEnabled, true)
	assert.Equal(forward1.VirusChecksEnabled, true)
	assert.Equal(forward1.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(forward1.Link.Collection.Href, "/domains/example.com/emailAliases")

	forward2 := response.Forwards[1]
	assert.Equal(forward2.Active, true)
	assert.Equal(forward2.Destination, "destination@example.com")
	assert.Equal(forward2.Source, "source@example.com")
	assert.Equal(forward2.SpamChecksEnabled, false)
	assert.Equal(forward2.VirusChecksEnabled, true)
	assert.Equal(forward2.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.org")
	assert.Equal(forward2.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingListForwardsPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"forwards": [
				{
					"active": true,
					"destination": "destination@example.com",
					"source": "source@example.com",
					"spamChecksEnabled": true,
					"virusChecksEnabled": true,
					"_links": {
						"self": {
							"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
						},
						"collection": {
							"href": "/domains/example.com/emailAliases"
						}
					}
				}
			],
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases"
				},
				"parent": {
					"href": "/domains/example.com"
				}
			},
			"_metadata": {
				"totalCount": 11,
				"limit": 10,
				"offset": 1
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := HostingApi{}.ListForwards(ctx, "example.com", "random@example.com", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Forwards), 1)
	assert.Equal(response.Link.Self.Href, "/domains/example.com/emailAliases")
	assert.Equal(response.Link.Parent.Href, "/domains/example.com")

	forward1 := response.Forwards[0]
	assert.Equal(forward1.Active, true)
	assert.Equal(forward1.Destination, "destination@example.com")
	assert.Equal(forward1.Source, "source@example.com")
	assert.Equal(forward1.SpamChecksEnabled, true)
	assert.Equal(forward1.VirusChecksEnabled, true)
	assert.Equal(forward1.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(forward1.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingListForwardsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListForwards(ctx, "example.com", "random@example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListForwards(ctx, "example.com", "random@example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.ListForwards(ctx, "example.com", "random@example.com", PaginationOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingCreateForward(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"destination": "destination@example.com",
			"source": "source@example.com",
			"spamChecksEnabled": true,
			"virusChecksEnabled": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
				},
				"collection": {
					"href": "/domains/example.com/emailAliases"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.CreateForward(ctx, "example.com", "random@example.com", map[string]interface{}{
		"active":             true,
		"destination":        "destination@example.com",
		"source":             "source@example.com",
		"spamChecksEnabled":  true,
		"virusChecksEnabled": true,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Destination, "destination@example.com")
	assert.Equal(resp.Source, "source@example.com")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, true)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingCreateForwardServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateForward(ctx, "example.com", "random@example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateForward(ctx, "example.com", "random@example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateForward(ctx, "example.com", "random@example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.CreateForward(ctx, "example.com", "random@example.com", map[string]interface{}{
					"active":             true,
					"destination":        "destination@example.com",
					"source":             "source@example.com",
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingGetForward(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"destination": "destination@example.com",
			"source": "source@example.com",
			"spamChecksEnabled": true,
			"virusChecksEnabled": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
				},
				"collection": {
					"href": "/domains/example.com/emailAliases"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.GetForward(ctx, "example.com", "random@example.com", "destination@example.com")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Destination, "destination@example.com")
	assert.Equal(resp.Source, "source@example.com")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, true)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingGetForwardServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetForward(ctx, "example.com", "random@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetForward(ctx, "example.com", "random@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetForward(ctx, "example.com", "random@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.GetForward(ctx, "example.com", "random@example.com", "destination@example.com")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingUpdateForward(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"active": true,
			"destination": "destination@example.com",
			"source": "source@example.com",
			"spamChecksEnabled": true,
			"virusChecksEnabled": true,
			"_links": {
				"self": {
					"href": "/domains/example.com/emailAliases/source.example.com/destination.example.com"
				},
				"collection": {
					"href": "/domains/example.com/emailAliases"
				}
			}
		}`)
	})
	defer teardown()

	ctx := context.Background()
	resp, err := HostingApi{}.UpdateForward(ctx, "example.com", "random@example.com", "destination@example.com", map[string]bool{
		"active":             true,
		"spamChecksEnabled":  true,
		"virusChecksEnabled": true,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp.Active, true)
	assert.Equal(resp.Destination, "destination@example.com")
	assert.Equal(resp.Source, "source@example.com")
	assert.Equal(resp.SpamChecksEnabled, true)
	assert.Equal(resp.VirusChecksEnabled, true)
	assert.Equal(resp.Link.Self.Href, "/domains/example.com/emailAliases/source.example.com/destination.example.com")
	assert.Equal(resp.Link.Collection.Href, "/domains/example.com/emailAliases")
}

func TestHostingUpdateForwardServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateForward(ctx, "example.com", "random@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateForward(ctx, "example.com", "random@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateForward(ctx, "example.com", "random@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return HostingApi{}.UpdateForward(ctx, "example.com", "random@example.com", "destination@example.com", map[string]bool{
					"active":             true,
					"spamChecksEnabled":  true,
					"virusChecksEnabled": true,
				})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestHostingDeleteForward(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	ctx := context.Background()
	err := HostingApi{}.DeleteForward(ctx, "example.com", "email", "dest")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestHostingDeleteForwardServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 401",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteForward(ctx, "example.com", "email", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "401",
				Message:       "You are not authorized to view this resource.",
			},
		},
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteForward(ctx, "example.com", "email", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "ACCESS_DENIED",
				Message:       "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteForward(ctx, "example.com", "email", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "SERVER_ERROR",
				Message:       "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, HostingApi{}.DeleteForward(ctx, "example.com", "email", "dest")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "TEMPORARILY_UNAVAILABLE",
				Message:       "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
