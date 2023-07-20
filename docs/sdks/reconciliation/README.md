# Reconciliation

### Available Operations

* [PostRecon](#postrecon) - PG Reconciliation
* [PostSettlementRecon](#postsettlementrecon) - Settlement Reconciliation

## PostRecon

Use this API to get the payment gateway reconciliation details with date range.

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
    res, err := s.Reconciliation.PostRecon(ctx, operations.PostReconRequest{
        FetchPGReconRequest: &shared.FetchPGReconRequest{
            Filters: shared.FetchPGReconRequestFilters{
                EndDate: "velit",
                StartDate: "error",
            },
            Pagination: shared.FetchPGReconRequestPagination{
                Cursor: pglatest.String("quia"),
                Limit: 338007,
            },
        },
        XAPIVersion: pglatest.String("vitae"),
        XClientID: "laborum",
        XClientSecret: "animi",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.PostReconRequest](../../models/operations/postreconrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostReconResponse](../../models/operations/postreconresponse.md), error**


## PostSettlementRecon

Use this API to get settlement reconciliation details using Settlement ID, settlement UTR or date range.

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
    res, err := s.Reconciliation.PostSettlementRecon(ctx, operations.PostSettlementReconRequest{
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    138183,
                    778346,
                },
                EndDate: pglatest.String("sequi"),
                SettlementUtrs: []string{
                    "ipsam",
                    "id",
                    "possimus",
                    "aut",
                },
                StartDate: pglatest.String("quasi"),
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Cursor: pglatest.String("error"),
                Limit: 837945,
            },
        },
        XAPIVersion: pglatest.String("laborum"),
        XClientID: "quasi",
        XClientSecret: "reiciendis",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PostSettlementReconRequest](../../models/operations/postsettlementreconrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PostSettlementReconResponse](../../models/operations/postsettlementreconresponse.md), error**

