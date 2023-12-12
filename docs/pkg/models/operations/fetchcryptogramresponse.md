# FetchCryptogramResponse


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ContentType`                                                        | *string*                                                             | :heavy_check_mark:                                                   | HTTP response content type for this operation                        |
| `Cryptogram`                                                         | [*shared.Cryptogram](../../../pkg/models/shared/cryptogram.md)       | :heavy_minus_sign:                                                   | OK                                                                   |
| `ErrorResponse`                                                      | [*shared.ErrorResponse](../../../pkg/models/shared/errorresponse.md) | :heavy_minus_sign:                                                   | Any bad or invalid request will lead to following error object       |
| `Headers`                                                            | map[string][]*string*                                                | :heavy_check_mark:                                                   | N/A                                                                  |
| `StatusCode`                                                         | *int*                                                                | :heavy_check_mark:                                                   | HTTP response status code for this operation                         |
| `RawResponse`                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)               | :heavy_check_mark:                                                   | Raw HTTP response; suitable for custom response parsing              |