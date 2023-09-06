// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaymentMethodCardlessEMIInPaymentsEntity struct {
	Channel  *string `json:"channel,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Provider *string `json:"provider,omitempty"`
}

func (o *PaymentMethodCardlessEMIInPaymentsEntity) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}

func (o *PaymentMethodCardlessEMIInPaymentsEntity) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *PaymentMethodCardlessEMIInPaymentsEntity) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}