# EMIOffer


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `BankName`                                                                                 | *string*                                                                                   | :heavy_check_mark:                                                                         | Bank Name                                                                                  | hdfc bank                                                                                  |
| `Tenures`                                                                                  | []*int64*                                                                                  | :heavy_minus_sign:                                                                         | N/A                                                                                        |                                                                                            |
| `Type`                                                                                     | *string*                                                                                   | :heavy_check_mark:                                                                         | Type of emi offer. Possible values are `credit_card_emi`, `debit_card_emi`, `cardless_emi` | cardless_emi                                                                               |