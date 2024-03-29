// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/utils"
	"net/http"
)

type EligibilityCardlessEMIRequest struct {
	EligibilityCardlessEMIRequest *shared.EligibilityCardlessEMIRequest `request:"mediaType=application/json"`
	XAPIVersion                   *string                               `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID                     string                                `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret                 string                                `header:"style=simple,explode=false,name=x-client-secret"`
}

func (e EligibilityCardlessEMIRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EligibilityCardlessEMIRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EligibilityCardlessEMIRequest) GetEligibilityCardlessEMIRequest() *shared.EligibilityCardlessEMIRequest {
	if o == nil {
		return nil
	}
	return o.EligibilityCardlessEMIRequest
}

func (o *EligibilityCardlessEMIRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *EligibilityCardlessEMIRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *EligibilityCardlessEMIRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type EligibilityCardlessEMIResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.EligibleCardlessEMIEntity
}

func (o *EligibilityCardlessEMIResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EligibilityCardlessEMIResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *EligibilityCardlessEMIResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EligibilityCardlessEMIResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EligibilityCardlessEMIResponse) GetClasses() []shared.EligibleCardlessEMIEntity {
	if o == nil {
		return nil
	}
	return o.Classes
}
