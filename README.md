# Leaseweb Go SDK
Leaseweb Go SDK provides a golang client for Leaseweb services.

## Installation
```bash
go get github.com/Leaseweb/leaseweb-go-sdk/publiccloud # Replace publiccloud with the package you want to use.
```

## Generate your API Key
You can generate your API key at the [Customer Portal](https://secure.leaseweb.com/).

## Example
```bash
go get github.com/Leaseweb/leaseweb-go-sdk/publiccloud
```

```golang
package main

import (
	"context"
	"fmt"
	"os"

	apiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func main() {
	ctx := context.WithValue(
		context.Background(),
		apiclient.ContextAPIKeys,
		map[string]apiclient.APIKey{
			"X-LSW-Auth": {Key: "API_KEY_STRING"},
		},
	)

	configuration := apiclient.NewConfiguration()
	apiClient := apiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubliccloudAPI.GetInstanceList(ctx).Execute()
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error when calling `PubliccloudAPI.GetInstanceList``: %v\n", err)
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		if err != nil {
			return
		}
	}
	// response from `GetInstanceList`: GetInstanceListResult
	_, err = fmt.Fprintf(os.Stdout, "Response from `PubliccloudAPI.GetInstanceList`: %v\n", resp)
}
```

### Documentation
Documentation for the Leaseweb Go SDK Services can be found [here](https://pkg.go.dev/github.com/leaseweb/leaseweb-go-sdk#section-directories).

## Issues
If you encounter an issue with the project, you are welcome to submit an [issue](https://github.com/Leaseweb/leaseweb-go-sdk/issues).

## Pull requests
We ask that an issue is always opened prior to a PR to give us the opportunity to discuss the best place for the change before investing your effort on a patch that we may not be able to accept.

The code in the repository is generated from the OpenAPI specification. This means PRs to code, tests and sometimes even markdown are often inappropriate, and we may need to make changes to the specification instead.

