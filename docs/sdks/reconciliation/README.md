# Reconciliation
(*Reconciliation*)

### Available Operations

* [PostRecon](#postrecon) - PG Reconciliation
* [PostSettlementRecon](#postsettlementrecon) - Settlement Reconciliation

## PostRecon

Use this API to get the payment gateway reconciliation details with date range.

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Reconciliation.PostRecon(ctx, operations.PostReconRequest{
        FetchPGReconRequest: &shared.FetchPGReconRequest{
            Filters: shared.Filters{
                EndDate: "string",
                StartDate: "string",
            },
            Pagination: shared.Pagination{
                Limit: 85382,
            },
        },
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchPGRecon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PostReconRequest](../../pkg/models/operations/postreconrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PostReconResponse](../../pkg/models/operations/postreconresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostSettlementRecon

Use this API to get settlement reconciliation details using Settlement ID, settlement UTR or date range.

### Example Usage

```go
package main

import(
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"context"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Reconciliation.PostSettlementRecon(ctx, operations.PostSettlementReconRequest{
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    956121,
                },
                SettlementUtrs: []string{
                    "string",
                },
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Limit: 71166,
            },
        },
        XClientID: "string",
        XClientSecret: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchSettlementRecon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostSettlementReconRequest](../../pkg/models/operations/postsettlementreconrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PostSettlementReconResponse](../../pkg/models/operations/postsettlementreconresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
