// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WHpayment struct {
	AuthID          *string          `json:"auth_id,omitempty"`
	BankReference   *string          `json:"bank_reference,omitempty"`
	CfPaymentID     *string          `json:"cf_payment_id,omitempty"`
	PaymentAmount   *float64         `json:"payment_amount,omitempty"`
	PaymentCurrency *string          `json:"payment_currency,omitempty"`
	PaymentGroup    *string          `json:"payment_group,omitempty"`
	PaymentMessage  *string          `json:"payment_message,omitempty"`
	PaymentMethod   *WHpaymentMethod `json:"payment_method,omitempty"`
	PaymentStatus   *string          `json:"payment_status,omitempty"`
	PaymentTime     *string          `json:"payment_time,omitempty"`
}

func (o *WHpayment) GetAuthID() *string {
	if o == nil {
		return nil
	}
	return o.AuthID
}

func (o *WHpayment) GetBankReference() *string {
	if o == nil {
		return nil
	}
	return o.BankReference
}

func (o *WHpayment) GetCfPaymentID() *string {
	if o == nil {
		return nil
	}
	return o.CfPaymentID
}

func (o *WHpayment) GetPaymentAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.PaymentAmount
}

func (o *WHpayment) GetPaymentCurrency() *string {
	if o == nil {
		return nil
	}
	return o.PaymentCurrency
}

func (o *WHpayment) GetPaymentGroup() *string {
	if o == nil {
		return nil
	}
	return o.PaymentGroup
}

func (o *WHpayment) GetPaymentMessage() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMessage
}

func (o *WHpayment) GetPaymentMethod() *WHpaymentMethod {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *WHpayment) GetPaymentStatus() *string {
	if o == nil {
		return nil
	}
	return o.PaymentStatus
}

func (o *WHpayment) GetPaymentTime() *string {
	if o == nil {
		return nil
	}
	return o.PaymentTime
}
