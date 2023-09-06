# Authentication

## Overview

The Authentication API allows merchants to show a native screen and capture OTP on their own page and submit to Cashfree. This feature is only available on request.

### Available Operations

* [OTPRequest](#otprequest) - Submit or Resend OTP

## OTPRequest

If you accept OTP on your own page, you can use the below API to send OTP to Cashfree.

### Example Usage

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
            Otp: "nulla",
        },
        PaymentID: "corrupti",
        XAPIVersion: pglatest.String("illum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OTPResponseEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.OTPRequestRequest](../../models/operations/otprequestrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.OTPRequestResponse](../../models/operations/otprequestresponse.md), error**
