package leaseweb

import (
	"fmt"
	"net/url"
)

const INVOICE_API_DEFAULT_VERSION = "v1"

type InvoiceApi struct {
	version string
}

func (ia *InvoiceApi) SetVersion(version string) {
	ia.version = version
}

func (ia InvoiceApi) GetVersion() string {
	if ia.version == "" {
		return INVOICE_API_DEFAULT_VERSION
	}
	return ia.version
}

func (ia InvoiceApi) GetPath() string {
	return "/invoices"
}

type Invoice struct {
	Currency                string   `json:"currency"`
	Date                    string   `json:"date"`
	DueDate                 string   `json:"dueDate"`
	Id                      string   `json:"id"`
	IsPartialPaymentAllowed bool     `json:"isPartialPaymentAllowed"`
	OpenAmount              float32  `json:"openAmount"`
	Status                  string   `json:"status"`
	TaxAmount               float32  `json:"taxAmount"`
	Total                   float32  `json:"total"`
	Credits                 []Credit `json:"credits"`
	Lines                   []Line   `json:"lineItems"`
}

type Credit struct {
	Date      string  `json:"date"`
	Id        string  `json:"id"`
	TaxAmount float32 `json:"taxAmount"`
	Total     float32 `json:"total"`
}

type Line struct {
	ContractId  string  `json:"contractId"`
	EquipmentId string  `json:"equipmentId"`
	Product     string  `json:"product"`
	Quantity    int     `json:"quantity"`
	Reference   string  `json:"reference"`
	TotalAmount float32 `json:"totalAmount"`
	UnitAmount  float32 `json:"unitAmount"`
}

type Metadata struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"totalCount"`
}

type Contract struct {
	ContractId  string  `json:"contractId"`
	Currency    string  `json:"currency"`
	EndDate     string  `json:"endDate"`
	EquipmentId string  `json:"equipmentId"`
	PoNumber    string  `json:"poNumber"`
	Price       float32 `json:"price"`
	Product     string  `json:"product"`
	Reference   string  `json:"reference"`
	StartDate   string  `json:"startDate"`
}

type Invoices struct {
	Invoices []Invoice `json:"invoices"`
	Metadata Metadata  `json:"_metadata"`
}

type ProForma struct {
	Currency     string     `json:"currency"`
	ProformaDate string     `json:"proformaDate"`
	SubTotal     float32    `json:"subTotal"`
	Total        float32    `json:"total"`
	VatAmount    float32    `json:"vatAmount"`
	Metadata     Metadata   `json:"_metadata"`
	Contracts    []Contract `json:"contractItems"`
}

func (ia InvoiceApi) ListInvoices(args ...int) (Invoices, error) {

	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	invoices := &Invoices{}
	r := leasewebRequest{
		response: invoices,
		method:   GET,
		endpoint: ia.GetPath() + "/" + ia.GetVersion() + "/invoices" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return Invoices{}, err
	}

	return *invoices, nil
}

func (ia InvoiceApi) GetProForma(args ...int) (ProForma, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	proForma := &ProForma{}
	r := leasewebRequest{
		response: proForma,
		method:   GET,
		endpoint: ia.GetPath() + "/" + ia.GetVersion() + "/invoices/proforma" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return ProForma{}, err
	}

	return *proForma, nil
}

func (ia InvoiceApi) GetInvoice(invoiceId string) (Invoice, error) {

	invoice := &Invoice{}
	r := leasewebRequest{
		response: invoice,
		method:   GET,
		endpoint: ia.GetPath() + "/" + ia.GetVersion() + "/invoices/" + invoiceId,
	}
	err := doRequest(r)
	if err != nil {
		return Invoice{}, err
	}

	return *invoice, nil
}
