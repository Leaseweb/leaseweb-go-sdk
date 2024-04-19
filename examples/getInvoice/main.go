package main

import (
	"context"
	"fmt"

	"github.com/Leaseweb/leaseweb-go-sdk/invoice"
)

func main() {

	ctx := context.WithValue(
		context.Background(),
		invoice.ContextAPIKeys,
		map[string]invoice.APIKey{
			"X-LSW-Auth": {Key: "YOUR_LEASEWEB_API_KEY"},
		},
	)

	lsw := invoice.NewAPIClient(invoice.NewConfiguration())
	resp, httpErr, err := lsw.InvoiceAPI.GetInvoice(ctx, "invoice id").Execute()
	fmt.Println(resp, httpErr, err)
}
