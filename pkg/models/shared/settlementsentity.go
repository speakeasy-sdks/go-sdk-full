// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SettlementsEntity struct {
	CfPaymentID        *string  `json:"cf_payment_id,omitempty"`
	CfSettlementID     *string  `json:"cf_settlement_id,omitempty"`
	Entity             *string  `json:"entity,omitempty"`
	OrderAmount        *float64 `json:"order_amount,omitempty"`
	OrderID            *string  `json:"order_id,omitempty"`
	PaymentTime        *string  `json:"payment_time,omitempty"`
	ServiceCharge      *float64 `json:"service_charge,omitempty"`
	ServiceTax         *float64 `json:"service_tax,omitempty"`
	SettlementAmount   *float64 `json:"settlement_amount,omitempty"`
	SettlementCurrency *string  `json:"settlement_currency,omitempty"`
	SettlementID       *int64   `json:"settlement_id,omitempty"`
	TransferID         *int64   `json:"transfer_id,omitempty"`
	TransferTime       *string  `json:"transfer_time,omitempty"`
	TransferUtr        *string  `json:"transfer_utr,omitempty"`
}

func (o *SettlementsEntity) GetCfPaymentID() *string {
	if o == nil {
		return nil
	}
	return o.CfPaymentID
}

func (o *SettlementsEntity) GetCfSettlementID() *string {
	if o == nil {
		return nil
	}
	return o.CfSettlementID
}

func (o *SettlementsEntity) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *SettlementsEntity) GetOrderAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.OrderAmount
}

func (o *SettlementsEntity) GetOrderID() *string {
	if o == nil {
		return nil
	}
	return o.OrderID
}

func (o *SettlementsEntity) GetPaymentTime() *string {
	if o == nil {
		return nil
	}
	return o.PaymentTime
}

func (o *SettlementsEntity) GetServiceCharge() *float64 {
	if o == nil {
		return nil
	}
	return o.ServiceCharge
}

func (o *SettlementsEntity) GetServiceTax() *float64 {
	if o == nil {
		return nil
	}
	return o.ServiceTax
}

func (o *SettlementsEntity) GetSettlementAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.SettlementAmount
}

func (o *SettlementsEntity) GetSettlementCurrency() *string {
	if o == nil {
		return nil
	}
	return o.SettlementCurrency
}

func (o *SettlementsEntity) GetSettlementID() *int64 {
	if o == nil {
		return nil
	}
	return o.SettlementID
}

func (o *SettlementsEntity) GetTransferID() *int64 {
	if o == nil {
		return nil
	}
	return o.TransferID
}

func (o *SettlementsEntity) GetTransferTime() *string {
	if o == nil {
		return nil
	}
	return o.TransferTime
}

func (o *SettlementsEntity) GetTransferUtr() *string {
	if o == nil {
		return nil
	}
	return o.TransferUtr
}
