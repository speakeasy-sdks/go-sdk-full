# Refunds

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
	"context"
	"log"
	"github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
)

func main() {
    s := pglatest.New()

    ctx := context.Background()
    res, err := s.Refunds.Createrefund(ctx, operations.CreaterefundRequest{
        CreateRefundRequest: &shared.CreateRefundRequest{
            RefundAmount: 2444.25,
            RefundID: "error",
            RefundNote: pglatest.String("quia"),
            RefundSpeed: shared.CreateRefundRequestRefundSpeedStandard.ToPointer(),
            RefundSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: pglatest.Float64(1103.75),
                    Percentage: pglatest.Float64(6747.52),
                    VendorID: pglatest.String("animi"),
                },
            },
        },
        OrderID: "enim",
        XAPIVersion: pglatest.String("odit"),
        XClientID: "quo",
        XClientSecret: "sequi",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreaterefundRequest](../../models/operations/createrefundrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreaterefundResponse](../../models/operations/createrefundresponse.md), error**


## GetRefund

Use this API to fetch a specific refund processed on your Cashfree Account.

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
    res, err := s.Refunds.GetRefund(ctx, operations.GetRefundRequest{
        OrderID: "tenetur",
        RefundID: "ipsam",
        XAPIVersion: pglatest.String("id"),
        XClientID: "possimus",
        XClientSecret: "aut",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetRefundRequest](../../models/operations/getrefundrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetRefundResponse](../../models/operations/getrefundresponse.md), error**


## Getallrefundsfororder

Use this API to fetch all refunds processed against an order.

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
    res, err := s.Refunds.Getallrefundsfororder(ctx, operations.GetallrefundsfororderRequest{
        OrderID: "quasi",
        XAPIVersion: pglatest.String("error"),
        XClientID: "temporibus",
        XClientSecret: "laborum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundsEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetallrefundsfororderRequest](../../models/operations/getallrefundsfororderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetallrefundsfororderResponse](../../models/operations/getallrefundsfororderresponse.md), error**

