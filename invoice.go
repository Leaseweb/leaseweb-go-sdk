package leaseweb

import (
	"fmt"
	"net/url"
)

const INVOICE_API_DEFAULT_VERSION = "v1"

type InvoiceApi struct {
	version string
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

func (ia *InvoiceApi) SetVersion(version string) {
	ia.version = version
}

func (ia InvoiceApi) getPath(endpoint string) string {
	version := ia.version
	if version == "" {
		version = INVOICE_API_DEFAULT_VERSION
	}
	return "/invoices/" + version + endpoint
}

func (ia InvoiceApi) ListInvoices(args ...int) (*Invoices, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ia.getPath("/invoices?" + v.Encode())
	result := &Invoices{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ia InvoiceApi) GetProForma(args ...int) (*ProForma, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := ia.getPath("/invoices/proforma?" + v.Encode())
	result := &ProForma{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ia InvoiceApi) GetInvoice(invoiceId string) (*Invoice, error) {
	path := ia.getPath("/invoices/" + invoiceId)
	result := &Invoice{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}
