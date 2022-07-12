package leaseweb

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const CUSTOMER_ACCOUNT_API_DEFAULT_VERSION = "v1"

type CustomerAccountApi struct {
	version string
}

func (cai *CustomerAccountApi) SetVersion(version string) {
	cai.version = version
}

func (cai CustomerAccountApi) GetVersion() string {
	if cai.version == "" {
		return CUSTOMER_ACCOUNT_API_DEFAULT_VERSION
	}
	return cai.version
}

func (cai CustomerAccountApi) GetPath() string {
	return "/account"
}

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

func (cai CustomerAccountApi) GetCustomerAccount() (CustomerAccount, error) {

	customerAccount := &CustomerAccount{}
	r := leasewebRequest{
		response: customerAccount,
		method:   GET,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/details",
	}
	err := doRequest(r)
	if err != nil {
		return CustomerAccount{}, err
	}

	return *customerAccount, nil
}

func (cai CustomerAccountApi) UpdateCustomerAccount(ad Address) error {

	r := leasewebRequest{
		payload:  map[string]Address{"address": ad},
		method:   PUT,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/details",
	}
	err := doRequest(r)
	if err != nil {
		return err
	}

	return nil
}

func (cai CustomerAccountApi) ListContacts(args ...interface{}) (Contacts, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	contacts := &Contacts{}
	r := leasewebRequest{
		response: contacts,
		method:   GET,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/contacts" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return Contacts{}, err
	}

	return *contacts, nil
}

func (cai CustomerAccountApi) CreateContact(newContact Contact) (Contact, error) {

	contact := &Contact{}
	r := leasewebRequest{
		payload:  newContact,
		response: contact,
		method:   POST,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/contacts",
	}
	err := doRequest(r)
	if err != nil {
		return Contact{}, err
	}

	return *contact, nil
}

func (cai CustomerAccountApi) DeleteContact(contactId string) error {

	r := leasewebRequest{
		method:   DELETE,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/contacts/" + contactId,
	}
	err := doRequest(r)
	if err != nil {
		return err
	}

	return nil
}

func (cai CustomerAccountApi) GetContact(contactId string) (Contact, error) {

	contact := &Contact{}
	r := leasewebRequest{
		response: contact,
		method:   GET,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/contacts/" + contactId,
	}
	err := doRequest(r)
	if err != nil {
		return Contact{}, err
	}

	return *contact, nil
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

	r := leasewebRequest{
		payload:  payload,
		method:   PUT,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/contacts/" + contactId,
	}
	err := doRequest(r)
	if err != nil {
		return err
	}

	return nil
}

func (cai CustomerAccountApi) AssignPrimaryRolesToContact(contactId string, roles []string) error {

	r := leasewebRequest{
		payload:  map[string][]string{"roles": roles},
		method:   POST,
		endpoint: cai.GetPath() + "/" + cai.GetVersion() + "/contacts/" + contactId,
	}
	err := doRequest(r)
	if err != nil {
		return err
	}

	return nil
}
