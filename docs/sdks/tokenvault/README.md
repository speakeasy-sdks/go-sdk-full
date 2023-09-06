# TokenVault

## Overview

Cashfree's token Vault helps you save cards and tokenize them in a PCI complaint manner. We support creation of network tokens which can be used across acquiring banks

### Available Operations

* [DeleteSpecificSavedInstrument](#deletespecificsavedinstrument) - Delete Saved Instrument
* [FetchAllSavedInstruments](#fetchallsavedinstruments) - Fetch All Saved Instruments
* [FetchCryptogram](#fetchcryptogram) - Fetch cryptogram for saved instrument
* [FetchSpecificSavedInstrument](#fetchspecificsavedinstrument) - Fetch Single Saved Instrument

## DeleteSpecificSavedInstrument

To delete a saved instrument for a customer id and instrument id

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
    res, err := s.TokenVault.DeleteSpecificSavedInstrument(ctx, operations.DeleteSpecificSavedInstrumentRequest{
        CustomerID: "reprehenderit",
        InstrumentID: "ut",
        XAPIVersion: pglatest.String("maiores"),
        XClientID: "dicta",
        XClientSecret: "corporis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchAllSavedInstruments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteSpecificSavedInstrumentRequest](../../models/operations/deletespecificsavedinstrumentrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.DeleteSpecificSavedInstrumentResponse](../../models/operations/deletespecificsavedinstrumentresponse.md), error**


## FetchAllSavedInstruments

To get all saved instruments for a customer id

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
    res, err := s.TokenVault.FetchAllSavedInstruments(ctx, operations.FetchAllSavedInstrumentsRequest{
        CustomerID: "dolore",
        InstrumentType: operations.FetchAllSavedInstrumentsInstrumentTypeCard,
        XAPIVersion: pglatest.String("iusto"),
        XClientID: "dicta",
        XClientSecret: "harum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchAllSavedInstruments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.FetchAllSavedInstrumentsRequest](../../models/operations/fetchallsavedinstrumentsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.FetchAllSavedInstrumentsResponse](../../models/operations/fetchallsavedinstrumentsresponse.md), error**


## FetchCryptogram

To get the card network token, token expiry and cryptogram for a saved instrument using instrument id

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
    res, err := s.TokenVault.FetchCryptogram(ctx, operations.FetchCryptogramRequest{
        CustomerID: "enim",
        InstrumentID: "accusamus",
        XAPIVersion: pglatest.String("commodi"),
        XClientID: "repudiandae",
        XClientSecret: "quae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Cryptogram != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.FetchCryptogramRequest](../../models/operations/fetchcryptogramrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.FetchCryptogramResponse](../../models/operations/fetchcryptogramresponse.md), error**


## FetchSpecificSavedInstrument

To get specific saved instrument for a customer id and instrument id

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
    res, err := s.TokenVault.FetchSpecificSavedInstrument(ctx, operations.FetchSpecificSavedInstrumentRequest{
        CustomerID: "ipsum",
        InstrumentID: "quidem",
        XAPIVersion: pglatest.String("molestias"),
        XClientID: "excepturi",
        XClientSecret: "pariatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchAllSavedInstruments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.FetchSpecificSavedInstrumentRequest](../../models/operations/fetchspecificsavedinstrumentrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.FetchSpecificSavedInstrumentResponse](../../models/operations/fetchspecificsavedinstrumentresponse.md), error**
