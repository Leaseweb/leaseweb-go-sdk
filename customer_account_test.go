package leaseweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomerAccount(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
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
	response, err := customerAccountApi.GetCustomerAccount()

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

func TestGetCustomerAccountError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetCustomerAccount()

	assert := assert.New(t)

	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestGetCustomerAccountError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetCustomerAccount()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestGetCustomerAccountError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetCustomerAccount()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestGetCustomerAccountError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetCustomerAccount()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestUpdateCustomerAccount(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateCustomerAccount(Address{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})

	assert := assert.New(t)
	assert.Nil(err)
}

func TestUpdateCustomerAccountError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateCustomerAccount(Address{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})

	assert := assert.New(t)

	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestUpdateCustomerAccountError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateCustomerAccount(Address{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestUpdateCustomerAccountError405(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "405", "errorMessage": "AccountDetails modifications are not permitted for USA residents, please contact support for any modification request."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateCustomerAccount(Address{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "AccountDetails modifications are not permitted for USA residents, please contact support for any modification request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "405")
}

func TestUpdateCustomerAccountError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateCustomerAccount(Address{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestUpdateCustomerAccountError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateCustomerAccount(Address{City: "amsterdam", HouseNumber: "800", PostalCode: "1105 AB", Street: "Hessenbergweg"})

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestListContacts(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
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
	resp, err := customerAccountApi.ListContacts()

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

func TestListContactsPaginateAndFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
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
	resp, err := customerAccountApi.ListContacts(1, 5, []string{"GENERAL", "SECURITY", "TECHNICAL", "BILLING"})

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

func TestListContactsError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.ListContacts()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestListContactsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.ListContacts()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestListContactsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.ListContacts()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestListContactsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.ListContacts()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestCreateContact(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
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
	newContact := Contact{
		Description: "Mr.",
		Email:       "john@doe.com",
		FirstName:   "John",
		LastName:    "Doe",
		Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
		Mobile:      Phone{CountryCode: "NL", Number: "682212341"},
		Phone:       Phone{CountryCode: "NL", Number: "682212342"},
	}
	resp, err := customerAccountApi.CreateContact(newContact)

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

func TestCreateContactError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	newContact := Contact{
		Description: "Mr.",
		Email:       "john@doe.com",
		FirstName:   "John",
		LastName:    "Doe",
		Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
		Mobile:      Phone{CountryCode: "NL", Number: "682212341"},
		Phone:       Phone{CountryCode: "NL", Number: "682212342"},
	}
	resp, err := customerAccountApi.CreateContact(newContact)

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestCreateContactError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	newContact := Contact{
		Description: "Mr.",
		Email:       "john@doe.com",
		FirstName:   "John",
		LastName:    "Doe",
		Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
		Mobile:      Phone{CountryCode: "NL", Number: "682212341"},
		Phone:       Phone{CountryCode: "NL", Number: "682212342"},
	}
	resp, err := customerAccountApi.CreateContact(newContact)

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestCreateContactError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	newContact := Contact{
		Description: "Mr.",
		Email:       "john@doe.com",
		FirstName:   "John",
		LastName:    "Doe",
		Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
		Mobile:      Phone{CountryCode: "NL", Number: "682212341"},
		Phone:       Phone{CountryCode: "NL", Number: "682212342"},
	}
	resp, err := customerAccountApi.CreateContact(newContact)

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestCreateContactError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	newContact := Contact{
		Description: "Mr.",
		Email:       "john@doe.com",
		FirstName:   "John",
		LastName:    "Doe",
		Roles:       []string{"GENERAL", "TECHNICAL", "BILLING"},
		Mobile:      Phone{CountryCode: "NL", Number: "682212341"},
		Phone:       Phone{CountryCode: "NL", Number: "682212342"},
	}
	resp, err := customerAccountApi.CreateContact(newContact)

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestDeleteContact(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.DeleteContact("contact-id")

	assert := assert.New(t)
	assert.Nil(err)
}

func TestDeleteContactError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.DeleteContact("contact-id")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestDeleteContactError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.DeleteContact("contact-id")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestDeleteContactError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.DeleteContact("contact-id")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "404")
}

func TestDeleteContactError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.DeleteContact("contact-id")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestDeleteContactError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.DeleteContact("contact-id")

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestGetContact(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
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
	response, err := customerAccountApi.GetContact("contact-id")

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

func TestGetContactError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetContact("contact-id")

	assert := assert.New(t)

	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestGetContactError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetContact("contact-id")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestGetContactError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetContact("contact-id")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestGetContactError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	resp, err := customerAccountApi.GetContact("contact-id")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestUpdateContact(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateContact(
		"contact-id",
		Phone{CountryCode: "NL", Number: "653388213"},
		[]string{"GENERAL", "TECHNICAL"},
		Phone{CountryCode: "NL", Number: "653388214"},
		"Description",
	)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestUpdateContactError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateContact(
		"contact-id",
		Phone{CountryCode: "NL", Number: "653388213"},
		[]string{"GENERAL", "TECHNICAL"},
		Phone{CountryCode: "NL", Number: "653388214"},
		"Description",
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestUpdateContactError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateContact(
		"contact-id",
		Phone{CountryCode: "NL", Number: "653388213"},
		[]string{"GENERAL", "TECHNICAL"},
		Phone{CountryCode: "NL", Number: "653388214"},
		"Description",
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestUpdateContactError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateContact(
		"contact-id",
		Phone{CountryCode: "NL", Number: "653388213"},
		[]string{"GENERAL", "TECHNICAL"},
		Phone{CountryCode: "NL", Number: "653388214"},
		"Description",
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "404")
}

func TestUpdateContactError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateContact(
		"contact-id",
		Phone{CountryCode: "NL", Number: "653388213"},
		[]string{"GENERAL", "TECHNICAL"},
		Phone{CountryCode: "NL", Number: "653388214"},
		"Description",
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestUpdateContactError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.UpdateContact(
		"contact-id",
		Phone{CountryCode: "NL", Number: "653388213"},
		[]string{"GENERAL", "TECHNICAL"},
		Phone{CountryCode: "NL", Number: "653388214"},
		"Description",
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestAssignPrimaryRolesToContact(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.AssignPrimaryRolesToContact(
		"contact-id",
		[]string{"GENERAL", "TECHNICAL"},
	)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestAssignPrimaryRolesToContactError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.AssignPrimaryRolesToContact(
		"contact-id",
		[]string{"GENERAL", "TECHNICAL"},
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestAssignPrimaryRolesToContactError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.AssignPrimaryRolesToContact(
		"contact-id",
		[]string{"GENERAL", "TECHNICAL"},
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestAssignPrimaryRolesToContactError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "404", "errorMessage": "Resource not found"}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.AssignPrimaryRolesToContact(
		"contact-id",
		[]string{"GENERAL", "TECHNICAL"},
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "404")
}

func TestAssignPrimaryRolesToContactError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.AssignPrimaryRolesToContact(
		"contact-id",
		[]string{"GENERAL", "TECHNICAL"},
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestAssignPrimaryRolesToContactError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	customerAccountApi := CustomerAccountApi{}
	err := customerAccountApi.AssignPrimaryRolesToContact(
		"contact-id",
		[]string{"GENERAL", "TECHNICAL"},
	)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}
