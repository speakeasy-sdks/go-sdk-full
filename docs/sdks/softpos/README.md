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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v3"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.SoftPOS.CreateTerminals(ctx, operations.CreateTerminalsRequest{
        CreateTerminalRequest: &shared.CreateTerminalRequest{
            TerminalID: gosdkfull.String("1"),
            TerminalName: "Jane Doe",
            TerminalPhoneNo: "9876543210",
        },
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateTerminalsRequest](../../pkg/models/operations/createterminalsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateTerminalsResponse](../../pkg/models/operations/createterminalsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTerminalByMobileNumber

Use this API to view all details of a terminal.

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v3"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.SoftPOS.GetTerminalByMobileNumber(ctx, operations.GetTerminalByMobileNumberRequest{
        TerminalPhoneNo: "string",
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetTerminalByMobileNumberRequest](../../pkg/models/operations/getterminalbymobilenumberrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetTerminalByMobileNumberResponse](../../pkg/models/operations/getterminalbymobilenumberresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
