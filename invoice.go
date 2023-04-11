package leaseweb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const INVOICE_API_VERSION = "v1"

type InvoiceApi struct{}

type Invoice struct {
	Currency                string          `json:"currency"`
	Date                    string          `json:"date"`
	DueDate                 string          `json:"dueDate"`
	Id                      string          `json:"id"`
	IsPartialPaymentAllowed bool            `json:"isPartialPaymentAllowed"`
	OpenAmount              json.Number     `json:"openAmount"`
	Status                  string          `json:"status"`
	TaxAmount               json.Number     `json:"taxAmount"`
	Total                   json.Number     `json:"total"`
	Credits                 []InvoiceCredit `json:"credits"`
	Lines                   []InvoiceLine   `json:"lineItems"`
}

type InvoiceCredit struct {
	Date      string      `json:"date"`
	Id        string      `json:"id"`
	TaxAmount json.Number `json:"taxAmount"`
	Total     json.Number `json:"total"`
}

type InvoiceLine struct {
	ContractId  string      `json:"contractId"`
	EquipmentId string      `json:"equipmentId"`
	Product     string      `json:"product"`
	Quantity    json.Number `json:"quantity"`
	Reference   string      `json:"reference"`
	TotalAmount json.Number `json:"totalAmount"`
	UnitAmount  json.Number `json:"unitAmount"`
}

type InvoiceContract struct {
	ContractId  string      `json:"contractId"`
	Currency    string      `json:"currency"`
	EndDate     string      `json:"endDate"`
	EquipmentId string      `json:"equipmentId"`
	PoNumber    string      `json:"poNumber"`
	Price       json.Number `json:"price"`
	Product     string      `json:"product"`
	Reference   string      `json:"reference"`
	StartDate   string      `json:"startDate"`
}

type Invoices struct {
	Invoices []Invoice `json:"invoices"`
	Metadata Metadata  `json:"_metadata"`
}

type InvoiceProForma struct {
	Currency     string            `json:"currency"`
	ProformaDate string            `json:"proformaDate"`
	SubTotal     json.Number       `json:"subTotal"`
	Total        json.Number       `json:"total"`
	VatAmount    json.Number       `json:"vatAmount"`
	Metadata     Metadata          `json:"_metadata"`
	Contracts    []InvoiceContract `json:"contractItems"`
}

func (ia InvoiceApi) getPath(endpoint string) string {
	return "/invoices/" + INVOICE_API_VERSION + endpoint
}

func (ia InvoiceApi) List(ctx context.Context, opts PaginationOptions) (*Invoices, error) {
	path := ia.getPath("/invoices")
	query := options.Encode(opts)
	result := &Invoices{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ia InvoiceApi) ListProForma(ctx context.Context, opts PaginationOptions) (*InvoiceProForma, error) {
	path := ia.getPath("/invoices/proforma")
	query := options.Encode(opts)
	result := &InvoiceProForma{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ia InvoiceApi) Get(ctx context.Context, invoiceId string) (*Invoice, error) {
	path := ia.getPath("/invoices/" + invoiceId)
	result := &Invoice{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}
