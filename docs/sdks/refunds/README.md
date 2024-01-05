# Refunds
(*Refunds*)

### Available Operations

* [Createrefund](#createrefund) - Create Refund
* [GetRefund](#getrefund) - Get Refund
* [Getallrefundsfororder](#getallrefundsfororder) - Get All Refunds for an Order

## Createrefund

Use this API to initiate refunds.

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Refunds.Createrefund(ctx, operations.CreaterefundRequest{
        CreateRefundRequest: &shared.CreateRefundRequest{
            RefundAmount: 567.71,
            RefundID: "string",
            RefundSplits: []shared.VendorSplit{
                shared.VendorSplit{},
            },
        },
        OrderID: "string",
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreaterefundRequest](../../pkg/models/operations/createrefundrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreaterefundResponse](../../pkg/models/operations/createrefundresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRefund

Use this API to fetch a specific refund processed on your Cashfree Account.

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Refunds.GetRefund(ctx, operations.GetRefundRequest{
        OrderID: "string",
        RefundID: "string",
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetRefundRequest](../../pkg/models/operations/getrefundrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetRefundResponse](../../pkg/models/operations/getrefundresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Getallrefundsfororder

Use this API to fetch all refunds processed against an order.

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Refunds.Getallrefundsfororder(ctx, operations.GetallrefundsfororderRequest{
        OrderID: "string",
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetallrefundsfororderRequest](../../pkg/models/operations/getallrefundsfororderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetallrefundsfororderResponse](../../pkg/models/operations/getallrefundsfororderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
