// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/utils"
	"net/http"
)

type EligibilityOfferRequest struct {
	EligibilityOffersRequest *shared.EligibilityOffersRequest `request:"mediaType=application/json"`
	XAPIVersion              *string                          `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID                string                           `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret            string                           `header:"style=simple,explode=false,name=x-client-secret"`
}

func (e EligibilityOfferRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EligibilityOfferRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EligibilityOfferRequest) GetEligibilityOffersRequest() *shared.EligibilityOffersRequest {
	if o == nil {
		return nil
	}
	return o.EligibilityOffersRequest
}

func (o *EligibilityOfferRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *EligibilityOfferRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *EligibilityOfferRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type EligibilityOfferResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.EligibleOffersEntity
}

func (o *EligibilityOfferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EligibilityOfferResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *EligibilityOfferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EligibilityOfferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EligibilityOfferResponse) GetClasses() []shared.EligibleOffersEntity {
	if o == nil {
		return nil
	}
	return o.Classes
}
