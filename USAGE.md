<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"log"
)

func main() {
	s := gosdkfull.New()

	ctx := context.Background()
	res, err := s.Authentication.OTPRequest(ctx, operations.OTPRequestRequest{
		OTPRequest: &shared.OTPRequest{
			Action: shared.OTPRequestActionSubmitOtp,
			Otp:    "string",
		},
		PaymentID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.OTPResponseEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->