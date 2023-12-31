// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Entity - Type of object
type Entity string

const (
	EntityRefund Entity = "refund"
)

func (e Entity) ToPointer() *Entity {
	return &e
}

func (e *Entity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "refund":
		*e = Entity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Entity: %v", v)
	}
}

// Metadata - Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
type Metadata struct {
}

// RefundMode - Method or speed of processing refund
type RefundMode string

const (
	RefundModeStandard RefundMode = "STANDARD"
	RefundModeInstant  RefundMode = "INSTANT"
)

func (e RefundMode) ToPointer() *RefundMode {
	return &e
}

func (e *RefundMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		fallthrough
	case "INSTANT":
		*e = RefundMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundMode: %v", v)
	}
}

// RefundStatus - This can be one of ["SUCCESS", "PENDING", "CANCELLED", "ONHOLD", "FAILED"]
type RefundStatus string

const (
	RefundStatusSuccess   RefundStatus = "SUCCESS"
	RefundStatusPending   RefundStatus = "PENDING"
	RefundStatusCancelled RefundStatus = "CANCELLED"
	RefundStatusOnhold    RefundStatus = "ONHOLD"
)

func (e RefundStatus) ToPointer() *RefundStatus {
	return &e
}

func (e *RefundStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		fallthrough
	case "PENDING":
		fallthrough
	case "CANCELLED":
		fallthrough
	case "ONHOLD":
		*e = RefundStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundStatus: %v", v)
	}
}

// RefundType - This can be one of ["PAYMENT_AUTO_REFUND", "MERCHANT_INITIATED", "UNRECONCILED_AUTO_REFUND"]
type RefundType string

const (
	RefundTypePaymentAutoRefund      RefundType = "PAYMENT_AUTO_REFUND"
	RefundTypeMerchantInitiated      RefundType = "MERCHANT_INITIATED"
	RefundTypeUnreconciledAutoRefund RefundType = "UNRECONCILED_AUTO_REFUND"
)

func (e RefundType) ToPointer() *RefundType {
	return &e
}

func (e *RefundType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PAYMENT_AUTO_REFUND":
		fallthrough
	case "MERCHANT_INITIATED":
		fallthrough
	case "UNRECONCILED_AUTO_REFUND":
		*e = RefundType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundType: %v", v)
	}
}

type RefundsEntity struct {
	// Cashfree Payments ID of the payment for which refund is initiated
	CfPaymentID *int64 `json:"cf_payment_id,omitempty"`
	// Cashfree Payments ID for a refund
	CfRefundID *string `json:"cf_refund_id,omitempty"`
	// Time of refund creation
	CreatedAt *string `json:"created_at,omitempty"`
	// Type of object
	Entity *Entity `json:"entity,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	Metadata *Metadata `json:"metadata,omitempty"`
	// Merchant’s order Id of the order for which refund is initiated
	OrderID *string `json:"order_id,omitempty"`
	// Time when refund was processed successfully
	ProcessedAt *string `json:"processed_at,omitempty"`
	// Amount that is refunded
	RefundAmount *float64 `json:"refund_amount,omitempty"`
	// The bank reference number for refund
	RefundArn *string `json:"refund_arn,omitempty"`
	// Charges in INR for processing refund
	RefundCharge *float64 `json:"refund_charge,omitempty"`
	// Currency of the refund amount
	RefundCurrency *string `json:"refund_currency,omitempty"`
	// Merchant’s refund ID of the refund
	RefundID *string `json:"refund_id,omitempty"`
	// Method or speed of processing refund
	RefundMode *RefundMode `json:"refund_mode,omitempty"`
	// Note added by merchant for the refund
	RefundNote   *string       `json:"refund_note,omitempty"`
	RefundSpeed  *RefundSpeed  `json:"refund_speed,omitempty"`
	RefundSplits []VendorSplit `json:"refund_splits,omitempty"`
	// This can be one of ["SUCCESS", "PENDING", "CANCELLED", "ONHOLD", "FAILED"]
	RefundStatus *RefundStatus `json:"refund_status,omitempty"`
	// This can be one of ["PAYMENT_AUTO_REFUND", "MERCHANT_INITIATED", "UNRECONCILED_AUTO_REFUND"]
	RefundType *RefundType `json:"refund_type,omitempty"`
	// Description of refund status
	StatusDescription *string `json:"status_description,omitempty"`
}

func (o *RefundsEntity) GetCfPaymentID() *int64 {
	if o == nil {
		return nil
	}
	return o.CfPaymentID
}

func (o *RefundsEntity) GetCfRefundID() *string {
	if o == nil {
		return nil
	}
	return o.CfRefundID
}

func (o *RefundsEntity) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RefundsEntity) GetEntity() *Entity {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *RefundsEntity) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *RefundsEntity) GetOrderID() *string {
	if o == nil {
		return nil
	}
	return o.OrderID
}

func (o *RefundsEntity) GetProcessedAt() *string {
	if o == nil {
		return nil
	}
	return o.ProcessedAt
}

func (o *RefundsEntity) GetRefundAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.RefundAmount
}

func (o *RefundsEntity) GetRefundArn() *string {
	if o == nil {
		return nil
	}
	return o.RefundArn
}

func (o *RefundsEntity) GetRefundCharge() *float64 {
	if o == nil {
		return nil
	}
	return o.RefundCharge
}

func (o *RefundsEntity) GetRefundCurrency() *string {
	if o == nil {
		return nil
	}
	return o.RefundCurrency
}

func (o *RefundsEntity) GetRefundID() *string {
	if o == nil {
		return nil
	}
	return o.RefundID
}

func (o *RefundsEntity) GetRefundMode() *RefundMode {
	if o == nil {
		return nil
	}
	return o.RefundMode
}

func (o *RefundsEntity) GetRefundNote() *string {
	if o == nil {
		return nil
	}
	return o.RefundNote
}

func (o *RefundsEntity) GetRefundSpeed() *RefundSpeed {
	if o == nil {
		return nil
	}
	return o.RefundSpeed
}

func (o *RefundsEntity) GetRefundSplits() []VendorSplit {
	if o == nil {
		return nil
	}
	return o.RefundSplits
}

func (o *RefundsEntity) GetRefundStatus() *RefundStatus {
	if o == nil {
		return nil
	}
	return o.RefundStatus
}

func (o *RefundsEntity) GetRefundType() *RefundType {
	if o == nil {
		return nil
	}
	return o.RefundType
}

func (o *RefundsEntity) GetStatusDescription() *string {
	if o == nil {
		return nil
	}
	return o.StatusDescription
}
