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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.CancelPaymentLink(ctx, operations.CancelPaymentLinkRequest{
        LinkID: "perferendis",
        XAPIVersion: gosdkfull.String("ad"),
        XClientID: "natus",
        XClientSecret: "sed",
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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.CreatePaymentLink(ctx, operations.CreatePaymentLinkRequest{
        CreateLinkRequest: &shared.CreateLinkRequest{
            CustomerDetails: shared.LinkCustomerDetailsEntity{
                CustomerEmail: gosdkfull.String("iste"),
                CustomerName: gosdkfull.String("dolor"),
                CustomerPhone: "natus",
            },
            LinkAmount: 3864.89,
            LinkAutoReminders: gosdkfull.Bool(false),
            LinkCurrency: "hic",
            LinkExpiryTime: gosdkfull.String("saepe"),
            LinkID: "fuga",
            LinkMeta: &shared.LinkMetaEntity{
                NotifyURL: gosdkfull.String("in"),
                PaymentMethods: gosdkfull.String("corporis"),
                ReturnURL: gosdkfull.String("iste"),
                UpiIntent: gosdkfull.Bool(false),
            },
            LinkMinimumPartialAmount: gosdkfull.Float64(4370.32),
            LinkNotes: map[string]string{
                "saepe": "quidem",
            },
            LinkNotify: &shared.LinkNotifyEntity{
                SendEmail: gosdkfull.Bool(false),
                SendSms: gosdkfull.Bool(false),
            },
            LinkPartialPayments: gosdkfull.Bool(false),
            LinkPurpose: "architecto",
        },
        XAPIVersion: gosdkfull.String("ipsa"),
        XClientID: "reiciendis",
        XClientSecret: "est",
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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.GetPaymentLinkDetails(ctx, operations.GetPaymentLinkDetailsRequest{
        LinkID: "mollitia",
        XAPIVersion: gosdkfull.String("laborum"),
        XClientID: "dolores",
        XClientSecret: "dolorem",
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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.GetPaymentLinkOrders(ctx, operations.GetPaymentLinkOrdersRequest{
        LinkID: "corporis",
        XAPIVersion: gosdkfull.String("explicabo"),
        XClientID: "nobis",
        XClientSecret: "enim",
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

