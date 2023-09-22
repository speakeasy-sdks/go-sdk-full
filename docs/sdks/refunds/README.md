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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Refunds.Createrefund(ctx, operations.CreaterefundRequest{
        CreateRefundRequest: &shared.CreateRefundRequest{
            RefundAmount: 1103.75,
            RefundID: "laborum",
            RefundNote: gosdkfull.String("animi"),
            RefundSpeed: shared.CreateRefundRequestRefundSpeedStandard.ToPointer(),
            RefundSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: gosdkfull.Float64(1381.83),
                    Percentage: gosdkfull.Float64(7783.46),
                    VendorID: gosdkfull.String("sequi"),
                },
            },
        },
        OrderID: "tenetur",
        XAPIVersion: gosdkfull.String("ipsam"),
        XClientID: "id",
        XClientSecret: "possimus",
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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Refunds.GetRefund(ctx, operations.GetRefundRequest{
        OrderID: "aut",
        RefundID: "quasi",
        XAPIVersion: gosdkfull.String("error"),
        XClientID: "temporibus",
        XClientSecret: "laborum",
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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Refunds.Getallrefundsfororder(ctx, operations.GetallrefundsfororderRequest{
        OrderID: "quasi",
        XAPIVersion: gosdkfull.String("reiciendis"),
        XClientID: "voluptatibus",
        XClientSecret: "vero",
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

