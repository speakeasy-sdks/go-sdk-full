# Orders
(*.Orders*)

### Available Operations

* [CreateOrder](#createorder) - Create Order
* [GetOrder](#getorder) - Get Order
* [OrderPay](#orderpay) - Order Pay
* [Preauthorization](#preauthorization) - Preauthorization

## CreateOrder

Use this API to create orders with Cashfree from your backend and get the payment link. To use this API S2S flag needs to be enabled from the backend. In case you want to use the cards payment option the PCI DSS flag is required, for more information email us at "care@cashfree.com".

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/callbacks"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
            CustomerDetails: shared.CustomerDetails{
                CustomerID: "string",
                CustomerPhone: "string",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: gosdkfull.String("2021-07-29T00:00:00Z"),
            OrderMeta: &shared.OrderMeta{},
            OrderNote: gosdkfull.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{},
            },
            OrderTags: map[string]string{
                "key": "string",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "string",
                TerminalPhoneNo: "string",
                TerminalType: "string",
            },
        },
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOrderRequest](../../models/operations/createorderrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**


## GetOrder

Use this API to view all details of an order.

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
    res, err := s.Orders.GetOrder(ctx, operations.GetOrderRequest{
        OrderID: "string",
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetOrderRequest](../../models/operations/getorderrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetOrderResponse](../../models/operations/getorderresponse.md), error**


## OrderPay

Use this API when you have already created the orders and want Cashfree to process the payment. To use this API S2S flag needs to be enabled from the backend. In case you want to use the cards payment option the PCI DSS flag is required, for more information send an email to "care@cashfree.com".

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
    res, err := s.Orders.OrderPay(ctx, operations.OrderPayRequest{
        OrderPayRequest: &shared.OrderPayRequest{
            OfferID: gosdkfull.String("faa6cc05-d1e2-401c-b0cf-0c9db3ff0f0b"),
            PaymentMethod: shared.CreateOrderPayRequestPaymentMethodCardlessEMIPaymentMethod(
                    shared.CardlessEMIPaymentMethod{
                        CardlessEmi: shared.CardlessEMI{},
                    },
            ),
            PaymentSessionID: "session__CvcEmNKDkmERQrxnx39ibhJ3Ii034pjc8ZVxf3qcgEXCWlgDDlHRgz2XYZCqpajDQSXMMtCusPgOIxYP2LZx0-05p39gC2Vgmq1RAj--gcn",
        },
        XAPIVersion: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrderPayResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.OrderPayRequest](../../models/operations/orderpayrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.OrderPayResponse](../../models/operations/orderpayresponse.md), error**


## Preauthorization

Use this API to capture or void a preauthorized payment

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
    res, err := s.Orders.Preauthorization(ctx, operations.PreauthorizationRequest{
        AuthorizationRequest: &shared.AuthorizationRequest{},
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
| `request`                                                                                | [operations.PreauthorizationRequest](../../models/operations/preauthorizationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PreauthorizationResponse](../../models/operations/preauthorizationresponse.md), error**

