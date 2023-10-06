# SoftPOS
(*SoftPOS*)

## Overview

softPOS' agent and order management system now supported by APIs

### Available Operations

* [CreateTerminals](#createterminals) - Create Terminal
* [GetTerminalByMobileNumber](#getterminalbymobilenumber) - Get terminal status using phone number

## CreateTerminals

Use this API to create new terminals to use softPOS.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.SoftPOS.CreateTerminals(ctx, operations.CreateTerminalsRequest{
        CreateTerminalRequest: &shared.CreateTerminalRequest{
            TerminalID: gosdkfull.String("opium navigating Schaden"),
            TerminalName: "parse Branding passage",
            TerminalPhoneNo: "siemens",
        },
        XAPIVersion: gosdkfull.String("Account"),
        XClientID: "with",
        XClientSecret: "ivory scalable Tricycle",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TerminalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateTerminalsRequest](../../models/operations/createterminalsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateTerminalsResponse](../../models/operations/createterminalsresponse.md), error**


## GetTerminalByMobileNumber

Use this API to view all details of a terminal.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.SoftPOS.GetTerminalByMobileNumber(ctx, operations.GetTerminalByMobileNumberRequest{
        TerminalPhoneNo: "coulomb Bedfordshire",
        XAPIVersion: gosdkfull.String("Producer Trial"),
        XClientID: "Maine Global",
        XClientSecret: "perspiciatis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TerminalDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetTerminalByMobileNumberRequest](../../models/operations/getterminalbymobilenumberrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetTerminalByMobileNumberResponse](../../models/operations/getterminalbymobilenumberresponse.md), error**

