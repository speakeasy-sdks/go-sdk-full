// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ErrorDetailsInPaymentsEntity - The error details are present only for failed payments
type ErrorDetailsInPaymentsEntity struct {
	ErrorCode        *string `json:"error_code,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
	ErrorReason      *string `json:"error_reason,omitempty"`
	ErrorSource      *string `json:"error_source,omitempty"`
}

func (o *ErrorDetailsInPaymentsEntity) GetErrorCode() *string {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *ErrorDetailsInPaymentsEntity) GetErrorDescription() *string {
	if o == nil {
		return nil
	}
	return o.ErrorDescription
}

func (o *ErrorDetailsInPaymentsEntity) GetErrorReason() *string {
	if o == nil {
		return nil
	}
	return o.ErrorReason
}

func (o *ErrorDetailsInPaymentsEntity) GetErrorSource() *string {
	if o == nil {
		return nil
	}
	return o.ErrorSource
}
