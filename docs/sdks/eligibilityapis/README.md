# EligibilityAPIs

### Available Operations

* [EligibilityCardlessEMI](#eligibilitycardlessemi) - Get eligible Cardless EMI
* [EligibilityOffer](#eligibilityoffer) - Get eligible Offers
* [EligibilityPaylater](#eligibilitypaylater) - Get eligible Paylater

## EligibilityCardlessEMI

Use this API to get eligible Cardless EMI Payment Methods for a customer on an order.

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
    res, err := s.EligibilityAPIs.EligibilityCardlessEMI(ctx, operations.EligibilityCardlessEMIRequest{
        EligibilityCardlessEMIRequest: &shared.EligibilityCardlessEMIRequest{
            Queries: shared.CardlessEMIQueries{
                Amount: pglatest.Float64(100),
                CustomerDetails: &shared.CustomerDetailsCardlessEMI{
                    CustomerPhone: "9898989898",
                },
                OrderID: pglatest.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
            },
        },
        XAPIVersion: pglatest.String("vel"),
        XClientID: "error",
        XClientSecret: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EligibleCardlessEMIEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.EligibilityCardlessEMIRequest](../../models/operations/eligibilitycardlessemirequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.EligibilityCardlessEMIResponse](../../models/operations/eligibilitycardlessemiresponse.md), error**


## EligibilityOffer

Use this API to get eligible offers for an order or amount.

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
    res, err := s.EligibilityAPIs.EligibilityOffer(ctx, operations.EligibilityOfferRequest{
        EligibilityOffersRequest: &shared.EligibilityOffersRequest{
            Filters: &shared.OfferFilters{
                OfferType: []shared.OfferType{
                    shared.OfferTypeCashback,
                    shared.OfferTypeCashback,
                },
            },
            Queries: shared.OfferQueries{
                Amount: pglatest.Float64(100),
                OrderID: pglatest.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
            },
        },
        XAPIVersion: pglatest.String("debitis"),
        XClientID: "ipsa",
        XClientSecret: "delectus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EligibleOffersEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.EligibilityOfferRequest](../../models/operations/eligibilityofferrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.EligibilityOfferResponse](../../models/operations/eligibilityofferresponse.md), error**


## EligibilityPaylater

Use this API to get eligible Paylater Payment Methods for a customer on an order.

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
    res, err := s.EligibilityAPIs.EligibilityPaylater(ctx, operations.EligibilityPaylaterRequest{
        EligibilityCardlessEMIRequest: &shared.EligibilityCardlessEMIRequest{
            Queries: shared.CardlessEMIQueries{
                Amount: pglatest.Float64(100),
                CustomerDetails: &shared.CustomerDetailsCardlessEMI{
                    CustomerPhone: "9898989898",
                },
                OrderID: pglatest.String("order_413462PK1RI1IwYB1X69LgzUQWiSxYDF"),
            },
        },
        XAPIVersion: pglatest.String("tempora"),
        XClientID: "suscipit",
        XClientSecret: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EligiblePaylaters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.EligibilityPaylaterRequest](../../models/operations/eligibilitypaylaterrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.EligibilityPaylaterResponse](../../models/operations/eligibilitypaylaterresponse.md), error**

