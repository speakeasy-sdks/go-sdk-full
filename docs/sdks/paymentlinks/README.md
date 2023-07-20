# PaymentLinks

### Available Operations

* [CancelPaymentLink](#cancelpaymentlink) - Cancel Payment Link
* [CreatePaymentLink](#createpaymentlink) - Create Payment Link
* [GetPaymentLinkDetails](#getpaymentlinkdetails) - Fetch Payment Link Details
* [GetPaymentLinkOrders](#getpaymentlinkorders) - Get Orders for a Payment Link

## CancelPaymentLink

Use this API to cancel a payment link. No further payments can be done against a cancelled link. Only a link in ACTIVE status can be cancelled.

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
    res, err := s.PaymentLinks.CancelPaymentLink(ctx, operations.CancelPaymentLinkRequest{
        LinkID: "laboriosam",
        XAPIVersion: pglatest.String("hic"),
        XClientID: "saepe",
        XClientSecret: "fuga",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkCancelledResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CancelPaymentLinkRequest](../../models/operations/cancelpaymentlinkrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CancelPaymentLinkResponse](../../models/operations/cancelpaymentlinkresponse.md), error**


## CreatePaymentLink

Use this API to create a new payment link. The created payment link url will be available in the API response parameter link_url.

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
    res, err := s.PaymentLinks.CreatePaymentLink(ctx, operations.CreatePaymentLinkRequest{
        CreateLinkRequest: &shared.CreateLinkRequest{
            CustomerDetails: shared.LinkCustomerDetailsEntity{
                CustomerEmail: pglatest.String("in"),
                CustomerName: pglatest.String("corporis"),
                CustomerPhone: "iste",
            },
            LinkAmount: 4370.32,
            LinkAutoReminders: pglatest.Bool(false),
            LinkCurrency: "saepe",
            LinkExpiryTime: pglatest.String("quidem"),
            LinkID: "architecto",
            LinkMeta: &shared.LinkMetaEntity{
                NotifyURL: pglatest.String("ipsa"),
                PaymentMethods: pglatest.String("reiciendis"),
                ReturnURL: pglatest.String("est"),
                UpiIntent: pglatest.Bool(false),
            },
            LinkMinimumPartialAmount: pglatest.Float64(6531.4),
            LinkNotes: map[string]string{
                "dolores": "dolorem",
                "corporis": "explicabo",
                "nobis": "enim",
            },
            LinkNotify: &shared.LinkNotifyEntity{
                SendEmail: pglatest.Bool(false),
                SendSms: pglatest.Bool(false),
            },
            LinkPartialPayments: pglatest.Bool(false),
            LinkPurpose: "omnis",
        },
        XAPIVersion: pglatest.String("nemo"),
        XClientID: "minima",
        XClientSecret: "excepturi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreatePaymentLinkRequest](../../models/operations/createpaymentlinkrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreatePaymentLinkResponse](../../models/operations/createpaymentlinkresponse.md), error**


## GetPaymentLinkDetails

Use this API to view all details and status of a payment link.

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
    res, err := s.PaymentLinks.GetPaymentLinkDetails(ctx, operations.GetPaymentLinkDetailsRequest{
        LinkID: "accusantium",
        XAPIVersion: pglatest.String("iure"),
        XClientID: "culpa",
        XClientSecret: "doloribus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPaymentLinkDetailsRequest](../../models/operations/getpaymentlinkdetailsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetPaymentLinkDetailsResponse](../../models/operations/getpaymentlinkdetailsresponse.md), error**


## GetPaymentLinkOrders

Use this API to view all order details for a payment link.

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
    res, err := s.PaymentLinks.GetPaymentLinkOrders(ctx, operations.GetPaymentLinkOrdersRequest{
        LinkID: "sapiente",
        XAPIVersion: pglatest.String("architecto"),
        XClientID: "mollitia",
        XClientSecret: "dolorem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkOrdersResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPaymentLinkOrdersRequest](../../models/operations/getpaymentlinkordersrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetPaymentLinkOrdersResponse](../../models/operations/getpaymentlinkordersresponse.md), error**

