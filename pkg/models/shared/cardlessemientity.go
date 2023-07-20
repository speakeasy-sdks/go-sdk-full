// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CardlessEMIEntity struct {
	EmiPlans      []EMIPlansArray `json:"emi_plans,omitempty"`
	PaymentMethod *string         `json:"payment_method,omitempty"`
}

func (o *CardlessEMIEntity) GetEmiPlans() []EMIPlansArray {
	if o == nil {
		return nil
	}
	return o.EmiPlans
}

func (o *CardlessEMIEntity) GetPaymentMethod() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}
