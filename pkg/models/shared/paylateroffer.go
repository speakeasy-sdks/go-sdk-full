// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaylaterOffer struct {
	Provider *string `json:"provider,omitempty"`
}

func (o *PaylaterOffer) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}
