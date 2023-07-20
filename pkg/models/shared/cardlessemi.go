// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CardlessEMIProvider - One of [`flexmoney`, `zestmoney`, `hdfc`, `icici`, `cashe`, `idfc`, `kotak`]
type CardlessEMIProvider string

const (
	CardlessEMIProviderFlexmoney CardlessEMIProvider = "flexmoney"
	CardlessEMIProviderZestmoney CardlessEMIProvider = "zestmoney"
	CardlessEMIProviderHdfc      CardlessEMIProvider = "hdfc"
	CardlessEMIProviderIcici     CardlessEMIProvider = "icici"
	CardlessEMIProviderCashe     CardlessEMIProvider = "cashe"
	CardlessEMIProviderIdfc      CardlessEMIProvider = "idfc"
	CardlessEMIProviderKotak     CardlessEMIProvider = "kotak"
)

func (e CardlessEMIProvider) ToPointer() *CardlessEMIProvider {
	return &e
}

func (e *CardlessEMIProvider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "flexmoney":
		fallthrough
	case "zestmoney":
		fallthrough
	case "hdfc":
		fallthrough
	case "icici":
		fallthrough
	case "cashe":
		fallthrough
	case "idfc":
		fallthrough
	case "kotak":
		*e = CardlessEMIProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardlessEMIProvider: %v", v)
	}
}

type CardlessEMI struct {
	// The channel for cardless EMI is always `link`
	Channel *string `json:"channel,omitempty"`
	// EMI tenure for the selected provider. This is mandatory when provider is one of [`hdfc`, `icici`, `cashe`, `idfc`, `kotak`]
	EmiTenure *int64 `json:"emi_tenure,omitempty"`
	// Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as 'invalid_request_error' and code as 'invalid_request_error'
	Phone *string `json:"phone,omitempty"`
	// One of [`flexmoney`, `zestmoney`, `hdfc`, `icici`, `cashe`, `idfc`, `kotak`]
	Provider *CardlessEMIProvider `json:"provider,omitempty"`
}

func (o *CardlessEMI) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}

func (o *CardlessEMI) GetEmiTenure() *int64 {
	if o == nil {
		return nil
	}
	return o.EmiTenure
}

func (o *CardlessEMI) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *CardlessEMI) GetProvider() *CardlessEMIProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}
