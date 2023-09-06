# SoftPOS

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
	"github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
)

func main() {
    s := pglatest.New()

    ctx := context.Background()
    res, err := s.SoftPOS.CreateTerminals(ctx, operations.CreateTerminalsRequest{
        CreateTerminalRequest: &shared.CreateTerminalRequest{
            TerminalID: pglatest.String("modi"),
            TerminalName: "praesentium",
            TerminalPhoneNo: "rem",
        },
        XAPIVersion: pglatest.String("voluptates"),
        XClientID: "quasi",
        XClientSecret: "repudiandae",
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
	"github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := pglatest.New()

    ctx := context.Background()
    res, err := s.SoftPOS.GetTerminalByMobileNumber(ctx, operations.GetTerminalByMobileNumberRequest{
        TerminalPhoneNo: "sint",
        XAPIVersion: pglatest.String("veritatis"),
        XClientID: "itaque",
        XClientSecret: "incidunt",
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

