# Payments

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
	"github.com/speakeasy-sdks/go-sdk-full"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/operations"
)

func main() {
    s := pglatest.New()

    ctx := context.Background()
    res, err := s.Payments.GetPaymentbyID(ctx, operations.GetPaymentbyIDRequest{
        CfPaymentID: 358152,
        OrderID: "explicabo",
        XAPIVersion: pglatest.String("nobis"),
        XClientID: "enim",
        XClientSecret: "omnis",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetPaymentbyIDRequest](../../models/operations/getpaymentbyidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetPaymentbyIDResponse](../../models/operations/getpaymentbyidresponse.md), error**


## GetPaymentsfororder

Use this API to view all payment details for an order.

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
    res, err := s.Payments.GetPaymentsfororder(ctx, operations.GetPaymentsfororderRequest{
        OrderID: "nemo",
        XAPIVersion: pglatest.String("minima"),
        XClientID: "excepturi",
        XClientSecret: "accusantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPaymentsfororder200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetPaymentsfororderRequest](../../models/operations/getpaymentsfororderrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetPaymentsfororderResponse](../../models/operations/getpaymentsfororderresponse.md), error**

