# Offers
(*Offers*)

### Available Operations

* [CreateOffer](#createoffer) - Create Offer
* [GetOffer](#getoffer) - Get Offer by ID

## CreateOffer

Use this API to create offers with Cashfree from your backend

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v3"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Offers.CreateOffer(ctx, operations.CreateOfferRequest{
        CreateOfferBackendRequest: &shared.CreateOfferBackendRequest{
            OfferDetails: shared.OfferDetails{
                CashbackDetails: &shared.CashbackDetails{
                    MaxCashbackAmount: "string",
                },
                DiscountDetails: &shared.DiscountDetails{
                    DiscountType: shared.DiscountTypePercentage,
                    DiscountValue: "string",
                    MaxDiscountAmount: "string",
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
                OfferTncType: shared.OfferTncTypePost,
                OfferTncValue: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
            },
            OfferValidations: shared.OfferValidations{
                MaxAllowed: "10",
                MinAmount: gosdkfull.String("1"),
                PaymentMethod: shared.CreatePaymentMethodOfferNB(
                        shared.OfferNB{
                            Netbanking: &shared.NBOffer{
                                BankName: "all",
                            },
                        },
                ),
            },
        },
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateOfferRequest](../../pkg/models/operations/createofferrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateOfferResponse](../../pkg/models/operations/createofferresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetOffer

Use this API to get offer by offer_id

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v3"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Offers.GetOffer(ctx, operations.GetOfferRequest{
        OfferID: "string",
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetOfferRequest](../../pkg/models/operations/getofferrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetOfferResponse](../../pkg/models/operations/getofferresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
