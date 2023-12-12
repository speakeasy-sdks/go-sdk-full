# CreateOrderBackendRequest


## Fields

| Field                                                                                                                                                                                                    | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              | Example                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CustomerDetails`                                                                                                                                                                                        | [shared.CustomerDetails](../../../pkg/models/shared/customerdetails.md)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                       | The customer details that are necessary. Note that you can pass dummy details if your use case does not require the customer details.                                                                    |                                                                                                                                                                                                          |
| `OrderAmount`                                                                                                                                                                                            | *float64*                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                       | Bill amount for the order. Provide upto two decimals. 10.15 means Rs 10 and 15 paisa                                                                                                                     | 10.15                                                                                                                                                                                                    |
| `OrderCurrency`                                                                                                                                                                                          | *string*                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                       | Currency for the order. INR if left empty. Contact care@cashfree.com to enable new currencies.                                                                                                           | INR                                                                                                                                                                                                      |
| `OrderExpiryTime`                                                                                                                                                                                        | **string*                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                       | Time after which the order expires. Customers will not be able to make the payment beyond the time specified here. We store timestamps in IST, but you can provide them in a valid ISO 8601 time format. | 2021-07-29T00:00:00Z                                                                                                                                                                                     |
| `OrderID`                                                                                                                                                                                                | **string*                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                       | Order identifier present in your system. Alphanumeric and only - and _ allowed.                                                                                                                          |                                                                                                                                                                                                          |
| `OrderMeta`                                                                                                                                                                                              | [*shared.OrderMeta](../../../pkg/models/shared/ordermeta.md)                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                       | N/A                                                                                                                                                                                                      |                                                                                                                                                                                                          |
| `OrderNote`                                                                                                                                                                                              | **string*                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                       | Order note for reference.                                                                                                                                                                                | Test order                                                                                                                                                                                               |
| `OrderSplits`                                                                                                                                                                                            | [][shared.VendorSplit](../../../pkg/models/shared/vendorsplit.md)                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                       | N/A                                                                                                                                                                                                      |                                                                                                                                                                                                          |
| `OrderTags`                                                                                                                                                                                              | map[string]*string*                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                       | Custom Tags which can be passed for an order. A maximum of 6 tags can be added                                                                                                                           |                                                                                                                                                                                                          |
| `Terminal`                                                                                                                                                                                               | [*shared.TerminalDetails](../../../pkg/models/shared/terminaldetails.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                       | N/A                                                                                                                                                                                                      |                                                                                                                                                                                                          |