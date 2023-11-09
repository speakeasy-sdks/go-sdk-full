# Payments
(*Payments*)

### Available Operations

* [GetPaymentbyID](#getpaymentbyid) - Get Payment by ID
* [GetPaymentsfororder](#getpaymentsfororder) - Get Payments for an Order

## GetPaymentbyID

Use this API to view payment details of an order for a payment ID.

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
    res, err := s.Payments.GetPaymentbyID(ctx, operations.GetPaymentbyIDRequest{
        CfPaymentID: 310675,
        OrderID: "string",
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPaymentbyIDRequest](../../pkg/models/operations/getpaymentbyidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetPaymentbyIDResponse](../../pkg/models/operations/getpaymentbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPaymentsfororder

Use this API to view all payment details for an order.

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
    res, err := s.Payments.GetPaymentsfororder(ctx, operations.GetPaymentsfororderRequest{
        OrderID: "string",
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPaymentsfororderRequest](../../pkg/models/operations/getpaymentsfororderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetPaymentsfororderResponse](../../pkg/models/operations/getpaymentsfororderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
