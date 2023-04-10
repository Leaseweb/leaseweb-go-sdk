package main

import (
	"context"
	"fmt"

	"github.com/LeaseWeb/leaseweb-go-sdk"
)

func main() {

	// exampleInitLeasewebClient()
	// exampleSetDefaultAccount()
	// exampleSetDefaultAccount()
}

func exampleInitLeasewebClient() {

	ctx := context.Background()

	leaseweb.InitLeasewebClient("API_KEY")
	result, _ := leaseweb.CustomerAccountApi{}.Get(ctx)
	fmt.Println(result)
}

func exampleSetDefaultAccount() {

	ctx := context.Background()

	account := leaseweb.Account{ApiKey: "API_KEY"}
	leaseweb.SetDefaultAccount(account)
	result, _ := leaseweb.CustomerAccountApi{}.Get(ctx)
	fmt.Println(result)
}

func exampleSetContextWithAccount() {

	ctx := context.Background()

	account := leaseweb.Account{ApiKey: "API_KEY_1"}
	leaseweb.SetDefaultAccount(account)
	result, _ := leaseweb.CustomerAccountApi{}.Get(ctx)
	fmt.Println(result)

	newAccount := leaseweb.Account{ApiKey: "API_KEY_2"}
	newCtx := leaseweb.SetContextWithAccount(ctx, newAccount)
	result, _ = leaseweb.CustomerAccountApi{}.Get(newCtx)
	fmt.Println(result)

	result, _ = leaseweb.CustomerAccountApi{}.Get(ctx)
	fmt.Println(result)
}
