// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaylaterEntity struct {
	PaymentMethod *string `json:"payment_method,omitempty"`
}

func (o *PaylaterEntity) GetPaymentMethod() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}
