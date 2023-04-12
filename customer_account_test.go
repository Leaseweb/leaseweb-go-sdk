package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerAccountGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"address": {
				"city": "Amsterdam",
				"country": "Netherlands",
				"houseNumber": "8",
				"postalCode": "1101 EC",
				"state": "foo state",
				"stateCode": "bar stateCode",
				"street": "luttenbergweg"
			},
			"name": "Leaseweb",
			"resellerTier": "resellerTier",
			"vatNumber": "vat1234134"
		}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	response, err := customerAccountApi.Get(ctx)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Name, "Leaseweb")
	assert.Equal(response.ResellerTier, "resellerTier")
	assert.Equal(response.VatNumber, "vat1234134")
	assert.Equal(response.Address.City, "Amsterdam")
	assert.Equal(response.Address.Country, "Netherlands")
	assert.Equal(response.Address.HouseNumber, "8")
	assert.Equal(response.Address.PostalCode, "1101 EC")
	assert.Equal(response.Address.State, "foo state")
	assert.Equal(response.Address.StateCode, "bar stateCode")
	assert.Equal(response.Address.Street, "luttenbergweg")
}

func TestCustomerAccountGetServerErrors(t *testing.T) {
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
				return CustomerAccountApi{}.Get(ctx)
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.Get(ctx)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.Get(ctx)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.Get(ctx)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestCustomerAccountUpdate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	err := customerAccountApi.Update(ctx, CustomerAccountAddress{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})

	assert := assert.New(t)
	assert.Nil(err)
}

func TestCustomerAccountUpdateServerErrors(t *testing.T) {
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
				return nil, CustomerAccountApi{}.Update(ctx, CustomerAccountAddress{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.Update(ctx, CustomerAccountAddress{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 405",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "405", "errorMessage": "AccountDetails modifications are not permitted for USA residents, please contact support for any modification request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.Update(ctx, CustomerAccountAddress{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "405",
				Message:       "AccountDetails modifications are not permitted for USA residents, please contact support for any modification request.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.Update(ctx, CustomerAccountAddress{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.Update(ctx, CustomerAccountAddress{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
func TestCustomerAccountListContacts(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 1}, "contacts": [
			{
				"description": "Mr.",
				"email": "john@doe.com",
				"firstName": "John",
				"id": "1",
				"lastName": "Doe",
				"mobile": {
					"countryCode": "NL",
					"dialCode": "+31",
					"number": "682212341"
				},
				"phone": {
					"countryCode": "NL",
					"dialCode": "+32",
					"number": "682212342"
				},
				"primaryRoles": [
					"GENERAL",
					"SECURITY",
					"TECHNICAL",
					"BILLING"
				],
				"roles": [
					"GENERAL",
					"TECHNICAL",
					"BILLING",
					"SECURITY"
				]
			}
		]}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	resp, err := customerAccountApi.ListContacts(ctx, CustomerAccountContactsListOptions{})

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Metadata.TotalCount, 1)
	assert.Equal(resp.Metadata.Offset, 0)
	assert.Equal(resp.Metadata.Limit, 10)
	assert.Equal(len(resp.Contacts), 1)

	contact := resp.Contacts[0]
	assert.Equal(contact.Description, "Mr.")
	assert.Equal(contact.Email, "john@doe.com")
	assert.Equal(contact.FirstName, "John")
	assert.Equal(contact.Id, "1")
	assert.Equal(contact.LastName, "Doe")
	assert.Subset(contact.Roles, []string{"GENERAL", "TECHNICAL", "BILLING", "SECURITY"})
	assert.Subset(contact.PrimaryRoles, []string{"GENERAL", "TECHNICAL", "BILLING", "SECURITY"})
	assert.Equal(contact.Mobile.CountryCode, "NL")
	assert.Equal(contact.Mobile.DialCode, "+31")
	assert.Equal(contact.Mobile.Number, "682212341")
	assert.Equal(contact.Phone.CountryCode, "NL")
	assert.Equal(contact.Phone.DialCode, "+32")
	assert.Equal(contact.Phone.Number, "682212342")
}

func TestCustomerAccountListContactsPaginateAndFilter(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 5, "offset": 1, "totalCount": 6}, "contacts": [
			{
				"description": "Mr.",
				"email": "john@doe.com",
				"firstName": "John",
				"id": "1",
				"lastName": "Doe",
				"mobile": {
					"countryCode": "NL",
					"dialCode": "+31",
					"number": "682212341"
				},
				"phone": {
					"countryCode": "NL",
					"dialCode": "+32",
					"number": "682212342"
				},
				"primaryRoles": [
					"GENERAL",
					"SECURITY",
					"TECHNICAL",
					"BILLING"
				],
				"roles": [
					"GENERAL",
					"TECHNICAL",
					"BILLING",
					"SECURITY"
				]
			}
		]}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	opts := CustomerAccountContactsListOptions{
		PrimaryRoles: []string{"GENERAL", "SECURITY", "TECHNICAL", "BILLING"},
	}
	resp, err := customerAccountApi.ListContacts(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Metadata.TotalCount, 6)
	assert.Equal(resp.Metadata.Offset, 1)
	assert.Equal(resp.Metadata.Limit, 5)
	assert.Equal(len(resp.Contacts), 1)

	contact := resp.Contacts[0]
	assert.Equal(contact.Description, "Mr.")
	assert.Equal(contact.Email, "john@doe.com")
	assert.Equal(contact.FirstName, "John")
	assert.Equal(contact.Id, "1")
	assert.Equal(contact.LastName, "Doe")
	assert.Subset(contact.Roles, []string{"GENERAL", "TECHNICAL", "BILLING", "SECURITY"})
	assert.Subset(contact.PrimaryRoles, []string{"GENERAL", "TECHNICAL", "BILLING", "SECURITY"})
	assert.Equal(contact.Mobile.CountryCode, "NL")
	assert.Equal(contact.Mobile.DialCode, "+31")
	assert.Equal(contact.Mobile.Number, "682212341")
	assert.Equal(contact.Phone.CountryCode, "NL")
	assert.Equal(contact.Phone.DialCode, "+32")
	assert.Equal(contact.Phone.Number, "682212342")
}

func TestCustomerAccountListContactsServerErrors(t *testing.T) {
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
				return CustomerAccountApi{}.ListContacts(ctx, CustomerAccountContactsListOptions{})
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.ListContacts(ctx, CustomerAccountContactsListOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.ListContacts(ctx, CustomerAccountContactsListOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.ListContacts(ctx, CustomerAccountContactsListOptions{})
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestCustomerAccountCreateContact(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"description": "Mr.",
			"email": "john@doe.com",
			"firstName": "John",
			"lastName": "Doe",
			"mobile": {
				"countryCode": "NL",
				"number": "682212341"
			},
			"phone": {
				"countryCode": "NL",
				"number": "682212342"
			},
			"roles": [
				"GENERAL",
				"TECHNICAL",
				"BILLING"
			]
		}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	newContact := CustomerAccountContact{
		Description: "Mr.",
		Email:       "john@doe.com",
		FirstName:   "John",
		LastName:    "Doe",
		Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
		Mobile:      CustomerAccountPhone{CountryCode: "NL", Number: "682212341"},
		Phone:       CustomerAccountPhone{CountryCode: "NL", Number: "682212342"},
	}
	ctx := context.Background()
	resp, err := customerAccountApi.CreateContact(ctx, newContact)

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(resp.Description, "Mr.")
	assert.Equal(resp.Email, "john@doe.com")
	assert.Equal(resp.FirstName, "John")
	assert.Equal(resp.LastName, "Doe")
	assert.Subset(resp.Roles, []string{"GENERAL", "TECHNICAL", "BILLING"})
	assert.Equal(resp.Mobile.CountryCode, "NL")
	assert.Equal(resp.Mobile.Number, "682212341")
	assert.Equal(resp.Phone.CountryCode, "NL")
	assert.Equal(resp.Phone.Number, "682212342")
}

func TestCustomerAccountCreateContactServerErrors(t *testing.T) {
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
				newContact := CustomerAccountContact{
					Description: "Mr.",
					Email:       "john@doe.com",
					FirstName:   "John",
					LastName:    "Doe",
					Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
					Mobile:      CustomerAccountPhone{CountryCode: "NL", Number: "682212341"},
					Phone:       CustomerAccountPhone{CountryCode: "NL", Number: "682212342"},
				}
				ctx := context.Background()
				return CustomerAccountApi{}.CreateContact(ctx, newContact)
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				newContact := CustomerAccountContact{
					Description: "Mr.",
					Email:       "john@doe.com",
					FirstName:   "John",
					LastName:    "Doe",
					Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
					Mobile:      CustomerAccountPhone{CountryCode: "NL", Number: "682212341"},
					Phone:       CustomerAccountPhone{CountryCode: "NL", Number: "682212342"},
				}
				ctx := context.Background()
				return CustomerAccountApi{}.CreateContact(ctx, newContact)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				newContact := CustomerAccountContact{
					Description: "Mr.",
					Email:       "john@doe.com",
					FirstName:   "John",
					LastName:    "Doe",
					Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
					Mobile:      CustomerAccountPhone{CountryCode: "NL", Number: "682212341"},
					Phone:       CustomerAccountPhone{CountryCode: "NL", Number: "682212342"},
				}
				ctx := context.Background()
				return CustomerAccountApi{}.CreateContact(ctx, newContact)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				newContact := CustomerAccountContact{
					Description: "Mr.",
					Email:       "john@doe.com",
					FirstName:   "John",
					LastName:    "Doe",
					Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
					Mobile:      CustomerAccountPhone{CountryCode: "NL", Number: "682212341"},
					Phone:       CustomerAccountPhone{CountryCode: "NL", Number: "682212342"},
				}
				ctx := context.Background()
				return CustomerAccountApi{}.CreateContact(ctx, newContact)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestCustomerAccountDeleteContact(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	err := customerAccountApi.DeleteContact(ctx, "contact-id")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestCustomerAccountDeleteContactServerErrors(t *testing.T) {
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
				return nil, CustomerAccountApi{}.DeleteContact(ctx, "contact-id")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.DeleteContact(ctx, "contact-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.DeleteContact(ctx, "contact-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodDelete, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.DeleteContact(ctx, "contact-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestCustomerAccountGetContact(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"description": "Mr.",
			"email": "john@doe.com",
			"firstName": "John",
			"id": "1",
			"lastName": "Doe",
			"mobile": {
				"countryCode": "NL",
				"dialCode": "+31",
				"number": "682212341"
			},
			"phone": {
				"countryCode": "NL",
				"dialCode": "+32",
				"number": "682212342"
			},
			"primaryRoles": [
				"GENERAL",
				"TECHNICAL",
				"BILLING",
				"SECURITY"
			],
			"roles": [
				"GENERAL",
				"TECHNICAL",
				"BILLING",
				"SECURITY"
			]
		}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	response, err := customerAccountApi.GetContact(ctx, "contact-id")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Description, "Mr.")
	assert.Equal(response.Email, "john@doe.com")
	assert.Equal(response.FirstName, "John")
	assert.Equal(response.Id, "1")
	assert.Equal(response.LastName, "Doe")
	assert.Subset(response.Roles, []string{"GENERAL", "TECHNICAL", "BILLING", "SECURITY"})
	assert.Subset(response.PrimaryRoles, []string{"GENERAL", "TECHNICAL", "BILLING", "SECURITY"})
	assert.Equal(response.Mobile.CountryCode, "NL")
	assert.Equal(response.Mobile.DialCode, "+31")
	assert.Equal(response.Mobile.Number, "682212341")
	assert.Equal(response.Phone.CountryCode, "NL")
	assert.Equal(response.Phone.DialCode, "+32")
	assert.Equal(response.Phone.Number, "682212342")
}

func TestCustomerAccountGetContactServerErrors(t *testing.T) {
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
				return CustomerAccountApi{}.GetContact(ctx, "contact-id")
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.GetContact(ctx, "contact-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.GetContact(ctx, "contact-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return CustomerAccountApi{}.GetContact(ctx, "contact-id")
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestCustomerAccountUpdateContact(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	err := customerAccountApi.UpdateContact(
		ctx,
		"contact-id",
		CustomerAccountPhone{CountryCode: "NL", Number: "653388213"},
		[]string{"GENERAL", "TECHNICAL"},
		CustomerAccountPhone{CountryCode: "NL", Number: "653388214"},
		"Description",
	)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestCustomerAccountUpdateContactServerErrors(t *testing.T) {
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
				return nil, CustomerAccountApi{}.UpdateContact(
					ctx,
					"contact-id",
					CustomerAccountPhone{CountryCode: "NL", Number: "653388213"},
					[]string{"GENERAL", "TECHNICAL"},
					CustomerAccountPhone{CountryCode: "NL", Number: "653388214"},
					"Description",
				)
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.UpdateContact(
					ctx,
					"contact-id",
					CustomerAccountPhone{CountryCode: "NL", Number: "653388213"},
					[]string{"GENERAL", "TECHNICAL"},
					CustomerAccountPhone{CountryCode: "NL", Number: "653388214"},
					"Description",
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.UpdateContact(
					ctx,
					"contact-id",
					CustomerAccountPhone{CountryCode: "NL", Number: "653388213"},
					[]string{"GENERAL", "TECHNICAL"},
					CustomerAccountPhone{CountryCode: "NL", Number: "653388214"},
					"Description",
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.UpdateContact(
					ctx,
					"contact-id",
					CustomerAccountPhone{CountryCode: "NL", Number: "653388213"},
					[]string{"GENERAL", "TECHNICAL"},
					CustomerAccountPhone{CountryCode: "NL", Number: "653388214"},
					"Description",
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.UpdateContact(
					ctx,
					"contact-id",
					CustomerAccountPhone{CountryCode: "NL", Number: "653388213"},
					[]string{"GENERAL", "TECHNICAL"},
					CustomerAccountPhone{CountryCode: "NL", Number: "653388214"},
					"Description",
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestCustomerAccountAssignPrimaryRolesToContact(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	ctx := context.Background()
	err := customerAccountApi.AssignPrimaryRolesToContact(
		ctx,
		"contact-id",
		[]string{"GENERAL", "TECHNICAL"},
	)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestCustomerAccountAssignPrimaryRolesToContactServerErrors(t *testing.T) {
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
				return nil, CustomerAccountApi{}.AssignPrimaryRolesToContact(
					ctx,
					"contact-id",
					[]string{"GENERAL", "TECHNICAL"},
				)
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
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.AssignPrimaryRolesToContact(
					ctx,
					"contact-id",
					[]string{"GENERAL", "TECHNICAL"},
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "403",
				Message:       "Access to the requested resource is forbidden.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.AssignPrimaryRolesToContact(
					ctx,
					"contact-id",
					[]string{"GENERAL", "TECHNICAL"},
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "404",
				Message:       "Resource not found",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.AssignPrimaryRolesToContact(
					ctx,
					"contact-id",
					[]string{"GENERAL", "TECHNICAL"},
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "500",
				Message:       "The API could not handle your request at this time.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, CustomerAccountApi{}.AssignPrimaryRolesToContact(
					ctx,
					"contact-id",
					[]string{"GENERAL", "TECHNICAL"},
				)
			},
			ExpectedError: ApiError{
				CorrelationId: "289346a1-3eaf-4da4-b707-62ef12eb08be",
				Code:          "503",
				Message:       "The API is not available at the moment.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
