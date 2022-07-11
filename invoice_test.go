package leaseweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListInvoices(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "invoices": [
			{
				"currency": "EUR",
				"date": "2022-01-06T00:00:00+00:00",
				"dueDate": "2022-02-30T00:00:00+00:00",
				"id": "00000001",
				"isPartialPaymentAllowed": true,
				"openAmount": 1756.21,
				"status": "OVERDUE",
				"taxAmount": 0,
				"total": 1756.21
			},
			{
				"currency": "EUR",
				"date": "2022-03-06T00:00:00+00:00",
				"dueDate": "2022-04-30T00:00:00+00:00",
				"id": "00000002",
				"isPartialPaymentAllowed": true,
				"openAmount": 34,
				"status": "OPEN",
				"taxAmount": 0,
				"total": 34
			}
		]}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	response, err := invoiceApi.ListInvoices()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Invoices), 2)

	invoices1 := response.Invoices[0]
	assert.Equal(invoices1.Currency, "EUR")
	assert.Equal(invoices1.Date, "2022-01-06T00:00:00+00:00")
	assert.Equal(invoices1.DueDate, "2022-02-30T00:00:00+00:00")
	assert.Equal(invoices1.IsPartialPaymentAllowed, true)
	assert.Equal(invoices1.OpenAmount, float32(1756.21))
	assert.Equal(invoices1.Status, "OVERDUE")
	assert.Equal(invoices1.TaxAmount, float32(0.0))
	assert.Equal(invoices1.Total, float32(1756.21))

	invoices2 := response.Invoices[1]
	assert.Equal(invoices2.Currency, "EUR")
	assert.Equal(invoices2.Date, "2022-03-06T00:00:00+00:00")
	assert.Equal(invoices2.DueDate, "2022-04-30T00:00:00+00:00")
	assert.Equal(invoices2.Id, "00000002")
	assert.Equal(invoices2.IsPartialPaymentAllowed, true)
	assert.Equal(invoices2.OpenAmount, float32(34))
	assert.Equal(invoices2.Status, "OPEN")
	assert.Equal(invoices2.TaxAmount, float32(0.0))
	assert.Equal(invoices2.Total, float32(34))
}

func TestListInvoicesBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "invoices": []}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	response, err := invoiceApi.ListInvoices()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Invoices), 0)
}

func TestListInvoicesPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "invoices": [{"id": "00000001"}]}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	response, err := invoiceApi.ListInvoices(1)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Invoices), 1)
}

func TestListInvoicesError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.ListInvoices()

	assert := assert.New(t)

	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestListInvoicesError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.ListInvoices()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestListInvoicesError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.ListInvoices()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestListInvoicesError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.ListInvoices()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestGetProForma(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w,
			`{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2},
			"contractItems": [
				{
					"contractId": "50000103",
					"currency": "EUR",
					"endDate": "2022-01-01T00:00:00+00:00",
					"equipmentId": "26430",
					"poNumber": "40002154000110",
					"price": 151.05,
					"product": "DEDICATED SERVER",
					"reference": "this is a reference 1",
					"startDate": "2022-03-01T00:00:00+00:00"
				},
				{
					"contractId": "50000104",
					"currency": "EUR",
					"endDate": "2021-01-01T00:00:00+00:00",
					"equipmentId": "26431",
					"poNumber": "40002154000111",
					"price": 150.05,
					"product": "ATS",
					"reference": "this is a reference 2",
					"startDate": "2020-02-01T00:00:00+00:00"
				}
			],
			"currency": "EUR",
			"proformaDate": "2021-07-01T00:00:00+00:00",
			"subTotal": 300.04,
			"total": 352.55,
			"vatAmount": 52.51
		}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	response, err := invoiceApi.GetProForma()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Contracts), 2)

	assert.Equal(response.Currency, "EUR")
	assert.Equal(response.ProformaDate, "2021-07-01T00:00:00+00:00")
	assert.Equal(response.SubTotal, float32(300.04))
	assert.Equal(response.Total, float32(352.55))
	assert.Equal(response.VatAmount, float32(52.51))

	contract1 := response.Contracts[0]
	assert.Equal(contract1.Currency, "EUR")
	assert.Equal(contract1.ContractId, "50000103")
	assert.Equal(contract1.EndDate, "2022-01-01T00:00:00+00:00")
	assert.Equal(contract1.EquipmentId, "26430")
	assert.Equal(contract1.PoNumber, "40002154000110")
	assert.Equal(contract1.Price, float32(151.05))
	assert.Equal(contract1.Product, "DEDICATED SERVER")
	assert.Equal(contract1.Reference, "this is a reference 1")
	assert.Equal(contract1.StartDate, "2022-03-01T00:00:00+00:00")

	contract2 := response.Contracts[1]
	assert.Equal(contract2.Currency, "EUR")
	assert.Equal(contract2.ContractId, "50000104")
	assert.Equal(contract2.EndDate, "2021-01-01T00:00:00+00:00")
	assert.Equal(contract2.EquipmentId, "26431")
	assert.Equal(contract2.PoNumber, "40002154000111")
	assert.Equal(contract2.Price, float32(150.05))
	assert.Equal(contract2.Product, "ATS")
	assert.Equal(contract2.Reference, "this is a reference 2")
	assert.Equal(contract2.StartDate, "2020-02-01T00:00:00+00:00")
}

func TestGetProFormaPaginate(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w,
			`{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11},
			"contractItems": [
				{
					"contractId": "50000103",
					"currency": "EUR",
					"endDate": "2022-01-01T00:00:00+00:00",
					"equipmentId": "26430",
					"poNumber": "40002154000110",
					"price": 151.05,
					"product": "DEDICATED SERVER",
					"reference": "this is a reference 1",
					"startDate": "2022-03-01T00:00:00+00:00"
				}
			],
			"currency": "EUR",
			"proformaDate": "2021-07-01T00:00:00+00:00",
			"subTotal": 300.04,
			"total": 352.55,
			"vatAmount": 52.51
		}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	response, err := invoiceApi.GetProForma(1)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Contracts), 1)

	assert.Equal(response.Currency, "EUR")
	assert.Equal(response.ProformaDate, "2021-07-01T00:00:00+00:00")
	assert.Equal(response.SubTotal, float32(300.04))
	assert.Equal(response.Total, float32(352.55))
	assert.Equal(response.VatAmount, float32(52.51))

	contract1 := response.Contracts[0]
	assert.Equal(contract1.Currency, "EUR")
	assert.Equal(contract1.ContractId, "50000103")
	assert.Equal(contract1.EndDate, "2022-01-01T00:00:00+00:00")
	assert.Equal(contract1.EquipmentId, "26430")
	assert.Equal(contract1.PoNumber, "40002154000110")
	assert.Equal(contract1.Price, float32(151.05))
	assert.Equal(contract1.Product, "DEDICATED SERVER")
	assert.Equal(contract1.Reference, "this is a reference 1")
	assert.Equal(contract1.StartDate, "2022-03-01T00:00:00+00:00")
}

func TestGetProFormaError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetProForma()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestGetProFormaError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetProForma()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestGetProFormaError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetProForma()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestGetProFormaError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetProForma()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestGetInvoice(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `
			{"credits": [
				{
					"date": "2019-05-06T00:00:00+00:00",
					"id": "00001211",
					"taxAmount": 1.5,
					"total": 15
				}
			],
			"currency": "EUR",
			"date": "2022-07-01T00:00:00+00:00",
			"dueDate": "2019-05-30T00:00:00+00:00",
			"id": "00000001",
			"isPartialPaymentAllowed": true,
			"lineItems": [
				{
					"contractId": "12345678",
					"equipmentId": "1234",
					"product": "Rackspace",
					"quantity": 1,
					"reference": "This is a reference",
					"totalAmount": 151.5,
					"unitAmount": 152.5
				}
			],
			"openAmount": 1751.21,
			"status": "OPEN",
			"taxAmount": 0,
			"total": 1756.21
		}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	response, err := invoiceApi.GetInvoice("00000001")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(len(response.Credits), 1)
	assert.Equal(len(response.Lines), 1)

	assert.Equal(response.Currency, "EUR")
	assert.Equal(response.Date, "2022-07-01T00:00:00+00:00")
	assert.Equal(response.DueDate, "2019-05-30T00:00:00+00:00")
	assert.Equal(response.Id, "00000001")
	assert.Equal(response.IsPartialPaymentAllowed, true)
	assert.Equal(response.OpenAmount, float32(1751.21))
	assert.Equal(response.Status, "OPEN")
	assert.Equal(response.TaxAmount, float32(0))
	assert.Equal(response.Total, float32(1756.21))

	assert.Equal(response.Credits[0].Date, "2019-05-06T00:00:00+00:00")
	assert.Equal(response.Credits[0].Id, "00001211")
	assert.Equal(response.Credits[0].TaxAmount, float32(1.5))
	assert.Equal(response.Credits[0].Total, float32(15))

	assert.Equal(response.Lines[0].ContractId, "12345678")
	assert.Equal(response.Lines[0].EquipmentId, "1234")
	assert.Equal(response.Lines[0].Product, "Rackspace")
	assert.Equal(response.Lines[0].Quantity, 1)
	assert.Equal(response.Lines[0].Reference, "This is a reference")
	assert.Equal(response.Lines[0].TotalAmount, float32(151.5))
	assert.Equal(response.Lines[0].UnitAmount, float32(152.5))
}

func TestGetInvoiceError401(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "401", "errorMessage": "You are not authorized to view this resource."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetInvoice("000000001")

	assert := assert.New(t)

	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "You are not authorized to view this resource.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "401")
}

func TestGetInvoiceError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "403", "errorMessage": "Access to the requested resource is forbidden."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetInvoice("000000001")

	assert := assert.New(t)

	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Access to the requested resource is forbidden.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "403")
}

func TestGetInvoiceError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The API could not handle your request at this time."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetInvoice("000000001")

	assert := assert.New(t)

	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API could not handle your request at this time.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestGetInvoiceError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"correlationId":"289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The API is not available at the moment."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.GetInvoice("000000001")

	assert := assert.New(t)

	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The API is not available at the moment.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}
