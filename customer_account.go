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
	Name         string                 `json:"name"`
	ResellerTier string                 `json:"resellerTier"`
	VatNumber    string                 `json:"vatNumber"`
	Address      CustomerAccountAddress `json:"address"`
}

type CustomerAccountAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalCode"`
	State       string `json:"state"`
	StateCode   string `json:"stateCode"`
	Street      string `json:"street"`
}

type CustomerAccountContacts struct {
	Contacts []CustomerAccountContact `json:"contacts"`
	Metadata Metadata                 `json:"_metadata"`
}

type CustomerAccountContact struct {
	Description  string               `json:"description"`
	Email        string               `json:"email"`
	FirstName    string               `json:"firstName"`
	Id           string               `json:"id"`
	LastName     string               `json:"lastName"`
	Roles        []string             `json:"roles"`
	PrimaryRoles []string             `json:"primaryRoles"`
	Mobile       CustomerAccountPhone `json:"mobile"`
	Phone        CustomerAccountPhone `json:"phone"`
}

type CustomerAccountPhone struct {
	CountryCode string `json:"countryCode"`
	DialCode    string `json:"dialCode"`
	Number      string `json:"number"`
}

func (cai CustomerAccountApi) getPath(endpoint string) string {
	return "/account/" + CUSTOMER_ACCOUNT_API_VERSION + endpoint
}

func (cai CustomerAccountApi) Get() (*CustomerAccount, error) {
	path := cai.getPath("/details")
	result := &CustomerAccount{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) Update(ad CustomerAccountAddress) error {
	path := cai.getPath("/details")
	payload := map[string]CustomerAccountAddress{"address": ad}
	return doRequest(http.MethodPut, path, nil, payload)
}

func (cai CustomerAccountApi) ListContacts(args ...interface{}) (*CustomerAccountContacts, error) {
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
	result := &CustomerAccountContacts{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) CreateContact(newContact CustomerAccountContact) (*CustomerAccountContact, error) {
	path := cai.getPath("/contacts")
	result := &CustomerAccountContact{}
	if err := doRequest(http.MethodPost, path, result, newContact); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) DeleteContact(contactId string) error {
	path := cai.getPath("/contacts/" + contactId)
	return doRequest(http.MethodDelete, path)
}

func (cai CustomerAccountApi) GetContact(contactId string) (*CustomerAccountContact, error) {
	path := cai.getPath("/contacts/" + contactId)
	result := &CustomerAccountContact{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) UpdateContact(contactId string, phone CustomerAccountPhone, roles []string, args ...interface{}) error {
	payload := make(map[string]interface{})
	payload["phone"] = phone
	payload["roles"] = roles
	if len(args) >= 1 {
		payload["mobile"] = args[0].(CustomerAccountPhone)
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
