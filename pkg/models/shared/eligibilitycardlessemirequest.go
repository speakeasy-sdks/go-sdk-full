// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EligibilityCardlessEMIRequest struct {
	Queries CardlessEMIQueries `json:"queries"`
}

func (o *EligibilityCardlessEMIRequest) GetQueries() CardlessEMIQueries {
	if o == nil {
		return CardlessEMIQueries{}
	}
	return o.Queries
}