// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateOfferBackendRequest struct {
	OfferDetails     OfferDetails     `json:"offer_details"`
	OfferMeta        OfferMeta        `json:"offer_meta"`
	OfferTnc         OfferTnc         `json:"offer_tnc"`
	OfferValidations OfferValidations `json:"offer_validations"`
}

func (o *CreateOfferBackendRequest) GetOfferDetails() OfferDetails {
	if o == nil {
		return OfferDetails{}
	}
	return o.OfferDetails
}

func (o *CreateOfferBackendRequest) GetOfferMeta() OfferMeta {
	if o == nil {
		return OfferMeta{}
	}
	return o.OfferMeta
}

func (o *CreateOfferBackendRequest) GetOfferTnc() OfferTnc {
	if o == nil {
		return OfferTnc{}
	}
	return o.OfferTnc
}

func (o *CreateOfferBackendRequest) GetOfferValidations() OfferValidations {
	if o == nil {
		return OfferValidations{}
	}
	return o.OfferValidations
}