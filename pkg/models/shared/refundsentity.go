// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RefundsEntityEntity - Type of object
type RefundsEntityEntity string

const (
	RefundsEntityEntityRefund RefundsEntityEntity = "refund"
)

func (e RefundsEntityEntity) ToPointer() *RefundsEntityEntity {
	return &e
}

func (e *RefundsEntityEntity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "refund":
		*e = RefundsEntityEntity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundsEntityEntity: %v", v)
	}
}

// RefundsEntityMetadata - Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
type RefundsEntityMetadata struct {
}

// RefundsEntityRefundMode - Method or speed of processing refund
type RefundsEntityRefundMode string

const (
	RefundsEntityRefundModeStandard RefundsEntityRefundMode = "STANDARD"
	RefundsEntityRefundModeInstant  RefundsEntityRefundMode = "INSTANT"
)

func (e RefundsEntityRefundMode) ToPointer() *RefundsEntityRefundMode {
	return &e
}

func (e *RefundsEntityRefundMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		fallthrough
	case "INSTANT":
		*e = RefundsEntityRefundMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundsEntityRefundMode: %v", v)
	}
}

// RefundsEntityRefundStatus - This can be one of ["SUCCESS", "PENDING", "CANCELLED", "ONHOLD", "FAILED"]
type RefundsEntityRefundStatus string

const (
	RefundsEntityRefundStatusSuccess   RefundsEntityRefundStatus = "SUCCESS"
	RefundsEntityRefundStatusPending   RefundsEntityRefundStatus = "PENDING"
	RefundsEntityRefundStatusCancelled RefundsEntityRefundStatus = "CANCELLED"
	RefundsEntityRefundStatusOnhold    RefundsEntityRefundStatus = "ONHOLD"
)

func (e RefundsEntityRefundStatus) ToPointer() *RefundsEntityRefundStatus {
	return &e
}

func (e *RefundsEntityRefundStatus) UnmarshalJSON(data []byte) error {
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
		*e = RefundsEntityRefundStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundsEntityRefundStatus: %v", v)
	}
}

// RefundsEntityRefundType - This can be one of ["PAYMENT_AUTO_REFUND", "MERCHANT_INITIATED", "UNRECONCILED_AUTO_REFUND"]
type RefundsEntityRefundType string

const (
	RefundsEntityRefundTypePaymentAutoRefund      RefundsEntityRefundType = "PAYMENT_AUTO_REFUND"
	RefundsEntityRefundTypeMerchantInitiated      RefundsEntityRefundType = "MERCHANT_INITIATED"
	RefundsEntityRefundTypeUnreconciledAutoRefund RefundsEntityRefundType = "UNRECONCILED_AUTO_REFUND"
)

func (e RefundsEntityRefundType) ToPointer() *RefundsEntityRefundType {
	return &e
}

func (e *RefundsEntityRefundType) UnmarshalJSON(data []byte) error {
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
		*e = RefundsEntityRefundType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundsEntityRefundType: %v", v)
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
	Entity *RefundsEntityEntity `json:"entity,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	Metadata *RefundsEntityMetadata `json:"metadata,omitempty"`
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
	RefundMode *RefundsEntityRefundMode `json:"refund_mode,omitempty"`
	// Note added by merchant for the refund
	RefundNote   *string       `json:"refund_note,omitempty"`
	RefundSpeed  *RefundSpeed  `json:"refund_speed,omitempty"`
	RefundSplits []VendorSplit `json:"refund_splits,omitempty"`
	// This can be one of ["SUCCESS", "PENDING", "CANCELLED", "ONHOLD", "FAILED"]
	RefundStatus *RefundsEntityRefundStatus `json:"refund_status,omitempty"`
	// This can be one of ["PAYMENT_AUTO_REFUND", "MERCHANT_INITIATED", "UNRECONCILED_AUTO_REFUND"]
	RefundType *RefundsEntityRefundType `json:"refund_type,omitempty"`
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

func (o *RefundsEntity) GetEntity() *RefundsEntityEntity {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *RefundsEntity) GetMetadata() *RefundsEntityMetadata {
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

func (o *RefundsEntity) GetRefundMode() *RefundsEntityRefundMode {
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

func (o *RefundsEntity) GetRefundStatus() *RefundsEntityRefundStatus {
	if o == nil {
		return nil
	}
	return o.RefundStatus
}

func (o *RefundsEntity) GetRefundType() *RefundsEntityRefundType {
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
