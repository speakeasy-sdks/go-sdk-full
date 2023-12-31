// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CashbackType - Type of discount
type CashbackType string

const (
	CashbackTypeFlat       CashbackType = "flat"
	CashbackTypePercentage CashbackType = "percentage"
)

func (e CashbackType) ToPointer() *CashbackType {
	return &e
}

func (e *CashbackType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "flat":
		fallthrough
	case "percentage":
		*e = CashbackType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CashbackType: %v", v)
	}
}

type CashbackDetails struct {
	// Type of discount
	CashbackType *CashbackType `json:"cashback_type,omitempty"`
	// Value of Discount.
	CashbackValue *string `json:"cashback_value,omitempty"`
	// Maximum Value of Cashback allowed.
	MaxCashbackAmount string `json:"max_cashback_amount"`
}

func (o *CashbackDetails) GetCashbackType() *CashbackType {
	if o == nil {
		return nil
	}
	return o.CashbackType
}

func (o *CashbackDetails) GetCashbackValue() *string {
	if o == nil {
		return nil
	}
	return o.CashbackValue
}

func (o *CashbackDetails) GetMaxCashbackAmount() string {
	if o == nil {
		return ""
	}
	return o.MaxCashbackAmount
}
