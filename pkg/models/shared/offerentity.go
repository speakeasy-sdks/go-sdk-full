// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// OfferEntity - OK
type OfferEntity struct {
	OfferDetails     *OfferDetails     `json:"offer_details,omitempty"`
	OfferID          *int64            `json:"offer_id,omitempty"`
	OfferMeta        *OfferMeta        `json:"offer_meta,omitempty"`
	OfferStatus      *string           `json:"offer_status,omitempty"`
	OfferTnc         *OfferTnc         `json:"offer_tnc,omitempty"`
	OfferValidations *OfferValidations `json:"offer_validations,omitempty"`
}

func (o *OfferEntity) GetOfferDetails() *OfferDetails {
	if o == nil {
		return nil
	}
	return o.OfferDetails
}

func (o *OfferEntity) GetOfferID() *int64 {
	if o == nil {
		return nil
	}
	return o.OfferID
}

func (o *OfferEntity) GetOfferMeta() *OfferMeta {
	if o == nil {
		return nil
	}
	return o.OfferMeta
}

func (o *OfferEntity) GetOfferStatus() *string {
	if o == nil {
		return nil
	}
	return o.OfferStatus
}

func (o *OfferEntity) GetOfferTnc() *OfferTnc {
	if o == nil {
		return nil
	}
	return o.OfferTnc
}

func (o *OfferEntity) GetOfferValidations() *OfferValidations {
	if o == nil {
		return nil
	}
	return o.OfferValidations
}
