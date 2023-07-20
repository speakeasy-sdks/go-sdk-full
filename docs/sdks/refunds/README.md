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
            RefundAmount: 9764.6,
            RefundID: "vero",
            RefundNote: pglatest.String("nihil"),
            RefundSpeed: shared.CreateRefundRequestRefundSpeedInstant.ToPointer(),
            RefundSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: pglatest.Float64(557.14),
                    Percentage: pglatest.Float64(6048.46),
                    VendorID: pglatest.String("voluptate"),
                },
                shared.VendorSplit{
                    Amount: pglatest.Float64(7392.64),
                    Percentage: pglatest.Float64(199.87),
                    VendorID: pglatest.String("doloremque"),
                },
                shared.VendorSplit{
                    Amount: pglatest.Float64(4417.11),
                    Percentage: pglatest.Float64(2828.07),
                    VendorID: pglatest.String("maiores"),
                },
                shared.VendorSplit{
                    Amount: pglatest.Float64(1201.96),
                    Percentage: pglatest.Float64(3594.44),
                    VendorID: pglatest.String("dolore"),
                },
            },
        },
        OrderID: "iusto",
        XAPIVersion: pglatest.String("dicta"),
        XClientID: "harum",
        XClientSecret: "enim",
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
        OrderID: "accusamus",
        RefundID: "commodi",
        XAPIVersion: pglatest.String("repudiandae"),
        XClientID: "quae",
        XClientSecret: "ipsum",
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
        OrderID: "quidem",
        XAPIVersion: pglatest.String("molestias"),
        XClientID: "excepturi",
        XClientSecret: "pariatur",
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

