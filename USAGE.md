<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
)

func main() {
    s := pglatest.New()

    ctx := context.Background()
    res, err := s.Authentication.OTPRequest(ctx, operations.OTPRequestRequest{
        OTPRequest: &shared.OTPRequest{
            Action: shared.OTPRequestActionResendOtp,
            Otp: "provident",
        },
        PaymentID: "distinctio",
        XAPIVersion: pglatest.String("quibusdam"),
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