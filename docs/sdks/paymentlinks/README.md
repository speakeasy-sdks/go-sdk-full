# PaymentLinks
(*PaymentLinks*)

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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.CancelPaymentLink(ctx, operations.CancelPaymentLinkRequest{
        LinkID: "string",
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CancelPaymentLinkRequest](../../pkg/models/operations/cancelpaymentlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CancelPaymentLinkResponse](../../pkg/models/operations/cancelpaymentlinkresponse.md), error**
| Error Object                 | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.LinkCancelledError | 400                          | application/json             |
| sdkerrors.SDKError           | 400-600                      | */*                          |

## CreatePaymentLink

Use this API to create a new payment link. The created payment link url will be available in the API response parameter link_url.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.CreatePaymentLink(ctx, operations.CreatePaymentLinkRequest{
        CreateLinkRequest: &shared.CreateLinkRequest{
            CustomerDetails: shared.LinkCustomerDetailsEntity{
                CustomerPhone: "string",
            },
            LinkAmount: 5411.24,
            LinkCurrency: "string",
            LinkID: "string",
            LinkMeta: &shared.LinkMetaEntity{},
            LinkNotes: map[string]string{
                "key_1": "value_1",
                "key_2": "value_2",
            },
            LinkNotify: &shared.LinkNotifyEntity{},
            LinkPurpose: "string",
        },
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreatePaymentLinkRequest](../../pkg/models/operations/createpaymentlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreatePaymentLinkResponse](../../pkg/models/operations/createpaymentlinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPaymentLinkDetails

Use this API to view all details and status of a payment link.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.GetPaymentLinkDetails(ctx, operations.GetPaymentLinkDetailsRequest{
        LinkID: "string",
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetPaymentLinkDetailsRequest](../../pkg/models/operations/getpaymentlinkdetailsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetPaymentLinkDetailsResponse](../../pkg/models/operations/getpaymentlinkdetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPaymentLinkOrders

Use this API to view all order details for a payment link.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.PaymentLinks.GetPaymentLinkOrders(ctx, operations.GetPaymentLinkOrdersRequest{
        LinkID: "string",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetPaymentLinkOrdersRequest](../../pkg/models/operations/getpaymentlinkordersrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetPaymentLinkOrdersResponse](../../pkg/models/operations/getpaymentlinkordersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
