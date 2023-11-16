// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CardlessEMIQueries struct {
	// Amount of the order. OrderId of the order. Either of `order_id` or `amount` is mandatory.
	Amount *float64 `json:"amount,omitempty"`
	// Details of the customer for whom eligibility is being checked.
	CustomerDetails *CustomerDetailsCardlessEMI `json:"customer_details,omitempty"`
	// OrderId of the order. Either of `order_id` or `amount` is mandatory.
	OrderID *string `json:"order_id,omitempty"`
}

func (o *CardlessEMIQueries) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CardlessEMIQueries) GetCustomerDetails() *CustomerDetailsCardlessEMI {
	if o == nil {
		return nil
	}
	return o.CustomerDetails
}

func (o *CardlessEMIQueries) GetOrderID() *string {
	if o == nil {
		return nil
	}
	return o.OrderID
}
