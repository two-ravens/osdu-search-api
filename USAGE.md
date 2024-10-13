<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.SearchAPI.QueryWithCursor(ctx, "<id>", components.CursorQueryRequest{
		Kind: components.Kind{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->