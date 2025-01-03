// Package leasewebgosdk provides a golang client for Leaseweb's REST API.
package leasewebgosdk

import (
	_ "github.com/leaseweb/leaseweb-go-sdk/abuse"
	_ "github.com/leaseweb/leaseweb-go-sdk/aggregationpack"
	_ "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
	_ "github.com/leaseweb/leaseweb-go-sdk/dns"
	_ "github.com/leaseweb/leaseweb-go-sdk/invoice"
	_ "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
}
