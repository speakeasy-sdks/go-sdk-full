// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CardBankName - One of ["Kotak", "ICICI", "RBL", "BOB", "Standard Chartered"]. Card bank name, required for EMI payments. This is the bank user has selected for EMI
type CardBankName string

const (
	CardBankNameKotak             CardBankName = "Kotak"
	CardBankNameIcici             CardBankName = "ICICI"
	CardBankNameRbl               CardBankName = "RBL"
	CardBankNameBob               CardBankName = "BOB"
	CardBankNameStandardChartered CardBankName = "Standard Chartered"
)

func (e CardBankName) ToPointer() *CardBankName {
	return &e
}

func (e *CardBankName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Kotak":
		fallthrough
	case "ICICI":
		fallthrough
	case "RBL":
		fallthrough
	case "BOB":
		fallthrough
	case "Standard Chartered":
		*e = CardBankName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardBankName: %v", v)
	}
}

// Channel - The channel for card payments can be "link" or "post". Post is used for seamless OTP payments where merchant captures OTP on their own page.
type Channel string

const (
	ChannelLink Channel = "link"
	ChannelPost Channel = "post"
)

func (e Channel) ToPointer() *Channel {
	return &e
}

func (e *Channel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		fallthrough
	case "post":
		*e = Channel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Channel: %v", v)
	}
}

type Card struct {
	// Card alias as returned by Cashfree Vault API.
	CardAlias *string `json:"card_alias,omitempty"`
	// One of ["Kotak", "ICICI", "RBL", "BOB", "Standard Chartered"]. Card bank name, required for EMI payments. This is the bank user has selected for EMI
	CardBankName *CardBankName `json:"card_bank_name,omitempty"`
	// CVV mentioned on the card.
	CardCvv *string `json:"card_cvv,omitempty"`
	// last 4 digits of original card number. Required only for tokenized card transactions.
	CardDisplay *string `json:"card_display,omitempty"`
	// Card expiry month for plain card transactions. Token expiry month for tokenized card transactions.
	CardExpiryMm *string `json:"card_expiry_mm,omitempty"`
	// Card expiry year for plain card transactions. Token expiry year for tokenized card transactions.
	CardExpiryYy *string `json:"card_expiry_yy,omitempty"`
	// Customer name mentioned on the card.
	CardHolderName *string `json:"card_holder_name,omitempty"`
	// Customer card number for plain card transactions. Token pan number for tokenized card transactions.
	CardNumber *string `json:"card_number,omitempty"`
	// The channel for card payments can be "link" or "post". Post is used for seamless OTP payments where merchant captures OTP on their own page.
	Channel *Channel `json:"channel,omitempty"`
	// cryptogram received from card network. Required only for tokenized card transactions.
	Cryptogram *string `json:"cryptogram,omitempty"`
	// EMI tenure selected by the user
	EmiTenure *int64 `json:"emi_tenure,omitempty"`
	// instrument id of saved card. Required only to make payment using saved instrument.
	InstrumentID *string `json:"instrument_id,omitempty"`
	// TRID issued by card networks. Required only for tokenized card transactions.
	TokenRequestorID *string `json:"token_requestor_id,omitempty"`
}

func (o *Card) GetCardAlias() *string {
	if o == nil {
		return nil
	}
	return o.CardAlias
}

func (o *Card) GetCardBankName() *CardBankName {
	if o == nil {
		return nil
	}
	return o.CardBankName
}

func (o *Card) GetCardCvv() *string {
	if o == nil {
		return nil
	}
	return o.CardCvv
}

func (o *Card) GetCardDisplay() *string {
	if o == nil {
		return nil
	}
	return o.CardDisplay
}

func (o *Card) GetCardExpiryMm() *string {
	if o == nil {
		return nil
	}
	return o.CardExpiryMm
}

func (o *Card) GetCardExpiryYy() *string {
	if o == nil {
		return nil
	}
	return o.CardExpiryYy
}

func (o *Card) GetCardHolderName() *string {
	if o == nil {
		return nil
	}
	return o.CardHolderName
}

func (o *Card) GetCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.CardNumber
}

func (o *Card) GetChannel() *Channel {
	if o == nil {
		return nil
	}
	return o.Channel
}

func (o *Card) GetCryptogram() *string {
	if o == nil {
		return nil
	}
	return o.Cryptogram
}

func (o *Card) GetEmiTenure() *int64 {
	if o == nil {
		return nil
	}
	return o.EmiTenure
}

func (o *Card) GetInstrumentID() *string {
	if o == nil {
		return nil
	}
	return o.InstrumentID
}

func (o *Card) GetTokenRequestorID() *string {
	if o == nil {
		return nil
	}
	return o.TokenRequestorID
}
