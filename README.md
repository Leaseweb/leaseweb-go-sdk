# Leaseweb Go SDK
Leaseweb Go SDK provides a golang client for Leaseweb's REST API.

### Installation

```bash
go get github.com/LeaseWeb/leaseweb-go-sdk/publicCloud # for publicCloud

go get github.com/LeaseWeb/leaseweb-go-sdk/dedicatedServer # for dedicatedServer
```

### Generate your API Key
You can generate your API key at the [Customer Portal](https://secure.leaseweb.com/)

### Example
```golang
package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "github.com/leaseweb/leaseweb-go-sdk/publicCloud"
)

func main() {

	ctx := context.WithValue(
		context.Background(),
		openapiclient.ContextAPIKeys,
		map[string]openapiclient.APIKey{
			"X-LSW-Auth": {Key: "API_KEY_STRING"},
		},
	)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicCloudAPI.GetInstanceList(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicCloudAPI.GetInstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceList`: GetInstanceListResult
	fmt.Fprintf(os.Stdout, "Response from `PublicCloudAPI.GetInstanceList`: %v\n", resp)
}
```

### Documentation
The Leaseweb Go SDK documentation based on product:

- [Public Cloud](publicCloud/README.md)
- [Dedicated Server](dedicatedServer/README.md)
- [Aggregation Pack](aggregationPack/README.md)
- [Abuse](abuse/README.md)
- [Invoice](invoice/README.md)

### Issues
If you encounter an issue with the project, you are welcome to submit an [issue](https://github.com/Leaseweb/leaseweb-go-sdk/issues).

### Pull requests
We ask that an issue is always opened prior to a PR to give us the opportunity to discuss the best place for the change before investing your effort on a patch that we may not be able to accept.

The code in the repository is generated from the OpenAPI specification. This means PRs to code, tests and sometimes even markdown are often inappropriate, and we may need to make changes to the specification instead.

