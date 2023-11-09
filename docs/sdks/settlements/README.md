# Settlements
(*Settlements*)

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
	gosdkfull "github.com/speakeasy-sdks/go-sdk-full/v2"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/operations"
)

func main() {
    s := gosdkfull.New()

    ctx := context.Background()
    res, err := s.Settlements.Getsettlements(ctx, operations.GetsettlementsRequest{
        OrderID: "string",
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetsettlementsRequest](../../pkg/models/operations/getsettlementsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetsettlementsResponse](../../pkg/models/operations/getsettlementsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSettlements

Use this API to get all settlement details by specifying the settlement ID, settlement UTR or date range.

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
    res, err := s.Settlements.PostSettlements(ctx, operations.PostSettlementsRequest{
        FetchSettlementReconRequest: &shared.FetchSettlementReconRequest{
            Filters: shared.FetchSettlementReconRequestFilters{
                CfSettlementIds: []int64{
                    297769,
                },
                SettlementUtrs: []string{
                    "string",
                },
            },
            Pagination: shared.FetchSettlementReconRequestPagination{
                Limit: 116172,
            },
        },
        XClientID: "string",
        XClientSecret: "string",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PostSettlementsRequest](../../pkg/models/operations/postsettlementsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PostSettlementsResponse](../../pkg/models/operations/postsettlementsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
