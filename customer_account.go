package leaseweb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
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

type CustomerAccountListContactsOptions struct {
	PaginationOptions
	PrimaryRoles []string `param:"primaryRoles"`
}

func (cai CustomerAccountApi) getPath(endpoint string) string {
	return "/account/" + CUSTOMER_ACCOUNT_API_VERSION + endpoint
}

func (cai CustomerAccountApi) Get(ctx context.Context) (*CustomerAccount, error) {
	path := cai.getPath("/details")
	result := &CustomerAccount{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) Update(ctx context.Context, ad CustomerAccountAddress) error {
	path := cai.getPath("/details")
	payload := map[string]CustomerAccountAddress{"address": ad}
	return doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (cai CustomerAccountApi) ListContacts(ctx context.Context, opts CustomerAccountListContactsOptions) (*CustomerAccountContacts, error) {
	path := cai.getPath("/contacts")
	query := options.Encode(opts)
	result := &CustomerAccountContacts{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) CreateContact(ctx context.Context, newContact CustomerAccountContact) (*CustomerAccountContact, error) {
	path := cai.getPath("/contacts")
	result := &CustomerAccountContact{}
	if err := doRequest(ctx, http.MethodPost, path, "", result, newContact); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) DeleteContact(ctx context.Context, contactId string) error {
	path := cai.getPath("/contacts/" + contactId)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (cai CustomerAccountApi) GetContact(ctx context.Context, contactId string) (*CustomerAccountContact, error) {
	path := cai.getPath("/contacts/" + contactId)
	result := &CustomerAccountContact{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cai CustomerAccountApi) UpdateContact(ctx context.Context, contactId string, phone CustomerAccountPhone, roles []string, args ...interface{}) error {
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
	return doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (cai CustomerAccountApi) AssignPrimaryRolesToContact(ctx context.Context, contactId string, roles []string) error {
	payload := map[string][]string{"roles": roles}
	path := cai.getPath("/contacts" + contactId)
	return doRequest(ctx, http.MethodPost, path, "", nil, payload)
}
