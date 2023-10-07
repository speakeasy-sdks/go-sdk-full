# github.com/speakeasy-sdks/go-sdk-full

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/go-sdk-full
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
    res, err := s.Authentication.OTPRequest(ctx, operations.OTPRequestRequest{
        OTPRequest: &shared.OTPRequest{
            Action: shared.OTPRequestActionSubmitOtp,
            Otp: "Tricycle pace",
        },
        PaymentID: "Nobelium Planner",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OTPResponseEntity != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Authentication](docs/sdks/authentication/README.md)

* [OTPRequest](docs/sdks/authentication/README.md#otprequest) - Submit or Resend OTP

### [EligibilityAPIs](docs/sdks/eligibilityapis/README.md)

* [EligibilityCardlessEMI](docs/sdks/eligibilityapis/README.md#eligibilitycardlessemi) - Get eligible Cardless EMI
* [EligibilityOffer](docs/sdks/eligibilityapis/README.md#eligibilityoffer) - Get eligible Offers
* [EligibilityPaylater](docs/sdks/eligibilityapis/README.md#eligibilitypaylater) - Get eligible Paylater

### [Offers](docs/sdks/offers/README.md)

* [CreateOffer](docs/sdks/offers/README.md#createoffer) - Create Offer
* [GetOffer](docs/sdks/offers/README.md#getoffer) - Get Offer by ID

### [Orders](docs/sdks/orders/README.md)

* [CreateOrder](docs/sdks/orders/README.md#createorder) - Create Order
* [GetOrder](docs/sdks/orders/README.md#getorder) - Get Order
* [OrderPay](docs/sdks/orders/README.md#orderpay) - Order Pay
* [Preauthorization](docs/sdks/orders/README.md#preauthorization) - Preauthorization

### [PaymentLinks](docs/sdks/paymentlinks/README.md)

* [CancelPaymentLink](docs/sdks/paymentlinks/README.md#cancelpaymentlink) - Cancel Payment Link
* [CreatePaymentLink](docs/sdks/paymentlinks/README.md#createpaymentlink) - Create Payment Link
* [GetPaymentLinkDetails](docs/sdks/paymentlinks/README.md#getpaymentlinkdetails) - Fetch Payment Link Details
* [GetPaymentLinkOrders](docs/sdks/paymentlinks/README.md#getpaymentlinkorders) - Get Orders for a Payment Link

### [Payments](docs/sdks/payments/README.md)

* [GetPaymentbyID](docs/sdks/payments/README.md#getpaymentbyid) - Get Payment by ID
* [GetPaymentsfororder](docs/sdks/payments/README.md#getpaymentsfororder) - Get Payments for an Order

### [Reconciliation](docs/sdks/reconciliation/README.md)

* [PostRecon](docs/sdks/reconciliation/README.md#postrecon) - PG Reconciliation
* [PostSettlementRecon](docs/sdks/reconciliation/README.md#postsettlementrecon) - Settlement Reconciliation

### [Refunds](docs/sdks/refunds/README.md)

* [Createrefund](docs/sdks/refunds/README.md#createrefund) - Create Refund
* [GetRefund](docs/sdks/refunds/README.md#getrefund) - Get Refund
* [Getallrefundsfororder](docs/sdks/refunds/README.md#getallrefundsfororder) - Get All Refunds for an Order

### [Settlements](docs/sdks/settlements/README.md)

* [Getsettlements](docs/sdks/settlements/README.md#getsettlements) - Get Settlements by Order ID
* [PostSettlements](docs/sdks/settlements/README.md#postsettlements) - Get All Settlements

### [TokenVault](docs/sdks/tokenvault/README.md)

* [DeleteSpecificSavedInstrument](docs/sdks/tokenvault/README.md#deletespecificsavedinstrument) - Delete Saved Instrument
* [FetchAllSavedInstruments](docs/sdks/tokenvault/README.md#fetchallsavedinstruments) - Fetch All Saved Instruments
* [FetchCryptogram](docs/sdks/tokenvault/README.md#fetchcryptogram) - Fetch cryptogram for saved instrument
* [FetchSpecificSavedInstrument](docs/sdks/tokenvault/README.md#fetchspecificsavedinstrument) - Fetch Single Saved Instrument

### [SoftPOS](docs/sdks/softpos/README.md)

* [CreateTerminals](docs/sdks/softpos/README.md#createterminals) - Create Terminal
* [GetTerminalByMobileNumber](docs/sdks/softpos/README.md#getterminalbymobilenumber) - Get terminal status using phone number
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
