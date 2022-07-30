package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const CUSTOMER_ACCOUNT_API_VERSION = "v1"

type CustomerAccountApi struct{}

type CustomerAccount struct {
	Name         string  `json:"name"`
	ResellerTier string  `json:"resellerTier"`
	VatNumber    string  `json:"vatNumber"`
	Address      Address `json:"address"`
}

type Address struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalCode"`
	State       string `json:"state"`
	StateCode   string `json:"stateCode"`
	Street      string `json:"street"`
}

type Contacts struct {
	Contacts []Contact `json:"contacts"`
	Metadata Metadata  `json:"_metadata"`
}

type Contact struct {
	Description  string   `json:"description"`
	Email        string   `json:"email"`
	FirstName    string   `json:"firstName"`
	Id           string   `json:"id"`
	LastName     string   `json:"lastName"`
	Roles        []string `json:"roles"`
	PrimaryRoles []string `json:"primaryRoles"`
	Mobile       Phone    `json:"mobile"`
	Phone        Phone    `json:"phone"`
}

type Phone struct {
	CountryCode string `json:"countryCode"`
	DialCode    string `json:"dialCode"`
	Number      string `json:"number"`
}

func (cai CustomerAccountApi) getPath(endpoint string) string {
	return "/account/" + CUSTOMER_ACCOUNT_API_VERSION + endpoint
}

func (cai CustomerAccountApi) GetCustomerAccount() (*CustomerAccount, error) {
	path := cai.getPath("/details")
	result := &CustomerAccount{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) UpdateCustomerAccount(ad Address) error {
	path := cai.getPath("/details")
	payload := map[string]Address{"address": ad}
	return doRequest(http.MethodPut, path, nil, payload)
}

func (cai CustomerAccountApi) ListContacts(args ...interface{}) (*Contacts, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		s := reflect.ValueOf(args[2])
		var primaryRoles []string
		for i := 0; i < s.Len(); i++ {
			primaryRoles = append(primaryRoles, s.Index(i).String())
		}
		v.Add("primaryRoles", strings.Join(primaryRoles, ","))
	}

	path := cai.getPath("/contacts?" + v.Encode())
	result := &Contacts{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) CreateContact(newContact Contact) (*Contact, error) {
	path := cai.getPath("/contacts")
	result := &Contact{}
	if err := doRequest(http.MethodPost, path, result, newContact); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) DeleteContact(contactId string) error {
	path := cai.getPath("/contacts/" + contactId)
	return doRequest(http.MethodDelete, path)
}

func (cai CustomerAccountApi) GetContact(contactId string) (*Contact, error) {
	path := cai.getPath("/contacts" + contactId)
	result := &Contact{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) UpdateContact(contactId string, phone Phone, roles []string, args ...interface{}) error {
	payload := make(map[string]interface{})
	payload["phone"] = phone
	payload["roles"] = roles
	if len(args) >= 1 {
		payload["mobile"] = args[0].(Phone)
	}
	if len(args) >= 2 {
		payload["description"] = fmt.Sprint(args[1])
	}

	path := cai.getPath("/contacts" + contactId)
	return doRequest(http.MethodPut, path, nil, payload)
}

func (cai CustomerAccountApi) AssignPrimaryRolesToContact(contactId string, roles []string) error {
	payload := map[string][]string{"roles": roles}
	path := cai.getPath("/contacts" + contactId)
	return doRequest(http.MethodPost, path, nil, payload)
}
