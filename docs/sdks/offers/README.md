# Offers

### Available Operations

* [CreateOffer](#createoffer) - Create Offer
* [GetOffer](#getoffer) - Get Offer by ID

## CreateOffer

Use this API to create offers with Cashfree from your backend

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
    res, err := s.Offers.CreateOffer(ctx, operations.CreateOfferRequest{
        CreateOfferBackendRequest: &shared.CreateOfferBackendRequest{
            OfferDetails: shared.OfferDetails{
                CashbackDetails: &shared.CashbackDetails{
                    CashbackType: shared.CashbackDetailsCashbackTypeFlat.ToPointer(),
                    CashbackValue: pglatest.String("molestiae"),
                    MaxCashbackAmount: "minus",
                },
                DiscountDetails: &shared.DiscountDetails{
                    DiscountType: shared.DiscountDetailsDiscountTypePercentage,
                    DiscountValue: "voluptatum",
                    MaxDiscountAmount: "iusto",
                },
                OfferType: shared.OfferDetailsOfferTypeDiscountAndCashback,
            },
            OfferMeta: shared.OfferMeta{
                OfferCode: "CFTESTOFFER",
                OfferDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
                OfferEndTime: "2023-03-29T08:09:51Z",
                OfferStartTime: "2023-03-21T08:09:51Z",
                OfferTitle: "Test Offer",
            },
            OfferTnc: shared.OfferTnc{
                OfferTncType: shared.OfferTncOfferTncTypePost,
                OfferTncValue: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
            },
            OfferValidations: shared.OfferValidations{
                MaxAllowed: "10",
                MinAmount: pglatest.String("1"),
                PaymentMethod: shared.OfferValidationsPaymentMethod{},
            },
        },
        XAPIVersion: pglatest.String("nisi"),
        XClientID: "recusandae",
        XClientSecret: "temporibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OfferEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOfferRequest](../../models/operations/createofferrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateOfferResponse](../../models/operations/createofferresponse.md), error**


## GetOffer

Use this API to get offer by offer_id

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
    res, err := s.Offers.GetOffer(ctx, operations.GetOfferRequest{
        OfferID: "ab",
        XAPIVersion: pglatest.String("quis"),
        XClientID: "veritatis",
        XClientSecret: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OfferEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetOfferRequest](../../models/operations/getofferrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetOfferResponse](../../models/operations/getofferresponse.md), error**
