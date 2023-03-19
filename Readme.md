# Leaseweb Go SDK

This package provides a golang client for Leaseweb's REST API.

### Installation

```bash
go get github.com/LeaseWeb/leaseweb-go-sdk
```

### Generate your API Key
You can generate your API key at the [Customer Portal](https://secure.leaseweb.com/)

### Authenticate and make first call
```golang
import LSW "github.com/LeaseWeb/leaseweb-go-sdk"
import "context"

func main() {
    LSW.InitLeasewebClient("api key here")
    ctx := context.Background()
    result, err := LSW.DedicatedServerApi{}.List(ctx)
    if err == nil {
        fmt.Println(result)
    }
}
```

### Documentation
The full documentation for Leaseweb go SDK is available in the [wiki](https://github.com/LeaseWeb/leaseweb-go-sdk/wiki).

