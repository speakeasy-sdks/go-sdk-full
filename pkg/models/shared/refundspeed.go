// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RefundSpeed struct {
	// Accepted speed of refund.
	Accepted *string `json:"accepted,omitempty"`
	// Error message, if any for refund_speed request
	Message *string `json:"message,omitempty"`
	// Processed speed of refund.
	Processed *string `json:"processed,omitempty"`
	// Requested speed of refund.
	Requested *string `json:"requested,omitempty"`
}

func (o *RefundSpeed) GetAccepted() *string {
	if o == nil {
		return nil
	}
	return o.Accepted
}

func (o *RefundSpeed) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RefundSpeed) GetProcessed() *string {
	if o == nil {
		return nil
	}
	return o.Processed
}

func (o *RefundSpeed) GetRequested() *string {
	if o == nil {
		return nil
	}
	return o.Requested
}
