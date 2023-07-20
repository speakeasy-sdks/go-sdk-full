# Settlements

### Available Operations

* [Getsettlements](#getsettlements) - Get Settlements by Order ID
* [PostSettlements](#postsettlements) - Get All Settlements

## Getsettlements

Use this API to view all the settlements of a particular order.

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
    res, err := s.Settlements.Getsettlements(ctx, operations.GetsettlementsRequest{
        OrderID: "modi",
        XAPIVersion: pglatest.String("praesentium"),
        XClientID: "rem",
        XClientSecret: "voluptates",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SettlementsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetsettlementsRequest](../../models/operations/getsettlementsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetsettlementsResponse](../../models/operations/getsettlementsresponse.md), error**


## PostSettlements

Use this API to get all settlement details by specifying the settlement ID, settlement UTR or date range.

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
    res, err := s.Settlements.PostSettlements(ctx, operations.PostSettlementsRequest{
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    921158,
                },
                EndDate: pglatest.String("sint"),
                SettlementUtrs: []string{
                    "itaque",
                },
                StartDate: pglatest.String("incidunt"),
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Cursor: pglatest.String("enim"),
                Limit: 9356,
            },
        },
        XAPIVersion: pglatest.String("est"),
        XClientID: "quibusdam",
        XClientSecret: "explicabo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchSettlement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PostSettlementsRequest](../../models/operations/postsettlementsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PostSettlementsResponse](../../models/operations/postsettlementsresponse.md), error**

