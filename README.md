# github.com/speakeasy-sdks/go-sdk-full

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/go-sdk-full
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
	s := gosdkfull.New()

	ctx := context.Background()
	res, err := s.TokenVault.DeleteSpecificSavedInstrument(ctx, operations.DeleteSpecificSavedInstrumentRequest{
		CustomerID:    "string",
		InstrumentID:  "string",
		XClientID:     "string",
		XClientSecret: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.FetchAllSavedInstruments != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [TokenVault](docs/sdks/tokenvault/README.md)

* [DeleteSpecificSavedInstrument](docs/sdks/tokenvault/README.md#deletespecificsavedinstrument) - Delete Saved Instrument
* [FetchAllSavedInstruments](docs/sdks/tokenvault/README.md#fetchallsavedinstruments) - Fetch All Saved Instruments
* [FetchCryptogram](docs/sdks/tokenvault/README.md#fetchcryptogram) - Fetch cryptogram for saved instrument
* [FetchSpecificSavedInstrument](docs/sdks/tokenvault/README.md#fetchspecificsavedinstrument) - Fetch Single Saved Instrument

### [EligibilityAPIs](docs/sdks/eligibilityapis/README.md)

* [EligibilityCardlessEMI](docs/sdks/eligibilityapis/README.md#eligibilitycardlessemi) - Get eligible Cardless EMI
* [EligibilityOffer](docs/sdks/eligibilityapis/README.md#eligibilityoffer) - Get eligible Offers
* [EligibilityPaylater](docs/sdks/eligibilityapis/README.md#eligibilitypaylater) - Get eligible Paylater

### [PaymentLinks](docs/sdks/paymentlinks/README.md)

* [CancelPaymentLink](docs/sdks/paymentlinks/README.md#cancelpaymentlink) - Cancel Payment Link
* [CreatePaymentLink](docs/sdks/paymentlinks/README.md#createpaymentlink) - Create Payment Link
* [GetPaymentLinkDetails](docs/sdks/paymentlinks/README.md#getpaymentlinkdetails) - Fetch Payment Link Details
* [GetPaymentLinkOrders](docs/sdks/paymentlinks/README.md#getpaymentlinkorders) - Get Orders for a Payment Link

### [Offers](docs/sdks/offers/README.md)

* [CreateOffer](docs/sdks/offers/README.md#createoffer) - Create Offer
* [GetOffer](docs/sdks/offers/README.md#getoffer) - Get Offer by ID

### [Orders](docs/sdks/orders/README.md)

* [CreateOrder](docs/sdks/orders/README.md#createorder) - Create Order
* [GetOrder](docs/sdks/orders/README.md#getorder) - Get Order
* [OrderPay](docs/sdks/orders/README.md#orderpay) - Order Pay
* [Preauthorization](docs/sdks/orders/README.md#preauthorization) - Preauthorization

### [Authentication](docs/sdks/authentication/README.md)

* [OTPRequest](docs/sdks/authentication/README.md#otprequest) - Submit or Resend OTP

### [Payments](docs/sdks/payments/README.md)

* [GetPaymentbyID](docs/sdks/payments/README.md#getpaymentbyid) - Get Payment by ID
* [GetPaymentsfororder](docs/sdks/payments/README.md#getpaymentsfororder) - Get Payments for an Order

### [Refunds](docs/sdks/refunds/README.md)

* [Createrefund](docs/sdks/refunds/README.md#createrefund) - Create Refund
* [GetRefund](docs/sdks/refunds/README.md#getrefund) - Get Refund
* [Getallrefundsfororder](docs/sdks/refunds/README.md#getallrefundsfororder) - Get All Refunds for an Order

### [Settlements](docs/sdks/settlements/README.md)

* [Getsettlements](docs/sdks/settlements/README.md#getsettlements) - Get Settlements by Order ID
* [PostSettlements](docs/sdks/settlements/README.md#postsettlements) - Get All Settlements

### [Reconciliation](docs/sdks/reconciliation/README.md)

* [PostRecon](docs/sdks/reconciliation/README.md#postrecon) - PG Reconciliation
* [PostSettlementRecon](docs/sdks/reconciliation/README.md#postsettlementrecon) - Settlement Reconciliation

### [SoftPOS](docs/sdks/softpos/README.md)

* [CreateTerminals](docs/sdks/softpos/README.md#createterminals) - Create Terminal
* [GetTerminalByMobileNumber](docs/sdks/softpos/README.md#getterminalbymobilenumber) - Get terminal status using phone number
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                 | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.LinkCancelledError | 400                          | application/json             |
| sdkerrors.SDKError           | 4xx-5xx                      | */*                          |

### Example

```go
package main

import (
	"context"
	"errors"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/sdkerrors"
	"log"
)

func main() {
	s := gosdkfull.New()

	ctx := context.Background()
	res, err := s.PaymentLinks.CancelPaymentLink(ctx, operations.CancelPaymentLinkRequest{
		LinkID:        "string",
		XClientID:     "string",
		XClientSecret: "string",
	})
	if err != nil {

		var e *sdkerrors.LinkCancelledError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://sandbox.cashfree.com/pg` | None |
| 1 | `https://api.cashfree.com/pg` | None |

#### Example

```go
package main

import (
	"context"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
	s := gosdkfull.New(
		gosdkfull.WithServerIndex(1),
	)

	ctx := context.Background()
	res, err := s.TokenVault.DeleteSpecificSavedInstrument(ctx, operations.DeleteSpecificSavedInstrumentRequest{
		CustomerID:    "string",
		InstrumentID:  "string",
		XClientID:     "string",
		XClientSecret: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.FetchAllSavedInstruments != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
	s := gosdkfull.New(
		gosdkfull.WithServerURL("https://sandbox.cashfree.com/pg"),
	)

	ctx := context.Background()
	res, err := s.TokenVault.DeleteSpecificSavedInstrument(ctx, operations.DeleteSpecificSavedInstrumentRequest{
		CustomerID:    "string",
		InstrumentID:  "string",
		XClientID:     "string",
		XClientSecret: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.FetchAllSavedInstruments != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
