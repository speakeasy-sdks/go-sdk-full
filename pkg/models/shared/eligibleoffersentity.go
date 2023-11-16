// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EligibleOffersEntity struct {
	Eligibility   *string      `json:"eligibility,omitempty"`
	EntityDetails *OfferEntity `json:"entity_details,omitempty"`
	EntityType    *string      `json:"entity_type,omitempty"`
	EntityValue   *string      `json:"entity_value,omitempty"`
}

func (o *EligibleOffersEntity) GetEligibility() *string {
	if o == nil {
		return nil
	}
	return o.Eligibility
}

func (o *EligibleOffersEntity) GetEntityDetails() *OfferEntity {
	if o == nil {
		return nil
	}
	return o.EntityDetails
}

func (o *EligibleOffersEntity) GetEntityType() *string {
	if o == nil {
		return nil
	}
	return o.EntityType
}

func (o *EligibleOffersEntity) GetEntityValue() *string {
	if o == nil {
		return nil
	}
	return o.EntityValue
}
