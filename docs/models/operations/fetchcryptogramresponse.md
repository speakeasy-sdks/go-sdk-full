# FetchCryptogramResponse


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ContentType`                                                  | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `Cryptogram`                                                   | [*shared.Cryptogram](../../models/shared/cryptogram.md)        | :heavy_minus_sign:                                             | OK                                                             |
| `ErrorResponse`                                                | [*shared.ErrorResponse](../../models/shared/errorresponse.md)  | :heavy_minus_sign:                                             | Any bad or invalid request will lead to following error object |
| `Headers`                                                      | map[string][]*string*                                          | :heavy_minus_sign:                                             | N/A                                                            |
| `StatusCode`                                                   | *int*                                                          | :heavy_check_mark:                                             | N/A                                                            |
| `RawResponse`                                                  | [*http.Response](https://pkg.go.dev/net/http#Response)         | :heavy_minus_sign:                                             | N/A                                                            |