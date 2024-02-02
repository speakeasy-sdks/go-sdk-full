<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v3"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/operations"
	"log"
)

func main() {
	s := gosdkfull.New()

	ctx := context.Background()
	res, err := s.TokenVault.DeleteSpecificSavedInstrument(ctx, operations.DeleteSpecificSavedInstrumentRequest{
		CustomerID:    "string",
		InstrumentID:  "string",
		XClientID:     "string",
		XClientSecret: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.FetchAllSavedInstruments != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->